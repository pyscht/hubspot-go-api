/*
CRM Objects

Testing AssociationsApiService

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

func Test_objects_AssociationsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AssociationsApiService DeleteCrmV3ObjectsObjectTypeObjectIdAssociationsToObjectTypeToObjectIdAssociationTypeArchive", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var objectType string
		var objectId string
		var toObjectType string
		var toObjectId string
		var associationType string

		httpRes, err := apiClient.AssociationsApi.DeleteCrmV3ObjectsObjectTypeObjectIdAssociationsToObjectTypeToObjectIdAssociationTypeArchive(context.Background(), objectType, objectId, toObjectType, toObjectId, associationType).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AssociationsApiService GetCrmV3ObjectsObjectTypeObjectIdAssociationsToObjectTypeGetAll", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var objectType string
		var objectId string
		var toObjectType string

		resp, httpRes, err := apiClient.AssociationsApi.GetCrmV3ObjectsObjectTypeObjectIdAssociationsToObjectTypeGetAll(context.Background(), objectType, objectId, toObjectType).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AssociationsApiService PutCrmV3ObjectsObjectTypeObjectIdAssociationsToObjectTypeToObjectIdAssociationTypeCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var objectType string
		var objectId string
		var toObjectType string
		var toObjectId string
		var associationType string

		resp, httpRes, err := apiClient.AssociationsApi.PutCrmV3ObjectsObjectTypeObjectIdAssociationsToObjectTypeToObjectIdAssociationTypeCreate(context.Background(), objectType, objectId, toObjectType, toObjectId, associationType).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}