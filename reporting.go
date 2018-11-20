package edgegrid

import (
	"fmt"
	"net/http"
)

type ReportingService struct {
	client *Client
}

type AkamaiReportingBody struct {
	ObjectType string   `json:"objectType"`
	ObjectIds  []string `json:"objectIds"`
	Metrics    []string `json:"metrics"`
}

type AkamaiReportingBodyAll struct {
	ObjectType string   `json:"objectType"`
	ObjectIds  string   `json:"objectIds"`
	Metrics    []string `json:"metrics"`
}

type AkamaiReportOptions struct {
	TypeOfReport string
	Interval     string
	DateRange    string
}

func (nls *ReportingService) GenerateReport(body interface{}, opts AkamaiReportOptions) (*ClientResponse, error) {

	apiURI := fmt.Sprintf("%s/%s/versions/1/report-data?%s&interval=%s", ReportingPathV1, opts.TypeOfReport, opts.DateRange, opts.Interval)

	resp, err := nls.client.NewRequest(http.MethodPost, apiURI, body, nil)

	return resp, err

}
