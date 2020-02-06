/*
 * Object Storage
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package api_client

// StorageUpdateBucketRequest struct for StorageUpdateBucketRequest
type StorageUpdateBucketRequest struct {
	// - PRIVATE: The bucket is private and only accessibly with credentials  - PUBLIC: The bucket is public and accessible over the internet
	Visibility string `json:"visibility,omitempty"`
}
