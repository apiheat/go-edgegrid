package edgegrid

import (
	"fmt"
	"net/http"
)

type SiteShieldService struct {
	client *Client
}

// AkamaiSiteShieldMapsResp response struct
type AkamaiSiteShieldMapsResp struct {
	SiteShieldMaps []AkamaiSiteShieldMap `json:"siteShieldMaps"`
}

// AkamaiSiteShieldMap struct
type AkamaiSiteShieldMap struct {
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

func (nls *SiteShieldService) ListMaps() (*AkamaiSiteShieldMapsResp, *ClientResponse, error) {
	var k *AkamaiSiteShieldMapsResp
	resp, err := nls.client.NewRequest(http.MethodGet, SiteshieldPathV1, nil, &k)

	return k, resp, err
}

func (nls *SiteShieldService) ListMap(id string) (*AkamaiSiteShieldMap, *ClientResponse, error) {
	apiURI := fmt.Sprintf("%s/%s", SiteshieldPathV1, id)

	var k *AkamaiSiteShieldMap
	resp, err := nls.client.NewRequest(http.MethodGet, apiURI, nil, &k)

	return k, resp, err
}

func (nls *SiteShieldService) AckMap(id string) (*AkamaiSiteShieldMap, *ClientResponse, error) {
	apiURI := fmt.Sprintf("%s/%s/acknowledge", SiteshieldPathV1, id)

	var k *AkamaiSiteShieldMap
	resp, err := nls.client.NewRequest(http.MethodPost, apiURI, nil, &k)

	return k, resp, err
}
