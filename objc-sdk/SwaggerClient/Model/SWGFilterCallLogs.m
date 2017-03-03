#import "SWGFilterCallLogs.h"

@implementation SWGFilterCallLogs

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
  return [[JSONKeyMapper alloc] initWithDictionary:@{ @"id": @"_id", @"start_time": @"startTime", @"created_at": @"createdAt", @"direction": @"direction", @"called_number": @"calledNumber", @"type": @"type" }];
}

/**
 * Indicates whether the property with the given name is optional.
 * If `propertyName` is optional, then return `YES`, otherwise return `NO`.
 * This method is used by `JSONModel`.
 */
+ (BOOL)propertyIsOptional:(NSString *)propertyName {

  NSArray *optionalProperties = @[@"_id", @"startTime", @"createdAt", @"direction", @"calledNumber", @"type"];
  return [optionalProperties containsObject:propertyName];
}

@end
