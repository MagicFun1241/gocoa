#import "view.h"
#include <stdlib.h>
#include <string.h>
#import <objc/runtime.h>

// View creation and basic functions
void* View_New(int x, int y, int width, int height)
{
    NSRect frame = NSMakeRect(x, y, width, height);
    NSView* view = [[NSView alloc] initWithFrame:frame];
    [view autorelease];
    return view;
}

void* View_NewWithFrame(int x, int y, int width, int height)
{
    return View_New(x, y, width, height);
}

void View_Dealloc(void *viewPtr)
{
    NSView* view = (NSView*)viewPtr;
    [view release];
}

// View hierarchy management
void View_AddSubview(void *viewPtr, void *subviewPtr)
{
    NSView* view = (NSView*)viewPtr;
    NSView* subview = (NSView*)subviewPtr;
    [view addSubview:subview];
}

void View_RemoveFromSuperview(void *viewPtr)
{
    NSView* view = (NSView*)viewPtr;
    [view removeFromSuperview];
}

void View_RemoveSubview(void *viewPtr, void *subviewPtr)
{
    NSView* view = (NSView*)viewPtr;
    NSView* subview = (NSView*)subviewPtr;
    if ([view.subviews containsObject:subview]) {
        [subview removeFromSuperview];
    }
}

void* View_GetSuperview(void *viewPtr)
{
    NSView* view = (NSView*)viewPtr;
    return view.superview;
}

void** View_GetSubviews(void *viewPtr, int *count)
{
    NSView* view = (NSView*)viewPtr;
    NSArray* subviews = view.subviews;
    *count = (int)subviews.count;
    
    if (*count == 0) {
        return NULL;
    }
    
    void** result = malloc(sizeof(void*) * (*count));
    for (int i = 0; i < *count; i++) {
        result[i] = subviews[i];
    }
    return result;
}

int View_IsDescendantOf(void *viewPtr, void *ancestorPtr)
{
    NSView* view = (NSView*)viewPtr;
    NSView* ancestor = (NSView*)ancestorPtr;
    return [view isDescendantOf:ancestor] ? 1 : 0;
}

// Frame and bounds management
void View_SetFrame(void *viewPtr, int x, int y, int width, int height)
{
    NSView* view = (NSView*)viewPtr;
    NSRect frame = NSMakeRect(x, y, width, height);
    [view setFrame:frame];
}

void View_GetFrame(void *viewPtr, int *x, int *y, int *width, int *height)
{
    NSView* view = (NSView*)viewPtr;
    NSRect frame = view.frame;
    *x = (int)frame.origin.x;
    *y = (int)frame.origin.y;
    *width = (int)frame.size.width;
    *height = (int)frame.size.height;
}

void View_SetBounds(void *viewPtr, int x, int y, int width, int height)
{
    NSView* view = (NSView*)viewPtr;
    NSRect bounds = NSMakeRect(x, y, width, height);
    [view setBounds:bounds];
}

void View_GetBounds(void *viewPtr, int *x, int *y, int *width, int *height)
{
    NSView* view = (NSView*)viewPtr;
    NSRect bounds = view.bounds;
    *x = (int)bounds.origin.x;
    *y = (int)bounds.origin.y;
    *width = (int)bounds.size.width;
    *height = (int)bounds.size.height;
}

void View_SetFrameOrigin(void *viewPtr, int x, int y)
{
    NSView* view = (NSView*)viewPtr;
    NSPoint origin = NSMakePoint(x, y);
    [view setFrameOrigin:origin];
}

void View_SetFrameSize(void *viewPtr, int width, int height)
{
    NSView* view = (NSView*)viewPtr;
    NSSize size = NSMakeSize(width, height);
    [view setFrameSize:size];
}

void View_SetBoundsOrigin(void *viewPtr, int x, int y)
{
    NSView* view = (NSView*)viewPtr;
    NSPoint origin = NSMakePoint(x, y);
    [view setBoundsOrigin:origin];
}

void View_SetBoundsSize(void *viewPtr, int width, int height)
{
    NSView* view = (NSView*)viewPtr;
    NSSize size = NSMakeSize(width, height);
    [view setBoundsSize:size];
}

// Display and drawing
void View_SetNeedsDisplay(void *viewPtr, int needsDisplay)
{
    NSView* view = (NSView*)viewPtr;
    [view setNeedsDisplay:(needsDisplay != 0)];
}

void View_SetNeedsDisplayInRect(void *viewPtr, int x, int y, int width, int height)
{
    NSView* view = (NSView*)viewPtr;
    NSRect rect = NSMakeRect(x, y, width, height);
    [view setNeedsDisplayInRect:rect];
}

