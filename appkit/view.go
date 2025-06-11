package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa -framework QuartzCore
// #include "view.h"
import "C"
import "unsafe"

// AutoresizingMask constants for view autoresizing behavior
type AutoresizingMask int

const (
	ViewNotSizable         AutoresizingMask = 0
	ViewMinXMarginFlexible AutoresizingMask = 1
	ViewWidthFlexible      AutoresizingMask = 2
	ViewMaxXMarginFlexible AutoresizingMask = 4
	ViewMinYMarginFlexible AutoresizingMask = 8
	ViewHeightFlexible     AutoresizingMask = 16
	ViewMaxYMarginFlexible AutoresizingMask = 32
)

// FocusRingType constants
type FocusRingType int

const (
	FocusRingTypeDefault  FocusRingType = 0
	FocusRingTypeNone     FocusRingType = 1
	FocusRingTypeExterior FocusRingType = 2
)

// AppearanceType constants
type AppearanceType int

const (
	AppearanceLight AppearanceType = 0
	AppearanceDark  AppearanceType = 1
)

// Rect represents a rectangle with origin and size
type Rect struct {
	X      int
	Y      int
	Width  int
	Height int
}

// Point represents a point with x and y coordinates
type Point struct {
	X int
	Y int
}

// Size represents a size with width and height
type Size struct {
	Width  int
	Height int
}

// View represents an NSView
type View struct {
	viewPtr unsafe.Pointer
}

// NewView creates a new NSView with the specified frame
func NewView(x, y, width, height int) *View {
	viewPtr := C.View_New(C.int(x), C.int(y), C.int(width), C.int(height))
	return &View{viewPtr: viewPtr}
}

// NewViewWithFrame creates a new NSView with the specified frame (alias for NewView)
func NewViewWithFrame(x, y, width, height int) *View {
	return NewView(x, y, width, height)
}

// Dealloc releases the view
func (v *View) Dealloc() {
	if v.viewPtr != nil {
		C.View_Dealloc(v.viewPtr)
		v.viewPtr = nil
	}
}

// View hierarchy management

// AddSubview adds a subview to this view
func (v *View) AddSubview(subview *View) {
	C.View_AddSubview(v.viewPtr, subview.viewPtr)
}

// RemoveFromSuperview removes this view from its superview
func (v *View) RemoveFromSuperview() {
	C.View_RemoveFromSuperview(v.viewPtr)
}

// RemoveSubview removes the specified subview from this view
func (v *View) RemoveSubview(subview *View) {
	C.View_RemoveSubview(v.viewPtr, subview.viewPtr)
}

// GetSuperview returns the superview of this view
func (v *View) GetSuperview() *View {
	superviewPtr := C.View_GetSuperview(v.viewPtr)
	if superviewPtr == nil {
		return nil
	}
	return &View{viewPtr: superviewPtr}
}

// GetSubviews returns all subviews of this view
func (v *View) GetSubviews() []*View {
	var count C.int
	subviewsPtrPtr := C.View_GetSubviews(v.viewPtr, &count)

	if count == 0 || subviewsPtrPtr == nil {
		return []*View{}
	}

	// Convert C array to Go slice
	subviewsSlice := (*[1 << 30]unsafe.Pointer)(unsafe.Pointer(subviewsPtrPtr))[:count:count]
	result := make([]*View, count)

	for i := 0; i < int(count); i++ {
		result[i] = &View{viewPtr: subviewsSlice[i]}
	}

	// Free the C array
	C.free(unsafe.Pointer(subviewsPtrPtr))
	return result
}

// IsDescendantOf returns true if this view is a descendant of the specified ancestor view
func (v *View) IsDescendantOf(ancestor *View) bool {
	return C.View_IsDescendantOf(v.viewPtr, ancestor.viewPtr) != 0
}

// Frame and bounds management

// SetFrame sets the view's frame rectangle
func (v *View) SetFrame(x, y, width, height int) {
	C.View_SetFrame(v.viewPtr, C.int(x), C.int(y), C.int(width), C.int(height))
}

