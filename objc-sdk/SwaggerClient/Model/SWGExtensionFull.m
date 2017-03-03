#import "SWGExtensionFull.h"

@implementation SWGExtensionFull

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
  return [[JSONKeyMapper alloc] initWithDictionary:@{ @"id": @"_id", @"name": @"name", @"extension": @"extension", @"full_name": @"fullName", @"usage_type": @"usageType", @"device_membership": @"deviceMembership", @"timezone": @"timezone", @"name_greeting": @"nameGreeting", @"include_in_directory": @"includeInDirectory", @"caller_id": @"callerId", @"local_area_code": @"localAreaCode", @"enable_call_waiting": @"enableCallWaiting", @"enable_outbound_calls": @"enableOutboundCalls", @"voicemail": @"voicemail", @"call_notifications": @"callNotifications", @"route": @"route" }];
}

/**
 * Indicates whether the property with the given name is optional.
 * If `propertyName` is optional, then return `YES`, otherwise return `NO`.
 * This method is used by `JSONModel`.
 */
+ (BOOL)propertyIsOptional:(NSString *)propertyName {

  NSArray *optionalProperties = @[@"_id", @"name", @"extension", @"fullName", @"usageType", @"deviceMembership", @"timezone", @"nameGreeting", @"includeInDirectory", @"callerId", @"localAreaCode", @"enableCallWaiting", @"enableOutboundCalls", @"voicemail", @"callNotifications", @"route"];
  return [optionalProperties containsObject:propertyName];
}

@end
