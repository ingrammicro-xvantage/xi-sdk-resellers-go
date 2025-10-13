# QuoteDetailsResponseProductsInnerPrice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QuotePrice** | Pointer to **NullableFloat32** | Ingram Micro quoted price specific to the reseller and quote. | [optional] 
**Msrp** | Pointer to **NullableFloat32** | Manufacturer Suggested Retail Price | [optional] 
**ExtendedMsrp** | Pointer to **NullableFloat32** | Extended MSRP - Manufacturer Suggested Retail Price X Quantity | [optional] 
**ExtendedQuotePrice** | Pointer to **NullableFloat32** | Extended reseller quoted price (cost to reseller) X Quantity | [optional] 
**RemainingQuantityExtendedMsrp** | Pointer to **NullableFloat32** |  | [optional] 
**RemainingQuantityExtendedQuotePrice** | Pointer to **NullableFloat32** |  | [optional] 
**DiscountOffList** | Pointer to **NullableString** | Discount off list percentage extended | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**RecurringPriceModel** | Pointer to **string** |  | [optional] 
**UnitOfMeasure** | Pointer to **string** |  | [optional] 
**Tax** | Pointer to **NullableFloat32** |  | [optional] 
**Extrafees** | Pointer to **NullableFloat32** |  | [optional] 
**ExtraFeesDetails** | Pointer to [**[]QuoteDetailsResponseProductsInnerPriceExtraFeesDetailsInner**](QuoteDetailsResponseProductsInnerPriceExtraFeesDetailsInner.md) |  | [optional] 
**Discounts** | Pointer to [**[]QuoteDetailsResponseProductsInnerPriceDiscountsInner**](QuoteDetailsResponseProductsInnerPriceDiscountsInner.md) |  | [optional] 

## Methods

### NewQuoteDetailsResponseProductsInnerPrice

`func NewQuoteDetailsResponseProductsInnerPrice() *QuoteDetailsResponseProductsInnerPrice`

NewQuoteDetailsResponseProductsInnerPrice instantiates a new QuoteDetailsResponseProductsInnerPrice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuoteDetailsResponseProductsInnerPriceWithDefaults

`func NewQuoteDetailsResponseProductsInnerPriceWithDefaults() *QuoteDetailsResponseProductsInnerPrice`

NewQuoteDetailsResponseProductsInnerPriceWithDefaults instantiates a new QuoteDetailsResponseProductsInnerPrice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuotePrice

`func (o *QuoteDetailsResponseProductsInnerPrice) GetQuotePrice() float32`

GetQuotePrice returns the QuotePrice field if non-nil, zero value otherwise.

### GetQuotePriceOk

`func (o *QuoteDetailsResponseProductsInnerPrice) GetQuotePriceOk() (*float32, bool)`

GetQuotePriceOk returns a tuple with the QuotePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotePrice

`func (o *QuoteDetailsResponseProductsInnerPrice) SetQuotePrice(v float32)`

SetQuotePrice sets QuotePrice field to given value.

### HasQuotePrice

`func (o *QuoteDetailsResponseProductsInnerPrice) HasQuotePrice() bool`

HasQuotePrice returns a boolean if a field has been set.

### SetQuotePriceNil

`func (o *QuoteDetailsResponseProductsInnerPrice) SetQuotePriceNil(b bool)`

 SetQuotePriceNil sets the value for QuotePrice to be an explicit nil

### UnsetQuotePrice
`func (o *QuoteDetailsResponseProductsInnerPrice) UnsetQuotePrice()`

UnsetQuotePrice ensures that no value is present for QuotePrice, not even an explicit nil
### GetMsrp

`func (o *QuoteDetailsResponseProductsInnerPrice) GetMsrp() float32`

GetMsrp returns the Msrp field if non-nil, zero value otherwise.

### GetMsrpOk

`func (o *QuoteDetailsResponseProductsInnerPrice) GetMsrpOk() (*float32, bool)`

GetMsrpOk returns a tuple with the Msrp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsrp

`func (o *QuoteDetailsResponseProductsInnerPrice) SetMsrp(v float32)`

SetMsrp sets Msrp field to given value.

### HasMsrp

`func (o *QuoteDetailsResponseProductsInnerPrice) HasMsrp() bool`

HasMsrp returns a boolean if a field has been set.

### SetMsrpNil

`func (o *QuoteDetailsResponseProductsInnerPrice) SetMsrpNil(b bool)`

 SetMsrpNil sets the value for Msrp to be an explicit nil

### UnsetMsrp
`func (o *QuoteDetailsResponseProductsInnerPrice) UnsetMsrp()`

