# OrderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserUid** | **string** | CUSTOMER must assign a unique order ID | 
**FaceValueId** | **int32** | Face value ID | 
**FaceValue** | **float32** | Face value | 
**FaceValueCurrency** | **string** | Face value currency alphabetic ISO 4217 code | 
**GiftCardCode** | Pointer to **string** |  | [optional] 
**Scenario** | Pointer to [**OrderScenario**](OrderScenario.md) |  | [optional] [default to ORDERSCENARIO_LIVE]
**CallbackUrl** | Pointer to **string** | CUSTOMER callback URL were order status will be send | [optional] 
**BeneficiaryInformation** | [**[]BeneficiaryInformation**](BeneficiaryInformation.md) | List of beneficiary information | 
**Notifications** | Pointer to [**[]Subscription**](Subscription.md) | Notification subscriptions | [optional] 

## Methods

### NewOrderRequest

`func NewOrderRequest(userUid string, faceValueId int32, faceValue float32, faceValueCurrency string, beneficiaryInformation []BeneficiaryInformation, ) *OrderRequest`

NewOrderRequest instantiates a new OrderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderRequestWithDefaults

`func NewOrderRequestWithDefaults() *OrderRequest`

NewOrderRequestWithDefaults instantiates a new OrderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserUid

`func (o *OrderRequest) GetUserUid() string`

GetUserUid returns the UserUid field if non-nil, zero value otherwise.

### GetUserUidOk

`func (o *OrderRequest) GetUserUidOk() (*string, bool)`

GetUserUidOk returns a tuple with the UserUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserUid

`func (o *OrderRequest) SetUserUid(v string)`

SetUserUid sets UserUid field to given value.


### GetFaceValueId

`func (o *OrderRequest) GetFaceValueId() int32`

GetFaceValueId returns the FaceValueId field if non-nil, zero value otherwise.

### GetFaceValueIdOk

`func (o *OrderRequest) GetFaceValueIdOk() (*int32, bool)`

GetFaceValueIdOk returns a tuple with the FaceValueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaceValueId

`func (o *OrderRequest) SetFaceValueId(v int32)`

SetFaceValueId sets FaceValueId field to given value.


### GetFaceValue

`func (o *OrderRequest) GetFaceValue() float32`

GetFaceValue returns the FaceValue field if non-nil, zero value otherwise.

### GetFaceValueOk

`func (o *OrderRequest) GetFaceValueOk() (*float32, bool)`

GetFaceValueOk returns a tuple with the FaceValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaceValue

`func (o *OrderRequest) SetFaceValue(v float32)`

SetFaceValue sets FaceValue field to given value.


### GetFaceValueCurrency

`func (o *OrderRequest) GetFaceValueCurrency() string`

GetFaceValueCurrency returns the FaceValueCurrency field if non-nil, zero value otherwise.

### GetFaceValueCurrencyOk

`func (o *OrderRequest) GetFaceValueCurrencyOk() (*string, bool)`

GetFaceValueCurrencyOk returns a tuple with the FaceValueCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaceValueCurrency

`func (o *OrderRequest) SetFaceValueCurrency(v string)`

SetFaceValueCurrency sets FaceValueCurrency field to given value.


### GetGiftCardCode

`func (o *OrderRequest) GetGiftCardCode() string`

GetGiftCardCode returns the GiftCardCode field if non-nil, zero value otherwise.

### GetGiftCardCodeOk

`func (o *OrderRequest) GetGiftCardCodeOk() (*string, bool)`

GetGiftCardCodeOk returns a tuple with the GiftCardCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiftCardCode

`func (o *OrderRequest) SetGiftCardCode(v string)`

SetGiftCardCode sets GiftCardCode field to given value.

### HasGiftCardCode

`func (o *OrderRequest) HasGiftCardCode() bool`

HasGiftCardCode returns a boolean if a field has been set.

### GetScenario

`func (o *OrderRequest) GetScenario() OrderScenario`

GetScenario returns the Scenario field if non-nil, zero value otherwise.

### GetScenarioOk

`func (o *OrderRequest) GetScenarioOk() (*OrderScenario, bool)`

GetScenarioOk returns a tuple with the Scenario field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScenario

`func (o *OrderRequest) SetScenario(v OrderScenario)`

SetScenario sets Scenario field to given value.

### HasScenario

`func (o *OrderRequest) HasScenario() bool`

HasScenario returns a boolean if a field has been set.

### GetCallbackUrl

`func (o *OrderRequest) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *OrderRequest) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *OrderRequest) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.

### HasCallbackUrl

`func (o *OrderRequest) HasCallbackUrl() bool`

HasCallbackUrl returns a boolean if a field has been set.

### GetBeneficiaryInformation

`func (o *OrderRequest) GetBeneficiaryInformation() []BeneficiaryInformation`

GetBeneficiaryInformation returns the BeneficiaryInformation field if non-nil, zero value otherwise.

### GetBeneficiaryInformationOk

`func (o *OrderRequest) GetBeneficiaryInformationOk() (*[]BeneficiaryInformation, bool)`

GetBeneficiaryInformationOk returns a tuple with the BeneficiaryInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiaryInformation

`func (o *OrderRequest) SetBeneficiaryInformation(v []BeneficiaryInformation)`

SetBeneficiaryInformation sets BeneficiaryInformation field to given value.


### GetNotifications

`func (o *OrderRequest) GetNotifications() []Subscription`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *OrderRequest) GetNotificationsOk() (*[]Subscription, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *OrderRequest) SetNotifications(v []Subscription)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *OrderRequest) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


