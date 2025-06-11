#import <Cocoa/Cocoa.h>
#import <Foundation/Foundation.h>

// View creation and basic functions
void* View_New(int x, int y, int width, int height);
void* View_NewWithFrame(int x, int y, int width, int height);
void View_Dealloc(void *viewPtr);

// View hierarchy management
void View_AddSubview(void *viewPtr, void *subviewPtr);
void View_RemoveFromSuperview(void *viewPtr);
void View_RemoveSubview(void *viewPtr, void *subviewPtr);
void* View_GetSuperview(void *viewPtr);
void** View_GetSubviews(void *viewPtr, int *count);
int View_IsDescendantOf(void *viewPtr, void *ancestorPtr);

// Frame and bounds management
void View_SetFrame(void *viewPtr, int x, int y, int width, int height);
void View_GetFrame(void *viewPtr, int *x, int *y, int *width, int *height);
void View_SetBounds(void *viewPtr, int x, int y, int width, int height);
void View_GetBounds(void *viewPtr, int *x, int *y, int *width, int *height);
void View_SetFrameOrigin(void *viewPtr, int x, int y);
void View_SetFrameSize(void *viewPtr, int width, int height);
void View_SetBoundsOrigin(void *viewPtr, int x, int y);
void View_SetBoundsSize(void *viewPtr, int width, int height);

// Display and drawing
void View_SetNeedsDisplay(void *viewPtr, int needsDisplay);
void View_SetNeedsDisplayInRect(void *viewPtr, int x, int y, int width, int height);
void View_Display(void *viewPtr);
void View_DisplayIfNeeded(void *viewPtr);
void View_DisplayRect(void *viewPtr, int x, int y, int width, int height);
int View_NeedsDisplay(void *viewPtr);

// Coordinate conversion
void View_ConvertPointToView(void *viewPtr, int x, int y, void *toViewPtr, int *outX, int *outY);
void View_ConvertPointFromView(void *viewPtr, int x, int y, void *fromViewPtr, int *outX, int *outY);
void View_ConvertRectToView(void *viewPtr, int x, int y, int width, int height, void *toViewPtr, int *outX, int *outY, int *outWidth, int *outHeight);
void View_ConvertRectFromView(void *viewPtr, int x, int y, int width, int height, void *fromViewPtr, int *outX, int *outY, int *outWidth, int *outHeight);

// Hit testing
void* View_HitTest(void *viewPtr, int x, int y);
int View_IsOpaque(void *viewPtr);
void View_SetOpaque(void *viewPtr, int opaque);

// Visibility and clipping
int View_IsHidden(void *viewPtr);
void View_SetHidden(void *viewPtr, int hidden);
int View_IsHiddenOrHasHiddenAncestor(void *viewPtr);
void View_SetClipsToBounds(void *viewPtr, int clips);
int View_ClipsToBounds(void *viewPtr);

// Auto layout and constraints
void View_SetTranslatesAutoresizingMaskIntoConstraints(void *viewPtr, int translates);
int View_TranslatesAutoresizingMaskIntoConstraints(void *viewPtr);
void View_SetAutoresizingMask(void *viewPtr, int mask);
int View_GetAutoresizingMask(void *viewPtr);

// Layer and appearance
void View_SetWantsLayer(void *viewPtr, int wantsLayer);
int View_WantsLayer(void *viewPtr);
void* View_GetLayer(void *viewPtr);
void View_SetLayer(void *viewPtr, void *layerPtr);
void View_SetAlphaValue(void *viewPtr, double alpha);
double View_GetAlphaValue(void *viewPtr);

// Responder chain and events
int View_AcceptsFirstResponder(void *viewPtr);
int View_BecomeFirstResponder(void *viewPtr);
int View_ResignFirstResponder(void *viewPtr);
void* View_NextResponder(void *viewPtr);

// Mouse events
void View_SetAcceptsMouseEvents(void *viewPtr, int accepts);
int View_AcceptsMouseEvents(void *viewPtr);

// Focus ring and highlighting
void View_SetKeyboardFocusRingNeedsDisplayInRect(void *viewPtr, int x, int y, int width, int height);
void View_SetFocusRingType(void *viewPtr, int type);
int View_GetFocusRingType(void *viewPtr);

// Scrolling support
int View_IsFlipped(void *viewPtr);
void View_SetFlipped(void *viewPtr, int flipped);
void View_ScrollRectToVisible(void *viewPtr, int x, int y, int width, int height);
void View_ScrollPoint(void *viewPtr, int x, int y);

// Window relationship
void* View_GetWindow(void *viewPtr);
int View_IsInLiveResize(void *viewPtr);

// Layout and sizing
void View_SetFrameCenterRotation(void *viewPtr, double angle);
double View_GetFrameCenterRotation(void *viewPtr);

// Tag and identification
void View_SetTag(void *viewPtr, int tag);
int View_GetTag(void *viewPtr);
void* View_ViewWithTag(void *viewPtr, int tag);

// Tooltip support
void View_SetToolTip(void *viewPtr, const char* tooltip);
char* View_GetToolTip(void *viewPtr);

// Appearance and material (macOS 10.14+)
void View_SetAppearance(void *viewPtr, int appearanceType);
int View_GetEffectiveAppearance(void *viewPtr);

// Animation support
void View_SetAnimations(void *viewPtr, void *animationsDict);
void* View_GetAnimations(void *viewPtr);

// Drawing context
void* View_BitmapImageRepForCachingDisplayInRect(void *viewPtr, int x, int y, int width, int height);
void View_CacheDisplayInRect(void *viewPtr, int x, int y, int width, int height, void *bitmapImageRep);

// Invalidation
void View_InvalidateIntrinsicContentSize(void *viewPtr);
void View_GetIntrinsicContentSize(void *viewPtr, int *width, int *height);
