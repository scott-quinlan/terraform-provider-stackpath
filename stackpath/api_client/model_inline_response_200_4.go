/*
 * Object Storage
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package api_client

// InlineResponse2004 A collection of metrics
type InlineResponse2004 struct {
	// A metrics query's resulting status
	Status string                 `json:"status,omitempty"`
	Data   InlineResponse2004Data `json:"data,omitempty"`
	// The type of error encountered when querying for metrics
	ErrorType string `json:"errorType,omitempty"`
	// The error encountered when querying for metrics
	Error string `json:"error,omitempty"`
	// Warnings encountered when querying for metrics
	Warnings []string `json:"warnings,omitempty"`
}
