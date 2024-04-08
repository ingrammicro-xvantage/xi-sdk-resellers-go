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

// checks if the OrderCreateResponseOrdersInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderCreateResponseOrdersInner{}

// OrderCreateResponseOrdersInner struct for OrderCreateResponseOrdersInner
type OrderCreateResponseOrdersInner struct {
	// The number of lines in the order that were successful.
	NumberOfLinesWithSuccess *int32 `json:"numberOfLinesWithSuccess,omitempty"`
	// The number of lines in the order that have errors.
	NumberOfLinesWithError *int32 `json:"numberOfLinesWithError,omitempty"`
	// The number of lines in the order that have a warning.
	NumberOfLinesWithWarning *int32 `json:"numberOfLinesWithWarning,omitempty"`
	// The Ingram Micro order number.
	IngramOrderNumber *string `json:"ingramOrderNumber,omitempty"`
	// The date in UTC format that the order was created in Ingram Micro's system.
	IngramOrderDate *string `json:"ingramOrderDate,omitempty"`
	// Order-level notes.
	Notes *string `json:"notes,omitempty"`
	// The order typer. One of: S=Stocked PO D=Direct Ship PO
	OrderType *string `json:"orderType,omitempty"`
	// The total price for the order.
	OrderTotal *float32 `json:"orderTotal,omitempty"`
	// The total freight charges for the order.
	FreightCharges *float32 `json:"freightCharges,omitempty"`
	// The total tax for the order.
	TotalTax *float32 `json:"totalTax,omitempty"`
	// The country-specific three character ISO 4217 currency code used for the order.
	CurrencyCode *string `json:"currencyCode,omitempty"`
	// The line-level details for the order.
	Lines []OrderCreateResponseOrdersInnerLinesInner `json:"lines,omitempty"`
	MiscellaneousCharges []OrderCreateResponseOrdersInnerMiscellaneousChargesInner `json:"miscellaneousCharges,omitempty"`
	// Link to Order Details for the order(s).
	Links []OrderCreateResponseOrdersInnerLinksInner `json:"links,omitempty"`
	// A list of rejected line items.
	RejectedLineItems []OrderCreateResponseOrdersInnerRejectedLineItemsInner `json:"rejectedLineItems,omitempty"`
	AdditionalAttributes []OrderCreateResponseOrdersInnerAdditionalAttributesInner `json:"additionalAttributes,omitempty"`
}

// NewOrderCreateResponseOrdersInner instantiates a new OrderCreateResponseOrdersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderCreateResponseOrdersInner() *OrderCreateResponseOrdersInner {
	this := OrderCreateResponseOrdersInner{}
	return &this
}

// NewOrderCreateResponseOrdersInnerWithDefaults instantiates a new OrderCreateResponseOrdersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderCreateResponseOrdersInnerWithDefaults() *OrderCreateResponseOrdersInner {
	this := OrderCreateResponseOrdersInner{}
	return &this
}

// GetNumberOfLinesWithSuccess returns the NumberOfLinesWithSuccess field value if set, zero value otherwise.
func (o *OrderCreateResponseOrdersInner) GetNumberOfLinesWithSuccess() int32 {
	if o == nil || IsNil(o.NumberOfLinesWithSuccess) {
		var ret int32
		return ret
	}
	return *o.NumberOfLinesWithSuccess
}

// GetNumberOfLinesWithSuccessOk returns a tuple with the NumberOfLinesWithSuccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateResponseOrdersInner) GetNumberOfLinesWithSuccessOk() (*int32, bool) {
	if o == nil || IsNil(o.NumberOfLinesWithSuccess) {
		return nil, false
	}
	return o.NumberOfLinesWithSuccess, true
}

// HasNumberOfLinesWithSuccess returns a boolean if a field has been set.
func (o *OrderCreateResponseOrdersInner) HasNumberOfLinesWithSuccess() bool {
	if o != nil && !IsNil(o.NumberOfLinesWithSuccess) {
		return true
	}

	return false
}

// SetNumberOfLinesWithSuccess gets a reference to the given int32 and assigns it to the NumberOfLinesWithSuccess field.
func (o *OrderCreateResponseOrdersInner) SetNumberOfLinesWithSuccess(v int32) {
	o.NumberOfLinesWithSuccess = &v
}

// GetNumberOfLinesWithError returns the NumberOfLinesWithError field value if set, zero value otherwise.
func (o *OrderCreateResponseOrdersInner) GetNumberOfLinesWithError() int32 {
	if o == nil || IsNil(o.NumberOfLinesWithError) {
		var ret int32
		return ret
	}
	return *o.NumberOfLinesWithError
}

