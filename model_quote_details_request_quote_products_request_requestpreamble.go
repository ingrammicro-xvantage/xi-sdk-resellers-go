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

// checks if the QuoteDetailsRequestQuoteProductsRequestRequestpreamble type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuoteDetailsRequestQuoteProductsRequestRequestpreamble{}

// QuoteDetailsRequestQuoteProductsRequestRequestpreamble struct for QuoteDetailsRequestQuoteProductsRequestRequestpreamble
type QuoteDetailsRequestQuoteProductsRequestRequestpreamble struct {
	// Reseller Number (referred to as the account BCN) is the unique identifier for an Ingram Micro customer account.
	CustomerNumber string `json:"customerNumber"`
	// The ISO country codes are internationally recognized codes designated for each country represented by a two-letter combination (alpha-2).
	IsoCountryCode string `json:"isoCountryCode"`
}

type _QuoteDetailsRequestQuoteProductsRequestRequestpreamble QuoteDetailsRequestQuoteProductsRequestRequestpreamble

// NewQuoteDetailsRequestQuoteProductsRequestRequestpreamble instantiates a new QuoteDetailsRequestQuoteProductsRequestRequestpreamble object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuoteDetailsRequestQuoteProductsRequestRequestpreamble(customerNumber string, isoCountryCode string) *QuoteDetailsRequestQuoteProductsRequestRequestpreamble {
	this := QuoteDetailsRequestQuoteProductsRequestRequestpreamble{}
	this.CustomerNumber = customerNumber
	this.IsoCountryCode = isoCountryCode
	return &this
}

// NewQuoteDetailsRequestQuoteProductsRequestRequestpreambleWithDefaults instantiates a new QuoteDetailsRequestQuoteProductsRequestRequestpreamble object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuoteDetailsRequestQuoteProductsRequestRequestpreambleWithDefaults() *QuoteDetailsRequestQuoteProductsRequestRequestpreamble {
	this := QuoteDetailsRequestQuoteProductsRequestRequestpreamble{}
	return &this
}

// GetCustomerNumber returns the CustomerNumber field value
func (o *QuoteDetailsRequestQuoteProductsRequestRequestpreamble) GetCustomerNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerNumber
}

// GetCustomerNumberOk returns a tuple with the CustomerNumber field value
// and a boolean to check if the value has been set.
func (o *QuoteDetailsRequestQuoteProductsRequestRequestpreamble) GetCustomerNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerNumber, true
}

// SetCustomerNumber sets field value
func (o *QuoteDetailsRequestQuoteProductsRequestRequestpreamble) SetCustomerNumber(v string) {
	o.CustomerNumber = v
}

// GetIsoCountryCode returns the IsoCountryCode field value
func (o *QuoteDetailsRequestQuoteProductsRequestRequestpreamble) GetIsoCountryCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IsoCountryCode
}

// GetIsoCountryCodeOk returns a tuple with the IsoCountryCode field value
// and a boolean to check if the value has been set.
func (o *QuoteDetailsRequestQuoteProductsRequestRequestpreamble) GetIsoCountryCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsoCountryCode, true
}

// SetIsoCountryCode sets field value
func (o *QuoteDetailsRequestQuoteProductsRequestRequestpreamble) SetIsoCountryCode(v string) {
	o.IsoCountryCode = v
}

func (o QuoteDetailsRequestQuoteProductsRequestRequestpreamble) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuoteDetailsRequestQuoteProductsRequestRequestpreamble) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["customerNumber"] = o.CustomerNumber
	toSerialize["isoCountryCode"] = o.IsoCountryCode
	return toSerialize, nil
}

func (o *QuoteDetailsRequestQuoteProductsRequestRequestpreamble) UnmarshalJSON(data []byte) (err error) {
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

	varQuoteDetailsRequestQuoteProductsRequestRequestpreamble := _QuoteDetailsRequestQuoteProductsRequestRequestpreamble{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varQuoteDetailsRequestQuoteProductsRequestRequestpreamble)

	if err != nil {
		return err
	}

	*o = QuoteDetailsRequestQuoteProductsRequestRequestpreamble(varQuoteDetailsRequestQuoteProductsRequestRequestpreamble)

	return err
}

type NullableQuoteDetailsRequestQuoteProductsRequestRequestpreamble struct {
	value *QuoteDetailsRequestQuoteProductsRequestRequestpreamble
	isSet bool
}

func (v NullableQuoteDetailsRequestQuoteProductsRequestRequestpreamble) Get() *QuoteDetailsRequestQuoteProductsRequestRequestpreamble {
	return v.value
}

func (v *NullableQuoteDetailsRequestQuoteProductsRequestRequestpreamble) Set(val *QuoteDetailsRequestQuoteProductsRequestRequestpreamble) {
	v.value = val
	v.isSet = true
}

func (v NullableQuoteDetailsRequestQuoteProductsRequestRequestpreamble) IsSet() bool {
	return v.isSet
}

func (v *NullableQuoteDetailsRequestQuoteProductsRequestRequestpreamble) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuoteDetailsRequestQuoteProductsRequestRequestpreamble(val *QuoteDetailsRequestQuoteProductsRequestRequestpreamble) *NullableQuoteDetailsRequestQuoteProductsRequestRequestpreamble {
	return &NullableQuoteDetailsRequestQuoteProductsRequestRequestpreamble{value: val, isSet: true}
}

func (v NullableQuoteDetailsRequestQuoteProductsRequestRequestpreamble) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuoteDetailsRequestQuoteProductsRequestRequestpreamble) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


