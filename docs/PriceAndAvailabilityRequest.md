# PriceAndAvailabilityRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShowAvailableDiscounts** | Pointer to **NullableBool** | Boolean value that will display Discount details in the response when true. | [optional] 
**ShowReserveInventoryDetails** | Pointer to **NullableBool** | Boolean value that will display reserve inventory details in the response when true. | [optional] 
**SpecialBidNumber** | Pointer to **NullableString** | Pre-approved special pricing/bid number provided to the reseller by the vendor for special pricing and discounts. Used to track the bid number where different line items have different bid numbers. | [optional] 
**AvailabilityByWarehouse** | Pointer to [**[]PriceAndAvailabilityRequestAvailabilityByWarehouseInner**](PriceAndAvailabilityRequestAvailabilityByWarehouseInner.md) |  | [optional] 
**Products** | Pointer to [**[]PriceAndAvailabilityRequestProductsInner**](PriceAndAvailabilityRequestProductsInner.md) |  | [optional] 
**AdditionalAttributes** | Pointer to [**[]PriceAndAvailabilityRequestAdditionalAttributesInner**](PriceAndAvailabilityRequestAdditionalAttributesInner.md) |  | [optional] 

## Methods

### NewPriceAndAvailabilityRequest

`func NewPriceAndAvailabilityRequest() *PriceAndAvailabilityRequest`

NewPriceAndAvailabilityRequest instantiates a new PriceAndAvailabilityRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceAndAvailabilityRequestWithDefaults

`func NewPriceAndAvailabilityRequestWithDefaults() *PriceAndAvailabilityRequest`

NewPriceAndAvailabilityRequestWithDefaults instantiates a new PriceAndAvailabilityRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShowAvailableDiscounts

`func (o *PriceAndAvailabilityRequest) GetShowAvailableDiscounts() bool`

GetShowAvailableDiscounts returns the ShowAvailableDiscounts field if non-nil, zero value otherwise.

### GetShowAvailableDiscountsOk

`func (o *PriceAndAvailabilityRequest) GetShowAvailableDiscountsOk() (*bool, bool)`

GetShowAvailableDiscountsOk returns a tuple with the ShowAvailableDiscounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowAvailableDiscounts

`func (o *PriceAndAvailabilityRequest) SetShowAvailableDiscounts(v bool)`

SetShowAvailableDiscounts sets ShowAvailableDiscounts field to given value.

### HasShowAvailableDiscounts

`func (o *PriceAndAvailabilityRequest) HasShowAvailableDiscounts() bool`

HasShowAvailableDiscounts returns a boolean if a field has been set.

### SetShowAvailableDiscountsNil

`func (o *PriceAndAvailabilityRequest) SetShowAvailableDiscountsNil(b bool)`

 SetShowAvailableDiscountsNil sets the value for ShowAvailableDiscounts to be an explicit nil

### UnsetShowAvailableDiscounts
`func (o *PriceAndAvailabilityRequest) UnsetShowAvailableDiscounts()`

UnsetShowAvailableDiscounts ensures that no value is present for ShowAvailableDiscounts, not even an explicit nil
### GetShowReserveInventoryDetails

`func (o *PriceAndAvailabilityRequest) GetShowReserveInventoryDetails() bool`

GetShowReserveInventoryDetails returns the ShowReserveInventoryDetails field if non-nil, zero value otherwise.

### GetShowReserveInventoryDetailsOk

`func (o *PriceAndAvailabilityRequest) GetShowReserveInventoryDetailsOk() (*bool, bool)`

GetShowReserveInventoryDetailsOk returns a tuple with the ShowReserveInventoryDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowReserveInventoryDetails

`func (o *PriceAndAvailabilityRequest) SetShowReserveInventoryDetails(v bool)`

SetShowReserveInventoryDetails sets ShowReserveInventoryDetails field to given value.

### HasShowReserveInventoryDetails

`func (o *PriceAndAvailabilityRequest) HasShowReserveInventoryDetails() bool`

HasShowReserveInventoryDetails returns a boolean if a field has been set.

### SetShowReserveInventoryDetailsNil

`func (o *PriceAndAvailabilityRequest) SetShowReserveInventoryDetailsNil(b bool)`

 SetShowReserveInventoryDetailsNil sets the value for ShowReserveInventoryDetails to be an explicit nil

### UnsetShowReserveInventoryDetails
`func (o *PriceAndAvailabilityRequest) UnsetShowReserveInventoryDetails()`

UnsetShowReserveInventoryDetails ensures that no value is present for ShowReserveInventoryDetails, not even an explicit nil
### GetSpecialBidNumber

`func (o *PriceAndAvailabilityRequest) GetSpecialBidNumber() string`

