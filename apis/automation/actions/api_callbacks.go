/*
Custom Workflow Actions

Create custom workflow actions

API version: v4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package actions

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// CallbacksApiService CallbacksApi service
type CallbacksApiService service

type ApiPostAutomationV4ActionsCallbacksCallbackIdCompleteCompleteRequest struct {
	ctx                       context.Context
	ApiService                *CallbacksApiService
	callbackId                string
	callbackCompletionRequest *CallbackCompletionRequest
}

// The result of the completed action.
func (r ApiPostAutomationV4ActionsCallbacksCallbackIdCompleteCompleteRequest) CallbackCompletionRequest(callbackCompletionRequest CallbackCompletionRequest) ApiPostAutomationV4ActionsCallbacksCallbackIdCompleteCompleteRequest {
	r.callbackCompletionRequest = &callbackCompletionRequest
	return r
}

func (r ApiPostAutomationV4ActionsCallbacksCallbackIdCompleteCompleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.PostAutomationV4ActionsCallbacksCallbackIdCompleteCompleteExecute(r)
}

/*
PostAutomationV4ActionsCallbacksCallbackIdCompleteComplete Complete a callback

Completes the given action callback.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param callbackId The ID of the target app.
	@return ApiPostAutomationV4ActionsCallbacksCallbackIdCompleteCompleteRequest
*/
func (a *CallbacksApiService) PostAutomationV4ActionsCallbacksCallbackIdCompleteComplete(ctx context.Context, callbackId string) ApiPostAutomationV4ActionsCallbacksCallbackIdCompleteCompleteRequest {
	return ApiPostAutomationV4ActionsCallbacksCallbackIdCompleteCompleteRequest{
		ApiService: a,
		ctx:        ctx,
		callbackId: callbackId,
	}
}

// Execute executes the request
func (a *CallbacksApiService) PostAutomationV4ActionsCallbacksCallbackIdCompleteCompleteExecute(r ApiPostAutomationV4ActionsCallbacksCallbackIdCompleteCompleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CallbacksApiService.PostAutomationV4ActionsCallbacksCallbackIdCompleteComplete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/automation/v4/actions/callbacks/{callbackId}/complete"
	localVarPath = strings.Replace(localVarPath, "{"+"callbackId"+"}", url.PathEscape(parameterValueToString(r.callbackId, "callbackId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.callbackCompletionRequest == nil {
		return nil, reportError("callbackCompletionRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.callbackCompletionRequest
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

type ApiPostAutomationV4ActionsCallbacksCompleteCompleteBatchRequest struct {
	ctx                                      context.Context
	ApiService                               *CallbacksApiService
	batchInputCallbackCompletionBatchRequest *BatchInputCallbackCompletionBatchRequest
}

// The result of the completed action.
func (r ApiPostAutomationV4ActionsCallbacksCompleteCompleteBatchRequest) BatchInputCallbackCompletionBatchRequest(batchInputCallbackCompletionBatchRequest BatchInputCallbackCompletionBatchRequest) ApiPostAutomationV4ActionsCallbacksCompleteCompleteBatchRequest {
	r.batchInputCallbackCompletionBatchRequest = &batchInputCallbackCompletionBatchRequest
	return r
}

func (r ApiPostAutomationV4ActionsCallbacksCompleteCompleteBatchRequest) Execute() (*http.Response, error) {
	return r.ApiService.PostAutomationV4ActionsCallbacksCompleteCompleteBatchExecute(r)
}

/*
PostAutomationV4ActionsCallbacksCompleteCompleteBatch Complete a batch of callbacks

Completes the given action callbacks.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiPostAutomationV4ActionsCallbacksCompleteCompleteBatchRequest
*/
func (a *CallbacksApiService) PostAutomationV4ActionsCallbacksCompleteCompleteBatch(ctx context.Context) ApiPostAutomationV4ActionsCallbacksCompleteCompleteBatchRequest {
	return ApiPostAutomationV4ActionsCallbacksCompleteCompleteBatchRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *CallbacksApiService) PostAutomationV4ActionsCallbacksCompleteCompleteBatchExecute(r ApiPostAutomationV4ActionsCallbacksCompleteCompleteBatchRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CallbacksApiService.PostAutomationV4ActionsCallbacksCompleteCompleteBatch")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/automation/v4/actions/callbacks/complete"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.batchInputCallbackCompletionBatchRequest == nil {
		return nil, reportError("batchInputCallbackCompletionBatchRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.batchInputCallbackCompletionBatchRequest
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
