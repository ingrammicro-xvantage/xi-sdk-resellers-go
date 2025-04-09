# \OrdersAPI

All URIs are relative to *https://api.ingrammicro.com:443*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteOrdercancel**](OrdersAPI.md#DeleteOrdercancel) | **Delete** /resellers/v6/orders/{OrderNumber} | Cancel your Order
[**GetOrderdetailsV61**](OrdersAPI.md#GetOrderdetailsV61) | **Get** /resellers/v6.1/orders/{ordernumber} | Get Order Details v6.1
[**GetResellersV6Ordersearch**](OrdersAPI.md#GetResellersV6Ordersearch) | **Get** /resellers/v6/orders/search | Search your Orders
[**PostCreateorderV6**](OrdersAPI.md#PostCreateorderV6) | **Post** /resellers/v6/orders | Create your Order
[**PostCreateorderV7**](OrdersAPI.md#PostCreateorderV7) | **Post** /resellers/v7/orders | Create your Order v7
[**PutOrdermodify**](OrdersAPI.md#PutOrdermodify) | **Put** /resellers/v6/orders/{orderNumber} | Modify your Order



## DeleteOrdercancel

> DeleteOrdercancel(ctx, orderNumber).IMCustomerNumber(iMCustomerNumber).IMCountryCode(iMCountryCode).IMCorrelationID(iMCorrelationID).RegionCode(regionCode).IMSenderID(iMSenderID).Execute()

Cancel your Order



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
	orderNumber := "20-RD128" // string | Ingram Micro sales order number.
	iMCustomerNumber := "20-222222" // string | Your unique Ingram Micro customer number.
	iMCountryCode := "US" // string | Two-character ISO country code.
	iMCorrelationID := "fbac82ba-cf0a-4bcf-fc03-0c5084" // string | Unique transaction number to identify each transaction accross all the systems.
	regionCode := "CS" // string | Region code for sandbox testing - Not for use in production. (optional)
	iMSenderID := "MyCompany" // string | Unique value used to identify the sender of the transaction. Example: MyCompany (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrdersAPI.DeleteOrdercancel(context.Background(), orderNumber).IMCustomerNumber(iMCustomerNumber).IMCountryCode(iMCountryCode).IMCorrelationID(iMCorrelationID).RegionCode(regionCode).IMSenderID(iMSenderID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAPI.DeleteOrdercancel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderNumber** | **string** | Ingram Micro sales order number. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrdercancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iMCustomerNumber** | **string** | Your unique Ingram Micro customer number. | 
 **iMCountryCode** | **string** | Two-character ISO country code. | 
 **iMCorrelationID** | **string** | Unique transaction number to identify each transaction accross all the systems. | 
 **regionCode** | **string** | Region code for sandbox testing - Not for use in production. | 
 **iMSenderID** | **string** | Unique value used to identify the sender of the transaction. Example: MyCompany | 

### Return type

 (empty response body)

### Authorization

[application](../README.md#application)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrderdetailsV61

> OrderDetailB2B GetOrderdetailsV61(ctx, ordernumber).IMCustomerNumber(iMCustomerNumber).IMCountryCode(iMCountryCode).IMCorrelationID(iMCorrelationID).IMSenderID(iMSenderID).IngramOrderDate(ingramOrderDate).VendorNumber(vendorNumber).SimulateStatus(simulateStatus).IsIml(isIml).RegionCode(regionCode).Execute()

Get Order Details v6.1



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	ordernumber := "20-RD3QV" // string | The Ingram Micro sales order number.
	iMCustomerNumber := "20-222222" // string | Your unique Ingram Micro customer number.
	iMCountryCode := "US" // string | Two-character ISO country code.
	iMCorrelationID := "fbac82ba-cf0a-4bcf-fc03-0c5084" // string | Unique transaction number to identify each transaction accross all the systems.
	iMSenderID := "MyCompany" // string | Unique value used to identify the sender of the transaction. Example: MyCompany. (optional)
	ingramOrderDate := time.Now() // string | The date and time in UTC format that the order was created. (optional)
	vendorNumber := "vendorNumber_example" // string | Vendor Number. (optional)
	simulateStatus := "simulateStatus_example" // string | Order response for various order statuses. Not for use in production. (optional)
	isIml := true // bool | True/False only for IML customers. (optional)
	regionCode := "regionCode_example" // string | Region code for sandbox testing - Not for use in production. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrdersAPI.GetOrderdetailsV61(context.Background(), ordernumber).IMCustomerNumber(iMCustomerNumber).IMCountryCode(iMCountryCode).IMCorrelationID(iMCorrelationID).IMSenderID(iMSenderID).IngramOrderDate(ingramOrderDate).VendorNumber(vendorNumber).SimulateStatus(simulateStatus).IsIml(isIml).RegionCode(regionCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAPI.GetOrderdetailsV61``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrderdetailsV61`: OrderDetailB2B
	fmt.Fprintf(os.Stdout, "Response from `OrdersAPI.GetOrderdetailsV61`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ordernumber** | **string** | The Ingram Micro sales order number. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrderdetailsV61Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iMCustomerNumber** | **string** | Your unique Ingram Micro customer number. | 
 **iMCountryCode** | **string** | Two-character ISO country code. | 
 **iMCorrelationID** | **string** | Unique transaction number to identify each transaction accross all the systems. | 
 **iMSenderID** | **string** | Unique value used to identify the sender of the transaction. Example: MyCompany. | 
 **ingramOrderDate** | **string** | The date and time in UTC format that the order was created. | 
 **vendorNumber** | **string** | Vendor Number. | 
 **simulateStatus** | **string** | Order response for various order statuses. Not for use in production. | 
 **isIml** | **bool** | True/False only for IML customers. | 
 **regionCode** | **string** | Region code for sandbox testing - Not for use in production. | 

### Return type

[**OrderDetailB2B**](OrderDetailB2B.md)

### Authorization

[application](../README.md#application)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetResellersV6Ordersearch

> OrderSearchResponse GetResellersV6Ordersearch(ctx).IMCustomerNumber(iMCustomerNumber).IMCountryCode(iMCountryCode).IMCorrelationID(iMCorrelationID).IngramOrderNumber(ingramOrderNumber).OrderStatus(orderStatus).OrderStatusIn(orderStatusIn).IngramOrderDate(ingramOrderDate).IngramOrderDateBt(ingramOrderDateBt).IMSenderID(iMSenderID).CustomerOrderNumber(customerOrderNumber).PageSize(pageSize).PageNumber(pageNumber).EndCustomerOrderNumber(endCustomerOrderNumber).InvoiceDateBt(invoiceDateBt).ShipDateBt(shipDateBt).DeliveryDateBt(deliveryDateBt).IngramPartNumber(ingramPartNumber).VendorPartNumber(vendorPartNumber).SerialNumber(serialNumber).TrackingNumber(trackingNumber).VendorName(vendorName).SpecialBidNumber(specialBidNumber).Execute()

Search your Orders



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
	iMCorrelationID := "fbac82ba-cf0a-4bcf-fc03-0c5084" // string | Unique transaction number to identify each transaction accross all the systems.
	ingramOrderNumber := "ingramOrderNumber_example" // string | The Ingram Micro order number. (optional)
	orderStatus := "orderStatus_example" // string | Ingram Micro order status. (optional)
	orderStatusIn := []string{"Inner_example"} // []string | Ingram Micro order status(can use it for multiple entries). (optional)
	ingramOrderDate := "2021-04-23" // string | Search by Order date(yyyy-MM-dd). (optional)
	ingramOrderDateBt := []string{"Inner_example"} // []string | Search with the start and end date(only 2 entries allowed). (optional)
	iMSenderID := "MyCompany" // string | Unique value used to identify the sender of the transaction. Example: MyCompany (optional)
	customerOrderNumber := "customerOrderNumber_example" // string | Search using your PO/Order number. (optional)
	pageSize := int32(56) // int32 | The number of records required in the call - max records 100 per page. (optional)
	pageNumber := int32(56) // int32 | The page number reference. (optional)
	endCustomerOrderNumber := "endCustomerOrderNumber_example" // string | End customer/user purchase order number. (optional)
	invoiceDateBt := []string{"Inner_example"} // []string | Invoice date of order, search with the start and end date(only 2 entries allowed).*Currently, this feature is not available in Australia. (optional)
	shipDateBt := []string{"Inner_example"} // []string | Shipment date of order, search with the start and end date(only 2 entries allowed). (optional)
	deliveryDateBt := []string{"Inner_example"} // []string | The delivery date of the order, search with the start and end date(only 2 entries allowed).*Currently, this feature is not available in Australia (optional)
	ingramPartNumber := "ingramPartNumber_example" // string | Ingram Micro unique part number for the product. (optional)
	vendorPartNumber := "vendorPartNumber_example" // string | Vendor’s part number for the product. (optional)
	serialNumber := "serialNumber_example" // string | A serial number of the product. (optional)
	trackingNumber := "trackingNumber_example" // string | The tracking number of the order.*Currently, this feature is not available in Australia (optional)
	vendorName := "vendorName_example" // string | Name of the vendor. (optional)
	specialBidNumber := "specialBidNumber_example" // string | The bid number provided to the reseller by the vendor for special pricing and discounts. Line-level bid numbers take precedence over header-level bid numbers.*Currently, this feature is not available in Australia (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrdersAPI.GetResellersV6Ordersearch(context.Background()).IMCustomerNumber(iMCustomerNumber).IMCountryCode(iMCountryCode).IMCorrelationID(iMCorrelationID).IngramOrderNumber(ingramOrderNumber).OrderStatus(orderStatus).OrderStatusIn(orderStatusIn).IngramOrderDate(ingramOrderDate).IngramOrderDateBt(ingramOrderDateBt).IMSenderID(iMSenderID).CustomerOrderNumber(customerOrderNumber).PageSize(pageSize).PageNumber(pageNumber).EndCustomerOrderNumber(endCustomerOrderNumber).InvoiceDateBt(invoiceDateBt).ShipDateBt(shipDateBt).DeliveryDateBt(deliveryDateBt).IngramPartNumber(ingramPartNumber).VendorPartNumber(vendorPartNumber).SerialNumber(serialNumber).TrackingNumber(trackingNumber).VendorName(vendorName).SpecialBidNumber(specialBidNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAPI.GetResellersV6Ordersearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetResellersV6Ordersearch`: OrderSearchResponse
	fmt.Fprintf(os.Stdout, "Response from `OrdersAPI.GetResellersV6Ordersearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetResellersV6OrdersearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iMCustomerNumber** | **string** | Your unique Ingram Micro customer number. | 
 **iMCountryCode** | **string** | Two-character ISO country code. | 
 **iMCorrelationID** | **string** | Unique transaction number to identify each transaction accross all the systems. | 
 **ingramOrderNumber** | **string** | The Ingram Micro order number. | 
 **orderStatus** | **string** | Ingram Micro order status. | 
 **orderStatusIn** | **[]string** | Ingram Micro order status(can use it for multiple entries). | 
 **ingramOrderDate** | **string** | Search by Order date(yyyy-MM-dd). | 
 **ingramOrderDateBt** | **[]string** | Search with the start and end date(only 2 entries allowed). | 
 **iMSenderID** | **string** | Unique value used to identify the sender of the transaction. Example: MyCompany | 
 **customerOrderNumber** | **string** | Search using your PO/Order number. | 
 **pageSize** | **int32** | The number of records required in the call - max records 100 per page. | 
 **pageNumber** | **int32** | The page number reference. | 
 **endCustomerOrderNumber** | **string** | End customer/user purchase order number. | 
 **invoiceDateBt** | **[]string** | Invoice date of order, search with the start and end date(only 2 entries allowed).*Currently, this feature is not available in Australia. | 
 **shipDateBt** | **[]string** | Shipment date of order, search with the start and end date(only 2 entries allowed). | 
 **deliveryDateBt** | **[]string** | The delivery date of the order, search with the start and end date(only 2 entries allowed).*Currently, this feature is not available in Australia | 
 **ingramPartNumber** | **string** | Ingram Micro unique part number for the product. | 
 **vendorPartNumber** | **string** | Vendor’s part number for the product. | 
 **serialNumber** | **string** | A serial number of the product. | 
 **trackingNumber** | **string** | The tracking number of the order.*Currently, this feature is not available in Australia | 
 **vendorName** | **string** | Name of the vendor. | 
 **specialBidNumber** | **string** | The bid number provided to the reseller by the vendor for special pricing and discounts. Line-level bid numbers take precedence over header-level bid numbers.*Currently, this feature is not available in Australia | 

### Return type

[**OrderSearchResponse**](OrderSearchResponse.md)

### Authorization

[application](../README.md#application)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCreateorderV6

> OrderCreateResponse PostCreateorderV6(ctx).IMCustomerNumber(iMCustomerNumber).IMCountryCode(iMCountryCode).IMCorrelationID(iMCorrelationID).OrderCreateRequest(orderCreateRequest).IMSenderID(iMSenderID).Execute()

Create your Order



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
	iMCorrelationID := "fbac82ba-cf0a-4bcf-fc03-0c5084" // string | Unique transaction number to identify each transaction accross all the systems.
	orderCreateRequest := *openapiclient.NewOrderCreateRequest("CustomerOrderNumber_example") // OrderCreateRequest | 
	iMSenderID := "MyCompany" // string | Unique value used to identify the sender of the transaction. Example: MyCompany (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrdersAPI.PostCreateorderV6(context.Background()).IMCustomerNumber(iMCustomerNumber).IMCountryCode(iMCountryCode).IMCorrelationID(iMCorrelationID).OrderCreateRequest(orderCreateRequest).IMSenderID(iMSenderID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAPI.PostCreateorderV6``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCreateorderV6`: OrderCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `OrdersAPI.PostCreateorderV6`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCreateorderV6Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iMCustomerNumber** | **string** | Your unique Ingram Micro customer number. | 
 **iMCountryCode** | **string** | Two-character ISO country code. | 
 **iMCorrelationID** | **string** | Unique transaction number to identify each transaction accross all the systems. | 
 **orderCreateRequest** | [**OrderCreateRequest**](OrderCreateRequest.md) |  | 
 **iMSenderID** | **string** | Unique value used to identify the sender of the transaction. Example: MyCompany | 

### Return type

[**OrderCreateResponse**](OrderCreateResponse.md)

### Authorization

[application](../README.md#application)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCreateorderV7

> OrderCreateV7Response201 PostCreateorderV7(ctx).IMCustomerNumber(iMCustomerNumber).IMCountryCode(iMCountryCode).IMCorrelationID(iMCorrelationID).OrderCreateV7Request(orderCreateV7Request).IMSenderID(iMSenderID).Execute()

Create your Order v7



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
	orderCreateV7Request := *openapiclient.NewOrderCreateV7Request() // OrderCreateV7Request | 
	iMSenderID := "MyCompany" // string | Unique value used to identify the sender of the transaction. Example: MyCompany (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrdersAPI.PostCreateorderV7(context.Background()).IMCustomerNumber(iMCustomerNumber).IMCountryCode(iMCountryCode).IMCorrelationID(iMCorrelationID).OrderCreateV7Request(orderCreateV7Request).IMSenderID(iMSenderID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAPI.PostCreateorderV7``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCreateorderV7`: OrderCreateV7Response201
	fmt.Fprintf(os.Stdout, "Response from `OrdersAPI.PostCreateorderV7`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCreateorderV7Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iMCustomerNumber** | **string** | Your unique Ingram Micro customer number. | 
 **iMCountryCode** | **string** | Two-character ISO country code. | 
 **iMCorrelationID** | **string** | Unique transaction number to identify each transaction across all the systems. | 
 **orderCreateV7Request** | [**OrderCreateV7Request**](OrderCreateV7Request.md) |  | 
 **iMSenderID** | **string** | Unique value used to identify the sender of the transaction. Example: MyCompany | 

### Return type

[**OrderCreateV7Response201**](OrderCreateV7Response201.md)

### Authorization

[application](../README.md#application)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutOrdermodify

> OrderModifyResponse PutOrdermodify(ctx, orderNumber).IMCustomerNumber(iMCustomerNumber).IMCountryCode(iMCountryCode).IMCorrelationID(iMCorrelationID).OrderModifyRequest(orderModifyRequest).ActionCode(actionCode).RegionCode(regionCode).IMSenderID(iMSenderID).Execute()

Modify your Order



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
	orderNumber := "20-RC1RD" // string | Ingram sales order number.
	iMCustomerNumber := "20-222222" // string | Your unique Ingram Micro customer number.
	iMCountryCode := "US" // string | Two-character ISO country code.
	iMCorrelationID := "fbac82ba-cf0a-4bcf-fc03-0c5084" // string | Unique transaction number to identify each transaction across all the systems.
	orderModifyRequest := *openapiclient.NewOrderModifyRequest() // OrderModifyRequest | 
	actionCode := "release" // string | Action code to be used for order release. (optional)
	regionCode := "CS" // string | Region code paramter to be used only for order release functionality.Region code is only for sandbox not for production (optional)
	iMSenderID := "MyCompany" // string | Unique value used to identify the sender of the transaction. Example: MyCompany (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrdersAPI.PutOrdermodify(context.Background(), orderNumber).IMCustomerNumber(iMCustomerNumber).IMCountryCode(iMCountryCode).IMCorrelationID(iMCorrelationID).OrderModifyRequest(orderModifyRequest).ActionCode(actionCode).RegionCode(regionCode).IMSenderID(iMSenderID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAPI.PutOrdermodify``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutOrdermodify`: OrderModifyResponse
	fmt.Fprintf(os.Stdout, "Response from `OrdersAPI.PutOrdermodify`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderNumber** | **string** | Ingram sales order number. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutOrdermodifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iMCustomerNumber** | **string** | Your unique Ingram Micro customer number. | 
 **iMCountryCode** | **string** | Two-character ISO country code. | 
 **iMCorrelationID** | **string** | Unique transaction number to identify each transaction across all the systems. | 
 **orderModifyRequest** | [**OrderModifyRequest**](OrderModifyRequest.md) |  | 
 **actionCode** | **string** | Action code to be used for order release. | 
 **regionCode** | **string** | Region code paramter to be used only for order release functionality.Region code is only for sandbox not for production | 
 **iMSenderID** | **string** | Unique value used to identify the sender of the transaction. Example: MyCompany | 

### Return type

[**OrderModifyResponse**](OrderModifyResponse.md)

### Authorization

[application](../README.md#application)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

