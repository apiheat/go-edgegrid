package edgeauth

//Credentials object represents items required for authentication
//towards Akamai' API
type Credentials struct {
	Host         string `ini:"host"`
	ClientToken  string `ini:"client_token"`
	ClientSecret string `ini:"client_secret"`
	AccessToken  string `ini:"access_token"`
}
