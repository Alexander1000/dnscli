package fz

import (
	"fmt"
	"net/url"

	"github.com/mixanemca/dnscli/models"
	"github.com/mixanemca/dnscli/pdnshttp"
)

// List returns list of a forwardign zones
func (c *client) List() (models.ForwardZones, error) {
	fzs := make(models.ForwardZones, 0)
	path := "/api/v1/servers/localhost/forward-zones"
	err := c.httpClient.Get(path, &fzs)
	if err != nil {
		return nil, err
	}
	return fzs, nil
}

// Get returns forwarding zone by name
func (c *client) Get(name string) (models.ForwardZone, error) {
	created := models.ForwardZone{}
	path := fmt.Sprintf("/api/v1/servers/localhost/forward-zones/%s", url.PathEscape(name))
	err := c.httpClient.Get(path, &created)
	if err != nil {
		return models.ForwardZone{}, err
	}
	return created, nil
}

// Add creates new forward zone
func (c *client) Add(fz models.ForwardZone) error {
	path := "/api/v1/servers/localhost/forward-zones"
	err := c.httpClient.Post(path, nil, pdnshttp.WithJSONRequestBody(&fz))
	if err != nil {
		return err
	}
	return nil
}

// Update updates forwarding zone
func (c *client) Update(fz models.ForwardZone) error {
	path := fmt.Sprintf("/api/v1/servers/localhost/forward-zones/%s", fz.Name)
	err := c.httpClient.Patch(path, nil, pdnshttp.WithJSONRequestBody(&fz))
	if err != nil {
		return err
	}
	return nil
}

// Delete delete forwarding zone
func (c *client) Delete(name string) error {
	path := fmt.Sprintf("/api/v1/servers/localhost/forward-zones/%s", url.PathEscape(name))
	err := c.httpClient.Delete(path, nil)
	if err != nil {
		return err
	}
	return nil
}
