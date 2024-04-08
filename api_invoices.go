/*
XI Sdk Resellers

For resellers seeking to innovate with Ingram Micro's API solutions, automate your eCommerce experience with our array of API's and webhooks to craft a seamless journey for your customers.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package xi_sdk_resellers

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"reflect"
)


// InvoicesAPIService InvoicesAPI service
type InvoicesAPIService service

type ApiGetInvoicedetailsV61Request struct {
	ctx context.Context
	ApiService *InvoicesAPIService
	invoiceNumber string
	iMCustomerNumber *string
	iMCountryCode *string
	iMCorrelationID *string
	iMApplicationID *string
	customerType *string
	includeSerialNumbers *bool
}

// Your unique Ingram Micro customer number.
func (r ApiGetInvoicedetailsV61Request) IMCustomerNumber(iMCustomerNumber string) ApiGetInvoicedetailsV61Request {
	r.iMCustomerNumber = &iMCustomerNumber
	return r
}

// Two-character ISO country code.
func (r ApiGetInvoicedetailsV61Request) IMCountryCode(iMCountryCode string) ApiGetInvoicedetailsV61Request {
	r.iMCountryCode = &iMCountryCode
	return r
}

// Unique transaction number to identify each transaction across all the systems.
func (r ApiGetInvoicedetailsV61Request) IMCorrelationID(iMCorrelationID string) ApiGetInvoicedetailsV61Request {
	r.iMCorrelationID = &iMCorrelationID
	return r
}

// Unique value used to identify the sender of the transaction. Example: MyCompany.
func (r ApiGetInvoicedetailsV61Request) IMApplicationID(iMApplicationID string) ApiGetInvoicedetailsV61Request {
	r.iMApplicationID = &iMApplicationID
	return r
}

// it should be invoice or order
func (r ApiGetInvoicedetailsV61Request) CustomerType(customerType string) ApiGetInvoicedetailsV61Request {
	r.customerType = &customerType
	return r
}

// if serial in the response send as true or else false
func (r ApiGetInvoicedetailsV61Request) IncludeSerialNumbers(includeSerialNumbers bool) ApiGetInvoicedetailsV61Request {
	r.includeSerialNumbers = &includeSerialNumbers
	return r
}

func (r ApiGetInvoicedetailsV61Request) Execute() (*InvoiceDetailsv61Response, *http.Response, error) {
	return r.ApiService.GetInvoicedetailsV61Execute(r)
}

/*
GetInvoicedetailsV61 Get Invoice Details v6.1

Use your Ingram Micro invoice number to search for existing invoices or retrieve existing invoice details.

The invoice number, IM-CustomerNumber, IM-CountryCode, IM-ApplicationId and IM-CorrelationID are required parameters.

.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param invoiceNumber The Ingram Micro invoice number.
 @return ApiGetInvoicedetailsV61Request
*/
func (a *InvoicesAPIService) GetInvoicedetailsV61(ctx context.Context, invoiceNumber string) ApiGetInvoicedetailsV61Request {
	return ApiGetInvoicedetailsV61Request{
		ApiService: a,
		ctx: ctx,
		invoiceNumber: invoiceNumber,
	}
}

