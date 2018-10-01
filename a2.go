package edgegrid

import "fmt"

type AdaptiveAccelerationService struct {
	client *Client
}

func (nls *AdaptiveAccelerationService) ReportProperty(id string) (*ClientResponse, error) {

	apiURI := fmt.Sprintf("%s/%s", A2PathV1, id)

	resp, err := nls.client.NewRequest("GET", apiURI, nil, nil)

	return resp, err

}

func (nls *AdaptiveAccelerationService) ResetProperty(id string) (*ClientResponse, error) {

	apiURI := fmt.Sprintf("%s/%s", A2PathV1, id)

	resp, err := nls.client.NewRequest("POST", apiURI, nil, nil)

	return resp, err

}
