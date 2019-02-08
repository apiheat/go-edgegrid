package edgegrid

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/go-ini/ini"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
)

// Akamai {OPEN} EdgeGrid Authentication Service
type AuthService struct {
	client *Client
}

//EdgercCredentials are items from config file
type EdgercCredentials struct {
	host         string `ini:"host"`
	clientToken  string `ini:"client_token"`
	clientSecret string `ini:"client_secret"`
	accessToken  string `ini:"access_token"`
}

type reader struct {
	*bytes.Buffer
}

func (m reader) Close() error { return nil }

// Init initializes using a configuration file in standard INI format
func InitEdgerc(edgercConfig, edgercSection string) (*EdgercCredentials, error) {

	log.WithFields(log.Fields{
		"edgercConfig":  edgercConfig,
		"edgercSection": edgercSection,
	}).Info("[InitEdgerc]::Initialize credentials")

	// We first want to check Env Variables
	log.Debug("[InitEdgerc]::Loading credentials from environment variables")
	var (
		requiredOptions = []string{"HOST", "CLIENT_TOKEN", "CLIENT_SECRET", "ACCESS_TOKEN"}
		missing         []string
	)

	prefix := "AKAMAI_"
	envCredentials := &EdgercCredentials{}

	for _, opt := range requiredOptions {
		val, ok := os.LookupEnv(prefix + opt)
		if val == "" {
			missing = append(missing, prefix+opt)
		}
		if !ok {
			missing = append(missing, prefix+opt)
		} else {
			switch {
			case opt == "HOST":
				envCredentials.host = val
			case opt == "CLIENT_TOKEN":
				envCredentials.clientToken = val
			case opt == "CLIENT_SECRET":
				envCredentials.clientSecret = val
			case opt == "ACCESS_TOKEN":
				envCredentials.accessToken = val
			}
		}
	}

	if len(missing) > 0 {
		log.Debug(fmt.Sprintf("[InitEdgerc]::Missing required environment variables: %s", missing))
	}

	if len(missing) == 0 {
		err := validateCreds(envCredentials, "environment variable")
		if err == nil {
			log.Debug("[InitEdgerc]::Return ENV credentials object")
			return envCredentials, nil
		}
		log.Debug(fmt.Sprintf("[InitEdgerc]::Environment variables are not correct: %s", err))
	}

	// Load the file based on our provided config
	log.Debug("[InitEdgerc]::Loading credentials file")
	edgerc, err := ini.Load(edgercConfig)
	if err != nil {
		return nil, fmt.Errorf("Error loading file? '%s'", err)
	}

	log.Debug("[InitEdgerc]::Loading section from credentials file")
	sectionNames := edgerc.SectionStrings()
	if !(stringInSlice(edgercSection, sectionNames)) {
		return nil, fmt.Errorf("Could not load section '%s'", edgercSection)
	}

	log.Debug("[InitEdgerc]::Lookup for credentials ( host/secrets etc)")
	edgercHost := edgerc.Section(edgercSection).Key("host").String()
	edgercclientToken := edgerc.Section(edgercSection).Key("client_token").String()
	edgercclientSecret := edgerc.Section(edgercSection).Key("client_secret").String()
	edgercaccessToken := edgerc.Section(edgercSection).Key("access_token").String()

	log.Debug("[InitEdgerc]::Create credentials object")
	loadedCredentials := &EdgercCredentials{
		host:         edgercHost,
		clientToken:  edgercclientToken,
		clientSecret: edgercclientSecret,
		accessToken:  edgercaccessToken,
	}

	err = validateCreds(loadedCredentials, fmt.Sprintf("section '%s'", edgercSection))
	if err != nil {
		return nil, err
	}

	log.Debug("[InitEdgerc]::Map credentials to appropiate object")
	err = edgerc.Section(edgercSection).MapTo(loadedCredentials)
	if err != nil {
		return nil, fmt.Errorf("Error loading file? %s", err)
	}

	log.Debug("[InitEdgerc]::Return credentials object")
	return loadedCredentials, nil

}

