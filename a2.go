package edgegrid

import "fmt"

type AdaptiveAccelerationAPIService struct {
	client *Client
}

func (nls *AdaptiveAccelerationAPIService) ReportProperty(ID string) (*ClientResponse, error) {

	apiURI := fmt.Sprintf("%s/%s", apiPaths["a2_v1"], ID)

	resp, err := nls.client.NewRequest("GET", apiURI, nil, nil)

	return resp, err

}

func (nls *AdaptiveAccelerationAPIService) ResetProperty(ID string) (*ClientResponse, error) {

	apiURI := fmt.Sprintf("%s/%s", apiPaths["a2_v1"], ID)

	resp, err := nls.client.NewRequest("POST", apiURI, nil, nil)

	return resp, err

}
