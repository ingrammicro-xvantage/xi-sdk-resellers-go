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

// checks if the PriceAndAvailabilityResponseInnerDiscountsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PriceAndAvailabilityResponseInnerDiscountsInner{}

// PriceAndAvailabilityResponseInnerDiscountsInner struct for PriceAndAvailabilityResponseInnerDiscountsInner
type PriceAndAvailabilityResponseInnerDiscountsInner struct {
	SpecialPricing []PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner `json:"specialPricing,omitempty"`
	QuantityDiscounts []PriceAndAvailabilityResponseInnerDiscountsInnerQuantityDiscountsInner `json:"quantityDiscounts,omitempty"`
}

// NewPriceAndAvailabilityResponseInnerDiscountsInner instantiates a new PriceAndAvailabilityResponseInnerDiscountsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPriceAndAvailabilityResponseInnerDiscountsInner() *PriceAndAvailabilityResponseInnerDiscountsInner {
	this := PriceAndAvailabilityResponseInnerDiscountsInner{}
	return &this
}

// NewPriceAndAvailabilityResponseInnerDiscountsInnerWithDefaults instantiates a new PriceAndAvailabilityResponseInnerDiscountsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceAndAvailabilityResponseInnerDiscountsInnerWithDefaults() *PriceAndAvailabilityResponseInnerDiscountsInner {
	this := PriceAndAvailabilityResponseInnerDiscountsInner{}
	return &this
}

// GetSpecialPricing returns the SpecialPricing field value if set, zero value otherwise.
func (o *PriceAndAvailabilityResponseInnerDiscountsInner) GetSpecialPricing() []PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner {
	if o == nil || IsNil(o.SpecialPricing) {
		var ret []PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner
		return ret
	}
	return o.SpecialPricing
}

// GetSpecialPricingOk returns a tuple with the SpecialPricing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceAndAvailabilityResponseInnerDiscountsInner) GetSpecialPricingOk() ([]PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner, bool) {
	if o == nil || IsNil(o.SpecialPricing) {
		return nil, false
	}
	return o.SpecialPricing, true
}

// HasSpecialPricing returns a boolean if a field has been set.
func (o *PriceAndAvailabilityResponseInnerDiscountsInner) HasSpecialPricing() bool {
	if o != nil && !IsNil(o.SpecialPricing) {
		return true
	}

	return false
}

// SetSpecialPricing gets a reference to the given []PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner and assigns it to the SpecialPricing field.
func (o *PriceAndAvailabilityResponseInnerDiscountsInner) SetSpecialPricing(v []PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) {
	o.SpecialPricing = v
}

// GetQuantityDiscounts returns the QuantityDiscounts field value if set, zero value otherwise.
func (o *PriceAndAvailabilityResponseInnerDiscountsInner) GetQuantityDiscounts() []PriceAndAvailabilityResponseInnerDiscountsInnerQuantityDiscountsInner {
	if o == nil || IsNil(o.QuantityDiscounts) {
		var ret []PriceAndAvailabilityResponseInnerDiscountsInnerQuantityDiscountsInner
		return ret
	}
	return o.QuantityDiscounts
}

// GetQuantityDiscountsOk returns a tuple with the QuantityDiscounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceAndAvailabilityResponseInnerDiscountsInner) GetQuantityDiscountsOk() ([]PriceAndAvailabilityResponseInnerDiscountsInnerQuantityDiscountsInner, bool) {
	if o == nil || IsNil(o.QuantityDiscounts) {
		return nil, false
	}
	return o.QuantityDiscounts, true
}

// HasQuantityDiscounts returns a boolean if a field has been set.
func (o *PriceAndAvailabilityResponseInnerDiscountsInner) HasQuantityDiscounts() bool {
	if o != nil && !IsNil(o.QuantityDiscounts) {
		return true
	}

	return false
}

// SetQuantityDiscounts gets a reference to the given []PriceAndAvailabilityResponseInnerDiscountsInnerQuantityDiscountsInner and assigns it to the QuantityDiscounts field.
func (o *PriceAndAvailabilityResponseInnerDiscountsInner) SetQuantityDiscounts(v []PriceAndAvailabilityResponseInnerDiscountsInnerQuantityDiscountsInner) {
	o.QuantityDiscounts = v
}

func (o PriceAndAvailabilityResponseInnerDiscountsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PriceAndAvailabilityResponseInnerDiscountsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SpecialPricing) {
		toSerialize["specialPricing"] = o.SpecialPricing
	}
	if !IsNil(o.QuantityDiscounts) {
		toSerialize["quantityDiscounts"] = o.QuantityDiscounts
	}
	return toSerialize, nil
}

type NullablePriceAndAvailabilityResponseInnerDiscountsInner struct {
	value *PriceAndAvailabilityResponseInnerDiscountsInner
	isSet bool
}

func (v NullablePriceAndAvailabilityResponseInnerDiscountsInner) Get() *PriceAndAvailabilityResponseInnerDiscountsInner {
	return v.value
}

func (v *NullablePriceAndAvailabilityResponseInnerDiscountsInner) Set(val *PriceAndAvailabilityResponseInnerDiscountsInner) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceAndAvailabilityResponseInnerDiscountsInner) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceAndAvailabilityResponseInnerDiscountsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceAndAvailabilityResponseInnerDiscountsInner(val *PriceAndAvailabilityResponseInnerDiscountsInner) *NullablePriceAndAvailabilityResponseInnerDiscountsInner {
	return &NullablePriceAndAvailabilityResponseInnerDiscountsInner{value: val, isSet: true}
}

func (v NullablePriceAndAvailabilityResponseInnerDiscountsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceAndAvailabilityResponseInnerDiscountsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


