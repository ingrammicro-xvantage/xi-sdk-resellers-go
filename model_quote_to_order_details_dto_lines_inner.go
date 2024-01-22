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

// checks if the QuoteToOrderDetailsDTOLinesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuoteToOrderDetailsDTOLinesInner{}

// QuoteToOrderDetailsDTOLinesInner struct for QuoteToOrderDetailsDTOLinesInner
type QuoteToOrderDetailsDTOLinesInner struct {
	// The reseller's line item number for reference in their system.
	CustomerLineNumber *string `json:"customerLineNumber,omitempty"`
	// Unique IngramMicro part number.
	IngramPartNumber *string `json:"ingramPartNumber,omitempty"`
	// The quantity of the line item.
	Quantity *string `json:"quantity,omitempty"`
	// The object containing the list of fields required at a line level by the vendor.
	VmfAdditionalAttributesLines []ValidateQuoteResponseLinesInnerVmfAdditionalAttributesLinesInner `json:"vmfAdditionalAttributesLines,omitempty"`
}

// NewQuoteToOrderDetailsDTOLinesInner instantiates a new QuoteToOrderDetailsDTOLinesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuoteToOrderDetailsDTOLinesInner() *QuoteToOrderDetailsDTOLinesInner {
	this := QuoteToOrderDetailsDTOLinesInner{}
	return &this
}

// NewQuoteToOrderDetailsDTOLinesInnerWithDefaults instantiates a new QuoteToOrderDetailsDTOLinesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuoteToOrderDetailsDTOLinesInnerWithDefaults() *QuoteToOrderDetailsDTOLinesInner {
	this := QuoteToOrderDetailsDTOLinesInner{}
	return &this
}

// GetCustomerLineNumber returns the CustomerLineNumber field value if set, zero value otherwise.
func (o *QuoteToOrderDetailsDTOLinesInner) GetCustomerLineNumber() string {
	if o == nil || IsNil(o.CustomerLineNumber) {
		var ret string
		return ret
	}
	return *o.CustomerLineNumber
}

// GetCustomerLineNumberOk returns a tuple with the CustomerLineNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteToOrderDetailsDTOLinesInner) GetCustomerLineNumberOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerLineNumber) {
		return nil, false
	}
	return o.CustomerLineNumber, true
}

// HasCustomerLineNumber returns a boolean if a field has been set.
func (o *QuoteToOrderDetailsDTOLinesInner) HasCustomerLineNumber() bool {
	if o != nil && !IsNil(o.CustomerLineNumber) {
		return true
	}

	return false
}

// SetCustomerLineNumber gets a reference to the given string and assigns it to the CustomerLineNumber field.
func (o *QuoteToOrderDetailsDTOLinesInner) SetCustomerLineNumber(v string) {
	o.CustomerLineNumber = &v
}

// GetIngramPartNumber returns the IngramPartNumber field value if set, zero value otherwise.
func (o *QuoteToOrderDetailsDTOLinesInner) GetIngramPartNumber() string {
	if o == nil || IsNil(o.IngramPartNumber) {
		var ret string
		return ret
	}
	return *o.IngramPartNumber
}

// GetIngramPartNumberOk returns a tuple with the IngramPartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteToOrderDetailsDTOLinesInner) GetIngramPartNumberOk() (*string, bool) {
	if o == nil || IsNil(o.IngramPartNumber) {
		return nil, false
	}
	return o.IngramPartNumber, true
}

// HasIngramPartNumber returns a boolean if a field has been set.
func (o *QuoteToOrderDetailsDTOLinesInner) HasIngramPartNumber() bool {
	if o != nil && !IsNil(o.IngramPartNumber) {
		return true
	}

	return false
}

// SetIngramPartNumber gets a reference to the given string and assigns it to the IngramPartNumber field.
func (o *QuoteToOrderDetailsDTOLinesInner) SetIngramPartNumber(v string) {
	o.IngramPartNumber = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *QuoteToOrderDetailsDTOLinesInner) GetQuantity() string {
	if o == nil || IsNil(o.Quantity) {
		var ret string
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteToOrderDetailsDTOLinesInner) GetQuantityOk() (*string, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *QuoteToOrderDetailsDTOLinesInner) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given string and assigns it to the Quantity field.
func (o *QuoteToOrderDetailsDTOLinesInner) SetQuantity(v string) {
	o.Quantity = &v
}

// GetVmfAdditionalAttributesLines returns the VmfAdditionalAttributesLines field value if set, zero value otherwise.
func (o *QuoteToOrderDetailsDTOLinesInner) GetVmfAdditionalAttributesLines() []ValidateQuoteResponseLinesInnerVmfAdditionalAttributesLinesInner {
	if o == nil || IsNil(o.VmfAdditionalAttributesLines) {
		var ret []ValidateQuoteResponseLinesInnerVmfAdditionalAttributesLinesInner
		return ret
	}
	return o.VmfAdditionalAttributesLines
}

// GetVmfAdditionalAttributesLinesOk returns a tuple with the VmfAdditionalAttributesLines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteToOrderDetailsDTOLinesInner) GetVmfAdditionalAttributesLinesOk() ([]ValidateQuoteResponseLinesInnerVmfAdditionalAttributesLinesInner, bool) {
	if o == nil || IsNil(o.VmfAdditionalAttributesLines) {
		return nil, false
	}
	return o.VmfAdditionalAttributesLines, true
}

// HasVmfAdditionalAttributesLines returns a boolean if a field has been set.
func (o *QuoteToOrderDetailsDTOLinesInner) HasVmfAdditionalAttributesLines() bool {
	if o != nil && !IsNil(o.VmfAdditionalAttributesLines) {
		return true
	}

	return false
}

// SetVmfAdditionalAttributesLines gets a reference to the given []ValidateQuoteResponseLinesInnerVmfAdditionalAttributesLinesInner and assigns it to the VmfAdditionalAttributesLines field.
func (o *QuoteToOrderDetailsDTOLinesInner) SetVmfAdditionalAttributesLines(v []ValidateQuoteResponseLinesInnerVmfAdditionalAttributesLinesInner) {
	o.VmfAdditionalAttributesLines = v
}

func (o QuoteToOrderDetailsDTOLinesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuoteToOrderDetailsDTOLinesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CustomerLineNumber) {
		toSerialize["customerLineNumber"] = o.CustomerLineNumber
	}
	if !IsNil(o.IngramPartNumber) {
		toSerialize["ingramPartNumber"] = o.IngramPartNumber
	}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if !IsNil(o.VmfAdditionalAttributesLines) {
		toSerialize["vmfAdditionalAttributesLines"] = o.VmfAdditionalAttributesLines
	}
	return toSerialize, nil
}

type NullableQuoteToOrderDetailsDTOLinesInner struct {
	value *QuoteToOrderDetailsDTOLinesInner
	isSet bool
}

func (v NullableQuoteToOrderDetailsDTOLinesInner) Get() *QuoteToOrderDetailsDTOLinesInner {
	return v.value
}

func (v *NullableQuoteToOrderDetailsDTOLinesInner) Set(val *QuoteToOrderDetailsDTOLinesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQuoteToOrderDetailsDTOLinesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQuoteToOrderDetailsDTOLinesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuoteToOrderDetailsDTOLinesInner(val *QuoteToOrderDetailsDTOLinesInner) *NullableQuoteToOrderDetailsDTOLinesInner {
	return &NullableQuoteToOrderDetailsDTOLinesInner{value: val, isSet: true}
}

func (v NullableQuoteToOrderDetailsDTOLinesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuoteToOrderDetailsDTOLinesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


