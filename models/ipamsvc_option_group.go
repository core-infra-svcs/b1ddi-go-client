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

// IpamsvcOptionGroup OptionGroup
//
// An __OptionGroup__ object (_dhcp/option_group_) is a named collection of options.
//
// swagger:model ipamsvcOptionGroup
type IpamsvcOptionGroup struct {

	// The description for the option group. May contain 0 to 1024 characters. Can include UTF-8.
	Comment string `json:"comment,omitempty"`

	// Time when the object has been created.
	// Read Only: true
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// The list of DHCP options for the option group. May be either a specific option or a group of options.
	DhcpOptions []*IpamsvcOptionItem `json:"dhcp_options"`

	// The resource identifier.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// The name of the option group. Must contain 1 to 256 characters. Can include UTF-8.
	// Required: true
	Name *string `json:"name,omitempty"`

	// The tags for the option group in JSON format.
	Tags interface{} `json:"tags,omitempty"`

	// Time when the object has been updated. Equals to _created_at_ if not updated after creation.
	// Read Only: true
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`
}

// Validate validates this ipamsvc option group
func (m *IpamsvcOptionGroup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDhcpOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IpamsvcOptionGroup) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *IpamsvcOptionGroup) validateDhcpOptions(formats strfmt.Registry) error {
	if swag.IsZero(m.DhcpOptions) { // not required
		return nil
	}

	for i := 0; i < len(m.DhcpOptions); i++ {
		if swag.IsZero(m.DhcpOptions[i]) { // not required
			continue
		}

		if m.DhcpOptions[i] != nil {
			if err := m.DhcpOptions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dhcp_options" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dhcp_options" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IpamsvcOptionGroup) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *IpamsvcOptionGroup) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this ipamsvc option group based on the context it is used
func (m *IpamsvcOptionGroup) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreatedAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDhcpOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUpdatedAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IpamsvcOptionGroup) contextValidateCreatedAt(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created_at", "body", strfmt.DateTime(m.CreatedAt)); err != nil {
		return err
	}

	return nil
}

func (m *IpamsvcOptionGroup) contextValidateDhcpOptions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DhcpOptions); i++ {

		if m.DhcpOptions[i] != nil {
			if err := m.DhcpOptions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dhcp_options" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dhcp_options" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IpamsvcOptionGroup) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *IpamsvcOptionGroup) contextValidateUpdatedAt(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "updated_at", "body", strfmt.DateTime(m.UpdatedAt)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IpamsvcOptionGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IpamsvcOptionGroup) UnmarshalBinary(b []byte) error {
	var res IpamsvcOptionGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
