# QuoteCreateRequestProductsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerLineNumber** | Pointer to **string** | The reseller&#39;s line item number for reference in their system. | [optional] 
**IngramPartNumber** | Pointer to **string** | Ingram Micro SKU (stock keeping unit). An identification, usually alphanumeric, of a particular product that allows it to be tracked for inventory purposes | [optional] 
**VendorPartNumber** | Pointer to **string** | Vendor Part Number | [optional] 
**Quantity** | Pointer to **string** | Quantity of product line item quoted. | [optional] 
**SpecialBid** | Pointer to **string** | Special bid associated with product line item | [optional] 
**LineLevelNotes** | Pointer to **string** | Product line-item comments. | [optional] 
**PricingType** | Pointer to **string** | Pricing type of the quote | [optional] 

## Methods

### NewQuoteCreateRequestProductsInner

`func NewQuoteCreateRequestProductsInner() *QuoteCreateRequestProductsInner`

NewQuoteCreateRequestProductsInner instantiates a new QuoteCreateRequestProductsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuoteCreateRequestProductsInnerWithDefaults

`func NewQuoteCreateRequestProductsInnerWithDefaults() *QuoteCreateRequestProductsInner`

NewQuoteCreateRequestProductsInnerWithDefaults instantiates a new QuoteCreateRequestProductsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerLineNumber

`func (o *QuoteCreateRequestProductsInner) GetCustomerLineNumber() string`

GetCustomerLineNumber returns the CustomerLineNumber field if non-nil, zero value otherwise.

### GetCustomerLineNumberOk

`func (o *QuoteCreateRequestProductsInner) GetCustomerLineNumberOk() (*string, bool)`

GetCustomerLineNumberOk returns a tuple with the CustomerLineNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerLineNumber

`func (o *QuoteCreateRequestProductsInner) SetCustomerLineNumber(v string)`

SetCustomerLineNumber sets CustomerLineNumber field to given value.

### HasCustomerLineNumber

`func (o *QuoteCreateRequestProductsInner) HasCustomerLineNumber() bool`

HasCustomerLineNumber returns a boolean if a field has been set.

### GetIngramPartNumber

`func (o *QuoteCreateRequestProductsInner) GetIngramPartNumber() string`

GetIngramPartNumber returns the IngramPartNumber field if non-nil, zero value otherwise.

### GetIngramPartNumberOk

`func (o *QuoteCreateRequestProductsInner) GetIngramPartNumberOk() (*string, bool)`

GetIngramPartNumberOk returns a tuple with the IngramPartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngramPartNumber

`func (o *QuoteCreateRequestProductsInner) SetIngramPartNumber(v string)`

SetIngramPartNumber sets IngramPartNumber field to given value.

### HasIngramPartNumber

`func (o *QuoteCreateRequestProductsInner) HasIngramPartNumber() bool`

HasIngramPartNumber returns a boolean if a field has been set.

### GetVendorPartNumber

`func (o *QuoteCreateRequestProductsInner) GetVendorPartNumber() string`

GetVendorPartNumber returns the VendorPartNumber field if non-nil, zero value otherwise.

### GetVendorPartNumberOk

`func (o *QuoteCreateRequestProductsInner) GetVendorPartNumberOk() (*string, bool)`

GetVendorPartNumberOk returns a tuple with the VendorPartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorPartNumber

`func (o *QuoteCreateRequestProductsInner) SetVendorPartNumber(v string)`

SetVendorPartNumber sets VendorPartNumber field to given value.

### HasVendorPartNumber

`func (o *QuoteCreateRequestProductsInner) HasVendorPartNumber() bool`

HasVendorPartNumber returns a boolean if a field has been set.

### GetQuantity

`func (o *QuoteCreateRequestProductsInner) GetQuantity() string`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *QuoteCreateRequestProductsInner) GetQuantityOk() (*string, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *QuoteCreateRequestProductsInner) SetQuantity(v string)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *QuoteCreateRequestProductsInner) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetSpecialBid

`func (o *QuoteCreateRequestProductsInner) GetSpecialBid() string`

GetSpecialBid returns the SpecialBid field if non-nil, zero value otherwise.

### GetSpecialBidOk

`func (o *QuoteCreateRequestProductsInner) GetSpecialBidOk() (*string, bool)`

GetSpecialBidOk returns a tuple with the SpecialBid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialBid

`func (o *QuoteCreateRequestProductsInner) SetSpecialBid(v string)`

SetSpecialBid sets SpecialBid field to given value.

### HasSpecialBid

`func (o *QuoteCreateRequestProductsInner) HasSpecialBid() bool`

HasSpecialBid returns a boolean if a field has been set.

### GetLineLevelNotes

`func (o *QuoteCreateRequestProductsInner) GetLineLevelNotes() string`

GetLineLevelNotes returns the LineLevelNotes field if non-nil, zero value otherwise.

### GetLineLevelNotesOk

`func (o *QuoteCreateRequestProductsInner) GetLineLevelNotesOk() (*string, bool)`

GetLineLevelNotesOk returns a tuple with the LineLevelNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineLevelNotes

`func (o *QuoteCreateRequestProductsInner) SetLineLevelNotes(v string)`

SetLineLevelNotes sets LineLevelNotes field to given value.

### HasLineLevelNotes

`func (o *QuoteCreateRequestProductsInner) HasLineLevelNotes() bool`

HasLineLevelNotes returns a boolean if a field has been set.

### GetPricingType

`func (o *QuoteCreateRequestProductsInner) GetPricingType() string`

GetPricingType returns the PricingType field if non-nil, zero value otherwise.

### GetPricingTypeOk

`func (o *QuoteCreateRequestProductsInner) GetPricingTypeOk() (*string, bool)`

GetPricingTypeOk returns a tuple with the PricingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingType

`func (o *QuoteCreateRequestProductsInner) SetPricingType(v string)`

SetPricingType sets PricingType field to given value.

### HasPricingType

`func (o *QuoteCreateRequestProductsInner) HasPricingType() bool`

HasPricingType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


