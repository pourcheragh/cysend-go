# VerificationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | [**VerificationMethod**](VerificationMethod.md) |  | 
**Number** | **string** | Phone number to verify. The number must be in E.164 format | 
**TemplateId** | Pointer to **int32** | Template ID for SMS verification | [optional] 
**Language** | Pointer to **string** | Language code in ISO 639-1 format (default &#39;en&#39; &#x3D; English) | [optional] 

## Methods

### NewVerificationRequest

`func NewVerificationRequest(method VerificationMethod, number string, ) *VerificationRequest`

NewVerificationRequest instantiates a new VerificationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerificationRequestWithDefaults

`func NewVerificationRequestWithDefaults() *VerificationRequest`

NewVerificationRequestWithDefaults instantiates a new VerificationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethod

`func (o *VerificationRequest) GetMethod() VerificationMethod`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *VerificationRequest) GetMethodOk() (*VerificationMethod, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *VerificationRequest) SetMethod(v VerificationMethod)`

SetMethod sets Method field to given value.


### GetNumber

`func (o *VerificationRequest) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *VerificationRequest) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *VerificationRequest) SetNumber(v string)`

SetNumber sets Number field to given value.


### GetTemplateId

`func (o *VerificationRequest) GetTemplateId() int32`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *VerificationRequest) GetTemplateIdOk() (*int32, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *VerificationRequest) SetTemplateId(v int32)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *VerificationRequest) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetLanguage

`func (o *VerificationRequest) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *VerificationRequest) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *VerificationRequest) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *VerificationRequest) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


