package main

import (
	"fmt"
	"os"
	"time"
)

// @formatter:off

func GregoryLeibniz(fyneFunc func(string), done chan bool) {
	// Open a log file 
	fileHandle, err1 := os.OpenFile("dataLog-From_calculate-pi-and-friends.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file
	check(err1)                              // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
		defer func(fileHandle *os.File) {   // It’s idiomatic to defer a Close immediately after opening a file.
			err := fileHandle.Close()
			if err != nil {}
		}(fileHandle)

	usingBigFloats = false // π = (4/1) - (4/3) + (4/5) - (4/7) + (4/9) - (4/11) + (4/13) - (4/15) ...
		fyneFunc(fmt.Sprintf("\n\nYou selected Gregory-Leibniz formula  :  π = 4 * ( 1 - 1/3 + 1/5 - 1/7 + 1/9 ...) \n"))
		fyneFunc(fmt.Sprintf("   Infinitesimal calculus was developed independently in the late 17th century by Isaac Newton\n"))
		fyneFunc(fmt.Sprintf("    ... James Gregory, and Gottfried Wilhelm Leibniz\n"))
		fyneFunc(fmt.Sprintf("   4 Billion iterations will (initially) be executed ... \n\n"))
		fyneFunc(fmt.Sprintf(" ... working ...\n\n"))
	
	start := time.Now()

	var denom float64
	var sum float64
	denom = 3
	sum = 1 - (1 / denom)
	
	iterInt64 = 1   // global 
	iterFloat64 = 0 // global
	
	for iterInt64 < 4000000000 {
		select {
		case <-done: // ::: here an attempt is made to read from the channel (a closed channel can be read from successfully; but what is read will be the null/zero value of the type of chan (0, false, "", 0.0, etc.)
			// in the case of this particular channel (which is of type bool) we get the value false from having received from the channel when it is already closed. 
			// ::: if the channel known by the moniker "done" is already closed, that/it is to be interpreted as the abort signal by all listening processes. 
			fmt.Println("Goroutine Gregory-Leibniz for-loop (1 of 2) is being terminated by select case finding the done channel to be already closed")
			return // Exit the goroutine
		default:
		iterFloat64++
		iterInt64++
		
		denom = denom + 2
		
		if iterInt64%2 == 0 {
			sum = sum + 1/denom
		} else {
			sum = sum - 1/denom
		}
		
		π = 4 * sum // calculate ::: pi : π
		
			if iterInt64 == 100000000 {
			fyneFunc(fmt.Sprintf("... 100,000,000 completed iterations ...\n"))
			fyneFunc(fmt.Sprintf("   %0.7f was calculated by the Gottfried Wilhelm Leibniz formula\n\n", π))
			t := time.Now()
			elapsed := t.Sub(start)
			fyneFunc(fmt.Sprintf("  100,000,000 iterations in %s yields 8 digits of π\n\n", elapsed))
			}
			if iterInt64 == 200000000 {
			fyneFunc(fmt.Sprintf("... 200,000,000 gets another digit ...\n"))
			fyneFunc(fmt.Sprintf("   %0.9f was calculated by the Gottfried Wilhelm Leibniz formula\n\n", π))
			t := time.Now()
			elapsed := t.Sub(start)
			fyneFunc(fmt.Sprintf("  200,000,000 iterations in %s yields 9 digits of π\n\n", elapsed))
			}
			if iterInt64 == 400000000 {
			fyneFunc(fmt.Sprintf("... 400,000,000 iterations completed, still at nine ...\n"))
			fyneFunc(fmt.Sprintf("   %0.10f was calculated by the Gottfried Wilhelm Leibniz formula\n\n", π))
			t := time.Now()
			elapsed := t.Sub(start)
			fyneFunc(fmt.Sprintf("  400,000,000 iterations in %s yields 9 digits of π\n\n", elapsed))
			}
			if iterInt64 == 600000000 {
			fyneFunc(fmt.Sprintf("... 600,000,000 iterations, still at nine ...\n"))
			fyneFunc(fmt.Sprintf("   %0.5f was calculated by the Gottfried Wilhelm Leibniz formula\n\n", π))
			t := time.Now()
			elapsed := t.Sub(start)
			fyneFunc(fmt.Sprintf("  600,000,000 iterations in %s yields 9 digits of π\n\n", elapsed))
			}
			if iterInt64 == 1000000000 {
			fyneFunc(fmt.Sprintf("... 1 Billion iterations completed, still nine ...\n"))
			fyneFunc(fmt.Sprintf("   %0.5f was calculated by the Gottfried Wilhelm Leibniz formula\n\n", π))
			t := time.Now()
			elapsed := t.Sub(start)
			fyneFunc(fmt.Sprintf("  1,000,000,000 iterations in %s yields 9 digits of π\n\n", elapsed))
			}
			if iterInt64 == 2000000000 {
			fyneFunc(fmt.Sprintf("... 2 Billion, and still just nine ...\n"))
			fyneFunc(fmt.Sprintf("   %0.5f was calculated by the Gottfried Wilhelm Leibniz formula\n\n", π))
			t := time.Now()
			elapsed := t.Sub(start)
			fyneFunc(fmt.Sprintf("  2,000,000,000 iterations in %s yields 9 digits of π\n\n", elapsed))
			}
			if iterInt64 == 4000000000 { // ::: last one
				fyneFunc(fmt.Sprintf("\n... 4 Billion, gets us ten digits  ..."))
				fyneFunc(fmt.Sprintf("   %0.5f was calculated by the Gottfried Wilhelm Leibniz formula\n\n", π))
				t := time.Now()
				elapsed := t.Sub(start)
				fyneFunc(fmt.Sprintf("  4,000,000,000 iterations in %s yields 10 digits of π\n\n", elapsed))
				fyneFunc(fmt.Sprintf(" per the Gottfried Wilhelm Leibniz formula\n"))
		
				LinesPerIter = 14
				fyneFunc(fmt.Sprintf("at aprox %0.2f lines of code per iteration ...\n", LinesPerIter))
				LinesPerSecond = (LinesPerIter * iterFloat64) / elapsed.Seconds() // .Seconds() returns a float64
				fyneFunc(fmt.Sprintf("Aprox %.0f lines of code were executed per second \n", LinesPerSecond))
				
				// store results in a log file 
					Hostname, _ := os.Hostname()
					current_time := time.Now()
					TotalRun := elapsed.String()   // cast time duration to a String type for Fprintf "formatted print"
				
					// to ::: file
						_, err0 := fmt.Fprintf(fileHandle, "\n  -- Gottfried Wilhelm Leibniz --  on %s \n", Hostname)
							check(err0)
						_, err6 := fmt.Fprint(fileHandle, "was run on: ", current_time.Format(time.ANSIC), "\n")
							check(err6)
						_, err2 := fmt.Fprintf(fileHandle, "%.0f was Lines/Second  \n", LinesPerSecond)
							check(err2)
						_, err4 := fmt.Fprintf(fileHandle, "%e was Iterations/Seconds  \n", iterFloat64/elapsed.Seconds())
							check(err4)
						_, err5 := fmt.Fprintf(fileHandle, "%e was total Iterations  \n", iterFloat64)
							check(err5)
						_, err7 := fmt.Fprintf(fileHandle, "Total runTime was %s \n", TotalRun) // add total runtime of this calculation
							check(err7)
			} // end of last if
			}
	} // end of first for loop
	

		// print to ::: file
			fyneFunc(fmt.Sprintf( "\n\nWe continue the Gottfried Wilhelm Leibniz formula  :  π = 4 * ( 1 - 1/3 + 1/5 - 1/7 + 1/9 ... \n"))
			fyneFunc(fmt.Sprintf("    π = 3 + 4/(2*3*4) - 4/(4*5*6) + 4/(6*7*8) - 4/(8*9*10) + 4/(10*11*12) ..."))
			
			fyneFunc(fmt.Sprintf("   Infinitesimal calculus was developed independently in the late 17th century by Isaac Newton"))
			fyneFunc(fmt.Sprintf("   and Gottfried Wilhelm Leibniz\n\n"))
			fyneFunc(fmt.Sprintf("   9 billion iterations will be executed \n\n   ... working ...\n"))

		start = time.Now()

		for iterInt64 < 9000000000 {
			select {
			case <-done: // ::: here an attempt is made to read from the channel (a closed channel can be read from successfully; but what is read will be the null/zero value of the type of chan (0, false, "", 0.0, etc.)
				// in the case of this particular channel (which is of type bool) we get the value false from having received from the channel when it is already closed. 
				// ::: if the channel known by the moniker "done" is already closed, that/it is to be interpreted as the abort signal by all listening processes. 
				fmt.Println("Goroutine Gregory-Leibniz for-loop (2 of 2) is being terminated by select case finding the done channel to be already closed")
				return // Exit the goroutine
			default:
			iterFloat64++
			iterInt64++
			denom = denom + 2
			if iterInt64%2 == 0 {
			sum = sum + 1/denom
			} else {
			sum = sum - 1/denom
			}
			π = 4 * sum
			
			if iterInt64 == 6000000000 {
			fyneFunc(fmt.Sprintf("... 6 Billion completed ... \n"))
			fyneFunc(fmt.Sprintf("   %0.13f was calculated by the Gottfried Wilhelm Leibniz formula", π))
			t := time.Now()
			elapsed := t.Sub(start)
			fyneFunc(fmt.Sprintf("  6,000,000,000 iterations in %s still yields 10 digits of π\n", elapsed))
			fyneFunc(fmt.Sprintf( "\n  ... working ...\n\n"))
			}
			if iterInt64 == 8000000000 {
			fyneFunc(fmt.Sprintf("... 8 Billion completed. still ten ...\n"))
			fyneFunc(fmt.Sprintf("   %0.13f was calculated by the Gottfried Wilhelm Leibniz formula", π))
			t := time.Now()
			elapsed := t.Sub(start)
			fyneFunc(fmt.Sprintf("  8,000,000,000 iterations in %s still yields 10 digits of π\n", elapsed))
			fyneFunc(fmt.Sprintf( "\n ... working ...\n"))
			}
			if iterInt64 == 9000000000 {
			fyneFunc(fmt.Sprintf("   %0.13f was calculated by the Gottfried Wilhelm Leibniz formula", π))
			// fyneFunc(fmt.Sprintf("   ", iter)
			t := time.Now()
			elapsed := t.Sub(start)
			fyneFunc(fmt.Sprintf("\n... 9B iterations in %s, but to get 10 digits we only needed 4B iterations\n\n", elapsed))
			fyneFunc(fmt.Sprintf(" per  --  the Gottfried Wilhelm Leibniz formula\n"))
			
				t = time.Now()
				elapsed = t.Sub(start)
				TotalRun := elapsed.String()          // cast time duration to a String type for Fprintf "formatted print"
			
			LinesPerIter = 14 // estimate
			// print to ::: screen
				fyneFunc(fmt.Sprintf("at aprox %0.2f lines of code per iteration ...\n\n", LinesPerIter))
			
					fyneFunc(fmt.Sprintf("\n%e was Iterations/Seconds", iterFloat64/elapsed.Seconds()))
					fyneFunc(fmt.Sprintf("   Infinitesimal calculus was developed independently in the late 17th century by Isaac Newton\n"))
					fyneFunc(fmt.Sprintf("   and Gottfried Wilhelm Leibniz. This implementaion was done entirely by Richard Woolley"))
					fyneFunc(fmt.Sprintf(""))
					fyneFunc(fmt.Sprintf(""))
			
				LinesPerSecond = (LinesPerIter * iterFloat64) / elapsed.Seconds() // .Seconds() returns a float64
			// to ::: screen
				fyneFunc(fmt.Sprintf("Aprox %.0f lines of code were executed per second \n", LinesPerSecond))
				fyneFunc(fmt.Sprintf("Total runTime was %s \n", TotalRun)) // add total runtime of this calculation
				
					Hostname, _ := os.Hostname()
					current_time := time.Now()
			
			// print to ::: file
				_, err0 := fmt.Fprintf(fileHandle, "\n  -- Gottfried Wilhelm Leibniz (cont.) -- on %s \n", Hostname)
					check(err0)
				_, err6 := fmt.Fprint(fileHandle, "was run on: ", current_time.Format(time.ANSIC), "\n")
					check(err6)
				_, err2 := fmt.Fprintf(fileHandle, "%.0f was Lines/Second  \n", LinesPerSecond)
					check(err2)
				_, err4 := fmt.Fprintf(fileHandle, "%e was Iterations/Seconds  \n", iterFloat64/elapsed.Seconds())
					check(err4)
				_, err5 := fmt.Fprintf(fileHandle, "%e was total Iterations  \n", iterFloat64)
					check(err5)
			
				_, err7 := fmt.Fprintf(fileHandle, "Total runTime was %s \n", TotalRun) // add total runtime of this calculation
					check(err7)
			}
			}
		}
		
	// ::: Prepare to exit the Gottfried method function
	calculating = false // Allow another method to be selected.
	for _, btn := range buttons1 { // ok to only Enable buttons1, because I expect to only ever execute this from window1
		btn.Enable() // ::: Enable
	}
} // written entirely by Richard Woolley
