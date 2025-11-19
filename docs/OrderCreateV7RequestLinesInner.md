# OrderCreateV7RequestLinesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerLineNumber** | Pointer to **string** | The reseller&#39;s line item number for reference in their system. The customer line number needs to be a unique numeric value between 1 and 884. In the event we receive duplicate values or alphanumeric values in the customer line number, we will re-sequence the customer line number. To prevent re-sequencing, please use a unique numeric value between 1 and 884 in the customer line number. | [optional] 
**IngramPartNumber** | Pointer to **string** | The unique IngramMicro part number. | [optional] 
**VendorPartNumber** | Pointer to **string** | The vendor&#39;s part number for the line item. | [optional] 
**Quantity** | Pointer to **int32** | The requested quantity of the line item. | [optional] 
**UnitPrice** | Pointer to **float32** | The reseller-requested unit price for the line item. The unit price is not guaranteed. | [optional] 
**SpecialBidNumber** | Pointer to **string** | The line-level bid number provided to the reseller by the vendor for special pricing and discounts. Used to track the bid number in the case of split orders or where different line items have different bid numbers. Line-level bid number take precedence over header-level bid numbers. | [optional] 
**EndUserPrice** | Pointer to **float32** | The end-user price. Required for Export Orders. | [optional] 
**Notes** | Pointer to **string** | The attribute field data. | [optional] 
**ResourceId** | Pointer to **string** | The resource id of the subscription | [optional] 
**Planid** | Pointer to **string** | ID of the subscription plan | [optional] 
**SubscriptionPeriod** | Pointer to [**[]OrderCreateV7RequestLinesInnerSubscriptionPeriodInner**](OrderCreateV7RequestLinesInnerSubscriptionPeriodInner.md) | The object containing the list of options related to the subscription period. | [optional] 
**BillingPeriod** | Pointer to [**[]OrderCreateV7RequestLinesInnerBillingPeriodInner**](OrderCreateV7RequestLinesInnerBillingPeriodInner.md) | The object containing the list of options related to the billing period. | [optional] 
**Margin** | Pointer to **float32** | Line-level margin requested by customer | [optional] 
**EndCustomerPrice** | Pointer to **float32** | Line-level end-customer price requsted by customer | [optional] 
**VriAdditionalAttributes** | Pointer to [**[]OrderCreateV7RequestVmfVendorAdditionalAttributesInner**](OrderCreateV7RequestVmfVendorAdditionalAttributesInner.md) | The object containing the list of Vendor Mandatory Fields required by the vendor for the subscription products. | [optional] 
**EndUserInfo** | Pointer to [**[]OrderCreateV7RequestLinesInnerEndUserInfoInner**](OrderCreateV7RequestLinesInnerEndUserInfoInner.md) |  | [optional] 
**AdditionalAttributes** | Pointer to [**[]OrderCreateV7RequestLinesInnerAdditionalAttributesInner**](OrderCreateV7RequestLinesInnerAdditionalAttributesInner.md) |  | [optional] 
**WarrantyInfo** | Pointer to [**OrderCreateV7RequestLinesInnerWarrantyInfo**](OrderCreateV7RequestLinesInnerWarrantyInfo.md) |  | [optional] 
**VmfAdditionalAttributesLines** | Pointer to [**[]OrderCreateV7RequestLinesInnerVmfAdditionalAttributesLinesInner**](OrderCreateV7RequestLinesInnerVmfAdditionalAttributesLinesInner.md) | The object containing the list of fields required at a line level by the vendor.&lt;br&gt; This a &lt;code&gt;Deprecated&lt;/code&gt; object. Kindly use &lt;b&gt;vmfVendorAdditionalAttributes&lt;/b&gt; object | [optional] 

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

### GetResourceId

