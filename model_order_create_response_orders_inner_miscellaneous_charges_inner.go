/*
XI Sdk Resellers

For Resellers. Who are looking to Innovate with Ingram Micro's API SolutionsAutomate your eCommerce with our offering of APIs and Webhooks to create a seamless experience for your customers.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package xi_sdk_resellers

import (
	"encoding/json"
)

// checks if the OrderCreateResponseOrdersInnerMiscellaneousChargesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderCreateResponseOrdersInnerMiscellaneousChargesInner{}

// OrderCreateResponseOrdersInnerMiscellaneousChargesInner struct for OrderCreateResponseOrdersInnerMiscellaneousChargesInner
type OrderCreateResponseOrdersInnerMiscellaneousChargesInner struct {
	// The sub order number. The two-digit prefix is the warehouse code of the warehouse nearest the reseller. The middle number is the order number. The two-digit suffix is the sub order number.
	SubOrderNumber *string `json:"subOrderNumber,omitempty"`
	// Impulse line number for the miscellaneous charge.
	ChargeLineReference *string `json:"chargeLineReference,omitempty"`
	// Description of the miscellaneous charges for the order.
	ChargeDescription *string `json:"chargeDescription,omitempty"`
	// The total amount of miscellaneous charges for the order.
	ChargeAmount *float32 `json:"chargeAmount,omitempty"`
}

// NewOrderCreateResponseOrdersInnerMiscellaneousChargesInner instantiates a new OrderCreateResponseOrdersInnerMiscellaneousChargesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderCreateResponseOrdersInnerMiscellaneousChargesInner() *OrderCreateResponseOrdersInnerMiscellaneousChargesInner {
	this := OrderCreateResponseOrdersInnerMiscellaneousChargesInner{}
	return &this
}

// NewOrderCreateResponseOrdersInnerMiscellaneousChargesInnerWithDefaults instantiates a new OrderCreateResponseOrdersInnerMiscellaneousChargesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderCreateResponseOrdersInnerMiscellaneousChargesInnerWithDefaults() *OrderCreateResponseOrdersInnerMiscellaneousChargesInner {
	this := OrderCreateResponseOrdersInnerMiscellaneousChargesInner{}
	return &this
}

// GetSubOrderNumber returns the SubOrderNumber field value if set, zero value otherwise.
func (o *OrderCreateResponseOrdersInnerMiscellaneousChargesInner) GetSubOrderNumber() string {
	if o == nil || IsNil(o.SubOrderNumber) {
		var ret string
		return ret
	}
	return *o.SubOrderNumber
}

// GetSubOrderNumberOk returns a tuple with the SubOrderNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateResponseOrdersInnerMiscellaneousChargesInner) GetSubOrderNumberOk() (*string, bool) {
	if o == nil || IsNil(o.SubOrderNumber) {
		return nil, false
	}
	return o.SubOrderNumber, true
}

// HasSubOrderNumber returns a boolean if a field has been set.
func (o *OrderCreateResponseOrdersInnerMiscellaneousChargesInner) HasSubOrderNumber() bool {
	if o != nil && !IsNil(o.SubOrderNumber) {
		return true
	}

	return false
}

// SetSubOrderNumber gets a reference to the given string and assigns it to the SubOrderNumber field.
func (o *OrderCreateResponseOrdersInnerMiscellaneousChargesInner) SetSubOrderNumber(v string) {
	o.SubOrderNumber = &v
}

// GetChargeLineReference returns the ChargeLineReference field value if set, zero value otherwise.
func (o *OrderCreateResponseOrdersInnerMiscellaneousChargesInner) GetChargeLineReference() string {
	if o == nil || IsNil(o.ChargeLineReference) {
		var ret string
		return ret
	}
	return *o.ChargeLineReference
}

// GetChargeLineReferenceOk returns a tuple with the ChargeLineReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateResponseOrdersInnerMiscellaneousChargesInner) GetChargeLineReferenceOk() (*string, bool) {
	if o == nil || IsNil(o.ChargeLineReference) {
		return nil, false
	}
	return o.ChargeLineReference, true
}

// HasChargeLineReference returns a boolean if a field has been set.
func (o *OrderCreateResponseOrdersInnerMiscellaneousChargesInner) HasChargeLineReference() bool {
	if o != nil && !IsNil(o.ChargeLineReference) {
		return true
	}

	return false
}

// SetChargeLineReference gets a reference to the given string and assigns it to the ChargeLineReference field.
func (o *OrderCreateResponseOrdersInnerMiscellaneousChargesInner) SetChargeLineReference(v string) {
	o.ChargeLineReference = &v
}

// GetChargeDescription returns the ChargeDescription field value if set, zero value otherwise.
func (o *OrderCreateResponseOrdersInnerMiscellaneousChargesInner) GetChargeDescription() string {
	if o == nil || IsNil(o.ChargeDescription) {
		var ret string
		return ret
	}
	return *o.ChargeDescription
}

// GetChargeDescriptionOk returns a tuple with the ChargeDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateResponseOrdersInnerMiscellaneousChargesInner) GetChargeDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.ChargeDescription) {
		return nil, false
	}
	return o.ChargeDescription, true
}

// HasChargeDescription returns a boolean if a field has been set.
func (o *OrderCreateResponseOrdersInnerMiscellaneousChargesInner) HasChargeDescription() bool {
	if o != nil && !IsNil(o.ChargeDescription) {
		return true
	}

	return false
}

// SetChargeDescription gets a reference to the given string and assigns it to the ChargeDescription field.
func (o *OrderCreateResponseOrdersInnerMiscellaneousChargesInner) SetChargeDescription(v string) {
	o.ChargeDescription = &v
}

// GetChargeAmount returns the ChargeAmount field value if set, zero value otherwise.
func (o *OrderCreateResponseOrdersInnerMiscellaneousChargesInner) GetChargeAmount() float32 {
	if o == nil || IsNil(o.ChargeAmount) {
		var ret float32
		return ret
	}
	return *o.ChargeAmount
}

// GetChargeAmountOk returns a tuple with the ChargeAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateResponseOrdersInnerMiscellaneousChargesInner) GetChargeAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.ChargeAmount) {
		return nil, false
	}
	return o.ChargeAmount, true
}

// HasChargeAmount returns a boolean if a field has been set.
func (o *OrderCreateResponseOrdersInnerMiscellaneousChargesInner) HasChargeAmount() bool {
	if o != nil && !IsNil(o.ChargeAmount) {
		return true
	}

	return false
}

// SetChargeAmount gets a reference to the given float32 and assigns it to the ChargeAmount field.
func (o *OrderCreateResponseOrdersInnerMiscellaneousChargesInner) SetChargeAmount(v float32) {
	o.ChargeAmount = &v
}

func (o OrderCreateResponseOrdersInnerMiscellaneousChargesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderCreateResponseOrdersInnerMiscellaneousChargesInner) ToMap() (map[string]interface{}, error) {
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

type NullableOrderCreateResponseOrdersInnerMiscellaneousChargesInner struct {
	value *OrderCreateResponseOrdersInnerMiscellaneousChargesInner
	isSet bool
}

func (v NullableOrderCreateResponseOrdersInnerMiscellaneousChargesInner) Get() *OrderCreateResponseOrdersInnerMiscellaneousChargesInner {
	return v.value
}

func (v *NullableOrderCreateResponseOrdersInnerMiscellaneousChargesInner) Set(val *OrderCreateResponseOrdersInnerMiscellaneousChargesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderCreateResponseOrdersInnerMiscellaneousChargesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderCreateResponseOrdersInnerMiscellaneousChargesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderCreateResponseOrdersInnerMiscellaneousChargesInner(val *OrderCreateResponseOrdersInnerMiscellaneousChargesInner) *NullableOrderCreateResponseOrdersInnerMiscellaneousChargesInner {
	return &NullableOrderCreateResponseOrdersInnerMiscellaneousChargesInner{value: val, isSet: true}
}

func (v NullableOrderCreateResponseOrdersInnerMiscellaneousChargesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderCreateResponseOrdersInnerMiscellaneousChargesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


