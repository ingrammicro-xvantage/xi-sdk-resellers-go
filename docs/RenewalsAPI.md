# \RenewalsAPI

All URIs are relative to *https://api.ingrammicro.com:443*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetResellersV6Renewalsdetails**](RenewalsAPI.md#GetResellersV6Renewalsdetails) | **Get** /resellers/v6/renewals/{renewalId} | Renewals Details
[**PostRenewalssearch**](RenewalsAPI.md#PostRenewalssearch) | **Post** /resellers/v6/renewals/search | Renewals Search



## GetResellersV6Renewalsdetails

> RenewalsDetailsResponse GetResellersV6Renewalsdetails(ctx, renewalId).IMCustomerNumber(iMCustomerNumber).IMCountryCode(iMCountryCode).IMCorrelationID(iMCorrelationID).IMSenderID(iMSenderID).Execute()

Renewals Details



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
	renewalId := "123456" // string | Unique Ingram renewal ID.
	iMSenderID := "MyCompany" // string | Unique value used to identify the sender of the transaction. Example: MyCompany (optional)

	configuration := xi_sdk_resellers.NewConfiguration()
	apiClient := xi_sdk_resellers.NewAPIClient(configuration)
	resp, r, err := apiClient.RenewalsAPI.GetResellersV6Renewalsdetails(context.Background(), renewalId).IMCustomerNumber(iMCustomerNumber).IMCountryCode(iMCountryCode).IMCorrelationID(iMCorrelationID).IMSenderID(iMSenderID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RenewalsAPI.GetResellersV6Renewalsdetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetResellersV6Renewalsdetails`: RenewalsDetailsResponse
	fmt.Fprintf(os.Stdout, "Response from `RenewalsAPI.GetResellersV6Renewalsdetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**renewalId** | **string** | Unique Ingram renewal ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetResellersV6RenewalsdetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iMCustomerNumber** | **string** | Your unique Ingram Micro customer number. | 
 **iMCountryCode** | **string** | Two-character ISO country code. | 
 **iMCorrelationID** | **string** | Unique transaction number to identify each transaction across all the systems. | 

 **iMSenderID** | **string** | Unique value used to identify the sender of the transaction. Example: MyCompany | 

### Return type

[**RenewalsDetailsResponse**](RenewalsDetailsResponse.md)

### Authorization

[application](../README.md#application)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostRenewalssearch

> RenewalsSearchResponse PostRenewalssearch(ctx).IMCustomerNumber(iMCustomerNumber).IMCountryCode(iMCountryCode).IMCorrelationID(iMCorrelationID).IMSenderID(iMSenderID).CustomerOrderNumber(customerOrderNumber).IngramPurchaseOrderNumber(ingramPurchaseOrderNumber).SerialNumber(serialNumber).Page(page).Size(size).Sort(sort).RenewalsSearchRequest(renewalsSearchRequest).Execute()

Renewals Search



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
	iMSenderID := "MyCompany" // string | Unique value used to identify the sender of the transaction. Example: MyCompany (optional)
	customerOrderNumber := "customerOrderNumber_example" // string | The reseller's unique PO/Order number. (optional)
	ingramPurchaseOrderNumber := "ingramPurchaseOrderNumber_example" // string | Sales order number. (optional)
	serialNumber := "serialNumber_example" // string | A serial number of the product. (optional)
	page := "page_example" // string | Number of page. (optional)
	size := "size_example" // string | The submitted pagesize, default is 25. (optional)
	sort := "sort_example" // string | Refers to the column selected to apply the sorting criteria. (optional)
	renewalsSearchRequest := *openapiclient.NewRenewalsSearchRequest() // RenewalsSearchRequest |  (optional)

	configuration := xi_sdk_resellers.NewConfiguration()
	apiClient := xi_sdk_resellers.NewAPIClient(configuration)
	resp, r, err := apiClient.RenewalsAPI.PostRenewalssearch(context.Background()).IMCustomerNumber(iMCustomerNumber).IMCountryCode(iMCountryCode).IMCorrelationID(iMCorrelationID).IMSenderID(iMSenderID).CustomerOrderNumber(customerOrderNumber).IngramPurchaseOrderNumber(ingramPurchaseOrderNumber).SerialNumber(serialNumber).Page(page).Size(size).Sort(sort).RenewalsSearchRequest(renewalsSearchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RenewalsAPI.PostRenewalssearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostRenewalssearch`: RenewalsSearchResponse
	fmt.Fprintf(os.Stdout, "Response from `RenewalsAPI.PostRenewalssearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostRenewalssearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iMCustomerNumber** | **string** | Your unique Ingram Micro customer number. | 
 **iMCountryCode** | **string** | Two-character ISO country code. | 
 **iMCorrelationID** | **string** | Unique transaction number to identify each transaction across all the systems. | 
 **iMSenderID** | **string** | Unique value used to identify the sender of the transaction. Example: MyCompany | 
 **customerOrderNumber** | **string** | The reseller&#39;s unique PO/Order number. | 
 **ingramPurchaseOrderNumber** | **string** | Sales order number. | 
 **serialNumber** | **string** | A serial number of the product. | 
 **page** | **string** | Number of page. | 
 **size** | **string** | The submitted pagesize, default is 25. | 
 **sort** | **string** | Refers to the column selected to apply the sorting criteria. | 
 **renewalsSearchRequest** | [**RenewalsSearchRequest**](RenewalsSearchRequest.md) |  | 

### Return type

[**RenewalsSearchResponse**](RenewalsSearchResponse.md)

### Authorization

[application](../README.md#application)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

