package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #import "imageview.h"
import "C"
import (
	"github.com/magicfun1241/gocoa/tools"
	"unsafe"
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
