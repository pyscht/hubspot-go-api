/*
Transactional Email

Testing SingleSendApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package transactional

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/pyscht/hubspot-go-api/apis/marketing/transactional"
)

func Test_transactional_SingleSendApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SingleSendApiService PostMarketingV3TransactionalSingleEmailSendSendEmail", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SingleSendApi.PostMarketingV3TransactionalSingleEmailSendSendEmail(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
