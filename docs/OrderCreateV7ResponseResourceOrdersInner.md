# OrderCreateV7ResponseResourceOrdersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumberOfLinesWithSuccess** | Pointer to **int32** | The number of lines in the order that were successful. | [optional] 
**NumberOfLinesWithError** | Pointer to **int32** | The number of lines in the order that have errors. | [optional] 
**NumberOfLinesWithWarning** | Pointer to **int32** | The number of lines in the order that have a warning. | [optional] 
**IngramOrderNumber** | Pointer to **string** | The Ingram Micro order number. | [optional] 
**IngramOrderDate** | Pointer to **string** | The date in UTC format that the order was created in Ingram Micro&#39;s system. | [optional] 
**Notes** | Pointer to **string** | Order-level notes. | [optional] 
**OrderType** | Pointer to **string** | The order typer. One of: S&#x3D;Stocked PO D&#x3D;Direct Ship PO | [optional] 
**OrderTotal** | Pointer to **float32** | The total price for the order. | [optional] 
**FreightCharges** | Pointer to **float32** | The total freight charges for the order. | [optional] 
**TotalTax** | Pointer to **float32** | The total tax for the order. | [optional] 
**CurrencyCode** | Pointer to **string** | The country-specific three character ISO 4217 currency code used for the order. | [optional] 
**Lines** | Pointer to [**[]OrderCreateV7ResponseResourceOrdersInnerLinesInner**](OrderCreateV7ResponseResourceOrdersInnerLinesInner.md) | The line-level details for the order. | [optional] 
**Links** | Pointer to [**[]OrderCreateResponseOrdersInnerLinksInner**](OrderCreateResponseOrdersInnerLinksInner.md) | Link to Order Details for the order(s). | [optional] 

## Methods

### NewOrderCreateV7ResponseResourceOrdersInner

`func NewOrderCreateV7ResponseResourceOrdersInner() *OrderCreateV7ResponseResourceOrdersInner`

NewOrderCreateV7ResponseResourceOrdersInner instantiates a new OrderCreateV7ResponseResourceOrdersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderCreateV7ResponseResourceOrdersInnerWithDefaults

`func NewOrderCreateV7ResponseResourceOrdersInnerWithDefaults() *OrderCreateV7ResponseResourceOrdersInner`

NewOrderCreateV7ResponseResourceOrdersInnerWithDefaults instantiates a new OrderCreateV7ResponseResourceOrdersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumberOfLinesWithSuccess

`func (o *OrderCreateV7ResponseResourceOrdersInner) GetNumberOfLinesWithSuccess() int32`

GetNumberOfLinesWithSuccess returns the NumberOfLinesWithSuccess field if non-nil, zero value otherwise.

### GetNumberOfLinesWithSuccessOk

`func (o *OrderCreateV7ResponseResourceOrdersInner) GetNumberOfLinesWithSuccessOk() (*int32, bool)`

GetNumberOfLinesWithSuccessOk returns a tuple with the NumberOfLinesWithSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfLinesWithSuccess

`func (o *OrderCreateV7ResponseResourceOrdersInner) SetNumberOfLinesWithSuccess(v int32)`

SetNumberOfLinesWithSuccess sets NumberOfLinesWithSuccess field to given value.

### HasNumberOfLinesWithSuccess

`func (o *OrderCreateV7ResponseResourceOrdersInner) HasNumberOfLinesWithSuccess() bool`

HasNumberOfLinesWithSuccess returns a boolean if a field has been set.

### GetNumberOfLinesWithError

`func (o *OrderCreateV7ResponseResourceOrdersInner) GetNumberOfLinesWithError() int32`

GetNumberOfLinesWithError returns the NumberOfLinesWithError field if non-nil, zero value otherwise.

### GetNumberOfLinesWithErrorOk

`func (o *OrderCreateV7ResponseResourceOrdersInner) GetNumberOfLinesWithErrorOk() (*int32, bool)`