// GetNumberOfLinesWithErrorOk returns a tuple with the NumberOfLinesWithError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateResponseOrdersInner) GetNumberOfLinesWithErrorOk() (*int32, bool) {
	if o == nil || IsNil(o.NumberOfLinesWithError) {
		return nil, false
	}
	return o.NumberOfLinesWithError, true
}

// HasNumberOfLinesWithError returns a boolean if a field has been set.
func (o *OrderCreateResponseOrdersInner) HasNumberOfLinesWithError() bool {
	if o != nil && !IsNil(o.NumberOfLinesWithError) {
		return true
	}

	return false
}

// SetNumberOfLinesWithError gets a reference to the given int32 and assigns it to the NumberOfLinesWithError field.
func (o *OrderCreateResponseOrdersInner) SetNumberOfLinesWithError(v int32) {
	o.NumberOfLinesWithError = &v
}

// GetNumberOfLinesWithWarning returns the NumberOfLinesWithWarning field value if set, zero value otherwise.
func (o *OrderCreateResponseOrdersInner) GetNumberOfLinesWithWarning() int32 {
	if o == nil || IsNil(o.NumberOfLinesWithWarning) {
		var ret int32
		return ret
	}
	return *o.NumberOfLinesWithWarning
}

// GetNumberOfLinesWithWarningOk returns a tuple with the NumberOfLinesWithWarning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateResponseOrdersInner) GetNumberOfLinesWithWarningOk() (*int32, bool) {
	if o == nil || IsNil(o.NumberOfLinesWithWarning) {
		return nil, false
	}
	return o.NumberOfLinesWithWarning, true
}

// HasNumberOfLinesWithWarning returns a boolean if a field has been set.
func (o *OrderCreateResponseOrdersInner) HasNumberOfLinesWithWarning() bool {
	if o != nil && !IsNil(o.NumberOfLinesWithWarning) {
		return true
	}

	return false
}

// SetNumberOfLinesWithWarning gets a reference to the given int32 and assigns it to the NumberOfLinesWithWarning field.
func (o *OrderCreateResponseOrdersInner) SetNumberOfLinesWithWarning(v int32) {
	o.NumberOfLinesWithWarning = &v
}

// GetIngramOrderNumber returns the IngramOrderNumber field value if set, zero value otherwise.
func (o *OrderCreateResponseOrdersInner) GetIngramOrderNumber() string {
	if o == nil || IsNil(o.IngramOrderNumber) {
		var ret string
		return ret
	}
	return *o.IngramOrderNumber
}

// GetIngramOrderNumberOk returns a tuple with the IngramOrderNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateResponseOrdersInner) GetIngramOrderNumberOk() (*string, bool) {
	if o == nil || IsNil(o.IngramOrderNumber) {
		return nil, false
	}
	return o.IngramOrderNumber, true
}

// HasIngramOrderNumber returns a boolean if a field has been set.
func (o *OrderCreateResponseOrdersInner) HasIngramOrderNumber() bool {
	if o != nil && !IsNil(o.IngramOrderNumber) {
		return true
	}

	return false
}

// SetIngramOrderNumber gets a reference to the given string and assigns it to the IngramOrderNumber field.
func (o *OrderCreateResponseOrdersInner) SetIngramOrderNumber(v string) {
	o.IngramOrderNumber = &v
}

// GetIngramOrderDate returns the IngramOrderDate field value if set, zero value otherwise.
func (o *OrderCreateResponseOrdersInner) GetIngramOrderDate() string {
	if o == nil || IsNil(o.IngramOrderDate) {
		var ret string
		return ret
	}
	return *o.IngramOrderDate
}

// GetIngramOrderDateOk returns a tuple with the IngramOrderDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateResponseOrdersInner) GetIngramOrderDateOk() (*string, bool) {
	if o == nil || IsNil(o.IngramOrderDate) {
		return nil, false
	}
	return o.IngramOrderDate, true
}

// HasIngramOrderDate returns a boolean if a field has been set.
func (o *OrderCreateResponseOrdersInner) HasIngramOrderDate() bool {
	if o != nil && !IsNil(o.IngramOrderDate) {
		return true
	}

	return false
}

