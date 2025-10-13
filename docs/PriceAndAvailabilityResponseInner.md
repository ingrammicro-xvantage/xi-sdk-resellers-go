# PriceAndAvailabilityResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | Pointer to **float32** |  | [optional] 
**ProductStatusCode** | Pointer to **NullableString** | Codes signifying whether the sku is active or not. | [optional] 
**ProductStatusMessage** | Pointer to **NullableString** | Message returned saying whether sku is active. | [optional] 
**IngramPartNumber** | Pointer to **NullableString** | Ingram Micro unique part number for the product. | [optional] 
**VendorPartNumber** | Pointer to **NullableString** | Vendor’s part number for the product. | [optional] 
**ExtendedVendorPartNumber** | Pointer to **NullableString** | Extended Vendor Part Number. *Currently, this feature is not available in these countries (Mexico, Turkey, New Zealand, Colombia, Chile, Brazil, Peru, Western Sahara). | [optional] 
**CustomerPartNumber** | Pointer to **NullableString** | Reseller / end-user’s part number for the product. | [optional] 
**Upc** | Pointer to **NullableString** | The UPC code for the product. Consists of 12 numeric digits that are uniquely assigned to each trade item. | [optional] 
**PartNumberType** | Pointer to **NullableString** | Number type of the part. | [optional] 
**VendorNumber** | Pointer to **NullableString** | Vendor number that identifies the product. | [optional] 
**VendorName** | Pointer to **NullableString** | Vendor name for the order. | [optional] 
**Description** | Pointer to **NullableString** | The description given for the product. | [optional] 
**ProductClass** | Pointer to **NullableString** | Indicates whether the product is directly shipped from the vendor’s warehouse or if the product ships from Ingram Micro’s warehouse. Class Codes are Ingram classifications on how skus are stocked A &#x3D; Product that is stocked usually in all IM warehouses and replenished on a regular basis. B &#x3D; Product that is stocked in limited IM warehouses and replenished on a regular basis C &#x3D; Product that is stocked in fewer IM warehouses and replenished on a regular basis. D &#x3D; Product that Ingram Micro has elected to discontinue. E &#x3D; Product that will be phased out later, according to the vendor. You may not want to replenish this product, but instead sell down what is in stock. F &#x3D; Product that we carry for a specific customer or supplier under a contractual agreement. N &#x3D; New Sku. Classification before first receipt O &#x3D; Discontinued product to be liquidated S&#x3D; Order for Specialized Demand (Order to backorder) X&#x3D; direct ship from Vendor V &#x3D; product that vendor has elected to discontinue. | [optional] 
**Uom** | Pointer to **NullableString** | The description given for the product. | [optional] 
**ProductStatus** | Pointer to **NullableString** | Status that gives whether the product is Active. | [optional] 
**AcceptBackOrder** | Pointer to **NullableBool** | Boolean that indicates if the product accepts backorder. | [optional] 
**ProductAuthorized** | Pointer to **NullableBool** | Boolean that indicates whether a product is authorized. | [optional] 
**ReturnableProduct** | Pointer to **NullableBool** | Boolean that indicates if the product can be returned. | [optional] 
**EndUserInfoRequired** | Pointer to **NullableBool** | Boolean that indicates  if end user information is required. | [optional] 
**GovtSpecialPriceAvailable** | Pointer to **NullableBool** | Boolean that indicates if special pricing is available for the product. | [optional] 
**GovtProgramType** | Pointer to **NullableString** | Program type, “PA” for government orders, “ED” for education order. | [optional] 
**GovtEndUserType** | Pointer to **NullableString** | Type of end user of the program. F &#x3D; Federal, S &#x3D; State, E &#x3D; Local, K &#x3D; K-12 school, and H &#x3D; Higher Education. | [optional] 
**Availability** | Pointer to [**NullablePriceAndAvailabilityResponseInnerAvailability**](PriceAndAvailabilityResponseInnerAvailability.md) |  | [optional] 
**ReserveInventoryDetails** | Pointer to [**[]PriceAndAvailabilityResponseInnerReserveInventoryDetailsInner**](PriceAndAvailabilityResponseInnerReserveInventoryDetailsInner.md) |  | [optional] 
**Pricing** | Pointer to [**NullablePriceAndAvailabilityResponseInnerPricing**](PriceAndAvailabilityResponseInnerPricing.md) |  | [optional] 
**Discounts** | Pointer to [**[]PriceAndAvailabilityResponseInnerDiscountsInner**](PriceAndAvailabilityResponseInnerDiscountsInner.md) |  | [optional] 
**BundlePartIndicator** | Pointer to **NullableBool** | True of false value to indicate whether it’s bundle part. *Currently, this feature is not available in these countries (Mexico, Turkey, New Zealand, Colombia, Chile, Brazil, Peru, Western Sahara). | [optional] 
**ServiceFees** | Pointer to [**[]PriceAndAvailabilityResponseInnerServiceFeesInner**](PriceAndAvailabilityResponseInnerServiceFeesInner.md) | *Currently, this feature is not available in these countries (Mexico, Turkey, New Zealand, Colombia, Chile, Brazil, Peru, Western Sahara). | [optional] 
**SubscriptionPrice** | Pointer to [**[]PriceAndAvailabilityResponseInnerSubscriptionPriceInner**](PriceAndAvailabilityResponseInnerSubscriptionPriceInner.md) |  | [optional] 

