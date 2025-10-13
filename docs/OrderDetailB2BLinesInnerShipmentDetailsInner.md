# OrderDetailB2BLinesInnerShipmentDetailsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Quantity** | Pointer to **NullableInt32** | The quantity shipped of the line item. | [optional] 
**DeliveryNumber** | Pointer to **string** | The actual date of delivery of the line item. | [optional] 
**EstimatedShipDate** | Pointer to **string** | The date the line item is expected to be shipped. | [optional] 
**ShippedDate** | Pointer to **string** |  | [optional] 
**EstimatedDeliveryDate** | Pointer to **string** |  | [optional] 
**ShipFromWarehouseId** | Pointer to **string** | The ID of the warehouse the product will ship from. | [optional] 
**ShipFromLocation** | Pointer to **string** | The city and state the line item ships from. | [optional] 
**InvoiceNumber** | Pointer to **string** | The Ingram Micro invoice number for the line item. | [optional] 
**InvoiceDate** | Pointer to **string** | The date the IngramMicro invoice was created for the line item. | [optional] 
**CarrierDetails** | Pointer to [**[]OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner**](OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner.md) | The shipment carrier details for the line item. | [optional] 

## Methods

### NewOrderDetailB2BLinesInnerShipmentDetailsInner

`func NewOrderDetailB2BLinesInnerShipmentDetailsInner() *OrderDetailB2BLinesInnerShipmentDetailsInner`

NewOrderDetailB2BLinesInnerShipmentDetailsInner instantiates a new OrderDetailB2BLinesInnerShipmentDetailsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderDetailB2BLinesInnerShipmentDetailsInnerWithDefaults

`func NewOrderDetailB2BLinesInnerShipmentDetailsInnerWithDefaults() *OrderDetailB2BLinesInnerShipmentDetailsInner`

NewOrderDetailB2BLinesInnerShipmentDetailsInnerWithDefaults instantiates a new OrderDetailB2BLinesInnerShipmentDetailsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuantity

