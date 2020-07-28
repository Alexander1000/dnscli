package models

import (
	"encoding/json"
	"fmt"
)

// FlushResult
type FlushResult struct {
	Count  int    `json:"count"`
	Result string `json:"result"`
}

// JSON returns FlushResult as a JSON string
func (fr FlushResult) JSON() string {
	j, err := json.Marshal(fr)
	if err != nil {
		return ""
	}
	return string(j)
}

// PrettyString returns FlushResult as a pretty formatted string
func (fr FlushResult) PrettyString() string {
	return fmt.Sprintf("wiped %d records\n", fr.Count)
}