## Methods

### NewPriceAndAvailabilityResponseInner

`func NewPriceAndAvailabilityResponseInner() *PriceAndAvailabilityResponseInner`

NewPriceAndAvailabilityResponseInner instantiates a new PriceAndAvailabilityResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceAndAvailabilityResponseInnerWithDefaults

`func NewPriceAndAvailabilityResponseInnerWithDefaults() *PriceAndAvailabilityResponseInner`

NewPriceAndAvailabilityResponseInnerWithDefaults instantiates a new PriceAndAvailabilityResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *PriceAndAvailabilityResponseInner) GetIndex() float32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *PriceAndAvailabilityResponseInner) GetIndexOk() (*float32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *PriceAndAvailabilityResponseInner) SetIndex(v float32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *PriceAndAvailabilityResponseInner) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetProductStatusCode

`func (o *PriceAndAvailabilityResponseInner) GetProductStatusCode() string`

GetProductStatusCode returns the ProductStatusCode field if non-nil, zero value otherwise.

### GetProductStatusCodeOk

`func (o *PriceAndAvailabilityResponseInner) GetProductStatusCodeOk() (*string, bool)`

GetProductStatusCodeOk returns a tuple with the ProductStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductStatusCode

`func (o *PriceAndAvailabilityResponseInner) SetProductStatusCode(v string)`

SetProductStatusCode sets ProductStatusCode field to given value.

### HasProductStatusCode

`func (o *PriceAndAvailabilityResponseInner) HasProductStatusCode() bool`

HasProductStatusCode returns a boolean if a field has been set.

### SetProductStatusCodeNil

`func (o *PriceAndAvailabilityResponseInner) SetProductStatusCodeNil(b bool)`

 SetProductStatusCodeNil sets the value for ProductStatusCode to be an explicit nil

### UnsetProductStatusCode
`func (o *PriceAndAvailabilityResponseInner) UnsetProductStatusCode()`

UnsetProductStatusCode ensures that no value is present for ProductStatusCode, not even an explicit nil
### GetProductStatusMessage

`func (o *PriceAndAvailabilityResponseInner) GetProductStatusMessage() string`

GetProductStatusMessage returns the ProductStatusMessage field if non-nil, zero value otherwise.

### GetProductStatusMessageOk

`func (o *PriceAndAvailabilityResponseInner) GetProductStatusMessageOk() (*string, bool)`

GetProductStatusMessageOk returns a tuple with the ProductStatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductStatusMessage

`func (o *PriceAndAvailabilityResponseInner) SetProductStatusMessage(v string)`

SetProductStatusMessage sets ProductStatusMessage field to given value.

### HasProductStatusMessage

`func (o *PriceAndAvailabilityResponseInner) HasProductStatusMessage() bool`

HasProductStatusMessage returns a boolean if a field has been set.

### SetProductStatusMessageNil

`func (o *PriceAndAvailabilityResponseInner) SetProductStatusMessageNil(b bool)`

 SetProductStatusMessageNil sets the value for ProductStatusMessage to be an explicit nil

### UnsetProductStatusMessage
`func (o *PriceAndAvailabilityResponseInner) UnsetProductStatusMessage()`

UnsetProductStatusMessage ensures that no value is present for ProductStatusMessage, not even an explicit nil
### GetIngramPartNumber

`func (o *PriceAndAvailabilityResponseInner) GetIngramPartNumber() string`

GetIngramPartNumber returns the IngramPartNumber field if non-nil, zero value otherwise.

### GetIngramPartNumberOk

`func (o *PriceAndAvailabilityResponseInner) GetIngramPartNumberOk() (*string, bool)`

GetIngramPartNumberOk returns a tuple with the IngramPartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngramPartNumber

`func (o *PriceAndAvailabilityResponseInner) SetIngramPartNumber(v string)`

SetIngramPartNumber sets IngramPartNumber field to given value.

### HasIngramPartNumber

`func (o *PriceAndAvailabilityResponseInner) HasIngramPartNumber() bool`

HasIngramPartNumber returns a boolean if a field has been set.

### SetIngramPartNumberNil

`func (o *PriceAndAvailabilityResponseInner) SetIngramPartNumberNil(b bool)`

 SetIngramPartNumberNil sets the value for IngramPartNumber to be an explicit nil

