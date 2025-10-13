# OrderCreateV7RequestLinesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerLineNumber** | Pointer to **string** | The reseller&#39;s line item number for reference in their system. The customer line number needs to be a unique numeric value between 1 and 884. In the event we receive duplicate values or alphanumeric values in the customer line number, we will re-sequence the customer line number. To prevent re-sequencing, please use a unique numeric value between 1 and 884 in the customer line number. | [optional] 
**IngramPartNumber** | Pointer to **NullableString** | The unique IngramMicro part number. | [optional] 
**VendorPartNumber** | Pointer to **NullableString** | The vendor&#39;s part number for the line item. | [optional] 
**Quantity** | Pointer to **int32** | The requested quantity of the line item. | [optional] 
**UnitPrice** | Pointer to **NullableFloat32** | The reseller-requested unit price for the line item. The unit price is not guaranteed. | [optional] 
**SpecialBidNumber** | Pointer to **NullableString** | The line-level bid number provided to the reseller by the vendor for special pricing and discounts. Used to track the bid number in the case of split orders or where different line items have different bid numbers. Line-level bid number take precedence over header-level bid numbers. | [optional] 
**EndUserPrice** | Pointer to **NullableFloat32** | The end-user price. Required for Export Orders. | [optional] 
**Notes** | Pointer to **NullableString** | The attribute field data. | [optional] 
**EndUserInfo** | Pointer to [**[]OrderCreateV7RequestLinesInnerEndUserInfoInner**](OrderCreateV7RequestLinesInnerEndUserInfoInner.md) |  | [optional] 
**AdditionalAttributes** | Pointer to [**[]OrderCreateV7RequestLinesInnerAdditionalAttributesInner**](OrderCreateV7RequestLinesInnerAdditionalAttributesInner.md) |  | [optional] 
**VmfAdditionalAttributesLines** | Pointer to [**[]OrderCreateV7RequestLinesInnerVmfAdditionalAttributesLinesInner**](OrderCreateV7RequestLinesInnerVmfAdditionalAttributesLinesInner.md) |  | [optional] 

## Methods

### NewOrderCreateV7RequestLinesInner

`func NewOrderCreateV7RequestLinesInner() *OrderCreateV7RequestLinesInner`

NewOrderCreateV7RequestLinesInner instantiates a new OrderCreateV7RequestLinesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderCreateV7RequestLinesInnerWithDefaults

`func NewOrderCreateV7RequestLinesInnerWithDefaults() *OrderCreateV7RequestLinesInner`

NewOrderCreateV7RequestLinesInnerWithDefaults instantiates a new OrderCreateV7RequestLinesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerLineNumber

`func (o *OrderCreateV7RequestLinesInner) GetCustomerLineNumber() string`

GetCustomerLineNumber returns the CustomerLineNumber field if non-nil, zero value otherwise.

### GetCustomerLineNumberOk

`func (o *OrderCreateV7RequestLinesInner) GetCustomerLineNumberOk() (*string, bool)`

GetCustomerLineNumberOk returns a tuple with the CustomerLineNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerLineNumber

`func (o *OrderCreateV7RequestLinesInner) SetCustomerLineNumber(v string)`

SetCustomerLineNumber sets CustomerLineNumber field to given value.

### HasCustomerLineNumber

`func (o *OrderCreateV7RequestLinesInner) HasCustomerLineNumber() bool`

HasCustomerLineNumber returns a boolean if a field has been set.

### GetIngramPartNumber

`func (o *OrderCreateV7RequestLinesInner) GetIngramPartNumber() string`

GetIngramPartNumber returns the IngramPartNumber field if non-nil, zero value otherwise.

### GetIngramPartNumberOk

`func (o *OrderCreateV7RequestLinesInner) GetIngramPartNumberOk() (*string, bool)`

GetIngramPartNumberOk returns a tuple with the IngramPartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngramPartNumber

`func (o *OrderCreateV7RequestLinesInner) SetIngramPartNumber(v string)`

SetIngramPartNumber sets IngramPartNumber field to given value.

### HasIngramPartNumber

`func (o *OrderCreateV7RequestLinesInner) HasIngramPartNumber() bool`

HasIngramPartNumber returns a boolean if a field has been set.

### SetIngramPartNumberNil

`func (o *OrderCreateV7RequestLinesInner) SetIngramPartNumberNil(b bool)`

 SetIngramPartNumberNil sets the value for IngramPartNumber to be an explicit nil

### UnsetIngramPartNumber
`func (o *OrderCreateV7RequestLinesInner) UnsetIngramPartNumber()`

