#import "SWGContactFull.h"

@implementation SWGContactFull

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
  return [[JSONKeyMapper alloc] initWithDictionary:@{ @"id": @"_id", @"prefix": @"prefix", @"first_name": @"firstName", @"middle_name": @"middleName", @"last_name": @"lastName", @"suffix": @"suffix", @"nickname": @"nickname", @"company": @"company", @"phonetic_first_name": @"phoneticFirstName", @"phonetic_middle_name": @"phoneticMiddleName", @"phonetic_last_name": @"phoneticLastName", @"department": @"department", @"job_title": @"jobTitle", @"emails": @"emails", @"phone_numbers": @"phoneNumbers", @"addresses": @"addresses", @"group": @"group", @"created_at": @"createdAt", @"updated_at": @"updatedAt" }];
}

/**
 * Indicates whether the property with the given name is optional.
 * If `propertyName` is optional, then return `YES`, otherwise return `NO`.
 * This method is used by `JSONModel`.
 */
+ (BOOL)propertyIsOptional:(NSString *)propertyName {

  NSArray *optionalProperties = @[@"_id", @"prefix", @"firstName", @"middleName", @"lastName", @"suffix", @"nickname", @"company", @"phoneticFirstName", @"phoneticMiddleName", @"phoneticLastName", @"department", @"jobTitle", @"emails", @"phoneNumbers", @"addresses", @"group", @"createdAt", @"updatedAt"];
  return [optionalProperties containsObject:propertyName];
}

@end
