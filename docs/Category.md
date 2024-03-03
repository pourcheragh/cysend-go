# Category

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CategoryId** | **int32** | CY.SEND category ID | 
**CategoryName** | **string** | Category name | 
**Promotion** | **bool** | True if there is a running promotion | 
**ProductIds** | Pointer to **[]int32** | List of available products | [optional] 

## Methods

### NewCategory

`func NewCategory(categoryId int32, categoryName string, promotion bool, ) *Category`

NewCategory instantiates a new Category object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategoryWithDefaults

`func NewCategoryWithDefaults() *Category`

NewCategoryWithDefaults instantiates a new Category object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategoryId

`func (o *Category) GetCategoryId() int32`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *Category) GetCategoryIdOk() (*int32, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *Category) SetCategoryId(v int32)`

SetCategoryId sets CategoryId field to given value.


### GetCategoryName

`func (o *Category) GetCategoryName() string`

GetCategoryName returns the CategoryName field if non-nil, zero value otherwise.

### GetCategoryNameOk

`func (o *Category) GetCategoryNameOk() (*string, bool)`

GetCategoryNameOk returns a tuple with the CategoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryName

`func (o *Category) SetCategoryName(v string)`

SetCategoryName sets CategoryName field to given value.


### GetPromotion

`func (o *Category) GetPromotion() bool`

GetPromotion returns the Promotion field if non-nil, zero value otherwise.

### GetPromotionOk

`func (o *Category) GetPromotionOk() (*bool, bool)`

GetPromotionOk returns a tuple with the Promotion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromotion

`func (o *Category) SetPromotion(v bool)`

SetPromotion sets Promotion field to given value.


### GetProductIds

`func (o *Category) GetProductIds() []int32`

GetProductIds returns the ProductIds field if non-nil, zero value otherwise.

### GetProductIdsOk

`func (o *Category) GetProductIdsOk() (*[]int32, bool)`

GetProductIdsOk returns a tuple with the ProductIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductIds

`func (o *Category) SetProductIds(v []int32)`

SetProductIds sets ProductIds field to given value.

### HasProductIds

`func (o *Category) HasProductIds() bool`

HasProductIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


