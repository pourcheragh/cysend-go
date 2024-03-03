# TokenOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpRestricted** | **bool** | If set to true, only IPs specified in &#39;authorized_ips&#39; parameter will be allowed to use this token. | 
**AuthorizedIps** | Pointer to **string** | Coma separated list of IPv4 addresses that will be authorized to use this token. The variable &#39;mine&#39; can be used to authorized the IP that requested the token. | [optional] [default to "mine"]
**TTL** | Pointer to **int32** | Token life in seconds. By default 3600 seconds (1 hour). | [optional] [default to 3600]

## Methods

### NewTokenOptions

`func NewTokenOptions(ipRestricted bool, ) *TokenOptions`

NewTokenOptions instantiates a new TokenOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenOptionsWithDefaults

`func NewTokenOptionsWithDefaults() *TokenOptions`

NewTokenOptionsWithDefaults instantiates a new TokenOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpRestricted

`func (o *TokenOptions) GetIpRestricted() bool`

GetIpRestricted returns the IpRestricted field if non-nil, zero value otherwise.

### GetIpRestrictedOk

`func (o *TokenOptions) GetIpRestrictedOk() (*bool, bool)`

GetIpRestrictedOk returns a tuple with the IpRestricted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRestricted

`func (o *TokenOptions) SetIpRestricted(v bool)`

SetIpRestricted sets IpRestricted field to given value.


### GetAuthorizedIps

`func (o *TokenOptions) GetAuthorizedIps() string`

GetAuthorizedIps returns the AuthorizedIps field if non-nil, zero value otherwise.

### GetAuthorizedIpsOk

`func (o *TokenOptions) GetAuthorizedIpsOk() (*string, bool)`

GetAuthorizedIpsOk returns a tuple with the AuthorizedIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedIps

`func (o *TokenOptions) SetAuthorizedIps(v string)`

SetAuthorizedIps sets AuthorizedIps field to given value.

### HasAuthorizedIps

`func (o *TokenOptions) HasAuthorizedIps() bool`

HasAuthorizedIps returns a boolean if a field has been set.

### GetTTL

`func (o *TokenOptions) GetTTL() int32`

GetTTL returns the TTL field if non-nil, zero value otherwise.

### GetTTLOk

`func (o *TokenOptions) GetTTLOk() (*int32, bool)`

GetTTLOk returns a tuple with the TTL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTTL

`func (o *TokenOptions) SetTTL(v int32)`

SetTTL sets TTL field to given value.

### HasTTL

`func (o *TokenOptions) HasTTL() bool`

HasTTL returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