void View_Display(void *viewPtr)
{
    NSView* view = (NSView*)viewPtr;
    [view display];
}

void View_DisplayIfNeeded(void *viewPtr)
{
    NSView* view = (NSView*)viewPtr;
    [view displayIfNeeded];
}

void View_DisplayRect(void *viewPtr, int x, int y, int width, int height)
{
    NSView* view = (NSView*)viewPtr;
    NSRect rect = NSMakeRect(x, y, width, height);
    [view displayRect:rect];
}

int View_NeedsDisplay(void *viewPtr)
{
    NSView* view = (NSView*)viewPtr;
    return [view needsDisplay] ? 1 : 0;
}

// Coordinate conversion
void View_ConvertPointToView(void *viewPtr, int x, int y, void *toViewPtr, int *outX, int *outY)
{
    NSView* view = (NSView*)viewPtr;
    NSView* toView = (NSView*)toViewPtr;
    NSPoint point = NSMakePoint(x, y);
    NSPoint converted = [view convertPoint:point toView:toView];
    *outX = (int)converted.x;
    *outY = (int)converted.y;
}

void View_ConvertPointFromView(void *viewPtr, int x, int y, void *fromViewPtr, int *outX, int *outY)
{
    NSView* view = (NSView*)viewPtr;
    NSView* fromView = (NSView*)fromViewPtr;
    NSPoint point = NSMakePoint(x, y);
    NSPoint converted = [view convertPoint:point fromView:fromView];
    *outX = (int)converted.x;
    *outY = (int)converted.y;
}

void View_ConvertRectToView(void *viewPtr, int x, int y, int width, int height, void *toViewPtr, int *outX, int *outY, int *outWidth, int *outHeight)
{
    NSView* view = (NSView*)viewPtr;
    NSView* toView = (NSView*)toViewPtr;
    NSRect rect = NSMakeRect(x, y, width, height);
    NSRect converted = [view convertRect:rect toView:toView];
    *outX = (int)converted.origin.x;
    *outY = (int)converted.origin.y;
    *outWidth = (int)converted.size.width;
    *outHeight = (int)converted.size.height;
}

void View_ConvertRectFromView(void *viewPtr, int x, int y, int width, int height, void *fromViewPtr, int *outX, int *outY, int *outWidth, int *outHeight)
{
    NSView* view = (NSView*)viewPtr;
    NSView* fromView = (NSView*)fromViewPtr;
    NSRect rect = NSMakeRect(x, y, width, height);
    NSRect converted = [view convertRect:rect fromView:fromView];
    *outX = (int)converted.origin.x;
    *outY = (int)converted.origin.y;
    *outWidth = (int)converted.size.width;
    *outHeight = (int)converted.size.height;
}

// Hit testing
void* View_HitTest(void *viewPtr, int x, int y)
{
    NSView* view = (NSView*)viewPtr;
    NSPoint point = NSMakePoint(x, y);
    return [view hitTest:point];
}

int View_IsOpaque(void *viewPtr)
{
    NSView* view = (NSView*)viewPtr;
    return [view isOpaque] ? 1 : 0;
}

void View_SetOpaque(void *viewPtr, int opaque)
{
    // Note: NSView doesn't have a direct setOpaque method
    // This would typically be handled by subclasses
    // For now, we'll implement as a no-op or store in associated object
}

// Visibility and clipping
int View_IsHidden(void *viewPtr)
{
    NSView* view = (NSView*)viewPtr;
    return [view isHidden] ? 1 : 0;
}

void View_SetHidden(void *viewPtr, int hidden)
{
    NSView* view = (NSView*)viewPtr;
    [view setHidden:(hidden != 0)];
}

int View_IsHiddenOrHasHiddenAncestor(void *viewPtr)
{
    NSView* view = (NSView*)viewPtr;
    return [view isHiddenOrHasHiddenAncestor] ? 1 : 0;
}

void View_SetClipsToBounds(void *viewPtr, int clips)
{
    NSView* view = (NSView*)viewPtr;
    if (clips) {
        view.layer.masksToBounds = YES;
    } else {
        view.layer.masksToBounds = NO;
    }
}

int View_ClipsToBounds(void *viewPtr)
{
    NSView* view = (NSView*)viewPtr;
    return view.layer.masksToBounds ? 1 : 0;
}

// Auto layout and constraints
void View_SetTranslatesAutoresizingMaskIntoConstraints(void *viewPtr, int translates)
{
    NSView* view = (NSView*)viewPtr;
    [view setTranslatesAutoresizingMaskIntoConstraints:(translates != 0)];
}

