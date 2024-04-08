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

// checks if the AsyncOrderCreateDTOShipmentDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AsyncOrderCreateDTOShipmentDetails{}

// AsyncOrderCreateDTOShipmentDetails Shipping details for the order provided by the customer.
type AsyncOrderCreateDTOShipmentDetails struct {
	// The code for the shipping carrier for the line item.
	CarrierCode *string `json:"carrierCode,omitempty"`
	// The reseller-requested delivery date in UTC format. Delivery date is not guaranteed.
	RequestedDeliveryDate *string `json:"requestedDeliveryDate,omitempty"`
	// Specifies whether the shipment will be shipped only when all products are fulfilled. The value of this field along with acceptBackOrder field decides the value of backorderflag. If this field is set, acceptBackOrder field is ignored. Possible values for this field are true, C, P, E.
	ShipComplete *string `json:"shipComplete,omitempty"`
	// Any special shipping instructions for the order.
	ShippingInstructions *string `json:"shippingInstructions,omitempty"`
	// The reseller 's shipping account number with carrier. Used to bill the shipping carrier directly from the reseller's account with the carrier.
	FreightAccountNumber *string `json:"freightAccountNumber,omitempty"`
	// Specifies whether a signature is required for delivery. Default is False.
	SignatureRequired *bool `json:"signatureRequired,omitempty"`
}

// NewAsyncOrderCreateDTOShipmentDetails instantiates a new AsyncOrderCreateDTOShipmentDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAsyncOrderCreateDTOShipmentDetails() *AsyncOrderCreateDTOShipmentDetails {
	this := AsyncOrderCreateDTOShipmentDetails{}
	return &this
}

// NewAsyncOrderCreateDTOShipmentDetailsWithDefaults instantiates a new AsyncOrderCreateDTOShipmentDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAsyncOrderCreateDTOShipmentDetailsWithDefaults() *AsyncOrderCreateDTOShipmentDetails {
	this := AsyncOrderCreateDTOShipmentDetails{}
	return &this
}

// GetCarrierCode returns the CarrierCode field value if set, zero value otherwise.
func (o *AsyncOrderCreateDTOShipmentDetails) GetCarrierCode() string {
	if o == nil || IsNil(o.CarrierCode) {
		var ret string
		return ret
	}
	return *o.CarrierCode
}

// GetCarrierCodeOk returns a tuple with the CarrierCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsyncOrderCreateDTOShipmentDetails) GetCarrierCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CarrierCode) {
		return nil, false
	}
	return o.CarrierCode, true
}

// HasCarrierCode returns a boolean if a field has been set.
func (o *AsyncOrderCreateDTOShipmentDetails) HasCarrierCode() bool {
	if o != nil && !IsNil(o.CarrierCode) {
		return true
	}

	return false
}

// SetCarrierCode gets a reference to the given string and assigns it to the CarrierCode field.
func (o *AsyncOrderCreateDTOShipmentDetails) SetCarrierCode(v string) {
	o.CarrierCode = &v
}

// GetRequestedDeliveryDate returns the RequestedDeliveryDate field value if set, zero value otherwise.
func (o *AsyncOrderCreateDTOShipmentDetails) GetRequestedDeliveryDate() string {
	if o == nil || IsNil(o.RequestedDeliveryDate) {
		var ret string
		return ret
	}
	return *o.RequestedDeliveryDate
}

// GetRequestedDeliveryDateOk returns a tuple with the RequestedDeliveryDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsyncOrderCreateDTOShipmentDetails) GetRequestedDeliveryDateOk() (*string, bool) {
	if o == nil || IsNil(o.RequestedDeliveryDate) {
		return nil, false
	}
	return o.RequestedDeliveryDate, true
}

// HasRequestedDeliveryDate returns a boolean if a field has been set.
func (o *AsyncOrderCreateDTOShipmentDetails) HasRequestedDeliveryDate() bool {
	if o != nil && !IsNil(o.RequestedDeliveryDate) {
		return true
	}

	return false
}

// SetRequestedDeliveryDate gets a reference to the given string and assigns it to the RequestedDeliveryDate field.
func (o *AsyncOrderCreateDTOShipmentDetails) SetRequestedDeliveryDate(v string) {
	o.RequestedDeliveryDate = &v
}

// GetShipComplete returns the ShipComplete field value if set, zero value otherwise.
func (o *AsyncOrderCreateDTOShipmentDetails) GetShipComplete() string {
	if o == nil || IsNil(o.ShipComplete) {
		var ret string
		return ret
	}
	return *o.ShipComplete
}

// GetShipCompleteOk returns a tuple with the ShipComplete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsyncOrderCreateDTOShipmentDetails) GetShipCompleteOk() (*string, bool) {
	if o == nil || IsNil(o.ShipComplete) {
		return nil, false
	}
	return o.ShipComplete, true
}

// HasShipComplete returns a boolean if a field has been set.
func (o *AsyncOrderCreateDTOShipmentDetails) HasShipComplete() bool {
	if o != nil && !IsNil(o.ShipComplete) {
		return true
	}

	return false
}

// SetShipComplete gets a reference to the given string and assigns it to the ShipComplete field.
func (o *AsyncOrderCreateDTOShipmentDetails) SetShipComplete(v string) {
	o.ShipComplete = &v
}