// Execute executes the request
//  @return InvoiceDetailsv61Response
func (a *InvoicesAPIService) GetInvoicedetailsV61Execute(r ApiGetInvoicedetailsV61Request) (*InvoiceDetailsv61Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InvoiceDetailsv61Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InvoicesAPIService.GetInvoicedetailsV61")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/resellers/v6.1/invoices/{invoiceNumber}"
	localVarPath = strings.Replace(localVarPath, "{"+"invoiceNumber"+"}", url.PathEscape(parameterValueToString(r.invoiceNumber, "invoiceNumber")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.invoiceNumber) > 12 {
		return localVarReturnValue, nil, reportError("invoiceNumber must have less than 12 elements")
	}
	if r.iMCustomerNumber == nil {
		return localVarReturnValue, nil, reportError("iMCustomerNumber is required and must be specified")
	}
	if strlen(*r.iMCustomerNumber) > 10 {
		return localVarReturnValue, nil, reportError("iMCustomerNumber must have less than 10 elements")
	}
	if r.iMCountryCode == nil {
		return localVarReturnValue, nil, reportError("iMCountryCode is required and must be specified")
	}
	if strlen(*r.iMCountryCode) < 2 {
		return localVarReturnValue, nil, reportError("iMCountryCode must have at least 2 elements")
	}
	if strlen(*r.iMCountryCode) > 2 {
		return localVarReturnValue, nil, reportError("iMCountryCode must have less than 2 elements")
	}
	if r.iMCorrelationID == nil {
		return localVarReturnValue, nil, reportError("iMCorrelationID is required and must be specified")
	}
	if strlen(*r.iMCorrelationID) > 32 {
		return localVarReturnValue, nil, reportError("iMCorrelationID must have less than 32 elements")
	}
	if r.iMApplicationID == nil {
		return localVarReturnValue, nil, reportError("iMApplicationID is required and must be specified")
	}
	if strlen(*r.iMApplicationID) > 32 {
		return localVarReturnValue, nil, reportError("iMApplicationID must have less than 32 elements")
	}

	if r.customerType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "customerType", r.customerType, "")
	}
	if r.includeSerialNumbers != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeSerialNumbers", r.includeSerialNumbers, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "IM-CustomerNumber", r.iMCustomerNumber, "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "IM-CountryCode", r.iMCountryCode, "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "IM-CorrelationID", r.iMCorrelationID, "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "IM-ApplicationID", r.iMApplicationID, "")
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetResellersV6InvoicesearchRequest struct {
	ctx context.Context
	ApiService *InvoicesAPIService
	iMApplicationID *string
	iMCustomerNumber *string
	iMCountryCode *string
	iMCorrelationID *string
	paymentTermsNetDate *string
	invoiceDate *string
	invoiceDueDate *string
	orderDate *string
	orderFromDate *string
	orderToDate *string
	orderNumber *string
	deliveryNumber *string
	invoiceNumber *string
	invoiceStatus *string
	invoiceType *string
	customerOrderNumber *string
	endCustomerOrderNumber *string
	specialBidNumber *string
	invoiceFromDueDate *string
	invoiceToDueDate *string
	invoiceFromDate *[]string
	invoiceToDate *[]string
	pageSize *int32
	pageNumber *int32
	orderby *string
	direction *string
	serialNumber *string
}

// Unique value used to identify the sender of the transaction. Example: MyCompany
func (r ApiGetResellersV6InvoicesearchRequest) IMApplicationID(iMApplicationID string) ApiGetResellersV6InvoicesearchRequest {
	r.iMApplicationID = &iMApplicationID
	return r
}

// Your unique Ingram Micro customer number.
func (r ApiGetResellersV6InvoicesearchRequest) IMCustomerNumber(iMCustomerNumber string) ApiGetResellersV6InvoicesearchRequest {
	r.iMCustomerNumber = &iMCustomerNumber
	return r
}

// Two-character ISO country code.
func (r ApiGetResellersV6InvoicesearchRequest) IMCountryCode(iMCountryCode string) ApiGetResellersV6InvoicesearchRequest {
	r.iMCountryCode = &iMCountryCode
	return r
}

// Unique transaction number to identify each transaction across all the systems.
func (r ApiGetResellersV6InvoicesearchRequest) IMCorrelationID(iMCorrelationID string) ApiGetResellersV6InvoicesearchRequest {
	r.iMCorrelationID = &iMCorrelationID
	return r
}

// Search by payment terms net date(yyyy-MM-dd).
func (r ApiGetResellersV6InvoicesearchRequest) PaymentTermsNetDate(paymentTermsNetDate string) ApiGetResellersV6InvoicesearchRequest {
	r.paymentTermsNetDate = &paymentTermsNetDate
	return r
}

// Search by invoice date(yyyy-MM-dd).
func (r ApiGetResellersV6InvoicesearchRequest) InvoiceDate(invoiceDate string) ApiGetResellersV6InvoicesearchRequest {
	r.invoiceDate = &invoiceDate
	return r
}

// Search by invoice date from(yyyy-MM-dd).
func (r ApiGetResellersV6InvoicesearchRequest) InvoiceDueDate(invoiceDueDate string) ApiGetResellersV6InvoicesearchRequest {
	r.invoiceDueDate = &invoiceDueDate
	return r
}

// Search by OrderDate date(yyyy-MM-dd).
func (r ApiGetResellersV6InvoicesearchRequest) OrderDate(orderDate string) ApiGetResellersV6InvoicesearchRequest {
	r.orderDate = &orderDate
	return r
}

// Search by OrderFromDate date(yyyy-MM-dd).
func (r ApiGetResellersV6InvoicesearchRequest) OrderFromDate(orderFromDate string) ApiGetResellersV6InvoicesearchRequest {
	r.orderFromDate = &orderFromDate
	return r
}

// Search by OrderToDate date(yyyy-MM-dd).
func (r ApiGetResellersV6InvoicesearchRequest) OrderToDate(orderToDate string) ApiGetResellersV6InvoicesearchRequest {
	r.orderToDate = &orderToDate
	return r
}

