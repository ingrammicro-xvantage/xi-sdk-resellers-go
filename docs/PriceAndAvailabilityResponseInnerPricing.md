# PriceAndAvailabilityResponseInnerPricing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrencyCode** | Pointer to **string** | The 3-digit ISO currency code. | [optional] 
**RetailPrice** | Pointer to **NullableFloat32** | The retail price of the product. | [optional] 
**MapPrice** | Pointer to **NullableFloat32** | Minimum Advertised Price (MAP). If required by the vendor, resellers can not sell below MAP price. | [optional] 
**CustomerPrice** | Pointer to **NullableFloat32** | The price customer pays after all special pricing and discounts have been applied. | [optional] 
**SpecialBidPricingAvailable** | Pointer to **NullableBool** | Boolean values specifies whether special Bid discounts are available for the product. | [optional] 
**WebDiscountsAvailable** | Pointer to **NullableBool** | Boolean values specifies whether web Discounts are available for the product. | [optional] 

## Methods

### NewPriceAndAvailabilityResponseInnerPricing

`func NewPriceAndAvailabilityResponseInnerPricing() *PriceAndAvailabilityResponseInnerPricing`

NewPriceAndAvailabilityResponseInnerPricing instantiates a new PriceAndAvailabilityResponseInnerPricing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceAndAvailabilityResponseInnerPricingWithDefaults

`func NewPriceAndAvailabilityResponseInnerPricingWithDefaults() *PriceAndAvailabilityResponseInnerPricing`

NewPriceAndAvailabilityResponseInnerPricingWithDefaults instantiates a new PriceAndAvailabilityResponseInnerPricing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrencyCode

