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

// checks if the OrderDetailB2BLinesInnerServiceContractInfoLicenseInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderDetailB2BLinesInnerServiceContractInfoLicenseInfo{}

// OrderDetailB2BLinesInnerServiceContractInfoLicenseInfo struct for OrderDetailB2BLinesInnerServiceContractInfoLicenseInfo
type OrderDetailB2BLinesInnerServiceContractInfoLicenseInfo struct {
	// License numbers.
	LicenseNumber []string `json:"licenseNumber,omitempty"`
	// Start Date of the license.
	LicenseStartDate *string `json:"licenseStartDate,omitempty"`
	// End Date of the license.
	LicenseEndDate *string `json:"licenseEndDate,omitempty"`
	// Description of the license.
	Description *string `json:"description,omitempty"`
	// Quantity of the license.
	Quantity *string `json:"quantity,omitempty"`
}

// NewOrderDetailB2BLinesInnerServiceContractInfoLicenseInfo instantiates a new OrderDetailB2BLinesInnerServiceContractInfoLicenseInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderDetailB2BLinesInnerServiceContractInfoLicenseInfo() *OrderDetailB2BLinesInnerServiceContractInfoLicenseInfo {
	this := OrderDetailB2BLinesInnerServiceContractInfoLicenseInfo{}
	return &this
}

// NewOrderDetailB2BLinesInnerServiceContractInfoLicenseInfoWithDefaults instantiates a new OrderDetailB2BLinesInnerServiceContractInfoLicenseInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderDetailB2BLinesInnerServiceContractInfoLicenseInfoWithDefaults() *OrderDetailB2BLinesInnerServiceContractInfoLicenseInfo {
	this := OrderDetailB2BLinesInnerServiceContractInfoLicenseInfo{}
	return &this
}

// GetLicenseNumber returns the LicenseNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrderDetailB2BLinesInnerServiceContractInfoLicenseInfo) GetLicenseNumber() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.LicenseNumber
}

// GetLicenseNumberOk returns a tuple with the LicenseNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrderDetailB2BLinesInnerServiceContractInfoLicenseInfo) GetLicenseNumberOk() ([]string, bool) {
	if o == nil || IsNil(o.LicenseNumber) {
		return nil, false
	}
	return o.LicenseNumber, true
}

// HasLicenseNumber returns a boolean if a field has been set.
func (o *OrderDetailB2BLinesInnerServiceContractInfoLicenseInfo) HasLicenseNumber() bool {
	if o != nil && IsNil(o.LicenseNumber) {
		return true
	}

	return false
}

// SetLicenseNumber gets a reference to the given []string and assigns it to the LicenseNumber field.
func (o *OrderDetailB2BLinesInnerServiceContractInfoLicenseInfo) SetLicenseNumber(v []string) {
	o.LicenseNumber = v
}

// GetLicenseStartDate returns the LicenseStartDate field value if set, zero value otherwise.
func (o *OrderDetailB2BLinesInnerServiceContractInfoLicenseInfo) GetLicenseStartDate() string {
	if o == nil || IsNil(o.LicenseStartDate) {
		var ret string
		return ret
	}
	return *o.LicenseStartDate
}

// GetLicenseStartDateOk returns a tuple with the LicenseStartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailB2BLinesInnerServiceContractInfoLicenseInfo) GetLicenseStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.LicenseStartDate) {
		return nil, false
	}
	return o.LicenseStartDate, true
}

// HasLicenseStartDate returns a boolean if a field has been set.
func (o *OrderDetailB2BLinesInnerServiceContractInfoLicenseInfo) HasLicenseStartDate() bool {
	if o != nil && !IsNil(o.LicenseStartDate) {
		return true
	}

	return false
}

// SetLicenseStartDate gets a reference to the given string and assigns it to the LicenseStartDate field.
func (o *OrderDetailB2BLinesInnerServiceContractInfoLicenseInfo) SetLicenseStartDate(v string) {
	o.LicenseStartDate = &v
}

