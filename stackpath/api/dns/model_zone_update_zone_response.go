/*
DNS

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dns

import (
	"encoding/json"
)

// checks if the ZoneUpdateZoneResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ZoneUpdateZoneResponse{}

// ZoneUpdateZoneResponse A response from a request to replace a DNS zone
type ZoneUpdateZoneResponse struct {
	Zone *ZoneZone `json:"zone,omitempty"`
}

// NewZoneUpdateZoneResponse instantiates a new ZoneUpdateZoneResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZoneUpdateZoneResponse() *ZoneUpdateZoneResponse {
	this := ZoneUpdateZoneResponse{}
	return &this
}

// NewZoneUpdateZoneResponseWithDefaults instantiates a new ZoneUpdateZoneResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewZoneUpdateZoneResponseWithDefaults() *ZoneUpdateZoneResponse {
	this := ZoneUpdateZoneResponse{}
	return &this
}

// GetZone returns the Zone field value if set, zero value otherwise.
func (o *ZoneUpdateZoneResponse) GetZone() ZoneZone {
	if o == nil || IsNil(o.Zone) {
		var ret ZoneZone
		return ret
	}
	return *o.Zone
}

// GetZoneOk returns a tuple with the Zone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneUpdateZoneResponse) GetZoneOk() (*ZoneZone, bool) {
	if o == nil || IsNil(o.Zone) {
		return nil, false
	}
	return o.Zone, true
}

// HasZone returns a boolean if a field has been set.
func (o *ZoneUpdateZoneResponse) HasZone() bool {
	if o != nil && !IsNil(o.Zone) {
		return true
	}

	return false
}

// SetZone gets a reference to the given ZoneZone and assigns it to the Zone field.
func (o *ZoneUpdateZoneResponse) SetZone(v ZoneZone) {
	o.Zone = &v
}

func (o ZoneUpdateZoneResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ZoneUpdateZoneResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Zone) {
		toSerialize["zone"] = o.Zone
	}
	return toSerialize, nil
}

type NullableZoneUpdateZoneResponse struct {
	value *ZoneUpdateZoneResponse
	isSet bool
}

func (v NullableZoneUpdateZoneResponse) Get() *ZoneUpdateZoneResponse {
	return v.value
}

func (v *NullableZoneUpdateZoneResponse) Set(val *ZoneUpdateZoneResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableZoneUpdateZoneResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableZoneUpdateZoneResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZoneUpdateZoneResponse(val *ZoneUpdateZoneResponse) *NullableZoneUpdateZoneResponse {
	return &NullableZoneUpdateZoneResponse{value: val, isSet: true}
}

func (v NullableZoneUpdateZoneResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZoneUpdateZoneResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

