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

// checks if the AddressType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddressType{}

// AddressType Address type object
type AddressType struct {
	Attention *string `json:"attention,omitempty"`
	Name1 *string `json:"name1,omitempty"`
	Name2 *string `json:"name2,omitempty"`
	Addressline1 *string `json:"addressline1,omitempty"`
	Addressline2 *string `json:"addressline2,omitempty"`
	Addressline3 *string `json:"addressline3,omitempty"`
	City *string `json:"city,omitempty"`
	State *string `json:"state,omitempty"`
	Postalcode *string `json:"postalcode,omitempty"`
	Countrycode *string `json:"countrycode,omitempty"`
	Fax *string `json:"fax,omitempty"`
	Phonenumber *string `json:"phonenumber,omitempty"`
	Email *string `json:"email,omitempty"`
}

// NewAddressType instantiates a new AddressType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressType() *AddressType {
	this := AddressType{}
	return &this
}

// NewAddressTypeWithDefaults instantiates a new AddressType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressTypeWithDefaults() *AddressType {
	this := AddressType{}
	return &this
}

// GetAttention returns the Attention field value if set, zero value otherwise.
func (o *AddressType) GetAttention() string {
	if o == nil || IsNil(o.Attention) {
		var ret string
		return ret
	}
	return *o.Attention
}

// GetAttentionOk returns a tuple with the Attention field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressType) GetAttentionOk() (*string, bool) {
	if o == nil || IsNil(o.Attention) {
		return nil, false
	}
	return o.Attention, true
}

// HasAttention returns a boolean if a field has been set.
func (o *AddressType) HasAttention() bool {
	if o != nil && !IsNil(o.Attention) {
		return true
	}

	return false
}

// SetAttention gets a reference to the given string and assigns it to the Attention field.
func (o *AddressType) SetAttention(v string) {
	o.Attention = &v
}

// GetName1 returns the Name1 field value if set, zero value otherwise.
func (o *AddressType) GetName1() string {
	if o == nil || IsNil(o.Name1) {
		var ret string
		return ret
	}
	return *o.Name1
}

// GetName1Ok returns a tuple with the Name1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressType) GetName1Ok() (*string, bool) {
	if o == nil || IsNil(o.Name1) {
		return nil, false
	}
	return o.Name1, true
}

// HasName1 returns a boolean if a field has been set.
func (o *AddressType) HasName1() bool {
	if o != nil && !IsNil(o.Name1) {
		return true
	}

	return false
}

// SetName1 gets a reference to the given string and assigns it to the Name1 field.
func (o *AddressType) SetName1(v string) {
	o.Name1 = &v
}

// GetName2 returns the Name2 field value if set, zero value otherwise.
func (o *AddressType) GetName2() string {
	if o == nil || IsNil(o.Name2) {
		var ret string
		return ret
	}
	return *o.Name2
}

// GetName2Ok returns a tuple with the Name2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressType) GetName2Ok() (*string, bool) {
	if o == nil || IsNil(o.Name2) {
		return nil, false
	}
	return o.Name2, true
}

// HasName2 returns a boolean if a field has been set.
func (o *AddressType) HasName2() bool {
	if o != nil && !IsNil(o.Name2) {
		return true
	}

	return false
}

// SetName2 gets a reference to the given string and assigns it to the Name2 field.
func (o *AddressType) SetName2(v string) {
	o.Name2 = &v
}

// GetAddressline1 returns the Addressline1 field value if set, zero value otherwise.
func (o *AddressType) GetAddressline1() string {
	if o == nil || IsNil(o.Addressline1) {
		var ret string
		return ret
	}
	return *o.Addressline1
}

// GetAddressline1Ok returns a tuple with the Addressline1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressType) GetAddressline1Ok() (*string, bool) {
	if o == nil || IsNil(o.Addressline1) {
		return nil, false
	}
	return o.Addressline1, true
}

// HasAddressline1 returns a boolean if a field has been set.
func (o *AddressType) HasAddressline1() bool {
	if o != nil && !IsNil(o.Addressline1) {
		return true
	}

	return false
}

// SetAddressline1 gets a reference to the given string and assigns it to the Addressline1 field.
func (o *AddressType) SetAddressline1(v string) {
	o.Addressline1 = &v
}

// GetAddressline2 returns the Addressline2 field value if set, zero value otherwise.
func (o *AddressType) GetAddressline2() string {
	if o == nil || IsNil(o.Addressline2) {
		var ret string
		return ret
	}
	return *o.Addressline2
}

// GetAddressline2Ok returns a tuple with the Addressline2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressType) GetAddressline2Ok() (*string, bool) {
	if o == nil || IsNil(o.Addressline2) {
		return nil, false
	}
	return o.Addressline2, true
}

// HasAddressline2 returns a boolean if a field has been set.
func (o *AddressType) HasAddressline2() bool {
	if o != nil && !IsNil(o.Addressline2) {
		return true
	}

	return false
}

// SetAddressline2 gets a reference to the given string and assigns it to the Addressline2 field.
func (o *AddressType) SetAddressline2(v string) {
	o.Addressline2 = &v
}

// GetAddressline3 returns the Addressline3 field value if set, zero value otherwise.
func (o *AddressType) GetAddressline3() string {
	if o == nil || IsNil(o.Addressline3) {
		var ret string
		return ret
	}
	return *o.Addressline3
}

// GetAddressline3Ok returns a tuple with the Addressline3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressType) GetAddressline3Ok() (*string, bool) {
	if o == nil || IsNil(o.Addressline3) {
		return nil, false
	}
	return o.Addressline3, true
}

