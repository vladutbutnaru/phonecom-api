package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// GroupSummary The Group Summary Object is used to briefly represent a contact group. It can occur in several places throughout this API. Here are the properties:
// swagger:model GroupSummary
type GroupSummary struct {

	// Integer ID. Read-only.
	ID int64 `json:"id,omitempty"`

	// Name
	Name string `json:"name,omitempty"`
}

// Validate validates this group summary
func (m *GroupSummary) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}