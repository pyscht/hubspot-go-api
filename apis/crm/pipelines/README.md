# Go API client for pipelines

Pipelines represent distinct stages in a workflow, like closing a deal or servicing a support ticket. These endpoints provide access to read and modify pipelines in HubSpot. Pipelines support `deals` and `tickets` object types.

## Pipeline ID validation

When calling endpoints that take pipelineId as a parameter, that ID must correspond to an existing, un-archived pipeline. Otherwise the request will fail with a `404 Not Found` response.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: v3
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import pipelines "github.com/pyscht/hubspot-go-api/apis/crm/pipelines"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), pipelines.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), pipelines.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), pipelines.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), pipelines.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.hubapi.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*PipelineAuditsApi* | [**GetCrmV3PipelinesObjectTypePipelineIdAuditGetAudit**](docs/PipelineAuditsApi.md#getcrmv3pipelinesobjecttypepipelineidauditgetaudit) | **Get** /crm/v3/pipelines/{objectType}/{pipelineId}/audit | Return an audit of all changes to the pipeline
*PipelineStageAuditsApi* | [**GetCrmV3PipelinesObjectTypePipelineIdStagesStageIdAuditGetAudit**](docs/PipelineStageAuditsApi.md#getcrmv3pipelinesobjecttypepipelineidstagesstageidauditgetaudit) | **Get** /crm/v3/pipelines/{objectType}/{pipelineId}/stages/{stageId}/audit | Return an audit of all changes to the pipeline stage
*PipelineStagesApi* | [**DeleteCrmV3PipelinesObjectTypePipelineIdStagesStageIdArchive**](docs/PipelineStagesApi.md#deletecrmv3pipelinesobjecttypepipelineidstagesstageidarchive) | **Delete** /crm/v3/pipelines/{objectType}/{pipelineId}/stages/{stageId} | Delete a pipeline stage
*PipelineStagesApi* | [**GetCrmV3PipelinesObjectTypePipelineIdStagesGetAll**](docs/PipelineStagesApi.md#getcrmv3pipelinesobjecttypepipelineidstagesgetall) | **Get** /crm/v3/pipelines/{objectType}/{pipelineId}/stages | Return all stages of a pipeline
*PipelineStagesApi* | [**GetCrmV3PipelinesObjectTypePipelineIdStagesStageIdGetById**](docs/PipelineStagesApi.md#getcrmv3pipelinesobjecttypepipelineidstagesstageidgetbyid) | **Get** /crm/v3/pipelines/{objectType}/{pipelineId}/stages/{stageId} | Return a pipeline stage by ID
*PipelineStagesApi* | [**PatchCrmV3PipelinesObjectTypePipelineIdStagesStageIdUpdate**](docs/PipelineStagesApi.md#patchcrmv3pipelinesobjecttypepipelineidstagesstageidupdate) | **Patch** /crm/v3/pipelines/{objectType}/{pipelineId}/stages/{stageId} | Update a pipeline stage
*PipelineStagesApi* | [**PostCrmV3PipelinesObjectTypePipelineIdStagesCreate**](docs/PipelineStagesApi.md#postcrmv3pipelinesobjecttypepipelineidstagescreate) | **Post** /crm/v3/pipelines/{objectType}/{pipelineId}/stages | Create a pipeline stage
*PipelineStagesApi* | [**PutCrmV3PipelinesObjectTypePipelineIdStagesStageIdReplace**](docs/PipelineStagesApi.md#putcrmv3pipelinesobjecttypepipelineidstagesstageidreplace) | **Put** /crm/v3/pipelines/{objectType}/{pipelineId}/stages/{stageId} | Replace a pipeline stage
*PipelinesApi* | [**DeleteCrmV3PipelinesObjectTypePipelineIdArchive**](docs/PipelinesApi.md#deletecrmv3pipelinesobjecttypepipelineidarchive) | **Delete** /crm/v3/pipelines/{objectType}/{pipelineId} | Delete a pipeline
*PipelinesApi* | [**GetCrmV3PipelinesObjectTypeGetAll**](docs/PipelinesApi.md#getcrmv3pipelinesobjecttypegetall) | **Get** /crm/v3/pipelines/{objectType} | Retrieve all pipelines
*PipelinesApi* | [**GetCrmV3PipelinesObjectTypePipelineIdGetById**](docs/PipelinesApi.md#getcrmv3pipelinesobjecttypepipelineidgetbyid) | **Get** /crm/v3/pipelines/{objectType}/{pipelineId} | Return a pipeline by ID
*PipelinesApi* | [**PatchCrmV3PipelinesObjectTypePipelineIdUpdate**](docs/PipelinesApi.md#patchcrmv3pipelinesobjecttypepipelineidupdate) | **Patch** /crm/v3/pipelines/{objectType}/{pipelineId} | Update a pipeline
*PipelinesApi* | [**PostCrmV3PipelinesObjectTypeCreate**](docs/PipelinesApi.md#postcrmv3pipelinesobjecttypecreate) | **Post** /crm/v3/pipelines/{objectType} | Create a pipeline
*PipelinesApi* | [**PutCrmV3PipelinesObjectTypePipelineIdReplace**](docs/PipelinesApi.md#putcrmv3pipelinesobjecttypepipelineidreplace) | **Put** /crm/v3/pipelines/{objectType}/{pipelineId} | Replace a pipeline


## Documentation For Models

 - [CollectionResponsePipelineNoPaging](docs/CollectionResponsePipelineNoPaging.md)
 - [CollectionResponsePipelineStageNoPaging](docs/CollectionResponsePipelineStageNoPaging.md)
 - [CollectionResponsePublicAuditInfoNoPaging](docs/CollectionResponsePublicAuditInfoNoPaging.md)
 - [Error](docs/Error.md)
 - [ErrorDetail](docs/ErrorDetail.md)
 - [Pipeline](docs/Pipeline.md)
 - [PipelineInput](docs/PipelineInput.md)
 - [PipelinePatchInput](docs/PipelinePatchInput.md)
 - [PipelineStage](docs/PipelineStage.md)
 - [PipelineStageInput](docs/PipelineStageInput.md)
 - [PipelineStagePatchInput](docs/PipelineStagePatchInput.md)
 - [PublicAuditInfo](docs/PublicAuditInfo.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### oauth2_legacy


- **Type**: OAuth
- **Flow**: accessCode
- **Authorization URL**: https://app.hubspot.com/oauth/authorize
- **Scopes**: 
 - **crm.schemas.custom.read**: View custom object definitions
 - **tickets**: Read and write tickets
 - **crm.objects.goals.read**: Read goals
 - **media_bridge.read**: Read media and media events
 - **e-commerce**: e-commerce
 - **timeline**: Create timeline events

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```golang
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, sw.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```

### oauth2


- **Type**: OAuth
- **Flow**: accessCode
- **Authorization URL**: https://app.hubspot.com/oauth/authorize
- **Scopes**: 
 - **crm.schemas.companies.read**:  
 - **crm.objects.deals.read**:  
 - **crm.schemas.line_items.read**: Line Items schemas
 - **crm.objects.deals.write**:  
 - **crm.schemas.deals.read**:  
 - **crm.objects.contacts.read**:  
 - **crm.schemas.quotes.read**: Quotes schemas
 - **crm.objects.contacts.write**:  
 - **crm.schemas.contacts.read**:  
 - **crm.objects.companies.write**:  
 - **crm.objects.companies.read**:  
 - **crm.schemas.companies.write**:  
 - **crm.schemas.deals.write**:  
 - **crm.schemas.contacts.write**:  

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```golang
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, sw.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```

### private_apps_legacy

- **Type**: API key
- **API key parameter name**: private-app-legacy
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: private-app-legacy and passed in as the auth context for each request.

### private_apps

- **Type**: API key
- **API key parameter name**: private-app
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: private-app and passed in as the auth context for each request.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author


