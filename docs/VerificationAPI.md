# \VerificationAPI

All URIs are relative to *https://www.cysend.com/openapi/prod/v5.2.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PhoneNumberVerificationRequest**](VerificationAPI.md#PhoneNumberVerificationRequest) | **Post** /verification/phone-number | Creates a phone number verification request
[**PhoneNumberVerify**](VerificationAPI.md#PhoneNumberVerify) | **Put** /verification/phone-number | Confirm a verification request



## PhoneNumberVerificationRequest

> VerificationStatus PhoneNumberVerificationRequest(ctx).VerificationRequest(verificationRequest).Execute()

Creates a phone number verification request



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
	verificationRequest := *openapiclient.NewVerificationRequest(openapiclient.VerificationMethod("sms"), "Number_example") // VerificationRequest | Verification request parameters

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VerificationAPI.PhoneNumberVerificationRequest(context.Background()).VerificationRequest(verificationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VerificationAPI.PhoneNumberVerificationRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PhoneNumberVerificationRequest`: VerificationStatus
	fmt.Fprintf(os.Stdout, "Response from `VerificationAPI.PhoneNumberVerificationRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPhoneNumberVerificationRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **verificationRequest** | [**VerificationRequest**](VerificationRequest.md) | Verification request parameters | 

### Return type

[**VerificationStatus**](VerificationStatus.md)

### Authorization

[BasicAuthentication](../README.md#BasicAuthentication), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PhoneNumberVerify

> VerificationStatus PhoneNumberVerify(ctx).Verification(verification).Execute()

Confirm a verification request



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
	verification := *openapiclient.NewVerification("Uid_example", openapiclient.VerificationMethod("sms"), "Code_example") // Verification | Verification parameters

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VerificationAPI.PhoneNumberVerify(context.Background()).Verification(verification).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VerificationAPI.PhoneNumberVerify``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PhoneNumberVerify`: VerificationStatus
	fmt.Fprintf(os.Stdout, "Response from `VerificationAPI.PhoneNumberVerify`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPhoneNumberVerifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **verification** | [**Verification**](Verification.md) | Verification parameters | 

### Return type

[**VerificationStatus**](VerificationStatus.md)

### Authorization

[BasicAuthentication](../README.md#BasicAuthentication), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