UnsetMsrp ensures that no value is present for Msrp, not even an explicit nil
### GetExtendedMsrp

`func (o *QuoteDetailsResponseProductsInnerPrice) GetExtendedMsrp() float32`

GetExtendedMsrp returns the ExtendedMsrp field if non-nil, zero value otherwise.

### GetExtendedMsrpOk

`func (o *QuoteDetailsResponseProductsInnerPrice) GetExtendedMsrpOk() (*float32, bool)`

GetExtendedMsrpOk returns a tuple with the ExtendedMsrp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedMsrp

`func (o *QuoteDetailsResponseProductsInnerPrice) SetExtendedMsrp(v float32)`

SetExtendedMsrp sets ExtendedMsrp field to given value.

### HasExtendedMsrp

`func (o *QuoteDetailsResponseProductsInnerPrice) HasExtendedMsrp() bool`

HasExtendedMsrp returns a boolean if a field has been set.

### SetExtendedMsrpNil

`func (o *QuoteDetailsResponseProductsInnerPrice) SetExtendedMsrpNil(b bool)`

 SetExtendedMsrpNil sets the value for ExtendedMsrp to be an explicit nil

### UnsetExtendedMsrp
`func (o *QuoteDetailsResponseProductsInnerPrice) UnsetExtendedMsrp()`

UnsetExtendedMsrp ensures that no value is present for ExtendedMsrp, not even an explicit nil
### GetExtendedQuotePrice

`func (o *QuoteDetailsResponseProductsInnerPrice) GetExtendedQuotePrice() float32`

GetExtendedQuotePrice returns the ExtendedQuotePrice field if non-nil, zero value otherwise.

### GetExtendedQuotePriceOk

`func (o *QuoteDetailsResponseProductsInnerPrice) GetExtendedQuotePriceOk() (*float32, bool)`

GetExtendedQuotePriceOk returns a tuple with the ExtendedQuotePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedQuotePrice

`func (o *QuoteDetailsResponseProductsInnerPrice) SetExtendedQuotePrice(v float32)`

SetExtendedQuotePrice sets ExtendedQuotePrice field to given value.

### HasExtendedQuotePrice

`func (o *QuoteDetailsResponseProductsInnerPrice) HasExtendedQuotePrice() bool`

HasExtendedQuotePrice returns a boolean if a field has been set.

### SetExtendedQuotePriceNil

`func (o *QuoteDetailsResponseProductsInnerPrice) SetExtendedQuotePriceNil(b bool)`

 SetExtendedQuotePriceNil sets the value for ExtendedQuotePrice to be an explicit nil

### UnsetExtendedQuotePrice
`func (o *QuoteDetailsResponseProductsInnerPrice) UnsetExtendedQuotePrice()`

UnsetExtendedQuotePrice ensures that no value is present for ExtendedQuotePrice, not even an explicit nil
### GetRemainingQuantityExtendedMsrp

`func (o *QuoteDetailsResponseProductsInnerPrice) GetRemainingQuantityExtendedMsrp() float32`

GetRemainingQuantityExtendedMsrp returns the RemainingQuantityExtendedMsrp field if non-nil, zero value otherwise.

### GetRemainingQuantityExtendedMsrpOk

`func (o *QuoteDetailsResponseProductsInnerPrice) GetRemainingQuantityExtendedMsrpOk() (*float32, bool)`

GetRemainingQuantityExtendedMsrpOk returns a tuple with the RemainingQuantityExtendedMsrp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingQuantityExtendedMsrp

`func (o *QuoteDetailsResponseProductsInnerPrice) SetRemainingQuantityExtendedMsrp(v float32)`

SetRemainingQuantityExtendedMsrp sets RemainingQuantityExtendedMsrp field to given value.

### HasRemainingQuantityExtendedMsrp

`func (o *QuoteDetailsResponseProductsInnerPrice) HasRemainingQuantityExtendedMsrp() bool`

HasRemainingQuantityExtendedMsrp returns a boolean if a field has been set.

### SetRemainingQuantityExtendedMsrpNil

`func (o *QuoteDetailsResponseProductsInnerPrice) SetRemainingQuantityExtendedMsrpNil(b bool)`

 SetRemainingQuantityExtendedMsrpNil sets the value for RemainingQuantityExtendedMsrp to be an explicit nil

### UnsetRemainingQuantityExtendedMsrp
`func (o *QuoteDetailsResponseProductsInnerPrice) UnsetRemainingQuantityExtendedMsrp()`

UnsetRemainingQuantityExtendedMsrp ensures that no value is present for RemainingQuantityExtendedMsrp, not even an explicit nil
### GetRemainingQuantityExtendedQuotePrice

