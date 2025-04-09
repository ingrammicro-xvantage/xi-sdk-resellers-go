# QuoteCreateWebhookResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Topic** | Pointer to **string** | Field for identifying whether it is a reseller or vendor event. For eg, resellers/orders | [optional] 
**Event** | Pointer to **string** | The event sent in the request. For eg, im::create. | [optional] 
**EventTimeStamp** | Pointer to **string** | The timestamp at which the event was sent. | [optional] 
**EventId** | Pointer to **string** | A unique id used as identifier for the sepcific event and used for generating the x-hub signature. | [optional] 
**Resource** | Pointer to [**QuoteCreateWebhookResponseResource**](QuoteCreateWebhookResponseResource.md) |  | [optional] 

## Methods

### NewQuoteCreateWebhookResponse

`func NewQuoteCreateWebhookResponse() *QuoteCreateWebhookResponse`

NewQuoteCreateWebhookResponse instantiates a new QuoteCreateWebhookResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuoteCreateWebhookResponseWithDefaults

`func NewQuoteCreateWebhookResponseWithDefaults() *QuoteCreateWebhookResponse`

NewQuoteCreateWebhookResponseWithDefaults instantiates a new QuoteCreateWebhookResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTopic

`func (o *QuoteCreateWebhookResponse) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *QuoteCreateWebhookResponse) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *QuoteCreateWebhookResponse) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *QuoteCreateWebhookResponse) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### GetEvent

`func (o *QuoteCreateWebhookResponse) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *QuoteCreateWebhookResponse) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *QuoteCreateWebhookResponse) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *QuoteCreateWebhookResponse) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetEventTimeStamp

`func (o *QuoteCreateWebhookResponse) GetEventTimeStamp() string`

GetEventTimeStamp returns the EventTimeStamp field if non-nil, zero value otherwise.

### GetEventTimeStampOk

`func (o *QuoteCreateWebhookResponse) GetEventTimeStampOk() (*string, bool)`

GetEventTimeStampOk returns a tuple with the EventTimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTimeStamp

`func (o *QuoteCreateWebhookResponse) SetEventTimeStamp(v string)`

SetEventTimeStamp sets EventTimeStamp field to given value.

### HasEventTimeStamp

`func (o *QuoteCreateWebhookResponse) HasEventTimeStamp() bool`

HasEventTimeStamp returns a boolean if a field has been set.

### GetEventId

`func (o *QuoteCreateWebhookResponse) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *QuoteCreateWebhookResponse) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *QuoteCreateWebhookResponse) SetEventId(v string)`

SetEventId sets EventId field to given value.

### HasEventId

`func (o *QuoteCreateWebhookResponse) HasEventId() bool`

HasEventId returns a boolean if a field has been set.

### GetResource

`func (o *QuoteCreateWebhookResponse) GetResource() QuoteCreateWebhookResponseResource`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *QuoteCreateWebhookResponse) GetResourceOk() (*QuoteCreateWebhookResponseResource, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *QuoteCreateWebhookResponse) SetResource(v QuoteCreateWebhookResponseResource)`

SetResource sets Resource field to given value.

### HasResource

`func (o *QuoteCreateWebhookResponse) HasResource() bool`

HasResource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


