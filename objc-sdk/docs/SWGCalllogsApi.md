# SWGCalllogsApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**getAccountCallLog**](SWGCalllogsApi.md#getaccountcalllog) | **GET** /accounts/{account_id}/call-logs/{callLog_id} | Show details of an individual Call Log entry
[**listAccountCallLogs**](SWGCalllogsApi.md#listaccountcalllogs) | **GET** /accounts/{account_id}/call-logs | Get a list of call details associated with your account


# **getAccountCallLog**
```objc
-(NSNumber*) getAccountCallLogWithAccountId: (NSNumber*) accountId
    callLogId: (NSNumber*) callLogId
        completionHandler: (void (^)(SWGCallLogFull* output, NSError* error)) handler;
```

Show details of an individual Call Log entry

This service shows the details of an individual Call Logs entry.

### Example 
```objc
SWGConfiguration *apiConfig = [SWGConfiguration sharedConfig];

// Configure API key authorization: (authentication scheme: apiKey)
[apiConfig setApiKey:@"YOUR_API_KEY" forApiKeyIdentifier:@"Authorization"];
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
//[apiConfig setApiKeyPrefix:@"Bearer" forApiKeyIdentifier:@"Authorization"];


NSNumber* accountId = @56; // Account ID
NSNumber* callLogId = @56; // Call Log ID

SWGCalllogsApi*apiInstance = [[SWGCalllogsApi alloc] init];

// Show details of an individual Call Log entry
[apiInstance getAccountCallLogWithAccountId:accountId
              callLogId:callLogId
          completionHandler: ^(SWGCallLogFull* output, NSError* error) {
                        if (output) {
                            NSLog(@"%@", output);
                        }
                        if (error) {
                            NSLog(@"Error calling SWGCalllogsApi->getAccountCallLog: %@", error);
                        }
                    }];
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **NSNumber***| Account ID | 
 **callLogId** | **NSNumber***| Call Log ID | 

### Return type

[**SWGCallLogFull***](SWGCallLogFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **listAccountCallLogs**
```objc
-(NSNumber*) listAccountCallLogsWithAccountId: (NSNumber*) accountId
    filtersId: (NSArray<NSString*>*) filtersId
    filtersStartTime: (NSArray<NSString*>*) filtersStartTime
    filtersCreatedAt: (NSString*) filtersCreatedAt
    filtersDirection: (NSString*) filtersDirection
    filtersCalledNumber: (NSString*) filtersCalledNumber
    filtersType: (NSString*) filtersType
    filtersExtension: (NSArray<NSString*>*) filtersExtension
    sortId: (NSString*) sortId
    sortStartTime: (NSString*) sortStartTime
    sortCreatedAt: (NSString*) sortCreatedAt
    limit: (NSNumber*) limit
    offset: (NSNumber*) offset
    fields: (NSString*) fields
        completionHandler: (void (^)(SWGListCallLogsFull* output, NSError* error)) handler;
```

Get a list of call details associated with your account

Fetch a list of call logs associated with your account.

### Example 
```objc
SWGConfiguration *apiConfig = [SWGConfiguration sharedConfig];

// Configure API key authorization: (authentication scheme: apiKey)
[apiConfig setApiKey:@"YOUR_API_KEY" forApiKeyIdentifier:@"Authorization"];
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
//[apiConfig setApiKeyPrefix:@"Bearer" forApiKeyIdentifier:@"Authorization"];


NSNumber* accountId = @56; // Account ID
NSArray<NSString*>* filtersId = @[@"filtersId_example"]; // ID filter (optional)
NSArray<NSString*>* filtersStartTime = @[@"filtersStartTime_example"]; // Call start time filter (optional)
NSString* filtersCreatedAt = @"filtersCreatedAt_example"; // Call log creation time filter (optional)
NSString* filtersDirection = @"filtersDirection_example"; // Call direction filter: in or out (optional)
NSString* filtersCalledNumber = @"filtersCalledNumber_example"; // Called number (optional)
NSString* filtersType = @"filtersType_example"; // Call type, such as 'call', 'fax', 'audiogram' (optional)
NSArray<NSString*>* filtersExtension = @[@"filtersExtension_example"]; // Extension filter (optional)
NSString* sortId = @"sortId_example"; // ID sorting (optional)
NSString* sortStartTime = @"sortStartTime_example"; // Sorting by call start time: asc or desc (optional)
NSString* sortCreatedAt = @"sortCreatedAt_example"; // Sorting by call log creation time: asc or desc (optional)
NSNumber* limit = @56; // Max results (optional)
NSNumber* offset = @56; // Results to skip (optional)
NSString* fields = @"fields_example"; // Field set (optional)

SWGCalllogsApi*apiInstance = [[SWGCalllogsApi alloc] init];

// Get a list of call details associated with your account
[apiInstance listAccountCallLogsWithAccountId:accountId
              filtersId:filtersId
              filtersStartTime:filtersStartTime
              filtersCreatedAt:filtersCreatedAt
              filtersDirection:filtersDirection
              filtersCalledNumber:filtersCalledNumber
              filtersType:filtersType
              filtersExtension:filtersExtension
              sortId:sortId
              sortStartTime:sortStartTime
              sortCreatedAt:sortCreatedAt
              limit:limit
              offset:offset
              fields:fields
          completionHandler: ^(SWGListCallLogsFull* output, NSError* error) {
                        if (output) {
                            NSLog(@"%@", output);
                        }
                        if (error) {
                            NSLog(@"Error calling SWGCalllogsApi->listAccountCallLogs: %@", error);
                        }
                    }];
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **NSNumber***| Account ID | 
 **filtersId** | [**NSArray&lt;NSString*&gt;***](NSString*.md)| ID filter | [optional] 
 **filtersStartTime** | [**NSArray&lt;NSString*&gt;***](NSString*.md)| Call start time filter | [optional] 
 **filtersCreatedAt** | **NSString***| Call log creation time filter | [optional] 
 **filtersDirection** | **NSString***| Call direction filter: in or out | [optional] 
 **filtersCalledNumber** | **NSString***| Called number | [optional] 
 **filtersType** | **NSString***| Call type, such as &#39;call&#39;, &#39;fax&#39;, &#39;audiogram&#39; | [optional] 
 **filtersExtension** | [**NSArray&lt;NSString*&gt;***](NSString*.md)| Extension filter | [optional] 
 **sortId** | **NSString***| ID sorting | [optional] 
 **sortStartTime** | **NSString***| Sorting by call start time: asc or desc | [optional] 
 **sortCreatedAt** | **NSString***| Sorting by call log creation time: asc or desc | [optional] 
 **limit** | **NSNumber***| Max results | [optional] 
 **offset** | **NSNumber***| Results to skip | [optional] 
 **fields** | **NSString***| Field set | [optional] 

### Return type

[**SWGListCallLogsFull***](SWGListCallLogsFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

