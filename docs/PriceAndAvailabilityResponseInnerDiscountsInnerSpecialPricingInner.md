# PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiscountType** | Pointer to **NullableString** | The type of discount being given to the customer. | [optional] 
**SpecialBidNumer** | Pointer to **NullableString** | Pre-approved special pricing/bid number provided to the reseller by the vendor for special pricing and discounts. Used to track the bid number where different line items have different bid numbers. Line-level bid numbers take precedence over header-level bid numbers. | [optional] 
**SpecialPricingDiscount** | Pointer to **NullableFloat32** | Special pricing discount amount given to the customer. | [optional] 
**SpecialPricingEffectiveDate** | Pointer to **NullableString** | The effective date of the special pricing available to the customer. | [optional] 
**SpecialPricingExpirationDate** | Pointer to **NullableString** | The expiration date of the special pricing available to the customer. | [optional] 
**SpecialPricingAvailableQuantity** | Pointer to **NullableInt32** | The available quantity of products with discounts. | [optional] 
**SpecialPricingMinQuantity** | Pointer to **NullableInt32** | The minimum quantity of products that have to be purchased to ensure the discount is applied. | [optional] 
**GovernmentDiscountType** | Pointer to **NullableString** | Type of Government Discount. *Currently, this discount is only available in the USA. | [optional] 
**GovernmentDiscountedCustomerPrice** | Pointer to **NullableFloat32** | Government Discounted Customer Price. *Currently, this discount is only available in the USA. | [optional] 

## Methods

### NewPriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner

`func NewPriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner() *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner`

NewPriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner instantiates a new PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInnerWithDefaults

`func NewPriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInnerWithDefaults() *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner`

NewPriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInnerWithDefaults instantiates a new PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiscountType

`func (o *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) GetDiscountType() string`

GetDiscountType returns the DiscountType field if non-nil, zero value otherwise.

### GetDiscountTypeOk

`func (o *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) GetDiscountTypeOk() (*string, bool)`

GetDiscountTypeOk returns a tuple with the DiscountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountType

`func (o *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) SetDiscountType(v string)`

SetDiscountType sets DiscountType field to given value.

### HasDiscountType

`func (o *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) HasDiscountType() bool`

HasDiscountType returns a boolean if a field has been set.

### SetDiscountTypeNil

`func (o *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) SetDiscountTypeNil(b bool)`

 SetDiscountTypeNil sets the value for DiscountType to be an explicit nil

### UnsetDiscountType
`func (o *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) UnsetDiscountType()`

UnsetDiscountType ensures that no value is present for DiscountType, not even an explicit nil
### GetSpecialBidNumer

`func (o *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) GetSpecialBidNumer() string`

GetSpecialBidNumer returns the SpecialBidNumer field if non-nil, zero value otherwise.

### GetSpecialBidNumerOk

`func (o *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) GetSpecialBidNumerOk() (*string, bool)`

GetSpecialBidNumerOk returns a tuple with the SpecialBidNumer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialBidNumer

`func (o *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) SetSpecialBidNumer(v string)`

SetSpecialBidNumer sets SpecialBidNumer field to given value.

### HasSpecialBidNumer

`func (o *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) HasSpecialBidNumer() bool`

HasSpecialBidNumer returns a boolean if a field has been set.

### SetSpecialBidNumerNil

`func (o *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) SetSpecialBidNumerNil(b bool)`

 SetSpecialBidNumerNil sets the value for SpecialBidNumer to be an explicit nil

### UnsetSpecialBidNumer
`func (o *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) UnsetSpecialBidNumer()`

UnsetSpecialBidNumer ensures that no value is present for SpecialBidNumer, not even an explicit nil
### GetSpecialPricingDiscount

`func (o *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) GetSpecialPricingDiscount() float32`

GetSpecialPricingDiscount returns the SpecialPricingDiscount field if non-nil, zero value otherwise.

### GetSpecialPricingDiscountOk

`func (o *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) GetSpecialPricingDiscountOk() (*float32, bool)`

GetSpecialPricingDiscountOk returns a tuple with the SpecialPricingDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialPricingDiscount

`func (o *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) SetSpecialPricingDiscount(v float32)`

SetSpecialPricingDiscount sets SpecialPricingDiscount field to given value.

### HasSpecialPricingDiscount

