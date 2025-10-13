# ProductDetailResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IngramPartNumber** | Pointer to **NullableString** | Ingram Micro unique part number for the product. | [optional] 
**VendorPartNumber** | Pointer to **NullableString** | Vendor’s part number for the product. | [optional] 
**ProductAuthorized** | Pointer to **NullableString** | Boolean that indicates whether a product is authorized. | [optional] 
**Description** | Pointer to **NullableString** | The description given for the product. | [optional] 
**Upc** | Pointer to **NullableString** | The UPC code for the product. Consists of 12 numeric digits that are uniquely assigned to each trade item. | [optional] 
**ProductCategory** | Pointer to **NullableString** | The category of the product. | [optional] 
**ProductSubcategory** | Pointer to **NullableString** | The sub-category of the product. | [optional] 
**VendorName** | Pointer to **NullableString** | Vendor name for the order. | [optional] 
**VendorNumber** | Pointer to **NullableString** | Vendor number that identifies the product. | [optional] 
**ProductStatusCode** | Pointer to **NullableString** | Status code of the product. | [optional] 
**ProductClass** | Pointer to **NullableString** | Indicates whether the product is directly shipped from the vendor’s warehouse or if the product ships from Ingram Micro’s warehouse. Class Codes are Ingram classifications on how skus are stocked A &#x3D; Product that is stocked usually in all IM warehouses and replenished on a regular basis. B &#x3D; Product that is stocked in limited IM warehouses and replenished on a regular basis C &#x3D; Product that is stocked in fewer IM warehouses and replenished on a regular basis. D &#x3D; Product that Ingram Micro has elected to discontinue. E &#x3D; Product that will be phased out later, according to the vendor. You may not want to replenish this product, but instead sell down what is in stock. F &#x3D; Product that we carry for a specific customer or supplier under a contractual agreement. N &#x3D; New Sku. Classification before first receipt O &#x3D; Discontinued product to be liquidated S&#x3D; Order for Specialized Demand (Order to backorder) X&#x3D; direct ship from Vendor V &#x3D; product that vendor has elected to discontinue. | [optional] 
**CustomerPartNumber** | Pointer to **NullableString** | Reseller / end-user’s part number for the product. | [optional] 
**Indicators** | Pointer to [**ProductDetailResponseIndicators**](ProductDetailResponseIndicators.md) |  | [optional] 
**CiscoFields** | Pointer to [**ProductDetailResponseCiscoFields**](ProductDetailResponseCiscoFields.md) |  | [optional] 
**WarrantyInformation** | Pointer to **[]string** | Warranty codes related to the product. | [optional] 
**AdditionalInformation** | Pointer to [**ProductDetailResponseAdditionalInformation**](ProductDetailResponseAdditionalInformation.md) |  | [optional] 
**SubscriptionDetails** | Pointer to [**[]ProductDetailResponseSubscriptionDetailsInner**](ProductDetailResponseSubscriptionDetailsInner.md) | Subscription product Details | [optional] 

## Methods

### NewProductDetailResponse

`func NewProductDetailResponse() *ProductDetailResponse`

NewProductDetailResponse instantiates a new ProductDetailResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductDetailResponseWithDefaults

`func NewProductDetailResponseWithDefaults() *ProductDetailResponse`

NewProductDetailResponseWithDefaults instantiates a new ProductDetailResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIngramPartNumber

`func (o *ProductDetailResponse) GetIngramPartNumber() string`

GetIngramPartNumber returns the IngramPartNumber field if non-nil, zero value otherwise.

### GetIngramPartNumberOk

`func (o *ProductDetailResponse) GetIngramPartNumberOk() (*string, bool)`

GetIngramPartNumberOk returns a tuple with the IngramPartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngramPartNumber

`func (o *ProductDetailResponse) SetIngramPartNumber(v string)`

SetIngramPartNumber sets IngramPartNumber field to given value.

### HasIngramPartNumber

`func (o *ProductDetailResponse) HasIngramPartNumber() bool`

HasIngramPartNumber returns a boolean if a field has been set.

### SetIngramPartNumberNil

`func (o *ProductDetailResponse) SetIngramPartNumberNil(b bool)`

 SetIngramPartNumberNil sets the value for IngramPartNumber to be an explicit nil

### UnsetIngramPartNumber
`func (o *ProductDetailResponse) UnsetIngramPartNumber()`

UnsetIngramPartNumber ensures that no value is present for IngramPartNumber, not even an explicit nil
### GetVendorPartNumber

`func (o *ProductDetailResponse) GetVendorPartNumber() string`

GetVendorPartNumber returns the VendorPartNumber field if non-nil, zero value otherwise.

### GetVendorPartNumberOk

