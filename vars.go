package gobam

import "net/http"

var sessionCookies []*http.Cookie

// EnableOrDisable contains the valid options for settings that must be Enabled or Disabled
var EnableOrDisable = []string{
	"Enable",
	"Disable",
}

// TraversalMethodology contains all valid traversal methodologies
var TraversalMethodology = []string{
	"NO_TRAVERSAL",
	"DEPTH_FIRST",
	"BREADTH_FIRST",
}

// IPAssignmentActions contains all valid action values for IP Assignment
var IPAssignmentActions = []string{
	"MAKE_STATIC",
	"MAKE_RESERVED",
	"MAKE_DHCP_RESERVED",
}

// ObjectTypes contains all valid object types in the BlueCat API
var ObjectTypes = []string{
	"Entity",
	"Configuration",
	"View",
	"Zone",
	"InternalRootZone",
	"ZoneTemplate",
	"EnumZone",
	"EnumNumber",
	"HostRecord",
	"AliasRecord",
	"MXRecord",
	"TXTRecord",
	"SRVRecord",
	"GenericRecord",
	"HINFORecord",
	"NAPTRRecord",
	"RecordWithLink",
	"ExternalHostRecord",
	"StartOfAuthority",
	"IP4Block",
	"IP4Network",
	"IP6Block",
	"IP6Network",
	"IP6Address",
	"IP4NetworkTemplate",
	"DHCP4Range",
	"DHCP6Range",
	"IP4Address",
	"MACPool",
	"DenyMACPool",
	"MACAddress",
	"TagGroup",
	"Tag",
	"User",
	"UserGroup",
	"Server",
	"ServerGroup",
	"NetworkServerInterface",
	"PublishedServerInterface",
	"NetworkInterface",
	"VirtualInterface",
	"LDAP",
	"Kerberos",
	"KerberosRealm",
	"Radius",
	"TFTPGroup",
	"TFTPFolder",
	"TFTPFile",
	"TFTPDeploymentRole",
	"DNSDeploymentRole",
	"DHCPDeploymentRole",
	"DNSOption",
	"DHCPV4ClientOption",
	"DHCPServiceOption",
	"DHCPRawOption",
	"DNSRawOption",
	"DHCPV6ClientOption",
	"DHCPV6ServiceOption",
	"DHCPV6RawOption",
	"VendorProfile",
	"VendorOptionDef",
	"VendorClientOption",
	"CustomOptionDef",
	"DHCPMatchClass",
	"DHCPSubClass",
	"Device",
	"DeviceType",
	"DeviceSubtype",
	"DeploymentScheduler",
	"IP4ReconciliationPolicy",
	"DNSSECSigningPolicy",
	"IP4IPGroup",
	"ResponsePolicy",
	"TSIGKey",
	"RPZone",
	"Location",
	"InterfaceID",
}
