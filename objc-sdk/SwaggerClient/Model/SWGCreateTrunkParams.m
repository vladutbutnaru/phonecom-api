#import "SWGCreateTrunkParams.h"

@implementation SWGCreateTrunkParams

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
  return [[JSONKeyMapper alloc] initWithDictionary:@{ @"name": @"name", @"url": @"url", @"max_concurrent_calls": @"maxConcurrentCalls", @"max_minutes_per_month": @"maxMinutesPerMonth", @"greeting": @"greeting", @"error_message": @"errorMessage", @"codecs": @"codecs" }];
}

/**
 * Indicates whether the property with the given name is optional.
 * If `propertyName` is optional, then return `YES`, otherwise return `NO`.
 * This method is used by `JSONModel`.
 */
+ (BOOL)propertyIsOptional:(NSString *)propertyName {

  NSArray *optionalProperties = @[@"maxConcurrentCalls", @"maxMinutesPerMonth", @"greeting", @"errorMessage", @"codecs"];
  return [optionalProperties containsObject:propertyName];
}

@end
