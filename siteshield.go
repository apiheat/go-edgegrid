package edgegrid

import (
	"fmt"
	"net/http"
)

type SiteShieldService struct {
	client *Client
}

// SiteShieldMapsResp response struct
type SiteShieldMapsResp struct {
	SiteShieldMaps []SiteShieldMap `json:"siteShieldMaps"`
}

// SiteShieldMap struct
type SiteShieldMap struct {
	AcknowledgeRequiredBy int64    `json:"acknowledgeRequiredBy"`
	Acknowledged          bool     `json:"acknowledged"`
	AcknowledgedBy        string   `json:"acknowledgedBy"`
	AcknowledgedOn        int64    `json:"acknowledgedOn"`
	Contacts              []string `json:"contacts"`
	CurrentCidrs          []string `json:"currentCidrs"`
	ID                    int      `json:"id"`
	LatestTicketID        int      `json:"latestTicketId"`
	MapAlias              string   `json:"mapAlias"`
	McmMapRuleID          int      `json:"mcmMapRuleId"`
	ProposedCidrs         []string `json:"proposedCidrs"`
	RuleName              string   `json:"ruleName"`
	Service               string   `json:"service"`
	Shared                bool     `json:"shared"`
	Type                  string   `json:"type"`
}

//ListMaps Lists siteshield maps
func (nls *SiteShieldService) ListMaps() (*SiteShieldMapsResp, *ClientResponse, error) {
	qParams := QStrNetworkList{}

	var respStruct *SiteShieldMapsResp
	resp, err := nls.client.makeAPIRequest(http.MethodGet, SiteshieldPathV1, qParams, &respStruct, nil, nil)

	return respStruct, resp, err
}

//ListMap Retrieves specific map based on ID
func (nls *SiteShieldService) ListMap(id string) (*SiteShieldMap, *ClientResponse, error) {
	qParams := QStrNetworkList{}
	path := fmt.Sprintf("%s/%s", SiteshieldPathV1, id)

	var respStruct *SiteShieldMap
	resp, err := nls.client.makeAPIRequest(http.MethodGet, path, qParams, &respStruct, nil, nil)

	return respStruct, resp, err
}

//AckMap Acknowledges specific map based on ID
func (nls *SiteShieldService) AckMap(id string) (*SiteShieldMap, *ClientResponse, error) {
	qParams := QStrNetworkList{}
	path := fmt.Sprintf("%s/%s/acknowledge", SiteshieldPathV1, id)

	var respStruct *SiteShieldMap
	resp, err := nls.client.makeAPIRequest(http.MethodPost, path, qParams, &respStruct, nil, nil)

	return respStruct, resp, err
}
