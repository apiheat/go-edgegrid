package diagnosticv2

import (
	"fmt"
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

// TranslateErrorAsync start async translation for given Akamai error code reference
func (dts *Diagnosticv2) TranslateErrorAsync(errorCode string) (*TranslateErrorAsync, error) {

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