// GetFrame returns the view's frame rectangle
func (v *View) GetFrame() Rect {
	var x, y, width, height C.int
	C.View_GetFrame(v.viewPtr, &x, &y, &width, &height)
	return Rect{X: int(x), Y: int(y), Width: int(width), Height: int(height)}
}

// SetBounds sets the view's bounds rectangle
func (v *View) SetBounds(x, y, width, height int) {
	C.View_SetBounds(v.viewPtr, C.int(x), C.int(y), C.int(width), C.int(height))
}

// GetBounds returns the view's bounds rectangle
func (v *View) GetBounds() Rect {
	var x, y, width, height C.int
	C.View_GetBounds(v.viewPtr, &x, &y, &width, &height)
	return Rect{X: int(x), Y: int(y), Width: int(width), Height: int(height)}
}

// SetFrameOrigin sets the view's frame origin
func (v *View) SetFrameOrigin(x, y int) {
	C.View_SetFrameOrigin(v.viewPtr, C.int(x), C.int(y))
}

// SetFrameSize sets the view's frame size
func (v *View) SetFrameSize(width, height int) {
	C.View_SetFrameSize(v.viewPtr, C.int(width), C.int(height))
}

// SetBoundsOrigin sets the view's bounds origin
func (v *View) SetBoundsOrigin(x, y int) {
	C.View_SetBoundsOrigin(v.viewPtr, C.int(x), C.int(y))
}

// SetBoundsSize sets the view's bounds size
func (v *View) SetBoundsSize(width, height int) {
	C.View_SetBoundsSize(v.viewPtr, C.int(width), C.int(height))
}

// Display and drawing

// SetNeedsDisplay marks the view as needing display
func (v *View) SetNeedsDisplay(needsDisplay bool) {
	var needs C.int
	if needsDisplay {
		needs = 1
	}
	C.View_SetNeedsDisplay(v.viewPtr, needs)
}

// SetNeedsDisplayInRect marks a specific rect as needing display
func (v *View) SetNeedsDisplayInRect(x, y, width, height int) {
	C.View_SetNeedsDisplayInRect(v.viewPtr, C.int(x), C.int(y), C.int(width), C.int(height))
}

// Display forces the view to redraw immediately
func (v *View) Display() {
	C.View_Display(v.viewPtr)
}

// DisplayIfNeeded redraws the view if it needs display
func (v *View) DisplayIfNeeded() {
	C.View_DisplayIfNeeded(v.viewPtr)
}

// DisplayRect redraws a specific rectangle
func (v *View) DisplayRect(x, y, width, height int) {
	C.View_DisplayRect(v.viewPtr, C.int(x), C.int(y), C.int(width), C.int(height))
}

// NeedsDisplay returns true if the view needs display
func (v *View) NeedsDisplay() bool {
	return C.View_NeedsDisplay(v.viewPtr) != 0
}

// Coordinate conversion

// ConvertPointToView converts a point from this view's coordinate system to another view's
func (v *View) ConvertPointToView(x, y int, toView *View) Point {
	var outX, outY C.int
	var toViewPtr unsafe.Pointer
	if toView != nil {
		toViewPtr = toView.viewPtr
	}
	C.View_ConvertPointToView(v.viewPtr, C.int(x), C.int(y), toViewPtr, &outX, &outY)
	return Point{X: int(outX), Y: int(outY)}
}

// ConvertPointFromView converts a point from another view's coordinate system to this view's
func (v *View) ConvertPointFromView(x, y int, fromView *View) Point {
	var outX, outY C.int
	var fromViewPtr unsafe.Pointer
	if fromView != nil {
		fromViewPtr = fromView.viewPtr
	}
	C.View_ConvertPointFromView(v.viewPtr, C.int(x), C.int(y), fromViewPtr, &outX, &outY)
	return Point{X: int(outX), Y: int(outY)}
}

// ConvertRectToView converts a rectangle from this view's coordinate system to another view's
func (v *View) ConvertRectToView(x, y, width, height int, toView *View) Rect {
	var outX, outY, outWidth, outHeight C.int
	var toViewPtr unsafe.Pointer
	if toView != nil {
		toViewPtr = toView.viewPtr
	}
	C.View_ConvertRectToView(v.viewPtr, C.int(x), C.int(y), C.int(width), C.int(height), toViewPtr, &outX, &outY, &outWidth, &outHeight)
	return Rect{X: int(outX), Y: int(outY), Width: int(outWidth), Height: int(outHeight)}
}

