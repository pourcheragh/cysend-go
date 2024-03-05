# OrderStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uid** | **string** | CY.SEND order unique ID | 
**UserUid** | **string** | CUSTOMER unique order ID | 
**Date** | **string** | Order date and time in ISO8601 format | 
**FaceValueId** | **int32** | Face value ID | 
**FaceValue** | **decimal.Decimal** | Face value | 
**FaceValueCurrency** | **string** | Face value currency alphabetic ISO 4217 code | 
**GiftCardCode** | Pointer to **string** |  | [optional] 
**Scenario** | [**OrderScenario**](OrderScenario.md) |  | [default to ORDERSCENARIO_LIVE]
**CallbackUrl** | Pointer to **string** | CUSTOMER callback URL were order status will be send | [optional] 
**Cost** | **decimal.Decimal** | Order cost | 
**Currency** | **string** | Order currency alphabetic ISO 4217 code | 
**BeneficiaryInformation** | [**[]BeneficiaryInformation**](BeneficiaryInformation.md) | List of beneficiary information | 
**IssuerReference** | Pointer to **string** | Issuer reference | [optional] 
**PrepaidCode** | Pointer to [**PrepaidCode**](PrepaidCode.md) |  | [optional] 
**Status** | **string** | Order status | 
**ResponseCode** | **string** | Response code as per ISO8583 | 
**ResponseTitle** | **string** | Response code title | 
**ResponseDescription** | **string** | Response code description | 

## Methods

### NewOrderStatus

`func NewOrderStatus(uid string, userUid string, date string, faceValueId int32, faceValue decimal.Decimal, faceValueCurrency string, scenario OrderScenario, cost decimal.Decimal, currency string, beneficiaryInformation []BeneficiaryInformation, status string, responseCode string, responseTitle string, responseDescription string, ) *OrderStatus`

NewOrderStatus instantiates a new OrderStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderStatusWithDefaults

`func NewOrderStatusWithDefaults() *OrderStatus`

NewOrderStatusWithDefaults instantiates a new OrderStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUid

`func (o *OrderStatus) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *OrderStatus) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *OrderStatus) SetUid(v string)`

SetUid sets Uid field to given value.


### GetUserUid

`func (o *OrderStatus) GetUserUid() string`

GetUserUid returns the UserUid field if non-nil, zero value otherwise.

### GetUserUidOk

`func (o *OrderStatus) GetUserUidOk() (*string, bool)`

GetUserUidOk returns a tuple with the UserUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserUid

`func (o *OrderStatus) SetUserUid(v string)`

SetUserUid sets UserUid field to given value.


### GetDate

`func (o *OrderStatus) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *OrderStatus) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *OrderStatus) SetDate(v string)`

SetDate sets Date field to given value.


### GetFaceValueId

`func (o *OrderStatus) GetFaceValueId() int32`

GetFaceValueId returns the FaceValueId field if non-nil, zero value otherwise.

### GetFaceValueIdOk

`func (o *OrderStatus) GetFaceValueIdOk() (*int32, bool)`

GetFaceValueIdOk returns a tuple with the FaceValueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaceValueId

`func (o *OrderStatus) SetFaceValueId(v int32)`

SetFaceValueId sets FaceValueId field to given value.


### GetFaceValue

`func (o *OrderStatus) GetFaceValue() decimal.Decimal`

GetFaceValue returns the FaceValue field if non-nil, zero value otherwise.

### GetFaceValueOk

`func (o *OrderStatus) GetFaceValueOk() (*decimal.Decimal, bool)`

GetFaceValueOk returns a tuple with the FaceValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaceValue

`func (o *OrderStatus) SetFaceValue(v decimal.Decimal)`

SetFaceValue sets FaceValue field to given value.


### GetFaceValueCurrency

`func (o *OrderStatus) GetFaceValueCurrency() string`

GetFaceValueCurrency returns the FaceValueCurrency field if non-nil, zero value otherwise.

### GetFaceValueCurrencyOk

`func (o *OrderStatus) GetFaceValueCurrencyOk() (*string, bool)`

GetFaceValueCurrencyOk returns a tuple with the FaceValueCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaceValueCurrency

`func (o *OrderStatus) SetFaceValueCurrency(v string)`

SetFaceValueCurrency sets FaceValueCurrency field to given value.


### GetGiftCardCode

`func (o *OrderStatus) GetGiftCardCode() string`

GetGiftCardCode returns the GiftCardCode field if non-nil, zero value otherwise.

### GetGiftCardCodeOk

`func (o *OrderStatus) GetGiftCardCodeOk() (*string, bool)`

GetGiftCardCodeOk returns a tuple with the GiftCardCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiftCardCode

`func (o *OrderStatus) SetGiftCardCode(v string)`

