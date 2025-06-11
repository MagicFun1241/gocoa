#import "window.h"
#import "windowdelegate.h"
#include "_cgo_export.h"

WindowDelegate *gocoa_windowDelegate = nil;

void* Window_New(int goWindowID, int x, int y, int width, int height, const char* title) 
{
    NSRect windowRect = NSMakeRect(x, y, width, height);
    id window = [[NSWindow alloc] initWithContentRect:windowRect 
        styleMask:(NSWindowStyleMaskTitled | NSWindowStyleMaskClosable | NSWindowStyleMaskResizable | NSWindowStyleMaskMiniaturizable)
        backing:NSBackingStoreBuffered
        defer:NO];
    [window setTitle:[NSString stringWithUTF8String:title]];
    [window autorelease];

    gocoa_windowDelegate = [[WindowDelegate alloc] init];
    [gocoa_windowDelegate setGoWindowID:goWindowID];
    [window setDelegate:gocoa_windowDelegate];
    return window;
}

void* Centered_Window_New(int goWindowID, int width, int height, const char* title) 
{
    NSRect windowRect = NSMakeRect(0, 0, width, height);
    id window = [[NSWindow alloc] initWithContentRect:windowRect 
        styleMask:(NSWindowStyleMaskTitled | NSWindowStyleMaskClosable | NSWindowStyleMaskResizable | NSWindowStyleMaskMiniaturizable)
        backing:NSBackingStoreBuffered
        defer:NO];
    [window setTitle:[NSString stringWithUTF8String:title]];
    [window autorelease];
    CGFloat xPos = NSWidth([[window screen] frame])/2 - NSWidth([window frame])/2;
    CGFloat yPos = NSHeight([[window screen] frame])/2 - NSHeight([window frame])/2;
    gocoa_windowDelegate = [[WindowDelegate alloc] init];
    [gocoa_windowDelegate setGoWindowID:goWindowID];
    [window setDelegate:gocoa_windowDelegate];
    [window setFrame:NSMakeRect(xPos, yPos, NSWidth([window frame]), NSHeight([window frame])) display:YES];
    
    return window;
}

int Screen_Center_X(void *wndPtr)
{
    NSWindow* window = (NSWindow*)wndPtr;
    CGFloat xPos = NSWidth([[window screen] frame])/2 - NSWidth([window frame])/2;
    return (int)(xPos);
}

int Screen_Center_Y(void *wndPtr)
{
    NSWindow* window = (NSWindow*)wndPtr;
    CGFloat yPos = NSHeight([[window screen] frame])/2 - NSHeight([window frame])/2;
    return (int)(yPos);
}

int Screen_X(void *wndPtr)
{
    NSWindow* window = (NSWindow*)wndPtr;
    CGFloat xPos = NSWidth([[window screen] frame]);
    return (int)(xPos);
}

int Screen_Y(void *wndPtr)
{
    NSWindow* window = (NSWindow*)wndPtr;
    CGFloat yPos = NSHeight([[window screen] frame]);
    return (int)(yPos);
}

// Window ordering and display
void Window_MakeKeyAndOrderFront(void *wndPtr)
{
    NSWindow* window = (NSWindow*)wndPtr;
    [window makeKeyAndOrderFront:nil];
}

void Window_OrderFront(void *wndPtr)
{
    NSWindow* window = (NSWindow*)wndPtr;
    [window orderFront:nil];
}

void Window_OrderBack(void *wndPtr)
{
    NSWindow* window = (NSWindow*)wndPtr;
    [window orderBack:nil];
}

void Window_OrderOut(void *wndPtr)
{
    NSWindow* window = (NSWindow*)wndPtr;
    [window orderOut:nil];
}

void Window_OrderFrontRegardless(void *wndPtr)
{
    NSWindow* window = (NSWindow*)wndPtr;
    [window orderFrontRegardless];
}

void Window_Close(void *wndPtr)
{
    NSWindow* window = (NSWindow*)wndPtr;
    [window close];
}

void Window_PerformClose(void *wndPtr)
{
    NSWindow* window = (NSWindow*)wndPtr;
    [window performClose:nil];
}

void Window_Center(void *wndPtr)
{
    NSWindow* window = (NSWindow*)wndPtr;
    [window center];
}

void Window_MakeKeyWindow(void *wndPtr)
{
    NSWindow* window = (NSWindow*)wndPtr;
    [window makeKeyWindow];
}

void Window_MakeMainWindow(void *wndPtr)
{
    NSWindow* window = (NSWindow*)wndPtr;
    [window makeMainWindow];
}

