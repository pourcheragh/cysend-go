# CostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FaceValueId** | **int32** | Face value ID | 
**FaceValue** | **float32** | Face value | 
**FaceValueCurrency** | **string** | Face value currency alphabetic ISO 4217 code | 
**GiftCardCode** | Pointer to **string** |  | [optional] 
**BeneficiaryInformation** | [**[]BeneficiaryInformation**](BeneficiaryInformation.md) | List of beneficiary information | 

## Methods

### NewCostRequest

`func NewCostRequest(faceValueId int32, faceValue float32, faceValueCurrency string, beneficiaryInformation []BeneficiaryInformation, ) *CostRequest`

NewCostRequest instantiates a new CostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCostRequestWithDefaults

`func NewCostRequestWithDefaults() *CostRequest`

NewCostRequestWithDefaults instantiates a new CostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFaceValueId

`func (o *CostRequest) GetFaceValueId() int32`

GetFaceValueId returns the FaceValueId field if non-nil, zero value otherwise.

### GetFaceValueIdOk

`func (o *CostRequest) GetFaceValueIdOk() (*int32, bool)`

GetFaceValueIdOk returns a tuple with the FaceValueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaceValueId

`func (o *CostRequest) SetFaceValueId(v int32)`

SetFaceValueId sets FaceValueId field to given value.


### GetFaceValue

`func (o *CostRequest) GetFaceValue() float32`

GetFaceValue returns the FaceValue field if non-nil, zero value otherwise.

### GetFaceValueOk

`func (o *CostRequest) GetFaceValueOk() (*float32, bool)`

GetFaceValueOk returns a tuple with the FaceValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaceValue

`func (o *CostRequest) SetFaceValue(v float32)`

SetFaceValue sets FaceValue field to given value.


### GetFaceValueCurrency

`func (o *CostRequest) GetFaceValueCurrency() string`

GetFaceValueCurrency returns the FaceValueCurrency field if non-nil, zero value otherwise.

### GetFaceValueCurrencyOk

`func (o *CostRequest) GetFaceValueCurrencyOk() (*string, bool)`

GetFaceValueCurrencyOk returns a tuple with the FaceValueCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaceValueCurrency

`func (o *CostRequest) SetFaceValueCurrency(v string)`

SetFaceValueCurrency sets FaceValueCurrency field to given value.


### GetGiftCardCode

`func (o *CostRequest) GetGiftCardCode() string`

GetGiftCardCode returns the GiftCardCode field if non-nil, zero value otherwise.

### GetGiftCardCodeOk

`func (o *CostRequest) GetGiftCardCodeOk() (*string, bool)`

GetGiftCardCodeOk returns a tuple with the GiftCardCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiftCardCode

`func (o *CostRequest) SetGiftCardCode(v string)`

SetGiftCardCode sets GiftCardCode field to given value.

### HasGiftCardCode

`func (o *CostRequest) HasGiftCardCode() bool`

HasGiftCardCode returns a boolean if a field has been set.

### GetBeneficiaryInformation

`func (o *CostRequest) GetBeneficiaryInformation() []BeneficiaryInformation`

GetBeneficiaryInformation returns the BeneficiaryInformation field if non-nil, zero value otherwise.

### GetBeneficiaryInformationOk

`func (o *CostRequest) GetBeneficiaryInformationOk() (*[]BeneficiaryInformation, bool)`

GetBeneficiaryInformationOk returns a tuple with the BeneficiaryInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiaryInformation

`func (o *CostRequest) SetBeneficiaryInformation(v []BeneficiaryInformation)`

SetBeneficiaryInformation sets BeneficiaryInformation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


