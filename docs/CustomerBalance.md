# CustomerBalance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Balance** | **float32** | CUSTOMER balance | 
**Currency** | **string** | CUSTOMER balance currency alphabetic ISO 4217 code | 

## Methods

### NewCustomerBalance

`func NewCustomerBalance(balance float32, currency string, ) *CustomerBalance`

NewCustomerBalance instantiates a new CustomerBalance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerBalanceWithDefaults

`func NewCustomerBalanceWithDefaults() *CustomerBalance`

NewCustomerBalanceWithDefaults instantiates a new CustomerBalance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalance

`func (o *CustomerBalance) GetBalance() float32`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *CustomerBalance) GetBalanceOk() (*float32, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *CustomerBalance) SetBalance(v float32)`

SetBalance sets Balance field to given value.


### GetCurrency

`func (o *CustomerBalance) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CustomerBalance) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CustomerBalance) SetCurrency(v string)`

SetCurrency sets Currency field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


