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

package cache

import (
	"github.com/mixanemca/dnscli/models"
	"github.com/mixanemca/dnscli/pdnshttp"
)

// Flush flush a cache-entry by name
func (c *client) Flush(name string) (*models.FlushResult, error) {
	result := &models.FlushResult{}
	path := "/api/v1/servers/localhost/cache/flush"
	err := c.httpClient.Put(path, result, pdnshttp.WithQueryValue("domain", name))
	if err != nil {
		return nil, err
	}
	return result, nil
}
