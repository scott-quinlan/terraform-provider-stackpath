/*
 * Object Storage
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package api_client

// StorageBucketVisibility - PRIVATE: The bucket is private and only accessibly with credentials  - PUBLIC: The bucket is public and accessible over the internet
type StorageBucketVisibility string

// List of storageBucketVisibility
const (
	PRIVATE StorageBucketVisibility = "PRIVATE"
	PUBLIC  StorageBucketVisibility = "PUBLIC"
)
