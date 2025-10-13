# PriceAndAvailabilityRequestProductsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IngramPartNumber** | Pointer to **NullableString** | Ingram Micro unique part number for the product. | [optional] 
**VendorPartNumber** | Pointer to **NullableString** | Vendor’s part number for the product. | [optional] 
**CustomerPartNumber** | Pointer to **NullableString** | Reseller/end-user’s part number for the product. | [optional] 
**Upc** | Pointer to **NullableString** | The UPC code for the product. Consists of 12 numeric digits that are uniquely assigned to each trade item. | [optional] 
**QuantityRequested** | Pointer to **NullableString** | Number of quantity of the Product. | [optional] 
**PlanID** | Pointer to **NullableString** | Id of the plan | [optional] 
**AdditionalAttributes** | Pointer to [**[]PriceAndAvailabilityRequestProductsInnerAdditionalAttributesInner**](PriceAndAvailabilityRequestProductsInnerAdditionalAttributesInner.md) |  | [optional] 

## Methods

### NewPriceAndAvailabilityRequestProductsInner

`func NewPriceAndAvailabilityRequestProductsInner() *PriceAndAvailabilityRequestProductsInner`

NewPriceAndAvailabilityRequestProductsInner instantiates a new PriceAndAvailabilityRequestProductsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceAndAvailabilityRequestProductsInnerWithDefaults

`func NewPriceAndAvailabilityRequestProductsInnerWithDefaults() *PriceAndAvailabilityRequestProductsInner`

NewPriceAndAvailabilityRequestProductsInnerWithDefaults instantiates a new PriceAndAvailabilityRequestProductsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIngramPartNumber

`func (o *PriceAndAvailabilityRequestProductsInner) GetIngramPartNumber() string`

GetIngramPartNumber returns the IngramPartNumber field if non-nil, zero value otherwise.

### GetIngramPartNumberOk

`func (o *PriceAndAvailabilityRequestProductsInner) GetIngramPartNumberOk() (*string, bool)`

GetIngramPartNumberOk returns a tuple with the IngramPartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngramPartNumber

`func (o *PriceAndAvailabilityRequestProductsInner) SetIngramPartNumber(v string)`

SetIngramPartNumber sets IngramPartNumber field to given value.

### HasIngramPartNumber

`func (o *PriceAndAvailabilityRequestProductsInner) HasIngramPartNumber() bool`

HasIngramPartNumber returns a boolean if a field has been set.

### SetIngramPartNumberNil

`func (o *PriceAndAvailabilityRequestProductsInner) SetIngramPartNumberNil(b bool)`

 SetIngramPartNumberNil sets the value for IngramPartNumber to be an explicit nil

### UnsetIngramPartNumber
`func (o *PriceAndAvailabilityRequestProductsInner) UnsetIngramPartNumber()`

UnsetIngramPartNumber ensures that no value is present for IngramPartNumber, not even an explicit nil
### GetVendorPartNumber

`func (o *PriceAndAvailabilityRequestProductsInner) GetVendorPartNumber() string`

GetVendorPartNumber returns the VendorPartNumber field if non-nil, zero value otherwise.

### GetVendorPartNumberOk

`func (o *PriceAndAvailabilityRequestProductsInner) GetVendorPartNumberOk() (*string, bool)`

GetVendorPartNumberOk returns a tuple with the VendorPartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorPartNumber

`func (o *PriceAndAvailabilityRequestProductsInner) SetVendorPartNumber(v string)`

SetVendorPartNumber sets VendorPartNumber field to given value.

### HasVendorPartNumber

`func (o *PriceAndAvailabilityRequestProductsInner) HasVendorPartNumber() bool`

HasVendorPartNumber returns a boolean if a field has been set.

### SetVendorPartNumberNil

`func (o *PriceAndAvailabilityRequestProductsInner) SetVendorPartNumberNil(b bool)`

 SetVendorPartNumberNil sets the value for VendorPartNumber to be an explicit nil

### UnsetVendorPartNumber
`func (o *PriceAndAvailabilityRequestProductsInner) UnsetVendorPartNumber()`

UnsetVendorPartNumber ensures that no value is present for VendorPartNumber, not even an explicit nil
### GetCustomerPartNumber

`func (o *PriceAndAvailabilityRequestProductsInner) GetCustomerPartNumber() string`

GetCustomerPartNumber returns the CustomerPartNumber field if non-nil, zero value otherwise.

### GetCustomerPartNumberOk

`func (o *PriceAndAvailabilityRequestProductsInner) GetCustomerPartNumberOk() (*string, bool)`

GetCustomerPartNumberOk returns a tuple with the CustomerPartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerPartNumber

`func (o *PriceAndAvailabilityRequestProductsInner) SetCustomerPartNumber(v string)`

SetCustomerPartNumber sets CustomerPartNumber field to given value.

### HasCustomerPartNumber

`func (o *PriceAndAvailabilityRequestProductsInner) HasCustomerPartNumber() bool`

HasCustomerPartNumber returns a boolean if a field has been set.

### SetCustomerPartNumberNil

`func (o *PriceAndAvailabilityRequestProductsInner) SetCustomerPartNumberNil(b bool)`

 SetCustomerPartNumberNil sets the value for CustomerPartNumber to be an explicit nil

