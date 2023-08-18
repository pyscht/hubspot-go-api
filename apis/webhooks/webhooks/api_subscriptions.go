/*
Webhooks API

Provides a way for apps to subscribe to certain change events in HubSpot. Once configured, apps will receive event payloads containing details about the changes at a specified target URL. There can only be one target URL for receiving event notifications per app.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package webhooks

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// SubscriptionsApiService SubscriptionsApi service
type SubscriptionsApiService service

type ApiDeleteWebhooksV3AppIdSubscriptionsSubscriptionIdArchiveRequest struct {
	ctx context.Context
	ApiService *SubscriptionsApiService
	subscriptionId int32
	appId int32
}

func (r ApiDeleteWebhooksV3AppIdSubscriptionsSubscriptionIdArchiveRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteWebhooksV3AppIdSubscriptionsSubscriptionIdArchiveExecute(r)
}

/*
DeleteWebhooksV3AppIdSubscriptionsSubscriptionIdArchive Method for DeleteWebhooksV3AppIdSubscriptionsSubscriptionIdArchive

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param subscriptionId
 @param appId
 @return ApiDeleteWebhooksV3AppIdSubscriptionsSubscriptionIdArchiveRequest
*/
func (a *SubscriptionsApiService) DeleteWebhooksV3AppIdSubscriptionsSubscriptionIdArchive(ctx context.Context, subscriptionId int32, appId int32) ApiDeleteWebhooksV3AppIdSubscriptionsSubscriptionIdArchiveRequest {
	return ApiDeleteWebhooksV3AppIdSubscriptionsSubscriptionIdArchiveRequest{
		ApiService: a,
		ctx: ctx,
		subscriptionId: subscriptionId,
		appId: appId,
	}
}

