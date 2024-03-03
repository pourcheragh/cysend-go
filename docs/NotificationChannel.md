# NotificationChannel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uid** | **string** | Unique ID | 
**Type** | **string** | Type | 
**Name** | **string** | Name | 
**Whitelabel** | **bool** | If true means notifications will be send using your templates; If false CY.SEND templates will be used. | 
**Destination** | [**[]NotificationDestinationParameterDescription**](NotificationDestinationParameterDescription.md) | List of destination parameters | 

## Methods

### NewNotificationChannel

`func NewNotificationChannel(uid string, type_ string, name string, whitelabel bool, destination []NotificationDestinationParameterDescription, ) *NotificationChannel`

NewNotificationChannel instantiates a new NotificationChannel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationChannelWithDefaults

`func NewNotificationChannelWithDefaults() *NotificationChannel`

NewNotificationChannelWithDefaults instantiates a new NotificationChannel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUid

`func (o *NotificationChannel) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *NotificationChannel) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *NotificationChannel) SetUid(v string)`

SetUid sets Uid field to given value.


### GetType

`func (o *NotificationChannel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NotificationChannel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NotificationChannel) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *NotificationChannel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NotificationChannel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NotificationChannel) SetName(v string)`

SetName sets Name field to given value.


### GetWhitelabel

`func (o *NotificationChannel) GetWhitelabel() bool`

GetWhitelabel returns the Whitelabel field if non-nil, zero value otherwise.

### GetWhitelabelOk

`func (o *NotificationChannel) GetWhitelabelOk() (*bool, bool)`

GetWhitelabelOk returns a tuple with the Whitelabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhitelabel

`func (o *NotificationChannel) SetWhitelabel(v bool)`

SetWhitelabel sets Whitelabel field to given value.


### GetDestination

`func (o *NotificationChannel) GetDestination() []NotificationDestinationParameterDescription`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *NotificationChannel) GetDestinationOk() (*[]NotificationDestinationParameterDescription, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *NotificationChannel) SetDestination(v []NotificationDestinationParameterDescription)`

SetDestination sets Destination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


