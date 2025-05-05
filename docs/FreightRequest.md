# FreightRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillToAddressId** | Pointer to **interface{}** | Suffix used to identify billing address. Created during onboarding. Resellers are provided with one or more address IDs depending on how many bill to addresses they need for various flooring companies they are using for credit. | [optional] 
**ShipToAddressId** | Pointer to **string** | The ID references the reseller&#39;s address in Ingram Micro&#39;s system for shipping. Provided to resellers during the onboarding process. | [optional] 
**ShipToAddress** | Pointer to [**FreightRequestShipToAddress**](FreightRequestShipToAddress.md) |  | [optional] 
**Lines** | Pointer to [**[]FreightRequestLinesInner**](FreightRequestLinesInner.md) |  | [optional] 

## Methods

### NewFreightRequest

`func NewFreightRequest() *FreightRequest`

NewFreightRequest instantiates a new FreightRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFreightRequestWithDefaults

`func NewFreightRequestWithDefaults() *FreightRequest`

NewFreightRequestWithDefaults instantiates a new FreightRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillToAddressId

`func (o *FreightRequest) GetBillToAddressId() interface{}`

GetBillToAddressId returns the BillToAddressId field if non-nil, zero value otherwise.

### GetBillToAddressIdOk

`func (o *FreightRequest) GetBillToAddressIdOk() (*interface{}, bool)`

GetBillToAddressIdOk returns a tuple with the BillToAddressId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillToAddressId

`func (o *FreightRequest) SetBillToAddressId(v interface{})`

SetBillToAddressId sets BillToAddressId field to given value.

### HasBillToAddressId

`func (o *FreightRequest) HasBillToAddressId() bool`

HasBillToAddressId returns a boolean if a field has been set.

### SetBillToAddressIdNil

`func (o *FreightRequest) SetBillToAddressIdNil(b bool)`

 SetBillToAddressIdNil sets the value for BillToAddressId to be an explicit nil

### UnsetBillToAddressId
`func (o *FreightRequest) UnsetBillToAddressId()`

UnsetBillToAddressId ensures that no value is present for BillToAddressId, not even an explicit nil
### GetShipToAddressId

`func (o *FreightRequest) GetShipToAddressId() string`

GetShipToAddressId returns the ShipToAddressId field if non-nil, zero value otherwise.

### GetShipToAddressIdOk

`func (o *FreightRequest) GetShipToAddressIdOk() (*string, bool)`

GetShipToAddressIdOk returns a tuple with the ShipToAddressId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipToAddressId

`func (o *FreightRequest) SetShipToAddressId(v string)`

SetShipToAddressId sets ShipToAddressId field to given value.

### HasShipToAddressId

`func (o *FreightRequest) HasShipToAddressId() bool`

HasShipToAddressId returns a boolean if a field has been set.

### GetShipToAddress

`func (o *FreightRequest) GetShipToAddress() FreightRequestShipToAddress`

GetShipToAddress returns the ShipToAddress field if non-nil, zero value otherwise.

### GetShipToAddressOk

`func (o *FreightRequest) GetShipToAddressOk() (*FreightRequestShipToAddress, bool)`

GetShipToAddressOk returns a tuple with the ShipToAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipToAddress

`func (o *FreightRequest) SetShipToAddress(v FreightRequestShipToAddress)`

SetShipToAddress sets ShipToAddress field to given value.

### HasShipToAddress

`func (o *FreightRequest) HasShipToAddress() bool`

HasShipToAddress returns a boolean if a field has been set.

### GetLines

`func (o *FreightRequest) GetLines() []FreightRequestLinesInner`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *FreightRequest) GetLinesOk() (*[]FreightRequestLinesInner, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *FreightRequest) SetLines(v []FreightRequestLinesInner)`

SetLines sets Lines field to given value.

### HasLines

`func (o *FreightRequest) HasLines() bool`

HasLines returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


