/*
DNS

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dns

import (
	"encoding/json"
	"fmt"
)

// ZoneDnsProvider The types of DNS providers that StackPath can detect  DNS provider types can trigger different logic when StackPath scans a domain
type ZoneDnsProvider string

// List of zoneDnsProvider
const (
	GENERAL ZoneDnsProvider = "GENERAL"
	CLOUDFLARE ZoneDnsProvider = "CLOUDFLARE"
)

// All allowed values of ZoneDnsProvider enum
var AllowedZoneDnsProviderEnumValues = []ZoneDnsProvider{
	"GENERAL",
	"CLOUDFLARE",
}

func (v *ZoneDnsProvider) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ZoneDnsProvider(value)
	for _, existing := range AllowedZoneDnsProviderEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ZoneDnsProvider", value)
}

// NewZoneDnsProviderFromValue returns a pointer to a valid ZoneDnsProvider
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewZoneDnsProviderFromValue(v string) (*ZoneDnsProvider, error) {
	ev := ZoneDnsProvider(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ZoneDnsProvider: valid values are %v", v, AllowedZoneDnsProviderEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ZoneDnsProvider) IsValid() bool {
	for _, existing := range AllowedZoneDnsProviderEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to zoneDnsProvider value
func (v ZoneDnsProvider) Ptr() *ZoneDnsProvider {
	return &v
}

type NullableZoneDnsProvider struct {
	value *ZoneDnsProvider
	isSet bool
}

func (v NullableZoneDnsProvider) Get() *ZoneDnsProvider {
	return v.value
}

func (v *NullableZoneDnsProvider) Set(val *ZoneDnsProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableZoneDnsProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableZoneDnsProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZoneDnsProvider(val *ZoneDnsProvider) *NullableZoneDnsProvider {
	return &NullableZoneDnsProvider{value: val, isSet: true}
}

func (v NullableZoneDnsProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZoneDnsProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

