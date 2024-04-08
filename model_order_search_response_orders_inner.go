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

// checks if the OrderSearchResponseOrdersInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderSearchResponseOrdersInner{}

// OrderSearchResponseOrdersInner struct for OrderSearchResponseOrdersInner
type OrderSearchResponseOrdersInner struct {
	// The Ingram Micro order number.
	IngramOrderNumber *string `json:"ingramOrderNumber,omitempty"`
	// The date the order was created(UTC).
	IngramOrderDate *string `json:"ingramOrderDate,omitempty"`
	// The reseller's order number for reference in their system.
	CustomerOrderNumber *string `json:"customerOrderNumber,omitempty"`
	// The vendor's order number.(only for D-Type Orders)
	VendorSalesOrderNumber *string `json:"vendorSalesOrderNumber,omitempty"`
	// The name of the vendor.
	VendorName *string `json:"vendorName,omitempty"`
	// The company name of the end user/customer.
	EndUserCompanyName *string `json:"endUserCompanyName,omitempty"`
	// The total of the order.
	OrderTotal *float32 `json:"orderTotal,omitempty"`
	// The header-level status of the order.(OPEN/CLOSED/CANCELLED)
	OrderStatus *string `json:"orderStatus,omitempty"`
	// Individual Ingram Micro order numbers associated with a single reseller PO.
	SubOrders []OrderSearchResponseOrdersInnerSubOrdersInner `json:"subOrders,omitempty"`
	Links *OrderSearchResponseOrdersInnerLinks `json:"links,omitempty"`
}

// NewOrderSearchResponseOrdersInner instantiates a new OrderSearchResponseOrdersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderSearchResponseOrdersInner() *OrderSearchResponseOrdersInner {
	this := OrderSearchResponseOrdersInner{}
	return &this
}

// NewOrderSearchResponseOrdersInnerWithDefaults instantiates a new OrderSearchResponseOrdersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderSearchResponseOrdersInnerWithDefaults() *OrderSearchResponseOrdersInner {
	this := OrderSearchResponseOrdersInner{}
	return &this
}

// GetIngramOrderNumber returns the IngramOrderNumber field value if set, zero value otherwise.
func (o *OrderSearchResponseOrdersInner) GetIngramOrderNumber() string {
	if o == nil || IsNil(o.IngramOrderNumber) {
		var ret string
		return ret
	}
	return *o.IngramOrderNumber
}

// GetIngramOrderNumberOk returns a tuple with the IngramOrderNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSearchResponseOrdersInner) GetIngramOrderNumberOk() (*string, bool) {
	if o == nil || IsNil(o.IngramOrderNumber) {
		return nil, false
	}
	return o.IngramOrderNumber, true
}

// HasIngramOrderNumber returns a boolean if a field has been set.
func (o *OrderSearchResponseOrdersInner) HasIngramOrderNumber() bool {
	if o != nil && !IsNil(o.IngramOrderNumber) {
		return true
	}

	return false
}

// SetIngramOrderNumber gets a reference to the given string and assigns it to the IngramOrderNumber field.
func (o *OrderSearchResponseOrdersInner) SetIngramOrderNumber(v string) {
	o.IngramOrderNumber = &v
}

// GetIngramOrderDate returns the IngramOrderDate field value if set, zero value otherwise.
func (o *OrderSearchResponseOrdersInner) GetIngramOrderDate() string {
	if o == nil || IsNil(o.IngramOrderDate) {
		var ret string
		return ret
	}
	return *o.IngramOrderDate
}

// GetIngramOrderDateOk returns a tuple with the IngramOrderDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSearchResponseOrdersInner) GetIngramOrderDateOk() (*string, bool) {
	if o == nil || IsNil(o.IngramOrderDate) {
		return nil, false
	}
	return o.IngramOrderDate, true
}

// HasIngramOrderDate returns a boolean if a field has been set.
func (o *OrderSearchResponseOrdersInner) HasIngramOrderDate() bool {
	if o != nil && !IsNil(o.IngramOrderDate) {
		return true
	}

	return false
}

// SetIngramOrderDate gets a reference to the given string and assigns it to the IngramOrderDate field.
func (o *OrderSearchResponseOrdersInner) SetIngramOrderDate(v string) {
	o.IngramOrderDate = &v
}

// GetCustomerOrderNumber returns the CustomerOrderNumber field value if set, zero value otherwise.
func (o *OrderSearchResponseOrdersInner) GetCustomerOrderNumber() string {
	if o == nil || IsNil(o.CustomerOrderNumber) {
		var ret string
		return ret
	}
	return *o.CustomerOrderNumber
}

// GetCustomerOrderNumberOk returns a tuple with the CustomerOrderNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSearchResponseOrdersInner) GetCustomerOrderNumberOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerOrderNumber) {
		return nil, false
	}
	return o.CustomerOrderNumber, true
}

