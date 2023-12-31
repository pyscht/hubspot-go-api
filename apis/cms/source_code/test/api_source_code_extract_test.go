/*
CMS Source Code

Testing SourceCodeExtractApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package source_code

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/pyscht/hubspot-go-api/apis/cms/source_code"
)

func Test_source_code_SourceCodeExtractApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SourceCodeExtractApiService GetCmsV3SourceCodeExtractAsyncTasksTaskIdStatusGetAsyncStatus", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var taskId int32

		resp, httpRes, err := apiClient.SourceCodeExtractApi.GetCmsV3SourceCodeExtractAsyncTasksTaskIdStatusGetAsyncStatus(context.Background(), taskId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SourceCodeExtractApiService PostCmsV3SourceCodeExtractAsyncDoAsync", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SourceCodeExtractApi.PostCmsV3SourceCodeExtractAsyncDoAsync(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
