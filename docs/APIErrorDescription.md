# APIErrorDescription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Http** | Pointer to **int32** | HTTP status code | [optional] 
**Code** | Pointer to **string** | For more information, share this error code with the support team | [optional] 
**Message** | Pointer to **string** | Error code description | [optional] 

## Methods

### NewAPIErrorDescription

`func NewAPIErrorDescription() *APIErrorDescription`

NewAPIErrorDescription instantiates a new APIErrorDescription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPIErrorDescriptionWithDefaults

`func NewAPIErrorDescriptionWithDefaults() *APIErrorDescription`

NewAPIErrorDescriptionWithDefaults instantiates a new APIErrorDescription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHttp

`func (o *APIErrorDescription) GetHttp() int32`

GetHttp returns the Http field if non-nil, zero value otherwise.

### GetHttpOk

`func (o *APIErrorDescription) GetHttpOk() (*int32, bool)`

GetHttpOk returns a tuple with the Http field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttp

`func (o *APIErrorDescription) SetHttp(v int32)`

SetHttp sets Http field to given value.

### HasHttp

`func (o *APIErrorDescription) HasHttp() bool`

HasHttp returns a boolean if a field has been set.

### GetCode

`func (o *APIErrorDescription) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *APIErrorDescription) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *APIErrorDescription) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *APIErrorDescription) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *APIErrorDescription) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *APIErrorDescription) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *APIErrorDescription) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *APIErrorDescription) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


