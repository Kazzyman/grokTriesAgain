package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

// Adapted nifty_scoreBoard for Fyne GUI
func nifty_scoreBoardG(fyneFunc func(string), done chan bool) float64 {
	usingBigFloats = false // Unused in this version, kept for compatibility
	ticker := time.NewTicker(time.Millisecond * 108)

	// Launch Pi calculation
	go pi_nfG(5000, done)

	// Update scoreboard periodically
	go func() {
		for range ticker.C {
			select {
			case <-done:
				ticker.Stop()
				return
			default:
				printCalculationSummaryG(fyneFunc, done)
			}
		}
	}()

	for {
		select {
		case <-computationDone:
			ticker.Stop()
			fmt.Println("Program done calculating Pi.")
			return <-pichan
		case <-done:
			ticker.Stop()
			fyneFunc("Terminal Output:\n\nAborted by user.")
			fmt.Println("Scoreboard aborted via done channel")
			return 0.0 // Indicate abort
		}
	}
}

// Updated printCalculationSummary for Fyne
func printCalculationSummaryG(fyneFunc func(string), done chan bool) {
	select {
	case piValue := <-pichan:
		output := fmt.Sprintf("Terminal Output:\n\nComputed Value of Pi:\t\t%.11f\n# of Nilakantha Terms:\t\t%d", piValue, termsCount)
		fyneFunc(output)
	case <-done:
		// Do nothing, let the abort case handle it
	}
}

// Nilakantha series calculation
func pi_nfG(n int, done chan bool) float64 {
	ch := make(chan float64)
	f := 3.0

	for k := 1; k <= n; k++ {
		select {
		case <-done:
			fmt.Println("pi_nf aborted")
			return f
		default:
			go nilakanthaTermG(ch, float64(k))
		}
	}

	for k := 1; k <= n; k++ {
		select {
		case <-done:
			fmt.Println("pi_nf aborted during summation")
			return f
		default:
			termsCount++
			f += <-ch
			pichan <- f
		}
	}

	computationDone <- true
	return f
}

// Nilakantha term calculation
func nilakanthaTermG(ch chan float64, k float64) {
	j := 2 * k
	if int64(k)%2 == 1 {
		ch <- 4.0 / (j * (j + 1) * (j + 2))
	} else {
		ch <- -4.0 / (j * (j + 1) * (j + 2))
	}
}

/*
.
.
.
.
. below are the originals for comparison and documentation
.
.
.
.
.
*/

// case 40: -- AMFnifty_scoreBoardA
// A concurrent computation of pi using Nilakantha's formula.
// by Diego Brener diegosilva13 on Github
// ******* nifty scoreboard ***********************************
func nifty_scoreBoard(fyneFunc func(string), done chan bool) float64 { // case 40:
	usingBigFloats = false
	// We use a ticker to specify the interval to update the values on the scoreboard
	ticker := time.NewTicker(time.Millisecond * 108)

	// If the user wants to prematurely break out of the program by issuing a Ctrl+C, we tell the
	// ... signal package to notify us over this interrupt channel
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	go pi_nf(5000, done)

	// This anonymous function is responsible for updating the scoreboard
	// as per the interval specified by the ticker
	go func() {
		for range ticker.C {
			printCalculationSummary(updateOutput1, done)
		}
	}()

	for {
		select {

		// Handle the case when the computation has ended, we can
		// end the program (exit out of this infinite loop)
		case <-computationDone:
			ticker.Stop()
			fmt.Println("Program done calculating Pi.")
			os.Exit(0)
		case <-done: // ::: here an attempt is made to read from the channel (a closed channel can be read from successfully; but what is read will be the null/zero value of the type of chan (0, false, "", 0.0, etc.)
			// in the case of this particular channel (which is of type bool) we get the value false from having received from the channel when it is already closed. 
			// ::: if the channel known by the moniker "done" is already closed, that/it is to be interpreted as the abort signal by all listening processes. 
			fmt.Println("Goroutine Monty for-loop (1 of 2) is being terminated by select case finding the done channel to be already closed")
			return 0.0 // Exit the goroutine
		// If the user interrupts the program (Ctrl+C) we end the
		// program (exit out of this infinite loop)
		case <-interrupt:
			ticker.Stop()
			fmt.Println("Program interrupted by the user.")
			return <-pichan
		}
	}
}

