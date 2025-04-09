# \DealsAPI

All URIs are relative to *https://api.ingrammicro.com:443*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetResellersV6Dealsdetails**](DealsAPI.md#GetResellersV6Dealsdetails) | **Get** /resellers/v6/deals/{dealId} | Deals Details
[**GetResellersV6Dealssearch**](DealsAPI.md#GetResellersV6Dealssearch) | **Get** /resellers/v6/deals/search | Deals Search



## GetResellersV6Dealsdetails

> DealsDetailsResponse GetResellersV6Dealsdetails(ctx, dealId).IMCustomerNumber(iMCustomerNumber).IMCountryCode(iMCountryCode).IMCorrelationID(iMCorrelationID).IMApplicationId(iMApplicationId).Execute()

Deals Details



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	iMCustomerNumber := "20-222222" // string | Your unique Ingram Micro customer number.
	iMCountryCode := "US" // string | Two-character ISO country code.
	iMCorrelationID := "fbac82ba-cf0a-4bcf-fc03-0c5084" // string | Unique transaction number to identify each transaction across all the systems.
	iMApplicationId := "MyCompany" // string | Unique value used to identify the sender of the transaction. Example: MyCompany
	dealId := "12345678" // string | Unique deal ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DealsAPI.GetResellersV6Dealsdetails(context.Background(), dealId).IMCustomerNumber(iMCustomerNumber).IMCountryCode(iMCountryCode).IMCorrelationID(iMCorrelationID).IMApplicationId(iMApplicationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DealsAPI.GetResellersV6Dealsdetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetResellersV6Dealsdetails`: DealsDetailsResponse
	fmt.Fprintf(os.Stdout, "Response from `DealsAPI.GetResellersV6Dealsdetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dealId** | **string** | Unique deal ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetResellersV6DealsdetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iMCustomerNumber** | **string** | Your unique Ingram Micro customer number. | 
 **iMCountryCode** | **string** | Two-character ISO country code. | 
 **iMCorrelationID** | **string** | Unique transaction number to identify each transaction across all the systems. | 
 **iMApplicationId** | **string** | Unique value used to identify the sender of the transaction. Example: MyCompany | 


### Return type

[**DealsDetailsResponse**](DealsDetailsResponse.md)

### Authorization

[application](../README.md#application)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetResellersV6Dealssearch

> DealsSearchResponse GetResellersV6Dealssearch(ctx).IMCustomerNumber(iMCustomerNumber).IMCountryCode(iMCountryCode).IMCorrelationID(iMCorrelationID).IMSenderID(iMSenderID).EndUser(endUser).Vendor(vendor).DealId(dealId).Size(size).Page(page).Execute()

Deals Search



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	iMCustomerNumber := "20-222222" // string | Your unique Ingram Micro customer number.
	iMCountryCode := "US" // string | Two-character ISO country code.
	iMCorrelationID := "fbac82ba-cf0a-4bcf-fc03-0c5084" // string | Unique transaction number to identify each transaction across all the systems.
	iMSenderID := "MyCompany" // string | Unique value used to identify the sender of the transaction. Example: MyCompany (optional)
	endUser := "EnduserCompany" // string | The end user/customer's name. (optional)
	vendor := "Cisco" // string | The vendor's name. (optional)
	dealId := "12345678" // string | Deal/Special bid number. (optional)
	size := int32(56) // int32 | The number of records required in the call - max records 100 per page. (optional)
	page := int32(56) // int32 | The page number reference. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DealsAPI.GetResellersV6Dealssearch(context.Background()).IMCustomerNumber(iMCustomerNumber).IMCountryCode(iMCountryCode).IMCorrelationID(iMCorrelationID).IMSenderID(iMSenderID).EndUser(endUser).Vendor(vendor).DealId(dealId).Size(size).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DealsAPI.GetResellersV6Dealssearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetResellersV6Dealssearch`: DealsSearchResponse
	fmt.Fprintf(os.Stdout, "Response from `DealsAPI.GetResellersV6Dealssearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetResellersV6DealssearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iMCustomerNumber** | **string** | Your unique Ingram Micro customer number. | 
 **iMCountryCode** | **string** | Two-character ISO country code. | 
 **iMCorrelationID** | **string** | Unique transaction number to identify each transaction across all the systems. | 
 **iMSenderID** | **string** | Unique value used to identify the sender of the transaction. Example: MyCompany | 
 **endUser** | **string** | The end user/customer&#39;s name. | 
 **vendor** | **string** | The vendor&#39;s name. | 
 **dealId** | **string** | Deal/Special bid number. | 
 **size** | **int32** | The number of records required in the call - max records 100 per page. | 
 **page** | **int32** | The page number reference. | 

### Return type

[**DealsSearchResponse**](DealsSearchResponse.md)

### Authorization

[application](../README.md#application)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

