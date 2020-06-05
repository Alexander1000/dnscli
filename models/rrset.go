package models

// ResourceRecordSet represents a Resource Record Set (all records with the same name and type).
type ResourceRecordSet struct {
	Name       string              `json:"name"`
	Type       string              `json:"type"`
	TTL        int                 `json:"ttl"`
	ChangeType RecordSetChangeType `json:"changetype,omitempty"`
	Records    []Record            `json:"records"`
	Comments   []Comment           `json:"comments"`
}

// Record represents a single record
type Record struct {
	Content  string `json:"content"`
	Disabled bool   `json:"disabled"`
	SetPTR   bool   `json:"set-ptr,omitempty"`
}

// Comment a comment about an Resource Record Set
type Comment struct {
	Content    string `json:"content"`
	Account    string `json:"account"`
	ModifiedAt int    `json:"modified_at"`
}
