# NotificationDestinationParameterDescription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name | 
**Required** | **bool** | If true means that the destination parameter is mandatory. | 

## Methods

### NewNotificationDestinationParameterDescription

`func NewNotificationDestinationParameterDescription(name string, required bool, ) *NotificationDestinationParameterDescription`

NewNotificationDestinationParameterDescription instantiates a new NotificationDestinationParameterDescription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationDestinationParameterDescriptionWithDefaults

`func NewNotificationDestinationParameterDescriptionWithDefaults() *NotificationDestinationParameterDescription`

NewNotificationDestinationParameterDescriptionWithDefaults instantiates a new NotificationDestinationParameterDescription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NotificationDestinationParameterDescription) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NotificationDestinationParameterDescription) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NotificationDestinationParameterDescription) SetName(v string)`

SetName sets Name field to given value.


### GetRequired

`func (o *NotificationDestinationParameterDescription) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *NotificationDestinationParameterDescription) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *NotificationDestinationParameterDescription) SetRequired(v bool)`

SetRequired sets Required field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


