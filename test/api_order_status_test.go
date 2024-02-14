/*
XI Sdk Resellers

Testing OrderStatusAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package xi_sdk_resellers

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_xi_sdk_resellers_OrderStatusAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test OrderStatusAPIService ResellersV1WebhooksOrderstatuseventPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.OrderStatusAPI.ResellersV1WebhooksOrderstatuseventPost(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
