# \PipelineStageAuditsApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCrmV3PipelinesObjectTypePipelineIdStagesStageIdAuditGetAudit**](PipelineStageAuditsApi.md#GetCrmV3PipelinesObjectTypePipelineIdStagesStageIdAuditGetAudit) | **Get** /crm/v3/pipelines/{objectType}/{pipelineId}/stages/{stageId}/audit | Return an audit of all changes to the pipeline stage



## GetCrmV3PipelinesObjectTypePipelineIdStagesStageIdAuditGetAudit

> CollectionResponsePublicAuditInfoNoPaging GetCrmV3PipelinesObjectTypePipelineIdStagesStageIdAuditGetAudit(ctx, objectType, stageId).Execute()

Return an audit of all changes to the pipeline stage



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/pyscht/hubspot-go-api/apis/crm/pipelines"
)

func main() {
    objectType := "objectType_example" // string | 
    stageId := "stageId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PipelineStageAuditsApi.GetCrmV3PipelinesObjectTypePipelineIdStagesStageIdAuditGetAudit(context.Background(), objectType, stageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelineStageAuditsApi.GetCrmV3PipelinesObjectTypePipelineIdStagesStageIdAuditGetAudit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCrmV3PipelinesObjectTypePipelineIdStagesStageIdAuditGetAudit`: CollectionResponsePublicAuditInfoNoPaging
    fmt.Fprintf(os.Stdout, "Response from `PipelineStageAuditsApi.GetCrmV3PipelinesObjectTypePipelineIdStagesStageIdAuditGetAudit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** |  | 
**stageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmV3PipelinesObjectTypePipelineIdStagesStageIdAuditGetAuditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CollectionResponsePublicAuditInfoNoPaging**](CollectionResponsePublicAuditInfoNoPaging.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [oauth2](../README.md#oauth2), [private_apps_legacy](../README.md#private_apps_legacy), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

