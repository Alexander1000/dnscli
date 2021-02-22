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

// ZoneKind DNS zone kind, one of `Native`, `Master`, `Slave`
type ZoneKind int

const (
	_ = iota
	// ZoneKindNative represenrs `Native` zone kind
	ZoneKindNative ZoneKind = iota
	// ZoneKindMaster represenrs `Master` zone kind
	ZoneKindMaster
	// ZoneKindSlave represenrs `Slave` zone kind
	ZoneKindSlave
)

// MarshalJSON implements the `json.Marshaler` interface
func (k ZoneKind) MarshalJSON() ([]byte, error) {
	switch k {
	case ZoneKindNative:
		return []byte(`"Native"`), nil
	case ZoneKindMaster:
		return []byte(`"Master"`), nil
	case ZoneKindSlave:
		return []byte(`"Slave"`), nil
	default:
		return nil, fmt.Errorf("unsupported zone kind: %d", k)
	}
}

// UnmarshalJSON implements the `json.Unmarshaler` interface
func (k *ZoneKind) UnmarshalJSON(input []byte) error {
	switch string(input) {
	case `"Native"`:
		*k = ZoneKindNative
	case `"Master"`:
		*k = ZoneKindMaster
	case `"Slave"`:
		*k = ZoneKindSlave
	default:
		return fmt.Errorf("unsupported zone kind: %s", string(input))
	}

	return nil
}
