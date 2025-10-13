# OrderCreateV7RequestShipToInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressId** | Pointer to **string** | The ID references the resellers address in Ingram Micro&#39;s system for shipping. Provided to resellers during the onboarding process. | [optional] 
**Contact** | Pointer to **NullableString** | The company contact provided by the reseller. | [optional] 
**CompanyName** | Pointer to **NullableString** | The reseller&#39;s company name or the End-User&#39;s Name | [optional] 
**AddressLine1** | Pointer to **NullableString** | The street address and building or house number the order will be shipped to. | [optional] 
**AddressLine2** | Pointer to **NullableString** | The apartment number the order will be shipped to. | [optional] 
**City** | Pointer to **NullableString** | The city the order will be shipped to. | [optional] 
**State** | Pointer to **NullableString** | The state the order will be shipped to. | [optional] 
**PostalCode** | Pointer to **NullableString** | End User Name | [optional] 
**CountryCode** | Pointer to **NullableString** | The zip or postal code the order will be shipped to. | [optional] 
**Email** | Pointer to **NullableString** | The company contact email address. | [optional] 
**ShippingNotes** | Pointer to **NullableString** |  | [optional] 
**PhoneNumber** | Pointer to **NullableString** | The company contact phone number. | [optional] 

## Methods

### NewOrderCreateV7RequestShipToInfo

`func NewOrderCreateV7RequestShipToInfo() *OrderCreateV7RequestShipToInfo`

NewOrderCreateV7RequestShipToInfo instantiates a new OrderCreateV7RequestShipToInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderCreateV7RequestShipToInfoWithDefaults

`func NewOrderCreateV7RequestShipToInfoWithDefaults() *OrderCreateV7RequestShipToInfo`

NewOrderCreateV7RequestShipToInfoWithDefaults instantiates a new OrderCreateV7RequestShipToInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressId

`func (o *OrderCreateV7RequestShipToInfo) GetAddressId() string`

GetAddressId returns the AddressId field if non-nil, zero value otherwise.

### GetAddressIdOk

`func (o *OrderCreateV7RequestShipToInfo) GetAddressIdOk() (*string, bool)`

GetAddressIdOk returns a tuple with the AddressId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressId

`func (o *OrderCreateV7RequestShipToInfo) SetAddressId(v string)`

SetAddressId sets AddressId field to given value.

### HasAddressId

`func (o *OrderCreateV7RequestShipToInfo) HasAddressId() bool`

HasAddressId returns a boolean if a field has been set.

### GetContact

`func (o *OrderCreateV7RequestShipToInfo) GetContact() string`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *OrderCreateV7RequestShipToInfo) GetContactOk() (*string, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *OrderCreateV7RequestShipToInfo) SetContact(v string)`

SetContact sets Contact field to given value.

### HasContact

`func (o *OrderCreateV7RequestShipToInfo) HasContact() bool`

HasContact returns a boolean if a field has been set.

### SetContactNil

`func (o *OrderCreateV7RequestShipToInfo) SetContactNil(b bool)`

 SetContactNil sets the value for Contact to be an explicit nil

### UnsetContact
`func (o *OrderCreateV7RequestShipToInfo) UnsetContact()`

UnsetContact ensures that no value is present for Contact, not even an explicit nil
### GetCompanyName

`func (o *OrderCreateV7RequestShipToInfo) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *OrderCreateV7RequestShipToInfo) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *OrderCreateV7RequestShipToInfo) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *OrderCreateV7RequestShipToInfo) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### SetCompanyNameNil

`func (o *OrderCreateV7RequestShipToInfo) SetCompanyNameNil(b bool)`

 SetCompanyNameNil sets the value for CompanyName to be an explicit nil

### UnsetCompanyName
`func (o *OrderCreateV7RequestShipToInfo) UnsetCompanyName()`

UnsetCompanyName ensures that no value is present for CompanyName, not even an explicit nil
### GetAddressLine1

`func (o *OrderCreateV7RequestShipToInfo) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *OrderCreateV7RequestShipToInfo) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *OrderCreateV7RequestShipToInfo) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.

### HasAddressLine1

`func (o *OrderCreateV7RequestShipToInfo) HasAddressLine1() bool`

HasAddressLine1 returns a boolean if a field has been set.

### SetAddressLine1Nil

`func (o *OrderCreateV7RequestShipToInfo) SetAddressLine1Nil(b bool)`

 SetAddressLine1Nil sets the value for AddressLine1 to be an explicit nil

### UnsetAddressLine1
`func (o *OrderCreateV7RequestShipToInfo) UnsetAddressLine1()`

UnsetAddressLine1 ensures that no value is present for AddressLine1, not even an explicit nil
### GetAddressLine2

`func (o *OrderCreateV7RequestShipToInfo) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *OrderCreateV7RequestShipToInfo) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *OrderCreateV7RequestShipToInfo) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *OrderCreateV7RequestShipToInfo) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### SetAddressLine2Nil

