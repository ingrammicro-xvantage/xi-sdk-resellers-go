# QuoteCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QuoteName** | Pointer to **string** | Quote Name given to quote by sales team or system generated. Generally used as a reference to identify the quote. | [optional] 
**Firstname** | Pointer to **string** | Logged in Users firstname | [optional] 
**Lastname** | Pointer to **string** | Logged in Users Lastname | [optional] 
**CustomerContact** | Pointer to **string** | Logged in Users email address contact. | [optional] 
**QuoteExpiryDate** | Pointer to **string** | The date on which a quote will expire. | [optional] 
**CustomerNeed** | Pointer to **string** | Any special need from the customer. | [optional] 
**EndUserInfo** | Pointer to [**QuoteCreateRequestEndUserInfo**](QuoteCreateRequestEndUserInfo.md) |  | [optional] 
**DealId** | Pointer to **string** | Price discount identifyer to specify a pricing discount that has been applied to the quote. | [optional] 
**PricingType** | Pointer to **string** | Pricing type of the quote. | [optional] 
**SendQuoteCopy** | Pointer to **string** | List of email addressed to whom the quote will be emailed after it&#39;s created. (Max 10 email ids) | [optional] 
**Products** | Pointer to [**[]QuoteCreateRequestProductsInner**](QuoteCreateRequestProductsInner.md) |  | [optional] 

## Methods

### NewQuoteCreateRequest

`func NewQuoteCreateRequest() *QuoteCreateRequest`

NewQuoteCreateRequest instantiates a new QuoteCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuoteCreateRequestWithDefaults

`func NewQuoteCreateRequestWithDefaults() *QuoteCreateRequest`

NewQuoteCreateRequestWithDefaults instantiates a new QuoteCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuoteName

`func (o *QuoteCreateRequest) GetQuoteName() string`

GetQuoteName returns the QuoteName field if non-nil, zero value otherwise.

### GetQuoteNameOk

`func (o *QuoteCreateRequest) GetQuoteNameOk() (*string, bool)`

GetQuoteNameOk returns a tuple with the QuoteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteName

`func (o *QuoteCreateRequest) SetQuoteName(v string)`

SetQuoteName sets QuoteName field to given value.

### HasQuoteName

`func (o *QuoteCreateRequest) HasQuoteName() bool`

HasQuoteName returns a boolean if a field has been set.

### GetFirstname

`func (o *QuoteCreateRequest) GetFirstname() string`

GetFirstname returns the Firstname field if non-nil, zero value otherwise.

### GetFirstnameOk

`func (o *QuoteCreateRequest) GetFirstnameOk() (*string, bool)`

GetFirstnameOk returns a tuple with the Firstname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstname

`func (o *QuoteCreateRequest) SetFirstname(v string)`

SetFirstname sets Firstname field to given value.

### HasFirstname

`func (o *QuoteCreateRequest) HasFirstname() bool`

HasFirstname returns a boolean if a field has been set.

### GetLastname

`func (o *QuoteCreateRequest) GetLastname() string`

GetLastname returns the Lastname field if non-nil, zero value otherwise.

### GetLastnameOk

`func (o *QuoteCreateRequest) GetLastnameOk() (*string, bool)`

GetLastnameOk returns a tuple with the Lastname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastname

`func (o *QuoteCreateRequest) SetLastname(v string)`

SetLastname sets Lastname field to given value.

### HasLastname

`func (o *QuoteCreateRequest) HasLastname() bool`

HasLastname returns a boolean if a field has been set.

### GetCustomerContact

`func (o *QuoteCreateRequest) GetCustomerContact() string`

GetCustomerContact returns the CustomerContact field if non-nil, zero value otherwise.

### GetCustomerContactOk

`func (o *QuoteCreateRequest) GetCustomerContactOk() (*string, bool)`

GetCustomerContactOk returns a tuple with the CustomerContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerContact

`func (o *QuoteCreateRequest) SetCustomerContact(v string)`

SetCustomerContact sets CustomerContact field to given value.

### HasCustomerContact

`func (o *QuoteCreateRequest) HasCustomerContact() bool`

HasCustomerContact returns a boolean if a field has been set.

### GetQuoteExpiryDate

`func (o *QuoteCreateRequest) GetQuoteExpiryDate() string`

GetQuoteExpiryDate returns the QuoteExpiryDate field if non-nil, zero value otherwise.

### GetQuoteExpiryDateOk

