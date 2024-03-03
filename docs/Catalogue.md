# Catalogue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Countries** | [**[]Country**](Country.md) |  | 
**Categories** | [**[]Category**](Category.md) |  | 
**Products** | [**[]Product**](Product.md) |  | 
**FaceValues** | [**FaceValue**](FaceValue.md) |  | 
**Maintenance** | [**[]Maintenance**](Maintenance.md) |  | 
**Promotions** | Pointer to [**[]Promotion**](Promotion.md) |  | [optional] 

## Methods

### NewCatalogue

`func NewCatalogue(countries []Country, categories []Category, products []Product, faceValues FaceValue, maintenance []Maintenance, ) *Catalogue`

NewCatalogue instantiates a new Catalogue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogueWithDefaults

`func NewCatalogueWithDefaults() *Catalogue`

NewCatalogueWithDefaults instantiates a new Catalogue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountries

`func (o *Catalogue) GetCountries() []Country`

GetCountries returns the Countries field if non-nil, zero value otherwise.

### GetCountriesOk

`func (o *Catalogue) GetCountriesOk() (*[]Country, bool)`

GetCountriesOk returns a tuple with the Countries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountries

`func (o *Catalogue) SetCountries(v []Country)`

SetCountries sets Countries field to given value.


### GetCategories

`func (o *Catalogue) GetCategories() []Category`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *Catalogue) GetCategoriesOk() (*[]Category, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *Catalogue) SetCategories(v []Category)`

SetCategories sets Categories field to given value.


### GetProducts

`func (o *Catalogue) GetProducts() []Product`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *Catalogue) GetProductsOk() (*[]Product, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *Catalogue) SetProducts(v []Product)`

SetProducts sets Products field to given value.


### GetFaceValues

`func (o *Catalogue) GetFaceValues() FaceValue`

GetFaceValues returns the FaceValues field if non-nil, zero value otherwise.

### GetFaceValuesOk

`func (o *Catalogue) GetFaceValuesOk() (*FaceValue, bool)`

GetFaceValuesOk returns a tuple with the FaceValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaceValues

`func (o *Catalogue) SetFaceValues(v FaceValue)`

SetFaceValues sets FaceValues field to given value.


### GetMaintenance

`func (o *Catalogue) GetMaintenance() []Maintenance`

GetMaintenance returns the Maintenance field if non-nil, zero value otherwise.

### GetMaintenanceOk

`func (o *Catalogue) GetMaintenanceOk() (*[]Maintenance, bool)`

GetMaintenanceOk returns a tuple with the Maintenance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenance

`func (o *Catalogue) SetMaintenance(v []Maintenance)`

SetMaintenance sets Maintenance field to given value.


### GetPromotions

`func (o *Catalogue) GetPromotions() []Promotion`

GetPromotions returns the Promotions field if non-nil, zero value otherwise.

### GetPromotionsOk

`func (o *Catalogue) GetPromotionsOk() (*[]Promotion, bool)`

GetPromotionsOk returns a tuple with the Promotions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromotions

`func (o *Catalogue) SetPromotions(v []Promotion)`

SetPromotions sets Promotions field to given value.

### HasPromotions

`func (o *Catalogue) HasPromotions() bool`

HasPromotions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


