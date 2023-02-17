// Code generated by go-swagger; DO NOT EDIT.

package workload_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Workloadv1NetworkInterfaceStatus Network interfaces that are attached to an instance in a workload
//
// swagger:model workloadv1NetworkInterfaceStatus
type Workloadv1NetworkInterfaceStatus struct {

	// A network interface's IPv4 gateway address
	Gateway string `json:"gateway,omitempty"`

	// A network interface's primary IPv4 address
	IPAddress string `json:"ipAddress,omitempty"`

	// Additional IPv4 addresses bound to a network interface
	IPAddressAliases []string `json:"ipAddressAliases"`

	// A network interface's primary IPv6 address
	IPV6Address string `json:"ipv6Address,omitempty"`

	// Additional IPv6 addresses bound to a network interface
	IPV6AddressAliases []string `json:"ipv6AddressAliases"`

	// A network interface's IPv6 gateway address
	IPV6Gateway string `json:"ipv6Gateway,omitempty"`

	// A network interface's name
	Network string `json:"network,omitempty"`
}

// Validate validates this workloadv1 network interface status
func (m *Workloadv1NetworkInterfaceStatus) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this workloadv1 network interface status based on context it is used
func (m *Workloadv1NetworkInterfaceStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Workloadv1NetworkInterfaceStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Workloadv1NetworkInterfaceStatus) UnmarshalBinary(b []byte) error {
	var res Workloadv1NetworkInterfaceStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
