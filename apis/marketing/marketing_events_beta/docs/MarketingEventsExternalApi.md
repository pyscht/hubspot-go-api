# \MarketingEventsExternalApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMarketingV3MarketingEventsEventsExternalEventIdArchive**](MarketingEventsExternalApi.md#DeleteMarketingV3MarketingEventsEventsExternalEventIdArchive) | **Delete** /marketing/v3/marketing-events/events/{externalEventId} | 
[**GetMarketingV3MarketingEventsEventsExternalEventIdGetById**](MarketingEventsExternalApi.md#GetMarketingV3MarketingEventsEventsExternalEventIdGetById) | **Get** /marketing/v3/marketing-events/events/{externalEventId} | 
[**PatchMarketingV3MarketingEventsEventsExternalEventIdUpdate**](MarketingEventsExternalApi.md#PatchMarketingV3MarketingEventsEventsExternalEventIdUpdate) | **Patch** /marketing/v3/marketing-events/events/{externalEventId} | 
[**PostMarketingV3MarketingEventsEventsCreate**](MarketingEventsExternalApi.md#PostMarketingV3MarketingEventsEventsCreate) | **Post** /marketing/v3/marketing-events/events | 
[**PostMarketingV3MarketingEventsEventsDeleteArchiveBatch**](MarketingEventsExternalApi.md#PostMarketingV3MarketingEventsEventsDeleteArchiveBatch) | **Post** /marketing/v3/marketing-events/events/delete | 
[**PostMarketingV3MarketingEventsEventsExternalEventIdCancelDoCancel**](MarketingEventsExternalApi.md#PostMarketingV3MarketingEventsEventsExternalEventIdCancelDoCancel) | **Post** /marketing/v3/marketing-events/events/{externalEventId}/cancel | 
[**PostMarketingV3MarketingEventsEventsExternalEventIdCompleteComplete**](MarketingEventsExternalApi.md#PostMarketingV3MarketingEventsEventsExternalEventIdCompleteComplete) | **Post** /marketing/v3/marketing-events/events/{externalEventId}/complete | 
[**PostMarketingV3MarketingEventsEventsExternalEventIdSubscriberStateEmailUpsertDoEmailUpsertById**](MarketingEventsExternalApi.md#PostMarketingV3MarketingEventsEventsExternalEventIdSubscriberStateEmailUpsertDoEmailUpsertById) | **Post** /marketing/v3/marketing-events/events/{externalEventId}/{subscriberState}/email-upsert | 
[**PostMarketingV3MarketingEventsEventsExternalEventIdSubscriberStateUpsertDoUpsertById**](MarketingEventsExternalApi.md#PostMarketingV3MarketingEventsEventsExternalEventIdSubscriberStateUpsertDoUpsertById) | **Post** /marketing/v3/marketing-events/events/{externalEventId}/{subscriberState}/upsert | 
[**PostMarketingV3MarketingEventsEventsUpsertDoUpsert**](MarketingEventsExternalApi.md#PostMarketingV3MarketingEventsEventsUpsertDoUpsert) | **Post** /marketing/v3/marketing-events/events/upsert | 
[**PutMarketingV3MarketingEventsEventsExternalEventIdReplace**](MarketingEventsExternalApi.md#PutMarketingV3MarketingEventsEventsExternalEventIdReplace) | **Put** /marketing/v3/marketing-events/events/{externalEventId} | 



## DeleteMarketingV3MarketingEventsEventsExternalEventIdArchive

> DeleteMarketingV3MarketingEventsEventsExternalEventIdArchive(ctx, externalEventId).ExternalAccountId(externalAccountId).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/pyscht/hubspot-go-api/apis/marketing/marketing_events_beta"
)

func main() {
    externalEventId := "externalEventId_example" // string | 
    externalAccountId := "externalAccountId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MarketingEventsExternalApi.DeleteMarketingV3MarketingEventsEventsExternalEventIdArchive(context.Background(), externalEventId).ExternalAccountId(externalAccountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketingEventsExternalApi.DeleteMarketingV3MarketingEventsEventsExternalEventIdArchive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalEventId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMarketingV3MarketingEventsEventsExternalEventIdArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **externalAccountId** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketingV3MarketingEventsEventsExternalEventIdGetById

> MarketingEventPublicReadResponse GetMarketingV3MarketingEventsEventsExternalEventIdGetById(ctx, externalEventId).ExternalAccountId(externalAccountId).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/pyscht/hubspot-go-api/apis/marketing/marketing_events_beta"
)

func main() {
    externalEventId := "externalEventId_example" // string | 
    externalAccountId := "externalAccountId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MarketingEventsExternalApi.GetMarketingV3MarketingEventsEventsExternalEventIdGetById(context.Background(), externalEventId).ExternalAccountId(externalAccountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketingEventsExternalApi.GetMarketingV3MarketingEventsEventsExternalEventIdGetById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMarketingV3MarketingEventsEventsExternalEventIdGetById`: MarketingEventPublicReadResponse
    fmt.Fprintf(os.Stdout, "Response from `MarketingEventsExternalApi.GetMarketingV3MarketingEventsEventsExternalEventIdGetById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalEventId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketingV3MarketingEventsEventsExternalEventIdGetByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **externalAccountId** | **string** |  | 

### Return type

[**MarketingEventPublicReadResponse**](MarketingEventPublicReadResponse.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchMarketingV3MarketingEventsEventsExternalEventIdUpdate

> MarketingEventPublicDefaultResponse PatchMarketingV3MarketingEventsEventsExternalEventIdUpdate(ctx, externalEventId).ExternalAccountId(externalAccountId).MarketingEventUpdateRequestParams(marketingEventUpdateRequestParams).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/pyscht/hubspot-go-api/apis/marketing/marketing_events_beta"
)

func main() {
    externalEventId := "externalEventId_example" // string | 
    externalAccountId := "externalAccountId_example" // string | 
    marketingEventUpdateRequestParams := *openapiclient.NewMarketingEventUpdateRequestParams() // MarketingEventUpdateRequestParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MarketingEventsExternalApi.PatchMarketingV3MarketingEventsEventsExternalEventIdUpdate(context.Background(), externalEventId).ExternalAccountId(externalAccountId).MarketingEventUpdateRequestParams(marketingEventUpdateRequestParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketingEventsExternalApi.PatchMarketingV3MarketingEventsEventsExternalEventIdUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchMarketingV3MarketingEventsEventsExternalEventIdUpdate`: MarketingEventPublicDefaultResponse
    fmt.Fprintf(os.Stdout, "Response from `MarketingEventsExternalApi.PatchMarketingV3MarketingEventsEventsExternalEventIdUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalEventId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchMarketingV3MarketingEventsEventsExternalEventIdUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **externalAccountId** | **string** |  | 
 **marketingEventUpdateRequestParams** | [**MarketingEventUpdateRequestParams**](MarketingEventUpdateRequestParams.md) |  | 

### Return type

[**MarketingEventPublicDefaultResponse**](MarketingEventPublicDefaultResponse.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMarketingV3MarketingEventsEventsCreate

> MarketingEventDefaultResponse PostMarketingV3MarketingEventsEventsCreate(ctx).MarketingEventCreateRequestParams(marketingEventCreateRequestParams).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/pyscht/hubspot-go-api/apis/marketing/marketing_events_beta"
)

func main() {
    marketingEventCreateRequestParams := *openapiclient.NewMarketingEventCreateRequestParams("EventName_example", "EventOrganizer_example", "ExternalAccountId_example", "ExternalEventId_example") // MarketingEventCreateRequestParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MarketingEventsExternalApi.PostMarketingV3MarketingEventsEventsCreate(context.Background()).MarketingEventCreateRequestParams(marketingEventCreateRequestParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketingEventsExternalApi.PostMarketingV3MarketingEventsEventsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostMarketingV3MarketingEventsEventsCreate`: MarketingEventDefaultResponse
    fmt.Fprintf(os.Stdout, "Response from `MarketingEventsExternalApi.PostMarketingV3MarketingEventsEventsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostMarketingV3MarketingEventsEventsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **marketingEventCreateRequestParams** | [**MarketingEventCreateRequestParams**](MarketingEventCreateRequestParams.md) |  | 

### Return type

[**MarketingEventDefaultResponse**](MarketingEventDefaultResponse.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMarketingV3MarketingEventsEventsDeleteArchiveBatch

> Error PostMarketingV3MarketingEventsEventsDeleteArchiveBatch(ctx).BatchInputMarketingEventExternalUniqueIdentifier(batchInputMarketingEventExternalUniqueIdentifier).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/pyscht/hubspot-go-api/apis/marketing/marketing_events_beta"
)

func main() {
    batchInputMarketingEventExternalUniqueIdentifier := *openapiclient.NewBatchInputMarketingEventExternalUniqueIdentifier([]openapiclient.MarketingEventExternalUniqueIdentifier{*openapiclient.NewMarketingEventExternalUniqueIdentifier(int32(123), "ExternalAccountId_example", "ExternalEventId_example")}) // BatchInputMarketingEventExternalUniqueIdentifier | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MarketingEventsExternalApi.PostMarketingV3MarketingEventsEventsDeleteArchiveBatch(context.Background()).BatchInputMarketingEventExternalUniqueIdentifier(batchInputMarketingEventExternalUniqueIdentifier).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketingEventsExternalApi.PostMarketingV3MarketingEventsEventsDeleteArchiveBatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostMarketingV3MarketingEventsEventsDeleteArchiveBatch`: Error
    fmt.Fprintf(os.Stdout, "Response from `MarketingEventsExternalApi.PostMarketingV3MarketingEventsEventsDeleteArchiveBatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostMarketingV3MarketingEventsEventsDeleteArchiveBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputMarketingEventExternalUniqueIdentifier** | [**BatchInputMarketingEventExternalUniqueIdentifier**](BatchInputMarketingEventExternalUniqueIdentifier.md) |  | 

### Return type

[**Error**](Error.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMarketingV3MarketingEventsEventsExternalEventIdCancelDoCancel

> MarketingEventDefaultResponse PostMarketingV3MarketingEventsEventsExternalEventIdCancelDoCancel(ctx, externalEventId).ExternalAccountId(externalAccountId).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/pyscht/hubspot-go-api/apis/marketing/marketing_events_beta"
)

func main() {
    externalEventId := "externalEventId_example" // string | 
    externalAccountId := "externalAccountId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MarketingEventsExternalApi.PostMarketingV3MarketingEventsEventsExternalEventIdCancelDoCancel(context.Background(), externalEventId).ExternalAccountId(externalAccountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketingEventsExternalApi.PostMarketingV3MarketingEventsEventsExternalEventIdCancelDoCancel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostMarketingV3MarketingEventsEventsExternalEventIdCancelDoCancel`: MarketingEventDefaultResponse
    fmt.Fprintf(os.Stdout, "Response from `MarketingEventsExternalApi.PostMarketingV3MarketingEventsEventsExternalEventIdCancelDoCancel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalEventId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMarketingV3MarketingEventsEventsExternalEventIdCancelDoCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **externalAccountId** | **string** |  | 

### Return type

[**MarketingEventDefaultResponse**](MarketingEventDefaultResponse.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMarketingV3MarketingEventsEventsExternalEventIdCompleteComplete

> MarketingEventDefaultResponse PostMarketingV3MarketingEventsEventsExternalEventIdCompleteComplete(ctx, externalEventId).ExternalAccountId(externalAccountId).MarketingEventCompleteRequestParams(marketingEventCompleteRequestParams).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/pyscht/hubspot-go-api/apis/marketing/marketing_events_beta"
)

func main() {
    externalEventId := "externalEventId_example" // string | 
    externalAccountId := "externalAccountId_example" // string | 
    marketingEventCompleteRequestParams := *openapiclient.NewMarketingEventCompleteRequestParams(time.Now(), time.Now()) // MarketingEventCompleteRequestParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MarketingEventsExternalApi.PostMarketingV3MarketingEventsEventsExternalEventIdCompleteComplete(context.Background(), externalEventId).ExternalAccountId(externalAccountId).MarketingEventCompleteRequestParams(marketingEventCompleteRequestParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketingEventsExternalApi.PostMarketingV3MarketingEventsEventsExternalEventIdCompleteComplete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostMarketingV3MarketingEventsEventsExternalEventIdCompleteComplete`: MarketingEventDefaultResponse
    fmt.Fprintf(os.Stdout, "Response from `MarketingEventsExternalApi.PostMarketingV3MarketingEventsEventsExternalEventIdCompleteComplete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalEventId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMarketingV3MarketingEventsEventsExternalEventIdCompleteCompleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **externalAccountId** | **string** |  | 
 **marketingEventCompleteRequestParams** | [**MarketingEventCompleteRequestParams**](MarketingEventCompleteRequestParams.md) |  | 

### Return type

[**MarketingEventDefaultResponse**](MarketingEventDefaultResponse.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMarketingV3MarketingEventsEventsExternalEventIdSubscriberStateEmailUpsertDoEmailUpsertById

> Error PostMarketingV3MarketingEventsEventsExternalEventIdSubscriberStateEmailUpsertDoEmailUpsertById(ctx, externalEventId, subscriberState).ExternalAccountId(externalAccountId).BatchInputMarketingEventEmailSubscriber(batchInputMarketingEventEmailSubscriber).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/pyscht/hubspot-go-api/apis/marketing/marketing_events_beta"
)

func main() {
    externalEventId := "externalEventId_example" // string | 
    subscriberState := "subscriberState_example" // string | 
    externalAccountId := "externalAccountId_example" // string | 
    batchInputMarketingEventEmailSubscriber := *openapiclient.NewBatchInputMarketingEventEmailSubscriber([]openapiclient.MarketingEventEmailSubscriber{*openapiclient.NewMarketingEventEmailSubscriber(int64(123), "Email_example")}) // BatchInputMarketingEventEmailSubscriber | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MarketingEventsExternalApi.PostMarketingV3MarketingEventsEventsExternalEventIdSubscriberStateEmailUpsertDoEmailUpsertById(context.Background(), externalEventId, subscriberState).ExternalAccountId(externalAccountId).BatchInputMarketingEventEmailSubscriber(batchInputMarketingEventEmailSubscriber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketingEventsExternalApi.PostMarketingV3MarketingEventsEventsExternalEventIdSubscriberStateEmailUpsertDoEmailUpsertById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostMarketingV3MarketingEventsEventsExternalEventIdSubscriberStateEmailUpsertDoEmailUpsertById`: Error
    fmt.Fprintf(os.Stdout, "Response from `MarketingEventsExternalApi.PostMarketingV3MarketingEventsEventsExternalEventIdSubscriberStateEmailUpsertDoEmailUpsertById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalEventId** | **string** |  | 
**subscriberState** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMarketingV3MarketingEventsEventsExternalEventIdSubscriberStateEmailUpsertDoEmailUpsertByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **externalAccountId** | **string** |  | 
 **batchInputMarketingEventEmailSubscriber** | [**BatchInputMarketingEventEmailSubscriber**](BatchInputMarketingEventEmailSubscriber.md) |  | 

### Return type

[**Error**](Error.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMarketingV3MarketingEventsEventsExternalEventIdSubscriberStateUpsertDoUpsertById

> Error PostMarketingV3MarketingEventsEventsExternalEventIdSubscriberStateUpsertDoUpsertById(ctx, externalEventId, subscriberState).ExternalAccountId(externalAccountId).BatchInputMarketingEventSubscriber(batchInputMarketingEventSubscriber).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/pyscht/hubspot-go-api/apis/marketing/marketing_events_beta"
)

func main() {
    externalEventId := "externalEventId_example" // string | 
    subscriberState := "subscriberState_example" // string | 
    externalAccountId := "externalAccountId_example" // string | 
    batchInputMarketingEventSubscriber := *openapiclient.NewBatchInputMarketingEventSubscriber([]openapiclient.MarketingEventSubscriber{*openapiclient.NewMarketingEventSubscriber(int64(123))}) // BatchInputMarketingEventSubscriber | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MarketingEventsExternalApi.PostMarketingV3MarketingEventsEventsExternalEventIdSubscriberStateUpsertDoUpsertById(context.Background(), externalEventId, subscriberState).ExternalAccountId(externalAccountId).BatchInputMarketingEventSubscriber(batchInputMarketingEventSubscriber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketingEventsExternalApi.PostMarketingV3MarketingEventsEventsExternalEventIdSubscriberStateUpsertDoUpsertById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostMarketingV3MarketingEventsEventsExternalEventIdSubscriberStateUpsertDoUpsertById`: Error
    fmt.Fprintf(os.Stdout, "Response from `MarketingEventsExternalApi.PostMarketingV3MarketingEventsEventsExternalEventIdSubscriberStateUpsertDoUpsertById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalEventId** | **string** |  | 
**subscriberState** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMarketingV3MarketingEventsEventsExternalEventIdSubscriberStateUpsertDoUpsertByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **externalAccountId** | **string** |  | 
 **batchInputMarketingEventSubscriber** | [**BatchInputMarketingEventSubscriber**](BatchInputMarketingEventSubscriber.md) |  | 

### Return type

[**Error**](Error.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMarketingV3MarketingEventsEventsUpsertDoUpsert

> BatchResponseMarketingEventPublicDefaultResponse PostMarketingV3MarketingEventsEventsUpsertDoUpsert(ctx).BatchInputMarketingEventCreateRequestParams(batchInputMarketingEventCreateRequestParams).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/pyscht/hubspot-go-api/apis/marketing/marketing_events_beta"
)

func main() {
    batchInputMarketingEventCreateRequestParams := *openapiclient.NewBatchInputMarketingEventCreateRequestParams([]openapiclient.MarketingEventCreateRequestParams{*openapiclient.NewMarketingEventCreateRequestParams("EventName_example", "EventOrganizer_example", "ExternalAccountId_example", "ExternalEventId_example")}) // BatchInputMarketingEventCreateRequestParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MarketingEventsExternalApi.PostMarketingV3MarketingEventsEventsUpsertDoUpsert(context.Background()).BatchInputMarketingEventCreateRequestParams(batchInputMarketingEventCreateRequestParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketingEventsExternalApi.PostMarketingV3MarketingEventsEventsUpsertDoUpsert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostMarketingV3MarketingEventsEventsUpsertDoUpsert`: BatchResponseMarketingEventPublicDefaultResponse
    fmt.Fprintf(os.Stdout, "Response from `MarketingEventsExternalApi.PostMarketingV3MarketingEventsEventsUpsertDoUpsert`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostMarketingV3MarketingEventsEventsUpsertDoUpsertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputMarketingEventCreateRequestParams** | [**BatchInputMarketingEventCreateRequestParams**](BatchInputMarketingEventCreateRequestParams.md) |  | 

### Return type

[**BatchResponseMarketingEventPublicDefaultResponse**](BatchResponseMarketingEventPublicDefaultResponse.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutMarketingV3MarketingEventsEventsExternalEventIdReplace

> MarketingEventPublicDefaultResponse PutMarketingV3MarketingEventsEventsExternalEventIdReplace(ctx, externalEventId).MarketingEventCreateRequestParams(marketingEventCreateRequestParams).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/pyscht/hubspot-go-api/apis/marketing/marketing_events_beta"
)

func main() {
    externalEventId := "externalEventId_example" // string | 
    marketingEventCreateRequestParams := *openapiclient.NewMarketingEventCreateRequestParams("EventName_example", "EventOrganizer_example", "ExternalAccountId_example", "ExternalEventId_example") // MarketingEventCreateRequestParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MarketingEventsExternalApi.PutMarketingV3MarketingEventsEventsExternalEventIdReplace(context.Background(), externalEventId).MarketingEventCreateRequestParams(marketingEventCreateRequestParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketingEventsExternalApi.PutMarketingV3MarketingEventsEventsExternalEventIdReplace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutMarketingV3MarketingEventsEventsExternalEventIdReplace`: MarketingEventPublicDefaultResponse
    fmt.Fprintf(os.Stdout, "Response from `MarketingEventsExternalApi.PutMarketingV3MarketingEventsEventsExternalEventIdReplace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalEventId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutMarketingV3MarketingEventsEventsExternalEventIdReplaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **marketingEventCreateRequestParams** | [**MarketingEventCreateRequestParams**](MarketingEventCreateRequestParams.md) |  | 

### Return type

[**MarketingEventPublicDefaultResponse**](MarketingEventPublicDefaultResponse.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

