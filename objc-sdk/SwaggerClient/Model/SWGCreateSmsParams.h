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



@protocol SWGCreateSmsParams
@end

@interface SWGCreateSmsParams : SWGObject

/* Phone number of sender 
 */
@property(nonatomic) NSString* from;
/* Outgoing destination numbers 
 */
@property(nonatomic) NSString* to;
/* Text body of the outgoing SMS message. Maximum 160 characters per message. 
 */
@property(nonatomic) NSString* text;

@end
