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
		Type              string                     `json:"type"`
		DeliveryFrequency ConfigurationParameterElem `json:"deliveryFrequency"`
	} `json:"aggregationDetails"`
	ContactDetails struct {
		MailAddresses []string                   `json:"mailAddresses"`
		Contact       ConfigurationParameterElem `json:"contact"`
	} `json:"contactDetails"`
	DeliveryDetails struct {
		Type         string `json:"type"`
		DomainPrefix string `json:"domainPrefix"`
		CpcodeID     int    `json:"cpcodeId"`
		Directory    string `json:"directory"`
	} `json:"deliveryDetails"`
	EncodingDetails struct {
		Encoding ConfigurationParameterElem `json:"encoding"`
	} `json:"encodingDetails"`
	LogFormatDetails struct {
		LogIdentifier string                     `json:"logIdentifier"`
		LogFormat     ConfigurationParameterElem `json:"logFormat"`
	} `json:"logFormatDetails"`
	MessageSize ConfigurationParameterElem `json:"messageSize"`
	Links       []struct {
		Rel    string `json:"rel"`
		Href   string `json:"href"`
		Title  string `json:"title,omitempty"`
		Method string `json:"method,omitempty"`
	} `json:"links"`
}

// ConfigurationParameterResp
type ConfigurationParameterResp []ConfigurationParameterElem

// ConfigurationParameterElem
type ConfigurationParameterElem struct {
	ID    string `json:"id"`
	Value string `json:"value"`
	Links []struct {
		Rel  string `json:"rel"`
		Href string `json:"href"`
	} `json:"links"`
}

// LogRedeliveryResp
type LogRedeliveryResp []LogRedeliveryElem

// LogRedeliveryElem
type LogRedeliveryElem struct {
	LogConfiguration struct {
		ID    string `json:"id"`
		Links []struct {
			Rel  string `json:"rel"`
			Href string `json:"href"`
		} `json:"links"`
	} `json:"logConfiguration"`
	ID             string `json:"id"`
	BeginTime      int    `json:"beginTime"`
	EndTime        int    `json:"endTime"`
	RedeliveryDate string `json:"redeliveryDate"`
	Status         string `json:"status"`
	CreatedDate    string `json:"createdDate"`
	ModifiedDate   string `json:"modifiedDate"`
	Links          []struct {
		Rel    string `json:"rel"`
		Href   string `json:"href"`
		Title  string `json:"title,omitempty"`
		Method string `json:"method,omitempty"`
	} `json:"links"`
}

// LogConfigurationOptions
type ConfigurationOptions struct {
	StartDate          string                          `json:"startDate"`
	LogSource          ConfigurationMember             `json:"logSource"`
	ContactDetails     ConfigurationContactDetails     `json:"contactDetails"`
	LogFormatDetails   ConfigurationLogFormatDetails   `json:"logFormatDetails"`
	MessageSize        ConfigurationMember             `json:"messageSize"`
	AggregationDetails ConfigurationAggregationDetails `json:"aggregationDetails"`
	EncodingDetails    ConfigurationEncodingDetails    `json:"encodingDetails"`
	DeliveryDetails    ConfigurationDeliveryDetails    `json:"deliveryDetails"`
}

// Generic Configuration type
type ConfigurationMember struct {
	Type string `json:"type,omitempty"`
	ID   string `json:"id,omitempty"`
}

type ConfigurationContactDetails struct {
	Contact       ConfigurationMember `json:"contact"`
	MailAddresses []string            `json:"mailAddresses"`
}

type ConfigurationLogFormatDetails struct {
	LogFormat     ConfigurationMember `json:"logFormat"`
	LogIdentifier string              `json:"logIdentifier"`
}

type ConfigurationAggregationDetails struct {
	Type              string              `json:"type"`
	DeliveryFrequency ConfigurationMember `json:"deliveryFrequency"`
}

type ConfigurationEncodingDetails struct {
	Encoding ConfigurationMember `json:"encoding"`
}

type ConfigurationDeliveryDetails struct {
	Type         string `json:"type"`
	EmailAddress string `json:"emailAddress,omitempty"`
	DomainPrefix string `json:"domainPrefix,omitempty"`
	CpcodeID     int    `json:"cpcodeId,omitempty"`
	Directory    string `json:"directory,omitempty"`
}

type ConfigurationCopyOptions struct {
	CopyTarget ConfigurationCopyTarget `json:"copyTarget"`
}

type ConfigurationCopyTarget struct {
	LogSource ConfigurationMember `json:"logSource"`
}
