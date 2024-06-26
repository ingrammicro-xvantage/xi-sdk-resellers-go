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

// checks if the OrderStatusAsyncNotificationRequestResourceInnerLinesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderStatusAsyncNotificationRequestResourceInnerLinesInner{}

// OrderStatusAsyncNotificationRequestResourceInnerLinesInner struct for OrderStatusAsyncNotificationRequestResourceInnerLinesInner
type OrderStatusAsyncNotificationRequestResourceInnerLinesInner struct {
	// The Ingram Micro line number for the product
	LineNumber *string `json:"LineNumber,omitempty"`
	// The sub order number. The two-digit prefix is the warehouse code of the warehouse nearest the reseller. The middle number is the order number. The two-digit suffix is the sub order number.
	SubOrderNumber *string `json:"subOrderNumber,omitempty"`
	// The status for the line item in the order. One of: Backordered, Open, Shipped
	LineStatus *string `json:"lineStatus,omitempty"`
	// The Ingram Micro part number for the line item.
	IngramPartNumber *string `json:"ingramPartNumber,omitempty"`
	// The vendor part number for the line item.
	VendorPartNumber *string `json:"vendorPartNumber,omitempty"`
	// The quantity of the line item requested.
	RequestedQuantity *string `json:"requestedQuantity,omitempty"`
	// The quantity of the line item that has been shipped.
	ShippedQuantity *string `json:"shippedQuantity,omitempty"`
	// The quantity of the line item that is backordered.
	BackorderedQuantity *string `json:"backorderedQuantity,omitempty"`
	ShipmentDetails []OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner `json:"shipmentDetails,omitempty"`
	SerialNumberDetails []OrderStatusAsyncNotificationRequestResourceInnerLinesInnerSerialNumberDetailsInner `json:"serialNumberDetails,omitempty"`
}

// NewOrderStatusAsyncNotificationRequestResourceInnerLinesInner instantiates a new OrderStatusAsyncNotificationRequestResourceInnerLinesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderStatusAsyncNotificationRequestResourceInnerLinesInner() *OrderStatusAsyncNotificationRequestResourceInnerLinesInner {
	this := OrderStatusAsyncNotificationRequestResourceInnerLinesInner{}
	return &this
}

// NewOrderStatusAsyncNotificationRequestResourceInnerLinesInnerWithDefaults instantiates a new OrderStatusAsyncNotificationRequestResourceInnerLinesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderStatusAsyncNotificationRequestResourceInnerLinesInnerWithDefaults() *OrderStatusAsyncNotificationRequestResourceInnerLinesInner {
	this := OrderStatusAsyncNotificationRequestResourceInnerLinesInner{}
	return &this
}

// GetLineNumber returns the LineNumber field value if set, zero value otherwise.
func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInner) GetLineNumber() string {
	if o == nil || IsNil(o.LineNumber) {
		var ret string
		return ret
	}
	return *o.LineNumber
}

// GetLineNumberOk returns a tuple with the LineNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInner) GetLineNumberOk() (*string, bool) {
	if o == nil || IsNil(o.LineNumber) {
		return nil, false
	}
	return o.LineNumber, true
}

// HasLineNumber returns a boolean if a field has been set.
func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInner) HasLineNumber() bool {
	if o != nil && !IsNil(o.LineNumber) {
		return true
	}

	return false
}

// SetLineNumber gets a reference to the given string and assigns it to the LineNumber field.
func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInner) SetLineNumber(v string) {
	o.LineNumber = &v
}

// GetSubOrderNumber returns the SubOrderNumber field value if set, zero value otherwise.
func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInner) GetSubOrderNumber() string {
	if o == nil || IsNil(o.SubOrderNumber) {
		var ret string
		return ret
	}
	return *o.SubOrderNumber
}

// GetSubOrderNumberOk returns a tuple with the SubOrderNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInner) GetSubOrderNumberOk() (*string, bool) {
	if o == nil || IsNil(o.SubOrderNumber) {
		return nil, false
	}
	return o.SubOrderNumber, true
}

// HasSubOrderNumber returns a boolean if a field has been set.
func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInner) HasSubOrderNumber() bool {
	if o != nil && !IsNil(o.SubOrderNumber) {
		return true
	}

	return false
}

// SetSubOrderNumber gets a reference to the given string and assigns it to the SubOrderNumber field.
func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInner) SetSubOrderNumber(v string) {
	o.SubOrderNumber = &v
}

// GetLineStatus returns the LineStatus field value if set, zero value otherwise.
func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInner) GetLineStatus() string {
	if o == nil || IsNil(o.LineStatus) {
		var ret string
		return ret
	}
	return *o.LineStatus
}

// GetLineStatusOk returns a tuple with the LineStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInner) GetLineStatusOk() (*string, bool) {
	if o == nil || IsNil(o.LineStatus) {
		return nil, false
	}
	return o.LineStatus, true
}

// HasLineStatus returns a boolean if a field has been set.
func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInner) HasLineStatus() bool {
	if o != nil && !IsNil(o.LineStatus) {
		return true
	}

	return false
}