// ConvertRectFromView converts a rectangle from another view's coordinate system to this view's
func (v *View) ConvertRectFromView(x, y, width, height int, fromView *View) Rect {
	var outX, outY, outWidth, outHeight C.int
	var fromViewPtr unsafe.Pointer
	if fromView != nil {
		fromViewPtr = fromView.viewPtr
	}
	C.View_ConvertRectFromView(v.viewPtr, C.int(x), C.int(y), C.int(width), C.int(height), fromViewPtr, &outX, &outY, &outWidth, &outHeight)
	return Rect{X: int(outX), Y: int(outY), Width: int(outWidth), Height: int(outHeight)}
}

// Hit testing

// HitTest returns the deepest descendant view that contains the specified point
func (v *View) HitTest(x, y int) *View {
	viewPtr := C.View_HitTest(v.viewPtr, C.int(x), C.int(y))
	if viewPtr == nil {
		return nil
	}
	return &View{viewPtr: viewPtr}
}

// IsOpaque returns true if the view is opaque
func (v *View) IsOpaque() bool {
	return C.View_IsOpaque(v.viewPtr) != 0
}

// SetOpaque sets whether the view is opaque (note: limited implementation)
func (v *View) SetOpaque(opaque bool) {
	var opaqueVal C.int
	if opaque {
		opaqueVal = 1
	}
	C.View_SetOpaque(v.viewPtr, opaqueVal)
}

// Visibility and clipping

// IsHidden returns true if the view is hidden
func (v *View) IsHidden() bool {
	return C.View_IsHidden(v.viewPtr) != 0
}

// SetHidden sets whether the view is hidden
func (v *View) SetHidden(hidden bool) {
	var hiddenVal C.int
	if hidden {
		hiddenVal = 1
	}
	C.View_SetHidden(v.viewPtr, hiddenVal)
}

// IsHiddenOrHasHiddenAncestor returns true if the view or any ancestor is hidden
func (v *View) IsHiddenOrHasHiddenAncestor() bool {
	return C.View_IsHiddenOrHasHiddenAncestor(v.viewPtr) != 0
}

// SetClipsToBounds sets whether the view clips its content to its bounds
func (v *View) SetClipsToBounds(clips bool) {
	var clipsVal C.int
	if clips {
		clipsVal = 1
	}
	C.View_SetClipsToBounds(v.viewPtr, clipsVal)
}

// ClipsToBounds returns true if the view clips its content to its bounds
func (v *View) ClipsToBounds() bool {
	return C.View_ClipsToBounds(v.viewPtr) != 0
}

// Auto layout and constraints

// SetTranslatesAutoresizingMaskIntoConstraints sets whether the view's autoresizing mask is translated into constraints
func (v *View) SetTranslatesAutoresizingMaskIntoConstraints(translates bool) {
	var translatesVal C.int
	if translates {
		translatesVal = 1
	}
	C.View_SetTranslatesAutoresizingMaskIntoConstraints(v.viewPtr, translatesVal)
}

// TranslatesAutoresizingMaskIntoConstraints returns true if the view's autoresizing mask is translated into constraints
func (v *View) TranslatesAutoresizingMaskIntoConstraints() bool {
	return C.View_TranslatesAutoresizingMaskIntoConstraints(v.viewPtr) != 0
}

// SetAutoresizingMask sets the view's autoresizing mask
func (v *View) SetAutoresizingMask(mask AutoresizingMask) {
	C.View_SetAutoresizingMask(v.viewPtr, C.int(mask))
}

// GetAutoresizingMask returns the view's autoresizing mask
func (v *View) GetAutoresizingMask() AutoresizingMask {
	return AutoresizingMask(C.View_GetAutoresizingMask(v.viewPtr))
}

// Layer and appearance

