// Code generated by go-swagger; DO NOT EDIT.

package ipam_models

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

// NetworkRoute network route
//
// swagger:model networkRoute
type NetworkRoute struct {

	// A route's destination network prefix, in CIDR format
	DestinationPrefixes []string `json:"destinationPrefixes"`

	// Selectors that determine which EdgeCompute workload network interfaces should be be used as the route's gateway
	//
	// Selectors are applied to the route in order, with the highest priority assigned to the first gateway.
	GatewaySelectors []*RouteGatewaySelector `json:"gatewaySelectors"`

	// A route's unique identifier
	// Read Only: true
	ID string `json:"id,omitempty"`

	// metadata
	Metadata *NetworkMetadata `json:"metadata,omitempty"`

	// A route's human-readable name
	Name string `json:"name,omitempty"`

	// The ID or slug of the VPC network that a route belongs to
	// Read Only: true
	NetworkID string `json:"networkId,omitempty"`

	// A route's machine-readable name
	Slug string `json:"slug,omitempty"`

	// The ID of the stack that a route belongs to
	// Read Only: true
	StackID string `json:"stackId,omitempty"`

	// The status of a route in different regions
	// Read Only: true
	Statuses []*NetworkRouteStatus `json:"statuses"`
}

// Validate validates this network route
func (m *NetworkRoute) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGatewaySelectors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatuses(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkRoute) validateGatewaySelectors(formats strfmt.Registry) error {
	if swag.IsZero(m.GatewaySelectors) { // not required
		return nil
	}

	for i := 0; i < len(m.GatewaySelectors); i++ {
		if swag.IsZero(m.GatewaySelectors[i]) { // not required
			continue
		}

		if m.GatewaySelectors[i] != nil {
			if err := m.GatewaySelectors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("gatewaySelectors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *NetworkRoute) validateMetadata(formats strfmt.Registry) error {
	if swag.IsZero(m.Metadata) { // not required
		return nil
	}

	if m.Metadata != nil {
		if err := m.Metadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

func (m *NetworkRoute) validateStatuses(formats strfmt.Registry) error {
	if swag.IsZero(m.Statuses) { // not required
		return nil
	}

	for i := 0; i < len(m.Statuses); i++ {
		if swag.IsZero(m.Statuses[i]) { // not required
			continue
		}

		if m.Statuses[i] != nil {
			if err := m.Statuses[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("statuses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this network route based on the context it is used
func (m *NetworkRoute) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateGatewaySelectors(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMetadata(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNetworkID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStackID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatuses(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkRoute) contextValidateGatewaySelectors(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.GatewaySelectors); i++ {

		if m.GatewaySelectors[i] != nil {
			if err := m.GatewaySelectors[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("gatewaySelectors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *NetworkRoute) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *NetworkRoute) contextValidateMetadata(ctx context.Context, formats strfmt.Registry) error {

	if m.Metadata != nil {
		if err := m.Metadata.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

func (m *NetworkRoute) contextValidateNetworkID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "networkId", "body", string(m.NetworkID)); err != nil {
		return err
	}

	return nil
}

func (m *NetworkRoute) contextValidateStackID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "stackId", "body", string(m.StackID)); err != nil {
		return err
	}

	return nil
}

func (m *NetworkRoute) contextValidateStatuses(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "statuses", "body", []*NetworkRouteStatus(m.Statuses)); err != nil {
		return err
	}

	for i := 0; i < len(m.Statuses); i++ {

		if m.Statuses[i] != nil {
			if err := m.Statuses[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("statuses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NetworkRoute) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkRoute) UnmarshalBinary(b []byte) error {
	var res NetworkRoute
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}