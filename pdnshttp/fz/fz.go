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
func (c *client) Add(fz models.ForwardZones) error {
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

// DeleteByName deletes forwarding zone by name
func (c *client) DeleteByName(name string) error {
	path := fmt.Sprintf("/api/v1/servers/localhost/forward-zones/%s", url.PathEscape(name))
	err := c.httpClient.Delete(path, nil)
	if err != nil {
		return err
	}
	return nil
}

// Delete deletes forwarding zones
func (c *client) Delete(fz models.ForwardZones) error {
	path := "/api/v1/servers/localhost/forward-zones"
	err := c.httpClient.Delete(path, nil, pdnshttp.WithJSONRequestBody(&fz))
	if err != nil {
		return err
	}
	return nil
}
