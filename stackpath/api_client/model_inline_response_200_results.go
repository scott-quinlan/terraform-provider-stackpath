/*
 * Object Storage
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package api_client

import (
	"time"
)

// InlineResponse200Results struct for InlineResponse200Results
type InlineResponse200Results struct {
	// The ID for the bucket
	Id string `json:"id,omitempty"`
	// The name of the bucket
	Label string `json:"label,omitempty"`
	// The URL used to access the bucket
	EndpointUrl string `json:"endpointUrl,omitempty"`
	// - PRIVATE: The bucket is private and only accessibly with credentials  - PUBLIC: The bucket is public and accessible over the internet
	Visibility string `json:"visibility,omitempty"`
	// The date when the bucket was created
	CreatedAt time.Time `json:"createdAt,omitempty"`
	// The date when the bucket was last updated
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	// The region in which the bucket is created. Available regions are: us-east-1, us-west-1, eu-central-1
	Region string `json:"region,omitempty"`
}
