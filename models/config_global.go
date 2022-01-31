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

// ConfigGlobal Global
//
// A Global configuration (_dns/global_). Used by default unless more specific configuration exists. There is only one instance of this object.
//
// swagger:model configGlobal
type ConfigGlobal struct {

	// Optional. List of custom root nameservers. The order does not matter.
	//
	// Error if empty while _custom_root_ns_enabled_ is _true_.
	// Error if there are duplicate items in the list.
	//
	// Defaults to empty.
	CustomRootNs []*ConfigRootNS `json:"custom_root_ns"`

	// Optional. _true_ to use custom root nameservers instead of the default ones.
	//
	// The _custom_root_ns_ field is validated when enabled.
	//
	// Defaults to false.
	CustomRootNsEnabled bool `json:"custom_root_ns_enabled,omitempty"`

	// Optional. _true_ to perform DNSSEC validation.
	// Ignored if _dnssec_enabled_ is _false_.
	//
	// Defaults to _true_.
	DnssecEnableValidation bool `json:"dnssec_enable_validation,omitempty"`

	// Optional. Master toggle for all DNSSEC processing.
	// Other _dnssec_*_ configuration is unused if this is disabled.
	//
	// Defaults to _true_.
	DnssecEnabled bool `json:"dnssec_enabled,omitempty"`

	// DNSSEC root keys. The root keys are not configurable.
	//
	// A default list is provided by cloud management and included here for config generation.
	// Read Only: true
	DnssecRootKeys []*ConfigTrustAnchor `json:"dnssec_root_keys"`

	// Optional. DNSSEC trust anchors.
	//
	// Error if there are list items with duplicate (_zone_, _sep_, _algorithm_) combinations.
	//
	// Defaults to empty.
	DnssecTrustAnchors []*ConfigTrustAnchor `json:"dnssec_trust_anchors"`

	// Optional. _true_ to reject expired DNSSEC keys.
	// Ignored if either _dnssec_enabled_ or _dnssec_enable_validation_ is _false_.
	//
	// Defaults to _true_.
	DnssecValidateExpiry bool `json:"dnssec_validate_expiry,omitempty"`

	// Optional. _true_ to enable EDNS client subnet for recursive queries.
	// Other _ecs_*_ fields are ignored if this field is not enabled.
	//
	// Defaults to _false_.
	EcsEnabled bool `json:"ecs_enabled,omitempty"`

	// Optional. _true_ to enable ECS options in outbound queries. This functionality has additional overhead so it is disabled by default.
	//
	// Defaults to _false_.
	EcsForwarding bool `json:"ecs_forwarding,omitempty"`

	// Optional. Maximum scope length for v4 ECS.
	//
	// Unsigned integer, min 1 max 24.
	//
	// Defaults to 24.
	EcsPrefixV4 int64 `json:"ecs_prefix_v4,omitempty"`

	// Optional. Maximum scope length for v6 ECS.
	//
	// Unsigned integer, min 1 max 56.
	//
	// Defaults to 56.
	EcsPrefixV6 int64 `json:"ecs_prefix_v6,omitempty"`

	// Optional. List of zones where ECS queries may be sent.
	//
	// Error if empty while _ecs_enabled_ is true.
	// Error if there are duplicate FQDNs in the list.
	//
	// Defaults to empty.
	EcsZones []*ConfigECSZone `json:"ecs_zones"`

	// Optional. _edns_udp_size_ represents the edns UDP size.
	// The size a querying DNS server advertises to the DNS server it’s sending a query to.
	//
	// Defaults to 1232 bytes.
	EdnsUDPSize int64 `json:"edns_udp_size,omitempty"`

	// Optional. List of forwarders.
	//
	// Error if empty while _forwarders_only_ is _true_.
	// Error if there are items in the list with duplicate addresses.
	//
	// Defaults to empty.
	Forwarders []*ConfigForwarder `json:"forwarders"`

	// Optional. _true_ to only forward.
	//
	// Defaults to _false_.
	ForwardersOnly bool `json:"forwarders_only,omitempty"`

	// _gss_tsig_enabled_ enables/disables GSS-TSIG signed dynamic updates.
	//
	// Defaults to _false_.
	GssTsigEnabled bool `json:"gss_tsig_enabled,omitempty"`

	// The resource identifier.
	// Required: true
	// Read Only: true
	ID string `json:"id,omitempty"`

	// _kerberos_keys_ contains a list of keys for GSS-TSIG signed dynamic updates.
	//
	// Defaults to empty.
	KerberosKeys []*ConfigKerberosKey `json:"kerberos_keys"`

	// Optional. Unused in the current on-prem DNS server implementation.
	//
	// Unsigned integer, min 0 max 3600 (1h).
	//
	// Defaults to 600.
	LameTTL int64 `json:"lame_ttl,omitempty"`

	// Optional. Control DNS query/response logging functionality.
	//
	// Defaults to _true_.
	LogQueryResponse bool `json:"log_query_response,omitempty"`

	// Optional. If _true_ only recursive queries from matching clients access the view.
	//
	// Defaults to _false_.
	MatchRecursiveOnly bool `json:"match_recursive_only,omitempty"`

	// Optional. Seconds to cache positive responses.
	//
	// Unsigned integer, min 1 max 604800 (7d).
	//
	// Defaults to 604800 (7d).
	MaxCacheTTL int64 `json:"max_cache_ttl,omitempty"`

	// Optional. Seconds to cache negative responses.
	//
	// Unsigned integer, min 1 max 604800 (7d).
	//
	// Defaults to 10800 (3h).
	MaxNegativeTTL int64 `json:"max_negative_ttl,omitempty"`

	// Optional. _max_udp_size_ represents maximum UDP payload size.
	// The maximum number of bytes a responding DNS server will send to a UDP datagram.
	//
	// Defaults to 1232 bytes.
	MaxUDPSize int64 `json:"max_udp_size,omitempty"`

	// Optional. When enabled, the DNS server will only add records to the authority and additional data sections when they are required.
	//
	// Defaults to _false_.
	MinimalResponses bool `json:"minimal_responses,omitempty"`

	// _notify_ all external secondary DNS servers.
	//
	// Defaults to _false_.
	Notify bool `json:"notify,omitempty"`

	// Optional. Clients must match this ACL to make authoritative queries.
	// Also used for recursive queries if that ACL is unset.
	//
	// Defaults to empty.
	QueryACL []*ConfigACLItem `json:"query_acl"`

	// Optional. Source port for outbound DNS queries.
	// When set to 0 the port is unspecified and the implementation may randomize it using any available ports.
	//
	// Defaults to 0.
	QueryPort int64 `json:"query_port,omitempty"`

	// Optional. Clients must match this ACL to make recursive queries. If this ACL is empty, then the _query_acl_ field will be used instead.
	//
	// Defaults to empty.
	RecursionACL []*ConfigACLItem `json:"recursion_acl"`

	// Optional. _true_ to allow recursive DNS queries.
	//
	// Defaults to _true_.
	RecursionEnabled bool `json:"recursion_enabled,omitempty"`

	// Optional. Defines the number of simultaneous recursive lookups the server will perform on behalf of its clients.
	//
	// Defaults to 1000.
	RecursiveClients int64 `json:"recursive_clients,omitempty"`

	// Optional. Seconds before a recursive query times out.
	//
	// Unsigned integer, min 10 max 30.
	//
	// Defaults to 10.
	ResolverQueryTimeout int64 `json:"resolver_query_timeout,omitempty"`

	// Optional. Maximum concurrent inbound AXFRs.
	// When set to 0 a host-dependent default will be used.
	//
	// Defaults to 0.
	SecondaryAxfrQueryLimit int64 `json:"secondary_axfr_query_limit,omitempty"`

	// Optional. Maximum concurrent outbound SOA queries.
	// When set to 0 a host-dependent default will be used.
	//
	// Defaults to 0.
	SecondarySoaQueryLimit int64 `json:"secondary_soa_query_limit,omitempty"`

	// Optional. Clients must match this ACL to receive zone transfers.
	//
	// Defaults to "deny any".
	TransferACL []*ConfigACLItem `json:"transfer_acl"`

	// Optional. Specifies which hosts are allowed to issue Dynamic DNS updates for authoritative zones of _primary_type_ _cloud_.
	//
	// Defaults to empty.
	UpdateACL []*ConfigACLItem `json:"update_acl"`

	// Optional. Use default forwarders to resolve queries for subzones.
	//
	// Defaults to _true_.
	UseForwardersForSubzones bool `json:"use_forwarders_for_subzones,omitempty"`

	// Optional. ZoneAuthority.
	ZoneAuthority *ConfigZoneAuthority `json:"zone_authority,omitempty"`
}

