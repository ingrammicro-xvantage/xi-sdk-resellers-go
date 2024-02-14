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

// checks if the QuoteToOrderDetailsDTOShipToInfoInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuoteToOrderDetailsDTOShipToInfoInner{}

// QuoteToOrderDetailsDTOShipToInfoInner struct for QuoteToOrderDetailsDTOShipToInfoInner
type QuoteToOrderDetailsDTOShipToInfoInner struct {
	// The company contact provided by the reseller.
	AddressId *string `json:"addressId,omitempty"`
	// The name of the company the order will be shipped to.
	CompanyName *string `json:"companyName,omitempty"`
	// The contact name for the order will be shipped to.
	Contact *string `json:"contact,omitempty"`
	// The address line 1 the order will be shipped to.
	AddressLine1 *string `json:"addressLine1,omitempty"`
	// The address line 2 the order will be shipped to.
	AddressLine2 *string `json:"addressLine2,omitempty"`
	// The address line 3 the order will be shipped to.
	AddressLine3 NullableString `json:"addressLine3,omitempty"`
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
}

// NewQuoteToOrderDetailsDTOShipToInfoInner instantiates a new QuoteToOrderDetailsDTOShipToInfoInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuoteToOrderDetailsDTOShipToInfoInner() *QuoteToOrderDetailsDTOShipToInfoInner {
	this := QuoteToOrderDetailsDTOShipToInfoInner{}
	return &this
}

// NewQuoteToOrderDetailsDTOShipToInfoInnerWithDefaults instantiates a new QuoteToOrderDetailsDTOShipToInfoInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuoteToOrderDetailsDTOShipToInfoInnerWithDefaults() *QuoteToOrderDetailsDTOShipToInfoInner {
	this := QuoteToOrderDetailsDTOShipToInfoInner{}
	return &this
}

// GetAddressId returns the AddressId field value if set, zero value otherwise.
func (o *QuoteToOrderDetailsDTOShipToInfoInner) GetAddressId() string {
	if o == nil || IsNil(o.AddressId) {
		var ret string
		return ret
	}
	return *o.AddressId
}

// GetAddressIdOk returns a tuple with the AddressId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteToOrderDetailsDTOShipToInfoInner) GetAddressIdOk() (*string, bool) {
	if o == nil || IsNil(o.AddressId) {
		return nil, false
	}
	return o.AddressId, true
}

// HasAddressId returns a boolean if a field has been set.
func (o *QuoteToOrderDetailsDTOShipToInfoInner) HasAddressId() bool {
	if o != nil && !IsNil(o.AddressId) {
		return true
	}

	return false
}

// SetAddressId gets a reference to the given string and assigns it to the AddressId field.
func (o *QuoteToOrderDetailsDTOShipToInfoInner) SetAddressId(v string) {
	o.AddressId = &v
}

// GetCompanyName returns the CompanyName field value if set, zero value otherwise.
func (o *QuoteToOrderDetailsDTOShipToInfoInner) GetCompanyName() string {
	if o == nil || IsNil(o.CompanyName) {
		var ret string
		return ret
	}
	return *o.CompanyName
}

// GetCompanyNameOk returns a tuple with the CompanyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteToOrderDetailsDTOShipToInfoInner) GetCompanyNameOk() (*string, bool) {
	if o == nil || IsNil(o.CompanyName) {
		return nil, false
	}
	return o.CompanyName, true
}

// HasCompanyName returns a boolean if a field has been set.
func (o *QuoteToOrderDetailsDTOShipToInfoInner) HasCompanyName() bool {
	if o != nil && !IsNil(o.CompanyName) {
		return true
	}

	return false
}

// SetCompanyName gets a reference to the given string and assigns it to the CompanyName field.
func (o *QuoteToOrderDetailsDTOShipToInfoInner) SetCompanyName(v string) {
	o.CompanyName = &v
}

// GetContact returns the Contact field value if set, zero value otherwise.
func (o *QuoteToOrderDetailsDTOShipToInfoInner) GetContact() string {
	if o == nil || IsNil(o.Contact) {
		var ret string
		return ret
	}
	return *o.Contact
}

