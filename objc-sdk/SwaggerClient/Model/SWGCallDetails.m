#import "SWGCallDetails.h"

@implementation SWGCallDetails

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
  return [[JSONKeyMapper alloc] initWithDictionary:@{ @"start_time": @"startTime", @"type": @"type", @"id_value": @"idValue", @"voip_id": @"voipId", @"voip_phone_id": @"voipPhoneId" }];
}

/**
 * Indicates whether the property with the given name is optional.
 * If `propertyName` is optional, then return `YES`, otherwise return `NO`.
 * This method is used by `JSONModel`.
 */
+ (BOOL)propertyIsOptional:(NSString *)propertyName {

  NSArray *optionalProperties = @[@"startTime", @"type", @"idValue", @"voipId", @"voipPhoneId"];
  return [optionalProperties containsObject:propertyName];
}

@end
