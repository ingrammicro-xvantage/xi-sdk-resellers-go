/*
XI Sdk Resellers

For resellers seeking to innovate with Ingram Micro's API solutions, automate your eCommerce experience with our array of API's and webhooks to craft a seamless journey for your customers.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package xi_sdk_resellers

import (
	"encoding/json"
)

// checks if the GetResellerV6ValidateQuote400ResponseFieldsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetResellerV6ValidateQuote400ResponseFieldsInner{}

// GetResellerV6ValidateQuote400ResponseFieldsInner struct for GetResellerV6ValidateQuote400ResponseFieldsInner
type GetResellerV6ValidateQuote400ResponseFieldsInner struct {
	// Contains the name of the field.
	Field *string `json:"field,omitempty"`
	// Gives the description of the field message.
	Message *string `json:"message,omitempty"`
	// Value sent in the input for the specific field.
	Value *string `json:"value,omitempty"`
}

// NewGetResellerV6ValidateQuote400ResponseFieldsInner instantiates a new GetResellerV6ValidateQuote400ResponseFieldsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetResellerV6ValidateQuote400ResponseFieldsInner() *GetResellerV6ValidateQuote400ResponseFieldsInner {
	this := GetResellerV6ValidateQuote400ResponseFieldsInner{}
	return &this
}

// NewGetResellerV6ValidateQuote400ResponseFieldsInnerWithDefaults instantiates a new GetResellerV6ValidateQuote400ResponseFieldsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetResellerV6ValidateQuote400ResponseFieldsInnerWithDefaults() *GetResellerV6ValidateQuote400ResponseFieldsInner {
	this := GetResellerV6ValidateQuote400ResponseFieldsInner{}
	return &this
}

// GetField returns the Field field value if set, zero value otherwise.
func (o *GetResellerV6ValidateQuote400ResponseFieldsInner) GetField() string {
	if o == nil || IsNil(o.Field) {
		var ret string
		return ret
	}
	return *o.Field
}

// GetFieldOk returns a tuple with the Field field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetResellerV6ValidateQuote400ResponseFieldsInner) GetFieldOk() (*string, bool) {
	if o == nil || IsNil(o.Field) {
		return nil, false
	}
	return o.Field, true
}

// HasField returns a boolean if a field has been set.
func (o *GetResellerV6ValidateQuote400ResponseFieldsInner) HasField() bool {
	if o != nil && !IsNil(o.Field) {
		return true
	}

	return false
}

// SetField gets a reference to the given string and assigns it to the Field field.
func (o *GetResellerV6ValidateQuote400ResponseFieldsInner) SetField(v string) {
	o.Field = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *GetResellerV6ValidateQuote400ResponseFieldsInner) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetResellerV6ValidateQuote400ResponseFieldsInner) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *GetResellerV6ValidateQuote400ResponseFieldsInner) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *GetResellerV6ValidateQuote400ResponseFieldsInner) SetMessage(v string) {
	o.Message = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *GetResellerV6ValidateQuote400ResponseFieldsInner) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetResellerV6ValidateQuote400ResponseFieldsInner) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *GetResellerV6ValidateQuote400ResponseFieldsInner) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *GetResellerV6ValidateQuote400ResponseFieldsInner) SetValue(v string) {
	o.Value = &v
}

func (o GetResellerV6ValidateQuote400ResponseFieldsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetResellerV6ValidateQuote400ResponseFieldsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Field) {
		toSerialize["field"] = o.Field
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableGetResellerV6ValidateQuote400ResponseFieldsInner struct {
	value *GetResellerV6ValidateQuote400ResponseFieldsInner
	isSet bool
}

func (v NullableGetResellerV6ValidateQuote400ResponseFieldsInner) Get() *GetResellerV6ValidateQuote400ResponseFieldsInner {
	return v.value
}

func (v *NullableGetResellerV6ValidateQuote400ResponseFieldsInner) Set(val *GetResellerV6ValidateQuote400ResponseFieldsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetResellerV6ValidateQuote400ResponseFieldsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetResellerV6ValidateQuote400ResponseFieldsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetResellerV6ValidateQuote400ResponseFieldsInner(val *GetResellerV6ValidateQuote400ResponseFieldsInner) *NullableGetResellerV6ValidateQuote400ResponseFieldsInner {
	return &NullableGetResellerV6ValidateQuote400ResponseFieldsInner{value: val, isSet: true}
}

func (v NullableGetResellerV6ValidateQuote400ResponseFieldsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetResellerV6ValidateQuote400ResponseFieldsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


