#import "textview.h"
#include "_cgo_export.h"

@implementation TextViewHandler
@end

TextViewPtr TextView_New(int goTextViewId, int x, int y, int w, int h) {
	/* create the NSTextView and add it to the window */
	id nsTextView = [[[NSTextView alloc] initWithFrame:NSMakeRect(x, y, w, h)] autorelease];

	return (TextViewPtr)nsTextView;
}

void TextView_SetText(TextViewPtr textViewPtr, const char* text) {
	NSTextView* tv = (NSTextView*)textViewPtr;
	[tv setString:[NSString stringWithUTF8String:text]];
}

void TextView_SetBackgroundColor(TextViewPtr textViewPtr, int r, int g, int b, int a) {
     NSTextView* tv = (NSTextView*)textViewPtr;
     //[button setBordered:false];
     [tv setBackgroundColor:[NSColor colorWithCalibratedRed:r/255.f green:g/255.f blue:b/255.f alpha:a/255.f]];
}

void TextView_Remove(TextViewPtr textViewPtr) {
    NSTextView* tv = (NSTextView*)textViewPtr;
    [tv removeFromSuperview];
}
