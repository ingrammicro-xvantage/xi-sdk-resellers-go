# QuoteDetailsResponseProductsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QuoteProductGuid** | Pointer to **string** | Quote Product GUID  is the primary quote key in Ingram Micro&#39;s CRM - needed to retrieve quote details. | [optional] 
**LineNumber** | Pointer to **string** | Line number which the product will appear in the quote.  Line number is manditory when unique configurations are included in a quote and mainting the item line order is required. | [optional] 
**Quantity** | Pointer to **int32** | Quantity of product line item quoted. | [optional] 
**RemainingQuoteQty** | Pointer to **NullableInt32** |  | [optional] 
**MinimumOrderAllowedQty** | Pointer to **NullableInt32** |  | [optional] 
**Notes** | Pointer to **string** | Product line item comments. | [optional] 
**Ean** | Pointer to **string** | EANUPC | [optional] 
**Coo** | Pointer to **string** | Country of Origin. | [optional] 
**IngramPartNumber** | Pointer to **string** | Ingram Micro SKU (stock keeping unit). An identification, usually alphanumeric, of a particular product that allows it to be tracked for inventory purposes | [optional] 
**VendorPartNumber** | Pointer to **string** | Vendor Part Number | [optional] 
**Description** | Pointer to **string** | Product description.  Note - The quote view api returns only the product short description as maintained in Ingram Micro&#39;s crm system.  For long descriptions, please refer to alternative information sources. | [optional] 
**Weight** | Pointer to **float32** | Weight is provided based on country standard.  For countries following Imperial standards - weight is presented as pounds with decimal.  In countries following metric standards, weight is provided as kilograms with decimal. | [optional] 
**WeightUom** | Pointer to **string** | Unit of measure | [optional] 
**IsSuggestionProduct** | Pointer to **bool** | Flag to indicate if a product line item is a suggested product.  The suggested product is provided in addition to the requested quoted products and a suggested option.  Suggested products are grouped together for subtotal and total calculations. | [optional] 
**VpnCategory** | Pointer to **string** | Vendor product category specific to Cisco. HWDW (hardware) or service. | [optional] 
**QuoteProductsSupplierPartAuxiliaryId** | Pointer to **string** | Vendor product configuration ID specific to Cisco. | [optional] 
**VendorName** | Pointer to **string** | Vendor name of the product | [optional] 
**Terms** | Pointer to **string** | Terms of the quote | [optional] 
**PlanDescription** | Pointer to **string** |  | [optional] 
**IsSubscription** | Pointer to **bool** |  | [optional] 
**ResellerMargin** | Pointer to **string** |  | [optional] 
**RequestedStartDate** | Pointer to **string** |  | [optional] 
**StartDate** | Pointer to **string** |  | [optional] 
**EndDate** | Pointer to **string** |  | [optional] 
**SerialNumber** | Pointer to **string** |  | [optional] 
**Price** | Pointer to [**QuoteDetailsResponseProductsInnerPrice**](QuoteDetailsResponseProductsInnerPrice.md) |  | [optional] 
**BillDetails** | Pointer to [**[]QuoteDetailsResponseProductsInnerBillDetailsInner**](QuoteDetailsResponseProductsInnerBillDetailsInner.md) |  | [optional] 

## Methods

### NewQuoteDetailsResponseProductsInner

`func NewQuoteDetailsResponseProductsInner() *QuoteDetailsResponseProductsInner`

NewQuoteDetailsResponseProductsInner instantiates a new QuoteDetailsResponseProductsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuoteDetailsResponseProductsInnerWithDefaults

`func NewQuoteDetailsResponseProductsInnerWithDefaults() *QuoteDetailsResponseProductsInner`

NewQuoteDetailsResponseProductsInnerWithDefaults instantiates a new QuoteDetailsResponseProductsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuoteProductGuid

`func (o *QuoteDetailsResponseProductsInner) GetQuoteProductGuid() string`

GetQuoteProductGuid returns the QuoteProductGuid field if non-nil, zero value otherwise.

### GetQuoteProductGuidOk

`func (o *QuoteDetailsResponseProductsInner) GetQuoteProductGuidOk() (*string, bool)`

GetQuoteProductGuidOk returns a tuple with the QuoteProductGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteProductGuid

`func (o *QuoteDetailsResponseProductsInner) SetQuoteProductGuid(v string)`

