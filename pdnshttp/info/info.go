package info

import (
	"github.com/mixanemca/dnscli/models"
)

// Get returns server version and build info
func (c *client) Get() (*models.InfoResult, error) {
	result := &models.InfoResult{}
	path := "/api/v1/version"
	err := c.httpClient.Get(path, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
