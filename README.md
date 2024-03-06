# Go API client for xi_sdk_resellers

For resellers seeking to innovate with Ingram Micro's API solutions, automate your eCommerce experience with our array of APIs and webhooks to craft a seamless journey for your customers.

## Installation
To install the package use:
```sh
go get -u github.com/ingrammicro-xvantage/xi-sdk-resellers-go

```
Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```


Put the package under your project folder and add the following in import:

```go
import xi_sdk_resellers "github.com/ingrammicro-xvantage/xi-sdk-resellers-go"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `xi_sdk_resellers.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), xi_sdk_resellers.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `xi_sdk_resellers.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), xi_sdk_resellers.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `xi_sdk_resellers.ContextOperationServerIndices` and `xi_sdk_resellers.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), xi_sdk_resellers.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), xi_sdk_resellers.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```
## Getting Started

Quickstart on creating an application can be found [here](getting-started.md)

## Documentation for API Endpoints

All URIs are relative to *https://api.ingrammicro.com:443*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AccesstokenAPI* | [**GetAccesstoken**](docs/AccesstokenAPI.md#getaccesstoken) | **Get** /oauth/oauth20/token | Accesstoken
*DealsAPI* | [**GetResellersV6Dealsdetails**](docs/DealsAPI.md#getresellersv6dealsdetails) | **Get** /resellers/v6/deals/{dealId} | Deals Details
*DealsAPI* | [**GetResellersV6Dealssearch**](docs/DealsAPI.md#getresellersv6dealssearch) | **Get** /resellers/v6/deals/search | Deals Search
*FreightEstimateAPI* | [**PostFreightestimate**](docs/FreightEstimateAPI.md#postfreightestimate) | **Post** /resellers/v6/freightestimate | Freight Estimate
*InvoicesAPI* | [**GetInvoicedetailsV61**](docs/InvoicesAPI.md#getinvoicedetailsv61) | **Get** /resellers/v6.1/invoices/{invoiceNumber} | Get Invoice Details v6.1
*InvoicesAPI* | [**GetResellersV6Invoicesearch**](docs/InvoicesAPI.md#getresellersv6invoicesearch) | **Get** /resellers/v6/invoices | Search your invoice
*OrderStatusAPI* | [**ResellersV1WebhooksOrderstatuseventPost**](docs/OrderStatusAPI.md#resellersv1webhooksorderstatuseventpost) | **Post** /resellers/v1/webhooks/orderstatusevent | Order Status
*OrdersAPI* | [**DeleteOrdercancel**](docs/OrdersAPI.md#deleteordercancel) | **Delete** /resellers/v6/orders/{OrderNumber} | Cancel your Order
*OrdersAPI* | [**GetOrderdetailsV61**](docs/OrdersAPI.md#getorderdetailsv61) | **Get** /resellers/v6.1/orders/{ordernumber} | Get Order Details v6.1
*OrdersAPI* | [**GetResellersV6Ordersearch**](docs/OrdersAPI.md#getresellersv6ordersearch) | **Get** /resellers/v6/orders/search | Search your Orders
*OrdersAPI* | [**PostCreateorderV6**](docs/OrdersAPI.md#postcreateorderv6) | **Post** /resellers/v6/orders | Create your Order
*OrdersAPI* | [**PutOrdermodify**](docs/OrdersAPI.md#putordermodify) | **Put** /resellers/v6/orders/{orderNumber} | Modify your Order
*ProductCatalogAPI* | [**GetResellerV6Productdetail**](docs/ProductCatalogAPI.md#getresellerv6productdetail) | **Get** /resellers/v6/catalog/details/{ingramPartNumber} | Product Details
*ProductCatalogAPI* | [**GetResellerV6Productsearch**](docs/ProductCatalogAPI.md#getresellerv6productsearch) | **Get** /resellers/v6/catalog | Search Products
*ProductCatalogAPI* | [**PostPriceandavailability**](docs/ProductCatalogAPI.md#postpriceandavailability) | **Post** /resellers/v6/catalog/priceandavailability | Price and Availability
*QuoteToOrderAPI* | [**PostQuoteToOrderV6**](docs/QuoteToOrderAPI.md#postquotetoorderv6) | **Post** /resellers/v6/q2o/orders | Quote To Order
*QuotesAPI* | [**GetQuotessearchV6**](docs/QuotesAPI.md#getquotessearchv6) | **Get** /resellers/v6/quotes/search | Quote Search
*QuotesAPI* | [**GetResellerV6ValidateQuote**](docs/QuotesAPI.md#getresellerv6validatequote) | **Get** /resellers/v6/q2o/validatequote | Validate Quote
*QuotesAPI* | [**GetResellersV6Quotes**](docs/QuotesAPI.md#getresellersv6quotes) | **Get** /resellers/v6/quotes/{quoteNumber} | Get Quote Details
*RenewalsAPI* | [**GetResellersV6Renewalsdetails**](docs/RenewalsAPI.md#getresellersv6renewalsdetails) | **Get** /resellers/v6/renewals/{renewalId} | Renewals Details
*RenewalsAPI* | [**PostRenewalssearch**](docs/RenewalsAPI.md#postrenewalssearch) | **Post** /resellers/v6/renewals/search | Renewals Search
*ReturnsAPI* | [**GetResellersV6Returnsdetails**](docs/ReturnsAPI.md#getresellersv6returnsdetails) | **Get** /resellers/v6/returns/{caseRequestNumber} | Returns Details
*ReturnsAPI* | [**GetResellersV6Returnssearch**](docs/ReturnsAPI.md#getresellersv6returnssearch) | **Get** /resellers/v6/returns/search | Returns Search
*ReturnsAPI* | [**PostReturnscreate**](docs/ReturnsAPI.md#postreturnscreate) | **Post** /resellers/v6/returns/create | Returns Create
*StockUpdateAPI* | [**ResellersV1WebhooksAvailabilityupdatePost**](docs/StockUpdateAPI.md#resellersv1webhooksavailabilityupdatepost) | **Post** /resellers/v1/webhooks/availabilityupdate | Stock Update


## Documentation For Models

 - [AccesstokenResponse](docs/AccesstokenResponse.md)
 - [AvailabilityAsyncNotificationRequest](docs/AvailabilityAsyncNotificationRequest.md)
 - [AvailabilityAsyncNotificationRequestResourceInner](docs/AvailabilityAsyncNotificationRequestResourceInner.md)
 - [AvailabilityAsyncNotificationRequestResourceInnerLinksInner](docs/AvailabilityAsyncNotificationRequestResourceInnerLinksInner.md)
 - [DealsDetailsResponse](docs/DealsDetailsResponse.md)
 - [DealsDetailsResponseProductsInner](docs/DealsDetailsResponseProductsInner.md)
 - [DealsSearchResponse](docs/DealsSearchResponse.md)
 - [DealsSearchResponseDealsInner](docs/DealsSearchResponseDealsInner.md)
 - [Error](docs/Error.md)
 - [ErrorResponse](docs/ErrorResponse.md)
 - [ErrorResponseDTO](docs/ErrorResponseDTO.md)
 - [ErrorResponseErrorsInner](docs/ErrorResponseErrorsInner.md)
 - [ErrorResponseErrorsInnerFieldsInner](docs/ErrorResponseErrorsInnerFieldsInner.md)
 - [Fields](docs/Fields.md)
 - [FreightRequest](docs/FreightRequest.md)
 - [FreightRequestLinesInner](docs/FreightRequestLinesInner.md)
 - [FreightRequestShipToAddressInner](docs/FreightRequestShipToAddressInner.md)
 - [FreightResponse](docs/FreightResponse.md)
 - [FreightResponseFreightEstimateResponse](docs/FreightResponseFreightEstimateResponse.md)
 - [FreightResponseFreightEstimateResponseDistributionInner](docs/FreightResponseFreightEstimateResponseDistributionInner.md)
 - [FreightResponseFreightEstimateResponseDistributionInnerCarrierListInner](docs/FreightResponseFreightEstimateResponseDistributionInnerCarrierListInner.md)
 - [FreightResponseFreightEstimateResponseLinesInner](docs/FreightResponseFreightEstimateResponseLinesInner.md)
 - [GetAccesstoken400Response](docs/GetAccesstoken400Response.md)
 - [GetAccesstoken500Response](docs/GetAccesstoken500Response.md)
 - [GetAccesstoken500ResponseFault](docs/GetAccesstoken500ResponseFault.md)
 - [GetAccesstoken500ResponseFaultDetail](docs/GetAccesstoken500ResponseFaultDetail.md)
 - [GetResellerV6ValidateQuote400Response](docs/GetResellerV6ValidateQuote400Response.md)
 - [GetResellerV6ValidateQuote400ResponseFieldsInner](docs/GetResellerV6ValidateQuote400ResponseFieldsInner.md)
 - [GetResellerV6ValidateQuote500Response](docs/GetResellerV6ValidateQuote500Response.md)
 - [InvoiceDetailsv61Response](docs/InvoiceDetailsv61Response.md)
 - [InvoiceDetailsv61ResponseBillToInfo](docs/InvoiceDetailsv61ResponseBillToInfo.md)
 - [InvoiceDetailsv61ResponseFxRateInfo](docs/InvoiceDetailsv61ResponseFxRateInfo.md)
 - [InvoiceDetailsv61ResponseLinesInner](docs/InvoiceDetailsv61ResponseLinesInner.md)
 - [InvoiceDetailsv61ResponseLinesInnerSerialNumbersInner](docs/InvoiceDetailsv61ResponseLinesInnerSerialNumbersInner.md)
 - [InvoiceDetailsv61ResponsePaymentTermsInfo](docs/InvoiceDetailsv61ResponsePaymentTermsInfo.md)
 - [InvoiceDetailsv61ResponseShipToInfo](docs/InvoiceDetailsv61ResponseShipToInfo.md)
 - [InvoiceDetailsv61ResponseSummary](docs/InvoiceDetailsv61ResponseSummary.md)
 - [InvoiceDetailsv61ResponseSummaryForeignFxTotals](docs/InvoiceDetailsv61ResponseSummaryForeignFxTotals.md)
 - [InvoiceDetailsv61ResponseSummaryLines](docs/InvoiceDetailsv61ResponseSummaryLines.md)
 - [InvoiceDetailsv61ResponseSummaryMiscChargesInner](docs/InvoiceDetailsv61ResponseSummaryMiscChargesInner.md)
 - [InvoiceDetailsv61ResponseSummaryTotals](docs/InvoiceDetailsv61ResponseSummaryTotals.md)
 - [InvoiceSearchResponse](docs/InvoiceSearchResponse.md)
 - [InvoiceSearchResponseInvoicesInner](docs/InvoiceSearchResponseInvoicesInner.md)
 - [OrderCreateRequest](docs/OrderCreateRequest.md)
 - [OrderCreateRequestAdditionalAttributesInner](docs/OrderCreateRequestAdditionalAttributesInner.md)
 - [OrderCreateRequestEndUserInfo](docs/OrderCreateRequestEndUserInfo.md)
 - [OrderCreateRequestLinesInner](docs/OrderCreateRequestLinesInner.md)
 - [OrderCreateRequestLinesInnerAdditionalAttributesInner](docs/OrderCreateRequestLinesInnerAdditionalAttributesInner.md)
 - [OrderCreateRequestLinesInnerEndUserInfoInner](docs/OrderCreateRequestLinesInnerEndUserInfoInner.md)
 - [OrderCreateRequestLinesInnerWarrantyInfoInner](docs/OrderCreateRequestLinesInnerWarrantyInfoInner.md)
 - [OrderCreateRequestLinesInnerWarrantyInfoInnerSerialInfoInner](docs/OrderCreateRequestLinesInnerWarrantyInfoInnerSerialInfoInner.md)
 - [OrderCreateRequestResellerInfo](docs/OrderCreateRequestResellerInfo.md)
 - [OrderCreateRequestShipToInfo](docs/OrderCreateRequestShipToInfo.md)
 - [OrderCreateRequestShipmentDetails](docs/OrderCreateRequestShipmentDetails.md)
 - [OrderCreateRequestVmf](docs/OrderCreateRequestVmf.md)
 - [OrderCreateResponse](docs/OrderCreateResponse.md)
 - [OrderCreateResponseEndUserInfo](docs/OrderCreateResponseEndUserInfo.md)
 - [OrderCreateResponseOrdersInner](docs/OrderCreateResponseOrdersInner.md)
 - [OrderCreateResponseOrdersInnerAdditionalAttributesInner](docs/OrderCreateResponseOrdersInnerAdditionalAttributesInner.md)
 - [OrderCreateResponseOrdersInnerLinesInner](docs/OrderCreateResponseOrdersInnerLinesInner.md)
 - [OrderCreateResponseOrdersInnerLinesInnerAdditionalAttributesInner](docs/OrderCreateResponseOrdersInnerLinesInnerAdditionalAttributesInner.md)
 - [OrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner](docs/OrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner.md)
 - [OrderCreateResponseOrdersInnerLinksInner](docs/OrderCreateResponseOrdersInnerLinksInner.md)
 - [OrderCreateResponseOrdersInnerMiscellaneousChargesInner](docs/OrderCreateResponseOrdersInnerMiscellaneousChargesInner.md)
 - [OrderCreateResponseOrdersInnerRejectedLineItemsInner](docs/OrderCreateResponseOrdersInnerRejectedLineItemsInner.md)
 - [OrderCreateResponseShipToInfo](docs/OrderCreateResponseShipToInfo.md)
 - [OrderDetailB2B](docs/OrderDetailB2B.md)
 - [OrderDetailB2BAdditionalAttributesInner](docs/OrderDetailB2BAdditionalAttributesInner.md)
 - [OrderDetailB2BBillToInfo](docs/OrderDetailB2BBillToInfo.md)
 - [OrderDetailB2BEndUserInfo](docs/OrderDetailB2BEndUserInfo.md)
 - [OrderDetailB2BLinesInner](docs/OrderDetailB2BLinesInner.md)
 - [OrderDetailB2BLinesInnerAdditionalAttributesInner](docs/OrderDetailB2BLinesInnerAdditionalAttributesInner.md)
 - [OrderDetailB2BLinesInnerEstimatedDatesInner](docs/OrderDetailB2BLinesInnerEstimatedDatesInner.md)
 - [OrderDetailB2BLinesInnerEstimatedDatesInnerDelivery](docs/OrderDetailB2BLinesInnerEstimatedDatesInnerDelivery.md)
 - [OrderDetailB2BLinesInnerEstimatedDatesInnerDeliveryDeliveryDateRange](docs/OrderDetailB2BLinesInnerEstimatedDatesInnerDeliveryDeliveryDateRange.md)
 - [OrderDetailB2BLinesInnerEstimatedDatesInnerShip](docs/OrderDetailB2BLinesInnerEstimatedDatesInnerShip.md)
 - [OrderDetailB2BLinesInnerEstimatedDatesInnerShipShipDateRange](docs/OrderDetailB2BLinesInnerEstimatedDatesInnerShipShipDateRange.md)
 - [OrderDetailB2BLinesInnerLinksInner](docs/OrderDetailB2BLinesInnerLinksInner.md)
 - [OrderDetailB2BLinesInnerMultipleShipmentsInner](docs/OrderDetailB2BLinesInnerMultipleShipmentsInner.md)
 - [OrderDetailB2BLinesInnerScheduleLinesInner](docs/OrderDetailB2BLinesInnerScheduleLinesInner.md)
 - [OrderDetailB2BLinesInnerServiceContractInfo](docs/OrderDetailB2BLinesInnerServiceContractInfo.md)
 - [OrderDetailB2BLinesInnerServiceContractInfoContractInfo](docs/OrderDetailB2BLinesInnerServiceContractInfoContractInfo.md)
 - [OrderDetailB2BLinesInnerServiceContractInfoLicenseInfo](docs/OrderDetailB2BLinesInnerServiceContractInfoLicenseInfo.md)
 - [OrderDetailB2BLinesInnerServiceContractInfoSubscriptions](docs/OrderDetailB2BLinesInnerServiceContractInfoSubscriptions.md)
 - [OrderDetailB2BLinesInnerShipmentDetailsInner](docs/OrderDetailB2BLinesInnerShipmentDetailsInner.md)
 - [OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner](docs/OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner.md)
 - [OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInnerTrackingDetailsInner](docs/OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInnerTrackingDetailsInner.md)
 - [OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInnerTrackingDetailsInnerSerialNumbersInner](docs/OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInnerTrackingDetailsInnerSerialNumbersInner.md)
 - [OrderDetailB2BMiscellaneousChargesInner](docs/OrderDetailB2BMiscellaneousChargesInner.md)
 - [OrderDetailB2BShipToInfo](docs/OrderDetailB2BShipToInfo.md)
 - [OrderModifyRequest](docs/OrderModifyRequest.md)
 - [OrderModifyRequestAdditionalAttributesInner](docs/OrderModifyRequestAdditionalAttributesInner.md)
 - [OrderModifyRequestLinesInner](docs/OrderModifyRequestLinesInner.md)
 - [OrderModifyRequestShipToInfo](docs/OrderModifyRequestShipToInfo.md)
 - [OrderModifyResponse](docs/OrderModifyResponse.md)
 - [OrderModifyResponseLinesInner](docs/OrderModifyResponseLinesInner.md)
 - [OrderModifyResponseLinesInnerAdditionalAttributesInner](docs/OrderModifyResponseLinesInnerAdditionalAttributesInner.md)
 - [OrderModifyResponseLinesInnerShipmentDetails](docs/OrderModifyResponseLinesInnerShipmentDetails.md)
 - [OrderModifyResponseRejectedLineItemsInner](docs/OrderModifyResponseRejectedLineItemsInner.md)
 - [OrderModifyResponseShipToInfo](docs/OrderModifyResponseShipToInfo.md)
 - [OrderSearchResponse](docs/OrderSearchResponse.md)
 - [OrderSearchResponseOrdersInner](docs/OrderSearchResponseOrdersInner.md)
 - [OrderSearchResponseOrdersInnerLinks](docs/OrderSearchResponseOrdersInnerLinks.md)
 - [OrderSearchResponseOrdersInnerSubOrdersInner](docs/OrderSearchResponseOrdersInnerSubOrdersInner.md)
 - [OrderSearchResponseOrdersInnerSubOrdersInnerLinksInner](docs/OrderSearchResponseOrdersInnerSubOrdersInnerLinksInner.md)
 - [OrderStatusAsyncNotificationRequest](docs/OrderStatusAsyncNotificationRequest.md)
 - [OrderStatusAsyncNotificationRequestResourceInner](docs/OrderStatusAsyncNotificationRequestResourceInner.md)
 - [OrderStatusAsyncNotificationRequestResourceInnerLinesInner](docs/OrderStatusAsyncNotificationRequestResourceInnerLinesInner.md)
 - [OrderStatusAsyncNotificationRequestResourceInnerLinesInnerSerialNumberDetailsInner](docs/OrderStatusAsyncNotificationRequestResourceInnerLinesInnerSerialNumberDetailsInner.md)
 - [OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner](docs/OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner.md)
 - [OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInnerPackageDetailsInner](docs/OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInnerPackageDetailsInner.md)
 - [OrderStatusAsyncNotificationRequestResourceInnerLinksInner](docs/OrderStatusAsyncNotificationRequestResourceInnerLinksInner.md)
 - [PostQuoteToOrderV6400Response](docs/PostQuoteToOrderV6400Response.md)
 - [PostQuoteToOrderV6400ResponseFieldsInner](docs/PostQuoteToOrderV6400ResponseFieldsInner.md)
 - [PostRenewalssearch400Response](docs/PostRenewalssearch400Response.md)
 - [PriceAndAvailabilityRequest](docs/PriceAndAvailabilityRequest.md)
 - [PriceAndAvailabilityRequestAdditionalAttributesInner](docs/PriceAndAvailabilityRequestAdditionalAttributesInner.md)
 - [PriceAndAvailabilityRequestAvailabilityByWarehouseInner](docs/PriceAndAvailabilityRequestAvailabilityByWarehouseInner.md)
 - [PriceAndAvailabilityRequestProductsInner](docs/PriceAndAvailabilityRequestProductsInner.md)
 - [PriceAndAvailabilityRequestProductsInnerAdditionalAttributesInner](docs/PriceAndAvailabilityRequestProductsInnerAdditionalAttributesInner.md)
 - [PriceAndAvailabilityResponseInner](docs/PriceAndAvailabilityResponseInner.md)
 - [PriceAndAvailabilityResponseInnerAvailability](docs/PriceAndAvailabilityResponseInnerAvailability.md)
 - [PriceAndAvailabilityResponseInnerAvailabilityAvailabilityByWarehouseInner](docs/PriceAndAvailabilityResponseInnerAvailabilityAvailabilityByWarehouseInner.md)
 - [PriceAndAvailabilityResponseInnerAvailabilityAvailabilityByWarehouseInnerBackOrderInfoInner](docs/PriceAndAvailabilityResponseInnerAvailabilityAvailabilityByWarehouseInnerBackOrderInfoInner.md)
 - [PriceAndAvailabilityResponseInnerDiscountsInner](docs/PriceAndAvailabilityResponseInnerDiscountsInner.md)
 - [PriceAndAvailabilityResponseInnerDiscountsInnerQuantityDiscountsInner](docs/PriceAndAvailabilityResponseInnerDiscountsInnerQuantityDiscountsInner.md)
 - [PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner](docs/PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner.md)
 - [PriceAndAvailabilityResponseInnerPricing](docs/PriceAndAvailabilityResponseInnerPricing.md)
 - [PriceAndAvailabilityResponseInnerReserveInventoryDetailsInner](docs/PriceAndAvailabilityResponseInnerReserveInventoryDetailsInner.md)
 - [PriceAndAvailabilityResponseInnerServiceFeesInner](docs/PriceAndAvailabilityResponseInnerServiceFeesInner.md)
 - [ProductDetailResponse](docs/ProductDetailResponse.md)
 - [ProductDetailResponseAdditionalInformation](docs/ProductDetailResponseAdditionalInformation.md)
 - [ProductDetailResponseAdditionalInformationProductWeightInner](docs/ProductDetailResponseAdditionalInformationProductWeightInner.md)
 - [ProductDetailResponseCiscoFields](docs/ProductDetailResponseCiscoFields.md)
 - [ProductDetailResponseIndicators](docs/ProductDetailResponseIndicators.md)
 - [ProductDetailResponseTechnicalSpecificationsInner](docs/ProductDetailResponseTechnicalSpecificationsInner.md)
 - [ProductSearchResponse](docs/ProductSearchResponse.md)
 - [ProductSearchResponseCatalogInner](docs/ProductSearchResponseCatalogInner.md)
 - [ProductSearchResponseCatalogInnerLinksInner](docs/ProductSearchResponseCatalogInnerLinksInner.md)
 - [QuoteDetailsResponse](docs/QuoteDetailsResponse.md)
 - [QuoteDetailsResponseAdditionalAttributesInner](docs/QuoteDetailsResponseAdditionalAttributesInner.md)
 - [QuoteDetailsResponseEndUserInfo](docs/QuoteDetailsResponseEndUserInfo.md)
 - [QuoteDetailsResponseProductsInner](docs/QuoteDetailsResponseProductsInner.md)
 - [QuoteDetailsResponseProductsInnerPrice](docs/QuoteDetailsResponseProductsInnerPrice.md)
 - [QuoteDetailsResponseResellerInfo](docs/QuoteDetailsResponseResellerInfo.md)
 - [QuoteSearchResponse](docs/QuoteSearchResponse.md)
 - [QuoteSearchResponseQuotesInner](docs/QuoteSearchResponseQuotesInner.md)
 - [QuoteSearchResponseQuotesInnerLinks](docs/QuoteSearchResponseQuotesInnerLinks.md)
 - [QuoteToOrderDetailsDTO](docs/QuoteToOrderDetailsDTO.md)
 - [QuoteToOrderDetailsDTOAdditionalAttributesInner](docs/QuoteToOrderDetailsDTOAdditionalAttributesInner.md)
 - [QuoteToOrderDetailsDTOEndUserInfo](docs/QuoteToOrderDetailsDTOEndUserInfo.md)
 - [QuoteToOrderDetailsDTOLinesInner](docs/QuoteToOrderDetailsDTOLinesInner.md)
 - [QuoteToOrderDetailsDTOLinesInnerVmfAdditionalAttributesLinesInner](docs/QuoteToOrderDetailsDTOLinesInnerVmfAdditionalAttributesLinesInner.md)
 - [QuoteToOrderDetailsDTOShipToInfo](docs/QuoteToOrderDetailsDTOShipToInfo.md)
 - [QuoteToOrderDetailsDTOVmfadditionalAttributesInner](docs/QuoteToOrderDetailsDTOVmfadditionalAttributesInner.md)
 - [QuoteToOrderResponse](docs/QuoteToOrderResponse.md)
 - [RenewalsDetailsResponse](docs/RenewalsDetailsResponse.md)
 - [RenewalsDetailsResponseAdditionalAttributesInner](docs/RenewalsDetailsResponseAdditionalAttributesInner.md)
 - [RenewalsDetailsResponseEndUserInfo](docs/RenewalsDetailsResponseEndUserInfo.md)
 - [RenewalsDetailsResponseProductsInner](docs/RenewalsDetailsResponseProductsInner.md)
 - [RenewalsDetailsResponseReferenceNumber](docs/RenewalsDetailsResponseReferenceNumber.md)
 - [RenewalsSearchRequest](docs/RenewalsSearchRequest.md)
 - [RenewalsSearchRequestDateType](docs/RenewalsSearchRequestDateType.md)
 - [RenewalsSearchRequestDateTypeEndDate](docs/RenewalsSearchRequestDateTypeEndDate.md)
 - [RenewalsSearchRequestDateTypeExpirationDate](docs/RenewalsSearchRequestDateTypeExpirationDate.md)
 - [RenewalsSearchRequestDateTypeInvoiceDate](docs/RenewalsSearchRequestDateTypeInvoiceDate.md)
 - [RenewalsSearchRequestDateTypeStartDate](docs/RenewalsSearchRequestDateTypeStartDate.md)
 - [RenewalsSearchRequestStatus](docs/RenewalsSearchRequestStatus.md)
 - [RenewalsSearchRequestStatusOpporutinyStatus](docs/RenewalsSearchRequestStatusOpporutinyStatus.md)
 - [RenewalsSearchResponse](docs/RenewalsSearchResponse.md)
 - [RenewalsSearchResponseRenewalsInner](docs/RenewalsSearchResponseRenewalsInner.md)
 - [RenewalsSearchResponseRenewalsInnerLinksInner](docs/RenewalsSearchResponseRenewalsInnerLinksInner.md)
 - [ReturnsCreateRequest](docs/ReturnsCreateRequest.md)
 - [ReturnsCreateRequestListInner](docs/ReturnsCreateRequestListInner.md)
 - [ReturnsCreateRequestListInnerShipFromInfoInner](docs/ReturnsCreateRequestListInnerShipFromInfoInner.md)
 - [ReturnsCreateResponse](docs/ReturnsCreateResponse.md)
 - [ReturnsCreateResponseReturnsClaimsInner](docs/ReturnsCreateResponseReturnsClaimsInner.md)
 - [ReturnsDetailsResponse](docs/ReturnsDetailsResponse.md)
 - [ReturnsDetailsResponseProductsInner](docs/ReturnsDetailsResponseProductsInner.md)
 - [ReturnsSearchResponse](docs/ReturnsSearchResponse.md)
 - [ReturnsSearchResponseReturnsClaimsInner](docs/ReturnsSearchResponseReturnsClaimsInner.md)
 - [ReturnsSearchResponseReturnsClaimsInnerLinksInner](docs/ReturnsSearchResponseReturnsClaimsInnerLinksInner.md)
 - [ValidateQuoteResponse](docs/ValidateQuoteResponse.md)
 - [ValidateQuoteResponseLinesInner](docs/ValidateQuoteResponseLinesInner.md)
 - [ValidateQuoteResponseVmfAdditionalAttributesInner](docs/ValidateQuoteResponseVmfAdditionalAttributesInner.md)


## Documentation for Authorization


Authentication schemes defined for the API:
### application

- **Type**: OAuth
- **Flow**: application
- **Authorization URL**: https://api.ingrammicro.com:443/oauth/oauth20/token?grant_type=client_credentials&client_id={ClientId}&client_secret={clientSecret}
- **Method**: Get
- **Scopes**: 
  - write: allows modifying resources
  - read: allows reading resources

Example

```go
auth := context.WithValue(context.Background(), xi_sdk_resellers.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```go
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, xi_sdk_resellers.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author
-[Ingram Micro Xvantage](https://github.com/ingrammicro-xvantage)

## Contact

For any inquiries or support, please feel free to contact us at:

- Email: xi_support@ingrammicro.com