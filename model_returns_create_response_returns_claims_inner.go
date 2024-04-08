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

// checks if the ReturnsCreateResponseReturnsClaimsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReturnsCreateResponseReturnsClaimsInner{}

// ReturnsCreateResponseReturnsClaimsInner struct for ReturnsCreateResponseReturnsClaimsInner
type ReturnsCreateResponseReturnsClaimsInner struct {
	// The rmaClaimId claim id.
	RmaClaimId *string `json:"rmaClaimId,omitempty"`
	// A unique return request number.
	CaseRequestNumber *string `json:"caseRequestNumber,omitempty"`
	// The reference number for the return.
	ReferenceNumber *string `json:"referenceNumber,omitempty"`
	// The date on which the return request was created. 
	CreatedOn *string `json:"createdOn,omitempty"`
	// Type of request.
	Type *string `json:"type,omitempty"`
	// The reason for the return.
	ReturnReason *string `json:"returnReason,omitempty"`
	// Unique line number from Ingram.
	IngramPartNumber *string `json:"ingramPartNumber,omitempty"`
	// Vendor Part Number.
	VendorPartNumber *string `json:"vendorPartNumber,omitempty"`
	// Return quantity of the product.
	Quantity *int32 `json:"quantity,omitempty"`
	// Return notes.
	Notes *string `json:"notes,omitempty"`
	// The estimated total value of the return.
	EstimatedTotalValue *float32 `json:"estimatedTotalValue,omitempty"`
	// The amount of credit.
	Credit *float32 `json:"credit,omitempty"`
	// The status of the request.
	Status *string `json:"status,omitempty"`
	Links []ReturnsSearchResponseReturnsClaimsInnerLinksInner `json:"links,omitempty"`
}

// NewReturnsCreateResponseReturnsClaimsInner instantiates a new ReturnsCreateResponseReturnsClaimsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReturnsCreateResponseReturnsClaimsInner() *ReturnsCreateResponseReturnsClaimsInner {
	this := ReturnsCreateResponseReturnsClaimsInner{}
	return &this
}

// NewReturnsCreateResponseReturnsClaimsInnerWithDefaults instantiates a new ReturnsCreateResponseReturnsClaimsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReturnsCreateResponseReturnsClaimsInnerWithDefaults() *ReturnsCreateResponseReturnsClaimsInner {
	this := ReturnsCreateResponseReturnsClaimsInner{}
	return &this
}

// GetRmaClaimId returns the RmaClaimId field value if set, zero value otherwise.
func (o *ReturnsCreateResponseReturnsClaimsInner) GetRmaClaimId() string {
	if o == nil || IsNil(o.RmaClaimId) {
		var ret string
		return ret
	}
	return *o.RmaClaimId
}

// GetRmaClaimIdOk returns a tuple with the RmaClaimId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnsCreateResponseReturnsClaimsInner) GetRmaClaimIdOk() (*string, bool) {
	if o == nil || IsNil(o.RmaClaimId) {
		return nil, false
	}
	return o.RmaClaimId, true
}

// HasRmaClaimId returns a boolean if a field has been set.
func (o *ReturnsCreateResponseReturnsClaimsInner) HasRmaClaimId() bool {
	if o != nil && !IsNil(o.RmaClaimId) {
		return true
	}

	return false
}

// SetRmaClaimId gets a reference to the given string and assigns it to the RmaClaimId field.
func (o *ReturnsCreateResponseReturnsClaimsInner) SetRmaClaimId(v string) {
	o.RmaClaimId = &v
}

// GetCaseRequestNumber returns the CaseRequestNumber field value if set, zero value otherwise.
func (o *ReturnsCreateResponseReturnsClaimsInner) GetCaseRequestNumber() string {
	if o == nil || IsNil(o.CaseRequestNumber) {
		var ret string
		return ret
	}
	return *o.CaseRequestNumber
}

// GetCaseRequestNumberOk returns a tuple with the CaseRequestNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnsCreateResponseReturnsClaimsInner) GetCaseRequestNumberOk() (*string, bool) {
	if o == nil || IsNil(o.CaseRequestNumber) {
		return nil, false
	}
	return o.CaseRequestNumber, true
}

