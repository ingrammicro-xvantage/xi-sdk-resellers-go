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

// checks if the FreightResponseFreightEstimateResponseDistributionInnerCarrierListInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FreightResponseFreightEstimateResponseDistributionInnerCarrierListInner{}

// FreightResponseFreightEstimateResponseDistributionInnerCarrierListInner struct for FreightResponseFreightEstimateResponseDistributionInnerCarrierListInner
type FreightResponseFreightEstimateResponseDistributionInnerCarrierListInner struct {
	// The code for the shipping carrier for the line item.
	CarrierCode *string `json:"carrierCode,omitempty"`
	// The name of the shipping carrier.
	ShipVia *string `json:"shipVia,omitempty"`
	// Mode of the carrier.
	CarrierMode *string `json:"carrierMode,omitempty"`
	// Estimated freight charge.
	EstimatedFreightCharge *string `json:"estimatedFreightCharge,omitempty"`
	// Number of transit days.
	DaysInTransit *string `json:"daysInTransit,omitempty"`
}

// NewFreightResponseFreightEstimateResponseDistributionInnerCarrierListInner instantiates a new FreightResponseFreightEstimateResponseDistributionInnerCarrierListInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFreightResponseFreightEstimateResponseDistributionInnerCarrierListInner() *FreightResponseFreightEstimateResponseDistributionInnerCarrierListInner {
	this := FreightResponseFreightEstimateResponseDistributionInnerCarrierListInner{}
	return &this
}

// NewFreightResponseFreightEstimateResponseDistributionInnerCarrierListInnerWithDefaults instantiates a new FreightResponseFreightEstimateResponseDistributionInnerCarrierListInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFreightResponseFreightEstimateResponseDistributionInnerCarrierListInnerWithDefaults() *FreightResponseFreightEstimateResponseDistributionInnerCarrierListInner {
	this := FreightResponseFreightEstimateResponseDistributionInnerCarrierListInner{}
	return &this
}

// GetCarrierCode returns the CarrierCode field value if set, zero value otherwise.
func (o *FreightResponseFreightEstimateResponseDistributionInnerCarrierListInner) GetCarrierCode() string {
	if o == nil || IsNil(o.CarrierCode) {
		var ret string
		return ret
	}
	return *o.CarrierCode
}

// GetCarrierCodeOk returns a tuple with the CarrierCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FreightResponseFreightEstimateResponseDistributionInnerCarrierListInner) GetCarrierCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CarrierCode) {
		return nil, false
	}
	return o.CarrierCode, true
}

// HasCarrierCode returns a boolean if a field has been set.
func (o *FreightResponseFreightEstimateResponseDistributionInnerCarrierListInner) HasCarrierCode() bool {
	if o != nil && !IsNil(o.CarrierCode) {
		return true
	}

	return false
}

// SetCarrierCode gets a reference to the given string and assigns it to the CarrierCode field.
func (o *FreightResponseFreightEstimateResponseDistributionInnerCarrierListInner) SetCarrierCode(v string) {
	o.CarrierCode = &v
}

// GetShipVia returns the ShipVia field value if set, zero value otherwise.
func (o *FreightResponseFreightEstimateResponseDistributionInnerCarrierListInner) GetShipVia() string {
	if o == nil || IsNil(o.ShipVia) {
		var ret string
		return ret
	}
	return *o.ShipVia
}

// GetShipViaOk returns a tuple with the ShipVia field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FreightResponseFreightEstimateResponseDistributionInnerCarrierListInner) GetShipViaOk() (*string, bool) {
	if o == nil || IsNil(o.ShipVia) {
		return nil, false
	}
	return o.ShipVia, true
}

// HasShipVia returns a boolean if a field has been set.
func (o *FreightResponseFreightEstimateResponseDistributionInnerCarrierListInner) HasShipVia() bool {
	if o != nil && !IsNil(o.ShipVia) {
		return true
	}

	return false
}

// SetShipVia gets a reference to the given string and assigns it to the ShipVia field.
func (o *FreightResponseFreightEstimateResponseDistributionInnerCarrierListInner) SetShipVia(v string) {
	o.ShipVia = &v
}

// GetCarrierMode returns the CarrierMode field value if set, zero value otherwise.
func (o *FreightResponseFreightEstimateResponseDistributionInnerCarrierListInner) GetCarrierMode() string {
	if o == nil || IsNil(o.CarrierMode) {
		var ret string
		return ret
	}
	return *o.CarrierMode
}

// GetCarrierModeOk returns a tuple with the CarrierMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FreightResponseFreightEstimateResponseDistributionInnerCarrierListInner) GetCarrierModeOk() (*string, bool) {
	if o == nil || IsNil(o.CarrierMode) {
		return nil, false
	}
	return o.CarrierMode, true
}

// HasCarrierMode returns a boolean if a field has been set.
func (o *FreightResponseFreightEstimateResponseDistributionInnerCarrierListInner) HasCarrierMode() bool {
	if o != nil && !IsNil(o.CarrierMode) {
		return true
	}

	return false
}

