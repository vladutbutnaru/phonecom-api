package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// CallLogSummary The Call Log Summary Object is used to briefly represent a call. It consists of properties as below:
// swagger:model CallLogSummary
type CallLogSummary struct {

	// Call duration in seconds
	CallDuration int64 `json:"call_duration,omitempty"`

	// Internal system call reference number
	CallNumber string `json:"call_number,omitempty"`

	// URL of call recording if available. Empty string if call recording does not exist
	CallRecording string `json:"call_recording,omitempty"`

	// Call made to this phone number
	CalledNumber string `json:"called_number,omitempty"`

	// Call made from this phone number
	CallerID string `json:"caller_id,omitempty"`

	// Call log creation time. Same as call end time + time needed to create call log
	CreatedAt string `json:"created_at,omitempty"`

	// Call direction: in or out
	Direction string `json:"direction,omitempty"`

	// Account extension
	Extension *ExtensionSummary `json:"extension,omitempty"`

	// Last action of call flow
	FinalAction string `json:"final_action,omitempty"`

	// ID
	ID string `json:"id,omitempty"`

	// Was call being monitored?
	IsMonitored string `json:"is_monitored,omitempty"`

	// Call start time
	StartTime string `json:"start_time,omitempty"`

	// Call type: call, fax, audiogram ...
	Type string `json:"type,omitempty"`

	// Internal system id, may be null
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this call log summary
func (m *CallLogSummary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExtension(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CallLogSummary) validateExtension(formats strfmt.Registry) error {

	if swag.IsZero(m.Extension) { // not required
		return nil
	}

	if m.Extension != nil {

		if err := m.Extension.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("extension")
			}
			return err
		}
	}

	return nil
}