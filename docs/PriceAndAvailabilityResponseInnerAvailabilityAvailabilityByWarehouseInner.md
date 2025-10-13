# PriceAndAvailabilityResponseInnerAvailabilityAvailabilityByWarehouseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Location** | Pointer to **NullableString** | Indicates where (location) the product is available. | [optional] 
**WarehouseId** | Pointer to **NullableString** | Indicates where (Ingram Warehouse Id) the product is available. | [optional] 
**QuantityAvailable** | Pointer to **NullableInt32** | The quantity of the product available in a given warehouse. | [optional] 
**QuantityBackordered** | Pointer to **NullableInt32** | The quantity of a product backordered in a given warehouse. | [optional] 
**QuantityBackorderedEta** | Pointer to **NullableString** | The estimated time of arrival of a product that has been backordered in a given warehouse. | [optional] 
**QuantityOnOrder** | Pointer to **NullableInt32** | The quantity of the product on order. | [optional] 
**BackOrderInfo** | Pointer to [**[]PriceAndAvailabilityResponseInnerAvailabilityAvailabilityByWarehouseInnerBackOrderInfoInner**](PriceAndAvailabilityResponseInnerAvailabilityAvailabilityByWarehouseInnerBackOrderInfoInner.md) | *Currently, this feature is not available in these countries (Mexico, Turkey, New Zealand, Colombia, Chile, Brazil, Peru, Western Sahara). | [optional] 

## Methods

### NewPriceAndAvailabilityResponseInnerAvailabilityAvailabilityByWarehouseInner

`func NewPriceAndAvailabilityResponseInnerAvailabilityAvailabilityByWarehouseInner() *PriceAndAvailabilityResponseInnerAvailabilityAvailabilityByWarehouseInner`

NewPriceAndAvailabilityResponseInnerAvailabilityAvailabilityByWarehouseInner instantiates a new PriceAndAvailabilityResponseInnerAvailabilityAvailabilityByWarehouseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceAndAvailabilityResponseInnerAvailabilityAvailabilityByWarehouseInnerWithDefaults

`func NewPriceAndAvailabilityResponseInnerAvailabilityAvailabilityByWarehouseInnerWithDefaults() *PriceAndAvailabilityResponseInnerAvailabilityAvailabilityByWarehouseInner`

NewPriceAndAvailabilityResponseInnerAvailabilityAvailabilityByWarehouseInnerWithDefaults instantiates a new PriceAndAvailabilityResponseInnerAvailabilityAvailabilityByWarehouseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocation

