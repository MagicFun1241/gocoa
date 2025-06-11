package main

import (
	"fmt"

	"github.com/magicfun1241/gocoa"
	"github.com/magicfun1241/gocoa/appkit"
)

func main() {
	fmt.Println("Starting NSTextField Methods Demo...")

	appkit.InitApplication()

	gocoa.RunOnMainThread(func() {
		// Create a larger window to showcase all features
		window := appkit.NewWindow("NSTextField Methods Demo", 0, 0, 500, 600)
		if window == nil {
			fmt.Println("Failed to create window")
			return
		}

		window.Center()
		window.SetTitle("Complete NSTextField API Demo")
		window.MakeKeyAndOrderFront()
		window.AddDefaultQuitMenu()

		setupTextFieldDemo(window)
	})

	fmt.Println("NSTextField Demo running...")
	appkit.RunApplication()
}

func setupTextFieldDemo(window *appkit.Window) {
	y := 550 // Start from top

	// 1. Basic label with centered text
	label1 := appkit.NewLabel(20, y, 460, 30)
	label1.SetStringValue("1. Centered Text Label")
	label1.SetAlignment(appkit.TextAlignmentCenter)
	label1.SetFont("Helvetica-Bold", 16.0)
	label1.SetTextColor("#000080FF") // Dark blue
	window.AddTextField(label1)
	y -= 40

	// 2. Multi-line label with word wrapping
	label2 := appkit.NewLabel(20, y-40, 460, 80)
	label2.SetStringValue("2. Multi-line Label: This is a long text that demonstrates word wrapping and multiple lines. The text will wrap automatically when it reaches the edge of the text field.")
	label2.SetAlignment(appkit.TextAlignmentLeft)
	label2.SetMaximumNumberOfLines(0) // Unlimited lines
	label2.SetLineBreakMode(appkit.LineBreakByWordWrapping)
	label2.SetFont("Helvetica", 12.0)
	window.AddTextField(label2)
	y -= 100

	// 3. Text field with placeholder
	textField1 := appkit.NewTextField(20, y, 200, 25)
	textField1.SetPlaceholderString("Enter your name...")
	textField1.SetAlignment(appkit.TextAlignmentCenter)
	textField1.SetFont("Helvetica", 12.0)
	window.AddTextField(textField1)

	// Label for the above text field
	label3 := appkit.NewLabel(230, y, 250, 25)
	label3.SetStringValue("← Text field with placeholder")
	label3.SetAlignment(appkit.TextAlignmentLeft)
	label3.SetFont("Helvetica", 11.0)
	label3.SetTextColor("#666666FF") // Gray
	window.AddTextField(label3)
	y -= 40

	// 4. Single line mode demonstration
	textField2 := appkit.NewTextField(20, y, 200, 25)
	textField2.SetStringValue("Single line mode - no wrapping")
	textField2.SetUsesSingleLineMode(true)
	textField2.SetLineBreakMode(appkit.LineBreakByTruncatingTail)
	textField2.SetFont("Helvetica", 12.0)
	window.AddTextField(textField2)

	label4 := appkit.NewLabel(230, y, 250, 25)
	label4.SetStringValue("← Single line with truncation")
	label4.SetAlignment(appkit.TextAlignmentLeft)
	label4.SetFont("Helvetica", 11.0)
	label4.SetTextColor("#666666FF")
	window.AddTextField(label4)
	y -= 40

	// 5. Different alignment examples
	alignments := []struct {
		alignment int
		name      string
	}{
		{appkit.TextAlignmentLeft, "Left aligned text"},
		{appkit.TextAlignmentCenter, "Center aligned text"},
		{appkit.TextAlignmentRight, "Right aligned text"},
	}

	for i, align := range alignments {
		textField := appkit.NewTextField(20, y, 460, 25)
		textField.SetStringValue(align.name)
		textField.SetAlignment(align.alignment)
		textField.SetFont("Helvetica", 12.0)
		textField.SetEditable(false)
		textField.SetSelectable(false)

		// Different background colors for visibility
		colors := []string{"#FFE0E0FF", "#E0FFE0FF", "#E0E0FFFF"}
		textField.SetBackgroundColor(colors[i])
		textField.SetDrawsBackground(true)

		window.AddTextField(textField)
		y -= 30
	}

	// 6. Line break mode examples
	lineBreakModes := []struct {
		mode int
		text string
	}{
		{appkit.LineBreakByWordWrapping, "Word wrapping: This text will wrap at word boundaries when it's too long"},
		{appkit.LineBreakByCharWrapping, "Character wrapping: This text will wrap at any character when too long"},
		{appkit.LineBreakByClipping, "Clipping: This text will be clipped when it's too long for the field"},
		{appkit.LineBreakByTruncatingTail, "Truncating tail: This text will show ... at the end when too long"},
	}

	y -= 20
	titleLabel := appkit.NewLabel(20, y, 460, 25)
	titleLabel.SetStringValue("Line Break Modes:")
	titleLabel.SetFont("Helvetica-Bold", 14.0)
	titleLabel.SetTextColor("#000080FF")
	window.AddTextField(titleLabel)
	y -= 30

	for _, mode := range lineBreakModes {
		textField := appkit.NewTextField(20, y, 460, 40)
		textField.SetStringValue(mode.text)
		textField.SetLineBreakMode(mode.mode)
		textField.SetMaximumNumberOfLines(2)
		textField.SetFont("Helvetica", 11.0)
		textField.SetEditable(false)
		textField.SetBackgroundColor("#F5F5F5FF")
		textField.SetDrawsBackground(true)
		window.AddTextField(textField)
		y -= 50
	}

	// 7. Font demonstration
	y -= 20
	fontLabel := appkit.NewLabel(20, y, 460, 25)
	fontLabel.SetStringValue("Different Fonts and Sizes:")
	fontLabel.SetFont("Helvetica-Bold", 14.0)
	fontLabel.SetTextColor("#000080FF")
	window.AddTextField(fontLabel)
	y -= 30

	fonts := []struct {
		name string
		size float64
		text string
	}{
		{"Times-Roman", 14.0, "Times Roman 14pt"},
		{"Courier", 12.0, "Courier 12pt (monospace)"},
		{"Helvetica-Oblique", 13.0, "Helvetica Oblique 13pt"},
		{"Monaco", 11.0, "Monaco 11pt (monospace)"},
	}

	for _, font := range fonts {
		textField := appkit.NewTextField(20, y, 460, 25)
		textField.SetStringValue(font.text)
		textField.SetFont(font.name, font.size)
		textField.SetEditable(false)
		textField.SetAlignment(appkit.TextAlignmentCenter)
		window.AddTextField(textField)
		y -= 30
	}

	fmt.Println("All NSTextField methods demonstrated!")
}
