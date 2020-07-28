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
