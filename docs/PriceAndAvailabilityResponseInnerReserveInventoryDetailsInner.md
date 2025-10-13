# PriceAndAvailabilityResponseInnerReserveInventoryDetailsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QuantityReserved** | Pointer to **NullableInt32** | The quantity of the product reserved for the customer. | [optional] 
**QuantityAvailable** | Pointer to **NullableInt32** | The availability of the product reserved. | [optional] 
**Effectivedate** | Pointer to **NullableString** | The reservation date for the product in UTC format. | [optional] 
**Expirydate** | Pointer to **NullableString** | The expiration date for the reservation of the product in UTC format. | [optional] 

## Methods

### NewPriceAndAvailabilityResponseInnerReserveInventoryDetailsInner

`func NewPriceAndAvailabilityResponseInnerReserveInventoryDetailsInner() *PriceAndAvailabilityResponseInnerReserveInventoryDetailsInner`

NewPriceAndAvailabilityResponseInnerReserveInventoryDetailsInner instantiates a new PriceAndAvailabilityResponseInnerReserveInventoryDetailsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceAndAvailabilityResponseInnerReserveInventoryDetailsInnerWithDefaults

`func NewPriceAndAvailabilityResponseInnerReserveInventoryDetailsInnerWithDefaults() *PriceAndAvailabilityResponseInnerReserveInventoryDetailsInner`

NewPriceAndAvailabilityResponseInnerReserveInventoryDetailsInnerWithDefaults instantiates a new PriceAndAvailabilityResponseInnerReserveInventoryDetailsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuantityReserved

`func (o *PriceAndAvailabilityResponseInnerReserveInventoryDetailsInner) GetQuantityReserved() int32`

GetQuantityReserved returns the QuantityReserved field if non-nil, zero value otherwise.

### GetQuantityReservedOk

`func (o *PriceAndAvailabilityResponseInnerReserveInventoryDetailsInner) GetQuantityReservedOk() (*int32, bool)`

GetQuantityReservedOk returns a tuple with the QuantityReserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityReserved

`func (o *PriceAndAvailabilityResponseInnerReserveInventoryDetailsInner) SetQuantityReserved(v int32)`

SetQuantityReserved sets QuantityReserved field to given value.

### HasQuantityReserved

`func (o *PriceAndAvailabilityResponseInnerReserveInventoryDetailsInner) HasQuantityReserved() bool`

HasQuantityReserved returns a boolean if a field has been set.

### SetQuantityReservedNil

`func (o *PriceAndAvailabilityResponseInnerReserveInventoryDetailsInner) SetQuantityReservedNil(b bool)`

 SetQuantityReservedNil sets the value for QuantityReserved to be an explicit nil

### UnsetQuantityReserved
`func (o *PriceAndAvailabilityResponseInnerReserveInventoryDetailsInner) UnsetQuantityReserved()`

UnsetQuantityReserved ensures that no value is present for QuantityReserved, not even an explicit nil
### GetQuantityAvailable

`func (o *PriceAndAvailabilityResponseInnerReserveInventoryDetailsInner) GetQuantityAvailable() int32`

GetQuantityAvailable returns the QuantityAvailable field if non-nil, zero value otherwise.

### GetQuantityAvailableOk

`func (o *PriceAndAvailabilityResponseInnerReserveInventoryDetailsInner) GetQuantityAvailableOk() (*int32, bool)`

GetQuantityAvailableOk returns a tuple with the QuantityAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityAvailable

`func (o *PriceAndAvailabilityResponseInnerReserveInventoryDetailsInner) SetQuantityAvailable(v int32)`

SetQuantityAvailable sets QuantityAvailable field to given value.

### HasQuantityAvailable

`func (o *PriceAndAvailabilityResponseInnerReserveInventoryDetailsInner) HasQuantityAvailable() bool`

HasQuantityAvailable returns a boolean if a field has been set.

### SetQuantityAvailableNil

`func (o *PriceAndAvailabilityResponseInnerReserveInventoryDetailsInner) SetQuantityAvailableNil(b bool)`

 SetQuantityAvailableNil sets the value for QuantityAvailable to be an explicit nil

### UnsetQuantityAvailable
`func (o *PriceAndAvailabilityResponseInnerReserveInventoryDetailsInner) UnsetQuantityAvailable()`

UnsetQuantityAvailable ensures that no value is present for QuantityAvailable, not even an explicit nil
### GetEffectivedate

`func (o *PriceAndAvailabilityResponseInnerReserveInventoryDetailsInner) GetEffectivedate() string`

GetEffectivedate returns the Effectivedate field if non-nil, zero value otherwise.

### GetEffectivedateOk

`func (o *PriceAndAvailabilityResponseInnerReserveInventoryDetailsInner) GetEffectivedateOk() (*string, bool)`

GetEffectivedateOk returns a tuple with the Effectivedate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectivedate

`func (o *PriceAndAvailabilityResponseInnerReserveInventoryDetailsInner) SetEffectivedate(v string)`

SetEffectivedate sets Effectivedate field to given value.

### HasEffectivedate

`func (o *PriceAndAvailabilityResponseInnerReserveInventoryDetailsInner) HasEffectivedate() bool`

HasEffectivedate returns a boolean if a field has been set.

### SetEffectivedateNil

`func (o *PriceAndAvailabilityResponseInnerReserveInventoryDetailsInner) SetEffectivedateNil(b bool)`

 SetEffectivedateNil sets the value for Effectivedate to be an explicit nil

### UnsetEffectivedate
`func (o *PriceAndAvailabilityResponseInnerReserveInventoryDetailsInner) UnsetEffectivedate()`

UnsetEffectivedate ensures that no value is present for Effectivedate, not even an explicit nil
### GetExpirydate

`func (o *PriceAndAvailabilityResponseInnerReserveInventoryDetailsInner) GetExpirydate() string`

GetExpirydate returns the Expirydate field if non-nil, zero value otherwise.

### GetExpirydateOk

`func (o *PriceAndAvailabilityResponseInnerReserveInventoryDetailsInner) GetExpirydateOk() (*string, bool)`

GetExpirydateOk returns a tuple with the Expirydate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirydate

`func (o *PriceAndAvailabilityResponseInnerReserveInventoryDetailsInner) SetExpirydate(v string)`

SetExpirydate sets Expirydate field to given value.

### HasExpirydate

`func (o *PriceAndAvailabilityResponseInnerReserveInventoryDetailsInner) HasExpirydate() bool`

HasExpirydate returns a boolean if a field has been set.

### SetExpirydateNil

`func (o *PriceAndAvailabilityResponseInnerReserveInventoryDetailsInner) SetExpirydateNil(b bool)`

 SetExpirydateNil sets the value for Expirydate to be an explicit nil

### UnsetExpirydate
`func (o *PriceAndAvailabilityResponseInnerReserveInventoryDetailsInner) UnsetExpirydate()`

UnsetExpirydate ensures that no value is present for Expirydate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


