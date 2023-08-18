/*
CrmPublicAssociationsServiceV4

Testing BatchApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package crm_associations

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/pyscht/hubspot-go-api/apis/crm/crm_associations"
)

func Test_crm_associations_BatchApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test BatchApiService PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchArchiveArchive", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fromObjectType string
		var toObjectType string

		httpRes, err := apiClient.BatchApi.PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchArchiveArchive(context.Background(), fromObjectType, toObjectType).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BatchApiService PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchAssociateDefaultCreateDefault", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fromObjectType string
		var toObjectType string

		resp, httpRes, err := apiClient.BatchApi.PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchAssociateDefaultCreateDefault(context.Background(), fromObjectType, toObjectType).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BatchApiService PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchCreateCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fromObjectType string
		var toObjectType string

		resp, httpRes, err := apiClient.BatchApi.PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchCreateCreate(context.Background(), fromObjectType, toObjectType).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BatchApiService PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchLabelsArchiveArchiveLabels", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fromObjectType string
		var toObjectType string

		httpRes, err := apiClient.BatchApi.PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchLabelsArchiveArchiveLabels(context.Background(), fromObjectType, toObjectType).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BatchApiService PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchReadGetPage", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fromObjectType string
		var toObjectType string

		resp, httpRes, err := apiClient.BatchApi.PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchReadGetPage(context.Background(), fromObjectType, toObjectType).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
