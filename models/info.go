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
