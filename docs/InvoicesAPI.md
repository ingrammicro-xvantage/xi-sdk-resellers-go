# \InvoicesAPI

All URIs are relative to *https://api.ingrammicro.com:443*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetInvoicedetailsV61**](InvoicesAPI.md#GetInvoicedetailsV61) | **Get** /resellers/v6.1/invoices/{invoiceNumber} | Get Invoice Details v6.1
[**GetResellersV6Invoicesearch**](InvoicesAPI.md#GetResellersV6Invoicesearch) | **Get** /resellers/v6/invoices | Search your invoice



## GetInvoicedetailsV61

> InvoiceDetailsv61Response GetInvoicedetailsV61(ctx, invoiceNumber).IMCustomerNumber(iMCustomerNumber).IMCountryCode(iMCountryCode).IMCorrelationID(iMCorrelationID).IMApplicationID(iMApplicationID).CustomerType(customerType).IncludeSerialNumbers(includeSerialNumbers).Execute()

Get Invoice Details v6.1



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
	invoiceNumber := "335238411" // string | The Ingram Micro invoice number.
	iMCustomerNumber := "20-222222" // string | Your unique Ingram Micro customer number.
	iMCountryCode := "US" // string | Two-character ISO country code.
	iMCorrelationID := "fbac82ba-cf0a-4bcf-fc03-0c5084" // string | Unique transaction number to identify each transaction across all the systems.
	iMApplicationID := "MyCompany" // string | Unique value used to identify the sender of the transaction. Example: MyCompany.
	customerType := "invoice" // string | it should be invoice or order (optional)
	includeSerialNumbers := false // bool | if serial in the response send as true or else false (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoicesAPI.GetInvoicedetailsV61(context.Background(), invoiceNumber).IMCustomerNumber(iMCustomerNumber).IMCountryCode(iMCountryCode).IMCorrelationID(iMCorrelationID).IMApplicationID(iMApplicationID).CustomerType(customerType).IncludeSerialNumbers(includeSerialNumbers).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.GetInvoicedetailsV61``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInvoicedetailsV61`: InvoiceDetailsv61Response
	fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.GetInvoicedetailsV61`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceNumber** | **string** | The Ingram Micro invoice number. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInvoicedetailsV61Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iMCustomerNumber** | **string** | Your unique Ingram Micro customer number. | 
 **iMCountryCode** | **string** | Two-character ISO country code. | 
 **iMCorrelationID** | **string** | Unique transaction number to identify each transaction across all the systems. | 
 **iMApplicationID** | **string** | Unique value used to identify the sender of the transaction. Example: MyCompany. | 
 **customerType** | **string** | it should be invoice or order | 
 **includeSerialNumbers** | **bool** | if serial in the response send as true or else false | 

### Return type

[**InvoiceDetailsv61Response**](InvoiceDetailsv61Response.md)

### Authorization

[application](../README.md#application)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetResellersV6Invoicesearch

> InvoiceSearchResponse GetResellersV6Invoicesearch(ctx).IMApplicationID(iMApplicationID).IMCustomerNumber(iMCustomerNumber).IMCountryCode(iMCountryCode).IMCorrelationID(iMCorrelationID).PaymentTermsNetDate(paymentTermsNetDate).InvoiceDate(invoiceDate).InvoiceDueDate(invoiceDueDate).OrderDate(orderDate).OrderFromDate(orderFromDate).OrderToDate(orderToDate).OrderNumber(orderNumber).DeliveryNumber(deliveryNumber).InvoiceNumber(invoiceNumber).InvoiceStatus(invoiceStatus).InvoiceType(invoiceType).CustomerOrderNumber(customerOrderNumber).EndCustomerOrderNumber(endCustomerOrderNumber).SpecialBidNumber(specialBidNumber).InvoiceFromDueDate(invoiceFromDueDate).InvoiceToDueDate(invoiceToDueDate).InvoiceFromDate(invoiceFromDate).InvoiceToDate(invoiceToDate).PageSize(pageSize).PageNumber(pageNumber).Orderby(orderby).Direction(direction).SerialNumber(serialNumber).Execute()

Search your invoice



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
	iMApplicationID := "MyCompany" // string | Unique value used to identify the sender of the transaction. Example: MyCompany
	iMCustomerNumber := "20-222222" // string | Your unique Ingram Micro customer number.
	iMCountryCode := "US" // string | Two-character ISO country code.
	iMCorrelationID := "fbac82ba-cf0a-4bcf-fc03-0c5084" // string | Unique transaction number to identify each transaction across all the systems.
	paymentTermsNetDate := "2021-04-23" // string | Search by payment terms net date(yyyy-MM-dd). (optional)
	invoiceDate := "2021-04-23" // string | Search by invoice date(yyyy-MM-dd). (optional)
	invoiceDueDate := "2021-04-23" // string | Search by invoice date from(yyyy-MM-dd). (optional)
	orderDate := "2021-04-23" // string | Search by OrderDate date(yyyy-MM-dd). (optional)
	orderFromDate := "2021-04-23" // string | Search by OrderFromDate date(yyyy-MM-dd). (optional)
	orderToDate := "2021-04-23" // string | Search by OrderToDate date(yyyy-MM-dd). (optional)
	orderNumber := "2021-04-23" // string | Search by order number (optional)
	deliveryNumber := "335238411" // string | Search by delivery number. (optional)
	invoiceNumber := "invoiceNumber_example" // string | The Ingram Micro invoice number. (optional)
	invoiceStatus := "invoiceStatus_example" // string | Ingram Micro invoice status. (optional)
	invoiceType := "invoiceType_example" // string | Ingram Micro InvoiceType. (optional)
	customerOrderNumber := "customerOrderNumber_example" // string | Ingram Micro CustomerOrderNumber. (optional)
	endCustomerOrderNumber := "endCustomerOrderNumber_example" // string | Ingram Micro EndCustomerOrderNumber. (optional)
	specialBidNumber := "specialBidNumber_example" // string | Ingram Micro SpecialBidNumber. (optional)
	invoiceFromDueDate := "2021-04-23" // string | Search by invoice due date from(yyyy-MM-dd). (optional)
	invoiceToDueDate := "2021-04-23" // string | Search by invoice due date to(yyyy-MM-dd). (optional)
	invoiceFromDate := []string{"2021-04-23"} // []string | Search by invoice date from(yyyy-MM-dd). (optional)
	invoiceToDate := []string{"2021-04-23"} // []string | Search by invoice date To(yyyy-MM-dd). (optional)
	pageSize := int32(56) // int32 | Number of records required in the call - max records 100 per page. (optional)
	pageNumber := int32(56) // int32 | The page number reference. (optional)
	orderby := "InvoiceDate" // string | Column name with which we want to sort. (optional)
	direction := "desc" // string | asc or desc , along with orderby column result set will be sorted. (optional)
	serialNumber := "serialNumber_example" // string | Serial number of the product. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoicesAPI.GetResellersV6Invoicesearch(context.Background()).IMApplicationID(iMApplicationID).IMCustomerNumber(iMCustomerNumber).IMCountryCode(iMCountryCode).IMCorrelationID(iMCorrelationID).PaymentTermsNetDate(paymentTermsNetDate).InvoiceDate(invoiceDate).InvoiceDueDate(invoiceDueDate).OrderDate(orderDate).OrderFromDate(orderFromDate).OrderToDate(orderToDate).OrderNumber(orderNumber).DeliveryNumber(deliveryNumber).InvoiceNumber(invoiceNumber).InvoiceStatus(invoiceStatus).InvoiceType(invoiceType).CustomerOrderNumber(customerOrderNumber).EndCustomerOrderNumber(endCustomerOrderNumber).SpecialBidNumber(specialBidNumber).InvoiceFromDueDate(invoiceFromDueDate).InvoiceToDueDate(invoiceToDueDate).InvoiceFromDate(invoiceFromDate).InvoiceToDate(invoiceToDate).PageSize(pageSize).PageNumber(pageNumber).Orderby(orderby).Direction(direction).SerialNumber(serialNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.GetResellersV6Invoicesearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetResellersV6Invoicesearch`: InvoiceSearchResponse
	fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.GetResellersV6Invoicesearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetResellersV6InvoicesearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iMApplicationID** | **string** | Unique value used to identify the sender of the transaction. Example: MyCompany | 
 **iMCustomerNumber** | **string** | Your unique Ingram Micro customer number. | 
 **iMCountryCode** | **string** | Two-character ISO country code. | 
 **iMCorrelationID** | **string** | Unique transaction number to identify each transaction across all the systems. | 
 **paymentTermsNetDate** | **string** | Search by payment terms net date(yyyy-MM-dd). | 
 **invoiceDate** | **string** | Search by invoice date(yyyy-MM-dd). | 
 **invoiceDueDate** | **string** | Search by invoice date from(yyyy-MM-dd). | 
 **orderDate** | **string** | Search by OrderDate date(yyyy-MM-dd). | 
 **orderFromDate** | **string** | Search by OrderFromDate date(yyyy-MM-dd). | 
 **orderToDate** | **string** | Search by OrderToDate date(yyyy-MM-dd). | 
 **orderNumber** | **string** | Search by order number | 
 **deliveryNumber** | **string** | Search by delivery number. | 
 **invoiceNumber** | **string** | The Ingram Micro invoice number. | 
 **invoiceStatus** | **string** | Ingram Micro invoice status. | 
 **invoiceType** | **string** | Ingram Micro InvoiceType. | 
 **customerOrderNumber** | **string** | Ingram Micro CustomerOrderNumber. | 
 **endCustomerOrderNumber** | **string** | Ingram Micro EndCustomerOrderNumber. | 
 **specialBidNumber** | **string** | Ingram Micro SpecialBidNumber. | 
 **invoiceFromDueDate** | **string** | Search by invoice due date from(yyyy-MM-dd). | 
 **invoiceToDueDate** | **string** | Search by invoice due date to(yyyy-MM-dd). | 
 **invoiceFromDate** | **[]string** | Search by invoice date from(yyyy-MM-dd). | 
 **invoiceToDate** | **[]string** | Search by invoice date To(yyyy-MM-dd). | 
 **pageSize** | **int32** | Number of records required in the call - max records 100 per page. | 
 **pageNumber** | **int32** | The page number reference. | 
 **orderby** | **string** | Column name with which we want to sort. | 
 **direction** | **string** | asc or desc , along with orderby column result set will be sorted. | 
 **serialNumber** | **string** | Serial number of the product. | 

### Return type

[**InvoiceSearchResponse**](InvoiceSearchResponse.md)

### Authorization

[application](../README.md#application)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

