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

// AtcdfpDfpCreateOrUpdatePayload DNS Forwarding Proxy object.
//
// For remote office deployments or in cases where installing an endpoint agent is not desirable or possible, you can use the DNS forwarding proxy. It is a software that runs on bare-metal, VM infrastructures, or Infoblox NIOS appliances; and it embeds the client IPs in DNS queries before forwarding them to BloxOne Cloud. The communications are encrypted and client visibility is maintained. The proxy also provides DNS resolution to local DNS zones when you configure local resolvers. Once you set up a DNS forwarding proxy, it becomes the main DNS server for your remote site. It will also cache responses to speed resolution of future queries.
//
// Note that DNS Forwarding Proxy cannot be created (all information regarding DFP is synchronized from hostapp service).
//
// swagger:model atcdfpDfpCreateOrUpdatePayload
type AtcdfpDfpCreateOrUpdatePayload struct {

	// The list of default DNS resolvers that will be used in case if the BloxOne Cloud is unreachable. Deprecated DO NOT USE. Use resolvers_all.
	// Example: ["4.4.4.4","2001:cdba:9abc:5678:ffff:ffff:ffff:ffff"]
	DefaultResolvers []string `json:"default_resolvers"`

	// host information. For internal Use only.
	Host []*AtcdfpDfpHost `json:"host"`

	// The DNS Forwarding Proxy object identifier.
	// Read Only: true
	ID int32 `json:"id,omitempty"`

	// The list of internal domain list ids associated with this DFP (or resolvers)
	InternalDomainLists []int32 `json:"internal_domain_lists"`

	// The name of the DNS Forwarding Proxy.
	// Example: dfp_a
	Name string `json:"name,omitempty"`

	// Point of Presence (PoP) region
	PopRegionID int32 `json:"pop_region_id,omitempty"`

	// The list of internal or local DNS servers' IPv4 or IPv6 addresses that are used as DNS resolvers. Deprecated DO NOT USE. Use resolvers_all.
	// Example: ["1.1.1.1","2001:cdba:9abc:5678:ffff:ffff:ffff:ffff"]
	Resolvers []string `json:"resolvers"`

	// The DNS forwarding proxy additional resolvers used for fallback and local resolution. This field replaces resolvers and default_resolvers fields which are deprecated. Either deprecated fields or new field can be used, both can not be used at same time.
	// Example: {"address":"1.1.1.1","is_fallback":false,"is_local":false,"protocols":["DO53","DOT"]}
	ResolversAll []*AtcdfpResolver `json:"resolvers_all"`

	// The DNS Forwarding Proxy Service ID object identifier.
	ServiceID string `json:"service_id,omitempty"`

	// The name of the DNS Forwarding Proxy Service.
	// Example: dfp_a
	ServiceName string `json:"service_name,omitempty"`

	// The DNS Forwarding Proxy site identifier that is appended to DNS queries originating from this DNS Forwarding Proxy and subsequently used for policy lookup purposes.
	// Example: 134997289555407a8527bea7957ea7a0
	SiteID string `json:"site_id,omitempty"`
}

// Validate validates this atcdfp dfp create or update payload
func (m *AtcdfpDfpCreateOrUpdatePayload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHost(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResolversAll(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AtcdfpDfpCreateOrUpdatePayload) validateHost(formats strfmt.Registry) error {
	if swag.IsZero(m.Host) { // not required
		return nil
	}

	for i := 0; i < len(m.Host); i++ {
		if swag.IsZero(m.Host[i]) { // not required
			continue
		}

		if m.Host[i] != nil {
			if err := m.Host[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("host" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("host" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AtcdfpDfpCreateOrUpdatePayload) validateResolversAll(formats strfmt.Registry) error {
	if swag.IsZero(m.ResolversAll) { // not required
		return nil
	}

	for i := 0; i < len(m.ResolversAll); i++ {
		if swag.IsZero(m.ResolversAll[i]) { // not required
			continue
		}

		if m.ResolversAll[i] != nil {
			if err := m.ResolversAll[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("resolvers_all" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("resolvers_all" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this atcdfp dfp create or update payload based on the context it is used
func (m *AtcdfpDfpCreateOrUpdatePayload) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHost(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResolversAll(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AtcdfpDfpCreateOrUpdatePayload) contextValidateHost(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Host); i++ {

		if m.Host[i] != nil {

			if swag.IsZero(m.Host[i]) { // not required
				return nil
			}

			if err := m.Host[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("host" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("host" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AtcdfpDfpCreateOrUpdatePayload) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int32(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *AtcdfpDfpCreateOrUpdatePayload) contextValidateResolversAll(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ResolversAll); i++ {

		if m.ResolversAll[i] != nil {

			if swag.IsZero(m.ResolversAll[i]) { // not required
				return nil
			}

			if err := m.ResolversAll[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("resolvers_all" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("resolvers_all" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AtcdfpDfpCreateOrUpdatePayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AtcdfpDfpCreateOrUpdatePayload) UnmarshalBinary(b []byte) error {
	var res AtcdfpDfpCreateOrUpdatePayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
