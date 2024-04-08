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

// checks if the AsyncOrderCreateDTOWarrantyInfoInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AsyncOrderCreateDTOWarrantyInfoInner{}

// AsyncOrderCreateDTOWarrantyInfoInner struct for AsyncOrderCreateDTOWarrantyInfoInner
type AsyncOrderCreateDTOWarrantyInfoInner struct {
	HardwareLineLink *string `json:"hardwareLineLink,omitempty"`
	WarrantyLineLink *string `json:"warrantyLineLink,omitempty"`
	DirectLineLink *string `json:"directLineLink,omitempty"`
	SerialInfo []AsyncOrderCreateDTOWarrantyInfoInnerSerialInfoInner `json:"serialInfo,omitempty"`
	// The object containing the list of fields required at a line level by the vendor.
	VmfAdditionalAttributesLines []AsyncOrderCreateDTOWarrantyInfoInnerVmfAdditionalAttributesLinesInner `json:"vmfAdditionalAttributesLines,omitempty"`
}

// NewAsyncOrderCreateDTOWarrantyInfoInner instantiates a new AsyncOrderCreateDTOWarrantyInfoInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAsyncOrderCreateDTOWarrantyInfoInner() *AsyncOrderCreateDTOWarrantyInfoInner {
	this := AsyncOrderCreateDTOWarrantyInfoInner{}
	return &this
}

// NewAsyncOrderCreateDTOWarrantyInfoInnerWithDefaults instantiates a new AsyncOrderCreateDTOWarrantyInfoInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAsyncOrderCreateDTOWarrantyInfoInnerWithDefaults() *AsyncOrderCreateDTOWarrantyInfoInner {
	this := AsyncOrderCreateDTOWarrantyInfoInner{}
	return &this
}

// GetHardwareLineLink returns the HardwareLineLink field value if set, zero value otherwise.
func (o *AsyncOrderCreateDTOWarrantyInfoInner) GetHardwareLineLink() string {
	if o == nil || IsNil(o.HardwareLineLink) {
		var ret string
		return ret
	}
	return *o.HardwareLineLink
}

// GetHardwareLineLinkOk returns a tuple with the HardwareLineLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsyncOrderCreateDTOWarrantyInfoInner) GetHardwareLineLinkOk() (*string, bool) {
	if o == nil || IsNil(o.HardwareLineLink) {
		return nil, false
	}
	return o.HardwareLineLink, true
}

// HasHardwareLineLink returns a boolean if a field has been set.
func (o *AsyncOrderCreateDTOWarrantyInfoInner) HasHardwareLineLink() bool {
	if o != nil && !IsNil(o.HardwareLineLink) {
		return true
	}

	return false
}

// SetHardwareLineLink gets a reference to the given string and assigns it to the HardwareLineLink field.
func (o *AsyncOrderCreateDTOWarrantyInfoInner) SetHardwareLineLink(v string) {
	o.HardwareLineLink = &v
}

// GetWarrantyLineLink returns the WarrantyLineLink field value if set, zero value otherwise.
func (o *AsyncOrderCreateDTOWarrantyInfoInner) GetWarrantyLineLink() string {
	if o == nil || IsNil(o.WarrantyLineLink) {
		var ret string
		return ret
	}
	return *o.WarrantyLineLink
}

// GetWarrantyLineLinkOk returns a tuple with the WarrantyLineLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsyncOrderCreateDTOWarrantyInfoInner) GetWarrantyLineLinkOk() (*string, bool) {
	if o == nil || IsNil(o.WarrantyLineLink) {
		return nil, false
	}
	return o.WarrantyLineLink, true
}

// HasWarrantyLineLink returns a boolean if a field has been set.
func (o *AsyncOrderCreateDTOWarrantyInfoInner) HasWarrantyLineLink() bool {
	if o != nil && !IsNil(o.WarrantyLineLink) {
		return true
	}

	return false
}

// SetWarrantyLineLink gets a reference to the given string and assigns it to the WarrantyLineLink field.
func (o *AsyncOrderCreateDTOWarrantyInfoInner) SetWarrantyLineLink(v string) {
	o.WarrantyLineLink = &v
}

// GetDirectLineLink returns the DirectLineLink field value if set, zero value otherwise.
func (o *AsyncOrderCreateDTOWarrantyInfoInner) GetDirectLineLink() string {
	if o == nil || IsNil(o.DirectLineLink) {
		var ret string
		return ret
	}
	return *o.DirectLineLink
}

// GetDirectLineLinkOk returns a tuple with the DirectLineLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsyncOrderCreateDTOWarrantyInfoInner) GetDirectLineLinkOk() (*string, bool) {
	if o == nil || IsNil(o.DirectLineLink) {
		return nil, false
	}
	return o.DirectLineLink, true
}

// HasDirectLineLink returns a boolean if a field has been set.
func (o *AsyncOrderCreateDTOWarrantyInfoInner) HasDirectLineLink() bool {
	if o != nil && !IsNil(o.DirectLineLink) {
		return true
	}

	return false
}

// SetDirectLineLink gets a reference to the given string and assigns it to the DirectLineLink field.
func (o *AsyncOrderCreateDTOWarrantyInfoInner) SetDirectLineLink(v string) {
	o.DirectLineLink = &v
}

