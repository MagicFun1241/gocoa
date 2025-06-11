package main

import (
	"fmt"

	"github.com/magicfun1241/gocoa"
	"github.com/magicfun1241/gocoa/appkit"
)

func main() {
	fmt.Println("Starting NSImageView Methods Demo...")

	appkit.InitApplication()

	gocoa.RunOnMainThread(func() {
		// Create a larger window to showcase all features
		window := appkit.NewWindow("NSImageView Methods Demo", 0, 0, 800, 700)
		if window == nil {
			fmt.Println("Failed to create window")
			return
		}

		window.Center()
		window.SetTitle("Complete NSImageView API Demo")
		window.MakeKeyAndOrderFront()
		window.AddDefaultQuitMenu()

		setupImageViewDemo(window)
	})

	fmt.Println("NSImageView Demo running...")
	appkit.RunApplication()
}

func setupImageViewDemo(window *appkit.Window) {
	y := 650 // Start from top

	// Title
	titleLabel := appkit.NewLabel(20, y, 760, 30)
	titleLabel.SetStringValue("NSImageView API Demonstration")
	titleLabel.SetAlignment(appkit.TextAlignmentCenter)
	titleLabel.SetFont("Helvetica-Bold", 18.0)
	titleLabel.SetTextColor("#000080FF")
	window.AddTextField(titleLabel)
	y -= 50

	// 1. Basic ImageView with system image
	imageView1 := appkit.NewImageView(20, y-80, 100, 80)
	imageView1.SetImageName("NSComputer") // System image
	imageView1.SetImageFrameStyle(appkit.FrameStylePhoto)
	imageView1.SetImageAlignment(appkit.ImageAlignCenter)
	imageView1.SetImageScaling(appkit.ImageScalingScaleProportionallyDown)
	window.AddImageView(imageView1)

	label1 := appkit.NewLabel(140, y-40, 300, 40)
	label1.SetStringValue("1. System Image (NSComputer)\nwith Photo Frame Style")
	label1.SetFont("Helvetica", 12.0)
	label1.SetMaximumNumberOfLines(2)
	window.AddTextField(label1)
	y -= 100

	// 2. ImageView with different frame styles
	frameStyles := []struct {
		style appkit.FrameStyle
		name  string
	}{
		{appkit.FrameStyleNone, "None"},
		{appkit.FrameStylePhoto, "Photo"},
		{appkit.FrameStyleGrayBezel, "Gray Bezel"},
		{appkit.FrameStyleGroove, "Groove"},
		{appkit.FrameStyleButton, "Button"},
	}

	frameLabel := appkit.NewLabel(20, y, 760, 25)
	frameLabel.SetStringValue("2. Different Frame Styles:")
	frameLabel.SetFont("Helvetica-Bold", 14.0)
	frameLabel.SetTextColor("#000080FF")
	window.AddTextField(frameLabel)
	y -= 35

	for i, frame := range frameStyles {
		imageView := appkit.NewImageView(20+i*150, y-60, 80, 60)
		imageView.SetImageName("NSFolder")
		imageView.SetImageFrameStyle(frame.style)
		imageView.SetImageAlignment(appkit.ImageAlignCenter)
		imageView.SetImageScaling(appkit.ImageScalingScaleProportionallyDown)
		window.AddImageView(imageView)

		label := appkit.NewLabel(20+i*150, y-75, 80, 15)
		label.SetStringValue(frame.name)
		label.SetAlignment(appkit.TextAlignmentCenter)
		label.SetFont("Helvetica", 10.0)
		window.AddTextField(label)
	}
	y -= 100

	// 3. Image Alignment demonstration
	alignLabel := appkit.NewLabel(20, y, 760, 25)
	alignLabel.SetStringValue("3. Image Alignment Options:")
	alignLabel.SetFont("Helvetica-Bold", 14.0)
	alignLabel.SetTextColor("#000080FF")
	window.AddTextField(alignLabel)
	y -= 35

	alignments := []struct {
		alignment appkit.ImageAlignment
		name      string
	}{
		{appkit.ImageAlignCenter, "Center"},
		{appkit.ImageAlignLeft, "Left"},
		{appkit.ImageAlignRight, "Right"},
		{appkit.ImageAlignTop, "Top"},
		{appkit.ImageAlignBottom, "Bottom"},
	}

	for i, align := range alignments {
		imageView := appkit.NewImageView(20+i*150, y-80, 100, 80)
		imageView.SetImageName("NSStatusAvailable")
		imageView.SetImageAlignment(align.alignment)
		imageView.SetImageScaling(appkit.ImageScalingScaleNone)
		imageView.SetImageFrameStyle(appkit.FrameStyleGrayBezel)
		window.AddImageView(imageView)

		label := appkit.NewLabel(20+i*150, y-95, 100, 15)
		label.SetStringValue(align.name)
		label.SetAlignment(appkit.TextAlignmentCenter)
		label.SetFont("Helvetica", 10.0)
		window.AddTextField(label)
	}
	y -= 120

	// 4. Image Scaling demonstration
	scalingLabel := appkit.NewLabel(20, y, 760, 25)
	scalingLabel.SetStringValue("4. Image Scaling Options:")
	scalingLabel.SetFont("Helvetica-Bold", 14.0)
	scalingLabel.SetTextColor("#000080FF")
	window.AddTextField(scalingLabel)
	y -= 35

	scalings := []struct {
		scaling appkit.ImageScaling
		name    string
	}{
		{appkit.ImageScalingScaleProportionallyDown, "Proportionally Down"},
		{appkit.ImageScalingScaleAxesIndependently, "Axes Independently"},
		{appkit.ImageScalingScaleNone, "None"},
		{appkit.ImageScalingScaleProportionallyUpOrDown, "Proportionally Up/Down"},
	}

	for i, scale := range scalings {
		imageView := appkit.NewImageView(20+i*190, y-80, 120, 80)
		imageView.SetImageName("NSUserGuest")
		imageView.SetImageScaling(scale.scaling)
		imageView.SetImageAlignment(appkit.ImageAlignCenter)
		imageView.SetImageFrameStyle(appkit.FrameStyleGroove)
		window.AddImageView(imageView)

		label := appkit.NewLabel(20+i*190, y-95, 120, 15)
		label.SetStringValue(scale.name)
		label.SetAlignment(appkit.TextAlignmentCenter)
		label.SetFont("Helvetica", 9.0)
		window.AddTextField(label)
	}
	y -= 120

	// 5. Properties demonstration
	propLabel := appkit.NewLabel(20, y, 760, 25)
	propLabel.SetStringValue("5. ImageView Properties:")
	propLabel.SetFont("Helvetica-Bold", 14.0)
	propLabel.SetTextColor("#000080FF")
	window.AddTextField(propLabel)
	y -= 35

	// Editable ImageView
	editableImageView := appkit.NewImageView(20, y-60, 80, 60)
	editableImageView.SetImageName("NSActionTemplate")
	editableImageView.SetEditable(true)
	editableImageView.SetAllowsCutCopyPaste(true)
	editableImageView.SetImageFrameStyle(appkit.FrameStyleButton)
	window.AddImageView(editableImageView)

	editLabel := appkit.NewLabel(110, y-45, 150, 30)
	editLabel.SetStringValue("Editable & Allows\nCut/Copy/Paste")
	editLabel.SetFont("Helvetica", 11.0)
	editLabel.SetMaximumNumberOfLines(2)
	window.AddTextField(editLabel)

	// ImageView with background color
	bgImageView := appkit.NewImageView(280, y-60, 80, 60)
	bgImageView.SetImageName("NSInfo")
	bgImageView.SetBackgroundColor("#FFE0E0FF") // Light red background
	bgImageView.SetImageFrameStyle(appkit.FrameStyleNone)
	window.AddImageView(bgImageView)

	bgLabel := appkit.NewLabel(370, y-45, 150, 30)
	bgLabel.SetStringValue("Custom Background\nColor")
	bgLabel.SetFont("Helvetica", 11.0)
	bgLabel.SetMaximumNumberOfLines(2)
	window.AddTextField(bgLabel)

	// ImageView with alpha
	alphaImageView := appkit.NewImageView(540, y-60, 80, 60)
	alphaImageView.SetImageName("NSRefreshTemplate")
	alphaImageView.SetAlphaValue(0.5) // 50% transparency
	alphaImageView.SetImageFrameStyle(appkit.FrameStylePhoto)
	window.AddImageView(alphaImageView)

	alphaLabel := appkit.NewLabel(630, y-45, 150, 30)
	alphaLabel.SetStringValue("50% Alpha\n(Transparency)")
	alphaLabel.SetFont("Helvetica", 11.0)
	alphaLabel.SetMaximumNumberOfLines(2)
	window.AddTextField(alphaLabel)
	y -= 100

	// 6. Method testing display
	methodLabel := appkit.NewLabel(20, y, 760, 25)
	methodLabel.SetStringValue("6. Method Results (Check Console):")
	methodLabel.SetFont("Helvetica-Bold", 14.0)
	methodLabel.SetTextColor("#000080FF")
	window.AddTextField(methodLabel)
	y -= 35

	// Test and display method results
	testImageView := appkit.NewImageView(20, y-60, 80, 60)
	testImageView.SetImageName("NSCaution")
	testImageView.SetImageFrameStyle(appkit.FrameStylePhoto)
	testImageView.SetEditable(true)
	testImageView.SetAnimates(true)
	window.AddImageView(testImageView)

	// Display method results
	resultText := fmt.Sprintf(
		"Image Name: %s\nEditable: %t\nAnimates: %t\nFrame Style: %d\nAlignment: %d\nScaling: %d\nAlpha: %.2f",
		testImageView.ImageName(),
		testImageView.Editable(),
		testImageView.Animates(),
		int(testImageView.ImageFrameStyle()),
		int(testImageView.ImageAlignment()),
		int(testImageView.ImageScaling()),
		testImageView.AlphaValue(),
	)

	resultLabel := appkit.NewLabel(110, y-120, 650, 120)
	resultLabel.SetStringValue(resultText)
	resultLabel.SetFont("Monaco", 10.0) // Monospace font
	resultLabel.SetMaximumNumberOfLines(0)
	resultLabel.SetLineBreakMode(appkit.LineBreakByWordWrapping)
	resultLabel.SetBackgroundColor("#F0F0F0FF")
	resultLabel.SetDrawsBackground(true)
	window.AddTextField(resultLabel)

	fmt.Println("=== NSImageView Method Test Results ===")
	fmt.Printf("Image Name: %s\n", testImageView.ImageName())
	fmt.Printf("Editable: %t\n", testImageView.Editable())
	fmt.Printf("Animates: %t\n", testImageView.Animates())
	fmt.Printf("Frame Style: %d\n", int(testImageView.ImageFrameStyle()))
	fmt.Printf("Image Alignment: %d\n", int(testImageView.ImageAlignment()))
	fmt.Printf("Image Scaling: %d\n", int(testImageView.ImageScaling()))
	fmt.Printf("Alpha Value: %.2f\n", testImageView.AlphaValue())
	fmt.Printf("Allows Cut/Copy/Paste: %t\n", testImageView.AllowsCutCopyPaste())
	fmt.Printf("Hidden: %t\n", testImageView.Hidden())
	fmt.Println("=====================================")

	fmt.Println("All NSImageView methods demonstrated!")
}
