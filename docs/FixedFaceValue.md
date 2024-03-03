# FixedFaceValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductId** | **int32** | Product ID | 
**FaceValueId** | **int32** | Face value ID | 
**FaceValue** | **float32** | Face value. This amount will be received by the beneficiary. | 
**FaceValueCurrency** | **string** | Face value currency alphabetic ISO 4217 code | 
**Cost** | **float32** | Cost of this face value. | 
**CostCurrency** | **string** | The cost currency alphabetic ISO 4217 code. This currency will always be the same as your CY.SEND account currency. | 
**Definition** | **string** | The product definition. Some products have an additional definition. i.e: Some operator bundles include additional products (SMS, Data, ...) and are textually describes in this field. | 
**Promotion** | **bool** | True if there is a running promotion | 

## Methods

### NewFixedFaceValue

`func NewFixedFaceValue(productId int32, faceValueId int32, faceValue float32, faceValueCurrency string, cost float32, costCurrency string, definition string, promotion bool, ) *FixedFaceValue`

NewFixedFaceValue instantiates a new FixedFaceValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFixedFaceValueWithDefaults

`func NewFixedFaceValueWithDefaults() *FixedFaceValue`

NewFixedFaceValueWithDefaults instantiates a new FixedFaceValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductId

`func (o *FixedFaceValue) GetProductId() int32`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *FixedFaceValue) GetProductIdOk() (*int32, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *FixedFaceValue) SetProductId(v int32)`

SetProductId sets ProductId field to given value.


### GetFaceValueId

`func (o *FixedFaceValue) GetFaceValueId() int32`

GetFaceValueId returns the FaceValueId field if non-nil, zero value otherwise.

### GetFaceValueIdOk

`func (o *FixedFaceValue) GetFaceValueIdOk() (*int32, bool)`

GetFaceValueIdOk returns a tuple with the FaceValueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaceValueId

`func (o *FixedFaceValue) SetFaceValueId(v int32)`

SetFaceValueId sets FaceValueId field to given value.


### GetFaceValue

`func (o *FixedFaceValue) GetFaceValue() float32`

GetFaceValue returns the FaceValue field if non-nil, zero value otherwise.

### GetFaceValueOk

`func (o *FixedFaceValue) GetFaceValueOk() (*float32, bool)`

GetFaceValueOk returns a tuple with the FaceValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaceValue

`func (o *FixedFaceValue) SetFaceValue(v float32)`

SetFaceValue sets FaceValue field to given value.


### GetFaceValueCurrency

`func (o *FixedFaceValue) GetFaceValueCurrency() string`

GetFaceValueCurrency returns the FaceValueCurrency field if non-nil, zero value otherwise.

### GetFaceValueCurrencyOk

`func (o *FixedFaceValue) GetFaceValueCurrencyOk() (*string, bool)`

GetFaceValueCurrencyOk returns a tuple with the FaceValueCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaceValueCurrency

`func (o *FixedFaceValue) SetFaceValueCurrency(v string)`

SetFaceValueCurrency sets FaceValueCurrency field to given value.


### GetCost

`func (o *FixedFaceValue) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *FixedFaceValue) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *FixedFaceValue) SetCost(v float32)`

SetCost sets Cost field to given value.


### GetCostCurrency

`func (o *FixedFaceValue) GetCostCurrency() string`

GetCostCurrency returns the CostCurrency field if non-nil, zero value otherwise.

### GetCostCurrencyOk

`func (o *FixedFaceValue) GetCostCurrencyOk() (*string, bool)`

GetCostCurrencyOk returns a tuple with the CostCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostCurrency

`func (o *FixedFaceValue) SetCostCurrency(v string)`

SetCostCurrency sets CostCurrency field to given value.


### GetDefinition

`func (o *FixedFaceValue) GetDefinition() string`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *FixedFaceValue) GetDefinitionOk() (*string, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *FixedFaceValue) SetDefinition(v string)`

SetDefinition sets Definition field to given value.


### GetPromotion

`func (o *FixedFaceValue) GetPromotion() bool`

GetPromotion returns the Promotion field if non-nil, zero value otherwise.

### GetPromotionOk

`func (o *FixedFaceValue) GetPromotionOk() (*bool, bool)`

GetPromotionOk returns a tuple with the Promotion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromotion

`func (o *FixedFaceValue) SetPromotion(v bool)`

SetPromotion sets Promotion field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


