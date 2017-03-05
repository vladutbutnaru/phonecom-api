#import <Foundation/Foundation.h>
#import "SWGMediaFull.h"
#import "SWGListMediaFull.h"
#import "SWGApi.h"

/**
* Phone.com API
* This is a Phone.com api Swagger definition
*
* OpenAPI spec version: 1.0.0
* Contact: apisupport@phone.com
*
* NOTE: This class is auto generated by the swagger code generator program.
* https://github.com/swagger-api/swagger-codegen.git
* Do not edit the class manually.
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
* http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
*/


@interface SWGMediaApi: NSObject <SWGApi>

extern NSString* kSWGMediaApiErrorDomain;
extern NSInteger kSWGMediaApiMissingParamErrorCode;

+(instancetype) sharedAPI;

/// Show details of an individual media recording (Greeting or Hold Music)
/// Get individual media recording
///
/// @param accountId Account ID
/// @param recordingId Recording ID
/// 
///  code:200 message:"OK",
///  code:401 message:"Unauthorized access",
///  code:403 message:"Forbidden",
///  code:404 message:"Not Found"
///
/// @return SWGMediaFull*
-(NSNumber*) getAccountMediaWithAccountId: (NSNumber*) accountId
    recordingId: (NSNumber*) recordingId
    completionHandler: (void (^)(SWGMediaFull* output, NSError* error)) handler;


/// Get a list of media recordings for an account
/// See Account Menus for more info on the properties.
///
/// @param accountId Account ID
/// @param filtersId ID filter (optional)
/// @param filtersName Name filter (optional)
/// @param sortId ID sorting (optional)
/// @param sortName Name sorting (optional)
/// @param limit Max results (optional)
/// @param offset Results to skip (optional)
/// @param fields Field set (optional)
/// 
///  code:200 message:"OK",
///  code:401 message:"Unauthorized access",
///  code:403 message:"Forbidden"
///
/// @return SWGListMediaFull*
-(NSNumber*) listAccountMediaWithAccountId: (NSNumber*) accountId
    filtersId: (NSArray<NSString*>*) filtersId
    filtersName: (NSArray<NSString*>*) filtersName
    sortId: (NSString*) sortId
    sortName: (NSString*) sortName
    limit: (NSNumber*) limit
    offset: (NSNumber*) offset
    fields: (NSString*) fields
    completionHandler: (void (^)(SWGListMediaFull* output, NSError* error)) handler;



@end