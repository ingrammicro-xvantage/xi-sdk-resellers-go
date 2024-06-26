# PostCreateorderV7500Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Traceid** | Pointer to **string** | Unique Id to identify error. | [optional] 
**Type** | Pointer to **string** | Describes the type of the error. | [optional] 
**Message** | Pointer to **string** | Describes the error message. | [optional] 
**Fields** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewPostCreateorderV7500Response

`func NewPostCreateorderV7500Response() *PostCreateorderV7500Response`

NewPostCreateorderV7500Response instantiates a new PostCreateorderV7500Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostCreateorderV7500ResponseWithDefaults

`func NewPostCreateorderV7500ResponseWithDefaults() *PostCreateorderV7500Response`

NewPostCreateorderV7500ResponseWithDefaults instantiates a new PostCreateorderV7500Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTraceid

`func (o *PostCreateorderV7500Response) GetTraceid() string`

GetTraceid returns the Traceid field if non-nil, zero value otherwise.

### GetTraceidOk

`func (o *PostCreateorderV7500Response) GetTraceidOk() (*string, bool)`

GetTraceidOk returns a tuple with the Traceid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceid

`func (o *PostCreateorderV7500Response) SetTraceid(v string)`

SetTraceid sets Traceid field to given value.

### HasTraceid

`func (o *PostCreateorderV7500Response) HasTraceid() bool`

HasTraceid returns a boolean if a field has been set.

### GetType

`func (o *PostCreateorderV7500Response) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PostCreateorderV7500Response) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PostCreateorderV7500Response) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PostCreateorderV7500Response) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMessage

`func (o *PostCreateorderV7500Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PostCreateorderV7500Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PostCreateorderV7500Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *PostCreateorderV7500Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetFields

`func (o *PostCreateorderV7500Response) GetFields() []map[string]interface{}`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *PostCreateorderV7500Response) GetFieldsOk() (*[]map[string]interface{}, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *PostCreateorderV7500Response) SetFields(v []map[string]interface{})`

SetFields sets Fields field to given value.

### HasFields

`func (o *PostCreateorderV7500Response) HasFields() bool`

HasFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


