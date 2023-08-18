/*
CRM Objects

CRM objects such as companies, contacts, deals, line items, products, tickets, and quotes are standard objects in HubSpot’s CRM. These core building blocks support custom properties, store critical information, and play a central role in the HubSpot application.  ## Supported Object Types  This API provides access to collections of CRM objects, which return a map of property names to values. Each object type has its own set of default properties, which can be found by exploring the [CRM Object Properties API](https://developers.hubspot.com/docs/methods/crm-properties/crm-properties-overview).  |Object Type |Properties returned by default | |--|--| | `companies` | `name`, `domain` | | `contacts` | `firstname`, `lastname`, `email` | | `deals` | `dealname`, `amount`, `closedate`, `pipeline`, `dealstage` | | `products` | `name`, `description`, `price` | | `tickets` | `content`, `hs_pipeline`, `hs_pipeline_stage`, `hs_ticket_category`, `hs_ticket_priority`, `subject` |  Find a list of all properties for an object type using the [CRM Object Properties](https://developers.hubspot.com/docs/methods/crm-properties/get-properties) API. e.g. `GET https://api.hubapi.com/properties/v2/companies/properties`. Change the properties returned in the response using the `properties` array in the request body.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package objects

import (
	"encoding/json"
)

// checks if the PublicAssociationsForObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PublicAssociationsForObject{}

// PublicAssociationsForObject struct for PublicAssociationsForObject
type PublicAssociationsForObject struct {
	To    PublicObjectId    `json:"to"`
	Types []AssociationSpec `json:"types"`
}

// NewPublicAssociationsForObject instantiates a new PublicAssociationsForObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicAssociationsForObject(to PublicObjectId, types []AssociationSpec) *PublicAssociationsForObject {
	this := PublicAssociationsForObject{}
	this.To = to
	this.Types = types
	return &this
}

// NewPublicAssociationsForObjectWithDefaults instantiates a new PublicAssociationsForObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicAssociationsForObjectWithDefaults() *PublicAssociationsForObject {
	this := PublicAssociationsForObject{}
	return &this
}

// GetTo returns the To field value
func (o *PublicAssociationsForObject) GetTo() PublicObjectId {
	if o == nil {
		var ret PublicObjectId
		return ret
	}

	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *PublicAssociationsForObject) GetToOk() (*PublicObjectId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.To, true
}

// SetTo sets field value
func (o *PublicAssociationsForObject) SetTo(v PublicObjectId) {
	o.To = v
}

// GetTypes returns the Types field value
func (o *PublicAssociationsForObject) GetTypes() []AssociationSpec {
	if o == nil {
		var ret []AssociationSpec
		return ret
	}

	return o.Types
}

// GetTypesOk returns a tuple with the Types field value
// and a boolean to check if the value has been set.
func (o *PublicAssociationsForObject) GetTypesOk() ([]AssociationSpec, bool) {
	if o == nil {
		return nil, false
	}
	return o.Types, true
}

// SetTypes sets field value
func (o *PublicAssociationsForObject) SetTypes(v []AssociationSpec) {
	o.Types = v
}

func (o PublicAssociationsForObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PublicAssociationsForObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["to"] = o.To
	toSerialize["types"] = o.Types
	return toSerialize, nil
}

type NullablePublicAssociationsForObject struct {
	value *PublicAssociationsForObject
	isSet bool
}

func (v NullablePublicAssociationsForObject) Get() *PublicAssociationsForObject {
	return v.value
}

func (v *NullablePublicAssociationsForObject) Set(val *PublicAssociationsForObject) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicAssociationsForObject) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicAssociationsForObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicAssociationsForObject(val *PublicAssociationsForObject) *NullablePublicAssociationsForObject {
	return &NullablePublicAssociationsForObject{value: val, isSet: true}
}

func (v NullablePublicAssociationsForObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicAssociationsForObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
