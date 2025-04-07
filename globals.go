package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

// @formatter:off

// var skip_redoing_loop int // ::: the existence of this old var may be a clue to an older efficiency tactic, now sorely missing?

// from main.go
var (
	// Build objects to set background colors: bgsc for scroll area (light green), bgwc for entire window (light blue), layered via NewMax.
	// Build objects to set background colors for scroll-area1 (bgsc), and the entire window1 (bgwc)
	bgsc = canvas.NewRectangle(color.NRGBA{R: 150, G: 180, B: 160, A: 240}) // Light green
	bgwc = canvas.NewRectangle(color.NRGBA{R: 110, G: 160, B: 255, A: 150}) // Light blue, lower number for A: means less opaque, or more transparent

	pie float64

	// Create scrollContainer1 to display outputLabel1 with initial user prompt.
	// Build/define scrollContainer1 as containing outputLabel1 with its initial greeting message
	outputLabel1 = widget.NewLabel("\nSelect one of the brightly-colored panels to estimate π via featured method...\n\n")
	scrollContainer1 = container.NewVScroll(outputLabel1)

	// Create app and window1 which is an extension of myApp
	// Initialize the Fyne app (myApp) and create window1 as its main window. Technically not a case of "extension"
	myApp = app.New()
	window1 = myApp.NewWindow("Rick's Pi calculation Demo, set #1")
	currentDone    chan bool 
)

// from window2.go
var (
	bgsc2 = canvas.NewRectangle(color.NRGBA{R: 130, G: 160, B: 250, A: 140}) // Light blue // was: 130, 160, 250, 160 ::: - -
	bgwc2 = canvas.NewRectangle(color.NRGBA{R: 110, G: 255, B: 160, A: 150}) // Light green ::: - -

	outputLabel2 = widget.NewLabel("Classic Pi calculators, make a selection") // ::: - -
	scrollContainer2 = container.NewScroll(outputLabel2) // ::: - -
	window2 = myApp.NewWindow("Rick's Pi calculation Demo, set #2") // ::: - -
)

// from roots
var (
	pairsSlice []Pairs // a slice of two-element structs, i.e., pairs 
	mathSqrtCheat            float64
	mathCbrtCheat            float64
	mgr             = NewTrafficManager(outputLabel2) // ::: - -
)
// Pairs A struct to contain two related whole numbers: an identity product (perfect square or cube), e.g. 49; and its root, which in that case would be 7 
type Pairs struct {
	product int
	root int
}


var calculating bool

// But
var (
	archiBut []*ColoredButton
	walisBut []*ColoredButton
	chudBut []*ColoredButton
	spigotBut []*ColoredButton
	montBut []*ColoredButton
	gaussBut []*ColoredButton
	customBut []*ColoredButton
	gottfieBut []*ColoredButton
)

// But2
var (
	chudBut2 []*ColoredButton
	BPPbut2 []*ColoredButton
	nilaBut2 []*ColoredButton
	scoreBut2 []*ColoredButton
	rootBut2 []*ColoredButton
	bbpMaxBut2 []*ColoredButton
)

// buttons
var (
	buttons1 []*ColoredButton // Change to ColoredButton // ::: - -
	buttons2 []*ColoredButton // Change to ColoredButton // ::: - -
	buttons3 []*ColoredButton // Change to ColoredButton
	buttons4 []*ColoredButton // Change to ColoredButton
)

var copyOfLastPosition int // ::: - -

// convenience globals:
var usingBigFloats = false // a variable of type bool which is passed by many funcs to print Result Stats Long() // ::: - -

var iterationsForMonte16i int
var iterationsForMonte16j int
var iterationsForMonteTotal int

var four float64 // is initialized to 4 where needed // ::: - -
var π float64    // a var can be any character, as in this Pi symbol/character // ::: - -
var LinesPerSecond float64 // ::: - -
var LinesPerIter float64 // ::: - -
var iterInt64 int64     // to be used primarily in selections which require modulus calculations // ::: - -
var iterFloat64 float64 // to be used in selections which do not require modulus calculations // ::: - -

// The following, are used in multiple functions of roots
var (
	sortedResults = []Results{} // sortedResults is an array of type Results as defined at the top of this file // ::: - -
	precisionOfRoot int    // this being global means we do not need to pass it in to the read func            // ::: - -
	Tim_win float64      // Time Window // ::: - -
	diffOfLarger int                   // ::: - -
	diffOfSmaller int                 // ::: - -
)

