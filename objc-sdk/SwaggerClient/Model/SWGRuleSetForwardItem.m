#import "SWGRuleSetForwardItem.h"

@implementation SWGRuleSetForwardItem

- (instancetype)init {
  self = [super init];
  if (self) {
    // initialize property's default value, if any
    self.screening = @0;
    self.callerId = @"calling_number";
    
  }
  return self;
}


/**
 * Maps json key to property name.
 * This method is used by `JSONModel`.
 */
+ (JSONKeyMapper *)keyMapper {
  return [[JSONKeyMapper alloc] initWithDictionary:@{ @"type": @"type", @"extension": @"extension", @"number": @"number", @"screening": @"screening", @"caller_id": @"callerId", @"voice_tag": @"voiceTag", @"distinctive_ring": @"distinctiveRing" }];
}

/**
 * Indicates whether the property with the given name is optional.
 * If `propertyName` is optional, then return `YES`, otherwise return `NO`.
 * This method is used by `JSONModel`.
 */
+ (BOOL)propertyIsOptional:(NSString *)propertyName {

  NSArray *optionalProperties = @[@"type", @"extension", @"number", @"screening", @"callerId", @"voiceTag", @"distinctiveRing"];
  return [optionalProperties containsObject:propertyName];
}

@end
