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

// checks if the OrderModifyResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderModifyResponse{}

// OrderModifyResponse struct for OrderModifyResponse
type OrderModifyResponse struct {
	// The IngramMicro order number.
	IngramOrderNumber *string `json:"ingramOrderNumber,omitempty"`
	// The description of the change.
	ChangeDescription *string `json:"changeDescription,omitempty"`
	// The date the order was modified.
	OrderModifiedDate *string `json:"orderModifiedDate,omitempty"`
	// The reseller's order number for reference in their system.
	CustomerOrderNumber *string `json:"customerOrderNumber,omitempty"`
	// The end user/customer's order number for reference in their system.
	EndCustomerOrderNumber *string `json:"endCustomerOrderNumber,omitempty"`
	// The total for the order.
	OrderTotal *float32 `json:"orderTotal,omitempty"`
	// Order-level notes.
	Notes *string `json:"notes,omitempty"`
	// The sub total for the order.
	OrderSubTotal *float32 `json:"orderSubTotal,omitempty"`
	// The freight charges for the order.
	FreightCharges *float32 `json:"freightCharges,omitempty"`
	// The total tax for the order.
	TotalTax *float32 `json:"totalTax,omitempty"`
	// The status of the order. One of the following. Backordered, In Progress, Shipped, Delivered, Canceled, On Hold
	OrderStatus *string `json:"orderStatus,omitempty"`
	// Suffix used to identify billing address. Created during onboarding. Resellers are provided with one or more address IDs depending on how many bill to addresses they need for various flooring companies they are using for credit.
	BillToAddressId *string `json:"billToAddressId,omitempty"`
	ShipToInfo *OrderModifyResponseShipToInfo `json:"shipToInfo,omitempty"`
	// The line-level details for the order.
	Lines []OrderModifyResponseLinesInner `json:"lines,omitempty"`
	// Details for failed lines in the order.
	RejectedLineItems []OrderModifyResponseRejectedLineItemsInner `json:"rejectedLineItems,omitempty"`
	// Header-level additional attributes.
	AdditionalAttributes []OrderModifyResponseLinesInnerAdditionalAttributesInner `json:"additionalAttributes,omitempty"`
}

// NewOrderModifyResponse instantiates a new OrderModifyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderModifyResponse() *OrderModifyResponse {
	this := OrderModifyResponse{}
	return &this
}

// NewOrderModifyResponseWithDefaults instantiates a new OrderModifyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderModifyResponseWithDefaults() *OrderModifyResponse {
	this := OrderModifyResponse{}
	return &this
}

// GetIngramOrderNumber returns the IngramOrderNumber field value if set, zero value otherwise.
func (o *OrderModifyResponse) GetIngramOrderNumber() string {
	if o == nil || IsNil(o.IngramOrderNumber) {
		var ret string
		return ret
	}
	return *o.IngramOrderNumber
}

// GetIngramOrderNumberOk returns a tuple with the IngramOrderNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderModifyResponse) GetIngramOrderNumberOk() (*string, bool) {
	if o == nil || IsNil(o.IngramOrderNumber) {
		return nil, false
	}
	return o.IngramOrderNumber, true
}

// HasIngramOrderNumber returns a boolean if a field has been set.
func (o *OrderModifyResponse) HasIngramOrderNumber() bool {
	if o != nil && !IsNil(o.IngramOrderNumber) {
		return true
	}

	return false
}

// SetIngramOrderNumber gets a reference to the given string and assigns it to the IngramOrderNumber field.
func (o *OrderModifyResponse) SetIngramOrderNumber(v string) {
	o.IngramOrderNumber = &v
}

// GetChangeDescription returns the ChangeDescription field value if set, zero value otherwise.
func (o *OrderModifyResponse) GetChangeDescription() string {
	if o == nil || IsNil(o.ChangeDescription) {
		var ret string
		return ret
	}
	return *o.ChangeDescription
}

// GetChangeDescriptionOk returns a tuple with the ChangeDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderModifyResponse) GetChangeDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.ChangeDescription) {
		return nil, false
	}
	return o.ChangeDescription, true
}

// HasChangeDescription returns a boolean if a field has been set.
func (o *OrderModifyResponse) HasChangeDescription() bool {
	if o != nil && !IsNil(o.ChangeDescription) {
		return true
	}

	return false
}

