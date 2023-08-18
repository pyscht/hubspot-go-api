/*
Accounting Extension

These APIs allow you to interact with HubSpot's Accounting Extension. It allows you to: * Specify the URLs that HubSpot will use when making webhook requests to your external accounting system. * Respond to webhook calls made to your external accounting system by HubSpot

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package accounting

import (
	"encoding/json"
	"time"
)

// checks if the InvoiceUpdateResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvoiceUpdateResponse{}

// InvoiceUpdateResponse struct for InvoiceUpdateResponse
type InvoiceUpdateResponse struct {
	ExternalInvoiceNumber   *string    `json:"externalInvoiceNumber,omitempty"`
	TotalAmountBilled       float32    `json:"totalAmountBilled"`
	BalanceDue              float32    `json:"balanceDue"`
	CurrencyCode            string     `json:"currencyCode"`
	DueDate                 string     `json:"dueDate"`
	ExternalRecipientId     string     `json:"externalRecipientId"`
	ReceivedByRecipientDate *int64     `json:"receivedByRecipientDate,omitempty"`
	ExternalCreateDateTime  *int64     `json:"externalCreateDateTime,omitempty"`
	IsVoided                bool       `json:"isVoided"`
	CreatedAt               time.Time  `json:"createdAt"`
	UpdatedAt               time.Time  `json:"updatedAt"`
	ArchivedAt              *time.Time `json:"archivedAt,omitempty"`
	Archived                bool       `json:"archived"`
	ExternalAccountId       string     `json:"externalAccountId"`
	InvoiceStatus           string     `json:"invoiceStatus"`
	Id                      string     `json:"id"`
}

// NewInvoiceUpdateResponse instantiates a new InvoiceUpdateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoiceUpdateResponse(totalAmountBilled float32, balanceDue float32, currencyCode string, dueDate string, externalRecipientId string, isVoided bool, createdAt time.Time, updatedAt time.Time, archived bool, externalAccountId string, invoiceStatus string, id string) *InvoiceUpdateResponse {
	this := InvoiceUpdateResponse{}
	this.TotalAmountBilled = totalAmountBilled
	this.BalanceDue = balanceDue
	this.CurrencyCode = currencyCode
	this.DueDate = dueDate
	this.ExternalRecipientId = externalRecipientId
	this.IsVoided = isVoided
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.Archived = archived
	this.ExternalAccountId = externalAccountId
	this.InvoiceStatus = invoiceStatus
	this.Id = id
	return &this
}

// NewInvoiceUpdateResponseWithDefaults instantiates a new InvoiceUpdateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoiceUpdateResponseWithDefaults() *InvoiceUpdateResponse {
	this := InvoiceUpdateResponse{}
	return &this
}

// GetExternalInvoiceNumber returns the ExternalInvoiceNumber field value if set, zero value otherwise.
func (o *InvoiceUpdateResponse) GetExternalInvoiceNumber() string {
	if o == nil || IsNil(o.ExternalInvoiceNumber) {
		var ret string
		return ret
	}
	return *o.ExternalInvoiceNumber
}

// GetExternalInvoiceNumberOk returns a tuple with the ExternalInvoiceNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceUpdateResponse) GetExternalInvoiceNumberOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalInvoiceNumber) {
		return nil, false
	}
	return o.ExternalInvoiceNumber, true
}

// HasExternalInvoiceNumber returns a boolean if a field has been set.
func (o *InvoiceUpdateResponse) HasExternalInvoiceNumber() bool {
	if o != nil && !IsNil(o.ExternalInvoiceNumber) {
		return true
	}

	return false
}

// SetExternalInvoiceNumber gets a reference to the given string and assigns it to the ExternalInvoiceNumber field.
func (o *InvoiceUpdateResponse) SetExternalInvoiceNumber(v string) {
	o.ExternalInvoiceNumber = &v
}

// GetTotalAmountBilled returns the TotalAmountBilled field value
func (o *InvoiceUpdateResponse) GetTotalAmountBilled() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TotalAmountBilled
}

// GetTotalAmountBilledOk returns a tuple with the TotalAmountBilled field value
// and a boolean to check if the value has been set.
func (o *InvoiceUpdateResponse) GetTotalAmountBilledOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalAmountBilled, true
}

// SetTotalAmountBilled sets field value
func (o *InvoiceUpdateResponse) SetTotalAmountBilled(v float32) {
	o.TotalAmountBilled = v
}

// GetBalanceDue returns the BalanceDue field value
func (o *InvoiceUpdateResponse) GetBalanceDue() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.BalanceDue
}

// GetBalanceDueOk returns a tuple with the BalanceDue field value
// and a boolean to check if the value has been set.
func (o *InvoiceUpdateResponse) GetBalanceDueOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BalanceDue, true
}

// SetBalanceDue sets field value
func (o *InvoiceUpdateResponse) SetBalanceDue(v float32) {
	o.BalanceDue = v
}

// GetCurrencyCode returns the CurrencyCode field value
func (o *InvoiceUpdateResponse) GetCurrencyCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value
// and a boolean to check if the value has been set.
func (o *InvoiceUpdateResponse) GetCurrencyCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrencyCode, true
}

// SetCurrencyCode sets field value
func (o *InvoiceUpdateResponse) SetCurrencyCode(v string) {
	o.CurrencyCode = v
}

// GetDueDate returns the DueDate field value
func (o *InvoiceUpdateResponse) GetDueDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DueDate
}

// GetDueDateOk returns a tuple with the DueDate field value
// and a boolean to check if the value has been set.
func (o *InvoiceUpdateResponse) GetDueDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DueDate, true
}

// SetDueDate sets field value
func (o *InvoiceUpdateResponse) SetDueDate(v string) {
	o.DueDate = v
}

// GetExternalRecipientId returns the ExternalRecipientId field value
func (o *InvoiceUpdateResponse) GetExternalRecipientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalRecipientId
}

// GetExternalRecipientIdOk returns a tuple with the ExternalRecipientId field value
// and a boolean to check if the value has been set.
func (o *InvoiceUpdateResponse) GetExternalRecipientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalRecipientId, true
}

// SetExternalRecipientId sets field value
func (o *InvoiceUpdateResponse) SetExternalRecipientId(v string) {
	o.ExternalRecipientId = v
}

// GetReceivedByRecipientDate returns the ReceivedByRecipientDate field value if set, zero value otherwise.
func (o *InvoiceUpdateResponse) GetReceivedByRecipientDate() int64 {
	if o == nil || IsNil(o.ReceivedByRecipientDate) {
		var ret int64
		return ret
	}
	return *o.ReceivedByRecipientDate
}

// GetReceivedByRecipientDateOk returns a tuple with the ReceivedByRecipientDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceUpdateResponse) GetReceivedByRecipientDateOk() (*int64, bool) {
	if o == nil || IsNil(o.ReceivedByRecipientDate) {
		return nil, false
	}
	return o.ReceivedByRecipientDate, true
}

// HasReceivedByRecipientDate returns a boolean if a field has been set.
func (o *InvoiceUpdateResponse) HasReceivedByRecipientDate() bool {
	if o != nil && !IsNil(o.ReceivedByRecipientDate) {
		return true
	}

	return false
}

// SetReceivedByRecipientDate gets a reference to the given int64 and assigns it to the ReceivedByRecipientDate field.
func (o *InvoiceUpdateResponse) SetReceivedByRecipientDate(v int64) {
	o.ReceivedByRecipientDate = &v
}

// GetExternalCreateDateTime returns the ExternalCreateDateTime field value if set, zero value otherwise.
func (o *InvoiceUpdateResponse) GetExternalCreateDateTime() int64 {
	if o == nil || IsNil(o.ExternalCreateDateTime) {
		var ret int64
		return ret
	}
	return *o.ExternalCreateDateTime
}

// GetExternalCreateDateTimeOk returns a tuple with the ExternalCreateDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceUpdateResponse) GetExternalCreateDateTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.ExternalCreateDateTime) {
		return nil, false
	}
	return o.ExternalCreateDateTime, true
}

// HasExternalCreateDateTime returns a boolean if a field has been set.
func (o *InvoiceUpdateResponse) HasExternalCreateDateTime() bool {
	if o != nil && !IsNil(o.ExternalCreateDateTime) {
		return true
	}

	return false
}

// SetExternalCreateDateTime gets a reference to the given int64 and assigns it to the ExternalCreateDateTime field.
func (o *InvoiceUpdateResponse) SetExternalCreateDateTime(v int64) {
	o.ExternalCreateDateTime = &v
}

// GetIsVoided returns the IsVoided field value
func (o *InvoiceUpdateResponse) GetIsVoided() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsVoided
}

// GetIsVoidedOk returns a tuple with the IsVoided field value
// and a boolean to check if the value has been set.
func (o *InvoiceUpdateResponse) GetIsVoidedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsVoided, true
}

// SetIsVoided sets field value
func (o *InvoiceUpdateResponse) SetIsVoided(v bool) {
	o.IsVoided = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *InvoiceUpdateResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *InvoiceUpdateResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *InvoiceUpdateResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *InvoiceUpdateResponse) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *InvoiceUpdateResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *InvoiceUpdateResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetArchivedAt returns the ArchivedAt field value if set, zero value otherwise.
func (o *InvoiceUpdateResponse) GetArchivedAt() time.Time {
	if o == nil || IsNil(o.ArchivedAt) {
		var ret time.Time
		return ret
	}
	return *o.ArchivedAt
}

// GetArchivedAtOk returns a tuple with the ArchivedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceUpdateResponse) GetArchivedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ArchivedAt) {
		return nil, false
	}
	return o.ArchivedAt, true
}

// HasArchivedAt returns a boolean if a field has been set.
func (o *InvoiceUpdateResponse) HasArchivedAt() bool {
	if o != nil && !IsNil(o.ArchivedAt) {
		return true
	}

	return false
}

// SetArchivedAt gets a reference to the given time.Time and assigns it to the ArchivedAt field.
func (o *InvoiceUpdateResponse) SetArchivedAt(v time.Time) {
	o.ArchivedAt = &v
}

// GetArchived returns the Archived field value
func (o *InvoiceUpdateResponse) GetArchived() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Archived
}

// GetArchivedOk returns a tuple with the Archived field value
// and a boolean to check if the value has been set.
func (o *InvoiceUpdateResponse) GetArchivedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Archived, true
}

// SetArchived sets field value
func (o *InvoiceUpdateResponse) SetArchived(v bool) {
	o.Archived = v
}

// GetExternalAccountId returns the ExternalAccountId field value
func (o *InvoiceUpdateResponse) GetExternalAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalAccountId
}

// GetExternalAccountIdOk returns a tuple with the ExternalAccountId field value
// and a boolean to check if the value has been set.
func (o *InvoiceUpdateResponse) GetExternalAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalAccountId, true
}

// SetExternalAccountId sets field value
func (o *InvoiceUpdateResponse) SetExternalAccountId(v string) {
	o.ExternalAccountId = v
}

// GetInvoiceStatus returns the InvoiceStatus field value
func (o *InvoiceUpdateResponse) GetInvoiceStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InvoiceStatus
}

// GetInvoiceStatusOk returns a tuple with the InvoiceStatus field value
// and a boolean to check if the value has been set.
func (o *InvoiceUpdateResponse) GetInvoiceStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InvoiceStatus, true
}

// SetInvoiceStatus sets field value
func (o *InvoiceUpdateResponse) SetInvoiceStatus(v string) {
	o.InvoiceStatus = v
}

// GetId returns the Id field value
func (o *InvoiceUpdateResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *InvoiceUpdateResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *InvoiceUpdateResponse) SetId(v string) {
	o.Id = v
}

func (o InvoiceUpdateResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvoiceUpdateResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExternalInvoiceNumber) {
		toSerialize["externalInvoiceNumber"] = o.ExternalInvoiceNumber
	}
	toSerialize["totalAmountBilled"] = o.TotalAmountBilled
	toSerialize["balanceDue"] = o.BalanceDue
	toSerialize["currencyCode"] = o.CurrencyCode
	toSerialize["dueDate"] = o.DueDate
	toSerialize["externalRecipientId"] = o.ExternalRecipientId
	if !IsNil(o.ReceivedByRecipientDate) {
		toSerialize["receivedByRecipientDate"] = o.ReceivedByRecipientDate
	}
	if !IsNil(o.ExternalCreateDateTime) {
		toSerialize["externalCreateDateTime"] = o.ExternalCreateDateTime
	}
	toSerialize["isVoided"] = o.IsVoided
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["updatedAt"] = o.UpdatedAt
	if !IsNil(o.ArchivedAt) {
		toSerialize["archivedAt"] = o.ArchivedAt
	}
	toSerialize["archived"] = o.Archived
	toSerialize["externalAccountId"] = o.ExternalAccountId
	toSerialize["invoiceStatus"] = o.InvoiceStatus
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableInvoiceUpdateResponse struct {
	value *InvoiceUpdateResponse
	isSet bool
}

func (v NullableInvoiceUpdateResponse) Get() *InvoiceUpdateResponse {
	return v.value
}

func (v *NullableInvoiceUpdateResponse) Set(val *InvoiceUpdateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoiceUpdateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoiceUpdateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoiceUpdateResponse(val *InvoiceUpdateResponse) *NullableInvoiceUpdateResponse {
	return &NullableInvoiceUpdateResponse{value: val, isSet: true}
}

func (v NullableInvoiceUpdateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoiceUpdateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}