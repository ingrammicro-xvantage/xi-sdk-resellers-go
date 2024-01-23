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

// checks if the ProductLineTypeTrackingnumberdetailsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductLineTypeTrackingnumberdetailsInner{}

// ProductLineTypeTrackingnumberdetailsInner struct for ProductLineTypeTrackingnumberdetailsInner
type ProductLineTypeTrackingnumberdetailsInner struct {
	Trackingnumber *string `json:"trackingnumber,omitempty"`
}

// NewProductLineTypeTrackingnumberdetailsInner instantiates a new ProductLineTypeTrackingnumberdetailsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductLineTypeTrackingnumberdetailsInner() *ProductLineTypeTrackingnumberdetailsInner {
	this := ProductLineTypeTrackingnumberdetailsInner{}
	return &this
}

// NewProductLineTypeTrackingnumberdetailsInnerWithDefaults instantiates a new ProductLineTypeTrackingnumberdetailsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductLineTypeTrackingnumberdetailsInnerWithDefaults() *ProductLineTypeTrackingnumberdetailsInner {
	this := ProductLineTypeTrackingnumberdetailsInner{}
	return &this
}

// GetTrackingnumber returns the Trackingnumber field value if set, zero value otherwise.
func (o *ProductLineTypeTrackingnumberdetailsInner) GetTrackingnumber() string {
	if o == nil || IsNil(o.Trackingnumber) {
		var ret string
		return ret
	}
	return *o.Trackingnumber
}

// GetTrackingnumberOk returns a tuple with the Trackingnumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductLineTypeTrackingnumberdetailsInner) GetTrackingnumberOk() (*string, bool) {
	if o == nil || IsNil(o.Trackingnumber) {
		return nil, false
	}
	return o.Trackingnumber, true
}

// HasTrackingnumber returns a boolean if a field has been set.
func (o *ProductLineTypeTrackingnumberdetailsInner) HasTrackingnumber() bool {
	if o != nil && !IsNil(o.Trackingnumber) {
		return true
	}

	return false
}

// SetTrackingnumber gets a reference to the given string and assigns it to the Trackingnumber field.
func (o *ProductLineTypeTrackingnumberdetailsInner) SetTrackingnumber(v string) {
	o.Trackingnumber = &v
}

func (o ProductLineTypeTrackingnumberdetailsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductLineTypeTrackingnumberdetailsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Trackingnumber) {
		toSerialize["trackingnumber"] = o.Trackingnumber
	}
	return toSerialize, nil
}

type NullableProductLineTypeTrackingnumberdetailsInner struct {
	value *ProductLineTypeTrackingnumberdetailsInner
	isSet bool
}

func (v NullableProductLineTypeTrackingnumberdetailsInner) Get() *ProductLineTypeTrackingnumberdetailsInner {
	return v.value
}

func (v *NullableProductLineTypeTrackingnumberdetailsInner) Set(val *ProductLineTypeTrackingnumberdetailsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableProductLineTypeTrackingnumberdetailsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableProductLineTypeTrackingnumberdetailsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductLineTypeTrackingnumberdetailsInner(val *ProductLineTypeTrackingnumberdetailsInner) *NullableProductLineTypeTrackingnumberdetailsInner {
	return &NullableProductLineTypeTrackingnumberdetailsInner{value: val, isSet: true}
}

func (v NullableProductLineTypeTrackingnumberdetailsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductLineTypeTrackingnumberdetailsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