// HasAddressline3 returns a boolean if a field has been set.
func (o *AddressType) HasAddressline3() bool {
	if o != nil && !IsNil(o.Addressline3) {
		return true
	}

	return false
}

// SetAddressline3 gets a reference to the given string and assigns it to the Addressline3 field.
func (o *AddressType) SetAddressline3(v string) {
	o.Addressline3 = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *AddressType) GetCity() string {
	if o == nil || IsNil(o.City) {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressType) GetCityOk() (*string, bool) {
	if o == nil || IsNil(o.City) {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *AddressType) HasCity() bool {
	if o != nil && !IsNil(o.City) {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *AddressType) SetCity(v string) {
	o.City = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *AddressType) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressType) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *AddressType) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *AddressType) SetState(v string) {
	o.State = &v
}

// GetPostalcode returns the Postalcode field value if set, zero value otherwise.
func (o *AddressType) GetPostalcode() string {
	if o == nil || IsNil(o.Postalcode) {
		var ret string
		return ret
	}
	return *o.Postalcode
}

// GetPostalcodeOk returns a tuple with the Postalcode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressType) GetPostalcodeOk() (*string, bool) {
	if o == nil || IsNil(o.Postalcode) {
		return nil, false
	}
	return o.Postalcode, true
}

// HasPostalcode returns a boolean if a field has been set.
func (o *AddressType) HasPostalcode() bool {
	if o != nil && !IsNil(o.Postalcode) {
		return true
	}

	return false
}

// SetPostalcode gets a reference to the given string and assigns it to the Postalcode field.
func (o *AddressType) SetPostalcode(v string) {
	o.Postalcode = &v
}

// GetCountrycode returns the Countrycode field value if set, zero value otherwise.
func (o *AddressType) GetCountrycode() string {
	if o == nil || IsNil(o.Countrycode) {
		var ret string
		return ret
	}
	return *o.Countrycode
}

// GetCountrycodeOk returns a tuple with the Countrycode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressType) GetCountrycodeOk() (*string, bool) {
	if o == nil || IsNil(o.Countrycode) {
		return nil, false
	}
	return o.Countrycode, true
}

// HasCountrycode returns a boolean if a field has been set.
func (o *AddressType) HasCountrycode() bool {
	if o != nil && !IsNil(o.Countrycode) {
		return true
	}

	return false
}

// SetCountrycode gets a reference to the given string and assigns it to the Countrycode field.
func (o *AddressType) SetCountrycode(v string) {
	o.Countrycode = &v
}

// GetFax returns the Fax field value if set, zero value otherwise.
func (o *AddressType) GetFax() string {
	if o == nil || IsNil(o.Fax) {
		var ret string
		return ret
	}
	return *o.Fax
}

// GetFaxOk returns a tuple with the Fax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressType) GetFaxOk() (*string, bool) {
	if o == nil || IsNil(o.Fax) {
		return nil, false
	}
	return o.Fax, true
}

// HasFax returns a boolean if a field has been set.
func (o *AddressType) HasFax() bool {
	if o != nil && !IsNil(o.Fax) {
		return true
	}

	return false
}

// SetFax gets a reference to the given string and assigns it to the Fax field.
func (o *AddressType) SetFax(v string) {
	o.Fax = &v
}

// GetPhonenumber returns the Phonenumber field value if set, zero value otherwise.
func (o *AddressType) GetPhonenumber() string {
	if o == nil || IsNil(o.Phonenumber) {
		var ret string
		return ret
	}
	return *o.Phonenumber
}

// GetPhonenumberOk returns a tuple with the Phonenumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressType) GetPhonenumberOk() (*string, bool) {
	if o == nil || IsNil(o.Phonenumber) {
		return nil, false
	}
	return o.Phonenumber, true
}

// HasPhonenumber returns a boolean if a field has been set.
func (o *AddressType) HasPhonenumber() bool {
	if o != nil && !IsNil(o.Phonenumber) {
		return true
	}

	return false
}

// SetPhonenumber gets a reference to the given string and assigns it to the Phonenumber field.
func (o *AddressType) SetPhonenumber(v string) {
	o.Phonenumber = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *AddressType) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressType) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *AddressType) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *AddressType) SetEmail(v string) {
	o.Email = &v
}

func (o AddressType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddressType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Attention) {
		toSerialize["attention"] = o.Attention
	}
	if !IsNil(o.Name1) {
		toSerialize["name1"] = o.Name1
	}
	if !IsNil(o.Name2) {
		toSerialize["name2"] = o.Name2
	}
	if !IsNil(o.Addressline1) {
		toSerialize["addressline1"] = o.Addressline1
	}
	if !IsNil(o.Addressline2) {
		toSerialize["addressline2"] = o.Addressline2
	}
	if !IsNil(o.Addressline3) {
		toSerialize["addressline3"] = o.Addressline3
	}
	if !IsNil(o.City) {
		toSerialize["city"] = o.City
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.Postalcode) {
		toSerialize["postalcode"] = o.Postalcode
	}
	if !IsNil(o.Countrycode) {
		toSerialize["countrycode"] = o.Countrycode
	}
	if !IsNil(o.Fax) {
		toSerialize["fax"] = o.Fax
	}
	if !IsNil(o.Phonenumber) {
		toSerialize["phonenumber"] = o.Phonenumber
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	return toSerialize, nil
}

type NullableAddressType struct {
	value *AddressType
	isSet bool
}

func (v NullableAddressType) Get() *AddressType {
	return v.value
}

func (v *NullableAddressType) Set(val *AddressType) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressType) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressType(val *AddressType) *NullableAddressType {
	return &NullableAddressType{value: val, isSet: true}
}

func (v NullableAddressType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