// HasCaseRequestNumber returns a boolean if a field has been set.
func (o *ReturnsCreateResponseReturnsClaimsInner) HasCaseRequestNumber() bool {
	if o != nil && !IsNil(o.CaseRequestNumber) {
		return true
	}

	return false
}

// SetCaseRequestNumber gets a reference to the given string and assigns it to the CaseRequestNumber field.
func (o *ReturnsCreateResponseReturnsClaimsInner) SetCaseRequestNumber(v string) {
	o.CaseRequestNumber = &v
}

// GetReferenceNumber returns the ReferenceNumber field value if set, zero value otherwise.
func (o *ReturnsCreateResponseReturnsClaimsInner) GetReferenceNumber() string {
	if o == nil || IsNil(o.ReferenceNumber) {
		var ret string
		return ret
	}
	return *o.ReferenceNumber
}

// GetReferenceNumberOk returns a tuple with the ReferenceNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnsCreateResponseReturnsClaimsInner) GetReferenceNumberOk() (*string, bool) {
	if o == nil || IsNil(o.ReferenceNumber) {
		return nil, false
	}
	return o.ReferenceNumber, true
}

// HasReferenceNumber returns a boolean if a field has been set.
func (o *ReturnsCreateResponseReturnsClaimsInner) HasReferenceNumber() bool {
	if o != nil && !IsNil(o.ReferenceNumber) {
		return true
	}

	return false
}

// SetReferenceNumber gets a reference to the given string and assigns it to the ReferenceNumber field.
func (o *ReturnsCreateResponseReturnsClaimsInner) SetReferenceNumber(v string) {
	o.ReferenceNumber = &v
}

// GetCreatedOn returns the CreatedOn field value if set, zero value otherwise.
func (o *ReturnsCreateResponseReturnsClaimsInner) GetCreatedOn() string {
	if o == nil || IsNil(o.CreatedOn) {
		var ret string
		return ret
	}
	return *o.CreatedOn
}

// GetCreatedOnOk returns a tuple with the CreatedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnsCreateResponseReturnsClaimsInner) GetCreatedOnOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedOn) {
		return nil, false
	}
	return o.CreatedOn, true
}

// HasCreatedOn returns a boolean if a field has been set.
func (o *ReturnsCreateResponseReturnsClaimsInner) HasCreatedOn() bool {
	if o != nil && !IsNil(o.CreatedOn) {
		return true
	}

	return false
}

// SetCreatedOn gets a reference to the given string and assigns it to the CreatedOn field.
func (o *ReturnsCreateResponseReturnsClaimsInner) SetCreatedOn(v string) {
	o.CreatedOn = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ReturnsCreateResponseReturnsClaimsInner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnsCreateResponseReturnsClaimsInner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ReturnsCreateResponseReturnsClaimsInner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ReturnsCreateResponseReturnsClaimsInner) SetType(v string) {
	o.Type = &v
}

// GetReturnReason returns the ReturnReason field value if set, zero value otherwise.
func (o *ReturnsCreateResponseReturnsClaimsInner) GetReturnReason() string {
	if o == nil || IsNil(o.ReturnReason) {
		var ret string
		return ret
	}
	return *o.ReturnReason
}

// GetReturnReasonOk returns a tuple with the ReturnReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnsCreateResponseReturnsClaimsInner) GetReturnReasonOk() (*string, bool) {
	if o == nil || IsNil(o.ReturnReason) {
		return nil, false
	}
	return o.ReturnReason, true
}

// HasReturnReason returns a boolean if a field has been set.
func (o *ReturnsCreateResponseReturnsClaimsInner) HasReturnReason() bool {
	if o != nil && !IsNil(o.ReturnReason) {
		return true
	}

	return false
}

// SetReturnReason gets a reference to the given string and assigns it to the ReturnReason field.
func (o *ReturnsCreateResponseReturnsClaimsInner) SetReturnReason(v string) {
	o.ReturnReason = &v
}

// GetIngramPartNumber returns the IngramPartNumber field value if set, zero value otherwise.
func (o *ReturnsCreateResponseReturnsClaimsInner) GetIngramPartNumber() string {
	if o == nil || IsNil(o.IngramPartNumber) {
		var ret string
		return ret
	}
	return *o.IngramPartNumber
}

