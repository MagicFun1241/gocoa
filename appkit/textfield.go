package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #import "textfield.h"
import "C"
import (
	"unsafe"

	"github.com/magicfun1241/gocoa/tools"
)

// TextField -Button represents a button control that can trigger actions.
type TextField struct {
	textFieldPtr C.TextFieldPtr
	callback     func()
}

var textfields []*TextField

// NewTextField - This func is not thread safe.
func NewTextField(x int, y int, width int, height int) *TextField {
	textFieldID := len(textfields)
	textFieldPtr := C.TextField_New(C.int(textFieldID), C.int(x), C.int(y), C.int(width), C.int(height))

	tf := &TextField{
		textFieldPtr: textFieldPtr,
	}
	textfields = append(textfields, tf)
	return tf
}

// StringValue - returns the string value of the text field
func (textField *TextField) StringValue() string {
	return C.GoString(C.TextField_StringValue(textField.textFieldPtr))
}

// SetStringValue sets the string value of the text field
func (textField *TextField) SetStringValue(text string) {
	cText := C.CString(text)
	defer C.free(unsafe.Pointer(cText))

	C.TextField_SetStringValue(textField.textFieldPtr, cText)
}

func (textField *TextField) Enabled() bool {
	return C.TextField_Enabled(textField.textFieldPtr) == 1
}

// SetEnabled sets the enabled value of the text field. CANNOT BE CHANGED AT RUNTIME
func (textField *TextField) SetEnabled(enabled bool) {
	if enabled {
		C.TextField_SetEnabled(textField.textFieldPtr, 1)
	} else {
		C.TextField_SetEnabled(textField.textFieldPtr, 0)
	}
}

func (textField *TextField) Editable() bool {
	return C.TextField_Editable(textField.textFieldPtr) == 1
}

func (textField *TextField) SetEditable(editable bool) {
	if editable {
		C.TextField_SetEditable(textField.textFieldPtr, 1)
	} else {
		C.TextField_SetEditable(textField.textFieldPtr, 0)
	}
}

func (textField *TextField) SetFontFamily(fontFamily string) {
	cText := C.CString(fontFamily)
	defer C.free(unsafe.Pointer(cText))

	C.TextField_SetFontFamily(textField.textFieldPtr, cText)
}

func (textField *TextField) SetFontSize(fontSize int) {
	C.TextField_SetFontSize(textField.textFieldPtr, C.int(fontSize))
}

func (textField *TextField) SetColor(hexRGBA string) {
	var r, g, b, a = tools.ParseHexRGBA(hexRGBA)
	C.TextField_SetColor(textField.textFieldPtr, C.int(r), C.int(g), C.int(b), C.int(a))
}

func (textField *TextField) SetBackgroundColor(hexRGBA string) {
	var r, g, b, a = tools.ParseHexRGBA(hexRGBA)
	C.TextField_SetBackgroundColor(textField.textFieldPtr, C.int(r), C.int(g), C.int(b), C.int(a))
}

func (textField *TextField) SetBorderColor(hexRGBA string) {
	var r, g, b, a = tools.ParseHexRGBA(hexRGBA)
	C.TextField_SetBorderColor(textField.textFieldPtr, C.int(r), C.int(g), C.int(b), C.int(a))
}

func (textField *TextField) SetBorderWidth(borderWidth int) {
	C.TextField_SetBorderWidth(textField.textFieldPtr, C.int(borderWidth))
}

func (textField *TextField) SetBezeled(bezeled bool) {
	if bezeled {
		C.TextField_SetBezeled(textField.textFieldPtr, C.int(1))
	} else {
		C.TextField_SetBezeled(textField.textFieldPtr, C.int(0))
	}
}

func (textField *TextField) SetDrawsBackground(drawsBackground bool) {
	if drawsBackground {
		C.TextField_SetDrawsBackground(textField.textFieldPtr, C.int(1))
	} else {
		C.TextField_SetDrawsBackground(textField.textFieldPtr, C.int(0))
	}
}

func (textField *TextField) SetSelectable(selectable bool) {
	if selectable {
		C.TextField_SetSelectable(textField.textFieldPtr, C.int(1))
	} else {
		C.TextField_SetSelectable(textField.textFieldPtr, C.int(0))
	}
}

