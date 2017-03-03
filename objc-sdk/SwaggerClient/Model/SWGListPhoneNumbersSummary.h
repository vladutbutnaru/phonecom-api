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

#import "SWGFilterIdNamePhoneNumberArray.h"
#import "SWGPhoneNumbersSummary.h"
#import "SWGSortIdNamePhoneNumber.h"


@protocol SWGListPhoneNumbersSummary
@end

@interface SWGListPhoneNumbersSummary : SWGObject


@property(nonatomic) SWGFilterIdNamePhoneNumberArray* filters;

@property(nonatomic) SWGSortIdNamePhoneNumber* sort;

@property(nonatomic) NSNumber* total;

@property(nonatomic) NSNumber* offset;

@property(nonatomic) NSNumber* limit;

@property(nonatomic) SWGPhoneNumbersSummary* items;

@end
