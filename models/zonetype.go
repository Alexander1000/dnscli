package models

import "fmt"

// ZoneType zone type
type ZoneType int

const (
	// ZoneTypeZone represents zone type `Zone`
	ZoneTypeZone ZoneType = iota
)

// MarshalJSON implements the `json.Marshaler` interface
func (k ZoneType) MarshalJSON() ([]byte, error) {
	switch k {
	case ZoneTypeZone:
		return []byte(`"Zone"`), nil
	default:
		return nil, fmt.Errorf("unsupported zone type: %d", k)
	}
}

// UnmarshalJSON implements the `json.Unmarshaler` interface
func (k *ZoneType) UnmarshalJSON(input []byte) error {
	switch string(input) {
	case `"Zone"`:
		*k = ZoneTypeZone
	default:
		return fmt.Errorf("unsupported zone type: %s", string(input))
	}

	return nil
}
