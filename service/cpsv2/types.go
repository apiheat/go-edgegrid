package cpsv2

// OutputEnrollments identifies a collection of OutputEnrollmentElement objects
type OutputEnrollments struct {
	Enrollments []OutputEnrollmentElement `json:"enrollments"`
}

// OutputEnrollmentElement displays all the information about the process that your certificate goes through
// from the time you request it, through renewal, and as you obtain subsequent versions
// A version label indicates this member is introduced in that version.
// A pre-version label indicates this member is removed in that version.
// No version label indicates this member is present in all versions.
type OutputEnrollmentElement struct {
	// The URI path to the enrollment.
	// The last segment of the URI path serves as a unique identifier for the enrollment.
	Location string `json:"location,omitempty"`
	// The registration authority or certificate authority (CA) you want to use to obtain a certificate.
	// A CA is a trusted entity that signs certificates and can vouch for the identity of a website.
	// Either symantec, lets-encrypt, or third-party.
	Ra string `json:"ra"`
	// There are three types of validation. Domain Validation (DV), which is the lowest level of validation.
	// The CA validates that you have control of the domain.
	// CPS supports DV certificates issued by Let’s Encrypt, a free, automated, and open CA, run for public benefit.
	// Organization Validation (OV), which is the next level of validation. The CA validates that you have control of the domain.
	// Extended Validation (EV), which is the highest level of validation in which you must have signed letters and
	// notaries sent to the CA before signing. You can also specify third party as a type of validation,
	// if you want to use a signed certificate obtained by you from a CA not supported by CPS.
	// Either dv, ev, ov, or third-party.
	ValidationType string `json:"validationType"`
	// Either san, single, wildcard, wildcard-san, or third-party. See ValidationType Values for details.
	CertificateType string `json:"certificateType"`
	// v7. Certificate trust chain type.
	CertificateChainType string `json:"certificateChainType,omitempty"`
	// Settings that specify any network information and
	// TLS Metadata you want CPS to use to push the completed certificate to the network.
	NetworkConfiguration OutputEnrollmentNetworkConfiguration `json:"networkConfiguration"`
	// The SHA (Secure Hash Algorithm) function. Specify either SHA-1 or SHA-256. We recommend you use SHA–256.
	SignatureAlgorithm string `json:"signatureAlgorithm,omitempty"`
	// If you turn change management on for an enrollment, it stops CPS from deploying the certificate
	// to the network until you acknowledge that you are ready to deploy the certificate.
	ChangeManagement bool `json:"changeManagement"`
	// When you create an enrollment, you also generate a certificate signing request (CSR) using CPS.
	// CPS signs the CSR with the private key. The CSR contains all the information the CA needs to issue your certificate.
	Csr OutputEnrollmentCsr `json:"csr"`
	// Your organization information.
	Org OutputEnrollmentOrg `json:"org"`
	// Contact information for the certificate administrator that you want to use as a contact at your company.
	AdminContact OutputEnrollmentContact `json:"adminContact"`
	// Contact information for an administrator at Akamai.
	TechContact OutputEnrollmentContact `json:"techContact"`
	// Specifies that you want to use a third party certificate. This is any certificate that is not issued through CPS.
	ThirdParty *OutputEnrollmentThirdParty `json:"thirdParty,omitempty"`
	// v7. Enable Dual-Stacked certificate deployment for this enrollment.
	EnableMultiStackedCertificates bool `json:"enableMultiStackedCertificates,omitempty"`
	// v9. The specific date on which the renewal automatically starts for the enrollment.
	AutoRenewalStartTime string `json:"autoRenewalStartTime,omitempty"`
	// Returns the Changes currently pending in CPS. The last item in the array is the most recent change.
	PendingChanges []string `json:"pendingChanges,omitempty"`
	// v7. Maximum number of SAN names supported for this enrollment type.
	MaxAllowedSanNames int `json:"maxAllowedSanNames,omitempty"`
	// v7. Maximum number of Wildcard SAN names supported for this enrollment type.
	MaxAllowedWildcardSanNames int `json:"maxAllowedWildcardSanNames,omitempty"`
}