func validateCreds(creds *EdgercCredentials, location string) error {
	log.Debug("[InitEdgerc]::Validating credentials - 'host'")
	if creds.host == "" {
		return fmt.Errorf("'host' is empty in '%s'", location)
	}

	valid := govalidator.IsURL(creds.host)
	if !valid {
		return fmt.Errorf("'host' is not valid URL in '%s'", location)
	}

	u, err := url.Parse(creds.host)
	if err != nil {
		return fmt.Errorf("'host' cannot be parsed correctly in '%s'", location)
	}

	if u.Scheme != "" {
		return fmt.Errorf("'host' in '%s' contains URL scheme: '%s', please remove '%s//'", location, u.Scheme, u.Scheme)
	}

	log.Debug("[InitEdgerc]::Validating credentials - 'client_token'")
	if creds.clientToken == "" {
		return fmt.Errorf("'client_token' is empty in '%s'", location)
	}

	log.Debug("[InitEdgerc]::Validating credentials - 'client_secret'")
	if creds.clientSecret == "" {
		return fmt.Errorf("'client_secret' is empty in '%s'", location)
	}

	log.Debug("[InitEdgerc]::Validating credentials - 'access_token'")
	if creds.accessToken == "" {
		return fmt.Errorf("'access_token' is empty in '%s'", location)
	}

	return nil
}

// AuthString takes prm and returns a string that can be
// used as the `Authorization` header in making Akamai API requests.
//
// The string returned by Auth conforms to the
// Akamai {OPEN} EdgeGrid Authentication scheme.
// https://developer.akamai.com/introduction/Client_Auth.html
func AuthString(eprm *EdgercCredentials, request *http.Request, headersToSign []string) string {

	u := uuid.NewV4()

	nonce := u.String()

	timestamp := time.Now().UTC().Format("20060102T15:04:05+0000")

	var auth bytes.Buffer
	orderedKeys := []string{"client_token", "access_token", "timestamp", "nonce"}

	m := map[string]string{
		orderedKeys[0]: eprm.clientToken,
		orderedKeys[1]: eprm.accessToken,
		orderedKeys[2]: timestamp,
		orderedKeys[3]: nonce,
	}

	auth.WriteString("EG1-HMAC-SHA256 ")

	for _, each := range orderedKeys {
		auth.WriteString(concat([]string{
			each,
			"=",
			m[each],
			";",
		}))
	}

	auth.WriteString(signRequest(request, timestamp, eprm.clientSecret, auth.String(), headersToSign))

	return auth.String()
}

func signRequest(request *http.Request, timestamp, clientSecret, authHeader string, headersToSign []string) string {
	dataToSign := makeDataToSign(request, authHeader, headersToSign)
	signingKey := makeSigningKey(timestamp, clientSecret)

	return concat([]string{
		"signature=",
		base64HmacSha256(dataToSign, signingKey),
	})
}

func base64Sha256(str string) string {
	h := sha256.New()

	h.Write([]byte(str))

	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func base64HmacSha256(message, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))

	h.Write([]byte(message))

	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func makeDataToSign(request *http.Request, authHeader string, headersToSign []string) string {
	var data bytes.Buffer
	values := []string{
		request.Method,
		request.URL.Scheme,
		request.Host,
		urlPathWithQuery(request),
		canonicalizeHeaders(request, headersToSign),
		makeContentHash(request),
		authHeader,
	}

	data.WriteString(strings.Join(values, "\t"))

	return data.String()
}

func canonicalizeHeaders(request *http.Request, headersToSign []string) string {
	var canonicalized bytes.Buffer

	for key, values := range request.Header {
		if stringInSlice(key, headersToSign) {
			canonicalized.WriteString(concat([]string{
				strings.ToLower(key),
				":",
				strings.Join(strings.Fields(values[0]), " "),
				"\t",
			}))
		}
	}

	return canonicalized.String()
}

func makeContentHash(req *http.Request) string {
	if req.Method == "POST" {
		buf, err := ioutil.ReadAll(req.Body)
		rdr := reader{bytes.NewBuffer(buf)}

		if err != nil {
			panic(err)
		}

		req.Body = rdr

		return base64Sha256(string(buf))
	}

	return ""
}

func makeSigningKey(timestamp, clientSecret string) string {
	return base64HmacSha256(timestamp, clientSecret)
}

//#TODO: Move to common CLI
func urlPathWithQuery(req *http.Request) string {
	var query string

	if req.URL.RawQuery != "" {
		query = concat([]string{
			"?",
			req.URL.RawQuery,
		})
	} else {
		query = ""
	}

	return concat([]string{
		req.URL.Path,
		query,
	})
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func concat(arr []string) string {
	var buff bytes.Buffer

	for _, elem := range arr {
		buff.WriteString(elem)
	}

	return buff.String()
}
