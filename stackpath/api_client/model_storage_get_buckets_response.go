/*
 * Object Storage
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package api_client

// StorageGetBucketsResponse The buckets for the given stack
type StorageGetBucketsResponse struct {
	PageInfo InlineResponse200PageInfo `json:"pageInfo,omitempty"`
	// The requested buckets
	Results []InlineResponse200Results `json:"results,omitempty"`
}
