#import <Cocoa/Cocoa.h>

// typedef void (*callback)(void);

@interface TextFieldHandler : NSObject

@property (assign) int goTextFieldId;

@end

typedef void* TextFieldPtr;

TextFieldPtr TextField_New(int goTextFieldId, int x, int y, int w, int h);

const char* TextField_StringValue(TextFieldPtr textFieldPtr);
void TextField_SetStringValue(TextFieldPtr textFieldPtr, const char* text);
const int TextField_Enabled(TextFieldPtr textFieldPtr);
void TextField_SetEnabled(TextFieldPtr textFieldPtr, int enabled);
const int TextField_Editable(TextFieldPtr textFieldPtr);
void TextField_SetEditable(TextFieldPtr textFieldPtr, int editable);
void TextField_SetFontFamily(TextFieldPtr textFieldPtr, const char* fontFamily);
void TextField_SetFontSize(TextFieldPtr textFieldPtr, const int fontSize);
void TextField_SetColor(TextFieldPtr textFieldPtr, const int r, const int g, const int b, const int a);
void TextField_SetBackgroundColor(TextFieldPtr textFieldPtr, const int r, const int g, const int b, const int a);
void TextField_SetBorderColor(TextFieldPtr textFieldPtr, const int r, const int g, const int b, const int a);
void TextField_SetBorderWidth(TextFieldPtr textFieldPtr, const int borderWidth);
void TextField_SetBezeled(TextFieldPtr textFieldPtr, const int bezeled);
void TextField_SetDrawsBackground(TextFieldPtr textFieldPtr, const int drawsBackground);
void TextField_SetSelectable(TextFieldPtr textFieldPtr, const int selectable);
void TextField_SetAlignment(TextFieldPtr textFieldPtr, const int alignment);
const char* TextField_PlaceholderString(TextFieldPtr textFieldPtr);
void TextField_SetPlaceholderString(TextFieldPtr textFieldPtr, const char* placeholder);
const int TextField_MaximumNumberOfLines(TextFieldPtr textFieldPtr);
void TextField_SetMaximumNumberOfLines(TextFieldPtr textFieldPtr, const int maxLines);
const int TextField_UsesSingleLineMode(TextFieldPtr textFieldPtr);
void TextField_SetUsesSingleLineMode(TextFieldPtr textFieldPtr, const int singleLine);
const int TextField_LineBreakMode(TextFieldPtr textFieldPtr);
void TextField_SetLineBreakMode(TextFieldPtr textFieldPtr, const int lineBreakMode);
void TextField_SetTextColor(TextFieldPtr textFieldPtr, const int r, const int g, const int b, const int a);
void TextField_SetFont(TextFieldPtr textFieldPtr, const char* fontName, const double fontSize);
const int TextField_IsHighlighted(TextFieldPtr textFieldPtr);
void TextField_SetHighlighted(TextFieldPtr textFieldPtr, const int highlighted);
void TextField_Remove(TextFieldPtr textFieldPtr);