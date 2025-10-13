# OrderCreateV7RequestEndUserInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndUserId** | Pointer to **NullableString** | ID for the end user/customer in Ingram Micro&#39;s system. | [optional] 
**Contact** | Pointer to **NullableString** | The contact name for the end user/customer. | [optional] 
**CompanyName** | Pointer to **NullableString** | The company name for the end user/customer. Required for Impulse countries. | [optional] 
**AddressLine1** | Pointer to **NullableString** | The end user/customer&#39;s street address and building or house number. Required for Impulse countries. | [optional] 
**AddressLine2** | Pointer to **NullableString** | The end user/customer&#39;s apartment number. | [optional] 
**City** | Pointer to **NullableString** | The end user/customer&#39;s city. Required for Impulse countries. | [optional] 
**State** | Pointer to **NullableString** | The end user/customer&#39;s state. Required for Impulse countries but optional for EMEA countries. | [optional] 
**PostalCode** | Pointer to **NullableString** | The end user/customer&#39;s zip or postal code. Required for Impulse countries. | [optional] 
**CountryCode** | Pointer to **NullableString** | The end user/customer&#39;s two-character ISO country code. | [optional] 
**PhoneNumber** | Pointer to **NullableString** | The end user/customer&#39;s phone number. | [optional] 
**Email** | Pointer to **NullableString** | The end user/customer&#39;s email. | [optional] 

## Methods

### NewOrderCreateV7RequestEndUserInfo

`func NewOrderCreateV7RequestEndUserInfo() *OrderCreateV7RequestEndUserInfo`

NewOrderCreateV7RequestEndUserInfo instantiates a new OrderCreateV7RequestEndUserInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderCreateV7RequestEndUserInfoWithDefaults

`func NewOrderCreateV7RequestEndUserInfoWithDefaults() *OrderCreateV7RequestEndUserInfo`

NewOrderCreateV7RequestEndUserInfoWithDefaults instantiates a new OrderCreateV7RequestEndUserInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndUserId

`func (o *OrderCreateV7RequestEndUserInfo) GetEndUserId() string`

GetEndUserId returns the EndUserId field if non-nil, zero value otherwise.

### GetEndUserIdOk

`func (o *OrderCreateV7RequestEndUserInfo) GetEndUserIdOk() (*string, bool)`

GetEndUserIdOk returns a tuple with the EndUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndUserId

`func (o *OrderCreateV7RequestEndUserInfo) SetEndUserId(v string)`

SetEndUserId sets EndUserId field to given value.

### HasEndUserId

`func (o *OrderCreateV7RequestEndUserInfo) HasEndUserId() bool`

HasEndUserId returns a boolean if a field has been set.

### SetEndUserIdNil

`func (o *OrderCreateV7RequestEndUserInfo) SetEndUserIdNil(b bool)`

 SetEndUserIdNil sets the value for EndUserId to be an explicit nil

### UnsetEndUserId
`func (o *OrderCreateV7RequestEndUserInfo) UnsetEndUserId()`

UnsetEndUserId ensures that no value is present for EndUserId, not even an explicit nil
### GetContact

`func (o *OrderCreateV7RequestEndUserInfo) GetContact() string`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *OrderCreateV7RequestEndUserInfo) GetContactOk() (*string, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *OrderCreateV7RequestEndUserInfo) SetContact(v string)`

SetContact sets Contact field to given value.

### HasContact

`func (o *OrderCreateV7RequestEndUserInfo) HasContact() bool`

HasContact returns a boolean if a field has been set.

### SetContactNil

`func (o *OrderCreateV7RequestEndUserInfo) SetContactNil(b bool)`

 SetContactNil sets the value for Contact to be an explicit nil

### UnsetContact
`func (o *OrderCreateV7RequestEndUserInfo) UnsetContact()`

UnsetContact ensures that no value is present for Contact, not even an explicit nil
### GetCompanyName

`func (o *OrderCreateV7RequestEndUserInfo) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *OrderCreateV7RequestEndUserInfo) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *OrderCreateV7RequestEndUserInfo) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *OrderCreateV7RequestEndUserInfo) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### SetCompanyNameNil

`func (o *OrderCreateV7RequestEndUserInfo) SetCompanyNameNil(b bool)`

 SetCompanyNameNil sets the value for CompanyName to be an explicit nil

### UnsetCompanyName
`func (o *OrderCreateV7RequestEndUserInfo) UnsetCompanyName()`

