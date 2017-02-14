package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// FilterListPhoneNumbersRegions filter list phone numbers regions
// swagger:model FilterListPhoneNumbersRegions
type FilterListPhoneNumbersRegions struct {

	// city
	City []string `json:"city"`

	// country code
	CountryCode []string `json:"country_code"`

	// country postal code
	CountryPostalCode []string `json:"country_postal_code"`

	// is toll free
	IsTollFree []string `json:"is_toll_free"`

	// npa
	Npa []int64 `json:"npa"`

	// nxx
	Nxx []string `json:"nxx"`

	// province postal code
	ProvincePostalCode []string `json:"province_postal_code"`
}

// Validate validates this filter list phone numbers regions
func (m *FilterListPhoneNumbersRegions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCity(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCountryCode(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCountryPostalCode(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateIsTollFree(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateNpa(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateNxx(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateProvincePostalCode(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FilterListPhoneNumbersRegions) validateCity(formats strfmt.Registry) error {

	if swag.IsZero(m.City) { // not required
		return nil
	}

	return nil
}

func (m *FilterListPhoneNumbersRegions) validateCountryCode(formats strfmt.Registry) error {

	if swag.IsZero(m.CountryCode) { // not required
		return nil
	}

	return nil
}

func (m *FilterListPhoneNumbersRegions) validateCountryPostalCode(formats strfmt.Registry) error {

	if swag.IsZero(m.CountryPostalCode) { // not required
		return nil
	}

	return nil
}

func (m *FilterListPhoneNumbersRegions) validateIsTollFree(formats strfmt.Registry) error {

	if swag.IsZero(m.IsTollFree) { // not required
		return nil
	}

	return nil
}

func (m *FilterListPhoneNumbersRegions) validateNpa(formats strfmt.Registry) error {

	if swag.IsZero(m.Npa) { // not required
		return nil
	}

	return nil
}

func (m *FilterListPhoneNumbersRegions) validateNxx(formats strfmt.Registry) error {

	if swag.IsZero(m.Nxx) { // not required
		return nil
	}

	return nil
}

func (m *FilterListPhoneNumbersRegions) validateProvincePostalCode(formats strfmt.Registry) error {

	if swag.IsZero(m.ProvincePostalCode) { // not required
		return nil
	}

	return nil
}
