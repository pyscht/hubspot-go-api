/*
Visitor Identification

Testing GenerateApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package visitor_identification

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/pyscht/hubspot-go-api/apis/conversations/visitor_identification"
)

func Test_visitor_identification_GenerateApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test GenerateApiService PostVisitorIdentificationV3TokensCreateGenerateToken", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.GenerateApi.PostVisitorIdentificationV3TokensCreateGenerateToken(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}