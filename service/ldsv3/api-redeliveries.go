package ldsv3

import (
	"fmt"
	"net/url"
	"path"
)

// GetLogRedelivery retrieves a specific log redelivery request.
func (lds *Ldsv3) GetLogRedelivery(redeliveryID string) (*LogRedeliveryElem, error) {
	if redeliveryID == "" {
		return nil, fmt.Errorf("Please provide log redelivery ID")
	}

	apiURI := fmt.Sprintf("%s/log-redeliveries/%s", basePath, redeliveryID)

	// Create and execute request
	resp, err := lds.Client.Rclient.R().
		SetResult(LogRedeliveryElem{}).
		SetError(LsdErrorv3{}).
		Get(apiURI)

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		e := resp.Error().(*LsdErrorv3)

		return nil, e
	}

	return resp.Result().(*LogRedeliveryElem), nil
}

// ListLogRedeliveries retrieves a list of requests to redeliver logs.
func (lds *Ldsv3) ListLogRedeliveries() (*LogRedeliveryResp, error) {
	apiURI := fmt.Sprintf("%s/log-redeliveries", basePath)

	// Create and execute request
	resp, err := lds.Client.Rclient.R().
		SetResult(LogRedeliveryResp{}).
		SetError(LsdErrorv3{}).
		Get(apiURI)

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		e := resp.Error().(*LsdErrorv3)

		return nil, e
	}

	return resp.Result().(*LogRedeliveryResp), nil
}

// CreateLogRedeliveries creates a new request to resend a log.
func (lds *Ldsv3) CreateLogRedeliveries(body RedeliveryBody) (string, error) {
	apiURI := fmt.Sprintf("%s/log-redeliveries", basePath)

	// Create and execute request
	resp, err := lds.Client.Rclient.R().
		SetError(LsdErrorv3{}).
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Post(apiURI)

	fmt.Println(resp)
	if err != nil {
		return "", err
	}

	if resp.IsError() {
		e := resp.Error().(*LsdErrorv3)

		return "", e
	}

	headers := resp.Header()
	location := headers.Get("Location")
	newConfigurationURL, err := url.Parse(location)

	if err != nil {
		return location, err
	}

	return path.Base(newConfigurationURL.Path), nil
}