// Search by order number
func (r ApiGetResellersV6InvoicesearchRequest) OrderNumber(orderNumber string) ApiGetResellersV6InvoicesearchRequest {
	r.orderNumber = &orderNumber
	return r
}

// Search by delivery number.
func (r ApiGetResellersV6InvoicesearchRequest) DeliveryNumber(deliveryNumber string) ApiGetResellersV6InvoicesearchRequest {
	r.deliveryNumber = &deliveryNumber
	return r
}

// The Ingram Micro invoice number.
func (r ApiGetResellersV6InvoicesearchRequest) InvoiceNumber(invoiceNumber string) ApiGetResellersV6InvoicesearchRequest {
	r.invoiceNumber = &invoiceNumber
	return r
}

// Ingram Micro invoice status.
func (r ApiGetResellersV6InvoicesearchRequest) InvoiceStatus(invoiceStatus string) ApiGetResellersV6InvoicesearchRequest {
	r.invoiceStatus = &invoiceStatus
	return r
}

// Ingram Micro InvoiceType.
func (r ApiGetResellersV6InvoicesearchRequest) InvoiceType(invoiceType string) ApiGetResellersV6InvoicesearchRequest {
	r.invoiceType = &invoiceType
	return r
}

// Ingram Micro CustomerOrderNumber.
func (r ApiGetResellersV6InvoicesearchRequest) CustomerOrderNumber(customerOrderNumber string) ApiGetResellersV6InvoicesearchRequest {
	r.customerOrderNumber = &customerOrderNumber
	return r
}

// Ingram Micro EndCustomerOrderNumber.
func (r ApiGetResellersV6InvoicesearchRequest) EndCustomerOrderNumber(endCustomerOrderNumber string) ApiGetResellersV6InvoicesearchRequest {
	r.endCustomerOrderNumber = &endCustomerOrderNumber
	return r
}

// Ingram Micro SpecialBidNumber.
func (r ApiGetResellersV6InvoicesearchRequest) SpecialBidNumber(specialBidNumber string) ApiGetResellersV6InvoicesearchRequest {
	r.specialBidNumber = &specialBidNumber
	return r
}

// Search by invoice due date from(yyyy-MM-dd).
func (r ApiGetResellersV6InvoicesearchRequest) InvoiceFromDueDate(invoiceFromDueDate string) ApiGetResellersV6InvoicesearchRequest {
	r.invoiceFromDueDate = &invoiceFromDueDate
	return r
}

// Search by invoice due date to(yyyy-MM-dd).
func (r ApiGetResellersV6InvoicesearchRequest) InvoiceToDueDate(invoiceToDueDate string) ApiGetResellersV6InvoicesearchRequest {
	r.invoiceToDueDate = &invoiceToDueDate
	return r
}

// Search by invoice date from(yyyy-MM-dd).
func (r ApiGetResellersV6InvoicesearchRequest) InvoiceFromDate(invoiceFromDate []string) ApiGetResellersV6InvoicesearchRequest {
	r.invoiceFromDate = &invoiceFromDate
	return r
}

// Search by invoice date To(yyyy-MM-dd).
func (r ApiGetResellersV6InvoicesearchRequest) InvoiceToDate(invoiceToDate []string) ApiGetResellersV6InvoicesearchRequest {
	r.invoiceToDate = &invoiceToDate
	return r
}

// Number of records required in the call - max records 100 per page.
func (r ApiGetResellersV6InvoicesearchRequest) PageSize(pageSize int32) ApiGetResellersV6InvoicesearchRequest {
	r.pageSize = &pageSize
	return r
}

// The page number reference.
func (r ApiGetResellersV6InvoicesearchRequest) PageNumber(pageNumber int32) ApiGetResellersV6InvoicesearchRequest {
	r.pageNumber = &pageNumber
	return r
}

// Column name with which we want to sort.
func (r ApiGetResellersV6InvoicesearchRequest) Orderby(orderby string) ApiGetResellersV6InvoicesearchRequest {
	r.orderby = &orderby
	return r
}

// asc or desc , along with orderby column result set will be sorted.
func (r ApiGetResellersV6InvoicesearchRequest) Direction(direction string) ApiGetResellersV6InvoicesearchRequest {
	r.direction = &direction
	return r
}

// Serial number of the product.
func (r ApiGetResellersV6InvoicesearchRequest) SerialNumber(serialNumber string) ApiGetResellersV6InvoicesearchRequest {
	r.serialNumber = &serialNumber
	return r
}

