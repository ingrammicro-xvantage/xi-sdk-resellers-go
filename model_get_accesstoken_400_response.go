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

// checks if the GetAccesstoken400Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAccesstoken400Response{}

// GetAccesstoken400Response struct for GetAccesstoken400Response
type GetAccesstoken400Response struct {
	Message *string `json:"message,omitempty"`
}

// NewGetAccesstoken400Response instantiates a new GetAccesstoken400Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAccesstoken400Response() *GetAccesstoken400Response {
	this := GetAccesstoken400Response{}
	return &this
}

// NewGetAccesstoken400ResponseWithDefaults instantiates a new GetAccesstoken400Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAccesstoken400ResponseWithDefaults() *GetAccesstoken400Response {
	this := GetAccesstoken400Response{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *GetAccesstoken400Response) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccesstoken400Response) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *GetAccesstoken400Response) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *GetAccesstoken400Response) SetMessage(v string) {
	o.Message = &v
}

func (o GetAccesstoken400Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAccesstoken400Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableGetAccesstoken400Response struct {
	value *GetAccesstoken400Response
	isSet bool
}

func (v NullableGetAccesstoken400Response) Get() *GetAccesstoken400Response {
	return v.value
}

func (v *NullableGetAccesstoken400Response) Set(val *GetAccesstoken400Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAccesstoken400Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAccesstoken400Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAccesstoken400Response(val *GetAccesstoken400Response) *NullableGetAccesstoken400Response {
	return &NullableGetAccesstoken400Response{value: val, isSet: true}
}

func (v NullableGetAccesstoken400Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAccesstoken400Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


