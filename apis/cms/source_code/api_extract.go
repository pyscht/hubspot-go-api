/*
CMS Source Code

Endpoints for interacting with files in the CMS Developer File System. These files include HTML templates, CSS, JS, modules, and other assets which are used to create CMS content.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package source_code

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// ExtractApiService ExtractApi service
type ExtractApiService service

type ApiPostCmsV3SourceCodeExtractPathExtractByPathRequest struct {
	ctx        context.Context
	ApiService *ExtractApiService
	path       string
}

func (r ApiPostCmsV3SourceCodeExtractPathExtractByPathRequest) Execute() (*http.Response, error) {
	return r.ApiService.PostCmsV3SourceCodeExtractPathExtractByPathExecute(r)
}

/*
PostCmsV3SourceCodeExtractPathExtractByPath Extracts a zip file

Extracts a zip file in the file system. The zip file will be extracted in-place and not be deleted automatically.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param path The file system location of the zip file.
	@return ApiPostCmsV3SourceCodeExtractPathExtractByPathRequest

Deprecated
*/
func (a *ExtractApiService) PostCmsV3SourceCodeExtractPathExtractByPath(ctx context.Context, path string) ApiPostCmsV3SourceCodeExtractPathExtractByPathRequest {
	return ApiPostCmsV3SourceCodeExtractPathExtractByPathRequest{
		ApiService: a,
		ctx:        ctx,
		path:       path,
	}
}

// Execute executes the request
// Deprecated
func (a *ExtractApiService) PostCmsV3SourceCodeExtractPathExtractByPathExecute(r ApiPostCmsV3SourceCodeExtractPathExtractByPathRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExtractApiService.PostCmsV3SourceCodeExtractPathExtractByPath")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cms/v3/source-code/extract/{path}"
	localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", url.PathEscape(parameterValueToString(r.path, "path")), -1)

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