// OutputEnrollmentNetworkConfiguration specifies any network information and
// TLS Metadata you want CPS to use to push the completed certificate to the network.
type OutputEnrollmentNetworkConfiguration struct {
	// v3. Set to the enum core to specify worldwide (includes China and Russia).
	// Set to the enum china+core to specify worldwide and China.
	// Set to the enum russia+core to specify worldwide and Russia.
	// You can only use this setting to include China and Russia if your Akamai contract specifies your ability to do so and
	// you have approval from the Chinese and Russian government.
	Geography string `json:"geography"`
	// v2. Set the type of deployment network you want to use.
	// Set Standard TLS (standard-tls) to deploy your certificate to Akamai’s standard secure network. It is not PCI compliant.
	// Set Enhanced TLS (enhanced-tls) to deploy your certificate to Akamai’s more secure network with PCI compliance capability.
	SecureNetwork string `json:"secureNetwork"`
	// Ciphers that you definitely want to include for your enrollment while deploying it on the network.
	// Defaults to ak-akamai-default when it is not set. For more information on cipher profiles,
	// see https://community.akamai.com/customers/s/article/SSL-TLS-Cipher-Profiles-for-Akamai-Secure-CDNrxdxm?language=mk
	MustHaveCiphers string `json:"mustHaveCiphers"`
	// Ciphers that you preferably want to include for your enrollment while deploying it on the network.
	// Defaults to ak-akamai-default when it is not set. For more information on cipher profiles,
	// see https://community.akamai.com/customers/s/article/SSL-TLS-Cipher-Profiles-for-Akamai-Secure-CDNrxdxm?language=mk
	PreferredCiphers string `json:"preferredCiphers"`
	// v2. Specify the TLS protocol version to disallow. CPS uses the TLS protocols that Akamai currently supports as a best practice.
	DisallowedTLSVersions []string `json:"disallowedTlsVersions"`
	// v7. SNI settings for your enrollment. Set to true to enable SNI-only for the enrollment.
	// This setting cannot be changed once an enrollment is created.
	SniOnly bool `json:"sniOnly"`
	// v7. Set to true to enable QUIC protocol.
	QuicEnabled bool `json:"quicEnabled"`
	// v7. DNS name settings.
	DNSNameSettings *OutputEnrollmentDNSNameSettings `json:"dnsNameSettings,omitempty"`
	// v7. Enable OCSP stapling for the enrollment. OCSP Stapling improves performance by including a valid OCSP response in every TLS handshake.
	// Specify OCSP Stapling if you want to improve performance by allowing the visitors to your site to query
	// the Online Certificate Status Protocol (OCSP) server at regular intervals to obtain a signed time-stamped OCSP response.
	// This response must be signed by the CA, not the server, therefore ensuring security.
	// Disable OSCP Stapling if you want visitors to your site to contact the CA directly for an OSCP response.
	// OCSP allows you to obtain the revocation status of a certificate. We recommend all customers enable this feature.
	// Use on, off or not-set.
	OcspStapling string `json:"ocspStapling"`
	// v9. The configuration for client mutual authentication.
	// Specifies the trust chain that is used to verify client certificates and some configuration options.
	ClientMutualAuthentication OutputEnrollmentClientMutualAuth `json:"clientMutualAuthentication"`
}

// OutputEnrollmentClientMutualAuth The configuration for client mutual authentication.
type OutputEnrollmentClientMutualAuth struct {
	// v9. The identifier of the set of trust chains, created in the Trust Chain Manager.
	SetID string `json:"setId,omitempty"`
	// v9. Contains the configuration options for the selected trust chain.
	AuthenticationOptions *OutputEnrollmentClientMutualAuthOptions `json:"authenticationOptions,omitempty"`
}

// OutputEnrollmentClientMutualAuthOptions contains the configuration options for the selected trust chain.
type OutputEnrollmentClientMutualAuthOptions struct {
	// v9. Whether you want to enable the server to send the certificate authority (CA) list to the client.
	SendCaListToClient bool `json:"sendCaListToClient,omitempty"`
	// v9. Whether you want to enable ocsp stapling for client certificates.
	Ocsp *OutputEnrollmentClientMutualAuthOptionOcsp `json:"ocsp,omitempty"`
}

