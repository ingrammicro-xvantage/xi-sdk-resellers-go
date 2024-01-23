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

// checks if the QuoteListResponseQuoteSearchResponseQuoteListInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuoteListResponseQuoteSearchResponseQuoteListInner{}

// QuoteListResponseQuoteSearchResponseQuoteListInner struct for QuoteListResponseQuoteSearchResponseQuoteListInner
type QuoteListResponseQuoteSearchResponseQuoteListInner struct {
	// Quote Name given to quote by sales team or system generated. Generally used as a reference to identify the quote.
	QuoteName *string `json:"quoteName,omitempty"`
	// Unique identifier generated by Ingram Micro's CRM specific to each quote. When applying a filter to the quoteNumber and including a partial quote number in the filter, all quotes containing any information included in the filter can be retrieved as a subset of all available customer quotes.
	QuoteNumber *string `json:"quoteNumber,omitempty"`
	// When a quote has been revised and updated, the quote number remains the same throughout the lifecycle of the quote, however, a Revision number is updated for each revision of the quote. The revision numbers is associated with the Unique Quote Number.
	RevisionNumber *int32 `json:"revisionNumber,omitempty"`
	// End User Name is the end customer name that is associated with a quote in Ingram Micro's CRM
	EndUserName *string `json:"endUserName,omitempty"`
	// Special Pricing Bid Number, also refers to as Dart Number relates to a unique pricing deal associated with a vendor for the quote.
	BidNumber *string `json:"bidNumber,omitempty"`
	// Total amount of quoted price for all products in the quote.
	TotalAmount *string `json:"totalAmount,omitempty"`
	// This refers to the primary status of the quote. API responses will return: Active
	QuoteStatus *string `json:"quoteStatus,omitempty"`
	// Date the Quote was initially Created
	CreatedDate *string `json:"createdDate,omitempty"`
	// Date the Quote was last updated or modified.
	LastModifiedDate *string `json:"lastModifiedDate,omitempty"`
	// Date the Quote Expires
	QuoteExpiryDate *string `json:"quoteExpiryDate,omitempty"`
}

// NewQuoteListResponseQuoteSearchResponseQuoteListInner instantiates a new QuoteListResponseQuoteSearchResponseQuoteListInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuoteListResponseQuoteSearchResponseQuoteListInner() *QuoteListResponseQuoteSearchResponseQuoteListInner {
	this := QuoteListResponseQuoteSearchResponseQuoteListInner{}
	return &this
}

// NewQuoteListResponseQuoteSearchResponseQuoteListInnerWithDefaults instantiates a new QuoteListResponseQuoteSearchResponseQuoteListInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuoteListResponseQuoteSearchResponseQuoteListInnerWithDefaults() *QuoteListResponseQuoteSearchResponseQuoteListInner {
	this := QuoteListResponseQuoteSearchResponseQuoteListInner{}
	return &this
}

// GetQuoteName returns the QuoteName field value if set, zero value otherwise.
func (o *QuoteListResponseQuoteSearchResponseQuoteListInner) GetQuoteName() string {
	if o == nil || IsNil(o.QuoteName) {
		var ret string
		return ret
	}
	return *o.QuoteName
}

// GetQuoteNameOk returns a tuple with the QuoteName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteListResponseQuoteSearchResponseQuoteListInner) GetQuoteNameOk() (*string, bool) {
	if o == nil || IsNil(o.QuoteName) {
		return nil, false
	}
	return o.QuoteName, true
}

// HasQuoteName returns a boolean if a field has been set.
func (o *QuoteListResponseQuoteSearchResponseQuoteListInner) HasQuoteName() bool {
	if o != nil && !IsNil(o.QuoteName) {
		return true
	}

	return false
}

// SetQuoteName gets a reference to the given string and assigns it to the QuoteName field.
func (o *QuoteListResponseQuoteSearchResponseQuoteListInner) SetQuoteName(v string) {
	o.QuoteName = &v
}

