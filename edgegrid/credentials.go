package edgegrid

import (
	"fmt"
	"os"

	"github.com/asaskevich/govalidator"
	"github.com/go-ini/ini"
	log "github.com/sirupsen/logrus"
	"github.com/thedevsaddam/gojsonq"
)

// A Credentials provides values and string type of credentials being used.
// Type can be of `netstorage` or `api`
type Credentials struct {
	//API based credentials
	Host         string `ini:"host" json:"host" valid:"required~Host is empty/blank"`
	ClientToken  string `ini:"client_token" json:"client_token" valid:"required~ClientToken is blank/empty"`
	ClientSecret string `ini:"client_secret" json:"client_secret" valid:"required~ClientSecret name is blank/empty"`
	AccessToken  string `ini:"access_token" json:"access_token" valid:"required~AccessToken name is blank/empty"`

	//Netstorage based credentials
	HostName string `ini:"hostname"`
	Key      string `ini:"key"`
	KeyName  string `ini:"keyname"`
	CPCode   int    `ini:"cpcode"`
}

// ErrorCredentials represents an error caused during credentials retrieval
type ErrorCredentials struct {
	ErrorMessage string `json:"error_message"`
	ErrorType    string `json:"error_type"`
}

// ErrorCredentials implements the error interface.
func (e ErrorCredentials) Error() string {
	return e.ErrorMessage
}

//CredentialsBuilder provides method to build credentials using
//methods chaining for easy retrieval.
type CredentialsBuilder struct {
	*Credentials

	// edgercFile & edgercSection being used when
	// retrieving credentials from file
	edgercFile    string
	edgercSection string

	// credentialsType defines what type of credentials we are dealing with
	// Can be either `api` or `netstorage`
	credentialsType string
}

//AutoLoad Tries to load credentials automatically from a environment variables or from section file.
func (ea *CredentialsBuilder) AutoLoad(section string) *Credentials {
	var creds *Credentials
	var err error
	var homeDir string

	creds = &Credentials{}
	if section == "" {
		section = "default"
	}

	creds, err = NewCredentials().FromEnv()
	if err != nil {
		homeDir, err = os.UserHomeDir()
		if err != nil {
			return nil
		}
		log.WithFields(log.Fields{
			"homedir": homeDir,
		}).Debugln("Edgerc file location (.edgerc) ")

		log.WithFields(log.Fields{
			"section": section,
		}).Debugln("Section in credentials")

		creds, err = NewCredentials().FromFile(homeDir + "/.edgerc").Section(section)
		if err != nil {
			return nil
		}
	}

	return creds
}

// NewCredentials is used to create new object on which we can chain our methods
//
// Example for environment variables retrieval
// creds, err := credentials.NewCredentials().FromEnv()
// if err != nil {
// 	fmt.Println(err)
// }
func NewCredentials() *CredentialsBuilder {
	return &CredentialsBuilder{
		edgercSection: "default",
	}
}

// FromEnv Retrieves credentials from env variables which are prefixed with 'AKAMAI_'
// In order to sucesfully build credentials file we need the following variables:
//
// AKAMAI_HOST
// AKAMAI_CLIENT_TOKEN
// AKAMAI_CLIENT_SECRET
// AKAMAI_ACCESS_TOKEN
//
// Example of using the environment variable credentials.
//
//     credValue, err  := credentials.NewEnvCredentials().FromEnv()
//     if err != nil {
//         // handle error
//     }
func (ea *CredentialsBuilder) FromEnv() (*Credentials, error) {
	e := ErrorCredentials{}

	log.Debugln("Loading credentials from environment variables")
	var (
		requiredOptions = []string{"HOST", "CLIENT_TOKEN", "CLIENT_SECRET", "ACCESS_TOKEN"}
		missing         []string
	)

	prefix := "AKAMAI_"
	envCredentials := &Credentials{}

	for _, opt := range requiredOptions {
		val, ok := os.LookupEnv(prefix + opt)
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

	if len(missing) > 0 {
		e.ErrorMessage = fmt.Sprintf("Missing required environment variables: %s", missing)
		e.ErrorType = "ErrorCredentialsMissingField"
		// log.Debuglnln(e.ErrorMessage)

		return nil, e
	}

	result, err := govalidator.ValidateStruct(envCredentials)
	if err != nil {
		e.ErrorMessage = fmt.Sprintf("Environment variables are not correct: %s", err.Error())
		e.ErrorType = "ErrorCredentialValidation"
		// log.Debuglnln(e.ErrorMessage)

		return nil, e
	}

	log.Debugf("Credentials from environment variables validated to: %v", result)

	return envCredentials, nil
}

// FromJSON Retrieves credentials from a given JSON string.
// Example use:
// creds, err := credentials.NewCredentials().FromJSON(`{ "client_secret": "x", "host": "y", "access_token": "z", "client_token": "b" }`)
// if err != nil {
// 	fmt.Println(err)
// }
func (ea *CredentialsBuilder) FromJSON(json string) (*Credentials, error) {
	e := ErrorCredentials{}
	log.Debugln("Loading credentials from JSON string")

	credentials := &Credentials{}
	gojsonq.New().FromString(json).Out(credentials)

	result, err := govalidator.ValidateStruct(credentials)
	if err != nil {
		e.ErrorMessage = fmt.Sprintf("JSON credentials are not correct: %s", err.Error())
		e.ErrorType = "ErrorCredentialValidation"
		log.Error(e.ErrorMessage)

		return nil, e
	}

	log.Debugf("Credentials from JSON validated to: %v", result)

	return credentials, nil
}

// FromFile Retrieves credentials from the file ( and section )  specified
//
//	creds, err := credentials.NewCredentials().FromFile("/Users/username/.edgerc").Section("abc")
//	if err != nil {
// 		fmt.Println(err)
// 	}
func (ea *CredentialsBuilder) FromFile(fileName string) *CredentialsBuilder {
	ea.edgercFile = fileName
	log.Debugf("Set file name for retrieval: %s", fileName)

	return ea
}

//Section Should be used in conjuction with FromFile() and defines which section to read credentials from.
func (ea *CredentialsBuilder) Section(section string) (*Credentials, error) {
	e := ErrorCredentials{}

	ea.edgercSection = section

	log.WithFields(log.Fields{
		"file": ea.edgercFile,
	}).Debugln("Loading credentials file")

	edgerc, err := ini.Load(ea.edgercFile)
	if err != nil {
		e.ErrorMessage = fmt.Sprintf("%s", err.Error())
		e.ErrorType = "ErrorCredentialFile"

		return nil, e
	}

	log.WithFields(log.Fields{
		"section": section,
	}).Debugln("Loading section from credentials file")

	sectionNames := edgerc.SectionStrings()

	if !(stringInSlice(ea.edgercSection, sectionNames)) {
		e.ErrorMessage = fmt.Sprintf("%s", "Could not find respective section in credentials file")
		e.ErrorType = "ErrorCredentialSection"

		return nil, e
	}

	log.Debugln("Create & map credentials object")
	credentials := &Credentials{}
	edgerc.Section(ea.edgercSection).MapTo(credentials)

	log.Debugln("Validating credentials")

	result, err := govalidator.ValidateStruct(credentials)
	if err != nil {
		e.ErrorMessage = fmt.Sprintf("JSON credentials are not correct: %s", err.Error())
		e.ErrorType = "ErrorCredentialValidation"

		return nil, e
	}

	log.Debugf(fmt.Sprintf("Credentials from file validated to: %v", result))
	return credentials, nil

}

//stringInSlice is a private helper for string operations.
func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
