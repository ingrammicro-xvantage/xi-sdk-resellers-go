# PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceId** | Pointer to **NullableString** |  | [optional] 
**ResourceUId** | Pointer to **NullableString** | The resource id of the subscription product. | [optional] 
**ResourceName** | Pointer to **NullableString** | The name of the resource of the subscription product. | [optional] 
**VendorPartNumber** | Pointer to **NullableString** | Vendor’s part number for the subscription product. | [optional] 
**MinUnits** | Pointer to **NullableFloat32** | Minimum unit needs to purchased. | [optional] 
**MaxUnits** | Pointer to **NullableFloat32** | Maximum unit available for a purchase. | [optional] 
**Recurringpricemodel** | Pointer to **NullableString** | Recurring price model | [optional] 
**UnitOfMeasure** | Pointer to **NullableString** | Unit of mesaure for a subscription product. | [optional] 
**ResourcePricing** | Pointer to [**[]PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerResourcePricingInner**](PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerResourcePricingInner.md) |  | [optional] 
**Discounts** | Pointer to [**[]PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInner**](PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInner.md) |  | [optional] 
**Fees** | Pointer to [**[]PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerFeesInner**](PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerFeesInner.md) |  | [optional] 

## Methods

### NewPriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInner

`func NewPriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInner() *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInner`

NewPriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInner instantiates a new PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerWithDefaults

`func NewPriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerWithDefaults() *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInner`

NewPriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerWithDefaults instantiates a new PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceId

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInner) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInner) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInner) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInner) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### SetResourceIdNil

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInner) SetResourceIdNil(b bool)`

 SetResourceIdNil sets the value for ResourceId to be an explicit nil

### UnsetResourceId
`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInner) UnsetResourceId()`

UnsetResourceId ensures that no value is present for ResourceId, not even an explicit nil
### GetResourceUId

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInner) GetResourceUId() string`

GetResourceUId returns the ResourceUId field if non-nil, zero value otherwise.

### GetResourceUIdOk

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInner) GetResourceUIdOk() (*string, bool)`

GetResourceUIdOk returns a tuple with the ResourceUId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceUId

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInner) SetResourceUId(v string)`

SetResourceUId sets ResourceUId field to given value.

### HasResourceUId

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInner) HasResourceUId() bool`

HasResourceUId returns a boolean if a field has been set.

### SetResourceUIdNil

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInner) SetResourceUIdNil(b bool)`

 SetResourceUIdNil sets the value for ResourceUId to be an explicit nil

### UnsetResourceUId
`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInner) UnsetResourceUId()`

UnsetResourceUId ensures that no value is present for ResourceUId, not even an explicit nil
### GetResourceName

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInner) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInner) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInner) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInner) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### SetResourceNameNil

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInner) SetResourceNameNil(b bool)`

 SetResourceNameNil sets the value for ResourceName to be an explicit nil

### UnsetResourceName
`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInner) UnsetResourceName()`

UnsetResourceName ensures that no value is present for ResourceName, not even an explicit nil
### GetVendorPartNumber

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInner) GetVendorPartNumber() string`

GetVendorPartNumber returns the VendorPartNumber field if non-nil, zero value otherwise.

### GetVendorPartNumberOk

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInner) GetVendorPartNumberOk() (*string, bool)`

GetVendorPartNumberOk returns a tuple with the VendorPartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorPartNumber

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInner) SetVendorPartNumber(v string)`

SetVendorPartNumber sets VendorPartNumber field to given value.

### HasVendorPartNumber

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInner) HasVendorPartNumber() bool`

HasVendorPartNumber returns a boolean if a field has been set.

### SetVendorPartNumberNil

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInner) SetVendorPartNumberNil(b bool)`

 SetVendorPartNumberNil sets the value for VendorPartNumber to be an explicit nil

### UnsetVendorPartNumber
`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInner) UnsetVendorPartNumber()`

UnsetVendorPartNumber ensures that no value is present for VendorPartNumber, not even an explicit nil
### GetMinUnits

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInner) GetMinUnits() float32`

GetMinUnits returns the MinUnits field if non-nil, zero value otherwise.

### GetMinUnitsOk

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInner) GetMinUnitsOk() (*float32, bool)`

GetMinUnitsOk returns a tuple with the MinUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinUnits

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInner) SetMinUnits(v float32)`

SetMinUnits sets MinUnits field to given value.

### HasMinUnits

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInner) HasMinUnits() bool`

HasMinUnits returns a boolean if a field has been set.