SetQuoteProductGuid sets QuoteProductGuid field to given value.

### HasQuoteProductGuid

`func (o *QuoteDetailsResponseProductsInner) HasQuoteProductGuid() bool`

HasQuoteProductGuid returns a boolean if a field has been set.

### GetLineNumber

`func (o *QuoteDetailsResponseProductsInner) GetLineNumber() string`

GetLineNumber returns the LineNumber field if non-nil, zero value otherwise.

### GetLineNumberOk

`func (o *QuoteDetailsResponseProductsInner) GetLineNumberOk() (*string, bool)`

GetLineNumberOk returns a tuple with the LineNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineNumber

`func (o *QuoteDetailsResponseProductsInner) SetLineNumber(v string)`

SetLineNumber sets LineNumber field to given value.

### HasLineNumber

`func (o *QuoteDetailsResponseProductsInner) HasLineNumber() bool`

HasLineNumber returns a boolean if a field has been set.

### GetQuantity

`func (o *QuoteDetailsResponseProductsInner) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *QuoteDetailsResponseProductsInner) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *QuoteDetailsResponseProductsInner) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *QuoteDetailsResponseProductsInner) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetRemainingQuoteQty

`func (o *QuoteDetailsResponseProductsInner) GetRemainingQuoteQty() int32`

GetRemainingQuoteQty returns the RemainingQuoteQty field if non-nil, zero value otherwise.

### GetRemainingQuoteQtyOk

`func (o *QuoteDetailsResponseProductsInner) GetRemainingQuoteQtyOk() (*int32, bool)`

GetRemainingQuoteQtyOk returns a tuple with the RemainingQuoteQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingQuoteQty

`func (o *QuoteDetailsResponseProductsInner) SetRemainingQuoteQty(v int32)`

SetRemainingQuoteQty sets RemainingQuoteQty field to given value.

### HasRemainingQuoteQty

`func (o *QuoteDetailsResponseProductsInner) HasRemainingQuoteQty() bool`

HasRemainingQuoteQty returns a boolean if a field has been set.

### SetRemainingQuoteQtyNil

`func (o *QuoteDetailsResponseProductsInner) SetRemainingQuoteQtyNil(b bool)`

 SetRemainingQuoteQtyNil sets the value for RemainingQuoteQty to be an explicit nil

### UnsetRemainingQuoteQty
`func (o *QuoteDetailsResponseProductsInner) UnsetRemainingQuoteQty()`

UnsetRemainingQuoteQty ensures that no value is present for RemainingQuoteQty, not even an explicit nil
### GetMinimumOrderAllowedQty

`func (o *QuoteDetailsResponseProductsInner) GetMinimumOrderAllowedQty() int32`

GetMinimumOrderAllowedQty returns the MinimumOrderAllowedQty field if non-nil, zero value otherwise.

### GetMinimumOrderAllowedQtyOk

`func (o *QuoteDetailsResponseProductsInner) GetMinimumOrderAllowedQtyOk() (*int32, bool)`

GetMinimumOrderAllowedQtyOk returns a tuple with the MinimumOrderAllowedQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumOrderAllowedQty

`func (o *QuoteDetailsResponseProductsInner) SetMinimumOrderAllowedQty(v int32)`

SetMinimumOrderAllowedQty sets MinimumOrderAllowedQty field to given value.

### HasMinimumOrderAllowedQty

`func (o *QuoteDetailsResponseProductsInner) HasMinimumOrderAllowedQty() bool`

HasMinimumOrderAllowedQty returns a boolean if a field has been set.

### SetMinimumOrderAllowedQtyNil

`func (o *QuoteDetailsResponseProductsInner) SetMinimumOrderAllowedQtyNil(b bool)`

 SetMinimumOrderAllowedQtyNil sets the value for MinimumOrderAllowedQty to be an explicit nil

### UnsetMinimumOrderAllowedQty
`func (o *QuoteDetailsResponseProductsInner) UnsetMinimumOrderAllowedQty()`

UnsetMinimumOrderAllowedQty ensures that no value is present for MinimumOrderAllowedQty, not even an explicit nil
### GetNotes

`func (o *QuoteDetailsResponseProductsInner) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *QuoteDetailsResponseProductsInner) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *QuoteDetailsResponseProductsInner) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *QuoteDetailsResponseProductsInner) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetEan

