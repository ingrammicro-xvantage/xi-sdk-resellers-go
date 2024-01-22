/*
Reseller API

For Resellers. <br> Who are looking to Innovate with Ingram Micro's API SolutionsAutomate your eCommerce with our offering of APIs and Webhooks to create a seamless experience for your customers.

API version: 6.0
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


// QuotesAPIService QuotesAPI service
type QuotesAPIService service

type ApiGetQuotessearchV6Request struct {
	ctx context.Context
	ApiService *QuotesAPIService
	iMCustomerNumber *string
	iMCountryCode *string
	iMCustomerContact *string
	iMCorrelationID *string
	ingramOrderDateBt *[]string
	quoteNumber *string
	specialBidNumber *string
	endUserContact *string
	sortingOrder *string
	sortBy *string
	pageSize *int32
	pageNumber *int32
	vendorName *string
	quoteName *string
	status *string
	quoteCreateDateBt *string
	iMSenderID *string
}

// Your unique Ingram Micro customer number.
func (r ApiGetQuotessearchV6Request) IMCustomerNumber(iMCustomerNumber string) ApiGetQuotessearchV6Request {
	r.iMCustomerNumber = &iMCustomerNumber
	return r
}

// Two-character ISO country code.
func (r ApiGetQuotessearchV6Request) IMCountryCode(iMCountryCode string) ApiGetQuotessearchV6Request {
	r.iMCountryCode = &iMCountryCode
	return r
}

// Logged in Users email address contact.
func (r ApiGetQuotessearchV6Request) IMCustomerContact(iMCustomerContact string) ApiGetQuotessearchV6Request {
	r.iMCustomerContact = &iMCustomerContact
	return r
}

// Unique transaction number to identify each transaction across all the systems.
func (r ApiGetQuotessearchV6Request) IMCorrelationID(iMCorrelationID string) ApiGetQuotessearchV6Request {
	r.iMCorrelationID = &iMCorrelationID
	return r
}

// Search with start and end date(only 2 entries allowed).
func (r ApiGetQuotessearchV6Request) IngramOrderDateBt(ingramOrderDateBt []string) ApiGetQuotessearchV6Request {
	r.ingramOrderDateBt = &ingramOrderDateBt
	return r
}

// Unique identifier generated by Ingram Micros CRM specific to each quote.  When applying a filter to the quoteNumber and including a partial quote number in the filter, all quotes containing any information included in the filter can be retrieved as a subset of all available customer quotes.
func (r ApiGetQuotessearchV6Request) QuoteNumber(quoteNumber string) ApiGetQuotessearchV6Request {
	r.quoteNumber = &quoteNumber
	return r
}

// Special Pricing Bid Number, also referred to as a Dart Number by some vendors, is a unique identifier associated with vendor specific products and discounts.
func (r ApiGetQuotessearchV6Request) SpecialBidNumber(specialBidNumber string) ApiGetQuotessearchV6Request {
	r.specialBidNumber = &specialBidNumber
	return r
}

// End User Name is the end customer name that is associated with a quote in Ingram Micros CRM.
func (r ApiGetQuotessearchV6Request) EndUserContact(endUserContact string) ApiGetQuotessearchV6Request {
	r.endUserContact = &endUserContact
	return r
}

// Sort applies to the selected column (sortingColumnName) and may be specified in  Ascending (asc) or Descending (desc) order. The default sort is Descending (desc) - most recent first.
func (r ApiGetQuotessearchV6Request) SortingOrder(sortingOrder string) ApiGetQuotessearchV6Request {
	r.sortingOrder = &sortingOrder
	return r
}

// Refers to the column selected to apply the sorting criteria.  The default column is dateCreated and will sort by the most recently created quote first with the following in descending order.  The default filter retrieves quotes created within the last 30 days. Filtering allows user to select a specific column to sort: quoteNumber, createdDate, lastModifiedDate and expiryDate.
func (r ApiGetQuotessearchV6Request) SortBy(sortBy string) ApiGetQuotessearchV6Request {
	r.sortBy = &sortBy
	return r
}

// Number of records (quotes) to display per page in the quote list.  The default is 25, but may be decreased using the filter .
func (r ApiGetQuotessearchV6Request) PageSize(pageSize int32) ApiGetQuotessearchV6Request {
	r.pageSize = &pageSize
	return r
}

// Page index or page number for the list of quotes being returned.
func (r ApiGetQuotessearchV6Request) PageNumber(pageNumber int32) ApiGetQuotessearchV6Request {
	r.pageNumber = &pageNumber
	return r
}

// The name of the vendor.
func (r ApiGetQuotessearchV6Request) VendorName(vendorName string) ApiGetQuotessearchV6Request {
	r.vendorName = &vendorName
	return r
}

// The quote name was given by the customer while creating quote.
func (r ApiGetQuotessearchV6Request) QuoteName(quoteName string) ApiGetQuotessearchV6Request {
	r.quoteName = &quoteName
	return r
}

// The status of the quote.
func (r ApiGetQuotessearchV6Request) Status(status string) ApiGetQuotessearchV6Request {
	r.status = &status
	return r
}

// Search with start and end date(only 2 entries allowed).
func (r ApiGetQuotessearchV6Request) QuoteCreateDateBt(quoteCreateDateBt string) ApiGetQuotessearchV6Request {
	r.quoteCreateDateBt = &quoteCreateDateBt
	return r
}

// Unique value used to identify the sender of the transaction. Example: MyCompany
func (r ApiGetQuotessearchV6Request) IMSenderID(iMSenderID string) ApiGetQuotessearchV6Request {
	r.iMSenderID = &iMSenderID
	return r
}

func (r ApiGetQuotessearchV6Request) Execute() (*QuoteSearchResponse, *http.Response, error) {
	return r.ApiService.GetQuotessearchV6Execute(r)
}

/*
GetQuotessearchV6 Quote Search

The Quote Search API, by default, will retrieve quotes modified or created within the last 30 days. Quotes older than 365 days are excluded by default. The date filters enable the retrieval of quotes older than 30 days and up to 365 days when using date range criteria. The Quote Search API enables the retrieval and filtering of relevant quote list key criteria data such as Quote Number, Special Bid Numbers, End User Name, Quote Status, and Date Ranges from Ingram Micros CRM system. Only Active quotes are avaiable through the API.  Draft and Closed quotes are excluded and are not accessable through the Quote List Search API.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetQuotessearchV6Request
*/
func (a *QuotesAPIService) GetQuotessearchV6(ctx context.Context) ApiGetQuotessearchV6Request {
	return ApiGetQuotessearchV6Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return QuoteSearchResponse
func (a *QuotesAPIService) GetQuotessearchV6Execute(r ApiGetQuotessearchV6Request) (*QuoteSearchResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *QuoteSearchResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "QuotesAPIService.GetQuotessearchV6")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/resellers/v6/quotes/search"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
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
	if r.iMCustomerContact == nil {
		return localVarReturnValue, nil, reportError("iMCustomerContact is required and must be specified")
	}
	if strlen(*r.iMCustomerContact) > 64 {
		return localVarReturnValue, nil, reportError("iMCustomerContact must have less than 64 elements")
	}
	if r.iMCorrelationID == nil {
		return localVarReturnValue, nil, reportError("iMCorrelationID is required and must be specified")
	}
	if strlen(*r.iMCorrelationID) > 32 {
		return localVarReturnValue, nil, reportError("iMCorrelationID must have less than 32 elements")
	}

	if r.ingramOrderDateBt != nil {
		t := *r.ingramOrderDateBt
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "ingramOrderDate-bt", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "ingramOrderDate-bt", t, "multi")
		}
	}
	if r.quoteNumber != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "quoteNumber", r.quoteNumber, "")
	}
	if r.specialBidNumber != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "specialBidNumber", r.specialBidNumber, "")
	}
	if r.endUserContact != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "endUserContact", r.endUserContact, "")
	}
	if r.sortingOrder != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sortingOrder", r.sortingOrder, "")
	}
	if r.sortBy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sortBy", r.sortBy, "")
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageSize", r.pageSize, "")
	}
	if r.pageNumber != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageNumber", r.pageNumber, "")
	} else {
		var defaultValue int32 = 1
		r.pageNumber = &defaultValue
	}
	if r.vendorName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "vendorName", r.vendorName, "")
	}
	if r.quoteName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "quoteName", r.quoteName, "")
	}
	if r.status != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "status", r.status, "")
	}
	if r.quoteCreateDateBt != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "quoteCreateDate-bt", r.quoteCreateDateBt, "")
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
	parameterAddToHeaderOrQuery(localVarHeaderParams, "IM-CustomerContact", r.iMCustomerContact, "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "IM-CorrelationID", r.iMCorrelationID, "")
	if r.iMSenderID != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "IM-SenderID", r.iMSenderID, "")
	}
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
		if localVarHTTPResponse.StatusCode == 404 {
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

type ApiGetResellerV6ValidateQuoteRequest struct {
	ctx context.Context
	ApiService *QuotesAPIService
	quoteNumber *string
	iMCustomerNumber *string
	iMCountryCode *string
	iMCorrelationID *string
	iMSenderID *string
}

// A unique identifier generated by Ingram Micro&#39;s CRM specific to each quote.
func (r ApiGetResellerV6ValidateQuoteRequest) QuoteNumber(quoteNumber string) ApiGetResellerV6ValidateQuoteRequest {
	r.quoteNumber = &quoteNumber
	return r
}

// Your Ingram Micro unique customer number.
func (r ApiGetResellerV6ValidateQuoteRequest) IMCustomerNumber(iMCustomerNumber string) ApiGetResellerV6ValidateQuoteRequest {
	r.iMCustomerNumber = &iMCustomerNumber
	return r
}

// Two-character ISO country code.
func (r ApiGetResellerV6ValidateQuoteRequest) IMCountryCode(iMCountryCode string) ApiGetResellerV6ValidateQuoteRequest {
	r.iMCountryCode = &iMCountryCode
	return r
}

// Unique transaction number to identify each transaction accross all the systems.
func (r ApiGetResellerV6ValidateQuoteRequest) IMCorrelationID(iMCorrelationID string) ApiGetResellerV6ValidateQuoteRequest {
	r.iMCorrelationID = &iMCorrelationID
	return r
}

// Unique identifier used to identify the third party source accessing the services.
func (r ApiGetResellerV6ValidateQuoteRequest) IMSenderID(iMSenderID string) ApiGetResellerV6ValidateQuoteRequest {
	r.iMSenderID = &iMSenderID
	return r
}

func (r ApiGetResellerV6ValidateQuoteRequest) Execute() (*ValidateQuoteResponse, *http.Response, error) {
	return r.ApiService.GetResellerV6ValidateQuoteExecute(r)
}

/*
GetResellerV6ValidateQuote Validate Quote

The validate quote helps the customer validate the quote created in Ingram Micro's system. Apart from validating the quote, the endpoint also identifies all the mandatory fields that are required by the vendor at the header and line levels.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetResellerV6ValidateQuoteRequest
*/
func (a *QuotesAPIService) GetResellerV6ValidateQuote(ctx context.Context) ApiGetResellerV6ValidateQuoteRequest {
	return ApiGetResellerV6ValidateQuoteRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ValidateQuoteResponse
func (a *QuotesAPIService) GetResellerV6ValidateQuoteExecute(r ApiGetResellerV6ValidateQuoteRequest) (*ValidateQuoteResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ValidateQuoteResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "QuotesAPIService.GetResellerV6ValidateQuote")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/resellers/v6/q2o/validatequote"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.quoteNumber == nil {
		return localVarReturnValue, nil, reportError("quoteNumber is required and must be specified")
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

	parameterAddToHeaderOrQuery(localVarQueryParams, "quoteNumber", r.quoteNumber, "")
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
	if r.iMSenderID != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "IM-SenderID", r.iMSenderID, "")
	}
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
			var v GetResellerV6ValidateQuote400Response
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
			var v GetResellerV6ValidateQuote500Response
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