// HasCustomerOrderNumber returns a boolean if a field has been set.
func (o *OrderSearchResponseOrdersInner) HasCustomerOrderNumber() bool {
	if o != nil && !IsNil(o.CustomerOrderNumber) {
		return true
	}

	return false
}

// SetCustomerOrderNumber gets a reference to the given string and assigns it to the CustomerOrderNumber field.
func (o *OrderSearchResponseOrdersInner) SetCustomerOrderNumber(v string) {
	o.CustomerOrderNumber = &v
}

// GetVendorSalesOrderNumber returns the VendorSalesOrderNumber field value if set, zero value otherwise.
func (o *OrderSearchResponseOrdersInner) GetVendorSalesOrderNumber() string {
	if o == nil || IsNil(o.VendorSalesOrderNumber) {
		var ret string
		return ret
	}
	return *o.VendorSalesOrderNumber
}

// GetVendorSalesOrderNumberOk returns a tuple with the VendorSalesOrderNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSearchResponseOrdersInner) GetVendorSalesOrderNumberOk() (*string, bool) {
	if o == nil || IsNil(o.VendorSalesOrderNumber) {
		return nil, false
	}
	return o.VendorSalesOrderNumber, true
}

// HasVendorSalesOrderNumber returns a boolean if a field has been set.
func (o *OrderSearchResponseOrdersInner) HasVendorSalesOrderNumber() bool {
	if o != nil && !IsNil(o.VendorSalesOrderNumber) {
		return true
	}

	return false
}

// SetVendorSalesOrderNumber gets a reference to the given string and assigns it to the VendorSalesOrderNumber field.
func (o *OrderSearchResponseOrdersInner) SetVendorSalesOrderNumber(v string) {
	o.VendorSalesOrderNumber = &v
}

// GetVendorName returns the VendorName field value if set, zero value otherwise.
func (o *OrderSearchResponseOrdersInner) GetVendorName() string {
	if o == nil || IsNil(o.VendorName) {
		var ret string
		return ret
	}
	return *o.VendorName
}

// GetVendorNameOk returns a tuple with the VendorName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSearchResponseOrdersInner) GetVendorNameOk() (*string, bool) {
	if o == nil || IsNil(o.VendorName) {
		return nil, false
	}
	return o.VendorName, true
}

// HasVendorName returns a boolean if a field has been set.
func (o *OrderSearchResponseOrdersInner) HasVendorName() bool {
	if o != nil && !IsNil(o.VendorName) {
		return true
	}

	return false
}

// SetVendorName gets a reference to the given string and assigns it to the VendorName field.
func (o *OrderSearchResponseOrdersInner) SetVendorName(v string) {
	o.VendorName = &v
}

// GetEndUserCompanyName returns the EndUserCompanyName field value if set, zero value otherwise.
func (o *OrderSearchResponseOrdersInner) GetEndUserCompanyName() string {
	if o == nil || IsNil(o.EndUserCompanyName) {
		var ret string
		return ret
	}
	return *o.EndUserCompanyName
}

// GetEndUserCompanyNameOk returns a tuple with the EndUserCompanyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSearchResponseOrdersInner) GetEndUserCompanyNameOk() (*string, bool) {
	if o == nil || IsNil(o.EndUserCompanyName) {
		return nil, false
	}
	return o.EndUserCompanyName, true
}

// HasEndUserCompanyName returns a boolean if a field has been set.
func (o *OrderSearchResponseOrdersInner) HasEndUserCompanyName() bool {
	if o != nil && !IsNil(o.EndUserCompanyName) {
		return true
	}

	return false
}

// SetEndUserCompanyName gets a reference to the given string and assigns it to the EndUserCompanyName field.
func (o *OrderSearchResponseOrdersInner) SetEndUserCompanyName(v string) {
	o.EndUserCompanyName = &v
}

// GetOrderTotal returns the OrderTotal field value if set, zero value otherwise.
func (o *OrderSearchResponseOrdersInner) GetOrderTotal() float32 {
	if o == nil || IsNil(o.OrderTotal) {
		var ret float32
		return ret
	}
	return *o.OrderTotal
}

// GetOrderTotalOk returns a tuple with the OrderTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSearchResponseOrdersInner) GetOrderTotalOk() (*float32, bool) {
	if o == nil || IsNil(o.OrderTotal) {
		return nil, false
	}
	return o.OrderTotal, true
}

// HasOrderTotal returns a boolean if a field has been set.
func (o *OrderSearchResponseOrdersInner) HasOrderTotal() bool {
	if o != nil && !IsNil(o.OrderTotal) {
		return true
	}

	return false
}

// SetOrderTotal gets a reference to the given float32 and assigns it to the OrderTotal field.
func (o *OrderSearchResponseOrdersInner) SetOrderTotal(v float32) {
	o.OrderTotal = &v
}

