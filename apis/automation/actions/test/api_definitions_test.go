/*
Custom Workflow Actions

Testing DefinitionsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package actions

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/pyscht/hubspot-go-api/apis/automation/actions"
)

func Test_actions_DefinitionsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DefinitionsApiService DeleteAutomationV4ActionsAppIdDefinitionIdArchive", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var definitionId string
		var appId int32

		httpRes, err := apiClient.DefinitionsApi.DeleteAutomationV4ActionsAppIdDefinitionIdArchive(context.Background(), definitionId, appId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefinitionsApiService GetAutomationV4ActionsAppIdDefinitionIdGetById", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var definitionId string
		var appId int32

		resp, httpRes, err := apiClient.DefinitionsApi.GetAutomationV4ActionsAppIdDefinitionIdGetById(context.Background(), definitionId, appId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefinitionsApiService GetAutomationV4ActionsAppIdGetPage", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var appId int32

		resp, httpRes, err := apiClient.DefinitionsApi.GetAutomationV4ActionsAppIdGetPage(context.Background(), appId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefinitionsApiService PatchAutomationV4ActionsAppIdDefinitionIdUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var definitionId string
		var appId int32

		resp, httpRes, err := apiClient.DefinitionsApi.PatchAutomationV4ActionsAppIdDefinitionIdUpdate(context.Background(), definitionId, appId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefinitionsApiService PostAutomationV4ActionsAppIdCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var appId int32

		resp, httpRes, err := apiClient.DefinitionsApi.PostAutomationV4ActionsAppIdCreate(context.Background(), appId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
