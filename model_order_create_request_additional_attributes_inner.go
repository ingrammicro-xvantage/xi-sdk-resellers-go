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

// checks if the OrderCreateRequestAdditionalAttributesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderCreateRequestAdditionalAttributesInner{}

// OrderCreateRequestAdditionalAttributesInner struct for OrderCreateRequestAdditionalAttributesInner
type OrderCreateRequestAdditionalAttributesInner struct {
	// allowPartialOrder: Allow orders with failed lines. (SAP) Depends on backorder settings.    DpasRating: DX rating by Department of Defense is the highest rating by the highest offices and meant to be top priority. DO any other gov offices at the federal level to priotize.    DpasProgramId: Identifies the actual agency that signed off on the DPAS priority.    allowDuplicateCustomerOrderNumber: Allow orders with duplicate customer PO numbers. Enables resellers to have the same PO number for multiple orders.     channelCode: Determine storage location for Markeplace(SAP) for different orderTypes.    customerPOType: Used for pricing, similar to orderType. Possible SAP values- ZXML and ZWEB.    storageLocation: Determine the location of the product stock in SAP for Marketplaces.    soldTo: To be used in cases when Sold-To is different than Customer ID.    orderType: Order Type[SAP]- ZOR and ZLCN.    endUserSearchTerm: Search ID for a end user contact is used in SAP to determine the contact name.    Z101: Used for end customer details such as name, address, phone, etc. This information flows to SAP and is used by warehouse.    euDepId: DEP ID would be the 'End User DEP/ABM Organization ID' up to 32 characters and is assigned by Apple.    depOrderNbr: depordernbr is 'End User PO to reseller' Can appear in message lines or dedicated end user po#.    govtProgramType: Program type, “PA” for government orders, “ED” for education order.    govtEndUserType: Type of end user of the program. F = Federal, S = State, E = Local, K = K-12 school, and H = Higher Education    govtSolicitationNumber: Education order’s contract number    govtPublicPrivateCode: Determines TAX / NO TAX.   'P' PUBLIC SECTOR,   'R' PRIVATE SECTOR.  Value needs only to be provided for EDUCATION order.    govtEndUserData: Name of the End user of the program. For example, STATE OF OHIO, CHICAGO SCHOOLDISTRICT etc.    govtEndUserPostalCode: 9 CHAR FIELD / Zip Code of the End user of the order.    dynamicMessageLine1: Custom Dynamic Message line 1.    allowOrderOnCustomerHold: Boolean value flag which allows a customer to create an order with the hold status.
	AttributeName *string `json:"attributeName,omitempty"`
	// attributefield data
	AttributeValue *string `json:"attributeValue,omitempty"`
}

// NewOrderCreateRequestAdditionalAttributesInner instantiates a new OrderCreateRequestAdditionalAttributesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderCreateRequestAdditionalAttributesInner() *OrderCreateRequestAdditionalAttributesInner {
	this := OrderCreateRequestAdditionalAttributesInner{}
	return &this
}

// NewOrderCreateRequestAdditionalAttributesInnerWithDefaults instantiates a new OrderCreateRequestAdditionalAttributesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderCreateRequestAdditionalAttributesInnerWithDefaults() *OrderCreateRequestAdditionalAttributesInner {
	this := OrderCreateRequestAdditionalAttributesInner{}
	return &this
}

// GetAttributeName returns the AttributeName field value if set, zero value otherwise.
func (o *OrderCreateRequestAdditionalAttributesInner) GetAttributeName() string {
	if o == nil || IsNil(o.AttributeName) {
		var ret string
		return ret
	}
	return *o.AttributeName
}

// GetAttributeNameOk returns a tuple with the AttributeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateRequestAdditionalAttributesInner) GetAttributeNameOk() (*string, bool) {
	if o == nil || IsNil(o.AttributeName) {
		return nil, false
	}
	return o.AttributeName, true
}

// HasAttributeName returns a boolean if a field has been set.
func (o *OrderCreateRequestAdditionalAttributesInner) HasAttributeName() bool {
	if o != nil && !IsNil(o.AttributeName) {
		return true
	}

	return false
}

// SetAttributeName gets a reference to the given string and assigns it to the AttributeName field.
func (o *OrderCreateRequestAdditionalAttributesInner) SetAttributeName(v string) {
	o.AttributeName = &v
}

// GetAttributeValue returns the AttributeValue field value if set, zero value otherwise.
func (o *OrderCreateRequestAdditionalAttributesInner) GetAttributeValue() string {
	if o == nil || IsNil(o.AttributeValue) {
		var ret string
		return ret
	}
	return *o.AttributeValue
}

// GetAttributeValueOk returns a tuple with the AttributeValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateRequestAdditionalAttributesInner) GetAttributeValueOk() (*string, bool) {
	if o == nil || IsNil(o.AttributeValue) {
		return nil, false
	}
	return o.AttributeValue, true
}

// HasAttributeValue returns a boolean if a field has been set.
func (o *OrderCreateRequestAdditionalAttributesInner) HasAttributeValue() bool {
	if o != nil && !IsNil(o.AttributeValue) {
		return true
	}

	return false
}

// SetAttributeValue gets a reference to the given string and assigns it to the AttributeValue field.
func (o *OrderCreateRequestAdditionalAttributesInner) SetAttributeValue(v string) {
	o.AttributeValue = &v
}

func (o OrderCreateRequestAdditionalAttributesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderCreateRequestAdditionalAttributesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AttributeName) {
		toSerialize["attributeName"] = o.AttributeName
	}
	if !IsNil(o.AttributeValue) {
		toSerialize["attributeValue"] = o.AttributeValue
	}
	return toSerialize, nil
}

type NullableOrderCreateRequestAdditionalAttributesInner struct {
	value *OrderCreateRequestAdditionalAttributesInner
	isSet bool
}

func (v NullableOrderCreateRequestAdditionalAttributesInner) Get() *OrderCreateRequestAdditionalAttributesInner {
	return v.value
}

func (v *NullableOrderCreateRequestAdditionalAttributesInner) Set(val *OrderCreateRequestAdditionalAttributesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderCreateRequestAdditionalAttributesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderCreateRequestAdditionalAttributesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderCreateRequestAdditionalAttributesInner(val *OrderCreateRequestAdditionalAttributesInner) *NullableOrderCreateRequestAdditionalAttributesInner {
	return &NullableOrderCreateRequestAdditionalAttributesInner{value: val, isSet: true}
}

func (v NullableOrderCreateRequestAdditionalAttributesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderCreateRequestAdditionalAttributesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


