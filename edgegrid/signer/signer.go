package signer

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/apiheat/go-edgegrid/v6/edgegrid"
	uuid "github.com/satori/go.uuid"
)

const (
	moniker string = "EG1-HMAC-SHA256"
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

// SignRequest returns a string that can be used as the `Authorization` header
// when making Akamai API requests.
//
// The string returned by this method conforms to the
// Akamai {OPEN} EdgeGrid Authentication scheme.
// https://developer.akamai.com/introduction/Client_Auth.html
func (sr *SignatureRequest) SignRequest(rrq *http.Request, headersToSign []string) string {

	nonce := generateNonce()
	timestamp := generateTimestamp()

	var auth bytes.Buffer

	joinedPairs := []string{
		"client_token=" + sr.creds.ClientToken,
		"access_token=" + sr.creds.AccessToken,
		"timestamp=" + timestamp,
		"nonce=" + nonce,
	}

	auth.WriteString(moniker + " " + strings.Join(joinedPairs, ";") + ";")

	dataToSign := generateDataToSign(rrq, auth.String(), []string{})
	signingKey := generateSigningKey(timestamp, sr.creds.ClientSecret)

	signature := concat([]string{
		"signature=",
		base64HmacSha256(dataToSign, signingKey),
	})

	auth.WriteString(signature)

	return auth.String()
}

// generateTimestamp retrurns timestamp in the
// format of “yyyyMMddTHH:mm:ss+0000” as required by Akamai network
func generateTimestamp() string {
	timestamp := time.Now().UTC().Format("20060102T15:04:05+0000")

	return timestamp
}

// generateNonce is a random string used to detect replayed request messages.
// A UUID is always randomly generated
func generateNonce() string {
	return uuid.NewV4().String()
}

func base64Sha256(data []byte) string {
	h := sha256.New()

	h.Write(data)

	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func base64HmacSha256(message, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))

	h.Write([]byte(message))

	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func generateDataToSign(rrq *http.Request, authHeader string, headersToSign []string) string {

	var data bytes.Buffer
	values := []string{
		rrq.Method,
		rrq.URL.Scheme,
		rrq.Host,
		urlPathWithQuery(rrq.URL.Path, rrq.URL.RawQuery),
		"", //TODO: to be implemented - not required in initial stage
		makeContentHash(rrq),
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
		// Make sure we do have body to build content from
		if req.Body == nil {
			return ""
		}
		buf, err := ioutil.ReadAll(req.Body)
		req.Body.Close() //  must close

		if err != nil {
			//TODO: log here
			panic(err)
		}

		// Correct body setup based on https://github.com/go-resty/resty/issues/252
		req.Body = ioutil.NopCloser(bytes.NewBuffer(buf))
		return base64Sha256(buf)
	}

	return ""
}

func generateSigningKey(timestamp, clientSecret string) string {
	return base64HmacSha256(timestamp, clientSecret)
}

func urlPathWithQuery(path, queryParams string) string {

	if queryParams != "" {
		return fmt.Sprintf("%s?%s", path, queryParams)
	}

	return path
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
