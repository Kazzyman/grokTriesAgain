package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"strconv"
)

// @formatter:off


// Three Additional Windows: 
// ::: ------------------------------------------------------------------------------------------------------------------------------------------------------------
func createWindow2() fyne.Window {
	window2.Resize(fyne.NewSize(1900, 1600))
	outputLabel2.Wrapping = fyne.TextWrapWord
	scrollContainer2.SetMinSize(fyne.NewSize(1900, 1090))
	coloredScroll2 := container.NewMax(bgsc2, scrollContainer2) // Light blue-ish scroll bg

	// set up user data entry for RootsBtn2  
	radicalEntry := widget.NewEntry()
	radicalEntry.SetPlaceHolder("Enter radical index (e.g., 2 or 3)")
	workEntry := widget.NewEntry()
	workEntry.SetPlaceHolder("Enter number to find the root of")

	done := make(chan bool) // local, kill channel for all goroutines that are listening: ::: not entirely sure of this one ???

	// ::: Get single input dialog < - - - - - - - - - - - - - - - - - - - - - - - - < -
	getSingleInput2 := func(title, prompt, defaultValue string, callback func(string, bool)) {
		confirmed := false // Track if OK was clicked
		d := dialog.NewEntryDialog(title, prompt, func(value string) {
			confirmed = true
			callback(value, true)
		}, window2)
		d.SetText(defaultValue)
		d.SetOnClosed(func() {
			if !confirmed { // Only trigger cancel if OK wasn’t clicked
				callback("", false)
			}
		})
		d.Show()
	}
	// ::: Dual input dialog < - - - - - - - - - - - - - - - - - - - - - - - - < -
	getDualInput2 := func(title, prompt1, prompt2, default1, default2 string, callback func(string, string, bool)) {
		calculating = true
		for _, btn := range buttons2 {
			btn.Disable()
		}
		entry1 := widget.NewEntry()
		entry1.SetText(default1)
		entry2 := widget.NewEntry()
		entry2.SetText(default2)
		submitButton := widget.NewButton("Run with those values",
			func() {
				callback(entry1.Text, entry2.Text, true)
				dialog.NewInformation("Submitted", "Values submitted", window2).Hide() // Hack to close dialog
				calculating = true
				for _, btn := range buttons2 {
					btn.Disable()
				}
			})
		form := container.NewVBox(
			widget.NewLabel(prompt1), entry1,
			widget.NewLabel(prompt2), entry2,
			container.NewHBox(submitButton),
		)
		d := dialog.NewCustom(title, "Dismiss dialogBox", form, window2)

		d.Resize(fyne.NewSize(400, 300))
		d.Show()
	}


	// ::: Buttons2

	// This button handler is different from the other ones here in that it gets its input from a widget.NewEntry() 
	RootsBtn2 := SetupRootsDemo(mgr, done, radicalEntry, workEntry) // ::: just one line here, for this math button (see SetupRootsDemo)
	/*
		.
		.
	*/

	NilakanthaBtn2 := NewColoredButton(
		"Nilakantha -- input iterations\n" +
			"output up to 26 digits of pi",
		color.RGBA{255, 255, 100, 235},
		func() {
			if calculating {
				return
			}
			calculating = true
			for _, btn := range buttons2 {
				btn.Disable()
			}
			for _, btn := range nilaBut2 { // Refer to the comments in the initial assignment and creation of archimedesBtn1
				calculating = true
				btn.Enable()
			}
			getDualInput2("Input Required", "Number of iterations (suggest 300,000 -> 30,000,000  -> 300,000,000):", "Precision (suggest 128):",
				"30000000", "128", // 30,000,000
				func(itersStr, precStr string, ok bool) {
					calculating = true
					for _, btn := range buttons2 {
						btn.Disable()
					}
					if !ok {
						updateOutput2("Nilakantha calculation canceled")
						return
					}
					iters := 30000000 // 30,000,000
					precision := 128
					itersStr = removeCommasAndPeriods(itersStr) // ::: allow user to enter a number with a comma
					val1, err1 := strconv.Atoi(itersStr)
					if err1 != nil {
						fmt.Println("Error converting iterations val1:", err1) // handle error
						iters = 30000000
					} else {
						iters = val1
					}
					val2, err2 := strconv.Atoi(precStr)
					if err2 != nil {
						fmt.Println("Error converting precision val2:", err2) // handle error 
						updateOutput2("setting precision to 128")
						// fyneFunc(fmt.Sprintf("setting precision to 512")) //  ::: cannot do this instead because ??
						precision = 128
					} else {
						precision = val2
					}
					go NilakanthaBig(updateOutput2, iters, precision, done) // ::: probably want to add a done channel to this one
					calculating = false
					for _, btn := range buttons2 {
						btn.Enable()
					}
				})
		})

	// ::: Chud is temp here, Bailey concur will go here eventually 
	ChudnovskyBtn2 := NewColoredButton("chudnovsky -- takes input", color.RGBA{255, 255, 100, 235},
		func() {
			if calculating {
				return
			}
			calculating = true
			for _, btn := range buttons2 {
				btn.Disable()
			}
			for _, btn := range chudBut2 { // Refer to the comments in the initial assignment and creation of archimedesBtn1
				calculating = true
				btn.Enable()
			}
			getSingleInput2("Input Required", "Enter the number of digits for the chudnovsky calculation (e.g., 46):", "46",
				func(digitsStr string, ok bool) {
					var chudDigits int
					if !ok {
						updateOutput2("chudnovsky calculation canceled")
						return
					}
					chudDigits = 46
					val, err := strconv.Atoi(digitsStr)
					if err != nil {
						fmt.Println("Error converting input:", err) // handel error 
						updateOutput2("Invalid input, using default 46 digits")
					} else if val <= 0 {
						updateOutput2("Input must be positive, using default 46 digits")
					} else if val > 10000 {
						updateOutput2("Input must be less than 10,001, using default 46 digits")
					} else {
						chudDigits = val
					}
					go func() {
						chudnovskyBig(updateOutput2, chudDigits, done)
						calculating = false
						for _, btn := range buttons2 {
							btn.Enable()
						}
					}()
				})
		})

	/*
		.
		.
	*/

	rootBut2 = []*ColoredButton{RootsBtn2} // these are a slick trick/kluge 
	chudBut2 = []*ColoredButton{ChudnovskyBtn2} // All these are a trick/kluge used as bug preventions // to keep methods from being started or restarted in parallel (over-lapping)
	nilaBut2 = []*ColoredButton{NilakanthaBtn2}

	buttons2 = []*ColoredButton{RootsBtn2, NilakanthaBtn2, ChudnovskyBtn2} // array used only for range btn.Enable()

	// ::: page-2 Lay-out
	content2 := container.NewVBox(
		widget.NewLabel("\nSelect a method to estimate π:\n"),

		radicalEntry,
		workEntry,

		container.NewGridWithColumns(4, RootsBtn2, NilakanthaBtn2, ChudnovskyBtn2),

		coloredScroll2, // Use coloredScroll2 directly or windowContent2 if you want an extra layer
	)
	windowContent2 := container.NewMax(bgwc2, content2) // Light green window bg // containers withing containers, labels on labels, functions in functions. Yet still inert. 

	/*
	      	window2.Canvas().SetOnTypedRune(func(r rune) { // Main-thread update loop using Fyne's lifecycle -- here an empty loop ::: see below:
	      	})

	      Every Fyne window has a Canvas, which is the drawable surface where all widgets (buttons, labels, etc.) are rendered. Calling window2.Canvas() gives you access to this canvas,
	   letting you interact with its properties or events.

	      .SetOnTypedRune(func(r rune) { ... }):
	      This method sets a callback function that Fyne calls whenever a user types a character (a "rune") into the window, provided the window has focus.

	      A rune in Go is an alias for int32 and represents a Unicode code point—essentially a single character, like 'a', '5', or 'π'. It’s more general than a byte, allowing it to handle
	   all kinds of text input (e.g., emojis, non-Latin scripts).

	      func(r rune) { ... }:
	      This is the callback function you provide. It runs on the main thread whenever a key is typed, and it receives the typed character (r) as an argument. The body of this
	   function (which you’ve shown as empty {}) is where you’d define what happens when a key is pressed.

	      "Main-thread update loop using Fyne's lifecycle":
	      The comment suggests this is part of Fyne’s event-driven lifecycle. Fyne runs its GUI in a single-threaded, event-based model on the main thread. When you set this callback, it
	   hooks into that lifecycle, ensuring your response to keypresses happens synchronously with other GUI updates (like rendering or widget changes). This avoids concurrency issues that
	   could arise if you tried to update the GUI from another thread.
	   :::
	      In short, window2.Canvas().SetOnTypedRune(func(r rune) { ... }) lets you capture and respond to keyboard input in window2. For example:
	      If a user types 'q', the function runs with r = 'q'.

	      You could use this to close the window, update a label, or trigger a calculation based on the input.

	   :::    Since your example has an empty function body ({}), it currently does nothing—it’s just a placeholder. The real action depends on what you put inside the {}.

	*/

	window2.SetContent(windowContent2) // Set once with the full layout
	return window2
} // end of createWindow2 "it's only a label". "no Show, no go" 


