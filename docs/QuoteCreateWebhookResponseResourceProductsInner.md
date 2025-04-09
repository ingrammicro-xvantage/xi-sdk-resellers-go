# QuoteCreateWebhookResponseResourceProductsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QuoteProductGuid** | Pointer to **string** | Quote Product GUID  is the primary quote key in Ingram Micro&#39;s CRM - needed to retrieve quote details. | [optional] 
**LineNumber** | Pointer to **string** | Line number which the product will appear in the quote.  Line number is manditory when unique configurations are included in a quote and mainting the item line order is required. | [optional] 
**Quantity** | Pointer to **string** | Quantity of product line item quoted. | [optional] 
**Notes** | Pointer to **string** | Product line item comments. | [optional] 
**Ean** | Pointer to **string** | EANUPC | [optional] 
**Coo** | Pointer to **string** | Country of Origin. | [optional] 
**IngramPartNumber** | Pointer to **string** | Ingram Micro SKU (stock keeping unit). An identification, usually alphanumeric, of a particular product that allows it to be tracked for inventory purposes | [optional] 
**VendorPartNumber** | Pointer to **string** | Vendor Part Number | [optional] 
**Description** | Pointer to **string** | Product description.  Note - The quote view api returns only the product short description as maintained in Ingram Micro&#39;s crm system.  For long descriptions, please refer to alternative information sources. | [optional] 
**Weight** | Pointer to **string** | Weight is provided based on country standard.  For countries following Imperial standards - weight is presented as pounds with decimal.  In countries following metric standards, weight is provided as kilograms with decimal. | [optional] 
**WeightUom** | Pointer to **string** | Unit of measure | [optional] 
**IsSuggestionProduct** | Pointer to **string** | Flag to indicate if a product line item is a suggested product.  The suggested product is provided in addition to the requested quoted products and a suggested option.  Suggested products are grouped together for subtotal and total calculations. | [optional] 
**VpnCategory** | Pointer to **string** | Vendor product category specific to Cisco. HWDW (hardware) or service. | [optional] 
**QuoteProductsSupplierPartAuxiliaryId** | Pointer to **string** | Vendor product configuration ID specific to Cisco. | [optional] 
**VendorName** | Pointer to **string** | Vendor name of the product | [optional] 
**Terms** | Pointer to **string** | Terms of the quote | [optional] 
**Price** | Pointer to [**QuoteCreateWebhookResponseResourceProductsInnerPrice**](QuoteCreateWebhookResponseResourceProductsInnerPrice.md) |  | [optional] 

## Methods

### NewQuoteCreateWebhookResponseResourceProductsInner

`func NewQuoteCreateWebhookResponseResourceProductsInner() *QuoteCreateWebhookResponseResourceProductsInner`

NewQuoteCreateWebhookResponseResourceProductsInner instantiates a new QuoteCreateWebhookResponseResourceProductsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuoteCreateWebhookResponseResourceProductsInnerWithDefaults

`func NewQuoteCreateWebhookResponseResourceProductsInnerWithDefaults() *QuoteCreateWebhookResponseResourceProductsInner`

NewQuoteCreateWebhookResponseResourceProductsInnerWithDefaults instantiates a new QuoteCreateWebhookResponseResourceProductsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuoteProductGuid

`func (o *QuoteCreateWebhookResponseResourceProductsInner) GetQuoteProductGuid() string`

GetQuoteProductGuid returns the QuoteProductGuid field if non-nil, zero value otherwise.

### GetQuoteProductGuidOk

`func (o *QuoteCreateWebhookResponseResourceProductsInner) GetQuoteProductGuidOk() (*string, bool)`

GetQuoteProductGuidOk returns a tuple with the QuoteProductGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteProductGuid

`func (o *QuoteCreateWebhookResponseResourceProductsInner) SetQuoteProductGuid(v string)`

SetQuoteProductGuid sets QuoteProductGuid field to given value.

### HasQuoteProductGuid

`func (o *QuoteCreateWebhookResponseResourceProductsInner) HasQuoteProductGuid() bool`

HasQuoteProductGuid returns a boolean if a field has been set.

### GetLineNumber

`func (o *QuoteCreateWebhookResponseResourceProductsInner) GetLineNumber() string`

GetLineNumber returns the LineNumber field if non-nil, zero value otherwise.

### GetLineNumberOk

`func (o *QuoteCreateWebhookResponseResourceProductsInner) GetLineNumberOk() (*string, bool)`

GetLineNumberOk returns a tuple with the LineNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineNumber

`func (o *QuoteCreateWebhookResponseResourceProductsInner) SetLineNumber(v string)`

SetLineNumber sets LineNumber field to given value.

### HasLineNumber

`func (o *QuoteCreateWebhookResponseResourceProductsInner) HasLineNumber() bool`

HasLineNumber returns a boolean if a field has been set.

### GetQuantity

