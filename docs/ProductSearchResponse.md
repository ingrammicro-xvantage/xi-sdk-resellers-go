# ProductSearchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordsFound** | Pointer to **int32** | The number of recourds found for the search. | [optional] 
**PageSize** | Pointer to **int32** | The number of results per page. Default is 25. | [optional] 
**PageNumber** | Pointer to **int32** | current page number default is 1 | [optional] 
**Catalog** | Pointer to [**[]ProductSearchResponseCatalogInner**](ProductSearchResponseCatalogInner.md) |  | [optional] 
**SubscriptionCatalog** | Pointer to [**[]ProductSearchResponseSubscriptionCatalogInner**](ProductSearchResponseSubscriptionCatalogInner.md) |  | [optional] 
**NextPage** | Pointer to **string** | link/URL for accessing next page. | [optional] 
**PreviousPage** | Pointer to **string** | link/URL for accessing previous page. | [optional] 

## Methods

### NewProductSearchResponse

`func NewProductSearchResponse() *ProductSearchResponse`

NewProductSearchResponse instantiates a new ProductSearchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductSearchResponseWithDefaults

`func NewProductSearchResponseWithDefaults() *ProductSearchResponse`

NewProductSearchResponseWithDefaults instantiates a new ProductSearchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordsFound

`func (o *ProductSearchResponse) GetRecordsFound() int32`

GetRecordsFound returns the RecordsFound field if non-nil, zero value otherwise.

### GetRecordsFoundOk

`func (o *ProductSearchResponse) GetRecordsFoundOk() (*int32, bool)`

GetRecordsFoundOk returns a tuple with the RecordsFound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordsFound

`func (o *ProductSearchResponse) SetRecordsFound(v int32)`

SetRecordsFound sets RecordsFound field to given value.

### HasRecordsFound

`func (o *ProductSearchResponse) HasRecordsFound() bool`

HasRecordsFound returns a boolean if a field has been set.

### GetPageSize

`func (o *ProductSearchResponse) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ProductSearchResponse) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ProductSearchResponse) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *ProductSearchResponse) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetPageNumber

`func (o *ProductSearchResponse) GetPageNumber() int32`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *ProductSearchResponse) GetPageNumberOk() (*int32, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *ProductSearchResponse) SetPageNumber(v int32)`

SetPageNumber sets PageNumber field to given value.

### HasPageNumber

`func (o *ProductSearchResponse) HasPageNumber() bool`

HasPageNumber returns a boolean if a field has been set.

### GetCatalog

`func (o *ProductSearchResponse) GetCatalog() []ProductSearchResponseCatalogInner`

GetCatalog returns the Catalog field if non-nil, zero value otherwise.

### GetCatalogOk

`func (o *ProductSearchResponse) GetCatalogOk() (*[]ProductSearchResponseCatalogInner, bool)`

GetCatalogOk returns a tuple with the Catalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalog

`func (o *ProductSearchResponse) SetCatalog(v []ProductSearchResponseCatalogInner)`

SetCatalog sets Catalog field to given value.

### HasCatalog

`func (o *ProductSearchResponse) HasCatalog() bool`

HasCatalog returns a boolean if a field has been set.

### GetSubscriptionCatalog

`func (o *ProductSearchResponse) GetSubscriptionCatalog() []ProductSearchResponseSubscriptionCatalogInner`

GetSubscriptionCatalog returns the SubscriptionCatalog field if non-nil, zero value otherwise.

### GetSubscriptionCatalogOk

`func (o *ProductSearchResponse) GetSubscriptionCatalogOk() (*[]ProductSearchResponseSubscriptionCatalogInner, bool)`

GetSubscriptionCatalogOk returns a tuple with the SubscriptionCatalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionCatalog

`func (o *ProductSearchResponse) SetSubscriptionCatalog(v []ProductSearchResponseSubscriptionCatalogInner)`

SetSubscriptionCatalog sets SubscriptionCatalog field to given value.

### HasSubscriptionCatalog

`func (o *ProductSearchResponse) HasSubscriptionCatalog() bool`

HasSubscriptionCatalog returns a boolean if a field has been set.

### GetNextPage

`func (o *ProductSearchResponse) GetNextPage() string`

GetNextPage returns the NextPage field if non-nil, zero value otherwise.

### GetNextPageOk

`func (o *ProductSearchResponse) GetNextPageOk() (*string, bool)`

GetNextPageOk returns a tuple with the NextPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPage

`func (o *ProductSearchResponse) SetNextPage(v string)`

SetNextPage sets NextPage field to given value.

### HasNextPage

`func (o *ProductSearchResponse) HasNextPage() bool`

HasNextPage returns a boolean if a field has been set.

### GetPreviousPage

`func (o *ProductSearchResponse) GetPreviousPage() string`

GetPreviousPage returns the PreviousPage field if non-nil, zero value otherwise.

### GetPreviousPageOk

`func (o *ProductSearchResponse) GetPreviousPageOk() (*string, bool)`

GetPreviousPageOk returns a tuple with the PreviousPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousPage

`func (o *ProductSearchResponse) SetPreviousPage(v string)`

SetPreviousPage sets PreviousPage field to given value.

### HasPreviousPage

`func (o *ProductSearchResponse) HasPreviousPage() bool`

HasPreviousPage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


