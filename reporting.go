package edgegrid

import (
	"fmt"
	"strings"
)

type ReportingAPIService struct {
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

func strToStrArr(str string) (strArr []string) {
	for _, s := range strings.Split(str, ",") {
		strArr = append(strArr, s)
	}
	return strArr
}

func (nls *ReportingAPIService) GenerateReport(body interface{}, opts AkamaiReportOptions) (*ClientResponse, error) {

	apiURI := fmt.Sprintf("%s/%s/versions/1/report-data?%s&interval=%s", apiPaths["reporting_v1"], opts.TypeOfReport, opts.DateRange, opts.Interval)

	resp, err := nls.client.NewRequest("POST", apiURI, body, nil)

	return resp, err

}