### UnsetIngramPartNumber
`func (o *PriceAndAvailabilityResponseInner) UnsetIngramPartNumber()`

UnsetIngramPartNumber ensures that no value is present for IngramPartNumber, not even an explicit nil
### GetVendorPartNumber

`func (o *PriceAndAvailabilityResponseInner) GetVendorPartNumber() string`

GetVendorPartNumber returns the VendorPartNumber field if non-nil, zero value otherwise.

### GetVendorPartNumberOk

`func (o *PriceAndAvailabilityResponseInner) GetVendorPartNumberOk() (*string, bool)`

GetVendorPartNumberOk returns a tuple with the VendorPartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorPartNumber

`func (o *PriceAndAvailabilityResponseInner) SetVendorPartNumber(v string)`

SetVendorPartNumber sets VendorPartNumber field to given value.

### HasVendorPartNumber

`func (o *PriceAndAvailabilityResponseInner) HasVendorPartNumber() bool`

HasVendorPartNumber returns a boolean if a field has been set.

### SetVendorPartNumberNil

`func (o *PriceAndAvailabilityResponseInner) SetVendorPartNumberNil(b bool)`

 SetVendorPartNumberNil sets the value for VendorPartNumber to be an explicit nil

### UnsetVendorPartNumber
`func (o *PriceAndAvailabilityResponseInner) UnsetVendorPartNumber()`

UnsetVendorPartNumber ensures that no value is present for VendorPartNumber, not even an explicit nil
### GetExtendedVendorPartNumber

`func (o *PriceAndAvailabilityResponseInner) GetExtendedVendorPartNumber() string`

GetExtendedVendorPartNumber returns the ExtendedVendorPartNumber field if non-nil, zero value otherwise.

### GetExtendedVendorPartNumberOk

`func (o *PriceAndAvailabilityResponseInner) GetExtendedVendorPartNumberOk() (*string, bool)`

GetExtendedVendorPartNumberOk returns a tuple with the ExtendedVendorPartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedVendorPartNumber

`func (o *PriceAndAvailabilityResponseInner) SetExtendedVendorPartNumber(v string)`

SetExtendedVendorPartNumber sets ExtendedVendorPartNumber field to given value.

### HasExtendedVendorPartNumber

`func (o *PriceAndAvailabilityResponseInner) HasExtendedVendorPartNumber() bool`

HasExtendedVendorPartNumber returns a boolean if a field has been set.

### SetExtendedVendorPartNumberNil

`func (o *PriceAndAvailabilityResponseInner) SetExtendedVendorPartNumberNil(b bool)`

 SetExtendedVendorPartNumberNil sets the value for ExtendedVendorPartNumber to be an explicit nil

### UnsetExtendedVendorPartNumber
`func (o *PriceAndAvailabilityResponseInner) UnsetExtendedVendorPartNumber()`

UnsetExtendedVendorPartNumber ensures that no value is present for ExtendedVendorPartNumber, not even an explicit nil
### GetCustomerPartNumber

`func (o *PriceAndAvailabilityResponseInner) GetCustomerPartNumber() string`

GetCustomerPartNumber returns the CustomerPartNumber field if non-nil, zero value otherwise.

### GetCustomerPartNumberOk

`func (o *PriceAndAvailabilityResponseInner) GetCustomerPartNumberOk() (*string, bool)`

GetCustomerPartNumberOk returns a tuple with the CustomerPartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerPartNumber

`func (o *PriceAndAvailabilityResponseInner) SetCustomerPartNumber(v string)`

SetCustomerPartNumber sets CustomerPartNumber field to given value.

### HasCustomerPartNumber

`func (o *PriceAndAvailabilityResponseInner) HasCustomerPartNumber() bool`

HasCustomerPartNumber returns a boolean if a field has been set.

### SetCustomerPartNumberNil

`func (o *PriceAndAvailabilityResponseInner) SetCustomerPartNumberNil(b bool)`

 SetCustomerPartNumberNil sets the value for CustomerPartNumber to be an explicit nil

### UnsetCustomerPartNumber
`func (o *PriceAndAvailabilityResponseInner) UnsetCustomerPartNumber()`

UnsetCustomerPartNumber ensures that no value is present for CustomerPartNumber, not even an explicit nil
### GetUpc

`func (o *PriceAndAvailabilityResponseInner) GetUpc() string`

GetUpc returns the Upc field if non-nil, zero value otherwise.

### GetUpcOk

`func (o *PriceAndAvailabilityResponseInner) GetUpcOk() (*string, bool)`

GetUpcOk returns a tuple with the Upc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpc

`func (o *PriceAndAvailabilityResponseInner) SetUpc(v string)`

SetUpc sets Upc field to given value.

### HasUpc

`func (o *PriceAndAvailabilityResponseInner) HasUpc() bool`

HasUpc returns a boolean if a field has been set.

### SetUpcNil

