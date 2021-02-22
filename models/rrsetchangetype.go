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

// RecordSetChangeType MUST be added when updating the RRSet. Must be REPLACE or DELETE.
// With DELETE, all existing RRs matching name and type will be deleted, including all comments.
// With REPLACE: when records is present, all existing RRs matching name and type will be deleted,
// and then new records given in records will be created. If no records are left, any existing
// comments will be deleted as well. When comments is present, all existing comments for the RRs
// matching name and type will be deleted, and then new comments given in comments will be created.
type RecordSetChangeType int

const (
	_ = iota
	// ChangeTypeDelete represents DELETE RecordSetChangeType
	ChangeTypeDelete RecordSetChangeType = iota
	// ChangeTypeReplace represents REPLACE RecordSetChangeType
	ChangeTypeReplace
)

// MarshalJSON implements the `json.Marshaler` interface
func (k RecordSetChangeType) MarshalJSON() ([]byte, error) {
	switch k {
	case ChangeTypeDelete:
		return []byte(`"DELETE"`), nil
	case ChangeTypeReplace:
		return []byte(`"REPLACE"`), nil
	default:
		return nil, fmt.Errorf("unsupported change type: %d", k)
	}
}

// UnmarshalJSON implements the `json.Unmarshaler` interface
func (k *RecordSetChangeType) UnmarshalJSON(input []byte) error {
	switch string(input) {
	case `"DELETE"`:
		*k = ChangeTypeDelete
	case `"REPLACE"`:
		*k = ChangeTypeReplace
	default:
		return fmt.Errorf("unsupported change type: %s", string(input))
	}

	return nil
}
