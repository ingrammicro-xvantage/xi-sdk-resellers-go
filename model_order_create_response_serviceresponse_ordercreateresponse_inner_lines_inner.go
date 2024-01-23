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

// checks if the OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner{}

// OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner struct for OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner
type OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner struct {
	// “P”-Line or SKU Number “C”-Comment Line
	Linetype *string `json:"linetype,omitempty"`
	// Ingram generated line number
	Globallinenumber *string `json:"globallinenumber,omitempty"`
	// Ingram Micro Sku Number
	Partnumber *string `json:"partnumber,omitempty"`
	Globalskuid *string `json:"globalskuid,omitempty"`
	Linenumber *string `json:"linenumber,omitempty"`
	// Transportation 2 digit codes
	Carriercode *string `json:"carriercode,omitempty"`
	// Transportation Carrier Name
	Carrierdescription *string `json:"carrierdescription,omitempty"`
	// Price requested by reseller. Price Variance can be set up by Ingram Micro Sales Rep
	Requestedunitprice *float32 `json:"requestedunitprice,omitempty"`
	// Quanity Requested
	Requestedquantity *int32 `json:"requestedquantity,omitempty"`
	// Quanity Shipped
	Confirmedquantity *int32 `json:"confirmedquantity,omitempty"`
	// Quanity of units that didn’t ship
	Backorderedquantity *int32 `json:"backorderedquantity,omitempty"`
	// Price Per Unit
	Unitproductprice *float32 `json:"unitproductprice,omitempty"`
	// Total amount. Quantity X Unit Price
	Netamount *float32 `json:"netamount,omitempty"`
	Warehouseid *string `json:"warehouseid,omitempty"`
	// Use order suffix with the globalorderid for this line item.
	Ordersuffix *string `json:"ordersuffix,omitempty"`
}

// NewOrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner instantiates a new OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner() *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner {
	this := OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner{}
	return &this
}

// NewOrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInnerWithDefaults instantiates a new OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInnerWithDefaults() *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner {
	this := OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner{}
	return &this
}

// GetLinetype returns the Linetype field value if set, zero value otherwise.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) GetLinetype() string {
	if o == nil || IsNil(o.Linetype) {
		var ret string
		return ret
	}
	return *o.Linetype
}

// GetLinetypeOk returns a tuple with the Linetype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) GetLinetypeOk() (*string, bool) {
	if o == nil || IsNil(o.Linetype) {
		return nil, false
	}
	return o.Linetype, true
}

// HasLinetype returns a boolean if a field has been set.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) HasLinetype() bool {
	if o != nil && !IsNil(o.Linetype) {
		return true
	}

	return false
}

// SetLinetype gets a reference to the given string and assigns it to the Linetype field.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) SetLinetype(v string) {
	o.Linetype = &v
}

// GetGloballinenumber returns the Globallinenumber field value if set, zero value otherwise.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) GetGloballinenumber() string {
	if o == nil || IsNil(o.Globallinenumber) {
		var ret string
		return ret
	}
	return *o.Globallinenumber
}

// GetGloballinenumberOk returns a tuple with the Globallinenumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) GetGloballinenumberOk() (*string, bool) {
	if o == nil || IsNil(o.Globallinenumber) {
		return nil, false
	}
	return o.Globallinenumber, true
}

// HasGloballinenumber returns a boolean if a field has been set.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) HasGloballinenumber() bool {
	if o != nil && !IsNil(o.Globallinenumber) {
		return true
	}

	return false
}

// SetGloballinenumber gets a reference to the given string and assigns it to the Globallinenumber field.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) SetGloballinenumber(v string) {
	o.Globallinenumber = &v
}

// GetPartnumber returns the Partnumber field value if set, zero value otherwise.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) GetPartnumber() string {
	if o == nil || IsNil(o.Partnumber) {
		var ret string
		return ret
	}
	return *o.Partnumber
}

// GetPartnumberOk returns a tuple with the Partnumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) GetPartnumberOk() (*string, bool) {
	if o == nil || IsNil(o.Partnumber) {
		return nil, false
	}
	return o.Partnumber, true
}

