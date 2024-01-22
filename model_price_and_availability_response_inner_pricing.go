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

// checks if the PriceAndAvailabilityResponseInnerPricing type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PriceAndAvailabilityResponseInnerPricing{}

// PriceAndAvailabilityResponseInnerPricing struct for PriceAndAvailabilityResponseInnerPricing
type PriceAndAvailabilityResponseInnerPricing struct {
	// The 3-digit ISO currency code.
	CurrencyCode *string `json:"currencyCode,omitempty"`
	// The retail price of the product.
	RetailPrice *float32 `json:"retailPrice,omitempty"`
	// Minimum Advertised Price (MAP). If required by the vendor, resellers can not sell below MAP price.
	MapPrice *float32 `json:"mapPrice,omitempty"`
	// The price customer pays after all special pricing and discounts have been applied.
	CustomerPrice *float32 `json:"customerPrice,omitempty"`
	// Boolean values specifies whether special Bid discounts are available for the product.
	SpecialBidPricingAvailable *bool `json:"specialBidPricingAvailable,omitempty"`
	// Boolean values specifies whether web Discounts are available for the product.
	WebDiscountsAvailable *bool `json:"webDiscountsAvailable,omitempty"`
}

// NewPriceAndAvailabilityResponseInnerPricing instantiates a new PriceAndAvailabilityResponseInnerPricing object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPriceAndAvailabilityResponseInnerPricing() *PriceAndAvailabilityResponseInnerPricing {
	this := PriceAndAvailabilityResponseInnerPricing{}
	return &this
}

// NewPriceAndAvailabilityResponseInnerPricingWithDefaults instantiates a new PriceAndAvailabilityResponseInnerPricing object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceAndAvailabilityResponseInnerPricingWithDefaults() *PriceAndAvailabilityResponseInnerPricing {
	this := PriceAndAvailabilityResponseInnerPricing{}
	return &this
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *PriceAndAvailabilityResponseInnerPricing) GetCurrencyCode() string {
	if o == nil || IsNil(o.CurrencyCode) {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceAndAvailabilityResponseInnerPricing) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CurrencyCode) {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *PriceAndAvailabilityResponseInnerPricing) HasCurrencyCode() bool {
	if o != nil && !IsNil(o.CurrencyCode) {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *PriceAndAvailabilityResponseInnerPricing) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

// GetRetailPrice returns the RetailPrice field value if set, zero value otherwise.
func (o *PriceAndAvailabilityResponseInnerPricing) GetRetailPrice() float32 {
	if o == nil || IsNil(o.RetailPrice) {
		var ret float32
		return ret
	}
	return *o.RetailPrice
}

// GetRetailPriceOk returns a tuple with the RetailPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceAndAvailabilityResponseInnerPricing) GetRetailPriceOk() (*float32, bool) {
	if o == nil || IsNil(o.RetailPrice) {
		return nil, false
	}
	return o.RetailPrice, true
}

// HasRetailPrice returns a boolean if a field has been set.
func (o *PriceAndAvailabilityResponseInnerPricing) HasRetailPrice() bool {
	if o != nil && !IsNil(o.RetailPrice) {
		return true
	}

	return false
}

// SetRetailPrice gets a reference to the given float32 and assigns it to the RetailPrice field.
func (o *PriceAndAvailabilityResponseInnerPricing) SetRetailPrice(v float32) {
	o.RetailPrice = &v
}

// GetMapPrice returns the MapPrice field value if set, zero value otherwise.
func (o *PriceAndAvailabilityResponseInnerPricing) GetMapPrice() float32 {
	if o == nil || IsNil(o.MapPrice) {
		var ret float32
		return ret
	}
	return *o.MapPrice
}

// GetMapPriceOk returns a tuple with the MapPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceAndAvailabilityResponseInnerPricing) GetMapPriceOk() (*float32, bool) {
	if o == nil || IsNil(o.MapPrice) {
		return nil, false
	}
	return o.MapPrice, true
}

// HasMapPrice returns a boolean if a field has been set.
func (o *PriceAndAvailabilityResponseInnerPricing) HasMapPrice() bool {
	if o != nil && !IsNil(o.MapPrice) {
		return true
	}

	return false
}

// SetMapPrice gets a reference to the given float32 and assigns it to the MapPrice field.
func (o *PriceAndAvailabilityResponseInnerPricing) SetMapPrice(v float32) {
	o.MapPrice = &v
}

// GetCustomerPrice returns the CustomerPrice field value if set, zero value otherwise.
func (o *PriceAndAvailabilityResponseInnerPricing) GetCustomerPrice() float32 {
	if o == nil || IsNil(o.CustomerPrice) {
		var ret float32
		return ret
	}
	return *o.CustomerPrice
}

// GetCustomerPriceOk returns a tuple with the CustomerPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceAndAvailabilityResponseInnerPricing) GetCustomerPriceOk() (*float32, bool) {
	if o == nil || IsNil(o.CustomerPrice) {
		return nil, false
	}
	return o.CustomerPrice, true
}