`func (o *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) HasSpecialPricingDiscount() bool`

HasSpecialPricingDiscount returns a boolean if a field has been set.

### SetSpecialPricingDiscountNil

`func (o *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) SetSpecialPricingDiscountNil(b bool)`

 SetSpecialPricingDiscountNil sets the value for SpecialPricingDiscount to be an explicit nil

### UnsetSpecialPricingDiscount
`func (o *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) UnsetSpecialPricingDiscount()`

UnsetSpecialPricingDiscount ensures that no value is present for SpecialPricingDiscount, not even an explicit nil
### GetSpecialPricingEffectiveDate

`func (o *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) GetSpecialPricingEffectiveDate() string`

GetSpecialPricingEffectiveDate returns the SpecialPricingEffectiveDate field if non-nil, zero value otherwise.

### GetSpecialPricingEffectiveDateOk

`func (o *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) GetSpecialPricingEffectiveDateOk() (*string, bool)`

GetSpecialPricingEffectiveDateOk returns a tuple with the SpecialPricingEffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialPricingEffectiveDate

`func (o *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) SetSpecialPricingEffectiveDate(v string)`

SetSpecialPricingEffectiveDate sets SpecialPricingEffectiveDate field to given value.

### HasSpecialPricingEffectiveDate

`func (o *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) HasSpecialPricingEffectiveDate() bool`

HasSpecialPricingEffectiveDate returns a boolean if a field has been set.

### SetSpecialPricingEffectiveDateNil

`func (o *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) SetSpecialPricingEffectiveDateNil(b bool)`

 SetSpecialPricingEffectiveDateNil sets the value for SpecialPricingEffectiveDate to be an explicit nil

### UnsetSpecialPricingEffectiveDate
`func (o *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) UnsetSpecialPricingEffectiveDate()`

UnsetSpecialPricingEffectiveDate ensures that no value is present for SpecialPricingEffectiveDate, not even an explicit nil
### GetSpecialPricingExpirationDate

`func (o *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) GetSpecialPricingExpirationDate() string`

GetSpecialPricingExpirationDate returns the SpecialPricingExpirationDate field if non-nil, zero value otherwise.

### GetSpecialPricingExpirationDateOk

`func (o *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) GetSpecialPricingExpirationDateOk() (*string, bool)`

GetSpecialPricingExpirationDateOk returns a tuple with the SpecialPricingExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialPricingExpirationDate

`func (o *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) SetSpecialPricingExpirationDate(v string)`

SetSpecialPricingExpirationDate sets SpecialPricingExpirationDate field to given value.

### HasSpecialPricingExpirationDate

`func (o *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) HasSpecialPricingExpirationDate() bool`

HasSpecialPricingExpirationDate returns a boolean if a field has been set.

### SetSpecialPricingExpirationDateNil

`func (o *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) SetSpecialPricingExpirationDateNil(b bool)`

 SetSpecialPricingExpirationDateNil sets the value for SpecialPricingExpirationDate to be an explicit nil

### UnsetSpecialPricingExpirationDate
`func (o *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) UnsetSpecialPricingExpirationDate()`

UnsetSpecialPricingExpirationDate ensures that no value is present for SpecialPricingExpirationDate, not even an explicit nil
### GetSpecialPricingAvailableQuantity

`func (o *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) GetSpecialPricingAvailableQuantity() int32`

GetSpecialPricingAvailableQuantity returns the SpecialPricingAvailableQuantity field if non-nil, zero value otherwise.

### GetSpecialPricingAvailableQuantityOk

`func (o *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) GetSpecialPricingAvailableQuantityOk() (*int32, bool)`

GetSpecialPricingAvailableQuantityOk returns a tuple with the SpecialPricingAvailableQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialPricingAvailableQuantity

`func (o *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) SetSpecialPricingAvailableQuantity(v int32)`

SetSpecialPricingAvailableQuantity sets SpecialPricingAvailableQuantity field to given value.

### HasSpecialPricingAvailableQuantity

`func (o *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) HasSpecialPricingAvailableQuantity() bool`

HasSpecialPricingAvailableQuantity returns a boolean if a field has been set.

### SetSpecialPricingAvailableQuantityNil

`func (o *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) SetSpecialPricingAvailableQuantityNil(b bool)`

 SetSpecialPricingAvailableQuantityNil sets the value for SpecialPricingAvailableQuantity to be an explicit nil

