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

package models

import "fmt"

// ZoneType zone type
type ZoneType int

const (
	// ZoneTypeZone represents zone type `Zone`
	ZoneTypeZone ZoneType = iota
)

// MarshalJSON implements the `json.Marshaler` interface
func (k ZoneType) MarshalJSON() ([]byte, error) {
	switch k {
	case ZoneTypeZone:
		return []byte(`"Zone"`), nil
	default:
		return nil, fmt.Errorf("unsupported zone type: %d", k)
	}
}

// UnmarshalJSON implements the `json.Unmarshaler` interface
func (k *ZoneType) UnmarshalJSON(input []byte) error {
	switch string(input) {
	case `"Zone"`:
		*k = ZoneTypeZone
	default:
		return fmt.Errorf("unsupported zone type: %s", string(input))
	}

	return nil
}
