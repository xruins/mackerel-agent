package linux

import (
	"net"

	"github.com/mackerelio/golib/logging"
	mkr "github.com/mackerelio/mackerel-client-go"
)

// InterfaceGenerator XXX
type InterfaceGenerator struct {
}

var interfaceLogger = logging.GetLogger("spec.interface")

// Generate XXX
func (g *InterfaceGenerator) Generate() ([]mkr.Interface, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	var results []mkr.Interface
	for _, iface := range ifaces {
		addrs, err := iface.Addrs()
		if err != nil {
			return nil, err
		}

		name := iface.Name
		if name == "" {
			continue
		}

		var ipv4Addresses, ipv6Addresses []string
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPAddr:
				ip = v.IP
			case *net.IPNet:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			if ipv4 := ip.To4(); ipv4 != nil {
				ipv4Addresses = append(ipv4Addresses, ipv4.String())
				continue
			}
			if ipv6 := ip.To16(); ipv6 != nil {
				ipv6Addresses = append(ipv6Addresses, ipv6.String())
			}
		}

		if len(ipv4Addresses) > 0 || len(ipv6Addresses) > 0 {
			results = append(results, mkr.Interface{
				Name:          name,
				IPv4Addresses: ipv4Addresses,
				IPv6Addresses: ipv6Addresses,
				MacAddress:    iface.HardwareAddr.String(),
			})
		}
	}
	return results, nil
}
