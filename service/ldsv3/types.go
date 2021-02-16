package ldsv3

// OutputSources identifies a collection of SourcesRespElement objects
type OutputSources []OutputSourcesElement

// OutputSourcesElement identifies an object for which you can create a log delivery configuration. Read-Only.
type OutputSourcesElement struct {
	// Unique identifier for each log source.
	ID string `json:"id"`
	// Identifies the type of log source.
	// Possible types are answerx, cpcode-products, edns, gtm
	Type string `json:"type"`
	// Used by all. Number of days that Akamai should keep logs collected for that log source.
	// You can create log redelivery requests only for this retention period.
	LogRetentionDays int `json:"logRetentionDays"`
	// Used by all. Technical links for actions with source type
	Links []OutputLinks `json:"links"`
	// Used by answerx. The AnswerX object name.
	Name string `json:"name,omitempty"`
	// Used by cpcode-products. CP code ID and name for which logs will be collected.
	CpCode string `json:"cpCode,omitempty"`
	// Used by cpcode-products. Products for which logs will be collected.
	Products []string `json:"products,omitempty"`
	// Used by edns. The full name of the EDNS zone.
	ZoneName string `json:"zoneName,omitempty"`
	// Used by gtm. The full name of the Traffic Management domain property.
	PropertyName string `json:"propertyName,omitempty"`
}

// OutputLinks represents structure widely used in response for links. Read-Only.
type OutputLinks struct {
	Rel    string `json:"rel"`
	Href   string `json:"href"`
	Method string `json:"method,omitempty"`
	Title  string `json:"title,omitempty"`
}

// OutputConfigurations identifies a collection of SourcesRespElement objects
type OutputConfigurations []OutputConfigurationElement

// OutputConfigurationElement represents output for the main log delivery configuration object.
type OutputConfigurationElement struct {
	// Read-only. Unique identifier of this configuration.
	ID int `json:"id"`
	// Read-only. Log configuration status, either active, expired, or suspended.
	// Only active configurations are used in the actual log delivery process.
	Status string `json:"status"`
	// Start date from which logs will be collected.
	StartDate string `json:"startDate"`
	// End date to which logs will be collected.
	EndDate string `json:"endDate,omitempty"`
	// Read-only. This member appears in log configurations only in server responses.
	// For creating and modifying log configuration, all required information to identify configuration is in the URL
	LogSource OutputSourcesElement `json:"logSource"`
	// Defines how to aggregate logs: by log arrival or by hit time.
	AggregationDetails ConfigurationBodyAggregationDetails `json:"aggregationDetails"`
	// Contains details about contact person for this log delivery configuration.
	ContactDetails ConfigurationBodyContactDetails `json:"contactDetails"`
	// Either an email or ftp or netstorage object.
	DeliveryDetails ConfigurationBodyDeliveryDetails `json:"deliveryDetails"`
	// Describes the log encoding.
	EncodingDetails ConfigurationBodyEncodingDetails `json:"encodingDetails"`
	// Describes the log format.
	LogFormatDetails ConfigurationBodyLogFormatDetails `json:"logFormatDetails"`
	// Packed log message’s approximate size
	MessageSize GenericConfigurationParameterElement `json:"messageSize"`
	// Read-Only. Technical links for actions with object
	Links []OutputLinks `json:"links"`
}

// OutputLogRedelivery identifies a collection of LogRedeliveryElement objects
type OutputLogRedelivery []OutputLogRedeliveryElement

// OutputLogRedeliveryElement collects information needed to create a log redelivery request.
type OutputLogRedeliveryElement struct {
	// Log delivery configuration for which this redelivery request was created.
	LogConfiguration GenericConfigurationParameterElement `json:"logConfiguration"`
	// Unique ID of this redelivery request.
	ID string `json:"id"`
	// First hour of time range (0–23) for which log redelivery is requested.
	BeginTime int `json:"beginTime"`
	// Last hour of time range (1–24) for which log redelivery is requested.
	EndTime int `json:"endTime"`
	// Date from which log redelivery is requested.
	RedeliveryDate string `json:"redeliveryDate"`
	// Status of the redelivery, for example new, scheduled, success, or failed.
	Status string `json:"status"`
	// Date the request for redelivery was created.
	CreatedDate string `json:"createdDate"`
	// Date of the last time the redelivery request was modified.
	ModifiedDate string `json:"modifiedDate"`
	// Read-Only. Technical links for actions with object
	Links []OutputLinks `json:"links"`
}

// ConfigurationParameterResponse identifies a collection of GenericConfigurationParameterElement objects
type ConfigurationParameterResponse []GenericConfigurationParameterElement

