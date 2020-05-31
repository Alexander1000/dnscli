package app

import (
	"net/http"
	"time"

	"github.com/mixanemca/dnscli/pdnshttp"
	"github.com/mixanemca/dnscli/pdnshttp/fz"
)

const (
	// BaseURL of PowerDNS API
	BaseURL string = "http://127.0.0.1:8081"
	// DefaultClientTimeout is time to wait before cancelling the request
	DefaultClientTimeout time.Duration = 5 * time.Second
)

type app struct {
	baseURL    string
	httpClient *http.Client

	fz fz.Client
}

// Option options for app
type Option func(c *app) error

// New creates a new PowerDNS client. Various client options can be used to configure
// the PowerDNS client
func New(opt ...Option) (App, error) {
	a := app{
		baseURL: BaseURL,
		httpClient: &http.Client{
			Timeout: DefaultClientTimeout,
		},
	}

	for i := range opt {
		if err := opt[i](&a); err != nil {
			return nil, err
		}
	}

	hc := pdnshttp.NewPDNSClient(a.baseURL, a.httpClient.Timeout)
	a.fz = fz.New(hc)

	return &a, nil
}

// SetBaseURL overrides the default BaseURL
func (a *app) SetBaseURL(b string) {
	a.baseURL = b
}

// SetTimeout overrides the default ClientTimeout
func (a *app) SetTimeout(d time.Duration) {
	a.httpClient.Timeout = d
}

func (a *app) ForwardZones() fz.Client {
	return a.fz
}