UnsetCompanyName ensures that no value is present for CompanyName, not even an explicit nil
### GetAddressLine1

`func (o *OrderCreateV7RequestEndUserInfo) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *OrderCreateV7RequestEndUserInfo) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *OrderCreateV7RequestEndUserInfo) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.

### HasAddressLine1

`func (o *OrderCreateV7RequestEndUserInfo) HasAddressLine1() bool`

HasAddressLine1 returns a boolean if a field has been set.

### SetAddressLine1Nil

`func (o *OrderCreateV7RequestEndUserInfo) SetAddressLine1Nil(b bool)`

 SetAddressLine1Nil sets the value for AddressLine1 to be an explicit nil

### UnsetAddressLine1
`func (o *OrderCreateV7RequestEndUserInfo) UnsetAddressLine1()`

UnsetAddressLine1 ensures that no value is present for AddressLine1, not even an explicit nil
### GetAddressLine2

`func (o *OrderCreateV7RequestEndUserInfo) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *OrderCreateV7RequestEndUserInfo) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *OrderCreateV7RequestEndUserInfo) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *OrderCreateV7RequestEndUserInfo) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### SetAddressLine2Nil

`func (o *OrderCreateV7RequestEndUserInfo) SetAddressLine2Nil(b bool)`

 SetAddressLine2Nil sets the value for AddressLine2 to be an explicit nil

### UnsetAddressLine2
`func (o *OrderCreateV7RequestEndUserInfo) UnsetAddressLine2()`

UnsetAddressLine2 ensures that no value is present for AddressLine2, not even an explicit nil
### GetCity

`func (o *OrderCreateV7RequestEndUserInfo) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *OrderCreateV7RequestEndUserInfo) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *OrderCreateV7RequestEndUserInfo) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *OrderCreateV7RequestEndUserInfo) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *OrderCreateV7RequestEndUserInfo) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *OrderCreateV7RequestEndUserInfo) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetState

`func (o *OrderCreateV7RequestEndUserInfo) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *OrderCreateV7RequestEndUserInfo) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *OrderCreateV7RequestEndUserInfo) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *OrderCreateV7RequestEndUserInfo) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *OrderCreateV7RequestEndUserInfo) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *OrderCreateV7RequestEndUserInfo) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetPostalCode

`func (o *OrderCreateV7RequestEndUserInfo) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *OrderCreateV7RequestEndUserInfo) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *OrderCreateV7RequestEndUserInfo) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *OrderCreateV7RequestEndUserInfo) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### SetPostalCodeNil

`func (o *OrderCreateV7RequestEndUserInfo) SetPostalCodeNil(b bool)`

 SetPostalCodeNil sets the value for PostalCode to be an explicit nil

### UnsetPostalCode
`func (o *OrderCreateV7RequestEndUserInfo) UnsetPostalCode()`

UnsetPostalCode ensures that no value is present for PostalCode, not even an explicit nil
### GetCountryCode

`func (o *OrderCreateV7RequestEndUserInfo) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *OrderCreateV7RequestEndUserInfo) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *OrderCreateV7RequestEndUserInfo) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *OrderCreateV7RequestEndUserInfo) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### SetCountryCodeNil

`func (o *OrderCreateV7RequestEndUserInfo) SetCountryCodeNil(b bool)`

 SetCountryCodeNil sets the value for CountryCode to be an explicit nil

### UnsetCountryCode
`func (o *OrderCreateV7RequestEndUserInfo) UnsetCountryCode()`

UnsetCountryCode ensures that no value is present for CountryCode, not even an explicit nil
### GetPhoneNumber

`func (o *OrderCreateV7RequestEndUserInfo) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *OrderCreateV7RequestEndUserInfo) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *OrderCreateV7RequestEndUserInfo) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *OrderCreateV7RequestEndUserInfo) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### SetPhoneNumberNil

`func (o *OrderCreateV7RequestEndUserInfo) SetPhoneNumberNil(b bool)`

 SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil

### UnsetPhoneNumber
`func (o *OrderCreateV7RequestEndUserInfo) UnsetPhoneNumber()`

UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
### GetEmail

`func (o *OrderCreateV7RequestEndUserInfo) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *OrderCreateV7RequestEndUserInfo) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *OrderCreateV7RequestEndUserInfo) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *OrderCreateV7RequestEndUserInfo) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *OrderCreateV7RequestEndUserInfo) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *OrderCreateV7RequestEndUserInfo) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


