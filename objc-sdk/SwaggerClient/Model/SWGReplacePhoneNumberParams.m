#import "SWGReplacePhoneNumberParams.h"

@implementation SWGReplacePhoneNumberParams

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
  return [[JSONKeyMapper alloc] initWithDictionary:@{ @"route": @"route", @"name": @"name", @"block_incoming": @"blockIncoming", @"block_anonymous": @"blockAnonymous", @"caller_id[name]": @"callerIdName", @"caller_id[type]": @"callerIdType", @"sms_forwarding[type]": @"smsForwardingType", @"sms_forwarding[application]": @"smsForwardingApplication", @"sms_forwarding[extension]": @"smsForwardingExtension", @"pool_item": @"poolItem", @"call_notifications[emails]": @"callNotificationsEmails", @"call_notifications[sms]": @"callNotificationsSms" }];
}

/**
 * Indicates whether the property with the given name is optional.
 * If `propertyName` is optional, then return `YES`, otherwise return `NO`.
 * This method is used by `JSONModel`.
 */
+ (BOOL)propertyIsOptional:(NSString *)propertyName {

  NSArray *optionalProperties = @[@"route", @"name", @"blockIncoming", @"blockAnonymous", @"callerIdName", @"callerIdType", @"smsForwardingType", @"smsForwardingApplication", @"smsForwardingExtension", @"poolItem", @"callNotificationsEmails", @"callNotificationsSms"];
  return [optionalProperties containsObject:propertyName];
}

@end
