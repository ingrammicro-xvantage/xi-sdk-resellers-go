/*
Reseller API

For Resellers. <br> Who are looking to Innovate with Ingram Micro's API SolutionsAutomate your eCommerce with our offering of APIs and Webhooks to create a seamless experience for your customers.

API version: 6.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package xi_sdk_resellers

import (
	"encoding/json"
)

// checks if the OrderDetailResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderDetailResponse{}

// OrderDetailResponse Response schema for order details endpoint
type OrderDetailResponse struct {
	Serviceresponse *OrderDetailResponseServiceresponse `json:"serviceresponse,omitempty"`
}

// NewOrderDetailResponse instantiates a new OrderDetailResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderDetailResponse() *OrderDetailResponse {
	this := OrderDetailResponse{}
	return &this
}

// NewOrderDetailResponseWithDefaults instantiates a new OrderDetailResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderDetailResponseWithDefaults() *OrderDetailResponse {
	this := OrderDetailResponse{}
	return &this
}

// GetServiceresponse returns the Serviceresponse field value if set, zero value otherwise.
func (o *OrderDetailResponse) GetServiceresponse() OrderDetailResponseServiceresponse {
	if o == nil || IsNil(o.Serviceresponse) {
		var ret OrderDetailResponseServiceresponse
		return ret
	}
	return *o.Serviceresponse
}

// GetServiceresponseOk returns a tuple with the Serviceresponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailResponse) GetServiceresponseOk() (*OrderDetailResponseServiceresponse, bool) {
	if o == nil || IsNil(o.Serviceresponse) {
		return nil, false
	}
	return o.Serviceresponse, true
}

// HasServiceresponse returns a boolean if a field has been set.
func (o *OrderDetailResponse) HasServiceresponse() bool {
	if o != nil && !IsNil(o.Serviceresponse) {
		return true
	}

	return false
}

// SetServiceresponse gets a reference to the given OrderDetailResponseServiceresponse and assigns it to the Serviceresponse field.
func (o *OrderDetailResponse) SetServiceresponse(v OrderDetailResponseServiceresponse) {
	o.Serviceresponse = &v
}

func (o OrderDetailResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderDetailResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Serviceresponse) {
		toSerialize["serviceresponse"] = o.Serviceresponse
	}
	return toSerialize, nil
}

type NullableOrderDetailResponse struct {
	value *OrderDetailResponse
	isSet bool
}

func (v NullableOrderDetailResponse) Get() *OrderDetailResponse {
	return v.value
}

func (v *NullableOrderDetailResponse) Set(val *OrderDetailResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderDetailResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderDetailResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderDetailResponse(val *OrderDetailResponse) *NullableOrderDetailResponse {
	return &NullableOrderDetailResponse{value: val, isSet: true}
}

func (v NullableOrderDetailResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderDetailResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


