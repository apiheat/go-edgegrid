package edgegrid

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"

	"github.com/asaskevich/govalidator"
	"github.com/go-ini/ini"
	log "github.com/sirupsen/logrus"
)

// Akamai {OPEN} EdgeGrid Authentication Service
type AuthService struct {
	client *Client
}

// //EdgercCredentials are items from config file
// type EdgercCredentials struct {
// 	host         string `ini:"host"`
// 	clientToken  string `ini:"client_token"`
// 	clientSecret string `ini:"client_secret"`
// 	accessToken  string `ini:"access_token"`
// }

type reader struct {
	*bytes.Buffer
}

func (m reader) Close() error { return nil }

// Init initializes using a configuration file in standard INI format
func InitEdgerc(edgercConfig, edgercSection string) (*EdgercCredentials, error) {

	log.WithFields(log.Fields{
		"edgercConfig":  edgercConfig,
		"edgercSection": edgercSection,
	}).Info("[InitEdgerc]::Initialize credentials")

	// Load the file based on our provided config
	log.Debug("[InitEdgerc]::Loading credentials file")
	edgerc, err := ini.Load(edgercConfig)
	if err != nil {
		return nil, fmt.Errorf("Error loading file? '%s'", err)
	}

	log.Debug("[InitEdgerc]::Loading section from credentials file")
	sectionNames := edgerc.SectionStrings()
	if !(stringInSlice(edgercSection, sectionNames)) {
		return nil, fmt.Errorf("Could not load section '%s'", edgercSection)
	}

	log.Debug("[InitEdgerc]::Lookup for credentials ( host/secrets etc)")
	edgercHost := edgerc.Section(edgercSection).Key("host").String()
	edgercclientToken := edgerc.Section(edgercSection).Key("client_token").String()
	edgercclientSecret := edgerc.Section(edgercSection).Key("client_secret").String()
	edgercaccessToken := edgerc.Section(edgercSection).Key("access_token").String()

	log.Debug("[InitEdgerc]::Create credentials object")
	loadedCredentials := &EdgercCredentials{
		host:         edgercHost,
		clientToken:  edgercclientToken,
		clientSecret: edgercclientSecret,
		accessToken:  edgercaccessToken,
	}

	err = validateCreds(loadedCredentials, fmt.Sprintf("section '%s'", edgercSection))
	if err != nil {
		return nil, err
	}

	log.Debug("[InitEdgerc]::Map credentials to appropiate object")
	err = edgerc.Section(edgercSection).MapTo(loadedCredentials)
	if err != nil {
		return nil, fmt.Errorf("Error loading file? %s", err)
	}

	log.Debug("[InitEdgerc]::Return credentials object")
	return loadedCredentials, nil

}

// removeStringDuplicates removes duplicated elements from strings array
func removeStringDuplicates(str []string) []string {
	// Use map to record duplicates as we find them.
	encountered := map[string]bool{}
	result := []string{}

	for v := range str {
		if encountered[str[v]] == true {
			// Do not add duplicate.
		} else {
			// Record this element as an encountered element.
			encountered[str[v]] = true
			// Append to result slice.
			if str[v] != "" {
				result = append(result, str[v])
			}
		}
	}
	// Return the new slice.
	return result
}

func validateCreds(creds *EdgercCredentials, location string) error {
	log.Debug("[InitEdgerc]::Validating credentials - 'host'")
	if creds.host == "" {
		return fmt.Errorf("'host' is empty in '%s'", location)
	}

	valid := govalidator.IsURL(creds.host)
	if !valid {
		return fmt.Errorf("'host' is not valid URL in '%s'", location)
	}

	u, err := url.Parse(creds.host)
	if err != nil {
		return fmt.Errorf("'host' cannot be parsed correctly in '%s'", location)
	}

	if u.Scheme != "" {
		return fmt.Errorf("'host' in '%s' contains URL scheme: '%s', please remove '%s//'", location, u.Scheme, u.Scheme)
	}

	log.Debug("[InitEdgerc]::Validating credentials - 'client_token'")
	if creds.clientToken == "" {
		return fmt.Errorf("'client_token' is empty in '%s'", location)
	}

	log.Debug("[InitEdgerc]::Validating credentials - 'client_secret'")
	if creds.clientSecret == "" {
		return fmt.Errorf("'client_secret' is empty in '%s'", location)
	}

	log.Debug("[InitEdgerc]::Validating credentials - 'access_token'")
	if creds.accessToken == "" {
		return fmt.Errorf("'access_token' is empty in '%s'", location)
	}

	return nil
}

//#TODO: Move to common CLI
func urlPathWithQuery(req *http.Request) string {
	var query string

	if req.URL.RawQuery != "" {
		query = concat([]string{
			"?",
			req.URL.RawQuery,
		})
	} else {
		query = ""
	}

	return concat([]string{
		req.URL.Path,
		query,
	})
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func concat(arr []string) string {
	var buff bytes.Buffer

	for _, elem := range arr {
		buff.WriteString(elem)
	}

	return buff.String()
}