`func (o *PriceAndAvailabilityResponseInner) SetUpcNil(b bool)`

 SetUpcNil sets the value for Upc to be an explicit nil

### UnsetUpc
`func (o *PriceAndAvailabilityResponseInner) UnsetUpc()`

UnsetUpc ensures that no value is present for Upc, not even an explicit nil
### GetPartNumberType

`func (o *PriceAndAvailabilityResponseInner) GetPartNumberType() string`

GetPartNumberType returns the PartNumberType field if non-nil, zero value otherwise.

### GetPartNumberTypeOk

`func (o *PriceAndAvailabilityResponseInner) GetPartNumberTypeOk() (*string, bool)`

GetPartNumberTypeOk returns a tuple with the PartNumberType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumberType

`func (o *PriceAndAvailabilityResponseInner) SetPartNumberType(v string)`

SetPartNumberType sets PartNumberType field to given value.

### HasPartNumberType

`func (o *PriceAndAvailabilityResponseInner) HasPartNumberType() bool`

HasPartNumberType returns a boolean if a field has been set.

### SetPartNumberTypeNil

`func (o *PriceAndAvailabilityResponseInner) SetPartNumberTypeNil(b bool)`

 SetPartNumberTypeNil sets the value for PartNumberType to be an explicit nil

### UnsetPartNumberType
`func (o *PriceAndAvailabilityResponseInner) UnsetPartNumberType()`

UnsetPartNumberType ensures that no value is present for PartNumberType, not even an explicit nil
### GetVendorNumber

`func (o *PriceAndAvailabilityResponseInner) GetVendorNumber() string`

GetVendorNumber returns the VendorNumber field if non-nil, zero value otherwise.

### GetVendorNumberOk

`func (o *PriceAndAvailabilityResponseInner) GetVendorNumberOk() (*string, bool)`

GetVendorNumberOk returns a tuple with the VendorNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorNumber

`func (o *PriceAndAvailabilityResponseInner) SetVendorNumber(v string)`

SetVendorNumber sets VendorNumber field to given value.

### HasVendorNumber

`func (o *PriceAndAvailabilityResponseInner) HasVendorNumber() bool`

HasVendorNumber returns a boolean if a field has been set.

### SetVendorNumberNil

`func (o *PriceAndAvailabilityResponseInner) SetVendorNumberNil(b bool)`

 SetVendorNumberNil sets the value for VendorNumber to be an explicit nil

### UnsetVendorNumber
`func (o *PriceAndAvailabilityResponseInner) UnsetVendorNumber()`

UnsetVendorNumber ensures that no value is present for VendorNumber, not even an explicit nil
### GetVendorName

`func (o *PriceAndAvailabilityResponseInner) GetVendorName() string`

GetVendorName returns the VendorName field if non-nil, zero value otherwise.

### GetVendorNameOk

`func (o *PriceAndAvailabilityResponseInner) GetVendorNameOk() (*string, bool)`

GetVendorNameOk returns a tuple with the VendorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorName

`func (o *PriceAndAvailabilityResponseInner) SetVendorName(v string)`

SetVendorName sets VendorName field to given value.

### HasVendorName

`func (o *PriceAndAvailabilityResponseInner) HasVendorName() bool`

HasVendorName returns a boolean if a field has been set.

### SetVendorNameNil

`func (o *PriceAndAvailabilityResponseInner) SetVendorNameNil(b bool)`

 SetVendorNameNil sets the value for VendorName to be an explicit nil

### UnsetVendorName
`func (o *PriceAndAvailabilityResponseInner) UnsetVendorName()`

UnsetVendorName ensures that no value is present for VendorName, not even an explicit nil
### GetDescription

`func (o *PriceAndAvailabilityResponseInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PriceAndAvailabilityResponseInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PriceAndAvailabilityResponseInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PriceAndAvailabilityResponseInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PriceAndAvailabilityResponseInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PriceAndAvailabilityResponseInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetProductClass

`func (o *PriceAndAvailabilityResponseInner) GetProductClass() string`

GetProductClass returns the ProductClass field if non-nil, zero value otherwise.

### GetProductClassOk

`func (o *PriceAndAvailabilityResponseInner) GetProductClassOk() (*string, bool)`

GetProductClassOk returns a tuple with the ProductClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductClass

`func (o *PriceAndAvailabilityResponseInner) SetProductClass(v string)`

SetProductClass sets ProductClass field to given value.

### HasProductClass

`func (o *PriceAndAvailabilityResponseInner) HasProductClass() bool`

HasProductClass returns a boolean if a field has been set.

### SetProductClassNil

`func (o *PriceAndAvailabilityResponseInner) SetProductClassNil(b bool)`

 SetProductClassNil sets the value for ProductClass to be an explicit nil

### UnsetProductClass
`func (o *PriceAndAvailabilityResponseInner) UnsetProductClass()`

UnsetProductClass ensures that no value is present for ProductClass, not even an explicit nil
### GetUom

`func (o *PriceAndAvailabilityResponseInner) GetUom() string`

GetUom returns the Uom field if non-nil, zero value otherwise.

### GetUomOk

`func (o *PriceAndAvailabilityResponseInner) GetUomOk() (*string, bool)`

GetUomOk returns a tuple with the Uom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUom

`func (o *PriceAndAvailabilityResponseInner) SetUom(v string)`

SetUom sets Uom field to given value.

### HasUom

`func (o *PriceAndAvailabilityResponseInner) HasUom() bool`

HasUom returns a boolean if a field has been set.

### SetUomNil

`func (o *PriceAndAvailabilityResponseInner) SetUomNil(b bool)`

 SetUomNil sets the value for Uom to be an explicit nil

### UnsetUom
`func (o *PriceAndAvailabilityResponseInner) UnsetUom()`

UnsetUom ensures that no value is present for Uom, not even an explicit nil
### GetProductStatus

`func (o *PriceAndAvailabilityResponseInner) GetProductStatus() string`

GetProductStatus returns the ProductStatus field if non-nil, zero value otherwise.

### GetProductStatusOk

`func (o *PriceAndAvailabilityResponseInner) GetProductStatusOk() (*string, bool)`

GetProductStatusOk returns a tuple with the ProductStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductStatus

`func (o *PriceAndAvailabilityResponseInner) SetProductStatus(v string)`

SetProductStatus sets ProductStatus field to given value.

### HasProductStatus

`func (o *PriceAndAvailabilityResponseInner) HasProductStatus() bool`

HasProductStatus returns a boolean if a field has been set.

### SetProductStatusNil

`func (o *PriceAndAvailabilityResponseInner) SetProductStatusNil(b bool)`

 SetProductStatusNil sets the value for ProductStatus to be an explicit nil

### UnsetProductStatus
`func (o *PriceAndAvailabilityResponseInner) UnsetProductStatus()`

UnsetProductStatus ensures that no value is present for ProductStatus, not even an explicit nil
### GetAcceptBackOrder

`func (o *PriceAndAvailabilityResponseInner) GetAcceptBackOrder() bool`

GetAcceptBackOrder returns the AcceptBackOrder field if non-nil, zero value otherwise.

### GetAcceptBackOrderOk

`func (o *PriceAndAvailabilityResponseInner) GetAcceptBackOrderOk() (*bool, bool)`

GetAcceptBackOrderOk returns a tuple with the AcceptBackOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptBackOrder

`func (o *PriceAndAvailabilityResponseInner) SetAcceptBackOrder(v bool)`

SetAcceptBackOrder sets AcceptBackOrder field to given value.

### HasAcceptBackOrder

`func (o *PriceAndAvailabilityResponseInner) HasAcceptBackOrder() bool`

HasAcceptBackOrder returns a boolean if a field has been set.

### SetAcceptBackOrderNil

`func (o *PriceAndAvailabilityResponseInner) SetAcceptBackOrderNil(b bool)`

 SetAcceptBackOrderNil sets the value for AcceptBackOrder to be an explicit nil

### UnsetAcceptBackOrder
`func (o *PriceAndAvailabilityResponseInner) UnsetAcceptBackOrder()`

UnsetAcceptBackOrder ensures that no value is present for AcceptBackOrder, not even an explicit nil
### GetProductAuthorized

`func (o *PriceAndAvailabilityResponseInner) GetProductAuthorized() bool`

GetProductAuthorized returns the ProductAuthorized field if non-nil, zero value otherwise.

### GetProductAuthorizedOk

`func (o *PriceAndAvailabilityResponseInner) GetProductAuthorizedOk() (*bool, bool)`

GetProductAuthorizedOk returns a tuple with the ProductAuthorized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductAuthorized

`func (o *PriceAndAvailabilityResponseInner) SetProductAuthorized(v bool)`

SetProductAuthorized sets ProductAuthorized field to given value.

### HasProductAuthorized

`func (o *PriceAndAvailabilityResponseInner) HasProductAuthorized() bool`

HasProductAuthorized returns a boolean if a field has been set.

### SetProductAuthorizedNil

`func (o *PriceAndAvailabilityResponseInner) SetProductAuthorizedNil(b bool)`

 SetProductAuthorizedNil sets the value for ProductAuthorized to be an explicit nil

### UnsetProductAuthorized
`func (o *PriceAndAvailabilityResponseInner) UnsetProductAuthorized()`

UnsetProductAuthorized ensures that no value is present for ProductAuthorized, not even an explicit nil
### GetReturnableProduct

`func (o *PriceAndAvailabilityResponseInner) GetReturnableProduct() bool`

GetReturnableProduct returns the ReturnableProduct field if non-nil, zero value otherwise.

### GetReturnableProductOk

`func (o *PriceAndAvailabilityResponseInner) GetReturnableProductOk() (*bool, bool)`

GetReturnableProductOk returns a tuple with the ReturnableProduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnableProduct

`func (o *PriceAndAvailabilityResponseInner) SetReturnableProduct(v bool)`

SetReturnableProduct sets ReturnableProduct field to given value.

### HasReturnableProduct

`func (o *PriceAndAvailabilityResponseInner) HasReturnableProduct() bool`

HasReturnableProduct returns a boolean if a field has been set.

### SetReturnableProductNil

`func (o *PriceAndAvailabilityResponseInner) SetReturnableProductNil(b bool)`

 SetReturnableProductNil sets the value for ReturnableProduct to be an explicit nil

### UnsetReturnableProduct
`func (o *PriceAndAvailabilityResponseInner) UnsetReturnableProduct()`

UnsetReturnableProduct ensures that no value is present for ReturnableProduct, not even an explicit nil
### GetEndUserInfoRequired

`func (o *PriceAndAvailabilityResponseInner) GetEndUserInfoRequired() bool`

GetEndUserInfoRequired returns the EndUserInfoRequired field if non-nil, zero value otherwise.

### GetEndUserInfoRequiredOk

`func (o *PriceAndAvailabilityResponseInner) GetEndUserInfoRequiredOk() (*bool, bool)`

GetEndUserInfoRequiredOk returns a tuple with the EndUserInfoRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndUserInfoRequired

`func (o *PriceAndAvailabilityResponseInner) SetEndUserInfoRequired(v bool)`

SetEndUserInfoRequired sets EndUserInfoRequired field to given value.

### HasEndUserInfoRequired

`func (o *PriceAndAvailabilityResponseInner) HasEndUserInfoRequired() bool`

HasEndUserInfoRequired returns a boolean if a field has been set.

### SetEndUserInfoRequiredNil

`func (o *PriceAndAvailabilityResponseInner) SetEndUserInfoRequiredNil(b bool)`

 SetEndUserInfoRequiredNil sets the value for EndUserInfoRequired to be an explicit nil

### UnsetEndUserInfoRequired
`func (o *PriceAndAvailabilityResponseInner) UnsetEndUserInfoRequired()`

UnsetEndUserInfoRequired ensures that no value is present for EndUserInfoRequired, not even an explicit nil
### GetGovtSpecialPriceAvailable

`func (o *PriceAndAvailabilityResponseInner) GetGovtSpecialPriceAvailable() bool`

GetGovtSpecialPriceAvailable returns the GovtSpecialPriceAvailable field if non-nil, zero value otherwise.

### GetGovtSpecialPriceAvailableOk

`func (o *PriceAndAvailabilityResponseInner) GetGovtSpecialPriceAvailableOk() (*bool, bool)`

GetGovtSpecialPriceAvailableOk returns a tuple with the GovtSpecialPriceAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGovtSpecialPriceAvailable

`func (o *PriceAndAvailabilityResponseInner) SetGovtSpecialPriceAvailable(v bool)`

SetGovtSpecialPriceAvailable sets GovtSpecialPriceAvailable field to given value.

### HasGovtSpecialPriceAvailable

`func (o *PriceAndAvailabilityResponseInner) HasGovtSpecialPriceAvailable() bool`

HasGovtSpecialPriceAvailable returns a boolean if a field has been set.

### SetGovtSpecialPriceAvailableNil

`func (o *PriceAndAvailabilityResponseInner) SetGovtSpecialPriceAvailableNil(b bool)`

 SetGovtSpecialPriceAvailableNil sets the value for GovtSpecialPriceAvailable to be an explicit nil

### UnsetGovtSpecialPriceAvailable
`func (o *PriceAndAvailabilityResponseInner) UnsetGovtSpecialPriceAvailable()`

UnsetGovtSpecialPriceAvailable ensures that no value is present for GovtSpecialPriceAvailable, not even an explicit nil
### GetGovtProgramType

`func (o *PriceAndAvailabilityResponseInner) GetGovtProgramType() string`

GetGovtProgramType returns the GovtProgramType field if non-nil, zero value otherwise.

### GetGovtProgramTypeOk

`func (o *PriceAndAvailabilityResponseInner) GetGovtProgramTypeOk() (*string, bool)`

GetGovtProgramTypeOk returns a tuple with the GovtProgramType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGovtProgramType

`func (o *PriceAndAvailabilityResponseInner) SetGovtProgramType(v string)`

SetGovtProgramType sets GovtProgramType field to given value.

### HasGovtProgramType

`func (o *PriceAndAvailabilityResponseInner) HasGovtProgramType() bool`

HasGovtProgramType returns a boolean if a field has been set.

### SetGovtProgramTypeNil

`func (o *PriceAndAvailabilityResponseInner) SetGovtProgramTypeNil(b bool)`

 SetGovtProgramTypeNil sets the value for GovtProgramType to be an explicit nil

### UnsetGovtProgramType
`func (o *PriceAndAvailabilityResponseInner) UnsetGovtProgramType()`

UnsetGovtProgramType ensures that no value is present for GovtProgramType, not even an explicit nil
### GetGovtEndUserType

`func (o *PriceAndAvailabilityResponseInner) GetGovtEndUserType() string`

GetGovtEndUserType returns the GovtEndUserType field if non-nil, zero value otherwise.

### GetGovtEndUserTypeOk

`func (o *PriceAndAvailabilityResponseInner) GetGovtEndUserTypeOk() (*string, bool)`

GetGovtEndUserTypeOk returns a tuple with the GovtEndUserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGovtEndUserType

`func (o *PriceAndAvailabilityResponseInner) SetGovtEndUserType(v string)`

SetGovtEndUserType sets GovtEndUserType field to given value.

### HasGovtEndUserType

`func (o *PriceAndAvailabilityResponseInner) HasGovtEndUserType() bool`

HasGovtEndUserType returns a boolean if a field has been set.

### SetGovtEndUserTypeNil

`func (o *PriceAndAvailabilityResponseInner) SetGovtEndUserTypeNil(b bool)`

 SetGovtEndUserTypeNil sets the value for GovtEndUserType to be an explicit nil

### UnsetGovtEndUserType
`func (o *PriceAndAvailabilityResponseInner) UnsetGovtEndUserType()`

UnsetGovtEndUserType ensures that no value is present for GovtEndUserType, not even an explicit nil
### GetAvailability

`func (o *PriceAndAvailabilityResponseInner) GetAvailability() PriceAndAvailabilityResponseInnerAvailability`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *PriceAndAvailabilityResponseInner) GetAvailabilityOk() (*PriceAndAvailabilityResponseInnerAvailability, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *PriceAndAvailabilityResponseInner) SetAvailability(v PriceAndAvailabilityResponseInnerAvailability)`

