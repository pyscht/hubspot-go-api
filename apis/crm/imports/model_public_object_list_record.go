/*
CRM Imports

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package imports

import (
	"encoding/json"
)

// checks if the PublicObjectListRecord type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PublicObjectListRecord{}

// PublicObjectListRecord A summary detailing which list contains the imported objects.
type PublicObjectListRecord struct {
	// The ID of the list containing the imported objects.
	ListId string `json:"listId"`
	// The type of object contained in the list.
	ObjectType string `json:"objectType"`
}

// NewPublicObjectListRecord instantiates a new PublicObjectListRecord object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicObjectListRecord(listId string, objectType string) *PublicObjectListRecord {
	this := PublicObjectListRecord{}
	this.ListId = listId
	this.ObjectType = objectType
	return &this
}

// NewPublicObjectListRecordWithDefaults instantiates a new PublicObjectListRecord object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicObjectListRecordWithDefaults() *PublicObjectListRecord {
	this := PublicObjectListRecord{}
	return &this
}

// GetListId returns the ListId field value
func (o *PublicObjectListRecord) GetListId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ListId
}

// GetListIdOk returns a tuple with the ListId field value
// and a boolean to check if the value has been set.
func (o *PublicObjectListRecord) GetListIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ListId, true
}

// SetListId sets field value
func (o *PublicObjectListRecord) SetListId(v string) {
	o.ListId = v
}

// GetObjectType returns the ObjectType field value
func (o *PublicObjectListRecord) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *PublicObjectListRecord) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *PublicObjectListRecord) SetObjectType(v string) {
	o.ObjectType = v
}

func (o PublicObjectListRecord) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PublicObjectListRecord) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["listId"] = o.ListId
	toSerialize["objectType"] = o.ObjectType
	return toSerialize, nil
}

type NullablePublicObjectListRecord struct {
	value *PublicObjectListRecord
	isSet bool
}

func (v NullablePublicObjectListRecord) Get() *PublicObjectListRecord {
	return v.value
}

func (v *NullablePublicObjectListRecord) Set(val *PublicObjectListRecord) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicObjectListRecord) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicObjectListRecord) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicObjectListRecord(val *PublicObjectListRecord) *NullablePublicObjectListRecord {
	return &NullablePublicObjectListRecord{value: val, isSet: true}
}

func (v NullablePublicObjectListRecord) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicObjectListRecord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}