package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// CreateQueueParams create queue params
// swagger:model CreateQueueParams
type CreateQueueParams struct {

	// Type of caller id
	CallerIDType string `json:"caller_id_type,omitempty"`

	// Recording lookup object
	Greeting interface{} `json:"greeting,omitempty"`

	// Recording lookup object
	HoldMusic interface{} `json:"hold_music,omitempty"`

	// Max seconds for hold
	MaxHoldTime int64 `json:"max_hold_time,omitempty"`

	// Extensions or phone numbers
	Members []interface{} `json:"members"`

	// Name of queue
	Name string `json:"name,omitempty"`

	// Number of seconds to ring each member
	RingTime int64 `json:"ring_time,omitempty"`
}

// Validate validates this create queue params
func (m *CreateQueueParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMembers(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateQueueParams) validateMembers(formats strfmt.Registry) error {

	if swag.IsZero(m.Members) { // not required
		return nil
	}

	for i := 0; i < len(m.Members); i++ {

	}

	return nil
}
