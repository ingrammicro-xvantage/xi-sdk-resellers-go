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

// checks if the OrderDetailB2BLinesInnerMultipleShipmentsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderDetailB2BLinesInnerMultipleShipmentsInner{}

// OrderDetailB2BLinesInnerMultipleShipmentsInner struct for OrderDetailB2BLinesInnerMultipleShipmentsInner
type OrderDetailB2BLinesInnerMultipleShipmentsInner struct {
	// Line number.
	LineNumber *string `json:"lineNumber,omitempty"`
	// Requested quantity.
	RequestedQuantity *int32 `json:"requestedQuantity,omitempty"`
	// Confirmed quantity.
	ConfirmedQuantity *int32 `json:"confirmedQuantity,omitempty"`
	// Date type. Example Single or multiple dates.
	DateType *string `json:"dateType,omitempty"`
	DateRange *OrderDetailB2BLinesInnerEstimatedDatesInnerShipShipDateRange `json:"dateRange,omitempty"`
	// Source.
	Source *string `json:"source,omitempty"`
	// Description.
	Description *string `json:"description,omitempty"`
	// Date.
	Date *string `json:"date,omitempty"`
	// Delivery date.
	DeliveryDate *string `json:"deliveryDate,omitempty"`
}

// NewOrderDetailB2BLinesInnerMultipleShipmentsInner instantiates a new OrderDetailB2BLinesInnerMultipleShipmentsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderDetailB2BLinesInnerMultipleShipmentsInner() *OrderDetailB2BLinesInnerMultipleShipmentsInner {
	this := OrderDetailB2BLinesInnerMultipleShipmentsInner{}
	return &this
}

// NewOrderDetailB2BLinesInnerMultipleShipmentsInnerWithDefaults instantiates a new OrderDetailB2BLinesInnerMultipleShipmentsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderDetailB2BLinesInnerMultipleShipmentsInnerWithDefaults() *OrderDetailB2BLinesInnerMultipleShipmentsInner {
	this := OrderDetailB2BLinesInnerMultipleShipmentsInner{}
	return &this
}

// GetLineNumber returns the LineNumber field value if set, zero value otherwise.
func (o *OrderDetailB2BLinesInnerMultipleShipmentsInner) GetLineNumber() string {
	if o == nil || IsNil(o.LineNumber) {
		var ret string
		return ret
	}
	return *o.LineNumber
}

// GetLineNumberOk returns a tuple with the LineNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailB2BLinesInnerMultipleShipmentsInner) GetLineNumberOk() (*string, bool) {
	if o == nil || IsNil(o.LineNumber) {
		return nil, false
	}
	return o.LineNumber, true
}

// HasLineNumber returns a boolean if a field has been set.
func (o *OrderDetailB2BLinesInnerMultipleShipmentsInner) HasLineNumber() bool {
	if o != nil && !IsNil(o.LineNumber) {
		return true
	}

	return false
}

// SetLineNumber gets a reference to the given string and assigns it to the LineNumber field.
func (o *OrderDetailB2BLinesInnerMultipleShipmentsInner) SetLineNumber(v string) {
	o.LineNumber = &v
}

// GetRequestedQuantity returns the RequestedQuantity field value if set, zero value otherwise.
func (o *OrderDetailB2BLinesInnerMultipleShipmentsInner) GetRequestedQuantity() int32 {
	if o == nil || IsNil(o.RequestedQuantity) {
		var ret int32
		return ret
	}
	return *o.RequestedQuantity
}

// GetRequestedQuantityOk returns a tuple with the RequestedQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailB2BLinesInnerMultipleShipmentsInner) GetRequestedQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.RequestedQuantity) {
		return nil, false
	}
	return o.RequestedQuantity, true
}

// HasRequestedQuantity returns a boolean if a field has been set.
func (o *OrderDetailB2BLinesInnerMultipleShipmentsInner) HasRequestedQuantity() bool {
	if o != nil && !IsNil(o.RequestedQuantity) {
		return true
	}

	return false
}

