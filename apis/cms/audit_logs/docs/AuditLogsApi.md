# \AuditLogsApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCmsV3AuditLogsGetPage**](AuditLogsApi.md#GetCmsV3AuditLogsGetPage) | **Get** /cms/v3/audit-logs/ | Query audit logs



## GetCmsV3AuditLogsGetPage

> CollectionResponsePublicAuditLog GetCmsV3AuditLogsGetPage(ctx).ObjectId(objectId).UserId(userId).After(after).Before(before).Sort(sort).EventType(eventType).Limit(limit).ObjectType(objectType).Execute()

Query audit logs



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/pyscht/hubspot-go-api/apis/cms/audit_logs"
)

func main() {
    objectId := []string{"Inner_example"} // []string | Comma separated list of object ids to filter by. (optional)
    userId := []string{"Inner_example"} // []string | Comma separated list of user ids to filter by. (optional)
    after := "after_example" // string | Timestamp after which audit logs will be returned (optional)
    before := "before_example" // string | Timestamp before which audit logs will be returned (optional)
    sort := []string{"Inner_example"} // []string | The sort direction for the audit logs. (Can only sort by timestamp). (optional)
    eventType := []string{"Inner_example"} // []string | Comma separated list of event types to filter by (CREATED, UPDATED, PUBLISHED, DELETED, UNPUBLISHED). (optional)
    limit := int32(56) // int32 | The number of logs to return. (optional)
    objectType := []string{"Inner_example"} // []string | Comma separated list of object types to filter by (BLOG, LANDING_PAGE, DOMAIN, HUBDB_TABLE etc.) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuditLogsApi.GetCmsV3AuditLogsGetPage(context.Background()).ObjectId(objectId).UserId(userId).After(after).Before(before).Sort(sort).EventType(eventType).Limit(limit).ObjectType(objectType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuditLogsApi.GetCmsV3AuditLogsGetPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCmsV3AuditLogsGetPage`: CollectionResponsePublicAuditLog
    fmt.Fprintf(os.Stdout, "Response from `AuditLogsApi.GetCmsV3AuditLogsGetPage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCmsV3AuditLogsGetPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **objectId** | **[]string** | Comma separated list of object ids to filter by. | 
 **userId** | **[]string** | Comma separated list of user ids to filter by. | 
 **after** | **string** | Timestamp after which audit logs will be returned | 
 **before** | **string** | Timestamp before which audit logs will be returned | 
 **sort** | **[]string** | The sort direction for the audit logs. (Can only sort by timestamp). | 
 **eventType** | **[]string** | Comma separated list of event types to filter by (CREATED, UPDATED, PUBLISHED, DELETED, UNPUBLISHED). | 
 **limit** | **int32** | The number of logs to return. | 
 **objectType** | **[]string** | Comma separated list of object types to filter by (BLOG, LANDING_PAGE, DOMAIN, HUBDB_TABLE etc.) | 

### Return type

[**CollectionResponsePublicAuditLog**](CollectionResponsePublicAuditLog.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

