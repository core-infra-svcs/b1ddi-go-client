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

// IpamsvcIPSpace IPSpace
//
// An __IPSpace__ object (_ipam/ip_space_) allows customers to represent their entire managed address space with no collision. A collision arises when two or more block of addresses overlap partially or fully.
//
// swagger:model ipamsvcIPSpace
type IpamsvcIPSpace struct {

	// The Automated Scope Management configuration for the IP space.
	AsmConfig *IpamsvcASMConfig `json:"asm_config,omitempty"`

	// The number of times the automated scope management usage limits have been exceeded for any of the subnets in this IP space.
	// Read Only: true
	AsmScopeFlag int64 `json:"asm_scope_flag,omitempty"`

	// The description for the IP space. May contain 0 to 1024 characters. Can include UTF-8.
	Comment string `json:"comment,omitempty"`

	// Time when the object has been created.
	// Read Only: true
	// Format: date-time
	CreatedAt *strfmt.DateTime `json:"created_at,omitempty"`

	// Controls who does the DDNS updates.
	//
	// Valid values are:
	// * _client_: DHCP server updates DNS if requested by client.
	// * _server_: DHCP server always updates DNS, overriding an update request from the client, unless the client requests no updates.
	// * _ignore_: DHCP server always updates DNS, even if the client says not to.
	// * _over_client_update_: Same as _server_. DHCP server always updates DNS, overriding an update request from the client, unless the client requests no updates.
	// * _over_no_update_: DHCP server updates DNS even if the client requests that no updates be done. If the client requests to do the update, DHCP server allows it.
	//
	// Defaults to _client_.
	DdnsClientUpdate string `json:"ddns_client_update,omitempty"`

	// The domain suffix for DDNS updates. FQDN, may be empty.
	//
	// Defaults to empty.
	DdnsDomain string `json:"ddns_domain,omitempty"`

	// Indicates if DDNS needs to generate a hostname when not supplied by the client.
	//
	// Defaults to _false_.
	DdnsGenerateName bool `json:"ddns_generate_name,omitempty"`

	// The prefix used in the generation of an FQDN.
	//
	// When generating a name, DHCP server will construct the name in the format: [ddns-generated-prefix]-[address-text].[ddns-qualifying-suffix].
	// where address-text is simply the lease IP address converted to a hyphenated string.
	//
	// Defaults to "myhost".
	DdnsGeneratedPrefix string `json:"ddns_generated_prefix,omitempty"`

	// Determines if DDNS updates are enabled at the IP space level.
	// Defaults to _true_.
	DdnsSendUpdates *bool `json:"ddns_send_updates,omitempty"`

	// Instructs the DHCP server to always update the DNS information when a lease is renewed even if its DNS information has not changed.
	//
	// Defaults to _false_.
	DdnsUpdateOnRenew bool `json:"ddns_update_on_renew,omitempty"`

	// When true, DHCP server will apply conflict resolution, as described in RFC 4703, when attempting to fulfill the update request.
	//
	// When false, DHCP server will simply attempt to update the DNS entries per the request, regardless of whether or not they conflict with existing entries owned by other DHCP4 clients.
	//
	// Defaults to _true_.
	DdnsUseConflictResolution *bool `json:"ddns_use_conflict_resolution,omitempty"`

	// The shared DHCP configuration for the IP space that controls how leases are issued.
	DhcpConfig *IpamsvcDHCPConfig `json:"dhcp_config,omitempty"`

	// The list of DHCP options for the IP space. May be either a specific option or a group of options.
	DhcpOptions []*IpamsvcOptionItem `json:"dhcp_options,omitempty"`

	// The configuration for header option filename field.
	HeaderOptionFilename string `json:"header_option_filename,omitempty"`

	// The configuration for header option server address field.
	HeaderOptionServerAddress string `json:"header_option_server_address,omitempty"`

	// The configuration for header option server name field.
	HeaderOptionServerName string `json:"header_option_server_name,omitempty"`

	// The character to replace non-matching characters with, when hostname rewrite is enabled.
	//
	// Any single ASCII character.
	//
	// Defaults to "_".
	HostnameRewriteChar string `json:"hostname_rewrite_char,omitempty"`

	// Indicates if client supplied hostnames will be rewritten prior to DDNS update by replacing every character that does not match _hostname_rewrite_regex_ by _hostname_rewrite_char_.
	//
	// Defaults to _false_.
	HostnameRewriteEnabled bool `json:"hostname_rewrite_enabled,omitempty"`

	// The regex bracket expression to match valid characters.
	//
	// Must begin with "[" and end with "]" and be a compilable POSIX regex.
	//
	// Defaults to "[^a-zA-Z0-9_.]".
	HostnameRewriteRegex string `json:"hostname_rewrite_regex,omitempty"`

	// The resource identifier.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// inheritance assigned hosts
	InheritanceAssignedHosts interface{} `json:"inheritance_assigned_hosts,omitempty"`

	// The inheritance configuration.
	InheritanceSources *IpamsvcIPSpaceInheritance `json:"inheritance_sources,omitempty"`

	// The name of the IP space. Must contain 1 to 256 characters. Can include UTF-8.
	// Required: true
	Name *string `json:"name"`

	// The tags for the IP space in JSON format.
	Tags interface{} `json:"tags,omitempty"`

	// The utilization threshold settings for the IP space.
	// Read Only: true
	Threshold *IpamsvcUtilizationThreshold `json:"threshold,omitempty"`

	// Time when the object has been updated. Equals to _created_at_ if not updated after creation.
	// Read Only: true
	// Format: date-time
	UpdatedAt *strfmt.DateTime `json:"updated_at,omitempty"`

	// The utilization of IP addresses in the IP space.
	// Read Only: true
	Utilization *IpamsvcUtilization `json:"utilization,omitempty"`

	// The resource identifier.
	VendorSpecificOptionOptionSpace string `json:"vendor_specific_option_option_space,omitempty"`
}

