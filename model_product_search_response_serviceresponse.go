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

// checks if the ProductSearchResponseServiceresponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductSearchResponseServiceresponse{}

// ProductSearchResponseServiceresponse struct for ProductSearchResponseServiceresponse
type ProductSearchResponseServiceresponse struct {
	Responsepreamble *ProductSearchResponseServiceresponseResponsepreamble `json:"responsepreamble,omitempty"`
	Productsearchresponse []ProductSearchResponseServiceresponseProductsearchresponseInner `json:"productsearchresponse,omitempty"`
}

// NewProductSearchResponseServiceresponse instantiates a new ProductSearchResponseServiceresponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductSearchResponseServiceresponse() *ProductSearchResponseServiceresponse {
	this := ProductSearchResponseServiceresponse{}
	return &this
}

// NewProductSearchResponseServiceresponseWithDefaults instantiates a new ProductSearchResponseServiceresponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductSearchResponseServiceresponseWithDefaults() *ProductSearchResponseServiceresponse {
	this := ProductSearchResponseServiceresponse{}
	return &this
}

// GetResponsepreamble returns the Responsepreamble field value if set, zero value otherwise.
func (o *ProductSearchResponseServiceresponse) GetResponsepreamble() ProductSearchResponseServiceresponseResponsepreamble {
	if o == nil || IsNil(o.Responsepreamble) {
		var ret ProductSearchResponseServiceresponseResponsepreamble
		return ret
	}
	return *o.Responsepreamble
}

// GetResponsepreambleOk returns a tuple with the Responsepreamble field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductSearchResponseServiceresponse) GetResponsepreambleOk() (*ProductSearchResponseServiceresponseResponsepreamble, bool) {
	if o == nil || IsNil(o.Responsepreamble) {
		return nil, false
	}
	return o.Responsepreamble, true
}

// HasResponsepreamble returns a boolean if a field has been set.
func (o *ProductSearchResponseServiceresponse) HasResponsepreamble() bool {
	if o != nil && !IsNil(o.Responsepreamble) {
		return true
	}

	return false
}

// SetResponsepreamble gets a reference to the given ProductSearchResponseServiceresponseResponsepreamble and assigns it to the Responsepreamble field.
func (o *ProductSearchResponseServiceresponse) SetResponsepreamble(v ProductSearchResponseServiceresponseResponsepreamble) {
	o.Responsepreamble = &v
}

// GetProductsearchresponse returns the Productsearchresponse field value if set, zero value otherwise.
func (o *ProductSearchResponseServiceresponse) GetProductsearchresponse() []ProductSearchResponseServiceresponseProductsearchresponseInner {
	if o == nil || IsNil(o.Productsearchresponse) {
		var ret []ProductSearchResponseServiceresponseProductsearchresponseInner
		return ret
	}
	return o.Productsearchresponse
}

// GetProductsearchresponseOk returns a tuple with the Productsearchresponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductSearchResponseServiceresponse) GetProductsearchresponseOk() ([]ProductSearchResponseServiceresponseProductsearchresponseInner, bool) {
	if o == nil || IsNil(o.Productsearchresponse) {
		return nil, false
	}
	return o.Productsearchresponse, true
}

// HasProductsearchresponse returns a boolean if a field has been set.
func (o *ProductSearchResponseServiceresponse) HasProductsearchresponse() bool {
	if o != nil && !IsNil(o.Productsearchresponse) {
		return true
	}

	return false
}

// SetProductsearchresponse gets a reference to the given []ProductSearchResponseServiceresponseProductsearchresponseInner and assigns it to the Productsearchresponse field.
func (o *ProductSearchResponseServiceresponse) SetProductsearchresponse(v []ProductSearchResponseServiceresponseProductsearchresponseInner) {
	o.Productsearchresponse = v
}

func (o ProductSearchResponseServiceresponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductSearchResponseServiceresponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Responsepreamble) {
		toSerialize["responsepreamble"] = o.Responsepreamble
	}
	if !IsNil(o.Productsearchresponse) {
		toSerialize["productsearchresponse"] = o.Productsearchresponse
	}
	return toSerialize, nil
}

type NullableProductSearchResponseServiceresponse struct {
	value *ProductSearchResponseServiceresponse
	isSet bool
}

func (v NullableProductSearchResponseServiceresponse) Get() *ProductSearchResponseServiceresponse {
	return v.value
}

func (v *NullableProductSearchResponseServiceresponse) Set(val *ProductSearchResponseServiceresponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProductSearchResponseServiceresponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProductSearchResponseServiceresponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductSearchResponseServiceresponse(val *ProductSearchResponseServiceresponse) *NullableProductSearchResponseServiceresponse {
	return &NullableProductSearchResponseServiceresponse{value: val, isSet: true}
}

func (v NullableProductSearchResponseServiceresponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductSearchResponseServiceresponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


