# PrepaidCode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | Prepaid code | 
**Serial** | **string** | Prepaid code serial number | 
**PinCode** | Pointer to **string** | Security code | [optional] 
**Expiration** | **string** | Prepaid code expiry date and time in ISO8601 format | 
**PdfLink** | Pointer to **string** | Link to PDF file with pin and barcode | [optional] 
**UsageInstructions** | Pointer to [**[]ProductUsageInstructionsInner**](ProductUsageInstructionsInner.md) | List of prepaid code usage instructions in various languages | [optional] 

## Methods

### NewPrepaidCode

`func NewPrepaidCode(code string, serial string, expiration string, ) *PrepaidCode`

NewPrepaidCode instantiates a new PrepaidCode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrepaidCodeWithDefaults

`func NewPrepaidCodeWithDefaults() *PrepaidCode`

NewPrepaidCodeWithDefaults instantiates a new PrepaidCode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *PrepaidCode) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *PrepaidCode) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *PrepaidCode) SetCode(v string)`

SetCode sets Code field to given value.


### GetSerial

`func (o *PrepaidCode) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *PrepaidCode) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *PrepaidCode) SetSerial(v string)`

SetSerial sets Serial field to given value.


### GetPinCode

`func (o *PrepaidCode) GetPinCode() string`

GetPinCode returns the PinCode field if non-nil, zero value otherwise.

### GetPinCodeOk

`func (o *PrepaidCode) GetPinCodeOk() (*string, bool)`

GetPinCodeOk returns a tuple with the PinCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinCode

`func (o *PrepaidCode) SetPinCode(v string)`

SetPinCode sets PinCode field to given value.

### HasPinCode

`func (o *PrepaidCode) HasPinCode() bool`

HasPinCode returns a boolean if a field has been set.

### GetExpiration

`func (o *PrepaidCode) GetExpiration() string`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *PrepaidCode) GetExpirationOk() (*string, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *PrepaidCode) SetExpiration(v string)`

SetExpiration sets Expiration field to given value.


### GetPdfLink

`func (o *PrepaidCode) GetPdfLink() string`

GetPdfLink returns the PdfLink field if non-nil, zero value otherwise.

### GetPdfLinkOk

`func (o *PrepaidCode) GetPdfLinkOk() (*string, bool)`

GetPdfLinkOk returns a tuple with the PdfLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdfLink

`func (o *PrepaidCode) SetPdfLink(v string)`

SetPdfLink sets PdfLink field to given value.

### HasPdfLink

`func (o *PrepaidCode) HasPdfLink() bool`

HasPdfLink returns a boolean if a field has been set.

### GetUsageInstructions

`func (o *PrepaidCode) GetUsageInstructions() []ProductUsageInstructionsInner`

GetUsageInstructions returns the UsageInstructions field if non-nil, zero value otherwise.

### GetUsageInstructionsOk

`func (o *PrepaidCode) GetUsageInstructionsOk() (*[]ProductUsageInstructionsInner, bool)`

GetUsageInstructionsOk returns a tuple with the UsageInstructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageInstructions

`func (o *PrepaidCode) SetUsageInstructions(v []ProductUsageInstructionsInner)`

SetUsageInstructions sets UsageInstructions field to given value.

### HasUsageInstructions

`func (o *PrepaidCode) HasUsageInstructions() bool`

HasUsageInstructions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


