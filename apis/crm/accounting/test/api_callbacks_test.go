/*
Accounting Extension

Testing CallbacksApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package accounting

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/pyscht/hubspot-go-api/apis/crm/accounting"
)

func Test_accounting_CallbacksApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test CallbacksApiService PostCrmV3ExtensionsAccountingCallbackCustomerCreateRequestIdCreateCustomer", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var requestId string

		httpRes, err := apiClient.CallbacksApi.PostCrmV3ExtensionsAccountingCallbackCustomerCreateRequestIdCreateCustomer(context.Background(), requestId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CallbacksApiService PostCrmV3ExtensionsAccountingCallbackCustomerSearchRequestIdDoCustomerSearch", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var requestId string

		httpRes, err := apiClient.CallbacksApi.PostCrmV3ExtensionsAccountingCallbackCustomerSearchRequestIdDoCustomerSearch(context.Background(), requestId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CallbacksApiService PostCrmV3ExtensionsAccountingCallbackExchangeRateRequestIdCreateExchangeRate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var requestId string

		httpRes, err := apiClient.CallbacksApi.PostCrmV3ExtensionsAccountingCallbackExchangeRateRequestIdCreateExchangeRate(context.Background(), requestId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CallbacksApiService PostCrmV3ExtensionsAccountingCallbackInvoiceCreateRequestIdCreateInvoice", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var requestId string

		httpRes, err := apiClient.CallbacksApi.PostCrmV3ExtensionsAccountingCallbackInvoiceCreateRequestIdCreateInvoice(context.Background(), requestId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CallbacksApiService PostCrmV3ExtensionsAccountingCallbackInvoicePdfRequestIdInvoicePdf", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var requestId string

		httpRes, err := apiClient.CallbacksApi.PostCrmV3ExtensionsAccountingCallbackInvoicePdfRequestIdInvoicePdf(context.Background(), requestId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CallbacksApiService PostCrmV3ExtensionsAccountingCallbackInvoiceSearchRequestIdDoInvoiceSearch", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var requestId string

		httpRes, err := apiClient.CallbacksApi.PostCrmV3ExtensionsAccountingCallbackInvoiceSearchRequestIdDoInvoiceSearch(context.Background(), requestId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CallbacksApiService PostCrmV3ExtensionsAccountingCallbackInvoicesRequestIdGetById", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var requestId string

		httpRes, err := apiClient.CallbacksApi.PostCrmV3ExtensionsAccountingCallbackInvoicesRequestIdGetById(context.Background(), requestId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CallbacksApiService PostCrmV3ExtensionsAccountingCallbackProductSearchRequestIdDoProductSearch", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var requestId string

		httpRes, err := apiClient.CallbacksApi.PostCrmV3ExtensionsAccountingCallbackProductSearchRequestIdDoProductSearch(context.Background(), requestId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CallbacksApiService PostCrmV3ExtensionsAccountingCallbackTaxSearchRequestIdDoTaxSearch", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var requestId string

		httpRes, err := apiClient.CallbacksApi.PostCrmV3ExtensionsAccountingCallbackTaxSearchRequestIdDoTaxSearch(context.Background(), requestId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CallbacksApiService PostCrmV3ExtensionsAccountingCallbackTermsRequestIdCreateTerm", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var requestId string

		httpRes, err := apiClient.CallbacksApi.PostCrmV3ExtensionsAccountingCallbackTermsRequestIdCreateTerm(context.Background(), requestId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