`func (o *QuoteCreateWebhookResponseResourceProductsInner) GetQuantity() string`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *QuoteCreateWebhookResponseResourceProductsInner) GetQuantityOk() (*string, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *QuoteCreateWebhookResponseResourceProductsInner) SetQuantity(v string)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *QuoteCreateWebhookResponseResourceProductsInner) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetNotes

`func (o *QuoteCreateWebhookResponseResourceProductsInner) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *QuoteCreateWebhookResponseResourceProductsInner) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *QuoteCreateWebhookResponseResourceProductsInner) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *QuoteCreateWebhookResponseResourceProductsInner) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetEan

`func (o *QuoteCreateWebhookResponseResourceProductsInner) GetEan() string`

GetEan returns the Ean field if non-nil, zero value otherwise.

### GetEanOk

`func (o *QuoteCreateWebhookResponseResourceProductsInner) GetEanOk() (*string, bool)`

GetEanOk returns a tuple with the Ean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEan

`func (o *QuoteCreateWebhookResponseResourceProductsInner) SetEan(v string)`

SetEan sets Ean field to given value.

### HasEan

`func (o *QuoteCreateWebhookResponseResourceProductsInner) HasEan() bool`

HasEan returns a boolean if a field has been set.

### GetCoo

`func (o *QuoteCreateWebhookResponseResourceProductsInner) GetCoo() string`

GetCoo returns the Coo field if non-nil, zero value otherwise.

### GetCooOk

`func (o *QuoteCreateWebhookResponseResourceProductsInner) GetCooOk() (*string, bool)`

GetCooOk returns a tuple with the Coo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoo

`func (o *QuoteCreateWebhookResponseResourceProductsInner) SetCoo(v string)`

SetCoo sets Coo field to given value.

### HasCoo

`func (o *QuoteCreateWebhookResponseResourceProductsInner) HasCoo() bool`

HasCoo returns a boolean if a field has been set.

### GetIngramPartNumber

`func (o *QuoteCreateWebhookResponseResourceProductsInner) GetIngramPartNumber() string`

GetIngramPartNumber returns the IngramPartNumber field if non-nil, zero value otherwise.

### GetIngramPartNumberOk

`func (o *QuoteCreateWebhookResponseResourceProductsInner) GetIngramPartNumberOk() (*string, bool)`

GetIngramPartNumberOk returns a tuple with the IngramPartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngramPartNumber

`func (o *QuoteCreateWebhookResponseResourceProductsInner) SetIngramPartNumber(v string)`

SetIngramPartNumber sets IngramPartNumber field to given value.

### HasIngramPartNumber

`func (o *QuoteCreateWebhookResponseResourceProductsInner) HasIngramPartNumber() bool`

HasIngramPartNumber returns a boolean if a field has been set.

### GetVendorPartNumber

`func (o *QuoteCreateWebhookResponseResourceProductsInner) GetVendorPartNumber() string`

GetVendorPartNumber returns the VendorPartNumber field if non-nil, zero value otherwise.

### GetVendorPartNumberOk

`func (o *QuoteCreateWebhookResponseResourceProductsInner) GetVendorPartNumberOk() (*string, bool)`

GetVendorPartNumberOk returns a tuple with the VendorPartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorPartNumber

`func (o *QuoteCreateWebhookResponseResourceProductsInner) SetVendorPartNumber(v string)`

SetVendorPartNumber sets VendorPartNumber field to given value.

### HasVendorPartNumber

`func (o *QuoteCreateWebhookResponseResourceProductsInner) HasVendorPartNumber() bool`

HasVendorPartNumber returns a boolean if a field has been set.

### GetDescription

`func (o *QuoteCreateWebhookResponseResourceProductsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *QuoteCreateWebhookResponseResourceProductsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *QuoteCreateWebhookResponseResourceProductsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *QuoteCreateWebhookResponseResourceProductsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetWeight

`func (o *QuoteCreateWebhookResponseResourceProductsInner) GetWeight() string`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *QuoteCreateWebhookResponseResourceProductsInner) GetWeightOk() (*string, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *QuoteCreateWebhookResponseResourceProductsInner) SetWeight(v string)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *QuoteCreateWebhookResponseResourceProductsInner) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetWeightUom

`func (o *QuoteCreateWebhookResponseResourceProductsInner) GetWeightUom() string`

GetWeightUom returns the WeightUom field if non-nil, zero value otherwise.

### GetWeightUomOk

`func (o *QuoteCreateWebhookResponseResourceProductsInner) GetWeightUomOk() (*string, bool)`

GetWeightUomOk returns a tuple with the WeightUom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightUom

`func (o *QuoteCreateWebhookResponseResourceProductsInner) SetWeightUom(v string)`

SetWeightUom sets WeightUom field to given value.

### HasWeightUom

`func (o *QuoteCreateWebhookResponseResourceProductsInner) HasWeightUom() bool`

HasWeightUom returns a boolean if a field has been set.

### GetIsSuggestionProduct

