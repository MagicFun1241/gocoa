package main

import (
	"fmt"
	"log"
	"time"

	"github.com/magicfun1241/gocoa"
	"github.com/magicfun1241/gocoa/appkit"
)

func main() {
	fmt.Println("Starting Gocoa Demo Application...")

	// Initialize the application
	appkit.InitApplication()

	// Use RunOnMainThread to ensure all UI operations happen on the main thread
	gocoa.RunOnMainThread(func() {
		// Create main window - correct signature: NewWindow(title, x, y, width, height)
		window := appkit.NewWindow("Gocoa Demo - Testing Library Components", 100, 100, 800, 600)
		if window == nil {
			log.Fatal("Failed to create window")
		}

		// Set window properties
		window.SetTitle("Gocoa Component Demo")
		window.MakeKeyAndOrderFront()
		window.AddDefaultQuitMenu()

		// Demo different components
		setupWindowComponents(window)
	})

	// Run the application
	fmt.Println("Demo application running...")
	appkit.RunApplication()
}

func setupWindowComponents(window *appkit.Window) {
	// Create a button demo
	createButtonDemo(window)

	// Create text field demo
	createTextFieldDemo(window)

	// Create other component demos
	createProgressIndicatorDemo(window)
	createComboBoxDemo(window)
}

func createButtonDemo(window *appkit.Window) {
	// NewButton requires x, y, width, height
	button := appkit.NewButton(50, 500, 150, 30)
	if button != nil {
		button.SetTitle("Test Button")
		button.OnClick(func() {
			fmt.Println("Button clicked! Library is working.")
		})
		window.AddButton(button)
	}
}

func createTextFieldDemo(window *appkit.Window) {
	// NewTextField requires x, y, width, height
	textField := appkit.NewTextField(50, 450, 200, 25)
	if textField != nil {
		textField.SetStringValue("Demo text field")
		window.AddTextField(textField)
	}
}

func createProgressIndicatorDemo(window *appkit.Window) {
	progress := appkit.NewProgressIndicator(50, 350, 200, 25)
	if progress != nil {
		progress.SetLimits(0.0, 100.0)
		progress.SetValue(25.0)
		progress.Show()
		window.AddProgressIndicator(progress)

		// Animate progress for demo
		go func() {
			for i := 0; i <= 100; i += 5 {
				time.Sleep(100 * time.Millisecond)

				gocoa.RunOnMainThread(func() {
					progress.SetValue(float64(i))
				})
			}
		}()
	}
}

func createComboBoxDemo(window *appkit.Window) {
	comboBox := appkit.NewComboBox(50, 300, 200, 25)
	if comboBox != nil {
		comboBox.AddItem("Option 1")
		comboBox.AddItem("Option 2")
		comboBox.AddItem("Option 3")
		comboBox.SetSelectedIndex(0)
		window.AddComboBox(comboBox)
	}
}
