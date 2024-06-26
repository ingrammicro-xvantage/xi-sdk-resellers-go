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

// checks if the RenewalsSearchRequestStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RenewalsSearchRequestStatus{}

// RenewalsSearchRequestStatus struct for RenewalsSearchRequestStatus
type RenewalsSearchRequestStatus struct {
	OpporutinyStatus *RenewalsSearchRequestStatusOpporutinyStatus `json:"OpporutinyStatus,omitempty"`
}

// NewRenewalsSearchRequestStatus instantiates a new RenewalsSearchRequestStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRenewalsSearchRequestStatus() *RenewalsSearchRequestStatus {
	this := RenewalsSearchRequestStatus{}
	return &this
}

// NewRenewalsSearchRequestStatusWithDefaults instantiates a new RenewalsSearchRequestStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRenewalsSearchRequestStatusWithDefaults() *RenewalsSearchRequestStatus {
	this := RenewalsSearchRequestStatus{}
	return &this
}

// GetOpporutinyStatus returns the OpporutinyStatus field value if set, zero value otherwise.
func (o *RenewalsSearchRequestStatus) GetOpporutinyStatus() RenewalsSearchRequestStatusOpporutinyStatus {
	if o == nil || IsNil(o.OpporutinyStatus) {
		var ret RenewalsSearchRequestStatusOpporutinyStatus
		return ret
	}
	return *o.OpporutinyStatus
}

// GetOpporutinyStatusOk returns a tuple with the OpporutinyStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenewalsSearchRequestStatus) GetOpporutinyStatusOk() (*RenewalsSearchRequestStatusOpporutinyStatus, bool) {
	if o == nil || IsNil(o.OpporutinyStatus) {
		return nil, false
	}
	return o.OpporutinyStatus, true
}

// HasOpporutinyStatus returns a boolean if a field has been set.
func (o *RenewalsSearchRequestStatus) HasOpporutinyStatus() bool {
	if o != nil && !IsNil(o.OpporutinyStatus) {
		return true
	}

	return false
}

// SetOpporutinyStatus gets a reference to the given RenewalsSearchRequestStatusOpporutinyStatus and assigns it to the OpporutinyStatus field.
func (o *RenewalsSearchRequestStatus) SetOpporutinyStatus(v RenewalsSearchRequestStatusOpporutinyStatus) {
	o.OpporutinyStatus = &v
}

func (o RenewalsSearchRequestStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RenewalsSearchRequestStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OpporutinyStatus) {
		toSerialize["OpporutinyStatus"] = o.OpporutinyStatus
	}
	return toSerialize, nil
}

type NullableRenewalsSearchRequestStatus struct {
	value *RenewalsSearchRequestStatus
	isSet bool
}

func (v NullableRenewalsSearchRequestStatus) Get() *RenewalsSearchRequestStatus {
	return v.value
}

func (v *NullableRenewalsSearchRequestStatus) Set(val *RenewalsSearchRequestStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableRenewalsSearchRequestStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableRenewalsSearchRequestStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRenewalsSearchRequestStatus(val *RenewalsSearchRequestStatus) *NullableRenewalsSearchRequestStatus {
	return &NullableRenewalsSearchRequestStatus{value: val, isSet: true}
}

func (v NullableRenewalsSearchRequestStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRenewalsSearchRequestStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


