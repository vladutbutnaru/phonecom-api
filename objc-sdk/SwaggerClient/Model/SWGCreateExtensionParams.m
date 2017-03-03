#import "SWGCreateExtensionParams.h"

@implementation SWGCreateExtensionParams

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
  return [[JSONKeyMapper alloc] initWithDictionary:@{ @"caller_id": @"callerId", @"usage_type": @"usageType", @"allows_call_waiting": @"allowsCallWaiting", @"extension": @"extension", @"include_in_directory": @"includeInDirectory", @"name": @"name", @"full_name": @"fullName", @"timezone": @"timezone", @"name_greeting": @"nameGreeting", @"voicemail[greeting][alternate]": @"voicemailGreetingAlternate", @"local_area_code": @"localAreaCode", @"voicemail[greeting][enable_leave_message_prompt]": @"voicemailGreetingEnableLeaveMessagePrompt", @"voicemail[enabled]": @"voicemailEnabled", @"enable_outbound_calls": @"enableOutboundCalls", @"enable_call_waiting": @"enableCallWaiting", @"voicemail[password]": @"voicemailPassword", @"voicemail[greeting][type]": @"voicemailGreetingType", @"voicemail[greeting][standard]": @"voicemailGreetingStandard", @"voicemail[transcription]": @"voicemailTranscription", @"voicemail[notifications][emails]": @"voicemailNotificationsEmails", @"voicemail[notifications][sms]": @"voicemailNotificationsSms", @"call_notifications[emails]": @"callNotificationsEmails", @"call_notifications[sms]": @"callNotificationsSms" }];
}

/**
 * Indicates whether the property with the given name is optional.
 * If `propertyName` is optional, then return `YES`, otherwise return `NO`.
 * This method is used by `JSONModel`.
 */
+ (BOOL)propertyIsOptional:(NSString *)propertyName {

  NSArray *optionalProperties = @[@"callerId", @"usageType", @"allowsCallWaiting", @"extension", @"includeInDirectory", @"name", @"fullName", @"timezone", @"nameGreeting", @"voicemailGreetingAlternate", @"localAreaCode", @"voicemailGreetingEnableLeaveMessagePrompt", @"voicemailEnabled", @"enableOutboundCalls", @"enableCallWaiting", @"voicemailPassword", @"voicemailGreetingType", @"voicemailGreetingStandard", @"voicemailTranscription", @"voicemailNotificationsEmails", @"voicemailNotificationsSms", @"callNotificationsEmails", @"callNotificationsSms"];
  return [optionalProperties containsObject:propertyName];
}

@end
