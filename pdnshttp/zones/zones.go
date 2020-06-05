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

// Delete delete zone from authoritative server
func (c *client) Delete(name string) error {
	path := fmt.Sprintf("/api/v1/servers/localhost/zones/%s", url.PathEscape(name))
	err := c.httpClient.Delete(path, nil)
	if err != nil {
		return err
	}
	return nil
}

// AddRecordSet will add a new set of recorecord
// sets for the exact name/type combination
func (c *client) AddRecordSet(zone string, set models.ResourceRecordSet) error {
	path := fmt.Sprintf("/api/v1/servers/localhost/zones/%s", url.PathEscape(zone))

	set.ChangeType = models.ChangeTypeReplace
	patch := models.Zone{
		ResourceRecordSets: []models.ResourceRecordSet{
			set,
		},
	}

	return c.httpClient.Patch(path, nil, pdnshttp.WithJSONRequestBody(&patch))
}

// DeleteRecordSet removes a record set from a zone. The record set is matched
// by name and type.
func (c *client) DeleteRecordSet(zone, name, rrtype string) error {
	path := fmt.Sprintf("/api/v1/servers/localhost/zones/%s", url.PathEscape(zone))

	set := models.ResourceRecordSet{
		Name:       name,
		Type:       rrtype,
		ChangeType: models.ChangeTypeDelete,
	}

	patch := models.Zone{
		ResourceRecordSets: []models.ResourceRecordSet{set},
	}

	return c.httpClient.Patch(path, nil, pdnshttp.WithJSONRequestBody(&patch))
}
