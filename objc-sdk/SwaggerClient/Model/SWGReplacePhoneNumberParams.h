#import <Foundation/Foundation.h>
#import "SWGObject.h"

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



@protocol SWGReplacePhoneNumberParams
@end

@interface SWGReplacePhoneNumberParams : SWGObject

/* Route lookup object [optional]
 */
@property(nonatomic) NSObject* route;
/* Phone Name [optional]
 */
@property(nonatomic) NSString* name;
/* Block incoming calls [optional]
 */
@property(nonatomic) NSNumber* blockIncoming;
/* Block anonymous calls [optional]
 */
@property(nonatomic) NSNumber* blockAnonymous;
/* Caller ID name [optional]
 */
@property(nonatomic) NSString* callerIdName;
/* Caller ID type [optional]
 */
@property(nonatomic) NSString* callerIdType;
/* 'application' or 'extension' [optional]
 */
@property(nonatomic) NSString* smsForwardingType;
/* Application lookup object [optional]
 */
@property(nonatomic) NSObject* smsForwardingApplication;
/* Extension lookup object [optional]
 */
@property(nonatomic) NSObject* smsForwardingExtension;
/* Pool lookup object [optional]
 */
@property(nonatomic) NSObject* poolItem;
/* Call notifications for emails. Can be a single email or an array of emails [optional]
 */
@property(nonatomic) NSArray<NSString*>* callNotificationsEmails;
/* Call notification for SMS [optional]
 */
@property(nonatomic) NSString* callNotificationsSms;

@end