SetAvailability sets Availability field to given value.

### HasAvailability

`func (o *PriceAndAvailabilityResponseInner) HasAvailability() bool`

HasAvailability returns a boolean if a field has been set.

### SetAvailabilityNil

`func (o *PriceAndAvailabilityResponseInner) SetAvailabilityNil(b bool)`

 SetAvailabilityNil sets the value for Availability to be an explicit nil

### UnsetAvailability
`func (o *PriceAndAvailabilityResponseInner) UnsetAvailability()`

UnsetAvailability ensures that no value is present for Availability, not even an explicit nil
### GetReserveInventoryDetails

`func (o *PriceAndAvailabilityResponseInner) GetReserveInventoryDetails() []PriceAndAvailabilityResponseInnerReserveInventoryDetailsInner`

GetReserveInventoryDetails returns the ReserveInventoryDetails field if non-nil, zero value otherwise.

### GetReserveInventoryDetailsOk

`func (o *PriceAndAvailabilityResponseInner) GetReserveInventoryDetailsOk() (*[]PriceAndAvailabilityResponseInnerReserveInventoryDetailsInner, bool)`

GetReserveInventoryDetailsOk returns a tuple with the ReserveInventoryDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReserveInventoryDetails

`func (o *PriceAndAvailabilityResponseInner) SetReserveInventoryDetails(v []PriceAndAvailabilityResponseInnerReserveInventoryDetailsInner)`

