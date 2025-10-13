# \ReturnsAPI

All URIs are relative to *https://api.ingrammicro.com:443*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetResellersV6Returnsdetails**](ReturnsAPI.md#GetResellersV6Returnsdetails) | **Get** /resellers/v6/returns/{caseRequestNumber} | Returns Details
[**GetResellersV6Returnssearch**](ReturnsAPI.md#GetResellersV6Returnssearch) | **Get** /resellers/v6/returns/search | Returns Search
[**PostReturnscreate**](ReturnsAPI.md#PostReturnscreate) | **Post** /resellers/v6/returns/create | Returns Create



## GetResellersV6Returnsdetails

> ReturnsDetailsResponse GetResellersV6Returnsdetails(ctx, caseRequestNumber).IMCustomerNumber(iMCustomerNumber).IMCountryCode(iMCountryCode).IMCorrelationID(iMCorrelationID).IMSenderID(iMSenderID).Execute()

Returns Details



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
	caseRequestNumber := "12345678" // string | A unique return request number.
	iMSenderID := "MyCompany" // string | Unique value used to identify the sender of the transaction. Example: MyCompany (optional)

	configuration := xi_sdk_resellers.NewConfiguration()
	apiClient := xi_sdk_resellers.NewAPIClient(configuration)
	resp, r, err := apiClient.ReturnsAPI.GetResellersV6Returnsdetails(context.Background(), caseRequestNumber).IMCustomerNumber(iMCustomerNumber).IMCountryCode(iMCountryCode).IMCorrelationID(iMCorrelationID).IMSenderID(iMSenderID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReturnsAPI.GetResellersV6Returnsdetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetResellersV6Returnsdetails`: ReturnsDetailsResponse
	fmt.Fprintf(os.Stdout, "Response from `ReturnsAPI.GetResellersV6Returnsdetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**caseRequestNumber** | **string** | A unique return request number. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetResellersV6ReturnsdetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iMCustomerNumber** | **string** | Your unique Ingram Micro customer number. | 
 **iMCountryCode** | **string** | Two-character ISO country code. | 
 **iMCorrelationID** | **string** | Unique transaction number to identify each transaction across all the systems. | 

 **iMSenderID** | **string** | Unique value used to identify the sender of the transaction. Example: MyCompany | 

### Return type

[**ReturnsDetailsResponse**](ReturnsDetailsResponse.md)

### Authorization

[application](../README.md#application)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetResellersV6Returnssearch

> ReturnsSearchResponse GetResellersV6Returnssearch(ctx).IMCustomerNumber(iMCustomerNumber).IMCountryCode(iMCountryCode).IMCorrelationID(iMCorrelationID).IMSenderID(iMSenderID).CaseRequestNumber(caseRequestNumber).InvoiceNumber(invoiceNumber).ReturnClaimId(returnClaimId).ReferenceNumber(referenceNumber).IngramPartNumber(ingramPartNumber).VendorPartNumber(vendorPartNumber).ReturnStatusIn(returnStatusIn).ClaimStatusIn(claimStatusIn).CreatedOnBt(createdOnBt).ModifiedOnBt(modifiedOnBt).ReturnReasonIn(returnReasonIn).Page(page).Size(size).Sort(sort).SortingColumnName(sortingColumnName).Execute()

Returns Search



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
	caseRequestNumber := "caseRequestNumber_example" // string | A unique return request number. (optional)
	invoiceNumber := "invoiceNumber_example" // string | The Invoice number for the order. (optional)
	returnClaimId := "returnClaimId_example" // string | A unique return claim Id. (optional)
	referenceNumber := "referenceNumber_example" // string | The reference number for the return. (optional)
	ingramPartNumber := "ingramPartNumber_example" // string | Unique IngramMicro part number. (optional)
	vendorPartNumber := "vendorPartNumber_example" // string | The vendor's part number. (optional)
	returnStatusIn := "returnStatusIn_example" // string | Comma-separated values of pre-defined status. Open, Approved, Partially Approved, Denied, Voided. (optional)
	claimStatusIn := "claimStatusIn_example" // string | Comma-separated values of pre-defined status. Open, Approved, Partially Approved, Denied, Voided. (optional)
	createdOnBt := "createdOnBt_example" // string | The date on which the return request was created.  (optional)
	modifiedOnBt := "modifiedOnBt_example" // string | The date on which the return request was last updated. (optional)
	returnReasonIn := "returnReasonIn_example" // string | Comma separated Pre-defined value. test, (EW) Express Warehousing, (AR) Account Receivables, (BB) Buy Back, (BE) Stock Balance Exception, (BO) Bill Only, (CE) Credit Dept Use - Credit Exception, (CF) Configuration Fee, (CS ) Customer Service Discretion, (CS1) Customer Service Discretion CS Error, (DE) Defective Exception, (DF) Defective Items, (DI) Direct Credit, (DM) Damaged from Carrier, (DN) Damaged Without Product, (DT) Direct/ Special Order, (DT1) Direct Ship billed, not shipped., (FO) Freight Out, (FX) No-Scan, (IN) Incomplete, (LS) Lost Shipment, (MN) Minimum Order Fee Credit, (OS) Over Shipment, (PR) Pricing Error, (RF) Refusal Credit, (RI) Re-Invoice, (RP) Return For Repair, (RT) Return Not Credited, (RTD) RCN, (SB) Stock Balance, (SD) Sales Discretion, (SH) Incorrect Shipping And Handling, (SS) Short Shipment, (SSK) Short Ship kit, (SW) Sales Writeoff, (TE) Opened Return, (TR) Training Refund, (TX) Tax Credit, (WS) Wrong Sales Sealed, (WW) Wrong Warehouse, (FS) Warehouse Failed Serial# Capture, Latin America Vebdor Credits, Select Source, ITAD - Trade-in Credit, Withholding Tax (optional)
	page := "page_example" // string | Number of page. (optional)
	size := "size_example" // string | The submitted pagesize, default is 25 (optional)
	sort := "sort_example" // string | Refers to the column selected to apply the sorting criteria. (optional)
	sortingColumnName := "sortingColumnName_example" // string | The column name which will be sorted on. (optional)

	configuration := xi_sdk_resellers.NewConfiguration()
	apiClient := xi_sdk_resellers.NewAPIClient(configuration)
	resp, r, err := apiClient.ReturnsAPI.GetResellersV6Returnssearch(context.Background()).IMCustomerNumber(iMCustomerNumber).IMCountryCode(iMCountryCode).IMCorrelationID(iMCorrelationID).IMSenderID(iMSenderID).CaseRequestNumber(caseRequestNumber).InvoiceNumber(invoiceNumber).ReturnClaimId(returnClaimId).ReferenceNumber(referenceNumber).IngramPartNumber(ingramPartNumber).VendorPartNumber(vendorPartNumber).ReturnStatusIn(returnStatusIn).ClaimStatusIn(claimStatusIn).CreatedOnBt(createdOnBt).ModifiedOnBt(modifiedOnBt).ReturnReasonIn(returnReasonIn).Page(page).Size(size).Sort(sort).SortingColumnName(sortingColumnName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReturnsAPI.GetResellersV6Returnssearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetResellersV6Returnssearch`: ReturnsSearchResponse
	fmt.Fprintf(os.Stdout, "Response from `ReturnsAPI.GetResellersV6Returnssearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetResellersV6ReturnssearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iMCustomerNumber** | **string** | Your unique Ingram Micro customer number. | 
 **iMCountryCode** | **string** | Two-character ISO country code. | 
 **iMCorrelationID** | **string** | Unique transaction number to identify each transaction across all the systems. | 
 **iMSenderID** | **string** | Unique value used to identify the sender of the transaction. Example: MyCompany | 
 **caseRequestNumber** | **string** | A unique return request number. | 
 **invoiceNumber** | **string** | The Invoice number for the order. | 
 **returnClaimId** | **string** | A unique return claim Id. | 
 **referenceNumber** | **string** | The reference number for the return. | 
 **ingramPartNumber** | **string** | Unique IngramMicro part number. | 
 **vendorPartNumber** | **string** | The vendor&#39;s part number. | 
 **returnStatusIn** | **string** | Comma-separated values of pre-defined status. Open, Approved, Partially Approved, Denied, Voided. | 
 **claimStatusIn** | **string** | Comma-separated values of pre-defined status. Open, Approved, Partially Approved, Denied, Voided. | 
 **createdOnBt** | **string** | The date on which the return request was created.  | 
 **modifiedOnBt** | **string** | The date on which the return request was last updated. | 
 **returnReasonIn** | **string** | Comma separated Pre-defined value. test, (EW) Express Warehousing, (AR) Account Receivables, (BB) Buy Back, (BE) Stock Balance Exception, (BO) Bill Only, (CE) Credit Dept Use - Credit Exception, (CF) Configuration Fee, (CS ) Customer Service Discretion, (CS1) Customer Service Discretion CS Error, (DE) Defective Exception, (DF) Defective Items, (DI) Direct Credit, (DM) Damaged from Carrier, (DN) Damaged Without Product, (DT) Direct/ Special Order, (DT1) Direct Ship billed, not shipped., (FO) Freight Out, (FX) No-Scan, (IN) Incomplete, (LS) Lost Shipment, (MN) Minimum Order Fee Credit, (OS) Over Shipment, (PR) Pricing Error, (RF) Refusal Credit, (RI) Re-Invoice, (RP) Return For Repair, (RT) Return Not Credited, (RTD) RCN, (SB) Stock Balance, (SD) Sales Discretion, (SH) Incorrect Shipping And Handling, (SS) Short Shipment, (SSK) Short Ship kit, (SW) Sales Writeoff, (TE) Opened Return, (TR) Training Refund, (TX) Tax Credit, (WS) Wrong Sales Sealed, (WW) Wrong Warehouse, (FS) Warehouse Failed Serial# Capture, Latin America Vebdor Credits, Select Source, ITAD - Trade-in Credit, Withholding Tax | 
 **page** | **string** | Number of page. | 
 **size** | **string** | The submitted pagesize, default is 25 | 
 **sort** | **string** | Refers to the column selected to apply the sorting criteria. | 
 **sortingColumnName** | **string** | The column name which will be sorted on. | 

### Return type

[**ReturnsSearchResponse**](ReturnsSearchResponse.md)

### Authorization

[application](../README.md#application)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostReturnscreate

> ReturnsCreateResponse PostReturnscreate(ctx).IMCustomerNumber(iMCustomerNumber).IMCountryCode(iMCountryCode).IMCorrelationID(iMCorrelationID).IMSenderID(iMSenderID).ReturnsCreateRequest(returnsCreateRequest).Execute()

Returns Create



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
	returnsCreateRequest := *openapiclient.NewReturnsCreateRequest() // ReturnsCreateRequest |  (optional)

	configuration := xi_sdk_resellers.NewConfiguration()
	apiClient := xi_sdk_resellers.NewAPIClient(configuration)
	resp, r, err := apiClient.ReturnsAPI.PostReturnscreate(context.Background()).IMCustomerNumber(iMCustomerNumber).IMCountryCode(iMCountryCode).IMCorrelationID(iMCorrelationID).IMSenderID(iMSenderID).ReturnsCreateRequest(returnsCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReturnsAPI.PostReturnscreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostReturnscreate`: ReturnsCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `ReturnsAPI.PostReturnscreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostReturnscreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iMCustomerNumber** | **string** | Your unique Ingram Micro customer number. | 
 **iMCountryCode** | **string** | Two-character ISO country code. | 
 **iMCorrelationID** | **string** | Unique transaction number to identify each transaction across all the systems. | 
 **iMSenderID** | **string** | Unique value used to identify the sender of the transaction. Example: MyCompany | 
 **returnsCreateRequest** | [**ReturnsCreateRequest**](ReturnsCreateRequest.md) |  | 

### Return type

[**ReturnsCreateResponse**](ReturnsCreateResponse.md)

### Authorization

[application](../README.md#application)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

