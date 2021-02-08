package ldsv3

import (
	"fmt"
)

// ListConfigurations list LDS configurations
func (lds *Ldsv3) ListConfigurations(sourceType string) (*ConfigurationsResp, error) {
	if sourceType == "" {
		return nil, fmt.Errorf("Missing argument 'sourceType'. Most probably you need cpcode-products as argument")
	}

	apiURI := fmt.Sprintf("%s/log-sources/%s/log-configurations", basePath, sourceType)

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

// ListSources list LDS sources
func (lds *Ldsv3) ListSources(sourceType string) (*SourcesResp, error) {
	if sourceType == "" {
		return nil, fmt.Errorf("Missing argument 'sourceType'. Most probably you need cpcode-products as argument")
	}

	apiURI := fmt.Sprintf("%s/log-sources/%s", basePath, sourceType)

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
