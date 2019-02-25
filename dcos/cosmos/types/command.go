// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Command command
// swagger:model command
type Command struct {

	// [Deprecated v3.x] An array of strings representing of the requirements file to use for installing the subcommand for Pip. Each item is interpreted as a line in the requirements file.
	// Required: true
	Pip []string `json:"pip"`
}

// Validate validates this command
func (m *Command) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePip(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Command) validatePip(formats strfmt.Registry) error {

	if err := validate.Required("pip", "body", m.Pip); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Command) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Command) UnmarshalBinary(b []byte) error {
	var res Command
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
