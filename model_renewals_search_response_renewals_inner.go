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

// checks if the RenewalsSearchResponseRenewalsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RenewalsSearchResponseRenewalsInner{}

// RenewalsSearchResponseRenewalsInner struct for RenewalsSearchResponseRenewalsInner
type RenewalsSearchResponseRenewalsInner struct {
	// Unique renewal ID.
	RenewalId *int32 `json:"renewalId,omitempty"`
	// The reseller's order number for reference in their system.
	CustomerOrderNumber *string `json:"customerOrderNumber,omitempty"`
	// Renewal reference number. It could be notification id or quote number.
	ReferenceNumber *string `json:"referenceNumber,omitempty"`
	// The company name for the end user/customer.
	EndUser *string `json:"endUser,omitempty"`
	// The name of the vendor.
	Vendor *string `json:"vendor,omitempty"`
	// Renewal expiration date.
	ExpirationDate *string `json:"expirationDate,omitempty"`
	// The value of the renewal.
	RenewalValue *float64 `json:"renewalValue,omitempty"`
	// The status of the renewal.
	Status *string `json:"status,omitempty"`
	Links []RenewalsSearchResponseRenewalsInnerLinksInner `json:"links,omitempty"`
}

// NewRenewalsSearchResponseRenewalsInner instantiates a new RenewalsSearchResponseRenewalsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRenewalsSearchResponseRenewalsInner() *RenewalsSearchResponseRenewalsInner {
	this := RenewalsSearchResponseRenewalsInner{}
	return &this
}

// NewRenewalsSearchResponseRenewalsInnerWithDefaults instantiates a new RenewalsSearchResponseRenewalsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRenewalsSearchResponseRenewalsInnerWithDefaults() *RenewalsSearchResponseRenewalsInner {
	this := RenewalsSearchResponseRenewalsInner{}
	return &this
}

// GetRenewalId returns the RenewalId field value if set, zero value otherwise.
func (o *RenewalsSearchResponseRenewalsInner) GetRenewalId() int32 {
	if o == nil || IsNil(o.RenewalId) {
		var ret int32
		return ret
	}
	return *o.RenewalId
}

// GetRenewalIdOk returns a tuple with the RenewalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenewalsSearchResponseRenewalsInner) GetRenewalIdOk() (*int32, bool) {
	if o == nil || IsNil(o.RenewalId) {
		return nil, false
	}
	return o.RenewalId, true
}

// HasRenewalId returns a boolean if a field has been set.
func (o *RenewalsSearchResponseRenewalsInner) HasRenewalId() bool {
	if o != nil && !IsNil(o.RenewalId) {
		return true
	}

	return false
}

// SetRenewalId gets a reference to the given int32 and assigns it to the RenewalId field.
func (o *RenewalsSearchResponseRenewalsInner) SetRenewalId(v int32) {
	o.RenewalId = &v
}

// GetCustomerOrderNumber returns the CustomerOrderNumber field value if set, zero value otherwise.
func (o *RenewalsSearchResponseRenewalsInner) GetCustomerOrderNumber() string {
	if o == nil || IsNil(o.CustomerOrderNumber) {
		var ret string
		return ret
	}
	return *o.CustomerOrderNumber
}

// GetCustomerOrderNumberOk returns a tuple with the CustomerOrderNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenewalsSearchResponseRenewalsInner) GetCustomerOrderNumberOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerOrderNumber) {
		return nil, false
	}
	return o.CustomerOrderNumber, true
}

// HasCustomerOrderNumber returns a boolean if a field has been set.
func (o *RenewalsSearchResponseRenewalsInner) HasCustomerOrderNumber() bool {
	if o != nil && !IsNil(o.CustomerOrderNumber) {
		return true
	}

	return false
}

// SetCustomerOrderNumber gets a reference to the given string and assigns it to the CustomerOrderNumber field.
func (o *RenewalsSearchResponseRenewalsInner) SetCustomerOrderNumber(v string) {
	o.CustomerOrderNumber = &v
}

// GetReferenceNumber returns the ReferenceNumber field value if set, zero value otherwise.
func (o *RenewalsSearchResponseRenewalsInner) GetReferenceNumber() string {
	if o == nil || IsNil(o.ReferenceNumber) {
		var ret string
		return ret
	}
	return *o.ReferenceNumber
}

// GetReferenceNumberOk returns a tuple with the ReferenceNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenewalsSearchResponseRenewalsInner) GetReferenceNumberOk() (*string, bool) {
	if o == nil || IsNil(o.ReferenceNumber) {
		return nil, false
	}
	return o.ReferenceNumber, true
}

// HasReferenceNumber returns a boolean if a field has been set.
func (o *RenewalsSearchResponseRenewalsInner) HasReferenceNumber() bool {
	if o != nil && !IsNil(o.ReferenceNumber) {
		return true
	}

	return false
}

// SetReferenceNumber gets a reference to the given string and assigns it to the ReferenceNumber field.
func (o *RenewalsSearchResponseRenewalsInner) SetReferenceNumber(v string) {
	o.ReferenceNumber = &v
}

// GetEndUser returns the EndUser field value if set, zero value otherwise.
func (o *RenewalsSearchResponseRenewalsInner) GetEndUser() string {
	if o == nil || IsNil(o.EndUser) {
		var ret string
		return ret
	}
	return *o.EndUser
}

// GetEndUserOk returns a tuple with the EndUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenewalsSearchResponseRenewalsInner) GetEndUserOk() (*string, bool) {
	if o == nil || IsNil(o.EndUser) {
		return nil, false
	}
	return o.EndUser, true
}

// HasEndUser returns a boolean if a field has been set.
func (o *RenewalsSearchResponseRenewalsInner) HasEndUser() bool {
	if o != nil && !IsNil(o.EndUser) {
		return true
	}

	return false
}

