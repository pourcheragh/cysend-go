# Country

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CountryId** | **int32** | CY.SEND country unique ID | 
**CountryName** | **string** | Country name | 
**CountryCode** | **int32** | Country dialling code | 
**CountryIsoAlpha2** | **string** | Country ISO 3166-1 alpha-2 | 
**CountryIsoAlpha3** | **string** | Country ISO 3166-1 alpha-3 | 
**Promotion** | **bool** | True if there is a running promotion | 
**CategoryIds** | Pointer to **[]int32** | List of available categories | [optional] 

## Methods

### NewCountry

`func NewCountry(countryId int32, countryName string, countryCode int32, countryIsoAlpha2 string, countryIsoAlpha3 string, promotion bool, ) *Country`

NewCountry instantiates a new Country object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCountryWithDefaults

`func NewCountryWithDefaults() *Country`

NewCountryWithDefaults instantiates a new Country object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountryId

`func (o *Country) GetCountryId() int32`

GetCountryId returns the CountryId field if non-nil, zero value otherwise.

### GetCountryIdOk

`func (o *Country) GetCountryIdOk() (*int32, bool)`

GetCountryIdOk returns a tuple with the CountryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryId

`func (o *Country) SetCountryId(v int32)`

SetCountryId sets CountryId field to given value.


### GetCountryName

`func (o *Country) GetCountryName() string`

GetCountryName returns the CountryName field if non-nil, zero value otherwise.

### GetCountryNameOk

`func (o *Country) GetCountryNameOk() (*string, bool)`

GetCountryNameOk returns a tuple with the CountryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryName

`func (o *Country) SetCountryName(v string)`

SetCountryName sets CountryName field to given value.


### GetCountryCode

`func (o *Country) GetCountryCode() int32`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *Country) GetCountryCodeOk() (*int32, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *Country) SetCountryCode(v int32)`

SetCountryCode sets CountryCode field to given value.


### GetCountryIsoAlpha2

`func (o *Country) GetCountryIsoAlpha2() string`

GetCountryIsoAlpha2 returns the CountryIsoAlpha2 field if non-nil, zero value otherwise.

### GetCountryIsoAlpha2Ok

`func (o *Country) GetCountryIsoAlpha2Ok() (*string, bool)`

GetCountryIsoAlpha2Ok returns a tuple with the CountryIsoAlpha2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryIsoAlpha2

`func (o *Country) SetCountryIsoAlpha2(v string)`

SetCountryIsoAlpha2 sets CountryIsoAlpha2 field to given value.


### GetCountryIsoAlpha3

`func (o *Country) GetCountryIsoAlpha3() string`

GetCountryIsoAlpha3 returns the CountryIsoAlpha3 field if non-nil, zero value otherwise.

### GetCountryIsoAlpha3Ok

`func (o *Country) GetCountryIsoAlpha3Ok() (*string, bool)`

GetCountryIsoAlpha3Ok returns a tuple with the CountryIsoAlpha3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryIsoAlpha3

`func (o *Country) SetCountryIsoAlpha3(v string)`

SetCountryIsoAlpha3 sets CountryIsoAlpha3 field to given value.


### GetPromotion

`func (o *Country) GetPromotion() bool`

GetPromotion returns the Promotion field if non-nil, zero value otherwise.

### GetPromotionOk

`func (o *Country) GetPromotionOk() (*bool, bool)`

GetPromotionOk returns a tuple with the Promotion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromotion

`func (o *Country) SetPromotion(v bool)`

SetPromotion sets Promotion field to given value.


### GetCategoryIds

`func (o *Country) GetCategoryIds() []int32`

GetCategoryIds returns the CategoryIds field if non-nil, zero value otherwise.

### GetCategoryIdsOk

`func (o *Country) GetCategoryIdsOk() (*[]int32, bool)`

GetCategoryIdsOk returns a tuple with the CategoryIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryIds

`func (o *Country) SetCategoryIds(v []int32)`

SetCategoryIds sets CategoryIds field to given value.

### HasCategoryIds

`func (o *Country) HasCategoryIds() bool`

HasCategoryIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