SetGiftCardCode sets GiftCardCode field to given value.

### HasGiftCardCode

`func (o *OrderStatus) HasGiftCardCode() bool`

HasGiftCardCode returns a boolean if a field has been set.

### GetScenario

`func (o *OrderStatus) GetScenario() OrderScenario`

GetScenario returns the Scenario field if non-nil, zero value otherwise.

### GetScenarioOk

`func (o *OrderStatus) GetScenarioOk() (*OrderScenario, bool)`

GetScenarioOk returns a tuple with the Scenario field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScenario

`func (o *OrderStatus) SetScenario(v OrderScenario)`

SetScenario sets Scenario field to given value.


### GetCallbackUrl

`func (o *OrderStatus) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *OrderStatus) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *OrderStatus) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.

### HasCallbackUrl

`func (o *OrderStatus) HasCallbackUrl() bool`

HasCallbackUrl returns a boolean if a field has been set.

### GetCost

`func (o *OrderStatus) GetCost() decimal.Decimal`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *OrderStatus) GetCostOk() (*decimal.Decimal, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *OrderStatus) SetCost(v decimal.Decimal)`

SetCost sets Cost field to given value.


### GetCurrency

`func (o *OrderStatus) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *OrderStatus) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *OrderStatus) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetBeneficiaryInformation

`func (o *OrderStatus) GetBeneficiaryInformation() []BeneficiaryInformation`

GetBeneficiaryInformation returns the BeneficiaryInformation field if non-nil, zero value otherwise.

### GetBeneficiaryInformationOk

`func (o *OrderStatus) GetBeneficiaryInformationOk() (*[]BeneficiaryInformation, bool)`

GetBeneficiaryInformationOk returns a tuple with the BeneficiaryInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiaryInformation

`func (o *OrderStatus) SetBeneficiaryInformation(v []BeneficiaryInformation)`

SetBeneficiaryInformation sets BeneficiaryInformation field to given value.


### GetIssuerReference

`func (o *OrderStatus) GetIssuerReference() string`

GetIssuerReference returns the IssuerReference field if non-nil, zero value otherwise.

### GetIssuerReferenceOk

`func (o *OrderStatus) GetIssuerReferenceOk() (*string, bool)`

GetIssuerReferenceOk returns a tuple with the IssuerReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerReference

`func (o *OrderStatus) SetIssuerReference(v string)`

SetIssuerReference sets IssuerReference field to given value.

### HasIssuerReference

`func (o *OrderStatus) HasIssuerReference() bool`

HasIssuerReference returns a boolean if a field has been set.

### GetPrepaidCode

`func (o *OrderStatus) GetPrepaidCode() PrepaidCode`

GetPrepaidCode returns the PrepaidCode field if non-nil, zero value otherwise.

### GetPrepaidCodeOk

`func (o *OrderStatus) GetPrepaidCodeOk() (*PrepaidCode, bool)`

GetPrepaidCodeOk returns a tuple with the PrepaidCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrepaidCode

`func (o *OrderStatus) SetPrepaidCode(v PrepaidCode)`

SetPrepaidCode sets PrepaidCode field to given value.

### HasPrepaidCode

`func (o *OrderStatus) HasPrepaidCode() bool`

HasPrepaidCode returns a boolean if a field has been set.

### GetStatus

`func (o *OrderStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OrderStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OrderStatus) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetResponseCode

`func (o *OrderStatus) GetResponseCode() string`

GetResponseCode returns the ResponseCode field if non-nil, zero value otherwise.

### GetResponseCodeOk

`func (o *OrderStatus) GetResponseCodeOk() (*string, bool)`

GetResponseCodeOk returns a tuple with the ResponseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseCode

`func (o *OrderStatus) SetResponseCode(v string)`

SetResponseCode sets ResponseCode field to given value.


### GetResponseTitle

`func (o *OrderStatus) GetResponseTitle() string`

GetResponseTitle returns the ResponseTitle field if non-nil, zero value otherwise.

### GetResponseTitleOk

`func (o *OrderStatus) GetResponseTitleOk() (*string, bool)`

GetResponseTitleOk returns a tuple with the ResponseTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTitle

`func (o *OrderStatus) SetResponseTitle(v string)`

SetResponseTitle sets ResponseTitle field to given value.


### GetResponseDescription

`func (o *OrderStatus) GetResponseDescription() string`

GetResponseDescription returns the ResponseDescription field if non-nil, zero value otherwise.

### GetResponseDescriptionOk

`func (o *OrderStatus) GetResponseDescriptionOk() (*string, bool)`

GetResponseDescriptionOk returns a tuple with the ResponseDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseDescription

`func (o *OrderStatus) SetResponseDescription(v string)`

SetResponseDescription sets ResponseDescription field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


