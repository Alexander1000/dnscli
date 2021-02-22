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

package fz

import "github.com/mixanemca/dnscli/models"

// Client defines the interface for forwarding zone operations.
type Client interface {
	// Add creates new forwarding zone
	Add(models.ForwardZones) error
	// Delete deletes forwarding zones
	Delete(models.ForwardZones) error
	// DeleteByName delete forwarding zone
	DeleteByName(name string) error
	// Get returns forwarding zone by name
	Get(name string) (models.ForwardZone, error)
	// List known forwarding zones
	List() (models.ForwardZones, error)
	// Update updates forwarding zone
	Update(models.ForwardZone) error
}
