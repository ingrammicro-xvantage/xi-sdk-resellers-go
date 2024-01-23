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

// checks if the MultiSKUPriceAndStockRequestServicerequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MultiSKUPriceAndStockRequestServicerequest{}

// MultiSKUPriceAndStockRequestServicerequest struct for MultiSKUPriceAndStockRequestServicerequest
type MultiSKUPriceAndStockRequestServicerequest struct {
	Requestpreamble *MultiSKUPriceAndStockRequestServicerequestRequestpreamble `json:"requestpreamble,omitempty"`
	Priceandstockrequest *MultiSKUPriceAndStockRequestServicerequestPriceandstockrequest `json:"priceandstockrequest,omitempty"`
}

// NewMultiSKUPriceAndStockRequestServicerequest instantiates a new MultiSKUPriceAndStockRequestServicerequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultiSKUPriceAndStockRequestServicerequest() *MultiSKUPriceAndStockRequestServicerequest {
	this := MultiSKUPriceAndStockRequestServicerequest{}
	return &this
}

// NewMultiSKUPriceAndStockRequestServicerequestWithDefaults instantiates a new MultiSKUPriceAndStockRequestServicerequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultiSKUPriceAndStockRequestServicerequestWithDefaults() *MultiSKUPriceAndStockRequestServicerequest {
	this := MultiSKUPriceAndStockRequestServicerequest{}
	return &this
}

// GetRequestpreamble returns the Requestpreamble field value if set, zero value otherwise.
func (o *MultiSKUPriceAndStockRequestServicerequest) GetRequestpreamble() MultiSKUPriceAndStockRequestServicerequestRequestpreamble {
	if o == nil || IsNil(o.Requestpreamble) {
		var ret MultiSKUPriceAndStockRequestServicerequestRequestpreamble
		return ret
	}
	return *o.Requestpreamble
}

// GetRequestpreambleOk returns a tuple with the Requestpreamble field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiSKUPriceAndStockRequestServicerequest) GetRequestpreambleOk() (*MultiSKUPriceAndStockRequestServicerequestRequestpreamble, bool) {
	if o == nil || IsNil(o.Requestpreamble) {
		return nil, false
	}
	return o.Requestpreamble, true
}

// HasRequestpreamble returns a boolean if a field has been set.
func (o *MultiSKUPriceAndStockRequestServicerequest) HasRequestpreamble() bool {
	if o != nil && !IsNil(o.Requestpreamble) {
		return true
	}

	return false
}

// SetRequestpreamble gets a reference to the given MultiSKUPriceAndStockRequestServicerequestRequestpreamble and assigns it to the Requestpreamble field.
func (o *MultiSKUPriceAndStockRequestServicerequest) SetRequestpreamble(v MultiSKUPriceAndStockRequestServicerequestRequestpreamble) {
	o.Requestpreamble = &v
}

// GetPriceandstockrequest returns the Priceandstockrequest field value if set, zero value otherwise.
func (o *MultiSKUPriceAndStockRequestServicerequest) GetPriceandstockrequest() MultiSKUPriceAndStockRequestServicerequestPriceandstockrequest {
	if o == nil || IsNil(o.Priceandstockrequest) {
		var ret MultiSKUPriceAndStockRequestServicerequestPriceandstockrequest
		return ret
	}
	return *o.Priceandstockrequest
}

// GetPriceandstockrequestOk returns a tuple with the Priceandstockrequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiSKUPriceAndStockRequestServicerequest) GetPriceandstockrequestOk() (*MultiSKUPriceAndStockRequestServicerequestPriceandstockrequest, bool) {
	if o == nil || IsNil(o.Priceandstockrequest) {
		return nil, false
	}
	return o.Priceandstockrequest, true
}

// HasPriceandstockrequest returns a boolean if a field has been set.
func (o *MultiSKUPriceAndStockRequestServicerequest) HasPriceandstockrequest() bool {
	if o != nil && !IsNil(o.Priceandstockrequest) {
		return true
	}

	return false
}

// SetPriceandstockrequest gets a reference to the given MultiSKUPriceAndStockRequestServicerequestPriceandstockrequest and assigns it to the Priceandstockrequest field.
func (o *MultiSKUPriceAndStockRequestServicerequest) SetPriceandstockrequest(v MultiSKUPriceAndStockRequestServicerequestPriceandstockrequest) {
	o.Priceandstockrequest = &v
}

func (o MultiSKUPriceAndStockRequestServicerequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MultiSKUPriceAndStockRequestServicerequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Requestpreamble) {
		toSerialize["requestpreamble"] = o.Requestpreamble
	}
	if !IsNil(o.Priceandstockrequest) {
		toSerialize["priceandstockrequest"] = o.Priceandstockrequest
	}
	return toSerialize, nil
}

type NullableMultiSKUPriceAndStockRequestServicerequest struct {
	value *MultiSKUPriceAndStockRequestServicerequest
	isSet bool
}

func (v NullableMultiSKUPriceAndStockRequestServicerequest) Get() *MultiSKUPriceAndStockRequestServicerequest {
	return v.value
}

func (v *NullableMultiSKUPriceAndStockRequestServicerequest) Set(val *MultiSKUPriceAndStockRequestServicerequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMultiSKUPriceAndStockRequestServicerequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMultiSKUPriceAndStockRequestServicerequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultiSKUPriceAndStockRequestServicerequest(val *MultiSKUPriceAndStockRequestServicerequest) *NullableMultiSKUPriceAndStockRequestServicerequest {
	return &NullableMultiSKUPriceAndStockRequestServicerequest{value: val, isSet: true}
}

func (v NullableMultiSKUPriceAndStockRequestServicerequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultiSKUPriceAndStockRequestServicerequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


