# SmsApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**createAccountSms**](SmsApi.md#createAccountSms) | **POST** /accounts/{account_id}/sms | Send a SMS to one or a group of recipients
[**getAccountSms**](SmsApi.md#getAccountSms) | **GET** /accounts/{account_id}/sms/{sms_id} | Show details of an individual SMS
[**listAccountSms**](SmsApi.md#listAccountSms) | **GET** /accounts/{account_id}/sms | Get a list of SMS messages for an account


<a name="createAccountSms"></a>
# **createAccountSms**
> SmsFull createAccountSms(accountId, data)

Send a SMS to one or a group of recipients

For more on the input fields, see Intro to SMS.

### Example
```java
// Import classes:
//import io.swagger.client.api.SmsApi;

SmsApi apiInstance = new SmsApi();
Integer accountId = 56; // Integer | Account ID
CreateSmsParams data = new CreateSmsParams(); // CreateSmsParams | SMS data
try {
    SmsFull result = apiInstance.createAccountSms(accountId, data);
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling SmsApi#createAccountSms");
    e.printStackTrace();
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **Integer**| Account ID |
 **data** | [**CreateSmsParams**](CreateSmsParams.md)| SMS data |

### Return type

[**SmsFull**](SmsFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="getAccountSms"></a>
# **getAccountSms**
> SmsFull getAccountSms(accountId, smsId)

Show details of an individual SMS

This service shows the details of an individual sms. See Intro to SMS for more info on the properties.

### Example
```java
// Import classes:
//import io.swagger.client.api.SmsApi;

SmsApi apiInstance = new SmsApi();
Integer accountId = 56; // Integer | Account ID
Integer smsId = 56; // Integer | SMS ID
try {
    SmsFull result = apiInstance.getAccountSms(accountId, smsId);
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling SmsApi#getAccountSms");
    e.printStackTrace();
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **Integer**| Account ID |
 **smsId** | **Integer**| SMS ID |

### Return type

[**SmsFull**](SmsFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="listAccountSms"></a>
# **listAccountSms**
> ListSmsFull listAccountSms(accountId, filtersId, filtersDirection, filtersFrom, sortId, sortCreatedAt, limit, offset, fields)

Get a list of SMS messages for an account

See Intro to SMS for more info on the properties.

### Example
```java
// Import classes:
//import io.swagger.client.api.SmsApi;

SmsApi apiInstance = new SmsApi();
Integer accountId = 56; // Integer | Account ID
List<String> filtersId = Arrays.asList("filtersId_example"); // List<String> | ID filter
String filtersDirection = "filtersDirection_example"; // String | Direction filter
String filtersFrom = "filtersFrom_example"; // String | Caller ID filter
String sortId = "sortId_example"; // String | ID sorting
String sortCreatedAt = "sortCreatedAt_example"; // String | Sort by created time of message
Integer limit = 56; // Integer | Max results
Integer offset = 56; // Integer | Results to skip
String fields = "fields_example"; // String | Field set
try {
    ListSmsFull result = apiInstance.listAccountSms(accountId, filtersId, filtersDirection, filtersFrom, sortId, sortCreatedAt, limit, offset, fields);
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling SmsApi#listAccountSms");
    e.printStackTrace();
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **Integer**| Account ID |
 **filtersId** | [**List&lt;String&gt;**](String.md)| ID filter | [optional]
 **filtersDirection** | **String**| Direction filter | [optional]
 **filtersFrom** | **String**| Caller ID filter | [optional]
 **sortId** | **String**| ID sorting | [optional]
 **sortCreatedAt** | **String**| Sort by created time of message | [optional]
 **limit** | **Integer**| Max results | [optional]
 **offset** | **Integer**| Results to skip | [optional]
 **fields** | **String**| Field set | [optional]

### Return type

[**ListSmsFull**](ListSmsFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

