// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StorageGetCredentialsResponse A response with new credentials
// swagger:model storageGetCredentialsResponse
type StorageGetCredentialsResponse struct {

	// The list of active credentials on account
	Credentials []*StorageGetCredentialsResponseCredentialsItems0 `json:"credentials"`
}

// Validate validates this storage get credentials response
func (m *StorageGetCredentialsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCredentials(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StorageGetCredentialsResponse) validateCredentials(formats strfmt.Registry) error {

	if swag.IsZero(m.Credentials) { // not required
		return nil
	}

	for i := 0; i < len(m.Credentials); i++ {
		if swag.IsZero(m.Credentials[i]) { // not required
			continue
		}

		if m.Credentials[i] != nil {
			if err := m.Credentials[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("credentials" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *StorageGetCredentialsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StorageGetCredentialsResponse) UnmarshalBinary(b []byte) error {
	var res StorageGetCredentialsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// StorageGetCredentialsResponseCredentialsItems0 Storage credentials for a user
// swagger:model StorageGetCredentialsResponseCredentialsItems0
type StorageGetCredentialsResponseCredentialsItems0 struct {

	// The ID for the access key
	AccessKey string `json:"accessKey,omitempty"`
}

// Validate validates this storage get credentials response credentials items0
func (m *StorageGetCredentialsResponseCredentialsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StorageGetCredentialsResponseCredentialsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StorageGetCredentialsResponseCredentialsItems0) UnmarshalBinary(b []byte) error {
	var res StorageGetCredentialsResponseCredentialsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
