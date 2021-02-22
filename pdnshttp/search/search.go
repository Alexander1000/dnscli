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

package search

import (
	"fmt"

	"github.com/mixanemca/dnscli/models"
	"github.com/mixanemca/dnscli/pdnshttp"
)

// Search search the data inside PowerDNS
func (c *client) Search(query string, max int, objectType models.ObjectType) (models.SearchResults, error) {
	results := make(models.SearchResults, 0)
	path := "/api/v1/servers/localhost/search-data"
	err := c.httpClient.Get(
		path,
		&results,
		pdnshttp.WithQueryValue("q", models.DeCanonicalize(query)),
		pdnshttp.WithQueryValue("max", fmt.Sprintf("%d", max)),
		pdnshttp.WithQueryValue("object_type", objectType.String()),
	)
	if err != nil {
		return nil, err
	}

	return results, nil
}
