# RangeFaceValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductId** | **int32** | Product ID | 
**FaceValueId** | **int32** | Face value ID | 
**FaceValueFrom** | **float32** | Minimum face value | 
**FaceValueTo** | **float32** | Maximum face value | 
**FaceValueStep** | **float32** | Incremental step of the face value range | 
**FaceValueCurrency** | **string** | Face value currency alphabetic ISO 4217 code | 
**MinimumCost** | **float32** | Minimum face value cost. | 
**MaximumCost** | **float32** | Maximum face value cost. | 
**CostCurrency** | **string** | The cost currency alphabetic ISO 4217 code. This currency will always be the same as your CY.SEND account currency. | 
**Promotion** | **bool** | True if there is a running promotion | 

## Methods

### NewRangeFaceValue

`func NewRangeFaceValue(productId int32, faceValueId int32, faceValueFrom float32, faceValueTo float32, faceValueStep float32, faceValueCurrency string, minimumCost float32, maximumCost float32, costCurrency string, promotion bool, ) *RangeFaceValue`

NewRangeFaceValue instantiates a new RangeFaceValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRangeFaceValueWithDefaults

`func NewRangeFaceValueWithDefaults() *RangeFaceValue`

NewRangeFaceValueWithDefaults instantiates a new RangeFaceValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductId

`func (o *RangeFaceValue) GetProductId() int32`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *RangeFaceValue) GetProductIdOk() (*int32, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *RangeFaceValue) SetProductId(v int32)`

SetProductId sets ProductId field to given value.


### GetFaceValueId

`func (o *RangeFaceValue) GetFaceValueId() int32`

GetFaceValueId returns the FaceValueId field if non-nil, zero value otherwise.

### GetFaceValueIdOk

`func (o *RangeFaceValue) GetFaceValueIdOk() (*int32, bool)`

GetFaceValueIdOk returns a tuple with the FaceValueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaceValueId

`func (o *RangeFaceValue) SetFaceValueId(v int32)`

SetFaceValueId sets FaceValueId field to given value.


### GetFaceValueFrom

`func (o *RangeFaceValue) GetFaceValueFrom() float32`

GetFaceValueFrom returns the FaceValueFrom field if non-nil, zero value otherwise.

### GetFaceValueFromOk

`func (o *RangeFaceValue) GetFaceValueFromOk() (*float32, bool)`

GetFaceValueFromOk returns a tuple with the FaceValueFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaceValueFrom

`func (o *RangeFaceValue) SetFaceValueFrom(v float32)`

SetFaceValueFrom sets FaceValueFrom field to given value.


### GetFaceValueTo

`func (o *RangeFaceValue) GetFaceValueTo() float32`

GetFaceValueTo returns the FaceValueTo field if non-nil, zero value otherwise.

### GetFaceValueToOk

`func (o *RangeFaceValue) GetFaceValueToOk() (*float32, bool)`

GetFaceValueToOk returns a tuple with the FaceValueTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaceValueTo

`func (o *RangeFaceValue) SetFaceValueTo(v float32)`

SetFaceValueTo sets FaceValueTo field to given value.


### GetFaceValueStep

`func (o *RangeFaceValue) GetFaceValueStep() float32`

GetFaceValueStep returns the FaceValueStep field if non-nil, zero value otherwise.

### GetFaceValueStepOk

`func (o *RangeFaceValue) GetFaceValueStepOk() (*float32, bool)`

GetFaceValueStepOk returns a tuple with the FaceValueStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaceValueStep

`func (o *RangeFaceValue) SetFaceValueStep(v float32)`

SetFaceValueStep sets FaceValueStep field to given value.


### GetFaceValueCurrency

`func (o *RangeFaceValue) GetFaceValueCurrency() string`

GetFaceValueCurrency returns the FaceValueCurrency field if non-nil, zero value otherwise.

### GetFaceValueCurrencyOk

`func (o *RangeFaceValue) GetFaceValueCurrencyOk() (*string, bool)`

GetFaceValueCurrencyOk returns a tuple with the FaceValueCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaceValueCurrency

`func (o *RangeFaceValue) SetFaceValueCurrency(v string)`

SetFaceValueCurrency sets FaceValueCurrency field to given value.


### GetMinimumCost

`func (o *RangeFaceValue) GetMinimumCost() float32`

GetMinimumCost returns the MinimumCost field if non-nil, zero value otherwise.

### GetMinimumCostOk

`func (o *RangeFaceValue) GetMinimumCostOk() (*float32, bool)`

GetMinimumCostOk returns a tuple with the MinimumCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumCost

`func (o *RangeFaceValue) SetMinimumCost(v float32)`

SetMinimumCost sets MinimumCost field to given value.


### GetMaximumCost

`func (o *RangeFaceValue) GetMaximumCost() float32`

GetMaximumCost returns the MaximumCost field if non-nil, zero value otherwise.

### GetMaximumCostOk

`func (o *RangeFaceValue) GetMaximumCostOk() (*float32, bool)`

GetMaximumCostOk returns a tuple with the MaximumCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumCost

`func (o *RangeFaceValue) SetMaximumCost(v float32)`

SetMaximumCost sets MaximumCost field to given value.


### GetCostCurrency

`func (o *RangeFaceValue) GetCostCurrency() string`

GetCostCurrency returns the CostCurrency field if non-nil, zero value otherwise.

### GetCostCurrencyOk

`func (o *RangeFaceValue) GetCostCurrencyOk() (*string, bool)`

GetCostCurrencyOk returns a tuple with the CostCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostCurrency

`func (o *RangeFaceValue) SetCostCurrency(v string)`

SetCostCurrency sets CostCurrency field to given value.


### GetPromotion

`func (o *RangeFaceValue) GetPromotion() bool`

GetPromotion returns the Promotion field if non-nil, zero value otherwise.

### GetPromotionOk

`func (o *RangeFaceValue) GetPromotionOk() (*bool, bool)`

GetPromotionOk returns a tuple with the Promotion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromotion

`func (o *RangeFaceValue) SetPromotion(v bool)`

SetPromotion sets Promotion field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


