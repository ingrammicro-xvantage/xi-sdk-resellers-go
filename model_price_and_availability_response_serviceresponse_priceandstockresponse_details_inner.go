/*
Reseller API

For Resellers. <br> Who are looking to Innovate with Ingram Micro's API SolutionsAutomate your eCommerce with our offering of APIs and Webhooks to create a seamless experience for your customers.

API version: 6.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package xi_sdk_resellers

import (
	"encoding/json"
)

// checks if the PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner{}

// PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner struct for PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner
type PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner struct {
	// SUCCESS or FAILED
	Itemstatus *string `json:"itemstatus,omitempty"`
	// Description of itemstatus
	Statusmessage *string `json:"statusmessage,omitempty"`
	// Ingram Micro part number
	Ingrampartnumber *string `json:"ingrampartnumber,omitempty"`
	// Manufacturer/Vendor part number
	Vendorpartnumber *string `json:"vendorpartnumber,omitempty"`
	Globalskuid *string `json:"globalskuid,omitempty"`
	// Customer specific price for the product, excluding taxes
	Customerprice *float32 `json:"customerprice,omitempty"`
	// Product description part 1
	Partdescription1 *string `json:"partdescription1,omitempty"`
	// Product description part 2
	Partdescription2 *string `json:"partdescription2,omitempty"`
	Vendornumber *string `json:"vendornumber,omitempty"`
	// Name of the vendor
	Vendorname *string `json:"vendorname,omitempty"`
	Cpucode *string `json:"cpucode,omitempty"`
	// Ingram Micro assigned product classification -  A-Stocked product in all IM warehouses, B-Limited stock in IM warehouses, C-Stocked in fewer wareshouses, D-Ingram discontinued, E-Planned to be phased out as per the vendor, F-Carried for specific customer as per the contract, N-New SKU, O-Discontinued to be liquidated, S-Order for specialized demand, V-Discontinued by vendor, X-Direct Ship products from vendor
	Class *string `json:"class,omitempty"`
	// Identifies if the SKU has been discontinued.
	Skustatus *string `json:"skustatus,omitempty"`
	Mediacpu *string `json:"mediacpu,omitempty"`
	Categorysubcategory *string `json:"categorysubcategory,omitempty"`
	Retailprice *float32 `json:"retailprice,omitempty"`
	Newmedia *string `json:"newmedia,omitempty"`
	// Y - End user required N - Not required End user
	Enduserrequired *string `json:"enduserrequired,omitempty"`
	// Y- Allow Backorder Flag N- Not allowed
	Backorderflag *string `json:"backorderflag,omitempty"`
	Skuauthorized *string `json:"skuauthorized,omitempty"`
	Extendedvendorpartnumber *string `json:"extendedvendorpartnumber,omitempty"`
	Warehousedetails []WarehouseListType `json:"warehousedetails,omitempty"`
}

// NewPriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner instantiates a new PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner() *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner {
	this := PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner{}
	return &this
}

// NewPriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInnerWithDefaults instantiates a new PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInnerWithDefaults() *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner {
	this := PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner{}
	return &this
}

// GetItemstatus returns the Itemstatus field value if set, zero value otherwise.
func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) GetItemstatus() string {
	if o == nil || IsNil(o.Itemstatus) {
		var ret string
		return ret
	}
	return *o.Itemstatus
}

// GetItemstatusOk returns a tuple with the Itemstatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) GetItemstatusOk() (*string, bool) {
	if o == nil || IsNil(o.Itemstatus) {
		return nil, false
	}
	return o.Itemstatus, true
}

// HasItemstatus returns a boolean if a field has been set.
func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) HasItemstatus() bool {
	if o != nil && !IsNil(o.Itemstatus) {
		return true
	}

	return false
}

// SetItemstatus gets a reference to the given string and assigns it to the Itemstatus field.
func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) SetItemstatus(v string) {
	o.Itemstatus = &v
}

// GetStatusmessage returns the Statusmessage field value if set, zero value otherwise.
func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) GetStatusmessage() string {
	if o == nil || IsNil(o.Statusmessage) {
		var ret string
		return ret
	}
	return *o.Statusmessage
}

// GetStatusmessageOk returns a tuple with the Statusmessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) GetStatusmessageOk() (*string, bool) {
	if o == nil || IsNil(o.Statusmessage) {
		return nil, false
	}
	return o.Statusmessage, true
}

// HasStatusmessage returns a boolean if a field has been set.
func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) HasStatusmessage() bool {
	if o != nil && !IsNil(o.Statusmessage) {
		return true
	}

	return false
}

// SetStatusmessage gets a reference to the given string and assigns it to the Statusmessage field.
func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) SetStatusmessage(v string) {
	o.Statusmessage = &v
}

// GetIngrampartnumber returns the Ingrampartnumber field value if set, zero value otherwise.
func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) GetIngrampartnumber() string {
	if o == nil || IsNil(o.Ingrampartnumber) {
		var ret string
		return ret
	}
	return *o.Ingrampartnumber
}

// GetIngrampartnumberOk returns a tuple with the Ingrampartnumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) GetIngrampartnumberOk() (*string, bool) {
	if o == nil || IsNil(o.Ingrampartnumber) {
		return nil, false
	}
	return o.Ingrampartnumber, true
}

// HasIngrampartnumber returns a boolean if a field has been set.
func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) HasIngrampartnumber() bool {
	if o != nil && !IsNil(o.Ingrampartnumber) {
		return true
	}

	return false
}

// SetIngrampartnumber gets a reference to the given string and assigns it to the Ingrampartnumber field.
func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) SetIngrampartnumber(v string) {
	o.Ingrampartnumber = &v
}

// GetVendorpartnumber returns the Vendorpartnumber field value if set, zero value otherwise.
func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) GetVendorpartnumber() string {
	if o == nil || IsNil(o.Vendorpartnumber) {
		var ret string
		return ret
	}
	return *o.Vendorpartnumber
}

// GetVendorpartnumberOk returns a tuple with the Vendorpartnumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) GetVendorpartnumberOk() (*string, bool) {
	if o == nil || IsNil(o.Vendorpartnumber) {
		return nil, false
	}
	return o.Vendorpartnumber, true
}

// HasVendorpartnumber returns a boolean if a field has been set.
func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) HasVendorpartnumber() bool {
	if o != nil && !IsNil(o.Vendorpartnumber) {
		return true
	}

	return false
}

// SetVendorpartnumber gets a reference to the given string and assigns it to the Vendorpartnumber field.
func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) SetVendorpartnumber(v string) {
	o.Vendorpartnumber = &v
}

// GetGlobalskuid returns the Globalskuid field value if set, zero value otherwise.
func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) GetGlobalskuid() string {
	if o == nil || IsNil(o.Globalskuid) {
		var ret string
		return ret
	}
	return *o.Globalskuid
}

// GetGlobalskuidOk returns a tuple with the Globalskuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) GetGlobalskuidOk() (*string, bool) {
	if o == nil || IsNil(o.Globalskuid) {
		return nil, false
	}
	return o.Globalskuid, true
}

// HasGlobalskuid returns a boolean if a field has been set.
func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) HasGlobalskuid() bool {
	if o != nil && !IsNil(o.Globalskuid) {
		return true
	}

	return false
}

// SetGlobalskuid gets a reference to the given string and assigns it to the Globalskuid field.
func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) SetGlobalskuid(v string) {
	o.Globalskuid = &v
}

// GetCustomerprice returns the Customerprice field value if set, zero value otherwise.
func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) GetCustomerprice() float32 {
	if o == nil || IsNil(o.Customerprice) {
		var ret float32
		return ret
	}
	return *o.Customerprice
}

// GetCustomerpriceOk returns a tuple with the Customerprice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) GetCustomerpriceOk() (*float32, bool) {
	if o == nil || IsNil(o.Customerprice) {
		return nil, false
	}
	return o.Customerprice, true
}

// HasCustomerprice returns a boolean if a field has been set.
func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) HasCustomerprice() bool {
	if o != nil && !IsNil(o.Customerprice) {
		return true
	}

	return false
}

// SetCustomerprice gets a reference to the given float32 and assigns it to the Customerprice field.
func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) SetCustomerprice(v float32) {
	o.Customerprice = &v
}

// GetPartdescription1 returns the Partdescription1 field value if set, zero value otherwise.
func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) GetPartdescription1() string {
	if o == nil || IsNil(o.Partdescription1) {
		var ret string
		return ret
	}
	return *o.Partdescription1
}

// GetPartdescription1Ok returns a tuple with the Partdescription1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) GetPartdescription1Ok() (*string, bool) {
	if o == nil || IsNil(o.Partdescription1) {
		return nil, false
	}
	return o.Partdescription1, true
}

// HasPartdescription1 returns a boolean if a field has been set.
func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) HasPartdescription1() bool {
	if o != nil && !IsNil(o.Partdescription1) {
		return true
	}

	return false
}

// SetPartdescription1 gets a reference to the given string and assigns it to the Partdescription1 field.
func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) SetPartdescription1(v string) {
	o.Partdescription1 = &v
}

// GetPartdescription2 returns the Partdescription2 field value if set, zero value otherwise.
func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) GetPartdescription2() string {
	if o == nil || IsNil(o.Partdescription2) {
		var ret string
		return ret
	}
	return *o.Partdescription2
}

// GetPartdescription2Ok returns a tuple with the Partdescription2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) GetPartdescription2Ok() (*string, bool) {
	if o == nil || IsNil(o.Partdescription2) {
		return nil, false
	}
	return o.Partdescription2, true
}

// HasPartdescription2 returns a boolean if a field has been set.
func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) HasPartdescription2() bool {
	if o != nil && !IsNil(o.Partdescription2) {
		return true
	}

	return false
}

// SetPartdescription2 gets a reference to the given string and assigns it to the Partdescription2 field.
func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) SetPartdescription2(v string) {
	o.Partdescription2 = &v
}

// GetVendornumber returns the Vendornumber field value if set, zero value otherwise.
func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) GetVendornumber() string {
	if o == nil || IsNil(o.Vendornumber) {
		var ret string
		return ret
	}
	return *o.Vendornumber
}

// GetVendornumberOk returns a tuple with the Vendornumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) GetVendornumberOk() (*string, bool) {
	if o == nil || IsNil(o.Vendornumber) {
		return nil, false
	}
	return o.Vendornumber, true
}

// HasVendornumber returns a boolean if a field has been set.
func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) HasVendornumber() bool {
	if o != nil && !IsNil(o.Vendornumber) {
		return true
	}

	return false
}

// SetVendornumber gets a reference to the given string and assigns it to the Vendornumber field.
func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) SetVendornumber(v string) {
	o.Vendornumber = &v
}

// GetVendorname returns the Vendorname field value if set, zero value otherwise.
func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) GetVendorname() string {
	if o == nil || IsNil(o.Vendorname) {
		var ret string
		return ret
	}
	return *o.Vendorname
}

// GetVendornameOk returns a tuple with the Vendorname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) GetVendornameOk() (*string, bool) {
	if o == nil || IsNil(o.Vendorname) {
		return nil, false
	}
	return o.Vendorname, true
}

// HasVendorname returns a boolean if a field has been set.
func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) HasVendorname() bool {
	if o != nil && !IsNil(o.Vendorname) {
		return true
	}

	return false
}

// SetVendorname gets a reference to the given string and assigns it to the Vendorname field.
func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) SetVendorname(v string) {
	o.Vendorname = &v
}

// GetCpucode returns the Cpucode field value if set, zero value otherwise.
func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) GetCpucode() string {
	if o == nil || IsNil(o.Cpucode) {
		var ret string
		return ret
	}
	return *o.Cpucode
}

// GetCpucodeOk returns a tuple with the Cpucode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) GetCpucodeOk() (*string, bool) {
	if o == nil || IsNil(o.Cpucode) {
		return nil, false
	}
	return o.Cpucode, true
}

// HasCpucode returns a boolean if a field has been set.
func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) HasCpucode() bool {
	if o != nil && !IsNil(o.Cpucode) {
		return true
	}

	return false
}

// SetCpucode gets a reference to the given string and assigns it to the Cpucode field.
func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) SetCpucode(v string) {
	o.Cpucode = &v
}

// GetClass returns the Class field value if set, zero value otherwise.
func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) GetClass() string {
	if o == nil || IsNil(o.Class) {
		var ret string
		return ret
	}
	return *o.Class
}

// GetClassOk returns a tuple with the Class field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) GetClassOk() (*string, bool) {
	if o == nil || IsNil(o.Class) {
		return nil, false
	}
	return o.Class, true
}

// HasClass returns a boolean if a field has been set.
func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) HasClass() bool {
	if o != nil && !IsNil(o.Class) {
		return true
	}

	return false
}

// SetClass gets a reference to the given string and assigns it to the Class field.
func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) SetClass(v string) {
	o.Class = &v
}

// GetSkustatus returns the Skustatus field value if set, zero value otherwise.
func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) GetSkustatus() string {
	if o == nil || IsNil(o.Skustatus) {
		var ret string
		return ret
	}
	return *o.Skustatus
}

// GetSkustatusOk returns a tuple with the Skustatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) GetSkustatusOk() (*string, bool) {
	if o == nil || IsNil(o.Skustatus) {
		return nil, false
	}
	return o.Skustatus, true
}

// HasSkustatus returns a boolean if a field has been set.
func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) HasSkustatus() bool {
	if o != nil && !IsNil(o.Skustatus) {
		return true
	}

	return false
}

// SetSkustatus gets a reference to the given string and assigns it to the Skustatus field.
func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) SetSkustatus(v string) {
	o.Skustatus = &v
}

// GetMediacpu returns the Mediacpu field value if set, zero value otherwise.
func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) GetMediacpu() string {
	if o == nil || IsNil(o.Mediacpu) {
		var ret string
		return ret
	}
	return *o.Mediacpu
}

// GetMediacpuOk returns a tuple with the Mediacpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) GetMediacpuOk() (*string, bool) {
	if o == nil || IsNil(o.Mediacpu) {
		return nil, false
	}
	return o.Mediacpu, true
}

// HasMediacpu returns a boolean if a field has been set.
func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) HasMediacpu() bool {
	if o != nil && !IsNil(o.Mediacpu) {
		return true
	}

	return false
}

// SetMediacpu gets a reference to the given string and assigns it to the Mediacpu field.
func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) SetMediacpu(v string) {
	o.Mediacpu = &v
}

// GetCategorysubcategory returns the Categorysubcategory field value if set, zero value otherwise.
func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) GetCategorysubcategory() string {
	if o == nil || IsNil(o.Categorysubcategory) {
		var ret string
		return ret
	}
	return *o.Categorysubcategory
}

// GetCategorysubcategoryOk returns a tuple with the Categorysubcategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) GetCategorysubcategoryOk() (*string, bool) {
	if o == nil || IsNil(o.Categorysubcategory) {
		return nil, false
	}
	return o.Categorysubcategory, true
}

// HasCategorysubcategory returns a boolean if a field has been set.
func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) HasCategorysubcategory() bool {
	if o != nil && !IsNil(o.Categorysubcategory) {
		return true
	}

	return false
}

// SetCategorysubcategory gets a reference to the given string and assigns it to the Categorysubcategory field.
func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) SetCategorysubcategory(v string) {
	o.Categorysubcategory = &v
}

// GetRetailprice returns the Retailprice field value if set, zero value otherwise.
func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) GetRetailprice() float32 {
	if o == nil || IsNil(o.Retailprice) {
		var ret float32
		return ret
	}
	return *o.Retailprice
}

// GetRetailpriceOk returns a tuple with the Retailprice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) GetRetailpriceOk() (*float32, bool) {
	if o == nil || IsNil(o.Retailprice) {
		return nil, false
	}
	return o.Retailprice, true
}

// HasRetailprice returns a boolean if a field has been set.
func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) HasRetailprice() bool {
	if o != nil && !IsNil(o.Retailprice) {
		return true
	}

	return false
}

// SetRetailprice gets a reference to the given float32 and assigns it to the Retailprice field.
func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) SetRetailprice(v float32) {
	o.Retailprice = &v
}

// GetNewmedia returns the Newmedia field value if set, zero value otherwise.
func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) GetNewmedia() string {
	if o == nil || IsNil(o.Newmedia) {
		var ret string
		return ret
	}
	return *o.Newmedia
}

// GetNewmediaOk returns a tuple with the Newmedia field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) GetNewmediaOk() (*string, bool) {
	if o == nil || IsNil(o.Newmedia) {
		return nil, false
	}
	return o.Newmedia, true
}

// HasNewmedia returns a boolean if a field has been set.
func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) HasNewmedia() bool {
	if o != nil && !IsNil(o.Newmedia) {
		return true
	}

	return false
}

// SetNewmedia gets a reference to the given string and assigns it to the Newmedia field.
func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) SetNewmedia(v string) {
	o.Newmedia = &v
}

// GetEnduserrequired returns the Enduserrequired field value if set, zero value otherwise.
func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) GetEnduserrequired() string {
	if o == nil || IsNil(o.Enduserrequired) {
		var ret string
		return ret
	}
	return *o.Enduserrequired
}

// GetEnduserrequiredOk returns a tuple with the Enduserrequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) GetEnduserrequiredOk() (*string, bool) {
	if o == nil || IsNil(o.Enduserrequired) {
		return nil, false
	}
	return o.Enduserrequired, true
}

// HasEnduserrequired returns a boolean if a field has been set.
func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) HasEnduserrequired() bool {
	if o != nil && !IsNil(o.Enduserrequired) {
		return true
	}

	return false
}

// SetEnduserrequired gets a reference to the given string and assigns it to the Enduserrequired field.
func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) SetEnduserrequired(v string) {
	o.Enduserrequired = &v
}

// GetBackorderflag returns the Backorderflag field value if set, zero value otherwise.
func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) GetBackorderflag() string {
	if o == nil || IsNil(o.Backorderflag) {
		var ret string
		return ret
	}
	return *o.Backorderflag
}

// GetBackorderflagOk returns a tuple with the Backorderflag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) GetBackorderflagOk() (*string, bool) {
	if o == nil || IsNil(o.Backorderflag) {
		return nil, false
	}
	return o.Backorderflag, true
}

// HasBackorderflag returns a boolean if a field has been set.
func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) HasBackorderflag() bool {
	if o != nil && !IsNil(o.Backorderflag) {
		return true
	}

	return false
}

// SetBackorderflag gets a reference to the given string and assigns it to the Backorderflag field.
func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) SetBackorderflag(v string) {
	o.Backorderflag = &v
}

// GetSkuauthorized returns the Skuauthorized field value if set, zero value otherwise.
func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) GetSkuauthorized() string {
	if o == nil || IsNil(o.Skuauthorized) {
		var ret string
		return ret
	}
	return *o.Skuauthorized
}

// GetSkuauthorizedOk returns a tuple with the Skuauthorized field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) GetSkuauthorizedOk() (*string, bool) {
	if o == nil || IsNil(o.Skuauthorized) {
		return nil, false
	}
	return o.Skuauthorized, true
}

// HasSkuauthorized returns a boolean if a field has been set.
func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) HasSkuauthorized() bool {
	if o != nil && !IsNil(o.Skuauthorized) {
		return true
	}

	return false
}

// SetSkuauthorized gets a reference to the given string and assigns it to the Skuauthorized field.
func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) SetSkuauthorized(v string) {
	o.Skuauthorized = &v
}

// GetExtendedvendorpartnumber returns the Extendedvendorpartnumber field value if set, zero value otherwise.
func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) GetExtendedvendorpartnumber() string {
	if o == nil || IsNil(o.Extendedvendorpartnumber) {
		var ret string
		return ret
	}
	return *o.Extendedvendorpartnumber
}

// GetExtendedvendorpartnumberOk returns a tuple with the Extendedvendorpartnumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) GetExtendedvendorpartnumberOk() (*string, bool) {
	if o == nil || IsNil(o.Extendedvendorpartnumber) {
		return nil, false
	}
	return o.Extendedvendorpartnumber, true
}

// HasExtendedvendorpartnumber returns a boolean if a field has been set.
func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) HasExtendedvendorpartnumber() bool {
	if o != nil && !IsNil(o.Extendedvendorpartnumber) {
		return true
	}

	return false
}

// SetExtendedvendorpartnumber gets a reference to the given string and assigns it to the Extendedvendorpartnumber field.
func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) SetExtendedvendorpartnumber(v string) {
	o.Extendedvendorpartnumber = &v
}

// GetWarehousedetails returns the Warehousedetails field value if set, zero value otherwise.
func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) GetWarehousedetails() []WarehouseListType {
	if o == nil || IsNil(o.Warehousedetails) {
		var ret []WarehouseListType
		return ret
	}
	return o.Warehousedetails
}

// GetWarehousedetailsOk returns a tuple with the Warehousedetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) GetWarehousedetailsOk() ([]WarehouseListType, bool) {
	if o == nil || IsNil(o.Warehousedetails) {
		return nil, false
	}
	return o.Warehousedetails, true
}

// HasWarehousedetails returns a boolean if a field has been set.
func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) HasWarehousedetails() bool {
	if o != nil && !IsNil(o.Warehousedetails) {
		return true
	}

	return false
}

// SetWarehousedetails gets a reference to the given []WarehouseListType and assigns it to the Warehousedetails field.
func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) SetWarehousedetails(v []WarehouseListType) {
	o.Warehousedetails = v
}

func (o PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Itemstatus) {
		toSerialize["itemstatus"] = o.Itemstatus
	}
	if !IsNil(o.Statusmessage) {
		toSerialize["statusmessage"] = o.Statusmessage
	}
	if !IsNil(o.Ingrampartnumber) {
		toSerialize["ingrampartnumber"] = o.Ingrampartnumber
	}
	if !IsNil(o.Vendorpartnumber) {
		toSerialize["vendorpartnumber"] = o.Vendorpartnumber
	}
	if !IsNil(o.Globalskuid) {
		toSerialize["globalskuid"] = o.Globalskuid
	}
	if !IsNil(o.Customerprice) {
		toSerialize["customerprice"] = o.Customerprice
	}
	if !IsNil(o.Partdescription1) {
		toSerialize["partdescription1"] = o.Partdescription1
	}
	if !IsNil(o.Partdescription2) {
		toSerialize["partdescription2"] = o.Partdescription2
	}
	if !IsNil(o.Vendornumber) {
		toSerialize["vendornumber"] = o.Vendornumber
	}
	if !IsNil(o.Vendorname) {
		toSerialize["vendorname"] = o.Vendorname
	}
	if !IsNil(o.Cpucode) {
		toSerialize["cpucode"] = o.Cpucode
	}
	if !IsNil(o.Class) {
		toSerialize["class"] = o.Class
	}
	if !IsNil(o.Skustatus) {
		toSerialize["skustatus"] = o.Skustatus
	}
	if !IsNil(o.Mediacpu) {
		toSerialize["mediacpu"] = o.Mediacpu
	}
	if !IsNil(o.Categorysubcategory) {
		toSerialize["categorysubcategory"] = o.Categorysubcategory
	}
	if !IsNil(o.Retailprice) {
		toSerialize["retailprice"] = o.Retailprice
	}
	if !IsNil(o.Newmedia) {
		toSerialize["newmedia"] = o.Newmedia
	}
	if !IsNil(o.Enduserrequired) {
		toSerialize["enduserrequired"] = o.Enduserrequired
	}
	if !IsNil(o.Backorderflag) {
		toSerialize["backorderflag"] = o.Backorderflag
	}
	if !IsNil(o.Skuauthorized) {
		toSerialize["skuauthorized"] = o.Skuauthorized
	}
	if !IsNil(o.Extendedvendorpartnumber) {
		toSerialize["extendedvendorpartnumber"] = o.Extendedvendorpartnumber
	}
	if !IsNil(o.Warehousedetails) {
		toSerialize["warehousedetails"] = o.Warehousedetails
	}
	return toSerialize, nil
}

type NullablePriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner struct {
	value *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner
	isSet bool
}

func (v NullablePriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) Get() *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner {
	return v.value
}

func (v *NullablePriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) Set(val *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner(val *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) *NullablePriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner {
	return &NullablePriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner{value: val, isSet: true}
}

func (v NullablePriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