`func (o *QuoteDetailsResponseProductsInner) GetEan() string`

GetEan returns the Ean field if non-nil, zero value otherwise.

### GetEanOk

`func (o *QuoteDetailsResponseProductsInner) GetEanOk() (*string, bool)`

GetEanOk returns a tuple with the Ean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEan

`func (o *QuoteDetailsResponseProductsInner) SetEan(v string)`

SetEan sets Ean field to given value.

### HasEan

`func (o *QuoteDetailsResponseProductsInner) HasEan() bool`

HasEan returns a boolean if a field has been set.

### GetCoo

`func (o *QuoteDetailsResponseProductsInner) GetCoo() string`

GetCoo returns the Coo field if non-nil, zero value otherwise.

### GetCooOk

`func (o *QuoteDetailsResponseProductsInner) GetCooOk() (*string, bool)`

GetCooOk returns a tuple with the Coo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoo

`func (o *QuoteDetailsResponseProductsInner) SetCoo(v string)`

SetCoo sets Coo field to given value.

### HasCoo

`func (o *QuoteDetailsResponseProductsInner) HasCoo() bool`

HasCoo returns a boolean if a field has been set.

### GetIngramPartNumber

`func (o *QuoteDetailsResponseProductsInner) GetIngramPartNumber() string`

GetIngramPartNumber returns the IngramPartNumber field if non-nil, zero value otherwise.

### GetIngramPartNumberOk

`func (o *QuoteDetailsResponseProductsInner) GetIngramPartNumberOk() (*string, bool)`

GetIngramPartNumberOk returns a tuple with the IngramPartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngramPartNumber

`func (o *QuoteDetailsResponseProductsInner) SetIngramPartNumber(v string)`

SetIngramPartNumber sets IngramPartNumber field to given value.

### HasIngramPartNumber

`func (o *QuoteDetailsResponseProductsInner) HasIngramPartNumber() bool`

HasIngramPartNumber returns a boolean if a field has been set.

### GetVendorPartNumber

`func (o *QuoteDetailsResponseProductsInner) GetVendorPartNumber() string`

GetVendorPartNumber returns the VendorPartNumber field if non-nil, zero value otherwise.

### GetVendorPartNumberOk

`func (o *QuoteDetailsResponseProductsInner) GetVendorPartNumberOk() (*string, bool)`

GetVendorPartNumberOk returns a tuple with the VendorPartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorPartNumber

`func (o *QuoteDetailsResponseProductsInner) SetVendorPartNumber(v string)`

SetVendorPartNumber sets VendorPartNumber field to given value.

### HasVendorPartNumber

`func (o *QuoteDetailsResponseProductsInner) HasVendorPartNumber() bool`

HasVendorPartNumber returns a boolean if a field has been set.

### GetDescription

`func (o *QuoteDetailsResponseProductsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *QuoteDetailsResponseProductsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *QuoteDetailsResponseProductsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *QuoteDetailsResponseProductsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetWeight

`func (o *QuoteDetailsResponseProductsInner) GetWeight() float32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *QuoteDetailsResponseProductsInner) GetWeightOk() (*float32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *QuoteDetailsResponseProductsInner) SetWeight(v float32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *QuoteDetailsResponseProductsInner) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetWeightUom

`func (o *QuoteDetailsResponseProductsInner) GetWeightUom() string`

GetWeightUom returns the WeightUom field if non-nil, zero value otherwise.

### GetWeightUomOk

`func (o *QuoteDetailsResponseProductsInner) GetWeightUomOk() (*string, bool)`

GetWeightUomOk returns a tuple with the WeightUom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightUom

`func (o *QuoteDetailsResponseProductsInner) SetWeightUom(v string)`

SetWeightUom sets WeightUom field to given value.

### HasWeightUom

`func (o *QuoteDetailsResponseProductsInner) HasWeightUom() bool`

HasWeightUom returns a boolean if a field has been set.

### GetIsSuggestionProduct

`func (o *QuoteDetailsResponseProductsInner) GetIsSuggestionProduct() bool`

GetIsSuggestionProduct returns the IsSuggestionProduct field if non-nil, zero value otherwise.

### GetIsSuggestionProductOk

`func (o *QuoteDetailsResponseProductsInner) GetIsSuggestionProductOk() (*bool, bool)`

