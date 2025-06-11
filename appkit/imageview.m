#import "imageview.h"
#include "_cgo_export.h"

ImageViewPtr ImageView_New(int goImageViewID, int x, int y, int w, int h, const char* url) {
    id nsImageView = [[[NSImageView alloc] initWithFrame:NSMakeRect(x, y, w, h)] autorelease];

    if (url && strlen(url) > 0) {
        NSImage *theImage = [[NSImage alloc] initWithContentsOfURL:[NSURL URLWithString:[NSString stringWithUTF8String:url]]];
        [nsImageView setImage:theImage];
    }

    return (ImageViewPtr)nsImageView;
}

void ImageView_SetFrameStyle(ImageViewPtr imageViewPtr, int frameStyle) {
    NSImageView* nsImageView = (NSImageView*)imageViewPtr;
    [nsImageView setImageFrameStyle:frameStyle];
}

void ImageView_SetImageAlignment(ImageViewPtr imageViewPtr, int imageAlignment) {
    NSImageView* nsImageView = (NSImageView*)imageViewPtr;
    [nsImageView setImageAlignment:imageAlignment];
}

void ImageView_SetImageScaling(ImageViewPtr imageViewPtr, int imageScaling) {
    NSImageView* nsImageView = (NSImageView*)imageViewPtr;
    [nsImageView setImageScaling:imageScaling];
}

void ImageView_SetAnimates(ImageViewPtr imageViewPtr, int animates) {
    NSImageView* nsImageView = (NSImageView*)imageViewPtr;
    [nsImageView setAnimates:animates];
}

void ImageView_SetImageUrl(ImageViewPtr imageViewPtr, const char* url) {
    NSImage *theImage = [[NSImage alloc] initWithContentsOfURL:[NSURL URLWithString:[NSString stringWithUTF8String:url]]];

    NSImageView* nsImageView = (NSImageView*)imageViewPtr;
    [nsImageView setImage:theImage];
}

void ImageView_SetImageBytes(ImageViewPtr imageViewPtr, const unsigned char* imageData, int length) {
    NSData *data = [NSData dataWithBytes:imageData length:length];
    NSImage *theImage = [[NSImage alloc] initWithData:data];

    NSImageView* nsImageView = (NSImageView*)imageViewPtr;
    [nsImageView setImage:theImage];
}

void ImageView_SetContentTintColor(ImageViewPtr imageViewPtr, int r, int g, int b, int a) {
    NSImageView* nsImageView = (NSImageView*)imageViewPtr;
    [nsImageView setContentTintColor:[NSColor colorWithCalibratedRed:r/255.f green:g/255.f blue:b/255.f alpha:a/255.f]];
}

void ImageView_SetEditable(ImageViewPtr imageViewPtr, int editable) {
    NSImageView* nsImageView = (NSImageView*)imageViewPtr;
    [nsImageView setEditable:editable];
}

const int ImageView_Editable(ImageViewPtr imageViewPtr) {
    NSImageView* nsImageView = (NSImageView*)imageViewPtr;
    return [nsImageView isEditable] ? 1 : 0;
}

void ImageView_SetImageName(ImageViewPtr imageViewPtr, const char* imageName) {
    NSImageView* nsImageView = (NSImageView*)imageViewPtr;
    NSString* imageNameStr = [NSString stringWithUTF8String:imageName];
    NSImage* image = [NSImage imageNamed:imageNameStr];
    [nsImageView setImage:image];
}

const char* ImageView_ImageName(ImageViewPtr imageViewPtr) {
    NSImageView* nsImageView = (NSImageView*)imageViewPtr;
    NSImage* image = [nsImageView image];
    if (image && [image name]) {
        return [[image name] cStringUsingEncoding:NSISOLatin1StringEncoding];
    }
    return "";
}

const int ImageView_Animates(ImageViewPtr imageViewPtr) {
    NSImageView* nsImageView = (NSImageView*)imageViewPtr;
    return [nsImageView animates] ? 1 : 0;
}

const int ImageView_FrameStyle(ImageViewPtr imageViewPtr) {
    NSImageView* nsImageView = (NSImageView*)imageViewPtr;
    return (int)[nsImageView imageFrameStyle];
}

const int ImageView_ImageAlignment(ImageViewPtr imageViewPtr) {
    NSImageView* nsImageView = (NSImageView*)imageViewPtr;
    return (int)[nsImageView imageAlignment];
}

const int ImageView_ImageScaling(ImageViewPtr imageViewPtr) {
    NSImageView* nsImageView = (NSImageView*)imageViewPtr;
    return (int)[nsImageView imageScaling];
}

void ImageView_SetAllowsCutCopyPaste(ImageViewPtr imageViewPtr, int allowsCutCopyPaste) {
    NSImageView* nsImageView = (NSImageView*)imageViewPtr;
    [nsImageView setAllowsCutCopyPaste:allowsCutCopyPaste ? YES : NO];
}

const int ImageView_AllowsCutCopyPaste(ImageViewPtr imageViewPtr) {
    NSImageView* nsImageView = (NSImageView*)imageViewPtr;
    return [nsImageView allowsCutCopyPaste] ? 1 : 0;
}

void ImageView_SetImageSymbolConfiguration(ImageViewPtr imageViewPtr, const char* configName, double pointSize, int weight) {
    if (@available(macOS 11.0, *)) {
        NSImageView* nsImageView = (NSImageView*)imageViewPtr;
        NSImageSymbolConfiguration* config = [NSImageSymbolConfiguration configurationWithPointSize:pointSize weight:weight];
        [nsImageView setSymbolConfiguration:config];
    }
}

void ImageView_SetBackgroundColor(ImageViewPtr imageViewPtr, int r, int g, int b, int a) {
    NSImageView* nsImageView = (NSImageView*)imageViewPtr;
    nsImageView.wantsLayer = YES;
    nsImageView.layer.backgroundColor = [[NSColor colorWithCalibratedRed:r/255.0f green:g/255.0f blue:b/255.0f alpha:a/255.0f] CGColor];
}

void ImageView_SetHidden(ImageViewPtr imageViewPtr, int hidden) {
    NSImageView* nsImageView = (NSImageView*)imageViewPtr;
    [nsImageView setHidden:hidden ? YES : NO];
}

const int ImageView_Hidden(ImageViewPtr imageViewPtr) {
    NSImageView* nsImageView = (NSImageView*)imageViewPtr;
    return [nsImageView isHidden] ? 1 : 0;
}

void ImageView_SetAlphaValue(ImageViewPtr imageViewPtr, double alpha) {
    NSImageView* nsImageView = (NSImageView*)imageViewPtr;
    [nsImageView setAlphaValue:alpha];
}

double ImageView_AlphaValue(ImageViewPtr imageViewPtr) {
    NSImageView* nsImageView = (NSImageView*)imageViewPtr;
    return [nsImageView alphaValue];
}
