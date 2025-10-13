# \AccesstokenAPI

All URIs are relative to *https://api.ingrammicro.com:443*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccesstoken**](AccesstokenAPI.md#GetAccesstoken) | **Get** /oauth/oauth20/token | Accesstoken



## GetAccesstoken

> AccesstokenResponse GetAccesstoken(ctx).GrantType(grantType).ClientId(clientId).ClientSecret(clientSecret).Execute()

Accesstoken



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
	grantType := "client_credentials" // string | Keep grant_type as client_credentials only.
	clientId := "clientId_example" // string | 
	clientSecret := "clientSecret_example" // string | 

	configuration := xi_sdk_resellers.NewConfiguration()
	apiClient := xi_sdk_resellers.NewAPIClient(configuration)
	resp, r, err := apiClient.AccesstokenAPI.GetAccesstoken(context.Background()).GrantType(grantType).ClientId(clientId).ClientSecret(clientSecret).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccesstokenAPI.GetAccesstoken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccesstoken`: AccesstokenResponse
	fmt.Fprintf(os.Stdout, "Response from `AccesstokenAPI.GetAccesstoken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccesstokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **grantType** | **string** | Keep grant_type as client_credentials only. | 
 **clientId** | **string** |  | 
 **clientSecret** | **string** |  | 

### Return type

[**AccesstokenResponse**](AccesstokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

