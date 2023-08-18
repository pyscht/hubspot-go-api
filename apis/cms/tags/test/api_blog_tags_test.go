/*
Blog Post endpoints

Testing BlogTagsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package tags

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/pyscht/hubspot-go-api/apis/cms/tags"
)

func Test_tags_BlogTagsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test BlogTagsApiService DeleteCmsV3BlogsTagsObjectIdArchive", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var objectId string

		httpRes, err := apiClient.BlogTagsApi.DeleteCmsV3BlogsTagsObjectIdArchive(context.Background(), objectId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlogTagsApiService GetCmsV3BlogsTagsGetPage", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.BlogTagsApi.GetCmsV3BlogsTagsGetPage(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlogTagsApiService GetCmsV3BlogsTagsObjectIdGetById", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var objectId string

		resp, httpRes, err := apiClient.BlogTagsApi.GetCmsV3BlogsTagsObjectIdGetById(context.Background(), objectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlogTagsApiService PatchCmsV3BlogsTagsObjectIdUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var objectId string

		resp, httpRes, err := apiClient.BlogTagsApi.PatchCmsV3BlogsTagsObjectIdUpdate(context.Background(), objectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlogTagsApiService PostCmsV3BlogsTagsBatchArchiveArchiveBatch", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.BlogTagsApi.PostCmsV3BlogsTagsBatchArchiveArchiveBatch(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlogTagsApiService PostCmsV3BlogsTagsBatchCreateCreateBatch", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.BlogTagsApi.PostCmsV3BlogsTagsBatchCreateCreateBatch(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlogTagsApiService PostCmsV3BlogsTagsBatchReadReadBatch", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.BlogTagsApi.PostCmsV3BlogsTagsBatchReadReadBatch(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlogTagsApiService PostCmsV3BlogsTagsBatchUpdateUpdateBatch", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.BlogTagsApi.PostCmsV3BlogsTagsBatchUpdateUpdateBatch(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlogTagsApiService PostCmsV3BlogsTagsCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.BlogTagsApi.PostCmsV3BlogsTagsCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlogTagsApiService PostCmsV3BlogsTagsMultiLanguageAttachToLangGroupAttachToLangGroup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.BlogTagsApi.PostCmsV3BlogsTagsMultiLanguageAttachToLangGroupAttachToLangGroup(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlogTagsApiService PostCmsV3BlogsTagsMultiLanguageCreateLanguageVariationCreateLangVariation", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.BlogTagsApi.PostCmsV3BlogsTagsMultiLanguageCreateLanguageVariationCreateLangVariation(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlogTagsApiService PostCmsV3BlogsTagsMultiLanguageDetachFromLangGroupDetachFromLangGroup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.BlogTagsApi.PostCmsV3BlogsTagsMultiLanguageDetachFromLangGroupDetachFromLangGroup(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlogTagsApiService PostCmsV3BlogsTagsMultiLanguageUpdateLanguagesUpdateLangs", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.BlogTagsApi.PostCmsV3BlogsTagsMultiLanguageUpdateLanguagesUpdateLangs(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlogTagsApiService PutCmsV3BlogsTagsMultiLanguageSetNewLangPrimarySetLangPrimary", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.BlogTagsApi.PutCmsV3BlogsTagsMultiLanguageSetNewLangPrimarySetLangPrimary(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