// HasPartnumber returns a boolean if a field has been set.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) HasPartnumber() bool {
	if o != nil && !IsNil(o.Partnumber) {
		return true
	}

	return false
}

// SetPartnumber gets a reference to the given string and assigns it to the Partnumber field.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) SetPartnumber(v string) {
	o.Partnumber = &v
}

// GetGlobalskuid returns the Globalskuid field value if set, zero value otherwise.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) GetGlobalskuid() string {
	if o == nil || IsNil(o.Globalskuid) {
		var ret string
		return ret
	}
	return *o.Globalskuid
}

// GetGlobalskuidOk returns a tuple with the Globalskuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) GetGlobalskuidOk() (*string, bool) {
	if o == nil || IsNil(o.Globalskuid) {
		return nil, false
	}
	return o.Globalskuid, true
}

// HasGlobalskuid returns a boolean if a field has been set.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) HasGlobalskuid() bool {
	if o != nil && !IsNil(o.Globalskuid) {
		return true
	}

	return false
}

// SetGlobalskuid gets a reference to the given string and assigns it to the Globalskuid field.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) SetGlobalskuid(v string) {
	o.Globalskuid = &v
}

// GetLinenumber returns the Linenumber field value if set, zero value otherwise.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) GetLinenumber() string {
	if o == nil || IsNil(o.Linenumber) {
		var ret string
		return ret
	}
	return *o.Linenumber
}

// GetLinenumberOk returns a tuple with the Linenumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) GetLinenumberOk() (*string, bool) {
	if o == nil || IsNil(o.Linenumber) {
		return nil, false
	}
	return o.Linenumber, true
}

// HasLinenumber returns a boolean if a field has been set.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) HasLinenumber() bool {
	if o != nil && !IsNil(o.Linenumber) {
		return true
	}

	return false
}

// SetLinenumber gets a reference to the given string and assigns it to the Linenumber field.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) SetLinenumber(v string) {
	o.Linenumber = &v
}

// GetCarriercode returns the Carriercode field value if set, zero value otherwise.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) GetCarriercode() string {
	if o == nil || IsNil(o.Carriercode) {
		var ret string
		return ret
	}
	return *o.Carriercode
}

// GetCarriercodeOk returns a tuple with the Carriercode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) GetCarriercodeOk() (*string, bool) {
	if o == nil || IsNil(o.Carriercode) {
		return nil, false
	}
	return o.Carriercode, true
}

// HasCarriercode returns a boolean if a field has been set.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) HasCarriercode() bool {
	if o != nil && !IsNil(o.Carriercode) {
		return true
	}

	return false
}

// SetCarriercode gets a reference to the given string and assigns it to the Carriercode field.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) SetCarriercode(v string) {
	o.Carriercode = &v
}

// GetCarrierdescription returns the Carrierdescription field value if set, zero value otherwise.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) GetCarrierdescription() string {
	if o == nil || IsNil(o.Carrierdescription) {
		var ret string
		return ret
	}
	return *o.Carrierdescription
}

// GetCarrierdescriptionOk returns a tuple with the Carrierdescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) GetCarrierdescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Carrierdescription) {
		return nil, false
	}
	return o.Carrierdescription, true
}

// HasCarrierdescription returns a boolean if a field has been set.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) HasCarrierdescription() bool {
	if o != nil && !IsNil(o.Carrierdescription) {
		return true
	}

	return false
}

// SetCarrierdescription gets a reference to the given string and assigns it to the Carrierdescription field.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) SetCarrierdescription(v string) {
	o.Carrierdescription = &v
}

// GetRequestedunitprice returns the Requestedunitprice field value if set, zero value otherwise.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) GetRequestedunitprice() float32 {
	if o == nil || IsNil(o.Requestedunitprice) {
		var ret float32
		return ret
	}
	return *o.Requestedunitprice
}

// GetRequestedunitpriceOk returns a tuple with the Requestedunitprice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) GetRequestedunitpriceOk() (*float32, bool) {
	if o == nil || IsNil(o.Requestedunitprice) {
		return nil, false
	}
	return o.Requestedunitprice, true
}

// HasRequestedunitprice returns a boolean if a field has been set.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) HasRequestedunitprice() bool {
	if o != nil && !IsNil(o.Requestedunitprice) {
		return true
	}

	return false
}

