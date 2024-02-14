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

// checks if the QuoteListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuoteListResponse{}

// QuoteListResponse Response schema for get quote list endpoint
type QuoteListResponse struct {
	QuoteSearchResponse *QuoteListResponseQuoteSearchResponse `json:"quoteSearchResponse,omitempty"`
}

// NewQuoteListResponse instantiates a new QuoteListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuoteListResponse() *QuoteListResponse {
	this := QuoteListResponse{}
	return &this
}

// NewQuoteListResponseWithDefaults instantiates a new QuoteListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuoteListResponseWithDefaults() *QuoteListResponse {
	this := QuoteListResponse{}
	return &this
}

// GetQuoteSearchResponse returns the QuoteSearchResponse field value if set, zero value otherwise.
func (o *QuoteListResponse) GetQuoteSearchResponse() QuoteListResponseQuoteSearchResponse {
	if o == nil || IsNil(o.QuoteSearchResponse) {
		var ret QuoteListResponseQuoteSearchResponse
		return ret
	}
	return *o.QuoteSearchResponse
}

// GetQuoteSearchResponseOk returns a tuple with the QuoteSearchResponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteListResponse) GetQuoteSearchResponseOk() (*QuoteListResponseQuoteSearchResponse, bool) {
	if o == nil || IsNil(o.QuoteSearchResponse) {
		return nil, false
	}
	return o.QuoteSearchResponse, true
}

// HasQuoteSearchResponse returns a boolean if a field has been set.
func (o *QuoteListResponse) HasQuoteSearchResponse() bool {
	if o != nil && !IsNil(o.QuoteSearchResponse) {
		return true
	}

	return false
}

// SetQuoteSearchResponse gets a reference to the given QuoteListResponseQuoteSearchResponse and assigns it to the QuoteSearchResponse field.
func (o *QuoteListResponse) SetQuoteSearchResponse(v QuoteListResponseQuoteSearchResponse) {
	o.QuoteSearchResponse = &v
}

func (o QuoteListResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuoteListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.QuoteSearchResponse) {
		toSerialize["quoteSearchResponse"] = o.QuoteSearchResponse
	}
	return toSerialize, nil
}

type NullableQuoteListResponse struct {
	value *QuoteListResponse
	isSet bool
}

func (v NullableQuoteListResponse) Get() *QuoteListResponse {
	return v.value
}

func (v *NullableQuoteListResponse) Set(val *QuoteListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQuoteListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQuoteListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuoteListResponse(val *QuoteListResponse) *NullableQuoteListResponse {
	return &NullableQuoteListResponse{value: val, isSet: true}
}

func (v NullableQuoteListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuoteListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