type ApiGetResellersV6QuotesRequest struct {
	ctx context.Context
	ApiService *QuotesAPIService
	iMCustomerNumber *string
	iMCountryCode *string
	iMCorrelationID *string
	quoteNumber string
	iMSenderID *string
}

// Your Ingram Micro unique customer number
func (r ApiGetResellersV6QuotesRequest) IMCustomerNumber(iMCustomerNumber string) ApiGetResellersV6QuotesRequest {
	r.iMCustomerNumber = &iMCustomerNumber
	return r
}

// Two-character ISO country code.
func (r ApiGetResellersV6QuotesRequest) IMCountryCode(iMCountryCode string) ApiGetResellersV6QuotesRequest {
	r.iMCountryCode = &iMCountryCode
	return r
}

// Unique transaction number to identify each transaction accross all the systems.
func (r ApiGetResellersV6QuotesRequest) IMCorrelationID(iMCorrelationID string) ApiGetResellersV6QuotesRequest {
	r.iMCorrelationID = &iMCorrelationID
	return r
}

// Unique identifier used to identify the third party source accessing the services.
func (r ApiGetResellersV6QuotesRequest) IMSenderID(iMSenderID string) ApiGetResellersV6QuotesRequest {
	r.iMSenderID = &iMSenderID
	return r
}

