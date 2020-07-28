package info

import "github.com/mixanemca/dnscli/pdnshttp"

type client struct {
	httpClient *pdnshttp.PDNSClient
}

// New creates a new Info client
func New(hc *pdnshttp.PDNSClient) Client {
	return &client{
		httpClient: hc,
	}
}
