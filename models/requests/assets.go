package requests

import (
	"encoding/json"

	"github.com/timoth-y/chainmetric-core/models"
	"github.com/timoth-y/chainmetric-core/utils"
)

// AssetsQuery defines the structure of query for models.Asset.
type AssetsQuery struct {
	IDs      []string       `json:"type,omitempty"`
	Type     *string        `json:"type,omitempty"`
	Holder   *string        `json:"holder,omitempty"`
	State    *string        `json:"state,omitempty"`
	Location *LocationQuery `json:"location,omitempty"`
	Tag      *string        `json:"tag,omitempty"`
	Limit    int32          `json:"limit,omitempty"`
	ScrollID string         `json:"scroll_id,omitempty"`
}

// Satisfies checks whether the models.Asset satisfies given AssetsQuery.
func (q *AssetsQuery) Satisfies(asset *models.Asset) bool {
	if len(q.IDs) != 0 && !utils.ContainsString(asset.ID, q.IDs) {
		return false
	}
	if q.Type != nil && asset.Type != *q.Type {
		return false
	}
	if q.Holder != nil && asset.Holder != *q.Holder {
		return false
	}
	if q.Location != nil && asset.Location.IsNearBy(q.Location.GeoPoint, q.Location.Distance) {
		return false
	}
	if q.State != nil && asset.State != *q.State {
		return false
	}
	if q.Tag != nil && !utils.ContainsString(*q.Tag, asset.Tags) {
		return false
	}

	return true
}

// Decode deserializes DevicesQuery model.
func (q AssetsQuery) Encode() []byte {
	data, err := json.Marshal(q); if err != nil {
		return nil
	}
	return data
}

// Decode deserializes DevicesQuery model.
func (q AssetsQuery) Decode(b []byte) (*AssetsQuery, error) {
	err := json.Unmarshal(b, &q)
	return &q, err
}
