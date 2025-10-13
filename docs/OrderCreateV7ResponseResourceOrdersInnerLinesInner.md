# OrderCreateV7ResponseResourceOrdersInnerLinesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubOrderNumber** | Pointer to **string** | The sub order number. The two-digit prefix is the warehouse code of the warehouse nearest the reseller. The middle number is the order number. The two-digit suffix is the sub order number. | [optional] 
**IngramLineNumber** | Pointer to **string** | The Ingram Micro line number for the product. | [optional] 
**CustomerLineNumber** | Pointer to **string** | The reseller&#39;s line number for reference in their system. | [optional] 
**LineStatus** | Pointer to **string** | The status for the line item in the order. One of: Backordered, Open | [optional] 
**IngramPartNumber** | Pointer to **string** | The Ingram Micro part number for the line item. | [optional] 
**UnitPrice** | Pointer to **float32** | The unit price for the line item. | [optional] 
**ExtendedUnitPrice** | Pointer to **float32** | The extended list price (unit price X quantity) for the line item. | [optional] 
**QuantityOrdered** | Pointer to **int32** | The quantity of the line item ordered. | [optional] 
**QuantityConfirmed** | Pointer to **int32** | The quantity of the line item that has been confirmed. | [optional] 
**QuantityBackOrdered** | Pointer to **int32** | The quantity of the line item that is backordered. | [optional] 
**Notes** | Pointer to **string** | Line-level notes. | [optional] 
**ShipmentDetails** | Pointer to [**[]OrderCreateV7ResponseResourceOrdersInnerLinesInnerShipmentDetailsInner**](OrderCreateV7ResponseResourceOrdersInnerLinesInnerShipmentDetailsInner.md) | The shipment details for the line item. | [optional] 

## Methods

### NewOrderCreateV7ResponseResourceOrdersInnerLinesInner

`func NewOrderCreateV7ResponseResourceOrdersInnerLinesInner() *OrderCreateV7ResponseResourceOrdersInnerLinesInner`

NewOrderCreateV7ResponseResourceOrdersInnerLinesInner instantiates a new OrderCreateV7ResponseResourceOrdersInnerLinesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderCreateV7ResponseResourceOrdersInnerLinesInnerWithDefaults

`func NewOrderCreateV7ResponseResourceOrdersInnerLinesInnerWithDefaults() *OrderCreateV7ResponseResourceOrdersInnerLinesInner`

NewOrderCreateV7ResponseResourceOrdersInnerLinesInnerWithDefaults instantiates a new OrderCreateV7ResponseResourceOrdersInnerLinesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubOrderNumber

`func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) GetSubOrderNumber() string`

GetSubOrderNumber returns the SubOrderNumber field if non-nil, zero value otherwise.

### GetSubOrderNumberOk

`func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) GetSubOrderNumberOk() (*string, bool)`

GetSubOrderNumberOk returns a tuple with the SubOrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubOrderNumber

`func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) SetSubOrderNumber(v string)`

SetSubOrderNumber sets SubOrderNumber field to given value.

### HasSubOrderNumber

`func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) HasSubOrderNumber() bool`

HasSubOrderNumber returns a boolean if a field has been set.

### GetIngramLineNumber

`func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) GetIngramLineNumber() string`

GetIngramLineNumber returns the IngramLineNumber field if non-nil, zero value otherwise.

### GetIngramLineNumberOk

`func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) GetIngramLineNumberOk() (*string, bool)`

GetIngramLineNumberOk returns a tuple with the IngramLineNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngramLineNumber

`func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) SetIngramLineNumber(v string)`

SetIngramLineNumber sets IngramLineNumber field to given value.

### HasIngramLineNumber

`func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) HasIngramLineNumber() bool`

HasIngramLineNumber returns a boolean if a field has been set.

### GetCustomerLineNumber

`func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) GetCustomerLineNumber() string`

GetCustomerLineNumber returns the CustomerLineNumber field if non-nil, zero value otherwise.

### GetCustomerLineNumberOk

`func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) GetCustomerLineNumberOk() (*string, bool)`

GetCustomerLineNumberOk returns a tuple with the CustomerLineNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerLineNumber

`func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) SetCustomerLineNumber(v string)`

SetCustomerLineNumber sets CustomerLineNumber field to given value.

### HasCustomerLineNumber

`func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) HasCustomerLineNumber() bool`

HasCustomerLineNumber returns a boolean if a field has been set.

### GetLineStatus

`func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) GetLineStatus() string`

GetLineStatus returns the LineStatus field if non-nil, zero value otherwise.

### GetLineStatusOk

`func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) GetLineStatusOk() (*string, bool)`

GetLineStatusOk returns a tuple with the LineStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineStatus

`func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) SetLineStatus(v string)`

SetLineStatus sets LineStatus field to given value.