GetIsSuggestionProductOk returns a tuple with the IsSuggestionProduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSuggestionProduct

`func (o *QuoteDetailsResponseProductsInner) SetIsSuggestionProduct(v bool)`

SetIsSuggestionProduct sets IsSuggestionProduct field to given value.

### HasIsSuggestionProduct

`func (o *QuoteDetailsResponseProductsInner) HasIsSuggestionProduct() bool`

HasIsSuggestionProduct returns a boolean if a field has been set.

### GetVpnCategory

`func (o *QuoteDetailsResponseProductsInner) GetVpnCategory() string`

GetVpnCategory returns the VpnCategory field if non-nil, zero value otherwise.

### GetVpnCategoryOk

`func (o *QuoteDetailsResponseProductsInner) GetVpnCategoryOk() (*string, bool)`

GetVpnCategoryOk returns a tuple with the VpnCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnCategory

`func (o *QuoteDetailsResponseProductsInner) SetVpnCategory(v string)`

SetVpnCategory sets VpnCategory field to given value.

### HasVpnCategory

`func (o *QuoteDetailsResponseProductsInner) HasVpnCategory() bool`

HasVpnCategory returns a boolean if a field has been set.

### GetQuoteProductsSupplierPartAuxiliaryId

`func (o *QuoteDetailsResponseProductsInner) GetQuoteProductsSupplierPartAuxiliaryId() string`

GetQuoteProductsSupplierPartAuxiliaryId returns the QuoteProductsSupplierPartAuxiliaryId field if non-nil, zero value otherwise.

### GetQuoteProductsSupplierPartAuxiliaryIdOk

`func (o *QuoteDetailsResponseProductsInner) GetQuoteProductsSupplierPartAuxiliaryIdOk() (*string, bool)`

GetQuoteProductsSupplierPartAuxiliaryIdOk returns a tuple with the QuoteProductsSupplierPartAuxiliaryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteProductsSupplierPartAuxiliaryId

`func (o *QuoteDetailsResponseProductsInner) SetQuoteProductsSupplierPartAuxiliaryId(v string)`

SetQuoteProductsSupplierPartAuxiliaryId sets QuoteProductsSupplierPartAuxiliaryId field to given value.

### HasQuoteProductsSupplierPartAuxiliaryId

`func (o *QuoteDetailsResponseProductsInner) HasQuoteProductsSupplierPartAuxiliaryId() bool`

HasQuoteProductsSupplierPartAuxiliaryId returns a boolean if a field has been set.

### GetVendorName

`func (o *QuoteDetailsResponseProductsInner) GetVendorName() string`

GetVendorName returns the VendorName field if non-nil, zero value otherwise.

### GetVendorNameOk

`func (o *QuoteDetailsResponseProductsInner) GetVendorNameOk() (*string, bool)`

GetVendorNameOk returns a tuple with the VendorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorName

`func (o *QuoteDetailsResponseProductsInner) SetVendorName(v string)`

SetVendorName sets VendorName field to given value.

### HasVendorName

`func (o *QuoteDetailsResponseProductsInner) HasVendorName() bool`

HasVendorName returns a boolean if a field has been set.

### GetTerms

`func (o *QuoteDetailsResponseProductsInner) GetTerms() string`

GetTerms returns the Terms field if non-nil, zero value otherwise.

### GetTermsOk

`func (o *QuoteDetailsResponseProductsInner) GetTermsOk() (*string, bool)`

GetTermsOk returns a tuple with the Terms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerms

`func (o *QuoteDetailsResponseProductsInner) SetTerms(v string)`

SetTerms sets Terms field to given value.

### HasTerms

`func (o *QuoteDetailsResponseProductsInner) HasTerms() bool`

HasTerms returns a boolean if a field has been set.

### GetPlanDescription

`func (o *QuoteDetailsResponseProductsInner) GetPlanDescription() string`

GetPlanDescription returns the PlanDescription field if non-nil, zero value otherwise.

### GetPlanDescriptionOk

`func (o *QuoteDetailsResponseProductsInner) GetPlanDescriptionOk() (*string, bool)`

GetPlanDescriptionOk returns a tuple with the PlanDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanDescription

`func (o *QuoteDetailsResponseProductsInner) SetPlanDescription(v string)`

SetPlanDescription sets PlanDescription field to given value.

