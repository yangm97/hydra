// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GenericError generic error
//
// swagger:model genericError
type GenericError struct {

	// The status code
	Code int64 `json:"code,omitempty"`

	// Debug information
	//
	// This field is often not exposed to protect against leaking
	// sensitive information.
	Debug string `json:"debug,omitempty"`

	// Further error details
	Details interface{} `json:"details,omitempty"`

	// The error ID
	//
	// Useful when trying to identify various errors in application logic.
	ID string `json:"id,omitempty"`

	// Error message
	//
	// The error's message.
	// Required: true
	Message *string `json:"message"`

	// A human-readable reason for the error
	Reason string `json:"reason,omitempty"`

	// The request ID
	//
	// The request ID is often exposed internally in order to trace
	// errors across service architectures. This is often a UUID.
	Request string `json:"request,omitempty"`

	// The status description
	Status string `json:"status,omitempty"`
}

// Validate validates this generic error
func (m *GenericError) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GenericError) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("message", "body", m.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GenericError) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GenericError) UnmarshalBinary(b []byte) error {
	var res GenericError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
