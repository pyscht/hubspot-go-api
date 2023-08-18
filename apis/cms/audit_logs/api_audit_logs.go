/*
CMS Audit Logs

Use this endpoint to query audit logs of CMS changes that occurred on your HubSpot account.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package audit_logs

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"reflect"
)

// AuditLogsApiService AuditLogsApi service
type AuditLogsApiService service

type ApiGetCmsV3AuditLogsGetPageRequest struct {
	ctx        context.Context
	ApiService *AuditLogsApiService
	objectId   *[]string
	userId     *[]string
	after      *string
	before     *string
	sort       *[]string
	eventType  *[]string
	limit      *int32
	objectType *[]string
}

// Comma separated list of object ids to filter by.
func (r ApiGetCmsV3AuditLogsGetPageRequest) ObjectId(objectId []string) ApiGetCmsV3AuditLogsGetPageRequest {
	r.objectId = &objectId
	return r
}

// Comma separated list of user ids to filter by.
func (r ApiGetCmsV3AuditLogsGetPageRequest) UserId(userId []string) ApiGetCmsV3AuditLogsGetPageRequest {
	r.userId = &userId
	return r
}

// Timestamp after which audit logs will be returned
func (r ApiGetCmsV3AuditLogsGetPageRequest) After(after string) ApiGetCmsV3AuditLogsGetPageRequest {
	r.after = &after
	return r
}

// Timestamp before which audit logs will be returned
func (r ApiGetCmsV3AuditLogsGetPageRequest) Before(before string) ApiGetCmsV3AuditLogsGetPageRequest {
	r.before = &before
	return r
}

// The sort direction for the audit logs. (Can only sort by timestamp).
func (r ApiGetCmsV3AuditLogsGetPageRequest) Sort(sort []string) ApiGetCmsV3AuditLogsGetPageRequest {
	r.sort = &sort
	return r
}

// Comma separated list of event types to filter by (CREATED, UPDATED, PUBLISHED, DELETED, UNPUBLISHED).
func (r ApiGetCmsV3AuditLogsGetPageRequest) EventType(eventType []string) ApiGetCmsV3AuditLogsGetPageRequest {
	r.eventType = &eventType
	return r
}

// The number of logs to return.
func (r ApiGetCmsV3AuditLogsGetPageRequest) Limit(limit int32) ApiGetCmsV3AuditLogsGetPageRequest {
	r.limit = &limit
	return r
}

// Comma separated list of object types to filter by (BLOG, LANDING_PAGE, DOMAIN, HUBDB_TABLE etc.)
func (r ApiGetCmsV3AuditLogsGetPageRequest) ObjectType(objectType []string) ApiGetCmsV3AuditLogsGetPageRequest {
	r.objectType = &objectType
	return r
}

func (r ApiGetCmsV3AuditLogsGetPageRequest) Execute() (*CollectionResponsePublicAuditLog, *http.Response, error) {
	return r.ApiService.GetCmsV3AuditLogsGetPageExecute(r)
}

/*
GetCmsV3AuditLogsGetPage Query audit logs

Returns audit logs based on filters.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetCmsV3AuditLogsGetPageRequest
*/
func (a *AuditLogsApiService) GetCmsV3AuditLogsGetPage(ctx context.Context) ApiGetCmsV3AuditLogsGetPageRequest {
	return ApiGetCmsV3AuditLogsGetPageRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CollectionResponsePublicAuditLog
func (a *AuditLogsApiService) GetCmsV3AuditLogsGetPageExecute(r ApiGetCmsV3AuditLogsGetPageRequest) (*CollectionResponsePublicAuditLog, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CollectionResponsePublicAuditLog
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuditLogsApiService.GetCmsV3AuditLogsGetPage")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cms/v3/audit-logs/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.objectId != nil {
		t := *r.objectId
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "objectId", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "objectId", t, "multi")
		}
	}
	if r.userId != nil {
		t := *r.userId
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "userId", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "userId", t, "multi")
		}
	}
	if r.after != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "after", r.after, "")
	}
	if r.before != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "before", r.before, "")
	}
	if r.sort != nil {
		t := *r.sort
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "sort", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "sort", t, "multi")
		}
	}
	if r.eventType != nil {
		t := *r.eventType
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "eventType", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "eventType", t, "multi")
		}
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.objectType != nil {
		t := *r.objectType
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "objectType", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "objectType", t, "multi")
		}
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