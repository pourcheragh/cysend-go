# \NotificationsAPI

All URIs are relative to *https://www.cysend.com/openapi/prod/v5.2.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNotificationChannels**](NotificationsAPI.md#GetNotificationChannels) | **Get** /notification/channels | Retrieve the list of notification channels
[**GetNotificationEvents**](NotificationsAPI.md#GetNotificationEvents) | **Get** /notification/events | Retrieve the list of notification events
[**GetNotificationResourceTypes**](NotificationsAPI.md#GetNotificationResourceTypes) | **Get** /notification/resource_types | Retrieve the list of notification resource types



## GetNotificationChannels

> []NotificationChannel GetNotificationChannels(ctx).Execute()

Retrieve the list of notification channels



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
	resp, r, err := apiClient.NotificationsAPI.GetNotificationChannels(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.GetNotificationChannels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNotificationChannels`: []NotificationChannel
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.GetNotificationChannels`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationChannelsRequest struct via the builder pattern


### Return type

[**[]NotificationChannel**](NotificationChannel.md)

### Authorization

[BasicAuthentication](../README.md#BasicAuthentication), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNotificationEvents

> []NotificationEvent GetNotificationEvents(ctx).Execute()

Retrieve the list of notification events



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
	resp, r, err := apiClient.NotificationsAPI.GetNotificationEvents(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.GetNotificationEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNotificationEvents`: []NotificationEvent
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.GetNotificationEvents`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationEventsRequest struct via the builder pattern


### Return type

[**[]NotificationEvent**](NotificationEvent.md)

### Authorization

[BasicAuthentication](../README.md#BasicAuthentication), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNotificationResourceTypes

> []NotificationResourceType GetNotificationResourceTypes(ctx).Execute()

Retrieve the list of notification resource types



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
	resp, r, err := apiClient.NotificationsAPI.GetNotificationResourceTypes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.GetNotificationResourceTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNotificationResourceTypes`: []NotificationResourceType
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.GetNotificationResourceTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationResourceTypesRequest struct via the builder pattern


### Return type

[**[]NotificationResourceType**](NotificationResourceType.md)

### Authorization

[BasicAuthentication](../README.md#BasicAuthentication), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
