/*
 * Object Storage
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package api_client

// DataMatrix A set of time series containing a range of data points over time for each time series
type DataMatrix struct {
	// A data point's value
	Results []InlineResponse2004DataMatrixResults `json:"results,omitempty"`
}
