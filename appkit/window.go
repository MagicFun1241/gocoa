package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #include "window.h"
import "C"
import "unsafe"

// WindowEvent - different window delegate Events
type WindowEvent int

const (
	didResize        WindowEvent = 0
	didMove          WindowEvent = 1
	didMiniaturize   WindowEvent = 2
	didDeminiaturize WindowEvent = 3
)

// WindowLevel constants for different window levels
type WindowLevel int

const (
	NormalWindowLevel            WindowLevel = 0
	FloatingWindowLevel          WindowLevel = 3
	SubmenuWindowLevel           WindowLevel = 3
	TornOffMenuWindowLevel       WindowLevel = 3
	MainMenuWindowLevel          WindowLevel = 24
	StatusWindowLevel            WindowLevel = 25
	ModalPanelWindowLevel        WindowLevel = 8
	PopUpMenuWindowLevel         WindowLevel = 101
	DraggingWindowLevel          WindowLevel = 500
	ScreenSaverWindowLevel       WindowLevel = 1000
	AssistiveTechHighWindowLevel WindowLevel = 1500
	DockWindowLevel              WindowLevel = 20
	OverlayWindowLevel           WindowLevel = 102
)

// EventHandler - handler functions that accepts the updated window as parameter
type EventHandler func(wnd *Window)

// Window is just that.
type Window struct {
	title     string
	x         int
	y         int
	w         int
	h         int
	callbacks map[WindowEvent]EventHandler
	winPtr    unsafe.Pointer
}

// Screen the screen of the window.
type Screen struct {
	X int
	Y int
}

var windows []*Window

// NewWindow constructs and returns a new window.
func NewWindow(title string, x int, y int, w int, h int) *Window {
	windowID := len(windows)
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))
	wnd := &Window{
		title:     title,
		x:         x,
		y:         y,
		w:         w,
		h:         h,
		callbacks: make(map[WindowEvent]EventHandler),
		winPtr:    C.Window_New(C.int(windowID), C.int(x), C.int(y), C.int(w), C.int(h), cTitle)}
	windows = append(windows, wnd)
	return wnd
}

// NewCenteredWindow constructs and returns a new window.
func NewCenteredWindow(title string, w int, h int) *Window {
	windowID := len(windows)
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))
	wnd := &Window{
		title:     title,
		w:         w,
		h:         h,
		callbacks: make(map[WindowEvent]EventHandler),
		winPtr:    C.Centered_Window_New(C.int(windowID), C.int(w), C.int(h), cTitle)}
	wnd.x = int(C.Screen_Center_X(wnd.winPtr))
	wnd.y = int(C.Screen_Center_Y(wnd.winPtr))
	windows = append(windows, wnd)
	return wnd
}

// GetScreen - returns the screen dimensions
func (wnd *Window) GetScreen() *Screen {
	return &Screen{
		X: int(C.Screen_X(wnd.winPtr)),
		Y: int(C.Screen_Y(wnd.winPtr))}
}

// MakeKeyAndOrderFront moves the window to the front of the screen list, within its
// level, and it shows the window.
func (wnd *Window) MakeKeyAndOrderFront() {
	C.Window_MakeKeyAndOrderFront(wnd.winPtr)
}

// OrderFront moves the window to the front of its level without changing key or main window
func (wnd *Window) OrderFront() {
	C.Window_OrderFront(wnd.winPtr)
}

// OrderBack moves the window to the back of its level without changing key or main window
func (wnd *Window) OrderBack() {
	C.Window_OrderBack(wnd.winPtr)
}

// OrderOut removes the window from the screen (hides it)
func (wnd *Window) OrderOut() {
	C.Window_OrderOut(wnd.winPtr)
}

// OrderFrontRegardless moves window to front even if app isn't active
func (wnd *Window) OrderFrontRegardless() {
	C.Window_OrderFrontRegardless(wnd.winPtr)
}

// Close removes the window from screen and releases it if configured to do so
func (wnd *Window) Close() {
	C.Window_Close(wnd.winPtr)
}

// PerformClose simulates clicking the close button (with validation)
func (wnd *Window) PerformClose() {
	C.Window_PerformClose(wnd.winPtr)
}

// Center positions the window in the center of the screen
func (wnd *Window) Center() {
	C.Window_Center(wnd.winPtr)
}

// Miniaturize minimizes the window to the dock
func (wnd *Window) Miniaturize() {
	C.Window_Miniaturize(wnd.winPtr)
}

// Deminiaturize restores window from dock
func (wnd *Window) Deminiaturize() {
	C.Window_Deminiaturize(wnd.winPtr)
}

// Zoom toggles window zoom state
func (wnd *Window) Zoom() {
	C.Window_Zoom(wnd.winPtr)
}

