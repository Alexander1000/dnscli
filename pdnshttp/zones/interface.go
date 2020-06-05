package zones

import "github.com/mixanemca/dnscli/models"

// Client defines the interface for forwarding zone operations.
type Client interface {
	// Add creates new zone
	Add(zone models.Zone) (*models.Zone, error)
	// AddRecordSet will add a new set of records to a zone. Existing record
	// sets for the exact name/type combination will be replaced.
	AddRecordSet(zone string, set models.ResourceRecordSet) error
	// Delete delete zone from authoritative server
	Delete(name string) error
	// DeleteRecordSet removes a record set from a zone. The record set is matched
	// by name and type.
	DeleteRecordSet(zone, name, rrtype string) error
	// List known authoritative zones
	List() (models.Zones, error)
	// ListByName return list with one authoritative zone by name argument
	ListByName(name string) (models.Zones, error)
}
