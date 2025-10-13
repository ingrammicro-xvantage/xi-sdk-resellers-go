# ProductDetailResponseSubscriptionDetailsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlanId** | Pointer to **string** | ID of the subscription plan. | [optional] 
**PlanName** | Pointer to **string** | Name of the subscription plan. | [optional] 
**PlanDescription** | Pointer to **string** | Description of the subscription plan. | [optional] 
**Groups** | Pointer to [**[]ProductDetailResponseSubscriptionDetailsInnerGroupsInner**](ProductDetailResponseSubscriptionDetailsInnerGroupsInner.md) | Details of the groups subscription product part of. | [optional] 
**SubscriptionPeriod** | Pointer to [**[]ProductDetailResponseSubscriptionDetailsInnerSubscriptionPeriodInner**](ProductDetailResponseSubscriptionDetailsInnerSubscriptionPeriodInner.md) | Details of the subscription period. | [optional] 
**BillingPeriod** | Pointer to [**ProductDetailResponseSubscriptionDetailsInnerBillingPeriod**](ProductDetailResponseSubscriptionDetailsInnerBillingPeriod.md) |  | [optional] 
**Options** | Pointer to [**[]ProductDetailResponseSubscriptionDetailsInnerOptionsInner**](ProductDetailResponseSubscriptionDetailsInnerOptionsInner.md) | Details of the resources available. | [optional] 
**Links** | Pointer to [**[]ProductSearchResponseSubscriptionCatalogInnerPlansInnerLinksInner**](ProductSearchResponseSubscriptionCatalogInnerPlansInnerLinksInner.md) |  | [optional] 
**NextPage** | Pointer to **string** | link/URL for accessing next page. | [optional] 
**PreviousPage** | Pointer to **string** | link/URL for accessing previous page. | [optional] 

## Methods

### NewProductDetailResponseSubscriptionDetailsInner

`func NewProductDetailResponseSubscriptionDetailsInner() *ProductDetailResponseSubscriptionDetailsInner`

NewProductDetailResponseSubscriptionDetailsInner instantiates a new ProductDetailResponseSubscriptionDetailsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductDetailResponseSubscriptionDetailsInnerWithDefaults

`func NewProductDetailResponseSubscriptionDetailsInnerWithDefaults() *ProductDetailResponseSubscriptionDetailsInner`

NewProductDetailResponseSubscriptionDetailsInnerWithDefaults instantiates a new ProductDetailResponseSubscriptionDetailsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlanId

