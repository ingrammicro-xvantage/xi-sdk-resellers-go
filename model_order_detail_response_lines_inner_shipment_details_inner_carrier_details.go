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

// checks if the OrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetails{}

// OrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetails The shipment carrier details for the line item.
type OrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetails struct {
	// The carrier code for the shipment containing the line item.
	CarrierCode *string `json:"carrierCode,omitempty"`
	// The name of the carrier of the shipment containing the line item.
	CarrierName *string `json:"carrierName,omitempty"`
	TrackingDetails []OrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetailsTrackingDetailsInner `json:"trackingDetails,omitempty"`
}

// NewOrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetails instantiates a new OrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetails() *OrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetails {
	this := OrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetails{}
	return &this
}

// NewOrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetailsWithDefaults instantiates a new OrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetailsWithDefaults() *OrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetails {
	this := OrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetails{}
	return &this
}

// GetCarrierCode returns the CarrierCode field value if set, zero value otherwise.
func (o *OrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetails) GetCarrierCode() string {
	if o == nil || IsNil(o.CarrierCode) {
		var ret string
		return ret
	}
	return *o.CarrierCode
}

// GetCarrierCodeOk returns a tuple with the CarrierCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetails) GetCarrierCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CarrierCode) {
		return nil, false
	}
	return o.CarrierCode, true
}

// HasCarrierCode returns a boolean if a field has been set.
func (o *OrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetails) HasCarrierCode() bool {
	if o != nil && !IsNil(o.CarrierCode) {
		return true
	}

	return false
}

// SetCarrierCode gets a reference to the given string and assigns it to the CarrierCode field.
func (o *OrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetails) SetCarrierCode(v string) {
	o.CarrierCode = &v
}

// GetCarrierName returns the CarrierName field value if set, zero value otherwise.
func (o *OrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetails) GetCarrierName() string {
	if o == nil || IsNil(o.CarrierName) {
		var ret string
		return ret
	}
	return *o.CarrierName
}

// GetCarrierNameOk returns a tuple with the CarrierName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetails) GetCarrierNameOk() (*string, bool) {
	if o == nil || IsNil(o.CarrierName) {
		return nil, false
	}
	return o.CarrierName, true
}

// HasCarrierName returns a boolean if a field has been set.
func (o *OrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetails) HasCarrierName() bool {
	if o != nil && !IsNil(o.CarrierName) {
		return true
	}

	return false
}

// SetCarrierName gets a reference to the given string and assigns it to the CarrierName field.
func (o *OrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetails) SetCarrierName(v string) {
	o.CarrierName = &v
}

// GetTrackingDetails returns the TrackingDetails field value if set, zero value otherwise.
func (o *OrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetails) GetTrackingDetails() []OrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetailsTrackingDetailsInner {
	if o == nil || IsNil(o.TrackingDetails) {
		var ret []OrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetailsTrackingDetailsInner
		return ret
	}
	return o.TrackingDetails
}

// GetTrackingDetailsOk returns a tuple with the TrackingDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetails) GetTrackingDetailsOk() ([]OrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetailsTrackingDetailsInner, bool) {
	if o == nil || IsNil(o.TrackingDetails) {
		return nil, false
	}
	return o.TrackingDetails, true
}

// HasTrackingDetails returns a boolean if a field has been set.
func (o *OrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetails) HasTrackingDetails() bool {
	if o != nil && !IsNil(o.TrackingDetails) {
		return true
	}

	return false
}

// SetTrackingDetails gets a reference to the given []OrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetailsTrackingDetailsInner and assigns it to the TrackingDetails field.
func (o *OrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetails) SetTrackingDetails(v []OrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetailsTrackingDetailsInner) {
	o.TrackingDetails = v
}

func (o OrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CarrierCode) {
		toSerialize["carrierCode"] = o.CarrierCode
	}
	if !IsNil(o.CarrierName) {
		toSerialize["carrierName"] = o.CarrierName
	}
	if !IsNil(o.TrackingDetails) {
		toSerialize["trackingDetails"] = o.TrackingDetails
	}
	return toSerialize, nil
}

type NullableOrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetails struct {
	value *OrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetails
	isSet bool
}

func (v NullableOrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetails) Get() *OrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetails {
	return v.value
}

func (v *NullableOrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetails) Set(val *OrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetails(val *OrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetails) *NullableOrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetails {
	return &NullableOrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetails{value: val, isSet: true}
}

func (v NullableOrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


