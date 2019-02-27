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

// ServiceUpdateRequest service update request
// swagger:model serviceUpdateRequest
type ServiceUpdateRequest struct {

	// app Id
	// Required: true
	AppID *string `json:"appId"`

	// options
	Options interface{} `json:"options,omitempty"`

	// package name
	PackageName string `json:"packageName,omitempty"`

	// package version
	PackageVersion string `json:"packageVersion,omitempty"`

	// If true any stored configuration will be ignored when producing the updated service configuration.
	// Required: true
	Replace *bool `json:"replace"`
}

// Validate validates this service update request
func (m *ServiceUpdateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAppID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReplace(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceUpdateRequest) validateAppID(formats strfmt.Registry) error {

	if err := validate.Required("appId", "body", m.AppID); err != nil {
		return err
	}

	return nil
}

func (m *ServiceUpdateRequest) validateReplace(formats strfmt.Registry) error {

	if err := validate.Required("replace", "body", m.Replace); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceUpdateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceUpdateRequest) UnmarshalBinary(b []byte) error {
	var res ServiceUpdateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
