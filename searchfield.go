package gocoa

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #import "searchfield.h"
import "C"

// SearchField represents a SearchField control that can trigger actions.
type SearchField struct {
	searchFieldPtr C.SearchFieldPtr
	callback       func()
}

var searchfields []*SearchField

func NewSearchField(x int, y int, width int, height int) *SearchField {
	searchFieldID := len(searchfields)
	searchFieldPtr := C.SearchField_New(C.int(searchFieldID), C.int(x), C.int(y), C.int(width), C.int(height))

	sf := &SearchField{
		searchFieldPtr: searchFieldPtr,
	}

	searchfields = append(searchfields, sf)
	return sf
}
