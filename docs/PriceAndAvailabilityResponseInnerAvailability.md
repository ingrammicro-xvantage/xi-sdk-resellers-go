# PriceAndAvailabilityResponseInnerAvailability

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | Pointer to **bool** | Boolean that indicates if the product ordered is available | [optional] 
**TotalAvailability** | Pointer to **NullableInt32** | The total amount of available products | [optional] 
**AvailabilityByWarehouse** | Pointer to [**[]PriceAndAvailabilityResponseInnerAvailabilityAvailabilityByWarehouseInner**](PriceAndAvailabilityResponseInnerAvailabilityAvailabilityByWarehouseInner.md) |  | [optional] 

## Methods

### NewPriceAndAvailabilityResponseInnerAvailability

`func NewPriceAndAvailabilityResponseInnerAvailability() *PriceAndAvailabilityResponseInnerAvailability`

NewPriceAndAvailabilityResponseInnerAvailability instantiates a new PriceAndAvailabilityResponseInnerAvailability object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceAndAvailabilityResponseInnerAvailabilityWithDefaults

`func NewPriceAndAvailabilityResponseInnerAvailabilityWithDefaults() *PriceAndAvailabilityResponseInnerAvailability`

NewPriceAndAvailabilityResponseInnerAvailabilityWithDefaults instantiates a new PriceAndAvailabilityResponseInnerAvailability object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *PriceAndAvailabilityResponseInnerAvailability) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *PriceAndAvailabilityResponseInnerAvailability) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *PriceAndAvailabilityResponseInnerAvailability) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *PriceAndAvailabilityResponseInnerAvailability) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetTotalAvailability

`func (o *PriceAndAvailabilityResponseInnerAvailability) GetTotalAvailability() int32`

GetTotalAvailability returns the TotalAvailability field if non-nil, zero value otherwise.

### GetTotalAvailabilityOk

`func (o *PriceAndAvailabilityResponseInnerAvailability) GetTotalAvailabilityOk() (*int32, bool)`

GetTotalAvailabilityOk returns a tuple with the TotalAvailability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAvailability

`func (o *PriceAndAvailabilityResponseInnerAvailability) SetTotalAvailability(v int32)`

SetTotalAvailability sets TotalAvailability field to given value.

### HasTotalAvailability

`func (o *PriceAndAvailabilityResponseInnerAvailability) HasTotalAvailability() bool`

HasTotalAvailability returns a boolean if a field has been set.

### SetTotalAvailabilityNil

`func (o *PriceAndAvailabilityResponseInnerAvailability) SetTotalAvailabilityNil(b bool)`

 SetTotalAvailabilityNil sets the value for TotalAvailability to be an explicit nil

### UnsetTotalAvailability
`func (o *PriceAndAvailabilityResponseInnerAvailability) UnsetTotalAvailability()`

UnsetTotalAvailability ensures that no value is present for TotalAvailability, not even an explicit nil
### GetAvailabilityByWarehouse

`func (o *PriceAndAvailabilityResponseInnerAvailability) GetAvailabilityByWarehouse() []PriceAndAvailabilityResponseInnerAvailabilityAvailabilityByWarehouseInner`

GetAvailabilityByWarehouse returns the AvailabilityByWarehouse field if non-nil, zero value otherwise.

### GetAvailabilityByWarehouseOk

`func (o *PriceAndAvailabilityResponseInnerAvailability) GetAvailabilityByWarehouseOk() (*[]PriceAndAvailabilityResponseInnerAvailabilityAvailabilityByWarehouseInner, bool)`

GetAvailabilityByWarehouseOk returns a tuple with the AvailabilityByWarehouse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityByWarehouse

`func (o *PriceAndAvailabilityResponseInnerAvailability) SetAvailabilityByWarehouse(v []PriceAndAvailabilityResponseInnerAvailabilityAvailabilityByWarehouseInner)`

SetAvailabilityByWarehouse sets AvailabilityByWarehouse field to given value.

### HasAvailabilityByWarehouse

`func (o *PriceAndAvailabilityResponseInnerAvailability) HasAvailabilityByWarehouse() bool`

HasAvailabilityByWarehouse returns a boolean if a field has been set.

### SetAvailabilityByWarehouseNil

`func (o *PriceAndAvailabilityResponseInnerAvailability) SetAvailabilityByWarehouseNil(b bool)`

 SetAvailabilityByWarehouseNil sets the value for AvailabilityByWarehouse to be an explicit nil

### UnsetAvailabilityByWarehouse
`func (o *PriceAndAvailabilityResponseInnerAvailability) UnsetAvailabilityByWarehouse()`

UnsetAvailabilityByWarehouse ensures that no value is present for AvailabilityByWarehouse, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


