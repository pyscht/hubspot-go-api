/*
Line Items

Testing BatchApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package line_items

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/pyscht/hubspot-go-api/apis/crm/line_items"
)

func Test_line_items_BatchApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test BatchApiService PostCrmV3ObjectsLineItemsBatchArchiveArchive", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.BatchApi.PostCrmV3ObjectsLineItemsBatchArchiveArchive(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BatchApiService PostCrmV3ObjectsLineItemsBatchCreateCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.BatchApi.PostCrmV3ObjectsLineItemsBatchCreateCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BatchApiService PostCrmV3ObjectsLineItemsBatchReadRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.BatchApi.PostCrmV3ObjectsLineItemsBatchReadRead(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BatchApiService PostCrmV3ObjectsLineItemsBatchUpdateUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.BatchApi.PostCrmV3ObjectsLineItemsBatchUpdateUpdate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
