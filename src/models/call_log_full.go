package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// CallLogFull The Full Call Log Object includes the properties in the Call Log Summary Object, along with the following:
// swagger:model CallLogFull
type CallLogFull struct {

	// Call duration in seconds
	CallDuration int64 `json:"call_duration,omitempty"`

	// Internal system call reference number
	CallNumber string `json:"call_number,omitempty"`

	// URL of call recording if available. Empty string if call recording does not exist
	CallRecording string `json:"call_recording,omitempty"`

	// Call made to this phone number
	CalledNumber string `json:"called_number,omitempty"`

	// Internal system caller id / name
	CallerCnam string `json:"caller_cnam,omitempty"`

	// Call made from this phone number
	CallerID string `json:"caller_id,omitempty"`

	// Call log creation time. Same as call end time + time needed to create call log
	CreatedAt string `json:"created_at,omitempty"`

	// A list of call flows from beginning of call to end of call.
	Details []*CallDetails `json:"details"`

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

// Validate validates this call log full
func (m *CallLogFull) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDetails(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateExtension(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CallLogFull) validateDetails(formats strfmt.Registry) error {

	if swag.IsZero(m.Details) { // not required
		return nil
	}

	for i := 0; i < len(m.Details); i++ {

		if swag.IsZero(m.Details[i]) { // not required
			continue
		}

		if m.Details[i] != nil {

			if err := m.Details[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CallLogFull) validateExtension(formats strfmt.Registry) error {

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
