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

// checks if the QuoteDetailsResponseShippingInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuoteDetailsResponseShippingInfo{}

// QuoteDetailsResponseShippingInfo struct for QuoteDetailsResponseShippingInfo
type QuoteDetailsResponseShippingInfo struct {
	// Contact name  of shipping info associated with the quote.
	CompanyName *string `json:"companyName,omitempty"`
	// Address line 1 for shipping info associated with the quote
	AddressLine1 *string `json:"addressLine1,omitempty"`
	// Address line 2 for shipping info associated with the quote.
	AddressLine2 *string `json:"addressLine2,omitempty"`
	// Address line 3 for shipping info associated with the quote.
	AddressLine3 *string `json:"addressLine3,omitempty"`
	// City for shipping info associated with the quote
	City *string `json:"city,omitempty"`
	// Two letter state abreviation for shipping info associated with the quote
	State *string `json:"state,omitempty"`
	// Email of shipping info the quote associated with the quote.
	Email *string `json:"email,omitempty"`
	// Phone number of shipping info associated with the quote.
	PhoneNumber *string `json:"phoneNumber,omitempty"`
	// Zip code of shipping info associated with the quote.
	PostalCode *string `json:"postalCode,omitempty"`
	ShpToGstinNumber *string `json:"shpToGstinNumber,omitempty"`
}

// NewQuoteDetailsResponseShippingInfo instantiates a new QuoteDetailsResponseShippingInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuoteDetailsResponseShippingInfo() *QuoteDetailsResponseShippingInfo {
	this := QuoteDetailsResponseShippingInfo{}
	return &this
}

// NewQuoteDetailsResponseShippingInfoWithDefaults instantiates a new QuoteDetailsResponseShippingInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuoteDetailsResponseShippingInfoWithDefaults() *QuoteDetailsResponseShippingInfo {
	this := QuoteDetailsResponseShippingInfo{}
	return &this
}

// GetCompanyName returns the CompanyName field value if set, zero value otherwise.
func (o *QuoteDetailsResponseShippingInfo) GetCompanyName() string {
	if o == nil || IsNil(o.CompanyName) {
		var ret string
		return ret
	}
	return *o.CompanyName
}

// GetCompanyNameOk returns a tuple with the CompanyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteDetailsResponseShippingInfo) GetCompanyNameOk() (*string, bool) {
	if o == nil || IsNil(o.CompanyName) {
		return nil, false
	}
	return o.CompanyName, true
}

// HasCompanyName returns a boolean if a field has been set.
func (o *QuoteDetailsResponseShippingInfo) HasCompanyName() bool {
	if o != nil && !IsNil(o.CompanyName) {
		return true
	}

	return false
}

// SetCompanyName gets a reference to the given string and assigns it to the CompanyName field.
func (o *QuoteDetailsResponseShippingInfo) SetCompanyName(v string) {
	o.CompanyName = &v
}

// GetAddressLine1 returns the AddressLine1 field value if set, zero value otherwise.
func (o *QuoteDetailsResponseShippingInfo) GetAddressLine1() string {
	if o == nil || IsNil(o.AddressLine1) {
		var ret string
		return ret
	}
	return *o.AddressLine1
}

// GetAddressLine1Ok returns a tuple with the AddressLine1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteDetailsResponseShippingInfo) GetAddressLine1Ok() (*string, bool) {
	if o == nil || IsNil(o.AddressLine1) {
		return nil, false
	}
	return o.AddressLine1, true
}

// HasAddressLine1 returns a boolean if a field has been set.
func (o *QuoteDetailsResponseShippingInfo) HasAddressLine1() bool {
	if o != nil && !IsNil(o.AddressLine1) {
		return true
	}

	return false
}

// SetAddressLine1 gets a reference to the given string and assigns it to the AddressLine1 field.
func (o *QuoteDetailsResponseShippingInfo) SetAddressLine1(v string) {
	o.AddressLine1 = &v
}

// GetAddressLine2 returns the AddressLine2 field value if set, zero value otherwise.
func (o *QuoteDetailsResponseShippingInfo) GetAddressLine2() string {
	if o == nil || IsNil(o.AddressLine2) {
		var ret string
		return ret
	}
	return *o.AddressLine2
}

// GetAddressLine2Ok returns a tuple with the AddressLine2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteDetailsResponseShippingInfo) GetAddressLine2Ok() (*string, bool) {
	if o == nil || IsNil(o.AddressLine2) {
		return nil, false
	}
	return o.AddressLine2, true
}

// HasAddressLine2 returns a boolean if a field has been set.
func (o *QuoteDetailsResponseShippingInfo) HasAddressLine2() bool {
	if o != nil && !IsNil(o.AddressLine2) {
		return true
	}

	return false
}

// SetAddressLine2 gets a reference to the given string and assigns it to the AddressLine2 field.
func (o *QuoteDetailsResponseShippingInfo) SetAddressLine2(v string) {
	o.AddressLine2 = &v
}

// GetAddressLine3 returns the AddressLine3 field value if set, zero value otherwise.
func (o *QuoteDetailsResponseShippingInfo) GetAddressLine3() string {
	if o == nil || IsNil(o.AddressLine3) {
		var ret string
		return ret
	}
	return *o.AddressLine3
}

