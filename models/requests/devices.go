package requests

import (
	"encoding/json"

	"github.com/timoth-y/chainmetric-core/models"
	"github.com/timoth-y/chainmetric-core/utils"
)

// DeviceUpdateRequest defines update request for models.Device.
type DeviceUpdateRequest struct {
	Name     *string             `json:"name,omitempty"`
	IP       *string             `json:"ip"`
	Hostname *string             `json:"hostname"`
	Profile  *string             `json:"profile,omitempty"`
	Supports models.Metrics      `json:"supports,omitempty"`
	Holder   *string             `json:"holder,omitempty"`
	State    *models.DeviceState `json:"state,omitempty"`
	Location *models.Location    `json:"location,omitempty"`
}

// Update updates models.Device
func (u *DeviceUpdateRequest) Update(device *models.Device) {
	if u.Name != nil {
		device.Name = *u.Name
	}

	if u.IP != nil {
		device.IP = *u.IP
	}

	if u.Hostname != nil {
		device.Hostname = *u.Hostname
	}

	if u.Profile != nil {
		device.Profile = *u.Profile
	}

	if u.Supports != nil {
		device.Supports = u.Supports
	}

	if u.Holder != nil {
		device.Holder = *u.Holder
	}

	if u.State != nil {
		device.State = *u.State
	}

	if u.Location != nil {
		device.Location = *u.Location
	}
}

// Encode serializes DeviceUpdateRequest model.
func (u DeviceUpdateRequest) Encode() []byte {
	data, err := json.Marshal(u); if err != nil {
		return nil
	}
	return data
}

// Decode deserializes DeviceUpdateRequest model.
func (u DeviceUpdateRequest) Decode(b []byte) (*DeviceUpdateRequest, error) {
	err := json.Unmarshal(b, &u)
	return &u, err
}

type DevicesQuery struct {
	IDs      []string            `json:"type,omitempty"`
	Supports models.Metrics      `json:"supports"`
	Holder   *string             `json:"holder,omitempty"`
	State    *models.DeviceState `json:"state,omitempty"`
	Location *models.Location    `json:"location,omitempty"`
}

// Satisfies checks whether the models.Device satisfies given DevicesQuery.
func (q *DevicesQuery) Satisfies(dev *models.Device) bool {
	if len(q.IDs) != 0 && !utils.ContainsString(dev.ID, q.IDs) {
		return false
	}
	if q.Holder != nil && dev.Holder != *q.Holder {
		return false
	}
	if q.Location != nil && dev.Location != *q.Location {
		return false
	}
	if q.State != nil && dev.State != *q.State {
		return false
	}

	return true
}

// Decode deserializes DevicesQuery model.
func (q DevicesQuery) Encode() []byte {
	data, err := json.Marshal(q); if err != nil {
		return nil
	}
	return data
}

// Decode deserializes DevicesQuery model.
func (q DevicesQuery) Decode(b []byte) (*DevicesQuery, error) {
	err := json.Unmarshal(b, &q)
	return &q, err
}
