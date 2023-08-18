/*
Webhooks API

Provides a way for apps to subscribe to certain change events in HubSpot. Once configured, apps will receive event payloads containing details about the changes at a specified target URL. There can only be one target URL for receiving event notifications per app.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package webhooks

import (
	"encoding/json"
)

// checks if the SubscriptionCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionCreateRequest{}

// SubscriptionCreateRequest New webhook settings for an app.
type SubscriptionCreateRequest struct {
	// Type of event to listen for. Can be one of `create`, `delete`, `deletedForPrivacy`, or `propertyChange`.
	EventType string `json:"eventType"`
	// The internal name of the property to monitor for changes. Only applies when `eventType` is `propertyChange`.
	PropertyName *string `json:"propertyName,omitempty"`
	// Determines if the subscription is active or paused. Defaults to false.
	Active *bool `json:"active,omitempty"`
}

// NewSubscriptionCreateRequest instantiates a new SubscriptionCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionCreateRequest(eventType string) *SubscriptionCreateRequest {
	this := SubscriptionCreateRequest{}
	this.EventType = eventType
	return &this
}

// NewSubscriptionCreateRequestWithDefaults instantiates a new SubscriptionCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionCreateRequestWithDefaults() *SubscriptionCreateRequest {
	this := SubscriptionCreateRequest{}
	return &this
}

// GetEventType returns the EventType field value
func (o *SubscriptionCreateRequest) GetEventType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value
// and a boolean to check if the value has been set.
func (o *SubscriptionCreateRequest) GetEventTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventType, true
}

// SetEventType sets field value
func (o *SubscriptionCreateRequest) SetEventType(v string) {
	o.EventType = v
}

// GetPropertyName returns the PropertyName field value if set, zero value otherwise.
func (o *SubscriptionCreateRequest) GetPropertyName() string {
	if o == nil || IsNil(o.PropertyName) {
		var ret string
		return ret
	}
	return *o.PropertyName
}

// GetPropertyNameOk returns a tuple with the PropertyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionCreateRequest) GetPropertyNameOk() (*string, bool) {
	if o == nil || IsNil(o.PropertyName) {
		return nil, false
	}
	return o.PropertyName, true
}

// HasPropertyName returns a boolean if a field has been set.
func (o *SubscriptionCreateRequest) HasPropertyName() bool {
	if o != nil && !IsNil(o.PropertyName) {
		return true
	}

	return false
}

// SetPropertyName gets a reference to the given string and assigns it to the PropertyName field.
func (o *SubscriptionCreateRequest) SetPropertyName(v string) {
	o.PropertyName = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *SubscriptionCreateRequest) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionCreateRequest) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *SubscriptionCreateRequest) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *SubscriptionCreateRequest) SetActive(v bool) {
	o.Active = &v
}

func (o SubscriptionCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["eventType"] = o.EventType
	if !IsNil(o.PropertyName) {
		toSerialize["propertyName"] = o.PropertyName
	}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	return toSerialize, nil
}

type NullableSubscriptionCreateRequest struct {
	value *SubscriptionCreateRequest
	isSet bool
}

func (v NullableSubscriptionCreateRequest) Get() *SubscriptionCreateRequest {
	return v.value
}

func (v *NullableSubscriptionCreateRequest) Set(val *SubscriptionCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionCreateRequest(val *SubscriptionCreateRequest) *NullableSubscriptionCreateRequest {
	return &NullableSubscriptionCreateRequest{value: val, isSet: true}
}

func (v NullableSubscriptionCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


