package main

import (
	"fmt"
	"log"

	"github.com/magicfun1241/gocoa"
	"github.com/magicfun1241/gocoa/appkit"
)

func main() {
	fmt.Println("Starting Centered Button Demo with Auto Layout...")

	// Initialize the application
	appkit.InitApplication()

	// Use RunOnMainThread to ensure all UI operations happen on the main thread
	gocoa.RunOnMainThread(func() {
		// Create main window with 400x300 size (matching center.c exactly)
		window := appkit.NewWindow("Centered Button", 0, 0, 400, 300)
		if window == nil {
			log.Fatal("Failed to create window")
		}

		// Center the window on screen
		window.Center()

		// Set window properties
		window.SetTitle("Centered Button")
		window.MakeKeyAndOrderFront()
		window.AddDefaultQuitMenu()

		// Create and center the button using auto layout constraints
		setupCenteredButtonWithConstraints(window)
	})

	// Run the application
	fmt.Println("Centered Button Demo with Auto Layout running...")
	appkit.RunApplication()
}

func setupCenteredButtonWithConstraints(window *appkit.Window) {
	// Get the window's content view
	contentView := window.GetContentView()
	if contentView == nil {
		log.Fatal("Failed to get content view")
		return
	}

	// Create the button with arbitrary frame (auto layout will handle positioning)
	button := appkit.NewButton(0, 0, 300, 300)
	if button == nil {
		log.Fatal("Failed to create button")
		return
	}

	button.SetTitle("Click Me")
	button.SetBezelStyle(appkit.ButtonBezelStyleRounded)

	// Add click handler
	button.OnClick(func() {
		fmt.Println("Auto layout centered button clicked!")
	})

	// Add button to the window first
	window.AddButton(button)

	// Create center X constraint: button.centerX = contentView.centerX
	centerXConstraint := appkit.CreateConstraint(
		button.Ptr(),                  // button (as unsafe.Pointer)
		appkit.LayoutAttributeCenterX, // button's center X
		appkit.LayoutRelationEqual,    // equal to
		contentView.Ptr(),             // content view
		appkit.LayoutAttributeCenterX, // content view's center X
		1.0,                           // multiplier
		0.0,                           // constant
	)

	// Create center Y constraint: button.centerY = contentView.centerY
	centerYConstraint := appkit.CreateConstraint(
		button.Ptr(),                  // button (as unsafe.Pointer)
		appkit.LayoutAttributeCenterY, // button's center Y
		appkit.LayoutRelationEqual,    // equal to
		contentView.Ptr(),             // content view
		appkit.LayoutAttributeCenterY, // content view's center Y
		1.0,                           // multiplier
		0.0,                           // constant
	)

	// Activate the constraints
	appkit.ActivateConstraints([]*appkit.LayoutConstraint{
		centerXConstraint,
		centerYConstraint,
	})

	fmt.Println("Auto layout constraints applied - button should be perfectly centered!")
}
