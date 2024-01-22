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

// checks if the OrderModifyRequestLinesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderModifyRequestLinesInner{}

// OrderModifyRequestLinesInner struct for OrderModifyRequestLinesInner
type OrderModifyRequestLinesInner struct {
	// The unique IngramMicro part number.
	IngramPartNumber *string `json:"ingramPartNumber,omitempty"`
	// The IngramMicro line number.
	IngramLineNumber *string `json:"ingramLineNumber,omitempty"`
	// The reseller's line number for reference in their system.
	CustomerLineNumber *string `json:"customerLineNumber,omitempty"`
	// The line number that was added, updated, or deleted.
	AddUpdateDeleteLine *string `json:"addUpdateDeleteLine,omitempty"`
	// The quantity of the line item.
	Quantity *int32 `json:"quantity,omitempty"`
	// The line-level notes.
	Notes *string `json:"notes,omitempty"`
}

// NewOrderModifyRequestLinesInner instantiates a new OrderModifyRequestLinesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderModifyRequestLinesInner() *OrderModifyRequestLinesInner {
	this := OrderModifyRequestLinesInner{}
	return &this
}

// NewOrderModifyRequestLinesInnerWithDefaults instantiates a new OrderModifyRequestLinesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderModifyRequestLinesInnerWithDefaults() *OrderModifyRequestLinesInner {
	this := OrderModifyRequestLinesInner{}
	return &this
}

// GetIngramPartNumber returns the IngramPartNumber field value if set, zero value otherwise.
func (o *OrderModifyRequestLinesInner) GetIngramPartNumber() string {
	if o == nil || IsNil(o.IngramPartNumber) {
		var ret string
		return ret
	}
	return *o.IngramPartNumber
}

// GetIngramPartNumberOk returns a tuple with the IngramPartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderModifyRequestLinesInner) GetIngramPartNumberOk() (*string, bool) {
	if o == nil || IsNil(o.IngramPartNumber) {
		return nil, false
	}
	return o.IngramPartNumber, true
}

// HasIngramPartNumber returns a boolean if a field has been set.
func (o *OrderModifyRequestLinesInner) HasIngramPartNumber() bool {
	if o != nil && !IsNil(o.IngramPartNumber) {
		return true
	}

	return false
}

// SetIngramPartNumber gets a reference to the given string and assigns it to the IngramPartNumber field.
func (o *OrderModifyRequestLinesInner) SetIngramPartNumber(v string) {
	o.IngramPartNumber = &v
}

// GetIngramLineNumber returns the IngramLineNumber field value if set, zero value otherwise.
func (o *OrderModifyRequestLinesInner) GetIngramLineNumber() string {
	if o == nil || IsNil(o.IngramLineNumber) {
		var ret string
		return ret
	}
	return *o.IngramLineNumber
}

// GetIngramLineNumberOk returns a tuple with the IngramLineNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderModifyRequestLinesInner) GetIngramLineNumberOk() (*string, bool) {
	if o == nil || IsNil(o.IngramLineNumber) {
		return nil, false
	}
	return o.IngramLineNumber, true
}

// HasIngramLineNumber returns a boolean if a field has been set.
func (o *OrderModifyRequestLinesInner) HasIngramLineNumber() bool {
	if o != nil && !IsNil(o.IngramLineNumber) {
		return true
	}

	return false
}

// SetIngramLineNumber gets a reference to the given string and assigns it to the IngramLineNumber field.
func (o *OrderModifyRequestLinesInner) SetIngramLineNumber(v string) {
	o.IngramLineNumber = &v
}

// GetCustomerLineNumber returns the CustomerLineNumber field value if set, zero value otherwise.
func (o *OrderModifyRequestLinesInner) GetCustomerLineNumber() string {
	if o == nil || IsNil(o.CustomerLineNumber) {
		var ret string
		return ret
	}
	return *o.CustomerLineNumber
}