`func (o *ProductDetailResponse) GetVendorPartNumberOk() (*string, bool)`

GetVendorPartNumberOk returns a tuple with the VendorPartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorPartNumber

`func (o *ProductDetailResponse) SetVendorPartNumber(v string)`

SetVendorPartNumber sets VendorPartNumber field to given value.

### HasVendorPartNumber

`func (o *ProductDetailResponse) HasVendorPartNumber() bool`

HasVendorPartNumber returns a boolean if a field has been set.

### SetVendorPartNumberNil

`func (o *ProductDetailResponse) SetVendorPartNumberNil(b bool)`

 SetVendorPartNumberNil sets the value for VendorPartNumber to be an explicit nil

### UnsetVendorPartNumber
`func (o *ProductDetailResponse) UnsetVendorPartNumber()`

UnsetVendorPartNumber ensures that no value is present for VendorPartNumber, not even an explicit nil
### GetProductAuthorized

`func (o *ProductDetailResponse) GetProductAuthorized() string`

GetProductAuthorized returns the ProductAuthorized field if non-nil, zero value otherwise.

### GetProductAuthorizedOk

`func (o *ProductDetailResponse) GetProductAuthorizedOk() (*string, bool)`

GetProductAuthorizedOk returns a tuple with the ProductAuthorized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductAuthorized

`func (o *ProductDetailResponse) SetProductAuthorized(v string)`

SetProductAuthorized sets ProductAuthorized field to given value.

### HasProductAuthorized

`func (o *ProductDetailResponse) HasProductAuthorized() bool`

HasProductAuthorized returns a boolean if a field has been set.

### SetProductAuthorizedNil

`func (o *ProductDetailResponse) SetProductAuthorizedNil(b bool)`

 SetProductAuthorizedNil sets the value for ProductAuthorized to be an explicit nil

### UnsetProductAuthorized
`func (o *ProductDetailResponse) UnsetProductAuthorized()`

UnsetProductAuthorized ensures that no value is present for ProductAuthorized, not even an explicit nil
### GetDescription

