# CalleridsApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**getCallerIds**](CalleridsApi.md#getCallerIds) | **GET** /accounts/{account_id}/extensions/{extension_id}/caller-ids | Show the Caller ID options a given extension can use


<a name="getCallerIds"></a>
# **getCallerIds**
> ListCallerIdsFull getCallerIds(accountId, extensionId, filtersNumber, filtersName, sortNumber, sortName, limit, offset, fields)

Show the Caller ID options a given extension can use

Get Caller ID

### Example
```java
// Import classes:
//import io.swagger.client.api.CalleridsApi;

CalleridsApi apiInstance = new CalleridsApi();
Integer accountId = 56; // Integer | Account ID
Integer extensionId = 56; // Integer | Extension ID
List<String> filtersNumber = Arrays.asList("filtersNumber_example"); // List<String> | Number filter
List<String> filtersName = Arrays.asList("filtersName_example"); // List<String> | Name filter
String sortNumber = "sortNumber_example"; // String | Number sorting
String sortName = "sortName_example"; // String | Name sorting
Integer limit = 56; // Integer | Max results
Integer offset = 56; // Integer | Results to skip
String fields = "fields_example"; // String | Field set
try {
    ListCallerIdsFull result = apiInstance.getCallerIds(accountId, extensionId, filtersNumber, filtersName, sortNumber, sortName, limit, offset, fields);
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling CalleridsApi#getCallerIds");
    e.printStackTrace();
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **Integer**| Account ID |
 **extensionId** | **Integer**| Extension ID |
 **filtersNumber** | [**List&lt;String&gt;**](String.md)| Number filter | [optional]
 **filtersName** | [**List&lt;String&gt;**](String.md)| Name filter | [optional]
 **sortNumber** | **String**| Number sorting | [optional]
 **sortName** | **String**| Name sorting | [optional]
 **limit** | **Integer**| Max results | [optional]
 **offset** | **Integer**| Results to skip | [optional]
 **fields** | **String**| Field set | [optional]

### Return type

[**ListCallerIdsFull**](ListCallerIdsFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

