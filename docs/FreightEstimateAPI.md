# \FreightEstimateAPI

All URIs are relative to *https://api.ingrammicro.com:443*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostFreightestimate**](FreightEstimateAPI.md#PostFreightestimate) | **Post** /resellers/v6/freightestimate | Freight Estimate



## PostFreightestimate

> FreightResponse PostFreightestimate(ctx).IMCustomerNumber(iMCustomerNumber).IMCountryCode(iMCountryCode).IMCorrelationID(iMCorrelationID).IMCustomerContact(iMCustomerContact).IMSenderID(iMSenderID).FreightRequest(freightRequest).Execute()

Freight Estimate



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	xi_sdk_resellers "https://github.com/ingrammicro-xvantage/xi-sdk-resellers-go"
)

func main() {
	iMCustomerNumber := "20-222222" // string | Your unique Ingram Micro customer number.
	iMCountryCode := "US" // string | Two-character ISO country code.
	iMCorrelationID := "fbac82ba-cf0a-4bcf-fc03-0c5084" // string | Unique transaction number to identify each transaction across all the systems.
	iMCustomerContact := "John.Doe@reseller.com" // string | Logged in Users email address contact.
	iMSenderID := "MyCompany" // string | Unique value used to identify the sender of the transaction. (optional)
	freightRequest := *openapiclient.NewFreightRequest() // FreightRequest |  (optional)

	configuration := xi_sdk_resellers.NewConfiguration()
	apiClient := xi_sdk_resellers.NewAPIClient(configuration)
	resp, r, err := apiClient.FreightEstimateAPI.PostFreightestimate(context.Background()).IMCustomerNumber(iMCustomerNumber).IMCountryCode(iMCountryCode).IMCorrelationID(iMCorrelationID).IMCustomerContact(iMCustomerContact).IMSenderID(iMSenderID).FreightRequest(freightRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FreightEstimateAPI.PostFreightestimate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostFreightestimate`: FreightResponse
	fmt.Fprintf(os.Stdout, "Response from `FreightEstimateAPI.PostFreightestimate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostFreightestimateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iMCustomerNumber** | **string** | Your unique Ingram Micro customer number. | 
 **iMCountryCode** | **string** | Two-character ISO country code. | 
 **iMCorrelationID** | **string** | Unique transaction number to identify each transaction across all the systems. | 
 **iMCustomerContact** | **string** | Logged in Users email address contact. | 
 **iMSenderID** | **string** | Unique value used to identify the sender of the transaction. | 
 **freightRequest** | [**FreightRequest**](FreightRequest.md) |  | 

### Return type

[**FreightResponse**](FreightResponse.md)

### Authorization

[application](../README.md#application)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

