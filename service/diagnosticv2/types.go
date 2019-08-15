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
	TranslatedError TranslatedErrorClass `json:"translatedError"`
}

type TranslatedErrorClass struct {
	URL              string               `json:"url"`
	HTTPResponseCode int64                `json:"httpResponseCode"`
	Timestamp        string               `json:"timestamp"`
	EpochTime        int64                `json:"epochTime"`
	ClientIP         string               `json:"clientIp"`
	ConnectingIP     string               `json:"connectingIp"`
	ServerIP         string               `json:"serverIp"`
	OriginHostname   string               `json:"originHostname"`
	OriginIP         string               `json:"originIp"`
	UserAgent        string               `json:"userAgent"`
	RequestMethod    string               `json:"requestMethod"`
	ReasonForFailure string               `json:"reasonForFailure"`
	WafDetails       string               `json:"wafDetails"`
	Logs             []TranslatedErrorLog `json:"logs"`
}

type TranslatedErrorLog struct {
	Description string            `json:"description"`
	Fields      map[string]string `json:"fields"`
}

//VerifyIP represents information if given IP address belongs to Akamai platform
type VerifyIP struct {
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

/*
type DTGTMPropertiesResp struct {
	GtmProperties []struct {
		Property string `json:"property"`
		Domain   string `json:"domain"`
		HostName string `json:"hostName"`
	} `json:"gtmProperties"`
}

type DTGTMPropertyIpsResp struct {
	GtmPropertyIps struct {
		Property  string   `json:"property"`
		Domain    string   `json:"domain"`
		TestIps   []string `json:"testIps"`
		TargetIps []string `json:"targetIps"`
	} `json:"gtmPropertyIps"`
}

type DTDigResp struct {
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

type DTMtrResp struct {
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

type DTGeolocation struct {
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

type DTCurlResp struct {
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

type DTCurlReq struct {
	URL       string `json:"url"`
	UserAgent string `json:"userAgent"`
}

type DTGenerateDiagLinkResp struct {
	DiagnosticURL string `json:"diagnosticUrl"`
}

type DTListDiagLinkRequestsResp struct {
	EndUserIPRequests []struct {
		EndUserName string    `json:"name"`
		RequestID   uint32    `json:"requestId"`
		URL         string    `json:"url"`
		Timestamp   time.Time `json:"timestamp"`
	} `json:"endUserIpRequests"`
}







type DTTranslatedErrorResp struct {
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
				GhostIP                         string `json:"Ghost IP"`
				ForwardRequest                  string `json:"Forward Request"`
				Timestamp                       string `json:"timestamp"`
				ContentBytesReceived            string `json:"content bytes received"`
				TotalEstimatedBytesReceived     string `json:"total estimated bytes received"`
				ForwardIP                       string `json:"Forward IP"`
				ClientIPPPrefresh               string `json:"client IP (p-prefresh)"`
				HTTPMethodGETHEADEtc            string `json:"HTTP method (GET HEAD etc)"`
				ARL                             string `json:"ARL"`
				HTTPStatusCode                  string `json:"HTTP status code"`
				ContentType                     string `json:"content-type"`
				IMSIIms                         string `json:"IMS (i-ims)"`
				SSL                             string `json:"SSL"`
				RequestNumber                   string `json:"Request Number"`
				Edgescape                       string `json:"Edgescape"`
				ForwardHostname                 string `json:"Forward Hostname"`
				GhostRequestHeaderSize          string `json:"Ghost request header size"`
				GhostRequestSize                string `json:"Ghost request size"`
				SSLOverheadBytes                string `json:"SSL overhead bytes"`
				ForwardARLIfRewrittenInMetadata string `json:"Forward ARL (if rewritten in metadata)"`
				RequestID                       string `json:"Request id"`
				ReceivedB                       string `json:"received_b"`
				ObjectMaxAgeS                   string `json:"object-max-age_s"`
				Sureroute2Info                  string `json:"Sureroute2info"`
				Range                           string `json:"range"`
				SureRouteRaceStatIndirRoute     string `json:"SureRouteRaceStat-indirRoute"`
				SureRouteRaceStatDirRoute       string `json:"SureRouteRace-stat-dirRoute"`
				ForwardSideHTTPOverhead         string `json:"Forward-side-http-overhead"`
				ReasonForThrottling             string `json:"Reason for Throttling"`
				TimeSpentDeferringForwardRead   string `json:"Time spent deferring forward read"`
				ObjectStatus2                   string `json:"Object Status 2"`
				MultiFeatureStatusField         string `json:"Multi-Feature Status Field"`
				MultiPurposeKeyValueField       string `json:"Multi-Purpose Key/Value Field"`
				RealIPOfForwardGhostESSL        string `json:"Real IP of Forward Ghost (ESSL)"`
			} `json:"fields"`
		} `json:"logs"`
	} `json:"translatedError"`
}

*/
