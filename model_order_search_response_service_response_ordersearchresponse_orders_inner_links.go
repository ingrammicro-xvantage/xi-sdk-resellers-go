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

// checks if the OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerLinks{}

// OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerLinks HATEOAS links for the main order
type OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerLinks struct {
	// Topic being orders in this case, if it is orders then the link will provide details of the order.
	Topic *string `json:"topic,omitempty"`
	// The API endpoint for accessing the relevant data
	Href *string `json:"href,omitempty"`
	// The type of call that can be made to the href link
	Type *string `json:"type,omitempty"`
}

// NewOrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerLinks instantiates a new OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerLinks() *OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerLinks {
	this := OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerLinks{}
	return &this
}

// NewOrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerLinksWithDefaults instantiates a new OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerLinksWithDefaults() *OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerLinks {
	this := OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerLinks{}
	return &this
}

// GetTopic returns the Topic field value if set, zero value otherwise.
func (o *OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerLinks) GetTopic() string {
	if o == nil || IsNil(o.Topic) {
		var ret string
		return ret
	}
	return *o.Topic
}

// GetTopicOk returns a tuple with the Topic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerLinks) GetTopicOk() (*string, bool) {
	if o == nil || IsNil(o.Topic) {
		return nil, false
	}
	return o.Topic, true
}

// HasTopic returns a boolean if a field has been set.
func (o *OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerLinks) HasTopic() bool {
	if o != nil && !IsNil(o.Topic) {
		return true
	}

	return false
}

// SetTopic gets a reference to the given string and assigns it to the Topic field.
func (o *OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerLinks) SetTopic(v string) {
	o.Topic = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerLinks) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerLinks) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerLinks) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerLinks) SetHref(v string) {
	o.Href = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerLinks) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerLinks) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerLinks) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerLinks) SetType(v string) {
	o.Type = &v
}

func (o OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerLinks) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Topic) {
		toSerialize["topic"] = o.Topic
	}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableOrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerLinks struct {
	value *OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerLinks
	isSet bool
}

func (v NullableOrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerLinks) Get() *OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerLinks {
	return v.value
}

func (v *NullableOrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerLinks) Set(val *OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerLinks(val *OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerLinks) *NullableOrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerLinks {
	return &NullableOrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerLinks{value: val, isSet: true}
}

func (v NullableOrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


