/*
CMS Site Search

Use these endpoints for searching content on your HubSpot hosted CMS website(s).

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package site-search

import (
	"encoding/json"
)

// checks if the ContentSearchResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContentSearchResult{}

// ContentSearchResult An individual search result.
type ContentSearchResult struct {
	// The ID of the content.
	Id int32 `json:"id"`
	// The matching score of the document.
	Score float32 `json:"score"`
	// The type of document. Can be `SITE_PAGE`, `LANDING_PAGE`, `BLOG_POST`, `LISTING_PAGE`, or `KNOWLEDGE_ARTICLE`.
	Type string `json:"type"`
	// The domain the document is hosted on.
	Domain string `json:"domain"`
	// The url of the document.
	Url string `json:"url"`
	// URL of the featured image.
	FeaturedImageUrl *string `json:"featuredImageUrl,omitempty"`
	// The document's language.
	Language *string `json:"language,omitempty"`
	// The title of the returned document.
	Title *string `json:"title,omitempty"`
	// The result's description. The content will be determined by the value of `length` in the request.
	Description *string `json:"description,omitempty"`
	// For knowledge articles, the category of the article.
	Category *string `json:"category,omitempty"`
	// For knowledge articles, the subcategory of the article.
	Subcategory *string `json:"subcategory,omitempty"`
	// Name of the author.
	AuthorFullName *string `json:"authorFullName,omitempty"`
	// If a blog post, the tags associated with it.
	Tags []string `json:"tags,omitempty"`
	// If a dynamic page, the ID of the HubDB table.
	TableId *int64 `json:"tableId,omitempty"`
	// If a dynamic page, the row ID in the HubDB table.
	RowId *int64 `json:"rowId,omitempty"`
	// The date the content was published.
	PublishedDate *int64 `json:"publishedDate,omitempty"`
	// The ID of the document in HubSpot.
	CombinedId *string `json:"combinedId,omitempty"`
}

// NewContentSearchResult instantiates a new ContentSearchResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContentSearchResult(id int32, score float32, type_ string, domain string, url string) *ContentSearchResult {
	this := ContentSearchResult{}
	this.Id = id
	this.Score = score
	this.Type = type_
	this.Domain = domain
	this.Url = url
	return &this
}

// NewContentSearchResultWithDefaults instantiates a new ContentSearchResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContentSearchResultWithDefaults() *ContentSearchResult {
	this := ContentSearchResult{}
	return &this
}

// GetId returns the Id field value
func (o *ContentSearchResult) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ContentSearchResult) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ContentSearchResult) SetId(v int32) {
	o.Id = v
}

// GetScore returns the Score field value
func (o *ContentSearchResult) GetScore() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Score
}

// GetScoreOk returns a tuple with the Score field value
// and a boolean to check if the value has been set.
func (o *ContentSearchResult) GetScoreOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Score, true
}

// SetScore sets field value
func (o *ContentSearchResult) SetScore(v float32) {
	o.Score = v
}

// GetType returns the Type field value
func (o *ContentSearchResult) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ContentSearchResult) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ContentSearchResult) SetType(v string) {
	o.Type = v
}

// GetDomain returns the Domain field value
func (o *ContentSearchResult) GetDomain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Domain
}

// GetDomainOk returns a tuple with the Domain field value
// and a boolean to check if the value has been set.
func (o *ContentSearchResult) GetDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Domain, true
}

// SetDomain sets field value
func (o *ContentSearchResult) SetDomain(v string) {
	o.Domain = v
}

// GetUrl returns the Url field value
func (o *ContentSearchResult) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *ContentSearchResult) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *ContentSearchResult) SetUrl(v string) {
	o.Url = v
}

// GetFeaturedImageUrl returns the FeaturedImageUrl field value if set, zero value otherwise.
func (o *ContentSearchResult) GetFeaturedImageUrl() string {
	if o == nil || IsNil(o.FeaturedImageUrl) {
		var ret string
		return ret
	}
	return *o.FeaturedImageUrl
}

// GetFeaturedImageUrlOk returns a tuple with the FeaturedImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentSearchResult) GetFeaturedImageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.FeaturedImageUrl) {
		return nil, false
	}
	return o.FeaturedImageUrl, true
}

// HasFeaturedImageUrl returns a boolean if a field has been set.
func (o *ContentSearchResult) HasFeaturedImageUrl() bool {
	if o != nil && !IsNil(o.FeaturedImageUrl) {
		return true
	}

	return false
}

// SetFeaturedImageUrl gets a reference to the given string and assigns it to the FeaturedImageUrl field.
func (o *ContentSearchResult) SetFeaturedImageUrl(v string) {
	o.FeaturedImageUrl = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *ContentSearchResult) GetLanguage() string {
	if o == nil || IsNil(o.Language) {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentSearchResult) GetLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *ContentSearchResult) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *ContentSearchResult) SetLanguage(v string) {
	o.Language = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ContentSearchResult) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentSearchResult) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ContentSearchResult) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ContentSearchResult) SetTitle(v string) {
	o.Title = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ContentSearchResult) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentSearchResult) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ContentSearchResult) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ContentSearchResult) SetDescription(v string) {
	o.Description = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *ContentSearchResult) GetCategory() string {
	if o == nil || IsNil(o.Category) {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentSearchResult) GetCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *ContentSearchResult) HasCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *ContentSearchResult) SetCategory(v string) {
	o.Category = &v
}

// GetSubcategory returns the Subcategory field value if set, zero value otherwise.
func (o *ContentSearchResult) GetSubcategory() string {
	if o == nil || IsNil(o.Subcategory) {
		var ret string
		return ret
	}
	return *o.Subcategory
}

// GetSubcategoryOk returns a tuple with the Subcategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentSearchResult) GetSubcategoryOk() (*string, bool) {
	if o == nil || IsNil(o.Subcategory) {
		return nil, false
	}
	return o.Subcategory, true
}

// HasSubcategory returns a boolean if a field has been set.
func (o *ContentSearchResult) HasSubcategory() bool {
	if o != nil && !IsNil(o.Subcategory) {
		return true
	}

	return false
}

// SetSubcategory gets a reference to the given string and assigns it to the Subcategory field.
func (o *ContentSearchResult) SetSubcategory(v string) {
	o.Subcategory = &v
}

// GetAuthorFullName returns the AuthorFullName field value if set, zero value otherwise.
func (o *ContentSearchResult) GetAuthorFullName() string {
	if o == nil || IsNil(o.AuthorFullName) {
		var ret string
		return ret
	}
	return *o.AuthorFullName
}

// GetAuthorFullNameOk returns a tuple with the AuthorFullName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentSearchResult) GetAuthorFullNameOk() (*string, bool) {
	if o == nil || IsNil(o.AuthorFullName) {
		return nil, false
	}
	return o.AuthorFullName, true
}

// HasAuthorFullName returns a boolean if a field has been set.
func (o *ContentSearchResult) HasAuthorFullName() bool {
	if o != nil && !IsNil(o.AuthorFullName) {
		return true
	}

	return false
}

// SetAuthorFullName gets a reference to the given string and assigns it to the AuthorFullName field.
func (o *ContentSearchResult) SetAuthorFullName(v string) {
	o.AuthorFullName = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ContentSearchResult) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentSearchResult) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ContentSearchResult) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *ContentSearchResult) SetTags(v []string) {
	o.Tags = v
}

// GetTableId returns the TableId field value if set, zero value otherwise.
func (o *ContentSearchResult) GetTableId() int64 {
	if o == nil || IsNil(o.TableId) {
		var ret int64
		return ret
	}
	return *o.TableId
}

// GetTableIdOk returns a tuple with the TableId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentSearchResult) GetTableIdOk() (*int64, bool) {
	if o == nil || IsNil(o.TableId) {
		return nil, false
	}
	return o.TableId, true
}

// HasTableId returns a boolean if a field has been set.
func (o *ContentSearchResult) HasTableId() bool {
	if o != nil && !IsNil(o.TableId) {
		return true
	}

	return false
}

// SetTableId gets a reference to the given int64 and assigns it to the TableId field.
func (o *ContentSearchResult) SetTableId(v int64) {
	o.TableId = &v
}

// GetRowId returns the RowId field value if set, zero value otherwise.
func (o *ContentSearchResult) GetRowId() int64 {
	if o == nil || IsNil(o.RowId) {
		var ret int64
		return ret
	}
	return *o.RowId
}

// GetRowIdOk returns a tuple with the RowId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentSearchResult) GetRowIdOk() (*int64, bool) {
	if o == nil || IsNil(o.RowId) {
		return nil, false
	}
	return o.RowId, true
}

// HasRowId returns a boolean if a field has been set.
func (o *ContentSearchResult) HasRowId() bool {
	if o != nil && !IsNil(o.RowId) {
		return true
	}

	return false
}

// SetRowId gets a reference to the given int64 and assigns it to the RowId field.
func (o *ContentSearchResult) SetRowId(v int64) {
	o.RowId = &v
}

// GetPublishedDate returns the PublishedDate field value if set, zero value otherwise.
func (o *ContentSearchResult) GetPublishedDate() int64 {
	if o == nil || IsNil(o.PublishedDate) {
		var ret int64
		return ret
	}
	return *o.PublishedDate
}

// GetPublishedDateOk returns a tuple with the PublishedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentSearchResult) GetPublishedDateOk() (*int64, bool) {
	if o == nil || IsNil(o.PublishedDate) {
		return nil, false
	}
	return o.PublishedDate, true
}

// HasPublishedDate returns a boolean if a field has been set.
func (o *ContentSearchResult) HasPublishedDate() bool {
	if o != nil && !IsNil(o.PublishedDate) {
		return true
	}

	return false
}

// SetPublishedDate gets a reference to the given int64 and assigns it to the PublishedDate field.
func (o *ContentSearchResult) SetPublishedDate(v int64) {
	o.PublishedDate = &v
}

// GetCombinedId returns the CombinedId field value if set, zero value otherwise.
func (o *ContentSearchResult) GetCombinedId() string {
	if o == nil || IsNil(o.CombinedId) {
		var ret string
		return ret
	}
	return *o.CombinedId
}

// GetCombinedIdOk returns a tuple with the CombinedId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentSearchResult) GetCombinedIdOk() (*string, bool) {
	if o == nil || IsNil(o.CombinedId) {
		return nil, false
	}
	return o.CombinedId, true
}

// HasCombinedId returns a boolean if a field has been set.
func (o *ContentSearchResult) HasCombinedId() bool {
	if o != nil && !IsNil(o.CombinedId) {
		return true
	}

	return false
}

// SetCombinedId gets a reference to the given string and assigns it to the CombinedId field.
func (o *ContentSearchResult) SetCombinedId(v string) {
	o.CombinedId = &v
}

func (o ContentSearchResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContentSearchResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["score"] = o.Score
	toSerialize["type"] = o.Type
	toSerialize["domain"] = o.Domain
	toSerialize["url"] = o.Url
	if !IsNil(o.FeaturedImageUrl) {
		toSerialize["featuredImageUrl"] = o.FeaturedImageUrl
	}
	if !IsNil(o.Language) {
		toSerialize["language"] = o.Language
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	if !IsNil(o.Subcategory) {
		toSerialize["subcategory"] = o.Subcategory
	}
	if !IsNil(o.AuthorFullName) {
		toSerialize["authorFullName"] = o.AuthorFullName
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.TableId) {
		toSerialize["tableId"] = o.TableId
	}
	if !IsNil(o.RowId) {
		toSerialize["rowId"] = o.RowId
	}
	if !IsNil(o.PublishedDate) {
		toSerialize["publishedDate"] = o.PublishedDate
	}
	if !IsNil(o.CombinedId) {
		toSerialize["combinedId"] = o.CombinedId
	}
	return toSerialize, nil
}

type NullableContentSearchResult struct {
	value *ContentSearchResult
	isSet bool
}

func (v NullableContentSearchResult) Get() *ContentSearchResult {
	return v.value
}

func (v *NullableContentSearchResult) Set(val *ContentSearchResult) {
	v.value = val
	v.isSet = true
}

func (v NullableContentSearchResult) IsSet() bool {
	return v.isSet
}

func (v *NullableContentSearchResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentSearchResult(val *ContentSearchResult) *NullableContentSearchResult {
	return &NullableContentSearchResult{value: val, isSet: true}
}

func (v NullableContentSearchResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentSearchResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

