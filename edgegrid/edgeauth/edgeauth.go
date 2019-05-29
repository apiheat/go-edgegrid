package edgeauth

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
)

//Edgeauth
type Edgeauth struct {
}

//GetCredentialsFromEnv Retrieves credentials from env variables which are prefixed with 'AKAMAI_'
//In order to sucesfully build credentials file we need the following variables:
// AKAMAI_HOST
// AKAMAI_CLIENT_TOKEN
// AKAMAI_CLIENT_SECRET
// AKAMAI_ACCESS_TOKEN
// Returns new EdgercCredentials object or error
func (ea *Edgeauth) GetCredentialsFromEnv() (*EdgercCredentials, error) {

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
				envCredentials.Host = val
			case opt == "CLIENT_TOKEN":
				envCredentials.ClientToken = val
			case opt == "CLIENT_SECRET":
				envCredentials.ClientSecret = val
			case opt == "ACCESS_TOKEN":
				envCredentials.AccessToken = val
			}
		}
	}

	missing = removeStringDuplicates(missing)

	if len(missing) > 0 {
		log.Debug(fmt.Sprintf("[InitEdgerc]::Missing required environment variables: %s", missing))
	}

	err := validateCredentials(envCredentials)
	if err == nil {
		log.Debug("[InitEdgerc]::Return credentials object from env variables")
		return envCredentials, nil
	}

	log.Debug(fmt.Sprintf("[InitEdgerc]::Environment variables are not correct: %s", err))

	return nil, nil
}

//GetCredentialsFromFile Retrieves credentials from ( usually ~/.edgerc ) and returns
//new EdgercCredentials object or error
func (ea *Edgeauth) GetCredentialsFromFile() (*EdgercCredentials, error) {

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
				envCredentials.Host = val
			case opt == "CLIENT_TOKEN":
				envCredentials.ClientToken = val
			case opt == "CLIENT_SECRET":
				envCredentials.ClientSecret = val
			case opt == "ACCESS_TOKEN":
				envCredentials.AccessToken = val
			}
		}
	}

	missing = removeStringDuplicates(missing)

	if len(missing) > 0 {
		log.Debug(fmt.Sprintf("[InitEdgerc]::Missing required environment variables: %s", missing))
	}

	err := validateCredentials(envCredentials)
	if err == nil {
		log.Debug("[InitEdgerc]::Return credentials object from env variables")
		return envCredentials, nil
	}

	log.Debug(fmt.Sprintf("[InitEdgerc]::Environment variables are not correct: %s", err))

	return nil, nil
}

//GetCredentialsFromJSON Retrieves credentials from JSON string and returns
//new EdgercCredentials object or error
func (ea *Edgeauth) GetCredentialsFromJSON() (*EdgercCredentials, error) {

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
				envCredentials.Host = val
			case opt == "CLIENT_TOKEN":
				envCredentials.ClientToken = val
			case opt == "CLIENT_SECRET":
				envCredentials.ClientSecret = val
			case opt == "ACCESS_TOKEN":
				envCredentials.AccessToken = val
			}
		}
	}

	missing = removeStringDuplicates(missing)

	if len(missing) > 0 {
		log.Debug(fmt.Sprintf("[InitEdgerc]::Missing required environment variables: %s", missing))
	}

	err := validateCredentials(envCredentials)
	if err == nil {
		log.Debug("[InitEdgerc]::Return credentials object from env variables")
		return envCredentials, nil
	}

	log.Debug(fmt.Sprintf("[InitEdgerc]::Environment variables are not correct: %s", err))

	return nil, nil
}

//removeStringDuplicates removes duplicated elements from strings array
func removeStringDuplicates(str []string) []string {
	// Use map to record duplicates as we find them.
	encountered := map[string]bool{}
	result := []string{}

	for v := range str {
		if encountered[str[v]] == true {
			// Do not add duplicate.
		} else {
			// Record this element as an encountered element.
			encountered[str[v]] = true
			// Append to result slice.
			if str[v] != "" {
				result = append(result, str[v])
			}
		}
	}
	// Return the new slice.
	return result
}

//validateCredentials checks if credentials we have are correct
func validateCredentials(creds *EdgercCredentials) error {
	log.Debug("[validateCreds]::Validating credentials - 'host'")
	if creds.Host == "" {
		log.Error("[validateCreds]::'host' value is empty")
		return errors.New("'host' is empty")
	}

	valid := govalidator.IsURL(creds.Host)
	if !valid {
		log.Error("[validateCreds]::'host' is not a valid URL")
		return errors.New("'host' is not a valid URL")
	}

	u, err := url.Parse(creds.Host)
	if err != nil {
		log.Error("[validateCreds]::'host' cannot be parsed correctly")
		return errors.New("'host' cannot be parsed correctly")
	}

	if u.Scheme != "" {
		log.Error("[validateCreds]::contains URL scheme")
		return errors.New("'host' contains URL scheme")
	}

	log.Debug("[InitEdgerc]::Validating credentials - 'client_token'")
	if creds.ClientToken == "" {
		log.Error("[validateCreds]::'client_token' is empty")
		return errors.New("'client_token' is empty")
	}

	log.Debug("[InitEdgerc]::Validating credentials - 'client_secret'")
	if creds.ClientSecret == "" {
		log.Error("[validateCreds]::'client_secret' is empty")
		return errors.New("'client_secret' is empty")

	}

	log.Debug("[InitEdgerc]::Validating credentials - 'access_token'")
	if creds.AccessToken == "" {
		log.Error("[validateCreds]::'access_token' is empty")
		return errors.New("'access_token' is empty")
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
