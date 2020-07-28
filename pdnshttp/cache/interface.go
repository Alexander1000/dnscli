package cache

import "github.com/mixanemca/dnscli/models"

// Client defines the interface for operations with cache.
type Client interface {
	// Flush flush a cache-entry by name
	Flush(name string) (*models.FlushResult, error)
}
