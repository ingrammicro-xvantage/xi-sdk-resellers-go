# OrderCreateV7ResponseResourceShipToInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressId** | Pointer to **string** | The ID references the resellers address in Ingram Micro&#39;s system for shipping. Provided to resellers during the onboarding process. | [optional] 
**CompanyName** | Pointer to **string** | The reseller&#39;s company name or the End-User&#39;s Name | [optional] 
**AddressLine1** | Pointer to **string** | The street address and building or house number the order will be shipped to. | [optional] 
**AddressLine2** | Pointer to **string** | The apartment number the order will be shipped to. | [optional] 
**City** | Pointer to **string** | The city the order will be shipped to. | [optional] 
**State** | Pointer to **string** | The state the order will be shipped to. | [optional] 
**PostalCode** | Pointer to **string** | The zip or postal code the order will be shipped to. | [optional] 
**CountryCode** | Pointer to **string** | The two-character ISO country code the order will be shipped to. | [optional] 

## Methods

### NewOrderCreateV7ResponseResourceShipToInfo

`func NewOrderCreateV7ResponseResourceShipToInfo() *OrderCreateV7ResponseResourceShipToInfo`

NewOrderCreateV7ResponseResourceShipToInfo instantiates a new OrderCreateV7ResponseResourceShipToInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderCreateV7ResponseResourceShipToInfoWithDefaults

`func NewOrderCreateV7ResponseResourceShipToInfoWithDefaults() *OrderCreateV7ResponseResourceShipToInfo`

NewOrderCreateV7ResponseResourceShipToInfoWithDefaults instantiates a new OrderCreateV7ResponseResourceShipToInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressId

`func (o *OrderCreateV7ResponseResourceShipToInfo) GetAddressId() string`

GetAddressId returns the AddressId field if non-nil, zero value otherwise.

### GetAddressIdOk

`func (o *OrderCreateV7ResponseResourceShipToInfo) GetAddressIdOk() (*string, bool)`

GetAddressIdOk returns a tuple with the AddressId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressId

`func (o *OrderCreateV7ResponseResourceShipToInfo) SetAddressId(v string)`

SetAddressId sets AddressId field to given value.

### HasAddressId

`func (o *OrderCreateV7ResponseResourceShipToInfo) HasAddressId() bool`

HasAddressId returns a boolean if a field has been set.

### GetCompanyName

`func (o *OrderCreateV7ResponseResourceShipToInfo) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *OrderCreateV7ResponseResourceShipToInfo) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *OrderCreateV7ResponseResourceShipToInfo) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *OrderCreateV7ResponseResourceShipToInfo) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetAddressLine1

`func (o *OrderCreateV7ResponseResourceShipToInfo) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *OrderCreateV7ResponseResourceShipToInfo) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *OrderCreateV7ResponseResourceShipToInfo) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.

### HasAddressLine1

`func (o *OrderCreateV7ResponseResourceShipToInfo) HasAddressLine1() bool`

HasAddressLine1 returns a boolean if a field has been set.

### GetAddressLine2

`func (o *OrderCreateV7ResponseResourceShipToInfo) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *OrderCreateV7ResponseResourceShipToInfo) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *OrderCreateV7ResponseResourceShipToInfo) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *OrderCreateV7ResponseResourceShipToInfo) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### GetCity

`func (o *OrderCreateV7ResponseResourceShipToInfo) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *OrderCreateV7ResponseResourceShipToInfo) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *OrderCreateV7ResponseResourceShipToInfo) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *OrderCreateV7ResponseResourceShipToInfo) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetState

`func (o *OrderCreateV7ResponseResourceShipToInfo) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *OrderCreateV7ResponseResourceShipToInfo) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *OrderCreateV7ResponseResourceShipToInfo) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *OrderCreateV7ResponseResourceShipToInfo) HasState() bool`

HasState returns a boolean if a field has been set.

### GetPostalCode

`func (o *OrderCreateV7ResponseResourceShipToInfo) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *OrderCreateV7ResponseResourceShipToInfo) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *OrderCreateV7ResponseResourceShipToInfo) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *OrderCreateV7ResponseResourceShipToInfo) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetCountryCode

`func (o *OrderCreateV7ResponseResourceShipToInfo) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *OrderCreateV7ResponseResourceShipToInfo) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *OrderCreateV7ResponseResourceShipToInfo) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *OrderCreateV7ResponseResourceShipToInfo) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


