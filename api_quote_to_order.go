/*
XI Sdk Resellers

For Resellers. Who are looking to Innovate with Ingram Micro's API SolutionsAutomate your eCommerce with our offering of APIs and Webhooks to create a seamless experience for your customers.

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
)


// QuoteToOrderAPIService QuoteToOrderAPI service
type QuoteToOrderAPIService service

type ApiPostQuoteToOrderV6Request struct {
	ctx context.Context
	ApiService *QuoteToOrderAPIService
	iMCustomerNumber *string
	iMCountryCode *string
	iMCorrelationID *string
	quoteToOrderDetailsDTO *QuoteToOrderDetailsDTO
	iMSenderID *string
}

// Your unique Ingram Micro customer number.
func (r ApiPostQuoteToOrderV6Request) IMCustomerNumber(iMCustomerNumber string) ApiPostQuoteToOrderV6Request {
	r.iMCustomerNumber = &iMCustomerNumber
	return r
}

// Two-character ISO country code.
func (r ApiPostQuoteToOrderV6Request) IMCountryCode(iMCountryCode string) ApiPostQuoteToOrderV6Request {
	r.iMCountryCode = &iMCountryCode
	return r
}

// Unique transaction number to identify each transaction accross all the systems.
func (r ApiPostQuoteToOrderV6Request) IMCorrelationID(iMCorrelationID string) ApiPostQuoteToOrderV6Request {
	r.iMCorrelationID = &iMCorrelationID
	return r
}

func (r ApiPostQuoteToOrderV6Request) QuoteToOrderDetailsDTO(quoteToOrderDetailsDTO QuoteToOrderDetailsDTO) ApiPostQuoteToOrderV6Request {
	r.quoteToOrderDetailsDTO = &quoteToOrderDetailsDTO
	return r
}

// Unique value used to identify the sender of the transaction.
func (r ApiPostQuoteToOrderV6Request) IMSenderID(iMSenderID string) ApiPostQuoteToOrderV6Request {
	r.iMSenderID = &iMSenderID
	return r
}

func (r ApiPostQuoteToOrderV6Request) Execute() (*QuoteToOrderResponse, *http.Response, error) {
	return r.ApiService.PostQuoteToOrderV6Execute(r)
}

/*
PostQuoteToOrderV6 Quote To Order

The “Quote to Order” (QTO) endpoint allows a customer to create an order using the existing quote which is in “Ready to Order” status. A customer can create an order using Configure to order (CTO) quote or a non-configure to order (Non-CTO) quote. Upon successful submission of the order create request, a confirmation message will be returned as an API response. <br > <br >Ingram Micro offers webhooks as a method to send notifications to Resellers once the order creation request is received. All the updates related to Order creation will be pushed as a notification to the customer via a pre-defined callback URL as an HTTP post. <br > <br > **Prerequisite:** Pre-defined callback URL <br > <br > Before creating an order using the quote, it’s recommended to validate the quote using the “Validate Quote” endpoint. Validate Quote endpoint will not only validate the quote but also outline all the mandatory fields required by the vendor at a header level and at the line level which a customer need to pass to the Quote To Order endpoint request.  For a detailed understanding of the “Validate Quote” endpoint, review the “Validate Quote” endpoint documentation. <br ><br > **How it works:** <br ><br > - The customer validates the quote with a quote number from Validate Quote endpoint. <br > - The customer copies all the mandatory fields required by the vendor and adds them to the QTO request body. <br > - The customer provides all the values for Vendor mandatory fields along with other required information for QTO to create an order. <br > - After the order creation request receipt acknowledgment from the QTO endpoint, all further order creation updates will be provided via webhook push notification.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPostQuoteToOrderV6Request
*/
func (a *QuoteToOrderAPIService) PostQuoteToOrderV6(ctx context.Context) ApiPostQuoteToOrderV6Request {
	return ApiPostQuoteToOrderV6Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return QuoteToOrderResponse
func (a *QuoteToOrderAPIService) PostQuoteToOrderV6Execute(r ApiPostQuoteToOrderV6Request) (*QuoteToOrderResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *QuoteToOrderResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "QuoteToOrderAPIService.PostQuoteToOrderV6")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/resellers/v6/q2o/orders"

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
	if r.quoteToOrderDetailsDTO == nil {
		return localVarReturnValue, nil, reportError("quoteToOrderDetailsDTO is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.quoteToOrderDetailsDTO
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
			var v PostQuoteToOrderV6400Response
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
