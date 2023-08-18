# Go API client for source_code

Endpoints for interacting with files in the CMS Developer File System. These files include HTML templates, CSS, JS, modules, and other assets which are used to create CMS content.

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
import source_code "github.com/pyscht/hubspot-go-api/apis/cms/source_code"
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
ctx := context.WithValue(context.Background(), source_code.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), source_code.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), source_code.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), source_code.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.hubapi.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ContentApi* | [**DeleteCmsV3SourceCodeEnvironmentContentPathArchive**](docs/ContentApi.md#deletecmsv3sourcecodeenvironmentcontentpatharchive) | **Delete** /cms/v3/source-code/{environment}/content/{path} | Delete a file
*ContentApi* | [**GetCmsV3SourceCodeEnvironmentContentPathGet**](docs/ContentApi.md#getcmsv3sourcecodeenvironmentcontentpathget) | **Get** /cms/v3/source-code/{environment}/content/{path} | Download a file
*ContentApi* | [**PostCmsV3SourceCodeEnvironmentContentPathCreate**](docs/ContentApi.md#postcmsv3sourcecodeenvironmentcontentpathcreate) | **Post** /cms/v3/source-code/{environment}/content/{path} | Create a file
*ContentApi* | [**PutCmsV3SourceCodeEnvironmentContentPathReplace**](docs/ContentApi.md#putcmsv3sourcecodeenvironmentcontentpathreplace) | **Put** /cms/v3/source-code/{environment}/content/{path} | Create or update a file
*ExtractApi* | [**PostCmsV3SourceCodeExtractPathExtractByPath**](docs/ExtractApi.md#postcmsv3sourcecodeextractpathextractbypath) | **Post** /cms/v3/source-code/extract/{path} | Extracts a zip file
*MetadataApi* | [**GetCmsV3SourceCodeEnvironmentMetadataPathGet**](docs/MetadataApi.md#getcmsv3sourcecodeenvironmentmetadatapathget) | **Get** /cms/v3/source-code/{environment}/metadata/{path} | Get the metadata for a file
*SourceCodeExtractApi* | [**GetCmsV3SourceCodeExtractAsyncTasksTaskIdStatusGetAsyncStatus**](docs/SourceCodeExtractApi.md#getcmsv3sourcecodeextractasynctaskstaskidstatusgetasyncstatus) | **Get** /cms/v3/source-code/extract/async/tasks/{taskId}/status | 
*SourceCodeExtractApi* | [**PostCmsV3SourceCodeExtractAsyncDoAsync**](docs/SourceCodeExtractApi.md#postcmsv3sourcecodeextractasyncdoasync) | **Post** /cms/v3/source-code/extract/async | 
*ValidationApi* | [**PostCmsV3SourceCodeEnvironmentValidatePathDoValidate**](docs/ValidationApi.md#postcmsv3sourcecodeenvironmentvalidatepathdovalidate) | **Post** /cms/v3/source-code/{environment}/validate/{path} | Validate the contents of a file


## Documentation For Models

 - [ActionResponse](docs/ActionResponse.md)
 - [AssetFileMetadata](docs/AssetFileMetadata.md)
 - [Error](docs/Error.md)
 - [ErrorDetail](docs/ErrorDetail.md)
 - [FileExtractRequest](docs/FileExtractRequest.md)
 - [TaskLocator](docs/TaskLocator.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### oauth2_legacy


- **Type**: OAuth
- **Flow**: accessCode
- **Authorization URL**: https://app.hubspot.com/oauth/authorize
- **Scopes**: 
 - **content**: Read from and write to my Content

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