`func (o *QuoteCreateWebhookResponseResourceProductsInner) GetIsSuggestionProduct() string`

GetIsSuggestionProduct returns the IsSuggestionProduct field if non-nil, zero value otherwise.

### GetIsSuggestionProductOk

`func (o *QuoteCreateWebhookResponseResourceProductsInner) GetIsSuggestionProductOk() (*string, bool)`

GetIsSuggestionProductOk returns a tuple with the IsSuggestionProduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSuggestionProduct

`func (o *QuoteCreateWebhookResponseResourceProductsInner) SetIsSuggestionProduct(v string)`

SetIsSuggestionProduct sets IsSuggestionProduct field to given value.

### HasIsSuggestionProduct

`func (o *QuoteCreateWebhookResponseResourceProductsInner) HasIsSuggestionProduct() bool`

HasIsSuggestionProduct returns a boolean if a field has been set.

### GetVpnCategory

`func (o *QuoteCreateWebhookResponseResourceProductsInner) GetVpnCategory() string`

GetVpnCategory returns the VpnCategory field if non-nil, zero value otherwise.

### GetVpnCategoryOk

`func (o *QuoteCreateWebhookResponseResourceProductsInner) GetVpnCategoryOk() (*string, bool)`

GetVpnCategoryOk returns a tuple with the VpnCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnCategory

`func (o *QuoteCreateWebhookResponseResourceProductsInner) SetVpnCategory(v string)`

SetVpnCategory sets VpnCategory field to given value.

### HasVpnCategory

`func (o *QuoteCreateWebhookResponseResourceProductsInner) HasVpnCategory() bool`

HasVpnCategory returns a boolean if a field has been set.

### GetQuoteProductsSupplierPartAuxiliaryId

`func (o *QuoteCreateWebhookResponseResourceProductsInner) GetQuoteProductsSupplierPartAuxiliaryId() string`

GetQuoteProductsSupplierPartAuxiliaryId returns the QuoteProductsSupplierPartAuxiliaryId field if non-nil, zero value otherwise.

### GetQuoteProductsSupplierPartAuxiliaryIdOk

`func (o *QuoteCreateWebhookResponseResourceProductsInner) GetQuoteProductsSupplierPartAuxiliaryIdOk() (*string, bool)`

GetQuoteProductsSupplierPartAuxiliaryIdOk returns a tuple with the QuoteProductsSupplierPartAuxiliaryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteProductsSupplierPartAuxiliaryId

`func (o *QuoteCreateWebhookResponseResourceProductsInner) SetQuoteProductsSupplierPartAuxiliaryId(v string)`

SetQuoteProductsSupplierPartAuxiliaryId sets QuoteProductsSupplierPartAuxiliaryId field to given value.

### HasQuoteProductsSupplierPartAuxiliaryId

`func (o *QuoteCreateWebhookResponseResourceProductsInner) HasQuoteProductsSupplierPartAuxiliaryId() bool`

HasQuoteProductsSupplierPartAuxiliaryId returns a boolean if a field has been set.

### GetVendorName

`func (o *QuoteCreateWebhookResponseResourceProductsInner) GetVendorName() string`

GetVendorName returns the VendorName field if non-nil, zero value otherwise.

### GetVendorNameOk

`func (o *QuoteCreateWebhookResponseResourceProductsInner) GetVendorNameOk() (*string, bool)`

GetVendorNameOk returns a tuple with the VendorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorName

`func (o *QuoteCreateWebhookResponseResourceProductsInner) SetVendorName(v string)`

SetVendorName sets VendorName field to given value.

### HasVendorName

`func (o *QuoteCreateWebhookResponseResourceProductsInner) HasVendorName() bool`

HasVendorName returns a boolean if a field has been set.

### GetTerms

`func (o *QuoteCreateWebhookResponseResourceProductsInner) GetTerms() string`

GetTerms returns the Terms field if non-nil, zero value otherwise.

### GetTermsOk

`func (o *QuoteCreateWebhookResponseResourceProductsInner) GetTermsOk() (*string, bool)`

GetTermsOk returns a tuple with the Terms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerms

`func (o *QuoteCreateWebhookResponseResourceProductsInner) SetTerms(v string)`

SetTerms sets Terms field to given value.

### HasTerms

`func (o *QuoteCreateWebhookResponseResourceProductsInner) HasTerms() bool`

HasTerms returns a boolean if a field has been set.

### GetPrice

`func (o *QuoteCreateWebhookResponseResourceProductsInner) GetPrice() QuoteCreateWebhookResponseResourceProductsInnerPrice`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *QuoteCreateWebhookResponseResourceProductsInner) GetPriceOk() (*QuoteCreateWebhookResponseResourceProductsInnerPrice, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *QuoteCreateWebhookResponseResourceProductsInner) SetPrice(v QuoteCreateWebhookResponseResourceProductsInnerPrice)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *QuoteCreateWebhookResponseResourceProductsInner) HasPrice() bool`

HasPrice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


