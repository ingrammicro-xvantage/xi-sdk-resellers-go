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

// checks if the DealsDetailsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DealsDetailsResponse{}

// DealsDetailsResponse struct for DealsDetailsResponse
type DealsDetailsResponse struct {
	// Deal/Special bid number.
	DealId *string `json:"dealId,omitempty"`
	// Most recent version number of the deal.
	Version *string `json:"version,omitempty"`
	// The end user/customer's name.
	EndUser *string `json:"endUser,omitempty"`
	// Extended MSRP - Manufacturer Suggested Retail Price X Quantity.
	ExtendedMsrp *float32 `json:"extendedMsrp,omitempty"`
	// The vendor's name.
	Vendor *string `json:"vendor,omitempty"`
	// The date on which the deal starts.
	DealReceivedOn *string `json:"dealReceivedOn,omitempty"`
	// Expiration date of the deal/Special bid.
	DealExpiryDate *string `json:"dealExpiryDate,omitempty"`
	// The date on which the price protection will end.
	PriceProtectionEndDate *string `json:"priceProtectionEndDate,omitempty"`
	// Country specific currency code.
	CurrencyCode *string `json:"currencyCode,omitempty"`
	EndUserInfo []RenewalsDetailsResponseEndUserInfoInner `json:"endUserInfo,omitempty"`
	Products []DealsDetailsResponseProductsInner `json:"products,omitempty"`
}

// NewDealsDetailsResponse instantiates a new DealsDetailsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDealsDetailsResponse() *DealsDetailsResponse {
	this := DealsDetailsResponse{}
	return &this
}

// NewDealsDetailsResponseWithDefaults instantiates a new DealsDetailsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDealsDetailsResponseWithDefaults() *DealsDetailsResponse {
	this := DealsDetailsResponse{}
	return &this
}

// GetDealId returns the DealId field value if set, zero value otherwise.
func (o *DealsDetailsResponse) GetDealId() string {
	if o == nil || IsNil(o.DealId) {
		var ret string
		return ret
	}
	return *o.DealId
}

// GetDealIdOk returns a tuple with the DealId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DealsDetailsResponse) GetDealIdOk() (*string, bool) {
	if o == nil || IsNil(o.DealId) {
		return nil, false
	}
	return o.DealId, true
}

// HasDealId returns a boolean if a field has been set.
func (o *DealsDetailsResponse) HasDealId() bool {
	if o != nil && !IsNil(o.DealId) {
		return true
	}

	return false
}

// SetDealId gets a reference to the given string and assigns it to the DealId field.
func (o *DealsDetailsResponse) SetDealId(v string) {
	o.DealId = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *DealsDetailsResponse) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DealsDetailsResponse) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *DealsDetailsResponse) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *DealsDetailsResponse) SetVersion(v string) {
	o.Version = &v
}

// GetEndUser returns the EndUser field value if set, zero value otherwise.
func (o *DealsDetailsResponse) GetEndUser() string {
	if o == nil || IsNil(o.EndUser) {
		var ret string
		return ret
	}
	return *o.EndUser
}

// GetEndUserOk returns a tuple with the EndUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DealsDetailsResponse) GetEndUserOk() (*string, bool) {
	if o == nil || IsNil(o.EndUser) {
		return nil, false
	}
	return o.EndUser, true
}

// HasEndUser returns a boolean if a field has been set.
func (o *DealsDetailsResponse) HasEndUser() bool {
	if o != nil && !IsNil(o.EndUser) {
		return true
	}

	return false
}

// SetEndUser gets a reference to the given string and assigns it to the EndUser field.
func (o *DealsDetailsResponse) SetEndUser(v string) {
	o.EndUser = &v
}

// GetExtendedMsrp returns the ExtendedMsrp field value if set, zero value otherwise.
func (o *DealsDetailsResponse) GetExtendedMsrp() float32 {
	if o == nil || IsNil(o.ExtendedMsrp) {
		var ret float32
		return ret
	}
	return *o.ExtendedMsrp
}

// GetExtendedMsrpOk returns a tuple with the ExtendedMsrp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DealsDetailsResponse) GetExtendedMsrpOk() (*float32, bool) {
	if o == nil || IsNil(o.ExtendedMsrp) {
		return nil, false
	}
	return o.ExtendedMsrp, true
}

