# \OpenAPIAPI

All URIs are relative to *https://www.cysend.com/openapi/prod/v5.2.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetErrors**](OpenAPIAPI.md#GetErrors) | **Get** /openapi/errors | Retrieves the list of CY.SEND OpenAPI errors
[**Ping**](OpenAPIAPI.md#Ping) | **Get** /openapi/ping | CY.SEND OpenAPI connection status monitor



## GetErrors

> []APIErrorDescription GetErrors(ctx).Execute()

Retrieves the list of CY.SEND OpenAPI errors



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
	resp, r, err := apiClient.OpenAPIAPI.GetErrors(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpenAPIAPI.GetErrors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetErrors`: []APIErrorDescription
	fmt.Fprintf(os.Stdout, "Response from `OpenAPIAPI.GetErrors`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetErrorsRequest struct via the builder pattern


### Return type

[**[]APIErrorDescription**](APIErrorDescription.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Ping

> string Ping(ctx).Execute()

CY.SEND OpenAPI connection status monitor



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
	resp, r, err := apiClient.OpenAPIAPI.Ping(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpenAPIAPI.Ping``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Ping`: string
	fmt.Fprintf(os.Stdout, "Response from `OpenAPIAPI.Ping`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPingRequest struct via the builder pattern


### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