// SetCarrierMode gets a reference to the given string and assigns it to the CarrierMode field.
func (o *FreightResponseFreightEstimateResponseDistributionInnerCarrierListInner) SetCarrierMode(v string) {
	o.CarrierMode = &v
}

// GetEstimatedFreightCharge returns the EstimatedFreightCharge field value if set, zero value otherwise.
func (o *FreightResponseFreightEstimateResponseDistributionInnerCarrierListInner) GetEstimatedFreightCharge() string {
	if o == nil || IsNil(o.EstimatedFreightCharge) {
		var ret string
		return ret
	}
	return *o.EstimatedFreightCharge
}

// GetEstimatedFreightChargeOk returns a tuple with the EstimatedFreightCharge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FreightResponseFreightEstimateResponseDistributionInnerCarrierListInner) GetEstimatedFreightChargeOk() (*string, bool) {
	if o == nil || IsNil(o.EstimatedFreightCharge) {
		return nil, false
	}
	return o.EstimatedFreightCharge, true
}

// HasEstimatedFreightCharge returns a boolean if a field has been set.
func (o *FreightResponseFreightEstimateResponseDistributionInnerCarrierListInner) HasEstimatedFreightCharge() bool {
	if o != nil && !IsNil(o.EstimatedFreightCharge) {
		return true
	}

	return false
}

// SetEstimatedFreightCharge gets a reference to the given string and assigns it to the EstimatedFreightCharge field.
func (o *FreightResponseFreightEstimateResponseDistributionInnerCarrierListInner) SetEstimatedFreightCharge(v string) {
	o.EstimatedFreightCharge = &v
}

// GetDaysInTransit returns the DaysInTransit field value if set, zero value otherwise.
func (o *FreightResponseFreightEstimateResponseDistributionInnerCarrierListInner) GetDaysInTransit() string {
	if o == nil || IsNil(o.DaysInTransit) {
		var ret string
		return ret
	}
	return *o.DaysInTransit
}

// GetDaysInTransitOk returns a tuple with the DaysInTransit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FreightResponseFreightEstimateResponseDistributionInnerCarrierListInner) GetDaysInTransitOk() (*string, bool) {
	if o == nil || IsNil(o.DaysInTransit) {
		return nil, false
	}
	return o.DaysInTransit, true
}

// HasDaysInTransit returns a boolean if a field has been set.
func (o *FreightResponseFreightEstimateResponseDistributionInnerCarrierListInner) HasDaysInTransit() bool {
	if o != nil && !IsNil(o.DaysInTransit) {
		return true
	}

	return false
}

// SetDaysInTransit gets a reference to the given string and assigns it to the DaysInTransit field.
func (o *FreightResponseFreightEstimateResponseDistributionInnerCarrierListInner) SetDaysInTransit(v string) {
	o.DaysInTransit = &v
}

func (o FreightResponseFreightEstimateResponseDistributionInnerCarrierListInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FreightResponseFreightEstimateResponseDistributionInnerCarrierListInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CarrierCode) {
		toSerialize["carrierCode"] = o.CarrierCode
	}
	if !IsNil(o.ShipVia) {
		toSerialize["shipVia"] = o.ShipVia
	}
	if !IsNil(o.CarrierMode) {
		toSerialize["carrierMode"] = o.CarrierMode
	}
	if !IsNil(o.EstimatedFreightCharge) {
		toSerialize["estimatedFreightCharge"] = o.EstimatedFreightCharge
	}
	if !IsNil(o.DaysInTransit) {
		toSerialize["daysInTransit"] = o.DaysInTransit
	}
	return toSerialize, nil
}

type NullableFreightResponseFreightEstimateResponseDistributionInnerCarrierListInner struct {
	value *FreightResponseFreightEstimateResponseDistributionInnerCarrierListInner
	isSet bool
}

func (v NullableFreightResponseFreightEstimateResponseDistributionInnerCarrierListInner) Get() *FreightResponseFreightEstimateResponseDistributionInnerCarrierListInner {
	return v.value
}

func (v *NullableFreightResponseFreightEstimateResponseDistributionInnerCarrierListInner) Set(val *FreightResponseFreightEstimateResponseDistributionInnerCarrierListInner) {
	v.value = val
	v.isSet = true
}

func (v NullableFreightResponseFreightEstimateResponseDistributionInnerCarrierListInner) IsSet() bool {
	return v.isSet
}

func (v *NullableFreightResponseFreightEstimateResponseDistributionInnerCarrierListInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFreightResponseFreightEstimateResponseDistributionInnerCarrierListInner(val *FreightResponseFreightEstimateResponseDistributionInnerCarrierListInner) *NullableFreightResponseFreightEstimateResponseDistributionInnerCarrierListInner {
	return &NullableFreightResponseFreightEstimateResponseDistributionInnerCarrierListInner{value: val, isSet: true}
}

func (v NullableFreightResponseFreightEstimateResponseDistributionInnerCarrierListInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFreightResponseFreightEstimateResponseDistributionInnerCarrierListInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