func (r ApiGetResellersV6QuotesRequest) Execute() (*QuoteDetailsResponse, *http.Response, error) {
	return r.ApiService.GetResellersV6QuotesExecute(r)
}

/*
GetResellersV6Quotes Get Quote Details

The quote details API provides all quote details associated with the quote number provided.

The **“quoteNumber”**, **“isoCountryCode”** and **“customerNumber”** parameters are required. 

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param quoteNumber Unique identifier generated by Ingram Micro's CRM specific to each quote.  When applying a filter to the quoteNumber and including a partial quote number in the filter, all quotes containing any information included in the filter can be retrieved as a subset of all available customer quotes.
 @return ApiGetResellersV6QuotesRequest
*/
func (a *QuotesAPIService) GetResellersV6Quotes(ctx context.Context, quoteNumber string) ApiGetResellersV6QuotesRequest {
	return ApiGetResellersV6QuotesRequest{
		ApiService: a,
		ctx: ctx,
		quoteNumber: quoteNumber,
	}
}

// Execute executes the request
//  @return QuoteDetailsResponse
func (a *QuotesAPIService) GetResellersV6QuotesExecute(r ApiGetResellersV6QuotesRequest) (*QuoteDetailsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *QuoteDetailsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "QuotesAPIService.GetResellersV6Quotes")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/resellers/v6/quotes/{quoteNumber}"
	localVarPath = strings.Replace(localVarPath, "{"+"quoteNumber"+"}", url.PathEscape(parameterValueToString(r.quoteNumber, "quoteNumber")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
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
	if r.iMSenderID != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "IM-SenderID", r.iMSenderID, "")
	}
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