`func (o *PriceAndAvailabilityResponseInnerAvailabilityAvailabilityByWarehouseInner) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *PriceAndAvailabilityResponseInnerAvailabilityAvailabilityByWarehouseInner) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *PriceAndAvailabilityResponseInnerAvailabilityAvailabilityByWarehouseInner) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *PriceAndAvailabilityResponseInnerAvailabilityAvailabilityByWarehouseInner) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *PriceAndAvailabilityResponseInnerAvailabilityAvailabilityByWarehouseInner) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *PriceAndAvailabilityResponseInnerAvailabilityAvailabilityByWarehouseInner) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetWarehouseId

`func (o *PriceAndAvailabilityResponseInnerAvailabilityAvailabilityByWarehouseInner) GetWarehouseId() string`

GetWarehouseId returns the WarehouseId field if non-nil, zero value otherwise.

### GetWarehouseIdOk

`func (o *PriceAndAvailabilityResponseInnerAvailabilityAvailabilityByWarehouseInner) GetWarehouseIdOk() (*string, bool)`

GetWarehouseIdOk returns a tuple with the WarehouseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseId

`func (o *PriceAndAvailabilityResponseInnerAvailabilityAvailabilityByWarehouseInner) SetWarehouseId(v string)`

SetWarehouseId sets WarehouseId field to given value.

### HasWarehouseId

`func (o *PriceAndAvailabilityResponseInnerAvailabilityAvailabilityByWarehouseInner) HasWarehouseId() bool`

HasWarehouseId returns a boolean if a field has been set.

### SetWarehouseIdNil

`func (o *PriceAndAvailabilityResponseInnerAvailabilityAvailabilityByWarehouseInner) SetWarehouseIdNil(b bool)`

 SetWarehouseIdNil sets the value for WarehouseId to be an explicit nil

### UnsetWarehouseId
`func (o *PriceAndAvailabilityResponseInnerAvailabilityAvailabilityByWarehouseInner) UnsetWarehouseId()`

UnsetWarehouseId ensures that no value is present for WarehouseId, not even an explicit nil
### GetQuantityAvailable

`func (o *PriceAndAvailabilityResponseInnerAvailabilityAvailabilityByWarehouseInner) GetQuantityAvailable() int32`

GetQuantityAvailable returns the QuantityAvailable field if non-nil, zero value otherwise.

### GetQuantityAvailableOk

`func (o *PriceAndAvailabilityResponseInnerAvailabilityAvailabilityByWarehouseInner) GetQuantityAvailableOk() (*int32, bool)`

GetQuantityAvailableOk returns a tuple with the QuantityAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityAvailable

`func (o *PriceAndAvailabilityResponseInnerAvailabilityAvailabilityByWarehouseInner) SetQuantityAvailable(v int32)`

SetQuantityAvailable sets QuantityAvailable field to given value.

### HasQuantityAvailable

`func (o *PriceAndAvailabilityResponseInnerAvailabilityAvailabilityByWarehouseInner) HasQuantityAvailable() bool`

HasQuantityAvailable returns a boolean if a field has been set.

### SetQuantityAvailableNil

`func (o *PriceAndAvailabilityResponseInnerAvailabilityAvailabilityByWarehouseInner) SetQuantityAvailableNil(b bool)`

 SetQuantityAvailableNil sets the value for QuantityAvailable to be an explicit nil

### UnsetQuantityAvailable
`func (o *PriceAndAvailabilityResponseInnerAvailabilityAvailabilityByWarehouseInner) UnsetQuantityAvailable()`

UnsetQuantityAvailable ensures that no value is present for QuantityAvailable, not even an explicit nil
### GetQuantityBackordered

`func (o *PriceAndAvailabilityResponseInnerAvailabilityAvailabilityByWarehouseInner) GetQuantityBackordered() int32`

GetQuantityBackordered returns the QuantityBackordered field if non-nil, zero value otherwise.

### GetQuantityBackorderedOk

`func (o *PriceAndAvailabilityResponseInnerAvailabilityAvailabilityByWarehouseInner) GetQuantityBackorderedOk() (*int32, bool)`

GetQuantityBackorderedOk returns a tuple with the QuantityBackordered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityBackordered

`func (o *PriceAndAvailabilityResponseInnerAvailabilityAvailabilityByWarehouseInner) SetQuantityBackordered(v int32)`

SetQuantityBackordered sets QuantityBackordered field to given value.

### HasQuantityBackordered

`func (o *PriceAndAvailabilityResponseInnerAvailabilityAvailabilityByWarehouseInner) HasQuantityBackordered() bool`

HasQuantityBackordered returns a boolean if a field has been set.

### SetQuantityBackorderedNil

`func (o *PriceAndAvailabilityResponseInnerAvailabilityAvailabilityByWarehouseInner) SetQuantityBackorderedNil(b bool)`

 SetQuantityBackorderedNil sets the value for QuantityBackordered to be an explicit nil

### UnsetQuantityBackordered
`func (o *PriceAndAvailabilityResponseInnerAvailabilityAvailabilityByWarehouseInner) UnsetQuantityBackordered()`

UnsetQuantityBackordered ensures that no value is present for QuantityBackordered, not even an explicit nil
### GetQuantityBackorderedEta

`func (o *PriceAndAvailabilityResponseInnerAvailabilityAvailabilityByWarehouseInner) GetQuantityBackorderedEta() string`

GetQuantityBackorderedEta returns the QuantityBackorderedEta field if non-nil, zero value otherwise.

### GetQuantityBackorderedEtaOk

`func (o *PriceAndAvailabilityResponseInnerAvailabilityAvailabilityByWarehouseInner) GetQuantityBackorderedEtaOk() (*string, bool)`

GetQuantityBackorderedEtaOk returns a tuple with the QuantityBackorderedEta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityBackorderedEta

`func (o *PriceAndAvailabilityResponseInnerAvailabilityAvailabilityByWarehouseInner) SetQuantityBackorderedEta(v string)`

SetQuantityBackorderedEta sets QuantityBackorderedEta field to given value.

### HasQuantityBackorderedEta

`func (o *PriceAndAvailabilityResponseInnerAvailabilityAvailabilityByWarehouseInner) HasQuantityBackorderedEta() bool`

HasQuantityBackorderedEta returns a boolean if a field has been set.

### SetQuantityBackorderedEtaNil

`func (o *PriceAndAvailabilityResponseInnerAvailabilityAvailabilityByWarehouseInner) SetQuantityBackorderedEtaNil(b bool)`

 SetQuantityBackorderedEtaNil sets the value for QuantityBackorderedEta to be an explicit nil

### UnsetQuantityBackorderedEta
`func (o *PriceAndAvailabilityResponseInnerAvailabilityAvailabilityByWarehouseInner) UnsetQuantityBackorderedEta()`

UnsetQuantityBackorderedEta ensures that no value is present for QuantityBackorderedEta, not even an explicit nil
### GetQuantityOnOrder

`func (o *PriceAndAvailabilityResponseInnerAvailabilityAvailabilityByWarehouseInner) GetQuantityOnOrder() int32`

GetQuantityOnOrder returns the QuantityOnOrder field if non-nil, zero value otherwise.

### GetQuantityOnOrderOk

`func (o *PriceAndAvailabilityResponseInnerAvailabilityAvailabilityByWarehouseInner) GetQuantityOnOrderOk() (*int32, bool)`

GetQuantityOnOrderOk returns a tuple with the QuantityOnOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityOnOrder

`func (o *PriceAndAvailabilityResponseInnerAvailabilityAvailabilityByWarehouseInner) SetQuantityOnOrder(v int32)`

SetQuantityOnOrder sets QuantityOnOrder field to given value.

### HasQuantityOnOrder

`func (o *PriceAndAvailabilityResponseInnerAvailabilityAvailabilityByWarehouseInner) HasQuantityOnOrder() bool`

HasQuantityOnOrder returns a boolean if a field has been set.

### SetQuantityOnOrderNil

`func (o *PriceAndAvailabilityResponseInnerAvailabilityAvailabilityByWarehouseInner) SetQuantityOnOrderNil(b bool)`

 SetQuantityOnOrderNil sets the value for QuantityOnOrder to be an explicit nil

### UnsetQuantityOnOrder
`func (o *PriceAndAvailabilityResponseInnerAvailabilityAvailabilityByWarehouseInner) UnsetQuantityOnOrder()`

UnsetQuantityOnOrder ensures that no value is present for QuantityOnOrder, not even an explicit nil
### GetBackOrderInfo

`func (o *PriceAndAvailabilityResponseInnerAvailabilityAvailabilityByWarehouseInner) GetBackOrderInfo() []PriceAndAvailabilityResponseInnerAvailabilityAvailabilityByWarehouseInnerBackOrderInfoInner`

GetBackOrderInfo returns the BackOrderInfo field if non-nil, zero value otherwise.

### GetBackOrderInfoOk

`func (o *PriceAndAvailabilityResponseInnerAvailabilityAvailabilityByWarehouseInner) GetBackOrderInfoOk() (*[]PriceAndAvailabilityResponseInnerAvailabilityAvailabilityByWarehouseInnerBackOrderInfoInner, bool)`

GetBackOrderInfoOk returns a tuple with the BackOrderInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackOrderInfo

`func (o *PriceAndAvailabilityResponseInnerAvailabilityAvailabilityByWarehouseInner) SetBackOrderInfo(v []PriceAndAvailabilityResponseInnerAvailabilityAvailabilityByWarehouseInnerBackOrderInfoInner)`

SetBackOrderInfo sets BackOrderInfo field to given value.

### HasBackOrderInfo

`func (o *PriceAndAvailabilityResponseInnerAvailabilityAvailabilityByWarehouseInner) HasBackOrderInfo() bool`

HasBackOrderInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


