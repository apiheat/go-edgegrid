package edgeauth

import (
	"errors"
	"fmt"
	"net/url"
	"os"

	"github.com/asaskevich/govalidator"
	log "github.com/sirupsen/logrus"
)

//Edgeauth
type Edgeauth struct {
}

//GetCredentialsFromEnv Retrieves credentials from env variables which are prefixed with 'AKAMAI_'
//In order to sucesfully build credentials file we need the following variables:
// AKAMAI_HOST
// AKAMAI_CLIENT_TOKEN
// AKAMAI_CLIENT_SECRET
// AKAMAI_ACCESS_TOKEN
// Returns new EdgercCredentials or error
func (ea *Edgeauth) GetCredentialsFromEnv() (*EdgercCredentials, error) {

	log.Debug("[InitEdgerc]::Loading credentials from environment variables")
	var (
		requiredOptions = []string{"HOST", "CLIENT_TOKEN", "CLIENT_SECRET", "ACCESS_TOKEN"}
		missing         []string
	)

	prefix := "AKAMAI_"
	envCredentials := &EdgercCredentials{}

	for _, opt := range requiredOptions {
		val, ok := os.LookupEnv(prefix + opt)
		if val == "" {
			missing = append(missing, prefix+opt)
		}
		if !ok {
			missing = append(missing, prefix+opt)
		} else {
			switch {
			case opt == "HOST":
				envCredentials.Host = val
			case opt == "CLIENT_TOKEN":
				envCredentials.ClientToken = val
			case opt == "CLIENT_SECRET":
				envCredentials.ClientSecret = val
			case opt == "ACCESS_TOKEN":
				envCredentials.AccessToken = val
			}
		}
	}

	missing = removeStringDuplicates(missing)

	if len(missing) > 0 {
		log.Debug(fmt.Sprintf("[InitEdgerc]::Missing required environment variables: %s", missing))
	}

	err := validateCredentials(envCredentials)
	if err == nil {
		log.Debug("[InitEdgerc]::Return credentials object from env variables")
		return envCredentials, nil
	}

	log.Debug(fmt.Sprintf("[InitEdgerc]::Environment variables are not correct: %s", err))

	return nil, nil
}

//removeStringDuplicates removes duplicated elements from strings array
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

//validateCredentials checks if credentials we have are correct
func validateCredentials(creds *EdgercCredentials) error {
	log.Debug("[validateCreds]::Validating credentials - 'host'")
	if creds.Host == "" {
		log.Error("[validateCreds]::'host' value is empty")
		return errors.New("'host' is empty")
	}

	valid := govalidator.IsURL(creds.Host)
	if !valid {
		log.Error("[validateCreds]::'host' is not a valid URL")
		return errors.New("'host' is not a valid URL")
	}

	u, err := url.Parse(creds.Host)
	if err != nil {
		log.Error("[validateCreds]::'host' cannot be parsed correctly")
		return errors.New("'host' cannot be parsed correctly")
	}

	if u.Scheme != "" {
		log.Error("[validateCreds]::contains URL scheme")
		return errors.New("'host' contains URL scheme")
	}

	log.Debug("[InitEdgerc]::Validating credentials - 'client_token'")
	if creds.ClientToken == "" {
		log.Error("[validateCreds]::'client_token' is empty")
		return errors.New("'client_token' is empty")
	}

	log.Debug("[InitEdgerc]::Validating credentials - 'client_secret'")
	if creds.ClientSecret == "" {
		log.Error("[validateCreds]::'client_secret' is empty")
		return errors.New("'client_secret' is empty")

	}

	log.Debug("[InitEdgerc]::Validating credentials - 'access_token'")
	if creds.AccessToken == "" {
		log.Error("[validateCreds]::'access_token' is empty")
		return errors.New("'access_token' is empty")
	}

	return nil
}
