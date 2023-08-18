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

// checks if the BatchInputSubscriptionBatchUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BatchInputSubscriptionBatchUpdateRequest{}

// BatchInputSubscriptionBatchUpdateRequest struct for BatchInputSubscriptionBatchUpdateRequest
type BatchInputSubscriptionBatchUpdateRequest struct {
	Inputs []SubscriptionBatchUpdateRequest `json:"inputs"`
}

// NewBatchInputSubscriptionBatchUpdateRequest instantiates a new BatchInputSubscriptionBatchUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchInputSubscriptionBatchUpdateRequest(inputs []SubscriptionBatchUpdateRequest) *BatchInputSubscriptionBatchUpdateRequest {
	this := BatchInputSubscriptionBatchUpdateRequest{}
	this.Inputs = inputs
	return &this
}

// NewBatchInputSubscriptionBatchUpdateRequestWithDefaults instantiates a new BatchInputSubscriptionBatchUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchInputSubscriptionBatchUpdateRequestWithDefaults() *BatchInputSubscriptionBatchUpdateRequest {
	this := BatchInputSubscriptionBatchUpdateRequest{}
	return &this
}

// GetInputs returns the Inputs field value
func (o *BatchInputSubscriptionBatchUpdateRequest) GetInputs() []SubscriptionBatchUpdateRequest {
	if o == nil {
		var ret []SubscriptionBatchUpdateRequest
		return ret
	}

	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value
// and a boolean to check if the value has been set.
func (o *BatchInputSubscriptionBatchUpdateRequest) GetInputsOk() ([]SubscriptionBatchUpdateRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Inputs, true
}

// SetInputs sets field value
func (o *BatchInputSubscriptionBatchUpdateRequest) SetInputs(v []SubscriptionBatchUpdateRequest) {
	o.Inputs = v
}

func (o BatchInputSubscriptionBatchUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BatchInputSubscriptionBatchUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["inputs"] = o.Inputs
	return toSerialize, nil
}

type NullableBatchInputSubscriptionBatchUpdateRequest struct {
	value *BatchInputSubscriptionBatchUpdateRequest
	isSet bool
}

func (v NullableBatchInputSubscriptionBatchUpdateRequest) Get() *BatchInputSubscriptionBatchUpdateRequest {
	return v.value
}

func (v *NullableBatchInputSubscriptionBatchUpdateRequest) Set(val *BatchInputSubscriptionBatchUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchInputSubscriptionBatchUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchInputSubscriptionBatchUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchInputSubscriptionBatchUpdateRequest(val *BatchInputSubscriptionBatchUpdateRequest) *NullableBatchInputSubscriptionBatchUpdateRequest {
	return &NullableBatchInputSubscriptionBatchUpdateRequest{value: val, isSet: true}
}

func (v NullableBatchInputSubscriptionBatchUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchInputSubscriptionBatchUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


