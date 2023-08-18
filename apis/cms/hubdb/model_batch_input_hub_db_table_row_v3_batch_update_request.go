/*
HubDB endpoints

HubDB is a relational data store that presents data as rows, columns, and cells in a table, much like a spreadsheet. HubDB tables can be added or modified [in the HubSpot CMS](https://knowledge.hubspot.com/cos-general/how-to-edit-hubdb-tables), but you can also use the API endpoints documented here. For more information on HubDB tables and using their data on a HubSpot site, see the [CMS developers site](https://designers.hubspot.com/docs/tools/hubdb). You can also see the [documentation for dynamic pages](https://designers.hubspot.com/docs/tutorials/how-to-build-dynamic-pages-with-hubdb) for more details about the `useForPages` field.  HubDB tables support `draft` and `published` versions. This allows you to update data in the table, either for testing or to allow for a manual approval process, without affecting any live pages using the existing data. Draft data can be reviewed, and published by a user working in HubSpot or published via the API. Draft data can also be discarded, allowing users to go back to the published version of the data without disrupting it. If a table is set to be `allowed for public access`, you can access the published version of the table and rows without any authentication by specifying the portal id via the query parameter `portalId`.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package hubdb

import (
	"encoding/json"
)

// checks if the BatchInputHubDbTableRowV3BatchUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BatchInputHubDbTableRowV3BatchUpdateRequest{}

// BatchInputHubDbTableRowV3BatchUpdateRequest struct for BatchInputHubDbTableRowV3BatchUpdateRequest
type BatchInputHubDbTableRowV3BatchUpdateRequest struct {
	Inputs []HubDbTableRowV3BatchUpdateRequest `json:"inputs"`
}

// NewBatchInputHubDbTableRowV3BatchUpdateRequest instantiates a new BatchInputHubDbTableRowV3BatchUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchInputHubDbTableRowV3BatchUpdateRequest(inputs []HubDbTableRowV3BatchUpdateRequest) *BatchInputHubDbTableRowV3BatchUpdateRequest {
	this := BatchInputHubDbTableRowV3BatchUpdateRequest{}
	this.Inputs = inputs
	return &this
}

// NewBatchInputHubDbTableRowV3BatchUpdateRequestWithDefaults instantiates a new BatchInputHubDbTableRowV3BatchUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchInputHubDbTableRowV3BatchUpdateRequestWithDefaults() *BatchInputHubDbTableRowV3BatchUpdateRequest {
	this := BatchInputHubDbTableRowV3BatchUpdateRequest{}
	return &this
}

// GetInputs returns the Inputs field value
func (o *BatchInputHubDbTableRowV3BatchUpdateRequest) GetInputs() []HubDbTableRowV3BatchUpdateRequest {
	if o == nil {
		var ret []HubDbTableRowV3BatchUpdateRequest
		return ret
	}

	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value
// and a boolean to check if the value has been set.
func (o *BatchInputHubDbTableRowV3BatchUpdateRequest) GetInputsOk() ([]HubDbTableRowV3BatchUpdateRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Inputs, true
}

// SetInputs sets field value
func (o *BatchInputHubDbTableRowV3BatchUpdateRequest) SetInputs(v []HubDbTableRowV3BatchUpdateRequest) {
	o.Inputs = v
}

func (o BatchInputHubDbTableRowV3BatchUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BatchInputHubDbTableRowV3BatchUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["inputs"] = o.Inputs
	return toSerialize, nil
}

type NullableBatchInputHubDbTableRowV3BatchUpdateRequest struct {
	value *BatchInputHubDbTableRowV3BatchUpdateRequest
	isSet bool
}

func (v NullableBatchInputHubDbTableRowV3BatchUpdateRequest) Get() *BatchInputHubDbTableRowV3BatchUpdateRequest {
	return v.value
}

func (v *NullableBatchInputHubDbTableRowV3BatchUpdateRequest) Set(val *BatchInputHubDbTableRowV3BatchUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchInputHubDbTableRowV3BatchUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchInputHubDbTableRowV3BatchUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchInputHubDbTableRowV3BatchUpdateRequest(val *BatchInputHubDbTableRowV3BatchUpdateRequest) *NullableBatchInputHubDbTableRowV3BatchUpdateRequest {
	return &NullableBatchInputHubDbTableRowV3BatchUpdateRequest{value: val, isSet: true}
}

func (v NullableBatchInputHubDbTableRowV3BatchUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchInputHubDbTableRowV3BatchUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