### HasLineStatus

`func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) HasLineStatus() bool`

HasLineStatus returns a boolean if a field has been set.

### GetIngramPartNumber

`func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) GetIngramPartNumber() string`

GetIngramPartNumber returns the IngramPartNumber field if non-nil, zero value otherwise.

### GetIngramPartNumberOk

`func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) GetIngramPartNumberOk() (*string, bool)`

GetIngramPartNumberOk returns a tuple with the IngramPartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngramPartNumber

`func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) SetIngramPartNumber(v string)`

SetIngramPartNumber sets IngramPartNumber field to given value.

### HasIngramPartNumber

`func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) HasIngramPartNumber() bool`

HasIngramPartNumber returns a boolean if a field has been set.

### GetUnitPrice

`func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) GetUnitPrice() float32`

GetUnitPrice returns the UnitPrice field if non-nil, zero value otherwise.

### GetUnitPriceOk

`func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) GetUnitPriceOk() (*float32, bool)`

GetUnitPriceOk returns a tuple with the UnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPrice

`func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) SetUnitPrice(v float32)`

SetUnitPrice sets UnitPrice field to given value.

### HasUnitPrice

`func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) HasUnitPrice() bool`

HasUnitPrice returns a boolean if a field has been set.

### GetExtendedUnitPrice

`func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) GetExtendedUnitPrice() float32`

GetExtendedUnitPrice returns the ExtendedUnitPrice field if non-nil, zero value otherwise.

### GetExtendedUnitPriceOk

`func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) GetExtendedUnitPriceOk() (*float32, bool)`

GetExtendedUnitPriceOk returns a tuple with the ExtendedUnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedUnitPrice

`func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) SetExtendedUnitPrice(v float32)`

SetExtendedUnitPrice sets ExtendedUnitPrice field to given value.

### HasExtendedUnitPrice

`func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) HasExtendedUnitPrice() bool`

HasExtendedUnitPrice returns a boolean if a field has been set.

### GetQuantityOrdered

`func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) GetQuantityOrdered() int32`

GetQuantityOrdered returns the QuantityOrdered field if non-nil, zero value otherwise.

### GetQuantityOrderedOk

`func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) GetQuantityOrderedOk() (*int32, bool)`

GetQuantityOrderedOk returns a tuple with the QuantityOrdered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityOrdered

`func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) SetQuantityOrdered(v int32)`

SetQuantityOrdered sets QuantityOrdered field to given value.

### HasQuantityOrdered

`func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) HasQuantityOrdered() bool`

HasQuantityOrdered returns a boolean if a field has been set.

### GetQuantityConfirmed

`func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) GetQuantityConfirmed() int32`

GetQuantityConfirmed returns the QuantityConfirmed field if non-nil, zero value otherwise.

### GetQuantityConfirmedOk

`func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) GetQuantityConfirmedOk() (*int32, bool)`

GetQuantityConfirmedOk returns a tuple with the QuantityConfirmed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityConfirmed

`func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) SetQuantityConfirmed(v int32)`

SetQuantityConfirmed sets QuantityConfirmed field to given value.

### HasQuantityConfirmed

`func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) HasQuantityConfirmed() bool`

HasQuantityConfirmed returns a boolean if a field has been set.

### GetQuantityBackOrdered

`func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) GetQuantityBackOrdered() int32`

GetQuantityBackOrdered returns the QuantityBackOrdered field if non-nil, zero value otherwise.

### GetQuantityBackOrderedOk

`func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) GetQuantityBackOrderedOk() (*int32, bool)`

GetQuantityBackOrderedOk returns a tuple with the QuantityBackOrdered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityBackOrdered

`func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) SetQuantityBackOrdered(v int32)`

SetQuantityBackOrdered sets QuantityBackOrdered field to given value.

### HasQuantityBackOrdered

`func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) HasQuantityBackOrdered() bool`

HasQuantityBackOrdered returns a boolean if a field has been set.

### GetNotes

`func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetShipmentDetails

`func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) GetShipmentDetails() []OrderCreateV7ResponseResourceOrdersInnerLinesInnerShipmentDetailsInner`

GetShipmentDetails returns the ShipmentDetails field if non-nil, zero value otherwise.

### GetShipmentDetailsOk

`func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) GetShipmentDetailsOk() (*[]OrderCreateV7ResponseResourceOrdersInnerLinesInnerShipmentDetailsInner, bool)`

GetShipmentDetailsOk returns a tuple with the ShipmentDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentDetails

`func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) SetShipmentDetails(v []OrderCreateV7ResponseResourceOrdersInnerLinesInnerShipmentDetailsInner)`

SetShipmentDetails sets ShipmentDetails field to given value.

### HasShipmentDetails

`func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) HasShipmentDetails() bool`

HasShipmentDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


