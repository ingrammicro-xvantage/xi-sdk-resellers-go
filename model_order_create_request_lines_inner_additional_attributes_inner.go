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

// checks if the OrderCreateRequestLinesInnerAdditionalAttributesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderCreateRequestLinesInnerAdditionalAttributesInner{}

// OrderCreateRequestLinesInnerAdditionalAttributesInner struct for OrderCreateRequestLinesInnerAdditionalAttributesInner
type OrderCreateRequestLinesInnerAdditionalAttributesInner struct {
	// SAP requested and country-specific line level details.
	AttributeName *string `json:"attributeName,omitempty"`
	// Line-level additional attributes.
	AttributeValue *string `json:"attributeValue,omitempty"`
}

// NewOrderCreateRequestLinesInnerAdditionalAttributesInner instantiates a new OrderCreateRequestLinesInnerAdditionalAttributesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderCreateRequestLinesInnerAdditionalAttributesInner() *OrderCreateRequestLinesInnerAdditionalAttributesInner {
	this := OrderCreateRequestLinesInnerAdditionalAttributesInner{}
	return &this
}

// NewOrderCreateRequestLinesInnerAdditionalAttributesInnerWithDefaults instantiates a new OrderCreateRequestLinesInnerAdditionalAttributesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderCreateRequestLinesInnerAdditionalAttributesInnerWithDefaults() *OrderCreateRequestLinesInnerAdditionalAttributesInner {
	this := OrderCreateRequestLinesInnerAdditionalAttributesInner{}
	return &this
}

// GetAttributeName returns the AttributeName field value if set, zero value otherwise.
func (o *OrderCreateRequestLinesInnerAdditionalAttributesInner) GetAttributeName() string {
	if o == nil || IsNil(o.AttributeName) {
		var ret string
		return ret
	}
	return *o.AttributeName
}

// GetAttributeNameOk returns a tuple with the AttributeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateRequestLinesInnerAdditionalAttributesInner) GetAttributeNameOk() (*string, bool) {
	if o == nil || IsNil(o.AttributeName) {
		return nil, false
	}
	return o.AttributeName, true
}

// HasAttributeName returns a boolean if a field has been set.
func (o *OrderCreateRequestLinesInnerAdditionalAttributesInner) HasAttributeName() bool {
	if o != nil && !IsNil(o.AttributeName) {
		return true
	}

	return false
}

// SetAttributeName gets a reference to the given string and assigns it to the AttributeName field.
func (o *OrderCreateRequestLinesInnerAdditionalAttributesInner) SetAttributeName(v string) {
	o.AttributeName = &v
}

// GetAttributeValue returns the AttributeValue field value if set, zero value otherwise.
func (o *OrderCreateRequestLinesInnerAdditionalAttributesInner) GetAttributeValue() string {
	if o == nil || IsNil(o.AttributeValue) {
		var ret string
		return ret
	}
	return *o.AttributeValue
}

// GetAttributeValueOk returns a tuple with the AttributeValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateRequestLinesInnerAdditionalAttributesInner) GetAttributeValueOk() (*string, bool) {
	if o == nil || IsNil(o.AttributeValue) {
		return nil, false
	}
	return o.AttributeValue, true
}

// HasAttributeValue returns a boolean if a field has been set.
func (o *OrderCreateRequestLinesInnerAdditionalAttributesInner) HasAttributeValue() bool {
	if o != nil && !IsNil(o.AttributeValue) {
		return true
	}

	return false
}

// SetAttributeValue gets a reference to the given string and assigns it to the AttributeValue field.
func (o *OrderCreateRequestLinesInnerAdditionalAttributesInner) SetAttributeValue(v string) {
	o.AttributeValue = &v
}

func (o OrderCreateRequestLinesInnerAdditionalAttributesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderCreateRequestLinesInnerAdditionalAttributesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AttributeName) {
		toSerialize["attributeName"] = o.AttributeName
	}
	if !IsNil(o.AttributeValue) {
		toSerialize["attributeValue"] = o.AttributeValue
	}
	return toSerialize, nil
}

type NullableOrderCreateRequestLinesInnerAdditionalAttributesInner struct {
	value *OrderCreateRequestLinesInnerAdditionalAttributesInner
	isSet bool
}

func (v NullableOrderCreateRequestLinesInnerAdditionalAttributesInner) Get() *OrderCreateRequestLinesInnerAdditionalAttributesInner {
	return v.value
}

func (v *NullableOrderCreateRequestLinesInnerAdditionalAttributesInner) Set(val *OrderCreateRequestLinesInnerAdditionalAttributesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderCreateRequestLinesInnerAdditionalAttributesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderCreateRequestLinesInnerAdditionalAttributesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderCreateRequestLinesInnerAdditionalAttributesInner(val *OrderCreateRequestLinesInnerAdditionalAttributesInner) *NullableOrderCreateRequestLinesInnerAdditionalAttributesInner {
	return &NullableOrderCreateRequestLinesInnerAdditionalAttributesInner{value: val, isSet: true}
}

func (v NullableOrderCreateRequestLinesInnerAdditionalAttributesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderCreateRequestLinesInnerAdditionalAttributesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


