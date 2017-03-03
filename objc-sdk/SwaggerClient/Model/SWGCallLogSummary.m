#import "SWGCallLogSummary.h"

@implementation SWGCallLogSummary

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
  return [[JSONKeyMapper alloc] initWithDictionary:@{ @"id": @"_id", @"uuid": @"uuid", @"extension": @"extension", @"caller_id": @"callerId", @"called_number": @"calledNumber", @"start_time": @"startTime", @"created_at": @"createdAt", @"direction": @"direction", @"type": @"type", @"call_duration": @"callDuration", @"is_monitored": @"isMonitored", @"call_number": @"callNumber", @"final_action": @"finalAction", @"call_recording": @"callRecording" }];
}

/**
 * Indicates whether the property with the given name is optional.
 * If `propertyName` is optional, then return `YES`, otherwise return `NO`.
 * This method is used by `JSONModel`.
 */
+ (BOOL)propertyIsOptional:(NSString *)propertyName {

  NSArray *optionalProperties = @[@"_id", @"uuid", @"extension", @"callerId", @"calledNumber", @"startTime", @"createdAt", @"direction", @"type", @"callDuration", @"isMonitored", @"callNumber", @"finalAction", @"callRecording"];
  return [optionalProperties containsObject:propertyName];
}

@end
