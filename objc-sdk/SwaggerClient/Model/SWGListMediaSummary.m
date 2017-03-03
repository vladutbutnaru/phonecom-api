#import "SWGListMediaSummary.h"

@implementation SWGListMediaSummary

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
  return [[JSONKeyMapper alloc] initWithDictionary:@{ @"filters": @"filters", @"sort": @"sort", @"total": @"total", @"offset": @"offset", @"limit": @"limit", @"items": @"items" }];
}

/**
 * Indicates whether the property with the given name is optional.
 * If `propertyName` is optional, then return `YES`, otherwise return `NO`.
 * This method is used by `JSONModel`.
 */
+ (BOOL)propertyIsOptional:(NSString *)propertyName {

  NSArray *optionalProperties = @[@"filters", @"sort", @"total", @"offset", @"limit", @"items"];
  return [optionalProperties containsObject:propertyName];
}

@end
