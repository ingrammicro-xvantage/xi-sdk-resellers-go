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

// checks if the OrderCreateRequestResellerInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderCreateRequestResellerInfo{}

// OrderCreateRequestResellerInfo The address and contact information provided by the reseller.
type OrderCreateRequestResellerInfo struct {
	// The reseller's Id.
	ResellerId *string `json:"resellerId,omitempty"`
	// The reseller's company name.
	CompanyName *string `json:"companyName,omitempty"`
	// The reseller's contact name.
	Contact *string `json:"contact,omitempty"`
	// The reseller's street address.
	AddressLine1 *string `json:"addressLine1,omitempty"`
	// The reseller's building or apartment number.
	AddressLine2 *string `json:"addressLine2,omitempty"`
	// The reseller's address line 3.
	AddressLine3 *string `json:"addressLine3,omitempty"`
	// The reseller's city.
	City *string `json:"city,omitempty"`
	// The reseller's state.
	State *string `json:"state,omitempty"`
	// The reseller's zip or postal code.
	PostalCode *string `json:"postalCode,omitempty"`
	// The reseller's two-character ISO country code.
	CountryCode *string `json:"countryCode,omitempty"`
	// The reseller's phone number.
	PhoneNumber *int32 `json:"phoneNumber,omitempty"`
	// The reseller's email address.
	Email *string `json:"email,omitempty"`
}

// NewOrderCreateRequestResellerInfo instantiates a new OrderCreateRequestResellerInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderCreateRequestResellerInfo() *OrderCreateRequestResellerInfo {
	this := OrderCreateRequestResellerInfo{}
	return &this
}

// NewOrderCreateRequestResellerInfoWithDefaults instantiates a new OrderCreateRequestResellerInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderCreateRequestResellerInfoWithDefaults() *OrderCreateRequestResellerInfo {
	this := OrderCreateRequestResellerInfo{}
	return &this
}

// GetResellerId returns the ResellerId field value if set, zero value otherwise.
func (o *OrderCreateRequestResellerInfo) GetResellerId() string {
	if o == nil || IsNil(o.ResellerId) {
		var ret string
		return ret
	}
	return *o.ResellerId
}

// GetResellerIdOk returns a tuple with the ResellerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateRequestResellerInfo) GetResellerIdOk() (*string, bool) {
	if o == nil || IsNil(o.ResellerId) {
		return nil, false
	}
	return o.ResellerId, true
}

// HasResellerId returns a boolean if a field has been set.
func (o *OrderCreateRequestResellerInfo) HasResellerId() bool {
	if o != nil && !IsNil(o.ResellerId) {
		return true
	}

	return false
}

// SetResellerId gets a reference to the given string and assigns it to the ResellerId field.
func (o *OrderCreateRequestResellerInfo) SetResellerId(v string) {
	o.ResellerId = &v
}

// GetCompanyName returns the CompanyName field value if set, zero value otherwise.
func (o *OrderCreateRequestResellerInfo) GetCompanyName() string {
	if o == nil || IsNil(o.CompanyName) {
		var ret string
		return ret
	}
	return *o.CompanyName
}

// GetCompanyNameOk returns a tuple with the CompanyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateRequestResellerInfo) GetCompanyNameOk() (*string, bool) {
	if o == nil || IsNil(o.CompanyName) {
		return nil, false
	}
	return o.CompanyName, true
}

// HasCompanyName returns a boolean if a field has been set.
func (o *OrderCreateRequestResellerInfo) HasCompanyName() bool {
	if o != nil && !IsNil(o.CompanyName) {
		return true
	}

	return false
}

// SetCompanyName gets a reference to the given string and assigns it to the CompanyName field.
func (o *OrderCreateRequestResellerInfo) SetCompanyName(v string) {
	o.CompanyName = &v
}

// GetContact returns the Contact field value if set, zero value otherwise.
func (o *OrderCreateRequestResellerInfo) GetContact() string {
	if o == nil || IsNil(o.Contact) {
		var ret string
		return ret
	}
	return *o.Contact
}