### HasPlanDescription

`func (o *QuoteDetailsResponseProductsInner) HasPlanDescription() bool`

HasPlanDescription returns a boolean if a field has been set.

### GetIsSubscription

`func (o *QuoteDetailsResponseProductsInner) GetIsSubscription() bool`

GetIsSubscription returns the IsSubscription field if non-nil, zero value otherwise.

### GetIsSubscriptionOk

`func (o *QuoteDetailsResponseProductsInner) GetIsSubscriptionOk() (*bool, bool)`

GetIsSubscriptionOk returns a tuple with the IsSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSubscription

`func (o *QuoteDetailsResponseProductsInner) SetIsSubscription(v bool)`

SetIsSubscription sets IsSubscription field to given value.

### HasIsSubscription

`func (o *QuoteDetailsResponseProductsInner) HasIsSubscription() bool`

HasIsSubscription returns a boolean if a field has been set.

### GetResellerMargin

`func (o *QuoteDetailsResponseProductsInner) GetResellerMargin() string`

GetResellerMargin returns the ResellerMargin field if non-nil, zero value otherwise.

### GetResellerMarginOk

`func (o *QuoteDetailsResponseProductsInner) GetResellerMarginOk() (*string, bool)`

GetResellerMarginOk returns a tuple with the ResellerMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResellerMargin

`func (o *QuoteDetailsResponseProductsInner) SetResellerMargin(v string)`

SetResellerMargin sets ResellerMargin field to given value.

### HasResellerMargin

`func (o *QuoteDetailsResponseProductsInner) HasResellerMargin() bool`

HasResellerMargin returns a boolean if a field has been set.

### GetRequestedStartDate

`func (o *QuoteDetailsResponseProductsInner) GetRequestedStartDate() string`

GetRequestedStartDate returns the RequestedStartDate field if non-nil, zero value otherwise.

### GetRequestedStartDateOk

`func (o *QuoteDetailsResponseProductsInner) GetRequestedStartDateOk() (*string, bool)`

GetRequestedStartDateOk returns a tuple with the RequestedStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedStartDate

`func (o *QuoteDetailsResponseProductsInner) SetRequestedStartDate(v string)`

SetRequestedStartDate sets RequestedStartDate field to given value.

### HasRequestedStartDate

`func (o *QuoteDetailsResponseProductsInner) HasRequestedStartDate() bool`

HasRequestedStartDate returns a boolean if a field has been set.

### GetStartDate

`func (o *QuoteDetailsResponseProductsInner) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *QuoteDetailsResponseProductsInner) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *QuoteDetailsResponseProductsInner) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *QuoteDetailsResponseProductsInner) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *QuoteDetailsResponseProductsInner) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *QuoteDetailsResponseProductsInner) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *QuoteDetailsResponseProductsInner) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *QuoteDetailsResponseProductsInner) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetSerialNumber

`func (o *QuoteDetailsResponseProductsInner) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *QuoteDetailsResponseProductsInner) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *QuoteDetailsResponseProductsInner) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *QuoteDetailsResponseProductsInner) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetPrice

`func (o *QuoteDetailsResponseProductsInner) GetPrice() QuoteDetailsResponseProductsInnerPrice`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *QuoteDetailsResponseProductsInner) GetPriceOk() (*QuoteDetailsResponseProductsInnerPrice, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *QuoteDetailsResponseProductsInner) SetPrice(v QuoteDetailsResponseProductsInnerPrice)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *QuoteDetailsResponseProductsInner) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetBillDetails

`func (o *QuoteDetailsResponseProductsInner) GetBillDetails() []QuoteDetailsResponseProductsInnerBillDetailsInner`

GetBillDetails returns the BillDetails field if non-nil, zero value otherwise.

### GetBillDetailsOk

`func (o *QuoteDetailsResponseProductsInner) GetBillDetailsOk() (*[]QuoteDetailsResponseProductsInnerBillDetailsInner, bool)`

GetBillDetailsOk returns a tuple with the BillDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillDetails

`func (o *QuoteDetailsResponseProductsInner) SetBillDetails(v []QuoteDetailsResponseProductsInnerBillDetailsInner)`

SetBillDetails sets BillDetails field to given value.

### HasBillDetails

`func (o *QuoteDetailsResponseProductsInner) HasBillDetails() bool`

HasBillDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