### UnsetCustomerPartNumber
`func (o *PriceAndAvailabilityRequestProductsInner) UnsetCustomerPartNumber()`

UnsetCustomerPartNumber ensures that no value is present for CustomerPartNumber, not even an explicit nil
### GetUpc

`func (o *PriceAndAvailabilityRequestProductsInner) GetUpc() string`

GetUpc returns the Upc field if non-nil, zero value otherwise.

### GetUpcOk

`func (o *PriceAndAvailabilityRequestProductsInner) GetUpcOk() (*string, bool)`

GetUpcOk returns a tuple with the Upc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpc

`func (o *PriceAndAvailabilityRequestProductsInner) SetUpc(v string)`

SetUpc sets Upc field to given value.

### HasUpc

`func (o *PriceAndAvailabilityRequestProductsInner) HasUpc() bool`

HasUpc returns a boolean if a field has been set.

### SetUpcNil

`func (o *PriceAndAvailabilityRequestProductsInner) SetUpcNil(b bool)`

 SetUpcNil sets the value for Upc to be an explicit nil

### UnsetUpc
`func (o *PriceAndAvailabilityRequestProductsInner) UnsetUpc()`

UnsetUpc ensures that no value is present for Upc, not even an explicit nil
### GetQuantityRequested

`func (o *PriceAndAvailabilityRequestProductsInner) GetQuantityRequested() string`

GetQuantityRequested returns the QuantityRequested field if non-nil, zero value otherwise.

### GetQuantityRequestedOk

`func (o *PriceAndAvailabilityRequestProductsInner) GetQuantityRequestedOk() (*string, bool)`

GetQuantityRequestedOk returns a tuple with the QuantityRequested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityRequested

`func (o *PriceAndAvailabilityRequestProductsInner) SetQuantityRequested(v string)`

SetQuantityRequested sets QuantityRequested field to given value.

### HasQuantityRequested

`func (o *PriceAndAvailabilityRequestProductsInner) HasQuantityRequested() bool`

HasQuantityRequested returns a boolean if a field has been set.

### SetQuantityRequestedNil

`func (o *PriceAndAvailabilityRequestProductsInner) SetQuantityRequestedNil(b bool)`

 SetQuantityRequestedNil sets the value for QuantityRequested to be an explicit nil

### UnsetQuantityRequested
`func (o *PriceAndAvailabilityRequestProductsInner) UnsetQuantityRequested()`

UnsetQuantityRequested ensures that no value is present for QuantityRequested, not even an explicit nil
### GetPlanID

`func (o *PriceAndAvailabilityRequestProductsInner) GetPlanID() string`

GetPlanID returns the PlanID field if non-nil, zero value otherwise.

### GetPlanIDOk

`func (o *PriceAndAvailabilityRequestProductsInner) GetPlanIDOk() (*string, bool)`

GetPlanIDOk returns a tuple with the PlanID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanID

`func (o *PriceAndAvailabilityRequestProductsInner) SetPlanID(v string)`

SetPlanID sets PlanID field to given value.

### HasPlanID

`func (o *PriceAndAvailabilityRequestProductsInner) HasPlanID() bool`

HasPlanID returns a boolean if a field has been set.

### SetPlanIDNil

`func (o *PriceAndAvailabilityRequestProductsInner) SetPlanIDNil(b bool)`

 SetPlanIDNil sets the value for PlanID to be an explicit nil

### UnsetPlanID
`func (o *PriceAndAvailabilityRequestProductsInner) UnsetPlanID()`

UnsetPlanID ensures that no value is present for PlanID, not even an explicit nil
### GetAdditionalAttributes

`func (o *PriceAndAvailabilityRequestProductsInner) GetAdditionalAttributes() []PriceAndAvailabilityRequestProductsInnerAdditionalAttributesInner`

GetAdditionalAttributes returns the AdditionalAttributes field if non-nil, zero value otherwise.

### GetAdditionalAttributesOk

`func (o *PriceAndAvailabilityRequestProductsInner) GetAdditionalAttributesOk() (*[]PriceAndAvailabilityRequestProductsInnerAdditionalAttributesInner, bool)`

GetAdditionalAttributesOk returns a tuple with the AdditionalAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalAttributes

`func (o *PriceAndAvailabilityRequestProductsInner) SetAdditionalAttributes(v []PriceAndAvailabilityRequestProductsInnerAdditionalAttributesInner)`

SetAdditionalAttributes sets AdditionalAttributes field to given value.

### HasAdditionalAttributes

`func (o *PriceAndAvailabilityRequestProductsInner) HasAdditionalAttributes() bool`

HasAdditionalAttributes returns a boolean if a field has been set.

### SetAdditionalAttributesNil

`func (o *PriceAndAvailabilityRequestProductsInner) SetAdditionalAttributesNil(b bool)`

 SetAdditionalAttributesNil sets the value for AdditionalAttributes to be an explicit nil

### UnsetAdditionalAttributes
`func (o *PriceAndAvailabilityRequestProductsInner) UnsetAdditionalAttributes()`

UnsetAdditionalAttributes ensures that no value is present for AdditionalAttributes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