// ::: ------------------------------------------------------------------------------------------------------------------------------------------------------------
func createWindow3(myApp fyne.App) fyne.Window {
	// Similar structure to createWindow2
	window3 := myApp.NewWindow("Odd Pi calculators")
	window3.Resize(fyne.NewSize(1900, 1600))
	outputLabel3 := widget.NewLabel("Odd Pi calculators, make a selection")
	outputLabel3.Wrapping = fyne.TextWrapWord
	scrollContainer3 := container.NewScroll(outputLabel3)
	scrollContainer3.SetMinSize(fyne.NewSize(1900, 1300))
	buttonContainer3 := container.NewGridWithColumns(4,
		widget.NewButton("Button 9", func() {}),
		widget.NewButton("Button 10", func() {}),
		widget.NewButton("Button 11", func() {}),
		widget.NewButton("Button 12", func() {}),
		widget.NewButton("Button 13", func() {}),
		widget.NewButton("Button 14", func() {}),
		widget.NewButton("Button 15", func() {}),
		widget.NewButton("Button 16", func() {}),
	)
	content3 := container.NewVBox(buttonContainer3, scrollContainer3)
	window3.SetContent(content3)
	return window3
}

// ::: ------------------------------------------------------------------------------------------------------------------------------------------------------------
func createWindow4(myApp fyne.App) fyne.Window {
	// Similar structure to createWindow2
	window4 := myApp.NewWindow("Misc Maths")
	window4.Resize(fyne.NewSize(1900, 1600))
	outputLabel4 := widget.NewLabel("Misc Maths, make a selection")
	outputLabel4.Wrapping = fyne.TextWrapWord
	scrollContainer4 := container.NewScroll(outputLabel4)
	scrollContainer4.SetMinSize(fyne.NewSize(1900, 1300))
	buttonContainer4 := container.NewGridWithColumns(4,
		widget.NewButton("Button 17", func() {}), widget.NewButton("Button 18", func() {}), widget.NewButton("Button 19", func() {}), widget.NewButton("Button 20", func() {}),
		widget.NewButton("Button 21", func() {}), widget.NewButton("Button 22", func() {}), widget.NewButton("Button 23", func() {}), widget.NewButton("Button 24", func() {}),
	)
	content4 := container.NewVBox(buttonContainer4, scrollContainer4)
	window4.SetContent(content4)
	return window4
}
