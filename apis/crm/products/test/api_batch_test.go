/*
Products

Testing BatchApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package products

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/pyscht/hubspot-go-api/apis/crm/products"
)

func Test_products_BatchApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test BatchApiService PostCrmV3ObjectsProductsBatchArchiveArchive", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.BatchApi.PostCrmV3ObjectsProductsBatchArchiveArchive(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BatchApiService PostCrmV3ObjectsProductsBatchCreateCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.BatchApi.PostCrmV3ObjectsProductsBatchCreateCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BatchApiService PostCrmV3ObjectsProductsBatchReadRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.BatchApi.PostCrmV3ObjectsProductsBatchReadRead(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BatchApiService PostCrmV3ObjectsProductsBatchUpdateUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.BatchApi.PostCrmV3ObjectsProductsBatchUpdateUpdate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}