// GetQuoteNumber returns the QuoteNumber field value if set, zero value otherwise.
func (o *QuoteListResponseQuoteSearchResponseQuoteListInner) GetQuoteNumber() string {
	if o == nil || IsNil(o.QuoteNumber) {
		var ret string
		return ret
	}
	return *o.QuoteNumber
}

// GetQuoteNumberOk returns a tuple with the QuoteNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteListResponseQuoteSearchResponseQuoteListInner) GetQuoteNumberOk() (*string, bool) {
	if o == nil || IsNil(o.QuoteNumber) {
		return nil, false
	}
	return o.QuoteNumber, true
}

// HasQuoteNumber returns a boolean if a field has been set.
func (o *QuoteListResponseQuoteSearchResponseQuoteListInner) HasQuoteNumber() bool {
	if o != nil && !IsNil(o.QuoteNumber) {
		return true
	}

	return false
}

// SetQuoteNumber gets a reference to the given string and assigns it to the QuoteNumber field.
func (o *QuoteListResponseQuoteSearchResponseQuoteListInner) SetQuoteNumber(v string) {
	o.QuoteNumber = &v
}

// GetRevisionNumber returns the RevisionNumber field value if set, zero value otherwise.
func (o *QuoteListResponseQuoteSearchResponseQuoteListInner) GetRevisionNumber() int32 {
	if o == nil || IsNil(o.RevisionNumber) {
		var ret int32
		return ret
	}
	return *o.RevisionNumber
}

// GetRevisionNumberOk returns a tuple with the RevisionNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteListResponseQuoteSearchResponseQuoteListInner) GetRevisionNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.RevisionNumber) {
		return nil, false
	}
	return o.RevisionNumber, true
}

// HasRevisionNumber returns a boolean if a field has been set.
func (o *QuoteListResponseQuoteSearchResponseQuoteListInner) HasRevisionNumber() bool {
	if o != nil && !IsNil(o.RevisionNumber) {
		return true
	}

	return false
}

// SetRevisionNumber gets a reference to the given int32 and assigns it to the RevisionNumber field.
func (o *QuoteListResponseQuoteSearchResponseQuoteListInner) SetRevisionNumber(v int32) {
	o.RevisionNumber = &v
}

// GetEndUserName returns the EndUserName field value if set, zero value otherwise.
func (o *QuoteListResponseQuoteSearchResponseQuoteListInner) GetEndUserName() string {
	if o == nil || IsNil(o.EndUserName) {
		var ret string
		return ret
	}
	return *o.EndUserName
}

// GetEndUserNameOk returns a tuple with the EndUserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteListResponseQuoteSearchResponseQuoteListInner) GetEndUserNameOk() (*string, bool) {
	if o == nil || IsNil(o.EndUserName) {
		return nil, false
	}
	return o.EndUserName, true
}

// HasEndUserName returns a boolean if a field has been set.
func (o *QuoteListResponseQuoteSearchResponseQuoteListInner) HasEndUserName() bool {
	if o != nil && !IsNil(o.EndUserName) {
		return true
	}

	return false
}

// SetEndUserName gets a reference to the given string and assigns it to the EndUserName field.
func (o *QuoteListResponseQuoteSearchResponseQuoteListInner) SetEndUserName(v string) {
	o.EndUserName = &v
}

// GetBidNumber returns the BidNumber field value if set, zero value otherwise.
func (o *QuoteListResponseQuoteSearchResponseQuoteListInner) GetBidNumber() string {
	if o == nil || IsNil(o.BidNumber) {
		var ret string
		return ret
	}
	return *o.BidNumber
}

// GetBidNumberOk returns a tuple with the BidNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteListResponseQuoteSearchResponseQuoteListInner) GetBidNumberOk() (*string, bool) {
	if o == nil || IsNil(o.BidNumber) {
		return nil, false
	}
	return o.BidNumber, true
}

// HasBidNumber returns a boolean if a field has been set.
func (o *QuoteListResponseQuoteSearchResponseQuoteListInner) HasBidNumber() bool {
	if o != nil && !IsNil(o.BidNumber) {
		return true
	}

	return false
}

