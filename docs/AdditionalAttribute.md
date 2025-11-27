# AdditionalAttribute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttributeName** | Pointer to **string** | The name of the vendor mandatory field. | [optional] 
**AttributeValue** | Pointer to **string** | The value of the vendor mandatory field. | [optional] 
**AttributeDescription** | Pointer to **string** | The description of the vendor mandatory field. | [optional] 
**AttributeHint** | Pointer to **string** | The hint of the vendor mandatory field. | [optional] 
**AttributeRequired** | Pointer to **string** | Indicates if the attribute is mandatory (Added to align with C#). | [optional] 
**Choices** | Pointer to [**[]AdditionalAttribute**](AdditionalAttribute.md) | A list of possible choices for the attribute. | [optional] 

## Methods

### NewAdditionalAttribute

`func NewAdditionalAttribute() *AdditionalAttribute`

NewAdditionalAttribute instantiates a new AdditionalAttribute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdditionalAttributeWithDefaults

`func NewAdditionalAttributeWithDefaults() *AdditionalAttribute`

NewAdditionalAttributeWithDefaults instantiates a new AdditionalAttribute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributeName

`func (o *AdditionalAttribute) GetAttributeName() string`

GetAttributeName returns the AttributeName field if non-nil, zero value otherwise.

### GetAttributeNameOk

`func (o *AdditionalAttribute) GetAttributeNameOk() (*string, bool)`

GetAttributeNameOk returns a tuple with the AttributeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeName

`func (o *AdditionalAttribute) SetAttributeName(v string)`

SetAttributeName sets AttributeName field to given value.

### HasAttributeName

`func (o *AdditionalAttribute) HasAttributeName() bool`

HasAttributeName returns a boolean if a field has been set.

### GetAttributeValue

`func (o *AdditionalAttribute) GetAttributeValue() string`

GetAttributeValue returns the AttributeValue field if non-nil, zero value otherwise.

### GetAttributeValueOk

`func (o *AdditionalAttribute) GetAttributeValueOk() (*string, bool)`

GetAttributeValueOk returns a tuple with the AttributeValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeValue

`func (o *AdditionalAttribute) SetAttributeValue(v string)`

SetAttributeValue sets AttributeValue field to given value.

### HasAttributeValue

`func (o *AdditionalAttribute) HasAttributeValue() bool`

HasAttributeValue returns a boolean if a field has been set.

### GetAttributeDescription

`func (o *AdditionalAttribute) GetAttributeDescription() string`

GetAttributeDescription returns the AttributeDescription field if non-nil, zero value otherwise.

### GetAttributeDescriptionOk

`func (o *AdditionalAttribute) GetAttributeDescriptionOk() (*string, bool)`

GetAttributeDescriptionOk returns a tuple with the AttributeDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeDescription

`func (o *AdditionalAttribute) SetAttributeDescription(v string)`

SetAttributeDescription sets AttributeDescription field to given value.

### HasAttributeDescription

`func (o *AdditionalAttribute) HasAttributeDescription() bool`

HasAttributeDescription returns a boolean if a field has been set.

### GetAttributeHint

`func (o *AdditionalAttribute) GetAttributeHint() string`

GetAttributeHint returns the AttributeHint field if non-nil, zero value otherwise.

### GetAttributeHintOk

`func (o *AdditionalAttribute) GetAttributeHintOk() (*string, bool)`

GetAttributeHintOk returns a tuple with the AttributeHint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeHint

`func (o *AdditionalAttribute) SetAttributeHint(v string)`

SetAttributeHint sets AttributeHint field to given value.

### HasAttributeHint

`func (o *AdditionalAttribute) HasAttributeHint() bool`

HasAttributeHint returns a boolean if a field has been set.

### GetAttributeRequired

`func (o *AdditionalAttribute) GetAttributeRequired() string`

GetAttributeRequired returns the AttributeRequired field if non-nil, zero value otherwise.

### GetAttributeRequiredOk

`func (o *AdditionalAttribute) GetAttributeRequiredOk() (*string, bool)`

GetAttributeRequiredOk returns a tuple with the AttributeRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeRequired

`func (o *AdditionalAttribute) SetAttributeRequired(v string)`

SetAttributeRequired sets AttributeRequired field to given value.

### HasAttributeRequired

`func (o *AdditionalAttribute) HasAttributeRequired() bool`

HasAttributeRequired returns a boolean if a field has been set.

### GetChoices

`func (o *AdditionalAttribute) GetChoices() []AdditionalAttribute`

GetChoices returns the Choices field if non-nil, zero value otherwise.

### GetChoicesOk

`func (o *AdditionalAttribute) GetChoicesOk() (*[]AdditionalAttribute, bool)`

GetChoicesOk returns a tuple with the Choices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChoices

`func (o *AdditionalAttribute) SetChoices(v []AdditionalAttribute)`

SetChoices sets Choices field to given value.

### HasChoices

`func (o *AdditionalAttribute) HasChoices() bool`

HasChoices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