// GetShippingInstructions returns the ShippingInstructions field value if set, zero value otherwise.
func (o *AsyncOrderCreateDTOShipmentDetails) GetShippingInstructions() string {
	if o == nil || IsNil(o.ShippingInstructions) {
		var ret string
		return ret
	}
	return *o.ShippingInstructions
}

// GetShippingInstructionsOk returns a tuple with the ShippingInstructions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsyncOrderCreateDTOShipmentDetails) GetShippingInstructionsOk() (*string, bool) {
	if o == nil || IsNil(o.ShippingInstructions) {
		return nil, false
	}
	return o.ShippingInstructions, true
}

// HasShippingInstructions returns a boolean if a field has been set.
func (o *AsyncOrderCreateDTOShipmentDetails) HasShippingInstructions() bool {
	if o != nil && !IsNil(o.ShippingInstructions) {
		return true
	}

	return false
}

// SetShippingInstructions gets a reference to the given string and assigns it to the ShippingInstructions field.
func (o *AsyncOrderCreateDTOShipmentDetails) SetShippingInstructions(v string) {
	o.ShippingInstructions = &v
}

// GetFreightAccountNumber returns the FreightAccountNumber field value if set, zero value otherwise.
func (o *AsyncOrderCreateDTOShipmentDetails) GetFreightAccountNumber() string {
	if o == nil || IsNil(o.FreightAccountNumber) {
		var ret string
		return ret
	}
	return *o.FreightAccountNumber
}

// GetFreightAccountNumberOk returns a tuple with the FreightAccountNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsyncOrderCreateDTOShipmentDetails) GetFreightAccountNumberOk() (*string, bool) {
	if o == nil || IsNil(o.FreightAccountNumber) {
		return nil, false
	}
	return o.FreightAccountNumber, true
}

// HasFreightAccountNumber returns a boolean if a field has been set.
func (o *AsyncOrderCreateDTOShipmentDetails) HasFreightAccountNumber() bool {
	if o != nil && !IsNil(o.FreightAccountNumber) {
		return true
	}

	return false
}

// SetFreightAccountNumber gets a reference to the given string and assigns it to the FreightAccountNumber field.
func (o *AsyncOrderCreateDTOShipmentDetails) SetFreightAccountNumber(v string) {
	o.FreightAccountNumber = &v
}

// GetSignatureRequired returns the SignatureRequired field value if set, zero value otherwise.
func (o *AsyncOrderCreateDTOShipmentDetails) GetSignatureRequired() bool {
	if o == nil || IsNil(o.SignatureRequired) {
		var ret bool
		return ret
	}
	return *o.SignatureRequired
}

// GetSignatureRequiredOk returns a tuple with the SignatureRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsyncOrderCreateDTOShipmentDetails) GetSignatureRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.SignatureRequired) {
		return nil, false
	}
	return o.SignatureRequired, true
}

// HasSignatureRequired returns a boolean if a field has been set.
func (o *AsyncOrderCreateDTOShipmentDetails) HasSignatureRequired() bool {
	if o != nil && !IsNil(o.SignatureRequired) {
		return true
	}

	return false
}

// SetSignatureRequired gets a reference to the given bool and assigns it to the SignatureRequired field.
func (o *AsyncOrderCreateDTOShipmentDetails) SetSignatureRequired(v bool) {
	o.SignatureRequired = &v
}

func (o AsyncOrderCreateDTOShipmentDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AsyncOrderCreateDTOShipmentDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CarrierCode) {
		toSerialize["carrierCode"] = o.CarrierCode
	}
	if !IsNil(o.RequestedDeliveryDate) {
		toSerialize["requestedDeliveryDate"] = o.RequestedDeliveryDate
	}
	if !IsNil(o.ShipComplete) {
		toSerialize["shipComplete"] = o.ShipComplete
	}
	if !IsNil(o.ShippingInstructions) {
		toSerialize["shippingInstructions"] = o.ShippingInstructions
	}
	if !IsNil(o.FreightAccountNumber) {
		toSerialize["freightAccountNumber"] = o.FreightAccountNumber
	}
	if !IsNil(o.SignatureRequired) {
		toSerialize["signatureRequired"] = o.SignatureRequired
	}
	return toSerialize, nil
}

type NullableAsyncOrderCreateDTOShipmentDetails struct {
	value *AsyncOrderCreateDTOShipmentDetails
	isSet bool
}

func (v NullableAsyncOrderCreateDTOShipmentDetails) Get() *AsyncOrderCreateDTOShipmentDetails {
	return v.value
}

func (v *NullableAsyncOrderCreateDTOShipmentDetails) Set(val *AsyncOrderCreateDTOShipmentDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableAsyncOrderCreateDTOShipmentDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableAsyncOrderCreateDTOShipmentDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAsyncOrderCreateDTOShipmentDetails(val *AsyncOrderCreateDTOShipmentDetails) *NullableAsyncOrderCreateDTOShipmentDetails {
	return &NullableAsyncOrderCreateDTOShipmentDetails{value: val, isSet: true}
}

func (v NullableAsyncOrderCreateDTOShipmentDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAsyncOrderCreateDTOShipmentDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

