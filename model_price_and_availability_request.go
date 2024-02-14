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

// checks if the PriceAndAvailabilityRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PriceAndAvailabilityRequest{}

// PriceAndAvailabilityRequest Request object model for the multi sku price and stock API endpoint
type PriceAndAvailabilityRequest struct {
	Servicerequest *PriceAndAvailabilityRequestServicerequest `json:"servicerequest,omitempty"`
}

// NewPriceAndAvailabilityRequest instantiates a new PriceAndAvailabilityRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPriceAndAvailabilityRequest() *PriceAndAvailabilityRequest {
	this := PriceAndAvailabilityRequest{}
	return &this
}

// NewPriceAndAvailabilityRequestWithDefaults instantiates a new PriceAndAvailabilityRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceAndAvailabilityRequestWithDefaults() *PriceAndAvailabilityRequest {
	this := PriceAndAvailabilityRequest{}
	return &this
}

// GetServicerequest returns the Servicerequest field value if set, zero value otherwise.
func (o *PriceAndAvailabilityRequest) GetServicerequest() PriceAndAvailabilityRequestServicerequest {
	if o == nil || IsNil(o.Servicerequest) {
		var ret PriceAndAvailabilityRequestServicerequest
		return ret
	}
	return *o.Servicerequest
}

// GetServicerequestOk returns a tuple with the Servicerequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceAndAvailabilityRequest) GetServicerequestOk() (*PriceAndAvailabilityRequestServicerequest, bool) {
	if o == nil || IsNil(o.Servicerequest) {
		return nil, false
	}
	return o.Servicerequest, true
}

// HasServicerequest returns a boolean if a field has been set.
func (o *PriceAndAvailabilityRequest) HasServicerequest() bool {
	if o != nil && !IsNil(o.Servicerequest) {
		return true
	}

	return false
}

// SetServicerequest gets a reference to the given PriceAndAvailabilityRequestServicerequest and assigns it to the Servicerequest field.
func (o *PriceAndAvailabilityRequest) SetServicerequest(v PriceAndAvailabilityRequestServicerequest) {
	o.Servicerequest = &v
}

func (o PriceAndAvailabilityRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PriceAndAvailabilityRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Servicerequest) {
		toSerialize["servicerequest"] = o.Servicerequest
	}
	return toSerialize, nil
}

type NullablePriceAndAvailabilityRequest struct {
	value *PriceAndAvailabilityRequest
	isSet bool
}

func (v NullablePriceAndAvailabilityRequest) Get() *PriceAndAvailabilityRequest {
	return v.value
}

func (v *NullablePriceAndAvailabilityRequest) Set(val *PriceAndAvailabilityRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceAndAvailabilityRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceAndAvailabilityRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceAndAvailabilityRequest(val *PriceAndAvailabilityRequest) *NullablePriceAndAvailabilityRequest {
	return &NullablePriceAndAvailabilityRequest{value: val, isSet: true}
}

func (v NullablePriceAndAvailabilityRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceAndAvailabilityRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


