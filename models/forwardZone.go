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
	Name        string `json:"name"`
	Nameservers ZoneNameservers
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
