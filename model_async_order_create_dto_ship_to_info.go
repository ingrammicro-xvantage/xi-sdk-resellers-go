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

// checks if the AsyncOrderCreateDTOShipToInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AsyncOrderCreateDTOShipToInfo{}

// AsyncOrderCreateDTOShipToInfo The shipping information provided by the reseller for order delivery.
type AsyncOrderCreateDTOShipToInfo struct {
	// The company contact provided by the reseller.
	AddressId *string `json:"addressId,omitempty"`
	// The contact name for the order will be shipped to.
	Contact *string `json:"contact,omitempty"`
	// The name of the company the order will be shipped to.
	CompanyName *string `json:"companyName,omitempty"`
	// The address line 1 the order will be shipped to.
	AddressLine1 *string `json:"addressLine1,omitempty"`
	// The address line 2 the order will be shipped to.
	AddressLine2 *string `json:"addressLine2,omitempty"`
	// The address line 3 the order will be shipped to.
	AddressLine3 NullableString `json:"addressLine3,omitempty"`
	// The address line 4 the order will be shipped to.
	AddressLine4 NullableString `json:"addressLine4,omitempty"`
	// Need description
	Name1 *string `json:"name1,omitempty"`
	// Need description
	Name2 *string `json:"name2,omitempty"`
	// The city the order will be shipped to.
	City *string `json:"city,omitempty"`
	// The state the order will be shipped to.
	State *string `json:"state,omitempty"`
	// The zip or postal code the order will be shipped to.
	PostalCode *string `json:"postalCode,omitempty"`
	// The two-character ISO country code the order will be shipped to.
	CountryCode *string `json:"countryCode,omitempty"`
	// The company contact email address.
	Email NullableString `json:"email,omitempty"`
	// Shipping Notes
	ShippingNotes *string `json:"shippingNotes,omitempty"`
	// Phone number for shipping
	PhoneNumber *string `json:"phoneNumber,omitempty"`
}

// NewAsyncOrderCreateDTOShipToInfo instantiates a new AsyncOrderCreateDTOShipToInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAsyncOrderCreateDTOShipToInfo() *AsyncOrderCreateDTOShipToInfo {
	this := AsyncOrderCreateDTOShipToInfo{}
	return &this
}

// NewAsyncOrderCreateDTOShipToInfoWithDefaults instantiates a new AsyncOrderCreateDTOShipToInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAsyncOrderCreateDTOShipToInfoWithDefaults() *AsyncOrderCreateDTOShipToInfo {
	this := AsyncOrderCreateDTOShipToInfo{}
	return &this
}

// GetAddressId returns the AddressId field value if set, zero value otherwise.
func (o *AsyncOrderCreateDTOShipToInfo) GetAddressId() string {
	if o == nil || IsNil(o.AddressId) {
		var ret string
		return ret
	}
	return *o.AddressId
}

// GetAddressIdOk returns a tuple with the AddressId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsyncOrderCreateDTOShipToInfo) GetAddressIdOk() (*string, bool) {
	if o == nil || IsNil(o.AddressId) {
		return nil, false
	}
	return o.AddressId, true
}

// HasAddressId returns a boolean if a field has been set.
func (o *AsyncOrderCreateDTOShipToInfo) HasAddressId() bool {
	if o != nil && !IsNil(o.AddressId) {
		return true
	}

	return false
}

// SetAddressId gets a reference to the given string and assigns it to the AddressId field.
func (o *AsyncOrderCreateDTOShipToInfo) SetAddressId(v string) {
	o.AddressId = &v
}

// GetContact returns the Contact field value if set, zero value otherwise.
func (o *AsyncOrderCreateDTOShipToInfo) GetContact() string {
	if o == nil || IsNil(o.Contact) {
		var ret string
		return ret
	}
	return *o.Contact
}

// GetContactOk returns a tuple with the Contact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsyncOrderCreateDTOShipToInfo) GetContactOk() (*string, bool) {
	if o == nil || IsNil(o.Contact) {
		return nil, false
	}
	return o.Contact, true
}

// HasContact returns a boolean if a field has been set.
func (o *AsyncOrderCreateDTOShipToInfo) HasContact() bool {
	if o != nil && !IsNil(o.Contact) {
		return true
	}

	return false
}

// SetContact gets a reference to the given string and assigns it to the Contact field.
func (o *AsyncOrderCreateDTOShipToInfo) SetContact(v string) {
	o.Contact = &v
}

// GetCompanyName returns the CompanyName field value if set, zero value otherwise.
func (o *AsyncOrderCreateDTOShipToInfo) GetCompanyName() string {
	if o == nil || IsNil(o.CompanyName) {
		var ret string
		return ret
	}
	return *o.CompanyName
}