// HasCustomerPrice returns a boolean if a field has been set.
func (o *PriceAndAvailabilityResponseInnerPricing) HasCustomerPrice() bool {
	if o != nil && !IsNil(o.CustomerPrice) {
		return true
	}

	return false
}

// SetCustomerPrice gets a reference to the given float32 and assigns it to the CustomerPrice field.
func (o *PriceAndAvailabilityResponseInnerPricing) SetCustomerPrice(v float32) {
	o.CustomerPrice = &v
}

// GetSpecialBidPricingAvailable returns the SpecialBidPricingAvailable field value if set, zero value otherwise.
func (o *PriceAndAvailabilityResponseInnerPricing) GetSpecialBidPricingAvailable() bool {
	if o == nil || IsNil(o.SpecialBidPricingAvailable) {
		var ret bool
		return ret
	}
	return *o.SpecialBidPricingAvailable
}

// GetSpecialBidPricingAvailableOk returns a tuple with the SpecialBidPricingAvailable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceAndAvailabilityResponseInnerPricing) GetSpecialBidPricingAvailableOk() (*bool, bool) {
	if o == nil || IsNil(o.SpecialBidPricingAvailable) {
		return nil, false
	}
	return o.SpecialBidPricingAvailable, true
}

// HasSpecialBidPricingAvailable returns a boolean if a field has been set.
func (o *PriceAndAvailabilityResponseInnerPricing) HasSpecialBidPricingAvailable() bool {
	if o != nil && !IsNil(o.SpecialBidPricingAvailable) {
		return true
	}

	return false
}

// SetSpecialBidPricingAvailable gets a reference to the given bool and assigns it to the SpecialBidPricingAvailable field.
func (o *PriceAndAvailabilityResponseInnerPricing) SetSpecialBidPricingAvailable(v bool) {
	o.SpecialBidPricingAvailable = &v
}

// GetWebDiscountsAvailable returns the WebDiscountsAvailable field value if set, zero value otherwise.
func (o *PriceAndAvailabilityResponseInnerPricing) GetWebDiscountsAvailable() bool {
	if o == nil || IsNil(o.WebDiscountsAvailable) {
		var ret bool
		return ret
	}
	return *o.WebDiscountsAvailable
}

// GetWebDiscountsAvailableOk returns a tuple with the WebDiscountsAvailable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceAndAvailabilityResponseInnerPricing) GetWebDiscountsAvailableOk() (*bool, bool) {
	if o == nil || IsNil(o.WebDiscountsAvailable) {
		return nil, false
	}
	return o.WebDiscountsAvailable, true
}

// HasWebDiscountsAvailable returns a boolean if a field has been set.
func (o *PriceAndAvailabilityResponseInnerPricing) HasWebDiscountsAvailable() bool {
	if o != nil && !IsNil(o.WebDiscountsAvailable) {
		return true
	}

	return false
}

// SetWebDiscountsAvailable gets a reference to the given bool and assigns it to the WebDiscountsAvailable field.
func (o *PriceAndAvailabilityResponseInnerPricing) SetWebDiscountsAvailable(v bool) {
	o.WebDiscountsAvailable = &v
}

func (o PriceAndAvailabilityResponseInnerPricing) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PriceAndAvailabilityResponseInnerPricing) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CurrencyCode) {
		toSerialize["currencyCode"] = o.CurrencyCode
	}
	if !IsNil(o.RetailPrice) {
		toSerialize["retailPrice"] = o.RetailPrice
	}
	if !IsNil(o.MapPrice) {
		toSerialize["mapPrice"] = o.MapPrice
	}
	if !IsNil(o.CustomerPrice) {
		toSerialize["customerPrice"] = o.CustomerPrice
	}
	if !IsNil(o.SpecialBidPricingAvailable) {
		toSerialize["specialBidPricingAvailable"] = o.SpecialBidPricingAvailable
	}
	if !IsNil(o.WebDiscountsAvailable) {
		toSerialize["webDiscountsAvailable"] = o.WebDiscountsAvailable
	}
	return toSerialize, nil
}

type NullablePriceAndAvailabilityResponseInnerPricing struct {
	value *PriceAndAvailabilityResponseInnerPricing
	isSet bool
}

func (v NullablePriceAndAvailabilityResponseInnerPricing) Get() *PriceAndAvailabilityResponseInnerPricing {
	return v.value
}

func (v *NullablePriceAndAvailabilityResponseInnerPricing) Set(val *PriceAndAvailabilityResponseInnerPricing) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceAndAvailabilityResponseInnerPricing) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceAndAvailabilityResponseInnerPricing) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceAndAvailabilityResponseInnerPricing(val *PriceAndAvailabilityResponseInnerPricing) *NullablePriceAndAvailabilityResponseInnerPricing {
	return &NullablePriceAndAvailabilityResponseInnerPricing{value: val, isSet: true}
}

func (v NullablePriceAndAvailabilityResponseInnerPricing) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceAndAvailabilityResponseInnerPricing) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


