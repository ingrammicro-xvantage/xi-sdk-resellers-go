# OrderCreateV7RequestAdditionalAttributesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttributeName** | Pointer to **NullableString** | **allowPartialOrder:** Allow orders with failed lines.    **DpasRating:** DX rating by Department of Defense is the highest rating by the highest offices and meant to be top priority. DO any other gov offices at the federal level to priotize.    **DpasProgramId:** Identifies the actual agency that signed off on the DPAS priority.    **allowDuplicateCustomerOrderNumber:** Allow orders with duplicate customer PO numbers. Enables resellers to have the same PO number for multiple orders.     **euDepId:** DEP ID would be the &#39;End User DEP/ABM Organization ID&#39; up to 32 characters and is assigned by Apple.    **depOrderNbr:** depordernbr is &#39;End User PO to reseller&#39; Can appear in message lines or dedicated end user po#.    **govtProgramType:** Program type, “PA” for government orders, “ED” for education order.    **govtEndUserType:** Type of end user of the program. F &#x3D; Federal, S &#x3D; State, E &#x3D; Local, K &#x3D; K-12 school, and H &#x3D; Higher Education    **govtSolicitationNumber:** Education order’s contract number    **govtPublicPrivateCode:** Determines TAX / NO TAX.   &#39;P&#39; PUBLIC SECTOR,   &#39;R&#39; PRIVATE SECTOR.  Value needs only to be provided for EDUCATION order.    **govtEndUserData:** Name of the End user of the program. For example, STATE OF OHIO, CHICAGO SCHOOLDISTRICT etc.    **govtEndUserPostalCode:** 9 CHAR FIELD / Zip Code of the End user of the order.    **dynamicMessageLine1:** Custom Dynamic Message line 1.    **allowOrderOnCustomerHold:** Boolean value flag which allows a customer to create an order with the hold status. | [optional] 
**AttributeValue** | Pointer to **NullableString** | The attribute field data. | [optional] 

## Methods

### NewOrderCreateV7RequestAdditionalAttributesInner

`func NewOrderCreateV7RequestAdditionalAttributesInner() *OrderCreateV7RequestAdditionalAttributesInner`

NewOrderCreateV7RequestAdditionalAttributesInner instantiates a new OrderCreateV7RequestAdditionalAttributesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderCreateV7RequestAdditionalAttributesInnerWithDefaults

`func NewOrderCreateV7RequestAdditionalAttributesInnerWithDefaults() *OrderCreateV7RequestAdditionalAttributesInner`

NewOrderCreateV7RequestAdditionalAttributesInnerWithDefaults instantiates a new OrderCreateV7RequestAdditionalAttributesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributeName

`func (o *OrderCreateV7RequestAdditionalAttributesInner) GetAttributeName() string`

GetAttributeName returns the AttributeName field if non-nil, zero value otherwise.

### GetAttributeNameOk

`func (o *OrderCreateV7RequestAdditionalAttributesInner) GetAttributeNameOk() (*string, bool)`

GetAttributeNameOk returns a tuple with the AttributeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeName

`func (o *OrderCreateV7RequestAdditionalAttributesInner) SetAttributeName(v string)`

SetAttributeName sets AttributeName field to given value.

### HasAttributeName

`func (o *OrderCreateV7RequestAdditionalAttributesInner) HasAttributeName() bool`

HasAttributeName returns a boolean if a field has been set.

### SetAttributeNameNil

`func (o *OrderCreateV7RequestAdditionalAttributesInner) SetAttributeNameNil(b bool)`

 SetAttributeNameNil sets the value for AttributeName to be an explicit nil

### UnsetAttributeName
`func (o *OrderCreateV7RequestAdditionalAttributesInner) UnsetAttributeName()`

UnsetAttributeName ensures that no value is present for AttributeName, not even an explicit nil
### GetAttributeValue

`func (o *OrderCreateV7RequestAdditionalAttributesInner) GetAttributeValue() string`

GetAttributeValue returns the AttributeValue field if non-nil, zero value otherwise.

### GetAttributeValueOk

`func (o *OrderCreateV7RequestAdditionalAttributesInner) GetAttributeValueOk() (*string, bool)`

GetAttributeValueOk returns a tuple with the AttributeValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeValue

`func (o *OrderCreateV7RequestAdditionalAttributesInner) SetAttributeValue(v string)`

SetAttributeValue sets AttributeValue field to given value.

### HasAttributeValue

`func (o *OrderCreateV7RequestAdditionalAttributesInner) HasAttributeValue() bool`

HasAttributeValue returns a boolean if a field has been set.

### SetAttributeValueNil

`func (o *OrderCreateV7RequestAdditionalAttributesInner) SetAttributeValueNil(b bool)`

 SetAttributeValueNil sets the value for AttributeValue to be an explicit nil

### UnsetAttributeValue
`func (o *OrderCreateV7RequestAdditionalAttributesInner) UnsetAttributeValue()`

UnsetAttributeValue ensures that no value is present for AttributeValue, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


