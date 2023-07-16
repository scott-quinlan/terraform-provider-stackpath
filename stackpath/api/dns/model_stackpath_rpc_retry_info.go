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

// checks if the StackpathRpcRetryInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StackpathRpcRetryInfo{}

// StackpathRpcRetryInfo struct for StackpathRpcRetryInfo
type StackpathRpcRetryInfo struct {
	ApiStatusDetail
	RetryDelay *string `json:"retryDelay,omitempty"`
}

// NewStackpathRpcRetryInfo instantiates a new StackpathRpcRetryInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStackpathRpcRetryInfo(type_ string) *StackpathRpcRetryInfo {
	this := StackpathRpcRetryInfo{}
	this.Type = type_
	return &this
}

// NewStackpathRpcRetryInfoWithDefaults instantiates a new StackpathRpcRetryInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStackpathRpcRetryInfoWithDefaults() *StackpathRpcRetryInfo {
	this := StackpathRpcRetryInfo{}
	return &this
}

// GetRetryDelay returns the RetryDelay field value if set, zero value otherwise.
func (o *StackpathRpcRetryInfo) GetRetryDelay() string {
	if o == nil || IsNil(o.RetryDelay) {
		var ret string
		return ret
	}
	return *o.RetryDelay
}

// GetRetryDelayOk returns a tuple with the RetryDelay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StackpathRpcRetryInfo) GetRetryDelayOk() (*string, bool) {
	if o == nil || IsNil(o.RetryDelay) {
		return nil, false
	}
	return o.RetryDelay, true
}

// HasRetryDelay returns a boolean if a field has been set.
func (o *StackpathRpcRetryInfo) HasRetryDelay() bool {
	if o != nil && !IsNil(o.RetryDelay) {
		return true
	}

	return false
}

// SetRetryDelay gets a reference to the given string and assigns it to the RetryDelay field.
func (o *StackpathRpcRetryInfo) SetRetryDelay(v string) {
	o.RetryDelay = &v
}

func (o StackpathRpcRetryInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StackpathRpcRetryInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedApiStatusDetail, errApiStatusDetail := json.Marshal(o.ApiStatusDetail)
	if errApiStatusDetail != nil {
		return map[string]interface{}{}, errApiStatusDetail
	}
	errApiStatusDetail = json.Unmarshal([]byte(serializedApiStatusDetail), &toSerialize)
	if errApiStatusDetail != nil {
		return map[string]interface{}{}, errApiStatusDetail
	}
	if !IsNil(o.RetryDelay) {
		toSerialize["retryDelay"] = o.RetryDelay
	}
	return toSerialize, nil
}

type NullableStackpathRpcRetryInfo struct {
	value *StackpathRpcRetryInfo
	isSet bool
}

func (v NullableStackpathRpcRetryInfo) Get() *StackpathRpcRetryInfo {
	return v.value
}

func (v *NullableStackpathRpcRetryInfo) Set(val *StackpathRpcRetryInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableStackpathRpcRetryInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableStackpathRpcRetryInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStackpathRpcRetryInfo(val *StackpathRpcRetryInfo) *NullableStackpathRpcRetryInfo {
	return &NullableStackpathRpcRetryInfo{value: val, isSet: true}
}

func (v NullableStackpathRpcRetryInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStackpathRpcRetryInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


