package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// ScheduleFull The Full Schedule Object is identical to the Schedule Summary Object. See above for details.
// swagger:model ScheduleFull
type ScheduleFull struct {

	// Integer Schedule ID. Read-only.
	ID int64 `json:"id,omitempty"`

	// Name
	Name string `json:"name,omitempty"`
}

// Validate validates this schedule full
func (m *ScheduleFull) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
