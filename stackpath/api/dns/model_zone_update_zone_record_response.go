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

// checks if the ZoneUpdateZoneRecordResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ZoneUpdateZoneRecordResponse{}

// ZoneUpdateZoneRecordResponse A response from a request to replace a DNS zone resource record
type ZoneUpdateZoneRecordResponse struct {
	Record *ZoneZoneRecord `json:"record,omitempty"`
}

// NewZoneUpdateZoneRecordResponse instantiates a new ZoneUpdateZoneRecordResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZoneUpdateZoneRecordResponse() *ZoneUpdateZoneRecordResponse {
	this := ZoneUpdateZoneRecordResponse{}
	return &this
}

// NewZoneUpdateZoneRecordResponseWithDefaults instantiates a new ZoneUpdateZoneRecordResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewZoneUpdateZoneRecordResponseWithDefaults() *ZoneUpdateZoneRecordResponse {
	this := ZoneUpdateZoneRecordResponse{}
	return &this
}

// GetRecord returns the Record field value if set, zero value otherwise.
func (o *ZoneUpdateZoneRecordResponse) GetRecord() ZoneZoneRecord {
	if o == nil || IsNil(o.Record) {
		var ret ZoneZoneRecord
		return ret
	}
	return *o.Record
}

// GetRecordOk returns a tuple with the Record field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneUpdateZoneRecordResponse) GetRecordOk() (*ZoneZoneRecord, bool) {
	if o == nil || IsNil(o.Record) {
		return nil, false
	}
	return o.Record, true
}

// HasRecord returns a boolean if a field has been set.
func (o *ZoneUpdateZoneRecordResponse) HasRecord() bool {
	if o != nil && !IsNil(o.Record) {
		return true
	}

	return false
}

// SetRecord gets a reference to the given ZoneZoneRecord and assigns it to the Record field.
func (o *ZoneUpdateZoneRecordResponse) SetRecord(v ZoneZoneRecord) {
	o.Record = &v
}

func (o ZoneUpdateZoneRecordResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ZoneUpdateZoneRecordResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Record) {
		toSerialize["record"] = o.Record
	}
	return toSerialize, nil
}

type NullableZoneUpdateZoneRecordResponse struct {
	value *ZoneUpdateZoneRecordResponse
	isSet bool
}

func (v NullableZoneUpdateZoneRecordResponse) Get() *ZoneUpdateZoneRecordResponse {
	return v.value
}

func (v *NullableZoneUpdateZoneRecordResponse) Set(val *ZoneUpdateZoneRecordResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableZoneUpdateZoneRecordResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableZoneUpdateZoneRecordResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZoneUpdateZoneRecordResponse(val *ZoneUpdateZoneRecordResponse) *NullableZoneUpdateZoneRecordResponse {
	return &NullableZoneUpdateZoneRecordResponse{value: val, isSet: true}
}

func (v NullableZoneUpdateZoneRecordResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZoneUpdateZoneRecordResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


