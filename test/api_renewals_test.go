/*
XI Sdk Resellers

Testing RenewalsAPIService

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

func Test_xi_sdk_resellers_RenewalsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RenewalsAPIService GetResellersV6Renewalsdetails", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var renewalId string

		resp, httpRes, err := apiClient.RenewalsAPI.GetResellersV6Renewalsdetails(context.Background(), renewalId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RenewalsAPIService PostRenewalssearch", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.RenewalsAPI.PostRenewalssearch(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
