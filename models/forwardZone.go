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
	"sort"
	"strings"
	"text/tabwriter"
)

// ForwardZone represent forwarding zone for settings in PowerDNS Recursor
type ForwardZone struct {
	Name        string          `json:"name"`
	Nameservers ZoneNameservers `json:"nameservers"`
}

// PrettyString creates a pretty string of the forwarding zone
func (fz ForwardZone) PrettyString() string {
	buffer := new(bytes.Buffer)

	w := tabwriter.NewWriter(buffer, minwidth, tabwidth, padding, padchar, tabwriter.TabIndent)
	fmt.Fprintf(w, "zone\tnameservers\n")
	fmt.Fprintf(w, "----\t-----------\n")
	fmt.Fprintf(w, "%s\t%s\n", DeCanonicalize(fz.Name), strings.Join(fz.Nameservers, ", "))
	w.Flush()

	return buffer.String()
}

// JSON returns the JSON representation of the forwarding zone
func (fz ForwardZone) JSON() string {
	j, err := json.Marshal(fz)
	if err != nil {
		return ""
	}
	return string(j)
}

// ForwardZones represent list of zorwarding zones
type ForwardZones []*ForwardZone

// PrettyString cretes a pretty string of the forwarding zones list
func (fzs ForwardZones) PrettyString() string {
	// Sorting []forwardZone by Name
	sort.Slice(fzs, func(i, j int) bool {
		return fzs[i].Name < fzs[j].Name
	})

	buffer := new(bytes.Buffer)

	w := tabwriter.NewWriter(buffer, minwidth, tabwidth, padding, padchar, tabwriter.TabIndent)
	fmt.Fprintf(w, "zone\tnameservers\n")
	fmt.Fprintf(w, "----\t-----------\n")
	for _, fz := range fzs {
		fmt.Fprintf(w, "%s\t%s\n", DeCanonicalize(fz.Name), strings.Join(fz.Nameservers, ", "))
	}
	w.Flush()

	return buffer.String()
}

// JSON returns the JSON representation of the forwarding zones list
func (fzs ForwardZones) JSON() string {
	j, err := json.Marshal(fzs)
	if err != nil {
		return ""
	}
	return string(j)
}
