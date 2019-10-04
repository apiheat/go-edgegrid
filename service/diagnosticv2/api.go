package diagnosticv2

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
)

// IsStringInSlice returns TRUE is slice contains string and false if not
func isStringInSlice(a string, list []string) bool {
	// We need that to not filter for empty list
	if len(list) > 0 {
		for _, b := range list {
			if b == a {
				return true
			}
		}
		return false
	}
	return true
}

func executeFromSourceSupported(str string) bool {
	if isStringInSlice(str, []string{"ghost-locations", "ip-addresses"}) {
		return true
	}

	return false
}

//ListGhostLocations returns location for ghost servers
func (dts *Diagnosticv2) ListGhostLocations() (*GhostLocations, error) {

	// Create and execute request
	resp, err := dts.Client.Rclient.R().
		SetResult(GhostLocations{}).
		SetError(DiagnosticErrorv2{}).
		Get(fmt.Sprintf("%s/ghost-locations/available", basePath))

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		e := resp.Error().(*DiagnosticErrorv2)
		if e.Status != 0 {
			return nil, e
		}
	}

	return resp.Result().(*GhostLocations), nil

}

// LaunchTranslateErrorAsync start async translation for given Akamai error code reference
func (dts *Diagnosticv2) LaunchTranslateErrorAsync(errorCode string) (*TranslateErrorAsync, error) {

	// Create and execute request
	resp, err := dts.Client.Rclient.R().
		SetResult(TranslateErrorAsync{}).
		SetError(DiagnosticErrorv2{}).
		Post(fmt.Sprintf("%s/errors/%s/translate-error", basePath, errorCode))

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		e := resp.Error().(*DiagnosticErrorv2)
		if e.Status != 0 {
			return nil, e
		}
	}

	return resp.Result().(*TranslateErrorAsync), nil
}

// CheckTranslateErrorAsync polls for status of the StartTranslateErrorAsync returned request id
func (dts *Diagnosticv2) CheckTranslateErrorAsync(requestID string) (*TranslateErrorAsync, error) {

	// Create and execute request
	resp, err := dts.Client.Rclient.R().
		SetResult(TranslateErrorAsync{}).
		SetError(DiagnosticErrorv2{}).
		Get(fmt.Sprintf("%s/translate-error-requests/%s", basePath, requestID))

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		e := resp.Error().(*DiagnosticErrorv2)
		if e.Status != 0 {
			return nil, e
		}
	}

	return resp.Result().(*TranslateErrorAsync), nil

}

// RetrieveTranslateErrorAsync retrieves translated error message from Akamai platform
// https://developer.akamai.com/api/core_features/diagnostic_tools/v2.html#gettranslateerrorperrequest
func (dts *Diagnosticv2) RetrieveTranslateErrorAsync(requestID string) (*TranslatedError, error) {

	// Create and execute request
	resp, err := dts.Client.Rclient.R().
		SetResult(TranslatedError{}).
		SetError(DiagnosticErrorv2{}).
		Get(fmt.Sprintf("%s/translate-error-requests/%s/translated-error", basePath, requestID))

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		e := resp.Error().(*DiagnosticErrorv2)
		if e.Status != 0 {
			return nil, e
		}
	}

	return resp.Result().(*TranslatedError), nil

}

// TranslateErrorAsync will make request and wait for response
func (dts *Diagnosticv2) TranslateErrorAsync(errorCode string, retries int) (*TranslatedError, error) {
	count := retries
	// Create and execute request
	resp, err := dts.Client.Rclient.R().
		SetResult(TranslateErrorAsync{}).
		SetError(DiagnosticErrorv2{}).
		Post(fmt.Sprintf("%s/errors/%s/translate-error", basePath, errorCode))

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		e := resp.Error().(*DiagnosticErrorv2)
		if e.Status != 0 {
			return nil, e
		}
	}

	req := resp.Result().(*TranslateErrorAsync)
	requestID := req.RequestID
	log.Debugf("Request for error code translation was submitted. Request ID is %s", requestID)

	if resp.StatusCode() == http.StatusTooManyRequests {
		log.Debugf("Request limit per 60 seconds reached. Will wait for a minute")
		time.Sleep(61 * time.Second)
	}

	log.Debugf("Polling error code in %d seconds", req.RetryAfter)
	time.Sleep(time.Duration(req.RetryAfter+1) * time.Second)

	// Check request
	// With requestId and retryAfter data we can try to poll data
	log.Debugf("Making Translate Error request for ID: %s. Attempt 1 out of %d", requestID, retries)
	response, err := dts.Client.Rclient.R().
		SetResult(TranslatedError{}).
		SetError(DiagnosticErrorv2{}).
		Get(fmt.Sprintf("%s/translate-error-requests/%s/translated-error", basePath, requestID))

	count -= 2

	if response.StatusCode() == http.StatusBadRequest {
		return nil, response.Error().(*DiagnosticErrorv2)
	}

	if err != nil || response.StatusCode() != http.StatusOK {
		for {
			log.Debugf("Polling error code in %d seconds", req.RetryAfter)
			time.Sleep(time.Duration(req.RetryAfter+1) * time.Second)

			log.Debugf("Making Translate Error request for ID: %s. Attempt %d out of %d", requestID, retries-count, retries)

			count--

			response, err = dts.Client.Rclient.R().
				SetResult(TranslatedError{}).
				SetError(DiagnosticErrorv2{}).
				Get(fmt.Sprintf("%s/translate-error-requests/%s/translated-error", basePath, requestID))

			if err != nil {
				return nil, err
			}

			if response.StatusCode() == http.StatusBadRequest {
				return nil, response.Error().(*DiagnosticErrorv2)
			}

			if response.StatusCode() == http.StatusForbidden {
				return nil, response.Error().(*DiagnosticErrorv2)
			}

			if response.StatusCode() == http.StatusOK {
				break
			}

			if count == 0 {
				return nil, DiagnosticErrorv2{Detail: "Operation took too long. Exiting..."}
			}
		}
	}

	return response.Result().(*TranslatedError), nil
}

