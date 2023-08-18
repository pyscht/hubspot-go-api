# \AttendanceSubscriberStateChangesApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostMarketingV3MarketingEventsAttendanceExternalEventIdSubscriberStateCreateCreate**](AttendanceSubscriberStateChangesApi.md#PostMarketingV3MarketingEventsAttendanceExternalEventIdSubscriberStateCreateCreate) | **Post** /marketing/v3/marketing-events/attendance/{externalEventId}/{subscriberState}/create | Record
[**PostMarketingV3MarketingEventsAttendanceExternalEventIdSubscriberStateEmailCreateCreateByEmail**](AttendanceSubscriberStateChangesApi.md#PostMarketingV3MarketingEventsAttendanceExternalEventIdSubscriberStateEmailCreateCreateByEmail) | **Post** /marketing/v3/marketing-events/attendance/{externalEventId}/{subscriberState}/email-create | Record



## PostMarketingV3MarketingEventsAttendanceExternalEventIdSubscriberStateCreateCreate

> BatchResponseSubscriberVidResponse PostMarketingV3MarketingEventsAttendanceExternalEventIdSubscriberStateCreateCreate(ctx, externalEventId, subscriberState).BatchInputMarketingEventSubscriber(batchInputMarketingEventSubscriber).ExternalAccountId(externalAccountId).Execute()

Record



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
    externalEventId := "externalEventId_example" // string | The id of the marketing event
    subscriberState := "subscriberState_example" // string | The new subscriber state for the HubSpot contacts and the specified marketing event. For example: 'register', 'attend' or 'cancel'.
    batchInputMarketingEventSubscriber := *openapiclient.NewBatchInputMarketingEventSubscriber([]openapiclient.MarketingEventSubscriber{*openapiclient.NewMarketingEventSubscriber(int64(123))}) // BatchInputMarketingEventSubscriber | The details of the contacts to subscribe to the event. Parameters of join and left time if state is Attended.
    externalAccountId := "externalAccountId_example" // string | The account id associated with the marketing event (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AttendanceSubscriberStateChangesApi.PostMarketingV3MarketingEventsAttendanceExternalEventIdSubscriberStateCreateCreate(context.Background(), externalEventId, subscriberState).BatchInputMarketingEventSubscriber(batchInputMarketingEventSubscriber).ExternalAccountId(externalAccountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AttendanceSubscriberStateChangesApi.PostMarketingV3MarketingEventsAttendanceExternalEventIdSubscriberStateCreateCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostMarketingV3MarketingEventsAttendanceExternalEventIdSubscriberStateCreateCreate`: BatchResponseSubscriberVidResponse
    fmt.Fprintf(os.Stdout, "Response from `AttendanceSubscriberStateChangesApi.PostMarketingV3MarketingEventsAttendanceExternalEventIdSubscriberStateCreateCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalEventId** | **string** | The id of the marketing event | 
**subscriberState** | **string** | The new subscriber state for the HubSpot contacts and the specified marketing event. For example: &#39;register&#39;, &#39;attend&#39; or &#39;cancel&#39;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMarketingV3MarketingEventsAttendanceExternalEventIdSubscriberStateCreateCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **batchInputMarketingEventSubscriber** | [**BatchInputMarketingEventSubscriber**](BatchInputMarketingEventSubscriber.md) | The details of the contacts to subscribe to the event. Parameters of join and left time if state is Attended. | 
 **externalAccountId** | **string** | The account id associated with the marketing event | 

### Return type

[**BatchResponseSubscriberVidResponse**](BatchResponseSubscriberVidResponse.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMarketingV3MarketingEventsAttendanceExternalEventIdSubscriberStateEmailCreateCreateByEmail

> BatchResponseSubscriberEmailResponse PostMarketingV3MarketingEventsAttendanceExternalEventIdSubscriberStateEmailCreateCreateByEmail(ctx, externalEventId, subscriberState).BatchInputMarketingEventEmailSubscriber(batchInputMarketingEventEmailSubscriber).ExternalAccountId(externalAccountId).Execute()

Record



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
    externalEventId := "externalEventId_example" // string | The id of the marketing event
    subscriberState := "subscriberState_example" // string | The new subscriber state for the HubSpot contacts and the specified marketing event. For example: 'register', 'attend' or 'cancel'.
    batchInputMarketingEventEmailSubscriber := *openapiclient.NewBatchInputMarketingEventEmailSubscriber([]openapiclient.MarketingEventEmailSubscriber{*openapiclient.NewMarketingEventEmailSubscriber(int64(123), "Email_example")}) // BatchInputMarketingEventEmailSubscriber | The details of the contacts to subscribe to the event. Parameters of join and left time if state is Attended.
    externalAccountId := "externalAccountId_example" // string | The account id associated with the marketing event (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AttendanceSubscriberStateChangesApi.PostMarketingV3MarketingEventsAttendanceExternalEventIdSubscriberStateEmailCreateCreateByEmail(context.Background(), externalEventId, subscriberState).BatchInputMarketingEventEmailSubscriber(batchInputMarketingEventEmailSubscriber).ExternalAccountId(externalAccountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AttendanceSubscriberStateChangesApi.PostMarketingV3MarketingEventsAttendanceExternalEventIdSubscriberStateEmailCreateCreateByEmail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostMarketingV3MarketingEventsAttendanceExternalEventIdSubscriberStateEmailCreateCreateByEmail`: BatchResponseSubscriberEmailResponse
    fmt.Fprintf(os.Stdout, "Response from `AttendanceSubscriberStateChangesApi.PostMarketingV3MarketingEventsAttendanceExternalEventIdSubscriberStateEmailCreateCreateByEmail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalEventId** | **string** | The id of the marketing event | 
**subscriberState** | **string** | The new subscriber state for the HubSpot contacts and the specified marketing event. For example: &#39;register&#39;, &#39;attend&#39; or &#39;cancel&#39;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMarketingV3MarketingEventsAttendanceExternalEventIdSubscriberStateEmailCreateCreateByEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **batchInputMarketingEventEmailSubscriber** | [**BatchInputMarketingEventEmailSubscriber**](BatchInputMarketingEventEmailSubscriber.md) | The details of the contacts to subscribe to the event. Parameters of join and left time if state is Attended. | 
 **externalAccountId** | **string** | The account id associated with the marketing event | 

### Return type

[**BatchResponseSubscriberEmailResponse**](BatchResponseSubscriberEmailResponse.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

