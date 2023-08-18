/*
Properties

All HubSpot objects store data in default and custom properties. These endpoints provide access to read and modify object properties in HubSpot.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package properties

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// CoreApiService CoreApi service
type CoreApiService service

type ApiDeleteCrmV3PropertiesObjectTypePropertyNameArchiveRequest struct {
	ctx          context.Context
	ApiService   *CoreApiService
	objectType   string
	propertyName string
}

func (r ApiDeleteCrmV3PropertiesObjectTypePropertyNameArchiveRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteCrmV3PropertiesObjectTypePropertyNameArchiveExecute(r)
}

/*
DeleteCrmV3PropertiesObjectTypePropertyNameArchive Archive a property

Move a property identified by {propertyName} to the recycling bin.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param objectType
	@param propertyName
	@return ApiDeleteCrmV3PropertiesObjectTypePropertyNameArchiveRequest
*/
func (a *CoreApiService) DeleteCrmV3PropertiesObjectTypePropertyNameArchive(ctx context.Context, objectType string, propertyName string) ApiDeleteCrmV3PropertiesObjectTypePropertyNameArchiveRequest {
	return ApiDeleteCrmV3PropertiesObjectTypePropertyNameArchiveRequest{
		ApiService:   a,
		ctx:          ctx,
		objectType:   objectType,
		propertyName: propertyName,
	}
}

