/*
Quotes

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package quotes

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)

// SearchApiService SearchApi service
type SearchApiService service

type ApiPostCrmV3ObjectsQuotesSearchDoSearchRequest struct {
	ctx                       context.Context
	ApiService                *SearchApiService
	publicObjectSearchRequest *PublicObjectSearchRequest
}

func (r ApiPostCrmV3ObjectsQuotesSearchDoSearchRequest) PublicObjectSearchRequest(publicObjectSearchRequest PublicObjectSearchRequest) ApiPostCrmV3ObjectsQuotesSearchDoSearchRequest {
	r.publicObjectSearchRequest = &publicObjectSearchRequest
	return r
}

func (r ApiPostCrmV3ObjectsQuotesSearchDoSearchRequest) Execute() (*CollectionResponseWithTotalSimplePublicObjectForwardPaging, *http.Response, error) {
	return r.ApiService.PostCrmV3ObjectsQuotesSearchDoSearchExecute(r)
}

/*
PostCrmV3ObjectsQuotesSearchDoSearch Method for PostCrmV3ObjectsQuotesSearchDoSearch

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiPostCrmV3ObjectsQuotesSearchDoSearchRequest
*/
func (a *SearchApiService) PostCrmV3ObjectsQuotesSearchDoSearch(ctx context.Context) ApiPostCrmV3ObjectsQuotesSearchDoSearchRequest {
	return ApiPostCrmV3ObjectsQuotesSearchDoSearchRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CollectionResponseWithTotalSimplePublicObjectForwardPaging
func (a *SearchApiService) PostCrmV3ObjectsQuotesSearchDoSearchExecute(r ApiPostCrmV3ObjectsQuotesSearchDoSearchRequest) (*CollectionResponseWithTotalSimplePublicObjectForwardPaging, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CollectionResponseWithTotalSimplePublicObjectForwardPaging
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SearchApiService.PostCrmV3ObjectsQuotesSearchDoSearch")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/objects/quotes/search"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.publicObjectSearchRequest == nil {
		return localVarReturnValue, nil, reportError("publicObjectSearchRequest is required and must be specified")
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
	localVarPostBody = r.publicObjectSearchRequest
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