### SetMinUnitsNil

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInner) SetMinUnitsNil(b bool)`

 SetMinUnitsNil sets the value for MinUnits to be an explicit nil

### UnsetMinUnits
`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInner) UnsetMinUnits()`

UnsetMinUnits ensures that no value is present for MinUnits, not even an explicit nil
### GetMaxUnits

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInner) GetMaxUnits() float32`

GetMaxUnits returns the MaxUnits field if non-nil, zero value otherwise.

### GetMaxUnitsOk

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInner) GetMaxUnitsOk() (*float32, bool)`

GetMaxUnitsOk returns a tuple with the MaxUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUnits

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInner) SetMaxUnits(v float32)`

SetMaxUnits sets MaxUnits field to given value.

### HasMaxUnits

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInner) HasMaxUnits() bool`

HasMaxUnits returns a boolean if a field has been set.

### SetMaxUnitsNil

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInner) SetMaxUnitsNil(b bool)`

 SetMaxUnitsNil sets the value for MaxUnits to be an explicit nil

### UnsetMaxUnits
`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInner) UnsetMaxUnits()`

UnsetMaxUnits ensures that no value is present for MaxUnits, not even an explicit nil
### GetRecurringpricemodel

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInner) GetRecurringpricemodel() string`

GetRecurringpricemodel returns the Recurringpricemodel field if non-nil, zero value otherwise.

### GetRecurringpricemodelOk

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInner) GetRecurringpricemodelOk() (*string, bool)`

GetRecurringpricemodelOk returns a tuple with the Recurringpricemodel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringpricemodel

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInner) SetRecurringpricemodel(v string)`

SetRecurringpricemodel sets Recurringpricemodel field to given value.

### HasRecurringpricemodel

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInner) HasRecurringpricemodel() bool`

HasRecurringpricemodel returns a boolean if a field has been set.

### SetRecurringpricemodelNil

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInner) SetRecurringpricemodelNil(b bool)`

 SetRecurringpricemodelNil sets the value for Recurringpricemodel to be an explicit nil

### UnsetRecurringpricemodel
`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInner) UnsetRecurringpricemodel()`

UnsetRecurringpricemodel ensures that no value is present for Recurringpricemodel, not even an explicit nil
### GetUnitOfMeasure

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInner) GetUnitOfMeasure() string`

GetUnitOfMeasure returns the UnitOfMeasure field if non-nil, zero value otherwise.

### GetUnitOfMeasureOk

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInner) GetUnitOfMeasureOk() (*string, bool)`

GetUnitOfMeasureOk returns a tuple with the UnitOfMeasure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitOfMeasure

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInner) SetUnitOfMeasure(v string)`

SetUnitOfMeasure sets UnitOfMeasure field to given value.

### HasUnitOfMeasure

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInner) HasUnitOfMeasure() bool`

HasUnitOfMeasure returns a boolean if a field has been set.

### SetUnitOfMeasureNil

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInner) SetUnitOfMeasureNil(b bool)`

 SetUnitOfMeasureNil sets the value for UnitOfMeasure to be an explicit nil

### UnsetUnitOfMeasure
`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInner) UnsetUnitOfMeasure()`

UnsetUnitOfMeasure ensures that no value is present for UnitOfMeasure, not even an explicit nil
### GetResourcePricing

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInner) GetResourcePricing() []PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerResourcePricingInner`

GetResourcePricing returns the ResourcePricing field if non-nil, zero value otherwise.

### GetResourcePricingOk

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInner) GetResourcePricingOk() (*[]PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerResourcePricingInner, bool)`

GetResourcePricingOk returns a tuple with the ResourcePricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePricing

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInner) SetResourcePricing(v []PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerResourcePricingInner)`

SetResourcePricing sets ResourcePricing field to given value.

### HasResourcePricing

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInner) HasResourcePricing() bool`

HasResourcePricing returns a boolean if a field has been set.

### GetDiscounts

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInner) GetDiscounts() []PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInner`

GetDiscounts returns the Discounts field if non-nil, zero value otherwise.

### GetDiscountsOk

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInner) GetDiscountsOk() (*[]PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInner, bool)`

GetDiscountsOk returns a tuple with the Discounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscounts

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInner) SetDiscounts(v []PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInner)`

SetDiscounts sets Discounts field to given value.

### HasDiscounts

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInner) HasDiscounts() bool`

HasDiscounts returns a boolean if a field has been set.

### GetFees

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInner) GetFees() []PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerFeesInner`

GetFees returns the Fees field if non-nil, zero value otherwise.

### GetFeesOk

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInner) GetFeesOk() (*[]PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerFeesInner, bool)`

GetFeesOk returns a tuple with the Fees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFees

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInner) SetFees(v []PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerFeesInner)`

SetFees sets Fees field to given value.

### HasFees

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInner) HasFees() bool`

HasFees returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


