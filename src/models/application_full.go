package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// ApplicationFull The Full Application Object is identical to the Application Summary Object. See above for details.
// swagger:model ApplicationFull
type ApplicationFull struct {

	// Application ID. Read-only.
	ID int64 `json:"id,omitempty"`

	// Application name
	Name string `json:"name,omitempty"`
}

// Validate validates this application full
func (m *ApplicationFull) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
