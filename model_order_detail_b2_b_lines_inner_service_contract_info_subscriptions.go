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

// checks if the OrderDetailB2BLinesInnerServiceContractInfoSubscriptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderDetailB2BLinesInnerServiceContractInfoSubscriptions{}

// OrderDetailB2BLinesInnerServiceContractInfoSubscriptions struct for OrderDetailB2BLinesInnerServiceContractInfoSubscriptions
type OrderDetailB2BLinesInnerServiceContractInfoSubscriptions struct {
	// The ID of the subscription.
	SubscriptionId *string `json:"subscriptionId,omitempty"`
	// The term of the subscription.
	SubscriptionTerm *string `json:"subscriptionTerm,omitempty"`
	// The renewal term of the subscription.
	RenewalTerm *string `json:"renewalTerm,omitempty"`
	// The billing model of the billing.
	BillingModel *string `json:"billingModel,omitempty"`
	// Start date of the subcription.
	SubcriptionStartDate *string `json:"subcriptionStartDate,omitempty"`
	// End date of the subcription.
	SubcriptionEndDate *string `json:"subcriptionEndDate,omitempty"`
}

// NewOrderDetailB2BLinesInnerServiceContractInfoSubscriptions instantiates a new OrderDetailB2BLinesInnerServiceContractInfoSubscriptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderDetailB2BLinesInnerServiceContractInfoSubscriptions() *OrderDetailB2BLinesInnerServiceContractInfoSubscriptions {
	this := OrderDetailB2BLinesInnerServiceContractInfoSubscriptions{}
	return &this
}

// NewOrderDetailB2BLinesInnerServiceContractInfoSubscriptionsWithDefaults instantiates a new OrderDetailB2BLinesInnerServiceContractInfoSubscriptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderDetailB2BLinesInnerServiceContractInfoSubscriptionsWithDefaults() *OrderDetailB2BLinesInnerServiceContractInfoSubscriptions {
	this := OrderDetailB2BLinesInnerServiceContractInfoSubscriptions{}
	return &this
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise.
func (o *OrderDetailB2BLinesInnerServiceContractInfoSubscriptions) GetSubscriptionId() string {
	if o == nil || IsNil(o.SubscriptionId) {
		var ret string
		return ret
	}
	return *o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailB2BLinesInnerServiceContractInfoSubscriptions) GetSubscriptionIdOk() (*string, bool) {
	if o == nil || IsNil(o.SubscriptionId) {
		return nil, false
	}
	return o.SubscriptionId, true
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *OrderDetailB2BLinesInnerServiceContractInfoSubscriptions) HasSubscriptionId() bool {
	if o != nil && !IsNil(o.SubscriptionId) {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given string and assigns it to the SubscriptionId field.
func (o *OrderDetailB2BLinesInnerServiceContractInfoSubscriptions) SetSubscriptionId(v string) {
	o.SubscriptionId = &v
}

// GetSubscriptionTerm returns the SubscriptionTerm field value if set, zero value otherwise.
func (o *OrderDetailB2BLinesInnerServiceContractInfoSubscriptions) GetSubscriptionTerm() string {
	if o == nil || IsNil(o.SubscriptionTerm) {
		var ret string
		return ret
	}
	return *o.SubscriptionTerm
}

// GetSubscriptionTermOk returns a tuple with the SubscriptionTerm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailB2BLinesInnerServiceContractInfoSubscriptions) GetSubscriptionTermOk() (*string, bool) {
	if o == nil || IsNil(o.SubscriptionTerm) {
		return nil, false
	}
	return o.SubscriptionTerm, true
}

// HasSubscriptionTerm returns a boolean if a field has been set.
func (o *OrderDetailB2BLinesInnerServiceContractInfoSubscriptions) HasSubscriptionTerm() bool {
	if o != nil && !IsNil(o.SubscriptionTerm) {
		return true
	}

	return false
}

// SetSubscriptionTerm gets a reference to the given string and assigns it to the SubscriptionTerm field.
func (o *OrderDetailB2BLinesInnerServiceContractInfoSubscriptions) SetSubscriptionTerm(v string) {
	o.SubscriptionTerm = &v
}

// GetRenewalTerm returns the RenewalTerm field value if set, zero value otherwise.
func (o *OrderDetailB2BLinesInnerServiceContractInfoSubscriptions) GetRenewalTerm() string {
	if o == nil || IsNil(o.RenewalTerm) {
		var ret string
		return ret
	}
	return *o.RenewalTerm
}

// GetRenewalTermOk returns a tuple with the RenewalTerm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailB2BLinesInnerServiceContractInfoSubscriptions) GetRenewalTermOk() (*string, bool) {
	if o == nil || IsNil(o.RenewalTerm) {
		return nil, false
	}
	return o.RenewalTerm, true
}

// HasRenewalTerm returns a boolean if a field has been set.
func (o *OrderDetailB2BLinesInnerServiceContractInfoSubscriptions) HasRenewalTerm() bool {
	if o != nil && !IsNil(o.RenewalTerm) {
		return true
	}

	return false
}

// SetRenewalTerm gets a reference to the given string and assigns it to the RenewalTerm field.
func (o *OrderDetailB2BLinesInnerServiceContractInfoSubscriptions) SetRenewalTerm(v string) {
	o.RenewalTerm = &v
}

// GetBillingModel returns the BillingModel field value if set, zero value otherwise.
func (o *OrderDetailB2BLinesInnerServiceContractInfoSubscriptions) GetBillingModel() string {
	if o == nil || IsNil(o.BillingModel) {
		var ret string
		return ret
	}
	return *o.BillingModel
}

