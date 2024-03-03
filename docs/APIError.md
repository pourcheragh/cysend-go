# APIError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | For more information, share this error code with the support team | [optional] 
**Message** | Pointer to **string** | Error code description | [optional] 
**Information** | Pointer to **string** | Additional information | [optional] 

## Methods

### NewAPIError

`func NewAPIError() *APIError`

NewAPIError instantiates a new APIError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPIErrorWithDefaults

`func NewAPIErrorWithDefaults() *APIError`

NewAPIErrorWithDefaults instantiates a new APIError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *APIError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *APIError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *APIError) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *APIError) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *APIError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *APIError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *APIError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *APIError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetInformation

`func (o *APIError) GetInformation() string`

GetInformation returns the Information field if non-nil, zero value otherwise.

### GetInformationOk

`func (o *APIError) GetInformationOk() (*string, bool)`

GetInformationOk returns a tuple with the Information field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInformation

`func (o *APIError) SetInformation(v string)`

SetInformation sets Information field to given value.

### HasInformation

`func (o *APIError) HasInformation() bool`

HasInformation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


