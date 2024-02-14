/*
XI Sdk Resellers

For Resellers. Who are looking to Innovate with Ingram Micro's API SolutionsAutomate your eCommerce with our offering of APIs and Webhooks to create a seamless experience for your customers.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package xi_sdk_resellers

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the QuoteListRequestQuoteSearchRequestRequestPreamble type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuoteListRequestQuoteSearchRequestRequestPreamble{}

// QuoteListRequestQuoteSearchRequestRequestPreamble struct for QuoteListRequestQuoteSearchRequestRequestPreamble
type QuoteListRequestQuoteSearchRequestRequestPreamble struct {
	// Reseller Number (referred to as the account BCN) is the unique identifier for an Ingram Micro customer account.
	CustomerNumber string `json:"customerNumber"`
	// Logged in User's email address.
	CustomerContact *string `json:"customerContact,omitempty"`
	// The ISO country codes are internationally recognized codes designated for each country represented by a two-letter combination (alpha-2).
	IsoCountryCode string `json:"isoCountryCode"`
}

type _QuoteListRequestQuoteSearchRequestRequestPreamble QuoteListRequestQuoteSearchRequestRequestPreamble

// NewQuoteListRequestQuoteSearchRequestRequestPreamble instantiates a new QuoteListRequestQuoteSearchRequestRequestPreamble object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuoteListRequestQuoteSearchRequestRequestPreamble(customerNumber string, isoCountryCode string) *QuoteListRequestQuoteSearchRequestRequestPreamble {
	this := QuoteListRequestQuoteSearchRequestRequestPreamble{}
	this.CustomerNumber = customerNumber
	this.IsoCountryCode = isoCountryCode
	return &this
}

// NewQuoteListRequestQuoteSearchRequestRequestPreambleWithDefaults instantiates a new QuoteListRequestQuoteSearchRequestRequestPreamble object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuoteListRequestQuoteSearchRequestRequestPreambleWithDefaults() *QuoteListRequestQuoteSearchRequestRequestPreamble {
	this := QuoteListRequestQuoteSearchRequestRequestPreamble{}
	return &this
}

// GetCustomerNumber returns the CustomerNumber field value
func (o *QuoteListRequestQuoteSearchRequestRequestPreamble) GetCustomerNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerNumber
}

// GetCustomerNumberOk returns a tuple with the CustomerNumber field value
// and a boolean to check if the value has been set.
func (o *QuoteListRequestQuoteSearchRequestRequestPreamble) GetCustomerNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerNumber, true
}

// SetCustomerNumber sets field value
func (o *QuoteListRequestQuoteSearchRequestRequestPreamble) SetCustomerNumber(v string) {
	o.CustomerNumber = v
}

// GetCustomerContact returns the CustomerContact field value if set, zero value otherwise.
func (o *QuoteListRequestQuoteSearchRequestRequestPreamble) GetCustomerContact() string {
	if o == nil || IsNil(o.CustomerContact) {
		var ret string
		return ret
	}
	return *o.CustomerContact
}

// GetCustomerContactOk returns a tuple with the CustomerContact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteListRequestQuoteSearchRequestRequestPreamble) GetCustomerContactOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerContact) {
		return nil, false
	}
	return o.CustomerContact, true
}

// HasCustomerContact returns a boolean if a field has been set.
func (o *QuoteListRequestQuoteSearchRequestRequestPreamble) HasCustomerContact() bool {
	if o != nil && !IsNil(o.CustomerContact) {
		return true
	}

	return false
}

// SetCustomerContact gets a reference to the given string and assigns it to the CustomerContact field.
func (o *QuoteListRequestQuoteSearchRequestRequestPreamble) SetCustomerContact(v string) {
	o.CustomerContact = &v
}

// GetIsoCountryCode returns the IsoCountryCode field value
func (o *QuoteListRequestQuoteSearchRequestRequestPreamble) GetIsoCountryCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IsoCountryCode
}

// GetIsoCountryCodeOk returns a tuple with the IsoCountryCode field value
// and a boolean to check if the value has been set.
func (o *QuoteListRequestQuoteSearchRequestRequestPreamble) GetIsoCountryCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsoCountryCode, true
}

// SetIsoCountryCode sets field value
func (o *QuoteListRequestQuoteSearchRequestRequestPreamble) SetIsoCountryCode(v string) {
	o.IsoCountryCode = v
}

func (o QuoteListRequestQuoteSearchRequestRequestPreamble) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuoteListRequestQuoteSearchRequestRequestPreamble) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["customerNumber"] = o.CustomerNumber
	if !IsNil(o.CustomerContact) {
		toSerialize["customerContact"] = o.CustomerContact
	}
	toSerialize["isoCountryCode"] = o.IsoCountryCode
	return toSerialize, nil
}

func (o *QuoteListRequestQuoteSearchRequestRequestPreamble) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"customerNumber",
		"isoCountryCode",
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

	varQuoteListRequestQuoteSearchRequestRequestPreamble := _QuoteListRequestQuoteSearchRequestRequestPreamble{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varQuoteListRequestQuoteSearchRequestRequestPreamble)

	if err != nil {
		return err
	}

	*o = QuoteListRequestQuoteSearchRequestRequestPreamble(varQuoteListRequestQuoteSearchRequestRequestPreamble)

	return err
}

type NullableQuoteListRequestQuoteSearchRequestRequestPreamble struct {
	value *QuoteListRequestQuoteSearchRequestRequestPreamble
	isSet bool
}

func (v NullableQuoteListRequestQuoteSearchRequestRequestPreamble) Get() *QuoteListRequestQuoteSearchRequestRequestPreamble {
	return v.value
}

func (v *NullableQuoteListRequestQuoteSearchRequestRequestPreamble) Set(val *QuoteListRequestQuoteSearchRequestRequestPreamble) {
	v.value = val
	v.isSet = true
}

func (v NullableQuoteListRequestQuoteSearchRequestRequestPreamble) IsSet() bool {
	return v.isSet
}

func (v *NullableQuoteListRequestQuoteSearchRequestRequestPreamble) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuoteListRequestQuoteSearchRequestRequestPreamble(val *QuoteListRequestQuoteSearchRequestRequestPreamble) *NullableQuoteListRequestQuoteSearchRequestRequestPreamble {
	return &NullableQuoteListRequestQuoteSearchRequestRequestPreamble{value: val, isSet: true}
}

func (v NullableQuoteListRequestQuoteSearchRequestRequestPreamble) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuoteListRequestQuoteSearchRequestRequestPreamble) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


