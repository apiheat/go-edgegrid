package edgegrid

// AkamaiEnvironmentVar represents Akamai's env variables used
type AkamaiEnvironmentVar string

// AkamaiEnvironment represents Akamai's target environment type.
type AkamaiEnvironment string

// AkamaiRequestFrom represents Akamai's source for request.
type AkamaiRequestFrom string

// AkamaiSubscription represents Akamai's notification actions for subscriptions.
type AkamaiSubscription string

const (
	EnvVarEdgercPath        AkamaiEnvironmentVar = "AKAMAI_EDGERC_CONFIG"
	EnvVarEdgercSection     AkamaiEnvironmentVar = "AKAMAI_EDGERC_SECTION"
	EnvVarDebugLevelSection AkamaiEnvironmentVar = "AKAMAI_EDGERC_DEBUGLEVEL"

	Production AkamaiEnvironment = "production"
	Staging    AkamaiEnvironment = "staging"

	Ghost     AkamaiRequestFrom = "ghost-locations"
	IPAddress AkamaiRequestFrom = "ip-addresses"

	Subscribe   AkamaiSubscription = "subscribe"
	Unsubscribe AkamaiSubscription = "unsubscribe"
)
