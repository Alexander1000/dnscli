/*
Copyright © 2021 Michael Bruskov <mixanemca@yandex.ru>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

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
