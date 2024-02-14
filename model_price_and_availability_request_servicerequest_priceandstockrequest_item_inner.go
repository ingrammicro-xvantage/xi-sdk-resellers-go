/*
XI Sdk Resellers

For Resellers. Who are looking to Innovate with Ingram Micro's API SolutionsAutomate your eCommerce with our offering of APIs and Webhooks to create a seamless experience for your customers.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package xi_sdk_resellers

import (
	"encoding/json"
)

// checks if the PriceAndAvailabilityRequestServicerequestPriceandstockrequestItemInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PriceAndAvailabilityRequestServicerequestPriceandstockrequestItemInner{}

// PriceAndAvailabilityRequestServicerequestPriceandstockrequestItemInner struct for PriceAndAvailabilityRequestServicerequestPriceandstockrequestItemInner
type PriceAndAvailabilityRequestServicerequestPriceandstockrequestItemInner struct {
	Index *int32 `json:"index,omitempty"`
	// Ingram Micro SKU number
	Ingrampartnumber *string `json:"ingrampartnumber,omitempty"`
	// Vendor/Manufacture Part Number
	Vendorpartnumber *string `json:"vendorpartnumber,omitempty"`
	// Universal Product code
	Upc *string `json:"upc,omitempty"`
	// Unique identoifier for the customer, needs custom setup.
	Customerpartnumber *string `json:"customerpartnumber,omitempty"`
	// Unique identity for Ingram Micro warehouses against which stock details are returned.
	Warehouseidlist []string `json:"warehouseidlist,omitempty"`
	Extendedvendorpartnumber *string `json:"extendedvendorpartnumber,omitempty"`
	Quantity *float32 `json:"quantity,omitempty"`
	Enduserid *string `json:"enduserid,omitempty"`
	Govtprogramtype *string `json:"govtprogramtype,omitempty"`
	Govtendusertype *string `json:"govtendusertype,omitempty"`
	Specialbidnumber *string `json:"specialbidnumber,omitempty"`
}

// NewPriceAndAvailabilityRequestServicerequestPriceandstockrequestItemInner instantiates a new PriceAndAvailabilityRequestServicerequestPriceandstockrequestItemInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPriceAndAvailabilityRequestServicerequestPriceandstockrequestItemInner() *PriceAndAvailabilityRequestServicerequestPriceandstockrequestItemInner {
	this := PriceAndAvailabilityRequestServicerequestPriceandstockrequestItemInner{}
	return &this
}

// NewPriceAndAvailabilityRequestServicerequestPriceandstockrequestItemInnerWithDefaults instantiates a new PriceAndAvailabilityRequestServicerequestPriceandstockrequestItemInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceAndAvailabilityRequestServicerequestPriceandstockrequestItemInnerWithDefaults() *PriceAndAvailabilityRequestServicerequestPriceandstockrequestItemInner {
	this := PriceAndAvailabilityRequestServicerequestPriceandstockrequestItemInner{}
	return &this
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *PriceAndAvailabilityRequestServicerequestPriceandstockrequestItemInner) GetIndex() int32 {
	if o == nil || IsNil(o.Index) {
		var ret int32
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceAndAvailabilityRequestServicerequestPriceandstockrequestItemInner) GetIndexOk() (*int32, bool) {
	if o == nil || IsNil(o.Index) {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *PriceAndAvailabilityRequestServicerequestPriceandstockrequestItemInner) HasIndex() bool {
	if o != nil && !IsNil(o.Index) {
		return true
	}

	return false
}

// SetIndex gets a reference to the given int32 and assigns it to the Index field.
func (o *PriceAndAvailabilityRequestServicerequestPriceandstockrequestItemInner) SetIndex(v int32) {
	o.Index = &v
}

// GetIngrampartnumber returns the Ingrampartnumber field value if set, zero value otherwise.
func (o *PriceAndAvailabilityRequestServicerequestPriceandstockrequestItemInner) GetIngrampartnumber() string {
	if o == nil || IsNil(o.Ingrampartnumber) {
		var ret string
		return ret
	}
	return *o.Ingrampartnumber
}

// GetIngrampartnumberOk returns a tuple with the Ingrampartnumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceAndAvailabilityRequestServicerequestPriceandstockrequestItemInner) GetIngrampartnumberOk() (*string, bool) {
	if o == nil || IsNil(o.Ingrampartnumber) {
		return nil, false
	}
	return o.Ingrampartnumber, true
}

// HasIngrampartnumber returns a boolean if a field has been set.
func (o *PriceAndAvailabilityRequestServicerequestPriceandstockrequestItemInner) HasIngrampartnumber() bool {
	if o != nil && !IsNil(o.Ingrampartnumber) {
		return true
	}

	return false
}

// SetIngrampartnumber gets a reference to the given string and assigns it to the Ingrampartnumber field.
func (o *PriceAndAvailabilityRequestServicerequestPriceandstockrequestItemInner) SetIngrampartnumber(v string) {
	o.Ingrampartnumber = &v
}

// GetVendorpartnumber returns the Vendorpartnumber field value if set, zero value otherwise.
func (o *PriceAndAvailabilityRequestServicerequestPriceandstockrequestItemInner) GetVendorpartnumber() string {
	if o == nil || IsNil(o.Vendorpartnumber) {
		var ret string
		return ret
	}
	return *o.Vendorpartnumber
}

// GetVendorpartnumberOk returns a tuple with the Vendorpartnumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceAndAvailabilityRequestServicerequestPriceandstockrequestItemInner) GetVendorpartnumberOk() (*string, bool) {
	if o == nil || IsNil(o.Vendorpartnumber) {
		return nil, false
	}
	return o.Vendorpartnumber, true
}

// HasVendorpartnumber returns a boolean if a field has been set.
func (o *PriceAndAvailabilityRequestServicerequestPriceandstockrequestItemInner) HasVendorpartnumber() bool {
	if o != nil && !IsNil(o.Vendorpartnumber) {
		return true
	}

	return false
}

// SetVendorpartnumber gets a reference to the given string and assigns it to the Vendorpartnumber field.
func (o *PriceAndAvailabilityRequestServicerequestPriceandstockrequestItemInner) SetVendorpartnumber(v string) {
	o.Vendorpartnumber = &v
}

// GetUpc returns the Upc field value if set, zero value otherwise.
func (o *PriceAndAvailabilityRequestServicerequestPriceandstockrequestItemInner) GetUpc() string {
	if o == nil || IsNil(o.Upc) {
		var ret string
		return ret
	}
	return *o.Upc
}

// GetUpcOk returns a tuple with the Upc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceAndAvailabilityRequestServicerequestPriceandstockrequestItemInner) GetUpcOk() (*string, bool) {
	if o == nil || IsNil(o.Upc) {
		return nil, false
	}
	return o.Upc, true
}

// HasUpc returns a boolean if a field has been set.
func (o *PriceAndAvailabilityRequestServicerequestPriceandstockrequestItemInner) HasUpc() bool {
	if o != nil && !IsNil(o.Upc) {
		return true
	}

	return false
}

// SetUpc gets a reference to the given string and assigns it to the Upc field.
func (o *PriceAndAvailabilityRequestServicerequestPriceandstockrequestItemInner) SetUpc(v string) {
	o.Upc = &v
}

// GetCustomerpartnumber returns the Customerpartnumber field value if set, zero value otherwise.
func (o *PriceAndAvailabilityRequestServicerequestPriceandstockrequestItemInner) GetCustomerpartnumber() string {
	if o == nil || IsNil(o.Customerpartnumber) {
		var ret string
		return ret
	}
	return *o.Customerpartnumber
}

// GetCustomerpartnumberOk returns a tuple with the Customerpartnumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceAndAvailabilityRequestServicerequestPriceandstockrequestItemInner) GetCustomerpartnumberOk() (*string, bool) {
	if o == nil || IsNil(o.Customerpartnumber) {
		return nil, false
	}
	return o.Customerpartnumber, true
}

// HasCustomerpartnumber returns a boolean if a field has been set.
func (o *PriceAndAvailabilityRequestServicerequestPriceandstockrequestItemInner) HasCustomerpartnumber() bool {
	if o != nil && !IsNil(o.Customerpartnumber) {
		return true
	}

	return false
}

// SetCustomerpartnumber gets a reference to the given string and assigns it to the Customerpartnumber field.
func (o *PriceAndAvailabilityRequestServicerequestPriceandstockrequestItemInner) SetCustomerpartnumber(v string) {
	o.Customerpartnumber = &v
}

// GetWarehouseidlist returns the Warehouseidlist field value if set, zero value otherwise.
func (o *PriceAndAvailabilityRequestServicerequestPriceandstockrequestItemInner) GetWarehouseidlist() []string {
	if o == nil || IsNil(o.Warehouseidlist) {
		var ret []string
		return ret
	}
	return o.Warehouseidlist
}

// GetWarehouseidlistOk returns a tuple with the Warehouseidlist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceAndAvailabilityRequestServicerequestPriceandstockrequestItemInner) GetWarehouseidlistOk() ([]string, bool) {
	if o == nil || IsNil(o.Warehouseidlist) {
		return nil, false
	}
	return o.Warehouseidlist, true
}

// HasWarehouseidlist returns a boolean if a field has been set.
func (o *PriceAndAvailabilityRequestServicerequestPriceandstockrequestItemInner) HasWarehouseidlist() bool {
	if o != nil && !IsNil(o.Warehouseidlist) {
		return true
	}

	return false
}

// SetWarehouseidlist gets a reference to the given []string and assigns it to the Warehouseidlist field.
func (o *PriceAndAvailabilityRequestServicerequestPriceandstockrequestItemInner) SetWarehouseidlist(v []string) {
	o.Warehouseidlist = v
}

// GetExtendedvendorpartnumber returns the Extendedvendorpartnumber field value if set, zero value otherwise.
func (o *PriceAndAvailabilityRequestServicerequestPriceandstockrequestItemInner) GetExtendedvendorpartnumber() string {
	if o == nil || IsNil(o.Extendedvendorpartnumber) {
		var ret string
		return ret
	}
	return *o.Extendedvendorpartnumber
}

// GetExtendedvendorpartnumberOk returns a tuple with the Extendedvendorpartnumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceAndAvailabilityRequestServicerequestPriceandstockrequestItemInner) GetExtendedvendorpartnumberOk() (*string, bool) {
	if o == nil || IsNil(o.Extendedvendorpartnumber) {
		return nil, false
	}
	return o.Extendedvendorpartnumber, true
}

// HasExtendedvendorpartnumber returns a boolean if a field has been set.
func (o *PriceAndAvailabilityRequestServicerequestPriceandstockrequestItemInner) HasExtendedvendorpartnumber() bool {
	if o != nil && !IsNil(o.Extendedvendorpartnumber) {
		return true
	}

	return false
}

// SetExtendedvendorpartnumber gets a reference to the given string and assigns it to the Extendedvendorpartnumber field.
func (o *PriceAndAvailabilityRequestServicerequestPriceandstockrequestItemInner) SetExtendedvendorpartnumber(v string) {
	o.Extendedvendorpartnumber = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *PriceAndAvailabilityRequestServicerequestPriceandstockrequestItemInner) GetQuantity() float32 {
	if o == nil || IsNil(o.Quantity) {
		var ret float32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceAndAvailabilityRequestServicerequestPriceandstockrequestItemInner) GetQuantityOk() (*float32, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *PriceAndAvailabilityRequestServicerequestPriceandstockrequestItemInner) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given float32 and assigns it to the Quantity field.
func (o *PriceAndAvailabilityRequestServicerequestPriceandstockrequestItemInner) SetQuantity(v float32) {
	o.Quantity = &v
}

// GetEnduserid returns the Enduserid field value if set, zero value otherwise.
func (o *PriceAndAvailabilityRequestServicerequestPriceandstockrequestItemInner) GetEnduserid() string {
	if o == nil || IsNil(o.Enduserid) {
		var ret string
		return ret
	}
	return *o.Enduserid
}

// GetEnduseridOk returns a tuple with the Enduserid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceAndAvailabilityRequestServicerequestPriceandstockrequestItemInner) GetEnduseridOk() (*string, bool) {
	if o == nil || IsNil(o.Enduserid) {
		return nil, false
	}
	return o.Enduserid, true
}

// HasEnduserid returns a boolean if a field has been set.
func (o *PriceAndAvailabilityRequestServicerequestPriceandstockrequestItemInner) HasEnduserid() bool {
	if o != nil && !IsNil(o.Enduserid) {
		return true
	}

	return false
}

// SetEnduserid gets a reference to the given string and assigns it to the Enduserid field.
func (o *PriceAndAvailabilityRequestServicerequestPriceandstockrequestItemInner) SetEnduserid(v string) {
	o.Enduserid = &v
}

// GetGovtprogramtype returns the Govtprogramtype field value if set, zero value otherwise.
func (o *PriceAndAvailabilityRequestServicerequestPriceandstockrequestItemInner) GetGovtprogramtype() string {
	if o == nil || IsNil(o.Govtprogramtype) {
		var ret string
		return ret
	}
	return *o.Govtprogramtype
}

// GetGovtprogramtypeOk returns a tuple with the Govtprogramtype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceAndAvailabilityRequestServicerequestPriceandstockrequestItemInner) GetGovtprogramtypeOk() (*string, bool) {
	if o == nil || IsNil(o.Govtprogramtype) {
		return nil, false
	}
	return o.Govtprogramtype, true
}

// HasGovtprogramtype returns a boolean if a field has been set.
func (o *PriceAndAvailabilityRequestServicerequestPriceandstockrequestItemInner) HasGovtprogramtype() bool {
	if o != nil && !IsNil(o.Govtprogramtype) {
		return true
	}

	return false
}

// SetGovtprogramtype gets a reference to the given string and assigns it to the Govtprogramtype field.
func (o *PriceAndAvailabilityRequestServicerequestPriceandstockrequestItemInner) SetGovtprogramtype(v string) {
	o.Govtprogramtype = &v
}

// GetGovtendusertype returns the Govtendusertype field value if set, zero value otherwise.
func (o *PriceAndAvailabilityRequestServicerequestPriceandstockrequestItemInner) GetGovtendusertype() string {
	if o == nil || IsNil(o.Govtendusertype) {
		var ret string
		return ret
	}
	return *o.Govtendusertype
}

// GetGovtendusertypeOk returns a tuple with the Govtendusertype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceAndAvailabilityRequestServicerequestPriceandstockrequestItemInner) GetGovtendusertypeOk() (*string, bool) {
	if o == nil || IsNil(o.Govtendusertype) {
		return nil, false
	}
	return o.Govtendusertype, true
}

// HasGovtendusertype returns a boolean if a field has been set.
func (o *PriceAndAvailabilityRequestServicerequestPriceandstockrequestItemInner) HasGovtendusertype() bool {
	if o != nil && !IsNil(o.Govtendusertype) {
		return true
	}

	return false
}

// SetGovtendusertype gets a reference to the given string and assigns it to the Govtendusertype field.
func (o *PriceAndAvailabilityRequestServicerequestPriceandstockrequestItemInner) SetGovtendusertype(v string) {
	o.Govtendusertype = &v
}

// GetSpecialbidnumber returns the Specialbidnumber field value if set, zero value otherwise.
func (o *PriceAndAvailabilityRequestServicerequestPriceandstockrequestItemInner) GetSpecialbidnumber() string {
	if o == nil || IsNil(o.Specialbidnumber) {
		var ret string
		return ret
	}
	return *o.Specialbidnumber
}

// GetSpecialbidnumberOk returns a tuple with the Specialbidnumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceAndAvailabilityRequestServicerequestPriceandstockrequestItemInner) GetSpecialbidnumberOk() (*string, bool) {
	if o == nil || IsNil(o.Specialbidnumber) {
		return nil, false
	}
	return o.Specialbidnumber, true
}

// HasSpecialbidnumber returns a boolean if a field has been set.
func (o *PriceAndAvailabilityRequestServicerequestPriceandstockrequestItemInner) HasSpecialbidnumber() bool {
	if o != nil && !IsNil(o.Specialbidnumber) {
		return true
	}

	return false
}

// SetSpecialbidnumber gets a reference to the given string and assigns it to the Specialbidnumber field.
func (o *PriceAndAvailabilityRequestServicerequestPriceandstockrequestItemInner) SetSpecialbidnumber(v string) {
	o.Specialbidnumber = &v
}

func (o PriceAndAvailabilityRequestServicerequestPriceandstockrequestItemInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PriceAndAvailabilityRequestServicerequestPriceandstockrequestItemInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Index) {
		toSerialize["index"] = o.Index
	}
	if !IsNil(o.Ingrampartnumber) {
		toSerialize["ingrampartnumber"] = o.Ingrampartnumber
	}
	if !IsNil(o.Vendorpartnumber) {
		toSerialize["vendorpartnumber"] = o.Vendorpartnumber
	}
	if !IsNil(o.Upc) {
		toSerialize["upc"] = o.Upc
	}
	if !IsNil(o.Customerpartnumber) {
		toSerialize["customerpartnumber"] = o.Customerpartnumber
	}
	if !IsNil(o.Warehouseidlist) {
		toSerialize["warehouseidlist"] = o.Warehouseidlist
	}
	if !IsNil(o.Extendedvendorpartnumber) {
		toSerialize["extendedvendorpartnumber"] = o.Extendedvendorpartnumber
	}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if !IsNil(o.Enduserid) {
		toSerialize["enduserid"] = o.Enduserid
	}
	if !IsNil(o.Govtprogramtype) {
		toSerialize["govtprogramtype"] = o.Govtprogramtype
	}
	if !IsNil(o.Govtendusertype) {
		toSerialize["govtendusertype"] = o.Govtendusertype
	}
	if !IsNil(o.Specialbidnumber) {
		toSerialize["specialbidnumber"] = o.Specialbidnumber
	}
	return toSerialize, nil
}

type NullablePriceAndAvailabilityRequestServicerequestPriceandstockrequestItemInner struct {
	value *PriceAndAvailabilityRequestServicerequestPriceandstockrequestItemInner
	isSet bool
}

func (v NullablePriceAndAvailabilityRequestServicerequestPriceandstockrequestItemInner) Get() *PriceAndAvailabilityRequestServicerequestPriceandstockrequestItemInner {
	return v.value
}

func (v *NullablePriceAndAvailabilityRequestServicerequestPriceandstockrequestItemInner) Set(val *PriceAndAvailabilityRequestServicerequestPriceandstockrequestItemInner) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceAndAvailabilityRequestServicerequestPriceandstockrequestItemInner) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceAndAvailabilityRequestServicerequestPriceandstockrequestItemInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceAndAvailabilityRequestServicerequestPriceandstockrequestItemInner(val *PriceAndAvailabilityRequestServicerequestPriceandstockrequestItemInner) *NullablePriceAndAvailabilityRequestServicerequestPriceandstockrequestItemInner {
	return &NullablePriceAndAvailabilityRequestServicerequestPriceandstockrequestItemInner{value: val, isSet: true}
}

func (v NullablePriceAndAvailabilityRequestServicerequestPriceandstockrequestItemInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceAndAvailabilityRequestServicerequestPriceandstockrequestItemInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


