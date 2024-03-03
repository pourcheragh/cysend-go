# FaceValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fixed** | [**[]FixedFaceValue**](FixedFaceValue.md) |  | 
**Range** | [**[]RangeFaceValue**](RangeFaceValue.md) |  | 

## Methods

### NewFaceValue

`func NewFaceValue(fixed []FixedFaceValue, range_ []RangeFaceValue, ) *FaceValue`

NewFaceValue instantiates a new FaceValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFaceValueWithDefaults

`func NewFaceValueWithDefaults() *FaceValue`

NewFaceValueWithDefaults instantiates a new FaceValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFixed

`func (o *FaceValue) GetFixed() []FixedFaceValue`

GetFixed returns the Fixed field if non-nil, zero value otherwise.

### GetFixedOk

`func (o *FaceValue) GetFixedOk() (*[]FixedFaceValue, bool)`

GetFixedOk returns a tuple with the Fixed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixed

`func (o *FaceValue) SetFixed(v []FixedFaceValue)`

SetFixed sets Fixed field to given value.


### GetRange

`func (o *FaceValue) GetRange() []RangeFaceValue`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *FaceValue) GetRangeOk() (*[]RangeFaceValue, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *FaceValue) SetRange(v []RangeFaceValue)`

SetRange sets Range field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


