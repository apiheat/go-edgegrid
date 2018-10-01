package edgegrid

import "fmt"

type IdentityManagementService struct {
	client *Client
}

// AkamaiUser data
type AkamaiUser struct {
	UIIdentityID  string `json:"uiIdentityId"`
	FirstName     string `json:"firstName"`
	LastName      string `json:"lastName"`
	UIUserName    string `json:"uiUserName"`
	Email         string `json:"email"`
	AccountID     string `json:"accountId"`
	LastLoginDate string `json:"lastLoginDate"`
	TfaEnabled    bool   `json:"tfaEnabled"`
	TfaConfigured bool   `json:"tfaConfigured"`
}

func (nls *IdentityManagementService) ListUsers() (*[]AkamaiUser, *ClientResponse, error) {

	apiURI := fmt.Sprintf("%s/user-admin/ui-identities", IdentityManagementPathV1)

	var k *[]AkamaiUser
	resp, err := nls.client.NewRequest("GET", apiURI, nil, &k)

	return k, resp, err
}