// MakeKeyWindow makes this window the key window
func (wnd *Window) MakeKeyWindow() {
	C.Window_MakeKeyWindow(wnd.winPtr)
}

// MakeMainWindow makes this window the main window
func (wnd *Window) MakeMainWindow() {
	C.Window_MakeMainWindow(wnd.winPtr)
}

// AddButton adds a Button to the window.
func (wnd *Window) AddButton(btn *Button) {
	C.Window_AddButton(wnd.winPtr, btn.buttonPtr)
}

// AddWebView adds a WebView to the window.
func (wnd *Window) AddWebView(webview *WebView) {
	C.Window_AddWebView(wnd.winPtr, webview.webviewPtr)
}

// AddDatePicker adds a DatePicker to the window.
func (wnd *Window) AddDatePicker(datePicker *DatePicker) {
	C.Window_AddDatePicker(wnd.winPtr, datePicker.datePickerPtr)
}

// AddTextView - adds a TextView to the window.
func (wnd *Window) AddTextView(tv *TextView) {
	C.Window_AddTextView(wnd.winPtr, tv.textViewPtr)
}

// AddTextField - adds a TextField to the window.
func (wnd *Window) AddTextField(tv *TextField) {
	C.Window_AddTextField(wnd.winPtr, tv.textFieldPtr)
}

// AddSearchField - adds a SearchField to the window.
func (wnd *Window) AddSearchField(sf *SearchField) {
	C.Window_AddSearchField(wnd.winPtr, sf.searchFieldPtr)
}

// AddLabel - adds a TextField to the window.
func (wnd *Window) AddLabel(tv *TextField) {
	C.Window_AddTextField(wnd.winPtr, tv.textFieldPtr)
}

// AddProgressIndicator adds a ProgressIndicator to the window.
func (wnd *Window) AddProgressIndicator(indicator *ProgressIndicator) {
	C.Window_AddProgressIndicator(wnd.winPtr, indicator.progressIndicatorPtr)
}

// AddImageView adds an ImageView to the window.
func (wnd *Window) AddImageView(imageView *ImageView) {
	C.Window_AddImageView(wnd.winPtr, imageView.imageViewPtr)
}

// AddSlider adds an Slider to the window.
func (wnd *Window) AddSlider(slider *Slider) {
	C.Window_AddSlider(wnd.winPtr, slider.sliderPtr)
}

// AddComboBox adds an ComboBox to the window.
func (wnd *Window) AddComboBox(comboBox *ComboBox) {
	C.Window_AddComboBox(wnd.winPtr, comboBox.comboBoxPtr)
}

// Update - forces the whole window to repaint
func (wnd *Window) Update() {
	C.Window_Update(wnd.winPtr)
}

// Display forces the window to redraw
func (wnd *Window) Display() {
	C.Window_Display(wnd.winPtr)
}

// InvalidateShadow marks the window shadow for redraw
func (wnd *Window) InvalidateShadow() {
	C.Window_InvalidateShadow(wnd.winPtr)
}

// Position and size getters
func (wnd *Window) GetX() int {
	return wnd.x
}

func (wnd *Window) GetY() int {
	return wnd.y
}

func (wnd *Window) GetWidth() int {
	return wnd.w
}

func (wnd *Window) GetHeight() int {
	return wnd.h
}

// SetFrame sets the window's frame rectangle
func (wnd *Window) SetFrame(x, y, width, height int, display bool) {
	if display {
		C.Window_SetFrameDisplay(wnd.winPtr, C.int(x), C.int(y), C.int(width), C.int(height), C.int(1))
	} else {
		C.Window_SetFrameDisplay(wnd.winPtr, C.int(x), C.int(y), C.int(width), C.int(height), C.int(0))
	}
	wnd.x = x
	wnd.y = y
	wnd.w = width
	wnd.h = height
}

// SetFrameOrigin sets just the window's origin
func (wnd *Window) SetFrameOrigin(x, y int) {
	C.Window_SetFrameOrigin(wnd.winPtr, C.int(x), C.int(y))
	wnd.x = x
	wnd.y = y
}

// SetContentSize sets the size of the content area
func (wnd *Window) SetContentSize(width, height int) {
	C.Window_SetContentSize(wnd.winPtr, C.int(width), C.int(height))
}

// SetMinSize sets the minimum window size
func (wnd *Window) SetMinSize(width, height int) {
	C.Window_SetMinSize(wnd.winPtr, C.int(width), C.int(height))
}

// SetMaxSize sets the maximum window size
func (wnd *Window) SetMaxSize(width, height int) {
	C.Window_SetMaxSize(wnd.winPtr, C.int(width), C.int(height))
}