`func (o *OrderCreateV7RequestLinesInner) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *OrderCreateV7RequestLinesInner) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *OrderCreateV7RequestLinesInner) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *OrderCreateV7RequestLinesInner) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### GetPlanid

`func (o *OrderCreateV7RequestLinesInner) GetPlanid() string`

GetPlanid returns the Planid field if non-nil, zero value otherwise.

### GetPlanidOk

`func (o *OrderCreateV7RequestLinesInner) GetPlanidOk() (*string, bool)`

GetPlanidOk returns a tuple with the Planid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanid

`func (o *OrderCreateV7RequestLinesInner) SetPlanid(v string)`

SetPlanid sets Planid field to given value.

### HasPlanid

`func (o *OrderCreateV7RequestLinesInner) HasPlanid() bool`

HasPlanid returns a boolean if a field has been set.

### GetSubscriptionPeriod

`func (o *OrderCreateV7RequestLinesInner) GetSubscriptionPeriod() []OrderCreateV7RequestLinesInnerSubscriptionPeriodInner`

GetSubscriptionPeriod returns the SubscriptionPeriod field if non-nil, zero value otherwise.

### GetSubscriptionPeriodOk

`func (o *OrderCreateV7RequestLinesInner) GetSubscriptionPeriodOk() (*[]OrderCreateV7RequestLinesInnerSubscriptionPeriodInner, bool)`

GetSubscriptionPeriodOk returns a tuple with the SubscriptionPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionPeriod

`func (o *OrderCreateV7RequestLinesInner) SetSubscriptionPeriod(v []OrderCreateV7RequestLinesInnerSubscriptionPeriodInner)`

SetSubscriptionPeriod sets SubscriptionPeriod field to given value.

### HasSubscriptionPeriod

`func (o *OrderCreateV7RequestLinesInner) HasSubscriptionPeriod() bool`

HasSubscriptionPeriod returns a boolean if a field has been set.

### GetBillingPeriod

`func (o *OrderCreateV7RequestLinesInner) GetBillingPeriod() []OrderCreateV7RequestLinesInnerBillingPeriodInner`

GetBillingPeriod returns the BillingPeriod field if non-nil, zero value otherwise.

### GetBillingPeriodOk

`func (o *OrderCreateV7RequestLinesInner) GetBillingPeriodOk() (*[]OrderCreateV7RequestLinesInnerBillingPeriodInner, bool)`

GetBillingPeriodOk returns a tuple with the BillingPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingPeriod

`func (o *OrderCreateV7RequestLinesInner) SetBillingPeriod(v []OrderCreateV7RequestLinesInnerBillingPeriodInner)`

SetBillingPeriod sets BillingPeriod field to given value.

### HasBillingPeriod

`func (o *OrderCreateV7RequestLinesInner) HasBillingPeriod() bool`

HasBillingPeriod returns a boolean if a field has been set.

### GetMargin

`func (o *OrderCreateV7RequestLinesInner) GetMargin() float32`

GetMargin returns the Margin field if non-nil, zero value otherwise.

### GetMarginOk

`func (o *OrderCreateV7RequestLinesInner) GetMarginOk() (*float32, bool)`

GetMarginOk returns a tuple with the Margin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMargin

`func (o *OrderCreateV7RequestLinesInner) SetMargin(v float32)`

SetMargin sets Margin field to given value.

### HasMargin

`func (o *OrderCreateV7RequestLinesInner) HasMargin() bool`

HasMargin returns a boolean if a field has been set.

### GetEndCustomerPrice

`func (o *OrderCreateV7RequestLinesInner) GetEndCustomerPrice() float32`

GetEndCustomerPrice returns the EndCustomerPrice field if non-nil, zero value otherwise.

### GetEndCustomerPriceOk

`func (o *OrderCreateV7RequestLinesInner) GetEndCustomerPriceOk() (*float32, bool)`

GetEndCustomerPriceOk returns a tuple with the EndCustomerPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndCustomerPrice

`func (o *OrderCreateV7RequestLinesInner) SetEndCustomerPrice(v float32)`

SetEndCustomerPrice sets EndCustomerPrice field to given value.

### HasEndCustomerPrice

`func (o *OrderCreateV7RequestLinesInner) HasEndCustomerPrice() bool`

HasEndCustomerPrice returns a boolean if a field has been set.

### GetVriAdditionalAttributes

`func (o *OrderCreateV7RequestLinesInner) GetVriAdditionalAttributes() []OrderCreateV7RequestVmfVendorAdditionalAttributesInner`

GetVriAdditionalAttributes returns the VriAdditionalAttributes field if non-nil, zero value otherwise.

### GetVriAdditionalAttributesOk

`func (o *OrderCreateV7RequestLinesInner) GetVriAdditionalAttributesOk() (*[]OrderCreateV7RequestVmfVendorAdditionalAttributesInner, bool)`

GetVriAdditionalAttributesOk returns a tuple with the VriAdditionalAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVriAdditionalAttributes

`func (o *OrderCreateV7RequestLinesInner) SetVriAdditionalAttributes(v []OrderCreateV7RequestVmfVendorAdditionalAttributesInner)`

SetVriAdditionalAttributes sets VriAdditionalAttributes field to given value.

### HasVriAdditionalAttributes

`func (o *OrderCreateV7RequestLinesInner) HasVriAdditionalAttributes() bool`

HasVriAdditionalAttributes returns a boolean if a field has been set.

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

### GetWarrantyInfo

`func (o *OrderCreateV7RequestLinesInner) GetWarrantyInfo() OrderCreateV7RequestLinesInnerWarrantyInfo`

GetWarrantyInfo returns the WarrantyInfo field if non-nil, zero value otherwise.

### GetWarrantyInfoOk

`func (o *OrderCreateV7RequestLinesInner) GetWarrantyInfoOk() (*OrderCreateV7RequestLinesInnerWarrantyInfo, bool)`

GetWarrantyInfoOk returns a tuple with the WarrantyInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarrantyInfo

`func (o *OrderCreateV7RequestLinesInner) SetWarrantyInfo(v OrderCreateV7RequestLinesInnerWarrantyInfo)`

SetWarrantyInfo sets WarrantyInfo field to given value.

### HasWarrantyInfo

`func (o *OrderCreateV7RequestLinesInner) HasWarrantyInfo() bool`

HasWarrantyInfo returns a boolean if a field has been set.

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