// CheckIPAddress checks if given IP belongs to Akamai CDN
func (dts *Diagnosticv2) CheckIPAddress(ip string) (*CDNStatus, error) {

	// Create and execute request
	resp, err := dts.Client.Rclient.R().
		SetResult(CDNStatus{}).
		SetError(DiagnosticErrorv2{}).
		Get(fmt.Sprintf("%s/ip-addresses/%s/is-cdn-ip", basePath, ip))

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		e := resp.Error().(*DiagnosticErrorv2)
		if e.Status != 0 {
			return nil, e
		}
	}

	return resp.Result().(*CDNStatus), nil
}

// CreateDiagnosticLink generates user link and request
func (dts *Diagnosticv2) GenerateDiagnosticLink(username, testURL string) (*DiagnosticLinkURL, error) {

	diagnosticLinkRequest := DiagnosticLinkRequest{
		EndUserName: username,
		URL:         testURL,
	}

	// Create and execute request
	resp, err := dts.Client.Rclient.R().
		SetBody(diagnosticLinkRequest).
		SetResult(DiagnosticLinkURL{}).
		SetError(DiagnosticErrorv2{}).
		Post(fmt.Sprintf("%s/end-users/diagnostic-url", basePath))

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		e := resp.Error().(*DiagnosticErrorv2)
		if e.Status != 0 {
			return nil, e
		}
	}

	return resp.Result().(*DiagnosticLinkURL), nil
}

// ListDiagnosticLinkRequests lists all requests
func (dts *Diagnosticv2) ListDiagnosticLinkRequests() (*DiagnosticLinkRequests, error) {
	// Create and execute request
	resp, err := dts.Client.Rclient.R().
		SetResult(DiagnosticLinkRequests{}).
		SetError(DiagnosticErrorv2{}).
		Get(fmt.Sprintf("%s/end-users/ip-requests", basePath))

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		e := resp.Error().(*DiagnosticErrorv2)
		if e.Status != 0 {
			return nil, e
		}
	}

	return resp.Result().(*DiagnosticLinkRequests), nil
}

// RetrieveDiagnosticLinkRequest gets request details
func (dts *Diagnosticv2) RetrieveDiagnosticLinkRequest(id string) (*DiagnosticLinkResult, error) {

	// Create and execute request
	resp, err := dts.Client.Rclient.R().
		SetResult(DiagnosticLinkResult{}).
		SetError(DiagnosticErrorv2{}).
		Get(fmt.Sprintf("%s/end-users/ip-requests/%s/ip-details", basePath, id))

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		e := resp.Error().(*DiagnosticErrorv2)
		if e.Status != 0 {
			return nil, e
		}
	}

	return resp.Result().(*DiagnosticLinkResult), nil
}

// RetrieveIPGeolocation provides given IP geolocation details
func (dts *Diagnosticv2) RetrieveIPGeolocation(ip string) (*Geolocation, error) {
	// Create and execute request
	resp, err := dts.Client.Rclient.R().
		SetResult(Geolocation{}).
		SetError(DiagnosticErrorv2{}).
		Get(fmt.Sprintf("%s/ip-addresses/%s/geo-location", basePath, ip))

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		e := resp.Error().(*DiagnosticErrorv2)
		if e.Status != 0 {
			return nil, e
		}
	}

	return resp.Result().(*Geolocation), nil
}