// SetLineStatus gets a reference to the given string and assigns it to the LineStatus field.
func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInner) SetLineStatus(v string) {
	o.LineStatus = &v
}

// GetIngramPartNumber returns the IngramPartNumber field value if set, zero value otherwise.
func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInner) GetIngramPartNumber() string {
	if o == nil || IsNil(o.IngramPartNumber) {
		var ret string
		return ret
	}
	return *o.IngramPartNumber
}

// GetIngramPartNumberOk returns a tuple with the IngramPartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInner) GetIngramPartNumberOk() (*string, bool) {
	if o == nil || IsNil(o.IngramPartNumber) {
		return nil, false
	}
	return o.IngramPartNumber, true
}

// HasIngramPartNumber returns a boolean if a field has been set.
func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInner) HasIngramPartNumber() bool {
	if o != nil && !IsNil(o.IngramPartNumber) {
		return true
	}

	return false
}

// SetIngramPartNumber gets a reference to the given string and assigns it to the IngramPartNumber field.
func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInner) SetIngramPartNumber(v string) {
	o.IngramPartNumber = &v
}

// GetVendorPartNumber returns the VendorPartNumber field value if set, zero value otherwise.
func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInner) GetVendorPartNumber() string {
	if o == nil || IsNil(o.VendorPartNumber) {
		var ret string
		return ret
	}
	return *o.VendorPartNumber
}

// GetVendorPartNumberOk returns a tuple with the VendorPartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInner) GetVendorPartNumberOk() (*string, bool) {
	if o == nil || IsNil(o.VendorPartNumber) {
		return nil, false
	}
	return o.VendorPartNumber, true
}

// HasVendorPartNumber returns a boolean if a field has been set.
func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInner) HasVendorPartNumber() bool {
	if o != nil && !IsNil(o.VendorPartNumber) {
		return true
	}

	return false
}

// SetVendorPartNumber gets a reference to the given string and assigns it to the VendorPartNumber field.
func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInner) SetVendorPartNumber(v string) {
	o.VendorPartNumber = &v
}

// GetRequestedQuantity returns the RequestedQuantity field value if set, zero value otherwise.
func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInner) GetRequestedQuantity() string {
	if o == nil || IsNil(o.RequestedQuantity) {
		var ret string
		return ret
	}
	return *o.RequestedQuantity
}

// GetRequestedQuantityOk returns a tuple with the RequestedQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInner) GetRequestedQuantityOk() (*string, bool) {
	if o == nil || IsNil(o.RequestedQuantity) {
		return nil, false
	}
	return o.RequestedQuantity, true
}

// HasRequestedQuantity returns a boolean if a field has been set.
func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInner) HasRequestedQuantity() bool {
	if o != nil && !IsNil(o.RequestedQuantity) {
		return true
	}

	return false
}

// SetRequestedQuantity gets a reference to the given string and assigns it to the RequestedQuantity field.
func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInner) SetRequestedQuantity(v string) {
	o.RequestedQuantity = &v
}

// GetShippedQuantity returns the ShippedQuantity field value if set, zero value otherwise.
func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInner) GetShippedQuantity() string {
	if o == nil || IsNil(o.ShippedQuantity) {
		var ret string
		return ret
	}
	return *o.ShippedQuantity
}

// GetShippedQuantityOk returns a tuple with the ShippedQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInner) GetShippedQuantityOk() (*string, bool) {
	if o == nil || IsNil(o.ShippedQuantity) {
		return nil, false
	}
	return o.ShippedQuantity, true
}

// HasShippedQuantity returns a boolean if a field has been set.
func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInner) HasShippedQuantity() bool {
	if o != nil && !IsNil(o.ShippedQuantity) {
		return true
	}

	return false
}

// SetShippedQuantity gets a reference to the given string and assigns it to the ShippedQuantity field.
func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInner) SetShippedQuantity(v string) {
	o.ShippedQuantity = &v
}

// GetBackorderedQuantity returns the BackorderedQuantity field value if set, zero value otherwise.
func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInner) GetBackorderedQuantity() string {
	if o == nil || IsNil(o.BackorderedQuantity) {
		var ret string
		return ret
	}
	return *o.BackorderedQuantity
}

// GetBackorderedQuantityOk returns a tuple with the BackorderedQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInner) GetBackorderedQuantityOk() (*string, bool) {
	if o == nil || IsNil(o.BackorderedQuantity) {
		return nil, false
	}
	return o.BackorderedQuantity, true
}

// HasBackorderedQuantity returns a boolean if a field has been set.
func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInner) HasBackorderedQuantity() bool {
	if o != nil && !IsNil(o.BackorderedQuantity) {
		return true
	}

	return false
}

// SetBackorderedQuantity gets a reference to the given string and assigns it to the BackorderedQuantity field.
func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInner) SetBackorderedQuantity(v string) {
	o.BackorderedQuantity = &v
}