// SetWantsLayer sets whether the view wants a backing layer
func (v *View) SetWantsLayer(wantsLayer bool) {
	var wants C.int
	if wantsLayer {
		wants = 1
	}
	C.View_SetWantsLayer(v.viewPtr, wants)
}

// WantsLayer returns true if the view wants a backing layer
func (v *View) WantsLayer() bool {
	return C.View_WantsLayer(v.viewPtr) != 0
}

// GetLayer returns the view's backing layer
func (v *View) GetLayer() unsafe.Pointer {
	return C.View_GetLayer(v.viewPtr)
}

// SetLayer sets the view's backing layer
func (v *View) SetLayer(layer unsafe.Pointer) {
	C.View_SetLayer(v.viewPtr, layer)
}

// SetAlphaValue sets the view's alpha transparency
func (v *View) SetAlphaValue(alpha float64) {
	C.View_SetAlphaValue(v.viewPtr, C.double(alpha))
}

// GetAlphaValue returns the view's alpha transparency
func (v *View) GetAlphaValue() float64 {
	return float64(C.View_GetAlphaValue(v.viewPtr))
}

// Responder chain and events

// AcceptsFirstResponder returns true if the view accepts first responder status
func (v *View) AcceptsFirstResponder() bool {
	return C.View_AcceptsFirstResponder(v.viewPtr) != 0
}

// BecomeFirstResponder attempts to make the view the first responder
func (v *View) BecomeFirstResponder() bool {
	return C.View_BecomeFirstResponder(v.viewPtr) != 0
}

// ResignFirstResponder attempts to resign first responder status
func (v *View) ResignFirstResponder() bool {
	return C.View_ResignFirstResponder(v.viewPtr) != 0
}

// NextResponder returns the next responder in the responder chain
func (v *View) NextResponder() unsafe.Pointer {
	return C.View_NextResponder(v.viewPtr)
}

// Mouse events

// SetAcceptsMouseEvents sets whether the view accepts mouse events
func (v *View) SetAcceptsMouseEvents(accepts bool) {
	var acceptsVal C.int
	if accepts {
		acceptsVal = 1
	}
	C.View_SetAcceptsMouseEvents(v.viewPtr, acceptsVal)
}

// AcceptsMouseEvents returns true if the view accepts mouse events
func (v *View) AcceptsMouseEvents() bool {
	return C.View_AcceptsMouseEvents(v.viewPtr) != 0
}

// Focus ring and highlighting

// SetKeyboardFocusRingNeedsDisplayInRect marks the keyboard focus ring as needing display in a rect
func (v *View) SetKeyboardFocusRingNeedsDisplayInRect(x, y, width, height int) {
	C.View_SetKeyboardFocusRingNeedsDisplayInRect(v.viewPtr, C.int(x), C.int(y), C.int(width), C.int(height))
}

// SetFocusRingType sets the view's focus ring type
func (v *View) SetFocusRingType(ringType FocusRingType) {
	C.View_SetFocusRingType(v.viewPtr, C.int(ringType))
}

// GetFocusRingType returns the view's focus ring type
func (v *View) GetFocusRingType() FocusRingType {
	return FocusRingType(C.View_GetFocusRingType(v.viewPtr))
}

// Scrolling support

// IsFlipped returns true if the view uses flipped coordinates
func (v *View) IsFlipped() bool {
	return C.View_IsFlipped(v.viewPtr) != 0
}

// SetFlipped sets whether the view uses flipped coordinates (limited implementation)
func (v *View) SetFlipped(flipped bool) {
	var flippedVal C.int
	if flipped {
		flippedVal = 1
	}
	C.View_SetFlipped(v.viewPtr, flippedVal)
}

// ScrollRectToVisible scrolls the view to make the specified rect visible
func (v *View) ScrollRectToVisible(x, y, width, height int) {
	C.View_ScrollRectToVisible(v.viewPtr, C.int(x), C.int(y), C.int(width), C.int(height))
}

// ScrollPoint scrolls the view to the specified point
func (v *View) ScrollPoint(x, y int) {
	C.View_ScrollPoint(v.viewPtr, C.int(x), C.int(y))
}

// Window relationship

