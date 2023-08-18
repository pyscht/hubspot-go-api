/*
Schemas

The CRM uses schemas to define how custom objects should store and represent information in the HubSpot CRM. Schemas define details about an object's type, properties, and associations. The schema can be uniquely identified by its **object type ID**.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schemas

import (
	"encoding/json"
)

// checks if the Option type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Option{}

// Option The options available when a property is an enumeration
type Option struct {
	// A human-readable option label that will be shown in HubSpot.
	Label string `json:"label"`
	// The internal value of the option, which must be used when setting the property value through the API.
	Value string `json:"value"`
	// A description of the option.
	Description *string `json:"description,omitempty"`
	// Options are displayed in order starting with the lowest positive integer value. Values of -1 will cause the option to be displayed after any positive values.
	DisplayOrder *int32 `json:"displayOrder,omitempty"`
	// Hidden options will not be displayed in HubSpot.
	Hidden bool `json:"hidden"`
}

// NewOption instantiates a new Option object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOption(label string, value string, hidden bool) *Option {
	this := Option{}
	this.Label = label
	this.Value = value
	this.Hidden = hidden
	return &this
}

// NewOptionWithDefaults instantiates a new Option object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptionWithDefaults() *Option {
	this := Option{}
	return &this
}

// GetLabel returns the Label field value
func (o *Option) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *Option) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *Option) SetLabel(v string) {
	o.Label = v
}

// GetValue returns the Value field value
func (o *Option) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *Option) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *Option) SetValue(v string) {
	o.Value = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Option) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Option) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Option) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Option) SetDescription(v string) {
	o.Description = &v
}

// GetDisplayOrder returns the DisplayOrder field value if set, zero value otherwise.
func (o *Option) GetDisplayOrder() int32 {
	if o == nil || IsNil(o.DisplayOrder) {
		var ret int32
		return ret
	}
	return *o.DisplayOrder
}

// GetDisplayOrderOk returns a tuple with the DisplayOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Option) GetDisplayOrderOk() (*int32, bool) {
	if o == nil || IsNil(o.DisplayOrder) {
		return nil, false
	}
	return o.DisplayOrder, true
}

// HasDisplayOrder returns a boolean if a field has been set.
func (o *Option) HasDisplayOrder() bool {
	if o != nil && !IsNil(o.DisplayOrder) {
		return true
	}

	return false
}

// SetDisplayOrder gets a reference to the given int32 and assigns it to the DisplayOrder field.
func (o *Option) SetDisplayOrder(v int32) {
	o.DisplayOrder = &v
}

// GetHidden returns the Hidden field value
func (o *Option) GetHidden() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Hidden
}

// GetHiddenOk returns a tuple with the Hidden field value
// and a boolean to check if the value has been set.
func (o *Option) GetHiddenOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hidden, true
}

// SetHidden sets field value
func (o *Option) SetHidden(v bool) {
	o.Hidden = v
}

func (o Option) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Option) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["label"] = o.Label
	toSerialize["value"] = o.Value
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.DisplayOrder) {
		toSerialize["displayOrder"] = o.DisplayOrder
	}
	toSerialize["hidden"] = o.Hidden
	return toSerialize, nil
}

type NullableOption struct {
	value *Option
	isSet bool
}

func (v NullableOption) Get() *Option {
	return v.value
}

func (v *NullableOption) Set(val *Option) {
	v.value = val
	v.isSet = true
}

func (v NullableOption) IsSet() bool {
	return v.isSet
}

func (v *NullableOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOption(val *Option) *NullableOption {
	return &NullableOption{value: val, isSet: true}
}

func (v NullableOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
