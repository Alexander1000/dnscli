package fz

import "github.com/mixanemca/dnscli/models"

// Client defines the interface for forwarding zone operations.
type Client interface {
	// Add creates new forwarding zone
	Add(models.ForwardZone) error
	// Delete delete forwarding zone
	Delete(name string) error
	// Get returns forwarding zone by name
	Get(name string) (models.ForwardZone, error)
	// List known forwarding zones
	List() (models.ForwardZones, error)
	// Update updates forwarding zone
	Update(models.ForwardZone) error
}
