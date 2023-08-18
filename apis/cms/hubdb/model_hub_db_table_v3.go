/*
HubDB endpoints

HubDB is a relational data store that presents data as rows, columns, and cells in a table, much like a spreadsheet. HubDB tables can be added or modified [in the HubSpot CMS](https://knowledge.hubspot.com/cos-general/how-to-edit-hubdb-tables), but you can also use the API endpoints documented here. For more information on HubDB tables and using their data on a HubSpot site, see the [CMS developers site](https://designers.hubspot.com/docs/tools/hubdb). You can also see the [documentation for dynamic pages](https://designers.hubspot.com/docs/tutorials/how-to-build-dynamic-pages-with-hubdb) for more details about the `useForPages` field.  HubDB tables support `draft` and `published` versions. This allows you to update data in the table, either for testing or to allow for a manual approval process, without affecting any live pages using the existing data. Draft data can be reviewed, and published by a user working in HubSpot or published via the API. Draft data can also be discarded, allowing users to go back to the published version of the data without disrupting it. If a table is set to be `allowed for public access`, you can access the published version of the table and rows without any authentication by specifying the portal id via the query parameter `portalId`.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package hubdb

import (
	"encoding/json"
	"time"
)

// checks if the HubDbTableV3 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HubDbTableV3{}

// HubDbTableV3 Model for HubDB table
type HubDbTableV3 struct {
	// Id of the table
	Id *string `json:"id,omitempty"`
	// Name of the table
	Name string `json:"name"`
	// Label of the table
	Label string `json:"label"`
	// List of columns in the table
	Columns   []Column `json:"columns,omitempty"`
	Published *bool    `json:"published,omitempty"`
	Deleted   *bool    `json:"deleted,omitempty"`
	// Number of columns including deleted
	ColumnCount *int32 `json:"columnCount,omitempty"`
	// Number of rows in the table
	RowCount  *int32      `json:"rowCount,omitempty"`
	CreatedBy *SimpleUser `json:"createdBy,omitempty"`
	UpdatedBy *SimpleUser `json:"updatedBy,omitempty"`
	// Specifies whether the table can be used for creation of dynamic pages
	UseForPages *bool `json:"useForPages,omitempty"`
	// Specifies whether child tables can be created
	AllowChildTables *bool `json:"allowChildTables,omitempty"`
	// Specifies creation of multi-level dynamic pages using child tables
	EnableChildTablePages *bool `json:"enableChildTablePages,omitempty"`
	IsOrderedManually     *bool `json:"isOrderedManually,omitempty"`
	// Specifies the key value pairs of the metadata fields with the associated column ids
	DynamicMetaTags *map[string]int32 `json:"dynamicMetaTags,omitempty"`
	// Specifies whether the table can be read by public without authorization
	AllowPublicApiAccess *bool `json:"allowPublicApiAccess,omitempty"`
	// Timestamp at which the table is created
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// Timestamp at which the table is published recently
	PublishedAt *time.Time `json:"publishedAt,omitempty"`
	// Timestamp at which the table is updated recently
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// NewHubDbTableV3 instantiates a new HubDbTableV3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHubDbTableV3(name string, label string) *HubDbTableV3 {
	this := HubDbTableV3{}
	this.Name = name
	this.Label = label
	return &this
}

// NewHubDbTableV3WithDefaults instantiates a new HubDbTableV3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHubDbTableV3WithDefaults() *HubDbTableV3 {
	this := HubDbTableV3{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *HubDbTableV3) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HubDbTableV3) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *HubDbTableV3) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *HubDbTableV3) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *HubDbTableV3) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *HubDbTableV3) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *HubDbTableV3) SetName(v string) {
	o.Name = v
}

// GetLabel returns the Label field value
func (o *HubDbTableV3) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *HubDbTableV3) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *HubDbTableV3) SetLabel(v string) {
	o.Label = v
}

// GetColumns returns the Columns field value if set, zero value otherwise.
func (o *HubDbTableV3) GetColumns() []Column {
	if o == nil || IsNil(o.Columns) {
		var ret []Column
		return ret
	}
	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HubDbTableV3) GetColumnsOk() ([]Column, bool) {
	if o == nil || IsNil(o.Columns) {
		return nil, false
	}
	return o.Columns, true
}

// HasColumns returns a boolean if a field has been set.
func (o *HubDbTableV3) HasColumns() bool {
	if o != nil && !IsNil(o.Columns) {
		return true
	}

	return false
}

// SetColumns gets a reference to the given []Column and assigns it to the Columns field.
func (o *HubDbTableV3) SetColumns(v []Column) {
	o.Columns = v
}

// GetPublished returns the Published field value if set, zero value otherwise.
func (o *HubDbTableV3) GetPublished() bool {
	if o == nil || IsNil(o.Published) {
		var ret bool
		return ret
	}
	return *o.Published
}

// GetPublishedOk returns a tuple with the Published field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HubDbTableV3) GetPublishedOk() (*bool, bool) {
	if o == nil || IsNil(o.Published) {
		return nil, false
	}
	return o.Published, true
}

// HasPublished returns a boolean if a field has been set.
func (o *HubDbTableV3) HasPublished() bool {
	if o != nil && !IsNil(o.Published) {
		return true
	}

	return false
}

// SetPublished gets a reference to the given bool and assigns it to the Published field.
func (o *HubDbTableV3) SetPublished(v bool) {
	o.Published = &v
}

// GetDeleted returns the Deleted field value if set, zero value otherwise.
func (o *HubDbTableV3) GetDeleted() bool {
	if o == nil || IsNil(o.Deleted) {
		var ret bool
		return ret
	}
	return *o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HubDbTableV3) GetDeletedOk() (*bool, bool) {
	if o == nil || IsNil(o.Deleted) {
		return nil, false
	}
	return o.Deleted, true
}

// HasDeleted returns a boolean if a field has been set.
func (o *HubDbTableV3) HasDeleted() bool {
	if o != nil && !IsNil(o.Deleted) {
		return true
	}

	return false
}

// SetDeleted gets a reference to the given bool and assigns it to the Deleted field.
func (o *HubDbTableV3) SetDeleted(v bool) {
	o.Deleted = &v
}

// GetColumnCount returns the ColumnCount field value if set, zero value otherwise.
func (o *HubDbTableV3) GetColumnCount() int32 {
	if o == nil || IsNil(o.ColumnCount) {
		var ret int32
		return ret
	}
	return *o.ColumnCount
}

// GetColumnCountOk returns a tuple with the ColumnCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HubDbTableV3) GetColumnCountOk() (*int32, bool) {
	if o == nil || IsNil(o.ColumnCount) {
		return nil, false
	}
	return o.ColumnCount, true
}

// HasColumnCount returns a boolean if a field has been set.
func (o *HubDbTableV3) HasColumnCount() bool {
	if o != nil && !IsNil(o.ColumnCount) {
		return true
	}

	return false
}

// SetColumnCount gets a reference to the given int32 and assigns it to the ColumnCount field.
func (o *HubDbTableV3) SetColumnCount(v int32) {
	o.ColumnCount = &v
}

// GetRowCount returns the RowCount field value if set, zero value otherwise.
func (o *HubDbTableV3) GetRowCount() int32 {
	if o == nil || IsNil(o.RowCount) {
		var ret int32
		return ret
	}
	return *o.RowCount
}

// GetRowCountOk returns a tuple with the RowCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HubDbTableV3) GetRowCountOk() (*int32, bool) {
	if o == nil || IsNil(o.RowCount) {
		return nil, false
	}
	return o.RowCount, true
}

// HasRowCount returns a boolean if a field has been set.
func (o *HubDbTableV3) HasRowCount() bool {
	if o != nil && !IsNil(o.RowCount) {
		return true
	}

	return false
}

// SetRowCount gets a reference to the given int32 and assigns it to the RowCount field.
func (o *HubDbTableV3) SetRowCount(v int32) {
	o.RowCount = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *HubDbTableV3) GetCreatedBy() SimpleUser {
	if o == nil || IsNil(o.CreatedBy) {
		var ret SimpleUser
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HubDbTableV3) GetCreatedByOk() (*SimpleUser, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *HubDbTableV3) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given SimpleUser and assigns it to the CreatedBy field.
func (o *HubDbTableV3) SetCreatedBy(v SimpleUser) {
	o.CreatedBy = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *HubDbTableV3) GetUpdatedBy() SimpleUser {
	if o == nil || IsNil(o.UpdatedBy) {
		var ret SimpleUser
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HubDbTableV3) GetUpdatedByOk() (*SimpleUser, bool) {
	if o == nil || IsNil(o.UpdatedBy) {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *HubDbTableV3) HasUpdatedBy() bool {
	if o != nil && !IsNil(o.UpdatedBy) {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given SimpleUser and assigns it to the UpdatedBy field.
func (o *HubDbTableV3) SetUpdatedBy(v SimpleUser) {
	o.UpdatedBy = &v
}

// GetUseForPages returns the UseForPages field value if set, zero value otherwise.
func (o *HubDbTableV3) GetUseForPages() bool {
	if o == nil || IsNil(o.UseForPages) {
		var ret bool
		return ret
	}
	return *o.UseForPages
}

// GetUseForPagesOk returns a tuple with the UseForPages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HubDbTableV3) GetUseForPagesOk() (*bool, bool) {
	if o == nil || IsNil(o.UseForPages) {
		return nil, false
	}
	return o.UseForPages, true
}

// HasUseForPages returns a boolean if a field has been set.
func (o *HubDbTableV3) HasUseForPages() bool {
	if o != nil && !IsNil(o.UseForPages) {
		return true
	}

	return false
}

// SetUseForPages gets a reference to the given bool and assigns it to the UseForPages field.
func (o *HubDbTableV3) SetUseForPages(v bool) {
	o.UseForPages = &v
}

// GetAllowChildTables returns the AllowChildTables field value if set, zero value otherwise.
func (o *HubDbTableV3) GetAllowChildTables() bool {
	if o == nil || IsNil(o.AllowChildTables) {
		var ret bool
		return ret
	}
	return *o.AllowChildTables
}

// GetAllowChildTablesOk returns a tuple with the AllowChildTables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HubDbTableV3) GetAllowChildTablesOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowChildTables) {
		return nil, false
	}
	return o.AllowChildTables, true
}

// HasAllowChildTables returns a boolean if a field has been set.
func (o *HubDbTableV3) HasAllowChildTables() bool {
	if o != nil && !IsNil(o.AllowChildTables) {
		return true
	}

	return false
}

// SetAllowChildTables gets a reference to the given bool and assigns it to the AllowChildTables field.
func (o *HubDbTableV3) SetAllowChildTables(v bool) {
	o.AllowChildTables = &v
}

// GetEnableChildTablePages returns the EnableChildTablePages field value if set, zero value otherwise.
func (o *HubDbTableV3) GetEnableChildTablePages() bool {
	if o == nil || IsNil(o.EnableChildTablePages) {
		var ret bool
		return ret
	}
	return *o.EnableChildTablePages
}

// GetEnableChildTablePagesOk returns a tuple with the EnableChildTablePages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HubDbTableV3) GetEnableChildTablePagesOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableChildTablePages) {
		return nil, false
	}
	return o.EnableChildTablePages, true
}

// HasEnableChildTablePages returns a boolean if a field has been set.
func (o *HubDbTableV3) HasEnableChildTablePages() bool {
	if o != nil && !IsNil(o.EnableChildTablePages) {
		return true
	}

	return false
}

// SetEnableChildTablePages gets a reference to the given bool and assigns it to the EnableChildTablePages field.
func (o *HubDbTableV3) SetEnableChildTablePages(v bool) {
	o.EnableChildTablePages = &v
}

// GetIsOrderedManually returns the IsOrderedManually field value if set, zero value otherwise.
func (o *HubDbTableV3) GetIsOrderedManually() bool {
	if o == nil || IsNil(o.IsOrderedManually) {
		var ret bool
		return ret
	}
	return *o.IsOrderedManually
}

// GetIsOrderedManuallyOk returns a tuple with the IsOrderedManually field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HubDbTableV3) GetIsOrderedManuallyOk() (*bool, bool) {
	if o == nil || IsNil(o.IsOrderedManually) {
		return nil, false
	}
	return o.IsOrderedManually, true
}

// HasIsOrderedManually returns a boolean if a field has been set.
func (o *HubDbTableV3) HasIsOrderedManually() bool {
	if o != nil && !IsNil(o.IsOrderedManually) {
		return true
	}

	return false
}

// SetIsOrderedManually gets a reference to the given bool and assigns it to the IsOrderedManually field.
func (o *HubDbTableV3) SetIsOrderedManually(v bool) {
	o.IsOrderedManually = &v
}

// GetDynamicMetaTags returns the DynamicMetaTags field value if set, zero value otherwise.
func (o *HubDbTableV3) GetDynamicMetaTags() map[string]int32 {
	if o == nil || IsNil(o.DynamicMetaTags) {
		var ret map[string]int32
		return ret
	}
	return *o.DynamicMetaTags
}

// GetDynamicMetaTagsOk returns a tuple with the DynamicMetaTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HubDbTableV3) GetDynamicMetaTagsOk() (*map[string]int32, bool) {
	if o == nil || IsNil(o.DynamicMetaTags) {
		return nil, false
	}
	return o.DynamicMetaTags, true
}

// HasDynamicMetaTags returns a boolean if a field has been set.
func (o *HubDbTableV3) HasDynamicMetaTags() bool {
	if o != nil && !IsNil(o.DynamicMetaTags) {
		return true
	}

	return false
}

// SetDynamicMetaTags gets a reference to the given map[string]int32 and assigns it to the DynamicMetaTags field.
func (o *HubDbTableV3) SetDynamicMetaTags(v map[string]int32) {
	o.DynamicMetaTags = &v
}

// GetAllowPublicApiAccess returns the AllowPublicApiAccess field value if set, zero value otherwise.
func (o *HubDbTableV3) GetAllowPublicApiAccess() bool {
	if o == nil || IsNil(o.AllowPublicApiAccess) {
		var ret bool
		return ret
	}
	return *o.AllowPublicApiAccess
}

// GetAllowPublicApiAccessOk returns a tuple with the AllowPublicApiAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HubDbTableV3) GetAllowPublicApiAccessOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowPublicApiAccess) {
		return nil, false
	}
	return o.AllowPublicApiAccess, true
}

// HasAllowPublicApiAccess returns a boolean if a field has been set.
func (o *HubDbTableV3) HasAllowPublicApiAccess() bool {
	if o != nil && !IsNil(o.AllowPublicApiAccess) {
		return true
	}

	return false
}

// SetAllowPublicApiAccess gets a reference to the given bool and assigns it to the AllowPublicApiAccess field.
func (o *HubDbTableV3) SetAllowPublicApiAccess(v bool) {
	o.AllowPublicApiAccess = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *HubDbTableV3) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HubDbTableV3) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *HubDbTableV3) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *HubDbTableV3) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetPublishedAt returns the PublishedAt field value if set, zero value otherwise.
func (o *HubDbTableV3) GetPublishedAt() time.Time {
	if o == nil || IsNil(o.PublishedAt) {
		var ret time.Time
		return ret
	}
	return *o.PublishedAt
}

// GetPublishedAtOk returns a tuple with the PublishedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HubDbTableV3) GetPublishedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PublishedAt) {
		return nil, false
	}
	return o.PublishedAt, true
}

// HasPublishedAt returns a boolean if a field has been set.
func (o *HubDbTableV3) HasPublishedAt() bool {
	if o != nil && !IsNil(o.PublishedAt) {
		return true
	}

	return false
}

// SetPublishedAt gets a reference to the given time.Time and assigns it to the PublishedAt field.
func (o *HubDbTableV3) SetPublishedAt(v time.Time) {
	o.PublishedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *HubDbTableV3) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HubDbTableV3) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *HubDbTableV3) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *HubDbTableV3) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o HubDbTableV3) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HubDbTableV3) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["name"] = o.Name
	toSerialize["label"] = o.Label
	if !IsNil(o.Columns) {
		toSerialize["columns"] = o.Columns
	}
	if !IsNil(o.Published) {
		toSerialize["published"] = o.Published
	}
	if !IsNil(o.Deleted) {
		toSerialize["deleted"] = o.Deleted
	}
	if !IsNil(o.ColumnCount) {
		toSerialize["columnCount"] = o.ColumnCount
	}
	if !IsNil(o.RowCount) {
		toSerialize["rowCount"] = o.RowCount
	}
	if !IsNil(o.CreatedBy) {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if !IsNil(o.UpdatedBy) {
		toSerialize["updatedBy"] = o.UpdatedBy
	}
	if !IsNil(o.UseForPages) {
		toSerialize["useForPages"] = o.UseForPages
	}
	if !IsNil(o.AllowChildTables) {
		toSerialize["allowChildTables"] = o.AllowChildTables
	}
	if !IsNil(o.EnableChildTablePages) {
		toSerialize["enableChildTablePages"] = o.EnableChildTablePages
	}
	if !IsNil(o.IsOrderedManually) {
		toSerialize["isOrderedManually"] = o.IsOrderedManually
	}
	if !IsNil(o.DynamicMetaTags) {
		toSerialize["dynamicMetaTags"] = o.DynamicMetaTags
	}
	if !IsNil(o.AllowPublicApiAccess) {
		toSerialize["allowPublicApiAccess"] = o.AllowPublicApiAccess
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.PublishedAt) {
		toSerialize["publishedAt"] = o.PublishedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableHubDbTableV3 struct {
	value *HubDbTableV3
	isSet bool
}

func (v NullableHubDbTableV3) Get() *HubDbTableV3 {
	return v.value
}

func (v *NullableHubDbTableV3) Set(val *HubDbTableV3) {
	v.value = val
	v.isSet = true
}

func (v NullableHubDbTableV3) IsSet() bool {
	return v.isSet
}

func (v *NullableHubDbTableV3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHubDbTableV3(val *HubDbTableV3) *NullableHubDbTableV3 {
	return &NullableHubDbTableV3{value: val, isSet: true}
}

func (v NullableHubDbTableV3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHubDbTableV3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}