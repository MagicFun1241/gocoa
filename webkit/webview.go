package webkit

import "github.com/magicfun1241/gocoa/appkit"

func NewWebView(x int, y int, width int, height int) *appkit.WebView {
	return appkit.NewWebView(x, y, width, height)
}

func NewWebViewWithUrl(x int, y int, width int, height int, url string) *appkit.WebView {
	return appkit.NewWebViewWithUrl(x, y, width, height, url)
}
