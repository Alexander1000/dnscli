package fz

import "github.com/mixanemca/dnscli/models"

// Client defines the interface for forwarding zone operations.
type Client interface {
	// List known forwarding zones
	List() (models.ForwardZones, error)
	// Add creates new forwarding zone
	Add(models.ForwardZone) error
	// Update updates forwarding zone
	Update(models.ForwardZone) error
	// Delete delete forwarding zone
	Delete(name string) error
}