// GetContactOk returns a tuple with the Contact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteToOrderDetailsDTOShipToInfoInner) GetContactOk() (*string, bool) {
	if o == nil || IsNil(o.Contact) {
		return nil, false
	}
	return o.Contact, true
}

// HasContact returns a boolean if a field has been set.
func (o *QuoteToOrderDetailsDTOShipToInfoInner) HasContact() bool {
	if o != nil && !IsNil(o.Contact) {
		return true
	}

	return false
}

// SetContact gets a reference to the given string and assigns it to the Contact field.
func (o *QuoteToOrderDetailsDTOShipToInfoInner) SetContact(v string) {
	o.Contact = &v
}

// GetAddressLine1 returns the AddressLine1 field value if set, zero value otherwise.
func (o *QuoteToOrderDetailsDTOShipToInfoInner) GetAddressLine1() string {
	if o == nil || IsNil(o.AddressLine1) {
		var ret string
		return ret
	}
	return *o.AddressLine1
}

// GetAddressLine1Ok returns a tuple with the AddressLine1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteToOrderDetailsDTOShipToInfoInner) GetAddressLine1Ok() (*string, bool) {
	if o == nil || IsNil(o.AddressLine1) {
		return nil, false
	}
	return o.AddressLine1, true
}

// HasAddressLine1 returns a boolean if a field has been set.
func (o *QuoteToOrderDetailsDTOShipToInfoInner) HasAddressLine1() bool {
	if o != nil && !IsNil(o.AddressLine1) {
		return true
	}

	return false
}

// SetAddressLine1 gets a reference to the given string and assigns it to the AddressLine1 field.
func (o *QuoteToOrderDetailsDTOShipToInfoInner) SetAddressLine1(v string) {
	o.AddressLine1 = &v
}

// GetAddressLine2 returns the AddressLine2 field value if set, zero value otherwise.
func (o *QuoteToOrderDetailsDTOShipToInfoInner) GetAddressLine2() string {
	if o == nil || IsNil(o.AddressLine2) {
		var ret string
		return ret
	}
	return *o.AddressLine2
}

// GetAddressLine2Ok returns a tuple with the AddressLine2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteToOrderDetailsDTOShipToInfoInner) GetAddressLine2Ok() (*string, bool) {
	if o == nil || IsNil(o.AddressLine2) {
		return nil, false
	}
	return o.AddressLine2, true
}

// HasAddressLine2 returns a boolean if a field has been set.
func (o *QuoteToOrderDetailsDTOShipToInfoInner) HasAddressLine2() bool {
	if o != nil && !IsNil(o.AddressLine2) {
		return true
	}

	return false
}

// SetAddressLine2 gets a reference to the given string and assigns it to the AddressLine2 field.
func (o *QuoteToOrderDetailsDTOShipToInfoInner) SetAddressLine2(v string) {
	o.AddressLine2 = &v
}

// GetAddressLine3 returns the AddressLine3 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QuoteToOrderDetailsDTOShipToInfoInner) GetAddressLine3() string {
	if o == nil || IsNil(o.AddressLine3.Get()) {
		var ret string
		return ret
	}
	return *o.AddressLine3.Get()
}

// GetAddressLine3Ok returns a tuple with the AddressLine3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QuoteToOrderDetailsDTOShipToInfoInner) GetAddressLine3Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AddressLine3.Get(), o.AddressLine3.IsSet()
}

// HasAddressLine3 returns a boolean if a field has been set.
func (o *QuoteToOrderDetailsDTOShipToInfoInner) HasAddressLine3() bool {
	if o != nil && o.AddressLine3.IsSet() {
		return true
	}

	return false
}

// SetAddressLine3 gets a reference to the given NullableString and assigns it to the AddressLine3 field.
func (o *QuoteToOrderDetailsDTOShipToInfoInner) SetAddressLine3(v string) {
	o.AddressLine3.Set(&v)
}
// SetAddressLine3Nil sets the value for AddressLine3 to be an explicit nil
func (o *QuoteToOrderDetailsDTOShipToInfoInner) SetAddressLine3Nil() {
	o.AddressLine3.Set(nil)
}

