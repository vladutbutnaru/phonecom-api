#import "SWGReplaceMenuParams.h"

@implementation SWGReplaceMenuParams

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
  return [[JSONKeyMapper alloc] initWithDictionary:@{ @"name": @"name", @"greeting": @"greeting", @"invalid_keypress": @"invalidKeypress", @"allow_extension_dial": @"allowExtensionDial", @"keypress_wait_time": @"keypressWaitTime", @"timeout_handler": @"timeoutHandler", @"options": @"options" }];
}

/**
 * Indicates whether the property with the given name is optional.
 * If `propertyName` is optional, then return `YES`, otherwise return `NO`.
 * This method is used by `JSONModel`.
 */
+ (BOOL)propertyIsOptional:(NSString *)propertyName {

  NSArray *optionalProperties = @[@"name", @"greeting", @"invalidKeypress", @"allowExtensionDial", @"keypressWaitTime", @"timeoutHandler", @"options"];
  return [optionalProperties containsObject:propertyName];
}

@end
