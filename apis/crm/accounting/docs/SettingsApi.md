# \SettingsApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCrmV3ExtensionsAccountingSettingsAppIdGetById**](SettingsApi.md#GetCrmV3ExtensionsAccountingSettingsAppIdGetById) | **Get** /crm/v3/extensions/accounting/settings/{appId} | Get URL settings
[**PutCrmV3ExtensionsAccountingSettingsAppIdReplace**](SettingsApi.md#PutCrmV3ExtensionsAccountingSettingsAppIdReplace) | **Put** /crm/v3/extensions/accounting/settings/{appId} | Add/Update URL Settings



## GetCrmV3ExtensionsAccountingSettingsAppIdGetById

> AccountingAppSettings GetCrmV3ExtensionsAccountingSettingsAppIdGetById(ctx, appId).Execute()

Get URL settings



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/pyscht/hubspot-go-api/apis/crm/accounting"
)

func main() {
    appId := int32(56) // int32 | The ID of the accounting app. This is the identifier of the application created in your HubSpot developer portal.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SettingsApi.GetCrmV3ExtensionsAccountingSettingsAppIdGetById(context.Background(), appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.GetCrmV3ExtensionsAccountingSettingsAppIdGetById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCrmV3ExtensionsAccountingSettingsAppIdGetById`: AccountingAppSettings
    fmt.Fprintf(os.Stdout, "Response from `SettingsApi.GetCrmV3ExtensionsAccountingSettingsAppIdGetById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** | The ID of the accounting app. This is the identifier of the application created in your HubSpot developer portal. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmV3ExtensionsAccountingSettingsAppIdGetByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccountingAppSettings**](AccountingAppSettings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCrmV3ExtensionsAccountingSettingsAppIdReplace

> PutCrmV3ExtensionsAccountingSettingsAppIdReplace(ctx, appId).AccountingAppSettings(accountingAppSettings).Execute()

Add/Update URL Settings



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/pyscht/hubspot-go-api/apis/crm/accounting"
)

func main() {
    appId := int32(56) // int32 | The ID of the accounting app. This is the identifier of the application created in your HubSpot developer portal.
    accountingAppSettings := *openapiclient.NewAccountingAppSettings(int32(123), *openapiclient.NewAccountingAppUrls("GetInvoiceUrl_example", "SearchCustomerUrl_example", "GetInvoicePdfUrl_example", "CustomerUrlTemplate_example", "ProductUrlTemplate_example", "InvoiceUrlTemplate_example")) // AccountingAppSettings | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SettingsApi.PutCrmV3ExtensionsAccountingSettingsAppIdReplace(context.Background(), appId).AccountingAppSettings(accountingAppSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.PutCrmV3ExtensionsAccountingSettingsAppIdReplace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** | The ID of the accounting app. This is the identifier of the application created in your HubSpot developer portal. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCrmV3ExtensionsAccountingSettingsAppIdReplaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountingAppSettings** | [**AccountingAppSettings**](AccountingAppSettings.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

