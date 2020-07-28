package search

import "github.com/mixanemca/dnscli/models"

// Client defines the interface for search operations.
type Client interface {
	// Search search the data inside PowerDNS
	Search(query string, max int, objectType models.ObjectType) (models.SearchResults, error)
}
