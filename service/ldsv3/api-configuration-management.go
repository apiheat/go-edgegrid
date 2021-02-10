package ldsv3

import (
	"fmt"
	"net/url"
	"path"
)

// UpdateLogConfiguration modifies a specific log delivery.
// You need to specify all the data members in the request, or missing members are removed from the configuration.
// The response’s Location header reflects where you can access the new configuration.
func (lds *Ldsv3) UpdateLogConfiguration(logConfigurationID string, opts ConfigurationOptions) (string, error) {
	if logConfigurationID == "" {
		return "", fmt.Errorf("Please provide log configuration ID")
	}

	apiURI := fmt.Sprintf("%s/log-configurations/%s", basePath, logConfigurationID)

	// Create and execute request
	resp, err := lds.Client.Rclient.R().
		//SetResult(ConfigurationsRespElem{}).
		SetError(LsdErrorv3{}).
		SetHeader("Content-Type", "application/json").
		SetBody(opts).
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
func (lds *Ldsv3) CopyLogConfiguration(logConfigurationID string, opts ConfigurationCopyOptions) (string, error) {
	if logConfigurationID == "" {
		return "", fmt.Errorf("Please provide log configuration ID")
	}

	apiURI := fmt.Sprintf("%s/log-configurations/%s/copy", basePath, logConfigurationID)

	// Create and execute request
	resp, err := lds.Client.Rclient.R().
		//SetResult(ConfigurationsRespElem{}).
		SetError(LsdErrorv3{}).
		SetHeader("Content-Type", "application/json").
		SetBody(opts).
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
