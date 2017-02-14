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

// RuleSetAction Filter Object. Optional. See below for details.
// swagger:model RuleSetAction
type RuleSetAction struct {

	// Required.
	// Pattern: directory|disconnect|fax|forward|greeting|hold|menu|queue|trunk|voicemail
	Action string `json:"action,omitempty"`

	// Required. Seconds that the caller should be placed on hold before being moved onto the next action. Must be an integer between 1 and 60 seconds.
	Duration int64 `json:"duration,omitempty"`

	// Extension that this action refers to. Output is an Extension Summary Object. Input must be an Extension Lookup Object. Required.
	Extension *ExtensionSummary `json:"extension,omitempty"`

	// Greeting that this action refers to. Output is a Media Summary Object. Input must be a Media Lookup Object. Required. Must refer to a media recording that has is_hold_music set to FALSE.
	Greeting *MediaSummary `json:"greeting,omitempty"`

	// Hold Music to be played while callers are waiting. Output is a Media Summary Object. Input must be a Media Lookup Object. Optional. Must refer to a media recording that has is_hold_music set to TRUE. Default is to play a standard ring tone.
	HoldMusic *MediaSummary `json:"hold_music,omitempty"`

	// This action is for forwarding calls to any number of extensions or phone numbers. The forwarding is handled in parallel, meaning that all phone numbers and/or extensions will ring simultaneously. When the call is answered by any single phone number or extension, ringing will stop for all of them. Subsequent actions in this rule set will be performed if the call is not answered before the timeout period is reached, or if it is forwarded to an extension that has its own route and that route does not result in any actions that disconnect the call or take over call handling.
	Items []*RuleSetForwardItem `json:"items"`

	// Menu that this action refers to. Required. Output is a Menu Summary Object. Input must be a Menu Lookup Object.
	Menu *MenuSummary `json:"menu,omitempty"`

	// Queue that this action refers to. Required. Output is a Queue Summary Object. Input must be a Queue Lookup Object.
	Queue *QueueSummary `json:"queue,omitempty"`

	// Seconds that our routing engine should wait until moving on. Optional. Must be an integer between 5 and 90. Default is 5 seconds.
	Timeout *int64 `json:"timeout,omitempty"`

	// Trunk that this action refers to. Required. Output is a Trunk Summary Object. Input must be a Trunk Lookup Object.
	Trunk *TrunkSummary `json:"trunk,omitempty"`
}

// Validate validates this rule set action
func (m *RuleSetAction) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateExtension(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateGreeting(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateHoldMusic(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateItems(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMenu(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateQueue(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTrunk(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RuleSetAction) validateAction(formats strfmt.Registry) error {

	if swag.IsZero(m.Action) { // not required
		return nil
	}

	if err := validate.Pattern("action", "body", string(m.Action), `directory|disconnect|fax|forward|greeting|hold|menu|queue|trunk|voicemail`); err != nil {
		return err
	}

	return nil
}

func (m *RuleSetAction) validateExtension(formats strfmt.Registry) error {

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

func (m *RuleSetAction) validateGreeting(formats strfmt.Registry) error {

	if swag.IsZero(m.Greeting) { // not required
		return nil
	}

	if m.Greeting != nil {

		if err := m.Greeting.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("greeting")
			}
			return err
		}
	}

	return nil
}

func (m *RuleSetAction) validateHoldMusic(formats strfmt.Registry) error {

	if swag.IsZero(m.HoldMusic) { // not required
		return nil
	}

	if m.HoldMusic != nil {

		if err := m.HoldMusic.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hold_music")
			}
			return err
		}
	}

	return nil
}

func (m *RuleSetAction) validateItems(formats strfmt.Registry) error {

	if swag.IsZero(m.Items) { // not required
		return nil
	}

	for i := 0; i < len(m.Items); i++ {

		if swag.IsZero(m.Items[i]) { // not required
			continue
		}

		if m.Items[i] != nil {

			if err := m.Items[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RuleSetAction) validateMenu(formats strfmt.Registry) error {

	if swag.IsZero(m.Menu) { // not required
		return nil
	}

	if m.Menu != nil {

		if err := m.Menu.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("menu")
			}
			return err
		}
	}

	return nil
}

func (m *RuleSetAction) validateQueue(formats strfmt.Registry) error {

	if swag.IsZero(m.Queue) { // not required
		return nil
	}

	if m.Queue != nil {

		if err := m.Queue.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("queue")
			}
			return err
		}
	}

	return nil
}

func (m *RuleSetAction) validateTrunk(formats strfmt.Registry) error {

	if swag.IsZero(m.Trunk) { // not required
		return nil
	}

	if m.Trunk != nil {

		if err := m.Trunk.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("trunk")
			}
			return err
		}
	}

	return nil
}