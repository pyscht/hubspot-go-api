# \PublicObjectApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostCrmV3ObjectsFeedbackSubmissionsMergeMerge**](PublicObjectApi.md#PostCrmV3ObjectsFeedbackSubmissionsMergeMerge) | **Post** /crm/v3/objects/feedback_submissions/merge | Merge two feedback submissions with same type



## PostCrmV3ObjectsFeedbackSubmissionsMergeMerge

> SimplePublicObject PostCrmV3ObjectsFeedbackSubmissionsMergeMerge(ctx).PublicMergeInput(publicMergeInput).Execute()

Merge two feedback submissions with same type

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/pyscht/hubspot-go-api/apis/crm/feedback_submissions"
)

func main() {
    publicMergeInput := *openapiclient.NewPublicMergeInput("PrimaryObjectId_example", "ObjectIdToMerge_example") // PublicMergeInput | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PublicObjectApi.PostCrmV3ObjectsFeedbackSubmissionsMergeMerge(context.Background()).PublicMergeInput(publicMergeInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicObjectApi.PostCrmV3ObjectsFeedbackSubmissionsMergeMerge``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCrmV3ObjectsFeedbackSubmissionsMergeMerge`: SimplePublicObject
    fmt.Fprintf(os.Stdout, "Response from `PublicObjectApi.PostCrmV3ObjectsFeedbackSubmissionsMergeMerge`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmV3ObjectsFeedbackSubmissionsMergeMergeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **publicMergeInput** | [**PublicMergeInput**](PublicMergeInput.md) |  | 

### Return type

[**SimplePublicObject**](SimplePublicObject.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