// SetRequestedQuantity gets a reference to the given int32 and assigns it to the RequestedQuantity field.
func (o *OrderDetailB2BLinesInnerMultipleShipmentsInner) SetRequestedQuantity(v int32) {
	o.RequestedQuantity = &v
}

// GetConfirmedQuantity returns the ConfirmedQuantity field value if set, zero value otherwise.
func (o *OrderDetailB2BLinesInnerMultipleShipmentsInner) GetConfirmedQuantity() int32 {
	if o == nil || IsNil(o.ConfirmedQuantity) {
		var ret int32
		return ret
	}
	return *o.ConfirmedQuantity
}

// GetConfirmedQuantityOk returns a tuple with the ConfirmedQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailB2BLinesInnerMultipleShipmentsInner) GetConfirmedQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.ConfirmedQuantity) {
		return nil, false
	}
	return o.ConfirmedQuantity, true
}

// HasConfirmedQuantity returns a boolean if a field has been set.
func (o *OrderDetailB2BLinesInnerMultipleShipmentsInner) HasConfirmedQuantity() bool {
	if o != nil && !IsNil(o.ConfirmedQuantity) {
		return true
	}

	return false
}

// SetConfirmedQuantity gets a reference to the given int32 and assigns it to the ConfirmedQuantity field.
func (o *OrderDetailB2BLinesInnerMultipleShipmentsInner) SetConfirmedQuantity(v int32) {
	o.ConfirmedQuantity = &v
}

// GetDateType returns the DateType field value if set, zero value otherwise.
func (o *OrderDetailB2BLinesInnerMultipleShipmentsInner) GetDateType() string {
	if o == nil || IsNil(o.DateType) {
		var ret string
		return ret
	}
	return *o.DateType
}

// GetDateTypeOk returns a tuple with the DateType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailB2BLinesInnerMultipleShipmentsInner) GetDateTypeOk() (*string, bool) {
	if o == nil || IsNil(o.DateType) {
		return nil, false
	}
	return o.DateType, true
}

// HasDateType returns a boolean if a field has been set.
func (o *OrderDetailB2BLinesInnerMultipleShipmentsInner) HasDateType() bool {
	if o != nil && !IsNil(o.DateType) {
		return true
	}

	return false
}

// SetDateType gets a reference to the given string and assigns it to the DateType field.
func (o *OrderDetailB2BLinesInnerMultipleShipmentsInner) SetDateType(v string) {
	o.DateType = &v
}

// GetDateRange returns the DateRange field value if set, zero value otherwise.
func (o *OrderDetailB2BLinesInnerMultipleShipmentsInner) GetDateRange() OrderDetailB2BLinesInnerEstimatedDatesInnerShipShipDateRange {
	if o == nil || IsNil(o.DateRange) {
		var ret OrderDetailB2BLinesInnerEstimatedDatesInnerShipShipDateRange
		return ret
	}
	return *o.DateRange
}

// GetDateRangeOk returns a tuple with the DateRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailB2BLinesInnerMultipleShipmentsInner) GetDateRangeOk() (*OrderDetailB2BLinesInnerEstimatedDatesInnerShipShipDateRange, bool) {
	if o == nil || IsNil(o.DateRange) {
		return nil, false
	}
	return o.DateRange, true
}

// HasDateRange returns a boolean if a field has been set.
func (o *OrderDetailB2BLinesInnerMultipleShipmentsInner) HasDateRange() bool {
	if o != nil && !IsNil(o.DateRange) {
		return true
	}

	return false
}

