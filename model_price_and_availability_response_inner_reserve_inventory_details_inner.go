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

// checks if the PriceAndAvailabilityResponseInnerReserveInventoryDetailsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PriceAndAvailabilityResponseInnerReserveInventoryDetailsInner{}

// PriceAndAvailabilityResponseInnerReserveInventoryDetailsInner struct for PriceAndAvailabilityResponseInnerReserveInventoryDetailsInner
type PriceAndAvailabilityResponseInnerReserveInventoryDetailsInner struct {
	// The quantity of the product reserved for the customer.
	QuantityReserved *int32 `json:"quantityReserved,omitempty"`
	// The availability of the product reserved.
	QuantityAvailable *int32 `json:"quantityAvailable,omitempty"`
	// The reservation date for the product in UTC format.
	Effectivedate *string `json:"effectivedate,omitempty"`
	// The expiration date for the reservation of the product in UTC format.
	Expirydate *string `json:"expirydate,omitempty"`
}

// NewPriceAndAvailabilityResponseInnerReserveInventoryDetailsInner instantiates a new PriceAndAvailabilityResponseInnerReserveInventoryDetailsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPriceAndAvailabilityResponseInnerReserveInventoryDetailsInner() *PriceAndAvailabilityResponseInnerReserveInventoryDetailsInner {
	this := PriceAndAvailabilityResponseInnerReserveInventoryDetailsInner{}
	return &this
}

// NewPriceAndAvailabilityResponseInnerReserveInventoryDetailsInnerWithDefaults instantiates a new PriceAndAvailabilityResponseInnerReserveInventoryDetailsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceAndAvailabilityResponseInnerReserveInventoryDetailsInnerWithDefaults() *PriceAndAvailabilityResponseInnerReserveInventoryDetailsInner {
	this := PriceAndAvailabilityResponseInnerReserveInventoryDetailsInner{}
	return &this
}

// GetQuantityReserved returns the QuantityReserved field value if set, zero value otherwise.
func (o *PriceAndAvailabilityResponseInnerReserveInventoryDetailsInner) GetQuantityReserved() int32 {
	if o == nil || IsNil(o.QuantityReserved) {
		var ret int32
		return ret
	}
	return *o.QuantityReserved
}

// GetQuantityReservedOk returns a tuple with the QuantityReserved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceAndAvailabilityResponseInnerReserveInventoryDetailsInner) GetQuantityReservedOk() (*int32, bool) {
	if o == nil || IsNil(o.QuantityReserved) {
		return nil, false
	}
	return o.QuantityReserved, true
}

// HasQuantityReserved returns a boolean if a field has been set.
func (o *PriceAndAvailabilityResponseInnerReserveInventoryDetailsInner) HasQuantityReserved() bool {
	if o != nil && !IsNil(o.QuantityReserved) {
		return true
	}

	return false
}

// SetQuantityReserved gets a reference to the given int32 and assigns it to the QuantityReserved field.
func (o *PriceAndAvailabilityResponseInnerReserveInventoryDetailsInner) SetQuantityReserved(v int32) {
	o.QuantityReserved = &v
}

// GetQuantityAvailable returns the QuantityAvailable field value if set, zero value otherwise.
func (o *PriceAndAvailabilityResponseInnerReserveInventoryDetailsInner) GetQuantityAvailable() int32 {
	if o == nil || IsNil(o.QuantityAvailable) {
		var ret int32
		return ret
	}
	return *o.QuantityAvailable
}

// GetQuantityAvailableOk returns a tuple with the QuantityAvailable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceAndAvailabilityResponseInnerReserveInventoryDetailsInner) GetQuantityAvailableOk() (*int32, bool) {
	if o == nil || IsNil(o.QuantityAvailable) {
		return nil, false
	}
	return o.QuantityAvailable, true
}

// HasQuantityAvailable returns a boolean if a field has been set.
func (o *PriceAndAvailabilityResponseInnerReserveInventoryDetailsInner) HasQuantityAvailable() bool {
	if o != nil && !IsNil(o.QuantityAvailable) {
		return true
	}

	return false
}

// SetQuantityAvailable gets a reference to the given int32 and assigns it to the QuantityAvailable field.
func (o *PriceAndAvailabilityResponseInnerReserveInventoryDetailsInner) SetQuantityAvailable(v int32) {
	o.QuantityAvailable = &v
}