// Execute executes the request
func (a *CoreApiService) DeleteCrmV3PropertiesObjectTypePropertyNameArchiveExecute(r ApiDeleteCrmV3PropertiesObjectTypePropertyNameArchiveRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CoreApiService.DeleteCrmV3PropertiesObjectTypePropertyNameArchive")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/properties/{objectType}/{propertyName}"
	localVarPath = strings.Replace(localVarPath, "{"+"objectType"+"}", url.PathEscape(parameterValueToString(r.objectType, "objectType")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"propertyName"+"}", url.PathEscape(parameterValueToString(r.propertyName, "propertyName")), -1)

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

type ApiGetCrmV3PropertiesObjectTypeGetAllRequest struct {
	ctx        context.Context
	ApiService *CoreApiService
	objectType string
	archived   *bool
	properties *string
}

// Whether to return only results that have been archived.
func (r ApiGetCrmV3PropertiesObjectTypeGetAllRequest) Archived(archived bool) ApiGetCrmV3PropertiesObjectTypeGetAllRequest {
	r.archived = &archived
	return r
}

func (r ApiGetCrmV3PropertiesObjectTypeGetAllRequest) Properties(properties string) ApiGetCrmV3PropertiesObjectTypeGetAllRequest {
	r.properties = &properties
	return r
}

func (r ApiGetCrmV3PropertiesObjectTypeGetAllRequest) Execute() (*CollectionResponsePropertyNoPaging, *http.Response, error) {
	return r.ApiService.GetCrmV3PropertiesObjectTypeGetAllExecute(r)
}

/*
GetCrmV3PropertiesObjectTypeGetAll Read all properties

Read all existing properties for the specified object type and HubSpot account.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param objectType
	@return ApiGetCrmV3PropertiesObjectTypeGetAllRequest
*/
func (a *CoreApiService) GetCrmV3PropertiesObjectTypeGetAll(ctx context.Context, objectType string) ApiGetCrmV3PropertiesObjectTypeGetAllRequest {
	return ApiGetCrmV3PropertiesObjectTypeGetAllRequest{
		ApiService: a,
		ctx:        ctx,
		objectType: objectType,
	}
}

// Execute executes the request
//
//	@return CollectionResponsePropertyNoPaging
func (a *CoreApiService) GetCrmV3PropertiesObjectTypeGetAllExecute(r ApiGetCrmV3PropertiesObjectTypeGetAllRequest) (*CollectionResponsePropertyNoPaging, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CollectionResponsePropertyNoPaging
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CoreApiService.GetCrmV3PropertiesObjectTypeGetAll")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/properties/{objectType}"
	localVarPath = strings.Replace(localVarPath, "{"+"objectType"+"}", url.PathEscape(parameterValueToString(r.objectType, "objectType")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.archived != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "archived", r.archived, "")
	}
	if r.properties != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "properties", r.properties, "")
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

type ApiGetCrmV3PropertiesObjectTypePropertyNameGetByNameRequest struct {
	ctx          context.Context
	ApiService   *CoreApiService
	objectType   string
	propertyName string
	archived     *bool
	properties   *string
}

// Whether to return only results that have been archived.
func (r ApiGetCrmV3PropertiesObjectTypePropertyNameGetByNameRequest) Archived(archived bool) ApiGetCrmV3PropertiesObjectTypePropertyNameGetByNameRequest {
	r.archived = &archived
	return r
}

func (r ApiGetCrmV3PropertiesObjectTypePropertyNameGetByNameRequest) Properties(properties string) ApiGetCrmV3PropertiesObjectTypePropertyNameGetByNameRequest {
	r.properties = &properties
	return r
}

func (r ApiGetCrmV3PropertiesObjectTypePropertyNameGetByNameRequest) Execute() (*Property, *http.Response, error) {
	return r.ApiService.GetCrmV3PropertiesObjectTypePropertyNameGetByNameExecute(r)
}

/*
GetCrmV3PropertiesObjectTypePropertyNameGetByName Read a property

Read a property identified by {propertyName}.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param objectType
	@param propertyName
	@return ApiGetCrmV3PropertiesObjectTypePropertyNameGetByNameRequest
*/
func (a *CoreApiService) GetCrmV3PropertiesObjectTypePropertyNameGetByName(ctx context.Context, objectType string, propertyName string) ApiGetCrmV3PropertiesObjectTypePropertyNameGetByNameRequest {
	return ApiGetCrmV3PropertiesObjectTypePropertyNameGetByNameRequest{
		ApiService:   a,
		ctx:          ctx,
		objectType:   objectType,
		propertyName: propertyName,
	}
}

// Execute executes the request
//
//	@return Property
func (a *CoreApiService) GetCrmV3PropertiesObjectTypePropertyNameGetByNameExecute(r ApiGetCrmV3PropertiesObjectTypePropertyNameGetByNameRequest) (*Property, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Property
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CoreApiService.GetCrmV3PropertiesObjectTypePropertyNameGetByName")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/properties/{objectType}/{propertyName}"
	localVarPath = strings.Replace(localVarPath, "{"+"objectType"+"}", url.PathEscape(parameterValueToString(r.objectType, "objectType")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"propertyName"+"}", url.PathEscape(parameterValueToString(r.propertyName, "propertyName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.archived != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "archived", r.archived, "")
	}
	if r.properties != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "properties", r.properties, "")
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

type ApiPatchCrmV3PropertiesObjectTypePropertyNameUpdateRequest struct {
	ctx            context.Context
	ApiService     *CoreApiService
	objectType     string
	propertyName   string
	propertyUpdate *PropertyUpdate
}

func (r ApiPatchCrmV3PropertiesObjectTypePropertyNameUpdateRequest) PropertyUpdate(propertyUpdate PropertyUpdate) ApiPatchCrmV3PropertiesObjectTypePropertyNameUpdateRequest {
	r.propertyUpdate = &propertyUpdate
	return r
}

func (r ApiPatchCrmV3PropertiesObjectTypePropertyNameUpdateRequest) Execute() (*Property, *http.Response, error) {
	return r.ApiService.PatchCrmV3PropertiesObjectTypePropertyNameUpdateExecute(r)
}

/*
PatchCrmV3PropertiesObjectTypePropertyNameUpdate Update a property

Perform a partial update of a property identified by {propertyName}. Provided fields will be overwritten.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param objectType
	@param propertyName
	@return ApiPatchCrmV3PropertiesObjectTypePropertyNameUpdateRequest
*/
func (a *CoreApiService) PatchCrmV3PropertiesObjectTypePropertyNameUpdate(ctx context.Context, objectType string, propertyName string) ApiPatchCrmV3PropertiesObjectTypePropertyNameUpdateRequest {
	return ApiPatchCrmV3PropertiesObjectTypePropertyNameUpdateRequest{
		ApiService:   a,
		ctx:          ctx,
		objectType:   objectType,
		propertyName: propertyName,
	}
}

// Execute executes the request
//
//	@return Property
func (a *CoreApiService) PatchCrmV3PropertiesObjectTypePropertyNameUpdateExecute(r ApiPatchCrmV3PropertiesObjectTypePropertyNameUpdateRequest) (*Property, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Property
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CoreApiService.PatchCrmV3PropertiesObjectTypePropertyNameUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/properties/{objectType}/{propertyName}"
	localVarPath = strings.Replace(localVarPath, "{"+"objectType"+"}", url.PathEscape(parameterValueToString(r.objectType, "objectType")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"propertyName"+"}", url.PathEscape(parameterValueToString(r.propertyName, "propertyName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.propertyUpdate == nil {
		return localVarReturnValue, nil, reportError("propertyUpdate is required and must be specified")
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
	localVarPostBody = r.propertyUpdate
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

type ApiPostCrmV3PropertiesObjectTypeCreateRequest struct {
	ctx            context.Context
	ApiService     *CoreApiService
	objectType     string
	propertyCreate *PropertyCreate
}

func (r ApiPostCrmV3PropertiesObjectTypeCreateRequest) PropertyCreate(propertyCreate PropertyCreate) ApiPostCrmV3PropertiesObjectTypeCreateRequest {
	r.propertyCreate = &propertyCreate
	return r
}

func (r ApiPostCrmV3PropertiesObjectTypeCreateRequest) Execute() (*Property, *http.Response, error) {
	return r.ApiService.PostCrmV3PropertiesObjectTypeCreateExecute(r)
}

/*
PostCrmV3PropertiesObjectTypeCreate Create a property

Create and return a copy of a new property for the specified object type.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param objectType
	@return ApiPostCrmV3PropertiesObjectTypeCreateRequest
*/
func (a *CoreApiService) PostCrmV3PropertiesObjectTypeCreate(ctx context.Context, objectType string) ApiPostCrmV3PropertiesObjectTypeCreateRequest {
	return ApiPostCrmV3PropertiesObjectTypeCreateRequest{
		ApiService: a,
		ctx:        ctx,
		objectType: objectType,
	}
}

// Execute executes the request
//
//	@return Property
func (a *CoreApiService) PostCrmV3PropertiesObjectTypeCreateExecute(r ApiPostCrmV3PropertiesObjectTypeCreateRequest) (*Property, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Property
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CoreApiService.PostCrmV3PropertiesObjectTypeCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/properties/{objectType}"
	localVarPath = strings.Replace(localVarPath, "{"+"objectType"+"}", url.PathEscape(parameterValueToString(r.objectType, "objectType")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.propertyCreate == nil {
		return localVarReturnValue, nil, reportError("propertyCreate is required and must be specified")
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
	localVarPostBody = r.propertyCreate
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
