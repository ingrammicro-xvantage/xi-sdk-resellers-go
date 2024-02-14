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

// checks if the OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInner{}

// OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInner struct for OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInner
type OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInner struct {
	// A sub order number
	Subordernumber *string `json:"subordernumber,omitempty"`
	// Order status code
	Statuscode *string `json:"statuscode,omitempty"`
	// Details of the order statuscode - i.e. statuscode = 4 then status = SHIPPED
	Status *string `json:"status,omitempty"`
	// Will be returned in case of order on hold
	Holdreasoncode *string `json:"holdreasoncode,omitempty"`
	// Reason for order hold - will be returned if the order is on hold
	Holdreason *string `json:"holdreason,omitempty"`
	// HATEOAS links for the details and invoices of the sub-orders if available
	Links []OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInnerLinksInner `json:"links,omitempty"`
}

// NewOrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInner instantiates a new OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInner() *OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInner {
	this := OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInner{}
	return &this
}

// NewOrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInnerWithDefaults instantiates a new OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInnerWithDefaults() *OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInner {
	this := OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInner{}
	return &this
}

// GetSubordernumber returns the Subordernumber field value if set, zero value otherwise.
func (o *OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInner) GetSubordernumber() string {
	if o == nil || IsNil(o.Subordernumber) {
		var ret string
		return ret
	}
	return *o.Subordernumber
}

// GetSubordernumberOk returns a tuple with the Subordernumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInner) GetSubordernumberOk() (*string, bool) {
	if o == nil || IsNil(o.Subordernumber) {
		return nil, false
	}
	return o.Subordernumber, true
}

// HasSubordernumber returns a boolean if a field has been set.
func (o *OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInner) HasSubordernumber() bool {
	if o != nil && !IsNil(o.Subordernumber) {
		return true
	}

	return false
}

// SetSubordernumber gets a reference to the given string and assigns it to the Subordernumber field.
func (o *OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInner) SetSubordernumber(v string) {
	o.Subordernumber = &v
}

// GetStatuscode returns the Statuscode field value if set, zero value otherwise.
func (o *OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInner) GetStatuscode() string {
	if o == nil || IsNil(o.Statuscode) {
		var ret string
		return ret
	}
	return *o.Statuscode
}

// GetStatuscodeOk returns a tuple with the Statuscode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInner) GetStatuscodeOk() (*string, bool) {
	if o == nil || IsNil(o.Statuscode) {
		return nil, false
	}
	return o.Statuscode, true
}

// HasStatuscode returns a boolean if a field has been set.
func (o *OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInner) HasStatuscode() bool {
	if o != nil && !IsNil(o.Statuscode) {
		return true
	}

	return false
}

// SetStatuscode gets a reference to the given string and assigns it to the Statuscode field.
func (o *OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInner) SetStatuscode(v string) {
	o.Statuscode = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInner) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInner) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInner) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInner) SetStatus(v string) {
	o.Status = &v
}

// GetHoldreasoncode returns the Holdreasoncode field value if set, zero value otherwise.
func (o *OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInner) GetHoldreasoncode() string {
	if o == nil || IsNil(o.Holdreasoncode) {
		var ret string
		return ret
	}
	return *o.Holdreasoncode
}

// GetHoldreasoncodeOk returns a tuple with the Holdreasoncode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInner) GetHoldreasoncodeOk() (*string, bool) {
	if o == nil || IsNil(o.Holdreasoncode) {
		return nil, false
	}
	return o.Holdreasoncode, true
}

// HasHoldreasoncode returns a boolean if a field has been set.
func (o *OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInner) HasHoldreasoncode() bool {
	if o != nil && !IsNil(o.Holdreasoncode) {
		return true
	}

	return false
}

// SetHoldreasoncode gets a reference to the given string and assigns it to the Holdreasoncode field.
func (o *OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInner) SetHoldreasoncode(v string) {
	o.Holdreasoncode = &v
}

// GetHoldreason returns the Holdreason field value if set, zero value otherwise.
func (o *OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInner) GetHoldreason() string {
	if o == nil || IsNil(o.Holdreason) {
		var ret string
		return ret
	}
	return *o.Holdreason
}

// GetHoldreasonOk returns a tuple with the Holdreason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInner) GetHoldreasonOk() (*string, bool) {
	if o == nil || IsNil(o.Holdreason) {
		return nil, false
	}
	return o.Holdreason, true
}

// HasHoldreason returns a boolean if a field has been set.
func (o *OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInner) HasHoldreason() bool {
	if o != nil && !IsNil(o.Holdreason) {
		return true
	}

	return false
}

// SetHoldreason gets a reference to the given string and assigns it to the Holdreason field.
func (o *OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInner) SetHoldreason(v string) {
	o.Holdreason = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInner) GetLinks() []OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInnerLinksInner {
	if o == nil || IsNil(o.Links) {
		var ret []OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInnerLinksInner
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInner) GetLinksOk() ([]OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInnerLinksInner, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInner) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInnerLinksInner and assigns it to the Links field.
func (o *OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInner) SetLinks(v []OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInnerLinksInner) {
	o.Links = v
}

func (o OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Subordernumber) {
		toSerialize["subordernumber"] = o.Subordernumber
	}
	if !IsNil(o.Statuscode) {
		toSerialize["statuscode"] = o.Statuscode
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Holdreasoncode) {
		toSerialize["holdreasoncode"] = o.Holdreasoncode
	}
	if !IsNil(o.Holdreason) {
		toSerialize["holdreason"] = o.Holdreason
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	return toSerialize, nil
}

type NullableOrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInner struct {
	value *OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInner
	isSet bool
}

func (v NullableOrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInner) Get() *OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInner {
	return v.value
}

func (v *NullableOrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInner) Set(val *OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInner(val *OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInner) *NullableOrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInner {
	return &NullableOrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInner{value: val, isSet: true}
}

func (v NullableOrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


