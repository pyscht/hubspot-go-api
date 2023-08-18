/*
HubSpot Events API

API for accessing CRM object events.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package events

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"time"
)

// EventsApiService EventsApi service
type EventsApiService service

type ApiGetEventsV3EventsGetPageRequest struct {
	ctx            context.Context
	ApiService     *EventsApiService
	occurredAfter  *time.Time
	occurredBefore *time.Time
	objectType     *string
	objectId       *int64
	eventType      *string
	after          *string
	before         *string
	limit          *int32
	sort           *[]string
}

// The starting time as an ISO 8601 timestamp.
func (r ApiGetEventsV3EventsGetPageRequest) OccurredAfter(occurredAfter time.Time) ApiGetEventsV3EventsGetPageRequest {
	r.occurredAfter = &occurredAfter
	return r
}

// The ending time as an ISO 8601 timestamp.
func (r ApiGetEventsV3EventsGetPageRequest) OccurredBefore(occurredBefore time.Time) ApiGetEventsV3EventsGetPageRequest {
	r.occurredBefore = &occurredBefore
	return r
}

// The type of object being selected. Valid values are hubspot named object types (e.g. &#x60;contact&#x60;).
func (r ApiGetEventsV3EventsGetPageRequest) ObjectType(objectType string) ApiGetEventsV3EventsGetPageRequest {
	r.objectType = &objectType
	return r
}

// The id of the selected object. If not present, then the &#x60;objectProperty&#x60; parameter is required.
func (r ApiGetEventsV3EventsGetPageRequest) ObjectId(objectId int64) ApiGetEventsV3EventsGetPageRequest {
	r.objectId = &objectId
	return r
}

// Limits the response to the specified event type.  For example &#x60;&amp;eventType&#x3D;e_visited_page&#x60; returns only &#x60;e_visited_page&#x60; events.  If not present all event types are returned.
func (r ApiGetEventsV3EventsGetPageRequest) EventType(eventType string) ApiGetEventsV3EventsGetPageRequest {
	r.eventType = &eventType
	return r
}

// An additional parameter that may be used to get the next &#x60;limit&#x60; set of results.
func (r ApiGetEventsV3EventsGetPageRequest) After(after string) ApiGetEventsV3EventsGetPageRequest {
	r.after = &after
	return r
}

func (r ApiGetEventsV3EventsGetPageRequest) Before(before string) ApiGetEventsV3EventsGetPageRequest {
	r.before = &before
	return r
}

// The maximum number of events to return, defaults to 20.
func (r ApiGetEventsV3EventsGetPageRequest) Limit(limit int32) ApiGetEventsV3EventsGetPageRequest {
	r.limit = &limit
	return r
}

// Selects the sort field and order. Defaults to ascending, prefix with &#x60;-&#x60; for descending order. &#x60;occurredAt&#x60; is the only field supported for sorting.
func (r ApiGetEventsV3EventsGetPageRequest) Sort(sort []string) ApiGetEventsV3EventsGetPageRequest {
	r.sort = &sort
	return r
}

func (r ApiGetEventsV3EventsGetPageRequest) Execute() (*CollectionResponseExternalUnifiedEvent, *http.Response, error) {
	return r.ApiService.GetEventsV3EventsGetPageExecute(r)
}

/*
GetEventsV3EventsGetPage Returns a collection of events matching a query.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetEventsV3EventsGetPageRequest
*/
func (a *EventsApiService) GetEventsV3EventsGetPage(ctx context.Context) ApiGetEventsV3EventsGetPageRequest {
	return ApiGetEventsV3EventsGetPageRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CollectionResponseExternalUnifiedEvent
func (a *EventsApiService) GetEventsV3EventsGetPageExecute(r ApiGetEventsV3EventsGetPageRequest) (*CollectionResponseExternalUnifiedEvent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CollectionResponseExternalUnifiedEvent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EventsApiService.GetEventsV3EventsGetPage")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/events/v3/events"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.occurredAfter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "occurredAfter", r.occurredAfter, "")
	}
	if r.occurredBefore != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "occurredBefore", r.occurredBefore, "")
	}
	if r.objectType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "objectType", r.objectType, "")
	}
	if r.objectId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "objectId", r.objectId, "")
	}
	if r.eventType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "eventType", r.eventType, "")
	}
	if r.after != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "after", r.after, "")
	}
	if r.before != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "before", r.before, "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
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
