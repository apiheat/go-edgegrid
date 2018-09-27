package edgegrid

import (
	log "github.com/sirupsen/logrus"
)

// DebugService allows to interact with client debugging options
//
//
type DebugService struct {
	client *Client
}

// SetDebugLevel Function used to set appropiate
//
// Akamai API docs: https://developer.akamai.com/api/luna/papi/resources.html#getgroups
func (dbs *DebugService) SetDebugLevel(debugLevel string) {

	switch debugLevel {
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	case "fatal":
		log.SetLevel(log.FatalLevel)
	case "panic":
		log.SetLevel(log.PanicLevel)
	}
}