func (r ApiGetResellersV6InvoicesearchRequest) Execute() (*InvoiceSearchResponse, *http.Response, error) {
	return r.ApiService.GetResellersV6InvoicesearchExecute(r)
}

/*
GetResellersV6Invoicesearch Search your invoice

Search your Ingram Micro invoices. This endpoint searches by multiple invoice parameters and supports pagination of results.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetResellersV6InvoicesearchRequest
*/
func (a *InvoicesAPIService) GetResellersV6Invoicesearch(ctx context.Context) ApiGetResellersV6InvoicesearchRequest {
	return ApiGetResellersV6InvoicesearchRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return InvoiceSearchResponse
func (a *InvoicesAPIService) GetResellersV6InvoicesearchExecute(r ApiGetResellersV6InvoicesearchRequest) (*InvoiceSearchResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InvoiceSearchResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InvoicesAPIService.GetResellersV6Invoicesearch")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/resellers/v6/invoices"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.iMApplicationID == nil {
		return localVarReturnValue, nil, reportError("iMApplicationID is required and must be specified")
	}
	if strlen(*r.iMApplicationID) > 32 {
		return localVarReturnValue, nil, reportError("iMApplicationID must have less than 32 elements")
	}
	if r.iMCustomerNumber == nil {
		return localVarReturnValue, nil, reportError("iMCustomerNumber is required and must be specified")
	}
	if strlen(*r.iMCustomerNumber) > 10 {
		return localVarReturnValue, nil, reportError("iMCustomerNumber must have less than 10 elements")
	}
	if r.iMCountryCode == nil {
		return localVarReturnValue, nil, reportError("iMCountryCode is required and must be specified")
	}
	if strlen(*r.iMCountryCode) < 2 {
		return localVarReturnValue, nil, reportError("iMCountryCode must have at least 2 elements")
	}
	if strlen(*r.iMCountryCode) > 2 {
		return localVarReturnValue, nil, reportError("iMCountryCode must have less than 2 elements")
	}
	if r.iMCorrelationID == nil {
		return localVarReturnValue, nil, reportError("iMCorrelationID is required and must be specified")
	}
	if strlen(*r.iMCorrelationID) > 32 {
		return localVarReturnValue, nil, reportError("iMCorrelationID must have less than 32 elements")
	}

	if r.paymentTermsNetDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "paymentTermsNetDate", r.paymentTermsNetDate, "")
	}
	if r.invoiceDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "invoiceDate", r.invoiceDate, "")
	}
	if r.invoiceDueDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "invoiceDueDate", r.invoiceDueDate, "")
	}
	if r.orderDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "orderDate", r.orderDate, "")
	}
	if r.orderFromDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "orderFromDate", r.orderFromDate, "")
	}
	if r.orderToDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "orderToDate", r.orderToDate, "")
	}
	if r.orderNumber != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "orderNumber", r.orderNumber, "")
	}
	if r.deliveryNumber != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "DeliveryNumber", r.deliveryNumber, "")
	}
	if r.invoiceNumber != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "invoiceNumber", r.invoiceNumber, "")
	}
	if r.invoiceStatus != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "invoiceStatus", r.invoiceStatus, "")
	}
	if r.invoiceType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "invoiceType", r.invoiceType, "")
	}
	if r.customerOrderNumber != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "customerOrderNumber", r.customerOrderNumber, "")
	}
	if r.endCustomerOrderNumber != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "endCustomerOrderNumber", r.endCustomerOrderNumber, "")
	}
	if r.specialBidNumber != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "specialBidNumber", r.specialBidNumber, "")
	}
	if r.invoiceFromDueDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "invoiceFromDueDate", r.invoiceFromDueDate, "")
	}
	if r.invoiceToDueDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "invoiceToDueDate", r.invoiceToDueDate, "")
	}
	if r.invoiceFromDate != nil {
		t := *r.invoiceFromDate
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "invoiceFromDate", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "invoiceFromDate", t, "multi")
		}
	}
	if r.invoiceToDate != nil {
		t := *r.invoiceToDate
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "invoiceToDate", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "invoiceToDate", t, "multi")
		}
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageSize", r.pageSize, "")
	}
	if r.pageNumber != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageNumber", r.pageNumber, "")
	}
	if r.orderby != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "orderby", r.orderby, "")
	}
	if r.direction != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "direction", r.direction, "")
	}
	if r.serialNumber != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "serialNumber", r.serialNumber, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "IM-ApplicationID", r.iMApplicationID, "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "IM-CustomerNumber", r.iMCustomerNumber, "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "IM-CountryCode", r.iMCountryCode, "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "IM-CorrelationID", r.iMCorrelationID, "")
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
