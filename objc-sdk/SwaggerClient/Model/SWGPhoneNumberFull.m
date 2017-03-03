#import "SWGPhoneNumberFull.h"

@implementation SWGPhoneNumberFull

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
  return [[JSONKeyMapper alloc] initWithDictionary:@{ @"id": @"_id", @"name": @"name", @"phone_number": @"phoneNumber", @"block_incoming": @"blockIncoming", @"block_anonymous": @"blockAnonymous", @"route": @"route", @"caller_id": @"callerId", @"sms_forwarding": @"smsForwarding", @"call_notifications": @"callNotifications" }];
}

/**
 * Indicates whether the property with the given name is optional.
 * If `propertyName` is optional, then return `YES`, otherwise return `NO`.
 * This method is used by `JSONModel`.
 */
+ (BOOL)propertyIsOptional:(NSString *)propertyName {

  NSArray *optionalProperties = @[@"_id", @"name", @"phoneNumber", @"blockIncoming", @"blockAnonymous", @"route", @"callerId", @"smsForwarding", @"callNotifications"];
  return [optionalProperties containsObject:propertyName];
}

@end
