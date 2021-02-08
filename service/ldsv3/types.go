package ldsv3

// SourcesResp
type SourcesResp []SourcesRespElem

// SourcesRespElem
type SourcesRespElem struct {
	Type     string   `json:"type"`
	ID       string   `json:"id"`
	CpCode   string   `json:"cpCode"`
	Products []string `json:"products"`
	Links    []struct {
		Rel    string `json:"rel"`
		Href   string `json:"href"`
		Method string `json:"method,omitempty"`
		Title  string `json:"title,omitempty"`
	} `json:"links"`
}

// ConfigurationsResp
type ConfigurationsResp []ConfigurationsRespElem

// ConfigurationsRespElem
type ConfigurationsRespElem struct {
	ID        int    `json:"id"`
	Status    string `json:"status"`
	StartDate string `json:"startDate"`
	LogSource struct {
		Links []struct {
			Rel  string `json:"rel"`
			Href string `json:"href"`
		} `json:"links"`
		Type             string   `json:"type"`
		ID               string   `json:"id"`
		CpCode           string   `json:"cpCode"`
		Products         []string `json:"products"`
		LogRetentionDays int      `json:"logRetentionDays"`
	} `json:"logSource"`
	AggregationDetails struct {
		Type              string `json:"type"`
		DeliveryFrequency struct {
			ID    string `json:"id"`
			Value string `json:"value"`
			Links []struct {
				Rel  string `json:"rel"`
				Href string `json:"href"`
			} `json:"links"`
		} `json:"deliveryFrequency"`
	} `json:"aggregationDetails"`
	ContactDetails struct {
		MailAddresses []string `json:"mailAddresses"`
		Contact       struct {
			ID    string `json:"id"`
			Value string `json:"value"`
			Links []struct {
				Rel  string `json:"rel"`
				Href string `json:"href"`
			} `json:"links"`
		} `json:"contact"`
	} `json:"contactDetails"`
	DeliveryDetails struct {
		Type         string `json:"type"`
		DomainPrefix string `json:"domainPrefix"`
		CpcodeID     int    `json:"cpcodeId"`
		Directory    string `json:"directory"`
	} `json:"deliveryDetails"`
	EncodingDetails struct {
		Encoding struct {
			ID    string `json:"id"`
			Value string `json:"value"`
			Links []struct {
				Rel  string `json:"rel"`
				Href string `json:"href"`
			} `json:"links"`
		} `json:"encoding"`
	} `json:"encodingDetails"`
	LogFormatDetails struct {
		LogIdentifier string `json:"logIdentifier"`
		LogFormat     struct {
			ID    string `json:"id"`
			Value string `json:"value"`
			Links []struct {
				Rel  string `json:"rel"`
				Href string `json:"href"`
			} `json:"links"`
		} `json:"logFormat"`
	} `json:"logFormatDetails"`
	MessageSize struct {
		ID    string `json:"id"`
		Value string `json:"value"`
		Links []struct {
			Rel  string `json:"rel"`
			Href string `json:"href"`
		} `json:"links"`
	} `json:"messageSize"`
	Links []struct {
		Rel    string `json:"rel"`
		Href   string `json:"href"`
		Title  string `json:"title,omitempty"`
		Method string `json:"method,omitempty"`
	} `json:"links"`
}
