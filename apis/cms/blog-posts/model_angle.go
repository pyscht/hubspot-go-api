/*
Blog Post endpoints

Use these endpoints for interacting with Blog Posts, Blog Authors, and Blog Tags

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package blog-posts

import (
	"encoding/json"
)

// checks if the Angle type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Angle{}

// Angle struct for Angle
type Angle struct {
	Value float32 `json:"value"`
	Units string `json:"units"`
}

// NewAngle instantiates a new Angle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAngle(value float32, units string) *Angle {
	this := Angle{}
	this.Value = value
	this.Units = units
	return &this
}

// NewAngleWithDefaults instantiates a new Angle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAngleWithDefaults() *Angle {
	this := Angle{}
	return &this
}

// GetValue returns the Value field value
func (o *Angle) GetValue() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *Angle) GetValueOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *Angle) SetValue(v float32) {
	o.Value = v
}

// GetUnits returns the Units field value
func (o *Angle) GetUnits() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Units
}

// GetUnitsOk returns a tuple with the Units field value
// and a boolean to check if the value has been set.
func (o *Angle) GetUnitsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Units, true
}

// SetUnits sets field value
func (o *Angle) SetUnits(v string) {
	o.Units = v
}

func (o Angle) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Angle) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["value"] = o.Value
	toSerialize["units"] = o.Units
	return toSerialize, nil
}

type NullableAngle struct {
	value *Angle
	isSet bool
}

func (v NullableAngle) Get() *Angle {
	return v.value
}

func (v *NullableAngle) Set(val *Angle) {
	v.value = val
	v.isSet = true
}

func (v NullableAngle) IsSet() bool {
	return v.isSet
}

func (v *NullableAngle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAngle(val *Angle) *NullableAngle {
	return &NullableAngle{value: val, isSet: true}
}

func (v NullableAngle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAngle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


