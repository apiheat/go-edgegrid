package edgegrid

import (
	"fmt"
	"net/http"
	"time"
)

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

type APICredentialDetails struct {
	CredentialID int       `json:"credentialId"`
	ClientToken  string    `json:"clientToken"`
	Status       string    `json:"status"`
	CreatedOn    time.Time `json:"createdOn"`
	Description  string    `json:"description"`
	ExpiresOn    time.Time `json:"expiresOn"`
	Actions      struct {
		Deactivate      bool `json:"deactivate"`
		Delete          bool `json:"delete"`
		Activate        bool `json:"activate"`
		EditDescription bool `json:"editDescription"`
		EditExpiration  bool `json:"editExpiration"`
	} `json:"actions"`
}

type APIAccountSwitchKey struct {
	AccountSwitchKey string `json:"accountSwitchKey"`
	AccountName      string `json:"accountName"`
}

type QStrAPIClientCredentials struct {
	Actions bool   `url:"actions,omitempty"`
	Search  string `url:"search,omitempty"`
}

// ██╗   ██╗███████╗███████╗██████╗
// ██║   ██║██╔════╝██╔════╝██╔══██╗
// ██║   ██║███████╗█████╗  ██████╔╝
// ██║   ██║╚════██║██╔══╝  ██╔══██╗
// ╚██████╔╝███████║███████╗██║  ██║
//  ╚═════╝ ╚══════╝╚══════╝╚═╝  ╚═╝

func (nls *IdentityManagementService) ListUsers() (*[]AkamaiUser, *ClientResponse, error) {

	path := fmt.Sprintf("%s/user-admin/ui-identities", IdentityManagementPathV2)

	var k *[]AkamaiUser
	resp, err := nls.client.NewRequest(http.MethodGet, path, nil, &k)

	return k, resp, err
}

//  █████╗ ██████╗ ██╗
// ██╔══██╗██╔══██╗██║
// ███████║██████╔╝██║
// ██╔══██║██╔═══╝ ██║
// ██║  ██║██║     ██║
// ╚═╝  ╚═╝╚═╝     ╚═╝

// GetAPIClientCreds Lists API credentials
//
// Akamai API docs: https://developer.akamai.com/api/core_features/identity_management/v1.html#getcredentials
func (nls *IdentityManagementService) GetAPIClientCreds(openIdentityID string, includeActions bool) (*[]APICredentialDetails, *ClientResponse, error) {

	qParams := QStrAPIClientCredentials{Actions: includeActions}

	path := fmt.Sprintf("%s/open-identities/%s/credentials", IdentityManagementPathV1, openIdentityID)

	var respStruct *[]APICredentialDetails
	resp, err := nls.client.makeAPIRequest(http.MethodGet, path, qParams, &respStruct, nil, nil)

	return respStruct, resp, err
}

// ListAPISwitchKeys Lists account switch keys
//
// Akamai API docs: https://developer.akamai.com/api/core_features/identity_management/v1.html#getaccountswitchkeys
func (nls *IdentityManagementService) ListAPISwitchKeys(openIdentityID, searchPattern string) (*[]APIAccountSwitchKey, *ClientResponse, error) {
	qParams := QStrAPIClientCredentials{}
	path := fmt.Sprintf("%s/open-identities/%s/account-switch-keys", IdentityManagementPathV1, openIdentityID)

	if searchPattern != "" {
		qParams = QStrAPIClientCredentials{Search: searchPattern}
	}

	var respStruct *[]APIAccountSwitchKey
	resp, err := nls.client.makeAPIRequest(http.MethodGet, path, qParams, &respStruct, nil, nil)

	return respStruct, resp, err
}
