/*
Feedback Submissions

Testing SearchApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package feedback_submissions

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/pyscht/hubspot-go-api/apis/crm/feedback_submissions"
)

func Test_feedback_submissions_SearchApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SearchApiService PostCrmV3ObjectsFeedbackSubmissionsSearchDoSearch", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SearchApi.PostCrmV3ObjectsFeedbackSubmissionsSearchDoSearch(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