GetNumberOfLinesWithErrorOk returns a tuple with the NumberOfLinesWithError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfLinesWithError

`func (o *OrderCreateV7ResponseResourceOrdersInner) SetNumberOfLinesWithError(v int32)`

SetNumberOfLinesWithError sets NumberOfLinesWithError field to given value.

### HasNumberOfLinesWithError

`func (o *OrderCreateV7ResponseResourceOrdersInner) HasNumberOfLinesWithError() bool`

HasNumberOfLinesWithError returns a boolean if a field has been set.

### GetNumberOfLinesWithWarning

`func (o *OrderCreateV7ResponseResourceOrdersInner) GetNumberOfLinesWithWarning() int32`

GetNumberOfLinesWithWarning returns the NumberOfLinesWithWarning field if non-nil, zero value otherwise.

### GetNumberOfLinesWithWarningOk

`func (o *OrderCreateV7ResponseResourceOrdersInner) GetNumberOfLinesWithWarningOk() (*int32, bool)`

GetNumberOfLinesWithWarningOk returns a tuple with the NumberOfLinesWithWarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfLinesWithWarning

`func (o *OrderCreateV7ResponseResourceOrdersInner) SetNumberOfLinesWithWarning(v int32)`

SetNumberOfLinesWithWarning sets NumberOfLinesWithWarning field to given value.

### HasNumberOfLinesWithWarning

`func (o *OrderCreateV7ResponseResourceOrdersInner) HasNumberOfLinesWithWarning() bool`

HasNumberOfLinesWithWarning returns a boolean if a field has been set.

### GetIngramOrderNumber

`func (o *OrderCreateV7ResponseResourceOrdersInner) GetIngramOrderNumber() string`

GetIngramOrderNumber returns the IngramOrderNumber field if non-nil, zero value otherwise.

### GetIngramOrderNumberOk

`func (o *OrderCreateV7ResponseResourceOrdersInner) GetIngramOrderNumberOk() (*string, bool)`

GetIngramOrderNumberOk returns a tuple with the IngramOrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngramOrderNumber

`func (o *OrderCreateV7ResponseResourceOrdersInner) SetIngramOrderNumber(v string)`

SetIngramOrderNumber sets IngramOrderNumber field to given value.

### HasIngramOrderNumber

`func (o *OrderCreateV7ResponseResourceOrdersInner) HasIngramOrderNumber() bool`

HasIngramOrderNumber returns a boolean if a field has been set.

### GetIngramOrderDate

`func (o *OrderCreateV7ResponseResourceOrdersInner) GetIngramOrderDate() string`

GetIngramOrderDate returns the IngramOrderDate field if non-nil, zero value otherwise.

### GetIngramOrderDateOk

`func (o *OrderCreateV7ResponseResourceOrdersInner) GetIngramOrderDateOk() (*string, bool)`

GetIngramOrderDateOk returns a tuple with the IngramOrderDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngramOrderDate

`func (o *OrderCreateV7ResponseResourceOrdersInner) SetIngramOrderDate(v string)`

SetIngramOrderDate sets IngramOrderDate field to given value.

### HasIngramOrderDate

`func (o *OrderCreateV7ResponseResourceOrdersInner) HasIngramOrderDate() bool`

HasIngramOrderDate returns a boolean if a field has been set.

### GetNotes

