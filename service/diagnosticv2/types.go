package diagnosticv2

import "time"

type GhostLocations struct {
	Locations []struct {
		ID    string `json:"id"`
		Value string `json:"value"`
	} `json:"locations"`
}

//TranslateErrorAsync is returned during async call to retrieve error from Akamai platform.
//It contains information needed to further retrieve error
type TranslateErrorAsync struct {
	RequestID  string `json:"requestId"`
	Link       string `json:"link"`
	RetryAfter int    `json:"retryAfter"`
}

type TranslatedError struct {
	TranslatedError struct {
		URL              string `json:"url"`
		HTTPResponseCode int    `json:"httpResponseCode"`
		Timestamp        string `json:"timestamp"`
		EpochTime        int    `json:"epochTime"`
		ClientIP         string `json:"clientIp"`
		ConnectingIP     string `json:"connectingIp"`
		ServerIP         string `json:"serverIp"`
		OriginHostname   string `json:"originHostname"`
		OriginIP         string `json:"originIp"`
		UserAgent        string `json:"userAgent"`
		RequestMethod    string `json:"requestMethod"`
		ReasonForFailure string `json:"reasonForFailure"`
		WafDetails       string `json:"wafDetails"`
		Logs             []struct {
			Description string `json:"description"`
			Fields      struct {
				EdgeServerIP                      string `json:"Edge server IP"`
				ClientRequestR                    string `json:"client request (r)"`
				DateTime                          string `json:"Date & Time"`
				EpochTime                         string `json:"Epoch Time"`
				ObjectSize                        string `json:"object size"`
				ContentBytesServed                string `json:"content bytes served"`
				TotalEstimatedBytesServed         string `json:"total estimated bytes served"`
				ClientIP                          string `json:"client IP"`
				HTTPMethod                        string `json:"HTTP method"`
				ARL                               string `json:"ARL"`
				HTTPStatusCode                    string `json:"HTTP status code"`
				Error                             string `json:"error"`
				ContentType                       string `json:"content-type"`
				HostHeader                        string `json:"host header"`
				Cookie                            string `json:"cookie"`
				Referrer                          string `json:"referrer"`
				UserAgent                         string `json:"user-agent"`
				IMS                               string `json:"IMS"`
				SSL                               string `json:"SSL"`
				PersistentRequestNumber           string `json:"persistent request number"`
				ClientRequestHeaderSize           string `json:"Client request header size"`
				AcceptLanguage                    string `json:"Accept-Language"`
				SSLOverheadBytes                  string `json:"SSL overhead bytes"`
				SerialNumberAndMap                string `json:"Serial number and map"`
				RequestByteRange                  string `json:"Request byte-range"`
				UncompressedLength                string `json:"Uncompressed length"`
				OtherErrorIndication              string `json:"Other-Error-Indication"`
				DcaData                           string `json:"dca-data"`
				XForwardedFor                     string `json:"X-Forwarded-For"`
				XAkamaiEdgeLog                    string `json:"X-Akamai-Edge-Log"`
				ObjectMaxAgeS                     string `json:"object-max-age_s"`
				CustomField                       string `json:"custom-field"`
				ObjectStatus2                     string `json:"object-status-2"`
				SslByte                           string `json:"ssl-byte"`
				CHTTPOverhead                     string `json:"c-http-overhead"`
				ClientRateLimiting                string `json:"Client-rate-limiting"`
				ClientRequestBodySize             string `json:"Client-request-body-size"`
				FlvSeekProcessingInfo             string `json:"flv seek processing info"`
				TrueClientIP                      string `json:"True client ip"`
				WebApplicationFirewallInformation string `json:"Web Application Firewall Information"`
				EdgeTokenizationInformation       string `json:"Edge Tokenization Information"`
				OriginFileSize                    string `json:"Origin File Size"`
				HTTPStreamingInfo                 string `json:"HTTP Streaming info"`
				ReasonForNotCachingPrivReleased   string `json:"Reason for not caching (priv/released)"`
				RateAccountingInfo                string `json:"Rate Accounting info"`
				RequestBodyInspection             string `json:"Request body inspection"`
				ResponseBodyInspection            string `json:"Response body inspection"`
			} `json:"fields"`
		} `json:"logs"`
	} `json:"translatedError"`
}

//CDNStatus represents information if given IP address belongs to Akamai platform
type CDNStatus struct {
	IsAkamai bool `json:"isCdnIp"`
}

type DiagnosticLinkURL struct {
	URL string `json:"diagnosticUrl"`
}

type DiagnosticLinkRequest struct {
	EndUserName string `json:"endUserName"`
	URL         string `json:"url"`
}

