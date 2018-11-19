package edgegrid

import (
	"fmt"
	"reflect"
	"time"

	log "github.com/sirupsen/logrus"
)

type DiagToolsService struct {
	client *Client
}

type AkamaiDTUserLinkReq struct {
	EndUserName string `json:"endUserName"`
	URL         string `json:"url"`
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
		Name      string      `json:"name"`
		Email     interface{} `json:"email"`
		Timestamp time.Time   `json:"timestamp"`
		URL       string      `json:"url"`
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

// LaunchErrorTranslationRequest async request creation for Error Translation
func (nls *DiagToolsService) LaunchErrorTranslationRequest(errorCode string) (*AkamaiDTErrorTranslationResp, *ClientResponse, error) {
	apiURI := fmt.Sprintf("%s/errors/%s/translate-error", DTPathV2, errorCode)

	var k *AkamaiDTErrorTranslationResp
	resp, err := nls.client.NewRequest("POST", apiURI, nil, &k)

	log.Debug(fmt.Sprintf("[%s]::Rate limit for Error Translation requests: %s per 60 seconds", reflect.TypeOf(nls), resp.Response.Header["X-Ratelimit-Limit"]))
	log.Debug(fmt.Sprintf("[%s]::Remaining allowed number of requests: %s", reflect.TypeOf(nls), resp.Response.Header["X-Ratelimit-Remaining"]))

	return k, resp, err
}

// CheckAnErrorTranslationRequest makes polling requests for status of request
// Looks like not working properly
func (nls *DiagToolsService) CheckAnErrorTranslationRequest(requestID string) (*AkamaiDTErrorTranslationResp, *ClientResponse, error) {
	apiURI := fmt.Sprintf("%s/translate-error-requests/%s", DTPathV2, requestID)

	var k *AkamaiDTErrorTranslationResp
	resp, err := nls.client.NewRequest("GET", apiURI, nil, &k)

	if resp.Response.StatusCode == 303 {
		return nil, resp, err
	}
	return k, resp, err
}

// TranslateAnError gets translated error message
func (nls *DiagToolsService) TranslateAnError(requestID string) (*AkamaiDTTranslatedErrorResp, *ClientResponse, error) {
	apiURI := fmt.Sprintf("%s/translate-error-requests/%s/translated-error", DTPathV2, requestID)

	var k *AkamaiDTTranslatedErrorResp
	resp, err := nls.client.NewRequest("GET", apiURI, nil, &k)

	return k, resp, err
}

// CDNStatus checks if given IP belongs to Akamai CDN
// TODO: migrate to async if required
func (nls *DiagToolsService) CDNStatus(ip string) (*AkamaiDTCDNStatusResp, *ClientResponse, error) {
	apiURI := fmt.Sprintf("%s/ip-addresses/%s/is-cdn-ip", DTPathV2, ip)

	var k *AkamaiDTCDNStatusResp
	resp, err := nls.client.NewRequest("GET", apiURI, nil, &k)

	log.Debug(fmt.Sprintf("[%s]::Rate limit for Error Translation requests: %s per 60 seconds", reflect.TypeOf(nls), resp.Response.Header["X-Ratelimit-Limit"]))
	log.Debug(fmt.Sprintf("[%s]::Remaining allowed number of requests: %s", reflect.TypeOf(nls), resp.Response.Header["X-Ratelimit-Remaining"]))

	return k, resp, err
}

// GenerateDiagnosticLink generates user link and request
func (nls *DiagToolsService) GenerateDiagnosticLink(username, testURL string) (*AkamaiDTGenerateDiagLinkResp, *ClientResponse, error) {
	apiURI := fmt.Sprintf("%s/end-users/diagnostic-url", DTPathV2)

	var k *AkamaiDTGenerateDiagLinkResp

	body := AkamaiDTUserLinkReq{
		EndUserName: username,
		URL:         testURL,
	}

	resp, err := nls.client.NewRequest("POST", apiURI, body, &k)

	return k, resp, err
}

// ListDiagnosticLinkRequests lists all requests
func (nls *DiagToolsService) ListDiagnosticLinkRequests() (*AkamaiDTListDiagLinkRequestsResp, *ClientResponse, error) {
	apiURI := fmt.Sprintf("%s/end-users/ip-requests", DTPathV2)

	var k *AkamaiDTListDiagLinkRequestsResp
	resp, err := nls.client.NewRequest("GET", apiURI, nil, &k)

	return k, resp, err
}

// GetDiagnosticLinkRequest gets request details
func (nls *DiagToolsService) GetDiagnosticLinkRequest(id string) (*AkamaiDTDiagLinkRequestResp, *ClientResponse, error) {
	apiURI := fmt.Sprintf("%s/end-users/ip-requests/%s/ip-details", DTPathV2, id)

	var k *AkamaiDTDiagLinkRequestResp
	resp, err := nls.client.NewRequest("GET", apiURI, nil, &k)

	return k, resp, err
}
