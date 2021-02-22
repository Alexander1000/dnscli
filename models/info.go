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

import (
	"bytes"
	"encoding/json"
	"fmt"
	"text/tabwriter"
)

// InfoResult represents the info result
type InfoResult struct {
	Version string `json:"version"`
	Commit  string `json:"commit"`
	Go      string `json:"go"`
}

// JSON returns InfoResult as a JSON string
func (i InfoResult) JSON() string {
	j, err := json.Marshal(i)
	if err != nil {
		return ""
	}
	return string(j)
}

// PrettyString returns the InfoResult as a pretty formatted string
func (i InfoResult) PrettyString() string {
	buffer := new(bytes.Buffer)
	w := tabwriter.NewWriter(buffer, minwidth, tabwidth, padding, padchar, tabwriter.TabIndent)

	fmt.Fprintf(w, "version\tcommit\tgo\n")
	fmt.Fprintf(w, "-------\t------\t--\n")
	fmt.Fprintf(w, "%s\t%s\t%s\n", i.Version, i.Commit, i.Go)
	w.Flush()

	return buffer.String()
}