`func (o *QuoteDetailsResponseProductsInnerPrice) GetRemainingQuantityExtendedQuotePrice() float32`

GetRemainingQuantityExtendedQuotePrice returns the RemainingQuantityExtendedQuotePrice field if non-nil, zero value otherwise.

### GetRemainingQuantityExtendedQuotePriceOk

`func (o *QuoteDetailsResponseProductsInnerPrice) GetRemainingQuantityExtendedQuotePriceOk() (*float32, bool)`

GetRemainingQuantityExtendedQuotePriceOk returns a tuple with the RemainingQuantityExtendedQuotePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingQuantityExtendedQuotePrice

`func (o *QuoteDetailsResponseProductsInnerPrice) SetRemainingQuantityExtendedQuotePrice(v float32)`

SetRemainingQuantityExtendedQuotePrice sets RemainingQuantityExtendedQuotePrice field to given value.

### HasRemainingQuantityExtendedQuotePrice

`func (o *QuoteDetailsResponseProductsInnerPrice) HasRemainingQuantityExtendedQuotePrice() bool`

HasRemainingQuantityExtendedQuotePrice returns a boolean if a field has been set.

### SetRemainingQuantityExtendedQuotePriceNil

`func (o *QuoteDetailsResponseProductsInnerPrice) SetRemainingQuantityExtendedQuotePriceNil(b bool)`

 SetRemainingQuantityExtendedQuotePriceNil sets the value for RemainingQuantityExtendedQuotePrice to be an explicit nil

### UnsetRemainingQuantityExtendedQuotePrice
`func (o *QuoteDetailsResponseProductsInnerPrice) UnsetRemainingQuantityExtendedQuotePrice()`

UnsetRemainingQuantityExtendedQuotePrice ensures that no value is present for RemainingQuantityExtendedQuotePrice, not even an explicit nil
### GetDiscountOffList

`func (o *QuoteDetailsResponseProductsInnerPrice) GetDiscountOffList() string`

GetDiscountOffList returns the DiscountOffList field if non-nil, zero value otherwise.

### GetDiscountOffListOk

`func (o *QuoteDetailsResponseProductsInnerPrice) GetDiscountOffListOk() (*string, bool)`

GetDiscountOffListOk returns a tuple with the DiscountOffList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountOffList

`func (o *QuoteDetailsResponseProductsInnerPrice) SetDiscountOffList(v string)`

SetDiscountOffList sets DiscountOffList field to given value.

### HasDiscountOffList

`func (o *QuoteDetailsResponseProductsInnerPrice) HasDiscountOffList() bool`

HasDiscountOffList returns a boolean if a field has been set.

### SetDiscountOffListNil

`func (o *QuoteDetailsResponseProductsInnerPrice) SetDiscountOffListNil(b bool)`

 SetDiscountOffListNil sets the value for DiscountOffList to be an explicit nil

### UnsetDiscountOffList
`func (o *QuoteDetailsResponseProductsInnerPrice) UnsetDiscountOffList()`

UnsetDiscountOffList ensures that no value is present for DiscountOffList, not even an explicit nil
### GetType

