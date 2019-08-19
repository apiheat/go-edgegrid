package siteshieldv1

import (
	"fmt"
	"net/http"
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
func (sss *Siteshieldv1) GetMap(id string) (*SiteShieldMap, *ClientResponse, error) {
	qParams := QStrSiteShield{}
	path := fmt.Sprintf("%s/%s", SiteshieldPathV1, id)

	var respStruct *SiteShieldMap
	resp, err := nls.client.makeAPIRequest(http.MethodGet, path, qParams, &respStruct, nil, nil)

	return respStruct, resp, err
}

//AcknowledgeMap Acknowledges specific map based on ID
func (sss *Siteshieldv1) AcknowledgeMap(id string) (*SiteShieldMap, *ClientResponse, error) {
	qParams := QStrSiteShield{}
	path := fmt.Sprintf("%s/%s/acknowledge", SiteshieldPathV1, id)

	var respStruct *SiteShieldMap
	resp, err := nls.client.makeAPIRequest(http.MethodPost, path, qParams, &respStruct, nil, nil)

	return respStruct, resp, err
}
