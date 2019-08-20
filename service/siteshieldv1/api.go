package siteshieldv1

import (
	"fmt"
)

//ListMaps Lists siteshield maps available in the account
func (sss *Siteshieldv1) ListMaps() (*SiteShieldMaps, error) {
	// Create and execute request
	resp, err := sss.Client.Rclient.R().
		SetResult(SiteShieldMaps{}).
		SetError(SiteshieldErrorv1{}).
		Get(basePath)

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		e := resp.Error().(*SiteshieldErrorv1)
		if e.Status != 0 {
			return nil, e
		}
	}

	return resp.Result().(*SiteShieldMaps), nil
}

//GetMap Retrieves specific map based on ID
func (sss *Siteshieldv1) GetMap(id string) (*SiteShieldMap, error) {
	// Create and execute request
	resp, err := sss.Client.Rclient.R().
		SetResult(SiteShieldMap{}).
		SetError(SiteshieldErrorv1{}).
		Get(fmt.Sprintf("%s/%s", basePath, id))

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		e := resp.Error().(*SiteshieldErrorv1)
		if e.Status != 0 {
			return nil, e
		}
	}

	return resp.Result().(*SiteShieldMap), nil
}

//AcknowledgeMap Acknowledges specific map based on ID
func (sss *Siteshieldv1) AcknowledgeMap(id string) (*SiteShieldMap, error) {
	// Create and execute request
	resp, err := sss.Client.Rclient.R().
		SetResult(SiteShieldMap{}).
		SetError(SiteshieldErrorv1{}).
		Post(fmt.Sprintf("%s/%s/acknowledge", basePath, id))

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		e := resp.Error().(*SiteshieldErrorv1)
		if e.Status != 0 {
			return nil, e
		}
	}

	return resp.Result().(*SiteShieldMap), nil

}
