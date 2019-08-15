package main

import (
	"fmt"
	"time"

	"github.com/mojbro/gocoa"
)

var wnd *gocoa.Window
var indicator *gocoa.ProgressIndicator

const maxValue = 100.00

func main() {
	gocoa.InitApplication()
	gocoa.OnApplicationDidFinishLaunching(func() {
		fmt.Println("App running!")
	})
	wnd = gocoa.NewWindow("Hello World!", 150, 150, 300, 200)

	// TextField
	mainTextField := gocoa.NewTextField(75, 145, 150, 25)
	wnd.AddTextField(mainTextField)

	// Change me button
	currentTitle, nextTitle := "Change me!", "Change me again!"
	changeButton := gocoa.NewButton(75, 120, 150, 25)
	changeButton.SetTitle(currentTitle)
	changeButton.OnClick(func() {
		changeButton.SetTitle(nextTitle)
		currentTitle, nextTitle = nextTitle, currentTitle
	})
	wnd.AddButton(changeButton)

	// Download button
	loadButton := gocoa.NewButton(75, 75, 150, 25)
	loadButton.SetTitle("download")

	loadButton.OnClick(func() {
		indicator.Show()
		go func() {
			increase()
		}()
	})
	wnd.AddButton(loadButton)

	// ProgressIndicator
	indicator = gocoa.NewProgressIndicator(75, 55, 150, 25)
	indicator.Hide()

	indicator.SetLimits(0.00, maxValue)
	indicator.SetValue(0.00)
	indicator.SetAutohide(true)
	wnd.AddProgressIndicator(indicator)

	// Quit button
	quitButton := gocoa.NewButton(75, 25, 150, 25)
	quitButton.SetTitle("Quit")
	quitButton.OnClick(func() {
		gocoa.TerminateApplication()
	})
	wnd.AddButton(quitButton)

	wnd.MakeKeyAndOrderFront()

	gocoa.RunApplication()

}

func increase() {
	time.Sleep(1 * time.Second)
	value := indicator.GetValue()
	inc := 17.00
	indicator.IncrementBy(inc)
	fmt.Printf("Increasing… %f by %f\n", value, inc)
	if indicator.GetValue() < maxValue {
		increase()
	}
}