// HasExtendedMsrp returns a boolean if a field has been set.
func (o *DealsDetailsResponse) HasExtendedMsrp() bool {
	if o != nil && !IsNil(o.ExtendedMsrp) {
		return true
	}

	return false
}

// SetExtendedMsrp gets a reference to the given float32 and assigns it to the ExtendedMsrp field.
func (o *DealsDetailsResponse) SetExtendedMsrp(v float32) {
	o.ExtendedMsrp = &v
}

// GetVendor returns the Vendor field value if set, zero value otherwise.
func (o *DealsDetailsResponse) GetVendor() string {
	if o == nil || IsNil(o.Vendor) {
		var ret string
		return ret
	}
	return *o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DealsDetailsResponse) GetVendorOk() (*string, bool) {
	if o == nil || IsNil(o.Vendor) {
		return nil, false
	}
	return o.Vendor, true
}

// HasVendor returns a boolean if a field has been set.
func (o *DealsDetailsResponse) HasVendor() bool {
	if o != nil && !IsNil(o.Vendor) {
		return true
	}

	return false
}

// SetVendor gets a reference to the given string and assigns it to the Vendor field.
func (o *DealsDetailsResponse) SetVendor(v string) {
	o.Vendor = &v
}

// GetDealReceivedOn returns the DealReceivedOn field value if set, zero value otherwise.
func (o *DealsDetailsResponse) GetDealReceivedOn() string {
	if o == nil || IsNil(o.DealReceivedOn) {
		var ret string
		return ret
	}
	return *o.DealReceivedOn
}

// GetDealReceivedOnOk returns a tuple with the DealReceivedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DealsDetailsResponse) GetDealReceivedOnOk() (*string, bool) {
	if o == nil || IsNil(o.DealReceivedOn) {
		return nil, false
	}
	return o.DealReceivedOn, true
}

// HasDealReceivedOn returns a boolean if a field has been set.
func (o *DealsDetailsResponse) HasDealReceivedOn() bool {
	if o != nil && !IsNil(o.DealReceivedOn) {
		return true
	}

	return false
}

// SetDealReceivedOn gets a reference to the given string and assigns it to the DealReceivedOn field.
func (o *DealsDetailsResponse) SetDealReceivedOn(v string) {
	o.DealReceivedOn = &v
}

// GetDealExpiryDate returns the DealExpiryDate field value if set, zero value otherwise.
func (o *DealsDetailsResponse) GetDealExpiryDate() string {
	if o == nil || IsNil(o.DealExpiryDate) {
		var ret string
		return ret
	}
	return *o.DealExpiryDate
}

// GetDealExpiryDateOk returns a tuple with the DealExpiryDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DealsDetailsResponse) GetDealExpiryDateOk() (*string, bool) {
	if o == nil || IsNil(o.DealExpiryDate) {
		return nil, false
	}
	return o.DealExpiryDate, true
}

// HasDealExpiryDate returns a boolean if a field has been set.
func (o *DealsDetailsResponse) HasDealExpiryDate() bool {
	if o != nil && !IsNil(o.DealExpiryDate) {
		return true
	}

	return false
}

// SetDealExpiryDate gets a reference to the given string and assigns it to the DealExpiryDate field.
func (o *DealsDetailsResponse) SetDealExpiryDate(v string) {
	o.DealExpiryDate = &v
}

// GetPriceProtectionEndDate returns the PriceProtectionEndDate field value if set, zero value otherwise.
func (o *DealsDetailsResponse) GetPriceProtectionEndDate() string {
	if o == nil || IsNil(o.PriceProtectionEndDate) {
		var ret string
		return ret
	}
	return *o.PriceProtectionEndDate
}

// GetPriceProtectionEndDateOk returns a tuple with the PriceProtectionEndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DealsDetailsResponse) GetPriceProtectionEndDateOk() (*string, bool) {
	if o == nil || IsNil(o.PriceProtectionEndDate) {
		return nil, false
	}
	return o.PriceProtectionEndDate, true
}

// HasPriceProtectionEndDate returns a boolean if a field has been set.
func (o *DealsDetailsResponse) HasPriceProtectionEndDate() bool {
	if o != nil && !IsNil(o.PriceProtectionEndDate) {
		return true
	}

	return false
}

