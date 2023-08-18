/*
Domains

Testing DomainsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package domains

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/pyscht/hubspot-go-api/apis/cms/domains"
)

func Test_domains_DomainsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DomainsApiService GetCmsV3DomainsDomainIdGetById", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domainId string

		resp, httpRes, err := apiClient.DomainsApi.GetCmsV3DomainsDomainIdGetById(context.Background(), domainId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DomainsApiService GetCmsV3DomainsGetPage", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DomainsApi.GetCmsV3DomainsGetPage(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}