/*
Custom Workflow Actions

Create custom workflow actions

API version: v4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package actions

import (
	"encoding/json"
)

// checks if the CollectionResponseActionRevisionForwardPaging type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CollectionResponseActionRevisionForwardPaging{}

// CollectionResponseActionRevisionForwardPaging struct for CollectionResponseActionRevisionForwardPaging
type CollectionResponseActionRevisionForwardPaging struct {
	Results []ActionRevision `json:"results"`
	Paging  *ForwardPaging   `json:"paging,omitempty"`
}

// NewCollectionResponseActionRevisionForwardPaging instantiates a new CollectionResponseActionRevisionForwardPaging object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionResponseActionRevisionForwardPaging(results []ActionRevision) *CollectionResponseActionRevisionForwardPaging {
	this := CollectionResponseActionRevisionForwardPaging{}
	this.Results = results
	return &this
}

// NewCollectionResponseActionRevisionForwardPagingWithDefaults instantiates a new CollectionResponseActionRevisionForwardPaging object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionResponseActionRevisionForwardPagingWithDefaults() *CollectionResponseActionRevisionForwardPaging {
	this := CollectionResponseActionRevisionForwardPaging{}
	return &this
}

// GetResults returns the Results field value
func (o *CollectionResponseActionRevisionForwardPaging) GetResults() []ActionRevision {
	if o == nil {
		var ret []ActionRevision
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *CollectionResponseActionRevisionForwardPaging) GetResultsOk() ([]ActionRevision, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *CollectionResponseActionRevisionForwardPaging) SetResults(v []ActionRevision) {
	o.Results = v
}

// GetPaging returns the Paging field value if set, zero value otherwise.
func (o *CollectionResponseActionRevisionForwardPaging) GetPaging() ForwardPaging {
	if o == nil || IsNil(o.Paging) {
		var ret ForwardPaging
		return ret
	}
	return *o.Paging
}

// GetPagingOk returns a tuple with the Paging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionResponseActionRevisionForwardPaging) GetPagingOk() (*ForwardPaging, bool) {
	if o == nil || IsNil(o.Paging) {
		return nil, false
	}
	return o.Paging, true
}

// HasPaging returns a boolean if a field has been set.
func (o *CollectionResponseActionRevisionForwardPaging) HasPaging() bool {
	if o != nil && !IsNil(o.Paging) {
		return true
	}

	return false
}

// SetPaging gets a reference to the given ForwardPaging and assigns it to the Paging field.
func (o *CollectionResponseActionRevisionForwardPaging) SetPaging(v ForwardPaging) {
	o.Paging = &v
}

func (o CollectionResponseActionRevisionForwardPaging) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CollectionResponseActionRevisionForwardPaging) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["results"] = o.Results
	if !IsNil(o.Paging) {
		toSerialize["paging"] = o.Paging
	}
	return toSerialize, nil
}

type NullableCollectionResponseActionRevisionForwardPaging struct {
	value *CollectionResponseActionRevisionForwardPaging
	isSet bool
}

func (v NullableCollectionResponseActionRevisionForwardPaging) Get() *CollectionResponseActionRevisionForwardPaging {
	return v.value
}

func (v *NullableCollectionResponseActionRevisionForwardPaging) Set(val *CollectionResponseActionRevisionForwardPaging) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionResponseActionRevisionForwardPaging) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionResponseActionRevisionForwardPaging) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionResponseActionRevisionForwardPaging(val *CollectionResponseActionRevisionForwardPaging) *NullableCollectionResponseActionRevisionForwardPaging {
	return &NullableCollectionResponseActionRevisionForwardPaging{value: val, isSet: true}
}

func (v NullableCollectionResponseActionRevisionForwardPaging) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionResponseActionRevisionForwardPaging) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