// GetContactOk returns a tuple with the Contact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateRequestResellerInfo) GetContactOk() (*string, bool) {
	if o == nil || IsNil(o.Contact) {
		return nil, false
	}
	return o.Contact, true
}

// HasContact returns a boolean if a field has been set.
func (o *OrderCreateRequestResellerInfo) HasContact() bool {
	if o != nil && !IsNil(o.Contact) {
		return true
	}

	return false
}

// SetContact gets a reference to the given string and assigns it to the Contact field.
func (o *OrderCreateRequestResellerInfo) SetContact(v string) {
	o.Contact = &v
}

// GetAddressLine1 returns the AddressLine1 field value if set, zero value otherwise.
func (o *OrderCreateRequestResellerInfo) GetAddressLine1() string {
	if o == nil || IsNil(o.AddressLine1) {
		var ret string
		return ret
	}
	return *o.AddressLine1
}

// GetAddressLine1Ok returns a tuple with the AddressLine1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateRequestResellerInfo) GetAddressLine1Ok() (*string, bool) {
	if o == nil || IsNil(o.AddressLine1) {
		return nil, false
	}
	return o.AddressLine1, true
}

// HasAddressLine1 returns a boolean if a field has been set.
func (o *OrderCreateRequestResellerInfo) HasAddressLine1() bool {
	if o != nil && !IsNil(o.AddressLine1) {
		return true
	}

	return false
}

// SetAddressLine1 gets a reference to the given string and assigns it to the AddressLine1 field.
func (o *OrderCreateRequestResellerInfo) SetAddressLine1(v string) {
	o.AddressLine1 = &v
}

// GetAddressLine2 returns the AddressLine2 field value if set, zero value otherwise.
func (o *OrderCreateRequestResellerInfo) GetAddressLine2() string {
	if o == nil || IsNil(o.AddressLine2) {
		var ret string
		return ret
	}
	return *o.AddressLine2
}

// GetAddressLine2Ok returns a tuple with the AddressLine2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateRequestResellerInfo) GetAddressLine2Ok() (*string, bool) {
	if o == nil || IsNil(o.AddressLine2) {
		return nil, false
	}
	return o.AddressLine2, true
}

// HasAddressLine2 returns a boolean if a field has been set.
func (o *OrderCreateRequestResellerInfo) HasAddressLine2() bool {
	if o != nil && !IsNil(o.AddressLine2) {
		return true
	}

	return false
}

// SetAddressLine2 gets a reference to the given string and assigns it to the AddressLine2 field.
func (o *OrderCreateRequestResellerInfo) SetAddressLine2(v string) {
	o.AddressLine2 = &v
}

// GetAddressLine3 returns the AddressLine3 field value if set, zero value otherwise.
func (o *OrderCreateRequestResellerInfo) GetAddressLine3() string {
	if o == nil || IsNil(o.AddressLine3) {
		var ret string
		return ret
	}
	return *o.AddressLine3
}

// GetAddressLine3Ok returns a tuple with the AddressLine3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateRequestResellerInfo) GetAddressLine3Ok() (*string, bool) {
	if o == nil || IsNil(o.AddressLine3) {
		return nil, false
	}
	return o.AddressLine3, true
}

// HasAddressLine3 returns a boolean if a field has been set.
func (o *OrderCreateRequestResellerInfo) HasAddressLine3() bool {
	if o != nil && !IsNil(o.AddressLine3) {
		return true
	}

	return false
}

