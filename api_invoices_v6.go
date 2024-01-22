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
)


// InvoicesV6APIService InvoicesV6API service
type InvoicesV6APIService service

type ApiGetInvoicedetailsV6Request struct {
	ctx context.Context
	ApiService *InvoicesV6APIService
	invoicenumber string
	version *string
	iMCustomerNumber *string
	iMCountryCode *string
	iMCorrelationID *string
	iMApplicationID *string
	customerType *string
	includeSerialNumbers *bool
}

// Version of codebase.
func (r ApiGetInvoicedetailsV6Request) Version(version string) ApiGetInvoicedetailsV6Request {
	r.version = &version
	return r
}

// Your unique Ingram Micro customer number.
func (r ApiGetInvoicedetailsV6Request) IMCustomerNumber(iMCustomerNumber string) ApiGetInvoicedetailsV6Request {
	r.iMCustomerNumber = &iMCustomerNumber
	return r
}

// Two-character ISO country code.
func (r ApiGetInvoicedetailsV6Request) IMCountryCode(iMCountryCode string) ApiGetInvoicedetailsV6Request {
	r.iMCountryCode = &iMCountryCode
	return r
}

// Unique transaction number to identify each transaction across all the systems.
func (r ApiGetInvoicedetailsV6Request) IMCorrelationID(iMCorrelationID string) ApiGetInvoicedetailsV6Request {
	r.iMCorrelationID = &iMCorrelationID
	return r
}

// Unique value used to identify the sender of the transaction. Example: MyCompany.
func (r ApiGetInvoicedetailsV6Request) IMApplicationID(iMApplicationID string) ApiGetInvoicedetailsV6Request {
	r.iMApplicationID = &iMApplicationID
	return r
}

// it should be invoice or order
func (r ApiGetInvoicedetailsV6Request) CustomerType(customerType string) ApiGetInvoicedetailsV6Request {
	r.customerType = &customerType
	return r
}

// if serial in the response send as true or else false
func (r ApiGetInvoicedetailsV6Request) IncludeSerialNumbers(includeSerialNumbers bool) ApiGetInvoicedetailsV6Request {
	r.includeSerialNumbers = &includeSerialNumbers
	return r
}

func (r ApiGetInvoicedetailsV6Request) Execute() (*InvoiceDetailResponse, *http.Response, error) {
	return r.ApiService.GetInvoicedetailsV6Execute(r)
}

/*
GetInvoicedetailsV6 Get Invoice Details v6

Use your Ingram Micro invoice number to search for existing invoices or retrieve existing invoice details.

The invoice number, IM-CustomerNumber, IM-CountryCode, IM-ApplicationId and IM-CorrelationID are required parameters.

.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param invoicenumber The Ingram Micro invoice number.
 @return ApiGetInvoicedetailsV6Request
*/
func (a *InvoicesV6APIService) GetInvoicedetailsV6(ctx context.Context, invoicenumber string) ApiGetInvoicedetailsV6Request {
	return ApiGetInvoicedetailsV6Request{
		ApiService: a,
		ctx: ctx,
		invoicenumber: invoicenumber,
	}
}

// Execute executes the request
//  @return InvoiceDetailResponse
func (a *InvoicesV6APIService) GetInvoicedetailsV6Execute(r ApiGetInvoicedetailsV6Request) (*InvoiceDetailResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InvoiceDetailResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InvoicesV6APIService.GetInvoicedetailsV6")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/resellers/v6/invoices/{invoicenumber}"
	localVarPath = strings.Replace(localVarPath, "{"+"invoicenumber"+"}", url.PathEscape(parameterValueToString(r.invoicenumber, "invoicenumber")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.invoicenumber) > 12 {
		return localVarReturnValue, nil, reportError("invoicenumber must have less than 12 elements")
	}
	if r.version == nil {
		return localVarReturnValue, nil, reportError("version is required and must be specified")
	}
	if strlen(*r.version) > 10 {
		return localVarReturnValue, nil, reportError("version must have less than 10 elements")
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
	parameterAddToHeaderOrQuery(localVarHeaderParams, "version", r.version, "")
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
