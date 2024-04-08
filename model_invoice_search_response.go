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

// checks if the InvoiceSearchResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvoiceSearchResponse{}

// InvoiceSearchResponse struct for InvoiceSearchResponse
type InvoiceSearchResponse struct {
	// Total count of quotes retrieved in the request response.
	RecordsFound *int32 `json:"recordsFound,omitempty"`
	// Number of records (quotes) displayed per page in the quote list.
	PageSize *int32 `json:"pageSize,omitempty"`
	// Page index or page number for the list of quotes being returned.
	PageNumber *int32 `json:"pageNumber,omitempty"`
	// The Invoices details for the requested criteria.
	Invoices []InvoiceSearchResponseInvoicesInner `json:"invoices,omitempty"`
	// Next page of the pagination.
	NextPage *string `json:"nextPage,omitempty"`
}

// NewInvoiceSearchResponse instantiates a new InvoiceSearchResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoiceSearchResponse() *InvoiceSearchResponse {
	this := InvoiceSearchResponse{}
	return &this
}

// NewInvoiceSearchResponseWithDefaults instantiates a new InvoiceSearchResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoiceSearchResponseWithDefaults() *InvoiceSearchResponse {
	this := InvoiceSearchResponse{}
	return &this
}

// GetRecordsFound returns the RecordsFound field value if set, zero value otherwise.
func (o *InvoiceSearchResponse) GetRecordsFound() int32 {
	if o == nil || IsNil(o.RecordsFound) {
		var ret int32
		return ret
	}
	return *o.RecordsFound
}

// GetRecordsFoundOk returns a tuple with the RecordsFound field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceSearchResponse) GetRecordsFoundOk() (*int32, bool) {
	if o == nil || IsNil(o.RecordsFound) {
		return nil, false
	}
	return o.RecordsFound, true
}

// HasRecordsFound returns a boolean if a field has been set.
func (o *InvoiceSearchResponse) HasRecordsFound() bool {
	if o != nil && !IsNil(o.RecordsFound) {
		return true
	}

	return false
}

// SetRecordsFound gets a reference to the given int32 and assigns it to the RecordsFound field.
func (o *InvoiceSearchResponse) SetRecordsFound(v int32) {
	o.RecordsFound = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *InvoiceSearchResponse) GetPageSize() int32 {
	if o == nil || IsNil(o.PageSize) {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceSearchResponse) GetPageSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *InvoiceSearchResponse) HasPageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *InvoiceSearchResponse) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetPageNumber returns the PageNumber field value if set, zero value otherwise.
func (o *InvoiceSearchResponse) GetPageNumber() int32 {
	if o == nil || IsNil(o.PageNumber) {
		var ret int32
		return ret
	}
	return *o.PageNumber
}

// GetPageNumberOk returns a tuple with the PageNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceSearchResponse) GetPageNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.PageNumber) {
		return nil, false
	}
	return o.PageNumber, true
}

// HasPageNumber returns a boolean if a field has been set.
func (o *InvoiceSearchResponse) HasPageNumber() bool {
	if o != nil && !IsNil(o.PageNumber) {
		return true
	}

	return false
}

// SetPageNumber gets a reference to the given int32 and assigns it to the PageNumber field.
func (o *InvoiceSearchResponse) SetPageNumber(v int32) {
	o.PageNumber = &v
}

// GetInvoices returns the Invoices field value if set, zero value otherwise.
func (o *InvoiceSearchResponse) GetInvoices() []InvoiceSearchResponseInvoicesInner {
	if o == nil || IsNil(o.Invoices) {
		var ret []InvoiceSearchResponseInvoicesInner
		return ret
	}
	return o.Invoices
}

// GetInvoicesOk returns a tuple with the Invoices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceSearchResponse) GetInvoicesOk() ([]InvoiceSearchResponseInvoicesInner, bool) {
	if o == nil || IsNil(o.Invoices) {
		return nil, false
	}
	return o.Invoices, true
}

// HasInvoices returns a boolean if a field has been set.
func (o *InvoiceSearchResponse) HasInvoices() bool {
	if o != nil && !IsNil(o.Invoices) {
		return true
	}

	return false
}

// SetInvoices gets a reference to the given []InvoiceSearchResponseInvoicesInner and assigns it to the Invoices field.
func (o *InvoiceSearchResponse) SetInvoices(v []InvoiceSearchResponseInvoicesInner) {
	o.Invoices = v
}

// GetNextPage returns the NextPage field value if set, zero value otherwise.
func (o *InvoiceSearchResponse) GetNextPage() string {
	if o == nil || IsNil(o.NextPage) {
		var ret string
		return ret
	}
	return *o.NextPage
}

// GetNextPageOk returns a tuple with the NextPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceSearchResponse) GetNextPageOk() (*string, bool) {
	if o == nil || IsNil(o.NextPage) {
		return nil, false
	}
	return o.NextPage, true
}

// HasNextPage returns a boolean if a field has been set.
func (o *InvoiceSearchResponse) HasNextPage() bool {
	if o != nil && !IsNil(o.NextPage) {
		return true
	}

	return false
}

// SetNextPage gets a reference to the given string and assigns it to the NextPage field.
func (o *InvoiceSearchResponse) SetNextPage(v string) {
	o.NextPage = &v
}

func (o InvoiceSearchResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvoiceSearchResponse) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Invoices) {
		toSerialize["invoices"] = o.Invoices
	}
	if !IsNil(o.NextPage) {
		toSerialize["nextPage"] = o.NextPage
	}
	return toSerialize, nil
}

type NullableInvoiceSearchResponse struct {
	value *InvoiceSearchResponse
	isSet bool
}

func (v NullableInvoiceSearchResponse) Get() *InvoiceSearchResponse {
	return v.value
}

func (v *NullableInvoiceSearchResponse) Set(val *InvoiceSearchResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoiceSearchResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoiceSearchResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoiceSearchResponse(val *InvoiceSearchResponse) *NullableInvoiceSearchResponse {
	return &NullableInvoiceSearchResponse{value: val, isSet: true}
}

func (v NullableInvoiceSearchResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoiceSearchResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


