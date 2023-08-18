/*
Business Unit

Retrieve Business Unit information.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package business_units

import (
	"encoding/json"
)

// checks if the PublicBusinessUnit type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PublicBusinessUnit{}

// PublicBusinessUnit A Business Unit
type PublicBusinessUnit struct {
	// The Business Unit's unique ID
	Id string `json:"id"`
	// The Business Unit's name
	Name         string                          `json:"name"`
	LogoMetadata *PublicBusinessUnitLogoMetadata `json:"logoMetadata,omitempty"`
}

// NewPublicBusinessUnit instantiates a new PublicBusinessUnit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicBusinessUnit(id string, name string) *PublicBusinessUnit {
	this := PublicBusinessUnit{}
	this.Id = id
	this.Name = name
	return &this
}

// NewPublicBusinessUnitWithDefaults instantiates a new PublicBusinessUnit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicBusinessUnitWithDefaults() *PublicBusinessUnit {
	this := PublicBusinessUnit{}
	return &this
}

// GetId returns the Id field value
func (o *PublicBusinessUnit) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PublicBusinessUnit) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PublicBusinessUnit) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *PublicBusinessUnit) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PublicBusinessUnit) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PublicBusinessUnit) SetName(v string) {
	o.Name = v
}

// GetLogoMetadata returns the LogoMetadata field value if set, zero value otherwise.
func (o *PublicBusinessUnit) GetLogoMetadata() PublicBusinessUnitLogoMetadata {
	if o == nil || IsNil(o.LogoMetadata) {
		var ret PublicBusinessUnitLogoMetadata
		return ret
	}
	return *o.LogoMetadata
}

// GetLogoMetadataOk returns a tuple with the LogoMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicBusinessUnit) GetLogoMetadataOk() (*PublicBusinessUnitLogoMetadata, bool) {
	if o == nil || IsNil(o.LogoMetadata) {
		return nil, false
	}
	return o.LogoMetadata, true
}

// HasLogoMetadata returns a boolean if a field has been set.
func (o *PublicBusinessUnit) HasLogoMetadata() bool {
	if o != nil && !IsNil(o.LogoMetadata) {
		return true
	}

	return false
}

// SetLogoMetadata gets a reference to the given PublicBusinessUnitLogoMetadata and assigns it to the LogoMetadata field.
func (o *PublicBusinessUnit) SetLogoMetadata(v PublicBusinessUnitLogoMetadata) {
	o.LogoMetadata = &v
}

func (o PublicBusinessUnit) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PublicBusinessUnit) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	if !IsNil(o.LogoMetadata) {
		toSerialize["logoMetadata"] = o.LogoMetadata
	}
	return toSerialize, nil
}

type NullablePublicBusinessUnit struct {
	value *PublicBusinessUnit
	isSet bool
}

func (v NullablePublicBusinessUnit) Get() *PublicBusinessUnit {
	return v.value
}

func (v *NullablePublicBusinessUnit) Set(val *PublicBusinessUnit) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicBusinessUnit) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicBusinessUnit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicBusinessUnit(val *PublicBusinessUnit) *NullablePublicBusinessUnit {
	return &NullablePublicBusinessUnit{value: val, isSet: true}
}

func (v NullablePublicBusinessUnit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicBusinessUnit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
