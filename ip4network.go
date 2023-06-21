package gobam

import (
	"fmt"
	"strconv"
	"strings"
)

// IP4NetworkProperties contains all properties returned by an IP4Network
type IP4NetworkProperties struct {
	Name                      string
	CIDR                      string
	Template                  string
	Gateway                   string
	DefaultDomains            []string
	DefaultView               string
	DNSRestrictions           []string
	AllowDuplicateHost        string
	PingBeforeAssign          string
	InheritAllowDuplicateHost bool
	InheritPingBeforeAssign   bool
	InheritDNSRestrictions    bool
	InheritDefaultDomains     bool
	InheritDefaultView        bool
	LocationCode              string
	LocationInherited         bool
	CustomProperties          map[string]string
}

// IP4NetworkPropertiesString parses an IP4NetworkProperties struct into an IP4Network properties string
func IP4NetworkPropertiesString(networkProperties IP4NetworkProperties) string {
	properties := ""

	if networkProperties.Name != "" {
		properties = properties + "name=" + networkProperties.Name + "|"
	}

	if networkProperties.Gateway != "" {
		properties = properties + "gateway=" + networkProperties.Gateway + "|"
	}

	if len(networkProperties.DefaultDomains) > 0 {
		properties = properties + "defaultDomains=" + strings.Join(networkProperties.DefaultDomains, ",") + "|"
	}

	if networkProperties.DefaultView != "" {
		properties = properties + "defaultView=" + networkProperties.DefaultView + "|"
	}

	if len(networkProperties.DNSRestrictions) > 0 {
		properties = properties + "dnsRestrictions=" + strings.Join(networkProperties.DNSRestrictions, ",") + "|"
	}

	if networkProperties.AllowDuplicateHost != "" {
		properties = properties + "allowDuplicateHost=" + networkProperties.AllowDuplicateHost + "|"
	}

	if networkProperties.PingBeforeAssign != "" {
		properties = properties + "pingBeforeAssign=" + networkProperties.PingBeforeAssign + "|"
	}

	properties = properties + "inheritAllowDuplicateHost=" + strconv.FormatBool(networkProperties.InheritAllowDuplicateHost) + "|"
	properties = properties + "inheritAllowDuplicateHost=" + strconv.FormatBool(networkProperties.InheritPingBeforeAssign) + "|"
	properties = properties + "inheritAllowDuplicateHost=" + strconv.FormatBool(networkProperties.InheritDNSRestrictions) + "|"
	properties = properties + "inheritAllowDuplicateHost=" + strconv.FormatBool(networkProperties.InheritDefaultDomains) + "|"
	properties = properties + "inheritAllowDuplicateHost=" + strconv.FormatBool(networkProperties.InheritDefaultView) + "|"

	if networkProperties.LocationCode != "" {
		properties = properties + "locationCode=" + networkProperties.LocationCode + "|"
	}

	for k, v := range networkProperties.CustomProperties {
		properties = properties + k + "=" + v + "|"
	}

	return properties
}

// ParseIP4NetworkProperties parses an IP4Network properties string into an IP4NetworkProperties struct
func ParseIP4NetworkProperties(properties string) (IP4NetworkProperties, error) {
	var networkProperties IP4NetworkProperties
	networkProperties.CustomProperties = make(map[string]string)

	props := strings.Split(properties, "|")
	for x := range props {
		if len(props[x]) > 0 {
			prop := strings.Split(props[x], "=")[0]
			val := strings.Split(props[x], "=")[1]

			switch prop {
			case "name":
				networkProperties.Name = val
			case "CIDR":
				networkProperties.CIDR = val
			case "template":
				networkProperties.Template = val
			case "gateway":
				networkProperties.Gateway = val
			case "defaultDomains":
				defaultDomains := strings.Split(val, ",")
				for i := range defaultDomains {
					networkProperties.DefaultDomains = append(networkProperties.DefaultDomains, defaultDomains[i])
				}
			case "defaultView":
				networkProperties.DefaultView = val
			case "dnsRestrictions":
				dnsRestrictions := strings.Split(val, ",")
				for i := range dnsRestrictions {
					networkProperties.DNSRestrictions = append(networkProperties.DNSRestrictions, dnsRestrictions[i])
				}
			case "allowDuplicateHost":
				networkProperties.AllowDuplicateHost = val
			case "pingBeforeAssign":
				networkProperties.PingBeforeAssign = val
			case "inheritAllowDuplicateHost":
				b, err := strconv.ParseBool(val)
				if err != nil {
					return networkProperties, fmt.Errorf("error parsing inheritAllowDuplicateHost to bool")
				}
				networkProperties.InheritAllowDuplicateHost = b
			case "inheritPingBeforeAssign":
				b, err := strconv.ParseBool(val)
				if err != nil {
					return networkProperties, fmt.Errorf("error parsing inheritPingBeforeAssign to bool")
				}
				networkProperties.InheritPingBeforeAssign = b
			case "inheritDNSRestrictions":
				b, err := strconv.ParseBool(val)
				if err != nil {
					return networkProperties, fmt.Errorf("error parsing inheritDNSRestrictions to bool")
				}
				networkProperties.InheritDNSRestrictions = b
			case "inheritDefaultDomains":
				b, err := strconv.ParseBool(val)
				if err != nil {
					return networkProperties, fmt.Errorf("error parsing inheritDefaultDomains to bool")
				}
				networkProperties.InheritDefaultDomains = b
			case "inheritDefaultView":
				b, err := strconv.ParseBool(val)
				if err != nil {
					return networkProperties, fmt.Errorf("error parsing inheritDefaultView to bool")
				}
				networkProperties.InheritDefaultView = b
			case "locationCode":
				networkProperties.LocationCode = val
			case "locationInherited":
				b, err := strconv.ParseBool(val)
				if err != nil {
					return networkProperties, fmt.Errorf("error parsing locationInherited to bool")
				}
				networkProperties.LocationInherited = b
			default:
				networkProperties.CustomProperties[prop] = val
			}
		}
	}

	return networkProperties, nil
}
