/*
Marketing Events Extension

Testing AttendanceSubscriberStateChangesApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package marketing_events_beta

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/pyscht/hubspot-go-api/apis/marketing/marketing_events_beta"
)

func Test_marketing_events_beta_AttendanceSubscriberStateChangesApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AttendanceSubscriberStateChangesApiService PostMarketingV3MarketingEventsAttendanceExternalEventIdSubscriberStateCreateCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var externalEventId string
		var subscriberState string

		resp, httpRes, err := apiClient.AttendanceSubscriberStateChangesApi.PostMarketingV3MarketingEventsAttendanceExternalEventIdSubscriberStateCreateCreate(context.Background(), externalEventId, subscriberState).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AttendanceSubscriberStateChangesApiService PostMarketingV3MarketingEventsAttendanceExternalEventIdSubscriberStateEmailCreateCreateByEmail", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var externalEventId string
		var subscriberState string

		resp, httpRes, err := apiClient.AttendanceSubscriberStateChangesApi.PostMarketingV3MarketingEventsAttendanceExternalEventIdSubscriberStateEmailCreateCreateByEmail(context.Background(), externalEventId, subscriberState).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}