// GetIngramPartNumberOk returns a tuple with the IngramPartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnsCreateResponseReturnsClaimsInner) GetIngramPartNumberOk() (*string, bool) {
	if o == nil || IsNil(o.IngramPartNumber) {
		return nil, false
	}
	return o.IngramPartNumber, true
}

// HasIngramPartNumber returns a boolean if a field has been set.
func (o *ReturnsCreateResponseReturnsClaimsInner) HasIngramPartNumber() bool {
	if o != nil && !IsNil(o.IngramPartNumber) {
		return true
	}

	return false
}

// SetIngramPartNumber gets a reference to the given string and assigns it to the IngramPartNumber field.
func (o *ReturnsCreateResponseReturnsClaimsInner) SetIngramPartNumber(v string) {
	o.IngramPartNumber = &v
}

// GetVendorPartNumber returns the VendorPartNumber field value if set, zero value otherwise.
func (o *ReturnsCreateResponseReturnsClaimsInner) GetVendorPartNumber() string {
	if o == nil || IsNil(o.VendorPartNumber) {
		var ret string
		return ret
	}
	return *o.VendorPartNumber
}

// GetVendorPartNumberOk returns a tuple with the VendorPartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnsCreateResponseReturnsClaimsInner) GetVendorPartNumberOk() (*string, bool) {
	if o == nil || IsNil(o.VendorPartNumber) {
		return nil, false
	}
	return o.VendorPartNumber, true
}

// HasVendorPartNumber returns a boolean if a field has been set.
func (o *ReturnsCreateResponseReturnsClaimsInner) HasVendorPartNumber() bool {
	if o != nil && !IsNil(o.VendorPartNumber) {
		return true
	}

	return false
}

// SetVendorPartNumber gets a reference to the given string and assigns it to the VendorPartNumber field.
func (o *ReturnsCreateResponseReturnsClaimsInner) SetVendorPartNumber(v string) {
	o.VendorPartNumber = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *ReturnsCreateResponseReturnsClaimsInner) GetQuantity() int32 {
	if o == nil || IsNil(o.Quantity) {
		var ret int32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnsCreateResponseReturnsClaimsInner) GetQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *ReturnsCreateResponseReturnsClaimsInner) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int32 and assigns it to the Quantity field.
func (o *ReturnsCreateResponseReturnsClaimsInner) SetQuantity(v int32) {
	o.Quantity = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *ReturnsCreateResponseReturnsClaimsInner) GetNotes() string {
	if o == nil || IsNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnsCreateResponseReturnsClaimsInner) GetNotesOk() (*string, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *ReturnsCreateResponseReturnsClaimsInner) HasNotes() bool {
	if o != nil && !IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *ReturnsCreateResponseReturnsClaimsInner) SetNotes(v string) {
	o.Notes = &v
}

// GetEstimatedTotalValue returns the EstimatedTotalValue field value if set, zero value otherwise.
func (o *ReturnsCreateResponseReturnsClaimsInner) GetEstimatedTotalValue() float32 {
	if o == nil || IsNil(o.EstimatedTotalValue) {
		var ret float32
		return ret
	}
	return *o.EstimatedTotalValue
}

// GetEstimatedTotalValueOk returns a tuple with the EstimatedTotalValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnsCreateResponseReturnsClaimsInner) GetEstimatedTotalValueOk() (*float32, bool) {
	if o == nil || IsNil(o.EstimatedTotalValue) {
		return nil, false
	}
	return o.EstimatedTotalValue, true
}

// HasEstimatedTotalValue returns a boolean if a field has been set.
func (o *ReturnsCreateResponseReturnsClaimsInner) HasEstimatedTotalValue() bool {
	if o != nil && !IsNil(o.EstimatedTotalValue) {
		return true
	}

	return false
}

// SetEstimatedTotalValue gets a reference to the given float32 and assigns it to the EstimatedTotalValue field.
func (o *ReturnsCreateResponseReturnsClaimsInner) SetEstimatedTotalValue(v float32) {
	o.EstimatedTotalValue = &v
}

