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

// checks if the InvoiceDetailsv61ResponseSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvoiceDetailsv61ResponseSummary{}

// InvoiceDetailsv61ResponseSummary struct for InvoiceDetailsv61ResponseSummary
type InvoiceDetailsv61ResponseSummary struct {
	Lines *InvoiceDetailsv61ResponseSummaryLines `json:"lines,omitempty"`
	// Miscellaneous charges.
	MiscCharges []InvoiceDetailsv61ResponseSummaryMiscChargesInner `json:"miscCharges,omitempty"`
	Totals *InvoiceDetailsv61ResponseSummaryTotals `json:"totals,omitempty"`
	ForeignFxTotals *InvoiceDetailsv61ResponseSummaryForeignFxTotals `json:"foreignFxTotals,omitempty"`
}

// NewInvoiceDetailsv61ResponseSummary instantiates a new InvoiceDetailsv61ResponseSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoiceDetailsv61ResponseSummary() *InvoiceDetailsv61ResponseSummary {
	this := InvoiceDetailsv61ResponseSummary{}
	return &this
}

// NewInvoiceDetailsv61ResponseSummaryWithDefaults instantiates a new InvoiceDetailsv61ResponseSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoiceDetailsv61ResponseSummaryWithDefaults() *InvoiceDetailsv61ResponseSummary {
	this := InvoiceDetailsv61ResponseSummary{}
	return &this
}

// GetLines returns the Lines field value if set, zero value otherwise.
func (o *InvoiceDetailsv61ResponseSummary) GetLines() InvoiceDetailsv61ResponseSummaryLines {
	if o == nil || IsNil(o.Lines) {
		var ret InvoiceDetailsv61ResponseSummaryLines
		return ret
	}
	return *o.Lines
}

// GetLinesOk returns a tuple with the Lines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceDetailsv61ResponseSummary) GetLinesOk() (*InvoiceDetailsv61ResponseSummaryLines, bool) {
	if o == nil || IsNil(o.Lines) {
		return nil, false
	}
	return o.Lines, true
}

// HasLines returns a boolean if a field has been set.
func (o *InvoiceDetailsv61ResponseSummary) HasLines() bool {
	if o != nil && !IsNil(o.Lines) {
		return true
	}

	return false
}

// SetLines gets a reference to the given InvoiceDetailsv61ResponseSummaryLines and assigns it to the Lines field.
func (o *InvoiceDetailsv61ResponseSummary) SetLines(v InvoiceDetailsv61ResponseSummaryLines) {
	o.Lines = &v
}

// GetMiscCharges returns the MiscCharges field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvoiceDetailsv61ResponseSummary) GetMiscCharges() []InvoiceDetailsv61ResponseSummaryMiscChargesInner {
	if o == nil {
		var ret []InvoiceDetailsv61ResponseSummaryMiscChargesInner
		return ret
	}
	return o.MiscCharges
}

// GetMiscChargesOk returns a tuple with the MiscCharges field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceDetailsv61ResponseSummary) GetMiscChargesOk() ([]InvoiceDetailsv61ResponseSummaryMiscChargesInner, bool) {
	if o == nil || IsNil(o.MiscCharges) {
		return nil, false
	}
	return o.MiscCharges, true
}

// HasMiscCharges returns a boolean if a field has been set.
func (o *InvoiceDetailsv61ResponseSummary) HasMiscCharges() bool {
	if o != nil && IsNil(o.MiscCharges) {
		return true
	}

	return false
}

// SetMiscCharges gets a reference to the given []InvoiceDetailsv61ResponseSummaryMiscChargesInner and assigns it to the MiscCharges field.
func (o *InvoiceDetailsv61ResponseSummary) SetMiscCharges(v []InvoiceDetailsv61ResponseSummaryMiscChargesInner) {
	o.MiscCharges = v
}