// GetOrderStatus returns the OrderStatus field value if set, zero value otherwise.
func (o *OrderSearchResponseOrdersInner) GetOrderStatus() string {
	if o == nil || IsNil(o.OrderStatus) {
		var ret string
		return ret
	}
	return *o.OrderStatus
}

// GetOrderStatusOk returns a tuple with the OrderStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSearchResponseOrdersInner) GetOrderStatusOk() (*string, bool) {
	if o == nil || IsNil(o.OrderStatus) {
		return nil, false
	}
	return o.OrderStatus, true
}

// HasOrderStatus returns a boolean if a field has been set.
func (o *OrderSearchResponseOrdersInner) HasOrderStatus() bool {
	if o != nil && !IsNil(o.OrderStatus) {
		return true
	}

	return false
}

// SetOrderStatus gets a reference to the given string and assigns it to the OrderStatus field.
func (o *OrderSearchResponseOrdersInner) SetOrderStatus(v string) {
	o.OrderStatus = &v
}

// GetSubOrders returns the SubOrders field value if set, zero value otherwise.
func (o *OrderSearchResponseOrdersInner) GetSubOrders() []OrderSearchResponseOrdersInnerSubOrdersInner {
	if o == nil || IsNil(o.SubOrders) {
		var ret []OrderSearchResponseOrdersInnerSubOrdersInner
		return ret
	}
	return o.SubOrders
}

// GetSubOrdersOk returns a tuple with the SubOrders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSearchResponseOrdersInner) GetSubOrdersOk() ([]OrderSearchResponseOrdersInnerSubOrdersInner, bool) {
	if o == nil || IsNil(o.SubOrders) {
		return nil, false
	}
	return o.SubOrders, true
}

// HasSubOrders returns a boolean if a field has been set.
func (o *OrderSearchResponseOrdersInner) HasSubOrders() bool {
	if o != nil && !IsNil(o.SubOrders) {
		return true
	}

	return false
}

// SetSubOrders gets a reference to the given []OrderSearchResponseOrdersInnerSubOrdersInner and assigns it to the SubOrders field.
func (o *OrderSearchResponseOrdersInner) SetSubOrders(v []OrderSearchResponseOrdersInnerSubOrdersInner) {
	o.SubOrders = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *OrderSearchResponseOrdersInner) GetLinks() OrderSearchResponseOrdersInnerLinks {
	if o == nil || IsNil(o.Links) {
		var ret OrderSearchResponseOrdersInnerLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSearchResponseOrdersInner) GetLinksOk() (*OrderSearchResponseOrdersInnerLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *OrderSearchResponseOrdersInner) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given OrderSearchResponseOrdersInnerLinks and assigns it to the Links field.
func (o *OrderSearchResponseOrdersInner) SetLinks(v OrderSearchResponseOrdersInnerLinks) {
	o.Links = &v
}

func (o OrderSearchResponseOrdersInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderSearchResponseOrdersInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IngramOrderNumber) {
		toSerialize["ingramOrderNumber"] = o.IngramOrderNumber
	}
	if !IsNil(o.IngramOrderDate) {
		toSerialize["ingramOrderDate"] = o.IngramOrderDate
	}
	if !IsNil(o.CustomerOrderNumber) {
		toSerialize["customerOrderNumber"] = o.CustomerOrderNumber
	}
	if !IsNil(o.VendorSalesOrderNumber) {
		toSerialize["vendorSalesOrderNumber"] = o.VendorSalesOrderNumber
	}
	if !IsNil(o.VendorName) {
		toSerialize["vendorName"] = o.VendorName
	}
	if !IsNil(o.EndUserCompanyName) {
		toSerialize["endUserCompanyName"] = o.EndUserCompanyName
	}
	if !IsNil(o.OrderTotal) {
		toSerialize["orderTotal"] = o.OrderTotal
	}
	if !IsNil(o.OrderStatus) {
		toSerialize["orderStatus"] = o.OrderStatus
	}
	if !IsNil(o.SubOrders) {
		toSerialize["subOrders"] = o.SubOrders
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	return toSerialize, nil
}

type NullableOrderSearchResponseOrdersInner struct {
	value *OrderSearchResponseOrdersInner
	isSet bool
}

func (v NullableOrderSearchResponseOrdersInner) Get() *OrderSearchResponseOrdersInner {
	return v.value
}

func (v *NullableOrderSearchResponseOrdersInner) Set(val *OrderSearchResponseOrdersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderSearchResponseOrdersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderSearchResponseOrdersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderSearchResponseOrdersInner(val *OrderSearchResponseOrdersInner) *NullableOrderSearchResponseOrdersInner {
	return &NullableOrderSearchResponseOrdersInner{value: val, isSet: true}
}

func (v NullableOrderSearchResponseOrdersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderSearchResponseOrdersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


