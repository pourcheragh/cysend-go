# ProductParameter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Parameter name | 
**Description** | **string** | Parameter description | 
**Type** | **string** | Parameter type | 
**Pattern** | Pointer to **string** | Parameter pattern (only present for type input parameter) | [optional] 
**Elements** | Pointer to [**[]ProductParameterElementsInner**](ProductParameterElementsInner.md) | List of select box elements (only present for type select parameter) | [optional] 
**Required** | **bool** | If true, this parameter is mandatory. | 

## Methods

### NewProductParameter

`func NewProductParameter(name string, description string, type_ string, required bool, ) *ProductParameter`

NewProductParameter instantiates a new ProductParameter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductParameterWithDefaults

`func NewProductParameterWithDefaults() *ProductParameter`

NewProductParameterWithDefaults instantiates a new ProductParameter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ProductParameter) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProductParameter) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProductParameter) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ProductParameter) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProductParameter) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProductParameter) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetType

`func (o *ProductParameter) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProductParameter) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProductParameter) SetType(v string)`

SetType sets Type field to given value.


### GetPattern

`func (o *ProductParameter) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *ProductParameter) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *ProductParameter) SetPattern(v string)`

SetPattern sets Pattern field to given value.

### HasPattern

`func (o *ProductParameter) HasPattern() bool`

HasPattern returns a boolean if a field has been set.

### GetElements

`func (o *ProductParameter) GetElements() []ProductParameterElementsInner`

GetElements returns the Elements field if non-nil, zero value otherwise.

### GetElementsOk

`func (o *ProductParameter) GetElementsOk() (*[]ProductParameterElementsInner, bool)`

GetElementsOk returns a tuple with the Elements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElements

`func (o *ProductParameter) SetElements(v []ProductParameterElementsInner)`

SetElements sets Elements field to given value.

### HasElements

`func (o *ProductParameter) HasElements() bool`

HasElements returns a boolean if a field has been set.

### GetRequired

`func (o *ProductParameter) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *ProductParameter) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *ProductParameter) SetRequired(v bool)`

SetRequired sets Required field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


