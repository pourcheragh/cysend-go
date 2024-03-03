# Subscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TriggerUid** | **string** | Trigger unique ID | 
**ChannelUid** | **string** | Channel unique ID | 
**Destination** | [**[]NotificationDestinationParameter**](NotificationDestinationParameter.md) | List of destination parameters | 

## Methods

### NewSubscription

`func NewSubscription(triggerUid string, channelUid string, destination []NotificationDestinationParameter, ) *Subscription`

NewSubscription instantiates a new Subscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionWithDefaults

`func NewSubscriptionWithDefaults() *Subscription`

NewSubscriptionWithDefaults instantiates a new Subscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTriggerUid

`func (o *Subscription) GetTriggerUid() string`

GetTriggerUid returns the TriggerUid field if non-nil, zero value otherwise.

### GetTriggerUidOk

`func (o *Subscription) GetTriggerUidOk() (*string, bool)`

GetTriggerUidOk returns a tuple with the TriggerUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerUid

`func (o *Subscription) SetTriggerUid(v string)`

SetTriggerUid sets TriggerUid field to given value.


### GetChannelUid

`func (o *Subscription) GetChannelUid() string`

GetChannelUid returns the ChannelUid field if non-nil, zero value otherwise.

### GetChannelUidOk

`func (o *Subscription) GetChannelUidOk() (*string, bool)`

GetChannelUidOk returns a tuple with the ChannelUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelUid

`func (o *Subscription) SetChannelUid(v string)`

SetChannelUid sets ChannelUid field to given value.


### GetDestination

`func (o *Subscription) GetDestination() []NotificationDestinationParameter`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *Subscription) GetDestinationOk() (*[]NotificationDestinationParameter, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *Subscription) SetDestination(v []NotificationDestinationParameter)`

SetDestination sets Destination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


