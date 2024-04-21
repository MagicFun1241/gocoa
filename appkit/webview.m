#import "webview.h"
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

void WebView_Remove(WebViewPtr webviewPtr) {
    WKWebView* wv = (WKWebView*)webviewPtr;
    [wv removeFromSuperview];
}