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

// checks if the QuoteToOrderResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuoteToOrderResponse{}

// QuoteToOrderResponse struct for QuoteToOrderResponse
type QuoteToOrderResponse struct {
	// Unique identifier generated by Ingram Micro's CRM specific to each quote.
	QuoteNumber *string `json:"quoteNumber,omitempty"`
	// A unique confirmation number for tracking purposes.
	ConfirmationNumber *string `json:"confirmationNumber,omitempty"`
	// A confirmation message.
	Message *string `json:"message,omitempty"`
}

// NewQuoteToOrderResponse instantiates a new QuoteToOrderResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuoteToOrderResponse() *QuoteToOrderResponse {
	this := QuoteToOrderResponse{}
	return &this
}

// NewQuoteToOrderResponseWithDefaults instantiates a new QuoteToOrderResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuoteToOrderResponseWithDefaults() *QuoteToOrderResponse {
	this := QuoteToOrderResponse{}
	return &this
}

// GetQuoteNumber returns the QuoteNumber field value if set, zero value otherwise.
func (o *QuoteToOrderResponse) GetQuoteNumber() string {
	if o == nil || IsNil(o.QuoteNumber) {
		var ret string
		return ret
	}
	return *o.QuoteNumber
}

// GetQuoteNumberOk returns a tuple with the QuoteNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteToOrderResponse) GetQuoteNumberOk() (*string, bool) {
	if o == nil || IsNil(o.QuoteNumber) {
		return nil, false
	}
	return o.QuoteNumber, true
}

// HasQuoteNumber returns a boolean if a field has been set.
func (o *QuoteToOrderResponse) HasQuoteNumber() bool {
	if o != nil && !IsNil(o.QuoteNumber) {
		return true
	}

	return false
}

// SetQuoteNumber gets a reference to the given string and assigns it to the QuoteNumber field.
func (o *QuoteToOrderResponse) SetQuoteNumber(v string) {
	o.QuoteNumber = &v
}

// GetConfirmationNumber returns the ConfirmationNumber field value if set, zero value otherwise.
func (o *QuoteToOrderResponse) GetConfirmationNumber() string {
	if o == nil || IsNil(o.ConfirmationNumber) {
		var ret string
		return ret
	}
	return *o.ConfirmationNumber
}

// GetConfirmationNumberOk returns a tuple with the ConfirmationNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteToOrderResponse) GetConfirmationNumberOk() (*string, bool) {
	if o == nil || IsNil(o.ConfirmationNumber) {
		return nil, false
	}
	return o.ConfirmationNumber, true
}

// HasConfirmationNumber returns a boolean if a field has been set.
func (o *QuoteToOrderResponse) HasConfirmationNumber() bool {
	if o != nil && !IsNil(o.ConfirmationNumber) {
		return true
	}

	return false
}

// SetConfirmationNumber gets a reference to the given string and assigns it to the ConfirmationNumber field.
func (o *QuoteToOrderResponse) SetConfirmationNumber(v string) {
	o.ConfirmationNumber = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *QuoteToOrderResponse) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteToOrderResponse) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *QuoteToOrderResponse) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *QuoteToOrderResponse) SetMessage(v string) {
	o.Message = &v
}

func (o QuoteToOrderResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuoteToOrderResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.QuoteNumber) {
		toSerialize["quoteNumber"] = o.QuoteNumber
	}
	if !IsNil(o.ConfirmationNumber) {
		toSerialize["confirmationNumber"] = o.ConfirmationNumber
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableQuoteToOrderResponse struct {
	value *QuoteToOrderResponse
	isSet bool
}

func (v NullableQuoteToOrderResponse) Get() *QuoteToOrderResponse {
	return v.value
}

func (v *NullableQuoteToOrderResponse) Set(val *QuoteToOrderResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQuoteToOrderResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQuoteToOrderResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuoteToOrderResponse(val *QuoteToOrderResponse) *NullableQuoteToOrderResponse {
	return &NullableQuoteToOrderResponse{value: val, isSet: true}
}

func (v NullableQuoteToOrderResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuoteToOrderResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