// SetEndUser gets a reference to the given string and assigns it to the EndUser field.
func (o *RenewalsSearchResponseRenewalsInner) SetEndUser(v string) {
	o.EndUser = &v
}

// GetVendor returns the Vendor field value if set, zero value otherwise.
func (o *RenewalsSearchResponseRenewalsInner) GetVendor() string {
	if o == nil || IsNil(o.Vendor) {
		var ret string
		return ret
	}
	return *o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenewalsSearchResponseRenewalsInner) GetVendorOk() (*string, bool) {
	if o == nil || IsNil(o.Vendor) {
		return nil, false
	}
	return o.Vendor, true
}

// HasVendor returns a boolean if a field has been set.
func (o *RenewalsSearchResponseRenewalsInner) HasVendor() bool {
	if o != nil && !IsNil(o.Vendor) {
		return true
	}

	return false
}

// SetVendor gets a reference to the given string and assigns it to the Vendor field.
func (o *RenewalsSearchResponseRenewalsInner) SetVendor(v string) {
	o.Vendor = &v
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise.
func (o *RenewalsSearchResponseRenewalsInner) GetExpirationDate() string {
	if o == nil || IsNil(o.ExpirationDate) {
		var ret string
		return ret
	}
	return *o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenewalsSearchResponseRenewalsInner) GetExpirationDateOk() (*string, bool) {
	if o == nil || IsNil(o.ExpirationDate) {
		return nil, false
	}
	return o.ExpirationDate, true
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *RenewalsSearchResponseRenewalsInner) HasExpirationDate() bool {
	if o != nil && !IsNil(o.ExpirationDate) {
		return true
	}

	return false
}

// SetExpirationDate gets a reference to the given string and assigns it to the ExpirationDate field.
func (o *RenewalsSearchResponseRenewalsInner) SetExpirationDate(v string) {
	o.ExpirationDate = &v
}

// GetRenewalValue returns the RenewalValue field value if set, zero value otherwise.
func (o *RenewalsSearchResponseRenewalsInner) GetRenewalValue() float64 {
	if o == nil || IsNil(o.RenewalValue) {
		var ret float64
		return ret
	}
	return *o.RenewalValue
}

// GetRenewalValueOk returns a tuple with the RenewalValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenewalsSearchResponseRenewalsInner) GetRenewalValueOk() (*float64, bool) {
	if o == nil || IsNil(o.RenewalValue) {
		return nil, false
	}
	return o.RenewalValue, true
}

// HasRenewalValue returns a boolean if a field has been set.
func (o *RenewalsSearchResponseRenewalsInner) HasRenewalValue() bool {
	if o != nil && !IsNil(o.RenewalValue) {
		return true
	}

	return false
}

// SetRenewalValue gets a reference to the given float64 and assigns it to the RenewalValue field.
func (o *RenewalsSearchResponseRenewalsInner) SetRenewalValue(v float64) {
	o.RenewalValue = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *RenewalsSearchResponseRenewalsInner) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenewalsSearchResponseRenewalsInner) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *RenewalsSearchResponseRenewalsInner) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *RenewalsSearchResponseRenewalsInner) SetStatus(v string) {
	o.Status = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *RenewalsSearchResponseRenewalsInner) GetLinks() []RenewalsSearchResponseRenewalsInnerLinksInner {
	if o == nil || IsNil(o.Links) {
		var ret []RenewalsSearchResponseRenewalsInnerLinksInner
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenewalsSearchResponseRenewalsInner) GetLinksOk() ([]RenewalsSearchResponseRenewalsInnerLinksInner, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *RenewalsSearchResponseRenewalsInner) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []RenewalsSearchResponseRenewalsInnerLinksInner and assigns it to the Links field.
func (o *RenewalsSearchResponseRenewalsInner) SetLinks(v []RenewalsSearchResponseRenewalsInnerLinksInner) {
	o.Links = v
}

func (o RenewalsSearchResponseRenewalsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RenewalsSearchResponseRenewalsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RenewalId) {
		toSerialize["renewalId"] = o.RenewalId
	}
	if !IsNil(o.CustomerOrderNumber) {
		toSerialize["customerOrderNumber"] = o.CustomerOrderNumber
	}
	if !IsNil(o.ReferenceNumber) {
		toSerialize["referenceNumber"] = o.ReferenceNumber
	}
	if !IsNil(o.EndUser) {
		toSerialize["endUser"] = o.EndUser
	}
	if !IsNil(o.Vendor) {
		toSerialize["vendor"] = o.Vendor
	}
	if !IsNil(o.ExpirationDate) {
		toSerialize["expirationDate"] = o.ExpirationDate
	}
	if !IsNil(o.RenewalValue) {
		toSerialize["renewalValue"] = o.RenewalValue
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	return toSerialize, nil
}

type NullableRenewalsSearchResponseRenewalsInner struct {
	value *RenewalsSearchResponseRenewalsInner
	isSet bool
}

func (v NullableRenewalsSearchResponseRenewalsInner) Get() *RenewalsSearchResponseRenewalsInner {
	return v.value
}

func (v *NullableRenewalsSearchResponseRenewalsInner) Set(val *RenewalsSearchResponseRenewalsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableRenewalsSearchResponseRenewalsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableRenewalsSearchResponseRenewalsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRenewalsSearchResponseRenewalsInner(val *RenewalsSearchResponseRenewalsInner) *NullableRenewalsSearchResponseRenewalsInner {
	return &NullableRenewalsSearchResponseRenewalsInner{value: val, isSet: true}
}

func (v NullableRenewalsSearchResponseRenewalsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRenewalsSearchResponseRenewalsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


