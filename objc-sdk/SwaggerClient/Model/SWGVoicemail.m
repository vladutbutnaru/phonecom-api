#import "SWGVoicemail.h"

@implementation SWGVoicemail

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
  return [[JSONKeyMapper alloc] initWithDictionary:@{ @"enabled": @"enabled", @"password": @"password", @"greeting": @"greeting", @"attachments": @"attachments", @"notifications": @"notifications", @"transcription": @"transcription" }];
}

/**
 * Indicates whether the property with the given name is optional.
 * If `propertyName` is optional, then return `YES`, otherwise return `NO`.
 * This method is used by `JSONModel`.
 */
+ (BOOL)propertyIsOptional:(NSString *)propertyName {

  NSArray *optionalProperties = @[@"enabled", @"password", @"greeting", @"attachments", @"notifications", @"transcription"];
  return [optionalProperties containsObject:propertyName];
}

@end
