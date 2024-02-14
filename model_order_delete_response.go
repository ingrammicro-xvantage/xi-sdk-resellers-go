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

// checks if the OrderDeleteResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderDeleteResponse{}

// OrderDeleteResponse Response schema for order delete endpoint
type OrderDeleteResponse struct {
	Serviceresponse *OrderCancelResponseServiceresponse `json:"serviceresponse,omitempty"`
}

// NewOrderDeleteResponse instantiates a new OrderDeleteResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderDeleteResponse() *OrderDeleteResponse {
	this := OrderDeleteResponse{}
	return &this
}

// NewOrderDeleteResponseWithDefaults instantiates a new OrderDeleteResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderDeleteResponseWithDefaults() *OrderDeleteResponse {
	this := OrderDeleteResponse{}
	return &this
}

// GetServiceresponse returns the Serviceresponse field value if set, zero value otherwise.
func (o *OrderDeleteResponse) GetServiceresponse() OrderCancelResponseServiceresponse {
	if o == nil || IsNil(o.Serviceresponse) {
		var ret OrderCancelResponseServiceresponse
		return ret
	}
	return *o.Serviceresponse
}

// GetServiceresponseOk returns a tuple with the Serviceresponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDeleteResponse) GetServiceresponseOk() (*OrderCancelResponseServiceresponse, bool) {
	if o == nil || IsNil(o.Serviceresponse) {
		return nil, false
	}
	return o.Serviceresponse, true
}

// HasServiceresponse returns a boolean if a field has been set.
func (o *OrderDeleteResponse) HasServiceresponse() bool {
	if o != nil && !IsNil(o.Serviceresponse) {
		return true
	}

	return false
}

// SetServiceresponse gets a reference to the given OrderCancelResponseServiceresponse and assigns it to the Serviceresponse field.
func (o *OrderDeleteResponse) SetServiceresponse(v OrderCancelResponseServiceresponse) {
	o.Serviceresponse = &v
}

func (o OrderDeleteResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderDeleteResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Serviceresponse) {
		toSerialize["serviceresponse"] = o.Serviceresponse
	}
	return toSerialize, nil
}

type NullableOrderDeleteResponse struct {
	value *OrderDeleteResponse
	isSet bool
}

func (v NullableOrderDeleteResponse) Get() *OrderDeleteResponse {
	return v.value
}

func (v *NullableOrderDeleteResponse) Set(val *OrderDeleteResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderDeleteResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderDeleteResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderDeleteResponse(val *OrderDeleteResponse) *NullableOrderDeleteResponse {
	return &NullableOrderDeleteResponse{value: val, isSet: true}
}

func (v NullableOrderDeleteResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderDeleteResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