### UnsetSpecialPricingAvailableQuantity
`func (o *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) UnsetSpecialPricingAvailableQuantity()`

UnsetSpecialPricingAvailableQuantity ensures that no value is present for SpecialPricingAvailableQuantity, not even an explicit nil
### GetSpecialPricingMinQuantity

`func (o *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) GetSpecialPricingMinQuantity() int32`

GetSpecialPricingMinQuantity returns the SpecialPricingMinQuantity field if non-nil, zero value otherwise.

### GetSpecialPricingMinQuantityOk

`func (o *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) GetSpecialPricingMinQuantityOk() (*int32, bool)`

GetSpecialPricingMinQuantityOk returns a tuple with the SpecialPricingMinQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialPricingMinQuantity

`func (o *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) SetSpecialPricingMinQuantity(v int32)`

SetSpecialPricingMinQuantity sets SpecialPricingMinQuantity field to given value.

### HasSpecialPricingMinQuantity

`func (o *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) HasSpecialPricingMinQuantity() bool`

HasSpecialPricingMinQuantity returns a boolean if a field has been set.

### SetSpecialPricingMinQuantityNil

`func (o *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) SetSpecialPricingMinQuantityNil(b bool)`

 SetSpecialPricingMinQuantityNil sets the value for SpecialPricingMinQuantity to be an explicit nil

### UnsetSpecialPricingMinQuantity
`func (o *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) UnsetSpecialPricingMinQuantity()`

UnsetSpecialPricingMinQuantity ensures that no value is present for SpecialPricingMinQuantity, not even an explicit nil
### GetGovernmentDiscountType

`func (o *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) GetGovernmentDiscountType() string`

GetGovernmentDiscountType returns the GovernmentDiscountType field if non-nil, zero value otherwise.

### GetGovernmentDiscountTypeOk

`func (o *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) GetGovernmentDiscountTypeOk() (*string, bool)`

GetGovernmentDiscountTypeOk returns a tuple with the GovernmentDiscountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGovernmentDiscountType

`func (o *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) SetGovernmentDiscountType(v string)`

SetGovernmentDiscountType sets GovernmentDiscountType field to given value.

### HasGovernmentDiscountType

`func (o *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) HasGovernmentDiscountType() bool`

HasGovernmentDiscountType returns a boolean if a field has been set.

### SetGovernmentDiscountTypeNil

`func (o *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) SetGovernmentDiscountTypeNil(b bool)`

 SetGovernmentDiscountTypeNil sets the value for GovernmentDiscountType to be an explicit nil

### UnsetGovernmentDiscountType
`func (o *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) UnsetGovernmentDiscountType()`

UnsetGovernmentDiscountType ensures that no value is present for GovernmentDiscountType, not even an explicit nil
### GetGovernmentDiscountedCustomerPrice

`func (o *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) GetGovernmentDiscountedCustomerPrice() float32`

GetGovernmentDiscountedCustomerPrice returns the GovernmentDiscountedCustomerPrice field if non-nil, zero value otherwise.

### GetGovernmentDiscountedCustomerPriceOk

`func (o *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) GetGovernmentDiscountedCustomerPriceOk() (*float32, bool)`

GetGovernmentDiscountedCustomerPriceOk returns a tuple with the GovernmentDiscountedCustomerPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGovernmentDiscountedCustomerPrice

`func (o *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) SetGovernmentDiscountedCustomerPrice(v float32)`

SetGovernmentDiscountedCustomerPrice sets GovernmentDiscountedCustomerPrice field to given value.

### HasGovernmentDiscountedCustomerPrice

`func (o *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) HasGovernmentDiscountedCustomerPrice() bool`

HasGovernmentDiscountedCustomerPrice returns a boolean if a field has been set.

### SetGovernmentDiscountedCustomerPriceNil

`func (o *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) SetGovernmentDiscountedCustomerPriceNil(b bool)`

 SetGovernmentDiscountedCustomerPriceNil sets the value for GovernmentDiscountedCustomerPrice to be an explicit nil

### UnsetGovernmentDiscountedCustomerPrice
`func (o *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) UnsetGovernmentDiscountedCustomerPrice()`

UnsetGovernmentDiscountedCustomerPrice ensures that no value is present for GovernmentDiscountedCustomerPrice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


