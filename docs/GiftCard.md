# GiftCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | CY.SEND gift card PIN code | 
**Serial** | **string** | CY.SEND gift card serial number | 
**Name** | **string** | CY.SEND gift card name | 
**Logo** | **string** | CY.SEND gift card logo URL | 
**Active** | **bool** | Flag to see if CY.SEND gift card is active or not | 
**FaceValue** | **float32** | CY.SEND gift card face value | 
**Currency** | **string** | CY.SEND gift card face-value currency alphabetic ISO 4217 code | 
**RemainingBalance** | **float32** | CY.SEND gift card remaining balance | 
**Expiration** | **string** | CY.SEND gift card expiry date and time in ISO8601 format | 
**UsageInstruction** | **string** | CY.SEND gift card usage information | 

## Methods

### NewGiftCard

`func NewGiftCard(code string, serial string, name string, logo string, active bool, faceValue float32, currency string, remainingBalance float32, expiration string, usageInstruction string, ) *GiftCard`

NewGiftCard instantiates a new GiftCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGiftCardWithDefaults

`func NewGiftCardWithDefaults() *GiftCard`

NewGiftCardWithDefaults instantiates a new GiftCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *GiftCard) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GiftCard) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GiftCard) SetCode(v string)`

SetCode sets Code field to given value.


### GetSerial

`func (o *GiftCard) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *GiftCard) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *GiftCard) SetSerial(v string)`

SetSerial sets Serial field to given value.


### GetName

`func (o *GiftCard) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GiftCard) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GiftCard) SetName(v string)`

SetName sets Name field to given value.


### GetLogo

`func (o *GiftCard) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *GiftCard) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *GiftCard) SetLogo(v string)`

SetLogo sets Logo field to given value.


### GetActive

`func (o *GiftCard) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *GiftCard) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *GiftCard) SetActive(v bool)`

SetActive sets Active field to given value.


### GetFaceValue

`func (o *GiftCard) GetFaceValue() float32`

GetFaceValue returns the FaceValue field if non-nil, zero value otherwise.

### GetFaceValueOk

`func (o *GiftCard) GetFaceValueOk() (*float32, bool)`

GetFaceValueOk returns a tuple with the FaceValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaceValue

`func (o *GiftCard) SetFaceValue(v float32)`

SetFaceValue sets FaceValue field to given value.


### GetCurrency

`func (o *GiftCard) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *GiftCard) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *GiftCard) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetRemainingBalance

`func (o *GiftCard) GetRemainingBalance() float32`

GetRemainingBalance returns the RemainingBalance field if non-nil, zero value otherwise.

### GetRemainingBalanceOk

`func (o *GiftCard) GetRemainingBalanceOk() (*float32, bool)`

GetRemainingBalanceOk returns a tuple with the RemainingBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingBalance

`func (o *GiftCard) SetRemainingBalance(v float32)`

SetRemainingBalance sets RemainingBalance field to given value.


### GetExpiration

`func (o *GiftCard) GetExpiration() string`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *GiftCard) GetExpirationOk() (*string, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *GiftCard) SetExpiration(v string)`

SetExpiration sets Expiration field to given value.


### GetUsageInstruction

`func (o *GiftCard) GetUsageInstruction() string`

GetUsageInstruction returns the UsageInstruction field if non-nil, zero value otherwise.

### GetUsageInstructionOk

`func (o *GiftCard) GetUsageInstructionOk() (*string, bool)`

GetUsageInstructionOk returns a tuple with the UsageInstruction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageInstruction

`func (o *GiftCard) SetUsageInstruction(v string)`

SetUsageInstruction sets UsageInstruction field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


