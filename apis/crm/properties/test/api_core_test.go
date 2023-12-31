/*
Properties

Testing CoreApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package properties

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/pyscht/hubspot-go-api/apis/crm/properties"
)

func Test_properties_CoreApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test CoreApiService DeleteCrmV3PropertiesObjectTypePropertyNameArchive", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var objectType string
		var propertyName string

		httpRes, err := apiClient.CoreApi.DeleteCrmV3PropertiesObjectTypePropertyNameArchive(context.Background(), objectType, propertyName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CoreApiService GetCrmV3PropertiesObjectTypeGetAll", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var objectType string

		resp, httpRes, err := apiClient.CoreApi.GetCrmV3PropertiesObjectTypeGetAll(context.Background(), objectType).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CoreApiService GetCrmV3PropertiesObjectTypePropertyNameGetByName", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var objectType string
		var propertyName string

		resp, httpRes, err := apiClient.CoreApi.GetCrmV3PropertiesObjectTypePropertyNameGetByName(context.Background(), objectType, propertyName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CoreApiService PatchCrmV3PropertiesObjectTypePropertyNameUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var objectType string
		var propertyName string

		resp, httpRes, err := apiClient.CoreApi.PatchCrmV3PropertiesObjectTypePropertyNameUpdate(context.Background(), objectType, propertyName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CoreApiService PostCrmV3PropertiesObjectTypeCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var objectType string

		resp, httpRes, err := apiClient.CoreApi.PostCrmV3PropertiesObjectTypeCreate(context.Background(), objectType).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