`func (o *QuoteCreateRequest) GetQuoteExpiryDateOk() (*string, bool)`

GetQuoteExpiryDateOk returns a tuple with the QuoteExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteExpiryDate

`func (o *QuoteCreateRequest) SetQuoteExpiryDate(v string)`

SetQuoteExpiryDate sets QuoteExpiryDate field to given value.

### HasQuoteExpiryDate

`func (o *QuoteCreateRequest) HasQuoteExpiryDate() bool`

HasQuoteExpiryDate returns a boolean if a field has been set.

### GetCustomerNeed

`func (o *QuoteCreateRequest) GetCustomerNeed() string`

GetCustomerNeed returns the CustomerNeed field if non-nil, zero value otherwise.

### GetCustomerNeedOk

`func (o *QuoteCreateRequest) GetCustomerNeedOk() (*string, bool)`

GetCustomerNeedOk returns a tuple with the CustomerNeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerNeed

`func (o *QuoteCreateRequest) SetCustomerNeed(v string)`

SetCustomerNeed sets CustomerNeed field to given value.

### HasCustomerNeed

`func (o *QuoteCreateRequest) HasCustomerNeed() bool`

HasCustomerNeed returns a boolean if a field has been set.

### GetEndUserInfo

`func (o *QuoteCreateRequest) GetEndUserInfo() QuoteCreateRequestEndUserInfo`

GetEndUserInfo returns the EndUserInfo field if non-nil, zero value otherwise.

### GetEndUserInfoOk

`func (o *QuoteCreateRequest) GetEndUserInfoOk() (*QuoteCreateRequestEndUserInfo, bool)`

GetEndUserInfoOk returns a tuple with the EndUserInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndUserInfo

`func (o *QuoteCreateRequest) SetEndUserInfo(v QuoteCreateRequestEndUserInfo)`

SetEndUserInfo sets EndUserInfo field to given value.

### HasEndUserInfo

`func (o *QuoteCreateRequest) HasEndUserInfo() bool`

HasEndUserInfo returns a boolean if a field has been set.

### GetDealId

`func (o *QuoteCreateRequest) GetDealId() string`

GetDealId returns the DealId field if non-nil, zero value otherwise.

### GetDealIdOk

`func (o *QuoteCreateRequest) GetDealIdOk() (*string, bool)`

GetDealIdOk returns a tuple with the DealId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealId

`func (o *QuoteCreateRequest) SetDealId(v string)`

SetDealId sets DealId field to given value.

### HasDealId

`func (o *QuoteCreateRequest) HasDealId() bool`

HasDealId returns a boolean if a field has been set.

### GetPricingType

`func (o *QuoteCreateRequest) GetPricingType() string`

GetPricingType returns the PricingType field if non-nil, zero value otherwise.

### GetPricingTypeOk

`func (o *QuoteCreateRequest) GetPricingTypeOk() (*string, bool)`

GetPricingTypeOk returns a tuple with the PricingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingType

`func (o *QuoteCreateRequest) SetPricingType(v string)`

SetPricingType sets PricingType field to given value.

### HasPricingType

`func (o *QuoteCreateRequest) HasPricingType() bool`

HasPricingType returns a boolean if a field has been set.

### GetSendQuoteCopy

`func (o *QuoteCreateRequest) GetSendQuoteCopy() string`

GetSendQuoteCopy returns the SendQuoteCopy field if non-nil, zero value otherwise.

### GetSendQuoteCopyOk

`func (o *QuoteCreateRequest) GetSendQuoteCopyOk() (*string, bool)`

GetSendQuoteCopyOk returns a tuple with the SendQuoteCopy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendQuoteCopy

`func (o *QuoteCreateRequest) SetSendQuoteCopy(v string)`

SetSendQuoteCopy sets SendQuoteCopy field to given value.

### HasSendQuoteCopy

`func (o *QuoteCreateRequest) HasSendQuoteCopy() bool`

HasSendQuoteCopy returns a boolean if a field has been set.

### GetProducts

`func (o *QuoteCreateRequest) GetProducts() []QuoteCreateRequestProductsInner`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *QuoteCreateRequest) GetProductsOk() (*[]QuoteCreateRequestProductsInner, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *QuoteCreateRequest) SetProducts(v []QuoteCreateRequestProductsInner)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *QuoteCreateRequest) HasProducts() bool`

HasProducts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


