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

// ConfigurationBody represents the main log delivery configuration object
// that creates and updates a log delivery configuration.
type ConfigurationBody struct {
	// Start date from which logs will be collected.
	StartDate string `json:"startDate"`
	// (Optional) End date to which logs will be collected.
	EndDate string `json:"endDate,omitempty"`
	//LogSource GenericBodyMember `json:"logSource"`
	// Contains details about contact person for this log delivery configuration.
	ContactDetails ConfigurationBodyContactDetails `json:"contactDetails"`
	// Describes the log format.
	LogFormatDetails ConfigurationBodyLogFormatDetails `json:"logFormatDetails"`
	// Packed log message’s approximate size
	MessageSize GenericBodyMember `json:"messageSize"`
	// Defines how to aggregate logs: by log arrival or by hit time.
	AggregationDetails ConfigurationBodyAggregationDetails `json:"aggregationDetails"`
	// Describes the log encoding.
	EncodingDetails ConfigurationBodyEncodingDetails `json:"encodingDetails"`
	// Either an email or ftp or netstorage object.
	DeliveryDetails ConfigurationBodyDeliveryDetails `json:"deliveryDetails"`
}

// GenericBodyMember represents structure widely used in request body
type GenericBodyMember struct {
	// Unique identifier for the object
	ID string `json:"id"`
	// Type is required only if used in ConfigurationCopyBodyTarget
	Type string `json:"type,omitempty"`
}

// ConfigurationBodyContactDetails contains details about contact person for this log delivery configuration.
type ConfigurationBodyContactDetails struct {
	// Contact information provided as a simple GenericBodyMember object
	Contact GenericBodyMember `json:"contact"`
	// List of email addresses for contacts for this log format configuration.
	MailAddresses []string `json:"mailAddresses"`
}

// ConfigurationBodyLogFormatDetails describes the log format.
type ConfigurationBodyLogFormatDetails struct {
	// Selected format for log delivery, a simple GenericBodyMember object.
	LogFormat GenericBodyMember `json:"logFormat"`
	// Represents the first token of the log filename.
	LogIdentifier string `json:"logIdentifier"`
}

// ConfigurationBodyAggregationDetails defines how to aggregate logs: by log arrival or by hit time.
type ConfigurationBodyAggregationDetails struct {
	// Identifies the type of aggregation. Possible type are byLogArrival or byHitTime
	Type string `json:"type"`
	// Used with byLogArrival type
	// Period of time that will be covered by log delivery, provided as a simple GenericBodyMember object.
	DeliveryFrequency *GenericBodyMember `json:"deliveryFrequency,omitempty"`
	// Used with byHitTime type
	// Indicates whether residual data should be sent at regular intervals after each day.
	DeliverResidualData bool `json:"deliverResidualData,omitempty"`
	// Used with byHitTime type
	// Data completion threshold, or the percentage of expected logs to be processed
	// before the log data is sent to you, provided as a simple GenericBodyMember object.
	DeliveryThreshold *GenericBodyMember `json:"deliveryThreshold,omitempty"`
}

// ConfigurationBodyEncodingDetails describes the log encoding.
type ConfigurationBodyEncodingDetails struct {
	// Selected encoding option used to encode logs, a simple GenericBodyMember object.
	Encoding GenericBodyMember `json:"encoding"`
	// Public key value for encrypted encoding.
	// You need to set the public key value if GPG encrypted encoding is used.
	EncodingKey string `json:"encodingKey,omitempty"`
}

// ConfigurationBodyDeliveryDetails encapsulates log delivery sent by email, ftp or netstorage(httpsns4)
type ConfigurationBodyDeliveryDetails struct {
	// Identifies this type of delivery, Can be email, ftp, httpsns4
	Type string `json:"type"`
	// Used with email type. Email address to which log will be sent.
	EmailAddress string `json:"emailAddress,omitempty"`
	// Used with httpsns4. Prefix of the storage group, this is part of URLs.
	DomainPrefix string `json:"domainPrefix,omitempty"`
	// Used with httpsns4. Identifies CP code within this storage group.
	CpcodeID int `json:"cpcodeId,omitempty"`
	// Used with httpsns4 or ftp. Directory within CP code that logs will be uploaded to.
	Directory string `json:"directory,omitempty"`
	// Used with ftp. Login used to authenticate to FTP machine.
	Login string `json:"login,omitempty"`
	// Used with ftp. Machine to which log will be sent by FTP.
	Machine string `json:"machine,omitempty"`
	// Used with ftp. Password used to authenticate to FTP machine.
	// Keep in mind that this field will be empty in all responses from the server.
	Password string `json:"password,omitempty"`
}

// ConfigurationCopyBody encapsulates information needed to copy a log configuration
type ConfigurationCopyBody struct {
	// Represents target log source for configuration copy request.
	CopyTarget ConfigurationCopyBodyTarget `json:"copyTarget"`
}

// ConfigurationCopyBodyTarget represents target log source for configuration copy request.
type ConfigurationCopyBodyTarget struct {
	// Describes detailed log source information for configuration
	// Both type and ID of log source are required
	LogSource GenericBodyMember `json:"logSource"`
}

// RedeliveryBody collects information needed to create a log redelivery request.
type RedeliveryBody struct {
	// Log delivery configuration for which this redelivery request was created.
	// ID in GenericBodyMember should be unique ID for the log delivery configuration.
	LogConfiguration GenericBodyMember `json:"logConfiguration"`
	// First hour of time range (0–23) for which log redelivery is requested.
	BeginTime int `json:"beginTime"`
	// Last hour of time range (1–24) for which log redelivery is requested.
	EndTime int `json:"endTime"`
	// Date from which log redelivery is requested.
	RedeliveryDate string `json:"redeliveryDate"`
}
