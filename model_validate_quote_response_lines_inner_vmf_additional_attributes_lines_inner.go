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

// checks if the ValidateQuoteResponseLinesInnerVmfAdditionalAttributesLinesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ValidateQuoteResponseLinesInnerVmfAdditionalAttributesLinesInner{}

// ValidateQuoteResponseLinesInnerVmfAdditionalAttributesLinesInner struct for ValidateQuoteResponseLinesInnerVmfAdditionalAttributesLinesInner
type ValidateQuoteResponseLinesInnerVmfAdditionalAttributesLinesInner struct {
	// The name of the line level field.
	AttributeName *string `json:"attributeName,omitempty"`
	// The value of the line level field.
	AttributeValue *string `json:"attributeValue,omitempty"`
	// The description of the line level field.
	AttributeDescription *string `json:"attributeDescription,omitempty"`
}

// NewValidateQuoteResponseLinesInnerVmfAdditionalAttributesLinesInner instantiates a new ValidateQuoteResponseLinesInnerVmfAdditionalAttributesLinesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidateQuoteResponseLinesInnerVmfAdditionalAttributesLinesInner() *ValidateQuoteResponseLinesInnerVmfAdditionalAttributesLinesInner {
	this := ValidateQuoteResponseLinesInnerVmfAdditionalAttributesLinesInner{}
	return &this
}

// NewValidateQuoteResponseLinesInnerVmfAdditionalAttributesLinesInnerWithDefaults instantiates a new ValidateQuoteResponseLinesInnerVmfAdditionalAttributesLinesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidateQuoteResponseLinesInnerVmfAdditionalAttributesLinesInnerWithDefaults() *ValidateQuoteResponseLinesInnerVmfAdditionalAttributesLinesInner {
	this := ValidateQuoteResponseLinesInnerVmfAdditionalAttributesLinesInner{}
	return &this
}

// GetAttributeName returns the AttributeName field value if set, zero value otherwise.
func (o *ValidateQuoteResponseLinesInnerVmfAdditionalAttributesLinesInner) GetAttributeName() string {
	if o == nil || IsNil(o.AttributeName) {
		var ret string
		return ret
	}
	return *o.AttributeName
}

// GetAttributeNameOk returns a tuple with the AttributeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateQuoteResponseLinesInnerVmfAdditionalAttributesLinesInner) GetAttributeNameOk() (*string, bool) {
	if o == nil || IsNil(o.AttributeName) {
		return nil, false
	}
	return o.AttributeName, true
}

// HasAttributeName returns a boolean if a field has been set.
func (o *ValidateQuoteResponseLinesInnerVmfAdditionalAttributesLinesInner) HasAttributeName() bool {
	if o != nil && !IsNil(o.AttributeName) {
		return true
	}

	return false
}

// SetAttributeName gets a reference to the given string and assigns it to the AttributeName field.
func (o *ValidateQuoteResponseLinesInnerVmfAdditionalAttributesLinesInner) SetAttributeName(v string) {
	o.AttributeName = &v
}

// GetAttributeValue returns the AttributeValue field value if set, zero value otherwise.
func (o *ValidateQuoteResponseLinesInnerVmfAdditionalAttributesLinesInner) GetAttributeValue() string {
	if o == nil || IsNil(o.AttributeValue) {
		var ret string
		return ret
	}
	return *o.AttributeValue
}

// GetAttributeValueOk returns a tuple with the AttributeValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateQuoteResponseLinesInnerVmfAdditionalAttributesLinesInner) GetAttributeValueOk() (*string, bool) {
	if o == nil || IsNil(o.AttributeValue) {
		return nil, false
	}
	return o.AttributeValue, true
}

// HasAttributeValue returns a boolean if a field has been set.
func (o *ValidateQuoteResponseLinesInnerVmfAdditionalAttributesLinesInner) HasAttributeValue() bool {
	if o != nil && !IsNil(o.AttributeValue) {
		return true
	}

	return false
}

// SetAttributeValue gets a reference to the given string and assigns it to the AttributeValue field.
func (o *ValidateQuoteResponseLinesInnerVmfAdditionalAttributesLinesInner) SetAttributeValue(v string) {
	o.AttributeValue = &v
}

// GetAttributeDescription returns the AttributeDescription field value if set, zero value otherwise.
func (o *ValidateQuoteResponseLinesInnerVmfAdditionalAttributesLinesInner) GetAttributeDescription() string {
	if o == nil || IsNil(o.AttributeDescription) {
		var ret string
		return ret
	}
	return *o.AttributeDescription
}

// GetAttributeDescriptionOk returns a tuple with the AttributeDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateQuoteResponseLinesInnerVmfAdditionalAttributesLinesInner) GetAttributeDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.AttributeDescription) {
		return nil, false
	}
	return o.AttributeDescription, true
}

// HasAttributeDescription returns a boolean if a field has been set.
func (o *ValidateQuoteResponseLinesInnerVmfAdditionalAttributesLinesInner) HasAttributeDescription() bool {
	if o != nil && !IsNil(o.AttributeDescription) {
		return true
	}

	return false
}

// SetAttributeDescription gets a reference to the given string and assigns it to the AttributeDescription field.
func (o *ValidateQuoteResponseLinesInnerVmfAdditionalAttributesLinesInner) SetAttributeDescription(v string) {
	o.AttributeDescription = &v
}

func (o ValidateQuoteResponseLinesInnerVmfAdditionalAttributesLinesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ValidateQuoteResponseLinesInnerVmfAdditionalAttributesLinesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AttributeName) {
		toSerialize["attributeName"] = o.AttributeName
	}
	if !IsNil(o.AttributeValue) {
		toSerialize["attributeValue"] = o.AttributeValue
	}
	if !IsNil(o.AttributeDescription) {
		toSerialize["attributeDescription"] = o.AttributeDescription
	}
	return toSerialize, nil
}

type NullableValidateQuoteResponseLinesInnerVmfAdditionalAttributesLinesInner struct {
	value *ValidateQuoteResponseLinesInnerVmfAdditionalAttributesLinesInner
	isSet bool
}

func (v NullableValidateQuoteResponseLinesInnerVmfAdditionalAttributesLinesInner) Get() *ValidateQuoteResponseLinesInnerVmfAdditionalAttributesLinesInner {
	return v.value
}

func (v *NullableValidateQuoteResponseLinesInnerVmfAdditionalAttributesLinesInner) Set(val *ValidateQuoteResponseLinesInnerVmfAdditionalAttributesLinesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableValidateQuoteResponseLinesInnerVmfAdditionalAttributesLinesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableValidateQuoteResponseLinesInnerVmfAdditionalAttributesLinesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidateQuoteResponseLinesInnerVmfAdditionalAttributesLinesInner(val *ValidateQuoteResponseLinesInnerVmfAdditionalAttributesLinesInner) *NullableValidateQuoteResponseLinesInnerVmfAdditionalAttributesLinesInner {
	return &NullableValidateQuoteResponseLinesInnerVmfAdditionalAttributesLinesInner{value: val, isSet: true}
}

func (v NullableValidateQuoteResponseLinesInnerVmfAdditionalAttributesLinesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidateQuoteResponseLinesInnerVmfAdditionalAttributesLinesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


