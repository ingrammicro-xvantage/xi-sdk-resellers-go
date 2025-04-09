# QuoteCreateRequestEndUserInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyName** | Pointer to **string** | Name of the company associated with the quote. | [optional] 
**Contact** | Pointer to **string** | Contact name of end user associated with the quote. | [optional] 
**AddressLine1** | Pointer to **string** | Address line 1 for end user associated with the quote | [optional] 
**AddressLine2** | Pointer to **string** | Address line 2 for end user associated with the quote | [optional] 
**City** | Pointer to **string** | City for end user associated with the quote | [optional] 
**State** | Pointer to **string** | Two letter state abbreviation for end user associated with the quote. | [optional] 
**PostalCode** | Pointer to **string** | Zip code of end user associated with the quote. | [optional] 
**CountryCode** | Pointer to **string** | Two letter Country abbreviation for end user associated with the quote. | [optional] 
**Email** | Pointer to **string** | Email of end-user the quote associated with the quote. | [optional] 
**PhoneNumber** | Pointer to **string** | Phone number of end user associated with the quote. | [optional] 

## Methods

### NewQuoteCreateRequestEndUserInfo

`func NewQuoteCreateRequestEndUserInfo() *QuoteCreateRequestEndUserInfo`

NewQuoteCreateRequestEndUserInfo instantiates a new QuoteCreateRequestEndUserInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuoteCreateRequestEndUserInfoWithDefaults

`func NewQuoteCreateRequestEndUserInfoWithDefaults() *QuoteCreateRequestEndUserInfo`

NewQuoteCreateRequestEndUserInfoWithDefaults instantiates a new QuoteCreateRequestEndUserInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyName

`func (o *QuoteCreateRequestEndUserInfo) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *QuoteCreateRequestEndUserInfo) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *QuoteCreateRequestEndUserInfo) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *QuoteCreateRequestEndUserInfo) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetContact

`func (o *QuoteCreateRequestEndUserInfo) GetContact() string`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *QuoteCreateRequestEndUserInfo) GetContactOk() (*string, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *QuoteCreateRequestEndUserInfo) SetContact(v string)`

SetContact sets Contact field to given value.

### HasContact

`func (o *QuoteCreateRequestEndUserInfo) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetAddressLine1

`func (o *QuoteCreateRequestEndUserInfo) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *QuoteCreateRequestEndUserInfo) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *QuoteCreateRequestEndUserInfo) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.

### HasAddressLine1

`func (o *QuoteCreateRequestEndUserInfo) HasAddressLine1() bool`

HasAddressLine1 returns a boolean if a field has been set.

### GetAddressLine2

`func (o *QuoteCreateRequestEndUserInfo) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *QuoteCreateRequestEndUserInfo) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *QuoteCreateRequestEndUserInfo) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *QuoteCreateRequestEndUserInfo) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### GetCity

`func (o *QuoteCreateRequestEndUserInfo) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *QuoteCreateRequestEndUserInfo) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *QuoteCreateRequestEndUserInfo) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *QuoteCreateRequestEndUserInfo) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetState

`func (o *QuoteCreateRequestEndUserInfo) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *QuoteCreateRequestEndUserInfo) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *QuoteCreateRequestEndUserInfo) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *QuoteCreateRequestEndUserInfo) HasState() bool`

HasState returns a boolean if a field has been set.

### GetPostalCode

`func (o *QuoteCreateRequestEndUserInfo) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *QuoteCreateRequestEndUserInfo) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *QuoteCreateRequestEndUserInfo) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *QuoteCreateRequestEndUserInfo) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetCountryCode

`func (o *QuoteCreateRequestEndUserInfo) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *QuoteCreateRequestEndUserInfo) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *QuoteCreateRequestEndUserInfo) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *QuoteCreateRequestEndUserInfo) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetEmail

`func (o *QuoteCreateRequestEndUserInfo) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *QuoteCreateRequestEndUserInfo) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *QuoteCreateRequestEndUserInfo) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *QuoteCreateRequestEndUserInfo) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *QuoteCreateRequestEndUserInfo) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *QuoteCreateRequestEndUserInfo) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *QuoteCreateRequestEndUserInfo) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *QuoteCreateRequestEndUserInfo) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


