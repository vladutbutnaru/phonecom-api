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



@protocol SWGAddress
@end

@interface SWGAddress : SWGObject

/* Street address line 1. Required. 
 */
@property(nonatomic) NSString* line1;
/* Street address line 2 [optional]
 */
@property(nonatomic) NSString* line2;
/* City. Required. 
 */
@property(nonatomic) NSString* city;
/* Province. Required if country is US [optional]
 */
@property(nonatomic) NSString* province;
/* Postal code [optional]
 */
@property(nonatomic) NSString* postalCode;
/* 2-character country code. Required. 
 */
@property(nonatomic) NSString* country;

@end
