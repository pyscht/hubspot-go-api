/*
Video Conference Extension

Testing SettingsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package videoconferencing

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/pyscht/hubspot-go-api/apis/crm/videoconferencing"
)

func Test_videoconferencing_SettingsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SettingsApiService DeleteCrmV3ExtensionsVideoconferencingSettingsAppIdArchive", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var appId int32

		httpRes, err := apiClient.SettingsApi.DeleteCrmV3ExtensionsVideoconferencingSettingsAppIdArchive(context.Background(), appId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SettingsApiService GetCrmV3ExtensionsVideoconferencingSettingsAppIdGetById", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var appId int32

		resp, httpRes, err := apiClient.SettingsApi.GetCrmV3ExtensionsVideoconferencingSettingsAppIdGetById(context.Background(), appId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SettingsApiService PutCrmV3ExtensionsVideoconferencingSettingsAppIdReplace", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var appId int32

		resp, httpRes, err := apiClient.SettingsApi.PutCrmV3ExtensionsVideoconferencingSettingsAppIdReplace(context.Background(), appId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
