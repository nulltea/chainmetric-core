package requests

import (
	"encoding/json"

	"github.com/timoth-y/chainmetric-core/models"
)

// DeviceUpdateRequest defines update request for models.Device.
type DeviceUpdateRequest struct {
	Name     *string                `json:"name,omitempty"`
	Profile  *string                `json:"profile,omitempty"`
	Supports models.Metrics         `json:"supports,omitempty"`
	Holder   *string                `json:"holder,omitempty"`
	State    *models.DeviceState    `json:"state,omitempty"`
	Location *string                `json:"location,omitempty"`
}

// DeviceUpdateRequest defines request for models.DeviceCommand execution.
type DeviceCommandRequest struct {
	DeviceID string               `json:"device_id"`
	Command  models.DeviceCommand `json:"command"`
	Args     []interface{}        `json:"args,omitempty"`
}

// Update updates models.Device
func (u *DeviceUpdateRequest) Update(device *models.Device) {
	if u.Name != nil {
		device.Name = *u.Name
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

func (u DeviceUpdateRequest) Encode() []byte {
	data, err := json.Marshal(u); if err != nil {
		return nil
	}
	return data
}

func (u DeviceUpdateRequest) Decode(b []byte) (*DeviceUpdateRequest, error) {
	err := json.Unmarshal(b, &u)
	return &u, err
}
