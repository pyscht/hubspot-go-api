# \DomainsApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCmsV3DomainsDomainIdGetById**](DomainsApi.md#GetCmsV3DomainsDomainIdGetById) | **Get** /cms/v3/domains/{domainId} | Get a single domain
[**GetCmsV3DomainsGetPage**](DomainsApi.md#GetCmsV3DomainsGetPage) | **Get** /cms/v3/domains/ | Get current domains



## GetCmsV3DomainsDomainIdGetById

> Domain GetCmsV3DomainsDomainIdGetById(ctx, domainId).Execute()

Get a single domain



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/pyscht/hubspot-go-api/apis/cms/domains"
)

func main() {
    domainId := "domainId_example" // string | The unique ID of the domain.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DomainsApi.GetCmsV3DomainsDomainIdGetById(context.Background(), domainId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsApi.GetCmsV3DomainsDomainIdGetById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCmsV3DomainsDomainIdGetById`: Domain
    fmt.Fprintf(os.Stdout, "Response from `DomainsApi.GetCmsV3DomainsDomainIdGetById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **string** | The unique ID of the domain. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCmsV3DomainsDomainIdGetByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Domain**](Domain.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy), [oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCmsV3DomainsGetPage

> CollectionResponseWithTotalDomainForwardPaging GetCmsV3DomainsGetPage(ctx).CreatedAt(createdAt).CreatedAfter(createdAfter).CreatedBefore(createdBefore).UpdatedAt(updatedAt).UpdatedAfter(updatedAfter).UpdatedBefore(updatedBefore).Sort(sort).After(after).Limit(limit).Archived(archived).Execute()

Get current domains



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/pyscht/hubspot-go-api/apis/cms/domains"
)

func main() {
    createdAt := time.Now() // time.Time | Only return domains created at this date. (optional)
    createdAfter := time.Now() // time.Time | Only return domains created after this date. (optional)
    createdBefore := time.Now() // time.Time | Only return domains created before this date. (optional)
    updatedAt := time.Now() // time.Time | Only return domains updated at this date. (optional)
    updatedAfter := time.Now() // time.Time | Only return domains updated after this date. (optional)
    updatedBefore := time.Now() // time.Time | Only return domains updated before this date. (optional)
    sort := []string{"Inner_example"} // []string |  (optional)
    after := "after_example" // string | The paging cursor token of the last successfully read resource will be returned as the `paging.next.after` JSON property of a paged response containing more results. (optional)
    limit := int32(56) // int32 | Maximum number of results per page. (optional)
    archived := true // bool | Whether to return only results that have been archived. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DomainsApi.GetCmsV3DomainsGetPage(context.Background()).CreatedAt(createdAt).CreatedAfter(createdAfter).CreatedBefore(createdBefore).UpdatedAt(updatedAt).UpdatedAfter(updatedAfter).UpdatedBefore(updatedBefore).Sort(sort).After(after).Limit(limit).Archived(archived).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsApi.GetCmsV3DomainsGetPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCmsV3DomainsGetPage`: CollectionResponseWithTotalDomainForwardPaging
    fmt.Fprintf(os.Stdout, "Response from `DomainsApi.GetCmsV3DomainsGetPage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCmsV3DomainsGetPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createdAt** | **time.Time** | Only return domains created at this date. | 
 **createdAfter** | **time.Time** | Only return domains created after this date. | 
 **createdBefore** | **time.Time** | Only return domains created before this date. | 
 **updatedAt** | **time.Time** | Only return domains updated at this date. | 
 **updatedAfter** | **time.Time** | Only return domains updated after this date. | 
 **updatedBefore** | **time.Time** | Only return domains updated before this date. | 
 **sort** | **[]string** |  | 
 **after** | **string** | The paging cursor token of the last successfully read resource will be returned as the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results. | 
 **limit** | **int32** | Maximum number of results per page. | 
 **archived** | **bool** | Whether to return only results that have been archived. | 

### Return type

[**CollectionResponseWithTotalDomainForwardPaging**](CollectionResponseWithTotalDomainForwardPaging.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy), [oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