SetReserveInventoryDetails sets ReserveInventoryDetails field to given value.

### HasReserveInventoryDetails

`func (o *PriceAndAvailabilityResponseInner) HasReserveInventoryDetails() bool`

HasReserveInventoryDetails returns a boolean if a field has been set.

### GetPricing

`func (o *PriceAndAvailabilityResponseInner) GetPricing() PriceAndAvailabilityResponseInnerPricing`

GetPricing returns the Pricing field if non-nil, zero value otherwise.

### GetPricingOk

`func (o *PriceAndAvailabilityResponseInner) GetPricingOk() (*PriceAndAvailabilityResponseInnerPricing, bool)`

GetPricingOk returns a tuple with the Pricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricing

`func (o *PriceAndAvailabilityResponseInner) SetPricing(v PriceAndAvailabilityResponseInnerPricing)`

SetPricing sets Pricing field to given value.

### HasPricing

`func (o *PriceAndAvailabilityResponseInner) HasPricing() bool`

HasPricing returns a boolean if a field has been set.

### SetPricingNil

`func (o *PriceAndAvailabilityResponseInner) SetPricingNil(b bool)`

 SetPricingNil sets the value for Pricing to be an explicit nil

### UnsetPricing
`func (o *PriceAndAvailabilityResponseInner) UnsetPricing()`

UnsetPricing ensures that no value is present for Pricing, not even an explicit nil
### GetDiscounts

