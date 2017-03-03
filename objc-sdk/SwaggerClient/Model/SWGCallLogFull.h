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

#import "SWGCallDetails.h"
#import "SWGExtensionSummary.h"


@protocol SWGCallLogFull
@end

@interface SWGCallLogFull : SWGObject

/* ID [optional]
 */
@property(nonatomic) NSString* _id;
/* Internal system id, may be null [optional]
 */
@property(nonatomic) NSString* uuid;
/* Account extension [optional]
 */
@property(nonatomic) SWGExtensionSummary* extension;
/* Call made from this phone number [optional]
 */
@property(nonatomic) NSString* callerId;
/* Call made to this phone number [optional]
 */
@property(nonatomic) NSString* calledNumber;
/* Call start time [optional]
 */
@property(nonatomic) NSString* startTime;
/* Call log creation time. Same as call end time + time needed to create call log [optional]
 */
@property(nonatomic) NSString* createdAt;
/* Call direction: in or out [optional]
 */
@property(nonatomic) NSString* direction;
/* Call type: call, fax, audiogram ... [optional]
 */
@property(nonatomic) NSString* type;
/* Call duration in seconds [optional]
 */
@property(nonatomic) NSNumber* callDuration;
/* Was call being monitored? [optional]
 */
@property(nonatomic) NSString* isMonitored;
/* Internal system call reference number [optional]
 */
@property(nonatomic) NSString* callNumber;
/* Last action of call flow [optional]
 */
@property(nonatomic) NSString* finalAction;
/* URL of call recording if available. Empty string if call recording does not exist [optional]
 */
@property(nonatomic) NSString* callRecording;
/* A list of call flows from beginning of call to end of call. [optional]
 */
@property(nonatomic) NSArray<SWGCallDetails>* details;
/* Internal system caller id / name [optional]
 */
@property(nonatomic) NSString* callerCnam;

@end
