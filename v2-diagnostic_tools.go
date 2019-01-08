package edgegrid

import (
	"fmt"
	"net/http"
	"reflect"
	"time"

	log "github.com/sirupsen/logrus"
)

type DiagToolsService struct {
	client *Client
}

type AkamaiGhostLocationsResp struct {
	Locations []struct {
		ID    string `json:"id"`
		Value string `json:"value"`
	} `json:"locations"`
}

type AkamaiDTGTMPropertiesResp struct {
	GtmProperties []struct {
		Property string `json:"property"`
		Domain   string `json:"domain"`
		HostName string `json:"hostName"`
	} `json:"gtmProperties"`
}

type AkamaiDTGTMPropertyIpsResp struct {
	GtmPropertyIps struct {
		Property  string   `json:"property"`
		Domain    string   `json:"domain"`
		TestIps   []string `json:"testIps"`
		TargetIps []string `json:"targetIps"`
	} `json:"gtmPropertyIps"`
}

type AkamaiDTDigResp struct {
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

type AkamaiDTMtrResp struct {
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

type AkamaiDTGeolocation struct {
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

type AkamaiDTCurlResp struct {
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

type AkamaiDTUserLinkReq struct {
	EndUserName string `json:"endUserName"`
	URL         string `json:"url"`
}

type AkamaiDTCurlReq struct {
	URL       string `json:"url"`
	UserAgent string `json:"userAgent"`
}

type AkamaiDTGenerateDiagLinkResp struct {
	DiagnosticURL string `json:"diagnosticUrl"`
}

type AkamaiDTListDiagLinkRequestsResp struct {
	EndUserIPRequests []struct {
		EndUserName string    `json:"name"`
		RequestID   uint32    `json:"requestId"`
		URL         string    `json:"url"`
		Timestamp   time.Time `json:"timestamp"`
	} `json:"endUserIpRequests"`
}

type AkamaiDTDiagLinkRequestResp struct {
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

type AkamaiDTCDNStatusResp struct {
	IsCdnIP bool `json:"isCdnIp"`
}

type AkamaiDTErrorTranslationResp struct {
	RequestID  string `json:"requestId"`
	Link       string `json:"link"`
	RetryAfter int    `json:"retryAfter"`
}

type AkamaiDTTranslatedErrorResp struct {
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

// QStrDiagTools includes query params used for diagnostic tools
type QStrDiagTools struct {
	HostName          string `url:"hostName,omitempty"`
	QueryType         string `url:"queryType,omitempty"`
	ResolveDNS        bool   `url:"resolveDns,omitempty"`
	DestinationDomain string `url:"destinationDomain,omitempty"`
}

// LaunchErrorTranslationRequest async request creation for Error Translation
func (nls *DiagToolsService) LaunchErrorTranslationRequest(errorCode string) (*AkamaiDTErrorTranslationResp, *ClientResponse, error) {

	qParams := QStrDiagTools{}
	path := fmt.Sprintf("%s/errors/%s/translate-error", DTPathV2, errorCode)

	var respStruct *AkamaiDTErrorTranslationResp
	resp, err := nls.client.makeAPIRequest(http.MethodPost, path, qParams, &respStruct, nil, nil)

	log.Debug(fmt.Sprintf("[%s]::Rate limit for Error Translation requests: %s per 60 seconds", reflect.TypeOf(nls), resp.Response.Header["X-Ratelimit-Limit"]))
	log.Debug(fmt.Sprintf("[%s]::Remaining allowed number of requests: %s", reflect.TypeOf(nls), resp.Response.Header["X-Ratelimit-Remaining"]))

	return respStruct, resp, err
}

// CheckAnErrorTranslationRequest makes polling requests for status of request
// Looks like not working properly
func (nls *DiagToolsService) CheckAnErrorTranslationRequest(requestID string) (*AkamaiDTErrorTranslationResp, *ClientResponse, error) {
	qParams := QStrDiagTools{}
	path := fmt.Sprintf("%s/translate-error-requests/%s", DTPathV2, requestID)

	var respStruct *AkamaiDTErrorTranslationResp
	resp, err := nls.client.makeAPIRequest(http.MethodGet, path, qParams, &respStruct, nil, nil)

	if resp.Response.StatusCode == http.StatusSeeOther {
		return nil, resp, err
	}
	return respStruct, resp, err
}

// TranslateAnError gets translated error message
func (nls *DiagToolsService) TranslateAnError(requestID string) (*AkamaiDTTranslatedErrorResp, *ClientResponse, error) {
	qParams := QStrDiagTools{}
	path := fmt.Sprintf("%s/translate-error-requests/%s/translated-error", DTPathV2, requestID)

	var respStruct *AkamaiDTTranslatedErrorResp
	resp, err := nls.client.makeAPIRequest(http.MethodGet, path, qParams, &respStruct, nil, nil)

	return respStruct, resp, err
}

// CDNStatus checks if given IP belongs to Akamai CDN
// TODO: migrate to async if required
func (nls *DiagToolsService) CDNStatus(ip string) (*AkamaiDTCDNStatusResp, *ClientResponse, error) {
	qParams := QStrDiagTools{}
	path := fmt.Sprintf("%s/ip-addresses/%s/is-cdn-ip", DTPathV2, ip)

	var respStruct *AkamaiDTCDNStatusResp
	resp, err := nls.client.makeAPIRequest(http.MethodGet, path, qParams, &respStruct, nil, nil)

	log.Debug(fmt.Sprintf("[%s]::Rate limit for CDN status requests: %s per 60 seconds", reflect.TypeOf(nls), resp.Response.Header["X-Ratelimit-Limit"]))
	log.Debug(fmt.Sprintf("[%s]::Remaining allowed number of requests: %s", reflect.TypeOf(nls), resp.Response.Header["X-Ratelimit-Remaining"]))

	return respStruct, resp, err
}

// GenerateDiagnosticLink generates user link and request
func (nls *DiagToolsService) GenerateDiagnosticLink(username, testURL string) (*AkamaiDTGenerateDiagLinkResp, *ClientResponse, error) {
	qParams := QStrDiagTools{}
	path := fmt.Sprintf("%s/end-users/diagnostic-url", DTPathV2)

	var respStruct *AkamaiDTGenerateDiagLinkResp

	requestStruct := AkamaiDTUserLinkReq{
		EndUserName: username,
		URL:         testURL,
	}

	resp, err := nls.client.makeAPIRequest(http.MethodPost, path, qParams, &respStruct, requestStruct, nil)

	return respStruct, resp, err
}

// ListDiagnosticLinkRequests lists all requests
func (nls *DiagToolsService) ListDiagnosticLinkRequests() (*AkamaiDTListDiagLinkRequestsResp, *ClientResponse, error) {
	qParams := QStrDiagTools{}
	path := fmt.Sprintf("%s/end-users/ip-requests", DTPathV2)

	var respStruct *AkamaiDTListDiagLinkRequestsResp
	resp, err := nls.client.makeAPIRequest(http.MethodGet, path, qParams, &respStruct, nil, nil)

	return respStruct, resp, err
}

// GetDiagnosticLinkRequest gets request details
func (nls *DiagToolsService) GetDiagnosticLinkRequest(id string) (*AkamaiDTDiagLinkRequestResp, *ClientResponse, error) {
	qParams := QStrDiagTools{}
	path := fmt.Sprintf("%s/end-users/ip-requests/%s/ip-details", DTPathV2, id)

	var respStruct *AkamaiDTDiagLinkRequestResp
	resp, err := nls.client.makeAPIRequest(http.MethodGet, path, qParams, &respStruct, nil, nil)

	return respStruct, resp, err
}

// IPGeolocation provides given IP geolocation details
func (nls *DiagToolsService) IPGeolocation(ip string) (*AkamaiDTGeolocation, *ClientResponse, error) {
	qParams := QStrDiagTools{}
	path := fmt.Sprintf("%s/ip-addresses/%s/geo-location", DTPathV2, ip)

	var respStruct *AkamaiDTGeolocation
	resp, err := nls.client.makeAPIRequest(http.MethodGet, path, qParams, &respStruct, nil, nil)

	log.Debug(fmt.Sprintf("[%s]::Rate limit for IP geolacation requests: %s per 60 seconds, but maximum 500 per day", reflect.TypeOf(nls), resp.Response.Header["X-Ratelimit-Limit"]))
	log.Debug(fmt.Sprintf("[%s]::Remaining allowed number of requests: %s", reflect.TypeOf(nls), resp.Response.Header["X-Ratelimit-Remaining"]))

	return respStruct, resp, err
}

// Dig provides dig functionality
func (nls *DiagToolsService) Dig(obj string, requestFrom AkamaiRequestFrom, hostname, query string) (*AkamaiDTDigResp, *ClientResponse, error) {

	if hostname == "" {
		return nil, nil, fmt.Errorf("'hostname' is required parameter: '%s'", hostname)
	}

	qParams := QStrDiagTools{
		HostName:  hostname,
		QueryType: query,
	}
	path := fmt.Sprintf("%s/%s/%s/dig-info", DTPathV2, requestFrom, obj)

	var respStruct *AkamaiDTDigResp
	resp, err := nls.client.makeAPIRequest(http.MethodGet, path, qParams, &respStruct, nil, nil)

	log.Debug(fmt.Sprintf("[%s]::Rate limit for request: %s per 60 seconds", reflect.TypeOf(nls), resp.Response.Header["X-Ratelimit-Limit"]))
	log.Debug(fmt.Sprintf("[%s]::Remaining allowed number of requests: %s", reflect.TypeOf(nls), resp.Response.Header["X-Ratelimit-Remaining"]))

	return respStruct, resp, err
}

// Mtr provides mtr functionality
func (nls *DiagToolsService) Mtr(obj string, requestFrom AkamaiRequestFrom, destinationDomain string, resolveDNS bool) (*AkamaiDTMtrResp, *ClientResponse, error) {

	if destinationDomain == "" {
		return nil, nil, fmt.Errorf("'destinationDomain' is required parameter: '%s'", destinationDomain)
	}

	qParams := QStrDiagTools{
		DestinationDomain: destinationDomain,
		ResolveDNS:        resolveDNS,
	}

	path := fmt.Sprintf("%s/%s/%s/mtr-data", DTPathV2, requestFrom, obj)

	var respStruct *AkamaiDTMtrResp
	resp, err := nls.client.makeAPIRequest(http.MethodGet, path, qParams, &respStruct, nil, nil)

	return respStruct, resp, err
}

// Curl provides curl functionality
func (nls *DiagToolsService) Curl(obj string, requestFrom AkamaiRequestFrom, testURL, userAgent string) (*AkamaiDTCurlResp, *ClientResponse, error) {

	if testURL == "" {
		return nil, nil, fmt.Errorf("'testURL' is required parameter: '%s'", testURL)
	}

	qParams := QStrDiagTools{}
	path := fmt.Sprintf("%s/%s/%s/curl-results", DTPathV2, requestFrom, obj)

	var respStruct *AkamaiDTCurlResp

	requestStruct := AkamaiDTCurlReq{
		UserAgent: userAgent,
		URL:       testURL,
	}

	resp, err := nls.client.makeAPIRequest(http.MethodPost, path, qParams, &respStruct, requestStruct, nil)

	return respStruct, resp, err
}

// ListGhostLocations provides Ghost locations
func (nls *DiagToolsService) ListGhostLocations() (*AkamaiGhostLocationsResp, *ClientResponse, error) {
	qParams := QStrDiagTools{}
	path := fmt.Sprintf("%s/ghost-locations/available", DTPathV2)

	var respStruct *AkamaiGhostLocationsResp
	resp, err := nls.client.makeAPIRequest(http.MethodGet, path, qParams, &respStruct, nil, nil)

	return respStruct, resp, err
}

// ListGTMProperties provides available GTM properties
func (nls *DiagToolsService) ListGTMProperties() (*AkamaiDTGTMPropertiesResp, *ClientResponse, error) {
	qParams := QStrDiagTools{}
	path := fmt.Sprintf("%s/gtm/gtm-properties", DTPathV2)

	var respStruct *AkamaiDTGTMPropertiesResp
	resp, err := nls.client.makeAPIRequest(http.MethodGet, path, qParams, &respStruct, nil, nil)

	return respStruct, resp, err
}

// ListGTMPropertyIPs provides available GTM properties
func (nls *DiagToolsService) ListGTMPropertyIPs(property, domain string) (*AkamaiDTGTMPropertyIpsResp, *ClientResponse, error) {

	if property == "" {
		return nil, nil, fmt.Errorf("'property' is required parameter: '%s'", property)
	}

	if domain == "" {
		return nil, nil, fmt.Errorf("'domain' is required parameter: '%s'", domain)
	}
	qParams := QStrDiagTools{}
	path := fmt.Sprintf("%s/gtm/%s/%s/gtm-property-ips", DTPathV2, property, domain)

	var respStruct *AkamaiDTGTMPropertyIpsResp
	resp, err := nls.client.makeAPIRequest(http.MethodGet, path, qParams, &respStruct, nil, nil)

	return respStruct, resp, err
}
