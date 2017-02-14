package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// SortCallLogs sort call logs
// swagger:model SortCallLogs
type SortCallLogs struct {

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// start time
	StartTime string `json:"start_time,omitempty"`
}

// Validate validates this sort call logs
func (m *SortCallLogs) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}