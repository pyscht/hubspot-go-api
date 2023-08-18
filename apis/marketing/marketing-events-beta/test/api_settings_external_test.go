/*
Marketing Events Extension

Testing SettingsExternalApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package marketing-events-beta

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/pyscht/hubspot-go-api/apis/marketing/marketing-events-beta"
)

func Test_marketing-events-beta_SettingsExternalApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SettingsExternalApiService GetMarketingV3MarketingEventsAppIdSettingsGetAll", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var appId int32

		resp, httpRes, err := apiClient.SettingsExternalApi.GetMarketingV3MarketingEventsAppIdSettingsGetAll(context.Background(), appId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SettingsExternalApiService PostMarketingV3MarketingEventsAppIdSettingsCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var appId int32

		resp, httpRes, err := apiClient.SettingsExternalApi.PostMarketingV3MarketingEventsAppIdSettingsCreate(context.Background(), appId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
