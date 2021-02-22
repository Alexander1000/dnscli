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

// ResourceRecordSet represents a Resource Record Set (all records with the same name and type).
type ResourceRecordSet struct {
	Name       string              `json:"name"`
	Type       string              `json:"type"`
	TTL        int                 `json:"ttl"`
	ChangeType RecordSetChangeType `json:"changetype,omitempty"`
	Records    []Record            `json:"records"`
	Comments   []Comment           `json:"comments"`
}

// Record represents a single record
type Record struct {
	Content  string `json:"content"`
	Disabled bool   `json:"disabled"`
	SetPTR   bool   `json:"set-ptr,omitempty"`
}

// Comment a comment about an Resource Record Set
type Comment struct {
	Content    string `json:"content"`
	Account    string `json:"account"`
	ModifiedAt int    `json:"modified_at"`
}
