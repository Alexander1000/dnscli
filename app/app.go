package app

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/mixanemca/dnscli/pdnshttp"
	"github.com/mixanemca/dnscli/pdnshttp/cache"
	"github.com/mixanemca/dnscli/pdnshttp/fz"
	"github.com/mixanemca/dnscli/pdnshttp/health"
	"github.com/mixanemca/dnscli/pdnshttp/info"
	"github.com/mixanemca/dnscli/pdnshttp/search"
	"github.com/mixanemca/dnscli/pdnshttp/zones"
)

const (
	// BaseURL of PowerDNS API
	BaseURL string = "http://127.0.0.1:8081"
	// DefaultClientTimeout is time to wait before cancelling the request
	DefaultClientTimeout time.Duration = 5 * time.Second
)

type app struct {
	baseURL     string
	httpClient  *http.Client
	debugOutput io.Writer

	cache  cache.Client
	fz     fz.Client
	health health.Client
	info   info.Client
	search search.Client
	zones  zones.Client
}

// Option options for app
type Option func(c *app) error

// New creates a new PowerDNS client. Various client options can be used to configure
// the PowerDNS client
func New(opts ...Option) (App, error) {
	a := &app{
		baseURL: BaseURL,
		httpClient: &http.Client{
			Timeout: DefaultClientTimeout,
		},
		debugOutput: ioutil.Discard,
	}

	for _, opt := range opts {
		if err := opt(a); err != nil {
			return nil, err
		}
	}

	hc := pdnshttp.NewPDNSClient(a.baseURL, a.httpClient.Timeout, a.debugOutput)
	a.cache = cache.New(hc)
	a.fz = fz.New(hc)
	a.health = health.New(hc)
	a.info = info.New(hc)
	a.search = search.New(hc)
	a.zones = zones.New(hc)

	return a, nil
}

// SetBaseURL overrides the default BaseURL
func (a *app) SetBaseURL(b string) {
	a.baseURL = b
}

// SetTimeout overrides the default ClientTimeout
func (a *app) SetTimeout(d time.Duration) {
	a.httpClient.Timeout = d
}

// SetDebugOutput overrides the default debugOutput
func (a *app) SetDebugOutput(yes bool) {
	if yes {
		a.debugOutput = os.Stderr
	}
}

func (a *app) Cache() cache.Client {
	return a.cache
}

func (a *app) ForwardZones() fz.Client {
	return a.fz
}

func (a *app) Health() health.Client {
	return a.health
}

func (a *app) Info() info.Client {
	return a.info
}

func (a *app) Search() search.Client {
	return a.search
}

func (a *app) Zones() zones.Client {
	return a.zones
}
