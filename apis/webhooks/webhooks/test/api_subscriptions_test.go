/*
Webhooks API

Testing SubscriptionsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package webhooks

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/pyscht/hubspot-go-api/apis/webhooks/webhooks"
)

func Test_webhooks_SubscriptionsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SubscriptionsApiService DeleteWebhooksV3AppIdSubscriptionsSubscriptionIdArchive", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var subscriptionId int32
		var appId int32

		httpRes, err := apiClient.SubscriptionsApi.DeleteWebhooksV3AppIdSubscriptionsSubscriptionIdArchive(context.Background(), subscriptionId, appId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SubscriptionsApiService GetWebhooksV3AppIdSubscriptionsGetAll", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var appId int32

		resp, httpRes, err := apiClient.SubscriptionsApi.GetWebhooksV3AppIdSubscriptionsGetAll(context.Background(), appId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SubscriptionsApiService GetWebhooksV3AppIdSubscriptionsSubscriptionIdGetById", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var subscriptionId int32
		var appId int32

		resp, httpRes, err := apiClient.SubscriptionsApi.GetWebhooksV3AppIdSubscriptionsSubscriptionIdGetById(context.Background(), subscriptionId, appId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SubscriptionsApiService PatchWebhooksV3AppIdSubscriptionsSubscriptionIdUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var subscriptionId int32
		var appId int32

		resp, httpRes, err := apiClient.SubscriptionsApi.PatchWebhooksV3AppIdSubscriptionsSubscriptionIdUpdate(context.Background(), subscriptionId, appId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SubscriptionsApiService PostWebhooksV3AppIdSubscriptionsBatchUpdateUpdateBatch", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var appId int32

		resp, httpRes, err := apiClient.SubscriptionsApi.PostWebhooksV3AppIdSubscriptionsBatchUpdateUpdateBatch(context.Background(), appId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SubscriptionsApiService PostWebhooksV3AppIdSubscriptionsCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var appId int32

		resp, httpRes, err := apiClient.SubscriptionsApi.PostWebhooksV3AppIdSubscriptionsCreate(context.Background(), appId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
