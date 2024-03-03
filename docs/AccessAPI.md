# \AccessAPI

All URIs are relative to *https://www.cysend.com/openapi/prod/v5.2.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccessToken**](AccessAPI.md#CreateAccessToken) | **Post** /access/token | Creates a new access token
[**GetAccessToken**](AccessAPI.md#GetAccessToken) | **Get** /access/token | Retrieves an existing access token



## CreateAccessToken

> Token CreateAccessToken(ctx).TokenOptions(tokenOptions).Execute()

Creates a new access token



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pourcheragh/cysend-go"
)

func main() {
	tokenOptions := *openapiclient.NewTokenOptions(true) // TokenOptions | Token creation parameters

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessAPI.CreateAccessToken(context.Background()).TokenOptions(tokenOptions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessAPI.CreateAccessToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAccessToken`: Token
	fmt.Fprintf(os.Stdout, "Response from `AccessAPI.CreateAccessToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccessTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenOptions** | [**TokenOptions**](TokenOptions.md) | Token creation parameters | 

### Return type

[**Token**](Token.md)

### Authorization

[BasicAuthentication](../README.md#BasicAuthentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccessToken

> Token GetAccessToken(ctx).Execute()

Retrieves an existing access token



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pourcheragh/cysend-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessAPI.GetAccessToken(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessAPI.GetAccessToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccessToken`: Token
	fmt.Fprintf(os.Stdout, "Response from `AccessAPI.GetAccessToken`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccessTokenRequest struct via the builder pattern


### Return type

[**Token**](Token.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

