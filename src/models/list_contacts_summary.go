package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ListContactsSummary list contacts summary
// swagger:model ListContactsSummary
type ListContactsSummary struct {

	// filters
	Filters *FilterIDGroupIDUpdatedAtArray `json:"filters,omitempty"`

	// items
	Items []*ContactSummary `json:"items"`

	// limit
	Limit int64 `json:"limit,omitempty"`

	// offset
	Offset int64 `json:"offset,omitempty"`

	// sort
	Sort *SortIDUpdatedAt `json:"sort,omitempty"`

	// total
	Total int64 `json:"total,omitempty"`
}

// Validate validates this list contacts summary
func (m *ListContactsSummary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFilters(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateItems(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSort(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListContactsSummary) validateFilters(formats strfmt.Registry) error {

	if swag.IsZero(m.Filters) { // not required
		return nil
	}

	if m.Filters != nil {

		if err := m.Filters.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("filters")
			}
			return err
		}
	}

	return nil
}

func (m *ListContactsSummary) validateItems(formats strfmt.Registry) error {

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

func (m *ListContactsSummary) validateSort(formats strfmt.Registry) error {

	if swag.IsZero(m.Sort) { // not required
		return nil
	}

	if m.Sort != nil {

		if err := m.Sort.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sort")
			}
			return err
		}
	}

	return nil
}