// GetCompanyNameOk returns a tuple with the CompanyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsyncOrderCreateDTOShipToInfo) GetCompanyNameOk() (*string, bool) {
	if o == nil || IsNil(o.CompanyName) {
		return nil, false
	}
	return o.CompanyName, true
}

// HasCompanyName returns a boolean if a field has been set.
func (o *AsyncOrderCreateDTOShipToInfo) HasCompanyName() bool {
	if o != nil && !IsNil(o.CompanyName) {
		return true
	}

	return false
}

// SetCompanyName gets a reference to the given string and assigns it to the CompanyName field.
func (o *AsyncOrderCreateDTOShipToInfo) SetCompanyName(v string) {
	o.CompanyName = &v
}

// GetAddressLine1 returns the AddressLine1 field value if set, zero value otherwise.
func (o *AsyncOrderCreateDTOShipToInfo) GetAddressLine1() string {
	if o == nil || IsNil(o.AddressLine1) {
		var ret string
		return ret
	}
	return *o.AddressLine1
}

// GetAddressLine1Ok returns a tuple with the AddressLine1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsyncOrderCreateDTOShipToInfo) GetAddressLine1Ok() (*string, bool) {
	if o == nil || IsNil(o.AddressLine1) {
		return nil, false
	}
	return o.AddressLine1, true
}

// HasAddressLine1 returns a boolean if a field has been set.
func (o *AsyncOrderCreateDTOShipToInfo) HasAddressLine1() bool {
	if o != nil && !IsNil(o.AddressLine1) {
		return true
	}

	return false
}

// SetAddressLine1 gets a reference to the given string and assigns it to the AddressLine1 field.
func (o *AsyncOrderCreateDTOShipToInfo) SetAddressLine1(v string) {
	o.AddressLine1 = &v
}

// GetAddressLine2 returns the AddressLine2 field value if set, zero value otherwise.
func (o *AsyncOrderCreateDTOShipToInfo) GetAddressLine2() string {
	if o == nil || IsNil(o.AddressLine2) {
		var ret string
		return ret
	}
	return *o.AddressLine2
}

// GetAddressLine2Ok returns a tuple with the AddressLine2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsyncOrderCreateDTOShipToInfo) GetAddressLine2Ok() (*string, bool) {
	if o == nil || IsNil(o.AddressLine2) {
		return nil, false
	}
	return o.AddressLine2, true
}

// HasAddressLine2 returns a boolean if a field has been set.
func (o *AsyncOrderCreateDTOShipToInfo) HasAddressLine2() bool {
	if o != nil && !IsNil(o.AddressLine2) {
		return true
	}

	return false
}

// SetAddressLine2 gets a reference to the given string and assigns it to the AddressLine2 field.
func (o *AsyncOrderCreateDTOShipToInfo) SetAddressLine2(v string) {
	o.AddressLine2 = &v
}

// GetAddressLine3 returns the AddressLine3 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AsyncOrderCreateDTOShipToInfo) GetAddressLine3() string {
	if o == nil || IsNil(o.AddressLine3.Get()) {
		var ret string
		return ret
	}
	return *o.AddressLine3.Get()
}

// GetAddressLine3Ok returns a tuple with the AddressLine3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AsyncOrderCreateDTOShipToInfo) GetAddressLine3Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AddressLine3.Get(), o.AddressLine3.IsSet()
}

// HasAddressLine3 returns a boolean if a field has been set.
func (o *AsyncOrderCreateDTOShipToInfo) HasAddressLine3() bool {
	if o != nil && o.AddressLine3.IsSet() {
		return true
	}

	return false
}

// SetAddressLine3 gets a reference to the given NullableString and assigns it to the AddressLine3 field.
func (o *AsyncOrderCreateDTOShipToInfo) SetAddressLine3(v string) {
	o.AddressLine3.Set(&v)
}
// SetAddressLine3Nil sets the value for AddressLine3 to be an explicit nil
func (o *AsyncOrderCreateDTOShipToInfo) SetAddressLine3Nil() {
	o.AddressLine3.Set(nil)
}

// UnsetAddressLine3 ensures that no value is present for AddressLine3, not even an explicit nil
func (o *AsyncOrderCreateDTOShipToInfo) UnsetAddressLine3() {
	o.AddressLine3.Unset()
}

// GetAddressLine4 returns the AddressLine4 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AsyncOrderCreateDTOShipToInfo) GetAddressLine4() string {
	if o == nil || IsNil(o.AddressLine4.Get()) {
		var ret string
		return ret
	}
	return *o.AddressLine4.Get()
}

