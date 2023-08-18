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

// checks if the CollectionResponseWithTotalVersionBlogPost type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CollectionResponseWithTotalVersionBlogPost{}

// CollectionResponseWithTotalVersionBlogPost Response object for collections of blog post versions with pagination information.
type CollectionResponseWithTotalVersionBlogPost struct {
	// Total number of blog post versions.
	Total int32 `json:"total"`
	// Collection of blog post versions.
	Results []VersionBlogPost `json:"results"`
	Paging *Paging `json:"paging,omitempty"`
}

// NewCollectionResponseWithTotalVersionBlogPost instantiates a new CollectionResponseWithTotalVersionBlogPost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionResponseWithTotalVersionBlogPost(total int32, results []VersionBlogPost) *CollectionResponseWithTotalVersionBlogPost {
	this := CollectionResponseWithTotalVersionBlogPost{}
	this.Total = total
	this.Results = results
	return &this
}

// NewCollectionResponseWithTotalVersionBlogPostWithDefaults instantiates a new CollectionResponseWithTotalVersionBlogPost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionResponseWithTotalVersionBlogPostWithDefaults() *CollectionResponseWithTotalVersionBlogPost {
	this := CollectionResponseWithTotalVersionBlogPost{}
	return &this
}

// GetTotal returns the Total field value
func (o *CollectionResponseWithTotalVersionBlogPost) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *CollectionResponseWithTotalVersionBlogPost) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *CollectionResponseWithTotalVersionBlogPost) SetTotal(v int32) {
	o.Total = v
}

// GetResults returns the Results field value
func (o *CollectionResponseWithTotalVersionBlogPost) GetResults() []VersionBlogPost {
	if o == nil {
		var ret []VersionBlogPost
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *CollectionResponseWithTotalVersionBlogPost) GetResultsOk() ([]VersionBlogPost, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *CollectionResponseWithTotalVersionBlogPost) SetResults(v []VersionBlogPost) {
	o.Results = v
}

// GetPaging returns the Paging field value if set, zero value otherwise.
func (o *CollectionResponseWithTotalVersionBlogPost) GetPaging() Paging {
	if o == nil || IsNil(o.Paging) {
		var ret Paging
		return ret
	}
	return *o.Paging
}

// GetPagingOk returns a tuple with the Paging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionResponseWithTotalVersionBlogPost) GetPagingOk() (*Paging, bool) {
	if o == nil || IsNil(o.Paging) {
		return nil, false
	}
	return o.Paging, true
}

// HasPaging returns a boolean if a field has been set.
func (o *CollectionResponseWithTotalVersionBlogPost) HasPaging() bool {
	if o != nil && !IsNil(o.Paging) {
		return true
	}

	return false
}

// SetPaging gets a reference to the given Paging and assigns it to the Paging field.
func (o *CollectionResponseWithTotalVersionBlogPost) SetPaging(v Paging) {
	o.Paging = &v
}

func (o CollectionResponseWithTotalVersionBlogPost) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CollectionResponseWithTotalVersionBlogPost) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["total"] = o.Total
	toSerialize["results"] = o.Results
	if !IsNil(o.Paging) {
		toSerialize["paging"] = o.Paging
	}
	return toSerialize, nil
}

type NullableCollectionResponseWithTotalVersionBlogPost struct {
	value *CollectionResponseWithTotalVersionBlogPost
	isSet bool
}

func (v NullableCollectionResponseWithTotalVersionBlogPost) Get() *CollectionResponseWithTotalVersionBlogPost {
	return v.value
}

func (v *NullableCollectionResponseWithTotalVersionBlogPost) Set(val *CollectionResponseWithTotalVersionBlogPost) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionResponseWithTotalVersionBlogPost) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionResponseWithTotalVersionBlogPost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionResponseWithTotalVersionBlogPost(val *CollectionResponseWithTotalVersionBlogPost) *NullableCollectionResponseWithTotalVersionBlogPost {
	return &NullableCollectionResponseWithTotalVersionBlogPost{value: val, isSet: true}
}

func (v NullableCollectionResponseWithTotalVersionBlogPost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionResponseWithTotalVersionBlogPost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