// SetRequestedunitprice gets a reference to the given float32 and assigns it to the Requestedunitprice field.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) SetRequestedunitprice(v float32) {
	o.Requestedunitprice = &v
}

// GetRequestedquantity returns the Requestedquantity field value if set, zero value otherwise.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) GetRequestedquantity() int32 {
	if o == nil || IsNil(o.Requestedquantity) {
		var ret int32
		return ret
	}
	return *o.Requestedquantity
}

// GetRequestedquantityOk returns a tuple with the Requestedquantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) GetRequestedquantityOk() (*int32, bool) {
	if o == nil || IsNil(o.Requestedquantity) {
		return nil, false
	}
	return o.Requestedquantity, true
}

// HasRequestedquantity returns a boolean if a field has been set.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) HasRequestedquantity() bool {
	if o != nil && !IsNil(o.Requestedquantity) {
		return true
	}

	return false
}

// SetRequestedquantity gets a reference to the given int32 and assigns it to the Requestedquantity field.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) SetRequestedquantity(v int32) {
	o.Requestedquantity = &v
}

// GetConfirmedquantity returns the Confirmedquantity field value if set, zero value otherwise.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) GetConfirmedquantity() int32 {
	if o == nil || IsNil(o.Confirmedquantity) {
		var ret int32
		return ret
	}
	return *o.Confirmedquantity
}

// GetConfirmedquantityOk returns a tuple with the Confirmedquantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) GetConfirmedquantityOk() (*int32, bool) {
	if o == nil || IsNil(o.Confirmedquantity) {
		return nil, false
	}
	return o.Confirmedquantity, true
}

// HasConfirmedquantity returns a boolean if a field has been set.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) HasConfirmedquantity() bool {
	if o != nil && !IsNil(o.Confirmedquantity) {
		return true
	}

	return false
}

// SetConfirmedquantity gets a reference to the given int32 and assigns it to the Confirmedquantity field.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) SetConfirmedquantity(v int32) {
	o.Confirmedquantity = &v
}

// GetBackorderedquantity returns the Backorderedquantity field value if set, zero value otherwise.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) GetBackorderedquantity() int32 {
	if o == nil || IsNil(o.Backorderedquantity) {
		var ret int32
		return ret
	}
	return *o.Backorderedquantity
}

// GetBackorderedquantityOk returns a tuple with the Backorderedquantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) GetBackorderedquantityOk() (*int32, bool) {
	if o == nil || IsNil(o.Backorderedquantity) {
		return nil, false
	}
	return o.Backorderedquantity, true
}

// HasBackorderedquantity returns a boolean if a field has been set.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) HasBackorderedquantity() bool {
	if o != nil && !IsNil(o.Backorderedquantity) {
		return true
	}

	return false
}

// SetBackorderedquantity gets a reference to the given int32 and assigns it to the Backorderedquantity field.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) SetBackorderedquantity(v int32) {
	o.Backorderedquantity = &v
}

// GetUnitproductprice returns the Unitproductprice field value if set, zero value otherwise.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) GetUnitproductprice() float32 {
	if o == nil || IsNil(o.Unitproductprice) {
		var ret float32
		return ret
	}
	return *o.Unitproductprice
}

// GetUnitproductpriceOk returns a tuple with the Unitproductprice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) GetUnitproductpriceOk() (*float32, bool) {
	if o == nil || IsNil(o.Unitproductprice) {
		return nil, false
	}
	return o.Unitproductprice, true
}

// HasUnitproductprice returns a boolean if a field has been set.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) HasUnitproductprice() bool {
	if o != nil && !IsNil(o.Unitproductprice) {
		return true
	}

	return false
}

// SetUnitproductprice gets a reference to the given float32 and assigns it to the Unitproductprice field.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) SetUnitproductprice(v float32) {
	o.Unitproductprice = &v
}

// GetNetamount returns the Netamount field value if set, zero value otherwise.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) GetNetamount() float32 {
	if o == nil || IsNil(o.Netamount) {
		var ret float32
		return ret
	}
	return *o.Netamount
}

// GetNetamountOk returns a tuple with the Netamount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) GetNetamountOk() (*float32, bool) {
	if o == nil || IsNil(o.Netamount) {
		return nil, false
	}
	return o.Netamount, true
}

