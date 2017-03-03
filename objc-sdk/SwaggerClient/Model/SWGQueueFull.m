#import "SWGQueueFull.h"

@implementation SWGQueueFull

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
  return [[JSONKeyMapper alloc] initWithDictionary:@{ @"id": @"_id", @"name": @"name", @"greeting": @"greeting", @"hold_music": @"holdMusic", @"max_hold_time": @"maxHoldTime", @"caller_id_type": @"callerIdType", @"ring_time": @"ringTime", @"members": @"members" }];
}

/**
 * Indicates whether the property with the given name is optional.
 * If `propertyName` is optional, then return `YES`, otherwise return `NO`.
 * This method is used by `JSONModel`.
 */
+ (BOOL)propertyIsOptional:(NSString *)propertyName {

  NSArray *optionalProperties = @[@"greeting", @"holdMusic", @"maxHoldTime", @"callerIdType", @"ringTime", @"members"];
  return [optionalProperties containsObject:propertyName];
}

@end