// We want to create a virtual scoreboard where we can simutaneously show the current value of Pi and
// ... also how many Nilakantha terms we have calculated.

// We can create the virtual scoreboard by using some ANSI escape codes
// Here's a wikipedia article giving you the complete list of ANSI escape
// codes: https://en.wikipedia.org/wiki/ANSI_escape_code

// convenience globals:
const ANSIClearScreenSequence = "\033[H\033[2J"
const ANSIFirstSlotScreenSequence = "\033[2;0H"
const ANSISecondSlotScreenSequence = "\033[3;0H"

// The channel used to update the current value of pi on the scoreboard
var pichan chan float64 = make(chan float64)

// The channel that we use to indicate that the program can exit
var computationDone chan bool = make(chan bool, 1)

// Number of Nilakantha terms for the scoreboard
var termsCount int

// The following function serves as our virtual scoreboard to show the current
// computed value of Pi using Nilakantha's formula

func printCalculationSummary(fyneFunc func(string), done chan bool) {

	fmt.Print(ANSIClearScreenSequence)
	fmt.Println(ANSIFirstSlotScreenSequence, "Computed Value of Pi:\t\t", <-pichan)
	fmt.Println(ANSISecondSlotScreenSequence, "# of Nilakantha Terms:\t\t", termsCount)
}

// We are going to use Nilakantha's formula from the Tantrasamgraha (which is more than 500 years old)
// ... to compute the value of Pi to several decimal points
func pi_nf(n int, done chan bool) float64 {
	ch := make(chan float64)
	f := 3.0
	// The k variable is considered to be the current step
	for k := 1; k <= n; k++ {
		select {
		case <-done: // ::: here an attempt is made to read from the channel (a closed channel can be read from successfully; but what is read will be the null/zero value of the type of chan (0, false, "", 0.0, etc.)
			// in the case of this particular channel (which is of type bool) we get the value false from having received from the channel when it is already closed. 
			// ::: if the channel known by the moniker "done" is already closed, that/it is to be interpreted as the abort signal by all listening processes. 
			fmt.Println("Goroutine Monty for-loop (1 of 2) is being terminated by select case finding the done channel to be already closed")
			return f // Exit the goroutine
		default:
			// Each Nilakantha term is calculated in its own goroutine
			go nilakanthaTerm(ch, float64(k))
		}
	}
	// f := 3.0

	// We sum up the calculated Nilakantha terms for n steps
	for k := 1; k <= n; k++ {
		select {
		case <-done: // ::: here an attempt is made to read from the channel (a closed channel can be read from successfully; but what is read will be the null/zero value of the type of chan (0, false, "", 0.0, etc.)
			// in the case of this particular channel (which is of type bool) we get the value false from having received from the channel when it is already closed. 
			// ::: if the channel known by the moniker "done" is already closed, that/it is to be interpreted as the abort signal by all listening processes. 
			fmt.Println("Goroutine Monty for-loop (1 of 2) is being terminated by select case finding the done channel to be already closed")
			return f // Exit the goroutine
		default:
			termsCount++
			f += <-ch
			pichan <- f
		}
	}

	// We notify that the computation is done over the channel
	computationDone <- true
	return f
}

// This function gives us the nilakanthaTerm for the kth step
func nilakanthaTerm(ch chan float64, k float64) {
	j := 2 * k
	if int64(k)%2 == 1 {
		ch <- 4.0 / (j * (j + 1) * (j + 2))
	} else {
		ch <- -4.0 / (j * (j + 1) * (j + 2))
	}
	// adapted by Richard Woolley
} // end of nifty_scoreBoard() set
