// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// UnitHealth unit health
// swagger:model unitHealth
type UnitHealth struct {

	// description
	Description string `json:"description,omitempty"`

	// health
	Health int64 `json:"health,omitempty"`

	// help
	Help string `json:"help,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// output
	Output string `json:"output,omitempty"`
}

// Validate validates this unit health
func (m *UnitHealth) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UnitHealth) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UnitHealth) UnmarshalBinary(b []byte) error {
	var res UnitHealth
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
