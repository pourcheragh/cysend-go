# ForceCallbackRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uid** | **string** | CY.SEND order unique ID | 
**UserUid** | **string** | CUSTOMER unique order ID | 

## Methods

### NewForceCallbackRequest

`func NewForceCallbackRequest(uid string, userUid string, ) *ForceCallbackRequest`

NewForceCallbackRequest instantiates a new ForceCallbackRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewForceCallbackRequestWithDefaults

`func NewForceCallbackRequestWithDefaults() *ForceCallbackRequest`

NewForceCallbackRequestWithDefaults instantiates a new ForceCallbackRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUid

`func (o *ForceCallbackRequest) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *ForceCallbackRequest) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *ForceCallbackRequest) SetUid(v string)`

SetUid sets Uid field to given value.


### GetUserUid

`func (o *ForceCallbackRequest) GetUserUid() string`

GetUserUid returns the UserUid field if non-nil, zero value otherwise.

### GetUserUidOk

`func (o *ForceCallbackRequest) GetUserUidOk() (*string, bool)`

GetUserUidOk returns a tuple with the UserUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserUid

`func (o *ForceCallbackRequest) SetUserUid(v string)`

SetUserUid sets UserUid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