// SetChangeDescription gets a reference to the given string and assigns it to the ChangeDescription field.
func (o *OrderModifyResponse) SetChangeDescription(v string) {
	o.ChangeDescription = &v
}

// GetOrderModifiedDate returns the OrderModifiedDate field value if set, zero value otherwise.
func (o *OrderModifyResponse) GetOrderModifiedDate() string {
	if o == nil || IsNil(o.OrderModifiedDate) {
		var ret string
		return ret
	}
	return *o.OrderModifiedDate
}

// GetOrderModifiedDateOk returns a tuple with the OrderModifiedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderModifyResponse) GetOrderModifiedDateOk() (*string, bool) {
	if o == nil || IsNil(o.OrderModifiedDate) {
		return nil, false
	}
	return o.OrderModifiedDate, true
}

// HasOrderModifiedDate returns a boolean if a field has been set.
func (o *OrderModifyResponse) HasOrderModifiedDate() bool {
	if o != nil && !IsNil(o.OrderModifiedDate) {
		return true
	}

	return false
}

// SetOrderModifiedDate gets a reference to the given string and assigns it to the OrderModifiedDate field.
func (o *OrderModifyResponse) SetOrderModifiedDate(v string) {
	o.OrderModifiedDate = &v
}

// GetCustomerOrderNumber returns the CustomerOrderNumber field value if set, zero value otherwise.
func (o *OrderModifyResponse) GetCustomerOrderNumber() string {
	if o == nil || IsNil(o.CustomerOrderNumber) {
		var ret string
		return ret
	}
	return *o.CustomerOrderNumber
}

// GetCustomerOrderNumberOk returns a tuple with the CustomerOrderNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderModifyResponse) GetCustomerOrderNumberOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerOrderNumber) {
		return nil, false
	}
	return o.CustomerOrderNumber, true
}

// HasCustomerOrderNumber returns a boolean if a field has been set.
func (o *OrderModifyResponse) HasCustomerOrderNumber() bool {
	if o != nil && !IsNil(o.CustomerOrderNumber) {
		return true
	}

	return false
}

// SetCustomerOrderNumber gets a reference to the given string and assigns it to the CustomerOrderNumber field.
func (o *OrderModifyResponse) SetCustomerOrderNumber(v string) {
	o.CustomerOrderNumber = &v
}

// GetEndCustomerOrderNumber returns the EndCustomerOrderNumber field value if set, zero value otherwise.
func (o *OrderModifyResponse) GetEndCustomerOrderNumber() string {
	if o == nil || IsNil(o.EndCustomerOrderNumber) {
		var ret string
		return ret
	}
	return *o.EndCustomerOrderNumber
}

// GetEndCustomerOrderNumberOk returns a tuple with the EndCustomerOrderNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderModifyResponse) GetEndCustomerOrderNumberOk() (*string, bool) {
	if o == nil || IsNil(o.EndCustomerOrderNumber) {
		return nil, false
	}
	return o.EndCustomerOrderNumber, true
}

// HasEndCustomerOrderNumber returns a boolean if a field has been set.
func (o *OrderModifyResponse) HasEndCustomerOrderNumber() bool {
	if o != nil && !IsNil(o.EndCustomerOrderNumber) {
		return true
	}

	return false
}

// SetEndCustomerOrderNumber gets a reference to the given string and assigns it to the EndCustomerOrderNumber field.
func (o *OrderModifyResponse) SetEndCustomerOrderNumber(v string) {
	o.EndCustomerOrderNumber = &v
}

// GetOrderTotal returns the OrderTotal field value if set, zero value otherwise.
func (o *OrderModifyResponse) GetOrderTotal() float32 {
	if o == nil || IsNil(o.OrderTotal) {
		var ret float32
		return ret
	}
	return *o.OrderTotal
}

// GetOrderTotalOk returns a tuple with the OrderTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderModifyResponse) GetOrderTotalOk() (*float32, bool) {
	if o == nil || IsNil(o.OrderTotal) {
		return nil, false
	}
	return o.OrderTotal, true
}

