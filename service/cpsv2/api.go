package cpsv2

import "fmt"

const (
	enrollmentVersion = "application/vnd.akamai.cps.enrollments.v9+json"
)

// ListEnrollments retrieves all enrollments.
func (cps *Cpsv2) ListEnrollments(contractID string) (*OutputEnrollments, error) {
	query := map[string]string{}

	if contractID != "" {
		query["contractId"] = contractID
	}

	apiURI := fmt.Sprintf("%s/enrollments", basePath)

	// Create and execute request
	resp, err := cps.Client.Rclient.R().
		SetResult(OutputEnrollments{}).
		SetError(CpsErrorv2{}).
		SetHeader("Accept", enrollmentVersion).
		SetQueryParams(query).
		Get(apiURI)

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		e := resp.Error().(*CpsErrorv2)

		return nil, e
	}

	return resp.Result().(*OutputEnrollments), nil
}
