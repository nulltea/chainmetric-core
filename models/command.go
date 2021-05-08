package models

import (
	"encoding/json"
	"time"
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
	Args      []interface{}       `json:"args"`
	Status    DeviceCommandStatus `json:"status"`
	Error     string              `json:"error"`
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

// Encode serializes the DeviceCommandLogEntry model.
func (m DeviceCommandLogEntry) Encode() []byte {
	data, err := json.Marshal(m); if err != nil {
		return nil
	}
	return data
}

// Decode deserializes the DeviceCommandLogEntry model.
func (m DeviceCommandLogEntry) Decode(b []byte) (*DeviceCommandLogEntry, error) {
	err := json.Unmarshal(b, &m)
	return &m, err
}
