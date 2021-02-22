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
	"crypto/tls"
	"crypto/x509"
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
	tlsEnable   bool
	tlsCAPath   string
	tlsCertPath string
	tlsKeyPath  string
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
	// App with default values
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

	// Use TLS transport
	if a.tlsEnable {
		tlsConfig, err := newTLSConfig(a.tlsCAPath, a.tlsCertPath, a.tlsKeyPath)
		if err != nil {
			return nil, err
		}
		a.httpClient.Transport = &http.Transport{
			TLSClientConfig: tlsConfig,
		}
	}

	hc := pdnshttp.NewPDNSClient(a.baseURL, a.httpClient, a.debugOutput)
	a.cache = cache.New(hc)
	a.fz = fz.New(hc)
	a.health = health.New(hc)
	a.info = info.New(hc)
	a.search = search.New(hc)
	a.zones = zones.New(hc)

	return a, nil
}

// newTLSConfig creates new tls.Config with certificates and key
func newTLSConfig(tlsCAPath, tlsCertPath, tlsKeyPath string) (*tls.Config, error) {
	// create a Certificate pool to hold one or more CA certificates
	caCertPool := x509.NewCertPool()

	// read CA certificate(s) and add to the Certificate Pool
	if tlsCAPath != "" {
		caCert, err := ioutil.ReadFile(tlsCAPath)
		if err != nil {
			return nil, err
		}
		if ok := caCertPool.AppendCertsFromPEM(caCert); !ok {
			return nil, err
		}
	}

	// Make a tls.Config with CA certificates and client TLS key and certificate
	cfg := &tls.Config{
		// controls whether a client verifies the server's certificate chain and host name
		// If InsecureSkipVerify is true, crypto/tls accepts any certificate presented by the server
		// and any host name in that certificate
		InsecureSkipVerify: false,
		RootCAs:            caCertPool,
		GetClientCertificate: func(*tls.CertificateRequestInfo) (*tls.Certificate, error) {
			c, err := tls.LoadX509KeyPair(tlsCertPath, tlsKeyPath)
			if err != nil {
				return nil, err
			}
			return &c, nil
		},
	}
	// parse cfg.Certificates and builds cfg.NameToCertificate from the CommonName and SubjectAlternateName
	// fields of each of the leaf certificates
	cfg.BuildNameToCertificate()

	return cfg, nil
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
