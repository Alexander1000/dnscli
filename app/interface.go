package app

import "github.com/mixanemca/dnscli/pdnshttp/fz"

// App is the root-level interface for interacting with the PowerDNS API.
// You can instantiate an implementation of this interface using the "New" function.
type App interface {
	// ForwardZones returns a specialized API for interacting with PowerDNS forwarding zones
	ForwardZones() fz.Client
}