// Validate validates this config global
func (m *ConfigGlobal) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCustomRootNs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDnssecRootKeys(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDnssecTrustAnchors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEcsZones(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateForwarders(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKerberosKeys(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQueryACL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecursionACL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTransferACL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdateACL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateZoneAuthority(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigGlobal) validateCustomRootNs(formats strfmt.Registry) error {
	if swag.IsZero(m.CustomRootNs) { // not required
		return nil
	}

	for i := 0; i < len(m.CustomRootNs); i++ {
		if swag.IsZero(m.CustomRootNs[i]) { // not required
			continue
		}

		if m.CustomRootNs[i] != nil {
			if err := m.CustomRootNs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("custom_root_ns" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("custom_root_ns" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ConfigGlobal) validateDnssecRootKeys(formats strfmt.Registry) error {
	if swag.IsZero(m.DnssecRootKeys) { // not required
		return nil
	}

	for i := 0; i < len(m.DnssecRootKeys); i++ {
		if swag.IsZero(m.DnssecRootKeys[i]) { // not required
			continue
		}

		if m.DnssecRootKeys[i] != nil {
			if err := m.DnssecRootKeys[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dnssec_root_keys" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dnssec_root_keys" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ConfigGlobal) validateDnssecTrustAnchors(formats strfmt.Registry) error {
	if swag.IsZero(m.DnssecTrustAnchors) { // not required
		return nil
	}

	for i := 0; i < len(m.DnssecTrustAnchors); i++ {
		if swag.IsZero(m.DnssecTrustAnchors[i]) { // not required
			continue
		}

		if m.DnssecTrustAnchors[i] != nil {
			if err := m.DnssecTrustAnchors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dnssec_trust_anchors" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dnssec_trust_anchors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ConfigGlobal) validateEcsZones(formats strfmt.Registry) error {
	if swag.IsZero(m.EcsZones) { // not required
		return nil
	}

	for i := 0; i < len(m.EcsZones); i++ {
		if swag.IsZero(m.EcsZones[i]) { // not required
			continue
		}

		if m.EcsZones[i] != nil {
			if err := m.EcsZones[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ecs_zones" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ecs_zones" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ConfigGlobal) validateForwarders(formats strfmt.Registry) error {
	if swag.IsZero(m.Forwarders) { // not required
		return nil
	}

	for i := 0; i < len(m.Forwarders); i++ {
		if swag.IsZero(m.Forwarders[i]) { // not required
			continue
		}

		if m.Forwarders[i] != nil {
			if err := m.Forwarders[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("forwarders" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("forwarders" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ConfigGlobal) validateID(formats strfmt.Registry) error {

	if err := validate.RequiredString("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *ConfigGlobal) validateKerberosKeys(formats strfmt.Registry) error {
	if swag.IsZero(m.KerberosKeys) { // not required
		return nil
	}

	for i := 0; i < len(m.KerberosKeys); i++ {
		if swag.IsZero(m.KerberosKeys[i]) { // not required
			continue
		}

		if m.KerberosKeys[i] != nil {
			if err := m.KerberosKeys[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("kerberos_keys" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("kerberos_keys" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ConfigGlobal) validateQueryACL(formats strfmt.Registry) error {
	if swag.IsZero(m.QueryACL) { // not required
		return nil
	}

	for i := 0; i < len(m.QueryACL); i++ {
		if swag.IsZero(m.QueryACL[i]) { // not required
			continue
		}

		if m.QueryACL[i] != nil {
			if err := m.QueryACL[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("query_acl" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("query_acl" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ConfigGlobal) validateRecursionACL(formats strfmt.Registry) error {
	if swag.IsZero(m.RecursionACL) { // not required
		return nil
	}

	for i := 0; i < len(m.RecursionACL); i++ {
		if swag.IsZero(m.RecursionACL[i]) { // not required
			continue
		}

		if m.RecursionACL[i] != nil {
			if err := m.RecursionACL[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("recursion_acl" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("recursion_acl" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ConfigGlobal) validateTransferACL(formats strfmt.Registry) error {
	if swag.IsZero(m.TransferACL) { // not required
		return nil
	}

	for i := 0; i < len(m.TransferACL); i++ {
		if swag.IsZero(m.TransferACL[i]) { // not required
			continue
		}

		if m.TransferACL[i] != nil {
			if err := m.TransferACL[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("transfer_acl" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("transfer_acl" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ConfigGlobal) validateUpdateACL(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdateACL) { // not required
		return nil
	}

	for i := 0; i < len(m.UpdateACL); i++ {
		if swag.IsZero(m.UpdateACL[i]) { // not required
			continue
		}

		if m.UpdateACL[i] != nil {
			if err := m.UpdateACL[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("update_acl" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("update_acl" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ConfigGlobal) validateZoneAuthority(formats strfmt.Registry) error {
	if swag.IsZero(m.ZoneAuthority) { // not required
		return nil
	}

	if m.ZoneAuthority != nil {
		if err := m.ZoneAuthority.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("zone_authority")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("zone_authority")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this config global based on the context it is used
func (m *ConfigGlobal) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCustomRootNs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDnssecRootKeys(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDnssecTrustAnchors(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEcsZones(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateForwarders(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateKerberosKeys(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateQueryACL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRecursionACL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTransferACL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUpdateACL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateZoneAuthority(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigGlobal) contextValidateCustomRootNs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.CustomRootNs); i++ {

		if m.CustomRootNs[i] != nil {
			if err := m.CustomRootNs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("custom_root_ns" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("custom_root_ns" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ConfigGlobal) contextValidateDnssecRootKeys(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "dnssec_root_keys", "body", []*ConfigTrustAnchor(m.DnssecRootKeys)); err != nil {
		return err
	}

	for i := 0; i < len(m.DnssecRootKeys); i++ {

		if m.DnssecRootKeys[i] != nil {
			if err := m.DnssecRootKeys[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dnssec_root_keys" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dnssec_root_keys" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ConfigGlobal) contextValidateDnssecTrustAnchors(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DnssecTrustAnchors); i++ {

		if m.DnssecTrustAnchors[i] != nil {
			if err := m.DnssecTrustAnchors[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dnssec_trust_anchors" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dnssec_trust_anchors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ConfigGlobal) contextValidateEcsZones(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.EcsZones); i++ {

		if m.EcsZones[i] != nil {
			if err := m.EcsZones[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ecs_zones" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ecs_zones" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ConfigGlobal) contextValidateForwarders(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Forwarders); i++ {

		if m.Forwarders[i] != nil {
			if err := m.Forwarders[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("forwarders" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("forwarders" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ConfigGlobal) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *ConfigGlobal) contextValidateKerberosKeys(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.KerberosKeys); i++ {

		if m.KerberosKeys[i] != nil {
			if err := m.KerberosKeys[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("kerberos_keys" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("kerberos_keys" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ConfigGlobal) contextValidateQueryACL(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.QueryACL); i++ {

		if m.QueryACL[i] != nil {
			if err := m.QueryACL[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("query_acl" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("query_acl" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ConfigGlobal) contextValidateRecursionACL(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.RecursionACL); i++ {

		if m.RecursionACL[i] != nil {
			if err := m.RecursionACL[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("recursion_acl" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("recursion_acl" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ConfigGlobal) contextValidateTransferACL(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.TransferACL); i++ {

		if m.TransferACL[i] != nil {
			if err := m.TransferACL[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("transfer_acl" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("transfer_acl" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ConfigGlobal) contextValidateUpdateACL(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.UpdateACL); i++ {

		if m.UpdateACL[i] != nil {
			if err := m.UpdateACL[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("update_acl" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("update_acl" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ConfigGlobal) contextValidateZoneAuthority(ctx context.Context, formats strfmt.Registry) error {

	if m.ZoneAuthority != nil {
		if err := m.ZoneAuthority.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("zone_authority")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("zone_authority")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigGlobal) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigGlobal) UnmarshalBinary(b []byte) error {
	var res ConfigGlobal
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
