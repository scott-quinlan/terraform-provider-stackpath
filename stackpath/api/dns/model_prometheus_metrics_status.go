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

// PrometheusMetricsStatus A metrics query's resulting status
type PrometheusMetricsStatus string

// List of prometheusMetricsStatus
const (
	SUCCESS PrometheusMetricsStatus = "SUCCESS"
	ERROR PrometheusMetricsStatus = "ERROR"
)

// All allowed values of PrometheusMetricsStatus enum
var AllowedPrometheusMetricsStatusEnumValues = []PrometheusMetricsStatus{
	"SUCCESS",
	"ERROR",
}

func (v *PrometheusMetricsStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PrometheusMetricsStatus(value)
	for _, existing := range AllowedPrometheusMetricsStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PrometheusMetricsStatus", value)
}

// NewPrometheusMetricsStatusFromValue returns a pointer to a valid PrometheusMetricsStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPrometheusMetricsStatusFromValue(v string) (*PrometheusMetricsStatus, error) {
	ev := PrometheusMetricsStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PrometheusMetricsStatus: valid values are %v", v, AllowedPrometheusMetricsStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PrometheusMetricsStatus) IsValid() bool {
	for _, existing := range AllowedPrometheusMetricsStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to prometheusMetricsStatus value
func (v PrometheusMetricsStatus) Ptr() *PrometheusMetricsStatus {
	return &v
}

type NullablePrometheusMetricsStatus struct {
	value *PrometheusMetricsStatus
	isSet bool
}

func (v NullablePrometheusMetricsStatus) Get() *PrometheusMetricsStatus {
	return v.value
}

func (v *NullablePrometheusMetricsStatus) Set(val *PrometheusMetricsStatus) {
	v.value = val
	v.isSet = true
}

func (v NullablePrometheusMetricsStatus) IsSet() bool {
	return v.isSet
}

func (v *NullablePrometheusMetricsStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrometheusMetricsStatus(val *PrometheusMetricsStatus) *NullablePrometheusMetricsStatus {
	return &NullablePrometheusMetricsStatus{value: val, isSet: true}
}

func (v NullablePrometheusMetricsStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrometheusMetricsStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

