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

// checks if the StackpathRpcQuotaFailureViolation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StackpathRpcQuotaFailureViolation{}

// StackpathRpcQuotaFailureViolation struct for StackpathRpcQuotaFailureViolation
type StackpathRpcQuotaFailureViolation struct {
	Subject *string `json:"subject,omitempty"`
	Description *string `json:"description,omitempty"`
}

// NewStackpathRpcQuotaFailureViolation instantiates a new StackpathRpcQuotaFailureViolation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStackpathRpcQuotaFailureViolation() *StackpathRpcQuotaFailureViolation {
	this := StackpathRpcQuotaFailureViolation{}
	return &this
}

// NewStackpathRpcQuotaFailureViolationWithDefaults instantiates a new StackpathRpcQuotaFailureViolation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStackpathRpcQuotaFailureViolationWithDefaults() *StackpathRpcQuotaFailureViolation {
	this := StackpathRpcQuotaFailureViolation{}
	return &this
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *StackpathRpcQuotaFailureViolation) GetSubject() string {
	if o == nil || IsNil(o.Subject) {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StackpathRpcQuotaFailureViolation) GetSubjectOk() (*string, bool) {
	if o == nil || IsNil(o.Subject) {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *StackpathRpcQuotaFailureViolation) HasSubject() bool {
	if o != nil && !IsNil(o.Subject) {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *StackpathRpcQuotaFailureViolation) SetSubject(v string) {
	o.Subject = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *StackpathRpcQuotaFailureViolation) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StackpathRpcQuotaFailureViolation) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *StackpathRpcQuotaFailureViolation) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *StackpathRpcQuotaFailureViolation) SetDescription(v string) {
	o.Description = &v
}

func (o StackpathRpcQuotaFailureViolation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StackpathRpcQuotaFailureViolation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Subject) {
		toSerialize["subject"] = o.Subject
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

type NullableStackpathRpcQuotaFailureViolation struct {
	value *StackpathRpcQuotaFailureViolation
	isSet bool
}

func (v NullableStackpathRpcQuotaFailureViolation) Get() *StackpathRpcQuotaFailureViolation {
	return v.value
}

func (v *NullableStackpathRpcQuotaFailureViolation) Set(val *StackpathRpcQuotaFailureViolation) {
	v.value = val
	v.isSet = true
}

func (v NullableStackpathRpcQuotaFailureViolation) IsSet() bool {
	return v.isSet
}

func (v *NullableStackpathRpcQuotaFailureViolation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStackpathRpcQuotaFailureViolation(val *StackpathRpcQuotaFailureViolation) *NullableStackpathRpcQuotaFailureViolation {
	return &NullableStackpathRpcQuotaFailureViolation{value: val, isSet: true}
}

func (v NullableStackpathRpcQuotaFailureViolation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStackpathRpcQuotaFailureViolation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


