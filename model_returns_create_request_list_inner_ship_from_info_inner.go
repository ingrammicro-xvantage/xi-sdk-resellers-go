/*
XI Sdk Resellers

For resellers seeking to innovate with Ingram Micro's API solutions, automate your eCommerce experience with our array of API's and webhooks to craft a seamless journey for your customers.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package xi_sdk_resellers

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ReturnsCreateRequestListInnerShipFromInfoInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReturnsCreateRequestListInnerShipFromInfoInner{}

// ReturnsCreateRequestListInnerShipFromInfoInner struct for ReturnsCreateRequestListInnerShipFromInfoInner
type ReturnsCreateRequestListInnerShipFromInfoInner struct {
	// Name of the company from where the product will be shipped.
	CompanyName string `json:"companyName"`
	// Contact name of the person from where the product will be shipped.
	Contact string `json:"contact"`
	// Ship from Address Line1.
	AddressLine1 string `json:"addressLine1"`
	// Ship from Address Line2.
	AddressLine2 *string `json:"addressLine2,omitempty"`
	// Ship from Address Line3.
	AddressLine3 *string `json:"addressLine3,omitempty"`
	// Ship from City.
	City string `json:"city"`
	// Ship from state.
	State string `json:"state"`
	// Ship from postal code.
	PostalCode string `json:"postalCode"`
	// ship from country code.
	CountryCode string `json:"countryCode"`
	// Ship from email.
	Email string `json:"email"`
	// Ship from phone number.
	PhoneNumber *string `json:"phoneNumber,omitempty"`
}

type _ReturnsCreateRequestListInnerShipFromInfoInner ReturnsCreateRequestListInnerShipFromInfoInner

// NewReturnsCreateRequestListInnerShipFromInfoInner instantiates a new ReturnsCreateRequestListInnerShipFromInfoInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReturnsCreateRequestListInnerShipFromInfoInner(companyName string, contact string, addressLine1 string, city string, state string, postalCode string, countryCode string, email string) *ReturnsCreateRequestListInnerShipFromInfoInner {
	this := ReturnsCreateRequestListInnerShipFromInfoInner{}
	this.CompanyName = companyName
	this.Contact = contact
	this.AddressLine1 = addressLine1
	this.City = city
	this.State = state
	this.PostalCode = postalCode
	this.CountryCode = countryCode
	this.Email = email
	return &this
}

// NewReturnsCreateRequestListInnerShipFromInfoInnerWithDefaults instantiates a new ReturnsCreateRequestListInnerShipFromInfoInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReturnsCreateRequestListInnerShipFromInfoInnerWithDefaults() *ReturnsCreateRequestListInnerShipFromInfoInner {
	this := ReturnsCreateRequestListInnerShipFromInfoInner{}
	return &this
}

// GetCompanyName returns the CompanyName field value
func (o *ReturnsCreateRequestListInnerShipFromInfoInner) GetCompanyName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CompanyName
}

// GetCompanyNameOk returns a tuple with the CompanyName field value
// and a boolean to check if the value has been set.
func (o *ReturnsCreateRequestListInnerShipFromInfoInner) GetCompanyNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CompanyName, true
}

// SetCompanyName sets field value
func (o *ReturnsCreateRequestListInnerShipFromInfoInner) SetCompanyName(v string) {
	o.CompanyName = v
}

// GetContact returns the Contact field value
func (o *ReturnsCreateRequestListInnerShipFromInfoInner) GetContact() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Contact
}

// GetContactOk returns a tuple with the Contact field value
// and a boolean to check if the value has been set.
func (o *ReturnsCreateRequestListInnerShipFromInfoInner) GetContactOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Contact, true
}

// SetContact sets field value
func (o *ReturnsCreateRequestListInnerShipFromInfoInner) SetContact(v string) {
	o.Contact = v
}

// GetAddressLine1 returns the AddressLine1 field value
func (o *ReturnsCreateRequestListInnerShipFromInfoInner) GetAddressLine1() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AddressLine1
}

// GetAddressLine1Ok returns a tuple with the AddressLine1 field value
// and a boolean to check if the value has been set.
func (o *ReturnsCreateRequestListInnerShipFromInfoInner) GetAddressLine1Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AddressLine1, true
}

// SetAddressLine1 sets field value
func (o *ReturnsCreateRequestListInnerShipFromInfoInner) SetAddressLine1(v string) {
	o.AddressLine1 = v
}

// GetAddressLine2 returns the AddressLine2 field value if set, zero value otherwise.
func (o *ReturnsCreateRequestListInnerShipFromInfoInner) GetAddressLine2() string {
	if o == nil || IsNil(o.AddressLine2) {
		var ret string
		return ret
	}
	return *o.AddressLine2
}

// GetAddressLine2Ok returns a tuple with the AddressLine2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnsCreateRequestListInnerShipFromInfoInner) GetAddressLine2Ok() (*string, bool) {
	if o == nil || IsNil(o.AddressLine2) {
		return nil, false
	}
	return o.AddressLine2, true
}

// HasAddressLine2 returns a boolean if a field has been set.
func (o *ReturnsCreateRequestListInnerShipFromInfoInner) HasAddressLine2() bool {
	if o != nil && !IsNil(o.AddressLine2) {
		return true
	}

	return false
}

// SetAddressLine2 gets a reference to the given string and assigns it to the AddressLine2 field.
func (o *ReturnsCreateRequestListInnerShipFromInfoInner) SetAddressLine2(v string) {
	o.AddressLine2 = &v
}

// GetAddressLine3 returns the AddressLine3 field value if set, zero value otherwise.
func (o *ReturnsCreateRequestListInnerShipFromInfoInner) GetAddressLine3() string {
	if o == nil || IsNil(o.AddressLine3) {
		var ret string
		return ret
	}
	return *o.AddressLine3
}

// GetAddressLine3Ok returns a tuple with the AddressLine3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnsCreateRequestListInnerShipFromInfoInner) GetAddressLine3Ok() (*string, bool) {
	if o == nil || IsNil(o.AddressLine3) {
		return nil, false
	}
	return o.AddressLine3, true
}

// HasAddressLine3 returns a boolean if a field has been set.
func (o *ReturnsCreateRequestListInnerShipFromInfoInner) HasAddressLine3() bool {
	if o != nil && !IsNil(o.AddressLine3) {
		return true
	}

	return false
}

// SetAddressLine3 gets a reference to the given string and assigns it to the AddressLine3 field.
func (o *ReturnsCreateRequestListInnerShipFromInfoInner) SetAddressLine3(v string) {
	o.AddressLine3 = &v
}

// GetCity returns the City field value
func (o *ReturnsCreateRequestListInnerShipFromInfoInner) GetCity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.City
}

// GetCityOk returns a tuple with the City field value
// and a boolean to check if the value has been set.
func (o *ReturnsCreateRequestListInnerShipFromInfoInner) GetCityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.City, true
}

// SetCity sets field value
func (o *ReturnsCreateRequestListInnerShipFromInfoInner) SetCity(v string) {
	o.City = v
}

// GetState returns the State field value
func (o *ReturnsCreateRequestListInnerShipFromInfoInner) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *ReturnsCreateRequestListInnerShipFromInfoInner) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *ReturnsCreateRequestListInnerShipFromInfoInner) SetState(v string) {
	o.State = v
}

// GetPostalCode returns the PostalCode field value
func (o *ReturnsCreateRequestListInnerShipFromInfoInner) GetPostalCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value
// and a boolean to check if the value has been set.
func (o *ReturnsCreateRequestListInnerShipFromInfoInner) GetPostalCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PostalCode, true
}

// SetPostalCode sets field value
func (o *ReturnsCreateRequestListInnerShipFromInfoInner) SetPostalCode(v string) {
	o.PostalCode = v
}

// GetCountryCode returns the CountryCode field value
func (o *ReturnsCreateRequestListInnerShipFromInfoInner) GetCountryCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value
// and a boolean to check if the value has been set.
func (o *ReturnsCreateRequestListInnerShipFromInfoInner) GetCountryCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CountryCode, true
}

// SetCountryCode sets field value
func (o *ReturnsCreateRequestListInnerShipFromInfoInner) SetCountryCode(v string) {
	o.CountryCode = v
}

// GetEmail returns the Email field value
func (o *ReturnsCreateRequestListInnerShipFromInfoInner) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *ReturnsCreateRequestListInnerShipFromInfoInner) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *ReturnsCreateRequestListInnerShipFromInfoInner) SetEmail(v string) {
	o.Email = v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *ReturnsCreateRequestListInnerShipFromInfoInner) GetPhoneNumber() string {
	if o == nil || IsNil(o.PhoneNumber) {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnsCreateRequestListInnerShipFromInfoInner) GetPhoneNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PhoneNumber) {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *ReturnsCreateRequestListInnerShipFromInfoInner) HasPhoneNumber() bool {
	if o != nil && !IsNil(o.PhoneNumber) {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *ReturnsCreateRequestListInnerShipFromInfoInner) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

func (o ReturnsCreateRequestListInnerShipFromInfoInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReturnsCreateRequestListInnerShipFromInfoInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["companyName"] = o.CompanyName
	toSerialize["contact"] = o.Contact
	toSerialize["addressLine1"] = o.AddressLine1
	if !IsNil(o.AddressLine2) {
		toSerialize["addressLine2"] = o.AddressLine2
	}
	if !IsNil(o.AddressLine3) {
		toSerialize["addressLine3"] = o.AddressLine3
	}
	toSerialize["city"] = o.City
	toSerialize["state"] = o.State
	toSerialize["postalCode"] = o.PostalCode
	toSerialize["countryCode"] = o.CountryCode
	toSerialize["email"] = o.Email
	if !IsNil(o.PhoneNumber) {
		toSerialize["phoneNumber"] = o.PhoneNumber
	}
	return toSerialize, nil
}

func (o *ReturnsCreateRequestListInnerShipFromInfoInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"companyName",
		"contact",
		"addressLine1",
		"city",
		"state",
		"postalCode",
		"countryCode",
		"email",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varReturnsCreateRequestListInnerShipFromInfoInner := _ReturnsCreateRequestListInnerShipFromInfoInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varReturnsCreateRequestListInnerShipFromInfoInner)

	if err != nil {
		return err
	}

	*o = ReturnsCreateRequestListInnerShipFromInfoInner(varReturnsCreateRequestListInnerShipFromInfoInner)

	return err
}

type NullableReturnsCreateRequestListInnerShipFromInfoInner struct {
	value *ReturnsCreateRequestListInnerShipFromInfoInner
	isSet bool
}

func (v NullableReturnsCreateRequestListInnerShipFromInfoInner) Get() *ReturnsCreateRequestListInnerShipFromInfoInner {
	return v.value
}

func (v *NullableReturnsCreateRequestListInnerShipFromInfoInner) Set(val *ReturnsCreateRequestListInnerShipFromInfoInner) {
	v.value = val
	v.isSet = true
}

func (v NullableReturnsCreateRequestListInnerShipFromInfoInner) IsSet() bool {
	return v.isSet
}

func (v *NullableReturnsCreateRequestListInnerShipFromInfoInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReturnsCreateRequestListInnerShipFromInfoInner(val *ReturnsCreateRequestListInnerShipFromInfoInner) *NullableReturnsCreateRequestListInnerShipFromInfoInner {
	return &NullableReturnsCreateRequestListInnerShipFromInfoInner{value: val, isSet: true}
}

func (v NullableReturnsCreateRequestListInnerShipFromInfoInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReturnsCreateRequestListInnerShipFromInfoInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


