package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"strconv"
)

// @formatter:off

func main() {
	countAndLogSLOC() // this one func is the only actual call-to-action apart from the last line in main() : window1.ShowAndRun() -- the setting of variables notwithstanding! 
	calculating = false // set the global the-coast-is-clear flag
	myApp.Settings().SetTheme(theme.LightTheme()) // establish a Theme that will work well with dialog boxes
	window1.Resize(fyne.NewSize(1900, 1600))

	scrollContainer1 = container.NewVScroll(outputLabel1)
	
	scrollContainer1.SetMinSize(fyne.NewSize(1900, 930)) // was before adding the scoreBoard 1900, 1050
	
	outputLabel1.Wrapping = fyne.TextWrapWord // make the text in the scrollable area auto-wrap

		coloredScroll := container.NewMax(bgsc, scrollContainer1) // Combine background and scroll, Layer light green background behind scroll content.
	
		windowContent := container.NewMax(bgwc, coloredScroll) // Layer the background and content; Layer light blue background across the entire window content.

/*
.
.
 */
	// Terminal-like display
	terminalDisplay := widget.NewTextGrid()
	terminalDisplay.SetText("Terminal Output:\n\nWaiting for calculation...")

	// Button only being used as a title-label for nifty_scoreBoard
	calcButton := widget.NewButton("Also try our special menu option: Calculate Pi on a ScoreBoard", func() {
		updateOutput1("\n- * - * - that button does nothing - * - * -\n\n")
	})

	// Layout for scoreboard section
	contentForScoreBoard := container.NewVBox(
		calcButton,
		terminalDisplay,
	)
/*
.
.
 */
	// Custom colored ::: Buttons1 - - - - - - - - - follow - - - - - - - - - - - v v v v v v v v v - - - - - - 
	/*
	.
	.
	 */
	archimedesBtn1 := NewColoredButton(
	"Archimedes method for finding π, modified by Richard Woolley\n" +
		"easy to understand geometric method using big.Float variables\n" +
		"produces 3,012 digits of delicious Pi in under a minute, 230BCE\n" +
		"             -*-*-*- Rick's personal favorite -*-*-*-          ",
		color.RGBA{255, 110, 110, 215},
		
		func() {
			if calculating {
				return
			}
			calculating = true
			for _, btn := range buttons1 {
				btn.Disable()
			}
			// We want to cause the button that corresponds to the currently executing method to remain bright, while the other buttons remain dimmed...
			for _, btn := range archiBut { // This trick accomplishes that because the archiBut array comes after the creation of archimedesBtn1
				calculating = true // This keeps archimedesBtn1 from being restarted in parallel with itself...
				btn.Enable() // ... even though we herewith enable archimedesBtn1 ... note that simply doing: archimedesBtn1.Enable() would not work...
			} //  ... because, we are inside of the creation of archimedesBtn1 [ it is a timing and scoping issue ]
			currentDone = make(chan bool) // ::: New channel per run
			updateOutput1("\nRunning ArchimedesBig...\n\n")
			go func(done chan bool) { // ::: go func now takes an argument
					defer func() {       // ::: new defer func with global calculating flag set 
						calculating = false
						updateOutput1("Calculation definitely finished; possibly aborted\n")
					}()
				ArchimedesBig(updateOutput1, done) // ::: func < - - - - - - - - - - - - - < -
					calculating = false
					for _, btn := range buttons1 {
						btn.Enable()
					}
			}(currentDone) // ::: pass via closure
			/*
			passing the currentDone channel via a closure to the goroutine. This is a common and idiomatic way in Go to ensure that the goroutine uses the specific 
			channel instance you’ve just created (or assigned) within the button handler, rather than relying on some outer or potentially stale reference.
			*/
		},
	)
	/*
	.
	.
	 */
	
	JohnWallisBtn1 := NewColoredButton(
	"John Wallis infinite series -- 40 billion iterations -- runs 5m30s\n" +
		"π = 2 * ((2/1)*(2/3)) * ((4/3)*(4/5)) * ((6/5)*(6/7)) ...\n" +
		"only manages to bake 10 digits of Pi in well-over five minutes\n" +
		"an infinite series circa 1655    --- served here by Rick Woolley ---",
		color.RGBA{110, 110, 255, 185}, 
		
		func() {
			if calculating {
				return
			}
			calculating = true
			for _, btn := range buttons1 {
				btn.Disable()
			}
			for _, btn := range walisBut { // Refer to the comments in the initial assignment and creation of archimedesBtn1
				btn.Enable()
			}
			currentDone = make(chan bool) // ::: New channel per run
			updateOutput1("\nRunning John Wallis...\n\n")
			go func(done chan bool) { // ::: go func now takes an argument
				defer func() {       // ::: new defer func with global calculating flag set 
					calculating = false // this does not appear to work 
					updateOutput1("Calculation definitely finished; possibly aborted\n")
				}()
						fmt.Printf("here before JohnWallisBtn1 calculating is %t\n", calculating) // this executes 
				pie = JohnWallis(updateOutput1, done) // ::: func < - - - - - - - - - - - - - < -
						fmt.Printf("here after JohnWallisBtn1 calculating is %t\n", calculating) // this does not execute, not does the first line in JohnWallis()
						
					current := outputLabel1.Text
					outputLabel1.SetText(current + fmt.Sprintf("\n\nπ ≈ %.11f\n", pie))
				calculating = false
				for _, btn := range buttons1 {
					btn.Enable()
				}
			}(currentDone)
			fmt.Printf("here at the end of JohnWallisBtn1 calculating is %t\n", calculating) // this executes 
		},
	)
/*
.
.
 */

	SpigotBtn1 := NewColoredButton(
	"The Spigot Algorithm, a Leibniz series. Served hot, bite by byte\n" +
		"spits out a nearly-unlimited, continuous stream of Pi goodness\n" +
		"This trick made possible by a bit of code mooched off of GitHub\n" +
		"bakes π without using any floating-point arithmetic",
		color.RGBA{255, 255, 100, 235},
		
		func() {
			var spigotDigits int = 1460 // to resolve a scoping issue 
			if calculating {
				return
			}
			calculating = true
			for _, btn := range buttons1 {
				btn.Disable()
			}
			for _, btn := range spigotBut { // Refer to the comments in the initial assignment and creation of archimedesBtn1
				calculating = true
				btn.Enable()
			}
			currentDone = make(chan bool) // ::: New channel per run
			updateOutput1("\nRunning The Spigot...\n\n")
			
			showCustomEntryDialog(
				"Input Desired number of digits",
				"Any number less than 1461",
				func(input string) {
					if input != "" { // This if-else is part of the magic that allows us to dismiss a dialog and allow others to run after the dialog is canceled/dismissed.
						input = removeCommasAndPeriods(input) // allow user to enter a number with a comma
						val, err := strconv.Atoi(input)
						if err != nil { // we may force val to become 460, or leave it alone ...
							fmt.Println("Error converting input:", err)
							updateOutput1("\nInvalid input, using default 1460 digits\n")
							val = 1460
						} else if val <= 0 {
							updateOutput1("\nInput must be positive, using default 1460 digits\n")
							val = 1460
						} else if val > 1460 {
							updateOutput1("\nInput must be less than 1461 -- using default of 1460 digits\n")
							val = 1460 
						} else {
							spigotDigits = val // resolves a scoping issue 
						}
						
						go func(done chan bool) { // ::: go func now takes an argument
							defer func() {       // ::: new defer func with global calculating flag set 
								calculating = false // this does not appear to work 
								updateOutput1("\nCalculation definitely finished; possibly aborted\n")
							}()
							TheSpigot(updateOutput1, spigotDigits, done) // ::: func < - - - - - - - - - - - - - < -  NOT AMENABLE TO KILLING VIA A DONE CHANNEL 
							calculating = false
							for _, btn := range buttons1 {
								btn.Enable()
							}
						}(currentDone)
					} else {
						// dialog canceled 
						updateOutput1("\nspigot calculation canceled, make another selection\n")
						for _, btn := range buttons1 {
							btn.Enable()
						}
						calculating = false // ::: this is the trick to allow others to run after the dialog is canceled/dismissed.
					}
				},
			)
		}, 
	)
	/*
	.
	.
	pi = \frac{1}{12} \left[ \sum_{n=0}^{\infty} \frac{(-1)^n (6n)! (13591409 + 545140134n)}{(3n)! (n!)^3 (640320^{3n + 3/2})} \right]^{-1}      */ 
	ChudnovskyBtn1 := NewColoredButton(
	"Chudnovsky -- by David & Gregory Chudnovsky -- late 1980s\n" +
		"extremely efficient, quickly bakes world-record quantities of Pi\n" +
		"this algorithm is a rapidly converging infinite series which\n" +
		"leverages properties of j-invariant from elliptic function theory",
		color.RGBA{100, 255, 100, 215}, 
		
		func() {
			// 
			var chudDigits int
				if calculating {
					return
				}
				calculating = true
				for _, btn := range buttons1 {
					btn.Disable()
				}
				for _, btn := range chudBut { 
					calculating = true 
					btn.Enable() 
				}
			currentDone = make(chan bool) // ::: New channel per run
			updateOutput1("\nRunning Chudnovsky...\n\n")
	
			showCustomEntryDialog(
				"Input Desired number of digits",
				"Any number less than 49,999",
				func(input string) {
					if input != "" { // This if-else is part of the magic that allows us to dismiss a dialog and allow others to run after the dialog is canceled/dismissed.
						input = removeCommasAndPeriods(input) // allow user to enter a number with a comma
						val, err := strconv.Atoi(input)
						if err != nil {
							fmt.Println("Error converting input:", err)
							updateOutput1("Invalid input, using default 49,000 digits")
						} else if val <= 0 {
							updateOutput1("Input must be positive, using default 49000 digits")
						} else if val > 50000 {
							updateOutput1("Input must be less than 50,000 -- using default of 49,000 digits")
						} else {
							chudDigits = val
						}
						go func(done chan bool) { // ::: go func now takes an argument
							defer func() {       // ::: new defer func with global calculating flag set 
								calculating = false // this does not appear to work 
								updateOutput1("Calculation definitely finished; possibly aborted\n")
							}()
							chudnovskyBig(updateOutput1, chudDigits, done) // ::: func < - - - - - - - - - - - - - < -  NOT AMENABLE TO KILLING VIA A DONE CHANNEL 
							calculating = false
							for _, btn := range buttons1 {
								btn.Enable()
							}
						}(currentDone)
					} else {
						// dialog canceled 
							updateOutput1("chudnovsky calculation canceled, make another selection")
							for _, btn := range buttons1 {
								btn.Enable()
							}
							calculating = false // ::: this is the trick to allow others to run after the dialog is canceled/dismissed.
					}
				},
			)
		},
	)
	/*
	.
	.
	 */

	MontyBtn1 := NewColoredButton(
		"Monte Carlo method for converging on π  --  big floats, & float64\n" +
			"Flavor: no fancy equations are used, only Go's pure randomness\n" +
			"4 digits of pi in 21s; 7 digits possible in 1h30m with a 119k grid\n" +
			"                   -*-*- Rick's second-favorite method -*-*-     ",
		color.RGBA{255, 255, 100, 235},

		func() {
			var MontDigits string
			if calculating {
				return
			}
			calculating = true
			for _, btn := range buttons1 {
				btn.Disable()
			}
			for _, btn := range montBut {
				calculating = true
				btn.Enable()
			}
			currentDone = make(chan bool) // ::: New channel per run
			updateOutput1("\nRunning Monte Carlo ...\n\n")

			showCustomEntryDialog(
				"Input Desired number of grid elements",
				"max 120,000; 10,000 will produce 4 pi digits, 110,00 may get you 5 digits",
				func(input string) {
					if input != "" { // This if-else is part of the magic that allows us to dismiss a dialog and allow others to run after the dialog is canceled/dismissed.
						input = removeCommasAndPeriods(input) // ::: allow user to enter a number with a comma
						val, err := strconv.Atoi(input) // val is now an int and input is a string
						if err != nil {
							fmt.Println("Error converting input:", err)
							updateOutput1("Invalid input, using default 10,000 digits")
						} else if val <= 1 {
							updateOutput1("Input must be greater than 1, using default 10,000 digits")
						} else if val > 120000 {
							updateOutput1("Input must be less than 120,001 -- using default of 10,000 digits")
						} else {
							MontDigits = strconv.Itoa(val) // val here is a number, an int to be precise. So, we use strconv.Itoa to convert the int to a string and assign it to MontDigits. 
						}
						go func(done chan bool) { // ::: go func now takes an argument
							defer func() {       // ::: new defer func with global calculating flag set 
								calculating = false // this does not appear to work 
								updateOutput1("Calculation definitely finished; possibly aborted\n")
							}()
							Monty(updateOutput1, MontDigits, done) // ::: func < - - - - - - - - - - - - < -  NOT AMENABLE TO KILLING VIA A DONE CHANNEL 
							calculating = false
							for _, btn := range buttons1 {
								btn.Enable()
							}
						}(currentDone)
					} else {
						// dialog canceled 
						updateOutput1("Monte Carlo calculation canceled, make another selection")
						for _, btn := range buttons1 {
							btn.Enable()
						}
						calculating = false // ::: this is the trick to allow others to run after the dialog is canceled/dismissed.
					}
				},
			)
		},
	)
	/*
		.
		.
	*/
	
	GaussBtn1 := NewColoredButton(
	"Gauss-Legendre -- C F Gauss, refined by Adrien-Marie Legendre\n" +
		"π ≈ (aₙ + bₙ)² / (4 tₙ)\n" +
		"only manages to make 10 digits of Pi in well-over five minutes\n" +
		"an infinite series circa 1655    --- served here by Rick Woolley ---",
		color.RGBA{100, 255, 100, 215},
		
		func() {
			if calculating {
				return
			}
			calculating = true
			for _, btn := range buttons1 {
				btn.Disable()
			}
			for _, btn := range gaussBut { // Refer to the comments in the initial assignment and creation of archimedesBtn1
				calculating = true
				btn.Enable()
			}
			currentDone = make(chan bool) // ::: New channel per run
			updateOutput1("\nRunning Gauss...\n\n")
			go func(done chan bool) { // ::: go func now takes an argument
				defer func() {       // ::: new defer func with global calculating flag set 
					calculating = false // this does not appear to work 
					updateOutput1("Calculation definitely finished; possibly aborted\n")
				}()
				Gauss_Legendre(updateOutput1, done) // ::: func < - - - - - - - - - - - - - < -
				calculating = false
				for _, btn := range buttons1 {
					btn.Enable()
				}
			}(currentDone)
		},
	)
	/*
	.
	.
	 */
	
	CustomSeriesBtn1 := NewColoredButton(
	"Custom series -- I don't remember where it's from ... \n" +
		"but it is very quick -- 4s gets us 9 digits of sweet cherry Pi\n" +
		"π = (4/1) - (4/3) + (4/5) - (4/7) + (4/9) - (4/11) + (4/13) - (4/15) ...",
		color.RGBA{255, 120, 120, 215}, // Greenish for variety
		
		func() {
			// WallisParent <- "Dick"
			if calculating {
				return
			}
			calculating = true
			for _, btn := range buttons1 {
				btn.Disable()
			}
			for _, btn := range customBut { // Refer to the comments in the initial assignment and creation of archimedesBtn1
				calculating = true
				btn.Enable()
			}
			currentDone = make(chan bool) // ::: New channel per run
			updateOutput1("\nRunning Custom Series ...\n\n")
			go func(done chan bool) { // ::: go func now takes an argument
				defer func() {       // ::: new defer func with global calculating flag set 
					calculating = false // this does not appear to work 
					updateOutput1("Calculation definitely finished; possibly aborted\n")
				}()
				CustomSeries(updateOutput1, done) // ::: probably want to add a done channel to this one
				calculating = false
				for _, btn := range buttons1 {
					btn.Enable()
				}
			}(currentDone)
		},
	)
	/*
	.
	.
	 */
	
	GregoryLeibnizBtn1 := NewColoredButton(
	"Gregory-Leibniz -- runs 20sec -- serves up 10 digits of Pi\n" +
		"James Gregory 1638–1675  Gottfried Wilhelm Leibniz 1646-1716\n" +
		"π = 4 * ( 1 - 1/3 + 1/5 - 1/7 + 1/9 ...) ",
		color.RGBA{110, 110, 255, 185},
		
		func() {
			if calculating {
				return
			}
			calculating = true
			for _, btn := range buttons1 {
				btn.Disable()
			}
			for _, btn := range gottfieBut { // Refer to the comments in the initial assignment and creation of archimedesBtn1
				calculating = true
				btn.Enable()
			}
			currentDone = make(chan bool) // ::: New channel per run
			updateOutput1("\nRunning Gregory-Leibniz...\n\n")
			go func(done chan bool) { // ::: go func now takes an argument
				defer func() {       // ::: new defer func with global calculating flag set 
					calculating = false // this does not appear to work 
					updateOutput1("Calculation definitely finished; possibly aborted\n")
				}()
				GregoryLeibniz(updateOutput1, done) // ::: probably want to add a done channel to this one
				calculating = false
				for _, btn := range buttons1 {
					btn.Enable()
				}
			}(currentDone)
		},
	)
	/*
	.
	.
	 */
	
	// Eight buttons on home page, so eight kluges 
	archiBut = []*ColoredButton{archimedesBtn1} // All these are a trick/kluge used as bug preventions // to keep methods from being started or restarted in parallel (over-lapping) 
	walisBut = []*ColoredButton{JohnWallisBtn1} 
	spigotBut = []*ColoredButton{SpigotBtn1} 
	chudBut = []*ColoredButton{ChudnovskyBtn1} 
	montBut = []*ColoredButton{MontyBtn1} 
	gaussBut = []*ColoredButton{GaussBtn1}
	customBut = []*ColoredButton{CustomSeriesBtn1}
	gottfieBut = []*ColoredButton{GregoryLeibnizBtn1}
	
	// same eight again: 
	buttons1 = []*ColoredButton{archimedesBtn1, JohnWallisBtn1, SpigotBtn1, ChudnovskyBtn1, MontyBtn1, GaussBtn1, CustomSeriesBtn1, GregoryLeibnizBtn1,} // used only for range btn.Enable()

		// ::: page Lay-out
		content1 := container.NewVBox(widget.NewLabel("\nSelect a method to estimate π:\n"),
			container.NewGridWithColumns(4, archimedesBtn1, JohnWallisBtn1, SpigotBtn1,
				ChudnovskyBtn1, MontyBtn1, GaussBtn1, CustomSeriesBtn1, GregoryLeibnizBtn1, contentForScoreBoard),
			windowContent,
		)
/*
.
.
 */
	// ::: drop-down menus -- same for all windows  -  -  --  -  -  --  -  -  --  -  -  --  -  -  --  -  -  --  -  -  --  -  -  --  -  -  --  
	logFilesMenu := fyne.NewMenu("Log-Files",
		fyne.NewMenuItem("View Log 1", func() { dialog.ShowInformation("Log Files", "Viewing Log 1", window1) }),
		fyne.NewMenuItem("View Log 2", func() { dialog.ShowInformation("Log Files", "Viewing Log 2", window1) }),
	)
	additionalMethodsMenu := fyne.NewMenu("Other-Methods",
		fyne.NewMenuItem("Home-Page (Pi methods)", func() { window1.Show() }),
		fyne.NewMenuItem("Second-page of Pi methods", func() { createWindow2().Show() }), // it is only the Show method of createWindow2 that specifies/(calls for) action. 
		fyne.NewMenuItem("Odd Pi calculators", func() { createWindow3(myApp).Show() }),   // ... yet even these are inert until the final line of main() : window1.ShowAndRun()
		fyne.NewMenuItem("Misc Maths", func() { createWindow4(myApp).Show() }),
	)
	optionsMenu := fyne.NewMenu("Options",
		fyne.NewMenuItem("Begin the ScoreBoard of Pi", func() {
			
			// dialog.ShowInformation("ScoreBoard", "Use Abort in Menu\nPrior to dismissing with OK", window1)
			if calculating {
				fmt.Println("Calculation already in progress")
				return
			}
			calculating = true
			currentDone = make(chan bool)
			termsCount = 0

			go func(done chan bool) {
				defer func() {
					calculating = false
					terminalDisplay.SetText(fmt.Sprintf("Terminal Output:\n\nCalculation stopped.\nFinal Pi: %.11f\nTerms: %d", <-pichan, termsCount))
				}()

				pie := nifty_scoreBoardG(func(text string) {
					terminalDisplay.SetText(text)
				}, done)

				if pie != 0.0 {
					terminalDisplay.SetText(fmt.Sprintf("Terminal Output:\n\nComputed Value of Pi:\t\t%.11f\n# of Nilakantha Terms:\t\t%d", pie, termsCount))
				}
			}(currentDone)
		}),
		fyne.NewMenuItem("Abort any currently executing method", func() {
			if currentDone == nil {
				updateOutput1("\nNo active calculation to abort, no such currentDone channel exists\n")
				fmt.Println("No active calculation to abort, no such currentDone channel exists")
				return
			}
			select {
			case <-currentDone:
				updateOutput1("\nMenu select determined that currentDone-chan had already been closed; all Goroutines were PREVIOUSLY notified to terminate\n") // ::: via closed chan status 
				fmt.Println("Menu select determined that currentDone-chan had already been closed; all Goroutines were PREVIOUSLY notified to terminate")
			default:
				close(currentDone)
				updateOutput1("\nTermination signals were sent to all current processes that may be listening\n") // ::: ... by way of closed chan status 
				fmt.Println("Termination signals were sent to all current processes that may be listening")
			}
		}),
		fyne.NewMenuItem("Show the terminal -- Cmd+Tab to return", func() {
			err := openTerminal()
			if err != nil {
				fmt.Println(err)
				return
			}		}),
	)
	/* ::: more: 
	select { // select is a concurrency-specific channel-only construct used to handle multiple channel operations, see explanation in second comment-block below.
	// Check if the currentDone channel is already closed (chan receive [<-] succeeds on a closed chan (it receives/reads that the channel is closed, successfully) -- false is returned in the case of chan type bool)
		case <-currentDone: // chan syntax for receive on/from chan "currentDone"
			updateOutput1("\nMenu select determined that done-chan had already been closed; all Goroutines were PREVIOUSLY notified to terminate\n") // ::: via closed chan status 
						// fmt.Printf("\nMenu select-case determined that calculating is %t\n", calculating)
		default: // chan was open but empty, receive has "failed" (nothing to receive: "blocks"), case has "failed" (does not trigger), chan has blocked until a value is sent on the chan; default ensues
			close(currentDone) // "else" close the currentDone chan, which will be interpreted as a termination signal by all listening processes
			updateOutput1("\nTermination signals were sent to all current processes that may be listening\n") // ::: ... by way of closed chan status 
						// fmt.Printf("\nMenu select-default determined that calculating is %t\n", calculating)
		}
	/
		operation (<-ch) on a closed channel:
			Succeeds immediately (no blocking/waiting).
			Returns the zero value of the channel’s type (false for chan bool, 0 for chan int, "" for chan string, etc.).
		When you try <-ch on an empty, open channel, it doesn’t fail — it blocks. Blocking means the operation pauses (waits) until something is put into the pipe
		... but in the context of a select, waiting is not succeeding, hence the default case is run.
	/
	/
		Switch: Like picking a door based on a number you’re holding — door 1, 2, or 3 opens depending on your number. Your num matches no doors? You get the default door.
			vs
		Select: Like waiting at a row of mailboxes for a letter to arrive — you grab the first one you see, or immediately walk away if you see none (default).
	/
...
	 */

	mainMenu := fyne.NewMainMenu(logFilesMenu, additionalMethodsMenu, optionsMenu)
	window1.SetMainMenu(mainMenu) // does not actually put up the menus, no actual action is called for here !!
	
	// Apply window background to the entire content
	windowWithBackground := container.NewMax(bgwc, content1)
	
	window1.SetContent(windowWithBackground)
	
	window1.ShowAndRun() // ACTION !!!
}
