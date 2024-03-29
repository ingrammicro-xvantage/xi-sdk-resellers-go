/*
XI Sdk Resellers

For Resellers. Who are looking to Innovate with Ingram Micro's API SolutionsAutomate your eCommerce with our offering of APIs and Webhooks to create a seamless experience for your customers.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package xi_sdk_resellers

import (
	"encoding/json"
)

// checks if the PostQuoteToOrderV6400Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostQuoteToOrderV6400Response{}

// PostQuoteToOrderV6400Response struct for PostQuoteToOrderV6400Response
type PostQuoteToOrderV6400Response struct {
	// A unique trace id to identify the issue.
	Traceid *string `json:"traceid,omitempty"`
	// Type of the error message.
	Type *string `json:"type,omitempty"`
	// A detailed error message.
	Message *string `json:"message,omitempty"`
	Fields []PostQuoteToOrderV6400ResponseFieldsInner `json:"fields,omitempty"`
}

// NewPostQuoteToOrderV6400Response instantiates a new PostQuoteToOrderV6400Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostQuoteToOrderV6400Response() *PostQuoteToOrderV6400Response {
	this := PostQuoteToOrderV6400Response{}
	return &this
}

// NewPostQuoteToOrderV6400ResponseWithDefaults instantiates a new PostQuoteToOrderV6400Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostQuoteToOrderV6400ResponseWithDefaults() *PostQuoteToOrderV6400Response {
	this := PostQuoteToOrderV6400Response{}
	return &this
}

// GetTraceid returns the Traceid field value if set, zero value otherwise.
func (o *PostQuoteToOrderV6400Response) GetTraceid() string {
	if o == nil || IsNil(o.Traceid) {
		var ret string
		return ret
	}
	return *o.Traceid
}

// GetTraceidOk returns a tuple with the Traceid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostQuoteToOrderV6400Response) GetTraceidOk() (*string, bool) {
	if o == nil || IsNil(o.Traceid) {
		return nil, false
	}
	return o.Traceid, true
}

// HasTraceid returns a boolean if a field has been set.
func (o *PostQuoteToOrderV6400Response) HasTraceid() bool {
	if o != nil && !IsNil(o.Traceid) {
		return true
	}

	return false
}

// SetTraceid gets a reference to the given string and assigns it to the Traceid field.
func (o *PostQuoteToOrderV6400Response) SetTraceid(v string) {
	o.Traceid = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PostQuoteToOrderV6400Response) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostQuoteToOrderV6400Response) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PostQuoteToOrderV6400Response) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PostQuoteToOrderV6400Response) SetType(v string) {
	o.Type = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *PostQuoteToOrderV6400Response) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostQuoteToOrderV6400Response) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *PostQuoteToOrderV6400Response) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *PostQuoteToOrderV6400Response) SetMessage(v string) {
	o.Message = &v
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *PostQuoteToOrderV6400Response) GetFields() []PostQuoteToOrderV6400ResponseFieldsInner {
	if o == nil || IsNil(o.Fields) {
		var ret []PostQuoteToOrderV6400ResponseFieldsInner
		return ret
	}
	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostQuoteToOrderV6400Response) GetFieldsOk() ([]PostQuoteToOrderV6400ResponseFieldsInner, bool) {
	if o == nil || IsNil(o.Fields) {
		return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *PostQuoteToOrderV6400Response) HasFields() bool {
	if o != nil && !IsNil(o.Fields) {
		return true
	}

	return false
}

// SetFields gets a reference to the given []PostQuoteToOrderV6400ResponseFieldsInner and assigns it to the Fields field.
func (o *PostQuoteToOrderV6400Response) SetFields(v []PostQuoteToOrderV6400ResponseFieldsInner) {
	o.Fields = v
}

func (o PostQuoteToOrderV6400Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostQuoteToOrderV6400Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Traceid) {
		toSerialize["traceid"] = o.Traceid
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Fields) {
		toSerialize["fields"] = o.Fields
	}
	return toSerialize, nil
}

type NullablePostQuoteToOrderV6400Response struct {
	value *PostQuoteToOrderV6400Response
	isSet bool
}

func (v NullablePostQuoteToOrderV6400Response) Get() *PostQuoteToOrderV6400Response {
	return v.value
}

func (v *NullablePostQuoteToOrderV6400Response) Set(val *PostQuoteToOrderV6400Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePostQuoteToOrderV6400Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePostQuoteToOrderV6400Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostQuoteToOrderV6400Response(val *PostQuoteToOrderV6400Response) *NullablePostQuoteToOrderV6400Response {
	return &NullablePostQuoteToOrderV6400Response{value: val, isSet: true}
}

func (v NullablePostQuoteToOrderV6400Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostQuoteToOrderV6400Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


