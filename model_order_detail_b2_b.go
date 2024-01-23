/*
Reseller API

For Resellers. <br> Who are looking to Innovate with Ingram Micro's API SolutionsAutomate your eCommerce with our offering of APIs and Webhooks to create a seamless experience for your customers.

API version: 6.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package xi_sdk_resellers

import (
	"encoding/json"
	"time"
)

// checks if the OrderDetailB2B type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderDetailB2B{}

// OrderDetailB2B struct for OrderDetailB2B
type OrderDetailB2B struct {
	// The IngramMicro sales order number.
	IngramOrderNumber *string `json:"ingramOrderNumber,omitempty"`
	// The IngramMicro sales order date.
	IngramOrderDate *time.Time `json:"ingramOrderDate,omitempty"`
	// The IngramMicro sales order type.
	OrderType *string `json:"orderType,omitempty"`
	// The reseller's order number for reference in their system.
	CustomerOrderNumber *string `json:"customerOrderNumber,omitempty"`
	// The end customer's order number for reference in their system.
	EndCustomerOrderNumber *string `json:"endCustomerOrderNumber,omitempty"`
	// The web order id of the order.
	WebOrderId *string `json:"webOrderId,omitempty"`
	// The vendor's order number for reference in their system
	VendorSalesOrderNumber *string `json:"vendorSalesOrderNumber,omitempty"`
	// Ingram purchase order number.
	IngramPurchaseOrderNumber *string `json:"ingramPurchaseOrderNumber,omitempty"`
	// The header-level status of the order. One of- Shipped, Canceled, Backordered, Processing, On Hold, Delivered.
	OrderStatus *string `json:"orderStatus,omitempty"`
	// The total cost for the order, includes subtotal, freight charges, and tax.
	OrderTotal *float64 `json:"orderTotal,omitempty"`
	// The sub total cost for the order, not including tax and freight.
	OrderSubTotal *float64 `json:"orderSubTotal,omitempty"`
	// The freight charges for the order.
	FreightCharges *float64 `json:"freightCharges,omitempty"`
	// The country-specific three digit ISO 4217 currency code for the order.
	CurrencyCode *string `json:"currencyCode,omitempty"`
	// Total order weight. unit -- North america - Pounds , other countries will be KG.
	TotalWeight *float64 `json:"totalWeight,omitempty"`
	// Total tax on the orders placed.
	TotalTax *float64 `json:"totalTax,omitempty"`
	// The payment terms of the order. (Ex- Net 30 days).
	PaymentTerms *string `json:"paymentTerms,omitempty"`
	// The header-level notes for the order.
	Notes *string `json:"notes,omitempty"`
	BillToInfo *OrderDetailB2BBillToInfo `json:"billToInfo,omitempty"`
	ShipToInfo *OrderDetailB2BShipToInfo `json:"shipToInfo,omitempty"`
	EndUserInfo *OrderDetailB2BEndUserInfo `json:"endUserInfo,omitempty"`
	Lines []OrderDetailB2BLinesInner `json:"lines,omitempty"`
	MiscellaneousCharges []OrderDetailB2BMiscellaneousChargesInner `json:"miscellaneousCharges,omitempty"`
	AdditionalAttributes []OrderDetailB2BAdditionalAttributesInner `json:"additionalAttributes,omitempty"`
}

// NewOrderDetailB2B instantiates a new OrderDetailB2B object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderDetailB2B() *OrderDetailB2B {
	this := OrderDetailB2B{}
	return &this
}

// NewOrderDetailB2BWithDefaults instantiates a new OrderDetailB2B object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderDetailB2BWithDefaults() *OrderDetailB2B {
	this := OrderDetailB2B{}
	return &this
}

// GetIngramOrderNumber returns the IngramOrderNumber field value if set, zero value otherwise.
func (o *OrderDetailB2B) GetIngramOrderNumber() string {
	if o == nil || IsNil(o.IngramOrderNumber) {
		var ret string
		return ret
	}
	return *o.IngramOrderNumber
}

// GetIngramOrderNumberOk returns a tuple with the IngramOrderNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailB2B) GetIngramOrderNumberOk() (*string, bool) {
	if o == nil || IsNil(o.IngramOrderNumber) {
		return nil, false
	}
	return o.IngramOrderNumber, true
}

// HasIngramOrderNumber returns a boolean if a field has been set.
func (o *OrderDetailB2B) HasIngramOrderNumber() bool {
	if o != nil && !IsNil(o.IngramOrderNumber) {
		return true
	}

	return false
}

// SetIngramOrderNumber gets a reference to the given string and assigns it to the IngramOrderNumber field.
func (o *OrderDetailB2B) SetIngramOrderNumber(v string) {
	o.IngramOrderNumber = &v
}

// GetIngramOrderDate returns the IngramOrderDate field value if set, zero value otherwise.
func (o *OrderDetailB2B) GetIngramOrderDate() time.Time {
	if o == nil || IsNil(o.IngramOrderDate) {
		var ret time.Time
		return ret
	}
	return *o.IngramOrderDate
}

// GetIngramOrderDateOk returns a tuple with the IngramOrderDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailB2B) GetIngramOrderDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.IngramOrderDate) {
		return nil, false
	}
	return o.IngramOrderDate, true
}

// HasIngramOrderDate returns a boolean if a field has been set.
func (o *OrderDetailB2B) HasIngramOrderDate() bool {
	if o != nil && !IsNil(o.IngramOrderDate) {
		return true
	}

	return false
}

// SetIngramOrderDate gets a reference to the given time.Time and assigns it to the IngramOrderDate field.
func (o *OrderDetailB2B) SetIngramOrderDate(v time.Time) {
	o.IngramOrderDate = &v
}

// GetOrderType returns the OrderType field value if set, zero value otherwise.
func (o *OrderDetailB2B) GetOrderType() string {
	if o == nil || IsNil(o.OrderType) {
		var ret string
		return ret
	}
	return *o.OrderType
}

// GetOrderTypeOk returns a tuple with the OrderType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailB2B) GetOrderTypeOk() (*string, bool) {
	if o == nil || IsNil(o.OrderType) {
		return nil, false
	}
	return o.OrderType, true
}

// HasOrderType returns a boolean if a field has been set.
func (o *OrderDetailB2B) HasOrderType() bool {
	if o != nil && !IsNil(o.OrderType) {
		return true
	}

	return false
}

// SetOrderType gets a reference to the given string and assigns it to the OrderType field.
func (o *OrderDetailB2B) SetOrderType(v string) {
	o.OrderType = &v
}

// GetCustomerOrderNumber returns the CustomerOrderNumber field value if set, zero value otherwise.
func (o *OrderDetailB2B) GetCustomerOrderNumber() string {
	if o == nil || IsNil(o.CustomerOrderNumber) {
		var ret string
		return ret
	}
	return *o.CustomerOrderNumber
}

// GetCustomerOrderNumberOk returns a tuple with the CustomerOrderNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailB2B) GetCustomerOrderNumberOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerOrderNumber) {
		return nil, false
	}
	return o.CustomerOrderNumber, true
}

// HasCustomerOrderNumber returns a boolean if a field has been set.
func (o *OrderDetailB2B) HasCustomerOrderNumber() bool {
	if o != nil && !IsNil(o.CustomerOrderNumber) {
		return true
	}

	return false
}

// SetCustomerOrderNumber gets a reference to the given string and assigns it to the CustomerOrderNumber field.
func (o *OrderDetailB2B) SetCustomerOrderNumber(v string) {
	o.CustomerOrderNumber = &v
}

// GetEndCustomerOrderNumber returns the EndCustomerOrderNumber field value if set, zero value otherwise.
func (o *OrderDetailB2B) GetEndCustomerOrderNumber() string {
	if o == nil || IsNil(o.EndCustomerOrderNumber) {
		var ret string
		return ret
	}
	return *o.EndCustomerOrderNumber
}

// GetEndCustomerOrderNumberOk returns a tuple with the EndCustomerOrderNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailB2B) GetEndCustomerOrderNumberOk() (*string, bool) {
	if o == nil || IsNil(o.EndCustomerOrderNumber) {
		return nil, false
	}
	return o.EndCustomerOrderNumber, true
}

// HasEndCustomerOrderNumber returns a boolean if a field has been set.
func (o *OrderDetailB2B) HasEndCustomerOrderNumber() bool {
	if o != nil && !IsNil(o.EndCustomerOrderNumber) {
		return true
	}

	return false
}

// SetEndCustomerOrderNumber gets a reference to the given string and assigns it to the EndCustomerOrderNumber field.
func (o *OrderDetailB2B) SetEndCustomerOrderNumber(v string) {
	o.EndCustomerOrderNumber = &v
}

// GetWebOrderId returns the WebOrderId field value if set, zero value otherwise.
func (o *OrderDetailB2B) GetWebOrderId() string {
	if o == nil || IsNil(o.WebOrderId) {
		var ret string
		return ret
	}
	return *o.WebOrderId
}

// GetWebOrderIdOk returns a tuple with the WebOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailB2B) GetWebOrderIdOk() (*string, bool) {
	if o == nil || IsNil(o.WebOrderId) {
		return nil, false
	}
	return o.WebOrderId, true
}

// HasWebOrderId returns a boolean if a field has been set.
func (o *OrderDetailB2B) HasWebOrderId() bool {
	if o != nil && !IsNil(o.WebOrderId) {
		return true
	}

	return false
}

// SetWebOrderId gets a reference to the given string and assigns it to the WebOrderId field.
func (o *OrderDetailB2B) SetWebOrderId(v string) {
	o.WebOrderId = &v
}

// GetVendorSalesOrderNumber returns the VendorSalesOrderNumber field value if set, zero value otherwise.
func (o *OrderDetailB2B) GetVendorSalesOrderNumber() string {
	if o == nil || IsNil(o.VendorSalesOrderNumber) {
		var ret string
		return ret
	}
	return *o.VendorSalesOrderNumber
}

// GetVendorSalesOrderNumberOk returns a tuple with the VendorSalesOrderNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailB2B) GetVendorSalesOrderNumberOk() (*string, bool) {
	if o == nil || IsNil(o.VendorSalesOrderNumber) {
		return nil, false
	}
	return o.VendorSalesOrderNumber, true
}

// HasVendorSalesOrderNumber returns a boolean if a field has been set.
func (o *OrderDetailB2B) HasVendorSalesOrderNumber() bool {
	if o != nil && !IsNil(o.VendorSalesOrderNumber) {
		return true
	}

	return false
}

// SetVendorSalesOrderNumber gets a reference to the given string and assigns it to the VendorSalesOrderNumber field.
func (o *OrderDetailB2B) SetVendorSalesOrderNumber(v string) {
	o.VendorSalesOrderNumber = &v
}

// GetIngramPurchaseOrderNumber returns the IngramPurchaseOrderNumber field value if set, zero value otherwise.
func (o *OrderDetailB2B) GetIngramPurchaseOrderNumber() string {
	if o == nil || IsNil(o.IngramPurchaseOrderNumber) {
		var ret string
		return ret
	}
	return *o.IngramPurchaseOrderNumber
}

// GetIngramPurchaseOrderNumberOk returns a tuple with the IngramPurchaseOrderNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailB2B) GetIngramPurchaseOrderNumberOk() (*string, bool) {
	if o == nil || IsNil(o.IngramPurchaseOrderNumber) {
		return nil, false
	}
	return o.IngramPurchaseOrderNumber, true
}

// HasIngramPurchaseOrderNumber returns a boolean if a field has been set.
func (o *OrderDetailB2B) HasIngramPurchaseOrderNumber() bool {
	if o != nil && !IsNil(o.IngramPurchaseOrderNumber) {
		return true
	}

	return false
}

// SetIngramPurchaseOrderNumber gets a reference to the given string and assigns it to the IngramPurchaseOrderNumber field.
func (o *OrderDetailB2B) SetIngramPurchaseOrderNumber(v string) {
	o.IngramPurchaseOrderNumber = &v
}

// GetOrderStatus returns the OrderStatus field value if set, zero value otherwise.
func (o *OrderDetailB2B) GetOrderStatus() string {
	if o == nil || IsNil(o.OrderStatus) {
		var ret string
		return ret
	}
	return *o.OrderStatus
}

// GetOrderStatusOk returns a tuple with the OrderStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailB2B) GetOrderStatusOk() (*string, bool) {
	if o == nil || IsNil(o.OrderStatus) {
		return nil, false
	}
	return o.OrderStatus, true
}

// HasOrderStatus returns a boolean if a field has been set.
func (o *OrderDetailB2B) HasOrderStatus() bool {
	if o != nil && !IsNil(o.OrderStatus) {
		return true
	}

	return false
}

// SetOrderStatus gets a reference to the given string and assigns it to the OrderStatus field.
func (o *OrderDetailB2B) SetOrderStatus(v string) {
	o.OrderStatus = &v
}

// GetOrderTotal returns the OrderTotal field value if set, zero value otherwise.
func (o *OrderDetailB2B) GetOrderTotal() float64 {
	if o == nil || IsNil(o.OrderTotal) {
		var ret float64
		return ret
	}
	return *o.OrderTotal
}

// GetOrderTotalOk returns a tuple with the OrderTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailB2B) GetOrderTotalOk() (*float64, bool) {
	if o == nil || IsNil(o.OrderTotal) {
		return nil, false
	}
	return o.OrderTotal, true
}

// HasOrderTotal returns a boolean if a field has been set.
func (o *OrderDetailB2B) HasOrderTotal() bool {
	if o != nil && !IsNil(o.OrderTotal) {
		return true
	}

	return false
}

// SetOrderTotal gets a reference to the given float64 and assigns it to the OrderTotal field.
func (o *OrderDetailB2B) SetOrderTotal(v float64) {
	o.OrderTotal = &v
}

// GetOrderSubTotal returns the OrderSubTotal field value if set, zero value otherwise.
func (o *OrderDetailB2B) GetOrderSubTotal() float64 {
	if o == nil || IsNil(o.OrderSubTotal) {
		var ret float64
		return ret
	}
	return *o.OrderSubTotal
}

// GetOrderSubTotalOk returns a tuple with the OrderSubTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailB2B) GetOrderSubTotalOk() (*float64, bool) {
	if o == nil || IsNil(o.OrderSubTotal) {
		return nil, false
	}
	return o.OrderSubTotal, true
}

// HasOrderSubTotal returns a boolean if a field has been set.
func (o *OrderDetailB2B) HasOrderSubTotal() bool {
	if o != nil && !IsNil(o.OrderSubTotal) {
		return true
	}

	return false
}

// SetOrderSubTotal gets a reference to the given float64 and assigns it to the OrderSubTotal field.
func (o *OrderDetailB2B) SetOrderSubTotal(v float64) {
	o.OrderSubTotal = &v
}

// GetFreightCharges returns the FreightCharges field value if set, zero value otherwise.
func (o *OrderDetailB2B) GetFreightCharges() float64 {
	if o == nil || IsNil(o.FreightCharges) {
		var ret float64
		return ret
	}
	return *o.FreightCharges
}

// GetFreightChargesOk returns a tuple with the FreightCharges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailB2B) GetFreightChargesOk() (*float64, bool) {
	if o == nil || IsNil(o.FreightCharges) {
		return nil, false
	}
	return o.FreightCharges, true
}

// HasFreightCharges returns a boolean if a field has been set.
func (o *OrderDetailB2B) HasFreightCharges() bool {
	if o != nil && !IsNil(o.FreightCharges) {
		return true
	}

	return false
}

// SetFreightCharges gets a reference to the given float64 and assigns it to the FreightCharges field.
func (o *OrderDetailB2B) SetFreightCharges(v float64) {
	o.FreightCharges = &v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *OrderDetailB2B) GetCurrencyCode() string {
	if o == nil || IsNil(o.CurrencyCode) {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailB2B) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CurrencyCode) {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *OrderDetailB2B) HasCurrencyCode() bool {
	if o != nil && !IsNil(o.CurrencyCode) {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *OrderDetailB2B) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

// GetTotalWeight returns the TotalWeight field value if set, zero value otherwise.
func (o *OrderDetailB2B) GetTotalWeight() float64 {
	if o == nil || IsNil(o.TotalWeight) {
		var ret float64
		return ret
	}
	return *o.TotalWeight
}

// GetTotalWeightOk returns a tuple with the TotalWeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailB2B) GetTotalWeightOk() (*float64, bool) {
	if o == nil || IsNil(o.TotalWeight) {
		return nil, false
	}
	return o.TotalWeight, true
}

// HasTotalWeight returns a boolean if a field has been set.
func (o *OrderDetailB2B) HasTotalWeight() bool {
	if o != nil && !IsNil(o.TotalWeight) {
		return true
	}

	return false
}

// SetTotalWeight gets a reference to the given float64 and assigns it to the TotalWeight field.
func (o *OrderDetailB2B) SetTotalWeight(v float64) {
	o.TotalWeight = &v
}

// GetTotalTax returns the TotalTax field value if set, zero value otherwise.
func (o *OrderDetailB2B) GetTotalTax() float64 {
	if o == nil || IsNil(o.TotalTax) {
		var ret float64
		return ret
	}
	return *o.TotalTax
}

// GetTotalTaxOk returns a tuple with the TotalTax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailB2B) GetTotalTaxOk() (*float64, bool) {
	if o == nil || IsNil(o.TotalTax) {
		return nil, false
	}
	return o.TotalTax, true
}

// HasTotalTax returns a boolean if a field has been set.
func (o *OrderDetailB2B) HasTotalTax() bool {
	if o != nil && !IsNil(o.TotalTax) {
		return true
	}

	return false
}

// SetTotalTax gets a reference to the given float64 and assigns it to the TotalTax field.
func (o *OrderDetailB2B) SetTotalTax(v float64) {
	o.TotalTax = &v
}

// GetPaymentTerms returns the PaymentTerms field value if set, zero value otherwise.
func (o *OrderDetailB2B) GetPaymentTerms() string {
	if o == nil || IsNil(o.PaymentTerms) {
		var ret string
		return ret
	}
	return *o.PaymentTerms
}

// GetPaymentTermsOk returns a tuple with the PaymentTerms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailB2B) GetPaymentTermsOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentTerms) {
		return nil, false
	}
	return o.PaymentTerms, true
}

// HasPaymentTerms returns a boolean if a field has been set.
func (o *OrderDetailB2B) HasPaymentTerms() bool {
	if o != nil && !IsNil(o.PaymentTerms) {
		return true
	}

	return false
}

// SetPaymentTerms gets a reference to the given string and assigns it to the PaymentTerms field.
func (o *OrderDetailB2B) SetPaymentTerms(v string) {
	o.PaymentTerms = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *OrderDetailB2B) GetNotes() string {
	if o == nil || IsNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailB2B) GetNotesOk() (*string, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *OrderDetailB2B) HasNotes() bool {
	if o != nil && !IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *OrderDetailB2B) SetNotes(v string) {
	o.Notes = &v
}

// GetBillToInfo returns the BillToInfo field value if set, zero value otherwise.
func (o *OrderDetailB2B) GetBillToInfo() OrderDetailB2BBillToInfo {
	if o == nil || IsNil(o.BillToInfo) {
		var ret OrderDetailB2BBillToInfo
		return ret
	}
	return *o.BillToInfo
}

// GetBillToInfoOk returns a tuple with the BillToInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailB2B) GetBillToInfoOk() (*OrderDetailB2BBillToInfo, bool) {
	if o == nil || IsNil(o.BillToInfo) {
		return nil, false
	}
	return o.BillToInfo, true
}

// HasBillToInfo returns a boolean if a field has been set.
func (o *OrderDetailB2B) HasBillToInfo() bool {
	if o != nil && !IsNil(o.BillToInfo) {
		return true
	}

	return false
}

// SetBillToInfo gets a reference to the given OrderDetailB2BBillToInfo and assigns it to the BillToInfo field.
func (o *OrderDetailB2B) SetBillToInfo(v OrderDetailB2BBillToInfo) {
	o.BillToInfo = &v
}

// GetShipToInfo returns the ShipToInfo field value if set, zero value otherwise.
func (o *OrderDetailB2B) GetShipToInfo() OrderDetailB2BShipToInfo {
	if o == nil || IsNil(o.ShipToInfo) {
		var ret OrderDetailB2BShipToInfo
		return ret
	}
	return *o.ShipToInfo
}

// GetShipToInfoOk returns a tuple with the ShipToInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailB2B) GetShipToInfoOk() (*OrderDetailB2BShipToInfo, bool) {
	if o == nil || IsNil(o.ShipToInfo) {
		return nil, false
	}
	return o.ShipToInfo, true
}

// HasShipToInfo returns a boolean if a field has been set.
func (o *OrderDetailB2B) HasShipToInfo() bool {
	if o != nil && !IsNil(o.ShipToInfo) {
		return true
	}

	return false
}

// SetShipToInfo gets a reference to the given OrderDetailB2BShipToInfo and assigns it to the ShipToInfo field.
func (o *OrderDetailB2B) SetShipToInfo(v OrderDetailB2BShipToInfo) {
	o.ShipToInfo = &v
}

// GetEndUserInfo returns the EndUserInfo field value if set, zero value otherwise.
func (o *OrderDetailB2B) GetEndUserInfo() OrderDetailB2BEndUserInfo {
	if o == nil || IsNil(o.EndUserInfo) {
		var ret OrderDetailB2BEndUserInfo
		return ret
	}
	return *o.EndUserInfo
}

// GetEndUserInfoOk returns a tuple with the EndUserInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailB2B) GetEndUserInfoOk() (*OrderDetailB2BEndUserInfo, bool) {
	if o == nil || IsNil(o.EndUserInfo) {
		return nil, false
	}
	return o.EndUserInfo, true
}

// HasEndUserInfo returns a boolean if a field has been set.
func (o *OrderDetailB2B) HasEndUserInfo() bool {
	if o != nil && !IsNil(o.EndUserInfo) {
		return true
	}

	return false
}

// SetEndUserInfo gets a reference to the given OrderDetailB2BEndUserInfo and assigns it to the EndUserInfo field.
func (o *OrderDetailB2B) SetEndUserInfo(v OrderDetailB2BEndUserInfo) {
	o.EndUserInfo = &v
}

// GetLines returns the Lines field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrderDetailB2B) GetLines() []OrderDetailB2BLinesInner {
	if o == nil {
		var ret []OrderDetailB2BLinesInner
		return ret
	}
	return o.Lines
}

// GetLinesOk returns a tuple with the Lines field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrderDetailB2B) GetLinesOk() ([]OrderDetailB2BLinesInner, bool) {
	if o == nil || IsNil(o.Lines) {
		return nil, false
	}
	return o.Lines, true
}

// HasLines returns a boolean if a field has been set.
func (o *OrderDetailB2B) HasLines() bool {
	if o != nil && IsNil(o.Lines) {
		return true
	}

	return false
}

// SetLines gets a reference to the given []OrderDetailB2BLinesInner and assigns it to the Lines field.
func (o *OrderDetailB2B) SetLines(v []OrderDetailB2BLinesInner) {
	o.Lines = v
}

// GetMiscellaneousCharges returns the MiscellaneousCharges field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrderDetailB2B) GetMiscellaneousCharges() []OrderDetailB2BMiscellaneousChargesInner {
	if o == nil {
		var ret []OrderDetailB2BMiscellaneousChargesInner
		return ret
	}
	return o.MiscellaneousCharges
}

// GetMiscellaneousChargesOk returns a tuple with the MiscellaneousCharges field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrderDetailB2B) GetMiscellaneousChargesOk() ([]OrderDetailB2BMiscellaneousChargesInner, bool) {
	if o == nil || IsNil(o.MiscellaneousCharges) {
		return nil, false
	}
	return o.MiscellaneousCharges, true
}

// HasMiscellaneousCharges returns a boolean if a field has been set.
func (o *OrderDetailB2B) HasMiscellaneousCharges() bool {
	if o != nil && IsNil(o.MiscellaneousCharges) {
		return true
	}

	return false
}

// SetMiscellaneousCharges gets a reference to the given []OrderDetailB2BMiscellaneousChargesInner and assigns it to the MiscellaneousCharges field.
func (o *OrderDetailB2B) SetMiscellaneousCharges(v []OrderDetailB2BMiscellaneousChargesInner) {
	o.MiscellaneousCharges = v
}

// GetAdditionalAttributes returns the AdditionalAttributes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrderDetailB2B) GetAdditionalAttributes() []OrderDetailB2BAdditionalAttributesInner {
	if o == nil {
		var ret []OrderDetailB2BAdditionalAttributesInner
		return ret
	}
	return o.AdditionalAttributes
}

// GetAdditionalAttributesOk returns a tuple with the AdditionalAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrderDetailB2B) GetAdditionalAttributesOk() ([]OrderDetailB2BAdditionalAttributesInner, bool) {
	if o == nil || IsNil(o.AdditionalAttributes) {
		return nil, false
	}
	return o.AdditionalAttributes, true
}

// HasAdditionalAttributes returns a boolean if a field has been set.
func (o *OrderDetailB2B) HasAdditionalAttributes() bool {
	if o != nil && IsNil(o.AdditionalAttributes) {
		return true
	}

	return false
}

// SetAdditionalAttributes gets a reference to the given []OrderDetailB2BAdditionalAttributesInner and assigns it to the AdditionalAttributes field.
func (o *OrderDetailB2B) SetAdditionalAttributes(v []OrderDetailB2BAdditionalAttributesInner) {
	o.AdditionalAttributes = v
}

func (o OrderDetailB2B) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderDetailB2B) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IngramOrderNumber) {
		toSerialize["ingramOrderNumber"] = o.IngramOrderNumber
	}
	if !IsNil(o.IngramOrderDate) {
		toSerialize["ingramOrderDate"] = o.IngramOrderDate
	}
	if !IsNil(o.OrderType) {
		toSerialize["orderType"] = o.OrderType
	}
	if !IsNil(o.CustomerOrderNumber) {
		toSerialize["customerOrderNumber"] = o.CustomerOrderNumber
	}
	if !IsNil(o.EndCustomerOrderNumber) {
		toSerialize["endCustomerOrderNumber"] = o.EndCustomerOrderNumber
	}
	if !IsNil(o.WebOrderId) {
		toSerialize["webOrderId"] = o.WebOrderId
	}
	if !IsNil(o.VendorSalesOrderNumber) {
		toSerialize["vendorSalesOrderNumber"] = o.VendorSalesOrderNumber
	}
	if !IsNil(o.IngramPurchaseOrderNumber) {
		toSerialize["ingramPurchaseOrderNumber"] = o.IngramPurchaseOrderNumber
	}
	if !IsNil(o.OrderStatus) {
		toSerialize["orderStatus"] = o.OrderStatus
	}
	if !IsNil(o.OrderTotal) {
		toSerialize["orderTotal"] = o.OrderTotal
	}
	if !IsNil(o.OrderSubTotal) {
		toSerialize["orderSubTotal"] = o.OrderSubTotal
	}
	if !IsNil(o.FreightCharges) {
		toSerialize["freightCharges"] = o.FreightCharges
	}
	if !IsNil(o.CurrencyCode) {
		toSerialize["currencyCode"] = o.CurrencyCode
	}
	if !IsNil(o.TotalWeight) {
		toSerialize["totalWeight"] = o.TotalWeight
	}
	if !IsNil(o.TotalTax) {
		toSerialize["totalTax"] = o.TotalTax
	}
	if !IsNil(o.PaymentTerms) {
		toSerialize["paymentTerms"] = o.PaymentTerms
	}
	if !IsNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	if !IsNil(o.BillToInfo) {
		toSerialize["billToInfo"] = o.BillToInfo
	}
	if !IsNil(o.ShipToInfo) {
		toSerialize["shipToInfo"] = o.ShipToInfo
	}
	if !IsNil(o.EndUserInfo) {
		toSerialize["endUserInfo"] = o.EndUserInfo
	}
	if o.Lines != nil {
		toSerialize["lines"] = o.Lines
	}
	if o.MiscellaneousCharges != nil {
		toSerialize["miscellaneousCharges"] = o.MiscellaneousCharges
	}
	if o.AdditionalAttributes != nil {
		toSerialize["additionalAttributes"] = o.AdditionalAttributes
	}
	return toSerialize, nil
}

type NullableOrderDetailB2B struct {
	value *OrderDetailB2B
	isSet bool
}

func (v NullableOrderDetailB2B) Get() *OrderDetailB2B {
	return v.value
}

func (v *NullableOrderDetailB2B) Set(val *OrderDetailB2B) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderDetailB2B) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderDetailB2B) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderDetailB2B(val *OrderDetailB2B) *NullableOrderDetailB2B {
	return &NullableOrderDetailB2B{value: val, isSet: true}
}

func (v NullableOrderDetailB2B) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderDetailB2B) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