int View_TranslatesAutoresizingMaskIntoConstraints(void *viewPtr)
{
    NSView* view = (NSView*)viewPtr;
    return [view translatesAutoresizingMaskIntoConstraints] ? 1 : 0;
}

void View_SetAutoresizingMask(void *viewPtr, int mask)
{
    NSView* view = (NSView*)viewPtr;
    [view setAutoresizingMask:mask];
}

int View_GetAutoresizingMask(void *viewPtr)
{
    NSView* view = (NSView*)viewPtr;
    return (int)[view autoresizingMask];
}

// Layer and appearance
void View_SetWantsLayer(void *viewPtr, int wantsLayer)
{
    NSView* view = (NSView*)viewPtr;
    [view setWantsLayer:(wantsLayer != 0)];
}

int View_WantsLayer(void *viewPtr)
{
    NSView* view = (NSView*)viewPtr;
    return [view wantsLayer] ? 1 : 0;
}

void* View_GetLayer(void *viewPtr)
{
    NSView* view = (NSView*)viewPtr;
    return view.layer;
}

void View_SetLayer(void *viewPtr, void *layerPtr)
{
    NSView* view = (NSView*)viewPtr;
    CALayer* layer = (CALayer*)layerPtr;
    [view setLayer:layer];
}

void View_SetAlphaValue(void *viewPtr, double alpha)
{
    NSView* view = (NSView*)viewPtr;
    [view setAlphaValue:alpha];
}

double View_GetAlphaValue(void *viewPtr)
{
    NSView* view = (NSView*)viewPtr;
    return [view alphaValue];
}

// Responder chain and events
int View_AcceptsFirstResponder(void *viewPtr)
{
    NSView* view = (NSView*)viewPtr;
    return [view acceptsFirstResponder] ? 1 : 0;
}

int View_BecomeFirstResponder(void *viewPtr)
{
    NSView* view = (NSView*)viewPtr;
    return [view becomeFirstResponder] ? 1 : 0;
}

int View_ResignFirstResponder(void *viewPtr)
{
    NSView* view = (NSView*)viewPtr;
    return [view resignFirstResponder] ? 1 : 0;
}

void* View_NextResponder(void *viewPtr)
{
    NSView* view = (NSView*)viewPtr;
    return [view nextResponder];
}

// Mouse events
void View_SetAcceptsMouseEvents(void *viewPtr, int accepts)
{
    // This is typically handled by subclasses or gesture recognizers
    // For basic NSView, we'll implement as no-op
}

int View_AcceptsMouseEvents(void *viewPtr)
{
    // Default implementation - most views accept mouse events
    return 1;
}

// Focus ring and highlighting
void View_SetKeyboardFocusRingNeedsDisplayInRect(void *viewPtr, int x, int y, int width, int height)
{
    NSView* view = (NSView*)viewPtr;
    NSRect rect = NSMakeRect(x, y, width, height);
    [view setKeyboardFocusRingNeedsDisplayInRect:rect];
}

void View_SetFocusRingType(void *viewPtr, int type)
{
    NSView* view = (NSView*)viewPtr;
    [view setFocusRingType:type];
}

int View_GetFocusRingType(void *viewPtr)
{
    NSView* view = (NSView*)viewPtr;
    return (int)[view focusRingType];
}

// Scrolling support
int View_IsFlipped(void *viewPtr)
{
    NSView* view = (NSView*)viewPtr;
    return [view isFlipped] ? 1 : 0;
}

void View_SetFlipped(void *viewPtr, int flipped)
{
    // NSView doesn't have a direct setFlipped method
    // This is typically overridden in subclasses
}

void View_ScrollRectToVisible(void *viewPtr, int x, int y, int width, int height)
{
    NSView* view = (NSView*)viewPtr;
    NSRect rect = NSMakeRect(x, y, width, height);
    [view scrollRectToVisible:rect];
}

void View_ScrollPoint(void *viewPtr, int x, int y)
{
    NSView* view = (NSView*)viewPtr;
    NSPoint point = NSMakePoint(x, y);
    [view scrollPoint:point];
}

// Window relationship
void* View_GetWindow(void *viewPtr)
{
    NSView* view = (NSView*)viewPtr;
    return view.window;
}

int View_IsInLiveResize(void *viewPtr)
{
    NSView* view = (NSView*)viewPtr;
    return [view inLiveResize] ? 1 : 0;
}

// Layout and sizing
void View_SetFrameCenterRotation(void *viewPtr, double angle)
{
    NSView* view = (NSView*)viewPtr;
    [view setFrameCenterRotation:angle];
}

double View_GetFrameCenterRotation(void *viewPtr)
{
    NSView* view = (NSView*)viewPtr;
    return [view frameCenterRotation];
}

