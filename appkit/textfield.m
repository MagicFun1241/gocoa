#import "textfield.h"
#include "_cgo_export.h"

@implementation TextFieldHandler
@end

TextFieldPtr TextField_New(int goTextFieldId, int x, int y, int w, int h) {
	/* create the NSTextField and add it to the window */
	id nsTextField = [[[NSTextField alloc] initWithFrame:NSMakeRect(x, y, w, h)] autorelease];

	return (TextFieldPtr)nsTextField;
}

const char* TextField_StringValue(TextFieldPtr textFieldPtr) {
	NSTextField* textField = (NSTextField*)textFieldPtr;
	return [[textField stringValue] cStringUsingEncoding:NSISOLatin1StringEncoding];
}

void TextField_SetStringValue(TextFieldPtr textFieldPtr, const char* text) {
	NSTextField* textField = (NSTextField*)textFieldPtr;
	[textField setStringValue:[NSString stringWithUTF8String:text]];
}

const int TextField_Enabled(TextFieldPtr textFieldPtr) {
	NSTextField* textField = (NSTextField*)textFieldPtr;
	return textField.enabled ? 1 : 0;
}

void TextField_SetEnabled(TextFieldPtr textFieldPtr, int enabled) {
	NSTextField* textField = (NSTextField*)textFieldPtr;
	[textField becomeFirstResponder];
	textField.enabled = enabled ? YES : NO;
	[textField resignFirstResponder];
}

const int TextField_Editable(TextFieldPtr textFieldPtr) {
	NSTextField* textField = (NSTextField*)textFieldPtr;
	return textField.editable;
}

void TextField_SetEditable(TextFieldPtr textFieldPtr, int editable) {
	NSTextField* textField = (NSTextField*)textFieldPtr;
	[textField becomeFirstResponder];
	textField.editable = editable ? YES : NO;
	[textField resignFirstResponder];
}

void TextField_SetFontFamily(TextFieldPtr textFieldPtr, const char* fontFamily) {
	NSTextField* textField = (NSTextField*)textFieldPtr;
	textField.font = [NSFont fontWithName:[NSString stringWithUTF8String:fontFamily] size:textField.font.pointSize];
}

void TextField_SetFontSize(TextFieldPtr textFieldPtr, const int fontSize) {
	NSTextField* textField = (NSTextField*)textFieldPtr;
	textField.font = [NSFont fontWithName:textField.font.fontName size:fontSize];
}

void TextField_SetColor(TextFieldPtr textFieldPtr, const int r, const int g, const int b, const int a) {
	NSTextField* textField = (NSTextField*)textFieldPtr;
	textField.textColor = [NSColor colorWithCalibratedRed:r/255.f green:g/255.f blue:b/255.f alpha:a/255.f];
}

void TextField_SetBackgroundColor(TextFieldPtr textFieldPtr, const int r, const int g, const int b, const int a) {
	NSTextField* textField = (NSTextField*)textFieldPtr;
	textField.drawsBackground = true;
	textField.backgroundColor = [NSColor colorWithCalibratedRed:r/255.f green:g/255.f blue:b/255.f alpha:a/255.f];
}

void TextField_SetBorderColor(TextFieldPtr textFieldPtr, const int r, const int g, const int b, const int a) {
	NSTextField* textField = (NSTextField*)textFieldPtr;
	textField.wantsLayer = true;
	textField.layer.borderColor = [[NSColor colorWithCalibratedRed:r/255.f green:g/255.f blue:b/255.f alpha:a/255.f] CGColor];
}

void TextField_SetBorderWidth(TextFieldPtr textFieldPtr, const int borderWidth) {
	NSTextField* textField = (NSTextField*)textFieldPtr;
	textField.wantsLayer = true;
	textField.layer.borderWidth = borderWidth;
}

void TextField_SetBezeled(TextFieldPtr textFieldPtr, const int bezeled) {
	NSTextField* textField = (NSTextField*)textFieldPtr;
	[textField setBezeled:bezeled];
}

