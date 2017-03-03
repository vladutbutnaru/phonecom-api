#import "SWGCreateContactParams.h"

@implementation SWGCreateContactParams

- (instancetype)init {
  self = [super init];
  if (self) {
    // initialize property's default value, if any
    
  }
  return self;
}


/**
 * Maps json key to property name.
 * This method is used by `JSONModel`.
 */
+ (JSONKeyMapper *)keyMapper {
  return [[JSONKeyMapper alloc] initWithDictionary:@{ @"first_name": @"firstName", @"middle_name": @"middleName", @"last_name": @"lastName", @"prefix": @"prefix", @"phonetic_first_name": @"phoneticFirstName", @"phonetic_middle_name": @"phoneticMiddleName", @"phonetic_last_name": @"phoneticLastName", @"suffix": @"suffix", @"nickname": @"nickname", @"company": @"company", @"department": @"department", @"job_title": @"jobTitle", @"emails": @"emails", @"phone_numbers": @"phoneNumbers", @"addresses": @"addresses", @"group": @"group" }];
}

/**
 * Indicates whether the property with the given name is optional.
 * If `propertyName` is optional, then return `YES`, otherwise return `NO`.
 * This method is used by `JSONModel`.
 */
+ (BOOL)propertyIsOptional:(NSString *)propertyName {

  NSArray *optionalProperties = @[@"firstName", @"middleName", @"lastName", @"prefix", @"phoneticFirstName", @"phoneticMiddleName", @"phoneticLastName", @"suffix", @"nickname", @"company", @"department", @"jobTitle", @"emails", @"phoneNumbers", @"addresses", @"group"];
  return [optionalProperties containsObject:propertyName];
}

@end
