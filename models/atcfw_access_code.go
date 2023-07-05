// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AtcfwAccessCode atcfw access code
//
// swagger:model atcfwAccessCode
type AtcfwAccessCode struct {

	// Auto generated unique Bypass Code value
	AccessKey string `json:"access_key,omitempty"`

	// The time when the Bypass Code object was activated.
	// Format: date-time
	Activation strfmt.DateTime `json:"activation,omitempty"`

	// The time when the Bypass Code object was created.
	// Read Only: true
	// Format: date-time
	CreatedTime strfmt.DateTime `json:"created_time,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// The time when the Bypass Code object was expired.
	// Format: date-time
	Expiration strfmt.DateTime `json:"expiration,omitempty"`

	// The name of Bypass Code
	Name string `json:"name,omitempty"`

	// The list of SecurityPolicy object identifiers.
	// Example: [12345,53215]
	PolicyIds []int32 `json:"policy_ids"`

	// The list of selected security rules
	Rules []*AtcfwAccessCodeRule `json:"rules"`

	// The time when the Bypass Code object was last updated.
	// Read Only: true
	// Format: date-time
	UpdatedTime strfmt.DateTime `json:"updated_time,omitempty"`
}

// Validate validates this atcfw access code
func (m *AtcfwAccessCode) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActivation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpiration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRules(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AtcfwAccessCode) validateActivation(formats strfmt.Registry) error {
	if swag.IsZero(m.Activation) { // not required
		return nil
	}

	if err := validate.FormatOf("activation", "body", "date-time", m.Activation.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AtcfwAccessCode) validateCreatedTime(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedTime) { // not required
		return nil
	}

	if err := validate.FormatOf("created_time", "body", "date-time", m.CreatedTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AtcfwAccessCode) validateExpiration(formats strfmt.Registry) error {
	if swag.IsZero(m.Expiration) { // not required
		return nil
	}

	if err := validate.FormatOf("expiration", "body", "date-time", m.Expiration.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AtcfwAccessCode) validateRules(formats strfmt.Registry) error {
	if swag.IsZero(m.Rules) { // not required
		return nil
	}

	for i := 0; i < len(m.Rules); i++ {
		if swag.IsZero(m.Rules[i]) { // not required
			continue
		}

		if m.Rules[i] != nil {
			if err := m.Rules[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("rules" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("rules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AtcfwAccessCode) validateUpdatedTime(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedTime) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_time", "body", "date-time", m.UpdatedTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this atcfw access code based on the context it is used
func (m *AtcfwAccessCode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreatedTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRules(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUpdatedTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AtcfwAccessCode) contextValidateCreatedTime(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created_time", "body", strfmt.DateTime(m.CreatedTime)); err != nil {
		return err
	}

	return nil
}

func (m *AtcfwAccessCode) contextValidateRules(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Rules); i++ {

		if m.Rules[i] != nil {

			if swag.IsZero(m.Rules[i]) { // not required
				return nil
			}

			if err := m.Rules[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("rules" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("rules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AtcfwAccessCode) contextValidateUpdatedTime(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "updated_time", "body", strfmt.DateTime(m.UpdatedTime)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AtcfwAccessCode) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AtcfwAccessCode) UnmarshalBinary(b []byte) error {
	var res AtcfwAccessCode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