// HasOrderTotal returns a boolean if a field has been set.
func (o *OrderModifyResponse) HasOrderTotal() bool {
	if o != nil && !IsNil(o.OrderTotal) {
		return true
	}

	return false
}

// SetOrderTotal gets a reference to the given float32 and assigns it to the OrderTotal field.
func (o *OrderModifyResponse) SetOrderTotal(v float32) {
	o.OrderTotal = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *OrderModifyResponse) GetNotes() string {
	if o == nil || IsNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderModifyResponse) GetNotesOk() (*string, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *OrderModifyResponse) HasNotes() bool {
	if o != nil && !IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *OrderModifyResponse) SetNotes(v string) {
	o.Notes = &v
}

// GetOrderSubTotal returns the OrderSubTotal field value if set, zero value otherwise.
func (o *OrderModifyResponse) GetOrderSubTotal() float32 {
	if o == nil || IsNil(o.OrderSubTotal) {
		var ret float32
		return ret
	}
	return *o.OrderSubTotal
}

// GetOrderSubTotalOk returns a tuple with the OrderSubTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderModifyResponse) GetOrderSubTotalOk() (*float32, bool) {
	if o == nil || IsNil(o.OrderSubTotal) {
		return nil, false
	}
	return o.OrderSubTotal, true
}

// HasOrderSubTotal returns a boolean if a field has been set.
func (o *OrderModifyResponse) HasOrderSubTotal() bool {
	if o != nil && !IsNil(o.OrderSubTotal) {
		return true
	}

	return false
}

// SetOrderSubTotal gets a reference to the given float32 and assigns it to the OrderSubTotal field.
func (o *OrderModifyResponse) SetOrderSubTotal(v float32) {
	o.OrderSubTotal = &v
}

// GetFreightCharges returns the FreightCharges field value if set, zero value otherwise.
func (o *OrderModifyResponse) GetFreightCharges() float32 {
	if o == nil || IsNil(o.FreightCharges) {
		var ret float32
		return ret
	}
	return *o.FreightCharges
}

// GetFreightChargesOk returns a tuple with the FreightCharges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderModifyResponse) GetFreightChargesOk() (*float32, bool) {
	if o == nil || IsNil(o.FreightCharges) {
		return nil, false
	}
	return o.FreightCharges, true
}

// HasFreightCharges returns a boolean if a field has been set.
func (o *OrderModifyResponse) HasFreightCharges() bool {
	if o != nil && !IsNil(o.FreightCharges) {
		return true
	}

	return false
}

// SetFreightCharges gets a reference to the given float32 and assigns it to the FreightCharges field.
func (o *OrderModifyResponse) SetFreightCharges(v float32) {
	o.FreightCharges = &v
}

// GetTotalTax returns the TotalTax field value if set, zero value otherwise.
func (o *OrderModifyResponse) GetTotalTax() float32 {
	if o == nil || IsNil(o.TotalTax) {
		var ret float32
		return ret
	}
	return *o.TotalTax
}

// GetTotalTaxOk returns a tuple with the TotalTax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderModifyResponse) GetTotalTaxOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalTax) {
		return nil, false
	}
	return o.TotalTax, true
}

// HasTotalTax returns a boolean if a field has been set.
func (o *OrderModifyResponse) HasTotalTax() bool {
	if o != nil && !IsNil(o.TotalTax) {
		return true
	}

	return false
}

// SetTotalTax gets a reference to the given float32 and assigns it to the TotalTax field.
func (o *OrderModifyResponse) SetTotalTax(v float32) {
	o.TotalTax = &v
}

// GetOrderStatus returns the OrderStatus field value if set, zero value otherwise.
func (o *OrderModifyResponse) GetOrderStatus() string {
	if o == nil || IsNil(o.OrderStatus) {
		var ret string
		return ret
	}
	return *o.OrderStatus
}

// GetOrderStatusOk returns a tuple with the OrderStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderModifyResponse) GetOrderStatusOk() (*string, bool) {
	if o == nil || IsNil(o.OrderStatus) {
		return nil, false
	}
	return o.OrderStatus, true
}

// HasOrderStatus returns a boolean if a field has been set.
func (o *OrderModifyResponse) HasOrderStatus() bool {
	if o != nil && !IsNil(o.OrderStatus) {
		return true
	}

	return false
}