// SetIngramOrderDate gets a reference to the given string and assigns it to the IngramOrderDate field.
func (o *OrderCreateResponseOrdersInner) SetIngramOrderDate(v string) {
	o.IngramOrderDate = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *OrderCreateResponseOrdersInner) GetNotes() string {
	if o == nil || IsNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateResponseOrdersInner) GetNotesOk() (*string, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *OrderCreateResponseOrdersInner) HasNotes() bool {
	if o != nil && !IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *OrderCreateResponseOrdersInner) SetNotes(v string) {
	o.Notes = &v
}

// GetOrderType returns the OrderType field value if set, zero value otherwise.
func (o *OrderCreateResponseOrdersInner) GetOrderType() string {
	if o == nil || IsNil(o.OrderType) {
		var ret string
		return ret
	}
	return *o.OrderType
}

// GetOrderTypeOk returns a tuple with the OrderType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateResponseOrdersInner) GetOrderTypeOk() (*string, bool) {
	if o == nil || IsNil(o.OrderType) {
		return nil, false
	}
	return o.OrderType, true
}

// HasOrderType returns a boolean if a field has been set.
func (o *OrderCreateResponseOrdersInner) HasOrderType() bool {
	if o != nil && !IsNil(o.OrderType) {
		return true
	}

	return false
}

// SetOrderType gets a reference to the given string and assigns it to the OrderType field.
func (o *OrderCreateResponseOrdersInner) SetOrderType(v string) {
	o.OrderType = &v
}

// GetOrderTotal returns the OrderTotal field value if set, zero value otherwise.
func (o *OrderCreateResponseOrdersInner) GetOrderTotal() float32 {
	if o == nil || IsNil(o.OrderTotal) {
		var ret float32
		return ret
	}
	return *o.OrderTotal
}

// GetOrderTotalOk returns a tuple with the OrderTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateResponseOrdersInner) GetOrderTotalOk() (*float32, bool) {
	if o == nil || IsNil(o.OrderTotal) {
		return nil, false
	}
	return o.OrderTotal, true
}

// HasOrderTotal returns a boolean if a field has been set.
func (o *OrderCreateResponseOrdersInner) HasOrderTotal() bool {
	if o != nil && !IsNil(o.OrderTotal) {
		return true
	}

	return false
}

// SetOrderTotal gets a reference to the given float32 and assigns it to the OrderTotal field.
func (o *OrderCreateResponseOrdersInner) SetOrderTotal(v float32) {
	o.OrderTotal = &v
}

// GetFreightCharges returns the FreightCharges field value if set, zero value otherwise.
func (o *OrderCreateResponseOrdersInner) GetFreightCharges() float32 {
	if o == nil || IsNil(o.FreightCharges) {
		var ret float32
		return ret
	}
	return *o.FreightCharges
}

// GetFreightChargesOk returns a tuple with the FreightCharges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateResponseOrdersInner) GetFreightChargesOk() (*float32, bool) {
	if o == nil || IsNil(o.FreightCharges) {
		return nil, false
	}
	return o.FreightCharges, true
}

// HasFreightCharges returns a boolean if a field has been set.
func (o *OrderCreateResponseOrdersInner) HasFreightCharges() bool {
	if o != nil && !IsNil(o.FreightCharges) {
		return true
	}

	return false
}

// SetFreightCharges gets a reference to the given float32 and assigns it to the FreightCharges field.
func (o *OrderCreateResponseOrdersInner) SetFreightCharges(v float32) {
	o.FreightCharges = &v
}

// GetTotalTax returns the TotalTax field value if set, zero value otherwise.
func (o *OrderCreateResponseOrdersInner) GetTotalTax() float32 {
	if o == nil || IsNil(o.TotalTax) {
		var ret float32
		return ret
	}
	return *o.TotalTax
}

// GetTotalTaxOk returns a tuple with the TotalTax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateResponseOrdersInner) GetTotalTaxOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalTax) {
		return nil, false
	}
	return o.TotalTax, true
}

// HasTotalTax returns a boolean if a field has been set.
func (o *OrderCreateResponseOrdersInner) HasTotalTax() bool {
	if o != nil && !IsNil(o.TotalTax) {
		return true
	}

	return false
}

