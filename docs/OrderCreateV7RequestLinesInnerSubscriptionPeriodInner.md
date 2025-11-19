# OrderCreateV7RequestLinesInnerSubscriptionPeriodInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Unit period of the subscription. Example, Years, Months | [optional] 
**Duration** | Pointer to **int32** | Length of the subscription. Example 1, 3 | [optional] 
**StartDate** | Pointer to **string** | The date on which subscription will start. | [optional] 
**EndDateAlignmentType** | Pointer to **string** | Subscription period end date alignment. ENUM -- &#39;MATCH_END_OF_CALENDAR_MONTH&#39;, &#39;CO_TERM_ON_SUBSCRIPTION&#39; | [optional] 
**SubscriptionId** | Pointer to **string** | The ID of an existing active subscription. | [optional] 

## Methods

### NewOrderCreateV7RequestLinesInnerSubscriptionPeriodInner

`func NewOrderCreateV7RequestLinesInnerSubscriptionPeriodInner() *OrderCreateV7RequestLinesInnerSubscriptionPeriodInner`

NewOrderCreateV7RequestLinesInnerSubscriptionPeriodInner instantiates a new OrderCreateV7RequestLinesInnerSubscriptionPeriodInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderCreateV7RequestLinesInnerSubscriptionPeriodInnerWithDefaults

`func NewOrderCreateV7RequestLinesInnerSubscriptionPeriodInnerWithDefaults() *OrderCreateV7RequestLinesInnerSubscriptionPeriodInner`

NewOrderCreateV7RequestLinesInnerSubscriptionPeriodInnerWithDefaults instantiates a new OrderCreateV7RequestLinesInnerSubscriptionPeriodInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *OrderCreateV7RequestLinesInnerSubscriptionPeriodInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OrderCreateV7RequestLinesInnerSubscriptionPeriodInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OrderCreateV7RequestLinesInnerSubscriptionPeriodInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OrderCreateV7RequestLinesInnerSubscriptionPeriodInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDuration

`func (o *OrderCreateV7RequestLinesInnerSubscriptionPeriodInner) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *OrderCreateV7RequestLinesInnerSubscriptionPeriodInner) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *OrderCreateV7RequestLinesInnerSubscriptionPeriodInner) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *OrderCreateV7RequestLinesInnerSubscriptionPeriodInner) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetStartDate

`func (o *OrderCreateV7RequestLinesInnerSubscriptionPeriodInner) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *OrderCreateV7RequestLinesInnerSubscriptionPeriodInner) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *OrderCreateV7RequestLinesInnerSubscriptionPeriodInner) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *OrderCreateV7RequestLinesInnerSubscriptionPeriodInner) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDateAlignmentType

`func (o *OrderCreateV7RequestLinesInnerSubscriptionPeriodInner) GetEndDateAlignmentType() string`

GetEndDateAlignmentType returns the EndDateAlignmentType field if non-nil, zero value otherwise.

### GetEndDateAlignmentTypeOk

`func (o *OrderCreateV7RequestLinesInnerSubscriptionPeriodInner) GetEndDateAlignmentTypeOk() (*string, bool)`

GetEndDateAlignmentTypeOk returns a tuple with the EndDateAlignmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDateAlignmentType

`func (o *OrderCreateV7RequestLinesInnerSubscriptionPeriodInner) SetEndDateAlignmentType(v string)`

SetEndDateAlignmentType sets EndDateAlignmentType field to given value.

### HasEndDateAlignmentType

`func (o *OrderCreateV7RequestLinesInnerSubscriptionPeriodInner) HasEndDateAlignmentType() bool`

HasEndDateAlignmentType returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *OrderCreateV7RequestLinesInnerSubscriptionPeriodInner) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *OrderCreateV7RequestLinesInnerSubscriptionPeriodInner) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *OrderCreateV7RequestLinesInnerSubscriptionPeriodInner) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *OrderCreateV7RequestLinesInnerSubscriptionPeriodInner) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