// SetOrderStatus gets a reference to the given string and assigns it to the OrderStatus field.
func (o *OrderModifyResponse) SetOrderStatus(v string) {
	o.OrderStatus = &v
}

// GetBillToAddressId returns the BillToAddressId field value if set, zero value otherwise.
func (o *OrderModifyResponse) GetBillToAddressId() string {
	if o == nil || IsNil(o.BillToAddressId) {
		var ret string
		return ret
	}
	return *o.BillToAddressId
}

// GetBillToAddressIdOk returns a tuple with the BillToAddressId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderModifyResponse) GetBillToAddressIdOk() (*string, bool) {
	if o == nil || IsNil(o.BillToAddressId) {
		return nil, false
	}
	return o.BillToAddressId, true
}

// HasBillToAddressId returns a boolean if a field has been set.
func (o *OrderModifyResponse) HasBillToAddressId() bool {
	if o != nil && !IsNil(o.BillToAddressId) {
		return true
	}

	return false
}

// SetBillToAddressId gets a reference to the given string and assigns it to the BillToAddressId field.
func (o *OrderModifyResponse) SetBillToAddressId(v string) {
	o.BillToAddressId = &v
}

// GetShipToInfo returns the ShipToInfo field value if set, zero value otherwise.
func (o *OrderModifyResponse) GetShipToInfo() OrderModifyResponseShipToInfo {
	if o == nil || IsNil(o.ShipToInfo) {
		var ret OrderModifyResponseShipToInfo
		return ret
	}
	return *o.ShipToInfo
}

// GetShipToInfoOk returns a tuple with the ShipToInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderModifyResponse) GetShipToInfoOk() (*OrderModifyResponseShipToInfo, bool) {
	if o == nil || IsNil(o.ShipToInfo) {
		return nil, false
	}
	return o.ShipToInfo, true
}

// HasShipToInfo returns a boolean if a field has been set.
func (o *OrderModifyResponse) HasShipToInfo() bool {
	if o != nil && !IsNil(o.ShipToInfo) {
		return true
	}

	return false
}

// SetShipToInfo gets a reference to the given OrderModifyResponseShipToInfo and assigns it to the ShipToInfo field.
func (o *OrderModifyResponse) SetShipToInfo(v OrderModifyResponseShipToInfo) {
	o.ShipToInfo = &v
}

// GetLines returns the Lines field value if set, zero value otherwise.
func (o *OrderModifyResponse) GetLines() []OrderModifyResponseLinesInner {
	if o == nil || IsNil(o.Lines) {
		var ret []OrderModifyResponseLinesInner
		return ret
	}
	return o.Lines
}

// GetLinesOk returns a tuple with the Lines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderModifyResponse) GetLinesOk() ([]OrderModifyResponseLinesInner, bool) {
	if o == nil || IsNil(o.Lines) {
		return nil, false
	}
	return o.Lines, true
}

// HasLines returns a boolean if a field has been set.
func (o *OrderModifyResponse) HasLines() bool {
	if o != nil && !IsNil(o.Lines) {
		return true
	}

	return false
}

// SetLines gets a reference to the given []OrderModifyResponseLinesInner and assigns it to the Lines field.
func (o *OrderModifyResponse) SetLines(v []OrderModifyResponseLinesInner) {
	o.Lines = v
}

// GetRejectedLineItems returns the RejectedLineItems field value if set, zero value otherwise.
func (o *OrderModifyResponse) GetRejectedLineItems() []OrderModifyResponseRejectedLineItemsInner {
	if o == nil || IsNil(o.RejectedLineItems) {
		var ret []OrderModifyResponseRejectedLineItemsInner
		return ret
	}
	return o.RejectedLineItems
}

// GetRejectedLineItemsOk returns a tuple with the RejectedLineItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderModifyResponse) GetRejectedLineItemsOk() ([]OrderModifyResponseRejectedLineItemsInner, bool) {
	if o == nil || IsNil(o.RejectedLineItems) {
		return nil, false
	}
	return o.RejectedLineItems, true
}

// HasRejectedLineItems returns a boolean if a field has been set.
func (o *OrderModifyResponse) HasRejectedLineItems() bool {
	if o != nil && !IsNil(o.RejectedLineItems) {
		return true
	}

	return false
}

