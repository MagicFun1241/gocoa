# gocoa

[![GoDoc](https://godoc.org/github.com/magicfun1241/gocoa?status.svg)](https://godoc.org/github.com/magicfun1241/gocoa)
[![Go Report Card](https://goreportcard.com/badge/github.com/magicfun1241/gocoa)](https://goreportcard.com/report/github.com/magicfun1241/gocoa)

Go bindings for the Cocoa framework to build macOS applications.

[//]: # (<img src="resources/images/helloworld-screenshot.gif" width="480" />)

## How to use

[//]: # (The following is a basic [Hello World]&#40;examples/helloworld&#41; application.)

```go
package main

import (
	"fmt"

	"github.com/magicfun1241/gocoa/appkit"
)

func main() {
	gocoa.InitApplication()
	gocoa.OnApplicationDidFinishLaunching(func() {
		fmt.Println("App running!")
	})
	wnd := gocoa.NewWindow("Hello World!", 150, 150, 300, 200)

	// Change me button
	currentTitle, nextTitle := "Change me!", "Change me again!"
	changeButton := gocoa.NewButton(75, 125, 150, 25)
	changeButton.SetTitle(currentTitle)
	changeButton.OnClick(func() {
		changeButton.SetTitle(nextTitle)
		currentTitle, nextTitle = nextTitle, currentTitle
	})
	wnd.AddButton(changeButton)

	// Quit button
	quitButton := gocoa.NewButton(75, 50, 150, 25)
	quitButton.SetTitle("Quit")
	quitButton.OnClick(func() { gocoa.TerminateApplication() })
	wnd.AddButton(quitButton)

	wnd.MakeKeyAndOrderFront()
	gocoa.RunApplication()
}
```

## Status of this project

This package is very, very early and incomplete! It is mostly just an experiment and is not really
useful yet.

### Partial implementation

Not all methods are implemented

#### Components
- [x] TextView
- [x] TextField
- [x] SearchField
- [x] ImageView
- [x] Button
- [x] DatePicker
- [x] ComboBox
- [x] ProgressIndicator
- [x] Slider
- [x] WebView

#### Additional APIs
- [x] UserDefaults


## Contributors

This is fork of [gocoa](https://github.com/mojbro/gocoa) with already merged pull request and small refactor

### Notice from original repository: 
Big thanks to [Philipp Haussleiter](https://github.com/phaus) who has contributed a great deal to this project.
