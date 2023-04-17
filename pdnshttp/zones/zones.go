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

// GetByName return zone with one authoritative zone by name argument and include the “rrsets” in the response
func (c *client) GetByName(name string) (*models.Zone, error) {
	zone := models.Zone{}
	path := fmt.Sprintf("/api/v1/servers/localhost/zones/%s", url.PathEscape(name))
	err := c.httpClient.Get(path, &zone)
	if err != nil {
		return nil, err
	}
	return &zone, nil
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