// GetCustomerLineNumberOk returns a tuple with the CustomerLineNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderModifyRequestLinesInner) GetCustomerLineNumberOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerLineNumber) {
		return nil, false
	}
	return o.CustomerLineNumber, true
}

// HasCustomerLineNumber returns a boolean if a field has been set.
func (o *OrderModifyRequestLinesInner) HasCustomerLineNumber() bool {
	if o != nil && !IsNil(o.CustomerLineNumber) {
		return true
	}

	return false
}

// SetCustomerLineNumber gets a reference to the given string and assigns it to the CustomerLineNumber field.
func (o *OrderModifyRequestLinesInner) SetCustomerLineNumber(v string) {
	o.CustomerLineNumber = &v
}

// GetAddUpdateDeleteLine returns the AddUpdateDeleteLine field value if set, zero value otherwise.
func (o *OrderModifyRequestLinesInner) GetAddUpdateDeleteLine() string {
	if o == nil || IsNil(o.AddUpdateDeleteLine) {
		var ret string
		return ret
	}
	return *o.AddUpdateDeleteLine
}

// GetAddUpdateDeleteLineOk returns a tuple with the AddUpdateDeleteLine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderModifyRequestLinesInner) GetAddUpdateDeleteLineOk() (*string, bool) {
	if o == nil || IsNil(o.AddUpdateDeleteLine) {
		return nil, false
	}
	return o.AddUpdateDeleteLine, true
}

// HasAddUpdateDeleteLine returns a boolean if a field has been set.
func (o *OrderModifyRequestLinesInner) HasAddUpdateDeleteLine() bool {
	if o != nil && !IsNil(o.AddUpdateDeleteLine) {
		return true
	}

	return false
}

// SetAddUpdateDeleteLine gets a reference to the given string and assigns it to the AddUpdateDeleteLine field.
func (o *OrderModifyRequestLinesInner) SetAddUpdateDeleteLine(v string) {
	o.AddUpdateDeleteLine = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *OrderModifyRequestLinesInner) GetQuantity() int32 {
	if o == nil || IsNil(o.Quantity) {
		var ret int32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderModifyRequestLinesInner) GetQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *OrderModifyRequestLinesInner) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int32 and assigns it to the Quantity field.
func (o *OrderModifyRequestLinesInner) SetQuantity(v int32) {
	o.Quantity = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *OrderModifyRequestLinesInner) GetNotes() string {
	if o == nil || IsNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderModifyRequestLinesInner) GetNotesOk() (*string, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *OrderModifyRequestLinesInner) HasNotes() bool {
	if o != nil && !IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *OrderModifyRequestLinesInner) SetNotes(v string) {
	o.Notes = &v
}

func (o OrderModifyRequestLinesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderModifyRequestLinesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IngramPartNumber) {
		toSerialize["ingramPartNumber"] = o.IngramPartNumber
	}
	if !IsNil(o.IngramLineNumber) {
		toSerialize["ingramLineNumber"] = o.IngramLineNumber
	}
	if !IsNil(o.CustomerLineNumber) {
		toSerialize["customerLineNumber"] = o.CustomerLineNumber
	}
	if !IsNil(o.AddUpdateDeleteLine) {
		toSerialize["addUpdateDeleteLine"] = o.AddUpdateDeleteLine
	}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if !IsNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	return toSerialize, nil
}

type NullableOrderModifyRequestLinesInner struct {
	value *OrderModifyRequestLinesInner
	isSet bool
}

func (v NullableOrderModifyRequestLinesInner) Get() *OrderModifyRequestLinesInner {
	return v.value
}

func (v *NullableOrderModifyRequestLinesInner) Set(val *OrderModifyRequestLinesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderModifyRequestLinesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderModifyRequestLinesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderModifyRequestLinesInner(val *OrderModifyRequestLinesInner) *NullableOrderModifyRequestLinesInner {
	return &NullableOrderModifyRequestLinesInner{value: val, isSet: true}
}

func (v NullableOrderModifyRequestLinesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderModifyRequestLinesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