`func (o *ProductDetailResponseSubscriptionDetailsInner) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *ProductDetailResponseSubscriptionDetailsInner) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *ProductDetailResponseSubscriptionDetailsInner) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *ProductDetailResponseSubscriptionDetailsInner) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetPlanName

`func (o *ProductDetailResponseSubscriptionDetailsInner) GetPlanName() string`

GetPlanName returns the PlanName field if non-nil, zero value otherwise.

### GetPlanNameOk

`func (o *ProductDetailResponseSubscriptionDetailsInner) GetPlanNameOk() (*string, bool)`

GetPlanNameOk returns a tuple with the PlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanName

`func (o *ProductDetailResponseSubscriptionDetailsInner) SetPlanName(v string)`

SetPlanName sets PlanName field to given value.

### HasPlanName

`func (o *ProductDetailResponseSubscriptionDetailsInner) HasPlanName() bool`

HasPlanName returns a boolean if a field has been set.

### GetPlanDescription

`func (o *ProductDetailResponseSubscriptionDetailsInner) GetPlanDescription() string`

GetPlanDescription returns the PlanDescription field if non-nil, zero value otherwise.

### GetPlanDescriptionOk

`func (o *ProductDetailResponseSubscriptionDetailsInner) GetPlanDescriptionOk() (*string, bool)`

GetPlanDescriptionOk returns a tuple with the PlanDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanDescription

`func (o *ProductDetailResponseSubscriptionDetailsInner) SetPlanDescription(v string)`

SetPlanDescription sets PlanDescription field to given value.

### HasPlanDescription

`func (o *ProductDetailResponseSubscriptionDetailsInner) HasPlanDescription() bool`

HasPlanDescription returns a boolean if a field has been set.

### GetGroups

`func (o *ProductDetailResponseSubscriptionDetailsInner) GetGroups() []ProductDetailResponseSubscriptionDetailsInnerGroupsInner`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *ProductDetailResponseSubscriptionDetailsInner) GetGroupsOk() (*[]ProductDetailResponseSubscriptionDetailsInnerGroupsInner, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *ProductDetailResponseSubscriptionDetailsInner) SetGroups(v []ProductDetailResponseSubscriptionDetailsInnerGroupsInner)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *ProductDetailResponseSubscriptionDetailsInner) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetSubscriptionPeriod

`func (o *ProductDetailResponseSubscriptionDetailsInner) GetSubscriptionPeriod() []ProductDetailResponseSubscriptionDetailsInnerSubscriptionPeriodInner`

GetSubscriptionPeriod returns the SubscriptionPeriod field if non-nil, zero value otherwise.

### GetSubscriptionPeriodOk

`func (o *ProductDetailResponseSubscriptionDetailsInner) GetSubscriptionPeriodOk() (*[]ProductDetailResponseSubscriptionDetailsInnerSubscriptionPeriodInner, bool)`

GetSubscriptionPeriodOk returns a tuple with the SubscriptionPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionPeriod

`func (o *ProductDetailResponseSubscriptionDetailsInner) SetSubscriptionPeriod(v []ProductDetailResponseSubscriptionDetailsInnerSubscriptionPeriodInner)`

SetSubscriptionPeriod sets SubscriptionPeriod field to given value.

### HasSubscriptionPeriod

`func (o *ProductDetailResponseSubscriptionDetailsInner) HasSubscriptionPeriod() bool`

HasSubscriptionPeriod returns a boolean if a field has been set.

### GetBillingPeriod

`func (o *ProductDetailResponseSubscriptionDetailsInner) GetBillingPeriod() ProductDetailResponseSubscriptionDetailsInnerBillingPeriod`

GetBillingPeriod returns the BillingPeriod field if non-nil, zero value otherwise.

### GetBillingPeriodOk

`func (o *ProductDetailResponseSubscriptionDetailsInner) GetBillingPeriodOk() (*ProductDetailResponseSubscriptionDetailsInnerBillingPeriod, bool)`

GetBillingPeriodOk returns a tuple with the BillingPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingPeriod

`func (o *ProductDetailResponseSubscriptionDetailsInner) SetBillingPeriod(v ProductDetailResponseSubscriptionDetailsInnerBillingPeriod)`

SetBillingPeriod sets BillingPeriod field to given value.

### HasBillingPeriod

`func (o *ProductDetailResponseSubscriptionDetailsInner) HasBillingPeriod() bool`

HasBillingPeriod returns a boolean if a field has been set.

### GetOptions

`func (o *ProductDetailResponseSubscriptionDetailsInner) GetOptions() []ProductDetailResponseSubscriptionDetailsInnerOptionsInner`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ProductDetailResponseSubscriptionDetailsInner) GetOptionsOk() (*[]ProductDetailResponseSubscriptionDetailsInnerOptionsInner, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ProductDetailResponseSubscriptionDetailsInner) SetOptions(v []ProductDetailResponseSubscriptionDetailsInnerOptionsInner)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *ProductDetailResponseSubscriptionDetailsInner) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetLinks

`func (o *ProductDetailResponseSubscriptionDetailsInner) GetLinks() []ProductSearchResponseSubscriptionCatalogInnerPlansInnerLinksInner`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ProductDetailResponseSubscriptionDetailsInner) GetLinksOk() (*[]ProductSearchResponseSubscriptionCatalogInnerPlansInnerLinksInner, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ProductDetailResponseSubscriptionDetailsInner) SetLinks(v []ProductSearchResponseSubscriptionCatalogInnerPlansInnerLinksInner)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ProductDetailResponseSubscriptionDetailsInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetNextPage

`func (o *ProductDetailResponseSubscriptionDetailsInner) GetNextPage() string`

GetNextPage returns the NextPage field if non-nil, zero value otherwise.

### GetNextPageOk

`func (o *ProductDetailResponseSubscriptionDetailsInner) GetNextPageOk() (*string, bool)`

GetNextPageOk returns a tuple with the NextPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPage

`func (o *ProductDetailResponseSubscriptionDetailsInner) SetNextPage(v string)`

SetNextPage sets NextPage field to given value.

### HasNextPage

`func (o *ProductDetailResponseSubscriptionDetailsInner) HasNextPage() bool`

HasNextPage returns a boolean if a field has been set.

### GetPreviousPage

`func (o *ProductDetailResponseSubscriptionDetailsInner) GetPreviousPage() string`

GetPreviousPage returns the PreviousPage field if non-nil, zero value otherwise.

### GetPreviousPageOk

`func (o *ProductDetailResponseSubscriptionDetailsInner) GetPreviousPageOk() (*string, bool)`

GetPreviousPageOk returns a tuple with the PreviousPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousPage

`func (o *ProductDetailResponseSubscriptionDetailsInner) SetPreviousPage(v string)`

SetPreviousPage sets PreviousPage field to given value.

### HasPreviousPage

`func (o *ProductDetailResponseSubscriptionDetailsInner) HasPreviousPage() bool`

HasPreviousPage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


