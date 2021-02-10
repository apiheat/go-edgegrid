package ldsv3

import "fmt"

// ListLogEncodingsByType retrieves all allowable log encodings.
func (lds *Ldsv3) ListLogEncodingsByType(logSourceType, deliveryType string) (*ConfigurationParameterResp, error) {
	if logSourceType == "" {
		return nil, fmt.Errorf("Missing argument 'logSourceType'. Most probably you need cpcode-products as argument")
	}

	apiURI := fmt.Sprintf("%s/log-sources/%s/encodings", basePath, logSourceType)

	query := map[string]string{}

	if deliveryType != "" {
		query["deliveryType"] = deliveryType
	}

	// Create and execute request
	resp, err := lds.Client.Rclient.R().
		SetResult(ConfigurationParameterResp{}).
		SetError(LsdErrorv3{}).
		SetQueryParams(query).
		Get(apiURI)

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		e := resp.Error().(*LsdErrorv3)

		return nil, e
	}

	return resp.Result().(*ConfigurationParameterResp), nil
}

// ListLogFormatPerID gets log formats of given logSourceType and logSourceId.
func (lds *Ldsv3) ListLogFormatPerID(logSourceID, logSourceType string) (*ConfigurationParameterResp, error) {
	if logSourceType == "" {
		return nil, fmt.Errorf("Missing argument 'logSourceType'. Most probably you need cpcode-products as argument")
	}

	if logSourceID == "" {
		return nil, fmt.Errorf("Missing argument 'logSourceID'")
	}

	apiURI := fmt.Sprintf("%s/log-sources/%s/%s/log-formats", basePath, logSourceType, logSourceID)

	// Create and execute request
	resp, err := lds.Client.Rclient.R().
		SetResult(ConfigurationParameterResp{}).
		SetError(LsdErrorv3{}).
		Get(apiURI)

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		e := resp.Error().(*LsdErrorv3)

		return nil, e
	}

	return resp.Result().(*ConfigurationParameterResp), nil
}

// ListLogFormatByType returns all available log formats for the specified logSourceType type.
// You need the Id of log format to create new log delivery configurations.
func (lds *Ldsv3) ListLogFormatByType(logSourceType string) (*ConfigurationParameterResp, error) {
	if logSourceType == "" {
		return nil, fmt.Errorf("Missing argument 'logSourceType'. Most probably you need cpcode-products as argument")
	}

	apiURI := fmt.Sprintf("%s/log-sources/%s/log-formats", basePath, logSourceType)

	// Create and execute request
	resp, err := lds.Client.Rclient.R().
		SetResult(ConfigurationParameterResp{}).
		SetError(LsdErrorv3{}).
		Get(apiURI)

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		e := resp.Error().(*LsdErrorv3)

		return nil, e
	}

	return resp.Result().(*ConfigurationParameterResp), nil
}

// ListLogConfigurationsByType returns all log delivery configurations of a given logSourceType.
// You would need the logConfigurationId to modify a log delivery configuration.
func (lds *Ldsv3) ListLogConfigurationsByType(logSourceType string) (*ConfigurationsResp, error) {
	if logSourceType == "" {
		return nil, fmt.Errorf("Missing argument 'logSourceType'. Most probably you need cpcode-products as argument")
	}

	apiURI := fmt.Sprintf("%s/log-sources/%s/log-configurations", basePath, logSourceType)

	// Create and execute request
	resp, err := lds.Client.Rclient.R().
		SetResult(ConfigurationsResp{}).
		SetError(LsdErrorv3{}).
		Get(apiURI)

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		e := resp.Error().(*LsdErrorv3)

		return nil, e
	}

	return resp.Result().(*ConfigurationsResp), nil
}

// ListLogConfigurationsPerID gets all log configurations of given logSourceType and logSourceId.
func (lds *Ldsv3) ListLogConfigurationsPerID(logSourceID, logSourceType string) (*ConfigurationsResp, error) {
	if logSourceType == "" {
		return nil, fmt.Errorf("Missing argument 'logSourceType'. Most probably you need cpcode-products as argument")
	}

	if logSourceID == "" {
		return nil, fmt.Errorf("Missing argument 'logSourceID'")
	}

	apiURI := fmt.Sprintf("%s/log-sources/%s/%s/log-configurations", basePath, logSourceType, logSourceID)

	// Create and execute request
	resp, err := lds.Client.Rclient.R().
		SetResult(ConfigurationsResp{}).
		SetError(LsdErrorv3{}).
		Get(apiURI)

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		e := resp.Error().(*LsdErrorv3)

		return nil, e
	}

	return resp.Result().(*ConfigurationsResp), nil
}

// ListSources returns all log sources (logSourceType) and log source ID (logSourceId) to which the user has access.
// You need the logSourceType and logSourceId to create a log delivery configuration.
func (lds *Ldsv3) ListSources() (*SourcesResp, error) {
	apiURI := fmt.Sprintf("%s/log-sources", basePath)

	// Create and execute request
	resp, err := lds.Client.Rclient.R().
		SetResult(SourcesResp{}).
		SetError(LsdErrorv3{}).
		Get(apiURI)

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		e := resp.Error().(*LsdErrorv3)

		return nil, e
	}

	return resp.Result().(*SourcesResp), nil
}

// ListSourcesByType returns all log sources of the specified logSourceType,
// one of cpcode-products, gtm-properties, edns-zones, or answerx-objects.
func (lds *Ldsv3) ListSourcesByType(logSourceType string) (*SourcesResp, error) {
	if logSourceType == "" {
		return nil, fmt.Errorf("Missing argument 'logSourceType'. Most probably you need cpcode-products as argument")
	}

	apiURI := fmt.Sprintf("%s/log-sources/%s", basePath, logSourceType)

	// Create and execute request
	resp, err := lds.Client.Rclient.R().
		SetResult(SourcesResp{}).
		SetError(LsdErrorv3{}).
		Get(apiURI)

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		e := resp.Error().(*LsdErrorv3)

		return nil, e
	}

	return resp.Result().(*SourcesResp), nil
}

// GetLogSource gets a log source of a given logSourceType type and logSourceId.
func (lds *Ldsv3) GetLogSource(logSourceID, logSourceType string) (*SourcesRespElem, error) {
	if logSourceID == "" {
		return nil, fmt.Errorf("Please provide log source ID")
	}

	if logSourceType == "" {
		return nil, fmt.Errorf("Please provide log source type")
	}

	apiURI := fmt.Sprintf("%s/log-sources/%s/%s", basePath, logSourceType, logSourceID)

	// Create and execute request
	resp, err := lds.Client.Rclient.R().
		SetResult(SourcesRespElem{}).
		SetError(LsdErrorv3{}).
		Get(apiURI)

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		e := resp.Error().(*LsdErrorv3)

		return nil, e
	}

	return resp.Result().(*SourcesRespElem), nil
}
