# \BatchApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostCrmV3ObjectsQuotesBatchArchiveArchive**](BatchApi.md#PostCrmV3ObjectsQuotesBatchArchiveArchive) | **Post** /crm/v3/objects/quotes/batch/archive | Archive a batch of quotes by ID
[**PostCrmV3ObjectsQuotesBatchCreateCreate**](BatchApi.md#PostCrmV3ObjectsQuotesBatchCreateCreate) | **Post** /crm/v3/objects/quotes/batch/create | Create a batch of quotes
[**PostCrmV3ObjectsQuotesBatchReadRead**](BatchApi.md#PostCrmV3ObjectsQuotesBatchReadRead) | **Post** /crm/v3/objects/quotes/batch/read | Read a batch of quotes by internal ID, or unique property values
[**PostCrmV3ObjectsQuotesBatchUpdateUpdate**](BatchApi.md#PostCrmV3ObjectsQuotesBatchUpdateUpdate) | **Post** /crm/v3/objects/quotes/batch/update | Update a batch of quotes



## PostCrmV3ObjectsQuotesBatchArchiveArchive

> PostCrmV3ObjectsQuotesBatchArchiveArchive(ctx).BatchInputSimplePublicObjectId(batchInputSimplePublicObjectId).Execute()

Archive a batch of quotes by ID

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/pyscht/hubspot-go-api/apis/crm/quotes"
)

func main() {
    batchInputSimplePublicObjectId := *openapiclient.NewBatchInputSimplePublicObjectId([]openapiclient.SimplePublicObjectId{*openapiclient.NewSimplePublicObjectId("Id_example")}) // BatchInputSimplePublicObjectId | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.BatchApi.PostCrmV3ObjectsQuotesBatchArchiveArchive(context.Background()).BatchInputSimplePublicObjectId(batchInputSimplePublicObjectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BatchApi.PostCrmV3ObjectsQuotesBatchArchiveArchive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmV3ObjectsQuotesBatchArchiveArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputSimplePublicObjectId** | [**BatchInputSimplePublicObjectId**](BatchInputSimplePublicObjectId.md) |  | 

### Return type

 (empty response body)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [oauth2](../README.md#oauth2), [private_apps_legacy](../README.md#private_apps_legacy), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCrmV3ObjectsQuotesBatchCreateCreate

> BatchResponseSimplePublicObject PostCrmV3ObjectsQuotesBatchCreateCreate(ctx).BatchInputSimplePublicObjectInputForCreate(batchInputSimplePublicObjectInputForCreate).Execute()

Create a batch of quotes

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/pyscht/hubspot-go-api/apis/crm/quotes"
)

func main() {
    batchInputSimplePublicObjectInputForCreate := *openapiclient.NewBatchInputSimplePublicObjectInputForCreate([]openapiclient.SimplePublicObjectInputForCreate{*openapiclient.NewSimplePublicObjectInputForCreate(map[string]string{"key": "Inner_example"}, []openapiclient.PublicAssociationsForObject{*openapiclient.NewPublicAssociationsForObject(*openapiclient.NewPublicObjectId("Id_example"), []openapiclient.AssociationSpec{*openapiclient.NewAssociationSpec("AssociationCategory_example", int32(123))})})}) // BatchInputSimplePublicObjectInputForCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BatchApi.PostCrmV3ObjectsQuotesBatchCreateCreate(context.Background()).BatchInputSimplePublicObjectInputForCreate(batchInputSimplePublicObjectInputForCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BatchApi.PostCrmV3ObjectsQuotesBatchCreateCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCrmV3ObjectsQuotesBatchCreateCreate`: BatchResponseSimplePublicObject
    fmt.Fprintf(os.Stdout, "Response from `BatchApi.PostCrmV3ObjectsQuotesBatchCreateCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmV3ObjectsQuotesBatchCreateCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputSimplePublicObjectInputForCreate** | [**BatchInputSimplePublicObjectInputForCreate**](BatchInputSimplePublicObjectInputForCreate.md) |  | 

### Return type

[**BatchResponseSimplePublicObject**](BatchResponseSimplePublicObject.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [oauth2](../README.md#oauth2), [private_apps_legacy](../README.md#private_apps_legacy), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCrmV3ObjectsQuotesBatchReadRead

> BatchResponseSimplePublicObject PostCrmV3ObjectsQuotesBatchReadRead(ctx).BatchReadInputSimplePublicObjectId(batchReadInputSimplePublicObjectId).Archived(archived).Execute()

Read a batch of quotes by internal ID, or unique property values

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/pyscht/hubspot-go-api/apis/crm/quotes"
)

func main() {
    batchReadInputSimplePublicObjectId := *openapiclient.NewBatchReadInputSimplePublicObjectId([]string{"Properties_example"}, []string{"PropertiesWithHistory_example"}, []openapiclient.SimplePublicObjectId{*openapiclient.NewSimplePublicObjectId("Id_example")}) // BatchReadInputSimplePublicObjectId | 
    archived := true // bool | Whether to return only results that have been archived. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BatchApi.PostCrmV3ObjectsQuotesBatchReadRead(context.Background()).BatchReadInputSimplePublicObjectId(batchReadInputSimplePublicObjectId).Archived(archived).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BatchApi.PostCrmV3ObjectsQuotesBatchReadRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCrmV3ObjectsQuotesBatchReadRead`: BatchResponseSimplePublicObject
    fmt.Fprintf(os.Stdout, "Response from `BatchApi.PostCrmV3ObjectsQuotesBatchReadRead`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmV3ObjectsQuotesBatchReadReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchReadInputSimplePublicObjectId** | [**BatchReadInputSimplePublicObjectId**](BatchReadInputSimplePublicObjectId.md) |  | 
 **archived** | **bool** | Whether to return only results that have been archived. | [default to false]

### Return type

[**BatchResponseSimplePublicObject**](BatchResponseSimplePublicObject.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy), [oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCrmV3ObjectsQuotesBatchUpdateUpdate

> BatchResponseSimplePublicObject PostCrmV3ObjectsQuotesBatchUpdateUpdate(ctx).BatchInputSimplePublicObjectBatchInput(batchInputSimplePublicObjectBatchInput).Execute()

Update a batch of quotes

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/pyscht/hubspot-go-api/apis/crm/quotes"
)

func main() {
    batchInputSimplePublicObjectBatchInput := *openapiclient.NewBatchInputSimplePublicObjectBatchInput([]openapiclient.SimplePublicObjectBatchInput{*openapiclient.NewSimplePublicObjectBatchInput(map[string]string{"key": "Inner_example"}, "Id_example")}) // BatchInputSimplePublicObjectBatchInput | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BatchApi.PostCrmV3ObjectsQuotesBatchUpdateUpdate(context.Background()).BatchInputSimplePublicObjectBatchInput(batchInputSimplePublicObjectBatchInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BatchApi.PostCrmV3ObjectsQuotesBatchUpdateUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCrmV3ObjectsQuotesBatchUpdateUpdate`: BatchResponseSimplePublicObject
    fmt.Fprintf(os.Stdout, "Response from `BatchApi.PostCrmV3ObjectsQuotesBatchUpdateUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmV3ObjectsQuotesBatchUpdateUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputSimplePublicObjectBatchInput** | [**BatchInputSimplePublicObjectBatchInput**](BatchInputSimplePublicObjectBatchInput.md) |  | 

### Return type

[**BatchResponseSimplePublicObject**](BatchResponseSimplePublicObject.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [oauth2](../README.md#oauth2), [private_apps_legacy](../README.md#private_apps_legacy), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

