# \StoreAPI

All URIs are relative to *https://www.cysend.com/openapi/prod/v5.2.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrder**](StoreAPI.md#CreateOrder) | **Post** /store/order | Place a purchase order for CY.SEND products
[**ForceCallback**](StoreAPI.md#ForceCallback) | **Post** /store/order/callback | Force a callback request
[**GetCatalogue**](StoreAPI.md#GetCatalogue) | **Get** /store/catalogue | Retrieves the complete CY.SEND product catalogue
[**GetCategories**](StoreAPI.md#GetCategories) | **Get** /store/catalogue/categories | Retrieves the CY.SEND product categories list from the catalogue
[**GetCountries**](StoreAPI.md#GetCountries) | **Get** /store/catalogue/countries | Retrieves the CY.SEND product catalogue country list
[**GetFaceValues**](StoreAPI.md#GetFaceValues) | **Get** /store/catalogue/face-value | Retrieves the CY.SEND product face value list
[**GetMaintenance**](StoreAPI.md#GetMaintenance) | **Get** /store/catalogue/maintenance | Retrieves current and future product maintenances
[**GetOrder**](StoreAPI.md#GetOrder) | **Get** /store/order | (Optional) Retrieves order information by CY.SEND order unique ID or by CUSTOMER order unique ID
[**GetProductCost**](StoreAPI.md#GetProductCost) | **Put** /store/order/cost | Retrieve the cost of a face value
[**GetProducts**](StoreAPI.md#GetProducts) | **Get** /store/catalogue/products | Retrieves the CY.SEND product catalogue list
[**GetPromotions**](StoreAPI.md#GetPromotions) | **Get** /store/catalogue/promotion | Retrieves current and future product promotions
[**GetResponses**](StoreAPI.md#GetResponses) | **Get** /store/order/responses | Retrieve the list of possible order response codes



## CreateOrder

> []OrderStatus CreateOrder(ctx).OrderRequest(orderRequest).Execute()

Place a purchase order for CY.SEND products



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
	orderRequest := []openapiclient.OrderRequest{*openapiclient.NewOrderRequest("myTID", int32(1234), float32(10), "CHF", []openapiclient.BeneficiaryInformation{*openapiclient.NewBeneficiaryInformation("mobile", "+41791234567")})} // []OrderRequest | Create purchase order request parameters

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StoreAPI.CreateOrder(context.Background()).OrderRequest(orderRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.CreateOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrder`: []OrderStatus
	fmt.Fprintf(os.Stdout, "Response from `StoreAPI.CreateOrder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderRequest** | [**[]OrderRequest**](OrderRequest.md) | Create purchase order request parameters | 

### Return type

[**[]OrderStatus**](OrderStatus.md)

### Authorization

[BasicAuthentication](../README.md#BasicAuthentication), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ForceCallback

> ForceCallback(ctx).ForceCallbackRequest(forceCallbackRequest).Execute()

Force a callback request



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
	forceCallbackRequest := *openapiclient.NewForceCallbackRequest("0101-123456-01", "myTID") // ForceCallbackRequest | Force callback request parameters

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StoreAPI.ForceCallback(context.Background()).ForceCallbackRequest(forceCallbackRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.ForceCallback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiForceCallbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **forceCallbackRequest** | [**ForceCallbackRequest**](ForceCallbackRequest.md) | Force callback request parameters | 

### Return type

 (empty response body)

### Authorization

[BasicAuthentication](../README.md#BasicAuthentication), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCatalogue

> Catalogue GetCatalogue(ctx).GiftCardCode(giftCardCode).ReturnUsageInstruction(returnUsageInstruction).Languages(languages).Execute()

Retrieves the complete CY.SEND product catalogue



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
	giftCardCode := "123456789" // string | CY.SEND gift card PIN code (optional)
	returnUsageInstruction := true // bool | Set to true to receive PIN products usage instructions. (optional) (default to false)
	languages := "en,fr,ru" // string | List of language code in ISO 639-1 format separated by a coma. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StoreAPI.GetCatalogue(context.Background()).GiftCardCode(giftCardCode).ReturnUsageInstruction(returnUsageInstruction).Languages(languages).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.GetCatalogue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCatalogue`: Catalogue
	fmt.Fprintf(os.Stdout, "Response from `StoreAPI.GetCatalogue`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCatalogueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **giftCardCode** | **string** | CY.SEND gift card PIN code | 
 **returnUsageInstruction** | **bool** | Set to true to receive PIN products usage instructions. | [default to false]
 **languages** | **string** | List of language code in ISO 639-1 format separated by a coma. | 

### Return type

[**Catalogue**](Catalogue.md)

### Authorization

[BasicAuthentication](../README.md#BasicAuthentication), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCategories

> []Category GetCategories(ctx).CountryId(countryId).CountryIso2(countryIso2).CategoryId(categoryId).Promotion(promotion).GiftCardCode(giftCardCode).Execute()

Retrieves the CY.SEND product categories list from the catalogue



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
	countryId := int32(233) // int32 | Country ID filter (optional)
	countryIso2 := "CH" // string | Country ISO 3166-1 alpha-2 filter (optional)
	categoryId := int32(14) // int32 | Category ID filter (optional)
	promotion := false // bool | Promotion filter (optional)
	giftCardCode := "123456789" // string | CY.SEND gift card PIN code (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StoreAPI.GetCategories(context.Background()).CountryId(countryId).CountryIso2(countryIso2).CategoryId(categoryId).Promotion(promotion).GiftCardCode(giftCardCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.GetCategories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCategories`: []Category
	fmt.Fprintf(os.Stdout, "Response from `StoreAPI.GetCategories`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCategoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **countryId** | **int32** | Country ID filter | 
 **countryIso2** | **string** | Country ISO 3166-1 alpha-2 filter | 
 **categoryId** | **int32** | Category ID filter | 
 **promotion** | **bool** | Promotion filter | 
 **giftCardCode** | **string** | CY.SEND gift card PIN code | 

### Return type

[**[]Category**](Category.md)

### Authorization

[BasicAuthentication](../README.md#BasicAuthentication), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCountries

> []Country GetCountries(ctx).CountryId(countryId).CountryIso2(countryIso2).Promotion(promotion).GiftCardCode(giftCardCode).Execute()

Retrieves the CY.SEND product catalogue country list



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
	countryId := int32(233) // int32 | Country ID filter (optional)
	countryIso2 := "CH" // string | Country ISO 3166-1 alpha-2 filter (optional)
	promotion := false // bool | Promotion filter (optional)
	giftCardCode := "123456789" // string | CY.SEND gift card PIN code (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StoreAPI.GetCountries(context.Background()).CountryId(countryId).CountryIso2(countryIso2).Promotion(promotion).GiftCardCode(giftCardCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.GetCountries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCountries`: []Country
	fmt.Fprintf(os.Stdout, "Response from `StoreAPI.GetCountries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCountriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **countryId** | **int32** | Country ID filter | 
 **countryIso2** | **string** | Country ISO 3166-1 alpha-2 filter | 
 **promotion** | **bool** | Promotion filter | 
 **giftCardCode** | **string** | CY.SEND gift card PIN code | 

### Return type

[**[]Country**](Country.md)

### Authorization

[BasicAuthentication](../README.md#BasicAuthentication), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFaceValues

> FaceValue GetFaceValues(ctx).CountryId(countryId).CountryIso2(countryIso2).CategoryId(categoryId).ProductId(productId).FaceValueId(faceValueId).Promotion(promotion).GiftCardCode(giftCardCode).Mobile(mobile).Execute()

Retrieves the CY.SEND product face value list



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
	countryId := int32(233) // int32 | Country ID filter (optional)
	countryIso2 := "CH" // string | Country ISO 3166-1 alpha-2 filter (optional)
	categoryId := int32(14) // int32 | Category ID filter (optional)
	productId := int32(1234) // int32 | Product ID filter (optional)
	faceValueId := int32(1234) // int32 | Face value ID filter (optional)
	promotion := false // bool | Promotion filter (optional)
	giftCardCode := "123456789" // string | CY.SEND gift card PIN code (optional)
	mobile := "+41791234567" // string | Mobile number you want to recharge. CY.SEND will detect the mobile operator of the number and return it. This service have a cost and will be billed on this request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StoreAPI.GetFaceValues(context.Background()).CountryId(countryId).CountryIso2(countryIso2).CategoryId(categoryId).ProductId(productId).FaceValueId(faceValueId).Promotion(promotion).GiftCardCode(giftCardCode).Mobile(mobile).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.GetFaceValues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFaceValues`: FaceValue
	fmt.Fprintf(os.Stdout, "Response from `StoreAPI.GetFaceValues`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFaceValuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **countryId** | **int32** | Country ID filter | 
 **countryIso2** | **string** | Country ISO 3166-1 alpha-2 filter | 
 **categoryId** | **int32** | Category ID filter | 
 **productId** | **int32** | Product ID filter | 
 **faceValueId** | **int32** | Face value ID filter | 
 **promotion** | **bool** | Promotion filter | 
 **giftCardCode** | **string** | CY.SEND gift card PIN code | 
 **mobile** | **string** | Mobile number you want to recharge. CY.SEND will detect the mobile operator of the number and return it. This service have a cost and will be billed on this request. | 

### Return type

[**FaceValue**](FaceValue.md)

### Authorization

[BasicAuthentication](../README.md#BasicAuthentication), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMaintenance

> []Maintenance GetMaintenance(ctx).ProductId(productId).Execute()

Retrieves current and future product maintenances



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
	productId := int32(1234) // int32 | Product ID filter (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StoreAPI.GetMaintenance(context.Background()).ProductId(productId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.GetMaintenance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMaintenance`: []Maintenance
	fmt.Fprintf(os.Stdout, "Response from `StoreAPI.GetMaintenance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMaintenanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productId** | **int32** | Product ID filter | 

### Return type

[**[]Maintenance**](Maintenance.md)

### Authorization

[BasicAuthentication](../README.md#BasicAuthentication), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrder

> []OrderStatus GetOrder(ctx).Uid(uid).UserUid(userUid).Execute()

(Optional) Retrieves order information by CY.SEND order unique ID or by CUSTOMER order unique ID



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
	uid := "uid_example" // string | CY.SEND order unique ID filter (optional)
	userUid := "userUid_example" // string | CUSTOMER unique order ID filter (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StoreAPI.GetOrder(context.Background()).Uid(uid).UserUid(userUid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.GetOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrder`: []OrderStatus
	fmt.Fprintf(os.Stdout, "Response from `StoreAPI.GetOrder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uid** | **string** | CY.SEND order unique ID filter | 
 **userUid** | **string** | CUSTOMER unique order ID filter | 

### Return type

[**[]OrderStatus**](OrderStatus.md)

### Authorization

[BasicAuthentication](../README.md#BasicAuthentication), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProductCost

> []CostResponse GetProductCost(ctx).CostRequest(costRequest).Execute()

Retrieve the cost of a face value



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
	costRequest := []openapiclient.CostRequest{*openapiclient.NewCostRequest(int32(1234), float32(10), "CHF", []openapiclient.BeneficiaryInformation{*openapiclient.NewBeneficiaryInformation("mobile", "+41791234567")})} // []CostRequest | Order you want to calculate the cost

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StoreAPI.GetProductCost(context.Background()).CostRequest(costRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.GetProductCost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProductCost`: []CostResponse
	fmt.Fprintf(os.Stdout, "Response from `StoreAPI.GetProductCost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProductCostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **costRequest** | [**[]CostRequest**](CostRequest.md) | Order you want to calculate the cost | 

### Return type

[**[]CostResponse**](CostResponse.md)

### Authorization

[BasicAuthentication](../README.md#BasicAuthentication), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProducts

> []Product GetProducts(ctx).CountryId(countryId).CountryIso2(countryIso2).CategoryId(categoryId).ProductId(productId).Maintenance(maintenance).Promotion(promotion).GiftCardCode(giftCardCode).Mobile(mobile).ReturnUsageInstruction(returnUsageInstruction).Languages(languages).Execute()

Retrieves the CY.SEND product catalogue list



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
	countryId := int32(233) // int32 | Country ID filter (optional)
	countryIso2 := "CH" // string | Country ISO 3166-1 alpha-2 filter (optional)
	categoryId := int32(14) // int32 | Category ID filter (optional)
	productId := int32(1234) // int32 | Product ID filter (optional)
	maintenance := false // bool | Maintenance filter (optional)
	promotion := false // bool | Promotion filter (optional)
	giftCardCode := "123456789" // string | CY.SEND gift card PIN code (optional)
	mobile := "+41791234567" // string | This mobile lookup filter will attempt to detect the mobile operator (check local regulations to be sure it is fully supported) and return the corresponding product. This operation has a cost of 0.01 EUR per query. Fund your account before you use this operation. [Load your balance.](https://www.cysend.com/en/add-funds) (optional)
	returnUsageInstruction := true // bool | Set to true to receive PIN products usage instructions. (optional) (default to false)
	languages := "en,fr,ru" // string | List of language code in ISO 639-1 format separated by a coma. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StoreAPI.GetProducts(context.Background()).CountryId(countryId).CountryIso2(countryIso2).CategoryId(categoryId).ProductId(productId).Maintenance(maintenance).Promotion(promotion).GiftCardCode(giftCardCode).Mobile(mobile).ReturnUsageInstruction(returnUsageInstruction).Languages(languages).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.GetProducts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProducts`: []Product
	fmt.Fprintf(os.Stdout, "Response from `StoreAPI.GetProducts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProductsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **countryId** | **int32** | Country ID filter | 
 **countryIso2** | **string** | Country ISO 3166-1 alpha-2 filter | 
 **categoryId** | **int32** | Category ID filter | 
 **productId** | **int32** | Product ID filter | 
 **maintenance** | **bool** | Maintenance filter | 
 **promotion** | **bool** | Promotion filter | 
 **giftCardCode** | **string** | CY.SEND gift card PIN code | 
 **mobile** | **string** | This mobile lookup filter will attempt to detect the mobile operator (check local regulations to be sure it is fully supported) and return the corresponding product. This operation has a cost of 0.01 EUR per query. Fund your account before you use this operation. [Load your balance.](https://www.cysend.com/en/add-funds) | 
 **returnUsageInstruction** | **bool** | Set to true to receive PIN products usage instructions. | [default to false]
 **languages** | **string** | List of language code in ISO 639-1 format separated by a coma. | 

### Return type

[**[]Product**](Product.md)

### Authorization

[BasicAuthentication](../README.md#BasicAuthentication), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPromotions

> []Promotion GetPromotions(ctx).ProductId(productId).Execute()

Retrieves current and future product promotions



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
	productId := int32(1234) // int32 | Product ID filter (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StoreAPI.GetPromotions(context.Background()).ProductId(productId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.GetPromotions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPromotions`: []Promotion
	fmt.Fprintf(os.Stdout, "Response from `StoreAPI.GetPromotions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPromotionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productId** | **int32** | Product ID filter | 

### Return type

[**[]Promotion**](Promotion.md)

### Authorization

[BasicAuthentication](../README.md#BasicAuthentication), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetResponses

> []OrderResponseCode GetResponses(ctx).Execute()

Retrieve the list of possible order response codes



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
	resp, r, err := apiClient.StoreAPI.GetResponses(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.GetResponses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetResponses`: []OrderResponseCode
	fmt.Fprintf(os.Stdout, "Response from `StoreAPI.GetResponses`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetResponsesRequest struct via the builder pattern


### Return type

[**[]OrderResponseCode**](OrderResponseCode.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

