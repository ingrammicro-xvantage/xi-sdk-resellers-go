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

// checks if the OrderDetailB2BMiscellaneousChargesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderDetailB2BMiscellaneousChargesInner{}

// OrderDetailB2BMiscellaneousChargesInner struct for OrderDetailB2BMiscellaneousChargesInner
type OrderDetailB2BMiscellaneousChargesInner struct {
	// The sub order number. The two-digit prefix is the warehouse code of the warehouse nearest the reseller. The middle number is the order number. The two-digit suffix is the sub order number.
	SubOrderNumber *string `json:"subOrderNumber,omitempty"`
	// Impulse line number for the miscellaneous charge.
	ChargeLineReference *string `json:"chargeLineReference,omitempty"`
	// Description of the miscellaneous charges.
	ChargeDescription *string `json:"chargeDescription,omitempty"`
	// The amount of miscellaneous charges.
	ChargeAmount *string `json:"chargeAmount,omitempty"`
}

// NewOrderDetailB2BMiscellaneousChargesInner instantiates a new OrderDetailB2BMiscellaneousChargesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderDetailB2BMiscellaneousChargesInner() *OrderDetailB2BMiscellaneousChargesInner {
	this := OrderDetailB2BMiscellaneousChargesInner{}
	return &this
}

// NewOrderDetailB2BMiscellaneousChargesInnerWithDefaults instantiates a new OrderDetailB2BMiscellaneousChargesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderDetailB2BMiscellaneousChargesInnerWithDefaults() *OrderDetailB2BMiscellaneousChargesInner {
	this := OrderDetailB2BMiscellaneousChargesInner{}
	return &this
}

// GetSubOrderNumber returns the SubOrderNumber field value if set, zero value otherwise.
func (o *OrderDetailB2BMiscellaneousChargesInner) GetSubOrderNumber() string {
	if o == nil || IsNil(o.SubOrderNumber) {
		var ret string
		return ret
	}
	return *o.SubOrderNumber
}

// GetSubOrderNumberOk returns a tuple with the SubOrderNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailB2BMiscellaneousChargesInner) GetSubOrderNumberOk() (*string, bool) {
	if o == nil || IsNil(o.SubOrderNumber) {
		return nil, false
	}
	return o.SubOrderNumber, true
}

// HasSubOrderNumber returns a boolean if a field has been set.
func (o *OrderDetailB2BMiscellaneousChargesInner) HasSubOrderNumber() bool {
	if o != nil && !IsNil(o.SubOrderNumber) {
		return true
	}

	return false
}

// SetSubOrderNumber gets a reference to the given string and assigns it to the SubOrderNumber field.
func (o *OrderDetailB2BMiscellaneousChargesInner) SetSubOrderNumber(v string) {
	o.SubOrderNumber = &v
}

// GetChargeLineReference returns the ChargeLineReference field value if set, zero value otherwise.
func (o *OrderDetailB2BMiscellaneousChargesInner) GetChargeLineReference() string {
	if o == nil || IsNil(o.ChargeLineReference) {
		var ret string
		return ret
	}
	return *o.ChargeLineReference
}

// GetChargeLineReferenceOk returns a tuple with the ChargeLineReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailB2BMiscellaneousChargesInner) GetChargeLineReferenceOk() (*string, bool) {
	if o == nil || IsNil(o.ChargeLineReference) {
		return nil, false
	}
	return o.ChargeLineReference, true
}

// HasChargeLineReference returns a boolean if a field has been set.
func (o *OrderDetailB2BMiscellaneousChargesInner) HasChargeLineReference() bool {
	if o != nil && !IsNil(o.ChargeLineReference) {
		return true
	}

	return false
}

// SetChargeLineReference gets a reference to the given string and assigns it to the ChargeLineReference field.
func (o *OrderDetailB2BMiscellaneousChargesInner) SetChargeLineReference(v string) {
	o.ChargeLineReference = &v
}

// GetChargeDescription returns the ChargeDescription field value if set, zero value otherwise.
func (o *OrderDetailB2BMiscellaneousChargesInner) GetChargeDescription() string {
	if o == nil || IsNil(o.ChargeDescription) {
		var ret string
		return ret
	}
	return *o.ChargeDescription
}

// GetChargeDescriptionOk returns a tuple with the ChargeDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailB2BMiscellaneousChargesInner) GetChargeDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.ChargeDescription) {
		return nil, false
	}
	return o.ChargeDescription, true
}

// HasChargeDescription returns a boolean if a field has been set.
func (o *OrderDetailB2BMiscellaneousChargesInner) HasChargeDescription() bool {
	if o != nil && !IsNil(o.ChargeDescription) {
		return true
	}

	return false
}

// SetChargeDescription gets a reference to the given string and assigns it to the ChargeDescription field.
func (o *OrderDetailB2BMiscellaneousChargesInner) SetChargeDescription(v string) {
	o.ChargeDescription = &v
}

// GetChargeAmount returns the ChargeAmount field value if set, zero value otherwise.
func (o *OrderDetailB2BMiscellaneousChargesInner) GetChargeAmount() string {
	if o == nil || IsNil(o.ChargeAmount) {
		var ret string
		return ret
	}
	return *o.ChargeAmount
}

// GetChargeAmountOk returns a tuple with the ChargeAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailB2BMiscellaneousChargesInner) GetChargeAmountOk() (*string, bool) {
	if o == nil || IsNil(o.ChargeAmount) {
		return nil, false
	}
	return o.ChargeAmount, true
}

// HasChargeAmount returns a boolean if a field has been set.
func (o *OrderDetailB2BMiscellaneousChargesInner) HasChargeAmount() bool {
	if o != nil && !IsNil(o.ChargeAmount) {
		return true
	}

	return false
}

// SetChargeAmount gets a reference to the given string and assigns it to the ChargeAmount field.
func (o *OrderDetailB2BMiscellaneousChargesInner) SetChargeAmount(v string) {
	o.ChargeAmount = &v
}

func (o OrderDetailB2BMiscellaneousChargesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderDetailB2BMiscellaneousChargesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SubOrderNumber) {
		toSerialize["subOrderNumber"] = o.SubOrderNumber
	}
	if !IsNil(o.ChargeLineReference) {
		toSerialize["chargeLineReference"] = o.ChargeLineReference
	}
	if !IsNil(o.ChargeDescription) {
		toSerialize["chargeDescription"] = o.ChargeDescription
	}
	if !IsNil(o.ChargeAmount) {
		toSerialize["chargeAmount"] = o.ChargeAmount
	}
	return toSerialize, nil
}

type NullableOrderDetailB2BMiscellaneousChargesInner struct {
	value *OrderDetailB2BMiscellaneousChargesInner
	isSet bool
}

func (v NullableOrderDetailB2BMiscellaneousChargesInner) Get() *OrderDetailB2BMiscellaneousChargesInner {
	return v.value
}

func (v *NullableOrderDetailB2BMiscellaneousChargesInner) Set(val *OrderDetailB2BMiscellaneousChargesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderDetailB2BMiscellaneousChargesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderDetailB2BMiscellaneousChargesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderDetailB2BMiscellaneousChargesInner(val *OrderDetailB2BMiscellaneousChargesInner) *NullableOrderDetailB2BMiscellaneousChargesInner {
	return &NullableOrderDetailB2BMiscellaneousChargesInner{value: val, isSet: true}
}

func (v NullableOrderDetailB2BMiscellaneousChargesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderDetailB2BMiscellaneousChargesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


