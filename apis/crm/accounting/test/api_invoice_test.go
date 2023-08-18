/*
Accounting Extension

Testing InvoiceApiService

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

func Test_accounting_InvoiceApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test InvoiceApiService GetCrmV3ExtensionsAccountingInvoiceInvoiceIdGetById", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var invoiceId string

		resp, httpRes, err := apiClient.InvoiceApi.GetCrmV3ExtensionsAccountingInvoiceInvoiceIdGetById(context.Background(), invoiceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InvoiceApiService PatchCrmV3ExtensionsAccountingInvoiceInvoiceIdUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var invoiceId string

		resp, httpRes, err := apiClient.InvoiceApi.PatchCrmV3ExtensionsAccountingInvoiceInvoiceIdUpdate(context.Background(), invoiceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InvoiceApiService PostCrmV3ExtensionsAccountingInvoiceInvoiceIdPaymentCreatePayment", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var invoiceId string

		resp, httpRes, err := apiClient.InvoiceApi.PostCrmV3ExtensionsAccountingInvoiceInvoiceIdPaymentCreatePayment(context.Background(), invoiceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
