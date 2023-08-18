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

// checks if the ThrottlingSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ThrottlingSettings{}

// ThrottlingSettings Configuration details for webhook throttling.
type ThrottlingSettings struct {
	// The maximum number of HTTP requests HubSpot will attempt to make to your app in a given time frame determined by `period`.
	MaxConcurrentRequests int32 `json:"maxConcurrentRequests"`
	// Time scale for this setting. Can be either `SECONDLY` (per second) or `ROLLING_MINUTE` (per minute).
	Period string `json:"period"`
}

// NewThrottlingSettings instantiates a new ThrottlingSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThrottlingSettings(maxConcurrentRequests int32, period string) *ThrottlingSettings {
	this := ThrottlingSettings{}
	this.MaxConcurrentRequests = maxConcurrentRequests
	this.Period = period
	return &this
}

// NewThrottlingSettingsWithDefaults instantiates a new ThrottlingSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThrottlingSettingsWithDefaults() *ThrottlingSettings {
	this := ThrottlingSettings{}
	return &this
}

// GetMaxConcurrentRequests returns the MaxConcurrentRequests field value
func (o *ThrottlingSettings) GetMaxConcurrentRequests() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaxConcurrentRequests
}

// GetMaxConcurrentRequestsOk returns a tuple with the MaxConcurrentRequests field value
// and a boolean to check if the value has been set.
func (o *ThrottlingSettings) GetMaxConcurrentRequestsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxConcurrentRequests, true
}

// SetMaxConcurrentRequests sets field value
func (o *ThrottlingSettings) SetMaxConcurrentRequests(v int32) {
	o.MaxConcurrentRequests = v
}

// GetPeriod returns the Period field value
func (o *ThrottlingSettings) GetPeriod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Period
}

// GetPeriodOk returns a tuple with the Period field value
// and a boolean to check if the value has been set.
func (o *ThrottlingSettings) GetPeriodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Period, true
}

// SetPeriod sets field value
func (o *ThrottlingSettings) SetPeriod(v string) {
	o.Period = v
}

func (o ThrottlingSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ThrottlingSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["maxConcurrentRequests"] = o.MaxConcurrentRequests
	toSerialize["period"] = o.Period
	return toSerialize, nil
}

type NullableThrottlingSettings struct {
	value *ThrottlingSettings
	isSet bool
}

func (v NullableThrottlingSettings) Get() *ThrottlingSettings {
	return v.value
}

func (v *NullableThrottlingSettings) Set(val *ThrottlingSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableThrottlingSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableThrottlingSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThrottlingSettings(val *ThrottlingSettings) *NullableThrottlingSettings {
	return &NullableThrottlingSettings{value: val, isSet: true}
}

func (v NullableThrottlingSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThrottlingSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


