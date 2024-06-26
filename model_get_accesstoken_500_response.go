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

// checks if the GetAccesstoken500Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAccesstoken500Response{}

// GetAccesstoken500Response struct for GetAccesstoken500Response
type GetAccesstoken500Response struct {
	Fault *GetAccesstoken500ResponseFault `json:"fault,omitempty"`
}

// NewGetAccesstoken500Response instantiates a new GetAccesstoken500Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAccesstoken500Response() *GetAccesstoken500Response {
	this := GetAccesstoken500Response{}
	return &this
}

// NewGetAccesstoken500ResponseWithDefaults instantiates a new GetAccesstoken500Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAccesstoken500ResponseWithDefaults() *GetAccesstoken500Response {
	this := GetAccesstoken500Response{}
	return &this
}

// GetFault returns the Fault field value if set, zero value otherwise.
func (o *GetAccesstoken500Response) GetFault() GetAccesstoken500ResponseFault {
	if o == nil || IsNil(o.Fault) {
		var ret GetAccesstoken500ResponseFault
		return ret
	}
	return *o.Fault
}

// GetFaultOk returns a tuple with the Fault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccesstoken500Response) GetFaultOk() (*GetAccesstoken500ResponseFault, bool) {
	if o == nil || IsNil(o.Fault) {
		return nil, false
	}
	return o.Fault, true
}

// HasFault returns a boolean if a field has been set.
func (o *GetAccesstoken500Response) HasFault() bool {
	if o != nil && !IsNil(o.Fault) {
		return true
	}

	return false
}

// SetFault gets a reference to the given GetAccesstoken500ResponseFault and assigns it to the Fault field.
func (o *GetAccesstoken500Response) SetFault(v GetAccesstoken500ResponseFault) {
	o.Fault = &v
}

func (o GetAccesstoken500Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAccesstoken500Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Fault) {
		toSerialize["fault"] = o.Fault
	}
	return toSerialize, nil
}

type NullableGetAccesstoken500Response struct {
	value *GetAccesstoken500Response
	isSet bool
}

func (v NullableGetAccesstoken500Response) Get() *GetAccesstoken500Response {
	return v.value
}

func (v *NullableGetAccesstoken500Response) Set(val *GetAccesstoken500Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAccesstoken500Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAccesstoken500Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAccesstoken500Response(val *GetAccesstoken500Response) *NullableGetAccesstoken500Response {
	return &NullableGetAccesstoken500Response{value: val, isSet: true}
}

func (v NullableGetAccesstoken500Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAccesstoken500Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