// Text alignment constants
const (
	TextAlignmentLeft   = 0 // NSTextAlignmentLeft
	TextAlignmentCenter = 1 // NSTextAlignmentCenter
	TextAlignmentRight  = 2 // NSTextAlignmentRight
)

// Line break mode constants
const (
	LineBreakByWordWrapping     = 0 // NSLineBreakByWordWrapping
	LineBreakByCharWrapping     = 1 // NSLineBreakByCharWrapping
	LineBreakByClipping         = 2 // NSLineBreakByClipping
	LineBreakByTruncatingHead   = 3 // NSLineBreakByTruncatingHead
	LineBreakByTruncatingTail   = 4 // NSLineBreakByTruncatingTail
	LineBreakByTruncatingMiddle = 5 // NSLineBreakByTruncatingMiddle
)

func (textField *TextField) SetAlignment(alignment int) {
	C.TextField_SetAlignment(textField.textFieldPtr, C.int(alignment))
}

// PlaceholderString returns the placeholder string of the text field
func (textField *TextField) PlaceholderString() string {
	return C.GoString(C.TextField_PlaceholderString(textField.textFieldPtr))
}

// SetPlaceholderString sets the placeholder string of the text field
func (textField *TextField) SetPlaceholderString(placeholder string) {
	cText := C.CString(placeholder)
	defer C.free(unsafe.Pointer(cText))
	C.TextField_SetPlaceholderString(textField.textFieldPtr, cText)
}

// MaximumNumberOfLines returns the maximum number of lines
func (textField *TextField) MaximumNumberOfLines() int {
	return int(C.TextField_MaximumNumberOfLines(textField.textFieldPtr))
}

// SetMaximumNumberOfLines sets the maximum number of lines (0 = unlimited)
func (textField *TextField) SetMaximumNumberOfLines(maxLines int) {
	C.TextField_SetMaximumNumberOfLines(textField.textFieldPtr, C.int(maxLines))
}

// UsesSingleLineMode returns whether the text field uses single line mode
func (textField *TextField) UsesSingleLineMode() bool {
	return C.TextField_UsesSingleLineMode(textField.textFieldPtr) == 1
}

// SetUsesSingleLineMode sets whether the text field uses single line mode
func (textField *TextField) SetUsesSingleLineMode(singleLine bool) {
	if singleLine {
		C.TextField_SetUsesSingleLineMode(textField.textFieldPtr, C.int(1))
	} else {
		C.TextField_SetUsesSingleLineMode(textField.textFieldPtr, C.int(0))
	}
}

// LineBreakMode returns the line break mode
func (textField *TextField) LineBreakMode() int {
	return int(C.TextField_LineBreakMode(textField.textFieldPtr))
}

// SetLineBreakMode sets the line break mode
func (textField *TextField) SetLineBreakMode(lineBreakMode int) {
	C.TextField_SetLineBreakMode(textField.textFieldPtr, C.int(lineBreakMode))
}

// SetTextColor sets the text color (alternative to SetColor with same functionality)
func (textField *TextField) SetTextColor(hexRGBA string) {
	var r, g, b, a = tools.ParseHexRGBA(hexRGBA)
	C.TextField_SetTextColor(textField.textFieldPtr, C.int(r), C.int(g), C.int(b), C.int(a))
}

// SetFont sets the font with name and size
func (textField *TextField) SetFont(fontName string, fontSize float64) {
	cFontName := C.CString(fontName)
	defer C.free(unsafe.Pointer(cFontName))
	C.TextField_SetFont(textField.textFieldPtr, cFontName, C.double(fontSize))
}

// IsHighlighted returns whether the text field is highlighted
func (textField *TextField) IsHighlighted() bool {
	return C.TextField_IsHighlighted(textField.textFieldPtr) == 1
}

// SetHighlighted sets the highlighted state of the text field
func (textField *TextField) SetHighlighted(highlighted bool) {
	if highlighted {
		C.TextField_SetHighlighted(textField.textFieldPtr, C.int(1))
	} else {
		C.TextField_SetHighlighted(textField.textFieldPtr, C.int(0))
	}
}

func (textField *TextField) Remove() {
	C.TextField_Remove(textField.textFieldPtr)
}
