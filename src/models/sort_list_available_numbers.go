package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// SortListAvailableNumbers sort list available numbers
// swagger:model SortListAvailableNumbers
type SortListAvailableNumbers struct {

	// internal
	Internal string `json:"internal,omitempty"`

	// phone number
	PhoneNumber string `json:"phone_number,omitempty"`

	// price
	Price string `json:"price,omitempty"`
}

// Validate validates this sort list available numbers
func (m *SortListAvailableNumbers) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}