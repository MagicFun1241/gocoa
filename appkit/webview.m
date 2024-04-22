#import "webview.h"
#include <CoreGraphics/CGGeometry.h>
#include "_cgo_export.h"

@implementation WebViewHandler
@end

WebViewPtr WebView_New(int goWebViewId, int x, int y, int w, int h, const char* url) {
    WKWebView *webview = [[WKWebView alloc] initWithFrame:NSMakeRect(x, y, w, h)];

    if (url && strlen(url) > 0) {
        NSURL *nsUrl = [NSURL URLWithString:[NSString stringWithUTF8String:url]];
        NSURLRequest *request = [NSURLRequest requestWithURL:nsUrl];
        [webview loadRequest:request];
    }

    return (WebViewPtr)webview;
}

void WebView_SetUrl(WebViewPtr webviewPtr, const char* url) {
    WKWebView* wv = (WKWebView*)webviewPtr;

    NSURL *nsUrl = [NSURL URLWithString:[NSString stringWithUTF8String:url]];
    NSURLRequest *request = [NSURLRequest requestWithURL:nsUrl];
    [wv loadRequest:request];
}

void WebView_SetAllowsBackForwardNavigationGestures(WebViewPtr webviewPtr, int enabled) {
    WKWebView* wv = (WKWebView*)webviewPtr;

    [wv setAllowsBackForwardNavigationGestures:enabled];
}

void WebView_SetAllowsMagnification(WebViewPtr webviewPtr, int enabled) {
    WKWebView* wv = (WKWebView*)webviewPtr;

    [wv setAllowsMagnification:enabled];
}

void WebView_SetFrame(WebViewPtr webviewPtr, int x, int y, int w, int h) {
    WKWebView* wv = (WKWebView*)webviewPtr;

    CGRect newFrame = CGRectMake(x, y, w, h);
    [wv setFrame:newFrame];
}

void WebView_SetCustomUserAgent(WebViewPtr webviewPtr, const char* userAgent) {
    WKWebView* wv = (WKWebView*)webviewPtr;

    [wv setCustomUserAgent:[NSString stringWithUTF8String:userAgent]];
}

void WebView_Remove(WebViewPtr webviewPtr) {
    WKWebView* wv = (WKWebView*)webviewPtr;
    [wv removeFromSuperview];
}