// GetLicenseEndDate returns the LicenseEndDate field value if set, zero value otherwise.
func (o *OrderDetailB2BLinesInnerServiceContractInfoLicenseInfo) GetLicenseEndDate() string {
	if o == nil || IsNil(o.LicenseEndDate) {
		var ret string
		return ret
	}
	return *o.LicenseEndDate
}

// GetLicenseEndDateOk returns a tuple with the LicenseEndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailB2BLinesInnerServiceContractInfoLicenseInfo) GetLicenseEndDateOk() (*string, bool) {
	if o == nil || IsNil(o.LicenseEndDate) {
		return nil, false
	}
	return o.LicenseEndDate, true
}

// HasLicenseEndDate returns a boolean if a field has been set.
func (o *OrderDetailB2BLinesInnerServiceContractInfoLicenseInfo) HasLicenseEndDate() bool {
	if o != nil && !IsNil(o.LicenseEndDate) {
		return true
	}

	return false
}

// SetLicenseEndDate gets a reference to the given string and assigns it to the LicenseEndDate field.
func (o *OrderDetailB2BLinesInnerServiceContractInfoLicenseInfo) SetLicenseEndDate(v string) {
	o.LicenseEndDate = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *OrderDetailB2BLinesInnerServiceContractInfoLicenseInfo) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailB2BLinesInnerServiceContractInfoLicenseInfo) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *OrderDetailB2BLinesInnerServiceContractInfoLicenseInfo) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *OrderDetailB2BLinesInnerServiceContractInfoLicenseInfo) SetDescription(v string) {
	o.Description = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *OrderDetailB2BLinesInnerServiceContractInfoLicenseInfo) GetQuantity() string {
	if o == nil || IsNil(o.Quantity) {
		var ret string
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailB2BLinesInnerServiceContractInfoLicenseInfo) GetQuantityOk() (*string, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *OrderDetailB2BLinesInnerServiceContractInfoLicenseInfo) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given string and assigns it to the Quantity field.
func (o *OrderDetailB2BLinesInnerServiceContractInfoLicenseInfo) SetQuantity(v string) {
	o.Quantity = &v
}

func (o OrderDetailB2BLinesInnerServiceContractInfoLicenseInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderDetailB2BLinesInnerServiceContractInfoLicenseInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.LicenseNumber != nil {
		toSerialize["licenseNumber"] = o.LicenseNumber
	}
	if !IsNil(o.LicenseStartDate) {
		toSerialize["licenseStartDate"] = o.LicenseStartDate
	}
	if !IsNil(o.LicenseEndDate) {
		toSerialize["licenseEndDate"] = o.LicenseEndDate
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	return toSerialize, nil
}

type NullableOrderDetailB2BLinesInnerServiceContractInfoLicenseInfo struct {
	value *OrderDetailB2BLinesInnerServiceContractInfoLicenseInfo
	isSet bool
}

func (v NullableOrderDetailB2BLinesInnerServiceContractInfoLicenseInfo) Get() *OrderDetailB2BLinesInnerServiceContractInfoLicenseInfo {
	return v.value
}

func (v *NullableOrderDetailB2BLinesInnerServiceContractInfoLicenseInfo) Set(val *OrderDetailB2BLinesInnerServiceContractInfoLicenseInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderDetailB2BLinesInnerServiceContractInfoLicenseInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderDetailB2BLinesInnerServiceContractInfoLicenseInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderDetailB2BLinesInnerServiceContractInfoLicenseInfo(val *OrderDetailB2BLinesInnerServiceContractInfoLicenseInfo) *NullableOrderDetailB2BLinesInnerServiceContractInfoLicenseInfo {
	return &NullableOrderDetailB2BLinesInnerServiceContractInfoLicenseInfo{value: val, isSet: true}
}

func (v NullableOrderDetailB2BLinesInnerServiceContractInfoLicenseInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderDetailB2BLinesInnerServiceContractInfoLicenseInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