// Window state actions
void Window_Miniaturize(void *wndPtr)
{
    NSWindow* window = (NSWindow*)wndPtr;
    [window miniaturize:nil];
}

void Window_Deminiaturize(void *wndPtr)
{
    NSWindow* window = (NSWindow*)wndPtr;
    [window deminiaturize:nil];
}

void Window_Zoom(void *wndPtr)
{
    NSWindow* window = (NSWindow*)wndPtr;
    [window zoom:nil];
}

// Display and drawing
void Window_Display(void *wndPtr)
{
    NSWindow* window = (NSWindow*)wndPtr;
    [window display];
}

void Window_InvalidateShadow(void *wndPtr)
{
    NSWindow* window = (NSWindow*)wndPtr;
    [window invalidateShadow];
}

void Window_Update(void *wndPtr)
{
    NSWindow* window = (NSWindow*)wndPtr;
    [[window contentView] setNeedsDisplay:YES];
}

// Frame and size management
void Window_SetFrameDisplay(void *wndPtr, int x, int y, int width, int height, int display)
{
    NSWindow* window = (NSWindow*)wndPtr;
    NSRect frame = NSMakeRect(x, y, width, height);
    [window setFrame:frame display:(display != 0)];
}

void Window_SetFrameOrigin(void *wndPtr, int x, int y)
{
    NSWindow* window = (NSWindow*)wndPtr;
    NSPoint origin = NSMakePoint(x, y);
    [window setFrameOrigin:origin];
}

void Window_SetContentSize(void *wndPtr, int width, int height)
{
    NSWindow* window = (NSWindow*)wndPtr;
    NSSize size = NSMakeSize(width, height);
    [window setContentSize:size];
}

void Window_SetMinSize(void *wndPtr, int width, int height)
{
    NSWindow* window = (NSWindow*)wndPtr;
    NSSize size = NSMakeSize(width, height);
    [window setMinSize:size];
}

void Window_SetMaxSize(void *wndPtr, int width, int height)
{
    NSWindow* window = (NSWindow*)wndPtr;
    NSSize size = NSMakeSize(width, height);
    [window setMaxSize:size];
}

void Window_SetAspectRatio(void *wndPtr, double ratio)
{
    NSWindow* window = (NSWindow*)wndPtr;
    NSSize aspectRatio = NSMakeSize(ratio, 1.0);
    [window setAspectRatio:aspectRatio];
}

// Window appearance
void Window_SetLevel(void *wndPtr, int level)
{
    NSWindow* window = (NSWindow*)wndPtr;
    [window setLevel:level];
}

void Window_SetAlphaValue(void *wndPtr, double alpha)
{
    NSWindow* window = (NSWindow*)wndPtr;
    [window setAlphaValue:alpha];
}

double Window_GetAlphaValue(void *wndPtr)
{
    NSWindow* window = (NSWindow*)wndPtr;
    return [window alphaValue];
}

void Window_SetOpaque(void *wndPtr, int opaque)
{
    NSWindow* window = (NSWindow*)wndPtr;
    [window setOpaque:(opaque != 0)];
}

void Window_SetHasShadow(void *wndPtr, int hasShadow)
{
    NSWindow* window = (NSWindow*)wndPtr;
    [window setHasShadow:(hasShadow != 0)];
}

// Window behavior
void Window_SetMovable(void *wndPtr, int movable)
{
    NSWindow* window = (NSWindow*)wndPtr;
    [window setMovable:(movable != 0)];
}

void Window_SetMovableByWindowBackground(void *wndPtr, int movable)
{
    NSWindow* window = (NSWindow*)wndPtr;
    [window setMovableByWindowBackground:(movable != 0)];
}

void Window_SetReleasedWhenClosed(void *wndPtr, int released)
{
    NSWindow* window = (NSWindow*)wndPtr;
    [window setReleasedWhenClosed:(released != 0)];
}

void Window_SetHidesOnDeactivate(void *wndPtr, int hides)
{
    NSWindow* window = (NSWindow*)wndPtr;
    [window setHidesOnDeactivate:(hides != 0)];
}

void Window_SetIgnoresMouseEvents(void *wndPtr, int ignores)
{
    NSWindow* window = (NSWindow*)wndPtr;
    [window setIgnoresMouseEvents:(ignores != 0)];
}

void Window_SetAcceptsMouseMovedEvents(void *wndPtr, int accepts)
{
    NSWindow* window = (NSWindow*)wndPtr;
    [window setAcceptsMouseMovedEvents:(accepts != 0)];
}

