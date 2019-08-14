package diagnosticv2

import (
	"fmt"
	"net/http"
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

// VerifyIPAddress checks if given IP belongs to Akamai CDN
func (dts *Diagnosticv2) VerifyIPAddress(ip string) (*VerifyIP, error) {

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

// RetrieveIPGeolocation provides given geolocation details based on given IP
func (dts *Diagnosticv2) RetrieveIPGeolocation(ip string) (*DTGeolocation, *ClientResponse, error) {
	qParams := QStrDiagTools{}
	path := fmt.Sprintf("%s/ip-addresses/%s/geo-location", DTPathV2, ip)

	var respStruct *DTGeolocation
	resp, err := nls.client.makeAPIRequest(http.MethodGet, path, qParams, &respStruct, nil, nil)

	return respStruct, resp, err
}
