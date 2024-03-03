# Promotion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PromotionId** | **int32** | Unique promotion ID | 
**ProductId** | **int32** | Unique product ID | 
**StartTime** | **string** | Date and time of promotion start in ISO8601 format | 
**StopTime** | **string** | Date and time of promotion start in ISO8601 format | 
**TimeZone** | [**PromotionTimeZone**](PromotionTimeZone.md) |  | 
**Descriptions** | [**[]PromotionDescriptionsInner**](PromotionDescriptionsInner.md) | List of promotion descriptions in various languages | 
**FaceValueIds** | Pointer to **[]int32** | List of face values included in the promotion | [optional] 

## Methods

### NewPromotion

`func NewPromotion(promotionId int32, productId int32, startTime string, stopTime string, timeZone PromotionTimeZone, descriptions []PromotionDescriptionsInner, ) *Promotion`

NewPromotion instantiates a new Promotion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPromotionWithDefaults

`func NewPromotionWithDefaults() *Promotion`

NewPromotionWithDefaults instantiates a new Promotion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPromotionId

`func (o *Promotion) GetPromotionId() int32`

GetPromotionId returns the PromotionId field if non-nil, zero value otherwise.

### GetPromotionIdOk

`func (o *Promotion) GetPromotionIdOk() (*int32, bool)`

GetPromotionIdOk returns a tuple with the PromotionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromotionId

`func (o *Promotion) SetPromotionId(v int32)`

SetPromotionId sets PromotionId field to given value.


### GetProductId

`func (o *Promotion) GetProductId() int32`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *Promotion) GetProductIdOk() (*int32, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *Promotion) SetProductId(v int32)`

SetProductId sets ProductId field to given value.


### GetStartTime

`func (o *Promotion) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *Promotion) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *Promotion) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.


### GetStopTime

`func (o *Promotion) GetStopTime() string`

GetStopTime returns the StopTime field if non-nil, zero value otherwise.

### GetStopTimeOk

`func (o *Promotion) GetStopTimeOk() (*string, bool)`

GetStopTimeOk returns a tuple with the StopTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopTime

`func (o *Promotion) SetStopTime(v string)`

SetStopTime sets StopTime field to given value.


### GetTimeZone

`func (o *Promotion) GetTimeZone() PromotionTimeZone`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *Promotion) GetTimeZoneOk() (*PromotionTimeZone, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *Promotion) SetTimeZone(v PromotionTimeZone)`

SetTimeZone sets TimeZone field to given value.


### GetDescriptions

`func (o *Promotion) GetDescriptions() []PromotionDescriptionsInner`

GetDescriptions returns the Descriptions field if non-nil, zero value otherwise.

### GetDescriptionsOk

`func (o *Promotion) GetDescriptionsOk() (*[]PromotionDescriptionsInner, bool)`

GetDescriptionsOk returns a tuple with the Descriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptions

`func (o *Promotion) SetDescriptions(v []PromotionDescriptionsInner)`

SetDescriptions sets Descriptions field to given value.


### GetFaceValueIds

`func (o *Promotion) GetFaceValueIds() []int32`

GetFaceValueIds returns the FaceValueIds field if non-nil, zero value otherwise.

### GetFaceValueIdsOk

`func (o *Promotion) GetFaceValueIdsOk() (*[]int32, bool)`

GetFaceValueIdsOk returns a tuple with the FaceValueIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaceValueIds

`func (o *Promotion) SetFaceValueIds(v []int32)`

SetFaceValueIds sets FaceValueIds field to given value.

### HasFaceValueIds

`func (o *Promotion) HasFaceValueIds() bool`

HasFaceValueIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


