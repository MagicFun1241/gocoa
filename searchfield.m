#import "searchfield.h"
#include "_cgo_export.h"

@implementation SearchFieldHandler
@end

TextViewPtr SearchField_New(int goSearchFieldId, int x, int y, int w, int h) {
	/* create the NSSearchField and add it to the window */
	id nsSearchField = [[[NSSearchField alloc] initWithFrame:NSMakeRect(x, y, w, h)] autorelease];

	return (TextViewPtr)nsSearchField;
}