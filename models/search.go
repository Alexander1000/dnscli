package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"text/tabwriter"
)

// ObjectType epresents the object type for which a search should be performed
// Possible object types; according to the PowerDNS documentation, this list is exhaustive.
type ObjectType int

const (
	_                        = iota
	ObjectTypeAll ObjectType = iota
	ObjectTypeZone
	ObjectTypeRecord
	ObjectTypeComment
)

// String makes this type implement fmt.Stringer
func (t ObjectType) String() string {
	switch t {
	case ObjectTypeAll:
		return "all"
	case ObjectTypeZone:
		return "zone"
	case ObjectTypeRecord:
		return "record"
	case ObjectTypeComment:
		return "comment"
	}

	return ""
}

// UnmarshalJSON makes this type implement json.Unmarshaler
func (t *ObjectType) UnmarshalJSON(b []byte) error {
	// convert []byte to int
	objType, err := strconv.Atoi(string(b))
	if err != nil {
		return err
	}
	// convert int to ObjectType
	switch ObjectType(objType) {
	case ObjectTypeAll:
		*t = ObjectTypeAll
	case ObjectTypeZone:
		*t = ObjectTypeZone
	case ObjectTypeRecord:
		*t = ObjectTypeRecord
	case ObjectTypeComment:
		*t = ObjectTypeComment
	default:
		return fmt.Errorf(`unknown search type: %s'`, string(b))
	}

	return nil
}

// SearchResult represents a single search result.
type SearchResult struct {
	Content    string     `json:"content"`
	Disabled   bool       `json:"disabled"`
	Name       string     `json:"name"`
	ObjectType ObjectType `json:"object_type"`
	ZoneID     string     `json:"zone_id"`
	Zone       string     `json:"zone"`
	Type       string     `json:"type"`
	TTL        int        `json:"ttl"`
}

// SearchResults represents a list of search results.
type SearchResults []SearchResult

// JSON returns SearchResults as a JSON string
func (r SearchResults) JSON() string {
	j, err := json.Marshal(r)
	if err != nil {
		return ""
	}
	return string(j)
}

// PrettyString returns SearchResults as a pretty formatted string
func (searchReults SearchResults) PrettyString() string {
	// Sorting by Content field
	sort.Slice(searchReults, func(i, j int) bool {
		return searchReults[i].Content < searchReults[j].Content
	})

	buffer := new(bytes.Buffer)

	w := tabwriter.NewWriter(buffer, minwidth, tabwidth, padding, padchar, tabwriter.TabIndent)
	fmt.Fprintf(w, "content\tdisabled\tname\tobject_type\tzone_id\tzone\ttype\tttl\n")
	fmt.Fprintf(w, "-------\t--------\t----\t-----------\t-------\t----\t----\t---\n")
	for _, r := range searchReults {
		fmt.Fprintf(w, "%s\t%t\t%s\t%s\t%s\t%s\t%s\t%d\n",
			r.Content,
			r.Disabled,
			r.Name,
			r.ObjectType.String(),
			r.ZoneID,
			r.Zone,
			r.Type,
			r.TTL,
		)
	}
	w.Flush()

	return buffer.String()
}
