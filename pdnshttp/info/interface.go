package info

import "github.com/mixanemca/dnscli/models"

// Client defines the interface for server info.
type Client interface {
	// Get returns server version and build info
	Get() (*models.InfoResult, error)
}
