package health

import (
	"github.com/mixanemca/dnscli/models"
)

// Get returns server health info
func (c *client) Get() (*models.HealthResult, error) {
	result := &models.HealthResult{}
	path := "/api/v1/health"
	err := c.httpClient.Get(path, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