// SetBidNumber gets a reference to the given string and assigns it to the BidNumber field.
func (o *QuoteListResponseQuoteSearchResponseQuoteListInner) SetBidNumber(v string) {
	o.BidNumber = &v
}

// GetTotalAmount returns the TotalAmount field value if set, zero value otherwise.
func (o *QuoteListResponseQuoteSearchResponseQuoteListInner) GetTotalAmount() string {
	if o == nil || IsNil(o.TotalAmount) {
		var ret string
		return ret
	}
	return *o.TotalAmount
}

// GetTotalAmountOk returns a tuple with the TotalAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteListResponseQuoteSearchResponseQuoteListInner) GetTotalAmountOk() (*string, bool) {
	if o == nil || IsNil(o.TotalAmount) {
		return nil, false
	}
	return o.TotalAmount, true
}

// HasTotalAmount returns a boolean if a field has been set.
func (o *QuoteListResponseQuoteSearchResponseQuoteListInner) HasTotalAmount() bool {
	if o != nil && !IsNil(o.TotalAmount) {
		return true
	}

	return false
}

// SetTotalAmount gets a reference to the given string and assigns it to the TotalAmount field.
func (o *QuoteListResponseQuoteSearchResponseQuoteListInner) SetTotalAmount(v string) {
	o.TotalAmount = &v
}

// GetQuoteStatus returns the QuoteStatus field value if set, zero value otherwise.
func (o *QuoteListResponseQuoteSearchResponseQuoteListInner) GetQuoteStatus() string {
	if o == nil || IsNil(o.QuoteStatus) {
		var ret string
		return ret
	}
	return *o.QuoteStatus
}

// GetQuoteStatusOk returns a tuple with the QuoteStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteListResponseQuoteSearchResponseQuoteListInner) GetQuoteStatusOk() (*string, bool) {
	if o == nil || IsNil(o.QuoteStatus) {
		return nil, false
	}
	return o.QuoteStatus, true
}

// HasQuoteStatus returns a boolean if a field has been set.
func (o *QuoteListResponseQuoteSearchResponseQuoteListInner) HasQuoteStatus() bool {
	if o != nil && !IsNil(o.QuoteStatus) {
		return true
	}

	return false
}

// SetQuoteStatus gets a reference to the given string and assigns it to the QuoteStatus field.
func (o *QuoteListResponseQuoteSearchResponseQuoteListInner) SetQuoteStatus(v string) {
	o.QuoteStatus = &v
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise.
func (o *QuoteListResponseQuoteSearchResponseQuoteListInner) GetCreatedDate() string {
	if o == nil || IsNil(o.CreatedDate) {
		var ret string
		return ret
	}
	return *o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteListResponseQuoteSearchResponseQuoteListInner) GetCreatedDateOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedDate) {
		return nil, false
	}
	return o.CreatedDate, true
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *QuoteListResponseQuoteSearchResponseQuoteListInner) HasCreatedDate() bool {
	if o != nil && !IsNil(o.CreatedDate) {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given string and assigns it to the CreatedDate field.
func (o *QuoteListResponseQuoteSearchResponseQuoteListInner) SetCreatedDate(v string) {
	o.CreatedDate = &v
}

// GetLastModifiedDate returns the LastModifiedDate field value if set, zero value otherwise.
func (o *QuoteListResponseQuoteSearchResponseQuoteListInner) GetLastModifiedDate() string {
	if o == nil || IsNil(o.LastModifiedDate) {
		var ret string
		return ret
	}
	return *o.LastModifiedDate
}

// GetLastModifiedDateOk returns a tuple with the LastModifiedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteListResponseQuoteSearchResponseQuoteListInner) GetLastModifiedDateOk() (*string, bool) {
	if o == nil || IsNil(o.LastModifiedDate) {
		return nil, false
	}
	return o.LastModifiedDate, true
}

// HasLastModifiedDate returns a boolean if a field has been set.
func (o *QuoteListResponseQuoteSearchResponseQuoteListInner) HasLastModifiedDate() bool {
	if o != nil && !IsNil(o.LastModifiedDate) {
		return true
	}

	return false
}

// SetLastModifiedDate gets a reference to the given string and assigns it to the LastModifiedDate field.
func (o *QuoteListResponseQuoteSearchResponseQuoteListInner) SetLastModifiedDate(v string) {
	o.LastModifiedDate = &v
}

// GetQuoteExpiryDate returns the QuoteExpiryDate field value if set, zero value otherwise.
func (o *QuoteListResponseQuoteSearchResponseQuoteListInner) GetQuoteExpiryDate() string {
	if o == nil || IsNil(o.QuoteExpiryDate) {
		var ret string
		return ret
	}
	return *o.QuoteExpiryDate
}

// GetQuoteExpiryDateOk returns a tuple with the QuoteExpiryDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteListResponseQuoteSearchResponseQuoteListInner) GetQuoteExpiryDateOk() (*string, bool) {
	if o == nil || IsNil(o.QuoteExpiryDate) {
		return nil, false
	}
	return o.QuoteExpiryDate, true
}

