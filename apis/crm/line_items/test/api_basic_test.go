/*
Line Items

Testing BasicApiService

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

func Test_line_items_BasicApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test BasicApiService DeleteCrmV3ObjectsLineItemsLineItemIdArchive", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var lineItemId string

		httpRes, err := apiClient.BasicApi.DeleteCrmV3ObjectsLineItemsLineItemIdArchive(context.Background(), lineItemId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BasicApiService GetCrmV3ObjectsLineItemsGetPage", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.BasicApi.GetCrmV3ObjectsLineItemsGetPage(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BasicApiService GetCrmV3ObjectsLineItemsLineItemIdGetById", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var lineItemId string

		resp, httpRes, err := apiClient.BasicApi.GetCrmV3ObjectsLineItemsLineItemIdGetById(context.Background(), lineItemId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BasicApiService PatchCrmV3ObjectsLineItemsLineItemIdUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var lineItemId string

		resp, httpRes, err := apiClient.BasicApi.PatchCrmV3ObjectsLineItemsLineItemIdUpdate(context.Background(), lineItemId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BasicApiService PostCrmV3ObjectsLineItemsCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.BasicApi.PostCrmV3ObjectsLineItemsCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
