package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RuleSetForwardItem Array of Forward Item Objects. See below for details. Required.
// swagger:model RuleSetForwardItem
type RuleSetForwardItem struct {

	// Optional. Must equal calling_number or called_number. Defines which phone number should be used for Caller ID. Default is calling_number.
	// Pattern: calling_number|called_number
	CallerID *string `json:"caller_id,omitempty"`

	// Optional. Must equal one of: DEFAULT, STYLE_2, STYLE_3, STYLE_4, STYLE_5, STYLE_6, STYLE_7, STYLE_8, or STYLE_9. Identifies the style of ring tone you will hear when an incoming call is waiting.
	// Pattern: DEFAULT|STYLE_2|STYLE_3|STYLE_4|STYLE_5|STYLE_6|STYLE_7|STYLE_8|STYLE_9
	DistinctiveRing string `json:"distinctive_ring,omitempty"`

	// Required if type = "extension". Extension that callers should be directed to. Output is an Extension Summary Object. Input must be an Extension Lookup Object.
	Extension *ExtensionSummary `json:"extension,omitempty"`

	// Required if type = "phone_number". Phone number that callers should be directed to. Must be a string in E.164 format.
	Number string `json:"number,omitempty"`

	// Boolean. Optional. Default is FALSE. Use this to activate call screening. If TRUE, the timeout on the parent action should be at least 30 seconds.
	Screening *bool `json:"screening,omitempty"`

	// Required. Must equal phone_number or extension.
	// Pattern: phone_number|extension
	Type string `json:"type,omitempty"`

	// Optional string. If screening = TRUE, this value will be passed into our Text-To-Speech engine and used to inform the caller of who they have reached.
	VoiceTag string `json:"voice_tag,omitempty"`
}

// Validate validates this rule set forward item
func (m *RuleSetForwardItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCallerID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDistinctiveRing(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateExtension(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RuleSetForwardItem) validateCallerID(formats strfmt.Registry) error {

	if swag.IsZero(m.CallerID) { // not required
		return nil
	}

	if err := validate.Pattern("caller_id", "body", string(*m.CallerID), `calling_number|called_number`); err != nil {
		return err
	}

	return nil
}

func (m *RuleSetForwardItem) validateDistinctiveRing(formats strfmt.Registry) error {

	if swag.IsZero(m.DistinctiveRing) { // not required
		return nil
	}

	if err := validate.Pattern("distinctive_ring", "body", string(m.DistinctiveRing), `DEFAULT|STYLE_2|STYLE_3|STYLE_4|STYLE_5|STYLE_6|STYLE_7|STYLE_8|STYLE_9`); err != nil {
		return err
	}

	return nil
}

func (m *RuleSetForwardItem) validateExtension(formats strfmt.Registry) error {

	if swag.IsZero(m.Extension) { // not required
		return nil
	}

	if m.Extension != nil {

		if err := m.Extension.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("extension")
			}
			return err
		}
	}

	return nil
}

func (m *RuleSetForwardItem) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := validate.Pattern("type", "body", string(m.Type), `phone_number|extension`); err != nil {
		return err
	}

	return nil
}