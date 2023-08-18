/*
Tickets

Testing SearchApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package tickets

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/pyscht/hubspot-go-api/apis/crm/tickets"
)

func Test_tickets_SearchApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SearchApiService PostCrmV3ObjectsTicketsSearchDoSearch", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SearchApi.PostCrmV3ObjectsTicketsSearchDoSearch(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}