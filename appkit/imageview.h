#import <Cocoa/Cocoa.h>

// typedef void (*callback)(void);

typedef void* ImageViewPtr;

ImageViewPtr ImageView_New(int goImageViewID, int x, int y, int w, int h, const char* url);

void ImageView_SetImageUrl(ImageViewPtr imageViewPtr, const char* url);
void ImageView_SetImageBytes(ImageViewPtr imageViewPtr, const unsigned char* imageData, int length);
void ImageView_SetImageName(ImageViewPtr imageViewPtr, const char* imageName);

void ImageView_SetAnimates(ImageViewPtr imageViewPtr, int animates);
const int ImageView_Animates(ImageViewPtr imageViewPtr);

void ImageView_SetContentTintColor(ImageViewPtr imageViewPtr, int r, int g, int b, int a);
void ImageView_SetEditable(ImageViewPtr imageViewPtr, int editable);
const int ImageView_Editable(ImageViewPtr imageViewPtr);

void ImageView_SetFrameStyle(ImageViewPtr imageViewPtr, int frameStyle);
const int ImageView_FrameStyle(ImageViewPtr imageViewPtr);

void ImageView_SetImageAlignment(ImageViewPtr imageViewPtr, int imageAlignment);
const int ImageView_ImageAlignment(ImageViewPtr imageViewPtr);

void ImageView_SetImageScaling(ImageViewPtr imageViewPtr, int imageScaling);
const int ImageView_ImageScaling(ImageViewPtr imageViewPtr);

void ImageView_SetAllowsCutCopyPaste(ImageViewPtr imageViewPtr, int allowsCutCopyPaste);
const int ImageView_AllowsCutCopyPaste(ImageViewPtr imageViewPtr);

void ImageView_SetImageSymbolConfiguration(ImageViewPtr imageViewPtr, const char* configName, double pointSize, int weight);

const char* ImageView_ImageName(ImageViewPtr imageViewPtr);

void ImageView_SetBackgroundColor(ImageViewPtr imageViewPtr, int r, int g, int b, int a);

void ImageView_SetHidden(ImageViewPtr imageViewPtr, int hidden);
const int ImageView_Hidden(ImageViewPtr imageViewPtr);

void ImageView_SetAlphaValue(ImageViewPtr imageViewPtr, double alpha);
double ImageView_AlphaValue(ImageViewPtr imageViewPtr);