// GetSerialInfo returns the SerialInfo field value if set, zero value otherwise.
func (o *AsyncOrderCreateDTOWarrantyInfoInner) GetSerialInfo() []AsyncOrderCreateDTOWarrantyInfoInnerSerialInfoInner {
	if o == nil || IsNil(o.SerialInfo) {
		var ret []AsyncOrderCreateDTOWarrantyInfoInnerSerialInfoInner
		return ret
	}
	return o.SerialInfo
}

// GetSerialInfoOk returns a tuple with the SerialInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsyncOrderCreateDTOWarrantyInfoInner) GetSerialInfoOk() ([]AsyncOrderCreateDTOWarrantyInfoInnerSerialInfoInner, bool) {
	if o == nil || IsNil(o.SerialInfo) {
		return nil, false
	}
	return o.SerialInfo, true
}

// HasSerialInfo returns a boolean if a field has been set.
func (o *AsyncOrderCreateDTOWarrantyInfoInner) HasSerialInfo() bool {
	if o != nil && !IsNil(o.SerialInfo) {
		return true
	}

	return false
}

// SetSerialInfo gets a reference to the given []AsyncOrderCreateDTOWarrantyInfoInnerSerialInfoInner and assigns it to the SerialInfo field.
func (o *AsyncOrderCreateDTOWarrantyInfoInner) SetSerialInfo(v []AsyncOrderCreateDTOWarrantyInfoInnerSerialInfoInner) {
	o.SerialInfo = v
}

// GetVmfAdditionalAttributesLines returns the VmfAdditionalAttributesLines field value if set, zero value otherwise.
func (o *AsyncOrderCreateDTOWarrantyInfoInner) GetVmfAdditionalAttributesLines() []AsyncOrderCreateDTOWarrantyInfoInnerVmfAdditionalAttributesLinesInner {
	if o == nil || IsNil(o.VmfAdditionalAttributesLines) {
		var ret []AsyncOrderCreateDTOWarrantyInfoInnerVmfAdditionalAttributesLinesInner
		return ret
	}
	return o.VmfAdditionalAttributesLines
}

// GetVmfAdditionalAttributesLinesOk returns a tuple with the VmfAdditionalAttributesLines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsyncOrderCreateDTOWarrantyInfoInner) GetVmfAdditionalAttributesLinesOk() ([]AsyncOrderCreateDTOWarrantyInfoInnerVmfAdditionalAttributesLinesInner, bool) {
	if o == nil || IsNil(o.VmfAdditionalAttributesLines) {
		return nil, false
	}
	return o.VmfAdditionalAttributesLines, true
}

// HasVmfAdditionalAttributesLines returns a boolean if a field has been set.
func (o *AsyncOrderCreateDTOWarrantyInfoInner) HasVmfAdditionalAttributesLines() bool {
	if o != nil && !IsNil(o.VmfAdditionalAttributesLines) {
		return true
	}

	return false
}

// SetVmfAdditionalAttributesLines gets a reference to the given []AsyncOrderCreateDTOWarrantyInfoInnerVmfAdditionalAttributesLinesInner and assigns it to the VmfAdditionalAttributesLines field.
func (o *AsyncOrderCreateDTOWarrantyInfoInner) SetVmfAdditionalAttributesLines(v []AsyncOrderCreateDTOWarrantyInfoInnerVmfAdditionalAttributesLinesInner) {
	o.VmfAdditionalAttributesLines = v
}

func (o AsyncOrderCreateDTOWarrantyInfoInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AsyncOrderCreateDTOWarrantyInfoInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HardwareLineLink) {
		toSerialize["hardwareLineLink"] = o.HardwareLineLink
	}
	if !IsNil(o.WarrantyLineLink) {
		toSerialize["warrantyLineLink"] = o.WarrantyLineLink
	}
	if !IsNil(o.DirectLineLink) {
		toSerialize["directLineLink"] = o.DirectLineLink
	}
	if !IsNil(o.SerialInfo) {
		toSerialize["serialInfo"] = o.SerialInfo
	}
	if !IsNil(o.VmfAdditionalAttributesLines) {
		toSerialize["vmfAdditionalAttributesLines"] = o.VmfAdditionalAttributesLines
	}
	return toSerialize, nil
}

type NullableAsyncOrderCreateDTOWarrantyInfoInner struct {
	value *AsyncOrderCreateDTOWarrantyInfoInner
	isSet bool
}

func (v NullableAsyncOrderCreateDTOWarrantyInfoInner) Get() *AsyncOrderCreateDTOWarrantyInfoInner {
	return v.value
}

func (v *NullableAsyncOrderCreateDTOWarrantyInfoInner) Set(val *AsyncOrderCreateDTOWarrantyInfoInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAsyncOrderCreateDTOWarrantyInfoInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAsyncOrderCreateDTOWarrantyInfoInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAsyncOrderCreateDTOWarrantyInfoInner(val *AsyncOrderCreateDTOWarrantyInfoInner) *NullableAsyncOrderCreateDTOWarrantyInfoInner {
	return &NullableAsyncOrderCreateDTOWarrantyInfoInner{value: val, isSet: true}
}

func (v NullableAsyncOrderCreateDTOWarrantyInfoInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAsyncOrderCreateDTOWarrantyInfoInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