`func (o *OrderDetailB2BLinesInnerShipmentDetailsInner) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *OrderDetailB2BLinesInnerShipmentDetailsInner) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *OrderDetailB2BLinesInnerShipmentDetailsInner) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *OrderDetailB2BLinesInnerShipmentDetailsInner) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### SetQuantityNil

`func (o *OrderDetailB2BLinesInnerShipmentDetailsInner) SetQuantityNil(b bool)`

 SetQuantityNil sets the value for Quantity to be an explicit nil

### UnsetQuantity
`func (o *OrderDetailB2BLinesInnerShipmentDetailsInner) UnsetQuantity()`

UnsetQuantity ensures that no value is present for Quantity, not even an explicit nil
### GetDeliveryNumber

`func (o *OrderDetailB2BLinesInnerShipmentDetailsInner) GetDeliveryNumber() string`

GetDeliveryNumber returns the DeliveryNumber field if non-nil, zero value otherwise.

### GetDeliveryNumberOk

`func (o *OrderDetailB2BLinesInnerShipmentDetailsInner) GetDeliveryNumberOk() (*string, bool)`

GetDeliveryNumberOk returns a tuple with the DeliveryNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryNumber

`func (o *OrderDetailB2BLinesInnerShipmentDetailsInner) SetDeliveryNumber(v string)`

SetDeliveryNumber sets DeliveryNumber field to given value.

### HasDeliveryNumber

`func (o *OrderDetailB2BLinesInnerShipmentDetailsInner) HasDeliveryNumber() bool`

HasDeliveryNumber returns a boolean if a field has been set.

### GetEstimatedShipDate

`func (o *OrderDetailB2BLinesInnerShipmentDetailsInner) GetEstimatedShipDate() string`

GetEstimatedShipDate returns the EstimatedShipDate field if non-nil, zero value otherwise.

### GetEstimatedShipDateOk

`func (o *OrderDetailB2BLinesInnerShipmentDetailsInner) GetEstimatedShipDateOk() (*string, bool)`

GetEstimatedShipDateOk returns a tuple with the EstimatedShipDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedShipDate

`func (o *OrderDetailB2BLinesInnerShipmentDetailsInner) SetEstimatedShipDate(v string)`

SetEstimatedShipDate sets EstimatedShipDate field to given value.

### HasEstimatedShipDate

`func (o *OrderDetailB2BLinesInnerShipmentDetailsInner) HasEstimatedShipDate() bool`

HasEstimatedShipDate returns a boolean if a field has been set.

### GetShippedDate

`func (o *OrderDetailB2BLinesInnerShipmentDetailsInner) GetShippedDate() string`

GetShippedDate returns the ShippedDate field if non-nil, zero value otherwise.

### GetShippedDateOk

`func (o *OrderDetailB2BLinesInnerShipmentDetailsInner) GetShippedDateOk() (*string, bool)`

GetShippedDateOk returns a tuple with the ShippedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippedDate

`func (o *OrderDetailB2BLinesInnerShipmentDetailsInner) SetShippedDate(v string)`

SetShippedDate sets ShippedDate field to given value.

### HasShippedDate

`func (o *OrderDetailB2BLinesInnerShipmentDetailsInner) HasShippedDate() bool`

HasShippedDate returns a boolean if a field has been set.

### GetEstimatedDeliveryDate

`func (o *OrderDetailB2BLinesInnerShipmentDetailsInner) GetEstimatedDeliveryDate() string`

GetEstimatedDeliveryDate returns the EstimatedDeliveryDate field if non-nil, zero value otherwise.

### GetEstimatedDeliveryDateOk

`func (o *OrderDetailB2BLinesInnerShipmentDetailsInner) GetEstimatedDeliveryDateOk() (*string, bool)`

GetEstimatedDeliveryDateOk returns a tuple with the EstimatedDeliveryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedDeliveryDate

`func (o *OrderDetailB2BLinesInnerShipmentDetailsInner) SetEstimatedDeliveryDate(v string)`

SetEstimatedDeliveryDate sets EstimatedDeliveryDate field to given value.

### HasEstimatedDeliveryDate

`func (o *OrderDetailB2BLinesInnerShipmentDetailsInner) HasEstimatedDeliveryDate() bool`

HasEstimatedDeliveryDate returns a boolean if a field has been set.

### GetShipFromWarehouseId

`func (o *OrderDetailB2BLinesInnerShipmentDetailsInner) GetShipFromWarehouseId() string`

GetShipFromWarehouseId returns the ShipFromWarehouseId field if non-nil, zero value otherwise.

### GetShipFromWarehouseIdOk

`func (o *OrderDetailB2BLinesInnerShipmentDetailsInner) GetShipFromWarehouseIdOk() (*string, bool)`

GetShipFromWarehouseIdOk returns a tuple with the ShipFromWarehouseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipFromWarehouseId

`func (o *OrderDetailB2BLinesInnerShipmentDetailsInner) SetShipFromWarehouseId(v string)`

SetShipFromWarehouseId sets ShipFromWarehouseId field to given value.

### HasShipFromWarehouseId

`func (o *OrderDetailB2BLinesInnerShipmentDetailsInner) HasShipFromWarehouseId() bool`

HasShipFromWarehouseId returns a boolean if a field has been set.

### GetShipFromLocation

`func (o *OrderDetailB2BLinesInnerShipmentDetailsInner) GetShipFromLocation() string`

GetShipFromLocation returns the ShipFromLocation field if non-nil, zero value otherwise.

### GetShipFromLocationOk

`func (o *OrderDetailB2BLinesInnerShipmentDetailsInner) GetShipFromLocationOk() (*string, bool)`

GetShipFromLocationOk returns a tuple with the ShipFromLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipFromLocation

`func (o *OrderDetailB2BLinesInnerShipmentDetailsInner) SetShipFromLocation(v string)`

SetShipFromLocation sets ShipFromLocation field to given value.

### HasShipFromLocation

`func (o *OrderDetailB2BLinesInnerShipmentDetailsInner) HasShipFromLocation() bool`

HasShipFromLocation returns a boolean if a field has been set.

### GetInvoiceNumber

`func (o *OrderDetailB2BLinesInnerShipmentDetailsInner) GetInvoiceNumber() string`

GetInvoiceNumber returns the InvoiceNumber field if non-nil, zero value otherwise.

### GetInvoiceNumberOk

`func (o *OrderDetailB2BLinesInnerShipmentDetailsInner) GetInvoiceNumberOk() (*string, bool)`

GetInvoiceNumberOk returns a tuple with the InvoiceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceNumber

`func (o *OrderDetailB2BLinesInnerShipmentDetailsInner) SetInvoiceNumber(v string)`

SetInvoiceNumber sets InvoiceNumber field to given value.

### HasInvoiceNumber

`func (o *OrderDetailB2BLinesInnerShipmentDetailsInner) HasInvoiceNumber() bool`

HasInvoiceNumber returns a boolean if a field has been set.

### GetInvoiceDate

`func (o *OrderDetailB2BLinesInnerShipmentDetailsInner) GetInvoiceDate() string`

GetInvoiceDate returns the InvoiceDate field if non-nil, zero value otherwise.

### GetInvoiceDateOk

`func (o *OrderDetailB2BLinesInnerShipmentDetailsInner) GetInvoiceDateOk() (*string, bool)`

GetInvoiceDateOk returns a tuple with the InvoiceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceDate

`func (o *OrderDetailB2BLinesInnerShipmentDetailsInner) SetInvoiceDate(v string)`

SetInvoiceDate sets InvoiceDate field to given value.

### HasInvoiceDate

`func (o *OrderDetailB2BLinesInnerShipmentDetailsInner) HasInvoiceDate() bool`

HasInvoiceDate returns a boolean if a field has been set.

### GetCarrierDetails

`func (o *OrderDetailB2BLinesInnerShipmentDetailsInner) GetCarrierDetails() []OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner`

GetCarrierDetails returns the CarrierDetails field if non-nil, zero value otherwise.

### GetCarrierDetailsOk

`func (o *OrderDetailB2BLinesInnerShipmentDetailsInner) GetCarrierDetailsOk() (*[]OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner, bool)`

GetCarrierDetailsOk returns a tuple with the CarrierDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierDetails

`func (o *OrderDetailB2BLinesInnerShipmentDetailsInner) SetCarrierDetails(v []OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner)`

SetCarrierDetails sets CarrierDetails field to given value.

### HasCarrierDetails

`func (o *OrderDetailB2BLinesInnerShipmentDetailsInner) HasCarrierDetails() bool`

HasCarrierDetails returns a boolean if a field has been set.

### SetCarrierDetailsNil

`func (o *OrderDetailB2BLinesInnerShipmentDetailsInner) SetCarrierDetailsNil(b bool)`

 SetCarrierDetailsNil sets the value for CarrierDetails to be an explicit nil

### UnsetCarrierDetails
`func (o *OrderDetailB2BLinesInnerShipmentDetailsInner) UnsetCarrierDetails()`

UnsetCarrierDetails ensures that no value is present for CarrierDetails, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