`func (o *OrderCreateV7RequestShipToInfo) SetAddressLine2Nil(b bool)`

 SetAddressLine2Nil sets the value for AddressLine2 to be an explicit nil

### UnsetAddressLine2
`func (o *OrderCreateV7RequestShipToInfo) UnsetAddressLine2()`

UnsetAddressLine2 ensures that no value is present for AddressLine2, not even an explicit nil
### GetCity

`func (o *OrderCreateV7RequestShipToInfo) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *OrderCreateV7RequestShipToInfo) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *OrderCreateV7RequestShipToInfo) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *OrderCreateV7RequestShipToInfo) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *OrderCreateV7RequestShipToInfo) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *OrderCreateV7RequestShipToInfo) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetState

`func (o *OrderCreateV7RequestShipToInfo) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *OrderCreateV7RequestShipToInfo) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *OrderCreateV7RequestShipToInfo) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *OrderCreateV7RequestShipToInfo) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *OrderCreateV7RequestShipToInfo) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *OrderCreateV7RequestShipToInfo) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetPostalCode

`func (o *OrderCreateV7RequestShipToInfo) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *OrderCreateV7RequestShipToInfo) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *OrderCreateV7RequestShipToInfo) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *OrderCreateV7RequestShipToInfo) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### SetPostalCodeNil

`func (o *OrderCreateV7RequestShipToInfo) SetPostalCodeNil(b bool)`

 SetPostalCodeNil sets the value for PostalCode to be an explicit nil

### UnsetPostalCode
`func (o *OrderCreateV7RequestShipToInfo) UnsetPostalCode()`

UnsetPostalCode ensures that no value is present for PostalCode, not even an explicit nil
### GetCountryCode

`func (o *OrderCreateV7RequestShipToInfo) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *OrderCreateV7RequestShipToInfo) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *OrderCreateV7RequestShipToInfo) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *OrderCreateV7RequestShipToInfo) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### SetCountryCodeNil

`func (o *OrderCreateV7RequestShipToInfo) SetCountryCodeNil(b bool)`

 SetCountryCodeNil sets the value for CountryCode to be an explicit nil

### UnsetCountryCode
`func (o *OrderCreateV7RequestShipToInfo) UnsetCountryCode()`

UnsetCountryCode ensures that no value is present for CountryCode, not even an explicit nil
### GetEmail

`func (o *OrderCreateV7RequestShipToInfo) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *OrderCreateV7RequestShipToInfo) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *OrderCreateV7RequestShipToInfo) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *OrderCreateV7RequestShipToInfo) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *OrderCreateV7RequestShipToInfo) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *OrderCreateV7RequestShipToInfo) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetShippingNotes

`func (o *OrderCreateV7RequestShipToInfo) GetShippingNotes() string`

GetShippingNotes returns the ShippingNotes field if non-nil, zero value otherwise.

### GetShippingNotesOk

`func (o *OrderCreateV7RequestShipToInfo) GetShippingNotesOk() (*string, bool)`

GetShippingNotesOk returns a tuple with the ShippingNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingNotes

`func (o *OrderCreateV7RequestShipToInfo) SetShippingNotes(v string)`

SetShippingNotes sets ShippingNotes field to given value.

### HasShippingNotes

`func (o *OrderCreateV7RequestShipToInfo) HasShippingNotes() bool`

HasShippingNotes returns a boolean if a field has been set.

### SetShippingNotesNil

`func (o *OrderCreateV7RequestShipToInfo) SetShippingNotesNil(b bool)`

 SetShippingNotesNil sets the value for ShippingNotes to be an explicit nil

### UnsetShippingNotes
`func (o *OrderCreateV7RequestShipToInfo) UnsetShippingNotes()`

UnsetShippingNotes ensures that no value is present for ShippingNotes, not even an explicit nil
### GetPhoneNumber

`func (o *OrderCreateV7RequestShipToInfo) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *OrderCreateV7RequestShipToInfo) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *OrderCreateV7RequestShipToInfo) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *OrderCreateV7RequestShipToInfo) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### SetPhoneNumberNil

`func (o *OrderCreateV7RequestShipToInfo) SetPhoneNumberNil(b bool)`

 SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil

### UnsetPhoneNumber
`func (o *OrderCreateV7RequestShipToInfo) UnsetPhoneNumber()`

UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