`func (o *OrderCreateV7ResponseResourceOrdersInner) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *OrderCreateV7ResponseResourceOrdersInner) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *OrderCreateV7ResponseResourceOrdersInner) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *OrderCreateV7ResponseResourceOrdersInner) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetOrderType

`func (o *OrderCreateV7ResponseResourceOrdersInner) GetOrderType() string`

GetOrderType returns the OrderType field if non-nil, zero value otherwise.

### GetOrderTypeOk

`func (o *OrderCreateV7ResponseResourceOrdersInner) GetOrderTypeOk() (*string, bool)`

GetOrderTypeOk returns a tuple with the OrderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderType

`func (o *OrderCreateV7ResponseResourceOrdersInner) SetOrderType(v string)`

SetOrderType sets OrderType field to given value.

### HasOrderType

`func (o *OrderCreateV7ResponseResourceOrdersInner) HasOrderType() bool`

HasOrderType returns a boolean if a field has been set.

### GetOrderTotal

`func (o *OrderCreateV7ResponseResourceOrdersInner) GetOrderTotal() float32`

GetOrderTotal returns the OrderTotal field if non-nil, zero value otherwise.

### GetOrderTotalOk

`func (o *OrderCreateV7ResponseResourceOrdersInner) GetOrderTotalOk() (*float32, bool)`

GetOrderTotalOk returns a tuple with the OrderTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderTotal

`func (o *OrderCreateV7ResponseResourceOrdersInner) SetOrderTotal(v float32)`

SetOrderTotal sets OrderTotal field to given value.

### HasOrderTotal

`func (o *OrderCreateV7ResponseResourceOrdersInner) HasOrderTotal() bool`

HasOrderTotal returns a boolean if a field has been set.

### GetFreightCharges

`func (o *OrderCreateV7ResponseResourceOrdersInner) GetFreightCharges() float32`

GetFreightCharges returns the FreightCharges field if non-nil, zero value otherwise.

### GetFreightChargesOk

`func (o *OrderCreateV7ResponseResourceOrdersInner) GetFreightChargesOk() (*float32, bool)`

GetFreightChargesOk returns a tuple with the FreightCharges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreightCharges

`func (o *OrderCreateV7ResponseResourceOrdersInner) SetFreightCharges(v float32)`

SetFreightCharges sets FreightCharges field to given value.

### HasFreightCharges

`func (o *OrderCreateV7ResponseResourceOrdersInner) HasFreightCharges() bool`

HasFreightCharges returns a boolean if a field has been set.

### GetTotalTax

`func (o *OrderCreateV7ResponseResourceOrdersInner) GetTotalTax() float32`

GetTotalTax returns the TotalTax field if non-nil, zero value otherwise.

### GetTotalTaxOk

`func (o *OrderCreateV7ResponseResourceOrdersInner) GetTotalTaxOk() (*float32, bool)`

GetTotalTaxOk returns a tuple with the TotalTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTax

`func (o *OrderCreateV7ResponseResourceOrdersInner) SetTotalTax(v float32)`

SetTotalTax sets TotalTax field to given value.

### HasTotalTax

`func (o *OrderCreateV7ResponseResourceOrdersInner) HasTotalTax() bool`

HasTotalTax returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *OrderCreateV7ResponseResourceOrdersInner) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *OrderCreateV7ResponseResourceOrdersInner) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *OrderCreateV7ResponseResourceOrdersInner) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *OrderCreateV7ResponseResourceOrdersInner) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetLines

`func (o *OrderCreateV7ResponseResourceOrdersInner) GetLines() []OrderCreateV7ResponseResourceOrdersInnerLinesInner`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *OrderCreateV7ResponseResourceOrdersInner) GetLinesOk() (*[]OrderCreateV7ResponseResourceOrdersInnerLinesInner, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *OrderCreateV7ResponseResourceOrdersInner) SetLines(v []OrderCreateV7ResponseResourceOrdersInnerLinesInner)`

SetLines sets Lines field to given value.

### HasLines

`func (o *OrderCreateV7ResponseResourceOrdersInner) HasLines() bool`

HasLines returns a boolean if a field has been set.

### GetLinks

`func (o *OrderCreateV7ResponseResourceOrdersInner) GetLinks() []OrderCreateResponseOrdersInnerLinksInner`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *OrderCreateV7ResponseResourceOrdersInner) GetLinksOk() (*[]OrderCreateResponseOrdersInnerLinksInner, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *OrderCreateV7ResponseResourceOrdersInner) SetLinks(v []OrderCreateResponseOrdersInnerLinksInner)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *OrderCreateV7ResponseResourceOrdersInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


