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

// checks if the QuoteSearchResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuoteSearchResponse{}

// QuoteSearchResponse struct for QuoteSearchResponse
type QuoteSearchResponse struct {
	// Total count of quotes retrieved in the request response.
	RecordsFound *int32 `json:"recordsFound,omitempty"`
	// Number of records (quotes) displayed per page in the quote list.
	PageSize *int32 `json:"pageSize,omitempty"`
	// Page index or page number for the list of quotes being returned.
	PageNumber *int32 `json:"pageNumber,omitempty"`
	// The quote details for the requested criteria.
	Quotes []QuoteSearchResponseQuotesInner `json:"quotes,omitempty"`
	NextPage *string `json:"nextPage,omitempty"`
	PrevPage *string `json:"prevPage,omitempty"`
}

// NewQuoteSearchResponse instantiates a new QuoteSearchResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuoteSearchResponse() *QuoteSearchResponse {
	this := QuoteSearchResponse{}
	return &this
}

// NewQuoteSearchResponseWithDefaults instantiates a new QuoteSearchResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuoteSearchResponseWithDefaults() *QuoteSearchResponse {
	this := QuoteSearchResponse{}
	return &this
}

// GetRecordsFound returns the RecordsFound field value if set, zero value otherwise.
func (o *QuoteSearchResponse) GetRecordsFound() int32 {
	if o == nil || IsNil(o.RecordsFound) {
		var ret int32
		return ret
	}
	return *o.RecordsFound
}

// GetRecordsFoundOk returns a tuple with the RecordsFound field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteSearchResponse) GetRecordsFoundOk() (*int32, bool) {
	if o == nil || IsNil(o.RecordsFound) {
		return nil, false
	}
	return o.RecordsFound, true
}

// HasRecordsFound returns a boolean if a field has been set.
func (o *QuoteSearchResponse) HasRecordsFound() bool {
	if o != nil && !IsNil(o.RecordsFound) {
		return true
	}

	return false
}

// SetRecordsFound gets a reference to the given int32 and assigns it to the RecordsFound field.
func (o *QuoteSearchResponse) SetRecordsFound(v int32) {
	o.RecordsFound = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *QuoteSearchResponse) GetPageSize() int32 {
	if o == nil || IsNil(o.PageSize) {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteSearchResponse) GetPageSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *QuoteSearchResponse) HasPageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *QuoteSearchResponse) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetPageNumber returns the PageNumber field value if set, zero value otherwise.
func (o *QuoteSearchResponse) GetPageNumber() int32 {
	if o == nil || IsNil(o.PageNumber) {
		var ret int32
		return ret
	}
	return *o.PageNumber
}

// GetPageNumberOk returns a tuple with the PageNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteSearchResponse) GetPageNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.PageNumber) {
		return nil, false
	}
	return o.PageNumber, true
}

// HasPageNumber returns a boolean if a field has been set.
func (o *QuoteSearchResponse) HasPageNumber() bool {
	if o != nil && !IsNil(o.PageNumber) {
		return true
	}

	return false
}

// SetPageNumber gets a reference to the given int32 and assigns it to the PageNumber field.
func (o *QuoteSearchResponse) SetPageNumber(v int32) {
	o.PageNumber = &v
}

// GetQuotes returns the Quotes field value if set, zero value otherwise.
func (o *QuoteSearchResponse) GetQuotes() []QuoteSearchResponseQuotesInner {
	if o == nil || IsNil(o.Quotes) {
		var ret []QuoteSearchResponseQuotesInner
		return ret
	}
	return o.Quotes
}

// GetQuotesOk returns a tuple with the Quotes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteSearchResponse) GetQuotesOk() ([]QuoteSearchResponseQuotesInner, bool) {
	if o == nil || IsNil(o.Quotes) {
		return nil, false
	}
	return o.Quotes, true
}

// HasQuotes returns a boolean if a field has been set.
func (o *QuoteSearchResponse) HasQuotes() bool {
	if o != nil && !IsNil(o.Quotes) {
		return true
	}

	return false
}

// SetQuotes gets a reference to the given []QuoteSearchResponseQuotesInner and assigns it to the Quotes field.
func (o *QuoteSearchResponse) SetQuotes(v []QuoteSearchResponseQuotesInner) {
	o.Quotes = v
}

// GetNextPage returns the NextPage field value if set, zero value otherwise.
func (o *QuoteSearchResponse) GetNextPage() string {
	if o == nil || IsNil(o.NextPage) {
		var ret string
		return ret
	}
	return *o.NextPage
}

// GetNextPageOk returns a tuple with the NextPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteSearchResponse) GetNextPageOk() (*string, bool) {
	if o == nil || IsNil(o.NextPage) {
		return nil, false
	}
	return o.NextPage, true
}

// HasNextPage returns a boolean if a field has been set.
func (o *QuoteSearchResponse) HasNextPage() bool {
	if o != nil && !IsNil(o.NextPage) {
		return true
	}

	return false
}

// SetNextPage gets a reference to the given string and assigns it to the NextPage field.
func (o *QuoteSearchResponse) SetNextPage(v string) {
	o.NextPage = &v
}

// GetPrevPage returns the PrevPage field value if set, zero value otherwise.
func (o *QuoteSearchResponse) GetPrevPage() string {
	if o == nil || IsNil(o.PrevPage) {
		var ret string
		return ret
	}
	return *o.PrevPage
}

// GetPrevPageOk returns a tuple with the PrevPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteSearchResponse) GetPrevPageOk() (*string, bool) {
	if o == nil || IsNil(o.PrevPage) {
		return nil, false
	}
	return o.PrevPage, true
}

// HasPrevPage returns a boolean if a field has been set.
func (o *QuoteSearchResponse) HasPrevPage() bool {
	if o != nil && !IsNil(o.PrevPage) {
		return true
	}

	return false
}

// SetPrevPage gets a reference to the given string and assigns it to the PrevPage field.
func (o *QuoteSearchResponse) SetPrevPage(v string) {
	o.PrevPage = &v
}

func (o QuoteSearchResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuoteSearchResponse) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Quotes) {
		toSerialize["quotes"] = o.Quotes
	}
	if !IsNil(o.NextPage) {
		toSerialize["nextPage"] = o.NextPage
	}
	if !IsNil(o.PrevPage) {
		toSerialize["prevPage"] = o.PrevPage
	}
	return toSerialize, nil
}

type NullableQuoteSearchResponse struct {
	value *QuoteSearchResponse
	isSet bool
}

func (v NullableQuoteSearchResponse) Get() *QuoteSearchResponse {
	return v.value
}

func (v *NullableQuoteSearchResponse) Set(val *QuoteSearchResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQuoteSearchResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQuoteSearchResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuoteSearchResponse(val *QuoteSearchResponse) *NullableQuoteSearchResponse {
	return &NullableQuoteSearchResponse{value: val, isSet: true}
}

func (v NullableQuoteSearchResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuoteSearchResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