// GetAddressLine4Ok returns a tuple with the AddressLine4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AsyncOrderCreateDTOShipToInfo) GetAddressLine4Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AddressLine4.Get(), o.AddressLine4.IsSet()
}

// HasAddressLine4 returns a boolean if a field has been set.
func (o *AsyncOrderCreateDTOShipToInfo) HasAddressLine4() bool {
	if o != nil && o.AddressLine4.IsSet() {
		return true
	}

	return false
}

// SetAddressLine4 gets a reference to the given NullableString and assigns it to the AddressLine4 field.
func (o *AsyncOrderCreateDTOShipToInfo) SetAddressLine4(v string) {
	o.AddressLine4.Set(&v)
}
// SetAddressLine4Nil sets the value for AddressLine4 to be an explicit nil
func (o *AsyncOrderCreateDTOShipToInfo) SetAddressLine4Nil() {
	o.AddressLine4.Set(nil)
}

// UnsetAddressLine4 ensures that no value is present for AddressLine4, not even an explicit nil
func (o *AsyncOrderCreateDTOShipToInfo) UnsetAddressLine4() {
	o.AddressLine4.Unset()
}

// GetName1 returns the Name1 field value if set, zero value otherwise.
func (o *AsyncOrderCreateDTOShipToInfo) GetName1() string {
	if o == nil || IsNil(o.Name1) {
		var ret string
		return ret
	}
	return *o.Name1
}

// GetName1Ok returns a tuple with the Name1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsyncOrderCreateDTOShipToInfo) GetName1Ok() (*string, bool) {
	if o == nil || IsNil(o.Name1) {
		return nil, false
	}
	return o.Name1, true
}

// HasName1 returns a boolean if a field has been set.
func (o *AsyncOrderCreateDTOShipToInfo) HasName1() bool {
	if o != nil && !IsNil(o.Name1) {
		return true
	}

	return false
}

// SetName1 gets a reference to the given string and assigns it to the Name1 field.
func (o *AsyncOrderCreateDTOShipToInfo) SetName1(v string) {
	o.Name1 = &v
}

// GetName2 returns the Name2 field value if set, zero value otherwise.
func (o *AsyncOrderCreateDTOShipToInfo) GetName2() string {
	if o == nil || IsNil(o.Name2) {
		var ret string
		return ret
	}
	return *o.Name2
}

// GetName2Ok returns a tuple with the Name2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsyncOrderCreateDTOShipToInfo) GetName2Ok() (*string, bool) {
	if o == nil || IsNil(o.Name2) {
		return nil, false
	}
	return o.Name2, true
}

// HasName2 returns a boolean if a field has been set.
func (o *AsyncOrderCreateDTOShipToInfo) HasName2() bool {
	if o != nil && !IsNil(o.Name2) {
		return true
	}

	return false
}

// SetName2 gets a reference to the given string and assigns it to the Name2 field.
func (o *AsyncOrderCreateDTOShipToInfo) SetName2(v string) {
	o.Name2 = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *AsyncOrderCreateDTOShipToInfo) GetCity() string {
	if o == nil || IsNil(o.City) {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsyncOrderCreateDTOShipToInfo) GetCityOk() (*string, bool) {
	if o == nil || IsNil(o.City) {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *AsyncOrderCreateDTOShipToInfo) HasCity() bool {
	if o != nil && !IsNil(o.City) {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *AsyncOrderCreateDTOShipToInfo) SetCity(v string) {
	o.City = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *AsyncOrderCreateDTOShipToInfo) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsyncOrderCreateDTOShipToInfo) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *AsyncOrderCreateDTOShipToInfo) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *AsyncOrderCreateDTOShipToInfo) SetState(v string) {
	o.State = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *AsyncOrderCreateDTOShipToInfo) GetPostalCode() string {
	if o == nil || IsNil(o.PostalCode) {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsyncOrderCreateDTOShipToInfo) GetPostalCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PostalCode) {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *AsyncOrderCreateDTOShipToInfo) HasPostalCode() bool {
	if o != nil && !IsNil(o.PostalCode) {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *AsyncOrderCreateDTOShipToInfo) SetPostalCode(v string) {
	o.PostalCode = &v
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *AsyncOrderCreateDTOShipToInfo) GetCountryCode() string {
	if o == nil || IsNil(o.CountryCode) {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsyncOrderCreateDTOShipToInfo) GetCountryCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CountryCode) {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *AsyncOrderCreateDTOShipToInfo) HasCountryCode() bool {
	if o != nil && !IsNil(o.CountryCode) {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *AsyncOrderCreateDTOShipToInfo) SetCountryCode(v string) {
	o.CountryCode = &v
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AsyncOrderCreateDTOShipToInfo) GetEmail() string {
	if o == nil || IsNil(o.Email.Get()) {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AsyncOrderCreateDTOShipToInfo) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *AsyncOrderCreateDTOShipToInfo) HasEmail() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *AsyncOrderCreateDTOShipToInfo) SetEmail(v string) {
	o.Email.Set(&v)
}
// SetEmailNil sets the value for Email to be an explicit nil
func (o *AsyncOrderCreateDTOShipToInfo) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *AsyncOrderCreateDTOShipToInfo) UnsetEmail() {
	o.Email.Unset()
}

