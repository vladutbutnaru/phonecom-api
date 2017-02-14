package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// CallDetails Each item in the 'details' record may contain the following properties:
// swagger:model CallDetails
type CallDetails struct {

	// ID associated with this call item endpoint (type)
	IDValue int64 `json:"id_value,omitempty"`

	// UNIX time stamp representing the UTC time that this call item starts
	StartTime int64 `json:"start_time,omitempty"`

	// Endpoint for this call item, such as call, voicemail, fax, menu, menu item, queue ...
	Type string `json:"type,omitempty"`

	// Same as account id
	VoipID int64 `json:"voip_id,omitempty"`

	// Same as account extension id
	VoipPhoneID int64 `json:"voip_phone_id,omitempty"`
}

// Validate validates this call details
func (m *CallDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