// GetWindow returns the window containing this view
func (v *View) GetWindow() unsafe.Pointer {
	return C.View_GetWindow(v.viewPtr)
}

// IsInLiveResize returns true if the view is currently being resized
func (v *View) IsInLiveResize() bool {
	return C.View_IsInLiveResize(v.viewPtr) != 0
}

// Layout and sizing

// SetFrameCenterRotation sets the view's rotation angle around its center
func (v *View) SetFrameCenterRotation(angle float64) {
	C.View_SetFrameCenterRotation(v.viewPtr, C.double(angle))
}

// GetFrameCenterRotation returns the view's rotation angle around its center
func (v *View) GetFrameCenterRotation() float64 {
	return float64(C.View_GetFrameCenterRotation(v.viewPtr))
}

// Tag and identification

// SetTag sets the view's tag
func (v *View) SetTag(tag int) {
	C.View_SetTag(v.viewPtr, C.int(tag))
}

// GetTag returns the view's tag
func (v *View) GetTag() int {
	return int(C.View_GetTag(v.viewPtr))
}

// ViewWithTag returns the view with the specified tag
func (v *View) ViewWithTag(tag int) *View {
	viewPtr := C.View_ViewWithTag(v.viewPtr, C.int(tag))
	if viewPtr == nil {
		return nil
	}
	return &View{viewPtr: viewPtr}
}

// Tooltip support

// SetToolTip sets the view's tooltip
func (v *View) SetToolTip(tooltip string) {
	cTooltip := C.CString(tooltip)
	defer C.free(unsafe.Pointer(cTooltip))
	C.View_SetToolTip(v.viewPtr, cTooltip)
}

// GetToolTip returns the view's tooltip
func (v *View) GetToolTip() string {
	cTooltip := C.View_GetToolTip(v.viewPtr)
	if cTooltip == nil {
		return ""
	}
	defer C.free(unsafe.Pointer(cTooltip))
	return C.GoString(cTooltip)
}

// Appearance and material (macOS 10.14+)

// SetAppearance sets the view's appearance
func (v *View) SetAppearance(appearanceType AppearanceType) {
	C.View_SetAppearance(v.viewPtr, C.int(appearanceType))
}

// GetEffectiveAppearance returns the view's effective appearance
func (v *View) GetEffectiveAppearance() AppearanceType {
	return AppearanceType(C.View_GetEffectiveAppearance(v.viewPtr))
}

// Animation support

// SetAnimations sets the view's animations dictionary
func (v *View) SetAnimations(animations unsafe.Pointer) {
	C.View_SetAnimations(v.viewPtr, animations)
}

// GetAnimations returns the view's animations dictionary
func (v *View) GetAnimations() unsafe.Pointer {
	return C.View_GetAnimations(v.viewPtr)
}

// Drawing context

// BitmapImageRepForCachingDisplayInRect returns a bitmap image rep for caching display
func (v *View) BitmapImageRepForCachingDisplayInRect(x, y, width, height int) unsafe.Pointer {
	return C.View_BitmapImageRepForCachingDisplayInRect(v.viewPtr, C.int(x), C.int(y), C.int(width), C.int(height))
}

// CacheDisplayInRect caches the display in the specified rect
func (v *View) CacheDisplayInRect(x, y, width, height int, bitmapImageRep unsafe.Pointer) {
	C.View_CacheDisplayInRect(v.viewPtr, C.int(x), C.int(y), C.int(width), C.int(height), bitmapImageRep)
}

// Invalidation

// InvalidateIntrinsicContentSize invalidates the view's intrinsic content size
func (v *View) InvalidateIntrinsicContentSize() {
	C.View_InvalidateIntrinsicContentSize(v.viewPtr)
}

// GetIntrinsicContentSize returns the view's intrinsic content size
func (v *View) GetIntrinsicContentSize() Size {
	var width, height C.int
	C.View_GetIntrinsicContentSize(v.viewPtr, &width, &height)
	return Size{Width: int(width), Height: int(height)}
}

// Utility method to get the underlying pointer (for integration with other AppKit types)
func (v *View) Ptr() unsafe.Pointer {
	return v.viewPtr
}