// SetTotalTax gets a reference to the given float32 and assigns it to the TotalTax field.
func (o *OrderCreateResponseOrdersInner) SetTotalTax(v float32) {
	o.TotalTax = &v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *OrderCreateResponseOrdersInner) GetCurrencyCode() string {
	if o == nil || IsNil(o.CurrencyCode) {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateResponseOrdersInner) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CurrencyCode) {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *OrderCreateResponseOrdersInner) HasCurrencyCode() bool {
	if o != nil && !IsNil(o.CurrencyCode) {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *OrderCreateResponseOrdersInner) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

// GetLines returns the Lines field value if set, zero value otherwise.
func (o *OrderCreateResponseOrdersInner) GetLines() []OrderCreateResponseOrdersInnerLinesInner {
	if o == nil || IsNil(o.Lines) {
		var ret []OrderCreateResponseOrdersInnerLinesInner
		return ret
	}
	return o.Lines
}

// GetLinesOk returns a tuple with the Lines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateResponseOrdersInner) GetLinesOk() ([]OrderCreateResponseOrdersInnerLinesInner, bool) {
	if o == nil || IsNil(o.Lines) {
		return nil, false
	}
	return o.Lines, true
}

// HasLines returns a boolean if a field has been set.
func (o *OrderCreateResponseOrdersInner) HasLines() bool {
	if o != nil && !IsNil(o.Lines) {
		return true
	}

	return false
}

// SetLines gets a reference to the given []OrderCreateResponseOrdersInnerLinesInner and assigns it to the Lines field.
func (o *OrderCreateResponseOrdersInner) SetLines(v []OrderCreateResponseOrdersInnerLinesInner) {
	o.Lines = v
}

// GetMiscellaneousCharges returns the MiscellaneousCharges field value if set, zero value otherwise.
func (o *OrderCreateResponseOrdersInner) GetMiscellaneousCharges() []OrderCreateResponseOrdersInnerMiscellaneousChargesInner {
	if o == nil || IsNil(o.MiscellaneousCharges) {
		var ret []OrderCreateResponseOrdersInnerMiscellaneousChargesInner
		return ret
	}
	return o.MiscellaneousCharges
}

// GetMiscellaneousChargesOk returns a tuple with the MiscellaneousCharges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateResponseOrdersInner) GetMiscellaneousChargesOk() ([]OrderCreateResponseOrdersInnerMiscellaneousChargesInner, bool) {
	if o == nil || IsNil(o.MiscellaneousCharges) {
		return nil, false
	}
	return o.MiscellaneousCharges, true
}

// HasMiscellaneousCharges returns a boolean if a field has been set.
func (o *OrderCreateResponseOrdersInner) HasMiscellaneousCharges() bool {
	if o != nil && !IsNil(o.MiscellaneousCharges) {
		return true
	}

	return false
}

// SetMiscellaneousCharges gets a reference to the given []OrderCreateResponseOrdersInnerMiscellaneousChargesInner and assigns it to the MiscellaneousCharges field.
func (o *OrderCreateResponseOrdersInner) SetMiscellaneousCharges(v []OrderCreateResponseOrdersInnerMiscellaneousChargesInner) {
	o.MiscellaneousCharges = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *OrderCreateResponseOrdersInner) GetLinks() []OrderCreateResponseOrdersInnerLinksInner {
	if o == nil || IsNil(o.Links) {
		var ret []OrderCreateResponseOrdersInnerLinksInner
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateResponseOrdersInner) GetLinksOk() ([]OrderCreateResponseOrdersInnerLinksInner, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *OrderCreateResponseOrdersInner) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []OrderCreateResponseOrdersInnerLinksInner and assigns it to the Links field.
func (o *OrderCreateResponseOrdersInner) SetLinks(v []OrderCreateResponseOrdersInnerLinksInner) {
	o.Links = v
}

// GetRejectedLineItems returns the RejectedLineItems field value if set, zero value otherwise.
func (o *OrderCreateResponseOrdersInner) GetRejectedLineItems() []OrderCreateResponseOrdersInnerRejectedLineItemsInner {
	if o == nil || IsNil(o.RejectedLineItems) {
		var ret []OrderCreateResponseOrdersInnerRejectedLineItemsInner
		return ret
	}
	return o.RejectedLineItems
}

// GetRejectedLineItemsOk returns a tuple with the RejectedLineItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateResponseOrdersInner) GetRejectedLineItemsOk() ([]OrderCreateResponseOrdersInnerRejectedLineItemsInner, bool) {
	if o == nil || IsNil(o.RejectedLineItems) {
		return nil, false
	}
	return o.RejectedLineItems, true
}

// HasRejectedLineItems returns a boolean if a field has been set.
func (o *OrderCreateResponseOrdersInner) HasRejectedLineItems() bool {
	if o != nil && !IsNil(o.RejectedLineItems) {
		return true
	}

	return false
}

