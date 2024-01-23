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

// checks if the ProductDetailResponseTechnicalSpecificationsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductDetailResponseTechnicalSpecificationsInner{}

// ProductDetailResponseTechnicalSpecificationsInner struct for ProductDetailResponseTechnicalSpecificationsInner
type ProductDetailResponseTechnicalSpecificationsInner struct {
	// Example : 'Basic'
	Headername *string `json:"headername,omitempty"`
	// Example : 'LCD Monitor'
	Attributevalue *string `json:"attributevalue,omitempty"`
	// Example : 'Basic|Product Type|LCD Monitor'
	Attributedisplay *string `json:"attributedisplay,omitempty"`
	// Example : 'Product Type'
	Attributename *string `json:"attributename,omitempty"`
}

// NewProductDetailResponseTechnicalSpecificationsInner instantiates a new ProductDetailResponseTechnicalSpecificationsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductDetailResponseTechnicalSpecificationsInner() *ProductDetailResponseTechnicalSpecificationsInner {
	this := ProductDetailResponseTechnicalSpecificationsInner{}
	return &this
}

// NewProductDetailResponseTechnicalSpecificationsInnerWithDefaults instantiates a new ProductDetailResponseTechnicalSpecificationsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductDetailResponseTechnicalSpecificationsInnerWithDefaults() *ProductDetailResponseTechnicalSpecificationsInner {
	this := ProductDetailResponseTechnicalSpecificationsInner{}
	return &this
}

// GetHeadername returns the Headername field value if set, zero value otherwise.
func (o *ProductDetailResponseTechnicalSpecificationsInner) GetHeadername() string {
	if o == nil || IsNil(o.Headername) {
		var ret string
		return ret
	}
	return *o.Headername
}

// GetHeadernameOk returns a tuple with the Headername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductDetailResponseTechnicalSpecificationsInner) GetHeadernameOk() (*string, bool) {
	if o == nil || IsNil(o.Headername) {
		return nil, false
	}
	return o.Headername, true
}

// HasHeadername returns a boolean if a field has been set.
func (o *ProductDetailResponseTechnicalSpecificationsInner) HasHeadername() bool {
	if o != nil && !IsNil(o.Headername) {
		return true
	}

	return false
}

// SetHeadername gets a reference to the given string and assigns it to the Headername field.
func (o *ProductDetailResponseTechnicalSpecificationsInner) SetHeadername(v string) {
	o.Headername = &v
}

// GetAttributevalue returns the Attributevalue field value if set, zero value otherwise.
func (o *ProductDetailResponseTechnicalSpecificationsInner) GetAttributevalue() string {
	if o == nil || IsNil(o.Attributevalue) {
		var ret string
		return ret
	}
	return *o.Attributevalue
}

// GetAttributevalueOk returns a tuple with the Attributevalue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductDetailResponseTechnicalSpecificationsInner) GetAttributevalueOk() (*string, bool) {
	if o == nil || IsNil(o.Attributevalue) {
		return nil, false
	}
	return o.Attributevalue, true
}

// HasAttributevalue returns a boolean if a field has been set.
func (o *ProductDetailResponseTechnicalSpecificationsInner) HasAttributevalue() bool {
	if o != nil && !IsNil(o.Attributevalue) {
		return true
	}

	return false
}

// SetAttributevalue gets a reference to the given string and assigns it to the Attributevalue field.
func (o *ProductDetailResponseTechnicalSpecificationsInner) SetAttributevalue(v string) {
	o.Attributevalue = &v
}

// GetAttributedisplay returns the Attributedisplay field value if set, zero value otherwise.
func (o *ProductDetailResponseTechnicalSpecificationsInner) GetAttributedisplay() string {
	if o == nil || IsNil(o.Attributedisplay) {
		var ret string
		return ret
	}
	return *o.Attributedisplay
}

// GetAttributedisplayOk returns a tuple with the Attributedisplay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductDetailResponseTechnicalSpecificationsInner) GetAttributedisplayOk() (*string, bool) {
	if o == nil || IsNil(o.Attributedisplay) {
		return nil, false
	}
	return o.Attributedisplay, true
}

// HasAttributedisplay returns a boolean if a field has been set.
func (o *ProductDetailResponseTechnicalSpecificationsInner) HasAttributedisplay() bool {
	if o != nil && !IsNil(o.Attributedisplay) {
		return true
	}

	return false
}

// SetAttributedisplay gets a reference to the given string and assigns it to the Attributedisplay field.
func (o *ProductDetailResponseTechnicalSpecificationsInner) SetAttributedisplay(v string) {
	o.Attributedisplay = &v
}

// GetAttributename returns the Attributename field value if set, zero value otherwise.
func (o *ProductDetailResponseTechnicalSpecificationsInner) GetAttributename() string {
	if o == nil || IsNil(o.Attributename) {
		var ret string
		return ret
	}
	return *o.Attributename
}

// GetAttributenameOk returns a tuple with the Attributename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductDetailResponseTechnicalSpecificationsInner) GetAttributenameOk() (*string, bool) {
	if o == nil || IsNil(o.Attributename) {
		return nil, false
	}
	return o.Attributename, true
}

// HasAttributename returns a boolean if a field has been set.
func (o *ProductDetailResponseTechnicalSpecificationsInner) HasAttributename() bool {
	if o != nil && !IsNil(o.Attributename) {
		return true
	}

	return false
}

// SetAttributename gets a reference to the given string and assigns it to the Attributename field.
func (o *ProductDetailResponseTechnicalSpecificationsInner) SetAttributename(v string) {
	o.Attributename = &v
}

func (o ProductDetailResponseTechnicalSpecificationsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductDetailResponseTechnicalSpecificationsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Headername) {
		toSerialize["headername"] = o.Headername
	}
	if !IsNil(o.Attributevalue) {
		toSerialize["attributevalue"] = o.Attributevalue
	}
	if !IsNil(o.Attributedisplay) {
		toSerialize["attributedisplay"] = o.Attributedisplay
	}
	if !IsNil(o.Attributename) {
		toSerialize["attributename"] = o.Attributename
	}
	return toSerialize, nil
}

type NullableProductDetailResponseTechnicalSpecificationsInner struct {
	value *ProductDetailResponseTechnicalSpecificationsInner
	isSet bool
}

func (v NullableProductDetailResponseTechnicalSpecificationsInner) Get() *ProductDetailResponseTechnicalSpecificationsInner {
	return v.value
}

func (v *NullableProductDetailResponseTechnicalSpecificationsInner) Set(val *ProductDetailResponseTechnicalSpecificationsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableProductDetailResponseTechnicalSpecificationsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableProductDetailResponseTechnicalSpecificationsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductDetailResponseTechnicalSpecificationsInner(val *ProductDetailResponseTechnicalSpecificationsInner) *NullableProductDetailResponseTechnicalSpecificationsInner {
	return &NullableProductDetailResponseTechnicalSpecificationsInner{value: val, isSet: true}
}

func (v NullableProductDetailResponseTechnicalSpecificationsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductDetailResponseTechnicalSpecificationsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


