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

// checks if the QuoteDetailsRequestQuoteProductsRequestRetrieveQuoteProductsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuoteDetailsRequestQuoteProductsRequestRetrieveQuoteProductsRequest{}

// QuoteDetailsRequestQuoteProductsRequestRetrieveQuoteProductsRequest struct for QuoteDetailsRequestQuoteProductsRequestRetrieveQuoteProductsRequest
type QuoteDetailsRequestQuoteProductsRequestRetrieveQuoteProductsRequest struct {
	// Unique identifier generated by Ingram Micro's CRM specific to each quote. When applying a filter to the quoteNumber and including a partial quote number in the filter, all quotes containing any information included in the filter can be retrieved as a subset of all available customer quotes.
	QuoteNumber *string `json:"quoteNumber,omitempty"`
	// Unique identifier used to identify the third party source accessing the services.
	ThirdPartySource *string `json:"thirdPartySource,omitempty"`
}

// NewQuoteDetailsRequestQuoteProductsRequestRetrieveQuoteProductsRequest instantiates a new QuoteDetailsRequestQuoteProductsRequestRetrieveQuoteProductsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuoteDetailsRequestQuoteProductsRequestRetrieveQuoteProductsRequest() *QuoteDetailsRequestQuoteProductsRequestRetrieveQuoteProductsRequest {
	this := QuoteDetailsRequestQuoteProductsRequestRetrieveQuoteProductsRequest{}
	return &this
}

// NewQuoteDetailsRequestQuoteProductsRequestRetrieveQuoteProductsRequestWithDefaults instantiates a new QuoteDetailsRequestQuoteProductsRequestRetrieveQuoteProductsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuoteDetailsRequestQuoteProductsRequestRetrieveQuoteProductsRequestWithDefaults() *QuoteDetailsRequestQuoteProductsRequestRetrieveQuoteProductsRequest {
	this := QuoteDetailsRequestQuoteProductsRequestRetrieveQuoteProductsRequest{}
	return &this
}

// GetQuoteNumber returns the QuoteNumber field value if set, zero value otherwise.
func (o *QuoteDetailsRequestQuoteProductsRequestRetrieveQuoteProductsRequest) GetQuoteNumber() string {
	if o == nil || IsNil(o.QuoteNumber) {
		var ret string
		return ret
	}
	return *o.QuoteNumber
}

// GetQuoteNumberOk returns a tuple with the QuoteNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteDetailsRequestQuoteProductsRequestRetrieveQuoteProductsRequest) GetQuoteNumberOk() (*string, bool) {
	if o == nil || IsNil(o.QuoteNumber) {
		return nil, false
	}
	return o.QuoteNumber, true
}

// HasQuoteNumber returns a boolean if a field has been set.
func (o *QuoteDetailsRequestQuoteProductsRequestRetrieveQuoteProductsRequest) HasQuoteNumber() bool {
	if o != nil && !IsNil(o.QuoteNumber) {
		return true
	}

	return false
}

// SetQuoteNumber gets a reference to the given string and assigns it to the QuoteNumber field.
func (o *QuoteDetailsRequestQuoteProductsRequestRetrieveQuoteProductsRequest) SetQuoteNumber(v string) {
	o.QuoteNumber = &v
}

// GetThirdPartySource returns the ThirdPartySource field value if set, zero value otherwise.
func (o *QuoteDetailsRequestQuoteProductsRequestRetrieveQuoteProductsRequest) GetThirdPartySource() string {
	if o == nil || IsNil(o.ThirdPartySource) {
		var ret string
		return ret
	}
	return *o.ThirdPartySource
}

// GetThirdPartySourceOk returns a tuple with the ThirdPartySource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteDetailsRequestQuoteProductsRequestRetrieveQuoteProductsRequest) GetThirdPartySourceOk() (*string, bool) {
	if o == nil || IsNil(o.ThirdPartySource) {
		return nil, false
	}
	return o.ThirdPartySource, true
}

// HasThirdPartySource returns a boolean if a field has been set.
func (o *QuoteDetailsRequestQuoteProductsRequestRetrieveQuoteProductsRequest) HasThirdPartySource() bool {
	if o != nil && !IsNil(o.ThirdPartySource) {
		return true
	}

	return false
}

// SetThirdPartySource gets a reference to the given string and assigns it to the ThirdPartySource field.
func (o *QuoteDetailsRequestQuoteProductsRequestRetrieveQuoteProductsRequest) SetThirdPartySource(v string) {
	o.ThirdPartySource = &v
}

func (o QuoteDetailsRequestQuoteProductsRequestRetrieveQuoteProductsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuoteDetailsRequestQuoteProductsRequestRetrieveQuoteProductsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.QuoteNumber) {
		toSerialize["quoteNumber"] = o.QuoteNumber
	}
	if !IsNil(o.ThirdPartySource) {
		toSerialize["thirdPartySource"] = o.ThirdPartySource
	}
	return toSerialize, nil
}

type NullableQuoteDetailsRequestQuoteProductsRequestRetrieveQuoteProductsRequest struct {
	value *QuoteDetailsRequestQuoteProductsRequestRetrieveQuoteProductsRequest
	isSet bool
}

func (v NullableQuoteDetailsRequestQuoteProductsRequestRetrieveQuoteProductsRequest) Get() *QuoteDetailsRequestQuoteProductsRequestRetrieveQuoteProductsRequest {
	return v.value
}

func (v *NullableQuoteDetailsRequestQuoteProductsRequestRetrieveQuoteProductsRequest) Set(val *QuoteDetailsRequestQuoteProductsRequestRetrieveQuoteProductsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableQuoteDetailsRequestQuoteProductsRequestRetrieveQuoteProductsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableQuoteDetailsRequestQuoteProductsRequestRetrieveQuoteProductsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuoteDetailsRequestQuoteProductsRequestRetrieveQuoteProductsRequest(val *QuoteDetailsRequestQuoteProductsRequestRetrieveQuoteProductsRequest) *NullableQuoteDetailsRequestQuoteProductsRequestRetrieveQuoteProductsRequest {
	return &NullableQuoteDetailsRequestQuoteProductsRequestRetrieveQuoteProductsRequest{value: val, isSet: true}
}

func (v NullableQuoteDetailsRequestQuoteProductsRequestRetrieveQuoteProductsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuoteDetailsRequestQuoteProductsRequestRetrieveQuoteProductsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