// ExecuteDig against a hostname to get DNS information, associating hostnames and IP addresses, from an IP address within the Akamai network not local to you. Specify the hostName as a query parameter, and an optional DNS queryType. See the Dig object for details on the response data.
func (dts *Diagnosticv2) ExecuteDig(obj, requestFrom, hostname, query string) (*DigResult, error) {
	if !executeFromSourceSupported(requestFrom) {
		return nil, fmt.Errorf("requestFrom value should be one of ['ghost-locations', 'ip-addresses'], you provided %s", requestFrom)
	}

	// Create and execute request
	resp, err := dts.Client.Rclient.R().
		SetQueryParams(map[string]string{
			"hostName":  hostname,
			"queryType": query,
		}).
		SetResult(DigResult{}).
		SetError(DiagnosticErrorv2{}).
		Get(fmt.Sprintf("%s/%s/%s/dig-info", basePath, requestFrom, obj))

		// /diagnostic-tools/v2/ip-addresses/{ipAddress}/dig-info{?hostName,queryType}
		// /diagnostic-tools/v2/ghost-locations/{locationId}/dig-info{?hostName,queryType}

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		e := resp.Error().(*DiagnosticErrorv2)
		if e.Status != 0 {
			return nil, e
		}
	}

	return resp.Result().(*DigResult), nil
}

// ExecuteMtr provides mtr functionality
func (dts *Diagnosticv2) ExecuteMtr(obj, requestFrom, destinationDomain string, resolveDNS bool) (*MtrResult, error) {
	if !executeFromSourceSupported(requestFrom) {
		return nil, fmt.Errorf("requestFrom value should be one of ['ghost-locations', 'ip-addresses'], you provided %s", requestFrom)
	}

	// Create and execute request
	resp, err := dts.Client.Rclient.R().
		SetQueryParams(map[string]string{
			"resolveDns":        strconv.FormatBool(resolveDNS),
			"destinationDomain": destinationDomain,
		}).
		SetResult(MtrResult{}).
		SetError(DiagnosticErrorv2{}).
		Get(fmt.Sprintf("%s/%s/%s/mtr-data", basePath, requestFrom, obj))

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		e := resp.Error().(*DiagnosticErrorv2)
		if e.Status != 0 {
			return nil, e
		}
	}

	return resp.Result().(*MtrResult), nil
}

// ExecuteCurl provides curl functionality
func (dts *Diagnosticv2) ExecuteCurl(obj, requestFrom, testURL, userAgent string) (*CurlResult, error) {
	if !executeFromSourceSupported(requestFrom) {
		return nil, fmt.Errorf("requestFrom value should be one of ['ghost-locations', 'ip-addresses'], you provided %s", requestFrom)
	}

	curlRequest := CurlRequest{
		UserAgent: userAgent,
		URL:       testURL,
	}

	// Create and execute request
	resp, err := dts.Client.Rclient.R().
		SetBody(curlRequest).
		SetResult(CurlResult{}).
		SetError(DiagnosticErrorv2{}).
		Post(fmt.Sprintf("%s/%s/%s/curl-results", basePath, requestFrom, obj))

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		e := resp.Error().(*DiagnosticErrorv2)
		if e.Status != 0 {
			return nil, e
		}
	}

	return resp.Result().(*CurlResult), nil
}

// ListGTMProperties provides available GTM properties
func (dts *Diagnosticv2) ListGTMProperties() (*GTMPropertiesResult, error) {
	// Create and execute request
	resp, err := dts.Client.Rclient.R().
		SetResult(GTMPropertiesResult{}).
		SetError(DiagnosticErrorv2{}).
		Get(fmt.Sprintf("%s/gtm/gtm-properties", basePath))

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		e := resp.Error().(*DiagnosticErrorv2)
		if e.Status != 0 {
			return nil, e
		}
	}

	return resp.Result().(*GTMPropertiesResult), nil
}

// ListGTMPropertyIPs provides available GTM properties
func (dts *Diagnosticv2) ListGTMPropertyIPs(property, domain string) (*GTMPropertyIpsResult, error) {

	if property == "" {
		return nil, fmt.Errorf("'property' is required parameter: '%s'", property)
	}

	if domain == "" {
		return nil, fmt.Errorf("'domain' is required parameter: '%s'", domain)
	}

	resp, err := dts.Client.Rclient.R().
		SetResult(GTMPropertyIpsResult{}).
		SetError(DiagnosticErrorv2{}).
		Get(fmt.Sprintf("%s/gtm/%s/%s/gtm-property-ips", basePath, property, domain))

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		e := resp.Error().(*DiagnosticErrorv2)
		if e.Status != 0 {
			return nil, e
		}
	}

	return resp.Result().(*GTMPropertyIpsResult), nil
}
