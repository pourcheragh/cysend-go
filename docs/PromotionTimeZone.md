# PromotionTimeZone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Time zone name | 
**UtcOffset** | Pointer to **string** | UTC offset | [optional] 

## Methods

### NewPromotionTimeZone

`func NewPromotionTimeZone(name string, ) *PromotionTimeZone`

NewPromotionTimeZone instantiates a new PromotionTimeZone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPromotionTimeZoneWithDefaults

`func NewPromotionTimeZoneWithDefaults() *PromotionTimeZone`

NewPromotionTimeZoneWithDefaults instantiates a new PromotionTimeZone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PromotionTimeZone) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PromotionTimeZone) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PromotionTimeZone) SetName(v string)`

SetName sets Name field to given value.


### GetUtcOffset

`func (o *PromotionTimeZone) GetUtcOffset() string`

GetUtcOffset returns the UtcOffset field if non-nil, zero value otherwise.

### GetUtcOffsetOk

`func (o *PromotionTimeZone) GetUtcOffsetOk() (*string, bool)`

GetUtcOffsetOk returns a tuple with the UtcOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtcOffset

`func (o *PromotionTimeZone) SetUtcOffset(v string)`

SetUtcOffset sets UtcOffset field to given value.

### HasUtcOffset

`func (o *PromotionTimeZone) HasUtcOffset() bool`

HasUtcOffset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


