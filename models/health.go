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

// HealthResult represents the health info result
type HealthResult struct {
	Alive    bool   `json:"alive"`
	Hostname string `json:"hostname"`
}

// JSON returns HealthResult as a JSON string
func (h HealthResult) JSON() string {
	j, err := json.Marshal(h)
	if err != nil {
		return ""
	}
	return string(j)
}

// PrettyString returns the HealthResult as a pretty formatted string
func (h HealthResult) PrettyString() string {
	buffer := new(bytes.Buffer)
	w := tabwriter.NewWriter(buffer, minwidth, tabwidth, padding, padchar, tabwriter.TabIndent)

	fmt.Fprintf(w, "alive\thostname\n")
	fmt.Fprintf(w, "-----\t--------\n")
	fmt.Fprintf(w, "%t\t%s\n", h.Alive, h.Hostname)
	w.Flush()

	return buffer.String()
}