// HasQuoteExpiryDate returns a boolean if a field has been set.
func (o *QuoteListResponseQuoteSearchResponseQuoteListInner) HasQuoteExpiryDate() bool {
	if o != nil && !IsNil(o.QuoteExpiryDate) {
		return true
	}

	return false
}

// SetQuoteExpiryDate gets a reference to the given string and assigns it to the QuoteExpiryDate field.
func (o *QuoteListResponseQuoteSearchResponseQuoteListInner) SetQuoteExpiryDate(v string) {
	o.QuoteExpiryDate = &v
}

func (o QuoteListResponseQuoteSearchResponseQuoteListInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuoteListResponseQuoteSearchResponseQuoteListInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.QuoteName) {
		toSerialize["quoteName"] = o.QuoteName
	}
	if !IsNil(o.QuoteNumber) {
		toSerialize["quoteNumber"] = o.QuoteNumber
	}
	if !IsNil(o.RevisionNumber) {
		toSerialize["revisionNumber"] = o.RevisionNumber
	}
	if !IsNil(o.EndUserName) {
		toSerialize["endUserName"] = o.EndUserName
	}
	if !IsNil(o.BidNumber) {
		toSerialize["bidNumber"] = o.BidNumber
	}
	if !IsNil(o.TotalAmount) {
		toSerialize["totalAmount"] = o.TotalAmount
	}
	if !IsNil(o.QuoteStatus) {
		toSerialize["quoteStatus"] = o.QuoteStatus
	}
	if !IsNil(o.CreatedDate) {
		toSerialize["createdDate"] = o.CreatedDate
	}
	if !IsNil(o.LastModifiedDate) {
		toSerialize["lastModifiedDate"] = o.LastModifiedDate
	}
	if !IsNil(o.QuoteExpiryDate) {
		toSerialize["quoteExpiryDate"] = o.QuoteExpiryDate
	}
	return toSerialize, nil
}

type NullableQuoteListResponseQuoteSearchResponseQuoteListInner struct {
	value *QuoteListResponseQuoteSearchResponseQuoteListInner
	isSet bool
}

func (v NullableQuoteListResponseQuoteSearchResponseQuoteListInner) Get() *QuoteListResponseQuoteSearchResponseQuoteListInner {
	return v.value
}

func (v *NullableQuoteListResponseQuoteSearchResponseQuoteListInner) Set(val *QuoteListResponseQuoteSearchResponseQuoteListInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQuoteListResponseQuoteSearchResponseQuoteListInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQuoteListResponseQuoteSearchResponseQuoteListInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuoteListResponseQuoteSearchResponseQuoteListInner(val *QuoteListResponseQuoteSearchResponseQuoteListInner) *NullableQuoteListResponseQuoteSearchResponseQuoteListInner {
	return &NullableQuoteListResponseQuoteSearchResponseQuoteListInner{value: val, isSet: true}
}

func (v NullableQuoteListResponseQuoteSearchResponseQuoteListInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuoteListResponseQuoteSearchResponseQuoteListInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


