package siteshieldv1

// SiteShieldMaps
type SiteShieldMaps struct {
	SiteShieldMaps []SiteShieldMap `json:"siteShieldMaps"`
}

// SiteShieldMap
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
