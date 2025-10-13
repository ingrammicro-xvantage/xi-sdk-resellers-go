# \StockUpdateAPI

All URIs are relative to *https://api.ingrammicro.com:443*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ResellersV1WebhooksAvailabilityupdatePost**](StockUpdateAPI.md#ResellersV1WebhooksAvailabilityupdatePost) | **Post** /resellers/v1/webhooks/availabilityupdate | Stock Update



## ResellersV1WebhooksAvailabilityupdatePost

> ResellersV1WebhooksAvailabilityupdatePost(ctx).Targeturl(targeturl).XHubSignature(xHubSignature).AvailabilityAsyncNotificationRequest(availabilityAsyncNotificationRequest).Execute()

Stock Update

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
	targeturl := "https://59a2dc5368073ab42fd9a92e210a9fdb.m.pipedream.net/" // string | The webhook url where the request needs to sent.
	xHubSignature := "3LeaTfLE5FLj1FcYflwdwFosH4ADHmMbds6thtirGC3e9lEkF9/1pt4T2fQQGlxf40EznDBER0b60M75K6ZW0A==" // string | Ingram Micro creates a signature token by use of a secret key + Event ID. The algorithm to generate the secret ley is given at link https://developer.ingrammicro.com/reseller/article/how-use-webhook-secret-key. Use the event Id in the below sample along with your secret key to generate the key. Alternatively, to send try this out, use a random text to see how it works.
	availabilityAsyncNotificationRequest := *openapiclient.NewAvailabilityAsyncNotificationRequest() // AvailabilityAsyncNotificationRequest | 

	configuration := xi_sdk_resellers.NewConfiguration()
	apiClient := xi_sdk_resellers.NewAPIClient(configuration)
	r, err := apiClient.StockUpdateAPI.ResellersV1WebhooksAvailabilityupdatePost(context.Background()).Targeturl(targeturl).XHubSignature(xHubSignature).AvailabilityAsyncNotificationRequest(availabilityAsyncNotificationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StockUpdateAPI.ResellersV1WebhooksAvailabilityupdatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResellersV1WebhooksAvailabilityupdatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **targeturl** | **string** | The webhook url where the request needs to sent. | 
 **xHubSignature** | **string** | Ingram Micro creates a signature token by use of a secret key + Event ID. The algorithm to generate the secret ley is given at link https://developer.ingrammicro.com/reseller/article/how-use-webhook-secret-key. Use the event Id in the below sample along with your secret key to generate the key. Alternatively, to send try this out, use a random text to see how it works. | 
 **availabilityAsyncNotificationRequest** | [**AvailabilityAsyncNotificationRequest**](AvailabilityAsyncNotificationRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[application](../README.md#application)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

