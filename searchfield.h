#import <Cocoa/Cocoa.h>

@interface SearchFieldHandler : NSObject

@property (assign) int goSearchFieldId;

@end

typedef void* SearchFieldPtr;

SearchFieldPtr SearchField_New(int goSearchFieldId, int x, int y, int w, int h);