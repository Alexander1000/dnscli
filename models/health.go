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
