package diagnosticv2

import (
	"fmt"
	"strconv"
)

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

// StartTranslateErrorAsync start async translation for given Akamai error code reference
func (dts *Diagnosticv2) StartTranslateErrorAsync(errorCode string) (*TranslateErrorAsync, error) {

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

// CheckIPAddress checks if given IP belongs to Akamai CDN
func (dts *Diagnosticv2) CheckIPAddress(ip string) (*VerifyIP, error) {

	// Create and execute request
	resp, err := dts.Client.Rclient.R().
		SetResult(VerifyIP{}).
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

	return resp.Result().(*VerifyIP), nil
}

// CreateDiagnosticLink generates user link and request
func (dts *Diagnosticv2) CreateDiagnosticLink(username, testURL string) (*DiagnosticLinkURL, error) {

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
func (dts *Diagnosticv2) ExecuteDig(obj string, requestFrom AkamaiRequestFrom, hostname, query string) (*DigResult, error) {

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
func (dts *Diagnosticv2) ExecuteMtr(obj string, requestFrom AkamaiRequestFrom, destinationDomain string, resolveDNS bool) (*MtrResult, error) {
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
func (dts *Diagnosticv2) ExecuteCurl(obj string, requestFrom AkamaiRequestFrom, testURL, userAgent string) (*CurlResult, error) {

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