UnsetIngramPartNumber ensures that no value is present for IngramPartNumber, not even an explicit nil
### GetVendorPartNumber

`func (o *OrderCreateV7RequestLinesInner) GetVendorPartNumber() string`

GetVendorPartNumber returns the VendorPartNumber field if non-nil, zero value otherwise.

### GetVendorPartNumberOk

`func (o *OrderCreateV7RequestLinesInner) GetVendorPartNumberOk() (*string, bool)`

GetVendorPartNumberOk returns a tuple with the VendorPartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorPartNumber

`func (o *OrderCreateV7RequestLinesInner) SetVendorPartNumber(v string)`

SetVendorPartNumber sets VendorPartNumber field to given value.

### HasVendorPartNumber

`func (o *OrderCreateV7RequestLinesInner) HasVendorPartNumber() bool`

HasVendorPartNumber returns a boolean if a field has been set.

### SetVendorPartNumberNil

`func (o *OrderCreateV7RequestLinesInner) SetVendorPartNumberNil(b bool)`

 SetVendorPartNumberNil sets the value for VendorPartNumber to be an explicit nil

### UnsetVendorPartNumber
`func (o *OrderCreateV7RequestLinesInner) UnsetVendorPartNumber()`

UnsetVendorPartNumber ensures that no value is present for VendorPartNumber, not even an explicit nil
### GetQuantity

`func (o *OrderCreateV7RequestLinesInner) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *OrderCreateV7RequestLinesInner) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *OrderCreateV7RequestLinesInner) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *OrderCreateV7RequestLinesInner) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetUnitPrice

`func (o *OrderCreateV7RequestLinesInner) GetUnitPrice() float32`

GetUnitPrice returns the UnitPrice field if non-nil, zero value otherwise.

### GetUnitPriceOk

`func (o *OrderCreateV7RequestLinesInner) GetUnitPriceOk() (*float32, bool)`

GetUnitPriceOk returns a tuple with the UnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPrice

`func (o *OrderCreateV7RequestLinesInner) SetUnitPrice(v float32)`

SetUnitPrice sets UnitPrice field to given value.

### HasUnitPrice

`func (o *OrderCreateV7RequestLinesInner) HasUnitPrice() bool`

HasUnitPrice returns a boolean if a field has been set.

### SetUnitPriceNil

`func (o *OrderCreateV7RequestLinesInner) SetUnitPriceNil(b bool)`

 SetUnitPriceNil sets the value for UnitPrice to be an explicit nil

### UnsetUnitPrice
`func (o *OrderCreateV7RequestLinesInner) UnsetUnitPrice()`

UnsetUnitPrice ensures that no value is present for UnitPrice, not even an explicit nil
### GetSpecialBidNumber

`func (o *OrderCreateV7RequestLinesInner) GetSpecialBidNumber() string`

GetSpecialBidNumber returns the SpecialBidNumber field if non-nil, zero value otherwise.

### GetSpecialBidNumberOk

`func (o *OrderCreateV7RequestLinesInner) GetSpecialBidNumberOk() (*string, bool)`

GetSpecialBidNumberOk returns a tuple with the SpecialBidNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialBidNumber

`func (o *OrderCreateV7RequestLinesInner) SetSpecialBidNumber(v string)`

SetSpecialBidNumber sets SpecialBidNumber field to given value.

### HasSpecialBidNumber

`func (o *OrderCreateV7RequestLinesInner) HasSpecialBidNumber() bool`

HasSpecialBidNumber returns a boolean if a field has been set.

### SetSpecialBidNumberNil

`func (o *OrderCreateV7RequestLinesInner) SetSpecialBidNumberNil(b bool)`

 SetSpecialBidNumberNil sets the value for SpecialBidNumber to be an explicit nil

### UnsetSpecialBidNumber
`func (o *OrderCreateV7RequestLinesInner) UnsetSpecialBidNumber()`

UnsetSpecialBidNumber ensures that no value is present for SpecialBidNumber, not even an explicit nil
### GetEndUserPrice

`func (o *OrderCreateV7RequestLinesInner) GetEndUserPrice() float32`

GetEndUserPrice returns the EndUserPrice field if non-nil, zero value otherwise.

### GetEndUserPriceOk

`func (o *OrderCreateV7RequestLinesInner) GetEndUserPriceOk() (*float32, bool)`

GetEndUserPriceOk returns a tuple with the EndUserPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndUserPrice

`func (o *OrderCreateV7RequestLinesInner) SetEndUserPrice(v float32)`

SetEndUserPrice sets EndUserPrice field to given value.

### HasEndUserPrice

`func (o *OrderCreateV7RequestLinesInner) HasEndUserPrice() bool`

HasEndUserPrice returns a boolean if a field has been set.

### SetEndUserPriceNil

`func (o *OrderCreateV7RequestLinesInner) SetEndUserPriceNil(b bool)`

 SetEndUserPriceNil sets the value for EndUserPrice to be an explicit nil

### UnsetEndUserPrice
`func (o *OrderCreateV7RequestLinesInner) UnsetEndUserPrice()`

UnsetEndUserPrice ensures that no value is present for EndUserPrice, not even an explicit nil
### GetNotes

`func (o *OrderCreateV7RequestLinesInner) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *OrderCreateV7RequestLinesInner) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *OrderCreateV7RequestLinesInner) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *OrderCreateV7RequestLinesInner) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *OrderCreateV7RequestLinesInner) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *OrderCreateV7RequestLinesInner) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetEndUserInfo

