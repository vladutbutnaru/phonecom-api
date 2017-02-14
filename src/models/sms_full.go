package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SmsFull The Full SMS Object includes all of the properties in the SMS Summary Object, along with several more.
// swagger:model SmsFull
type SmsFull struct {

	// Date string representing the UTC time that the object was created in the Phone.com API system.
	// Required: true
	CreatedAt *strfmt.DateTime `json:"created_at"`

	// Unix time stamp representing the UTC time that the object was created in the Phone.com API system.
	// Required: true
	CreatedEpoch *int64 `json:"created_epoch"`

	// Direction of SMS. 'in' for Incoming messages, 'out' for Outgoing messages.
	// Required: true
	Direction *string `json:"direction"`

	// Caller ID number to display on the incoming/outgoing SMS message. For an outgoing message, it must be a phone number associated with your Phone.com account.
	// Required: true
	From *string `json:"from"`

	// Unique SMS ID. Read-only.
	// Required: true
	ID *string `json:"id"`

	// Body of the SMS text
	// Required: true
	Text *string `json:"text"`

	// An array of SMS recipients.
	// Required: true
	To []*Recipient `json:"to"`
}

// Validate validates this sms full
func (m *SmsFull) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCreatedEpoch(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDirection(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateFrom(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateText(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTo(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SmsFull) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("created_at", "body", m.CreatedAt); err != nil {
		return err
	}

	return nil
}

func (m *SmsFull) validateCreatedEpoch(formats strfmt.Registry) error {

	if err := validate.Required("created_epoch", "body", m.CreatedEpoch); err != nil {
		return err
	}

	return nil
}

func (m *SmsFull) validateDirection(formats strfmt.Registry) error {

	if err := validate.Required("direction", "body", m.Direction); err != nil {
		return err
	}

	return nil
}

func (m *SmsFull) validateFrom(formats strfmt.Registry) error {

	if err := validate.Required("from", "body", m.From); err != nil {
		return err
	}

	return nil
}

func (m *SmsFull) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *SmsFull) validateText(formats strfmt.Registry) error {

	if err := validate.Required("text", "body", m.Text); err != nil {
		return err
	}

	return nil
}

func (m *SmsFull) validateTo(formats strfmt.Registry) error {

	if err := validate.Required("to", "body", m.To); err != nil {
		return err
	}

	for i := 0; i < len(m.To); i++ {

		if swag.IsZero(m.To[i]) { // not required
			continue
		}

		if m.To[i] != nil {

			if err := m.To[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("to" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}