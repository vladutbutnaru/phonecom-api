# ExpressservicecodesApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**getAccountExpressSrvCode**](ExpressservicecodesApi.md#getAccountExpressSrvCode) | **GET** /accounts/{account_id}/express-service-codes/{code_id} | Show details of an account Express Service Code
[**listAccountExpressSrvCodes**](ExpressservicecodesApi.md#listAccountExpressSrvCodes) | **GET** /accounts/{account_id}/express-service-codes | Get the Express Service Code associated with your account in list format


<a name="getAccountExpressSrvCode"></a>
# **getAccountExpressSrvCode**
> ExpressServiceCodeFull getAccountExpressSrvCode(accountId, codeId)

Show details of an account Express Service Code

This service shows the details of an Account Express Service Code.

### Example
```java
// Import classes:
//import io.swagger.client.api.ExpressservicecodesApi;

ExpressservicecodesApi apiInstance = new ExpressservicecodesApi();
Integer accountId = 56; // Integer | Account ID
Integer codeId = 56; // Integer | Device ID
try {
    ExpressServiceCodeFull result = apiInstance.getAccountExpressSrvCode(accountId, codeId);
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling ExpressservicecodesApi#getAccountExpressSrvCode");
    e.printStackTrace();
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **Integer**| Account ID |
 **codeId** | **Integer**| Device ID |

### Return type

[**ExpressServiceCodeFull**](ExpressServiceCodeFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="listAccountExpressSrvCodes"></a>
# **listAccountExpressSrvCodes**
> ListExpressServiceCodesFull listAccountExpressSrvCodes(accountId, filtersId)

Get the Express Service Code associated with your account in list format

See Express Service Codes for more detail.

### Example
```java
// Import classes:
//import io.swagger.client.api.ExpressservicecodesApi;

ExpressservicecodesApi apiInstance = new ExpressservicecodesApi();
Integer accountId = 56; // Integer | Account ID
List<String> filtersId = Arrays.asList("filtersId_example"); // List<String> | ID filter
try {
    ListExpressServiceCodesFull result = apiInstance.listAccountExpressSrvCodes(accountId, filtersId);
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling ExpressservicecodesApi#listAccountExpressSrvCodes");
    e.printStackTrace();
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **Integer**| Account ID |
 **filtersId** | [**List&lt;String&gt;**](String.md)| ID filter | [optional]

### Return type

[**ListExpressServiceCodesFull**](ListExpressServiceCodesFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

