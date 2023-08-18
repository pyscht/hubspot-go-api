/*
CRM Objects

Testing BatchApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package objects

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/pyscht/hubspot-go-api/apis/crm/objects"
)

func Test_objects_BatchApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test BatchApiService PostCrmV3ObjectsObjectTypeBatchArchiveArchive", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var objectType string

		httpRes, err := apiClient.BatchApi.PostCrmV3ObjectsObjectTypeBatchArchiveArchive(context.Background(), objectType).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BatchApiService PostCrmV3ObjectsObjectTypeBatchCreateCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var objectType string

		resp, httpRes, err := apiClient.BatchApi.PostCrmV3ObjectsObjectTypeBatchCreateCreate(context.Background(), objectType).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BatchApiService PostCrmV3ObjectsObjectTypeBatchReadRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var objectType string

		resp, httpRes, err := apiClient.BatchApi.PostCrmV3ObjectsObjectTypeBatchReadRead(context.Background(), objectType).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BatchApiService PostCrmV3ObjectsObjectTypeBatchUpdateUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var objectType string

		resp, httpRes, err := apiClient.BatchApi.PostCrmV3ObjectsObjectTypeBatchUpdateUpdate(context.Background(), objectType).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