// Tag and identification
static const void* ViewTagAssociationKey = &ViewTagAssociationKey;

void View_SetTag(void *viewPtr, int tag)
{
    NSView* view = (NSView*)viewPtr;
    NSNumber* tagNumber = [NSNumber numberWithInt:tag];
    objc_setAssociatedObject(view, ViewTagAssociationKey, tagNumber, OBJC_ASSOCIATION_RETAIN_NONATOMIC);
}

int View_GetTag(void *viewPtr)
{
    NSView* view = (NSView*)viewPtr;
    NSNumber* tagNumber = objc_getAssociatedObject(view, ViewTagAssociationKey);
    return tagNumber ? [tagNumber intValue] : 0;
}

void* View_ViewWithTag(void *viewPtr, int tag)
{
    NSView* view = (NSView*)viewPtr;
    
    // Check if this view has the matching tag
    NSNumber* viewTagNumber = objc_getAssociatedObject(view, ViewTagAssociationKey);
    if (viewTagNumber && [viewTagNumber intValue] == tag) {
        return view;
    }
    
    // Recursively search subviews
    for (NSView* subview in view.subviews) {
        void* result = View_ViewWithTag(subview, tag);
        if (result != nil) {
            return result;
        }
    }
    
    return nil;
}

// Tooltip support
void View_SetToolTip(void *viewPtr, const char* tooltip)
{
    NSView* view = (NSView*)viewPtr;
    NSString* tooltipStr = [NSString stringWithUTF8String:tooltip];
    [view setToolTip:tooltipStr];
}

char* View_GetToolTip(void *viewPtr)
{
    NSView* view = (NSView*)viewPtr;
    NSString* tooltip = [view toolTip];
    if (tooltip == nil) {
        return NULL;
    }
    const char* cString = [tooltip UTF8String];
    char* result = malloc(strlen(cString) + 1);
    strcpy(result, cString);
    return result;
}

// Appearance and material (macOS 10.14+)
void View_SetAppearance(void *viewPtr, int appearanceType)
{
    NSView* view = (NSView*)viewPtr;
    NSAppearance* appearance = nil;
    
    switch (appearanceType) {
        case 0: // Light
            if (@available(macOS 10.14, *)) {
                appearance = [NSAppearance appearanceNamed:NSAppearanceNameAqua];
            }
            break;
        case 1: // Dark
            if (@available(macOS 10.14, *)) {
                appearance = [NSAppearance appearanceNamed:NSAppearanceNameDarkAqua];
            }
            break;
        default:
            appearance = nil;
            break;
    }
    
    if (@available(macOS 10.14, *)) {
        [view setAppearance:appearance];
    }
}

int View_GetEffectiveAppearance(void *viewPtr)
{
    NSView* view = (NSView*)viewPtr;
    if (@available(macOS 10.14, *)) {
        NSAppearance* appearance = [view effectiveAppearance];
        NSString* name = appearance.name;
        
        if ([name isEqualToString:NSAppearanceNameDarkAqua]) {
            return 1; // Dark
        } else {
            return 0; // Light
        }
    }
    return 0; // Default to light
}

// Animation support
void View_SetAnimations(void *viewPtr, void *animationsDict)
{
    NSView* view = (NSView*)viewPtr;
    NSDictionary* animations = (NSDictionary*)animationsDict;
    [view setAnimations:animations];
}

void* View_GetAnimations(void *viewPtr)
{
    NSView* view = (NSView*)viewPtr;
    return [view animations];
}

// Drawing context
void* View_BitmapImageRepForCachingDisplayInRect(void *viewPtr, int x, int y, int width, int height)
{
    NSView* view = (NSView*)viewPtr;
    NSRect rect = NSMakeRect(x, y, width, height);
    return [view bitmapImageRepForCachingDisplayInRect:rect];
}

void View_CacheDisplayInRect(void *viewPtr, int x, int y, int width, int height, void *bitmapImageRep)
{
    NSView* view = (NSView*)viewPtr;
    NSRect rect = NSMakeRect(x, y, width, height);
    NSBitmapImageRep* imageRep = (NSBitmapImageRep*)bitmapImageRep;
    [view cacheDisplayInRect:rect toBitmapImageRep:imageRep];
}

// Invalidation
void View_InvalidateIntrinsicContentSize(void *viewPtr)
{
    NSView* view = (NSView*)viewPtr;
    [view invalidateIntrinsicContentSize];
}

void View_GetIntrinsicContentSize(void *viewPtr, int *width, int *height)
{
    NSView* view = (NSView*)viewPtr;
    NSSize size = [view intrinsicContentSize];
    *width = (int)size.width;
    *height = (int)size.height;
}
