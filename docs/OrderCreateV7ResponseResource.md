# OrderCreateV7ResponseResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerOrderNumber** | Pointer to **string** | The reseller&#39;s unique PO/Order number. | [optional] 
**BillToAddressId** | Pointer to **string** | Suffix used to identify billing address. Created during onboarding. Resellers are provided with one or more address IDs depending on how many bill to addresses they need for various flooring companies they are using for credit | [optional] 
**OrderSplit** | Pointer to **bool** | true for multiple orders | [optional] 
**ProcessedPartially** | Pointer to **bool** | true for partial order succesfully placed | [optional] 
**PurchaseOrderTotal** | Pointer to **float32** | Total of all the orders including taxes and fees. | [optional] 
**ShipToInfo** | Pointer to [**OrderCreateV7ResponseResourceShipToInfo**](OrderCreateV7ResponseResourceShipToInfo.md) |  | [optional] 
**Orders** | Pointer to [**[]OrderCreateV7ResponseResourceOrdersInner**](OrderCreateV7ResponseResourceOrdersInner.md) | Order-level details. | [optional] 

## Methods

### NewOrderCreateV7ResponseResource

`func NewOrderCreateV7ResponseResource() *OrderCreateV7ResponseResource`

NewOrderCreateV7ResponseResource instantiates a new OrderCreateV7ResponseResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderCreateV7ResponseResourceWithDefaults

`func NewOrderCreateV7ResponseResourceWithDefaults() *OrderCreateV7ResponseResource`

NewOrderCreateV7ResponseResourceWithDefaults instantiates a new OrderCreateV7ResponseResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerOrderNumber

`func (o *OrderCreateV7ResponseResource) GetCustomerOrderNumber() string`

GetCustomerOrderNumber returns the CustomerOrderNumber field if non-nil, zero value otherwise.

### GetCustomerOrderNumberOk

`func (o *OrderCreateV7ResponseResource) GetCustomerOrderNumberOk() (*string, bool)`

GetCustomerOrderNumberOk returns a tuple with the CustomerOrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerOrderNumber

`func (o *OrderCreateV7ResponseResource) SetCustomerOrderNumber(v string)`

SetCustomerOrderNumber sets CustomerOrderNumber field to given value.

### HasCustomerOrderNumber

`func (o *OrderCreateV7ResponseResource) HasCustomerOrderNumber() bool`

HasCustomerOrderNumber returns a boolean if a field has been set.

### GetBillToAddressId

`func (o *OrderCreateV7ResponseResource) GetBillToAddressId() string`

GetBillToAddressId returns the BillToAddressId field if non-nil, zero value otherwise.

### GetBillToAddressIdOk

`func (o *OrderCreateV7ResponseResource) GetBillToAddressIdOk() (*string, bool)`

GetBillToAddressIdOk returns a tuple with the BillToAddressId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillToAddressId

`func (o *OrderCreateV7ResponseResource) SetBillToAddressId(v string)`

SetBillToAddressId sets BillToAddressId field to given value.

### HasBillToAddressId

`func (o *OrderCreateV7ResponseResource) HasBillToAddressId() bool`

HasBillToAddressId returns a boolean if a field has been set.

### GetOrderSplit

`func (o *OrderCreateV7ResponseResource) GetOrderSplit() bool`

GetOrderSplit returns the OrderSplit field if non-nil, zero value otherwise.

### GetOrderSplitOk

`func (o *OrderCreateV7ResponseResource) GetOrderSplitOk() (*bool, bool)`

GetOrderSplitOk returns a tuple with the OrderSplit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderSplit

`func (o *OrderCreateV7ResponseResource) SetOrderSplit(v bool)`

SetOrderSplit sets OrderSplit field to given value.

### HasOrderSplit

`func (o *OrderCreateV7ResponseResource) HasOrderSplit() bool`

HasOrderSplit returns a boolean if a field has been set.

### GetProcessedPartially

`func (o *OrderCreateV7ResponseResource) GetProcessedPartially() bool`

GetProcessedPartially returns the ProcessedPartially field if non-nil, zero value otherwise.

### GetProcessedPartiallyOk

`func (o *OrderCreateV7ResponseResource) GetProcessedPartiallyOk() (*bool, bool)`

GetProcessedPartiallyOk returns a tuple with the ProcessedPartially field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessedPartially

`func (o *OrderCreateV7ResponseResource) SetProcessedPartially(v bool)`

SetProcessedPartially sets ProcessedPartially field to given value.

### HasProcessedPartially

`func (o *OrderCreateV7ResponseResource) HasProcessedPartially() bool`

HasProcessedPartially returns a boolean if a field has been set.

### GetPurchaseOrderTotal

`func (o *OrderCreateV7ResponseResource) GetPurchaseOrderTotal() float32`

GetPurchaseOrderTotal returns the PurchaseOrderTotal field if non-nil, zero value otherwise.

### GetPurchaseOrderTotalOk

`func (o *OrderCreateV7ResponseResource) GetPurchaseOrderTotalOk() (*float32, bool)`

GetPurchaseOrderTotalOk returns a tuple with the PurchaseOrderTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseOrderTotal

`func (o *OrderCreateV7ResponseResource) SetPurchaseOrderTotal(v float32)`

SetPurchaseOrderTotal sets PurchaseOrderTotal field to given value.

### HasPurchaseOrderTotal

`func (o *OrderCreateV7ResponseResource) HasPurchaseOrderTotal() bool`

HasPurchaseOrderTotal returns a boolean if a field has been set.

### GetShipToInfo

`func (o *OrderCreateV7ResponseResource) GetShipToInfo() OrderCreateV7ResponseResourceShipToInfo`

GetShipToInfo returns the ShipToInfo field if non-nil, zero value otherwise.

### GetShipToInfoOk

`func (o *OrderCreateV7ResponseResource) GetShipToInfoOk() (*OrderCreateV7ResponseResourceShipToInfo, bool)`

GetShipToInfoOk returns a tuple with the ShipToInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipToInfo

`func (o *OrderCreateV7ResponseResource) SetShipToInfo(v OrderCreateV7ResponseResourceShipToInfo)`

SetShipToInfo sets ShipToInfo field to given value.

### HasShipToInfo

`func (o *OrderCreateV7ResponseResource) HasShipToInfo() bool`

HasShipToInfo returns a boolean if a field has been set.

### GetOrders

`func (o *OrderCreateV7ResponseResource) GetOrders() []OrderCreateV7ResponseResourceOrdersInner`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *OrderCreateV7ResponseResource) GetOrdersOk() (*[]OrderCreateV7ResponseResourceOrdersInner, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *OrderCreateV7ResponseResource) SetOrders(v []OrderCreateV7ResponseResourceOrdersInner)`

SetOrders sets Orders field to given value.

### HasOrders

`func (o *OrderCreateV7ResponseResource) HasOrders() bool`

HasOrders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


