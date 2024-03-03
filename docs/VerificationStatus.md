# VerificationStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uid** | Pointer to **string** | Verification request unique code | [optional] 
**Method** | Pointer to [**VerificationMethod**](VerificationMethod.md) |  | [optional] 
**Status** | **string** | Verification status | 

## Methods

### NewVerificationStatus

`func NewVerificationStatus(status string, ) *VerificationStatus`

NewVerificationStatus instantiates a new VerificationStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerificationStatusWithDefaults

`func NewVerificationStatusWithDefaults() *VerificationStatus`

NewVerificationStatusWithDefaults instantiates a new VerificationStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUid

`func (o *VerificationStatus) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *VerificationStatus) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *VerificationStatus) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *VerificationStatus) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetMethod

`func (o *VerificationStatus) GetMethod() VerificationMethod`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *VerificationStatus) GetMethodOk() (*VerificationMethod, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *VerificationStatus) SetMethod(v VerificationMethod)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *VerificationStatus) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetStatus

`func (o *VerificationStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VerificationStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VerificationStatus) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


