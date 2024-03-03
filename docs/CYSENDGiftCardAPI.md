# \CYSENDGiftCardAPI

All URIs are relative to *https://www.cysend.com/openapi/prod/v5.2.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetActive**](CYSENDGiftCardAPI.md#GetActive) | **Get** /gift-card/{gift_card_code}/active | Checks if a CY.SEND gift card is active
[**GetGiftCard**](CYSENDGiftCardAPI.md#GetGiftCard) | **Get** /gift-card/{gift_card_code} | Retrieves a CY.SEND gift card information by PIN code
[**PutActive**](CYSENDGiftCardAPI.md#PutActive) | **Put** /gift-card/{gift_card_code}/active | Activates or deactivates a CY.SEND gift card



## GetActive

> GiftCardActive GetActive(ctx, giftCardCode).Execute()

Checks if a CY.SEND gift card is active



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
	giftCardCode := "123456789" // string | CY.SEND gift card PIN code

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CYSENDGiftCardAPI.GetActive(context.Background(), giftCardCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CYSENDGiftCardAPI.GetActive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetActive`: GiftCardActive
	fmt.Fprintf(os.Stdout, "Response from `CYSENDGiftCardAPI.GetActive`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**giftCardCode** | **string** | CY.SEND gift card PIN code | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetActiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GiftCardActive**](GiftCardActive.md)

### Authorization

[BasicAuthentication](../README.md#BasicAuthentication), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGiftCard

> GiftCard GetGiftCard(ctx, giftCardCode).Execute()

Retrieves a CY.SEND gift card information by PIN code



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
	giftCardCode := "123456789" // string | CY.SEND gift card PIN code

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CYSENDGiftCardAPI.GetGiftCard(context.Background(), giftCardCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CYSENDGiftCardAPI.GetGiftCard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGiftCard`: GiftCard
	fmt.Fprintf(os.Stdout, "Response from `CYSENDGiftCardAPI.GetGiftCard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**giftCardCode** | **string** | CY.SEND gift card PIN code | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGiftCardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GiftCard**](GiftCard.md)

### Authorization

[BasicAuthentication](../README.md#BasicAuthentication), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutActive

> PutActive(ctx, giftCardCode).Active(active).Execute()

Activates or deactivates a CY.SEND gift card



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
	giftCardCode := "123456789" // string | CY.SEND gift card PIN code
	active := true // bool | Active flag

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CYSENDGiftCardAPI.PutActive(context.Background(), giftCardCode).Active(active).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CYSENDGiftCardAPI.PutActive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**giftCardCode** | **string** | CY.SEND gift card PIN code | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutActiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **active** | **bool** | Active flag | 

### Return type

 (empty response body)

### Authorization

[BasicAuthentication](../README.md#BasicAuthentication), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

