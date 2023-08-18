/*
Line Items

Testing PublicObjectApiService

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

func Test_line_items_PublicObjectApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PublicObjectApiService PostCrmV3ObjectsLineItemsMergeMerge", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.PublicObjectApi.PostCrmV3ObjectsLineItemsMergeMerge(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}