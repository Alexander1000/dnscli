/*
Copyright © 2021 Michael Bruskov <mixanemca@yandex.ru>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

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
