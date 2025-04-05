package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"sync"
)

// @formatter:off

// a struct is like a blueprint for an object that can hold fields of ralated data.
type TrafficManager struct {
	workPiece int
	calculating bool
	radical int
	// output is a pointer to a widget.Label. The asterisk (*) means it’s a pointer (a reference to memory).
	output *widget.Label // here * denotes/declares type only -- it is NOT the dereferencing operator in this sort of context!!
	stop chan bool   // stop is a channel of type bool. Channels are a Go feature for communication between goroutines (concurrent tasks).
	// mu is a mutex (short for mutual exclusion) from the "sync" package. It’s used to prevent multiple goroutines from accessing
	// the struct’s fields at the same time, ensuring thread safety.
	mu sync.Mutex   // Mutex for thread-safe access to fields. ::: use it via:   m.mu.Lock()   &   defer m.mu.Unlock()
}

/*
The following is termed a Constructor Function: when called, it returns a new TrafficManager instance. Meaning, an object/instance of type TrafficManager. 
I like to think that the return type of *TrafficManager kind of marks the transition into setting up for some OOP-style methods to be created next. But it 
is important to realize that Go isn’t a traditional object-oriented language — it doesn’t have formal constructors, classes or inheritance as actual language 
features or constraints. Returning a pointer to a struct (*TrafficManager) enables methods on that struct (via receivers), which mimics some OOP behavior 
(encapsulation, methods), but Go is really just about composition and interfaces -- not classic OOP --  which is why we love it. Returning *TrafficManager 
merely allows for method attachment, supporting object-like behavior in Go’s non-traditional OOP approach.
    This func takes a pointer to a widget.Label as an argument and returns a pointer to a TrafficManager structure: a custom type defined earlier. 
The *, in this context, is NOT the dereferencing operator. Here both widget.Label and TrafficManager are strictly value types (Go doesn’t have “reference 
types” in the same sense as languages like Java — rather, in Go, by obtaining a pointer we explicitly create a references to the data stored in a variable --
or (more succinctly) we explicitly create a references to a value.
    The *'s in the signature are only telling the compiler that the type of the argument must be a pointer to a widget.Label type and that the type of 
the return value must be a pointer to an instance of a TrafficManager type. 
    Prefacing the * to a var or an expression within the function body does resolve to the memory address of the object being referred to; but not in the 
context of the signature wherein only declarations are being made. Conversely, prefixing & to a regular non-reference type var or expression (withing the body 
of a func) resolves to the memory address of the named object: thereby obtaining a pointer (a reference) to said object. If at a later point we needed to 
access the values stored at &TrafficManager, and assuming 'TrafficManagerR := &TrafficManager'; we could then say: 'x := *TrafficManagerR' thereby dereferencing 
the pointer TrafficManagerR and obtaining the value referred to by way of our new variable x. 
*/

func NewTrafficManager(output *widget.Label) *TrafficManager {
	// "&TrafficManager{...}" allocates a new instance and returns a pointer to same.
	// Curly braces "{}" enclose literal field initializations which must mirror the fields of the referenced type.
	return &TrafficManager{
		// In this context, Go uses colons as separators for "key:value" pairs. Here "key" is the field name and "value" is its initial value.
		// Commas are required after each pair in a multi-pair struct literal; and also after the last one, if the pairs are formatted as one per line.
		workPiece:   0,          // Initializes workPiece to 0
		calculating: false,      // Sets calculating to false
		radical:     2,          // Assigns 2 to radical
		output:      output,     // Sets output to the labeled pointer 'output' per this function's signature. 
		stop:        make(chan bool, 1), // Initializes stop as a buffered channel. If the ', 1' had been omitted we would be making an unbuffered chan
		// ::: not so sure it is a good idea for this to be buffered. 
	}
}

// SetWorkPiece is a method on the TrafficManager type. It sets the workPiece field to a given value.
// The "(m *TrafficManager)" part means this is a method tied to a TrafficManager pointer (receiver).
func (m *TrafficManager) SetWorkPiece(val int) {
	// Locks the mutex to ensure only one goroutine can modify the struct at a time (thread safety).
	m.mu.Lock()
	// "defer" ensures Unlock() is called when the function exits, even if it panics. This unlocks the mutex.
	defer m.mu.Unlock()
	// Sets the workPiece field to the provided value.
	m.workPiece = val
}

// SetRadical is similar to SetWorkPiece but updates the radical field.
// Methods like this are called "setters" because they modify a field.
func (m *TrafficManager) SetRadical(val int) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.radical = val
}

// Reset is a method that resets some fields to their initial states.
func (m *TrafficManager) Reset() {
	m.mu.Lock()
	defer m.mu.Unlock()
	// Resets workPiece to 0, like starting fresh. ::: bad grok
	// m.workPiece = 0
	// Sets calculating to false, indicating no active computation.
	m.calculating = false
}

// IsCalculating is a "getter" method that returns the current value of calculating.
// It’s a way to safely check the state from outside the struct.
func (m *TrafficManager) IsCalculating() bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	// Returns the value of calculating (true or false).
	return m.calculating
}

// SetCalculating is a setter for the calculating field.
func (m *TrafficManager) SetCalculating(val bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.calculating = val
}

// UpdateOutput updates the text displayed in the GUI label (output field).
// It takes a string argument to set as the new text.
func (m *TrafficManager) UpdateOutput(text string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	// Sets the text of the widget.Label to the provided string.
	m.output.SetText(text)
	// Refreshes the GUI canvas where the label is displayed.
	// "fyne.CurrentApp().Driver().CanvasForObject(m.output)" gets the canvas (drawing area) for the label,
	// and Refresh redraws it to show the updated text.
	fyne.CurrentApp().Driver().CanvasForObject(m.output).Refresh(m.output)
}

// SetOutput changes the output field to a new widget.Label pointer.
// This might be used to redirect output to a different GUI element.
func (m *TrafficManager) SetOutput(output *widget.Label) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.output = output
}

// GetWorkPiece is a getter for the workPiece field.
func (m *TrafficManager) GetWorkPiece() int {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.workPiece
}

// GetRadical is a getter for the radical field.
func (m *TrafficManager) GetRadical() int {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.radical
}

// Stop sends a signal to the stop channel to halt some process.
// It only sends the signal if calculating is true (i.e., something is running).
func (m *TrafficManager) Stop() {
	m.mu.Lock()
	defer m.mu.Unlock()
	// Checks if a computation is active.
	if m.calculating {
		// Sends true to the stop channel. Since it’s buffered (capacity 1), this won’t block unless the buffer is full. ::: careful 
		m.stop <- true
	}
}

// ShouldStop checks if a stop signal has been received.
// It’s a non-blocking way to see if the process should terminate.
func (m *TrafficManager) ShouldStop() bool {
	// "select" is used to handle channel operations. It tries multiple cases and runs the first one that’s ready.
	select {
	// If a value can be read from the stop channel, it means stop was signaled, so return true.
	case <-m.stop:
		return true
	// If no value is ready in the stop channel, the "default" case runs, returning false.
	default:
		return false
	}
}