`func (o *PriceAndAvailabilityResponseInner) GetDiscounts() []PriceAndAvailabilityResponseInnerDiscountsInner`

GetDiscounts returns the Discounts field if non-nil, zero value otherwise.

### GetDiscountsOk

`func (o *PriceAndAvailabilityResponseInner) GetDiscountsOk() (*[]PriceAndAvailabilityResponseInnerDiscountsInner, bool)`

GetDiscountsOk returns a tuple with the Discounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscounts

`func (o *PriceAndAvailabilityResponseInner) SetDiscounts(v []PriceAndAvailabilityResponseInnerDiscountsInner)`

SetDiscounts sets Discounts field to given value.

### HasDiscounts

`func (o *PriceAndAvailabilityResponseInner) HasDiscounts() bool`

HasDiscounts returns a boolean if a field has been set.

### SetDiscountsNil

`func (o *PriceAndAvailabilityResponseInner) SetDiscountsNil(b bool)`

 SetDiscountsNil sets the value for Discounts to be an explicit nil

### UnsetDiscounts
`func (o *PriceAndAvailabilityResponseInner) UnsetDiscounts()`

UnsetDiscounts ensures that no value is present for Discounts, not even an explicit nil
### GetBundlePartIndicator

`func (o *PriceAndAvailabilityResponseInner) GetBundlePartIndicator() bool`

GetBundlePartIndicator returns the BundlePartIndicator field if non-nil, zero value otherwise.

### GetBundlePartIndicatorOk

`func (o *PriceAndAvailabilityResponseInner) GetBundlePartIndicatorOk() (*bool, bool)`

GetBundlePartIndicatorOk returns a tuple with the BundlePartIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundlePartIndicator

`func (o *PriceAndAvailabilityResponseInner) SetBundlePartIndicator(v bool)`

SetBundlePartIndicator sets BundlePartIndicator field to given value.

### HasBundlePartIndicator

`func (o *PriceAndAvailabilityResponseInner) HasBundlePartIndicator() bool`

HasBundlePartIndicator returns a boolean if a field has been set.

### SetBundlePartIndicatorNil

`func (o *PriceAndAvailabilityResponseInner) SetBundlePartIndicatorNil(b bool)`

 SetBundlePartIndicatorNil sets the value for BundlePartIndicator to be an explicit nil

### UnsetBundlePartIndicator
`func (o *PriceAndAvailabilityResponseInner) UnsetBundlePartIndicator()`

UnsetBundlePartIndicator ensures that no value is present for BundlePartIndicator, not even an explicit nil
### GetServiceFees

`func (o *PriceAndAvailabilityResponseInner) GetServiceFees() []PriceAndAvailabilityResponseInnerServiceFeesInner`

GetServiceFees returns the ServiceFees field if non-nil, zero value otherwise.

### GetServiceFeesOk

`func (o *PriceAndAvailabilityResponseInner) GetServiceFeesOk() (*[]PriceAndAvailabilityResponseInnerServiceFeesInner, bool)`

GetServiceFeesOk returns a tuple with the ServiceFees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceFees

`func (o *PriceAndAvailabilityResponseInner) SetServiceFees(v []PriceAndAvailabilityResponseInnerServiceFeesInner)`

SetServiceFees sets ServiceFees field to given value.

### HasServiceFees

`func (o *PriceAndAvailabilityResponseInner) HasServiceFees() bool`

HasServiceFees returns a boolean if a field has been set.

### SetServiceFeesNil

`func (o *PriceAndAvailabilityResponseInner) SetServiceFeesNil(b bool)`

 SetServiceFeesNil sets the value for ServiceFees to be an explicit nil

### UnsetServiceFees
`func (o *PriceAndAvailabilityResponseInner) UnsetServiceFees()`

UnsetServiceFees ensures that no value is present for ServiceFees, not even an explicit nil
### GetSubscriptionPrice

`func (o *PriceAndAvailabilityResponseInner) GetSubscriptionPrice() []PriceAndAvailabilityResponseInnerSubscriptionPriceInner`

GetSubscriptionPrice returns the SubscriptionPrice field if non-nil, zero value otherwise.

### GetSubscriptionPriceOk

`func (o *PriceAndAvailabilityResponseInner) GetSubscriptionPriceOk() (*[]PriceAndAvailabilityResponseInnerSubscriptionPriceInner, bool)`

GetSubscriptionPriceOk returns a tuple with the SubscriptionPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionPrice

`func (o *PriceAndAvailabilityResponseInner) SetSubscriptionPrice(v []PriceAndAvailabilityResponseInnerSubscriptionPriceInner)`

SetSubscriptionPrice sets SubscriptionPrice field to given value.

### HasSubscriptionPrice

`func (o *PriceAndAvailabilityResponseInner) HasSubscriptionPrice() bool`

HasSubscriptionPrice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


