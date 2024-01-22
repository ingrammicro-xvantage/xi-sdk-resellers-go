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

// checks if the ReturnsSearchResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReturnsSearchResponse{}

// ReturnsSearchResponse struct for ReturnsSearchResponse
type ReturnsSearchResponse struct {
	// Number of records found.
	RecordsFound *int32 `json:"recordsFound,omitempty"`
	// Number of records in a page.
	PageSize *int32 `json:"pageSize,omitempty"`
	// Number of page.
	PageNumber *int32 `json:"pageNumber,omitempty"`
	ReturnsClaims []ReturnsSearchResponseReturnsClaimsInner `json:"returnsClaims,omitempty"`
	// URL for the next page.
	NextPage *string `json:"nextPage,omitempty"`
	// URL for the previous page.
	PreviousPage *string `json:"previousPage,omitempty"`
}

// NewReturnsSearchResponse instantiates a new ReturnsSearchResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReturnsSearchResponse() *ReturnsSearchResponse {
	this := ReturnsSearchResponse{}
	return &this
}

// NewReturnsSearchResponseWithDefaults instantiates a new ReturnsSearchResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReturnsSearchResponseWithDefaults() *ReturnsSearchResponse {
	this := ReturnsSearchResponse{}
	return &this
}

// GetRecordsFound returns the RecordsFound field value if set, zero value otherwise.
func (o *ReturnsSearchResponse) GetRecordsFound() int32 {
	if o == nil || IsNil(o.RecordsFound) {
		var ret int32
		return ret
	}
	return *o.RecordsFound
}

// GetRecordsFoundOk returns a tuple with the RecordsFound field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnsSearchResponse) GetRecordsFoundOk() (*int32, bool) {
	if o == nil || IsNil(o.RecordsFound) {
		return nil, false
	}
	return o.RecordsFound, true
}

// HasRecordsFound returns a boolean if a field has been set.
func (o *ReturnsSearchResponse) HasRecordsFound() bool {
	if o != nil && !IsNil(o.RecordsFound) {
		return true
	}

	return false
}

// SetRecordsFound gets a reference to the given int32 and assigns it to the RecordsFound field.
func (o *ReturnsSearchResponse) SetRecordsFound(v int32) {
	o.RecordsFound = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *ReturnsSearchResponse) GetPageSize() int32 {
	if o == nil || IsNil(o.PageSize) {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnsSearchResponse) GetPageSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *ReturnsSearchResponse) HasPageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *ReturnsSearchResponse) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetPageNumber returns the PageNumber field value if set, zero value otherwise.
func (o *ReturnsSearchResponse) GetPageNumber() int32 {
	if o == nil || IsNil(o.PageNumber) {
		var ret int32
		return ret
	}
	return *o.PageNumber
}

// GetPageNumberOk returns a tuple with the PageNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnsSearchResponse) GetPageNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.PageNumber) {
		return nil, false
	}
	return o.PageNumber, true
}

// HasPageNumber returns a boolean if a field has been set.
func (o *ReturnsSearchResponse) HasPageNumber() bool {
	if o != nil && !IsNil(o.PageNumber) {
		return true
	}

	return false
}

// SetPageNumber gets a reference to the given int32 and assigns it to the PageNumber field.
func (o *ReturnsSearchResponse) SetPageNumber(v int32) {
	o.PageNumber = &v
}

// GetReturnsClaims returns the ReturnsClaims field value if set, zero value otherwise.
func (o *ReturnsSearchResponse) GetReturnsClaims() []ReturnsSearchResponseReturnsClaimsInner {
	if o == nil || IsNil(o.ReturnsClaims) {
		var ret []ReturnsSearchResponseReturnsClaimsInner
		return ret
	}
	return o.ReturnsClaims
}

