// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AtcfwAccessCodeReadResponse atcfw access code read response
//
// swagger:model atcfwAccessCodeReadResponse
type AtcfwAccessCodeReadResponse struct {

	// The Bypass Code object.
	Results *AtcfwAccessCode `json:"results,omitempty"`
}

// Validate validates this atcfw access code read response
func (m *AtcfwAccessCodeReadResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResults(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AtcfwAccessCodeReadResponse) validateResults(formats strfmt.Registry) error {
	if swag.IsZero(m.Results) { // not required
		return nil
	}

	if m.Results != nil {
		if err := m.Results.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("results")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("results")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this atcfw access code read response based on the context it is used
func (m *AtcfwAccessCodeReadResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateResults(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AtcfwAccessCodeReadResponse) contextValidateResults(ctx context.Context, formats strfmt.Registry) error {

	if m.Results != nil {

		if swag.IsZero(m.Results) { // not required
			return nil
		}

		if err := m.Results.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("results")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("results")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AtcfwAccessCodeReadResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AtcfwAccessCodeReadResponse) UnmarshalBinary(b []byte) error {
	var res AtcfwAccessCodeReadResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
