package signer

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/apiheat/go-edgegrid/edgegrid"
	uuid "github.com/satori/go.uuid"
)

// SignatureRequest represents object which is used to sign request
type SignatureRequest struct {
	creds  *edgegrid.Credentials
	host   string
	scheme string
}

//New takes all required parameters and returns the required auth header
func New(cr *edgegrid.Credentials, scheme, host string) SignatureRequest {
	signatureRequest := SignatureRequest{
		creds:  cr,
		host:   host,
		scheme: scheme,
	}

	return signatureRequest
}

// Akamai {OPEN} EdgeGrid Authentication Service
type reader struct {
	*bytes.Buffer
}

func (m reader) Close() error { return nil }

// AuthString takes prm and returns a string that can be
// used as the `Authorization` header in making Akamai API requests.
//
// The string returned by Auth conforms to the
// Akamai {OPEN} EdgeGrid Authentication scheme.
// https://developer.akamai.com/introduction/Client_Auth.html
func (sr *SignatureRequest) AuthString(ecr *edgegrid.Credentials, method, host, scheme, urlpath string, headersToSign []string) string {

	u := uuid.NewV4()

	nonce := u.String()

	timestamp := time.Now().UTC().Format("20060102T15:04:05+0000")

	var auth bytes.Buffer
	orderedKeys := []string{"client_token", "access_token", "timestamp", "nonce"}

	m := map[string]string{
		orderedKeys[0]: ecr.ClientToken,
		orderedKeys[1]: ecr.AccessToken,
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

	data2sign := makeDataToSignv2(method, host, scheme, urlpath, auth.String(), []string{})
	signingKey := makeSigningKey(timestamp, ecr.ClientSecret)
	ah := concat([]string{
		"signature=",
		base64HmacSha256(data2sign, signingKey),
	})

	auth.WriteString(ah)

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

func makeDataToSignv2(method, host, scheme, urlpath, authHeader string, headersToSign []string) string {
	var data bytes.Buffer
	values := []string{
		method,
		scheme,
		host,
		urlpath,
		"", // canonical headers 2 sign
		"", // makeContentHash(request),
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
