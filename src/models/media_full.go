package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MediaFull The Full Recording Object includes all of the properties from the Recording Summary Object, along with the following:
// swagger:model MediaFull
type MediaFull struct {

	// Recording ID. Read-only.
	ID int64 `json:"id,omitempty"`

	// Name of recording
	Name string `json:"name,omitempty"`

	// Can be hold_music or greeting. Indicates the purpose of this recording and where it can be used.
	// Pattern: hold_music|greeting
	Type string `json:"type,omitempty"`
}

// Validate validates this media full
func (m *MediaFull) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MediaFull) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := validate.Pattern("type", "body", string(m.Type), `hold_music|greeting`); err != nil {
		return err
	}

	return nil
}
