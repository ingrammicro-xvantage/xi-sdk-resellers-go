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

// checks if the RenewalsSearchRequestDataTypeEndDate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RenewalsSearchRequestDataTypeEndDate{}

// RenewalsSearchRequestDataTypeEndDate struct for RenewalsSearchRequestDataTypeEndDate
type RenewalsSearchRequestDataTypeEndDate struct {
	// Custom from date for Renewal end date.
	CustomStartDate *string `json:"customStartDate,omitempty"`
	// Custom to date for Renewal end date.
	CustomEndDate *string `json:"customEndDate,omitempty"`
}

// NewRenewalsSearchRequestDataTypeEndDate instantiates a new RenewalsSearchRequestDataTypeEndDate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRenewalsSearchRequestDataTypeEndDate() *RenewalsSearchRequestDataTypeEndDate {
	this := RenewalsSearchRequestDataTypeEndDate{}
	return &this
}

// NewRenewalsSearchRequestDataTypeEndDateWithDefaults instantiates a new RenewalsSearchRequestDataTypeEndDate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRenewalsSearchRequestDataTypeEndDateWithDefaults() *RenewalsSearchRequestDataTypeEndDate {
	this := RenewalsSearchRequestDataTypeEndDate{}
	return &this
}

// GetCustomStartDate returns the CustomStartDate field value if set, zero value otherwise.
func (o *RenewalsSearchRequestDataTypeEndDate) GetCustomStartDate() string {
	if o == nil || IsNil(o.CustomStartDate) {
		var ret string
		return ret
	}
	return *o.CustomStartDate
}

// GetCustomStartDateOk returns a tuple with the CustomStartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenewalsSearchRequestDataTypeEndDate) GetCustomStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.CustomStartDate) {
		return nil, false
	}
	return o.CustomStartDate, true
}

// HasCustomStartDate returns a boolean if a field has been set.
func (o *RenewalsSearchRequestDataTypeEndDate) HasCustomStartDate() bool {
	if o != nil && !IsNil(o.CustomStartDate) {
		return true
	}

	return false
}

// SetCustomStartDate gets a reference to the given string and assigns it to the CustomStartDate field.
func (o *RenewalsSearchRequestDataTypeEndDate) SetCustomStartDate(v string) {
	o.CustomStartDate = &v
}

// GetCustomEndDate returns the CustomEndDate field value if set, zero value otherwise.
func (o *RenewalsSearchRequestDataTypeEndDate) GetCustomEndDate() string {
	if o == nil || IsNil(o.CustomEndDate) {
		var ret string
		return ret
	}
	return *o.CustomEndDate
}

// GetCustomEndDateOk returns a tuple with the CustomEndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenewalsSearchRequestDataTypeEndDate) GetCustomEndDateOk() (*string, bool) {
	if o == nil || IsNil(o.CustomEndDate) {
		return nil, false
	}
	return o.CustomEndDate, true
}

// HasCustomEndDate returns a boolean if a field has been set.
func (o *RenewalsSearchRequestDataTypeEndDate) HasCustomEndDate() bool {
	if o != nil && !IsNil(o.CustomEndDate) {
		return true
	}

	return false
}

// SetCustomEndDate gets a reference to the given string and assigns it to the CustomEndDate field.
func (o *RenewalsSearchRequestDataTypeEndDate) SetCustomEndDate(v string) {
	o.CustomEndDate = &v
}

func (o RenewalsSearchRequestDataTypeEndDate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RenewalsSearchRequestDataTypeEndDate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CustomStartDate) {
		toSerialize["customStartDate"] = o.CustomStartDate
	}
	if !IsNil(o.CustomEndDate) {
		toSerialize["customEndDate"] = o.CustomEndDate
	}
	return toSerialize, nil
}

type NullableRenewalsSearchRequestDataTypeEndDate struct {
	value *RenewalsSearchRequestDataTypeEndDate
	isSet bool
}

func (v NullableRenewalsSearchRequestDataTypeEndDate) Get() *RenewalsSearchRequestDataTypeEndDate {
	return v.value
}

func (v *NullableRenewalsSearchRequestDataTypeEndDate) Set(val *RenewalsSearchRequestDataTypeEndDate) {
	v.value = val
	v.isSet = true
}

func (v NullableRenewalsSearchRequestDataTypeEndDate) IsSet() bool {
	return v.isSet
}

func (v *NullableRenewalsSearchRequestDataTypeEndDate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRenewalsSearchRequestDataTypeEndDate(val *RenewalsSearchRequestDataTypeEndDate) *NullableRenewalsSearchRequestDataTypeEndDate {
	return &NullableRenewalsSearchRequestDataTypeEndDate{value: val, isSet: true}
}

func (v NullableRenewalsSearchRequestDataTypeEndDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRenewalsSearchRequestDataTypeEndDate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