// SetDateRange gets a reference to the given OrderDetailB2BLinesInnerEstimatedDatesInnerShipShipDateRange and assigns it to the DateRange field.
func (o *OrderDetailB2BLinesInnerMultipleShipmentsInner) SetDateRange(v OrderDetailB2BLinesInnerEstimatedDatesInnerShipShipDateRange) {
	o.DateRange = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *OrderDetailB2BLinesInnerMultipleShipmentsInner) GetSource() string {
	if o == nil || IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailB2BLinesInnerMultipleShipmentsInner) GetSourceOk() (*string, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *OrderDetailB2BLinesInnerMultipleShipmentsInner) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *OrderDetailB2BLinesInnerMultipleShipmentsInner) SetSource(v string) {
	o.Source = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *OrderDetailB2BLinesInnerMultipleShipmentsInner) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailB2BLinesInnerMultipleShipmentsInner) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *OrderDetailB2BLinesInnerMultipleShipmentsInner) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *OrderDetailB2BLinesInnerMultipleShipmentsInner) SetDescription(v string) {
	o.Description = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *OrderDetailB2BLinesInnerMultipleShipmentsInner) GetDate() string {
	if o == nil || IsNil(o.Date) {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailB2BLinesInnerMultipleShipmentsInner) GetDateOk() (*string, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *OrderDetailB2BLinesInnerMultipleShipmentsInner) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *OrderDetailB2BLinesInnerMultipleShipmentsInner) SetDate(v string) {
	o.Date = &v
}

// GetDeliveryDate returns the DeliveryDate field value if set, zero value otherwise.
func (o *OrderDetailB2BLinesInnerMultipleShipmentsInner) GetDeliveryDate() string {
	if o == nil || IsNil(o.DeliveryDate) {
		var ret string
		return ret
	}
	return *o.DeliveryDate
}

// GetDeliveryDateOk returns a tuple with the DeliveryDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailB2BLinesInnerMultipleShipmentsInner) GetDeliveryDateOk() (*string, bool) {
	if o == nil || IsNil(o.DeliveryDate) {
		return nil, false
	}
	return o.DeliveryDate, true
}

// HasDeliveryDate returns a boolean if a field has been set.
func (o *OrderDetailB2BLinesInnerMultipleShipmentsInner) HasDeliveryDate() bool {
	if o != nil && !IsNil(o.DeliveryDate) {
		return true
	}

	return false
}

// SetDeliveryDate gets a reference to the given string and assigns it to the DeliveryDate field.
func (o *OrderDetailB2BLinesInnerMultipleShipmentsInner) SetDeliveryDate(v string) {
	o.DeliveryDate = &v
}

func (o OrderDetailB2BLinesInnerMultipleShipmentsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderDetailB2BLinesInnerMultipleShipmentsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LineNumber) {
		toSerialize["lineNumber"] = o.LineNumber
	}
	if !IsNil(o.RequestedQuantity) {
		toSerialize["requestedQuantity"] = o.RequestedQuantity
	}
	if !IsNil(o.ConfirmedQuantity) {
		toSerialize["confirmedQuantity"] = o.ConfirmedQuantity
	}
	if !IsNil(o.DateType) {
		toSerialize["dateType"] = o.DateType
	}
	if !IsNil(o.DateRange) {
		toSerialize["dateRange"] = o.DateRange
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Date) {
		toSerialize["date"] = o.Date
	}
	if !IsNil(o.DeliveryDate) {
		toSerialize["deliveryDate"] = o.DeliveryDate
	}
	return toSerialize, nil
}

type NullableOrderDetailB2BLinesInnerMultipleShipmentsInner struct {
	value *OrderDetailB2BLinesInnerMultipleShipmentsInner
	isSet bool
}

func (v NullableOrderDetailB2BLinesInnerMultipleShipmentsInner) Get() *OrderDetailB2BLinesInnerMultipleShipmentsInner {
	return v.value
}

func (v *NullableOrderDetailB2BLinesInnerMultipleShipmentsInner) Set(val *OrderDetailB2BLinesInnerMultipleShipmentsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderDetailB2BLinesInnerMultipleShipmentsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderDetailB2BLinesInnerMultipleShipmentsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderDetailB2BLinesInnerMultipleShipmentsInner(val *OrderDetailB2BLinesInnerMultipleShipmentsInner) *NullableOrderDetailB2BLinesInnerMultipleShipmentsInner {
	return &NullableOrderDetailB2BLinesInnerMultipleShipmentsInner{value: val, isSet: true}
}

func (v NullableOrderDetailB2BLinesInnerMultipleShipmentsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderDetailB2BLinesInnerMultipleShipmentsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