`func (o *PriceAndAvailabilityResponseInnerPricing) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *PriceAndAvailabilityResponseInnerPricing) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *PriceAndAvailabilityResponseInnerPricing) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *PriceAndAvailabilityResponseInnerPricing) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetRetailPrice

`func (o *PriceAndAvailabilityResponseInnerPricing) GetRetailPrice() float32`

GetRetailPrice returns the RetailPrice field if non-nil, zero value otherwise.

### GetRetailPriceOk

`func (o *PriceAndAvailabilityResponseInnerPricing) GetRetailPriceOk() (*float32, bool)`

GetRetailPriceOk returns a tuple with the RetailPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetailPrice

`func (o *PriceAndAvailabilityResponseInnerPricing) SetRetailPrice(v float32)`

SetRetailPrice sets RetailPrice field to given value.

### HasRetailPrice

`func (o *PriceAndAvailabilityResponseInnerPricing) HasRetailPrice() bool`

HasRetailPrice returns a boolean if a field has been set.

### SetRetailPriceNil

`func (o *PriceAndAvailabilityResponseInnerPricing) SetRetailPriceNil(b bool)`

 SetRetailPriceNil sets the value for RetailPrice to be an explicit nil

### UnsetRetailPrice
`func (o *PriceAndAvailabilityResponseInnerPricing) UnsetRetailPrice()`

UnsetRetailPrice ensures that no value is present for RetailPrice, not even an explicit nil
### GetMapPrice

`func (o *PriceAndAvailabilityResponseInnerPricing) GetMapPrice() float32`

GetMapPrice returns the MapPrice field if non-nil, zero value otherwise.

### GetMapPriceOk

`func (o *PriceAndAvailabilityResponseInnerPricing) GetMapPriceOk() (*float32, bool)`

GetMapPriceOk returns a tuple with the MapPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapPrice

`func (o *PriceAndAvailabilityResponseInnerPricing) SetMapPrice(v float32)`

SetMapPrice sets MapPrice field to given value.

### HasMapPrice

`func (o *PriceAndAvailabilityResponseInnerPricing) HasMapPrice() bool`

HasMapPrice returns a boolean if a field has been set.

### SetMapPriceNil

`func (o *PriceAndAvailabilityResponseInnerPricing) SetMapPriceNil(b bool)`

 SetMapPriceNil sets the value for MapPrice to be an explicit nil

### UnsetMapPrice
`func (o *PriceAndAvailabilityResponseInnerPricing) UnsetMapPrice()`

UnsetMapPrice ensures that no value is present for MapPrice, not even an explicit nil
### GetCustomerPrice

`func (o *PriceAndAvailabilityResponseInnerPricing) GetCustomerPrice() float32`

GetCustomerPrice returns the CustomerPrice field if non-nil, zero value otherwise.

### GetCustomerPriceOk

`func (o *PriceAndAvailabilityResponseInnerPricing) GetCustomerPriceOk() (*float32, bool)`

GetCustomerPriceOk returns a tuple with the CustomerPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerPrice

`func (o *PriceAndAvailabilityResponseInnerPricing) SetCustomerPrice(v float32)`

SetCustomerPrice sets CustomerPrice field to given value.

### HasCustomerPrice

`func (o *PriceAndAvailabilityResponseInnerPricing) HasCustomerPrice() bool`

HasCustomerPrice returns a boolean if a field has been set.

### SetCustomerPriceNil

`func (o *PriceAndAvailabilityResponseInnerPricing) SetCustomerPriceNil(b bool)`

 SetCustomerPriceNil sets the value for CustomerPrice to be an explicit nil

### UnsetCustomerPrice
`func (o *PriceAndAvailabilityResponseInnerPricing) UnsetCustomerPrice()`

UnsetCustomerPrice ensures that no value is present for CustomerPrice, not even an explicit nil
### GetSpecialBidPricingAvailable

`func (o *PriceAndAvailabilityResponseInnerPricing) GetSpecialBidPricingAvailable() bool`

GetSpecialBidPricingAvailable returns the SpecialBidPricingAvailable field if non-nil, zero value otherwise.

### GetSpecialBidPricingAvailableOk

`func (o *PriceAndAvailabilityResponseInnerPricing) GetSpecialBidPricingAvailableOk() (*bool, bool)`

GetSpecialBidPricingAvailableOk returns a tuple with the SpecialBidPricingAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialBidPricingAvailable

`func (o *PriceAndAvailabilityResponseInnerPricing) SetSpecialBidPricingAvailable(v bool)`

SetSpecialBidPricingAvailable sets SpecialBidPricingAvailable field to given value.

### HasSpecialBidPricingAvailable

`func (o *PriceAndAvailabilityResponseInnerPricing) HasSpecialBidPricingAvailable() bool`

HasSpecialBidPricingAvailable returns a boolean if a field has been set.

### SetSpecialBidPricingAvailableNil

`func (o *PriceAndAvailabilityResponseInnerPricing) SetSpecialBidPricingAvailableNil(b bool)`

 SetSpecialBidPricingAvailableNil sets the value for SpecialBidPricingAvailable to be an explicit nil

### UnsetSpecialBidPricingAvailable
`func (o *PriceAndAvailabilityResponseInnerPricing) UnsetSpecialBidPricingAvailable()`

UnsetSpecialBidPricingAvailable ensures that no value is present for SpecialBidPricingAvailable, not even an explicit nil
### GetWebDiscountsAvailable

`func (o *PriceAndAvailabilityResponseInnerPricing) GetWebDiscountsAvailable() bool`

GetWebDiscountsAvailable returns the WebDiscountsAvailable field if non-nil, zero value otherwise.

### GetWebDiscountsAvailableOk

`func (o *PriceAndAvailabilityResponseInnerPricing) GetWebDiscountsAvailableOk() (*bool, bool)`

GetWebDiscountsAvailableOk returns a tuple with the WebDiscountsAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebDiscountsAvailable

`func (o *PriceAndAvailabilityResponseInnerPricing) SetWebDiscountsAvailable(v bool)`

SetWebDiscountsAvailable sets WebDiscountsAvailable field to given value.

### HasWebDiscountsAvailable

`func (o *PriceAndAvailabilityResponseInnerPricing) HasWebDiscountsAvailable() bool`

HasWebDiscountsAvailable returns a boolean if a field has been set.

### SetWebDiscountsAvailableNil

`func (o *PriceAndAvailabilityResponseInnerPricing) SetWebDiscountsAvailableNil(b bool)`

 SetWebDiscountsAvailableNil sets the value for WebDiscountsAvailable to be an explicit nil

### UnsetWebDiscountsAvailable
`func (o *PriceAndAvailabilityResponseInnerPricing) UnsetWebDiscountsAvailable()`

UnsetWebDiscountsAvailable ensures that no value is present for WebDiscountsAvailable, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


