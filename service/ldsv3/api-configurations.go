package ldsv3

import (
	"fmt"
	"net/url"
	"path"
)

// GetLogConfiguration retrieves a specific log delivery configuration.
func (lds *Ldsv3) GetLogConfiguration(logConfigurationID string) (*ConfigurationsRespElem, error) {
	if logConfigurationID == "" {
		return nil, fmt.Errorf("Please provide log configuration ID")
	}

	apiURI := fmt.Sprintf("%s/log-configurations/%s", basePath, logConfigurationID)

	// Create and execute request
	resp, err := lds.Client.Rclient.R().
		SetResult(ConfigurationsRespElem{}).
		SetError(LsdErrorv3{}).
		Get(apiURI)

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		e := resp.Error().(*LsdErrorv3)

		return nil, e
	}

	return resp.Result().(*ConfigurationsRespElem), nil
}

// UpdateLogConfiguration modifies a specific log delivery.
// You need to specify all the data members in the request, or missing members are removed from the configuration.
// The response’s Location header reflects where you can access the new configuration.
// You have to have top group account permissions for this call
func (lds *Ldsv3) UpdateLogConfiguration(logConfigurationID string, body ConfigurationBody) (string, error) {
	if logConfigurationID == "" {
		return "", fmt.Errorf("Please provide log configuration ID")
	}

	apiURI := fmt.Sprintf("%s/log-configurations/%s", basePath, logConfigurationID)

	// Create and execute request
	resp, err := lds.Client.Rclient.R().
		//SetResult(ConfigurationsRespElem{}).
		SetError(LsdErrorv3{}).
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Put(apiURI)

	fmt.Println(resp)
	if err != nil {
		return "", err
	}

	if resp.IsError() {
		e := resp.Error().(*LsdErrorv3)

		return "", e
	}

	headers := resp.Header()

	fmt.Println(headers)
	location := headers.Get("Location")

	return location, nil
}

// RemoveLogConfiguration deletes a specific log delivery configuration.
func (lds *Ldsv3) RemoveLogConfiguration(logConfigurationID string) error {
	if logConfigurationID == "" {
		return fmt.Errorf("Please provide log configuration ID")
	}

	apiURI := fmt.Sprintf("%s/log-configurations/%s", basePath, logConfigurationID)

	// Create and execute request
	resp, err := lds.Client.Rclient.R().
		//SetResult(ConfigurationsRespElem{}).
		SetError(LsdErrorv3{}).
		Delete(apiURI)

	if err != nil {
		return err
	}

	if resp.IsError() {
		e := resp.Error().(*LsdErrorv3)

		return e
	}

	return nil
}

// CopyLogConfiguration copies a specific log delivery configuration to a target log source to produce a new log delivery configuration.
// You have to have top group account permissions for this call
func (lds *Ldsv3) CopyLogConfiguration(logConfigurationID string, body ConfigurationCopyBody) (string, error) {
	if logConfigurationID == "" {
		return "", fmt.Errorf("Please provide log configuration ID")
	}

	apiURI := fmt.Sprintf("%s/log-configurations/%s/copy", basePath, logConfigurationID)

	// Create and execute request
	resp, err := lds.Client.Rclient.R().
		//SetResult(ConfigurationsRespElem{}).
		SetError(LsdErrorv3{}).
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Post(apiURI)

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

// SuspendLogConfiguration suspends log delivery for a specific configuration.
// You will not receive logs for this configuration while it is suspended.
func (lds *Ldsv3) SuspendLogConfiguration(logConfigurationID string) error {
	if logConfigurationID == "" {
		return fmt.Errorf("Please provide log configuration ID")
	}

	apiURI := fmt.Sprintf("%s/log-configurations/%s/suspend", basePath, logConfigurationID)

	// Create and execute request
	resp, err := lds.Client.Rclient.R().
		SetError(LsdErrorv3{}).
		Post(apiURI)

	if err != nil {
		return err
	}

	if resp.IsError() {
		e := resp.Error().(*LsdErrorv3)

		return e
	}

	return nil
}

// ResumeLogConfiguration resumes log delivery for a specific configuration.
func (lds *Ldsv3) ResumeLogConfiguration(logConfigurationID string) error {
	if logConfigurationID == "" {
		return fmt.Errorf("Please provide log configuration ID")
	}

	apiURI := fmt.Sprintf("%s/log-configurations/%s/resume", basePath, logConfigurationID)

	// Create and execute request
	resp, err := lds.Client.Rclient.R().
		SetError(LsdErrorv3{}).
		Post(apiURI)

	if err != nil {
		return err
	}

	if resp.IsError() {
		e := resp.Error().(*LsdErrorv3)

		return e
	}

	return nil
}

// CreateLogConfiguration creates new log configuration.
// The response’s Location header reflects where you can access the new configuration.
// You have to have top group account permissions for this call
func (lds *Ldsv3) CreateLogConfiguration(logCSourceID, logSourceType string, body ConfigurationBody) (string, error) {
	if logCSourceID == "" {
		return "", fmt.Errorf("Please provide log source ID")
	}

	if logSourceType == "" {
		return "", fmt.Errorf("Please provide log source type")
	}

	apiURI := fmt.Sprintf("%s/log-sources/%s/%s/log-configurations", basePath, logSourceType, logCSourceID)

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
