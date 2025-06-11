#import <Cocoa/Cocoa.h>

#import <Foundation/Foundation.h>
#import <WebKit/WebKit.h>

#include "button.h"
#include "combobox.h"
#include "datepicker.h"
#include "imageview.h"
#include "textview.h"
#include "textfield.h"
#include "searchfield.h"
#include "webview.h"
#include "progressindicator.h"
#include "slider.h"
#include "view.h"

// Window creation and basic functions
void* Window_New(int goWindowID, int x, int y, int width, int height, const char* title);
void* Centered_Window_New(int goWindowID, int width, int height, const char* title);

// Screen utilities
int Screen_Center_X(void *wndPtr);
int Screen_Center_Y(void *wndPtr);
int Screen_X(void *wndPtr);
int Screen_Y(void *wndPtr);

// Window ordering and display
void Window_MakeKeyAndOrderFront(void *wndPtr);
void Window_OrderFront(void *wndPtr);
void Window_OrderBack(void *wndPtr);
void Window_OrderOut(void *wndPtr);
void Window_OrderFrontRegardless(void *wndPtr);
void Window_Close(void *wndPtr);
void Window_PerformClose(void *wndPtr);
void Window_Center(void *wndPtr);
void Window_MakeKeyWindow(void *wndPtr);
void Window_MakeMainWindow(void *wndPtr);

// Window state actions
void Window_Miniaturize(void *wndPtr);
void Window_Deminiaturize(void *wndPtr);
void Window_Zoom(void *wndPtr);

// Display and drawing
void Window_Display(void *wndPtr);
void Window_FlushWindow(void *wndPtr);
void Window_InvalidateShadow(void *wndPtr);
void Window_Update(void *wndPtr);

// Frame and size management
void Window_SetFrameDisplay(void *wndPtr, int x, int y, int width, int height, int display);
void Window_SetFrameOrigin(void *wndPtr, int x, int y);
void Window_SetContentSize(void *wndPtr, int width, int height);
void Window_SetMinSize(void *wndPtr, int width, int height);
void Window_SetMaxSize(void *wndPtr, int width, int height);
void Window_SetAspectRatio(void *wndPtr, double ratio);

// Window appearance
void Window_SetLevel(void *wndPtr, int level);
void Window_SetAlphaValue(void *wndPtr, double alpha);
double Window_GetAlphaValue(void *wndPtr);
void Window_SetOpaque(void *wndPtr, int opaque);
void Window_SetHasShadow(void *wndPtr, int hasShadow);

// Window behavior
void Window_SetMovable(void *wndPtr, int movable);
void Window_SetMovableByWindowBackground(void *wndPtr, int movable);
void Window_SetReleasedWhenClosed(void *wndPtr, int released);
void Window_SetHidesOnDeactivate(void *wndPtr, int hides);
void Window_SetIgnoresMouseEvents(void *wndPtr, int ignores);
void Window_SetAcceptsMouseMovedEvents(void *wndPtr, int accepts);

// Window state queries
int Window_IsVisible(void *wndPtr);
int Window_IsKeyWindow(void *wndPtr);
int Window_IsMainWindow(void *wndPtr);
int Window_IsMiniaturized(void *wndPtr);
int Window_IsZoomed(void *wndPtr);
int Window_CanBecomeKeyWindow(void *wndPtr);
int Window_CanBecomeMainWindow(void *wndPtr);
int Window_IsOpaque(void *wndPtr);
int Window_HasShadow(void *wndPtr);
int Window_IsMovable(void *wndPtr);
int Window_IsMovableByWindowBackground(void *wndPtr);
int Window_IsReleasedWhenClosed(void *wndPtr);
int Window_HidesOnDeactivate(void *wndPtr);
int Window_IgnoresMouseEvents(void *wndPtr);
int Window_AcceptsMouseMovedEvents(void *wndPtr);
int Window_GetLevel(void *wndPtr);

// Document handling
void Window_SetDocumentEdited(void *wndPtr, int edited);
int Window_IsDocumentEdited(void *wndPtr);
void Window_SetRepresentedFilename(void *wndPtr, const char* filename);
char* Window_GetRepresentedFilename(void *wndPtr);

// Title and controls
void Window_SetTitle(void *wndPtr, const char* title);
void Window_SetMiniaturizeButtonEnabled(void *wndPtr, int enabled);
void Window_SetZoomButtonEnabled(void *wndPtr, int enabled);
void Window_SetCloseButtonEnabled(void *wndPtr, int enabled);
void Window_SetAllowsResizing(void *wndPtr, int enabled);

// UI elements
void Window_AddButton(void *wndPtr, ButtonPtr btnPtr);
void Window_AddDatePicker(void *wndPtr, DatePickerPtr datePickerPtr);
void Window_AddTextView(void *wndPtr, TextViewPtr tvPtr);
void Window_AddTextField(void *wndPtr, TextFieldPtr tfPtr);
void Window_AddSearchField(void *wndPtr, SearchFieldPtr tfPtr);
void Window_AddProgressIndicator(void *wndPtr, ProgressIndicatorPtr progressIndicatorPtr);
void Window_AddImageView(void *wndPtr, ImageViewPtr imageViewPtr);
void Window_AddSlider(void *wndPtr, SliderPtr sliderPtr);
void Window_AddComboBox(void *wndPtr, ComboBoxPtr comboBoxPtr);
void Window_AddWebView(void *wndPtr, WebViewPtr wvPtr);

// Menu
void Window_AddDefaultQuitMenu(void *wndPtr);
