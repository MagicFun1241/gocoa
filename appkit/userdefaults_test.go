package appkit

import (
	"testing"
)

func TestNewUserDefaults(t *testing.T) {
	defaults := NewUserDefaults()
	if defaults.userDefaultsPtr == nil {
		t.Fatalf("pointer to C defaults is nil!")
	}
}

func TestStringValue(t *testing.T) {
	defaults := NewUserDefaults()

	value := "FooBar!"
	defaults.SetString("TestKey", value)

	readValue := defaults.GetString("TestKey")
	if readValue != value {
		t.Fatalf("read value should be %s, but is %s", value, readValue)
	}
}

func TestKeyRemove(t *testing.T) {
	defaults := NewUserDefaults()

	value := "FooBar!"
	defaults.SetString("TestKey", value)

	defaults.Remove("TestKey")

	readValue := defaults.GetString("TestKey")
	if readValue != "" {
		t.Fatalf("read value should be empty string but got %s", readValue)
	}
}
