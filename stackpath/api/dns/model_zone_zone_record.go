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

// checks if the ZoneZoneRecord type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ZoneZoneRecord{}

// ZoneZoneRecord A DNS zone's resource record  A zone record describes an individual piece of DNS functionality in a DNS zone.
type ZoneZoneRecord struct {
	// A zone record's unique ID
	Id *string `json:"id,omitempty"`
	// The ID of the zone that a zone record belongs to
	ZoneId *string `json:"zoneId,omitempty"`
	// A zone record's name  Use the value \"@\" to denote current root domain name.
	Name *string `json:"name,omitempty"`
	// A zone record's type  Zone record types describe the zone record's behavior. For instance, a zone record's type can say that the record is a name to IP address value, a name alias, or which mail exchanger is responsible for the domain. See https://support.stackpath.com/hc/en-us/articles/360001085563-What-DNS-record-types-does-StackPath-support for more information.
	Type *string `json:"type,omitempty"`
	// A zone record's class code  This is typically \"IN\" for Internet related resource records.
	Class *string `json:"class,omitempty"`
	// A zone record's time to live  A record's TTL is the number of seconds that the record should be cached by DNS resolvers. Use lower TTL values if you expect zone records to change often. Use higher TTL values for records that won't change to prevent extra DNS lookups by clients.
	Ttl *int32 `json:"ttl,omitempty"`
	// A zone record's value  MX record data follows the format `<priority> <value>`, without `<>`s.
	Data *string `json:"data,omitempty"`
	// A zone record's priority  A resource record is replicated in StackPath's DNS infrastructure the number of times of the record's weight, giving it a more likely response to queries if a zone has records with the same name and type.
	Weight *int32 `json:"weight,omitempty"`
	// A key/value pair of user-defined labels for a zone record  Zone record labels are not processed by StackPath and are solely used for users to organize their DNS zones.
	Labels *map[string]string `json:"labels,omitempty"`
	// The date a zone record was created
	Created *time.Time `json:"created,omitempty"`
	// The date a zone record was last updated
	Updated *time.Time `json:"updated,omitempty"`
}

// NewZoneZoneRecord instantiates a new ZoneZoneRecord object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZoneZoneRecord() *ZoneZoneRecord {
	this := ZoneZoneRecord{}
	return &this
}

// NewZoneZoneRecordWithDefaults instantiates a new ZoneZoneRecord object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewZoneZoneRecordWithDefaults() *ZoneZoneRecord {
	this := ZoneZoneRecord{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ZoneZoneRecord) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneZoneRecord) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ZoneZoneRecord) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ZoneZoneRecord) SetId(v string) {
	o.Id = &v
}

// GetZoneId returns the ZoneId field value if set, zero value otherwise.
func (o *ZoneZoneRecord) GetZoneId() string {
	if o == nil || IsNil(o.ZoneId) {
		var ret string
		return ret
	}
	return *o.ZoneId
}

// GetZoneIdOk returns a tuple with the ZoneId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneZoneRecord) GetZoneIdOk() (*string, bool) {
	if o == nil || IsNil(o.ZoneId) {
		return nil, false
	}
	return o.ZoneId, true
}

// HasZoneId returns a boolean if a field has been set.
func (o *ZoneZoneRecord) HasZoneId() bool {
	if o != nil && !IsNil(o.ZoneId) {
		return true
	}

	return false
}

// SetZoneId gets a reference to the given string and assigns it to the ZoneId field.
func (o *ZoneZoneRecord) SetZoneId(v string) {
	o.ZoneId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ZoneZoneRecord) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneZoneRecord) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ZoneZoneRecord) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ZoneZoneRecord) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ZoneZoneRecord) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneZoneRecord) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ZoneZoneRecord) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ZoneZoneRecord) SetType(v string) {
	o.Type = &v
}

// GetClass returns the Class field value if set, zero value otherwise.
func (o *ZoneZoneRecord) GetClass() string {
	if o == nil || IsNil(o.Class) {
		var ret string
		return ret
	}
	return *o.Class
}

// GetClassOk returns a tuple with the Class field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneZoneRecord) GetClassOk() (*string, bool) {
	if o == nil || IsNil(o.Class) {
		return nil, false
	}
	return o.Class, true
}