// SetAspectRatio sets the window's aspect ratio
func (wnd *Window) SetAspectRatio(ratio float64) {
	C.Window_SetAspectRatio(wnd.winPtr, C.double(ratio))
}

// SetLevel sets the window's level (determines layering)
func (wnd *Window) SetLevel(level WindowLevel) {
	C.Window_SetLevel(wnd.winPtr, C.int(level))
}

// SetAlphaValue sets the window's transparency (0.0 = transparent, 1.0 = opaque)
func (wnd *Window) SetAlphaValue(alpha float64) {
	C.Window_SetAlphaValue(wnd.winPtr, C.double(alpha))
}

// GetAlphaValue gets the window's current alpha value
func (wnd *Window) GetAlphaValue() float64 {
	return float64(C.Window_GetAlphaValue(wnd.winPtr))
}

// SetOpaque sets whether the window is opaque
func (wnd *Window) SetOpaque(opaque bool) {
	if opaque {
		C.Window_SetOpaque(wnd.winPtr, C.int(1))
	} else {
		C.Window_SetOpaque(wnd.winPtr, C.int(0))
	}
}

// SetHasShadow sets whether the window has a shadow
func (wnd *Window) SetHasShadow(hasShadow bool) {
	if hasShadow {
		C.Window_SetHasShadow(wnd.winPtr, C.int(1))
	} else {
		C.Window_SetHasShadow(wnd.winPtr, C.int(0))
	}
}

// SetMovable sets whether the window can be moved by the user
func (wnd *Window) SetMovable(movable bool) {
	if movable {
		C.Window_SetMovable(wnd.winPtr, C.int(1))
	} else {
		C.Window_SetMovable(wnd.winPtr, C.int(0))
	}
}

// SetMovableByWindowBackground sets whether window can be moved by dragging background
func (wnd *Window) SetMovableByWindowBackground(movable bool) {
	if movable {
		C.Window_SetMovableByWindowBackground(wnd.winPtr, C.int(1))
	} else {
		C.Window_SetMovableByWindowBackground(wnd.winPtr, C.int(0))
	}
}

// SetReleasedWhenClosed sets whether window is released when closed
func (wnd *Window) SetReleasedWhenClosed(released bool) {
	if released {
		C.Window_SetReleasedWhenClosed(wnd.winPtr, C.int(1))
	} else {
		C.Window_SetReleasedWhenClosed(wnd.winPtr, C.int(0))
	}
}

// SetHidesOnDeactivate sets whether window hides when app becomes inactive
func (wnd *Window) SetHidesOnDeactivate(hides bool) {
	if hides {
		C.Window_SetHidesOnDeactivate(wnd.winPtr, C.int(1))
	} else {
		C.Window_SetHidesOnDeactivate(wnd.winPtr, C.int(0))
	}
}

// SetIgnoresMouseEvents sets whether window ignores mouse events
func (wnd *Window) SetIgnoresMouseEvents(ignores bool) {
	if ignores {
		C.Window_SetIgnoresMouseEvents(wnd.winPtr, C.int(1))
	} else {
		C.Window_SetIgnoresMouseEvents(wnd.winPtr, C.int(0))
	}
}

// SetAcceptsMouseMovedEvents sets whether window accepts mouse moved events
func (wnd *Window) SetAcceptsMouseMovedEvents(accepts bool) {
	if accepts {
		C.Window_SetAcceptsMouseMovedEvents(wnd.winPtr, C.int(1))
	} else {
		C.Window_SetAcceptsMouseMovedEvents(wnd.winPtr, C.int(0))
	}
}

// State query methods
func (wnd *Window) IsVisible() bool {
	return C.Window_IsVisible(wnd.winPtr) != 0
}

func (wnd *Window) IsKeyWindow() bool {
	return C.Window_IsKeyWindow(wnd.winPtr) != 0
}

func (wnd *Window) IsMainWindow() bool {
	return C.Window_IsMainWindow(wnd.winPtr) != 0
}

func (wnd *Window) IsMiniaturized() bool {
	return C.Window_IsMiniaturized(wnd.winPtr) != 0
}

func (wnd *Window) IsZoomed() bool {
	return C.Window_IsZoomed(wnd.winPtr) != 0
}

func (wnd *Window) CanBecomeKeyWindow() bool {
	return C.Window_CanBecomeKeyWindow(wnd.winPtr) != 0
}

func (wnd *Window) CanBecomeMainWindow() bool {
	return C.Window_CanBecomeMainWindow(wnd.winPtr) != 0
}

func (wnd *Window) IsOpaque() bool {
	return C.Window_IsOpaque(wnd.winPtr) != 0
}

func (wnd *Window) HasShadow() bool {
	return C.Window_HasShadow(wnd.winPtr) != 0
}

