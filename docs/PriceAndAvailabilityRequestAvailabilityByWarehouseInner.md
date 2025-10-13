# PriceAndAvailabilityRequestAvailabilityByWarehouseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailabilityByWarehouseId** | Pointer to **NullableString** | Plant/warehouse Id of a particular location in order to get just the inventory of that location. | [optional] 
**AvailabilityForAllLocation** | Pointer to **NullableBool** | Pass boolean value as input, if true the response will contain warehouse location details, if false the response will not hold warehouse location details. By default value is true. | [optional] 

## Methods

### NewPriceAndAvailabilityRequestAvailabilityByWarehouseInner

`func NewPriceAndAvailabilityRequestAvailabilityByWarehouseInner() *PriceAndAvailabilityRequestAvailabilityByWarehouseInner`

NewPriceAndAvailabilityRequestAvailabilityByWarehouseInner instantiates a new PriceAndAvailabilityRequestAvailabilityByWarehouseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceAndAvailabilityRequestAvailabilityByWarehouseInnerWithDefaults

`func NewPriceAndAvailabilityRequestAvailabilityByWarehouseInnerWithDefaults() *PriceAndAvailabilityRequestAvailabilityByWarehouseInner`

NewPriceAndAvailabilityRequestAvailabilityByWarehouseInnerWithDefaults instantiates a new PriceAndAvailabilityRequestAvailabilityByWarehouseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailabilityByWarehouseId

`func (o *PriceAndAvailabilityRequestAvailabilityByWarehouseInner) GetAvailabilityByWarehouseId() string`

GetAvailabilityByWarehouseId returns the AvailabilityByWarehouseId field if non-nil, zero value otherwise.

### GetAvailabilityByWarehouseIdOk

`func (o *PriceAndAvailabilityRequestAvailabilityByWarehouseInner) GetAvailabilityByWarehouseIdOk() (*string, bool)`

GetAvailabilityByWarehouseIdOk returns a tuple with the AvailabilityByWarehouseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityByWarehouseId

`func (o *PriceAndAvailabilityRequestAvailabilityByWarehouseInner) SetAvailabilityByWarehouseId(v string)`

SetAvailabilityByWarehouseId sets AvailabilityByWarehouseId field to given value.

### HasAvailabilityByWarehouseId

`func (o *PriceAndAvailabilityRequestAvailabilityByWarehouseInner) HasAvailabilityByWarehouseId() bool`

HasAvailabilityByWarehouseId returns a boolean if a field has been set.

### SetAvailabilityByWarehouseIdNil

`func (o *PriceAndAvailabilityRequestAvailabilityByWarehouseInner) SetAvailabilityByWarehouseIdNil(b bool)`

 SetAvailabilityByWarehouseIdNil sets the value for AvailabilityByWarehouseId to be an explicit nil

### UnsetAvailabilityByWarehouseId
`func (o *PriceAndAvailabilityRequestAvailabilityByWarehouseInner) UnsetAvailabilityByWarehouseId()`

UnsetAvailabilityByWarehouseId ensures that no value is present for AvailabilityByWarehouseId, not even an explicit nil
### GetAvailabilityForAllLocation

`func (o *PriceAndAvailabilityRequestAvailabilityByWarehouseInner) GetAvailabilityForAllLocation() bool`

GetAvailabilityForAllLocation returns the AvailabilityForAllLocation field if non-nil, zero value otherwise.

### GetAvailabilityForAllLocationOk

`func (o *PriceAndAvailabilityRequestAvailabilityByWarehouseInner) GetAvailabilityForAllLocationOk() (*bool, bool)`

GetAvailabilityForAllLocationOk returns a tuple with the AvailabilityForAllLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityForAllLocation

`func (o *PriceAndAvailabilityRequestAvailabilityByWarehouseInner) SetAvailabilityForAllLocation(v bool)`

SetAvailabilityForAllLocation sets AvailabilityForAllLocation field to given value.

### HasAvailabilityForAllLocation

`func (o *PriceAndAvailabilityRequestAvailabilityByWarehouseInner) HasAvailabilityForAllLocation() bool`

HasAvailabilityForAllLocation returns a boolean if a field has been set.

### SetAvailabilityForAllLocationNil

`func (o *PriceAndAvailabilityRequestAvailabilityByWarehouseInner) SetAvailabilityForAllLocationNil(b bool)`

 SetAvailabilityForAllLocationNil sets the value for AvailabilityForAllLocation to be an explicit nil

### UnsetAvailabilityForAllLocation
`func (o *PriceAndAvailabilityRequestAvailabilityByWarehouseInner) UnsetAvailabilityForAllLocation()`

UnsetAvailabilityForAllLocation ensures that no value is present for AvailabilityForAllLocation, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


