# PriceAndAvailabilityResponseInnerSubscriptionPriceInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | Pointer to **NullableFloat32** |  | [optional] 
**PlanId** | Pointer to **NullableString** | Id of the plan. | [optional] 
**PlanUId** | Pointer to **NullableString** |  | [optional] 
**PlanName** | Pointer to **NullableString** | Name of the plan. | [optional] 
**PlanDescription** | Pointer to **NullableString** | The description of the plan. | [optional] 
**Groups** | Pointer to [**[]PriceAndAvailabilityResponseInnerSubscriptionPriceInnerGroupsInner**](PriceAndAvailabilityResponseInnerSubscriptionPriceInnerGroupsInner.md) |  | [optional] 
**BillingPeriod** | Pointer to [**PriceAndAvailabilityResponseInnerSubscriptionPriceInnerBillingPeriod**](PriceAndAvailabilityResponseInnerSubscriptionPriceInnerBillingPeriod.md) |  | [optional] 
**SubscriptionPeriod** | Pointer to [**[]PriceAndAvailabilityResponseInnerSubscriptionPriceInnerSubscriptionPeriodInner**](PriceAndAvailabilityResponseInnerSubscriptionPriceInnerSubscriptionPeriodInner.md) |  | [optional] 
**Options** | Pointer to [**[]PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInner**](PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInner.md) |  | [optional] 

## Methods

### NewPriceAndAvailabilityResponseInnerSubscriptionPriceInner

`func NewPriceAndAvailabilityResponseInnerSubscriptionPriceInner() *PriceAndAvailabilityResponseInnerSubscriptionPriceInner`

NewPriceAndAvailabilityResponseInnerSubscriptionPriceInner instantiates a new PriceAndAvailabilityResponseInnerSubscriptionPriceInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceAndAvailabilityResponseInnerSubscriptionPriceInnerWithDefaults

`func NewPriceAndAvailabilityResponseInnerSubscriptionPriceInnerWithDefaults() *PriceAndAvailabilityResponseInnerSubscriptionPriceInner`

NewPriceAndAvailabilityResponseInnerSubscriptionPriceInnerWithDefaults instantiates a new PriceAndAvailabilityResponseInnerSubscriptionPriceInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInner) GetIndex() float32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInner) GetIndexOk() (*float32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInner) SetIndex(v float32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInner) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### SetIndexNil

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInner) SetIndexNil(b bool)`

 SetIndexNil sets the value for Index to be an explicit nil

### UnsetIndex
`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInner) UnsetIndex()`

UnsetIndex ensures that no value is present for Index, not even an explicit nil
### GetPlanId

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInner) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInner) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInner) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInner) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### SetPlanIdNil

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInner) SetPlanIdNil(b bool)`

 SetPlanIdNil sets the value for PlanId to be an explicit nil

### UnsetPlanId
`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInner) UnsetPlanId()`

UnsetPlanId ensures that no value is present for PlanId, not even an explicit nil
### GetPlanUId

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInner) GetPlanUId() string`

GetPlanUId returns the PlanUId field if non-nil, zero value otherwise.

### GetPlanUIdOk

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInner) GetPlanUIdOk() (*string, bool)`

GetPlanUIdOk returns a tuple with the PlanUId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanUId

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInner) SetPlanUId(v string)`

SetPlanUId sets PlanUId field to given value.

### HasPlanUId

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInner) HasPlanUId() bool`

HasPlanUId returns a boolean if a field has been set.

### SetPlanUIdNil

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInner) SetPlanUIdNil(b bool)`

 SetPlanUIdNil sets the value for PlanUId to be an explicit nil

### UnsetPlanUId
`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInner) UnsetPlanUId()`

UnsetPlanUId ensures that no value is present for PlanUId, not even an explicit nil
### GetPlanName

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInner) GetPlanName() string`

GetPlanName returns the PlanName field if non-nil, zero value otherwise.

