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
		pdnshttp.WithQueryValue("q", query),
		pdnshttp.WithQueryValue("max", fmt.Sprintf("%d", max)),
		pdnshttp.WithQueryValue("object_type", objectType.String()),
	)
	if err != nil {
		return nil, err
	}

	return results, nil
}
