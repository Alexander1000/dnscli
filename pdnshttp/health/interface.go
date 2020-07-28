package health

import "github.com/mixanemca/dnscli/models"

// Client defines the interface for server health info.
type Client interface {
	// Get returns server health info
	Get() (*models.HealthResult, error)
}
