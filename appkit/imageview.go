package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #import "imageview.h"
import "C"
import (
	"unsafe"

	"github.com/magicfun1241/gocoa/tools"
)

// Represents an ImageView control that can display images.
type ImageView struct {
	imageViewPtr C.ImageViewPtr
	callback     func()
}

type FrameStyle int32
type ImageAlignment int32
type ImageScaling int32

const (
	FrameStyleNone      FrameStyle = 0
	FrameStylePhoto     FrameStyle = 1
	FrameStyleGrayBezel FrameStyle = 2
	FrameStyleGroove    FrameStyle = 3
	FrameStyleButton    FrameStyle = 4
)

const (
	ImageAlignCenter      ImageAlignment = 0
	ImageAlignTop         ImageAlignment = 1
	ImageAlignTopLeft     ImageAlignment = 2
	ImageAlignTopRight    ImageAlignment = 3
	ImageAlignLeft        ImageAlignment = 4
	ImageAlignBottom      ImageAlignment = 5
	ImageAlignBottomLeft  ImageAlignment = 6
	ImageAlignBottomRight ImageAlignment = 7
	ImageAlignRight       ImageAlignment = 8
)

const (
	ImageScalingScaleProportionallyDown     ImageScaling = 0
	ImageScalingScaleAxesIndependently      ImageScaling = 1
	ImageScalingScaleNone                   ImageScaling = 2
	ImageScalingScaleProportionallyUpOrDown ImageScaling = 3
)

var imageViews []*ImageView

func NewImageView(x int, y int, width int, height int) *ImageView {
	imageViewID := len(imageViews)

	cUrl := C.CString("")
	defer C.free(unsafe.Pointer(cUrl))

	imageViewPtr := C.ImageView_New(C.int(imageViewID), C.int(x), C.int(y), C.int(width), C.int(height), cUrl)

	img := &ImageView{
		imageViewPtr: imageViewPtr,
	}
	imageViews = append(imageViews, img)
	return img
}

func NewImageViewWithUrl(x int, y int, width int, height int, url string) *ImageView {
	imageViewID := len(imageViews)

	cUrl := C.CString(url)
	defer C.free(unsafe.Pointer(cUrl))

	imageViewPtr := C.ImageView_New(C.int(imageViewID), C.int(x), C.int(y), C.int(width), C.int(height), cUrl)

	img := &ImageView{
		imageViewPtr: imageViewPtr,
	}
	imageViews = append(imageViews, img)
	return img
}

func (imageView *ImageView) SetEditable(editable bool) {
	if editable {
		C.ImageView_SetEditable(imageView.imageViewPtr, 1)
	} else {
		C.ImageView_SetEditable(imageView.imageViewPtr, 0)
	}
}

func (imageView *ImageView) SetImageFrameStyle(frameStyle FrameStyle) {
	C.ImageView_SetFrameStyle(imageView.imageViewPtr, C.int(frameStyle))
}

func (imageView *ImageView) SetImageAlignment(imageAlignment ImageAlignment) {
	C.ImageView_SetImageAlignment(imageView.imageViewPtr, C.int(imageAlignment))
}

func (imageView *ImageView) SetImageScaling(imageScaling ImageScaling) {
	C.ImageView_SetImageScaling(imageView.imageViewPtr, C.int(imageScaling))
}

func (imageView *ImageView) SetImageUrl(url string) {
	cUrl := C.CString(url)
	defer C.free(unsafe.Pointer(cUrl))

	C.ImageView_SetImageUrl(imageView.imageViewPtr, cUrl)
}

func (imageView *ImageView) SetImageBytes(imageData []byte) {
	cChar := (*C.uchar)(unsafe.Pointer(&imageData[0]))

	C.ImageView_SetImageBytes(imageView.imageViewPtr, cChar, C.int(len(imageData)))
}

func (imageView *ImageView) SetAnimates(animates bool) {
	if animates {
		C.ImageView_SetAnimates(imageView.imageViewPtr, 1)
	} else {
		C.ImageView_SetAnimates(imageView.imageViewPtr, 0)
	}
}

func (imageView *ImageView) SetContentTintColor(hexRGBA string) {
	var r, g, b, a = tools.ParseHexRGBA(hexRGBA)
	C.ImageView_SetContentTintColor(imageView.imageViewPtr, C.int(r), C.int(g), C.int(b), C.int(a))
}

