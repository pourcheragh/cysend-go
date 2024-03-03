# NotificationEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uid** | **string** | Unique ID | 
**ResourceTypeUid** | **string** | Notification resource type unique ID | 
**Name** | **string** | Event name | 

## Methods

### NewNotificationEvent

`func NewNotificationEvent(uid string, resourceTypeUid string, name string, ) *NotificationEvent`

NewNotificationEvent instantiates a new NotificationEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationEventWithDefaults

`func NewNotificationEventWithDefaults() *NotificationEvent`

NewNotificationEventWithDefaults instantiates a new NotificationEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUid

`func (o *NotificationEvent) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *NotificationEvent) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *NotificationEvent) SetUid(v string)`

SetUid sets Uid field to given value.


### GetResourceTypeUid

`func (o *NotificationEvent) GetResourceTypeUid() string`

GetResourceTypeUid returns the ResourceTypeUid field if non-nil, zero value otherwise.

### GetResourceTypeUidOk

`func (o *NotificationEvent) GetResourceTypeUidOk() (*string, bool)`

GetResourceTypeUidOk returns a tuple with the ResourceTypeUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceTypeUid

`func (o *NotificationEvent) SetResourceTypeUid(v string)`

SetResourceTypeUid sets ResourceTypeUid field to given value.


### GetName

`func (o *NotificationEvent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NotificationEvent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NotificationEvent) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