// OutputEnrollmentClientMutualAuthOptionOcsp Whether you want to enable ocsp stapling for client certificates.
type OutputEnrollmentClientMutualAuthOptionOcsp struct {
	//Whether the ocsp stapling is enabled.
	Enabled bool `json:"enabled,omitempty"`
}

// OutputEnrollmentDNSNameSettings provides DNS name settings.
type OutputEnrollmentDNSNameSettings struct {
	// v7. Enable if you want CPS to direct traffic using all the SANs listed in the SANs parameter when you created your enrollment.
	CloneDNSNames bool `json:"cloneDnsNames"`
	// Names served by SNI-only enabled enrollments.
	DNSNames []string `json:"dnsNames,omitempty"`
}

// OutputEnrollmentCsr when you create an enrollment, you also generate a certificate signing request (CSR) using CPS.
// CPS signs the CSR with the private key. The CSR contains all the information the CA needs to issue your certificate.
type OutputEnrollmentCsr struct {
	// The common name (CN) you want to use for the certificate in the Common Name field.
	Cn string `json:"cn"`
	// The country code for the country where your organization is located.
	C string `json:"c,omitempty"`
	// Your state or province.
	St string `json:"st,omitempty"`
	// Your city in the locality (city).
	L string `json:"l,omitempty"`
	// The name of your company or organization.
	// Enter the name as it appears in all legal documents and as it appears in the legal entity filing.
	O string `json:"o,omitempty"`
	// Your organizational unit.
	Ou string `json:"ou,omitempty"`
	// Additional common names (CN) to create a Subject Alternative Names (SAN) list.
	Sans []string `json:"sans,omitempty"`
}

// OutputEnrollmentOrg provides your organization information.
type OutputEnrollmentOrg struct {
	// The name of your organization.
	Name string `json:"name,omitempty"`
	// The address of your organization.
	AddressLineOne string `json:"addressLineOne,omitempty"`
	// The address of your organization.
	AddressLineTwo string `json:"addressLineTwo,omitempty"`
	// The city where your organization resides.
	City string `json:"city,omitempty"`
	// The region where your organization resides.
	Region string `json:"region,omitempty"`
	// The postal code of your organization.
	PostalCode string `json:"postalCode,omitempty"`
	// The country where your organization resides.
	Country string `json:"country,omitempty"`
	// The phone number of the administrator who you want to use as a contact at your company.
	Phone string `json:"phone,omitempty"`
}

// OutputEnrollmentContact provides Admin or Tech contact information for the certificate
type OutputEnrollmentContact struct {
	// The first name of the contact
	FirstName string `json:"firstName,omitempty"`
	// The last name of the  contact
	LastName string `json:"lastName,omitempty"`
	// The phone number of the contact
	Phone string `json:"phone,omitempty"`
	// The email of the contact
	Email string `json:"email,omitempty"`
	// The address of the contact
	AddressLineOne string `json:"addressLineOne,omitempty"`
	// The address of the contact
	AddressLineTwo string `json:"addressLineTwo,omitempty"`
	// The city of the contact
	City string `json:"city,omitempty"`
	// The country of the contact
	Country string `json:"country,omitempty"`
	// The organization name of the contact
	OrganizationName string `json:"organizationName,omitempty"`
	// The postal code of the contact
	PostalCode string `json:"postalCode,omitempty"`
	// The region of the contact
	Region string `json:"region,omitempty"`
	// The title of the contact
	Title string `json:"title,omitempty"`
}

// OutputEnrollmentThirdParty specifies that you want to use a third party certificate.
// This is any certificate that is not issued through CPS.
type OutputEnrollmentThirdParty struct {
	// If this is true, then the SANs in the enrollment do not appear in the CSR that CPS submits to the CA.
	ExcludeSans bool `json:"excludeSans"`
}
