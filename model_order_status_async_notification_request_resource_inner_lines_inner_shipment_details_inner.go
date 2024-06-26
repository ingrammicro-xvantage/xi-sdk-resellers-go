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

// checks if the OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner{}

// OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner struct for OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner
type OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner struct {
	// The date the line item was shipped.
	ShipmentDate *string `json:"shipmentDate,omitempty"`
	// The ID of the warehouse the product will ship from.
	ShipFromWarehouseId *string `json:"shipFromWarehouseId,omitempty"`
	// \"\"
	WarehouseName *string `json:"warehouseName,omitempty"`
	// The carrier code for the shipment containing the  line item.
	CarrierCode *string `json:"carrierCode,omitempty"`
	// The name of the carrier of the shipment containing   the line item.
	CarrierName *string `json:"carrierName,omitempty"`
	PackageDetails []OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInnerPackageDetailsInner `json:"packageDetails,omitempty"`
}

// NewOrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner instantiates a new OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner() *OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner {
	this := OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner{}
	return &this
}

// NewOrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInnerWithDefaults instantiates a new OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInnerWithDefaults() *OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner {
	this := OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner{}
	return &this
}

// GetShipmentDate returns the ShipmentDate field value if set, zero value otherwise.
func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner) GetShipmentDate() string {
	if o == nil || IsNil(o.ShipmentDate) {
		var ret string
		return ret
	}
	return *o.ShipmentDate
}

// GetShipmentDateOk returns a tuple with the ShipmentDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner) GetShipmentDateOk() (*string, bool) {
	if o == nil || IsNil(o.ShipmentDate) {
		return nil, false
	}
	return o.ShipmentDate, true
}

// HasShipmentDate returns a boolean if a field has been set.
func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner) HasShipmentDate() bool {
	if o != nil && !IsNil(o.ShipmentDate) {
		return true
	}

	return false
}

// SetShipmentDate gets a reference to the given string and assigns it to the ShipmentDate field.
func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner) SetShipmentDate(v string) {
	o.ShipmentDate = &v
}

// GetShipFromWarehouseId returns the ShipFromWarehouseId field value if set, zero value otherwise.
func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner) GetShipFromWarehouseId() string {
	if o == nil || IsNil(o.ShipFromWarehouseId) {
		var ret string
		return ret
	}
	return *o.ShipFromWarehouseId
}

// GetShipFromWarehouseIdOk returns a tuple with the ShipFromWarehouseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner) GetShipFromWarehouseIdOk() (*string, bool) {
	if o == nil || IsNil(o.ShipFromWarehouseId) {
		return nil, false
	}
	return o.ShipFromWarehouseId, true
}

// HasShipFromWarehouseId returns a boolean if a field has been set.
func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner) HasShipFromWarehouseId() bool {
	if o != nil && !IsNil(o.ShipFromWarehouseId) {
		return true
	}

	return false
}

// SetShipFromWarehouseId gets a reference to the given string and assigns it to the ShipFromWarehouseId field.
func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner) SetShipFromWarehouseId(v string) {
	o.ShipFromWarehouseId = &v
}

// GetWarehouseName returns the WarehouseName field value if set, zero value otherwise.
func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner) GetWarehouseName() string {
	if o == nil || IsNil(o.WarehouseName) {
		var ret string
		return ret
	}
	return *o.WarehouseName
}

// GetWarehouseNameOk returns a tuple with the WarehouseName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner) GetWarehouseNameOk() (*string, bool) {
	if o == nil || IsNil(o.WarehouseName) {
		return nil, false
	}
	return o.WarehouseName, true
}

// HasWarehouseName returns a boolean if a field has been set.
func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner) HasWarehouseName() bool {
	if o != nil && !IsNil(o.WarehouseName) {
		return true
	}

	return false
}

// SetWarehouseName gets a reference to the given string and assigns it to the WarehouseName field.
func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner) SetWarehouseName(v string) {
	o.WarehouseName = &v
}

// GetCarrierCode returns the CarrierCode field value if set, zero value otherwise.
func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner) GetCarrierCode() string {
	if o == nil || IsNil(o.CarrierCode) {
		var ret string
		return ret
	}
	return *o.CarrierCode
}

// GetCarrierCodeOk returns a tuple with the CarrierCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner) GetCarrierCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CarrierCode) {
		return nil, false
	}
	return o.CarrierCode, true
}

// HasCarrierCode returns a boolean if a field has been set.
func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner) HasCarrierCode() bool {
	if o != nil && !IsNil(o.CarrierCode) {
		return true
	}

	return false
}

// SetCarrierCode gets a reference to the given string and assigns it to the CarrierCode field.
func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner) SetCarrierCode(v string) {
	o.CarrierCode = &v
}

// GetCarrierName returns the CarrierName field value if set, zero value otherwise.
func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner) GetCarrierName() string {
	if o == nil || IsNil(o.CarrierName) {
		var ret string
		return ret
	}
	return *o.CarrierName
}

// GetCarrierNameOk returns a tuple with the CarrierName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner) GetCarrierNameOk() (*string, bool) {
	if o == nil || IsNil(o.CarrierName) {
		return nil, false
	}
	return o.CarrierName, true
}

// HasCarrierName returns a boolean if a field has been set.
func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner) HasCarrierName() bool {
	if o != nil && !IsNil(o.CarrierName) {
		return true
	}

	return false
}

// SetCarrierName gets a reference to the given string and assigns it to the CarrierName field.
func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner) SetCarrierName(v string) {
	o.CarrierName = &v
}

// GetPackageDetails returns the PackageDetails field value if set, zero value otherwise.
func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner) GetPackageDetails() []OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInnerPackageDetailsInner {
	if o == nil || IsNil(o.PackageDetails) {
		var ret []OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInnerPackageDetailsInner
		return ret
	}
	return o.PackageDetails
}

// GetPackageDetailsOk returns a tuple with the PackageDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner) GetPackageDetailsOk() ([]OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInnerPackageDetailsInner, bool) {
	if o == nil || IsNil(o.PackageDetails) {
		return nil, false
	}
	return o.PackageDetails, true
}

// HasPackageDetails returns a boolean if a field has been set.
func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner) HasPackageDetails() bool {
	if o != nil && !IsNil(o.PackageDetails) {
		return true
	}

	return false
}

// SetPackageDetails gets a reference to the given []OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInnerPackageDetailsInner and assigns it to the PackageDetails field.
func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner) SetPackageDetails(v []OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInnerPackageDetailsInner) {
	o.PackageDetails = v
}

func (o OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ShipmentDate) {
		toSerialize["shipmentDate"] = o.ShipmentDate
	}
	if !IsNil(o.ShipFromWarehouseId) {
		toSerialize["shipFromWarehouseId"] = o.ShipFromWarehouseId
	}
	if !IsNil(o.WarehouseName) {
		toSerialize["warehouseName"] = o.WarehouseName
	}
	if !IsNil(o.CarrierCode) {
		toSerialize["carrierCode"] = o.CarrierCode
	}
	if !IsNil(o.CarrierName) {
		toSerialize["carrierName"] = o.CarrierName
	}
	if !IsNil(o.PackageDetails) {
		toSerialize["packageDetails"] = o.PackageDetails
	}
	return toSerialize, nil
}

type NullableOrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner struct {
	value *OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner
	isSet bool
}

func (v NullableOrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner) Get() *OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner {
	return v.value
}

func (v *NullableOrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner) Set(val *OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner(val *OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner) *NullableOrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner {
	return &NullableOrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner{value: val, isSet: true}
}

func (v NullableOrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


