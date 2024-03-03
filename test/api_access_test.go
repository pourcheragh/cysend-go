/*
CY.SEND® OpenAPI Documentation

Testing AccessAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package cysend

import (
	"context"
	"testing"

	openapiclient "github.com/pourcheragh/cysend-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_cysend_AccessAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AccessAPIService CreateAccessToken", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.AccessAPI.CreateAccessToken(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccessAPIService GetAccessToken", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.AccessAPI.GetAccessToken(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}