GetSpecialBidNumber returns the SpecialBidNumber field if non-nil, zero value otherwise.

### GetSpecialBidNumberOk

`func (o *PriceAndAvailabilityRequest) GetSpecialBidNumberOk() (*string, bool)`

GetSpecialBidNumberOk returns a tuple with the SpecialBidNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialBidNumber

`func (o *PriceAndAvailabilityRequest) SetSpecialBidNumber(v string)`

SetSpecialBidNumber sets SpecialBidNumber field to given value.

### HasSpecialBidNumber

`func (o *PriceAndAvailabilityRequest) HasSpecialBidNumber() bool`

HasSpecialBidNumber returns a boolean if a field has been set.

### SetSpecialBidNumberNil

`func (o *PriceAndAvailabilityRequest) SetSpecialBidNumberNil(b bool)`

 SetSpecialBidNumberNil sets the value for SpecialBidNumber to be an explicit nil

### UnsetSpecialBidNumber
`func (o *PriceAndAvailabilityRequest) UnsetSpecialBidNumber()`

UnsetSpecialBidNumber ensures that no value is present for SpecialBidNumber, not even an explicit nil
### GetAvailabilityByWarehouse

`func (o *PriceAndAvailabilityRequest) GetAvailabilityByWarehouse() []PriceAndAvailabilityRequestAvailabilityByWarehouseInner`

GetAvailabilityByWarehouse returns the AvailabilityByWarehouse field if non-nil, zero value otherwise.

### GetAvailabilityByWarehouseOk

`func (o *PriceAndAvailabilityRequest) GetAvailabilityByWarehouseOk() (*[]PriceAndAvailabilityRequestAvailabilityByWarehouseInner, bool)`

GetAvailabilityByWarehouseOk returns a tuple with the AvailabilityByWarehouse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityByWarehouse

`func (o *PriceAndAvailabilityRequest) SetAvailabilityByWarehouse(v []PriceAndAvailabilityRequestAvailabilityByWarehouseInner)`

SetAvailabilityByWarehouse sets AvailabilityByWarehouse field to given value.

### HasAvailabilityByWarehouse

`func (o *PriceAndAvailabilityRequest) HasAvailabilityByWarehouse() bool`

HasAvailabilityByWarehouse returns a boolean if a field has been set.

### SetAvailabilityByWarehouseNil

`func (o *PriceAndAvailabilityRequest) SetAvailabilityByWarehouseNil(b bool)`

 SetAvailabilityByWarehouseNil sets the value for AvailabilityByWarehouse to be an explicit nil

### UnsetAvailabilityByWarehouse
`func (o *PriceAndAvailabilityRequest) UnsetAvailabilityByWarehouse()`

UnsetAvailabilityByWarehouse ensures that no value is present for AvailabilityByWarehouse, not even an explicit nil
### GetProducts

`func (o *PriceAndAvailabilityRequest) GetProducts() []PriceAndAvailabilityRequestProductsInner`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *PriceAndAvailabilityRequest) GetProductsOk() (*[]PriceAndAvailabilityRequestProductsInner, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *PriceAndAvailabilityRequest) SetProducts(v []PriceAndAvailabilityRequestProductsInner)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *PriceAndAvailabilityRequest) HasProducts() bool`

HasProducts returns a boolean if a field has been set.

### GetAdditionalAttributes

`func (o *PriceAndAvailabilityRequest) GetAdditionalAttributes() []PriceAndAvailabilityRequestAdditionalAttributesInner`

GetAdditionalAttributes returns the AdditionalAttributes field if non-nil, zero value otherwise.

### GetAdditionalAttributesOk

`func (o *PriceAndAvailabilityRequest) GetAdditionalAttributesOk() (*[]PriceAndAvailabilityRequestAdditionalAttributesInner, bool)`

GetAdditionalAttributesOk returns a tuple with the AdditionalAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalAttributes

`func (o *PriceAndAvailabilityRequest) SetAdditionalAttributes(v []PriceAndAvailabilityRequestAdditionalAttributesInner)`

SetAdditionalAttributes sets AdditionalAttributes field to given value.

### HasAdditionalAttributes

`func (o *PriceAndAvailabilityRequest) HasAdditionalAttributes() bool`

HasAdditionalAttributes returns a boolean if a field has been set.

### SetAdditionalAttributesNil

`func (o *PriceAndAvailabilityRequest) SetAdditionalAttributesNil(b bool)`

 SetAdditionalAttributesNil sets the value for AdditionalAttributes to be an explicit nil

### UnsetAdditionalAttributes
`func (o *PriceAndAvailabilityRequest) UnsetAdditionalAttributes()`

UnsetAdditionalAttributes ensures that no value is present for AdditionalAttributes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


