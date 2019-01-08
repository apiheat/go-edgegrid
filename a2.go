package edgegrid

import (
	"fmt"
	"net/http"
)

type AdaptiveAccelerationService struct {
	client *Client
}

//QStrAdaptiveAcceleration includes query params used across AdaptiveAccelerationService
type QStrAdaptiveAcceleration struct{}

//ReportProperty reports property ID
func (nls *AdaptiveAccelerationService) ReportProperty(id string) (*ClientResponse, error) {
	qParams := QStrAdaptiveAcceleration{}
	path := fmt.Sprintf("%s/%s", A2PathV1, id)

	resp, err := nls.client.makeAPIRequest(http.MethodGet, path, qParams, nil, nil, nil)

	return resp, err
}

//ResetProperty Resets property AdaptiveAcceleration based on given ID
func (nls *AdaptiveAccelerationService) ResetProperty(id string) (*ClientResponse, error) {
	qParams := QStrAdaptiveAcceleration{}
	path := fmt.Sprintf("%s/%s", A2PathV1, id)

	resp, err := nls.client.makeAPIRequest(http.MethodPost, path, qParams, nil, nil, nil)

	return resp, err
}