type DiagnosticLinkRequests struct {
	EndUserIPRequests []struct {
		EndUserName string    `json:"name"`
		RequestID   uint32    `json:"requestId"`
		URL         string    `json:"url"`
		Timestamp   time.Time `json:"timestamp"`
	} `json:"endUserIpRequests"`
}

type DiagnosticLinkResult struct {
	EndUserIPDetails struct {
		Name      string    `json:"name"`
		Email     string    `json:"email"`
		Timestamp time.Time `json:"timestamp"`
		URL       string    `json:"url"`
		Ips       []struct {
			Description string `json:"description"`
			Location    string `json:"location"`
			IP          string `json:"ip"`
			IPType      string `json:"ipType"`
		} `json:"ips"`
		Browser string `json:"browser"`
	} `json:"endUserIpDetails"`
}

type Geolocation struct {
	GeoLocation struct {
		ClientIP    string  `json:"clientIp"`
		CountryCode string  `json:"countryCode"`
		RegionCode  string  `json:"regionCode"`
		City        string  `json:"city"`
		Dma         int     `json:"dma"`
		Msa         int     `json:"msa"`
		Pmsa        int     `json:"pmsa"`
		AreaCode    string  `json:"areaCode"`
		Latitude    float64 `json:"latitude"`
		Longitude   float64 `json:"longitude"`
		County      string  `json:"county"`
		Continent   string  `json:"continent"`
		Fips        string  `json:"fips"`
		TimeZone    string  `json:"timeZone"`
		Network     string  `json:"network"`
		NetworkType string  `json:"networkType"`
		ZipCode     string  `json:"zipCode"`
		Throughput  string  `json:"throughput"`
		AsNum       string  `json:"asNum"`
		Proxy       string  `json:"proxy"`
	} `json:"geoLocation"`
}

type DigResult struct {
	DigInfo struct {
		Hostname      string `json:"hostname"`
		QueryType     string `json:"queryType"`
		AnswerSection []struct {
			Domain           string      `json:"domain"`
			TTL              int         `json:"ttl"`
			RecordClass      string      `json:"recordClass"`
			RecordType       string      `json:"recordType"`
			PreferenceValues interface{} `json:"preferenceValues"`
			Value            string      `json:"value"`
		} `json:"answerSection"`
		AuthoritySection []struct {
			Domain           string      `json:"domain"`
			TTL              int         `json:"ttl"`
			RecordClass      string      `json:"recordClass"`
			RecordType       string      `json:"recordType"`
			PreferenceValues interface{} `json:"preferenceValues"`
			Value            string      `json:"value"`
		} `json:"authoritySection"`
		Result string `json:"result"`
	} `json:"digInfo"`
}

type MtrResult struct {
	Mtr struct {
		Source      string    `json:"source"`
		Destination string    `json:"destination"`
		StartTime   time.Time `json:"startTime"`
		Host        string    `json:"host"`
		PacketLoss  float64   `json:"packetLoss"`
		AvgLatency  float64   `json:"avgLatency"`
		Analysis    string    `json:"analysis"`
		Hops        []struct {
			Number int     `json:"number"`
			Host   string  `json:"host"`
			Loss   float64 `json:"loss"`
			Sent   int     `json:"sent"`
			Last   float64 `json:"last"`
			Avg    float64 `json:"avg"`
			Best   float64 `json:"best"`
			Worst  float64 `json:"worst"`
			StDev  float64 `json:"stDev"`
		} `json:"hops"`
		Result string `json:"result"`
	} `json:"mtr"`
}

type CurlResult struct {
	CurlResults struct {
		HTTPStatusCode  int `json:"httpStatusCode"`
		ResponseHeaders struct {
			Server        string `json:"Server"`
			Connection    string `json:"Connection"`
			Expires       string `json:"Expires"`
			MimeVersion   string `json:"Mime-Version"`
			ContentLength string `json:"Content-Length"`
			Date          string `json:"Date"`
			ContentType   string `json:"Content-Type"`
		} `json:"responseHeaders"`
		ResponseBody string `json:"responseBody"`
	} `json:"curlResults"`
}

type CurlRequest struct {
	URL       string `json:"url"`
	UserAgent string `json:"userAgent"`
}

type GTMPropertiesResult struct {
	GtmProperties []struct {
		Property string `json:"property"`
		Domain   string `json:"domain"`
		HostName string `json:"hostName"`
	} `json:"gtmProperties"`
}

type GTMPropertyIpsResult struct {
	GtmPropertyIps struct {
		Property  string   `json:"property"`
		Domain    string   `json:"domain"`
		TestIps   []string `json:"testIps"`
		TargetIps []string `json:"targetIps"`
	} `json:"gtmPropertyIps"`
}
