#import <Cocoa/Cocoa.h>

#import <Foundation/Foundation.h>
#import <WebKit/WebKit.h>

@interface WebViewHandler : NSObject

@property (assign) int goWebViewId;

@end

typedef void* WebViewPtr;

WebViewPtr WebView_New(int goWebViewId, int x, int y, int w, int h, const char *url);
void WebView_SetUrl(WebViewPtr webviewPtr, const char* url);
void WebView_SetAllowsBackForwardNavigationGestures(WebViewPtr webviewPtr, int enabled);
void WebView_SetAllowsMagnification(WebViewPtr webviewPtr, int enabled);
void WebView_SetCustomUserAgent(WebViewPtr webviewPtr, const char* userAgent);
void WebView_SetFrame(WebViewPtr webviewPtr, int x, int y, int w, int h);

void WebView_Remove(WebViewPtr webviewPtr);