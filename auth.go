package edgegrid

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/go-ini/ini"
	homedir "github.com/mitchellh/go-homedir"
	uuid "github.com/satori/go.uuid"
)

const (
	envSection      = "AKAMAI_EDGERC_SECTION"
	envEdgercConfig = "AKAMAI_EDGERC_CONFIG"
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

	// Sets default value for credentials configuration file
	// to be pointing to ~/.edgerc
	dir, _ := homedir.Dir()
	dir += string(os.PathSeparator) + ".edgerc"

	// Load the file based on our provided config
	edgerc, err := ini.Load(dir)
	if err != nil {
		return nil, fmt.Errorf("Error loading file? %s", err)
	}

	sectionNames := edgerc.SectionStrings()
	if !(stringInSlice(edgercSection, sectionNames)) {
		return nil, fmt.Errorf("Could not load section  %s", edgercSection)
	}

	edgercHost := edgerc.Section(edgercSection).Key("host").String()
	edgercclientToken := edgerc.Section(edgercSection).Key("client_token").String()
	edgercclientSecret := edgerc.Section(edgercSection).Key("client_secret").String()
	edgercaccessToken := edgerc.Section(edgercSection).Key("access_token").String()

	loadedCredentials := &EdgercCredentials{
		host:         edgercHost,
		clientToken:  edgercclientToken,
		clientSecret: edgercclientSecret,
		accessToken:  edgercaccessToken,
	}
	err = edgerc.Section(edgercSection).MapTo(loadedCredentials)
	if err != nil {
		return nil, fmt.Errorf("Error loading file? %s", err)
	}

	return loadedCredentials, nil

}

// AuthString takes prm and returns a string that can be
// used as the `Authorization` header in making Akamai API requests.
//
// The string returned by Auth conforms to the
// Akamai {OPEN} EdgeGrid Authentication scheme.
// https://developer.akamai.com/introduction/Client_Auth.html
func AuthString(eprm *EdgercCredentials, request *http.Request, headersToSign []string) string {

	u, err := uuid.NewV4()
	if err != nil {
		return ""
	}
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
