# Maintenance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaintenanceId** | **int32** | Unique maintenance ID | 
**ProductId** | **int32** | Unique product ID | 
**StartTime** | **string** | Date and time of maintenance start in ISO8601 format | 
**StopTime** | **string** | Date and time of maintenance start in ISO8601 format | 
**Descriptions** | [**[]MaintenanceDescriptionsInner**](MaintenanceDescriptionsInner.md) | List of maintenance descriptions in various languages | 

## Methods

### NewMaintenance

`func NewMaintenance(maintenanceId int32, productId int32, startTime string, stopTime string, descriptions []MaintenanceDescriptionsInner, ) *Maintenance`

NewMaintenance instantiates a new Maintenance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMaintenanceWithDefaults

`func NewMaintenanceWithDefaults() *Maintenance`

NewMaintenanceWithDefaults instantiates a new Maintenance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaintenanceId

`func (o *Maintenance) GetMaintenanceId() int32`

GetMaintenanceId returns the MaintenanceId field if non-nil, zero value otherwise.

### GetMaintenanceIdOk

`func (o *Maintenance) GetMaintenanceIdOk() (*int32, bool)`

GetMaintenanceIdOk returns a tuple with the MaintenanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceId

`func (o *Maintenance) SetMaintenanceId(v int32)`

SetMaintenanceId sets MaintenanceId field to given value.


### GetProductId

`func (o *Maintenance) GetProductId() int32`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *Maintenance) GetProductIdOk() (*int32, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *Maintenance) SetProductId(v int32)`

SetProductId sets ProductId field to given value.


### GetStartTime

`func (o *Maintenance) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *Maintenance) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *Maintenance) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.


### GetStopTime

`func (o *Maintenance) GetStopTime() string`

GetStopTime returns the StopTime field if non-nil, zero value otherwise.

### GetStopTimeOk

`func (o *Maintenance) GetStopTimeOk() (*string, bool)`

GetStopTimeOk returns a tuple with the StopTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopTime

`func (o *Maintenance) SetStopTime(v string)`

SetStopTime sets StopTime field to given value.


### GetDescriptions

`func (o *Maintenance) GetDescriptions() []MaintenanceDescriptionsInner`

GetDescriptions returns the Descriptions field if non-nil, zero value otherwise.

### GetDescriptionsOk

`func (o *Maintenance) GetDescriptionsOk() (*[]MaintenanceDescriptionsInner, bool)`

GetDescriptionsOk returns a tuple with the Descriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptions

`func (o *Maintenance) SetDescriptions(v []MaintenanceDescriptionsInner)`

SetDescriptions sets Descriptions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


