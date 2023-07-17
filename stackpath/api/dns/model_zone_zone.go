/*
DNS

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dns

import (
	"encoding/json"
	"time"
)

// checks if the ZoneZone type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ZoneZone{}

// ZoneZone A DNS zone  A zone represents an individual domain in StackPath's DNS infrastructure.
type ZoneZone struct {
	// The ID of the stack that a zone belongs to
	StackId *string `json:"stackId,omitempty"`
	// The ID of the StackPath account that owns a zone
	AccountId *string `json:"accountId,omitempty"`
	// A zone's unique ID
	Id *string `json:"id,omitempty"`
	// A zone's name
	Domain *string `json:"domain,omitempty"`
	// A zone's version number  Version numbers are incremented automatically when a zone is updated
	Version *string `json:"version,omitempty"`
	// A key/value pair of user-defined labels for a zone  Zone labels are not processed by StackPath and are solely used for users to organize their DNS zones.
	Labels *map[string]string `json:"labels,omitempty"`
	// The date a zone was created
	Created *time.Time `json:"created,omitempty"`
	// The date a zone was last updated
	Updated *time.Time `json:"updated,omitempty"`
	// The hostnames of the StackPath resolvers that host a zone  Every zone has multiple name servers assigned by StackPath upon creation for redundancy purposes.
	Nameservers []string `json:"nameservers,omitempty"`
	// The date a zone's nameservers were last audited by StackPath
	Verified *time.Time `json:"verified,omitempty"`
	Status *ZoneZoneStatus `json:"status,omitempty"`
	// Whether or not a zone has been disabled by the user
	Disabled *bool `json:"disabled,omitempty"`
}

// NewZoneZone instantiates a new ZoneZone object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZoneZone() *ZoneZone {
	this := ZoneZone{}
	var status ZoneZoneStatus = ACTIVE
	this.Status = &status
	return &this
}

// NewZoneZoneWithDefaults instantiates a new ZoneZone object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewZoneZoneWithDefaults() *ZoneZone {
	this := ZoneZone{}
	var status ZoneZoneStatus = ACTIVE
	this.Status = &status
	return &this
}

// GetStackId returns the StackId field value if set, zero value otherwise.
func (o *ZoneZone) GetStackId() string {
	if o == nil || IsNil(o.StackId) {
		var ret string
		return ret
	}
	return *o.StackId
}

// GetStackIdOk returns a tuple with the StackId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneZone) GetStackIdOk() (*string, bool) {
	if o == nil || IsNil(o.StackId) {
		return nil, false
	}
	return o.StackId, true
}

// HasStackId returns a boolean if a field has been set.
func (o *ZoneZone) HasStackId() bool {
	if o != nil && !IsNil(o.StackId) {
		return true
	}

	return false
}

// SetStackId gets a reference to the given string and assigns it to the StackId field.
func (o *ZoneZone) SetStackId(v string) {
	o.StackId = &v
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *ZoneZone) GetAccountId() string {
	if o == nil || IsNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneZone) GetAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *ZoneZone) HasAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *ZoneZone) SetAccountId(v string) {
	o.AccountId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ZoneZone) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneZone) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ZoneZone) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ZoneZone) SetId(v string) {
	o.Id = &v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *ZoneZone) GetDomain() string {
	if o == nil || IsNil(o.Domain) {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneZone) GetDomainOk() (*string, bool) {
	if o == nil || IsNil(o.Domain) {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *ZoneZone) HasDomain() bool {
	if o != nil && !IsNil(o.Domain) {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *ZoneZone) SetDomain(v string) {
	o.Domain = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ZoneZone) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneZone) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ZoneZone) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *ZoneZone) SetVersion(v string) {
	o.Version = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *ZoneZone) GetLabels() map[string]string {
	if o == nil || IsNil(o.Labels) {
		var ret map[string]string
		return ret
	}
	return *o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneZone) GetLabelsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *ZoneZone) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given map[string]string and assigns it to the Labels field.
func (o *ZoneZone) SetLabels(v map[string]string) {
	o.Labels = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *ZoneZone) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneZone) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *ZoneZone) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *ZoneZone) SetCreated(v time.Time) {
	o.Created = &v
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *ZoneZone) GetUpdated() time.Time {
	if o == nil || IsNil(o.Updated) {
		var ret time.Time
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneZone) GetUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Updated) {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *ZoneZone) HasUpdated() bool {
	if o != nil && !IsNil(o.Updated) {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given time.Time and assigns it to the Updated field.
func (o *ZoneZone) SetUpdated(v time.Time) {
	o.Updated = &v
}

// GetNameservers returns the Nameservers field value if set, zero value otherwise.
func (o *ZoneZone) GetNameservers() []string {
	if o == nil || IsNil(o.Nameservers) {
		var ret []string
		return ret
	}
	return o.Nameservers
}

// GetNameserversOk returns a tuple with the Nameservers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneZone) GetNameserversOk() ([]string, bool) {
	if o == nil || IsNil(o.Nameservers) {
		return nil, false
	}
	return o.Nameservers, true
}

// HasNameservers returns a boolean if a field has been set.
func (o *ZoneZone) HasNameservers() bool {
	if o != nil && !IsNil(o.Nameservers) {
		return true
	}

	return false
}

// SetNameservers gets a reference to the given []string and assigns it to the Nameservers field.
func (o *ZoneZone) SetNameservers(v []string) {
	o.Nameservers = v
}

// GetVerified returns the Verified field value if set, zero value otherwise.
func (o *ZoneZone) GetVerified() time.Time {
	if o == nil || IsNil(o.Verified) {
		var ret time.Time
		return ret
	}
	return *o.Verified
}

// GetVerifiedOk returns a tuple with the Verified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneZone) GetVerifiedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Verified) {
		return nil, false
	}
	return o.Verified, true
}

// HasVerified returns a boolean if a field has been set.
func (o *ZoneZone) HasVerified() bool {
	if o != nil && !IsNil(o.Verified) {
		return true
	}

	return false
}

// SetVerified gets a reference to the given time.Time and assigns it to the Verified field.
func (o *ZoneZone) SetVerified(v time.Time) {
	o.Verified = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ZoneZone) GetStatus() ZoneZoneStatus {
	if o == nil || IsNil(o.Status) {
		var ret ZoneZoneStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneZone) GetStatusOk() (*ZoneZoneStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ZoneZone) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ZoneZoneStatus and assigns it to the Status field.
func (o *ZoneZone) SetStatus(v ZoneZoneStatus) {
	o.Status = &v
}

// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *ZoneZone) GetDisabled() bool {
	if o == nil || IsNil(o.Disabled) {
		var ret bool
		return ret
	}
	return *o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneZone) GetDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Disabled) {
		return nil, false
	}
	return o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *ZoneZone) HasDisabled() bool {
	if o != nil && !IsNil(o.Disabled) {
		return true
	}

	return false
}

// SetDisabled gets a reference to the given bool and assigns it to the Disabled field.
func (o *ZoneZone) SetDisabled(v bool) {
	o.Disabled = &v
}

func (o ZoneZone) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ZoneZone) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StackId) {
		toSerialize["stackId"] = o.StackId
	}
	if !IsNil(o.AccountId) {
		toSerialize["accountId"] = o.AccountId
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Domain) {
		toSerialize["domain"] = o.Domain
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.Updated) {
		toSerialize["updated"] = o.Updated
	}
	if !IsNil(o.Nameservers) {
		toSerialize["nameservers"] = o.Nameservers
	}
	if !IsNil(o.Verified) {
		toSerialize["verified"] = o.Verified
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Disabled) {
		toSerialize["disabled"] = o.Disabled
	}
	return toSerialize, nil
}

type NullableZoneZone struct {
	value *ZoneZone
	isSet bool
}

func (v NullableZoneZone) Get() *ZoneZone {
	return v.value
}

func (v *NullableZoneZone) Set(val *ZoneZone) {
	v.value = val
	v.isSet = true
}

func (v NullableZoneZone) IsSet() bool {
	return v.isSet
}

func (v *NullableZoneZone) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZoneZone(val *ZoneZone) *NullableZoneZone {
	return &NullableZoneZone{value: val, isSet: true}
}

func (v NullableZoneZone) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZoneZone) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