// GetCredit returns the Credit field value if set, zero value otherwise.
func (o *ReturnsCreateResponseReturnsClaimsInner) GetCredit() float32 {
	if o == nil || IsNil(o.Credit) {
		var ret float32
		return ret
	}
	return *o.Credit
}

// GetCreditOk returns a tuple with the Credit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnsCreateResponseReturnsClaimsInner) GetCreditOk() (*float32, bool) {
	if o == nil || IsNil(o.Credit) {
		return nil, false
	}
	return o.Credit, true
}

// HasCredit returns a boolean if a field has been set.
func (o *ReturnsCreateResponseReturnsClaimsInner) HasCredit() bool {
	if o != nil && !IsNil(o.Credit) {
		return true
	}

	return false
}

// SetCredit gets a reference to the given float32 and assigns it to the Credit field.
func (o *ReturnsCreateResponseReturnsClaimsInner) SetCredit(v float32) {
	o.Credit = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ReturnsCreateResponseReturnsClaimsInner) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnsCreateResponseReturnsClaimsInner) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ReturnsCreateResponseReturnsClaimsInner) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ReturnsCreateResponseReturnsClaimsInner) SetStatus(v string) {
	o.Status = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ReturnsCreateResponseReturnsClaimsInner) GetLinks() []ReturnsSearchResponseReturnsClaimsInnerLinksInner {
	if o == nil || IsNil(o.Links) {
		var ret []ReturnsSearchResponseReturnsClaimsInnerLinksInner
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnsCreateResponseReturnsClaimsInner) GetLinksOk() ([]ReturnsSearchResponseReturnsClaimsInnerLinksInner, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ReturnsCreateResponseReturnsClaimsInner) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []ReturnsSearchResponseReturnsClaimsInnerLinksInner and assigns it to the Links field.
func (o *ReturnsCreateResponseReturnsClaimsInner) SetLinks(v []ReturnsSearchResponseReturnsClaimsInnerLinksInner) {
	o.Links = v
}

func (o ReturnsCreateResponseReturnsClaimsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReturnsCreateResponseReturnsClaimsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RmaClaimId) {
		toSerialize["rmaClaimId"] = o.RmaClaimId
	}
	if !IsNil(o.CaseRequestNumber) {
		toSerialize["caseRequestNumber"] = o.CaseRequestNumber
	}
	if !IsNil(o.ReferenceNumber) {
		toSerialize["referenceNumber"] = o.ReferenceNumber
	}
	if !IsNil(o.CreatedOn) {
		toSerialize["createdOn"] = o.CreatedOn
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.ReturnReason) {
		toSerialize["returnReason"] = o.ReturnReason
	}
	if !IsNil(o.IngramPartNumber) {
		toSerialize["ingramPartNumber"] = o.IngramPartNumber
	}
	if !IsNil(o.VendorPartNumber) {
		toSerialize["vendorPartNumber"] = o.VendorPartNumber
	}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if !IsNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	if !IsNil(o.EstimatedTotalValue) {
		toSerialize["estimatedTotalValue"] = o.EstimatedTotalValue
	}
	if !IsNil(o.Credit) {
		toSerialize["credit"] = o.Credit
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	return toSerialize, nil
}

type NullableReturnsCreateResponseReturnsClaimsInner struct {
	value *ReturnsCreateResponseReturnsClaimsInner
	isSet bool
}

func (v NullableReturnsCreateResponseReturnsClaimsInner) Get() *ReturnsCreateResponseReturnsClaimsInner {
	return v.value
}

func (v *NullableReturnsCreateResponseReturnsClaimsInner) Set(val *ReturnsCreateResponseReturnsClaimsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableReturnsCreateResponseReturnsClaimsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableReturnsCreateResponseReturnsClaimsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReturnsCreateResponseReturnsClaimsInner(val *ReturnsCreateResponseReturnsClaimsInner) *NullableReturnsCreateResponseReturnsClaimsInner {
	return &NullableReturnsCreateResponseReturnsClaimsInner{value: val, isSet: true}
}

func (v NullableReturnsCreateResponseReturnsClaimsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReturnsCreateResponseReturnsClaimsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


