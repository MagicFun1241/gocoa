#import "imageview.h"
#include "_cgo_export.h"

ButtonPtr ImageView_New(int goButtonID, int x, int y, int w, int h, const char* url) {
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
