# OrderCreateV7Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Topic** | Pointer to **string** | Field for identifying whether it is a reseller or vendor event. For eg, resellers/orders | [optional] 
**Event** | Pointer to **string** | The event sent in the request. For eg, im::create. | [optional] 
**EventTimeStamp** | Pointer to **string** | The timestamp at which the event was sent. | [optional] 
**EventId** | Pointer to **string** | A unique id used as identifier for the sepcific event and used for generating the x-hub signature. | [optional] 
**EventType** | Pointer to **string** | The event name sent in the event request. | [optional] 
**Resource** | Pointer to [**OrderCreateV7ResponseResource**](OrderCreateV7ResponseResource.md) |  | [optional] 

## Methods

### NewOrderCreateV7Response

`func NewOrderCreateV7Response() *OrderCreateV7Response`

NewOrderCreateV7Response instantiates a new OrderCreateV7Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderCreateV7ResponseWithDefaults

`func NewOrderCreateV7ResponseWithDefaults() *OrderCreateV7Response`

NewOrderCreateV7ResponseWithDefaults instantiates a new OrderCreateV7Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTopic

`func (o *OrderCreateV7Response) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *OrderCreateV7Response) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *OrderCreateV7Response) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *OrderCreateV7Response) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### GetEvent

`func (o *OrderCreateV7Response) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *OrderCreateV7Response) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *OrderCreateV7Response) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *OrderCreateV7Response) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetEventTimeStamp

`func (o *OrderCreateV7Response) GetEventTimeStamp() string`

GetEventTimeStamp returns the EventTimeStamp field if non-nil, zero value otherwise.

### GetEventTimeStampOk

`func (o *OrderCreateV7Response) GetEventTimeStampOk() (*string, bool)`

GetEventTimeStampOk returns a tuple with the EventTimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTimeStamp

`func (o *OrderCreateV7Response) SetEventTimeStamp(v string)`

SetEventTimeStamp sets EventTimeStamp field to given value.

### HasEventTimeStamp

`func (o *OrderCreateV7Response) HasEventTimeStamp() bool`

HasEventTimeStamp returns a boolean if a field has been set.

### GetEventId

`func (o *OrderCreateV7Response) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *OrderCreateV7Response) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *OrderCreateV7Response) SetEventId(v string)`

SetEventId sets EventId field to given value.

### HasEventId

`func (o *OrderCreateV7Response) HasEventId() bool`

HasEventId returns a boolean if a field has been set.

### GetEventType

`func (o *OrderCreateV7Response) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *OrderCreateV7Response) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *OrderCreateV7Response) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *OrderCreateV7Response) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetResource

`func (o *OrderCreateV7Response) GetResource() OrderCreateV7ResponseResource`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *OrderCreateV7Response) GetResourceOk() (*OrderCreateV7ResponseResource, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *OrderCreateV7Response) SetResource(v OrderCreateV7ResponseResource)`

SetResource sets Resource field to given value.

### HasResource

`func (o *OrderCreateV7Response) HasResource() bool`

HasResource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


