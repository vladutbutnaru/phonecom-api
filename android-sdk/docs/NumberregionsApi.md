# NumberregionsApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**listAvailablePhoneNumberRegions**](NumberregionsApi.md#listAvailablePhoneNumberRegions) | **GET** /phone-numbers/available/regions | 


<a name="listAvailablePhoneNumberRegions"></a>
# **listAvailablePhoneNumberRegions**
> ListPhoneNumbersRegionsFull listAvailablePhoneNumberRegions(filtersCountryCode, filtersNpa, filtersNxx, filtersIsTollFree, filtersCity, filtersProvincePostalCode, filtersCountryPostalCode, sortCountryCode, sortNpa, sortNxx, sortIsTollFree, sortCity, sortProvincePostalCode, sortCountryPostalCode, limit, offset, fields, groupBy)



This service lists the quantities of available phone numbers by region.

### Example
```java
// Import classes:
//import io.swagger.client.api.NumberregionsApi;

NumberregionsApi apiInstance = new NumberregionsApi();
List<String> filtersCountryCode = Arrays.asList("filtersCountryCode_example"); // List<String> | Country Code filter
List<String> filtersNpa = Arrays.asList("filtersNpa_example"); // List<String> | Area Code filter (North America only)
List<String> filtersNxx = Arrays.asList("filtersNxx_example"); // List<String> | 2nd set of 3 digits filter (North America only)
List<String> filtersIsTollFree = Arrays.asList("filtersIsTollFree_example"); // List<String> | Toll-free status filter
List<String> filtersCity = Arrays.asList("filtersCity_example"); // List<String> | City filter
List<String> filtersProvincePostalCode = Arrays.asList("filtersProvincePostalCode_example"); // List<String> | State or Province filter
List<String> filtersCountryPostalCode = Arrays.asList("filtersCountryPostalCode_example"); // List<String> | Country filter
String sortCountryCode = "sortCountryCode_example"; // String | International calling code sorting
String sortNpa = "sortNpa_example"; // String | Area Code sorting (North America only)
String sortNxx = "sortNxx_example"; // String | 2nd set of 3 digits sorting (North America)
String sortIsTollFree = "sortIsTollFree_example"; // String | Toll Free status sorting
String sortCity = "sortCity_example"; // String | City sorting
String sortProvincePostalCode = "sortProvincePostalCode_example"; // String | State or Province sorting
String sortCountryPostalCode = "sortCountryPostalCode_example"; // String | Country sorting
Integer limit = 56; // Integer | Max results
Integer offset = 56; // Integer | Results to skip
String fields = "fields_example"; // String | Field set
List<String> groupBy = Arrays.asList("groupBy_example"); // List<String> | Fields to group by (supports the same set of fields as the filters and sorting)
try {
    ListPhoneNumbersRegionsFull result = apiInstance.listAvailablePhoneNumberRegions(filtersCountryCode, filtersNpa, filtersNxx, filtersIsTollFree, filtersCity, filtersProvincePostalCode, filtersCountryPostalCode, sortCountryCode, sortNpa, sortNxx, sortIsTollFree, sortCity, sortProvincePostalCode, sortCountryPostalCode, limit, offset, fields, groupBy);
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling NumberregionsApi#listAvailablePhoneNumberRegions");
    e.printStackTrace();
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filtersCountryCode** | [**List&lt;String&gt;**](String.md)| Country Code filter | [optional]
 **filtersNpa** | [**List&lt;String&gt;**](String.md)| Area Code filter (North America only) | [optional]
 **filtersNxx** | [**List&lt;String&gt;**](String.md)| 2nd set of 3 digits filter (North America only) | [optional]
 **filtersIsTollFree** | [**List&lt;String&gt;**](String.md)| Toll-free status filter | [optional]
 **filtersCity** | [**List&lt;String&gt;**](String.md)| City filter | [optional]
 **filtersProvincePostalCode** | [**List&lt;String&gt;**](String.md)| State or Province filter | [optional]
 **filtersCountryPostalCode** | [**List&lt;String&gt;**](String.md)| Country filter | [optional]
 **sortCountryCode** | **String**| International calling code sorting | [optional]
 **sortNpa** | **String**| Area Code sorting (North America only) | [optional]
 **sortNxx** | **String**| 2nd set of 3 digits sorting (North America) | [optional]
 **sortIsTollFree** | **String**| Toll Free status sorting | [optional]
 **sortCity** | **String**| City sorting | [optional]
 **sortProvincePostalCode** | **String**| State or Province sorting | [optional]
 **sortCountryPostalCode** | **String**| Country sorting | [optional]
 **limit** | **Integer**| Max results | [optional]
 **offset** | **Integer**| Results to skip | [optional]
 **fields** | **String**| Field set | [optional]
 **groupBy** | [**List&lt;String&gt;**](String.md)| Fields to group by (supports the same set of fields as the filters and sorting) | [optional]

### Return type

[**ListPhoneNumbersRegionsFull**](ListPhoneNumbersRegionsFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