`func (o *QuoteDetailsResponseProductsInnerPrice) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *QuoteDetailsResponseProductsInnerPrice) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *QuoteDetailsResponseProductsInnerPrice) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *QuoteDetailsResponseProductsInnerPrice) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRecurringPriceModel

`func (o *QuoteDetailsResponseProductsInnerPrice) GetRecurringPriceModel() string`

GetRecurringPriceModel returns the RecurringPriceModel field if non-nil, zero value otherwise.

### GetRecurringPriceModelOk

`func (o *QuoteDetailsResponseProductsInnerPrice) GetRecurringPriceModelOk() (*string, bool)`

GetRecurringPriceModelOk returns a tuple with the RecurringPriceModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringPriceModel

`func (o *QuoteDetailsResponseProductsInnerPrice) SetRecurringPriceModel(v string)`

SetRecurringPriceModel sets RecurringPriceModel field to given value.

### HasRecurringPriceModel

`func (o *QuoteDetailsResponseProductsInnerPrice) HasRecurringPriceModel() bool`

HasRecurringPriceModel returns a boolean if a field has been set.

### GetUnitOfMeasure

`func (o *QuoteDetailsResponseProductsInnerPrice) GetUnitOfMeasure() string`

GetUnitOfMeasure returns the UnitOfMeasure field if non-nil, zero value otherwise.

### GetUnitOfMeasureOk

`func (o *QuoteDetailsResponseProductsInnerPrice) GetUnitOfMeasureOk() (*string, bool)`

GetUnitOfMeasureOk returns a tuple with the UnitOfMeasure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitOfMeasure

`func (o *QuoteDetailsResponseProductsInnerPrice) SetUnitOfMeasure(v string)`

SetUnitOfMeasure sets UnitOfMeasure field to given value.

### HasUnitOfMeasure

`func (o *QuoteDetailsResponseProductsInnerPrice) HasUnitOfMeasure() bool`

HasUnitOfMeasure returns a boolean if a field has been set.

### GetTax

`func (o *QuoteDetailsResponseProductsInnerPrice) GetTax() float32`

GetTax returns the Tax field if non-nil, zero value otherwise.

### GetTaxOk

`func (o *QuoteDetailsResponseProductsInnerPrice) GetTaxOk() (*float32, bool)`

GetTaxOk returns a tuple with the Tax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTax

`func (o *QuoteDetailsResponseProductsInnerPrice) SetTax(v float32)`

SetTax sets Tax field to given value.

### HasTax

`func (o *QuoteDetailsResponseProductsInnerPrice) HasTax() bool`

HasTax returns a boolean if a field has been set.

### SetTaxNil

`func (o *QuoteDetailsResponseProductsInnerPrice) SetTaxNil(b bool)`

 SetTaxNil sets the value for Tax to be an explicit nil

### UnsetTax
`func (o *QuoteDetailsResponseProductsInnerPrice) UnsetTax()`

UnsetTax ensures that no value is present for Tax, not even an explicit nil
### GetExtrafees

`func (o *QuoteDetailsResponseProductsInnerPrice) GetExtrafees() float32`

GetExtrafees returns the Extrafees field if non-nil, zero value otherwise.

### GetExtrafeesOk

`func (o *QuoteDetailsResponseProductsInnerPrice) GetExtrafeesOk() (*float32, bool)`

GetExtrafeesOk returns a tuple with the Extrafees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtrafees

`func (o *QuoteDetailsResponseProductsInnerPrice) SetExtrafees(v float32)`

SetExtrafees sets Extrafees field to given value.

### HasExtrafees

`func (o *QuoteDetailsResponseProductsInnerPrice) HasExtrafees() bool`

HasExtrafees returns a boolean if a field has been set.

### SetExtrafeesNil

`func (o *QuoteDetailsResponseProductsInnerPrice) SetExtrafeesNil(b bool)`

 SetExtrafeesNil sets the value for Extrafees to be an explicit nil

### UnsetExtrafees
`func (o *QuoteDetailsResponseProductsInnerPrice) UnsetExtrafees()`

UnsetExtrafees ensures that no value is present for Extrafees, not even an explicit nil
### GetExtraFeesDetails

`func (o *QuoteDetailsResponseProductsInnerPrice) GetExtraFeesDetails() []QuoteDetailsResponseProductsInnerPriceExtraFeesDetailsInner`

GetExtraFeesDetails returns the ExtraFeesDetails field if non-nil, zero value otherwise.

### GetExtraFeesDetailsOk

`func (o *QuoteDetailsResponseProductsInnerPrice) GetExtraFeesDetailsOk() (*[]QuoteDetailsResponseProductsInnerPriceExtraFeesDetailsInner, bool)`

GetExtraFeesDetailsOk returns a tuple with the ExtraFeesDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraFeesDetails

`func (o *QuoteDetailsResponseProductsInnerPrice) SetExtraFeesDetails(v []QuoteDetailsResponseProductsInnerPriceExtraFeesDetailsInner)`

SetExtraFeesDetails sets ExtraFeesDetails field to given value.

### HasExtraFeesDetails

`func (o *QuoteDetailsResponseProductsInnerPrice) HasExtraFeesDetails() bool`

HasExtraFeesDetails returns a boolean if a field has been set.

### GetDiscounts

`func (o *QuoteDetailsResponseProductsInnerPrice) GetDiscounts() []QuoteDetailsResponseProductsInnerPriceDiscountsInner`

GetDiscounts returns the Discounts field if non-nil, zero value otherwise.

### GetDiscountsOk

`func (o *QuoteDetailsResponseProductsInnerPrice) GetDiscountsOk() (*[]QuoteDetailsResponseProductsInnerPriceDiscountsInner, bool)`

GetDiscountsOk returns a tuple with the Discounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscounts

`func (o *QuoteDetailsResponseProductsInnerPrice) SetDiscounts(v []QuoteDetailsResponseProductsInnerPriceDiscountsInner)`

SetDiscounts sets Discounts field to given value.

### HasDiscounts

`func (o *QuoteDetailsResponseProductsInnerPrice) HasDiscounts() bool`

HasDiscounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