// SetImageName sets the image using a named system image or app bundle image
func (imageView *ImageView) SetImageName(imageName string) {
	cImageName := C.CString(imageName)
	defer C.free(unsafe.Pointer(cImageName))
	C.ImageView_SetImageName(imageView.imageViewPtr, cImageName)
}

// ImageName returns the name of the current image
func (imageView *ImageView) ImageName() string {
	return C.GoString(C.ImageView_ImageName(imageView.imageViewPtr))
}

// Editable returns whether the image view is editable
func (imageView *ImageView) Editable() bool {
	return C.ImageView_Editable(imageView.imageViewPtr) == 1
}

// Animates returns whether the image view animates
func (imageView *ImageView) Animates() bool {
	return C.ImageView_Animates(imageView.imageViewPtr) == 1
}

// ImageFrameStyle returns the current frame style
func (imageView *ImageView) ImageFrameStyle() FrameStyle {
	return FrameStyle(C.ImageView_FrameStyle(imageView.imageViewPtr))
}

// ImageAlignment returns the current image alignment
func (imageView *ImageView) ImageAlignment() ImageAlignment {
	return ImageAlignment(C.ImageView_ImageAlignment(imageView.imageViewPtr))
}

// ImageScaling returns the current image scaling
func (imageView *ImageView) ImageScaling() ImageScaling {
	return ImageScaling(C.ImageView_ImageScaling(imageView.imageViewPtr))
}

// SetAllowsCutCopyPaste sets whether the image view allows cut, copy, and paste operations
func (imageView *ImageView) SetAllowsCutCopyPaste(allowsCutCopyPaste bool) {
	if allowsCutCopyPaste {
		C.ImageView_SetAllowsCutCopyPaste(imageView.imageViewPtr, C.int(1))
	} else {
		C.ImageView_SetAllowsCutCopyPaste(imageView.imageViewPtr, C.int(0))
	}
}

// AllowsCutCopyPaste returns whether the image view allows cut, copy, and paste operations
func (imageView *ImageView) AllowsCutCopyPaste() bool {
	return C.ImageView_AllowsCutCopyPaste(imageView.imageViewPtr) == 1
}

// Symbol weight constants for macOS 11.0+
const (
	SymbolWeightUltraLight = 1
	SymbolWeightThin       = 2
	SymbolWeightLight      = 3
	SymbolWeightRegular    = 4
	SymbolWeightMedium     = 5
	SymbolWeightSemibold   = 6
	SymbolWeightBold       = 7
	SymbolWeightHeavy      = 8
	SymbolWeightBlack      = 9
)

// SetImageSymbolConfiguration sets the symbol configuration for SF Symbols (macOS 11.0+)
func (imageView *ImageView) SetImageSymbolConfiguration(configName string, pointSize float64, weight int) {
	cConfigName := C.CString(configName)
	defer C.free(unsafe.Pointer(cConfigName))
	C.ImageView_SetImageSymbolConfiguration(imageView.imageViewPtr, cConfigName, C.double(pointSize), C.int(weight))
}

// SetBackgroundColor sets the background color of the image view
func (imageView *ImageView) SetBackgroundColor(hexRGBA string) {
	var r, g, b, a = tools.ParseHexRGBA(hexRGBA)
	C.ImageView_SetBackgroundColor(imageView.imageViewPtr, C.int(r), C.int(g), C.int(b), C.int(a))
}

// SetHidden sets whether the image view is hidden
func (imageView *ImageView) SetHidden(hidden bool) {
	if hidden {
		C.ImageView_SetHidden(imageView.imageViewPtr, C.int(1))
	} else {
		C.ImageView_SetHidden(imageView.imageViewPtr, C.int(0))
	}
}

// Hidden returns whether the image view is hidden
func (imageView *ImageView) Hidden() bool {
	return C.ImageView_Hidden(imageView.imageViewPtr) == 1
}

// SetAlphaValue sets the alpha (transparency) value of the image view (0.0 = transparent, 1.0 = opaque)
func (imageView *ImageView) SetAlphaValue(alpha float64) {
	C.ImageView_SetAlphaValue(imageView.imageViewPtr, C.double(alpha))
}

// AlphaValue returns the alpha (transparency) value of the image view
func (imageView *ImageView) AlphaValue() float64 {
	return float64(C.ImageView_AlphaValue(imageView.imageViewPtr))
}
