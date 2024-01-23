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

// checks if the OrderSearchRequestServicerequestOrderLookupRequestCustomerOrderNumber type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderSearchRequestServicerequestOrderLookupRequestCustomerOrderNumber{}

// OrderSearchRequestServicerequestOrderLookupRequestCustomerOrderNumber struct for OrderSearchRequestServicerequestOrderLookupRequestCustomerOrderNumber
type OrderSearchRequestServicerequestOrderLookupRequestCustomerOrderNumber struct {
	CustomerOrderNumber *string `json:"customerOrderNumber,omitempty"`
}

// NewOrderSearchRequestServicerequestOrderLookupRequestCustomerOrderNumber instantiates a new OrderSearchRequestServicerequestOrderLookupRequestCustomerOrderNumber object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderSearchRequestServicerequestOrderLookupRequestCustomerOrderNumber() *OrderSearchRequestServicerequestOrderLookupRequestCustomerOrderNumber {
	this := OrderSearchRequestServicerequestOrderLookupRequestCustomerOrderNumber{}
	return &this
}

// NewOrderSearchRequestServicerequestOrderLookupRequestCustomerOrderNumberWithDefaults instantiates a new OrderSearchRequestServicerequestOrderLookupRequestCustomerOrderNumber object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderSearchRequestServicerequestOrderLookupRequestCustomerOrderNumberWithDefaults() *OrderSearchRequestServicerequestOrderLookupRequestCustomerOrderNumber {
	this := OrderSearchRequestServicerequestOrderLookupRequestCustomerOrderNumber{}
	return &this
}

// GetCustomerOrderNumber returns the CustomerOrderNumber field value if set, zero value otherwise.
func (o *OrderSearchRequestServicerequestOrderLookupRequestCustomerOrderNumber) GetCustomerOrderNumber() string {
	if o == nil || IsNil(o.CustomerOrderNumber) {
		var ret string
		return ret
	}
	return *o.CustomerOrderNumber
}

// GetCustomerOrderNumberOk returns a tuple with the CustomerOrderNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSearchRequestServicerequestOrderLookupRequestCustomerOrderNumber) GetCustomerOrderNumberOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerOrderNumber) {
		return nil, false
	}
	return o.CustomerOrderNumber, true
}

// HasCustomerOrderNumber returns a boolean if a field has been set.
func (o *OrderSearchRequestServicerequestOrderLookupRequestCustomerOrderNumber) HasCustomerOrderNumber() bool {
	if o != nil && !IsNil(o.CustomerOrderNumber) {
		return true
	}

	return false
}

// SetCustomerOrderNumber gets a reference to the given string and assigns it to the CustomerOrderNumber field.
func (o *OrderSearchRequestServicerequestOrderLookupRequestCustomerOrderNumber) SetCustomerOrderNumber(v string) {
	o.CustomerOrderNumber = &v
}

func (o OrderSearchRequestServicerequestOrderLookupRequestCustomerOrderNumber) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderSearchRequestServicerequestOrderLookupRequestCustomerOrderNumber) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CustomerOrderNumber) {
		toSerialize["customerOrderNumber"] = o.CustomerOrderNumber
	}
	return toSerialize, nil
}

type NullableOrderSearchRequestServicerequestOrderLookupRequestCustomerOrderNumber struct {
	value *OrderSearchRequestServicerequestOrderLookupRequestCustomerOrderNumber
	isSet bool
}

func (v NullableOrderSearchRequestServicerequestOrderLookupRequestCustomerOrderNumber) Get() *OrderSearchRequestServicerequestOrderLookupRequestCustomerOrderNumber {
	return v.value
}

func (v *NullableOrderSearchRequestServicerequestOrderLookupRequestCustomerOrderNumber) Set(val *OrderSearchRequestServicerequestOrderLookupRequestCustomerOrderNumber) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderSearchRequestServicerequestOrderLookupRequestCustomerOrderNumber) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderSearchRequestServicerequestOrderLookupRequestCustomerOrderNumber) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderSearchRequestServicerequestOrderLookupRequestCustomerOrderNumber(val *OrderSearchRequestServicerequestOrderLookupRequestCustomerOrderNumber) *NullableOrderSearchRequestServicerequestOrderLookupRequestCustomerOrderNumber {
	return &NullableOrderSearchRequestServicerequestOrderLookupRequestCustomerOrderNumber{value: val, isSet: true}
}

func (v NullableOrderSearchRequestServicerequestOrderLookupRequestCustomerOrderNumber) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderSearchRequestServicerequestOrderLookupRequestCustomerOrderNumber) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