// UnsetAddressLine3 ensures that no value is present for AddressLine3, not even an explicit nil
func (o *QuoteToOrderDetailsDTOShipToInfoInner) UnsetAddressLine3() {
	o.AddressLine3.Unset()
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *QuoteToOrderDetailsDTOShipToInfoInner) GetCity() string {
	if o == nil || IsNil(o.City) {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteToOrderDetailsDTOShipToInfoInner) GetCityOk() (*string, bool) {
	if o == nil || IsNil(o.City) {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *QuoteToOrderDetailsDTOShipToInfoInner) HasCity() bool {
	if o != nil && !IsNil(o.City) {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *QuoteToOrderDetailsDTOShipToInfoInner) SetCity(v string) {
	o.City = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *QuoteToOrderDetailsDTOShipToInfoInner) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteToOrderDetailsDTOShipToInfoInner) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *QuoteToOrderDetailsDTOShipToInfoInner) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *QuoteToOrderDetailsDTOShipToInfoInner) SetState(v string) {
	o.State = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *QuoteToOrderDetailsDTOShipToInfoInner) GetPostalCode() string {
	if o == nil || IsNil(o.PostalCode) {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteToOrderDetailsDTOShipToInfoInner) GetPostalCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PostalCode) {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *QuoteToOrderDetailsDTOShipToInfoInner) HasPostalCode() bool {
	if o != nil && !IsNil(o.PostalCode) {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *QuoteToOrderDetailsDTOShipToInfoInner) SetPostalCode(v string) {
	o.PostalCode = &v
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *QuoteToOrderDetailsDTOShipToInfoInner) GetCountryCode() string {
	if o == nil || IsNil(o.CountryCode) {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteToOrderDetailsDTOShipToInfoInner) GetCountryCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CountryCode) {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *QuoteToOrderDetailsDTOShipToInfoInner) HasCountryCode() bool {
	if o != nil && !IsNil(o.CountryCode) {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *QuoteToOrderDetailsDTOShipToInfoInner) SetCountryCode(v string) {
	o.CountryCode = &v
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QuoteToOrderDetailsDTOShipToInfoInner) GetEmail() string {
	if o == nil || IsNil(o.Email.Get()) {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QuoteToOrderDetailsDTOShipToInfoInner) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *QuoteToOrderDetailsDTOShipToInfoInner) HasEmail() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *QuoteToOrderDetailsDTOShipToInfoInner) SetEmail(v string) {
	o.Email.Set(&v)
}
// SetEmailNil sets the value for Email to be an explicit nil
func (o *QuoteToOrderDetailsDTOShipToInfoInner) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *QuoteToOrderDetailsDTOShipToInfoInner) UnsetEmail() {
	o.Email.Unset()
}

func (o QuoteToOrderDetailsDTOShipToInfoInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuoteToOrderDetailsDTOShipToInfoInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AddressId) {
		toSerialize["addressId"] = o.AddressId
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
	if o.AddressLine3.IsSet() {
		toSerialize["addressLine3"] = o.AddressLine3.Get()
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
	return toSerialize, nil
}

type NullableQuoteToOrderDetailsDTOShipToInfoInner struct {
	value *QuoteToOrderDetailsDTOShipToInfoInner
	isSet bool
}

func (v NullableQuoteToOrderDetailsDTOShipToInfoInner) Get() *QuoteToOrderDetailsDTOShipToInfoInner {
	return v.value
}

func (v *NullableQuoteToOrderDetailsDTOShipToInfoInner) Set(val *QuoteToOrderDetailsDTOShipToInfoInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQuoteToOrderDetailsDTOShipToInfoInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQuoteToOrderDetailsDTOShipToInfoInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuoteToOrderDetailsDTOShipToInfoInner(val *QuoteToOrderDetailsDTOShipToInfoInner) *NullableQuoteToOrderDetailsDTOShipToInfoInner {
	return &NullableQuoteToOrderDetailsDTOShipToInfoInner{value: val, isSet: true}
}

func (v NullableQuoteToOrderDetailsDTOShipToInfoInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuoteToOrderDetailsDTOShipToInfoInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


