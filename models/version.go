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
