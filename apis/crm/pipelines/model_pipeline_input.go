/*
CRM Pipelines

Pipelines represent distinct stages in a workflow, like closing a deal or servicing a support ticket. These endpoints provide access to read and modify pipelines in HubSpot. Pipelines support `deals` and `tickets` object types.  ## Pipeline ID validation  When calling endpoints that take pipelineId as a parameter, that ID must correspond to an existing, un-archived pipeline. Otherwise the request will fail with a `404 Not Found` response.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pipelines

import (
	"encoding/json"
)

// checks if the PipelineInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PipelineInput{}

// PipelineInput An input used to create or replace a pipeline's definition.
type PipelineInput struct {
	// A unique label used to organize pipelines in HubSpot's UI
	Label string `json:"label"`
	// The order for displaying this pipeline. If two pipelines have a matching `displayOrder`, they will be sorted alphabetically by label.
	DisplayOrder int32 `json:"displayOrder"`
	// Pipeline stage inputs used to create the new or replacement pipeline.
	Stages []PipelineStageInput `json:"stages"`
}

// NewPipelineInput instantiates a new PipelineInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPipelineInput(label string, displayOrder int32, stages []PipelineStageInput) *PipelineInput {
	this := PipelineInput{}
	this.Label = label
	this.DisplayOrder = displayOrder
	this.Stages = stages
	return &this
}

// NewPipelineInputWithDefaults instantiates a new PipelineInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPipelineInputWithDefaults() *PipelineInput {
	this := PipelineInput{}
	return &this
}

// GetLabel returns the Label field value
func (o *PipelineInput) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *PipelineInput) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *PipelineInput) SetLabel(v string) {
	o.Label = v
}

// GetDisplayOrder returns the DisplayOrder field value
func (o *PipelineInput) GetDisplayOrder() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DisplayOrder
}

// GetDisplayOrderOk returns a tuple with the DisplayOrder field value
// and a boolean to check if the value has been set.
func (o *PipelineInput) GetDisplayOrderOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayOrder, true
}

// SetDisplayOrder sets field value
func (o *PipelineInput) SetDisplayOrder(v int32) {
	o.DisplayOrder = v
}

// GetStages returns the Stages field value
func (o *PipelineInput) GetStages() []PipelineStageInput {
	if o == nil {
		var ret []PipelineStageInput
		return ret
	}

	return o.Stages
}

// GetStagesOk returns a tuple with the Stages field value
// and a boolean to check if the value has been set.
func (o *PipelineInput) GetStagesOk() ([]PipelineStageInput, bool) {
	if o == nil {
		return nil, false
	}
	return o.Stages, true
}

// SetStages sets field value
func (o *PipelineInput) SetStages(v []PipelineStageInput) {
	o.Stages = v
}

func (o PipelineInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PipelineInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["label"] = o.Label
	toSerialize["displayOrder"] = o.DisplayOrder
	toSerialize["stages"] = o.Stages
	return toSerialize, nil
}

type NullablePipelineInput struct {
	value *PipelineInput
	isSet bool
}

func (v NullablePipelineInput) Get() *PipelineInput {
	return v.value
}

func (v *NullablePipelineInput) Set(val *PipelineInput) {
	v.value = val
	v.isSet = true
}

func (v NullablePipelineInput) IsSet() bool {
	return v.isSet
}

func (v *NullablePipelineInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePipelineInput(val *PipelineInput) *NullablePipelineInput {
	return &NullablePipelineInput{value: val, isSet: true}
}

func (v NullablePipelineInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePipelineInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