// Window state queries
int Window_IsVisible(void *wndPtr)
{
    NSWindow* window = (NSWindow*)wndPtr;
    return [window isVisible] ? 1 : 0;
}

int Window_IsKeyWindow(void *wndPtr)
{
    NSWindow* window = (NSWindow*)wndPtr;
    return [window isKeyWindow] ? 1 : 0;
}

int Window_IsMainWindow(void *wndPtr)
{
    NSWindow* window = (NSWindow*)wndPtr;
    return [window isMainWindow] ? 1 : 0;
}

int Window_IsMiniaturized(void *wndPtr)
{
    NSWindow* window = (NSWindow*)wndPtr;
    return [window isMiniaturized] ? 1 : 0;
}

int Window_IsZoomed(void *wndPtr)
{
    NSWindow* window = (NSWindow*)wndPtr;
    return [window isZoomed] ? 1 : 0;
}

int Window_CanBecomeKeyWindow(void *wndPtr)
{
    NSWindow* window = (NSWindow*)wndPtr;
    return [window canBecomeKeyWindow] ? 1 : 0;
}

int Window_CanBecomeMainWindow(void *wndPtr)
{
    NSWindow* window = (NSWindow*)wndPtr;
    return [window canBecomeMainWindow] ? 1 : 0;
}

int Window_IsOpaque(void *wndPtr)
{
    NSWindow* window = (NSWindow*)wndPtr;
    return [window isOpaque] ? 1 : 0;
}

int Window_HasShadow(void *wndPtr)
{
    NSWindow* window = (NSWindow*)wndPtr;
    return [window hasShadow] ? 1 : 0;
}

int Window_IsMovable(void *wndPtr)
{
    NSWindow* window = (NSWindow*)wndPtr;
    return [window isMovable] ? 1 : 0;
}

int Window_IsMovableByWindowBackground(void *wndPtr)
{
    NSWindow* window = (NSWindow*)wndPtr;
    return [window isMovableByWindowBackground] ? 1 : 0;
}

int Window_IsReleasedWhenClosed(void *wndPtr)
{
    NSWindow* window = (NSWindow*)wndPtr;
    return [window isReleasedWhenClosed] ? 1 : 0;
}

int Window_HidesOnDeactivate(void *wndPtr)
{
    NSWindow* window = (NSWindow*)wndPtr;
    return [window hidesOnDeactivate] ? 1 : 0;
}

int Window_IgnoresMouseEvents(void *wndPtr)
{
    NSWindow* window = (NSWindow*)wndPtr;
    return [window ignoresMouseEvents] ? 1 : 0;
}

int Window_AcceptsMouseMovedEvents(void *wndPtr)
{
    NSWindow* window = (NSWindow*)wndPtr;
    return [window acceptsMouseMovedEvents] ? 1 : 0;
}

int Window_GetLevel(void *wndPtr)
{
    NSWindow* window = (NSWindow*)wndPtr;
    return (int)[window level];
}

// Document handling
void Window_SetDocumentEdited(void *wndPtr, int edited)
{
    NSWindow* window = (NSWindow*)wndPtr;
    [window setDocumentEdited:(edited != 0)];
}

int Window_IsDocumentEdited(void *wndPtr)
{
    NSWindow* window = (NSWindow*)wndPtr;
    return [window isDocumentEdited] ? 1 : 0;
}

void Window_SetRepresentedFilename(void *wndPtr, const char* filename)
{
    NSWindow* window = (NSWindow*)wndPtr;
    NSString* filenameStr = [NSString stringWithUTF8String:filename];
    [window setRepresentedFilename:filenameStr];
}

char* Window_GetRepresentedFilename(void *wndPtr)
{
    NSWindow* window = (NSWindow*)wndPtr;
    NSString* filename = [window representedFilename];
    const char* cString = [filename UTF8String];
    // Allocate memory for the string that Go can free
    char* result = malloc(strlen(cString) + 1);
    strcpy(result, cString);
    return result;
}

// Title and controls
void Window_SetTitle(void *wndPtr, const char* title)
{
    NSWindow* window = (NSWindow*)wndPtr;
    [window setTitle:[NSString stringWithUTF8String:title]];
}

void Window_SetMiniaturizeButtonEnabled(void *wndPtr, int enabled) {
    NSWindow* window = (NSWindow*)wndPtr;
    NSButton *button = [window standardWindowButton:NSWindowMiniaturizeButton];
    [button setEnabled: enabled];
}

