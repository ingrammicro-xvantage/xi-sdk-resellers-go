# OrderCreateRequestShipmentDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CarrierCode** | Pointer to **string** | The code for the shipping carrier for the line item. | [optional] 
**FreightAccountNumber** | Pointer to **string** | The reseller &#39;s shipping account number with carrier. Used to bill the shipping carrier directly from the reseller&#39;s account with the carrier. | [optional] 
**ShipComplete** | Pointer to **string** | Specifies whether the shipment will be shipped only when all products are fulfilled. The value of this field along with acceptBackOrder field decides the value of backorderflag. If this field is set, acceptBackOrder field is ignored. Possible values for this field are true, C, P, E.    With value as true or C, backorderflag will be set as C.    With value as P, or E, backorderflag will be set as P or E respectively.    C &#x3D; Ship complete at distribution level.    P &#x3D; ship complete at line level.    E &#x3D; ship complete across all distributions.  | [optional] 
**RequestedDeliveryDate** | Pointer to **string** | The reseller-requested delivery date in UTC format. Delivery date is not guaranteed. | [optional] 
**SignatureRequired** | Pointer to **bool** | Specifies whether a signature is required for delivery. Default is False. | [optional] 
**ShippingInstructions** | Pointer to **string** |  | [optional] 

## Methods

### NewOrderCreateRequestShipmentDetails

`func NewOrderCreateRequestShipmentDetails() *OrderCreateRequestShipmentDetails`

NewOrderCreateRequestShipmentDetails instantiates a new OrderCreateRequestShipmentDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderCreateRequestShipmentDetailsWithDefaults

`func NewOrderCreateRequestShipmentDetailsWithDefaults() *OrderCreateRequestShipmentDetails`

NewOrderCreateRequestShipmentDetailsWithDefaults instantiates a new OrderCreateRequestShipmentDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCarrierCode

`func (o *OrderCreateRequestShipmentDetails) GetCarrierCode() string`

GetCarrierCode returns the CarrierCode field if non-nil, zero value otherwise.

### GetCarrierCodeOk

`func (o *OrderCreateRequestShipmentDetails) GetCarrierCodeOk() (*string, bool)`

GetCarrierCodeOk returns a tuple with the CarrierCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierCode

`func (o *OrderCreateRequestShipmentDetails) SetCarrierCode(v string)`

SetCarrierCode sets CarrierCode field to given value.

### HasCarrierCode

`func (o *OrderCreateRequestShipmentDetails) HasCarrierCode() bool`

HasCarrierCode returns a boolean if a field has been set.

### GetFreightAccountNumber

`func (o *OrderCreateRequestShipmentDetails) GetFreightAccountNumber() string`

GetFreightAccountNumber returns the FreightAccountNumber field if non-nil, zero value otherwise.

### GetFreightAccountNumberOk

`func (o *OrderCreateRequestShipmentDetails) GetFreightAccountNumberOk() (*string, bool)`

GetFreightAccountNumberOk returns a tuple with the FreightAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreightAccountNumber

`func (o *OrderCreateRequestShipmentDetails) SetFreightAccountNumber(v string)`

SetFreightAccountNumber sets FreightAccountNumber field to given value.

### HasFreightAccountNumber

`func (o *OrderCreateRequestShipmentDetails) HasFreightAccountNumber() bool`

HasFreightAccountNumber returns a boolean if a field has been set.

### GetShipComplete

`func (o *OrderCreateRequestShipmentDetails) GetShipComplete() string`

GetShipComplete returns the ShipComplete field if non-nil, zero value otherwise.

### GetShipCompleteOk

`func (o *OrderCreateRequestShipmentDetails) GetShipCompleteOk() (*string, bool)`

GetShipCompleteOk returns a tuple with the ShipComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipComplete

`func (o *OrderCreateRequestShipmentDetails) SetShipComplete(v string)`

SetShipComplete sets ShipComplete field to given value.

### HasShipComplete

`func (o *OrderCreateRequestShipmentDetails) HasShipComplete() bool`

HasShipComplete returns a boolean if a field has been set.

### GetRequestedDeliveryDate

`func (o *OrderCreateRequestShipmentDetails) GetRequestedDeliveryDate() string`

GetRequestedDeliveryDate returns the RequestedDeliveryDate field if non-nil, zero value otherwise.

### GetRequestedDeliveryDateOk

`func (o *OrderCreateRequestShipmentDetails) GetRequestedDeliveryDateOk() (*string, bool)`

GetRequestedDeliveryDateOk returns a tuple with the RequestedDeliveryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedDeliveryDate

`func (o *OrderCreateRequestShipmentDetails) SetRequestedDeliveryDate(v string)`

SetRequestedDeliveryDate sets RequestedDeliveryDate field to given value.

### HasRequestedDeliveryDate

`func (o *OrderCreateRequestShipmentDetails) HasRequestedDeliveryDate() bool`

HasRequestedDeliveryDate returns a boolean if a field has been set.

### GetSignatureRequired

`func (o *OrderCreateRequestShipmentDetails) GetSignatureRequired() bool`

GetSignatureRequired returns the SignatureRequired field if non-nil, zero value otherwise.

### GetSignatureRequiredOk

`func (o *OrderCreateRequestShipmentDetails) GetSignatureRequiredOk() (*bool, bool)`

GetSignatureRequiredOk returns a tuple with the SignatureRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureRequired

`func (o *OrderCreateRequestShipmentDetails) SetSignatureRequired(v bool)`

SetSignatureRequired sets SignatureRequired field to given value.

### HasSignatureRequired

`func (o *OrderCreateRequestShipmentDetails) HasSignatureRequired() bool`

HasSignatureRequired returns a boolean if a field has been set.

### GetShippingInstructions

`func (o *OrderCreateRequestShipmentDetails) GetShippingInstructions() string`

GetShippingInstructions returns the ShippingInstructions field if non-nil, zero value otherwise.

### GetShippingInstructionsOk

`func (o *OrderCreateRequestShipmentDetails) GetShippingInstructionsOk() (*string, bool)`

GetShippingInstructionsOk returns a tuple with the ShippingInstructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingInstructions

`func (o *OrderCreateRequestShipmentDetails) SetShippingInstructions(v string)`

SetShippingInstructions sets ShippingInstructions field to given value.

### HasShippingInstructions

`func (o *OrderCreateRequestShipmentDetails) HasShippingInstructions() bool`

HasShippingInstructions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


