package requests

import (
	"encoding/json"
	"time"

	"github.com/pkg/errors"

	"github.com/timoth-y/chainmetric-core/models"
)

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

// DeviceCommandResultsSubmitRequest defines structure of the request to submit
// command execution results on the given device.
type DeviceCommandResultsSubmitRequest struct {
	Status    models.DeviceCommandStatus `json:"status"`
	Error     *string                     `json:"error,omitempty"`
	Timestamp time.Time                  `json:"timestamp"`
}

// Apply validates and applies DeviceCommandResultsSubmitRequest on models.DeviceCommandLogEntry record.
func (c DeviceCommandResultsSubmitRequest) Apply(entry *models.DeviceCommandLogEntry) error {
	switch c.Status {
	case models.DeviceCmdCompleted,
		models.DeviceCmdProcessing,
		models.DeviceCmdFailed:
	default:
		return errors.Errorf("command status '%s' is unknown", c.Status)
	}

	entry.Status = c.Status
	entry.Timestamp = c.Timestamp
	if entry.Status == models.DeviceCmdFailed && c.Error != nil && len(*c.Error) != 0 {
		entry.Error = *c.Error
	}

	return nil
}

// Encode serialises DeviceCommandResultsSubmitRequest model.
func (c DeviceCommandResultsSubmitRequest) Encode() []byte {
	data, err := json.Marshal(c); if err != nil {
		return nil
	}
	return data
}

// Decode deserializes DeviceCommandResultsSubmitRequest model.
func (c DeviceCommandResultsSubmitRequest) Decode(b []byte) (*DeviceCommandResultsSubmitRequest, error) {
	err := json.Unmarshal(b, &c)
	return &c, err
}