// HasNetamount returns a boolean if a field has been set.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) HasNetamount() bool {
	if o != nil && !IsNil(o.Netamount) {
		return true
	}

	return false
}

// SetNetamount gets a reference to the given float32 and assigns it to the Netamount field.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) SetNetamount(v float32) {
	o.Netamount = &v
}

// GetWarehouseid returns the Warehouseid field value if set, zero value otherwise.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) GetWarehouseid() string {
	if o == nil || IsNil(o.Warehouseid) {
		var ret string
		return ret
	}
	return *o.Warehouseid
}

// GetWarehouseidOk returns a tuple with the Warehouseid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) GetWarehouseidOk() (*string, bool) {
	if o == nil || IsNil(o.Warehouseid) {
		return nil, false
	}
	return o.Warehouseid, true
}

// HasWarehouseid returns a boolean if a field has been set.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) HasWarehouseid() bool {
	if o != nil && !IsNil(o.Warehouseid) {
		return true
	}

	return false
}

// SetWarehouseid gets a reference to the given string and assigns it to the Warehouseid field.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) SetWarehouseid(v string) {
	o.Warehouseid = &v
}

// GetOrdersuffix returns the Ordersuffix field value if set, zero value otherwise.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) GetOrdersuffix() string {
	if o == nil || IsNil(o.Ordersuffix) {
		var ret string
		return ret
	}
	return *o.Ordersuffix
}

// GetOrdersuffixOk returns a tuple with the Ordersuffix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) GetOrdersuffixOk() (*string, bool) {
	if o == nil || IsNil(o.Ordersuffix) {
		return nil, false
	}
	return o.Ordersuffix, true
}

// HasOrdersuffix returns a boolean if a field has been set.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) HasOrdersuffix() bool {
	if o != nil && !IsNil(o.Ordersuffix) {
		return true
	}

	return false
}

// SetOrdersuffix gets a reference to the given string and assigns it to the Ordersuffix field.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) SetOrdersuffix(v string) {
	o.Ordersuffix = &v
}

func (o OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Linetype) {
		toSerialize["linetype"] = o.Linetype
	}
	if !IsNil(o.Globallinenumber) {
		toSerialize["globallinenumber"] = o.Globallinenumber
	}
	if !IsNil(o.Partnumber) {
		toSerialize["partnumber"] = o.Partnumber
	}
	if !IsNil(o.Globalskuid) {
		toSerialize["globalskuid"] = o.Globalskuid
	}
	if !IsNil(o.Linenumber) {
		toSerialize["linenumber"] = o.Linenumber
	}
	if !IsNil(o.Carriercode) {
		toSerialize["carriercode"] = o.Carriercode
	}
	if !IsNil(o.Carrierdescription) {
		toSerialize["carrierdescription"] = o.Carrierdescription
	}
	if !IsNil(o.Requestedunitprice) {
		toSerialize["requestedunitprice"] = o.Requestedunitprice
	}
	if !IsNil(o.Requestedquantity) {
		toSerialize["requestedquantity"] = o.Requestedquantity
	}
	if !IsNil(o.Confirmedquantity) {
		toSerialize["confirmedquantity"] = o.Confirmedquantity
	}
	if !IsNil(o.Backorderedquantity) {
		toSerialize["backorderedquantity"] = o.Backorderedquantity
	}
	if !IsNil(o.Unitproductprice) {
		toSerialize["unitproductprice"] = o.Unitproductprice
	}
	if !IsNil(o.Netamount) {
		toSerialize["netamount"] = o.Netamount
	}
	if !IsNil(o.Warehouseid) {
		toSerialize["warehouseid"] = o.Warehouseid
	}
	if !IsNil(o.Ordersuffix) {
		toSerialize["ordersuffix"] = o.Ordersuffix
	}
	return toSerialize, nil
}

type NullableOrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner struct {
	value *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner
	isSet bool
}

func (v NullableOrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) Get() *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner {
	return v.value
}

func (v *NullableOrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) Set(val *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner(val *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) *NullableOrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner {
	return &NullableOrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner{value: val, isSet: true}
}

func (v NullableOrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