// SetRejectedLineItems gets a reference to the given []OrderModifyResponseRejectedLineItemsInner and assigns it to the RejectedLineItems field.
func (o *OrderModifyResponse) SetRejectedLineItems(v []OrderModifyResponseRejectedLineItemsInner) {
	o.RejectedLineItems = v
}

// GetAdditionalAttributes returns the AdditionalAttributes field value if set, zero value otherwise.
func (o *OrderModifyResponse) GetAdditionalAttributes() []OrderModifyResponseLinesInnerAdditionalAttributesInner {
	if o == nil || IsNil(o.AdditionalAttributes) {
		var ret []OrderModifyResponseLinesInnerAdditionalAttributesInner
		return ret
	}
	return o.AdditionalAttributes
}

// GetAdditionalAttributesOk returns a tuple with the AdditionalAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderModifyResponse) GetAdditionalAttributesOk() ([]OrderModifyResponseLinesInnerAdditionalAttributesInner, bool) {
	if o == nil || IsNil(o.AdditionalAttributes) {
		return nil, false
	}
	return o.AdditionalAttributes, true
}

// HasAdditionalAttributes returns a boolean if a field has been set.
func (o *OrderModifyResponse) HasAdditionalAttributes() bool {
	if o != nil && !IsNil(o.AdditionalAttributes) {
		return true
	}

	return false
}

// SetAdditionalAttributes gets a reference to the given []OrderModifyResponseLinesInnerAdditionalAttributesInner and assigns it to the AdditionalAttributes field.
func (o *OrderModifyResponse) SetAdditionalAttributes(v []OrderModifyResponseLinesInnerAdditionalAttributesInner) {
	o.AdditionalAttributes = v
}

func (o OrderModifyResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderModifyResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IngramOrderNumber) {
		toSerialize["ingramOrderNumber"] = o.IngramOrderNumber
	}
	if !IsNil(o.ChangeDescription) {
		toSerialize["changeDescription"] = o.ChangeDescription
	}
	if !IsNil(o.OrderModifiedDate) {
		toSerialize["orderModifiedDate"] = o.OrderModifiedDate
	}
	if !IsNil(o.CustomerOrderNumber) {
		toSerialize["customerOrderNumber"] = o.CustomerOrderNumber
	}
	if !IsNil(o.EndCustomerOrderNumber) {
		toSerialize["endCustomerOrderNumber"] = o.EndCustomerOrderNumber
	}
	if !IsNil(o.OrderTotal) {
		toSerialize["orderTotal"] = o.OrderTotal
	}
	if !IsNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	if !IsNil(o.OrderSubTotal) {
		toSerialize["orderSubTotal"] = o.OrderSubTotal
	}
	if !IsNil(o.FreightCharges) {
		toSerialize["freightCharges"] = o.FreightCharges
	}
	if !IsNil(o.TotalTax) {
		toSerialize["totalTax"] = o.TotalTax
	}
	if !IsNil(o.OrderStatus) {
		toSerialize["orderStatus"] = o.OrderStatus
	}
	if !IsNil(o.BillToAddressId) {
		toSerialize["billToAddressId"] = o.BillToAddressId
	}
	if !IsNil(o.ShipToInfo) {
		toSerialize["shipToInfo"] = o.ShipToInfo
	}
	if !IsNil(o.Lines) {
		toSerialize["lines"] = o.Lines
	}
	if !IsNil(o.RejectedLineItems) {
		toSerialize["rejectedLineItems"] = o.RejectedLineItems
	}
	if !IsNil(o.AdditionalAttributes) {
		toSerialize["additionalAttributes"] = o.AdditionalAttributes
	}
	return toSerialize, nil
}

type NullableOrderModifyResponse struct {
	value *OrderModifyResponse
	isSet bool
}

func (v NullableOrderModifyResponse) Get() *OrderModifyResponse {
	return v.value
}

func (v *NullableOrderModifyResponse) Set(val *OrderModifyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderModifyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderModifyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderModifyResponse(val *OrderModifyResponse) *NullableOrderModifyResponse {
	return &NullableOrderModifyResponse{value: val, isSet: true}
}

func (v NullableOrderModifyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderModifyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


