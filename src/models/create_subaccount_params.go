package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreateSubaccountParams create subaccount params
// swagger:model CreateSubaccountParams
type CreateSubaccountParams struct {

	// Contact Object for billing purposes. See below for details.
	BillingContact *ContactSubaccount `json:"billing_contact,omitempty"`

	// Contact Object. See below for details.
	Contact *ContactSubaccount `json:"contact,omitempty"`

	// Sub account password
	// Required: true
	Password *string `json:"password"`

	// Sub account user name
	// Required: true
	Username *string `json:"username"`
}

// Validate validates this create subaccount params
func (m *CreateSubaccountParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBillingContact(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateContact(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePassword(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateUsername(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateSubaccountParams) validateBillingContact(formats strfmt.Registry) error {

	if swag.IsZero(m.BillingContact) { // not required
		return nil
	}

	if m.BillingContact != nil {

		if err := m.BillingContact.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("billing_contact")
			}
			return err
		}
	}

	return nil
}

func (m *CreateSubaccountParams) validateContact(formats strfmt.Registry) error {

	if swag.IsZero(m.Contact) { // not required
		return nil
	}

	if m.Contact != nil {

		if err := m.Contact.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("contact")
			}
			return err
		}
	}

	return nil
}

func (m *CreateSubaccountParams) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("password", "body", m.Password); err != nil {
		return err
	}

	return nil
}

func (m *CreateSubaccountParams) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("username", "body", m.Username); err != nil {
		return err
	}

	return nil
}
