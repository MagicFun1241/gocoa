package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework WebKit
// #import "webview.h"
import "C"
import "unsafe"

// WebView represents a WebView control that can trigger actions.
type WebView struct {
	webviewPtr C.WebViewPtr
	callback   func()
}

var webviews []*WebView

func NewWebView(x int, y int, width int, height int) *WebView {
	wv := NewWebViewWithUrl(x, y, width, height, "")
	return wv
}

func NewWebViewWithUrl(x int, y int, width int, height int, url string) *WebView {
	webviewID := len(searchfields)

	cUrl := C.CString(url)
	defer C.free(unsafe.Pointer(cUrl))

	webviewPtr := C.WebView_New(C.int(webviewID), C.int(x), C.int(y), C.int(width), C.int(height), cUrl)

	wv := &WebView{
		webviewPtr: webviewPtr,
	}

	webviews = append(webviews, wv)
	return wv
}

func (webview *WebView) SetUrl(url string) {
	cUrl := C.CString(url)
	defer C.free(unsafe.Pointer(cUrl))

	C.WebView_SetUrl(webview.webviewPtr, cUrl)
}

func (webview *WebView) SetFrame(x, y, w, h int) {
	C.WebView_SetFrame(webview.webviewPtr, C.int(x), C.int(y), C.int(w), C.int(h))
}

func (webview *WebView) SetAllowsMagnification(allows bool) {
	if allows {
		C.WebView_SetAllowsMagnification(webview.webviewPtr, C.int(1))
	} else {
		C.WebView_SetAllowsMagnification(webview.webviewPtr, C.int(0))
	}
}

func (webview *WebView) SetAllowsBackForwardNavigationGestures(allows bool) {
	if allows {
		C.WebView_SetAllowsBackForwardNavigationGestures(webview.webviewPtr, C.int(1))
	} else {
		C.WebView_SetAllowsBackForwardNavigationGestures(webview.webviewPtr, C.int(0))
	}
}

func (webview *WebView) SetCustomUserAgent(userAgent string) {
	cUserAgent := C.CString(userAgent)
	defer C.free(unsafe.Pointer(cUserAgent))

	C.WebView_SetCustomUserAgent(webview.webviewPtr, cUserAgent)
}

func (webview *WebView) Remove() {
	C.WebView_Remove(webview.webviewPtr)
}
