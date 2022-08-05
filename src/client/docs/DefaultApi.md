# \DefaultApi

All URIs are relative to *http://127.0.0.1:8555/api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Read**](DefaultApi.md#Read) | **Get** /read | Read Basin resource
[**Register**](DefaultApi.md#Register) | **Post** /register | Register Basin resource
[**Subscribe**](DefaultApi.md#Subscribe) | **Post** /subscribe | Request subscription to Basin resource
[**Write**](DefaultApi.md#Write) | **Put** /write | Write Basin resource



## Read

> string Read(ctx).Url(url).Execute()

Read Basin resource



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    url := "url_example" // string | URL of Basin resource (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.Read(context.Background()).Url(url).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.Read``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Read`: string
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.Read`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **url** | **string** | URL of Basin resource | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Register

> bool Register(ctx).RegisterRequest(registerRequest).Execute()

Register Basin resource



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    registerRequest := *openapiclient.NewRegisterRequest("Url_example", []openapiclient.PermissionJson{*openapiclient.NewPermissionJson()}, *openapiclient.NewAdapterJson("AdapterName_example", map[string]interface{}(123)), map[string]interface{}(123)) // RegisterRequest | Register Basin Resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.Register(context.Background()).RegisterRequest(registerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.Register``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Register`: bool
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.Register`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **registerRequest** | [**RegisterRequest**](RegisterRequest.md) | Register Basin Resource | 

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Subscribe

> map[string]interface{} Subscribe(ctx).SubscribeRequest(subscribeRequest).Execute()

Request subscription to Basin resource



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    subscribeRequest := *openapiclient.NewSubscribeRequest("Url_example", []openapiclient.PermissionJson{*openapiclient.NewPermissionJson()}) // SubscribeRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.Subscribe(context.Background()).SubscribeRequest(subscribeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.Subscribe``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Subscribe`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.Subscribe`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscribeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscribeRequest** | [**SubscribeRequest**](SubscribeRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Write

> bool Write(ctx).WriteRequest(writeRequest).Execute()

Write Basin resource



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    writeRequest := *openapiclient.NewWriteRequest("Url_example", string(123)) // WriteRequest | Write Basin Resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.Write(context.Background()).WriteRequest(writeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.Write``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Write`: bool
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.Write`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWriteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **writeRequest** | [**WriteRequest**](WriteRequest.md) | Write Basin Resource | 

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