void Window_SetZoomButtonEnabled(void *wndPtr, int enabled) {
    NSWindow* window = (NSWindow*)wndPtr;
    NSButton *button = [window standardWindowButton:NSWindowZoomButton];
    [button setEnabled: enabled];
}

void Window_SetCloseButtonEnabled(void *wndPtr, int enabled) {
    NSWindow* window = (NSWindow*)wndPtr;
    NSButton *button = [window standardWindowButton:NSWindowCloseButton];
    [button setEnabled: enabled];
}

void Window_SetAllowsResizing(void *wndPtr, int allowsResizing) {
    NSWindow* window = (NSWindow*)wndPtr;
    if(allowsResizing) {
        window.styleMask |= NSWindowStyleMaskResizable;
    } else {
        window.styleMask &= ~NSWindowStyleMaskResizable;
    }
}

// UI elements
void Window_AddButton(void *wndPtr, ButtonPtr btnPtr) 
{
    NSButton* button = (NSButton*)btnPtr;
    NSWindow* window = (NSWindow*)wndPtr;
    [[window contentView] addSubview:button];
}

void Window_AddDatePicker(void *wndPtr, DatePickerPtr datePickerPtr)
{
    NSDatePicker* datePicker = (NSDatePicker*)datePickerPtr;
    NSWindow* window = (NSWindow*)wndPtr;
    [[window contentView] addSubview:datePicker];
}

void Window_AddTextView(void *wndPtr, TextViewPtr tvPtr)
{
    NSTextView* textview = (NSTextView*)tvPtr;
    NSWindow* window = (NSWindow*)wndPtr;
    [[window contentView] addSubview:textview];
}

void Window_AddTextField(void *wndPtr, TextFieldPtr tfPtr)
{
    NSTextField* textfield = (NSTextField*)tfPtr;
    NSWindow* window = (NSWindow*)wndPtr;
    [[window contentView] addSubview:textfield];
}

void Window_AddWebView(void *wndPtr, WebViewPtr wvPtr)
{
    WKWebView* webview = (WKWebView*)wvPtr;
    NSWindow* window = (NSWindow*)wndPtr;
    [[window contentView] addSubview:webview];
}

void Window_AddSearchField(void *wndPtr, SearchFieldPtr tfPtr)
{
    NSSearchField* searchfield = (NSSearchField*)tfPtr;
    NSWindow* window = (NSWindow*)wndPtr;
    [[window contentView] addSubview:searchfield];
}

void Window_AddProgressIndicator(void *wndPtr, ProgressIndicatorPtr progressIndicatorPtr)
{
    NSProgressIndicator* indicator = (NSProgressIndicator*)progressIndicatorPtr;
    NSWindow* window = (NSWindow*)wndPtr;
    [[window contentView] addSubview:indicator];
}

void Window_AddImageView(void *wndPtr, ImageViewPtr imageViewPtr)
{
    NSImageView* imageView = (NSImageView*)imageViewPtr;
    NSWindow* window = (NSWindow*)wndPtr;
    [[window contentView] addSubview:imageView];
}

void Window_AddSlider(void *wndPtr, SliderPtr sliderPtr)
{
    NSSlider* slider = (NSSlider*)sliderPtr;
    NSWindow* window = (NSWindow*)wndPtr;
    [[window contentView] addSubview:slider];
}

void Window_AddComboBox(void *wndPtr, ComboBoxPtr comboBoxPtr)
{
    NSComboBox* comboBox = (NSComboBox*)comboBoxPtr;
    NSWindow* window = (NSWindow*)wndPtr;
    [[window contentView] addSubview:comboBox];
}

// Menu
void Window_AddDefaultQuitMenu(void *wndPtr) {
    NSWindow* window = (NSWindow*)wndPtr;

    id menubar = [[NSMenu new] autorelease];
    id appMenuItem = [[NSMenuItem new] autorelease];
    [menubar addItem:appMenuItem];
    [NSApp setMainMenu:menubar];
    id appMenu = [[NSMenu new] autorelease];
    id appName = [[NSProcessInfo processInfo] processName];
    id quitTitle = [@"Quit " stringByAppendingString:appName];
    id quitMenuItem = [[[NSMenuItem alloc] initWithTitle:quitTitle
        action:@selector(terminate:) keyEquivalent:@"q"] autorelease];
    [appMenu addItem:quitMenuItem];
    [appMenuItem setSubmenu:appMenu];
}
