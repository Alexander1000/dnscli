package zones

import "github.com/mixanemca/dnscli/models"

// Client defines the interface for forwarding zone operations.
type Client interface {
	// List known authoritative zones
	List() (models.Zones, error)
	// ListByName return list with one authoritative zone by name argument
	ListByName(name string) (models.Zones, error)
	// Add creates new zone
	Add(models.Zone) (*models.Zone, error)
	// Update updates forwarding zone
	// Update(models.ForwardZone) error
	// Delete delete forwarding zone
	// Delete(name string) error
}
