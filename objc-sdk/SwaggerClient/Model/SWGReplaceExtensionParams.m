#import "SWGReplaceExtensionParams.h"

@implementation SWGReplaceExtensionParams

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
  return [[JSONKeyMapper alloc] initWithDictionary:@{ @"voicemail[greeting][alternate]": @"voicemailGreetingAlternate", @"name_greeting": @"nameGreeting", @"name": @"name", @"timezone": @"timezone", @"include_in_directory": @"includeInDirectory", @"extension": @"extension", @"enable_outbound_calls": @"enableOutboundCalls", @"usage_type": @"usageType", @"voicemail[password]": @"voicemailPassword", @"full_name": @"fullName", @"enable_call_waiting": @"enableCallWaiting", @"voicemail[greeting][standard]": @"voicemailGreetingStandard", @"voicemail[greeting][type]": @"voicemailGreetingType", @"caller_id": @"callerId", @"local_area_code": @"localAreaCode", @"voicemail[enabled]": @"voicemailEnabled", @"voicemail[greeting][enable_leave_message_prompt]": @"voicemailGreetingEnableLeaveMessagePrompt", @"voicemail[transcription]": @"voicemailTranscription", @"voicemail[notifications][emails]": @"voicemailNotificationsEmails", @"voicemail[notifications][sms]": @"voicemailNotificationsSms", @"call_notifications[emails]": @"callNotificationsEmails", @"call_notifications[sms]": @"callNotificationsSms", @"route": @"route" }];
}

/**
 * Indicates whether the property with the given name is optional.
 * If `propertyName` is optional, then return `YES`, otherwise return `NO`.
 * This method is used by `JSONModel`.
 */
+ (BOOL)propertyIsOptional:(NSString *)propertyName {

  NSArray *optionalProperties = @[@"voicemailGreetingAlternate", @"nameGreeting", @"name", @"timezone", @"includeInDirectory", @"extension", @"enableOutboundCalls", @"usageType", @"voicemailPassword", @"fullName", @"enableCallWaiting", @"voicemailGreetingStandard", @"voicemailGreetingType", @"callerId", @"localAreaCode", @"voicemailEnabled", @"voicemailGreetingEnableLeaveMessagePrompt", @"voicemailTranscription", @"voicemailNotificationsEmails", @"voicemailNotificationsSms", @"callNotificationsEmails", @"callNotificationsSms", @"route"];
  return [optionalProperties containsObject:propertyName];
}

@end
