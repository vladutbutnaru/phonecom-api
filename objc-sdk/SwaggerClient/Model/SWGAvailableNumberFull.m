#import "SWGAvailableNumberFull.h"

@implementation SWGAvailableNumberFull

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
  return [[JSONKeyMapper alloc] initWithDictionary:@{ @"phone_number": @"phoneNumber", @"formatted": @"formatted", @"price": @"price", @"is_toll_free": @"isTollFree", @"country_code": @"countryCode", @"npa": @"npa", @"nxx": @"nxx", @"xxxx": @"xxxx", @"city": @"city", @"province": @"province", @"country": @"country" }];
}

/**
 * Indicates whether the property with the given name is optional.
 * If `propertyName` is optional, then return `YES`, otherwise return `NO`.
 * This method is used by `JSONModel`.
 */
+ (BOOL)propertyIsOptional:(NSString *)propertyName {

  NSArray *optionalProperties = @[@"phoneNumber", @"formatted", @"price", @"isTollFree", @"countryCode", @"npa", @"nxx", @"xxxx", @"city", @"province", @"country"];
  return [optionalProperties containsObject:propertyName];
}

@end