// GetShippingNotes returns the ShippingNotes field value if set, zero value otherwise.
func (o *AsyncOrderCreateDTOShipToInfo) GetShippingNotes() string {
	if o == nil || IsNil(o.ShippingNotes) {
		var ret string
		return ret
	}
	return *o.ShippingNotes
}

// GetShippingNotesOk returns a tuple with the ShippingNotes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsyncOrderCreateDTOShipToInfo) GetShippingNotesOk() (*string, bool) {
	if o == nil || IsNil(o.ShippingNotes) {
		return nil, false
	}
	return o.ShippingNotes, true
}

// HasShippingNotes returns a boolean if a field has been set.
func (o *AsyncOrderCreateDTOShipToInfo) HasShippingNotes() bool {
	if o != nil && !IsNil(o.ShippingNotes) {
		return true
	}

	return false
}

// SetShippingNotes gets a reference to the given string and assigns it to the ShippingNotes field.
func (o *AsyncOrderCreateDTOShipToInfo) SetShippingNotes(v string) {
	o.ShippingNotes = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *AsyncOrderCreateDTOShipToInfo) GetPhoneNumber() string {
	if o == nil || IsNil(o.PhoneNumber) {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsyncOrderCreateDTOShipToInfo) GetPhoneNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PhoneNumber) {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *AsyncOrderCreateDTOShipToInfo) HasPhoneNumber() bool {
	if o != nil && !IsNil(o.PhoneNumber) {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *AsyncOrderCreateDTOShipToInfo) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

func (o AsyncOrderCreateDTOShipToInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AsyncOrderCreateDTOShipToInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AddressId) {
		toSerialize["addressId"] = o.AddressId
	}
	if !IsNil(o.Contact) {
		toSerialize["contact"] = o.Contact
	}
	if !IsNil(o.CompanyName) {
		toSerialize["companyName"] = o.CompanyName
	}
	if !IsNil(o.AddressLine1) {
		toSerialize["addressLine1"] = o.AddressLine1
	}
	if !IsNil(o.AddressLine2) {
		toSerialize["addressLine2"] = o.AddressLine2
	}
	if o.AddressLine3.IsSet() {
		toSerialize["addressLine3"] = o.AddressLine3.Get()
	}
	if o.AddressLine4.IsSet() {
		toSerialize["addressLine4"] = o.AddressLine4.Get()
	}
	if !IsNil(o.Name1) {
		toSerialize["name1"] = o.Name1
	}
	if !IsNil(o.Name2) {
		toSerialize["name2"] = o.Name2
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
	if o.Email.IsSet() {
		toSerialize["email"] = o.Email.Get()
	}
	if !IsNil(o.ShippingNotes) {
		toSerialize["shippingNotes"] = o.ShippingNotes
	}
	if !IsNil(o.PhoneNumber) {
		toSerialize["phoneNumber"] = o.PhoneNumber
	}
	return toSerialize, nil
}

type NullableAsyncOrderCreateDTOShipToInfo struct {
	value *AsyncOrderCreateDTOShipToInfo
	isSet bool
}

func (v NullableAsyncOrderCreateDTOShipToInfo) Get() *AsyncOrderCreateDTOShipToInfo {
	return v.value
}

func (v *NullableAsyncOrderCreateDTOShipToInfo) Set(val *AsyncOrderCreateDTOShipToInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableAsyncOrderCreateDTOShipToInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableAsyncOrderCreateDTOShipToInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAsyncOrderCreateDTOShipToInfo(val *AsyncOrderCreateDTOShipToInfo) *NullableAsyncOrderCreateDTOShipToInfo {
	return &NullableAsyncOrderCreateDTOShipToInfo{value: val, isSet: true}
}

func (v NullableAsyncOrderCreateDTOShipToInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAsyncOrderCreateDTOShipToInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