// GetShipmentDetails returns the ShipmentDetails field value if set, zero value otherwise.
func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInner) GetShipmentDetails() []OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner {
	if o == nil || IsNil(o.ShipmentDetails) {
		var ret []OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner
		return ret
	}
	return o.ShipmentDetails
}

// GetShipmentDetailsOk returns a tuple with the ShipmentDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInner) GetShipmentDetailsOk() ([]OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner, bool) {
	if o == nil || IsNil(o.ShipmentDetails) {
		return nil, false
	}
	return o.ShipmentDetails, true
}

// HasShipmentDetails returns a boolean if a field has been set.
func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInner) HasShipmentDetails() bool {
	if o != nil && !IsNil(o.ShipmentDetails) {
		return true
	}

	return false
}

// SetShipmentDetails gets a reference to the given []OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner and assigns it to the ShipmentDetails field.
func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInner) SetShipmentDetails(v []OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner) {
	o.ShipmentDetails = v
}

// GetSerialNumberDetails returns the SerialNumberDetails field value if set, zero value otherwise.
func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInner) GetSerialNumberDetails() []OrderStatusAsyncNotificationRequestResourceInnerLinesInnerSerialNumberDetailsInner {
	if o == nil || IsNil(o.SerialNumberDetails) {
		var ret []OrderStatusAsyncNotificationRequestResourceInnerLinesInnerSerialNumberDetailsInner
		return ret
	}
	return o.SerialNumberDetails
}

// GetSerialNumberDetailsOk returns a tuple with the SerialNumberDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInner) GetSerialNumberDetailsOk() ([]OrderStatusAsyncNotificationRequestResourceInnerLinesInnerSerialNumberDetailsInner, bool) {
	if o == nil || IsNil(o.SerialNumberDetails) {
		return nil, false
	}
	return o.SerialNumberDetails, true
}

// HasSerialNumberDetails returns a boolean if a field has been set.
func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInner) HasSerialNumberDetails() bool {
	if o != nil && !IsNil(o.SerialNumberDetails) {
		return true
	}

	return false
}

// SetSerialNumberDetails gets a reference to the given []OrderStatusAsyncNotificationRequestResourceInnerLinesInnerSerialNumberDetailsInner and assigns it to the SerialNumberDetails field.
func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInner) SetSerialNumberDetails(v []OrderStatusAsyncNotificationRequestResourceInnerLinesInnerSerialNumberDetailsInner) {
	o.SerialNumberDetails = v
}

func (o OrderStatusAsyncNotificationRequestResourceInnerLinesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderStatusAsyncNotificationRequestResourceInnerLinesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LineNumber) {
		toSerialize["LineNumber"] = o.LineNumber
	}
	if !IsNil(o.SubOrderNumber) {
		toSerialize["subOrderNumber"] = o.SubOrderNumber
	}
	if !IsNil(o.LineStatus) {
		toSerialize["lineStatus"] = o.LineStatus
	}
	if !IsNil(o.IngramPartNumber) {
		toSerialize["ingramPartNumber"] = o.IngramPartNumber
	}
	if !IsNil(o.VendorPartNumber) {
		toSerialize["vendorPartNumber"] = o.VendorPartNumber
	}
	if !IsNil(o.RequestedQuantity) {
		toSerialize["requestedQuantity"] = o.RequestedQuantity
	}
	if !IsNil(o.ShippedQuantity) {
		toSerialize["shippedQuantity"] = o.ShippedQuantity
	}
	if !IsNil(o.BackorderedQuantity) {
		toSerialize["backorderedQuantity"] = o.BackorderedQuantity
	}
	if !IsNil(o.ShipmentDetails) {
		toSerialize["shipmentDetails"] = o.ShipmentDetails
	}
	if !IsNil(o.SerialNumberDetails) {
		toSerialize["serialNumberDetails"] = o.SerialNumberDetails
	}
	return toSerialize, nil
}

type NullableOrderStatusAsyncNotificationRequestResourceInnerLinesInner struct {
	value *OrderStatusAsyncNotificationRequestResourceInnerLinesInner
	isSet bool
}

func (v NullableOrderStatusAsyncNotificationRequestResourceInnerLinesInner) Get() *OrderStatusAsyncNotificationRequestResourceInnerLinesInner {
	return v.value
}

func (v *NullableOrderStatusAsyncNotificationRequestResourceInnerLinesInner) Set(val *OrderStatusAsyncNotificationRequestResourceInnerLinesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderStatusAsyncNotificationRequestResourceInnerLinesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderStatusAsyncNotificationRequestResourceInnerLinesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderStatusAsyncNotificationRequestResourceInnerLinesInner(val *OrderStatusAsyncNotificationRequestResourceInnerLinesInner) *NullableOrderStatusAsyncNotificationRequestResourceInnerLinesInner {
	return &NullableOrderStatusAsyncNotificationRequestResourceInnerLinesInner{value: val, isSet: true}
}

func (v NullableOrderStatusAsyncNotificationRequestResourceInnerLinesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderStatusAsyncNotificationRequestResourceInnerLinesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