// GetTotals returns the Totals field value if set, zero value otherwise.
func (o *InvoiceDetailsv61ResponseSummary) GetTotals() InvoiceDetailsv61ResponseSummaryTotals {
	if o == nil || IsNil(o.Totals) {
		var ret InvoiceDetailsv61ResponseSummaryTotals
		return ret
	}
	return *o.Totals
}

// GetTotalsOk returns a tuple with the Totals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceDetailsv61ResponseSummary) GetTotalsOk() (*InvoiceDetailsv61ResponseSummaryTotals, bool) {
	if o == nil || IsNil(o.Totals) {
		return nil, false
	}
	return o.Totals, true
}

// HasTotals returns a boolean if a field has been set.
func (o *InvoiceDetailsv61ResponseSummary) HasTotals() bool {
	if o != nil && !IsNil(o.Totals) {
		return true
	}

	return false
}

// SetTotals gets a reference to the given InvoiceDetailsv61ResponseSummaryTotals and assigns it to the Totals field.
func (o *InvoiceDetailsv61ResponseSummary) SetTotals(v InvoiceDetailsv61ResponseSummaryTotals) {
	o.Totals = &v
}

// GetForeignFxTotals returns the ForeignFxTotals field value if set, zero value otherwise.
func (o *InvoiceDetailsv61ResponseSummary) GetForeignFxTotals() InvoiceDetailsv61ResponseSummaryForeignFxTotals {
	if o == nil || IsNil(o.ForeignFxTotals) {
		var ret InvoiceDetailsv61ResponseSummaryForeignFxTotals
		return ret
	}
	return *o.ForeignFxTotals
}

// GetForeignFxTotalsOk returns a tuple with the ForeignFxTotals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceDetailsv61ResponseSummary) GetForeignFxTotalsOk() (*InvoiceDetailsv61ResponseSummaryForeignFxTotals, bool) {
	if o == nil || IsNil(o.ForeignFxTotals) {
		return nil, false
	}
	return o.ForeignFxTotals, true
}

// HasForeignFxTotals returns a boolean if a field has been set.
func (o *InvoiceDetailsv61ResponseSummary) HasForeignFxTotals() bool {
	if o != nil && !IsNil(o.ForeignFxTotals) {
		return true
	}

	return false
}

// SetForeignFxTotals gets a reference to the given InvoiceDetailsv61ResponseSummaryForeignFxTotals and assigns it to the ForeignFxTotals field.
func (o *InvoiceDetailsv61ResponseSummary) SetForeignFxTotals(v InvoiceDetailsv61ResponseSummaryForeignFxTotals) {
	o.ForeignFxTotals = &v
}

func (o InvoiceDetailsv61ResponseSummary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvoiceDetailsv61ResponseSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Lines) {
		toSerialize["lines"] = o.Lines
	}
	if o.MiscCharges != nil {
		toSerialize["miscCharges"] = o.MiscCharges
	}
	if !IsNil(o.Totals) {
		toSerialize["totals"] = o.Totals
	}
	if !IsNil(o.ForeignFxTotals) {
		toSerialize["foreignFxTotals"] = o.ForeignFxTotals
	}
	return toSerialize, nil
}

type NullableInvoiceDetailsv61ResponseSummary struct {
	value *InvoiceDetailsv61ResponseSummary
	isSet bool
}

func (v NullableInvoiceDetailsv61ResponseSummary) Get() *InvoiceDetailsv61ResponseSummary {
	return v.value
}

func (v *NullableInvoiceDetailsv61ResponseSummary) Set(val *InvoiceDetailsv61ResponseSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoiceDetailsv61ResponseSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoiceDetailsv61ResponseSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoiceDetailsv61ResponseSummary(val *InvoiceDetailsv61ResponseSummary) *NullableInvoiceDetailsv61ResponseSummary {
	return &NullableInvoiceDetailsv61ResponseSummary{value: val, isSet: true}
}

func (v NullableInvoiceDetailsv61ResponseSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoiceDetailsv61ResponseSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