// GetAddressLine3Ok returns a tuple with the AddressLine3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteDetailsResponseShippingInfo) GetAddressLine3Ok() (*string, bool) {
	if o == nil || IsNil(o.AddressLine3) {
		return nil, false
	}
	return o.AddressLine3, true
}

// HasAddressLine3 returns a boolean if a field has been set.
func (o *QuoteDetailsResponseShippingInfo) HasAddressLine3() bool {
	if o != nil && !IsNil(o.AddressLine3) {
		return true
	}

	return false
}

// SetAddressLine3 gets a reference to the given string and assigns it to the AddressLine3 field.
func (o *QuoteDetailsResponseShippingInfo) SetAddressLine3(v string) {
	o.AddressLine3 = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *QuoteDetailsResponseShippingInfo) GetCity() string {
	if o == nil || IsNil(o.City) {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteDetailsResponseShippingInfo) GetCityOk() (*string, bool) {
	if o == nil || IsNil(o.City) {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *QuoteDetailsResponseShippingInfo) HasCity() bool {
	if o != nil && !IsNil(o.City) {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *QuoteDetailsResponseShippingInfo) SetCity(v string) {
	o.City = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *QuoteDetailsResponseShippingInfo) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteDetailsResponseShippingInfo) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *QuoteDetailsResponseShippingInfo) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *QuoteDetailsResponseShippingInfo) SetState(v string) {
	o.State = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *QuoteDetailsResponseShippingInfo) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteDetailsResponseShippingInfo) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *QuoteDetailsResponseShippingInfo) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *QuoteDetailsResponseShippingInfo) SetEmail(v string) {
	o.Email = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *QuoteDetailsResponseShippingInfo) GetPhoneNumber() string {
	if o == nil || IsNil(o.PhoneNumber) {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteDetailsResponseShippingInfo) GetPhoneNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PhoneNumber) {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *QuoteDetailsResponseShippingInfo) HasPhoneNumber() bool {
	if o != nil && !IsNil(o.PhoneNumber) {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *QuoteDetailsResponseShippingInfo) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *QuoteDetailsResponseShippingInfo) GetPostalCode() string {
	if o == nil || IsNil(o.PostalCode) {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteDetailsResponseShippingInfo) GetPostalCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PostalCode) {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *QuoteDetailsResponseShippingInfo) HasPostalCode() bool {
	if o != nil && !IsNil(o.PostalCode) {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *QuoteDetailsResponseShippingInfo) SetPostalCode(v string) {
	o.PostalCode = &v
}

// GetShpToGstinNumber returns the ShpToGstinNumber field value if set, zero value otherwise.
func (o *QuoteDetailsResponseShippingInfo) GetShpToGstinNumber() string {
	if o == nil || IsNil(o.ShpToGstinNumber) {
		var ret string
		return ret
	}
	return *o.ShpToGstinNumber
}

// GetShpToGstinNumberOk returns a tuple with the ShpToGstinNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteDetailsResponseShippingInfo) GetShpToGstinNumberOk() (*string, bool) {
	if o == nil || IsNil(o.ShpToGstinNumber) {
		return nil, false
	}
	return o.ShpToGstinNumber, true
}

// HasShpToGstinNumber returns a boolean if a field has been set.
func (o *QuoteDetailsResponseShippingInfo) HasShpToGstinNumber() bool {
	if o != nil && !IsNil(o.ShpToGstinNumber) {
		return true
	}

	return false
}

// SetShpToGstinNumber gets a reference to the given string and assigns it to the ShpToGstinNumber field.
func (o *QuoteDetailsResponseShippingInfo) SetShpToGstinNumber(v string) {
	o.ShpToGstinNumber = &v
}

func (o QuoteDetailsResponseShippingInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuoteDetailsResponseShippingInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CompanyName) {
		toSerialize["companyName"] = o.CompanyName
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
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.PhoneNumber) {
		toSerialize["phoneNumber"] = o.PhoneNumber
	}
	if !IsNil(o.PostalCode) {
		toSerialize["postalCode"] = o.PostalCode
	}
	if !IsNil(o.ShpToGstinNumber) {
		toSerialize["shpToGstinNumber"] = o.ShpToGstinNumber
	}
	return toSerialize, nil
}

type NullableQuoteDetailsResponseShippingInfo struct {
	value *QuoteDetailsResponseShippingInfo
	isSet bool
}

func (v NullableQuoteDetailsResponseShippingInfo) Get() *QuoteDetailsResponseShippingInfo {
	return v.value
}

func (v *NullableQuoteDetailsResponseShippingInfo) Set(val *QuoteDetailsResponseShippingInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableQuoteDetailsResponseShippingInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableQuoteDetailsResponseShippingInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuoteDetailsResponseShippingInfo(val *QuoteDetailsResponseShippingInfo) *NullableQuoteDetailsResponseShippingInfo {
	return &NullableQuoteDetailsResponseShippingInfo{value: val, isSet: true}
}

func (v NullableQuoteDetailsResponseShippingInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuoteDetailsResponseShippingInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


