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

// checks if the SideOrCorner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SideOrCorner{}

// SideOrCorner struct for SideOrCorner
type SideOrCorner struct {
	VerticalSide string `json:"verticalSide"`
	HorizontalSide string `json:"horizontalSide"`
}

// NewSideOrCorner instantiates a new SideOrCorner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSideOrCorner(verticalSide string, horizontalSide string) *SideOrCorner {
	this := SideOrCorner{}
	this.VerticalSide = verticalSide
	this.HorizontalSide = horizontalSide
	return &this
}

// NewSideOrCornerWithDefaults instantiates a new SideOrCorner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSideOrCornerWithDefaults() *SideOrCorner {
	this := SideOrCorner{}
	return &this
}

// GetVerticalSide returns the VerticalSide field value
func (o *SideOrCorner) GetVerticalSide() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerticalSide
}

// GetVerticalSideOk returns a tuple with the VerticalSide field value
// and a boolean to check if the value has been set.
func (o *SideOrCorner) GetVerticalSideOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerticalSide, true
}

// SetVerticalSide sets field value
func (o *SideOrCorner) SetVerticalSide(v string) {
	o.VerticalSide = v
}

// GetHorizontalSide returns the HorizontalSide field value
func (o *SideOrCorner) GetHorizontalSide() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HorizontalSide
}

// GetHorizontalSideOk returns a tuple with the HorizontalSide field value
// and a boolean to check if the value has been set.
func (o *SideOrCorner) GetHorizontalSideOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HorizontalSide, true
}

// SetHorizontalSide sets field value
func (o *SideOrCorner) SetHorizontalSide(v string) {
	o.HorizontalSide = v
}

func (o SideOrCorner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SideOrCorner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["verticalSide"] = o.VerticalSide
	toSerialize["horizontalSide"] = o.HorizontalSide
	return toSerialize, nil
}

type NullableSideOrCorner struct {
	value *SideOrCorner
	isSet bool
}

func (v NullableSideOrCorner) Get() *SideOrCorner {
	return v.value
}

func (v *NullableSideOrCorner) Set(val *SideOrCorner) {
	v.value = val
	v.isSet = true
}

func (v NullableSideOrCorner) IsSet() bool {
	return v.isSet
}

func (v *NullableSideOrCorner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSideOrCorner(val *SideOrCorner) *NullableSideOrCorner {
	return &NullableSideOrCorner{value: val, isSet: true}
}

func (v NullableSideOrCorner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSideOrCorner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