// HasClass returns a boolean if a field has been set.
func (o *ZoneZoneRecord) HasClass() bool {
	if o != nil && !IsNil(o.Class) {
		return true
	}

	return false
}

// SetClass gets a reference to the given string and assigns it to the Class field.
func (o *ZoneZoneRecord) SetClass(v string) {
	o.Class = &v
}

// GetTtl returns the Ttl field value if set, zero value otherwise.
func (o *ZoneZoneRecord) GetTtl() int32 {
	if o == nil || IsNil(o.Ttl) {
		var ret int32
		return ret
	}
	return *o.Ttl
}

// GetTtlOk returns a tuple with the Ttl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneZoneRecord) GetTtlOk() (*int32, bool) {
	if o == nil || IsNil(o.Ttl) {
		return nil, false
	}
	return o.Ttl, true
}

// HasTtl returns a boolean if a field has been set.
func (o *ZoneZoneRecord) HasTtl() bool {
	if o != nil && !IsNil(o.Ttl) {
		return true
	}

	return false
}

// SetTtl gets a reference to the given int32 and assigns it to the Ttl field.
func (o *ZoneZoneRecord) SetTtl(v int32) {
	o.Ttl = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ZoneZoneRecord) GetData() string {
	if o == nil || IsNil(o.Data) {
		var ret string
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneZoneRecord) GetDataOk() (*string, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ZoneZoneRecord) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given string and assigns it to the Data field.
func (o *ZoneZoneRecord) SetData(v string) {
	o.Data = &v
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *ZoneZoneRecord) GetWeight() int32 {
	if o == nil || IsNil(o.Weight) {
		var ret int32
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneZoneRecord) GetWeightOk() (*int32, bool) {
	if o == nil || IsNil(o.Weight) {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *ZoneZoneRecord) HasWeight() bool {
	if o != nil && !IsNil(o.Weight) {
		return true
	}

	return false
}

// SetWeight gets a reference to the given int32 and assigns it to the Weight field.
func (o *ZoneZoneRecord) SetWeight(v int32) {
	o.Weight = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *ZoneZoneRecord) GetLabels() map[string]string {
	if o == nil || IsNil(o.Labels) {
		var ret map[string]string
		return ret
	}
	return *o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneZoneRecord) GetLabelsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *ZoneZoneRecord) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given map[string]string and assigns it to the Labels field.
func (o *ZoneZoneRecord) SetLabels(v map[string]string) {
	o.Labels = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *ZoneZoneRecord) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneZoneRecord) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *ZoneZoneRecord) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *ZoneZoneRecord) SetCreated(v time.Time) {
	o.Created = &v
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *ZoneZoneRecord) GetUpdated() time.Time {
	if o == nil || IsNil(o.Updated) {
		var ret time.Time
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneZoneRecord) GetUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Updated) {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *ZoneZoneRecord) HasUpdated() bool {
	if o != nil && !IsNil(o.Updated) {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given time.Time and assigns it to the Updated field.
func (o *ZoneZoneRecord) SetUpdated(v time.Time) {
	o.Updated = &v
}

func (o ZoneZoneRecord) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ZoneZoneRecord) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.ZoneId) {
		toSerialize["zoneId"] = o.ZoneId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Class) {
		toSerialize["class"] = o.Class
	}
	if !IsNil(o.Ttl) {
		toSerialize["ttl"] = o.Ttl
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Weight) {
		toSerialize["weight"] = o.Weight
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
	return toSerialize, nil
}

type NullableZoneZoneRecord struct {
	value *ZoneZoneRecord
	isSet bool
}

func (v NullableZoneZoneRecord) Get() *ZoneZoneRecord {
	return v.value
}

func (v *NullableZoneZoneRecord) Set(val *ZoneZoneRecord) {
	v.value = val
	v.isSet = true
}

func (v NullableZoneZoneRecord) IsSet() bool {
	return v.isSet
}

func (v *NullableZoneZoneRecord) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZoneZoneRecord(val *ZoneZoneRecord) *NullableZoneZoneRecord {
	return &NullableZoneZoneRecord{value: val, isSet: true}
}

func (v NullableZoneZoneRecord) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZoneZoneRecord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


