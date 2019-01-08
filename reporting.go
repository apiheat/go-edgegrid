package edgegrid

import (
	"fmt"
	"net/http"
	"time"
)

type ReportingService struct {
	client *Client
}

type ReportingBody struct {
	ObjectType string   `json:"objectType"`
	ObjectIds  []string `json:"objectIds"`
	Metrics    []string `json:"metrics"`
}

type ReportingBodyAll struct {
	ObjectType string   `json:"objectType"`
	ObjectIds  string   `json:"objectIds"`
	Metrics    []string `json:"metrics"`
}

//TODO: Change TypeOfReport into string consts ?
//ReportOptions represents options available for report generation
type ReportOptions struct {
	TypeOfReport string
	Interval     string
	Start        time.Time
	End          time.Time
}

// QStrReporting includes query params used for reporting
type QStrReporting struct {
	Start    time.Time `url:"start,omitempty"`
	End      time.Time `url:"end,omitempty"`
	Interval string    `url:"interval,omitempty"`
}

//GenerateReport Calls reporing API to generate given report based on provided request
func (nls *ReportingService) GenerateReport(body interface{}, opts ReportOptions) (*ClientResponse, error) {
	// since this call is tricky with query params we are
	qParams := QStrReporting{
		Interval: opts.Interval,
		End:      opts.End,
		Start:    opts.Start,
	}

	path := fmt.Sprintf("%s/%s/versions/1/report-data", ReportingPathV1, opts.TypeOfReport)
	resp, err := nls.client.makeAPIRequest(http.MethodPost, path, qParams, nil, body, nil)

	return resp, err

}