// SetAddressLine3 gets a reference to the given string and assigns it to the AddressLine3 field.
func (o *OrderCreateRequestResellerInfo) SetAddressLine3(v string) {
	o.AddressLine3 = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *OrderCreateRequestResellerInfo) GetCity() string {
	if o == nil || IsNil(o.City) {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateRequestResellerInfo) GetCityOk() (*string, bool) {
	if o == nil || IsNil(o.City) {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *OrderCreateRequestResellerInfo) HasCity() bool {
	if o != nil && !IsNil(o.City) {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *OrderCreateRequestResellerInfo) SetCity(v string) {
	o.City = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *OrderCreateRequestResellerInfo) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateRequestResellerInfo) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *OrderCreateRequestResellerInfo) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *OrderCreateRequestResellerInfo) SetState(v string) {
	o.State = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *OrderCreateRequestResellerInfo) GetPostalCode() string {
	if o == nil || IsNil(o.PostalCode) {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateRequestResellerInfo) GetPostalCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PostalCode) {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *OrderCreateRequestResellerInfo) HasPostalCode() bool {
	if o != nil && !IsNil(o.PostalCode) {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *OrderCreateRequestResellerInfo) SetPostalCode(v string) {
	o.PostalCode = &v
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *OrderCreateRequestResellerInfo) GetCountryCode() string {
	if o == nil || IsNil(o.CountryCode) {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateRequestResellerInfo) GetCountryCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CountryCode) {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *OrderCreateRequestResellerInfo) HasCountryCode() bool {
	if o != nil && !IsNil(o.CountryCode) {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *OrderCreateRequestResellerInfo) SetCountryCode(v string) {
	o.CountryCode = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *OrderCreateRequestResellerInfo) GetPhoneNumber() int32 {
	if o == nil || IsNil(o.PhoneNumber) {
		var ret int32
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateRequestResellerInfo) GetPhoneNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.PhoneNumber) {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *OrderCreateRequestResellerInfo) HasPhoneNumber() bool {
	if o != nil && !IsNil(o.PhoneNumber) {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given int32 and assigns it to the PhoneNumber field.
func (o *OrderCreateRequestResellerInfo) SetPhoneNumber(v int32) {
	o.PhoneNumber = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *OrderCreateRequestResellerInfo) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateRequestResellerInfo) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *OrderCreateRequestResellerInfo) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *OrderCreateRequestResellerInfo) SetEmail(v string) {
	o.Email = &v
}

func (o OrderCreateRequestResellerInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderCreateRequestResellerInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ResellerId) {
		toSerialize["resellerId"] = o.ResellerId
	}
	if !IsNil(o.CompanyName) {
		toSerialize["companyName"] = o.CompanyName
	}
	if !IsNil(o.Contact) {
		toSerialize["contact"] = o.Contact
	}
	if !IsNil(o.AddressLine1) {
		toSerialize["addressLine1"] = o.AddressLine1
	}
	if !IsNil(o.AddressLine2) {
		toSerialize["addressLine2"] = o.AddressLine2
	}
	if !IsNil(o.AddressLine3) {
		toSerialize["addressLine3"] = o.AddressLine3
	}
	if !IsNil(o.City) {
		toSerialize["city"] = o.City
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.PostalCode) {
		toSerialize["postalCode"] = o.PostalCode
	}
	if !IsNil(o.CountryCode) {
		toSerialize["countryCode"] = o.CountryCode
	}
	if !IsNil(o.PhoneNumber) {
		toSerialize["phoneNumber"] = o.PhoneNumber
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	return toSerialize, nil
}

type NullableOrderCreateRequestResellerInfo struct {
	value *OrderCreateRequestResellerInfo
	isSet bool
}

func (v NullableOrderCreateRequestResellerInfo) Get() *OrderCreateRequestResellerInfo {
	return v.value
}

func (v *NullableOrderCreateRequestResellerInfo) Set(val *OrderCreateRequestResellerInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderCreateRequestResellerInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderCreateRequestResellerInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderCreateRequestResellerInfo(val *OrderCreateRequestResellerInfo) *NullableOrderCreateRequestResellerInfo {
	return &NullableOrderCreateRequestResellerInfo{value: val, isSet: true}
}

func (v NullableOrderCreateRequestResellerInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderCreateRequestResellerInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