`func (o *OrderCreateV7RequestLinesInner) GetEndUserInfo() []OrderCreateV7RequestLinesInnerEndUserInfoInner`

GetEndUserInfo returns the EndUserInfo field if non-nil, zero value otherwise.

### GetEndUserInfoOk

`func (o *OrderCreateV7RequestLinesInner) GetEndUserInfoOk() (*[]OrderCreateV7RequestLinesInnerEndUserInfoInner, bool)`

GetEndUserInfoOk returns a tuple with the EndUserInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndUserInfo

`func (o *OrderCreateV7RequestLinesInner) SetEndUserInfo(v []OrderCreateV7RequestLinesInnerEndUserInfoInner)`

SetEndUserInfo sets EndUserInfo field to given value.

### HasEndUserInfo

`func (o *OrderCreateV7RequestLinesInner) HasEndUserInfo() bool`

HasEndUserInfo returns a boolean if a field has been set.

### SetEndUserInfoNil

`func (o *OrderCreateV7RequestLinesInner) SetEndUserInfoNil(b bool)`

 SetEndUserInfoNil sets the value for EndUserInfo to be an explicit nil

### UnsetEndUserInfo
`func (o *OrderCreateV7RequestLinesInner) UnsetEndUserInfo()`

UnsetEndUserInfo ensures that no value is present for EndUserInfo, not even an explicit nil
### GetAdditionalAttributes

`func (o *OrderCreateV7RequestLinesInner) GetAdditionalAttributes() []OrderCreateV7RequestLinesInnerAdditionalAttributesInner`

GetAdditionalAttributes returns the AdditionalAttributes field if non-nil, zero value otherwise.

### GetAdditionalAttributesOk

`func (o *OrderCreateV7RequestLinesInner) GetAdditionalAttributesOk() (*[]OrderCreateV7RequestLinesInnerAdditionalAttributesInner, bool)`

GetAdditionalAttributesOk returns a tuple with the AdditionalAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalAttributes

`func (o *OrderCreateV7RequestLinesInner) SetAdditionalAttributes(v []OrderCreateV7RequestLinesInnerAdditionalAttributesInner)`

SetAdditionalAttributes sets AdditionalAttributes field to given value.

### HasAdditionalAttributes

`func (o *OrderCreateV7RequestLinesInner) HasAdditionalAttributes() bool`

HasAdditionalAttributes returns a boolean if a field has been set.

### GetVmfAdditionalAttributesLines

`func (o *OrderCreateV7RequestLinesInner) GetVmfAdditionalAttributesLines() []OrderCreateV7RequestLinesInnerVmfAdditionalAttributesLinesInner`

GetVmfAdditionalAttributesLines returns the VmfAdditionalAttributesLines field if non-nil, zero value otherwise.

### GetVmfAdditionalAttributesLinesOk

`func (o *OrderCreateV7RequestLinesInner) GetVmfAdditionalAttributesLinesOk() (*[]OrderCreateV7RequestLinesInnerVmfAdditionalAttributesLinesInner, bool)`

GetVmfAdditionalAttributesLinesOk returns a tuple with the VmfAdditionalAttributesLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmfAdditionalAttributesLines

`func (o *OrderCreateV7RequestLinesInner) SetVmfAdditionalAttributesLines(v []OrderCreateV7RequestLinesInnerVmfAdditionalAttributesLinesInner)`

SetVmfAdditionalAttributesLines sets VmfAdditionalAttributesLines field to given value.

### HasVmfAdditionalAttributesLines

`func (o *OrderCreateV7RequestLinesInner) HasVmfAdditionalAttributesLines() bool`

HasVmfAdditionalAttributesLines returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


