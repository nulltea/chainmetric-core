package models

import "encoding/json"

// Asset defines asset data models.
type Asset struct {
	ID       string   `json:"id,omitempty"`
	SKU      string   `json:"sku,omitempty"`
	Name     string   `json:"name,omitempty"`
	Type     string   `json:"type,omitempty"`
	Info     string   `json:"info,omitempty"`
	Cost     float64  `json:"cost,omitempty"`
	Amount   int      `json:"amount,omitempty"`
	Holder   string   `json:"holder,omitempty"`
	State    string   `json:"state,omitempty" metadata:",optional"`
	Location Location `json:"location,omitempty"`
	Tags     []string `json:"tags,omitempty" metadata:",optional"`
}

// Encode serializes the Asset model.
func (m Asset) Encode() []byte {
	data, err := json.Marshal(m); if err != nil {
		return nil
	}
	return data
}

// Decode deserializes the Asset model.
func (m Asset) Decode(b []byte) (*Asset, error) {
	err := json.Unmarshal(b, &m)
	return &m, err
}
