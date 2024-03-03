# Product

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CountryZone** | **string** | Indication of the geographical zone in which the product can be used. Countries in a zone might vary depending on product. | 
**Countries** | **[]int32** | List of country IDs. A product can be in multiple countries. | 
**Categories** | **[]int32** | List of category IDs. A product can be in multiple categories. | 
**ProductId** | **int32** | Unique product ID | 
**ProductName** | **string** | Product name | 
**ProductDescription** | [**[]ProductProductDescriptionInner**](ProductProductDescriptionInner.md) | List of prepaid product description in various languages | 
**LogoUrl** | **string** | URL to the product logo | 
**Type** | [**ProductType**](ProductType.md) |  | 
**Promotion** | **bool** | True if there is a running promotion | 
**Maintenance** | **bool** | True if the product is under maintenance | 
**BeneficiaryInformation** | [**[]ProductParameter**](ProductParameter.md) | Each product has a different list or required beneficiary information to place a purchase order. Some will require beneficiary mobile number, account number, and others won’t require anything. | 
**FaceValueIds** | **[]int32** | List of available face values | 
**UsageInstructions** | Pointer to [**[]ProductUsageInstructionsInner**](ProductUsageInstructionsInner.md) | List of prepaid code usage instructions in various languages | [optional] 

## Methods

### NewProduct

`func NewProduct(countryZone string, countries []int32, categories []int32, productId int32, productName string, productDescription []ProductProductDescriptionInner, logoUrl string, type_ ProductType, promotion bool, maintenance bool, beneficiaryInformation []ProductParameter, faceValueIds []int32, ) *Product`

NewProduct instantiates a new Product object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductWithDefaults

`func NewProductWithDefaults() *Product`

NewProductWithDefaults instantiates a new Product object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountryZone

`func (o *Product) GetCountryZone() string`

GetCountryZone returns the CountryZone field if non-nil, zero value otherwise.

### GetCountryZoneOk

`func (o *Product) GetCountryZoneOk() (*string, bool)`

GetCountryZoneOk returns a tuple with the CountryZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryZone

`func (o *Product) SetCountryZone(v string)`

SetCountryZone sets CountryZone field to given value.


### GetCountries

`func (o *Product) GetCountries() []int32`

GetCountries returns the Countries field if non-nil, zero value otherwise.

### GetCountriesOk

`func (o *Product) GetCountriesOk() (*[]int32, bool)`

GetCountriesOk returns a tuple with the Countries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountries

`func (o *Product) SetCountries(v []int32)`

SetCountries sets Countries field to given value.


### GetCategories

`func (o *Product) GetCategories() []int32`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *Product) GetCategoriesOk() (*[]int32, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *Product) SetCategories(v []int32)`

SetCategories sets Categories field to given value.


### GetProductId

`func (o *Product) GetProductId() int32`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *Product) GetProductIdOk() (*int32, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *Product) SetProductId(v int32)`

SetProductId sets ProductId field to given value.


### GetProductName

`func (o *Product) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *Product) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *Product) SetProductName(v string)`

SetProductName sets ProductName field to given value.


### GetProductDescription

`func (o *Product) GetProductDescription() []ProductProductDescriptionInner`

GetProductDescription returns the ProductDescription field if non-nil, zero value otherwise.

### GetProductDescriptionOk

`func (o *Product) GetProductDescriptionOk() (*[]ProductProductDescriptionInner, bool)`

GetProductDescriptionOk returns a tuple with the ProductDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductDescription

`func (o *Product) SetProductDescription(v []ProductProductDescriptionInner)`

SetProductDescription sets ProductDescription field to given value.


### GetLogoUrl

`func (o *Product) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *Product) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *Product) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.


### GetType

`func (o *Product) GetType() ProductType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Product) GetTypeOk() (*ProductType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Product) SetType(v ProductType)`

SetType sets Type field to given value.


### GetPromotion

`func (o *Product) GetPromotion() bool`

GetPromotion returns the Promotion field if non-nil, zero value otherwise.

### GetPromotionOk

`func (o *Product) GetPromotionOk() (*bool, bool)`

GetPromotionOk returns a tuple with the Promotion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromotion

`func (o *Product) SetPromotion(v bool)`

SetPromotion sets Promotion field to given value.


### GetMaintenance

`func (o *Product) GetMaintenance() bool`

GetMaintenance returns the Maintenance field if non-nil, zero value otherwise.

### GetMaintenanceOk

`func (o *Product) GetMaintenanceOk() (*bool, bool)`

GetMaintenanceOk returns a tuple with the Maintenance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenance

`func (o *Product) SetMaintenance(v bool)`

SetMaintenance sets Maintenance field to given value.


### GetBeneficiaryInformation

`func (o *Product) GetBeneficiaryInformation() []ProductParameter`

GetBeneficiaryInformation returns the BeneficiaryInformation field if non-nil, zero value otherwise.

### GetBeneficiaryInformationOk

`func (o *Product) GetBeneficiaryInformationOk() (*[]ProductParameter, bool)`

GetBeneficiaryInformationOk returns a tuple with the BeneficiaryInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiaryInformation

`func (o *Product) SetBeneficiaryInformation(v []ProductParameter)`

SetBeneficiaryInformation sets BeneficiaryInformation field to given value.


### GetFaceValueIds

`func (o *Product) GetFaceValueIds() []int32`

GetFaceValueIds returns the FaceValueIds field if non-nil, zero value otherwise.

### GetFaceValueIdsOk

`func (o *Product) GetFaceValueIdsOk() (*[]int32, bool)`

GetFaceValueIdsOk returns a tuple with the FaceValueIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaceValueIds

`func (o *Product) SetFaceValueIds(v []int32)`

SetFaceValueIds sets FaceValueIds field to given value.


### GetUsageInstructions

`func (o *Product) GetUsageInstructions() []ProductUsageInstructionsInner`

GetUsageInstructions returns the UsageInstructions field if non-nil, zero value otherwise.

### GetUsageInstructionsOk

`func (o *Product) GetUsageInstructionsOk() (*[]ProductUsageInstructionsInner, bool)`

GetUsageInstructionsOk returns a tuple with the UsageInstructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageInstructions

`func (o *Product) SetUsageInstructions(v []ProductUsageInstructionsInner)`

SetUsageInstructions sets UsageInstructions field to given value.

### HasUsageInstructions

`func (o *Product) HasUsageInstructions() bool`

HasUsageInstructions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


