package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #include "application.h"
import "C"

var appDidFinishLaunchingFunc func()

// InitApplication initializes the global application instance. Call this before using
// the rest of the gocoa package.
func InitApplication() {
	C.InitSharedApplication()
}

// RunApplication launches the main Cocoa runloop.
func RunApplication() {
	C.RunApplication()
}

// OnApplicationDidFinishLaunching - will be triggered after Application Launch is finished
func OnApplicationDidFinishLaunching(fn func()) {
	appDidFinishLaunchingFunc = fn
}

// TerminateApplication - will be triggered, when the Application terminates
func TerminateApplication() {
	C.TerminateApplication()
}

//export callOnApplicationDidFinishLaunchingHandler
func callOnApplicationDidFinishLaunchingHandler() {
	if appDidFinishLaunchingFunc != nil {
		appDidFinishLaunchingFunc()
	}
}
