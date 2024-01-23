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

// checks if the OrderDetailB2BAdditionalAttributesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderDetailB2BAdditionalAttributesInner{}

// OrderDetailB2BAdditionalAttributesInner struct for OrderDetailB2BAdditionalAttributesInner
type OrderDetailB2BAdditionalAttributesInner struct {
	// Header level custom field names.
	AttributeName *string `json:"attributeName,omitempty"`
	// Value of the custom fields.
	AttributeValue *string `json:"attributeValue,omitempty"`
}

// NewOrderDetailB2BAdditionalAttributesInner instantiates a new OrderDetailB2BAdditionalAttributesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderDetailB2BAdditionalAttributesInner() *OrderDetailB2BAdditionalAttributesInner {
	this := OrderDetailB2BAdditionalAttributesInner{}
	return &this
}

// NewOrderDetailB2BAdditionalAttributesInnerWithDefaults instantiates a new OrderDetailB2BAdditionalAttributesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderDetailB2BAdditionalAttributesInnerWithDefaults() *OrderDetailB2BAdditionalAttributesInner {
	this := OrderDetailB2BAdditionalAttributesInner{}
	return &this
}

// GetAttributeName returns the AttributeName field value if set, zero value otherwise.
func (o *OrderDetailB2BAdditionalAttributesInner) GetAttributeName() string {
	if o == nil || IsNil(o.AttributeName) {
		var ret string
		return ret
	}
	return *o.AttributeName
}

// GetAttributeNameOk returns a tuple with the AttributeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailB2BAdditionalAttributesInner) GetAttributeNameOk() (*string, bool) {
	if o == nil || IsNil(o.AttributeName) {
		return nil, false
	}
	return o.AttributeName, true
}

// HasAttributeName returns a boolean if a field has been set.
func (o *OrderDetailB2BAdditionalAttributesInner) HasAttributeName() bool {
	if o != nil && !IsNil(o.AttributeName) {
		return true
	}

	return false
}

// SetAttributeName gets a reference to the given string and assigns it to the AttributeName field.
func (o *OrderDetailB2BAdditionalAttributesInner) SetAttributeName(v string) {
	o.AttributeName = &v
}

// GetAttributeValue returns the AttributeValue field value if set, zero value otherwise.
func (o *OrderDetailB2BAdditionalAttributesInner) GetAttributeValue() string {
	if o == nil || IsNil(o.AttributeValue) {
		var ret string
		return ret
	}
	return *o.AttributeValue
}

// GetAttributeValueOk returns a tuple with the AttributeValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailB2BAdditionalAttributesInner) GetAttributeValueOk() (*string, bool) {
	if o == nil || IsNil(o.AttributeValue) {
		return nil, false
	}
	return o.AttributeValue, true
}

// HasAttributeValue returns a boolean if a field has been set.
func (o *OrderDetailB2BAdditionalAttributesInner) HasAttributeValue() bool {
	if o != nil && !IsNil(o.AttributeValue) {
		return true
	}

	return false
}

// SetAttributeValue gets a reference to the given string and assigns it to the AttributeValue field.
func (o *OrderDetailB2BAdditionalAttributesInner) SetAttributeValue(v string) {
	o.AttributeValue = &v
}

func (o OrderDetailB2BAdditionalAttributesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderDetailB2BAdditionalAttributesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AttributeName) {
		toSerialize["attributeName"] = o.AttributeName
	}
	if !IsNil(o.AttributeValue) {
		toSerialize["attributeValue"] = o.AttributeValue
	}
	return toSerialize, nil
}

type NullableOrderDetailB2BAdditionalAttributesInner struct {
	value *OrderDetailB2BAdditionalAttributesInner
	isSet bool
}

func (v NullableOrderDetailB2BAdditionalAttributesInner) Get() *OrderDetailB2BAdditionalAttributesInner {
	return v.value
}

func (v *NullableOrderDetailB2BAdditionalAttributesInner) Set(val *OrderDetailB2BAdditionalAttributesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderDetailB2BAdditionalAttributesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderDetailB2BAdditionalAttributesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderDetailB2BAdditionalAttributesInner(val *OrderDetailB2BAdditionalAttributesInner) *NullableOrderDetailB2BAdditionalAttributesInner {
	return &NullableOrderDetailB2BAdditionalAttributesInner{value: val, isSet: true}
}

func (v NullableOrderDetailB2BAdditionalAttributesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderDetailB2BAdditionalAttributesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