func (wnd *Window) IsMovable() bool {
	return C.Window_IsMovable(wnd.winPtr) != 0
}

func (wnd *Window) IsMovableByWindowBackground() bool {
	return C.Window_IsMovableByWindowBackground(wnd.winPtr) != 0
}

func (wnd *Window) IsReleasedWhenClosed() bool {
	return C.Window_IsReleasedWhenClosed(wnd.winPtr) != 0
}

func (wnd *Window) HidesOnDeactivate() bool {
	return C.Window_HidesOnDeactivate(wnd.winPtr) != 0
}

func (wnd *Window) IgnoresMouseEvents() bool {
	return C.Window_IgnoresMouseEvents(wnd.winPtr) != 0
}

func (wnd *Window) AcceptsMouseMovedEvents() bool {
	return C.Window_AcceptsMouseMovedEvents(wnd.winPtr) != 0
}

// GetLevel gets the current window level
func (wnd *Window) GetLevel() int {
	return int(C.Window_GetLevel(wnd.winPtr))
}

// Document handling
func (wnd *Window) SetDocumentEdited(edited bool) {
	if edited {
		C.Window_SetDocumentEdited(wnd.winPtr, C.int(1))
	} else {
		C.Window_SetDocumentEdited(wnd.winPtr, C.int(0))
	}
}

func (wnd *Window) IsDocumentEdited() bool {
	return C.Window_IsDocumentEdited(wnd.winPtr) != 0
}

func (wnd *Window) SetRepresentedFilename(filename string) {
	cFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cFilename))
	C.Window_SetRepresentedFilename(wnd.winPtr, cFilename)
}

func (wnd *Window) GetRepresentedFilename() string {
	cStr := C.Window_GetRepresentedFilename(wnd.winPtr)
	defer C.free(unsafe.Pointer(cStr))
	return C.GoString(cStr)
}

// GetContentView returns the window's content view
func (wnd *Window) GetContentView() *View {
	contentViewPtr := C.Window_GetContentView(wnd.winPtr)
	if contentViewPtr == nil {
		return nil
	}
	return &View{viewPtr: contentViewPtr}
}

// Title methods
func (wnd *Window) SetTitle(title string) {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))
	C.Window_SetTitle(wnd.winPtr, cTitle)
	wnd.title = title
}

func (wnd *Window) GetTitle() string {
	return wnd.title
}

// Button control methods
func (wnd *Window) SetMiniaturizeButtonEnabled(enabled bool) {
	if enabled {
		C.Window_SetMiniaturizeButtonEnabled(wnd.winPtr, C.int(1))
	} else {
		C.Window_SetMiniaturizeButtonEnabled(wnd.winPtr, C.int(0))
	}
}

func (wnd *Window) SetZoomButtonEnabled(enabled bool) {
	if enabled {
		C.Window_SetZoomButtonEnabled(wnd.winPtr, C.int(1))
	} else {
		C.Window_SetZoomButtonEnabled(wnd.winPtr, C.int(0))
	}
}

func (wnd *Window) SetCloseButtonEnabled(enabled bool) {
	if enabled {
		C.Window_SetCloseButtonEnabled(wnd.winPtr, C.int(1))
	} else {
		C.Window_SetCloseButtonEnabled(wnd.winPtr, C.int(0))
	}
}

func (wnd *Window) SetAllowsResizing(allowsResizing bool) {
	if allowsResizing {
		C.Window_SetAllowsResizing(wnd.winPtr, C.int(1))
	} else {
		C.Window_SetAllowsResizing(wnd.winPtr, C.int(0))
	}
}

// Menu handling
func (wnd *Window) AddDefaultQuitMenu() {
	C.Window_AddDefaultQuitMenu(wnd.winPtr)
}

// Event handlers
func (wnd *Window) OnDidResize(fn EventHandler) {
	wnd.callbacks[didResize] = fn
}

func (wnd *Window) OnDidMiniaturize(fn EventHandler) {
	wnd.callbacks[didMiniaturize] = fn
}

func (wnd *Window) OnDidDeminiaturize(fn EventHandler) {
	wnd.callbacks[didDeminiaturize] = fn
}

func (wnd *Window) OnDidMove(fn EventHandler) {
	wnd.callbacks[didMove] = fn
}

//export onWindowEvent
func onWindowEvent(id C.int, eventID C.int, x C.int, y C.int, w C.int, h C.int) {
	windowID := int(id)
	event := WindowEvent(eventID)
	if windowID < len(windows) && windows[windowID].callbacks[event] != nil {
		wnd := windows[windowID]
		windows[windowID].callbacks[event](&Window{
			title:  wnd.title,
			x:      int(x),
			y:      int(y),
			w:      int(w),
			h:      int(h),
			winPtr: wnd.winPtr})
	}
}
