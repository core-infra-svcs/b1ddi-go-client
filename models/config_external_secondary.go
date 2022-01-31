// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ConfigExternalSecondary ExternalSecondary
//
// External DNS secondary.
//
// swagger:model configExternalSecondary
type ConfigExternalSecondary struct {

	// IP Address of nameserver.
	// Required: true
	Address *string `json:"address,omitempty"`

	// FQDN of nameserver.
	// Required: true
	Fqdn *string `json:"fqdn,omitempty"`

	// FQDN of nameserver in punycode.
	// Read Only: true
	ProtocolFqdn string `json:"protocol_fqdn,omitempty"`

	// If enabled, the NS record and glue record will NOT be automatically generated
	// according to secondaries nameserver assignment.
	//
	// Default: _false_
	Stealth bool `json:"stealth,omitempty"`

	// If enabled, secondaries will use the configured TSIG key when requesting a zone transfer.
	//
	// Default: _false_
	TsigEnabled bool `json:"tsig_enabled,omitempty"`

	// TSIG key.
	//
	// Error if empty while _tsig_enabled_ is _true_.
	TsigKey *ConfigTSIGKey `json:"tsig_key,omitempty"`
}

// Validate validates this config external secondary
func (m *ConfigExternalSecondary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFqdn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTsigKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigExternalSecondary) validateAddress(formats strfmt.Registry) error {

	if err := validate.Required("address", "body", m.Address); err != nil {
		return err
	}

	return nil
}

func (m *ConfigExternalSecondary) validateFqdn(formats strfmt.Registry) error {

	if err := validate.Required("fqdn", "body", m.Fqdn); err != nil {
		return err
	}

	return nil
}

func (m *ConfigExternalSecondary) validateTsigKey(formats strfmt.Registry) error {
	if swag.IsZero(m.TsigKey) { // not required
		return nil
	}

	if m.TsigKey != nil {
		if err := m.TsigKey.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tsig_key")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tsig_key")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this config external secondary based on the context it is used
func (m *ConfigExternalSecondary) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateProtocolFqdn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTsigKey(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigExternalSecondary) contextValidateProtocolFqdn(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "protocol_fqdn", "body", string(m.ProtocolFqdn)); err != nil {
		return err
	}

	return nil
}

func (m *ConfigExternalSecondary) contextValidateTsigKey(ctx context.Context, formats strfmt.Registry) error {

	if m.TsigKey != nil {
		if err := m.TsigKey.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tsig_key")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tsig_key")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigExternalSecondary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigExternalSecondary) UnmarshalBinary(b []byte) error {
	var res ConfigExternalSecondary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
