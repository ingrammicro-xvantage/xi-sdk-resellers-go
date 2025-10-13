# ProductSearchResponseSubscriptionCatalogInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupName** | Pointer to **string** | The name of the group. (The info will shown if searched with showGroupInfo&#x3D;true) | [optional] 
**GroupDescription** | Pointer to **string** | The description of the group. (The info will shown if searched with showGroupInfo&#x3D;true) | [optional] 
**NumberOfPlans** | Pointer to **string** | The number of plans available for the group. (The info will shown if searched with showGroupInfo&#x3D;true) | [optional] 
**Link** | Pointer to **string** | URL to access more details about the group. (The info will shown if searched with showGroupInfo&#x3D;true) | [optional] 
**Plans** | Pointer to [**[]ProductSearchResponseSubscriptionCatalogInnerPlansInner**](ProductSearchResponseSubscriptionCatalogInnerPlansInner.md) |  | [optional] 

## Methods

### NewProductSearchResponseSubscriptionCatalogInner

`func NewProductSearchResponseSubscriptionCatalogInner() *ProductSearchResponseSubscriptionCatalogInner`

NewProductSearchResponseSubscriptionCatalogInner instantiates a new ProductSearchResponseSubscriptionCatalogInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductSearchResponseSubscriptionCatalogInnerWithDefaults

`func NewProductSearchResponseSubscriptionCatalogInnerWithDefaults() *ProductSearchResponseSubscriptionCatalogInner`

NewProductSearchResponseSubscriptionCatalogInnerWithDefaults instantiates a new ProductSearchResponseSubscriptionCatalogInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupName

`func (o *ProductSearchResponseSubscriptionCatalogInner) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *ProductSearchResponseSubscriptionCatalogInner) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *ProductSearchResponseSubscriptionCatalogInner) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *ProductSearchResponseSubscriptionCatalogInner) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetGroupDescription

`func (o *ProductSearchResponseSubscriptionCatalogInner) GetGroupDescription() string`

GetGroupDescription returns the GroupDescription field if non-nil, zero value otherwise.

### GetGroupDescriptionOk

`func (o *ProductSearchResponseSubscriptionCatalogInner) GetGroupDescriptionOk() (*string, bool)`

GetGroupDescriptionOk returns a tuple with the GroupDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupDescription

`func (o *ProductSearchResponseSubscriptionCatalogInner) SetGroupDescription(v string)`

SetGroupDescription sets GroupDescription field to given value.

### HasGroupDescription

`func (o *ProductSearchResponseSubscriptionCatalogInner) HasGroupDescription() bool`

HasGroupDescription returns a boolean if a field has been set.

### GetNumberOfPlans

`func (o *ProductSearchResponseSubscriptionCatalogInner) GetNumberOfPlans() string`

GetNumberOfPlans returns the NumberOfPlans field if non-nil, zero value otherwise.

### GetNumberOfPlansOk

`func (o *ProductSearchResponseSubscriptionCatalogInner) GetNumberOfPlansOk() (*string, bool)`

GetNumberOfPlansOk returns a tuple with the NumberOfPlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfPlans

`func (o *ProductSearchResponseSubscriptionCatalogInner) SetNumberOfPlans(v string)`

SetNumberOfPlans sets NumberOfPlans field to given value.

### HasNumberOfPlans

`func (o *ProductSearchResponseSubscriptionCatalogInner) HasNumberOfPlans() bool`

HasNumberOfPlans returns a boolean if a field has been set.

### GetLink

`func (o *ProductSearchResponseSubscriptionCatalogInner) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *ProductSearchResponseSubscriptionCatalogInner) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *ProductSearchResponseSubscriptionCatalogInner) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *ProductSearchResponseSubscriptionCatalogInner) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetPlans

`func (o *ProductSearchResponseSubscriptionCatalogInner) GetPlans() []ProductSearchResponseSubscriptionCatalogInnerPlansInner`

GetPlans returns the Plans field if non-nil, zero value otherwise.

### GetPlansOk

`func (o *ProductSearchResponseSubscriptionCatalogInner) GetPlansOk() (*[]ProductSearchResponseSubscriptionCatalogInnerPlansInner, bool)`

GetPlansOk returns a tuple with the Plans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlans

`func (o *ProductSearchResponseSubscriptionCatalogInner) SetPlans(v []ProductSearchResponseSubscriptionCatalogInnerPlansInner)`

SetPlans sets Plans field to given value.

### HasPlans

`func (o *ProductSearchResponseSubscriptionCatalogInner) HasPlans() bool`

HasPlans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


