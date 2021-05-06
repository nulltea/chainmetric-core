package requests

import (
	"encoding/json"

	"github.com/pkg/errors"

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

// DeviceUpdateRequest defines request for models.DeviceCommand execution.
type DeviceCommandRequest struct {
	DeviceID  string               `json:"device_id"`
	Command   models.DeviceCommand `json:"command"`
	Args      []interface{}        `json:"args,omitempty"`
}

// DeviceCommandEventPayload embeds DeviceCommandRequest with execution log entry id.
// so that it could be passed internally between Smart Contract and device via event.
type DeviceCommandEventPayload struct {
	ID string `json:"id"`
	DeviceCommandRequest
}

// Encode serialises DeviceCommandRequest model.
func (c DeviceCommandRequest) Encode() []byte {
	data, err := json.Marshal(c); if err != nil {
		return nil
	}
	return data
}

// Decode deserializes DeviceCommandRequest model.
func (c DeviceCommandRequest) Decode(b []byte) (*DeviceCommandRequest, error) {
	err := json.Unmarshal(b, &c)
	return &c, err
}

// Encode serialises DeviceCommandEventPayload model.
func (c DeviceCommandEventPayload) Encode() []byte {
	data, err := json.Marshal(c); if err != nil {
		return nil
	}
	return data
}

// Decode deserializes DeviceCommandEventPayload model.
func (c DeviceCommandEventPayload) Decode(b []byte) (*DeviceCommandEventPayload, error) {
	err := json.Unmarshal(b, &c)
	return &c, err
}

// Validate validates DeviceCommandRequest model.
func (c DeviceCommandRequest) Validate() error {
	if len(c.DeviceID) == 0 {
		return errors.New("device id is required and must be provided")
	}

	switch c.Command {
	case models.DevicePauseCmd,
		models.DeviceResumeCmd,
		models.DevicePairingCmd:
	default:
		return errors.Errorf("command '%s' is not supported", c.Command)
	}

	return nil
}
