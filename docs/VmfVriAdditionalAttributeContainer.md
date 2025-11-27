# VmfVriAdditionalAttributeContainer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VendorName** | Pointer to **string** | The name of vendor. | [optional] 
**ProductId** | Pointer to **string** | The ID of product. | [optional] 
**AdditionalAttributes** | Pointer to [**[]AdditionalAttribute**](AdditionalAttribute.md) | List of required attributes for the specific product. | [optional] 

## Methods

### NewVmfVriAdditionalAttributeContainer

`func NewVmfVriAdditionalAttributeContainer() *VmfVriAdditionalAttributeContainer`

NewVmfVriAdditionalAttributeContainer instantiates a new VmfVriAdditionalAttributeContainer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmfVriAdditionalAttributeContainerWithDefaults

`func NewVmfVriAdditionalAttributeContainerWithDefaults() *VmfVriAdditionalAttributeContainer`

NewVmfVriAdditionalAttributeContainerWithDefaults instantiates a new VmfVriAdditionalAttributeContainer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVendorName

`func (o *VmfVriAdditionalAttributeContainer) GetVendorName() string`

GetVendorName returns the VendorName field if non-nil, zero value otherwise.

### GetVendorNameOk

`func (o *VmfVriAdditionalAttributeContainer) GetVendorNameOk() (*string, bool)`

GetVendorNameOk returns a tuple with the VendorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorName

`func (o *VmfVriAdditionalAttributeContainer) SetVendorName(v string)`

SetVendorName sets VendorName field to given value.

### HasVendorName

`func (o *VmfVriAdditionalAttributeContainer) HasVendorName() bool`

HasVendorName returns a boolean if a field has been set.

### GetProductId

`func (o *VmfVriAdditionalAttributeContainer) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *VmfVriAdditionalAttributeContainer) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *VmfVriAdditionalAttributeContainer) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *VmfVriAdditionalAttributeContainer) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetAdditionalAttributes

`func (o *VmfVriAdditionalAttributeContainer) GetAdditionalAttributes() []AdditionalAttribute`

GetAdditionalAttributes returns the AdditionalAttributes field if non-nil, zero value otherwise.

### GetAdditionalAttributesOk

`func (o *VmfVriAdditionalAttributeContainer) GetAdditionalAttributesOk() (*[]AdditionalAttribute, bool)`

GetAdditionalAttributesOk returns a tuple with the AdditionalAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalAttributes

`func (o *VmfVriAdditionalAttributeContainer) SetAdditionalAttributes(v []AdditionalAttribute)`

SetAdditionalAttributes sets AdditionalAttributes field to given value.

### HasAdditionalAttributes

`func (o *VmfVriAdditionalAttributeContainer) HasAdditionalAttributes() bool`

HasAdditionalAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


