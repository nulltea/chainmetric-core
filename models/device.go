package models

import (
	"encoding/json"
	"time"

	"github.com/pkg/errors"
)

// Device defines device data models.
type Device struct {
	ID       string      `json:"id"`
	IP       string      `json:"ip"`
	MAC      string      `json:"mac,omitempty" metadata:",optional"`
	Name     string      `json:"name,omitempty" metadata:",optional"`
	Hostname string      `json:"hostname"`
	Profile  string      `json:"profile,omitempty" metadata:",optional"`
	Supports Metrics     `json:"supports"`
	Holder   string      `json:"holder"`
	State    DeviceState `json:"state,omitempty" metadata:",optional"`
	Location string      `json:"location"`
}

// Encode serializes the Device model.
func (m Device) Encode() []byte {
	data, err := json.Marshal(m); if err != nil {
		return nil
	}
	return data
}

// Decode deserializes the Device model.
func (m Device) Decode(b []byte) (*Device, error) {
	err := json.Unmarshal(b, &m)
	return &m, err
}

// Decode validates the Device model.
func (m *Device) Validate() error {
	if len(m.ID) == 0 {
		return errors.New("id must be assigned to device")
	}

	if len(m.Hostname) == 0 {
		return errors.New("hostname must be assigned to device")
	}

	if len(m.Supports) == 0 {
		return errors.New("device must to support at least one metric")
	}

	return nil
}


// DeviceState defines Device states enum.
type DeviceState string

var (
	// DeviceOnline defines Device "online" state.
	DeviceOnline DeviceState = "online"
	// DeviceOffline defines Device "offline" state.
	DeviceOffline DeviceState = "offline"
	// DevicePaused defines Device "paused" state.
	DevicePaused DeviceState = "paused"
)


// DeviceCommand defines Device commands enum.
type DeviceCommand string

const (
	// DevicePauseCmd defines DeviceCommand to "pause" Device.
	DevicePauseCmd DeviceCommand = "pause"
	// DevicePauseCmd defines DeviceCommand to "pause" Device.
	DeviceResumeCmd DeviceCommand = "resume"
	// DevicePairingCmd defines DeviceCommand to pair Device via Bluetooth.
	DevicePairingCmd DeviceCommand = "ble_pair"
)

// DeviceCommandLogEntry defines DeviceCommand execution historical log record.
type DeviceCommandLogEntry struct {
	DeviceID  string              `json:"device_id"`
	Command   DeviceCommand       `json:"command"`
	Args      DeviceCommand       `json:"args"`
	Status    DeviceCommandStatus `json:"status"`
	Error     string              `json:"err"`
	Timestamp time.Time           `json:"timestamp"`
}

// DeviceCommandStatus defines DeviceCommand execution status.
type DeviceCommandStatus string

const (
	// DeviceCmdCompleted defines "completed" status of DeviceCommand execution.
	DeviceCmdCompleted DeviceCommandStatus = "completed"
	// DeviceCmdProcessing defines "processing" status of DeviceCommand execution.
	DeviceCmdProcessing DeviceCommandStatus = "processing"
	// DeviceCmdFailed defines "failed" status of DeviceCommand execution.
	DeviceCmdFailed DeviceCommandStatus = "failed"
)
