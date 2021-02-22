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

// VersionInfo represents version info
type VersionInfo struct {
	Version string `json:"version"`
	Build   string `json:"build"`
}

// JSON returns VersionInfo as a JSON string
func (v VersionInfo) JSON() string {
	j, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return string(j)
}

// PrettyString returns the VersionInfo as a pretty formatted string
func (v VersionInfo) PrettyString() string {
	buffer := new(bytes.Buffer)
	w := tabwriter.NewWriter(buffer, minwidth, tabwidth, padding, padchar, tabwriter.TabIndent)

	fmt.Fprintf(w, "version\tbuild\n")
	fmt.Fprintf(w, "------\t-----\n")
	fmt.Fprintf(w, "%s\t%s\n", v.Version, v.Build)
	w.Flush()

	return buffer.String()
}