// GenericConfigurationParameterElement a simple data type used for pairs of IDs and values.
type GenericConfigurationParameterElement struct {
	// Unique identifier for the object
	ID string `json:"id"`
	// Read-Only. Human-readable value for the object.
	Value string `json:"value,omitempty"`
	// Read-Only. Technical links for actions with object
	Links *[]OutputLinks `json:"links,omitempty"`
}

// ConfigurationBody represents the main log delivery configuration object
// that creates and updates a log delivery configuration.
type ConfigurationBody struct {
	// Start date from which logs will be collected.
	// Start date has to be set at least one day after current date
	StartDate string `json:"startDate"`
	// (Optional) End date to which logs will be collected.
	EndDate string `json:"endDate,omitempty"`
	// (Optional) Used in Update call only.
	// Describes detailed log source information for configuration
	// Both type and ID of log source are required
	LogSource *GenericConfigurationParameterElement `json:"logSource,omitempty"`
	// Contains details about contact person for this log delivery configuration.
	ContactDetails ConfigurationBodyContactDetails `json:"contactDetails"`
	// Describes the log format.
	LogFormatDetails ConfigurationBodyLogFormatDetails `json:"logFormatDetails"`
	// Packed log message’s approximate size
	MessageSize GenericConfigurationParameterElement `json:"messageSize"`
	// Defines how to aggregate logs: by log arrival or by hit time.
	AggregationDetails ConfigurationBodyAggregationDetails `json:"aggregationDetails"`
	// Describes the log encoding.
	EncodingDetails ConfigurationBodyEncodingDetails `json:"encodingDetails"`
	// Either an email or ftp or netstorage object.
	DeliveryDetails ConfigurationBodyDeliveryDetails `json:"deliveryDetails"`
}

// LogSourceBodyMember represents structure for LogSource input
type LogSourceBodyMember struct {
	// Unique identifier for the object
	ID string `json:"id"`
	// Type is required only if used in ConfigurationCopyBodyTarget
	Type string `json:"type,omitempty"`
}

// ConfigurationBodyContactDetails contains details about contact person for this log delivery configuration.
type ConfigurationBodyContactDetails struct {
	// Contact information provided as a simple GenericConfigurationParameterElement object
	Contact GenericConfigurationParameterElement `json:"contact"`
	// List of email addresses for contacts for this log format configuration.
	MailAddresses []string `json:"mailAddresses"`
}

// ConfigurationBodyLogFormatDetails describes the log format.
type ConfigurationBodyLogFormatDetails struct {
	// Selected format for log delivery, a simple GenericConfigurationParameterElement object.
	LogFormat GenericConfigurationParameterElement `json:"logFormat"`
	// Represents the first token of the log filename.
	LogIdentifier string `json:"logIdentifier"`
}

// ConfigurationBodyAggregationDetails defines how to aggregate logs: by log arrival or by hit time.
type ConfigurationBodyAggregationDetails struct {
	// Identifies the type of aggregation. Possible type are byLogArrival or byHitTime
	Type string `json:"type"`
	// Used with byLogArrival type
	// Period of time that will be covered by log delivery, provided as a simple GenericConfigurationParameterElement object.
	DeliveryFrequency *GenericConfigurationParameterElement `json:"deliveryFrequency,omitempty"`
	// Used with byHitTime type
	// Indicates whether residual data should be sent at regular intervals after each day.
	DeliverResidualData bool `json:"deliverResidualData,omitempty"`
	// Used with byHitTime type
	// Data completion threshold, or the percentage of expected logs to be processed
	// before the log data is sent to you, provided as a simple GenericConfigurationParameterElement object.
	DeliveryThreshold *GenericConfigurationParameterElement `json:"deliveryThreshold,omitempty"`
}

// ConfigurationBodyEncodingDetails describes the log encoding.
type ConfigurationBodyEncodingDetails struct {
	// Selected encoding option used to encode logs, a simple GenericConfigurationParameterElement object.
	Encoding GenericConfigurationParameterElement `json:"encoding"`
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
	LogSource LogSourceBodyMember `json:"logSource"`
}

// RedeliveryBody collects information needed to create a log redelivery request.
type RedeliveryBody struct {
	// Log delivery configuration for which this redelivery request was created.
	// ID in LogSourceBodyMember should be unique ID for the log delivery configuration.
	LogConfiguration LogSourceBodyMember `json:"logConfiguration"`
	// First hour of time range (0–23) for which log redelivery is requested.
	BeginTime int `json:"beginTime"`
	// Last hour of time range (1–24) for which log redelivery is requested.
	EndTime int `json:"endTime"`
	// Date from which log redelivery is requested.
	RedeliveryDate string `json:"redeliveryDate"`
}