// SetPriceProtectionEndDate gets a reference to the given string and assigns it to the PriceProtectionEndDate field.
func (o *DealsDetailsResponse) SetPriceProtectionEndDate(v string) {
	o.PriceProtectionEndDate = &v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *DealsDetailsResponse) GetCurrencyCode() string {
	if o == nil || IsNil(o.CurrencyCode) {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DealsDetailsResponse) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CurrencyCode) {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *DealsDetailsResponse) HasCurrencyCode() bool {
	if o != nil && !IsNil(o.CurrencyCode) {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *DealsDetailsResponse) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

// GetEndUserInfo returns the EndUserInfo field value if set, zero value otherwise.
func (o *DealsDetailsResponse) GetEndUserInfo() []RenewalsDetailsResponseEndUserInfoInner {
	if o == nil || IsNil(o.EndUserInfo) {
		var ret []RenewalsDetailsResponseEndUserInfoInner
		return ret
	}
	return o.EndUserInfo
}

// GetEndUserInfoOk returns a tuple with the EndUserInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DealsDetailsResponse) GetEndUserInfoOk() ([]RenewalsDetailsResponseEndUserInfoInner, bool) {
	if o == nil || IsNil(o.EndUserInfo) {
		return nil, false
	}
	return o.EndUserInfo, true
}

// HasEndUserInfo returns a boolean if a field has been set.
func (o *DealsDetailsResponse) HasEndUserInfo() bool {
	if o != nil && !IsNil(o.EndUserInfo) {
		return true
	}

	return false
}

// SetEndUserInfo gets a reference to the given []RenewalsDetailsResponseEndUserInfoInner and assigns it to the EndUserInfo field.
func (o *DealsDetailsResponse) SetEndUserInfo(v []RenewalsDetailsResponseEndUserInfoInner) {
	o.EndUserInfo = v
}

// GetProducts returns the Products field value if set, zero value otherwise.
func (o *DealsDetailsResponse) GetProducts() []DealsDetailsResponseProductsInner {
	if o == nil || IsNil(o.Products) {
		var ret []DealsDetailsResponseProductsInner
		return ret
	}
	return o.Products
}

// GetProductsOk returns a tuple with the Products field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DealsDetailsResponse) GetProductsOk() ([]DealsDetailsResponseProductsInner, bool) {
	if o == nil || IsNil(o.Products) {
		return nil, false
	}
	return o.Products, true
}

// HasProducts returns a boolean if a field has been set.
func (o *DealsDetailsResponse) HasProducts() bool {
	if o != nil && !IsNil(o.Products) {
		return true
	}

	return false
}

// SetProducts gets a reference to the given []DealsDetailsResponseProductsInner and assigns it to the Products field.
func (o *DealsDetailsResponse) SetProducts(v []DealsDetailsResponseProductsInner) {
	o.Products = v
}

func (o DealsDetailsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DealsDetailsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DealId) {
		toSerialize["dealId"] = o.DealId
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.EndUser) {
		toSerialize["endUser"] = o.EndUser
	}
	if !IsNil(o.ExtendedMsrp) {
		toSerialize["extendedMsrp"] = o.ExtendedMsrp
	}
	if !IsNil(o.Vendor) {
		toSerialize["vendor"] = o.Vendor
	}
	if !IsNil(o.DealReceivedOn) {
		toSerialize["dealReceivedOn"] = o.DealReceivedOn
	}
	if !IsNil(o.DealExpiryDate) {
		toSerialize["dealExpiryDate"] = o.DealExpiryDate
	}
	if !IsNil(o.PriceProtectionEndDate) {
		toSerialize["priceProtectionEndDate"] = o.PriceProtectionEndDate
	}
	if !IsNil(o.CurrencyCode) {
		toSerialize["currencyCode"] = o.CurrencyCode
	}
	if !IsNil(o.EndUserInfo) {
		toSerialize["endUserInfo"] = o.EndUserInfo
	}
	if !IsNil(o.Products) {
		toSerialize["products"] = o.Products
	}
	return toSerialize, nil
}

type NullableDealsDetailsResponse struct {
	value *DealsDetailsResponse
	isSet bool
}

func (v NullableDealsDetailsResponse) Get() *DealsDetailsResponse {
	return v.value
}

func (v *NullableDealsDetailsResponse) Set(val *DealsDetailsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDealsDetailsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDealsDetailsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDealsDetailsResponse(val *DealsDetailsResponse) *NullableDealsDetailsResponse {
	return &NullableDealsDetailsResponse{value: val, isSet: true}
}

func (v NullableDealsDetailsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDealsDetailsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


