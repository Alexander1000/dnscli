package zones

import (
	"fmt"
	"net/url"

	"github.com/mixanemca/dnscli/models"
	"github.com/mixanemca/dnscli/pdnshttp"
)

// List return list of an authoritative zones
func (c *client) List() (models.Zones, error) {
	zones := make(models.Zones, 0)
	path := "/api/v1/servers/localhost/zones"
	err := c.httpClient.Get(path, &zones)
	if err != nil {
		return nil, err
	}
	return zones, nil
}

// ListByName return list with one authoritative zone by name argument
func (c *client) ListByName(name string) (models.Zones, error) {
	zones := make(models.Zones, 0)
	path := "/api/v1/servers/localhost/zones"
	err := c.httpClient.Get(path, &zones, pdnshttp.WithQueryValue("zone", name))
	if err != nil {
		return nil, err
	}
	return zones, nil
}

// Add creates new authoritative zone
func (c *client) Add(z models.Zone) (*models.Zone, error) {
	created := models.Zone{}
	path := "/api/v1/servers/localhost/zones"

	z.ID = ""
	z.Type = models.ZoneTypeZone

	if z.Kind == 0 {
		z.Kind = models.ZoneKindNative
	}

	err := c.httpClient.Post(path, &created, pdnshttp.WithJSONRequestBody(&z))
	if err != nil {
		return nil, err
	}
	return &created, nil
}

/*
// Update updates forwarding zone
func (c *client) Update(fz models.ForwardZone) error {
	path := fmt.Sprintf("/api/v1/servers/localhost/forward-zones/%s", fz.Name)
	err := c.httpClient.Patch(path, nil, pdnshttp.WithJSONRequestBody(&fz))
	if err != nil {
		return err
	}
	return nil
}
*/

// Delete delete zone from authoritative server
func (c *client) Delete(name string) error {
	path := fmt.Sprintf("/api/v1/servers/localhost/zones/%s", url.PathEscape(name))
	err := c.httpClient.Delete(path, nil)
	if err != nil {
		return err
	}
	return nil
}
