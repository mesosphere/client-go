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

// V3PackageDescribeResponse v3 package describe response
// swagger:model v3PackageDescribeResponse
type V3PackageDescribeResponse struct {

	// package
	// Required: true
	Package *V50PackageDefinition `json:"package"`
}

// Validate validates this v3 package describe response
func (m *V3PackageDescribeResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePackage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V3PackageDescribeResponse) validatePackage(formats strfmt.Registry) error {

	if err := validate.Required("package", "body", m.Package); err != nil {
		return err
	}

	if m.Package != nil {
		if err := m.Package.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("package")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V3PackageDescribeResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V3PackageDescribeResponse) UnmarshalBinary(b []byte) error {
	var res V3PackageDescribeResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
