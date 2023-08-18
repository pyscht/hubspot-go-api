/*
Companies

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package companies

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
)

// BasicApiService BasicApi service
type BasicApiService service

type ApiDeleteCrmV3ObjectsCompaniesCompanyIdArchiveRequest struct {
	ctx        context.Context
	ApiService *BasicApiService
	companyId  string
}

func (r ApiDeleteCrmV3ObjectsCompaniesCompanyIdArchiveRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteCrmV3ObjectsCompaniesCompanyIdArchiveExecute(r)
}

/*
DeleteCrmV3ObjectsCompaniesCompanyIdArchive Archive

Move an Object identified by `{companyId}` to the recycling bin.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param companyId
	@return ApiDeleteCrmV3ObjectsCompaniesCompanyIdArchiveRequest
*/
func (a *BasicApiService) DeleteCrmV3ObjectsCompaniesCompanyIdArchive(ctx context.Context, companyId string) ApiDeleteCrmV3ObjectsCompaniesCompanyIdArchiveRequest {
	return ApiDeleteCrmV3ObjectsCompaniesCompanyIdArchiveRequest{
		ApiService: a,
		ctx:        ctx,
		companyId:  companyId,
	}
}

// Execute executes the request
func (a *BasicApiService) DeleteCrmV3ObjectsCompaniesCompanyIdArchiveExecute(r ApiDeleteCrmV3ObjectsCompaniesCompanyIdArchiveRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BasicApiService.DeleteCrmV3ObjectsCompaniesCompanyIdArchive")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/objects/companies/{companyId}"
	localVarPath = strings.Replace(localVarPath, "{"+"companyId"+"}", url.PathEscape(parameterValueToString(r.companyId, "companyId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["private_apps_legacy"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["private-app-legacy"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["private_apps"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["private-app"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetCrmV3ObjectsCompaniesCompanyIdGetByIdRequest struct {
	ctx                   context.Context
	ApiService            *BasicApiService
	companyId             string
	properties            *[]string
	propertiesWithHistory *[]string
	associations          *[]string
	archived              *bool
	idProperty            *string
}

// A comma separated list of the properties to be returned in the response. If any of the specified properties are not present on the requested object(s), they will be ignored.
func (r ApiGetCrmV3ObjectsCompaniesCompanyIdGetByIdRequest) Properties(properties []string) ApiGetCrmV3ObjectsCompaniesCompanyIdGetByIdRequest {
	r.properties = &properties
	return r
}

// A comma separated list of the properties to be returned along with their history of previous values. If any of the specified properties are not present on the requested object(s), they will be ignored.
func (r ApiGetCrmV3ObjectsCompaniesCompanyIdGetByIdRequest) PropertiesWithHistory(propertiesWithHistory []string) ApiGetCrmV3ObjectsCompaniesCompanyIdGetByIdRequest {
	r.propertiesWithHistory = &propertiesWithHistory
	return r
}

// A comma separated list of object types to retrieve associated IDs for. If any of the specified associations do not exist, they will be ignored.
func (r ApiGetCrmV3ObjectsCompaniesCompanyIdGetByIdRequest) Associations(associations []string) ApiGetCrmV3ObjectsCompaniesCompanyIdGetByIdRequest {
	r.associations = &associations
	return r
}

// Whether to return only results that have been archived.
func (r ApiGetCrmV3ObjectsCompaniesCompanyIdGetByIdRequest) Archived(archived bool) ApiGetCrmV3ObjectsCompaniesCompanyIdGetByIdRequest {
	r.archived = &archived
	return r
}

// The name of a property whose values are unique for this object type
func (r ApiGetCrmV3ObjectsCompaniesCompanyIdGetByIdRequest) IdProperty(idProperty string) ApiGetCrmV3ObjectsCompaniesCompanyIdGetByIdRequest {
	r.idProperty = &idProperty
	return r
}

func (r ApiGetCrmV3ObjectsCompaniesCompanyIdGetByIdRequest) Execute() (*SimplePublicObjectWithAssociations, *http.Response, error) {
	return r.ApiService.GetCrmV3ObjectsCompaniesCompanyIdGetByIdExecute(r)
}

/*
GetCrmV3ObjectsCompaniesCompanyIdGetById Read

Read an Object identified by `{companyId}`. `{companyId}` refers to the internal object ID by default, or optionally any unique property value as specified by the `idProperty` query param.  Control what is returned via the `properties` query param.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param companyId
	@return ApiGetCrmV3ObjectsCompaniesCompanyIdGetByIdRequest
*/
func (a *BasicApiService) GetCrmV3ObjectsCompaniesCompanyIdGetById(ctx context.Context, companyId string) ApiGetCrmV3ObjectsCompaniesCompanyIdGetByIdRequest {
	return ApiGetCrmV3ObjectsCompaniesCompanyIdGetByIdRequest{
		ApiService: a,
		ctx:        ctx,
		companyId:  companyId,
	}
}

// Execute executes the request
//
//	@return SimplePublicObjectWithAssociations
func (a *BasicApiService) GetCrmV3ObjectsCompaniesCompanyIdGetByIdExecute(r ApiGetCrmV3ObjectsCompaniesCompanyIdGetByIdRequest) (*SimplePublicObjectWithAssociations, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SimplePublicObjectWithAssociations
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BasicApiService.GetCrmV3ObjectsCompaniesCompanyIdGetById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/objects/companies/{companyId}"
	localVarPath = strings.Replace(localVarPath, "{"+"companyId"+"}", url.PathEscape(parameterValueToString(r.companyId, "companyId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.properties != nil {
		t := *r.properties
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "properties", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "properties", t, "multi")
		}
	}
	if r.propertiesWithHistory != nil {
		t := *r.propertiesWithHistory
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "propertiesWithHistory", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "propertiesWithHistory", t, "multi")
		}
	}
	if r.associations != nil {
		t := *r.associations
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "associations", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "associations", t, "multi")
		}
	}
	if r.archived != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "archived", r.archived, "")
	}
	if r.idProperty != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "idProperty", r.idProperty, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["private_apps_legacy"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["private-app-legacy"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["private_apps"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["private-app"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetCrmV3ObjectsCompaniesGetPageRequest struct {
	ctx                   context.Context
	ApiService            *BasicApiService
	limit                 *int32
	after                 *string
	properties            *[]string
	propertiesWithHistory *[]string
	associations          *[]string
	archived              *bool
}

// The maximum number of results to display per page.
func (r ApiGetCrmV3ObjectsCompaniesGetPageRequest) Limit(limit int32) ApiGetCrmV3ObjectsCompaniesGetPageRequest {
	r.limit = &limit
	return r
}

// The paging cursor token of the last successfully read resource will be returned as the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results.
func (r ApiGetCrmV3ObjectsCompaniesGetPageRequest) After(after string) ApiGetCrmV3ObjectsCompaniesGetPageRequest {
	r.after = &after
	return r
}

// A comma separated list of the properties to be returned in the response. If any of the specified properties are not present on the requested object(s), they will be ignored.
func (r ApiGetCrmV3ObjectsCompaniesGetPageRequest) Properties(properties []string) ApiGetCrmV3ObjectsCompaniesGetPageRequest {
	r.properties = &properties
	return r
}

// A comma separated list of the properties to be returned along with their history of previous values. If any of the specified properties are not present on the requested object(s), they will be ignored. Usage of this parameter will reduce the maximum number of objects that can be read by a single request.
func (r ApiGetCrmV3ObjectsCompaniesGetPageRequest) PropertiesWithHistory(propertiesWithHistory []string) ApiGetCrmV3ObjectsCompaniesGetPageRequest {
	r.propertiesWithHistory = &propertiesWithHistory
	return r
}

// A comma separated list of object types to retrieve associated IDs for. If any of the specified associations do not exist, they will be ignored.
func (r ApiGetCrmV3ObjectsCompaniesGetPageRequest) Associations(associations []string) ApiGetCrmV3ObjectsCompaniesGetPageRequest {
	r.associations = &associations
	return r
}

// Whether to return only results that have been archived.
func (r ApiGetCrmV3ObjectsCompaniesGetPageRequest) Archived(archived bool) ApiGetCrmV3ObjectsCompaniesGetPageRequest {
	r.archived = &archived
	return r
}

func (r ApiGetCrmV3ObjectsCompaniesGetPageRequest) Execute() (*CollectionResponseSimplePublicObjectWithAssociationsForwardPaging, *http.Response, error) {
	return r.ApiService.GetCrmV3ObjectsCompaniesGetPageExecute(r)
}

/*
GetCrmV3ObjectsCompaniesGetPage List

Read a page of companies. Control what is returned via the `properties` query param.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetCrmV3ObjectsCompaniesGetPageRequest
*/
func (a *BasicApiService) GetCrmV3ObjectsCompaniesGetPage(ctx context.Context) ApiGetCrmV3ObjectsCompaniesGetPageRequest {
	return ApiGetCrmV3ObjectsCompaniesGetPageRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CollectionResponseSimplePublicObjectWithAssociationsForwardPaging
func (a *BasicApiService) GetCrmV3ObjectsCompaniesGetPageExecute(r ApiGetCrmV3ObjectsCompaniesGetPageRequest) (*CollectionResponseSimplePublicObjectWithAssociationsForwardPaging, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CollectionResponseSimplePublicObjectWithAssociationsForwardPaging
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BasicApiService.GetCrmV3ObjectsCompaniesGetPage")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/objects/companies"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.after != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "after", r.after, "")
	}
	if r.properties != nil {
		t := *r.properties
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "properties", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "properties", t, "multi")
		}
	}
	if r.propertiesWithHistory != nil {
		t := *r.propertiesWithHistory
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "propertiesWithHistory", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "propertiesWithHistory", t, "multi")
		}
	}
	if r.associations != nil {
		t := *r.associations
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "associations", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "associations", t, "multi")
		}
	}
	if r.archived != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "archived", r.archived, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["private_apps_legacy"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["private-app-legacy"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["private_apps"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["private-app"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPatchCrmV3ObjectsCompaniesCompanyIdUpdateRequest struct {
	ctx                     context.Context
	ApiService              *BasicApiService
	companyId               string
	simplePublicObjectInput *SimplePublicObjectInput
	idProperty              *string
}

func (r ApiPatchCrmV3ObjectsCompaniesCompanyIdUpdateRequest) SimplePublicObjectInput(simplePublicObjectInput SimplePublicObjectInput) ApiPatchCrmV3ObjectsCompaniesCompanyIdUpdateRequest {
	r.simplePublicObjectInput = &simplePublicObjectInput
	return r
}

// The name of a property whose values are unique for this object type
func (r ApiPatchCrmV3ObjectsCompaniesCompanyIdUpdateRequest) IdProperty(idProperty string) ApiPatchCrmV3ObjectsCompaniesCompanyIdUpdateRequest {
	r.idProperty = &idProperty
	return r
}

func (r ApiPatchCrmV3ObjectsCompaniesCompanyIdUpdateRequest) Execute() (*SimplePublicObject, *http.Response, error) {
	return r.ApiService.PatchCrmV3ObjectsCompaniesCompanyIdUpdateExecute(r)
}

/*
PatchCrmV3ObjectsCompaniesCompanyIdUpdate Update

Perform a partial update of an Object identified by `{companyId}`. `{companyId}` refers to the internal object ID by default, or optionally any unique property value as specified by the `idProperty` query param. Provided property values will be overwritten. Read-only and non-existent properties will be ignored. Properties values can be cleared by passing an empty string.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param companyId
	@return ApiPatchCrmV3ObjectsCompaniesCompanyIdUpdateRequest
*/
func (a *BasicApiService) PatchCrmV3ObjectsCompaniesCompanyIdUpdate(ctx context.Context, companyId string) ApiPatchCrmV3ObjectsCompaniesCompanyIdUpdateRequest {
	return ApiPatchCrmV3ObjectsCompaniesCompanyIdUpdateRequest{
		ApiService: a,
		ctx:        ctx,
		companyId:  companyId,
	}
}

// Execute executes the request
//
//	@return SimplePublicObject
func (a *BasicApiService) PatchCrmV3ObjectsCompaniesCompanyIdUpdateExecute(r ApiPatchCrmV3ObjectsCompaniesCompanyIdUpdateRequest) (*SimplePublicObject, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SimplePublicObject
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BasicApiService.PatchCrmV3ObjectsCompaniesCompanyIdUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/objects/companies/{companyId}"
	localVarPath = strings.Replace(localVarPath, "{"+"companyId"+"}", url.PathEscape(parameterValueToString(r.companyId, "companyId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.simplePublicObjectInput == nil {
		return localVarReturnValue, nil, reportError("simplePublicObjectInput is required and must be specified")
	}

	if r.idProperty != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "idProperty", r.idProperty, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.simplePublicObjectInput
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["private_apps_legacy"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["private-app-legacy"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["private_apps"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["private-app"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPostCrmV3ObjectsCompaniesCreateRequest struct {
	ctx                              context.Context
	ApiService                       *BasicApiService
	simplePublicObjectInputForCreate *SimplePublicObjectInputForCreate
}

func (r ApiPostCrmV3ObjectsCompaniesCreateRequest) SimplePublicObjectInputForCreate(simplePublicObjectInputForCreate SimplePublicObjectInputForCreate) ApiPostCrmV3ObjectsCompaniesCreateRequest {
	r.simplePublicObjectInputForCreate = &simplePublicObjectInputForCreate
	return r
}

func (r ApiPostCrmV3ObjectsCompaniesCreateRequest) Execute() (*SimplePublicObject, *http.Response, error) {
	return r.ApiService.PostCrmV3ObjectsCompaniesCreateExecute(r)
}

/*
PostCrmV3ObjectsCompaniesCreate Create

Create a company with the given properties and return a copy of the object, including the ID. Documentation and examples for creating standard companies is provided.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiPostCrmV3ObjectsCompaniesCreateRequest
*/
func (a *BasicApiService) PostCrmV3ObjectsCompaniesCreate(ctx context.Context) ApiPostCrmV3ObjectsCompaniesCreateRequest {
	return ApiPostCrmV3ObjectsCompaniesCreateRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SimplePublicObject
func (a *BasicApiService) PostCrmV3ObjectsCompaniesCreateExecute(r ApiPostCrmV3ObjectsCompaniesCreateRequest) (*SimplePublicObject, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SimplePublicObject
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BasicApiService.PostCrmV3ObjectsCompaniesCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/objects/companies"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.simplePublicObjectInputForCreate == nil {
		return localVarReturnValue, nil, reportError("simplePublicObjectInputForCreate is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.simplePublicObjectInputForCreate
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["private_apps_legacy"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["private-app-legacy"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["private_apps"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["private-app"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
