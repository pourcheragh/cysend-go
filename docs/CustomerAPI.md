# \CustomerAPI

All URIs are relative to *https://www.cysend.com/openapi/prod/v5.2.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Balance**](CustomerAPI.md#Balance) | **Get** /customer/balance | Retrieves customer account default balance
[**GetSubscriptions**](CustomerAPI.md#GetSubscriptions) | **Get** /customer/subscriptions | Retrieves the list of subscriptions



## Balance

> CustomerBalance Balance(ctx).Execute()

Retrieves customer account default balance



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
	resp, r, err := apiClient.CustomerAPI.Balance(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerAPI.Balance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Balance`: CustomerBalance
	fmt.Fprintf(os.Stdout, "Response from `CustomerAPI.Balance`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiBalanceRequest struct via the builder pattern


### Return type

[**CustomerBalance**](CustomerBalance.md)

### Authorization

[BasicAuthentication](../README.md#BasicAuthentication), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubscriptions

> []Subscription GetSubscriptions(ctx).Execute()

Retrieves the list of subscriptions



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
	resp, r, err := apiClient.CustomerAPI.GetSubscriptions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerAPI.GetSubscriptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSubscriptions`: []Subscription
	fmt.Fprintf(os.Stdout, "Response from `CustomerAPI.GetSubscriptions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubscriptionsRequest struct via the builder pattern


### Return type

[**[]Subscription**](Subscription.md)

### Authorization

[BasicAuthentication](../README.md#BasicAuthentication), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

