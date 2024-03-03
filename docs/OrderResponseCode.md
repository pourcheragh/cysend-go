# OrderResponseCode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseCode** | **string** | Response code as per ISO8583 | 
**ResponseTitle** | **string** | Response code title | 
**ResponseDescription** | **string** | Response code description | 

## Methods

### NewOrderResponseCode

`func NewOrderResponseCode(responseCode string, responseTitle string, responseDescription string, ) *OrderResponseCode`

NewOrderResponseCode instantiates a new OrderResponseCode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderResponseCodeWithDefaults

`func NewOrderResponseCodeWithDefaults() *OrderResponseCode`

NewOrderResponseCodeWithDefaults instantiates a new OrderResponseCode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseCode

`func (o *OrderResponseCode) GetResponseCode() string`

GetResponseCode returns the ResponseCode field if non-nil, zero value otherwise.

### GetResponseCodeOk

`func (o *OrderResponseCode) GetResponseCodeOk() (*string, bool)`

GetResponseCodeOk returns a tuple with the ResponseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseCode

`func (o *OrderResponseCode) SetResponseCode(v string)`

SetResponseCode sets ResponseCode field to given value.


### GetResponseTitle

`func (o *OrderResponseCode) GetResponseTitle() string`

GetResponseTitle returns the ResponseTitle field if non-nil, zero value otherwise.

### GetResponseTitleOk

`func (o *OrderResponseCode) GetResponseTitleOk() (*string, bool)`

GetResponseTitleOk returns a tuple with the ResponseTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTitle

`func (o *OrderResponseCode) SetResponseTitle(v string)`

SetResponseTitle sets ResponseTitle field to given value.


### GetResponseDescription

`func (o *OrderResponseCode) GetResponseDescription() string`

GetResponseDescription returns the ResponseDescription field if non-nil, zero value otherwise.

### GetResponseDescriptionOk

`func (o *OrderResponseCode) GetResponseDescriptionOk() (*string, bool)`

GetResponseDescriptionOk returns a tuple with the ResponseDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseDescription

`func (o *OrderResponseCode) SetResponseDescription(v string)`

SetResponseDescription sets ResponseDescription field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


