package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"text/tabwriter"
)

// ZoneNameservers is a special list type to represent the nameservers of a zone.
// When nil, this type will still serialize to an empty JSON list.
// See https://github.com/mittwald/go-powerdns/issues/4 for more information
type ZoneNameservers []string

// MarshalJSON implements the `json.Marshaler` interface
func (z ZoneNameservers) MarshalJSON() ([]byte, error) {
	if z == nil {
		return []byte("[]"), nil
	}

	return json.Marshal([]string(z))
}

// Zone represents an authoritative DNS Zone
type Zone struct {
	ID                 string              `json:"id,omitempty"`
	Name               string              `json:"name"`
	Type               ZoneType            `json:"type"`
	URL                string              `json:"url,omitempty"`
	Kind               ZoneKind            `json:"kind,omitempty"`
	ResourceRecordSets []ResourceRecordSet `json:"rrsets,omitempty"`
	Serial             int                 `json:"serial,omitempty"`
	NotifiedSerial     int                 `json:"notified_serial,omitempty"`
	Masters            []string            `json:"masters,omitempty"`
	DNSSec             bool                `json:"dnssec,omitempty"`
	NSec3Param         string              `json:"nsec3param,omitempty"`
	NSec3Narrow        bool                `json:"nsec3narrow,omitempty"`
	Presigned          bool                `json:"presigned,omitempty"`
	SOAEdit            string              `json:"soa_edit,omitempty"`
	SOAEditAPI         string              `json:"soa_edit_api,omitempty"`
	APIRectify         bool                `json:"api_rectify,omitempty"`
	Zone               string              `json:"zone,omitempty"`
	Account            string              `json:"account,omitempty"`
	Nameservers        ZoneNameservers     `json:"nameservers"`
	TSIGMasterKeyIDs   []string            `json:"tsig_master_key_ids,omitempty"`
	TSIGSlaveKeyIDs    []string            `json:"tsig_slave_key_ids,omitempty"`
}

// JSON returns zone as a JSON string
func (z Zone) JSON() string {
	j, err := json.Marshal(z)
	if err != nil {
		return ""
	}
	return string(j)
}

// PrettyString returns zone as a pretty formatted string
func (z Zone) PrettyString() string {
	return fmt.Sprintf("%s: %s\n", z.Name, strings.Join(z.Nameservers, ", "))
}

// Zones represents the list of an authoritative DNS Zones
type Zones []*Zone

// JSON returs the list of zone as a JSON string
func (zones Zones) JSON() string {
	j, err := json.Marshal(zones)
	if err != nil {
		return ""
	}
	return string(j)
}

// PrettyString returns the list of zones as a pretty formatted string
func (zones Zones) PrettyString() string {
	// Sorting Zones by Name
	sort.Slice(zones, func(i, j int) bool {
		return zones[i].Name < zones[j].Name
	})

	buffer := new(bytes.Buffer)

	w := tabwriter.NewWriter(buffer, minwidth, tabwidth, padding, padchar, tabwriter.TabIndent)
	fmt.Fprintf(w, "zone\tserial\n")
	fmt.Fprintf(w, "----\t------\n")
	for _, z := range zones {
		// fmt.Fprintf(w, "%s\t%s\n", z.Name, strings.Join(z.Nameservers, ", "))
		fmt.Fprintf(w, "%s\t%d\n", DeCanonicalize(z.Name), z.Serial)
	}
	w.Flush()

	return buffer.String()
}