// Execute executes the request
func (a *SubscriptionsApiService) DeleteWebhooksV3AppIdSubscriptionsSubscriptionIdArchiveExecute(r ApiDeleteWebhooksV3AppIdSubscriptionsSubscriptionIdArchiveRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SubscriptionsApiService.DeleteWebhooksV3AppIdSubscriptionsSubscriptionIdArchive")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/webhooks/v3/{appId}/subscriptions/{subscriptionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"subscriptionId"+"}", url.PathEscape(parameterValueToString(r.subscriptionId, "subscriptionId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"appId"+"}", url.PathEscape(parameterValueToString(r.appId, "appId")), -1)

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
			if apiKey, ok := auth["developer_hapikey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("hapikey", key)
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

type ApiGetWebhooksV3AppIdSubscriptionsGetAllRequest struct {
	ctx context.Context
	ApiService *SubscriptionsApiService
	appId int32
}

func (r ApiGetWebhooksV3AppIdSubscriptionsGetAllRequest) Execute() (*SubscriptionListResponse, *http.Response, error) {
	return r.ApiService.GetWebhooksV3AppIdSubscriptionsGetAllExecute(r)
}

/*
GetWebhooksV3AppIdSubscriptionsGetAll Method for GetWebhooksV3AppIdSubscriptionsGetAll

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param appId
 @return ApiGetWebhooksV3AppIdSubscriptionsGetAllRequest
*/
func (a *SubscriptionsApiService) GetWebhooksV3AppIdSubscriptionsGetAll(ctx context.Context, appId int32) ApiGetWebhooksV3AppIdSubscriptionsGetAllRequest {
	return ApiGetWebhooksV3AppIdSubscriptionsGetAllRequest{
		ApiService: a,
		ctx: ctx,
		appId: appId,
	}
}

// Execute executes the request
//  @return SubscriptionListResponse
func (a *SubscriptionsApiService) GetWebhooksV3AppIdSubscriptionsGetAllExecute(r ApiGetWebhooksV3AppIdSubscriptionsGetAllRequest) (*SubscriptionListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SubscriptionListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SubscriptionsApiService.GetWebhooksV3AppIdSubscriptionsGetAll")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/webhooks/v3/{appId}/subscriptions"
	localVarPath = strings.Replace(localVarPath, "{"+"appId"+"}", url.PathEscape(parameterValueToString(r.appId, "appId")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/json", "*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["developer_hapikey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("hapikey", key)
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

type ApiGetWebhooksV3AppIdSubscriptionsSubscriptionIdGetByIdRequest struct {
	ctx context.Context
	ApiService *SubscriptionsApiService
	subscriptionId int32
	appId int32
}

func (r ApiGetWebhooksV3AppIdSubscriptionsSubscriptionIdGetByIdRequest) Execute() (*SubscriptionResponse, *http.Response, error) {
	return r.ApiService.GetWebhooksV3AppIdSubscriptionsSubscriptionIdGetByIdExecute(r)
}

/*
GetWebhooksV3AppIdSubscriptionsSubscriptionIdGetById Method for GetWebhooksV3AppIdSubscriptionsSubscriptionIdGetById

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param subscriptionId
 @param appId
 @return ApiGetWebhooksV3AppIdSubscriptionsSubscriptionIdGetByIdRequest
*/
func (a *SubscriptionsApiService) GetWebhooksV3AppIdSubscriptionsSubscriptionIdGetById(ctx context.Context, subscriptionId int32, appId int32) ApiGetWebhooksV3AppIdSubscriptionsSubscriptionIdGetByIdRequest {
	return ApiGetWebhooksV3AppIdSubscriptionsSubscriptionIdGetByIdRequest{
		ApiService: a,
		ctx: ctx,
		subscriptionId: subscriptionId,
		appId: appId,
	}
}

// Execute executes the request
//  @return SubscriptionResponse
func (a *SubscriptionsApiService) GetWebhooksV3AppIdSubscriptionsSubscriptionIdGetByIdExecute(r ApiGetWebhooksV3AppIdSubscriptionsSubscriptionIdGetByIdRequest) (*SubscriptionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SubscriptionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SubscriptionsApiService.GetWebhooksV3AppIdSubscriptionsSubscriptionIdGetById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/webhooks/v3/{appId}/subscriptions/{subscriptionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"subscriptionId"+"}", url.PathEscape(parameterValueToString(r.subscriptionId, "subscriptionId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"appId"+"}", url.PathEscape(parameterValueToString(r.appId, "appId")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/json", "*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["developer_hapikey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("hapikey", key)
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

type ApiPatchWebhooksV3AppIdSubscriptionsSubscriptionIdUpdateRequest struct {
	ctx context.Context
	ApiService *SubscriptionsApiService
	subscriptionId int32
	appId int32
	subscriptionPatchRequest *SubscriptionPatchRequest
}

func (r ApiPatchWebhooksV3AppIdSubscriptionsSubscriptionIdUpdateRequest) SubscriptionPatchRequest(subscriptionPatchRequest SubscriptionPatchRequest) ApiPatchWebhooksV3AppIdSubscriptionsSubscriptionIdUpdateRequest {
	r.subscriptionPatchRequest = &subscriptionPatchRequest
	return r
}

func (r ApiPatchWebhooksV3AppIdSubscriptionsSubscriptionIdUpdateRequest) Execute() (*SubscriptionResponse, *http.Response, error) {
	return r.ApiService.PatchWebhooksV3AppIdSubscriptionsSubscriptionIdUpdateExecute(r)
}

/*
PatchWebhooksV3AppIdSubscriptionsSubscriptionIdUpdate Method for PatchWebhooksV3AppIdSubscriptionsSubscriptionIdUpdate

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param subscriptionId
 @param appId
 @return ApiPatchWebhooksV3AppIdSubscriptionsSubscriptionIdUpdateRequest
*/
func (a *SubscriptionsApiService) PatchWebhooksV3AppIdSubscriptionsSubscriptionIdUpdate(ctx context.Context, subscriptionId int32, appId int32) ApiPatchWebhooksV3AppIdSubscriptionsSubscriptionIdUpdateRequest {
	return ApiPatchWebhooksV3AppIdSubscriptionsSubscriptionIdUpdateRequest{
		ApiService: a,
		ctx: ctx,
		subscriptionId: subscriptionId,
		appId: appId,
	}
}

// Execute executes the request
//  @return SubscriptionResponse
func (a *SubscriptionsApiService) PatchWebhooksV3AppIdSubscriptionsSubscriptionIdUpdateExecute(r ApiPatchWebhooksV3AppIdSubscriptionsSubscriptionIdUpdateRequest) (*SubscriptionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SubscriptionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SubscriptionsApiService.PatchWebhooksV3AppIdSubscriptionsSubscriptionIdUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/webhooks/v3/{appId}/subscriptions/{subscriptionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"subscriptionId"+"}", url.PathEscape(parameterValueToString(r.subscriptionId, "subscriptionId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"appId"+"}", url.PathEscape(parameterValueToString(r.appId, "appId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.subscriptionPatchRequest == nil {
		return localVarReturnValue, nil, reportError("subscriptionPatchRequest is required and must be specified")
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
	localVarPostBody = r.subscriptionPatchRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["developer_hapikey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("hapikey", key)
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

type ApiPostWebhooksV3AppIdSubscriptionsBatchUpdateUpdateBatchRequest struct {
	ctx context.Context
	ApiService *SubscriptionsApiService
	appId int32
	batchInputSubscriptionBatchUpdateRequest *BatchInputSubscriptionBatchUpdateRequest
}

func (r ApiPostWebhooksV3AppIdSubscriptionsBatchUpdateUpdateBatchRequest) BatchInputSubscriptionBatchUpdateRequest(batchInputSubscriptionBatchUpdateRequest BatchInputSubscriptionBatchUpdateRequest) ApiPostWebhooksV3AppIdSubscriptionsBatchUpdateUpdateBatchRequest {
	r.batchInputSubscriptionBatchUpdateRequest = &batchInputSubscriptionBatchUpdateRequest
	return r
}

func (r ApiPostWebhooksV3AppIdSubscriptionsBatchUpdateUpdateBatchRequest) Execute() (*BatchResponseSubscriptionResponse, *http.Response, error) {
	return r.ApiService.PostWebhooksV3AppIdSubscriptionsBatchUpdateUpdateBatchExecute(r)
}

/*
PostWebhooksV3AppIdSubscriptionsBatchUpdateUpdateBatch Method for PostWebhooksV3AppIdSubscriptionsBatchUpdateUpdateBatch

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param appId
 @return ApiPostWebhooksV3AppIdSubscriptionsBatchUpdateUpdateBatchRequest
*/
func (a *SubscriptionsApiService) PostWebhooksV3AppIdSubscriptionsBatchUpdateUpdateBatch(ctx context.Context, appId int32) ApiPostWebhooksV3AppIdSubscriptionsBatchUpdateUpdateBatchRequest {
	return ApiPostWebhooksV3AppIdSubscriptionsBatchUpdateUpdateBatchRequest{
		ApiService: a,
		ctx: ctx,
		appId: appId,
	}
}

// Execute executes the request
//  @return BatchResponseSubscriptionResponse
func (a *SubscriptionsApiService) PostWebhooksV3AppIdSubscriptionsBatchUpdateUpdateBatchExecute(r ApiPostWebhooksV3AppIdSubscriptionsBatchUpdateUpdateBatchRequest) (*BatchResponseSubscriptionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BatchResponseSubscriptionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SubscriptionsApiService.PostWebhooksV3AppIdSubscriptionsBatchUpdateUpdateBatch")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/webhooks/v3/{appId}/subscriptions/batch/update"
	localVarPath = strings.Replace(localVarPath, "{"+"appId"+"}", url.PathEscape(parameterValueToString(r.appId, "appId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.batchInputSubscriptionBatchUpdateRequest == nil {
		return localVarReturnValue, nil, reportError("batchInputSubscriptionBatchUpdateRequest is required and must be specified")
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
	localVarPostBody = r.batchInputSubscriptionBatchUpdateRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["developer_hapikey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("hapikey", key)
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

type ApiPostWebhooksV3AppIdSubscriptionsCreateRequest struct {
	ctx context.Context
	ApiService *SubscriptionsApiService
	appId int32
	subscriptionCreateRequest *SubscriptionCreateRequest
}

func (r ApiPostWebhooksV3AppIdSubscriptionsCreateRequest) SubscriptionCreateRequest(subscriptionCreateRequest SubscriptionCreateRequest) ApiPostWebhooksV3AppIdSubscriptionsCreateRequest {
	r.subscriptionCreateRequest = &subscriptionCreateRequest
	return r
}

func (r ApiPostWebhooksV3AppIdSubscriptionsCreateRequest) Execute() (*SubscriptionResponse, *http.Response, error) {
	return r.ApiService.PostWebhooksV3AppIdSubscriptionsCreateExecute(r)
}

/*
PostWebhooksV3AppIdSubscriptionsCreate Method for PostWebhooksV3AppIdSubscriptionsCreate

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param appId
 @return ApiPostWebhooksV3AppIdSubscriptionsCreateRequest
*/
func (a *SubscriptionsApiService) PostWebhooksV3AppIdSubscriptionsCreate(ctx context.Context, appId int32) ApiPostWebhooksV3AppIdSubscriptionsCreateRequest {
	return ApiPostWebhooksV3AppIdSubscriptionsCreateRequest{
		ApiService: a,
		ctx: ctx,
		appId: appId,
	}
}

// Execute executes the request
//  @return SubscriptionResponse
func (a *SubscriptionsApiService) PostWebhooksV3AppIdSubscriptionsCreateExecute(r ApiPostWebhooksV3AppIdSubscriptionsCreateRequest) (*SubscriptionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SubscriptionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SubscriptionsApiService.PostWebhooksV3AppIdSubscriptionsCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/webhooks/v3/{appId}/subscriptions"
	localVarPath = strings.Replace(localVarPath, "{"+"appId"+"}", url.PathEscape(parameterValueToString(r.appId, "appId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.subscriptionCreateRequest == nil {
		return localVarReturnValue, nil, reportError("subscriptionCreateRequest is required and must be specified")
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
	localVarPostBody = r.subscriptionCreateRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["developer_hapikey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("hapikey", key)
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