// GetReturnsClaimsOk returns a tuple with the ReturnsClaims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnsSearchResponse) GetReturnsClaimsOk() ([]ReturnsSearchResponseReturnsClaimsInner, bool) {
	if o == nil || IsNil(o.ReturnsClaims) {
		return nil, false
	}
	return o.ReturnsClaims, true
}

// HasReturnsClaims returns a boolean if a field has been set.
func (o *ReturnsSearchResponse) HasReturnsClaims() bool {
	if o != nil && !IsNil(o.ReturnsClaims) {
		return true
	}

	return false
}

// SetReturnsClaims gets a reference to the given []ReturnsSearchResponseReturnsClaimsInner and assigns it to the ReturnsClaims field.
func (o *ReturnsSearchResponse) SetReturnsClaims(v []ReturnsSearchResponseReturnsClaimsInner) {
	o.ReturnsClaims = v
}

// GetNextPage returns the NextPage field value if set, zero value otherwise.
func (o *ReturnsSearchResponse) GetNextPage() string {
	if o == nil || IsNil(o.NextPage) {
		var ret string
		return ret
	}
	return *o.NextPage
}

// GetNextPageOk returns a tuple with the NextPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnsSearchResponse) GetNextPageOk() (*string, bool) {
	if o == nil || IsNil(o.NextPage) {
		return nil, false
	}
	return o.NextPage, true
}

// HasNextPage returns a boolean if a field has been set.
func (o *ReturnsSearchResponse) HasNextPage() bool {
	if o != nil && !IsNil(o.NextPage) {
		return true
	}

	return false
}

// SetNextPage gets a reference to the given string and assigns it to the NextPage field.
func (o *ReturnsSearchResponse) SetNextPage(v string) {
	o.NextPage = &v
}

// GetPreviousPage returns the PreviousPage field value if set, zero value otherwise.
func (o *ReturnsSearchResponse) GetPreviousPage() string {
	if o == nil || IsNil(o.PreviousPage) {
		var ret string
		return ret
	}
	return *o.PreviousPage
}

// GetPreviousPageOk returns a tuple with the PreviousPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnsSearchResponse) GetPreviousPageOk() (*string, bool) {
	if o == nil || IsNil(o.PreviousPage) {
		return nil, false
	}
	return o.PreviousPage, true
}

// HasPreviousPage returns a boolean if a field has been set.
func (o *ReturnsSearchResponse) HasPreviousPage() bool {
	if o != nil && !IsNil(o.PreviousPage) {
		return true
	}

	return false
}

// SetPreviousPage gets a reference to the given string and assigns it to the PreviousPage field.
func (o *ReturnsSearchResponse) SetPreviousPage(v string) {
	o.PreviousPage = &v
}

func (o ReturnsSearchResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReturnsSearchResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RecordsFound) {
		toSerialize["recordsFound"] = o.RecordsFound
	}
	if !IsNil(o.PageSize) {
		toSerialize["pageSize"] = o.PageSize
	}
	if !IsNil(o.PageNumber) {
		toSerialize["pageNumber"] = o.PageNumber
	}
	if !IsNil(o.ReturnsClaims) {
		toSerialize["returnsClaims"] = o.ReturnsClaims
	}
	if !IsNil(o.NextPage) {
		toSerialize["nextPage"] = o.NextPage
	}
	if !IsNil(o.PreviousPage) {
		toSerialize["previousPage"] = o.PreviousPage
	}
	return toSerialize, nil
}

type NullableReturnsSearchResponse struct {
	value *ReturnsSearchResponse
	isSet bool
}

func (v NullableReturnsSearchResponse) Get() *ReturnsSearchResponse {
	return v.value
}

func (v *NullableReturnsSearchResponse) Set(val *ReturnsSearchResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReturnsSearchResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReturnsSearchResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReturnsSearchResponse(val *ReturnsSearchResponse) *NullableReturnsSearchResponse {
	return &NullableReturnsSearchResponse{value: val, isSet: true}
}

func (v NullableReturnsSearchResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReturnsSearchResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


