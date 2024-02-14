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

// checks if the InvoiceDetailRequestServicerequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvoiceDetailRequestServicerequest{}

// InvoiceDetailRequestServicerequest struct for InvoiceDetailRequestServicerequest
type InvoiceDetailRequestServicerequest struct {
	Requestpreamble *InvoiceDetailRequestServicerequestRequestpreamble `json:"requestpreamble,omitempty"`
	Invoicedetailrequest *InvoiceDetailRequestServicerequestInvoicedetailrequest `json:"invoicedetailrequest,omitempty"`
}

// NewInvoiceDetailRequestServicerequest instantiates a new InvoiceDetailRequestServicerequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoiceDetailRequestServicerequest() *InvoiceDetailRequestServicerequest {
	this := InvoiceDetailRequestServicerequest{}
	return &this
}

// NewInvoiceDetailRequestServicerequestWithDefaults instantiates a new InvoiceDetailRequestServicerequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoiceDetailRequestServicerequestWithDefaults() *InvoiceDetailRequestServicerequest {
	this := InvoiceDetailRequestServicerequest{}
	return &this
}

// GetRequestpreamble returns the Requestpreamble field value if set, zero value otherwise.
func (o *InvoiceDetailRequestServicerequest) GetRequestpreamble() InvoiceDetailRequestServicerequestRequestpreamble {
	if o == nil || IsNil(o.Requestpreamble) {
		var ret InvoiceDetailRequestServicerequestRequestpreamble
		return ret
	}
	return *o.Requestpreamble
}

// GetRequestpreambleOk returns a tuple with the Requestpreamble field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceDetailRequestServicerequest) GetRequestpreambleOk() (*InvoiceDetailRequestServicerequestRequestpreamble, bool) {
	if o == nil || IsNil(o.Requestpreamble) {
		return nil, false
	}
	return o.Requestpreamble, true
}

// HasRequestpreamble returns a boolean if a field has been set.
func (o *InvoiceDetailRequestServicerequest) HasRequestpreamble() bool {
	if o != nil && !IsNil(o.Requestpreamble) {
		return true
	}

	return false
}

// SetRequestpreamble gets a reference to the given InvoiceDetailRequestServicerequestRequestpreamble and assigns it to the Requestpreamble field.
func (o *InvoiceDetailRequestServicerequest) SetRequestpreamble(v InvoiceDetailRequestServicerequestRequestpreamble) {
	o.Requestpreamble = &v
}

// GetInvoicedetailrequest returns the Invoicedetailrequest field value if set, zero value otherwise.
func (o *InvoiceDetailRequestServicerequest) GetInvoicedetailrequest() InvoiceDetailRequestServicerequestInvoicedetailrequest {
	if o == nil || IsNil(o.Invoicedetailrequest) {
		var ret InvoiceDetailRequestServicerequestInvoicedetailrequest
		return ret
	}
	return *o.Invoicedetailrequest
}

// GetInvoicedetailrequestOk returns a tuple with the Invoicedetailrequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceDetailRequestServicerequest) GetInvoicedetailrequestOk() (*InvoiceDetailRequestServicerequestInvoicedetailrequest, bool) {
	if o == nil || IsNil(o.Invoicedetailrequest) {
		return nil, false
	}
	return o.Invoicedetailrequest, true
}

// HasInvoicedetailrequest returns a boolean if a field has been set.
func (o *InvoiceDetailRequestServicerequest) HasInvoicedetailrequest() bool {
	if o != nil && !IsNil(o.Invoicedetailrequest) {
		return true
	}

	return false
}

// SetInvoicedetailrequest gets a reference to the given InvoiceDetailRequestServicerequestInvoicedetailrequest and assigns it to the Invoicedetailrequest field.
func (o *InvoiceDetailRequestServicerequest) SetInvoicedetailrequest(v InvoiceDetailRequestServicerequestInvoicedetailrequest) {
	o.Invoicedetailrequest = &v
}

func (o InvoiceDetailRequestServicerequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvoiceDetailRequestServicerequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Requestpreamble) {
		toSerialize["requestpreamble"] = o.Requestpreamble
	}
	if !IsNil(o.Invoicedetailrequest) {
		toSerialize["invoicedetailrequest"] = o.Invoicedetailrequest
	}
	return toSerialize, nil
}

type NullableInvoiceDetailRequestServicerequest struct {
	value *InvoiceDetailRequestServicerequest
	isSet bool
}

func (v NullableInvoiceDetailRequestServicerequest) Get() *InvoiceDetailRequestServicerequest {
	return v.value
}

func (v *NullableInvoiceDetailRequestServicerequest) Set(val *InvoiceDetailRequestServicerequest) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoiceDetailRequestServicerequest) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoiceDetailRequestServicerequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoiceDetailRequestServicerequest(val *InvoiceDetailRequestServicerequest) *NullableInvoiceDetailRequestServicerequest {
	return &NullableInvoiceDetailRequestServicerequest{value: val, isSet: true}
}

func (v NullableInvoiceDetailRequestServicerequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoiceDetailRequestServicerequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


