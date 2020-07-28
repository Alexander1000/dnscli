package app

import (
	"github.com/mixanemca/dnscli/pdnshttp/cache"
	"github.com/mixanemca/dnscli/pdnshttp/fz"
	"github.com/mixanemca/dnscli/pdnshttp/health"
	"github.com/mixanemca/dnscli/pdnshttp/info"
	"github.com/mixanemca/dnscli/pdnshttp/search"
	"github.com/mixanemca/dnscli/pdnshttp/zones"
)

// App is the root-level interface for interacting with the PowerDNS API.
// You can instantiate an implementation of this interface using the "New" function.
type App interface {
	// Cache returns a specialized API for caching
	Cache() cache.Client
	// Health returns a specialized API for health info
	Health() health.Client
	// Info returns a specialized API for info
	Info() info.Client
	// ForwardZones returns a specialized API for interacting with PowerDNS forwarding zones
	ForwardZones() fz.Client
	// Search returns a specialized API for search the data in PowerDNS
	Search() search.Client
	// Zones returns a specialized API for interacting with PowerDNS authoritative zones
	Zones() zones.Client
}
