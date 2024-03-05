# CostResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | **string** | Order date and time in ISO8601 format | 
**ProductId** | **int32** | CY.SEND product id | 
**FaceValueId** | **int32** | Face value ID | 
**FaceValue** | **decimal.Decimal** | Face value | 
**FaceValueCurrency** | **string** | Face value currency alphabetic ISO 4217 code | 
**GiftCardCode** | Pointer to **string** |  | [optional] 
**Cost** | **decimal.Decimal** | Order cost | 
**Currency** | **string** | Order currency alphabetic ISO 4217 code | 
**BeneficiaryInformation** | [**[]BeneficiaryInformation**](BeneficiaryInformation.md) | List of beneficiary information | 

## Methods

### NewCostResponse

`func NewCostResponse(date string, productId int32, faceValueId int32, faceValue decimal.Decimal, faceValueCurrency string, cost decimal.Decimal, currency string, beneficiaryInformation []BeneficiaryInformation, ) *CostResponse`

NewCostResponse instantiates a new CostResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCostResponseWithDefaults

`func NewCostResponseWithDefaults() *CostResponse`

NewCostResponseWithDefaults instantiates a new CostResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *CostResponse) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *CostResponse) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *CostResponse) SetDate(v string)`

SetDate sets Date field to given value.


### GetProductId

`func (o *CostResponse) GetProductId() int32`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *CostResponse) GetProductIdOk() (*int32, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *CostResponse) SetProductId(v int32)`

SetProductId sets ProductId field to given value.


### GetFaceValueId

`func (o *CostResponse) GetFaceValueId() int32`

GetFaceValueId returns the FaceValueId field if non-nil, zero value otherwise.

### GetFaceValueIdOk

`func (o *CostResponse) GetFaceValueIdOk() (*int32, bool)`

GetFaceValueIdOk returns a tuple with the FaceValueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaceValueId

`func (o *CostResponse) SetFaceValueId(v int32)`

SetFaceValueId sets FaceValueId field to given value.


### GetFaceValue

`func (o *CostResponse) GetFaceValue() decimal.Decimal`

GetFaceValue returns the FaceValue field if non-nil, zero value otherwise.

### GetFaceValueOk

`func (o *CostResponse) GetFaceValueOk() (*decimal.Decimal, bool)`

GetFaceValueOk returns a tuple with the FaceValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaceValue

`func (o *CostResponse) SetFaceValue(v decimal.Decimal)`

SetFaceValue sets FaceValue field to given value.


### GetFaceValueCurrency

`func (o *CostResponse) GetFaceValueCurrency() string`

GetFaceValueCurrency returns the FaceValueCurrency field if non-nil, zero value otherwise.

### GetFaceValueCurrencyOk

`func (o *CostResponse) GetFaceValueCurrencyOk() (*string, bool)`

GetFaceValueCurrencyOk returns a tuple with the FaceValueCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaceValueCurrency

`func (o *CostResponse) SetFaceValueCurrency(v string)`

SetFaceValueCurrency sets FaceValueCurrency field to given value.


### GetGiftCardCode

`func (o *CostResponse) GetGiftCardCode() string`

GetGiftCardCode returns the GiftCardCode field if non-nil, zero value otherwise.

### GetGiftCardCodeOk

`func (o *CostResponse) GetGiftCardCodeOk() (*string, bool)`

GetGiftCardCodeOk returns a tuple with the GiftCardCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiftCardCode

`func (o *CostResponse) SetGiftCardCode(v string)`

SetGiftCardCode sets GiftCardCode field to given value.

### HasGiftCardCode

`func (o *CostResponse) HasGiftCardCode() bool`

HasGiftCardCode returns a boolean if a field has been set.

### GetCost

`func (o *CostResponse) GetCost() decimal.Decimal`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *CostResponse) GetCostOk() (*decimal.Decimal, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *CostResponse) SetCost(v decimal.Decimal)`

SetCost sets Cost field to given value.


### GetCurrency

`func (o *CostResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CostResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CostResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetBeneficiaryInformation

`func (o *CostResponse) GetBeneficiaryInformation() []BeneficiaryInformation`

GetBeneficiaryInformation returns the BeneficiaryInformation field if non-nil, zero value otherwise.

### GetBeneficiaryInformationOk

`func (o *CostResponse) GetBeneficiaryInformationOk() (*[]BeneficiaryInformation, bool)`

GetBeneficiaryInformationOk returns a tuple with the BeneficiaryInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiaryInformation

`func (o *CostResponse) SetBeneficiaryInformation(v []BeneficiaryInformation)`

SetBeneficiaryInformation sets BeneficiaryInformation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


