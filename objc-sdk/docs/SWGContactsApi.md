# SWGContactsApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**createAccountExtensionContact**](SWGContactsApi.md#createaccountextensioncontact) | **POST** /accounts/{account_id}/extensions/{extension_id}/contacts | Add a new address book contact for an extension
[**deleteAccountExtensionContact**](SWGContactsApi.md#deleteaccountextensioncontact) | **DELETE** /accounts/{account_id}/extensions/{extension_id}/contacts/{contact_id} | 
[**getAccountExtensionContact**](SWGContactsApi.md#getaccountextensioncontact) | **GET** /accounts/{account_id}/extensions/{extension_id}/contacts/{contact_id} | Retrieve the details of an address book contact
[**listAccountExtensionContacts**](SWGContactsApi.md#listaccountextensioncontacts) | **GET** /accounts/{account_id}/extensions/{extension_id}/contacts | Show a list of address book contacts
[**replaceAccountExtensionContact**](SWGContactsApi.md#replaceaccountextensioncontact) | **PUT** /accounts/{account_id}/extensions/{extension_id}/contacts | 


# **createAccountExtensionContact**
```objc
-(NSNumber*) createAccountExtensionContactWithAccountId: (NSNumber*) accountId
    extensionId: (NSNumber*) extensionId
    data: (SWGCreateContactParams*) data
        completionHandler: (void (^)(SWGContactFull* output, NSError* error)) handler;
```

Add a new address book contact for an extension

For more on the input fields, see Account Contacts.

### Example 
```objc
SWGConfiguration *apiConfig = [SWGConfiguration sharedConfig];

// Configure API key authorization: (authentication scheme: apiKey)
[apiConfig setApiKey:@"YOUR_API_KEY" forApiKeyIdentifier:@"Authorization"];
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
//[apiConfig setApiKeyPrefix:@"Bearer" forApiKeyIdentifier:@"Authorization"];


NSNumber* accountId = @56; // Account ID
NSNumber* extensionId = @56; // Extension ID
SWGCreateContactParams* data = [[SWGCreateContactParams alloc] init]; // Contact data (optional)

SWGContactsApi*apiInstance = [[SWGContactsApi alloc] init];

// Add a new address book contact for an extension
[apiInstance createAccountExtensionContactWithAccountId:accountId
              extensionId:extensionId
              data:data
          completionHandler: ^(SWGContactFull* output, NSError* error) {
                        if (output) {
                            NSLog(@"%@", output);
                        }
                        if (error) {
                            NSLog(@"Error calling SWGContactsApi->createAccountExtensionContact: %@", error);
                        }
                    }];
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **NSNumber***| Account ID | 
 **extensionId** | **NSNumber***| Extension ID | 
 **data** | [**SWGCreateContactParams***](SWGCreateContactParams*.md)| Contact data | [optional] 

### Return type

[**SWGContactFull***](SWGContactFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **deleteAccountExtensionContact**
```objc
-(NSNumber*) deleteAccountExtensionContactWithAccountId: (NSNumber*) accountId
    extensionId: (NSNumber*) extensionId
    contactId: (NSNumber*) contactId
        completionHandler: (void (^)(SWGDeleteContact* output, NSError* error)) handler;
```





### Example 
```objc
SWGConfiguration *apiConfig = [SWGConfiguration sharedConfig];

// Configure API key authorization: (authentication scheme: apiKey)
[apiConfig setApiKey:@"YOUR_API_KEY" forApiKeyIdentifier:@"Authorization"];
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
//[apiConfig setApiKeyPrefix:@"Bearer" forApiKeyIdentifier:@"Authorization"];


NSNumber* accountId = @56; // Account ID
NSNumber* extensionId = @56; // Extension ID
NSNumber* contactId = @56; // Contact ID

SWGContactsApi*apiInstance = [[SWGContactsApi alloc] init];

// 
[apiInstance deleteAccountExtensionContactWithAccountId:accountId
              extensionId:extensionId
              contactId:contactId
          completionHandler: ^(SWGDeleteContact* output, NSError* error) {
                        if (output) {
                            NSLog(@"%@", output);
                        }
                        if (error) {
                            NSLog(@"Error calling SWGContactsApi->deleteAccountExtensionContact: %@", error);
                        }
                    }];
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **NSNumber***| Account ID | 
 **extensionId** | **NSNumber***| Extension ID | 
 **contactId** | **NSNumber***| Contact ID | 

### Return type

[**SWGDeleteContact***](SWGDeleteContact.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **getAccountExtensionContact**
```objc
-(NSNumber*) getAccountExtensionContactWithAccountId: (NSNumber*) accountId
    extensionId: (NSNumber*) extensionId
    contactId: (NSNumber*) contactId
        completionHandler: (void (^)(SWGContactFull* output, NSError* error)) handler;
```

Retrieve the details of an address book contact

For more info on the fields shown, see Account Contacts.

### Example 
```objc
SWGConfiguration *apiConfig = [SWGConfiguration sharedConfig];

// Configure API key authorization: (authentication scheme: apiKey)
[apiConfig setApiKey:@"YOUR_API_KEY" forApiKeyIdentifier:@"Authorization"];
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
//[apiConfig setApiKeyPrefix:@"Bearer" forApiKeyIdentifier:@"Authorization"];


NSNumber* accountId = @56; // Account ID
NSNumber* extensionId = @56; // Extension ID
NSNumber* contactId = @56; // Contact ID

SWGContactsApi*apiInstance = [[SWGContactsApi alloc] init];

// Retrieve the details of an address book contact
[apiInstance getAccountExtensionContactWithAccountId:accountId
              extensionId:extensionId
              contactId:contactId
          completionHandler: ^(SWGContactFull* output, NSError* error) {
                        if (output) {
                            NSLog(@"%@", output);
                        }
                        if (error) {
                            NSLog(@"Error calling SWGContactsApi->getAccountExtensionContact: %@", error);
                        }
                    }];
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **NSNumber***| Account ID | 
 **extensionId** | **NSNumber***| Extension ID | 
 **contactId** | **NSNumber***| Contact ID | 

### Return type

[**SWGContactFull***](SWGContactFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **listAccountExtensionContacts**
```objc
-(NSNumber*) listAccountExtensionContactsWithAccountId: (NSNumber*) accountId
    extensionId: (NSNumber*) extensionId
    filtersId: (NSArray<NSString*>*) filtersId
    filtersGroupId: (NSArray<NSString*>*) filtersGroupId
    filtersUpdatedAt: (NSArray<NSString*>*) filtersUpdatedAt
    sortId: (NSString*) sortId
    sortUpdatedAt: (NSString*) sortUpdatedAt
    limit: (NSNumber*) limit
    offset: (NSNumber*) offset
    fields: (NSString*) fields
        completionHandler: (void (^)(SWGListContactsFull* output, NSError* error)) handler;
```

Show a list of address book contacts

See Account Contacts for more info on the fields in each item.

### Example 
```objc
SWGConfiguration *apiConfig = [SWGConfiguration sharedConfig];

// Configure API key authorization: (authentication scheme: apiKey)
[apiConfig setApiKey:@"YOUR_API_KEY" forApiKeyIdentifier:@"Authorization"];
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
//[apiConfig setApiKeyPrefix:@"Bearer" forApiKeyIdentifier:@"Authorization"];


NSNumber* accountId = @56; // Account ID
NSNumber* extensionId = @56; // Extension ID
NSArray<NSString*>* filtersId = @[@"filtersId_example"]; // ID filter (optional)
NSArray<NSString*>* filtersGroupId = @[@"filtersGroupId_example"]; // Group filter (optional)
NSArray<NSString*>* filtersUpdatedAt = @[@"filtersUpdatedAt_example"]; // Updated At filter (optional)
NSString* sortId = @"sortId_example"; // ID sorting (optional)
NSString* sortUpdatedAt = @"sortUpdatedAt_example"; // Updated At sorting (optional)
NSNumber* limit = @56; // Max results (optional)
NSNumber* offset = @56; // Results to skip (optional)
NSString* fields = @"fields_example"; // Field set (optional)

SWGContactsApi*apiInstance = [[SWGContactsApi alloc] init];

// Show a list of address book contacts
[apiInstance listAccountExtensionContactsWithAccountId:accountId
              extensionId:extensionId
              filtersId:filtersId
              filtersGroupId:filtersGroupId
              filtersUpdatedAt:filtersUpdatedAt
              sortId:sortId
              sortUpdatedAt:sortUpdatedAt
              limit:limit
              offset:offset
              fields:fields
          completionHandler: ^(SWGListContactsFull* output, NSError* error) {
                        if (output) {
                            NSLog(@"%@", output);
                        }
                        if (error) {
                            NSLog(@"Error calling SWGContactsApi->listAccountExtensionContacts: %@", error);
                        }
                    }];
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **NSNumber***| Account ID | 
 **extensionId** | **NSNumber***| Extension ID | 
 **filtersId** | [**NSArray&lt;NSString*&gt;***](NSString*.md)| ID filter | [optional] 
 **filtersGroupId** | [**NSArray&lt;NSString*&gt;***](NSString*.md)| Group filter | [optional] 
 **filtersUpdatedAt** | [**NSArray&lt;NSString*&gt;***](NSString*.md)| Updated At filter | [optional] 
 **sortId** | **NSString***| ID sorting | [optional] 
 **sortUpdatedAt** | **NSString***| Updated At sorting | [optional] 
 **limit** | **NSNumber***| Max results | [optional] 
 **offset** | **NSNumber***| Results to skip | [optional] 
 **fields** | **NSString***| Field set | [optional] 

### Return type

[**SWGListContactsFull***](SWGListContactsFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **replaceAccountExtensionContact**
```objc
-(NSNumber*) replaceAccountExtensionContactWithAccountId: (NSNumber*) accountId
    extensionId: (NSNumber*) extensionId
    data: (SWGCreateContactParams*) data
        completionHandler: (void (^)(SWGContactFull* output, NSError* error)) handler;
```



For more on the input fields, see Account Contacts.

### Example 
```objc
SWGConfiguration *apiConfig = [SWGConfiguration sharedConfig];

// Configure API key authorization: (authentication scheme: apiKey)
[apiConfig setApiKey:@"YOUR_API_KEY" forApiKeyIdentifier:@"Authorization"];
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
//[apiConfig setApiKeyPrefix:@"Bearer" forApiKeyIdentifier:@"Authorization"];


NSNumber* accountId = @56; // Account ID
NSNumber* extensionId = @56; // Extension ID
SWGCreateContactParams* data = [[SWGCreateContactParams alloc] init]; // Contact data (optional)

SWGContactsApi*apiInstance = [[SWGContactsApi alloc] init];

// 
[apiInstance replaceAccountExtensionContactWithAccountId:accountId
              extensionId:extensionId
              data:data
          completionHandler: ^(SWGContactFull* output, NSError* error) {
                        if (output) {
                            NSLog(@"%@", output);
                        }
                        if (error) {
                            NSLog(@"Error calling SWGContactsApi->replaceAccountExtensionContact: %@", error);
                        }
                    }];
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **NSNumber***| Account ID | 
 **extensionId** | **NSNumber***| Extension ID | 
 **data** | [**SWGCreateContactParams***](SWGCreateContactParams*.md)| Contact data | [optional] 

### Return type

[**SWGContactFull***](SWGContactFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

