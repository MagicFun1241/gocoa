// Package gocoa - Go bindings for the Cocoa framework.
// It is currently in early stage development and not very usable yet.
package gocoa

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Foundation

#import <Foundation/Foundation.h>
#import <dispatch/dispatch.h>

extern void goCallbackTrigger(int callbackID);

static inline void gocoa_main_thread_exec(int callbackID) {
    dispatch_queue_t queue = dispatch_get_main_queue();

    dispatch_async(queue, ^{
        goCallbackTrigger(callbackID);
    });
}
*/
import "C"
import (
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	callbackMutex sync.RWMutex
	callbacks     = make(map[int]func())
	nextID        int64
)

//export goCallbackTrigger
func goCallbackTrigger(callbackID C.int) {
	id := int(callbackID)

	callbackMutex.RLock()
	fn, exists := callbacks[id]
	callbackMutex.RUnlock()

	if exists {
		fn()

		// Clean up the callback after execution
		callbackMutex.Lock()
		delete(callbacks, id)
		callbackMutex.Unlock()
	}
}

// RunOnMainThread executes the given function on the main thread.
func RunOnMainThread(fn func()) {
	// Generate unique callback ID
	id := int(atomic.AddInt64(&nextID, 1))

	// Store the callback in our registry
	callbackMutex.Lock()
	callbacks[id] = fn
	callbackMutex.Unlock()

	// Execute on main thread via GCD
	C.gocoa_main_thread_exec(C.int(id))
}

// isMainThread checks if the current goroutine is running on the main thread.
// Note: This is a simplified check and may not be 100% accurate in all scenarios.
func isMainThread() bool {
	// This is a heuristic - the main thread typically has ID 1
	// In practice, you might want to use a more robust method
	return runtime.NumGoroutine() == 1 || C.dispatch_queue_get_label(C.dispatch_get_main_queue()) != nil
}
