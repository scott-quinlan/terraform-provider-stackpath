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

// checks if the DataMatrix type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataMatrix{}

// DataMatrix A set of time series containing a range of data points over time for each time series
type DataMatrix struct {
	// A data point's value
	Results []DataMatrixResult `json:"results,omitempty"`
}

// NewDataMatrix instantiates a new DataMatrix object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataMatrix() *DataMatrix {
	this := DataMatrix{}
	return &this
}

// NewDataMatrixWithDefaults instantiates a new DataMatrix object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataMatrixWithDefaults() *DataMatrix {
	this := DataMatrix{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *DataMatrix) GetResults() []DataMatrixResult {
	if o == nil || IsNil(o.Results) {
		var ret []DataMatrixResult
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataMatrix) GetResultsOk() ([]DataMatrixResult, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *DataMatrix) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []DataMatrixResult and assigns it to the Results field.
func (o *DataMatrix) SetResults(v []DataMatrixResult) {
	o.Results = v
}

func (o DataMatrix) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataMatrix) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	return toSerialize, nil
}

type NullableDataMatrix struct {
	value *DataMatrix
	isSet bool
}

func (v NullableDataMatrix) Get() *DataMatrix {
	return v.value
}

func (v *NullableDataMatrix) Set(val *DataMatrix) {
	v.value = val
	v.isSet = true
}

func (v NullableDataMatrix) IsSet() bool {
	return v.isSet
}

func (v *NullableDataMatrix) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataMatrix(val *DataMatrix) *NullableDataMatrix {
	return &NullableDataMatrix{value: val, isSet: true}
}

func (v NullableDataMatrix) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataMatrix) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