// SetRejectedLineItems gets a reference to the given []OrderCreateResponseOrdersInnerRejectedLineItemsInner and assigns it to the RejectedLineItems field.
func (o *OrderCreateResponseOrdersInner) SetRejectedLineItems(v []OrderCreateResponseOrdersInnerRejectedLineItemsInner) {
	o.RejectedLineItems = v
}

// GetAdditionalAttributes returns the AdditionalAttributes field value if set, zero value otherwise.
func (o *OrderCreateResponseOrdersInner) GetAdditionalAttributes() []OrderCreateResponseOrdersInnerAdditionalAttributesInner {
	if o == nil || IsNil(o.AdditionalAttributes) {
		var ret []OrderCreateResponseOrdersInnerAdditionalAttributesInner
		return ret
	}
	return o.AdditionalAttributes
}

// GetAdditionalAttributesOk returns a tuple with the AdditionalAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateResponseOrdersInner) GetAdditionalAttributesOk() ([]OrderCreateResponseOrdersInnerAdditionalAttributesInner, bool) {
	if o == nil || IsNil(o.AdditionalAttributes) {
		return nil, false
	}
	return o.AdditionalAttributes, true
}

// HasAdditionalAttributes returns a boolean if a field has been set.
func (o *OrderCreateResponseOrdersInner) HasAdditionalAttributes() bool {
	if o != nil && !IsNil(o.AdditionalAttributes) {
		return true
	}

	return false
}

// SetAdditionalAttributes gets a reference to the given []OrderCreateResponseOrdersInnerAdditionalAttributesInner and assigns it to the AdditionalAttributes field.
func (o *OrderCreateResponseOrdersInner) SetAdditionalAttributes(v []OrderCreateResponseOrdersInnerAdditionalAttributesInner) {
	o.AdditionalAttributes = v
}

func (o OrderCreateResponseOrdersInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderCreateResponseOrdersInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NumberOfLinesWithSuccess) {
		toSerialize["numberOfLinesWithSuccess"] = o.NumberOfLinesWithSuccess
	}
	if !IsNil(o.NumberOfLinesWithError) {
		toSerialize["numberOfLinesWithError"] = o.NumberOfLinesWithError
	}
	if !IsNil(o.NumberOfLinesWithWarning) {
		toSerialize["numberOfLinesWithWarning"] = o.NumberOfLinesWithWarning
	}
	if !IsNil(o.IngramOrderNumber) {
		toSerialize["ingramOrderNumber"] = o.IngramOrderNumber
	}
	if !IsNil(o.IngramOrderDate) {
		toSerialize["ingramOrderDate"] = o.IngramOrderDate
	}
	if !IsNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	if !IsNil(o.OrderType) {
		toSerialize["orderType"] = o.OrderType
	}
	if !IsNil(o.OrderTotal) {
		toSerialize["orderTotal"] = o.OrderTotal
	}
	if !IsNil(o.FreightCharges) {
		toSerialize["freightCharges"] = o.FreightCharges
	}
	if !IsNil(o.TotalTax) {
		toSerialize["totalTax"] = o.TotalTax
	}
	if !IsNil(o.CurrencyCode) {
		toSerialize["currencyCode"] = o.CurrencyCode
	}
	if !IsNil(o.Lines) {
		toSerialize["lines"] = o.Lines
	}
	if !IsNil(o.MiscellaneousCharges) {
		toSerialize["miscellaneousCharges"] = o.MiscellaneousCharges
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.RejectedLineItems) {
		toSerialize["rejectedLineItems"] = o.RejectedLineItems
	}
	if !IsNil(o.AdditionalAttributes) {
		toSerialize["additionalAttributes"] = o.AdditionalAttributes
	}
	return toSerialize, nil
}

type NullableOrderCreateResponseOrdersInner struct {
	value *OrderCreateResponseOrdersInner
	isSet bool
}

func (v NullableOrderCreateResponseOrdersInner) Get() *OrderCreateResponseOrdersInner {
	return v.value
}

func (v *NullableOrderCreateResponseOrdersInner) Set(val *OrderCreateResponseOrdersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderCreateResponseOrdersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderCreateResponseOrdersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderCreateResponseOrdersInner(val *OrderCreateResponseOrdersInner) *NullableOrderCreateResponseOrdersInner {
	return &NullableOrderCreateResponseOrdersInner{value: val, isSet: true}
}

func (v NullableOrderCreateResponseOrdersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderCreateResponseOrdersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