// GetEffectivedate returns the Effectivedate field value if set, zero value otherwise.
func (o *PriceAndAvailabilityResponseInnerReserveInventoryDetailsInner) GetEffectivedate() string {
	if o == nil || IsNil(o.Effectivedate) {
		var ret string
		return ret
	}
	return *o.Effectivedate
}

// GetEffectivedateOk returns a tuple with the Effectivedate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceAndAvailabilityResponseInnerReserveInventoryDetailsInner) GetEffectivedateOk() (*string, bool) {
	if o == nil || IsNil(o.Effectivedate) {
		return nil, false
	}
	return o.Effectivedate, true
}

// HasEffectivedate returns a boolean if a field has been set.
func (o *PriceAndAvailabilityResponseInnerReserveInventoryDetailsInner) HasEffectivedate() bool {
	if o != nil && !IsNil(o.Effectivedate) {
		return true
	}

	return false
}

// SetEffectivedate gets a reference to the given string and assigns it to the Effectivedate field.
func (o *PriceAndAvailabilityResponseInnerReserveInventoryDetailsInner) SetEffectivedate(v string) {
	o.Effectivedate = &v
}

// GetExpirydate returns the Expirydate field value if set, zero value otherwise.
func (o *PriceAndAvailabilityResponseInnerReserveInventoryDetailsInner) GetExpirydate() string {
	if o == nil || IsNil(o.Expirydate) {
		var ret string
		return ret
	}
	return *o.Expirydate
}

// GetExpirydateOk returns a tuple with the Expirydate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceAndAvailabilityResponseInnerReserveInventoryDetailsInner) GetExpirydateOk() (*string, bool) {
	if o == nil || IsNil(o.Expirydate) {
		return nil, false
	}
	return o.Expirydate, true
}

// HasExpirydate returns a boolean if a field has been set.
func (o *PriceAndAvailabilityResponseInnerReserveInventoryDetailsInner) HasExpirydate() bool {
	if o != nil && !IsNil(o.Expirydate) {
		return true
	}

	return false
}

// SetExpirydate gets a reference to the given string and assigns it to the Expirydate field.
func (o *PriceAndAvailabilityResponseInnerReserveInventoryDetailsInner) SetExpirydate(v string) {
	o.Expirydate = &v
}

func (o PriceAndAvailabilityResponseInnerReserveInventoryDetailsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PriceAndAvailabilityResponseInnerReserveInventoryDetailsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.QuantityReserved) {
		toSerialize["quantityReserved"] = o.QuantityReserved
	}
	if !IsNil(o.QuantityAvailable) {
		toSerialize["quantityAvailable"] = o.QuantityAvailable
	}
	if !IsNil(o.Effectivedate) {
		toSerialize["effectivedate"] = o.Effectivedate
	}
	if !IsNil(o.Expirydate) {
		toSerialize["expirydate"] = o.Expirydate
	}
	return toSerialize, nil
}

type NullablePriceAndAvailabilityResponseInnerReserveInventoryDetailsInner struct {
	value *PriceAndAvailabilityResponseInnerReserveInventoryDetailsInner
	isSet bool
}

func (v NullablePriceAndAvailabilityResponseInnerReserveInventoryDetailsInner) Get() *PriceAndAvailabilityResponseInnerReserveInventoryDetailsInner {
	return v.value
}

func (v *NullablePriceAndAvailabilityResponseInnerReserveInventoryDetailsInner) Set(val *PriceAndAvailabilityResponseInnerReserveInventoryDetailsInner) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceAndAvailabilityResponseInnerReserveInventoryDetailsInner) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceAndAvailabilityResponseInnerReserveInventoryDetailsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceAndAvailabilityResponseInnerReserveInventoryDetailsInner(val *PriceAndAvailabilityResponseInnerReserveInventoryDetailsInner) *NullablePriceAndAvailabilityResponseInnerReserveInventoryDetailsInner {
	return &NullablePriceAndAvailabilityResponseInnerReserveInventoryDetailsInner{value: val, isSet: true}
}

func (v NullablePriceAndAvailabilityResponseInnerReserveInventoryDetailsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceAndAvailabilityResponseInnerReserveInventoryDetailsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


