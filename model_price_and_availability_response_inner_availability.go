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

// checks if the PriceAndAvailabilityResponseInnerAvailability type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PriceAndAvailabilityResponseInnerAvailability{}

// PriceAndAvailabilityResponseInnerAvailability struct for PriceAndAvailabilityResponseInnerAvailability
type PriceAndAvailabilityResponseInnerAvailability struct {
	// Boolean that indicates if the product ordered is available
	Available *bool `json:"available,omitempty"`
	// The total amount of available products
	TotalAvailability *int32 `json:"totalAvailability,omitempty"`
	AvailabilityByWarehouse []PriceAndAvailabilityResponseInnerAvailabilityAvailabilityByWarehouseInner `json:"availabilityByWarehouse,omitempty"`
}

// NewPriceAndAvailabilityResponseInnerAvailability instantiates a new PriceAndAvailabilityResponseInnerAvailability object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPriceAndAvailabilityResponseInnerAvailability() *PriceAndAvailabilityResponseInnerAvailability {
	this := PriceAndAvailabilityResponseInnerAvailability{}
	return &this
}

// NewPriceAndAvailabilityResponseInnerAvailabilityWithDefaults instantiates a new PriceAndAvailabilityResponseInnerAvailability object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceAndAvailabilityResponseInnerAvailabilityWithDefaults() *PriceAndAvailabilityResponseInnerAvailability {
	this := PriceAndAvailabilityResponseInnerAvailability{}
	return &this
}

// GetAvailable returns the Available field value if set, zero value otherwise.
func (o *PriceAndAvailabilityResponseInnerAvailability) GetAvailable() bool {
	if o == nil || IsNil(o.Available) {
		var ret bool
		return ret
	}
	return *o.Available
}

// GetAvailableOk returns a tuple with the Available field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceAndAvailabilityResponseInnerAvailability) GetAvailableOk() (*bool, bool) {
	if o == nil || IsNil(o.Available) {
		return nil, false
	}
	return o.Available, true
}

// HasAvailable returns a boolean if a field has been set.
func (o *PriceAndAvailabilityResponseInnerAvailability) HasAvailable() bool {
	if o != nil && !IsNil(o.Available) {
		return true
	}

	return false
}

// SetAvailable gets a reference to the given bool and assigns it to the Available field.
func (o *PriceAndAvailabilityResponseInnerAvailability) SetAvailable(v bool) {
	o.Available = &v
}

// GetTotalAvailability returns the TotalAvailability field value if set, zero value otherwise.
func (o *PriceAndAvailabilityResponseInnerAvailability) GetTotalAvailability() int32 {
	if o == nil || IsNil(o.TotalAvailability) {
		var ret int32
		return ret
	}
	return *o.TotalAvailability
}

// GetTotalAvailabilityOk returns a tuple with the TotalAvailability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceAndAvailabilityResponseInnerAvailability) GetTotalAvailabilityOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalAvailability) {
		return nil, false
	}
	return o.TotalAvailability, true
}

// HasTotalAvailability returns a boolean if a field has been set.
func (o *PriceAndAvailabilityResponseInnerAvailability) HasTotalAvailability() bool {
	if o != nil && !IsNil(o.TotalAvailability) {
		return true
	}

	return false
}

// SetTotalAvailability gets a reference to the given int32 and assigns it to the TotalAvailability field.
func (o *PriceAndAvailabilityResponseInnerAvailability) SetTotalAvailability(v int32) {
	o.TotalAvailability = &v
}

// GetAvailabilityByWarehouse returns the AvailabilityByWarehouse field value if set, zero value otherwise.
func (o *PriceAndAvailabilityResponseInnerAvailability) GetAvailabilityByWarehouse() []PriceAndAvailabilityResponseInnerAvailabilityAvailabilityByWarehouseInner {
	if o == nil || IsNil(o.AvailabilityByWarehouse) {
		var ret []PriceAndAvailabilityResponseInnerAvailabilityAvailabilityByWarehouseInner
		return ret
	}
	return o.AvailabilityByWarehouse
}

// GetAvailabilityByWarehouseOk returns a tuple with the AvailabilityByWarehouse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceAndAvailabilityResponseInnerAvailability) GetAvailabilityByWarehouseOk() ([]PriceAndAvailabilityResponseInnerAvailabilityAvailabilityByWarehouseInner, bool) {
	if o == nil || IsNil(o.AvailabilityByWarehouse) {
		return nil, false
	}
	return o.AvailabilityByWarehouse, true
}

// HasAvailabilityByWarehouse returns a boolean if a field has been set.
func (o *PriceAndAvailabilityResponseInnerAvailability) HasAvailabilityByWarehouse() bool {
	if o != nil && !IsNil(o.AvailabilityByWarehouse) {
		return true
	}

	return false
}

// SetAvailabilityByWarehouse gets a reference to the given []PriceAndAvailabilityResponseInnerAvailabilityAvailabilityByWarehouseInner and assigns it to the AvailabilityByWarehouse field.
func (o *PriceAndAvailabilityResponseInnerAvailability) SetAvailabilityByWarehouse(v []PriceAndAvailabilityResponseInnerAvailabilityAvailabilityByWarehouseInner) {
	o.AvailabilityByWarehouse = v
}

func (o PriceAndAvailabilityResponseInnerAvailability) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PriceAndAvailabilityResponseInnerAvailability) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Available) {
		toSerialize["available"] = o.Available
	}
	if !IsNil(o.TotalAvailability) {
		toSerialize["totalAvailability"] = o.TotalAvailability
	}
	if !IsNil(o.AvailabilityByWarehouse) {
		toSerialize["availabilityByWarehouse"] = o.AvailabilityByWarehouse
	}
	return toSerialize, nil
}

type NullablePriceAndAvailabilityResponseInnerAvailability struct {
	value *PriceAndAvailabilityResponseInnerAvailability
	isSet bool
}

func (v NullablePriceAndAvailabilityResponseInnerAvailability) Get() *PriceAndAvailabilityResponseInnerAvailability {
	return v.value
}

func (v *NullablePriceAndAvailabilityResponseInnerAvailability) Set(val *PriceAndAvailabilityResponseInnerAvailability) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceAndAvailabilityResponseInnerAvailability) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceAndAvailabilityResponseInnerAvailability) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceAndAvailabilityResponseInnerAvailability(val *PriceAndAvailabilityResponseInnerAvailability) *NullablePriceAndAvailabilityResponseInnerAvailability {
	return &NullablePriceAndAvailabilityResponseInnerAvailability{value: val, isSet: true}
}

func (v NullablePriceAndAvailabilityResponseInnerAvailability) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceAndAvailabilityResponseInnerAvailability) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


