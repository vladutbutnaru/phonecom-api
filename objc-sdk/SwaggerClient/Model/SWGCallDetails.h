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



@protocol SWGCallDetails
@end

@interface SWGCallDetails : SWGObject

/* UNIX time stamp representing the UTC time that this call item starts [optional]
 */
@property(nonatomic) NSNumber* startTime;
/* Endpoint for this call item, such as call, voicemail, fax, menu, menu item, queue ... [optional]
 */
@property(nonatomic) NSString* type;
/* ID associated with this call item endpoint (type) [optional]
 */
@property(nonatomic) NSNumber* idValue;
/* Same as account id [optional]
 */
@property(nonatomic) NSNumber* voipId;
/* Same as account extension id [optional]
 */
@property(nonatomic) NSNumber* voipPhoneId;

@end