### GetPlanNameOk

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInner) GetPlanNameOk() (*string, bool)`

GetPlanNameOk returns a tuple with the PlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanName

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInner) SetPlanName(v string)`

SetPlanName sets PlanName field to given value.

### HasPlanName

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInner) HasPlanName() bool`

HasPlanName returns a boolean if a field has been set.

### SetPlanNameNil

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInner) SetPlanNameNil(b bool)`

 SetPlanNameNil sets the value for PlanName to be an explicit nil

### UnsetPlanName
`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInner) UnsetPlanName()`

UnsetPlanName ensures that no value is present for PlanName, not even an explicit nil
### GetPlanDescription

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInner) GetPlanDescription() string`

GetPlanDescription returns the PlanDescription field if non-nil, zero value otherwise.

### GetPlanDescriptionOk

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInner) GetPlanDescriptionOk() (*string, bool)`

GetPlanDescriptionOk returns a tuple with the PlanDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanDescription

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInner) SetPlanDescription(v string)`

SetPlanDescription sets PlanDescription field to given value.

### HasPlanDescription

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInner) HasPlanDescription() bool`

HasPlanDescription returns a boolean if a field has been set.

### SetPlanDescriptionNil

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInner) SetPlanDescriptionNil(b bool)`

 SetPlanDescriptionNil sets the value for PlanDescription to be an explicit nil

### UnsetPlanDescription
`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInner) UnsetPlanDescription()`

UnsetPlanDescription ensures that no value is present for PlanDescription, not even an explicit nil
### GetGroups

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInner) GetGroups() []PriceAndAvailabilityResponseInnerSubscriptionPriceInnerGroupsInner`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInner) GetGroupsOk() (*[]PriceAndAvailabilityResponseInnerSubscriptionPriceInnerGroupsInner, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInner) SetGroups(v []PriceAndAvailabilityResponseInnerSubscriptionPriceInnerGroupsInner)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInner) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetBillingPeriod

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInner) GetBillingPeriod() PriceAndAvailabilityResponseInnerSubscriptionPriceInnerBillingPeriod`

GetBillingPeriod returns the BillingPeriod field if non-nil, zero value otherwise.

### GetBillingPeriodOk

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInner) GetBillingPeriodOk() (*PriceAndAvailabilityResponseInnerSubscriptionPriceInnerBillingPeriod, bool)`

GetBillingPeriodOk returns a tuple with the BillingPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingPeriod

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInner) SetBillingPeriod(v PriceAndAvailabilityResponseInnerSubscriptionPriceInnerBillingPeriod)`

SetBillingPeriod sets BillingPeriod field to given value.

### HasBillingPeriod

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInner) HasBillingPeriod() bool`

HasBillingPeriod returns a boolean if a field has been set.

### GetSubscriptionPeriod

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInner) GetSubscriptionPeriod() []PriceAndAvailabilityResponseInnerSubscriptionPriceInnerSubscriptionPeriodInner`

GetSubscriptionPeriod returns the SubscriptionPeriod field if non-nil, zero value otherwise.

### GetSubscriptionPeriodOk

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInner) GetSubscriptionPeriodOk() (*[]PriceAndAvailabilityResponseInnerSubscriptionPriceInnerSubscriptionPeriodInner, bool)`

GetSubscriptionPeriodOk returns a tuple with the SubscriptionPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionPeriod

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInner) SetSubscriptionPeriod(v []PriceAndAvailabilityResponseInnerSubscriptionPriceInnerSubscriptionPeriodInner)`

SetSubscriptionPeriod sets SubscriptionPeriod field to given value.

### HasSubscriptionPeriod

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInner) HasSubscriptionPeriod() bool`

HasSubscriptionPeriod returns a boolean if a field has been set.

### GetOptions

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInner) GetOptions() []PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInner`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInner) GetOptionsOk() (*[]PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInner, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInner) SetOptions(v []PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInner)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInner) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