void TextField_SetDrawsBackground(TextFieldPtr textFieldPtr, const int drawsBackground) {
	NSTextField* textField = (NSTextField*)textFieldPtr;
	[textField setDrawsBackground:drawsBackground];
}

void TextField_SetSelectable(TextFieldPtr textFieldPtr, const int selectable) {
	NSTextField* textField = (NSTextField*)textFieldPtr;
	[textField setSelectable:selectable];
}

void TextField_SetAlignment(TextFieldPtr textFieldPtr, const int alignment) {
	NSTextField* textField = (NSTextField*)textFieldPtr;
	[textField setAlignment:alignment];
}

const char* TextField_PlaceholderString(TextFieldPtr textFieldPtr) {
	NSTextField* textField = (NSTextField*)textFieldPtr;
	NSString* placeholder = [textField placeholderString];
	return placeholder ? [placeholder cStringUsingEncoding:NSISOLatin1StringEncoding] : "";
}

void TextField_SetPlaceholderString(TextFieldPtr textFieldPtr, const char* placeholder) {
	NSTextField* textField = (NSTextField*)textFieldPtr;
	[textField setPlaceholderString:[NSString stringWithUTF8String:placeholder]];
}

const int TextField_MaximumNumberOfLines(TextFieldPtr textFieldPtr) {
	NSTextField* textField = (NSTextField*)textFieldPtr;
	return (int)[textField maximumNumberOfLines];
}

void TextField_SetMaximumNumberOfLines(TextFieldPtr textFieldPtr, const int maxLines) {
	NSTextField* textField = (NSTextField*)textFieldPtr;
	[textField setMaximumNumberOfLines:maxLines];
}

const int TextField_UsesSingleLineMode(TextFieldPtr textFieldPtr) {
	NSTextField* textField = (NSTextField*)textFieldPtr;
	return [textField usesSingleLineMode] ? 1 : 0;
}

void TextField_SetUsesSingleLineMode(TextFieldPtr textFieldPtr, const int singleLine) {
	NSTextField* textField = (NSTextField*)textFieldPtr;
	[textField setUsesSingleLineMode:singleLine ? YES : NO];
}

const int TextField_LineBreakMode(TextFieldPtr textFieldPtr) {
	NSTextField* textField = (NSTextField*)textFieldPtr;
	return (int)[[textField cell] lineBreakMode];
}

void TextField_SetLineBreakMode(TextFieldPtr textFieldPtr, const int lineBreakMode) {
	NSTextField* textField = (NSTextField*)textFieldPtr;
	[[textField cell] setLineBreakMode:lineBreakMode];
}

void TextField_SetTextColor(TextFieldPtr textFieldPtr, const int r, const int g, const int b, const int a) {
	NSTextField* textField = (NSTextField*)textFieldPtr;
	[textField setTextColor:[NSColor colorWithCalibratedRed:r/255.0f green:g/255.0f blue:b/255.0f alpha:a/255.0f]];
}

void TextField_SetFont(TextFieldPtr textFieldPtr, const char* fontName, const double fontSize) {
	NSTextField* textField = (NSTextField*)textFieldPtr;
	NSString* fontNameStr = [NSString stringWithUTF8String:fontName];
	NSFont* font = [NSFont fontWithName:fontNameStr size:fontSize];
	if (font) {
		[textField setFont:font];
	}
}

const int TextField_IsHighlighted(TextFieldPtr textFieldPtr) {
	NSTextField* textField = (NSTextField*)textFieldPtr;
	return [textField isHighlighted] ? 1 : 0;
}

void TextField_SetHighlighted(TextFieldPtr textFieldPtr, const int highlighted) {
	NSTextField* textField = (NSTextField*)textFieldPtr;
	[textField setHighlighted:highlighted ? YES : NO];
}

void TextField_Remove(TextFieldPtr textFieldPtr) {
    NSTextField* tf = (NSTextField*)textFieldPtr;
    [tf removeFromSuperview];
}