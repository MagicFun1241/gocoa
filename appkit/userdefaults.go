package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Foundation
// #import "userdefaults.h"
import "C"
import "unsafe"

// UserDefaults - user default settings
type UserDefaults struct {
	userDefaultsPtr C.UserDefaultsPtr
}

// NewUserDefaults - creates a new UserDefaults object
func NewUserDefaults() *UserDefaults {
	defaultsPtr := C.UserDefaults_New()
	defaults := &UserDefaults{
		userDefaultsPtr: defaultsPtr,
	}
	return defaults
}

// GetString - reads a string from userDefaults.
func (defaults *UserDefaults) GetString(key string) string {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))

	value := C.UserDefaults_GetString(defaults.userDefaultsPtr, cKey)
	return C.GoString(value)
}

// SetString - writes a string to userDefaults.
func (defaults *UserDefaults) SetString(key string, value string) {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))

	cValue := C.CString(value)
	defer C.free(unsafe.Pointer(cValue))

	C.UserDefaults_SetString(defaults.userDefaultsPtr, cKey, cValue)
}

// Remove - deletes key from userDefaults.
func (defaults *UserDefaults) Remove(key string) {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))

	C.UserDefaults_Remove(defaults.userDefaultsPtr, cKey)
}