`func (o *ProductDetailResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProductDetailResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProductDetailResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProductDetailResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ProductDetailResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ProductDetailResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetUpc

`func (o *ProductDetailResponse) GetUpc() string`

GetUpc returns the Upc field if non-nil, zero value otherwise.

### GetUpcOk

`func (o *ProductDetailResponse) GetUpcOk() (*string, bool)`

GetUpcOk returns a tuple with the Upc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpc

`func (o *ProductDetailResponse) SetUpc(v string)`

SetUpc sets Upc field to given value.

### HasUpc

`func (o *ProductDetailResponse) HasUpc() bool`

HasUpc returns a boolean if a field has been set.

### SetUpcNil

`func (o *ProductDetailResponse) SetUpcNil(b bool)`

 SetUpcNil sets the value for Upc to be an explicit nil

### UnsetUpc
`func (o *ProductDetailResponse) UnsetUpc()`

UnsetUpc ensures that no value is present for Upc, not even an explicit nil
### GetProductCategory

`func (o *ProductDetailResponse) GetProductCategory() string`

GetProductCategory returns the ProductCategory field if non-nil, zero value otherwise.

### GetProductCategoryOk

`func (o *ProductDetailResponse) GetProductCategoryOk() (*string, bool)`

GetProductCategoryOk returns a tuple with the ProductCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductCategory

`func (o *ProductDetailResponse) SetProductCategory(v string)`

SetProductCategory sets ProductCategory field to given value.

### HasProductCategory

`func (o *ProductDetailResponse) HasProductCategory() bool`

HasProductCategory returns a boolean if a field has been set.

### SetProductCategoryNil

`func (o *ProductDetailResponse) SetProductCategoryNil(b bool)`

 SetProductCategoryNil sets the value for ProductCategory to be an explicit nil

### UnsetProductCategory
`func (o *ProductDetailResponse) UnsetProductCategory()`

UnsetProductCategory ensures that no value is present for ProductCategory, not even an explicit nil
### GetProductSubcategory

`func (o *ProductDetailResponse) GetProductSubcategory() string`

GetProductSubcategory returns the ProductSubcategory field if non-nil, zero value otherwise.

### GetProductSubcategoryOk

`func (o *ProductDetailResponse) GetProductSubcategoryOk() (*string, bool)`

GetProductSubcategoryOk returns a tuple with the ProductSubcategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductSubcategory

`func (o *ProductDetailResponse) SetProductSubcategory(v string)`

SetProductSubcategory sets ProductSubcategory field to given value.

### HasProductSubcategory

`func (o *ProductDetailResponse) HasProductSubcategory() bool`

HasProductSubcategory returns a boolean if a field has been set.

### SetProductSubcategoryNil

`func (o *ProductDetailResponse) SetProductSubcategoryNil(b bool)`

 SetProductSubcategoryNil sets the value for ProductSubcategory to be an explicit nil

### UnsetProductSubcategory
`func (o *ProductDetailResponse) UnsetProductSubcategory()`

UnsetProductSubcategory ensures that no value is present for ProductSubcategory, not even an explicit nil
### GetVendorName

`func (o *ProductDetailResponse) GetVendorName() string`

GetVendorName returns the VendorName field if non-nil, zero value otherwise.

### GetVendorNameOk

`func (o *ProductDetailResponse) GetVendorNameOk() (*string, bool)`

GetVendorNameOk returns a tuple with the VendorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorName

`func (o *ProductDetailResponse) SetVendorName(v string)`

SetVendorName sets VendorName field to given value.

### HasVendorName

`func (o *ProductDetailResponse) HasVendorName() bool`

HasVendorName returns a boolean if a field has been set.

### SetVendorNameNil

`func (o *ProductDetailResponse) SetVendorNameNil(b bool)`

 SetVendorNameNil sets the value for VendorName to be an explicit nil

### UnsetVendorName
`func (o *ProductDetailResponse) UnsetVendorName()`

UnsetVendorName ensures that no value is present for VendorName, not even an explicit nil
### GetVendorNumber

`func (o *ProductDetailResponse) GetVendorNumber() string`

GetVendorNumber returns the VendorNumber field if non-nil, zero value otherwise.

### GetVendorNumberOk

`func (o *ProductDetailResponse) GetVendorNumberOk() (*string, bool)`

GetVendorNumberOk returns a tuple with the VendorNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorNumber

`func (o *ProductDetailResponse) SetVendorNumber(v string)`

SetVendorNumber sets VendorNumber field to given value.

### HasVendorNumber

`func (o *ProductDetailResponse) HasVendorNumber() bool`

HasVendorNumber returns a boolean if a field has been set.

### SetVendorNumberNil

`func (o *ProductDetailResponse) SetVendorNumberNil(b bool)`

 SetVendorNumberNil sets the value for VendorNumber to be an explicit nil

### UnsetVendorNumber
`func (o *ProductDetailResponse) UnsetVendorNumber()`

UnsetVendorNumber ensures that no value is present for VendorNumber, not even an explicit nil
### GetProductStatusCode

`func (o *ProductDetailResponse) GetProductStatusCode() string`

GetProductStatusCode returns the ProductStatusCode field if non-nil, zero value otherwise.

### GetProductStatusCodeOk

`func (o *ProductDetailResponse) GetProductStatusCodeOk() (*string, bool)`

GetProductStatusCodeOk returns a tuple with the ProductStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductStatusCode

`func (o *ProductDetailResponse) SetProductStatusCode(v string)`

SetProductStatusCode sets ProductStatusCode field to given value.

### HasProductStatusCode

`func (o *ProductDetailResponse) HasProductStatusCode() bool`

HasProductStatusCode returns a boolean if a field has been set.

### SetProductStatusCodeNil

`func (o *ProductDetailResponse) SetProductStatusCodeNil(b bool)`

 SetProductStatusCodeNil sets the value for ProductStatusCode to be an explicit nil

### UnsetProductStatusCode
`func (o *ProductDetailResponse) UnsetProductStatusCode()`

UnsetProductStatusCode ensures that no value is present for ProductStatusCode, not even an explicit nil
### GetProductClass

`func (o *ProductDetailResponse) GetProductClass() string`

GetProductClass returns the ProductClass field if non-nil, zero value otherwise.

### GetProductClassOk

`func (o *ProductDetailResponse) GetProductClassOk() (*string, bool)`

GetProductClassOk returns a tuple with the ProductClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductClass

`func (o *ProductDetailResponse) SetProductClass(v string)`

SetProductClass sets ProductClass field to given value.

### HasProductClass

`func (o *ProductDetailResponse) HasProductClass() bool`

HasProductClass returns a boolean if a field has been set.

### SetProductClassNil

`func (o *ProductDetailResponse) SetProductClassNil(b bool)`

 SetProductClassNil sets the value for ProductClass to be an explicit nil

### UnsetProductClass
`func (o *ProductDetailResponse) UnsetProductClass()`

UnsetProductClass ensures that no value is present for ProductClass, not even an explicit nil
### GetCustomerPartNumber

`func (o *ProductDetailResponse) GetCustomerPartNumber() string`

GetCustomerPartNumber returns the CustomerPartNumber field if non-nil, zero value otherwise.

### GetCustomerPartNumberOk

`func (o *ProductDetailResponse) GetCustomerPartNumberOk() (*string, bool)`

GetCustomerPartNumberOk returns a tuple with the CustomerPartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerPartNumber

`func (o *ProductDetailResponse) SetCustomerPartNumber(v string)`

SetCustomerPartNumber sets CustomerPartNumber field to given value.

### HasCustomerPartNumber

`func (o *ProductDetailResponse) HasCustomerPartNumber() bool`

HasCustomerPartNumber returns a boolean if a field has been set.

### SetCustomerPartNumberNil

`func (o *ProductDetailResponse) SetCustomerPartNumberNil(b bool)`

 SetCustomerPartNumberNil sets the value for CustomerPartNumber to be an explicit nil

### UnsetCustomerPartNumber
`func (o *ProductDetailResponse) UnsetCustomerPartNumber()`

UnsetCustomerPartNumber ensures that no value is present for CustomerPartNumber, not even an explicit nil
### GetIndicators

`func (o *ProductDetailResponse) GetIndicators() ProductDetailResponseIndicators`

GetIndicators returns the Indicators field if non-nil, zero value otherwise.

### GetIndicatorsOk

`func (o *ProductDetailResponse) GetIndicatorsOk() (*ProductDetailResponseIndicators, bool)`

GetIndicatorsOk returns a tuple with the Indicators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndicators

`func (o *ProductDetailResponse) SetIndicators(v ProductDetailResponseIndicators)`

SetIndicators sets Indicators field to given value.

### HasIndicators

`func (o *ProductDetailResponse) HasIndicators() bool`

HasIndicators returns a boolean if a field has been set.

### GetCiscoFields

`func (o *ProductDetailResponse) GetCiscoFields() ProductDetailResponseCiscoFields`

GetCiscoFields returns the CiscoFields field if non-nil, zero value otherwise.

### GetCiscoFieldsOk

`func (o *ProductDetailResponse) GetCiscoFieldsOk() (*ProductDetailResponseCiscoFields, bool)`

GetCiscoFieldsOk returns a tuple with the CiscoFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiscoFields

`func (o *ProductDetailResponse) SetCiscoFields(v ProductDetailResponseCiscoFields)`

SetCiscoFields sets CiscoFields field to given value.

### HasCiscoFields

`func (o *ProductDetailResponse) HasCiscoFields() bool`

HasCiscoFields returns a boolean if a field has been set.

### GetWarrantyInformation

`func (o *ProductDetailResponse) GetWarrantyInformation() []string`

GetWarrantyInformation returns the WarrantyInformation field if non-nil, zero value otherwise.

### GetWarrantyInformationOk

`func (o *ProductDetailResponse) GetWarrantyInformationOk() (*[]string, bool)`

GetWarrantyInformationOk returns a tuple with the WarrantyInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarrantyInformation

`func (o *ProductDetailResponse) SetWarrantyInformation(v []string)`

SetWarrantyInformation sets WarrantyInformation field to given value.

### HasWarrantyInformation

`func (o *ProductDetailResponse) HasWarrantyInformation() bool`

HasWarrantyInformation returns a boolean if a field has been set.

### GetAdditionalInformation

`func (o *ProductDetailResponse) GetAdditionalInformation() ProductDetailResponseAdditionalInformation`

GetAdditionalInformation returns the AdditionalInformation field if non-nil, zero value otherwise.

### GetAdditionalInformationOk

`func (o *ProductDetailResponse) GetAdditionalInformationOk() (*ProductDetailResponseAdditionalInformation, bool)`

GetAdditionalInformationOk returns a tuple with the AdditionalInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInformation

`func (o *ProductDetailResponse) SetAdditionalInformation(v ProductDetailResponseAdditionalInformation)`

SetAdditionalInformation sets AdditionalInformation field to given value.

### HasAdditionalInformation

`func (o *ProductDetailResponse) HasAdditionalInformation() bool`

HasAdditionalInformation returns a boolean if a field has been set.

### GetSubscriptionDetails

`func (o *ProductDetailResponse) GetSubscriptionDetails() []ProductDetailResponseSubscriptionDetailsInner`

GetSubscriptionDetails returns the SubscriptionDetails field if non-nil, zero value otherwise.

### GetSubscriptionDetailsOk

`func (o *ProductDetailResponse) GetSubscriptionDetailsOk() (*[]ProductDetailResponseSubscriptionDetailsInner, bool)`

GetSubscriptionDetailsOk returns a tuple with the SubscriptionDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionDetails

`func (o *ProductDetailResponse) SetSubscriptionDetails(v []ProductDetailResponseSubscriptionDetailsInner)`

SetSubscriptionDetails sets SubscriptionDetails field to given value.

### HasSubscriptionDetails

`func (o *ProductDetailResponse) HasSubscriptionDetails() bool`

HasSubscriptionDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