// GetBillingModelOk returns a tuple with the BillingModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailB2BLinesInnerServiceContractInfoSubscriptions) GetBillingModelOk() (*string, bool) {
	if o == nil || IsNil(o.BillingModel) {
		return nil, false
	}
	return o.BillingModel, true
}

// HasBillingModel returns a boolean if a field has been set.
func (o *OrderDetailB2BLinesInnerServiceContractInfoSubscriptions) HasBillingModel() bool {
	if o != nil && !IsNil(o.BillingModel) {
		return true
	}

	return false
}

// SetBillingModel gets a reference to the given string and assigns it to the BillingModel field.
func (o *OrderDetailB2BLinesInnerServiceContractInfoSubscriptions) SetBillingModel(v string) {
	o.BillingModel = &v
}

// GetSubcriptionStartDate returns the SubcriptionStartDate field value if set, zero value otherwise.
func (o *OrderDetailB2BLinesInnerServiceContractInfoSubscriptions) GetSubcriptionStartDate() string {
	if o == nil || IsNil(o.SubcriptionStartDate) {
		var ret string
		return ret
	}
	return *o.SubcriptionStartDate
}

// GetSubcriptionStartDateOk returns a tuple with the SubcriptionStartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailB2BLinesInnerServiceContractInfoSubscriptions) GetSubcriptionStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.SubcriptionStartDate) {
		return nil, false
	}
	return o.SubcriptionStartDate, true
}

// HasSubcriptionStartDate returns a boolean if a field has been set.
func (o *OrderDetailB2BLinesInnerServiceContractInfoSubscriptions) HasSubcriptionStartDate() bool {
	if o != nil && !IsNil(o.SubcriptionStartDate) {
		return true
	}

	return false
}

// SetSubcriptionStartDate gets a reference to the given string and assigns it to the SubcriptionStartDate field.
func (o *OrderDetailB2BLinesInnerServiceContractInfoSubscriptions) SetSubcriptionStartDate(v string) {
	o.SubcriptionStartDate = &v
}

// GetSubcriptionEndDate returns the SubcriptionEndDate field value if set, zero value otherwise.
func (o *OrderDetailB2BLinesInnerServiceContractInfoSubscriptions) GetSubcriptionEndDate() string {
	if o == nil || IsNil(o.SubcriptionEndDate) {
		var ret string
		return ret
	}
	return *o.SubcriptionEndDate
}

// GetSubcriptionEndDateOk returns a tuple with the SubcriptionEndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailB2BLinesInnerServiceContractInfoSubscriptions) GetSubcriptionEndDateOk() (*string, bool) {
	if o == nil || IsNil(o.SubcriptionEndDate) {
		return nil, false
	}
	return o.SubcriptionEndDate, true
}

// HasSubcriptionEndDate returns a boolean if a field has been set.
func (o *OrderDetailB2BLinesInnerServiceContractInfoSubscriptions) HasSubcriptionEndDate() bool {
	if o != nil && !IsNil(o.SubcriptionEndDate) {
		return true
	}

	return false
}

// SetSubcriptionEndDate gets a reference to the given string and assigns it to the SubcriptionEndDate field.
func (o *OrderDetailB2BLinesInnerServiceContractInfoSubscriptions) SetSubcriptionEndDate(v string) {
	o.SubcriptionEndDate = &v
}

func (o OrderDetailB2BLinesInnerServiceContractInfoSubscriptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderDetailB2BLinesInnerServiceContractInfoSubscriptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SubscriptionId) {
		toSerialize["subscriptionId"] = o.SubscriptionId
	}
	if !IsNil(o.SubscriptionTerm) {
		toSerialize["subscriptionTerm"] = o.SubscriptionTerm
	}
	if !IsNil(o.RenewalTerm) {
		toSerialize["renewalTerm"] = o.RenewalTerm
	}
	if !IsNil(o.BillingModel) {
		toSerialize["billingModel"] = o.BillingModel
	}
	if !IsNil(o.SubcriptionStartDate) {
		toSerialize["subcriptionStartDate"] = o.SubcriptionStartDate
	}
	if !IsNil(o.SubcriptionEndDate) {
		toSerialize["subcriptionEndDate"] = o.SubcriptionEndDate
	}
	return toSerialize, nil
}

type NullableOrderDetailB2BLinesInnerServiceContractInfoSubscriptions struct {
	value *OrderDetailB2BLinesInnerServiceContractInfoSubscriptions
	isSet bool
}

func (v NullableOrderDetailB2BLinesInnerServiceContractInfoSubscriptions) Get() *OrderDetailB2BLinesInnerServiceContractInfoSubscriptions {
	return v.value
}

func (v *NullableOrderDetailB2BLinesInnerServiceContractInfoSubscriptions) Set(val *OrderDetailB2BLinesInnerServiceContractInfoSubscriptions) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderDetailB2BLinesInnerServiceContractInfoSubscriptions) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderDetailB2BLinesInnerServiceContractInfoSubscriptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderDetailB2BLinesInnerServiceContractInfoSubscriptions(val *OrderDetailB2BLinesInnerServiceContractInfoSubscriptions) *NullableOrderDetailB2BLinesInnerServiceContractInfoSubscriptions {
	return &NullableOrderDetailB2BLinesInnerServiceContractInfoSubscriptions{value: val, isSet: true}
}

func (v NullableOrderDetailB2BLinesInnerServiceContractInfoSubscriptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderDetailB2BLinesInnerServiceContractInfoSubscriptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


