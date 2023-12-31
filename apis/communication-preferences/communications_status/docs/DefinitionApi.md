# \DefinitionApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCommunicationPreferencesV3DefinitionsGetPage**](DefinitionApi.md#GetCommunicationPreferencesV3DefinitionsGetPage) | **Get** /communication-preferences/v3/definitions | Get subscription definitions



## GetCommunicationPreferencesV3DefinitionsGetPage

> SubscriptionDefinitionsResponse GetCommunicationPreferencesV3DefinitionsGetPage(ctx).Execute()

Get subscription definitions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/pyscht/hubspot-go-api/apis/communication-preferences/communications_status"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefinitionApi.GetCommunicationPreferencesV3DefinitionsGetPage(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefinitionApi.GetCommunicationPreferencesV3DefinitionsGetPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCommunicationPreferencesV3DefinitionsGetPage`: SubscriptionDefinitionsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefinitionApi.GetCommunicationPreferencesV3DefinitionsGetPage`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommunicationPreferencesV3DefinitionsGetPageRequest struct via the builder pattern


### Return type

[**SubscriptionDefinitionsResponse**](SubscriptionDefinitionsResponse.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