// Validate validates this ipamsvc IP space
func (m *IpamsvcIPSpace) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAsmConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDhcpConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDhcpOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInheritanceSources(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateThreshold(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUtilization(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IpamsvcIPSpace) validateAsmConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.AsmConfig) { // not required
		return nil
	}

	if m.AsmConfig != nil {
		if err := m.AsmConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("asm_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("asm_config")
			}
			return err
		}
	}

	return nil
}

func (m *IpamsvcIPSpace) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *IpamsvcIPSpace) validateDhcpConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.DhcpConfig) { // not required
		return nil
	}

	if m.DhcpConfig != nil {
		if err := m.DhcpConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dhcp_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dhcp_config")
			}
			return err
		}
	}

	return nil
}

func (m *IpamsvcIPSpace) validateDhcpOptions(formats strfmt.Registry) error {
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

func (m *IpamsvcIPSpace) validateInheritanceSources(formats strfmt.Registry) error {
	if swag.IsZero(m.InheritanceSources) { // not required
		return nil
	}

	if m.InheritanceSources != nil {
		if err := m.InheritanceSources.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("inheritance_sources")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("inheritance_sources")
			}
			return err
		}
	}

	return nil
}

func (m *IpamsvcIPSpace) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *IpamsvcIPSpace) validateThreshold(formats strfmt.Registry) error {
	if swag.IsZero(m.Threshold) { // not required
		return nil
	}

	if m.Threshold != nil {
		if err := m.Threshold.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("threshold")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("threshold")
			}
			return err
		}
	}

	return nil
}

func (m *IpamsvcIPSpace) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *IpamsvcIPSpace) validateUtilization(formats strfmt.Registry) error {
	if swag.IsZero(m.Utilization) { // not required
		return nil
	}

	if m.Utilization != nil {
		if err := m.Utilization.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("utilization")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("utilization")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this ipamsvc IP space based on the context it is used
func (m *IpamsvcIPSpace) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAsmConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAsmScopeFlag(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreatedAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDhcpConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDhcpOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInheritanceSources(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateThreshold(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUpdatedAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUtilization(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IpamsvcIPSpace) contextValidateAsmConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.AsmConfig != nil {
		if err := m.AsmConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("asm_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("asm_config")
			}
			return err
		}
	}

	return nil
}

func (m *IpamsvcIPSpace) contextValidateAsmScopeFlag(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "asm_scope_flag", "body", int64(m.AsmScopeFlag)); err != nil {
		return err
	}

	return nil
}

func (m *IpamsvcIPSpace) contextValidateCreatedAt(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created_at", "body", m.CreatedAt); err != nil {
		return err
	}

	return nil
}

func (m *IpamsvcIPSpace) contextValidateDhcpConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.DhcpConfig != nil {
		if err := m.DhcpConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dhcp_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dhcp_config")
			}
			return err
		}
	}

	return nil
}

func (m *IpamsvcIPSpace) contextValidateDhcpOptions(ctx context.Context, formats strfmt.Registry) error {

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

func (m *IpamsvcIPSpace) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *IpamsvcIPSpace) contextValidateInheritanceSources(ctx context.Context, formats strfmt.Registry) error {

	if m.InheritanceSources != nil {
		if err := m.InheritanceSources.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("inheritance_sources")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("inheritance_sources")
			}
			return err
		}
	}

	return nil
}

func (m *IpamsvcIPSpace) contextValidateThreshold(ctx context.Context, formats strfmt.Registry) error {

	if m.Threshold != nil {
		if err := m.Threshold.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("threshold")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("threshold")
			}
			return err
		}
	}

	return nil
}

func (m *IpamsvcIPSpace) contextValidateUpdatedAt(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "updated_at", "body", m.UpdatedAt); err != nil {
		return err
	}

	return nil
}

func (m *IpamsvcIPSpace) contextValidateUtilization(ctx context.Context, formats strfmt.Registry) error {

	if m.Utilization != nil {
		if err := m.Utilization.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("utilization")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("utilization")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IpamsvcIPSpace) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IpamsvcIPSpace) UnmarshalBinary(b []byte) error {
	var res IpamsvcIPSpace
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
