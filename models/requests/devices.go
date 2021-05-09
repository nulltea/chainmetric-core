package requests

import (
	"encoding/json"

	"github.com/timoth-y/chainmetric-core/models"
)

// DeviceUpdateRequest defines update request for models.Device.
type DeviceUpdateRequest struct {
	Name     *string             `json:"name,omitempty"`
	IP       *string              `json:"ip"`
	Hostname *string              `json:"hostname"`
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
