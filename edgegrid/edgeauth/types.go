package edgeauth

//EdgercCredentials are items from config file
type EdgercCredentials struct {
	Host         string `ini:"host"`
	ClientToken  string `ini:"client_token"`
	ClientSecret string `ini:"client_secret"`
	AccessToken  string `ini:"access_token"`
}
