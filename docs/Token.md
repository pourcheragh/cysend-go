# Token

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | Pointer to **string** | Token number | [optional] 
**TTL** | Pointer to **string** | Expiry date and time in ISO8601 format | [optional] 
**AuthorizedIps** | Pointer to **string** | Coma separated list of authorized IPs | [optional] 
**Permissions** | Pointer to [**[]Permission**](Permission.md) | Token permission list. The default permissions are identical as the API credentials used to create this token. | [optional] 

## Methods

### NewToken

`func NewToken() *Token`

NewToken instantiates a new Token object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenWithDefaults

`func NewTokenWithDefaults() *Token`

NewTokenWithDefaults instantiates a new Token object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *Token) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *Token) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *Token) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *Token) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetTTL

`func (o *Token) GetTTL() string`

GetTTL returns the TTL field if non-nil, zero value otherwise.

### GetTTLOk

`func (o *Token) GetTTLOk() (*string, bool)`

GetTTLOk returns a tuple with the TTL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTTL

`func (o *Token) SetTTL(v string)`

SetTTL sets TTL field to given value.

### HasTTL

`func (o *Token) HasTTL() bool`

HasTTL returns a boolean if a field has been set.

### GetAuthorizedIps

`func (o *Token) GetAuthorizedIps() string`

GetAuthorizedIps returns the AuthorizedIps field if non-nil, zero value otherwise.

### GetAuthorizedIpsOk

`func (o *Token) GetAuthorizedIpsOk() (*string, bool)`

GetAuthorizedIpsOk returns a tuple with the AuthorizedIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedIps

`func (o *Token) SetAuthorizedIps(v string)`

SetAuthorizedIps sets AuthorizedIps field to given value.

### HasAuthorizedIps

`func (o *Token) HasAuthorizedIps() bool`

HasAuthorizedIps returns a boolean if a field has been set.

### GetPermissions

`func (o *Token) GetPermissions() []Permission`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *Token) GetPermissionsOk() (*[]Permission, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *Token) SetPermissions(v []Permission)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *Token) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


