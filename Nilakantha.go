package main

import (
	"fmt"
	"math/big"
	"os"
	"strings"
	"time"
)

// @formatter:off

func NilakanthaBig(fyneFunc func(string), iters int, precision int, done chan bool) { // Changed signature ::: - -
var printThisThen string
var printThis []string
var lenOfPi int

	fyneFunc(fmt.Sprintf("\n... working ...\n"))

	if iters > 36111222 {
		fyneFunc(fmt.Sprintf("\n... working ... Nilakantha using big floats"))
	}
	if iters > 42000000 {
		fyneFunc(fmt.Sprintf("\n... werkin ..."))
	}
	if iters > 55111222 {
		fyneFunc(fmt.Sprintf("\n... working for a while ..."))
	}
	if iters > 69111222 {
		fyneFunc(fmt.Sprintf("\n... will be working for quite a while ..."))
	}
	if iters > 80111222 {
		fyneFunc(fmt.Sprintf("\n... a very long while ... working ...\n"))
	}

	start := time.Now()

	var iterBig int

	// big.Float "constants":

		twoBig := big.NewFloat(2)
		threeBig := big.NewFloat(3)
		fourBig := big.NewFloat(4)

	// big.Float variables:

		digitoneBig := new(big.Float)
		*digitoneBig = *twoBig
	
		digittwoBig := new(big.Float)
		*digittwoBig = *threeBig
	
		digitthreeBig := new(big.Float)
		*digitthreeBig = *fourBig
	
		sumBig := new(big.Float)
		nexttermBig := new(big.Float)

	// set precision to a user-specified value
		sumBig.SetPrec(uint(precision))
		twoBig.SetPrec(uint(precision))
		threeBig.SetPrec(uint(precision))
		fourBig.SetPrec(uint(precision))
		digitoneBig.SetPrec(uint(precision))
		digittwoBig.SetPrec(uint(precision))
		digitthreeBig.SetPrec(uint(precision))
		nexttermBig.SetPrec(uint(precision))

	// ::: calculate initial value  	
	sumBig.Add(threeBig, new(big.Float).Quo(fourBig, new(big.Float).Mul(digitoneBig, new(big.Float).Mul(digittwoBig, digitthreeBig))))

	fileHandleBig, err1prslc2c := os.OpenFile("big_pie_is_in_here.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file
	check(err1prslc2c)           // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt // It’s idiomatic to defer a Close immediately after opening a file.
	defer func(fileHandleBig *os.File) {
		err := fileHandleBig.Close()
		if err != nil {}
	}(fileHandleBig) 

	iterBig = 1
	for iterBig < iters {
		select {
		case <-done: // ::: here an attempt is made to read from the channel (a closed channel can be read from successfully; but what is read will be the null/zero value of the type of chan (0, false, "", 0.0, etc.)
			// in the case of this particular channel (which is of type bool) we get the value false from having received from the channel when it is already closed. 
			// ::: if the channel known by the moniker "done" is already closed, that/it is to be interpreted as the abort signal by all listening processes. 
			fmt.Println("Goroutine Nilakantha for-loop (1 of 1) is being terminated by select case finding the done channel to be already closed")
			return // Exit the goroutine
		default:
		
		/*
		  -- Nilakantha Somayaji -- on Mac-mini.local
		was run on: Sun Mar 23 21:08:37 2025
		100000000 was total Iterations; 512 was precision setting for the big.Float types
		Total run was 1m4.656298791s   25 verified digits   3.141592653589793238462643
		*/
		
		iterBig++
		
			// ::: Calculate: 
				digitoneBig.Add(digitoneBig, twoBig)
				digittwoBig.Add(digittwoBig, twoBig)
				digitthreeBig.Add(digitthreeBig, twoBig)
		
				nexttermBig.Quo(fourBig, new(big.Float).Mul(digitoneBig, new(big.Float).Mul(digittwoBig, digitthreeBig)))

				if iterBig%2 == 0 { // % is modulus operator
					sumBig.Sub(sumBig, nexttermBig)
				} else {
					sumBig.Add(sumBig, nexttermBig)
				}

		if iterBig == 20111222 {
			fyneFunc(fmt.Sprintf("\n ... doin some ... ")) // Send to channel
		}
		if iterBig == 36111222 {
			fyneFunc(fmt.Sprintf("\n ... werkin ... "))
		}
		if iterBig == 42000000 {
			fyneFunc(fmt.Sprintf("\n... still werkin ... Nilakantha Somayaji method using big.Float types \n  -- with some patience one can generate 31 correct digits of pi this way.\n"))
		}
		if iterBig == 55111222 {
			fyneFunc(fmt.Sprintf("\n... been working for a while ..."))
		}
		if iterBig == 69111222 {
			fyneFunc(fmt.Sprintf("\n... been working for quite a while ..."))
		}
		if iterBig == 80111222 {
			fyneFunc(fmt.Sprintf("\n... it's been a very long while ... but still working ..."))
		}
		if iterBig == 180111222 {
			fyneFunc(fmt.Sprintf("\n... it's been a very long while, 180,111,222 done, ... and still working ..."))
		}
		if iterBig == 280111222 {
			fyneFunc(fmt.Sprintf("\n... it's been a very long while, 280,111,222 done, ... and still working ..."))
		}
		if iterBig == 480111222 {
			fyneFunc(fmt.Sprintf("\n... it's been a very long while, 480,111,222 done, ... still working ..."))
		}
		if iterBig == 680111222 {
			fyneFunc(fmt.Sprintf("\n... it's been a very long while, 680,111,222 done, ...  working ..."))
		}
		if iterBig == 880111222 {
			fyneFunc(fmt.Sprintf("\n... it's been a very long while, done, 880,111,222, done ... still, working ..."))
		}
		if iterBig == 977111222 {
			fyneFunc(fmt.Sprintf("\n... it's been a very long while, 977,111,222 already ... why am I still working? ..."))
		}
		}
	} // end of the loop, the only calculating loop
	t := time.Now()
	elapsed := t.Sub(start)
	TotalRun := elapsed.String()
	
		// ::: bug hammer = do this just once; KISS
		printThis, lenOfPi = checkPiTo100(sumBig) // all local variables defined at the top of this function 
		printThisThen = strings.Join(printThis, "")

	if lenOfPi > 55000 { // if length of pi is > 55,000 digits we have something really big
		// print to ::: screen
			fyneFunc(fmt.Sprintf("\n\n\nWe have been tasked with making a lot of pie and it was sooo big it needed its own file ...\n"))
			fyneFunc(fmt.Sprintf("\n\n  After allowing this process to finish (you may have to continue prodding this thing along for a while) ... \n"))
			fyneFunc(fmt.Sprintf("... Go have a look in /.big_pie_is_in_here.txt to find all the digits of π you had requested. \n\n"))

		// print (log) to a special ::: file
			_, err2prslc2c := fmt.Fprintf(fileHandleBig, "\nThese are the %d verified digits we have calculated, dumped by rick  :: \n", lenOfPi)
				check(err2prslc2c)
	
			_, err8prslc2c := fmt.Fprint(fileHandleBig, printThisThen) // to a file
				check(err8prslc2c)

		err := fileHandleBig.Close()
		if err != nil {
			return
		}
	} else { 

		fyneFunc(fmt.Sprintf("\npi as calculated herein is: %s", printThisThen))

		floatIterBig := float64(iterBig)
		printableIterbigWithcommas := formatFloat64WithThousandSeparators(floatIterBig)

		fyneFunc(fmt.Sprintf("\n.... we have matched %d digits in %s iterations: ", lenOfPi, printableIterbigWithcommas))

		fileHandleNilakan, err1prslc2c := os.OpenFile("dataLog-From_Nilakantha_Method_lengthy_prints.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		check(err1prslc2c)         // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt  // It’s idiomatic to defer a Close immediately after opening a file.
		defer func(fileHandleNilakan *os.File) {
			err := fileHandleNilakan.Close()
				if err != nil {}
			}(fileHandleNilakan) 
		
		// print to ::: file
				_, err2prslc2c := fmt.Fprintf(fileHandleNilakan,
					"\n\nBelow rick are the %d verified digits we have calculated via Nilakantha using precision of %d and iterations of %d: \n", lenOfPi, precision, iterBig)
						check(err2prslc2c)

		fyneFunc(fmt.Sprintf("\nhey, rick, pi as calculated herein is: %s", printThisThen))
		
		_, err8prslc2c := fmt.Fprint(fileHandleNilakan, printThisThen) 
		check(err8prslc2c)

		err := fileHandleNilakan.Close()
		if err != nil {
			return
		}

		fileHandleDefault, err1prslc2d := os.OpenFile("dataLog-From_calculate-pi-and-friends.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file
		check(err1prslc2d)             // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
		defer fileHandleDefault.Close()        // It’s idiomatic to defer a Close immediately after opening a file.
		
		// to ::: file
			_, err2prslc2d := fmt.Fprintf(fileHandleDefault, "\nThese are the %d verified digits we have calculated  :: \n", lenOfPi)
				check(err2prslc2d)
			
		// to ::: screen
			fyneFunc(fmt.Sprintf("\n These are the %d verified digits we have calculated: \n", lenOfPi))

		_, err8prslc2c = fmt.Fprint(fileHandleBig, printThisThen) // to a file
			check(err8prslc2c)

				err = fileHandleDefault.Close()
				if err != nil {
					return
				}
	}
	fyneFunc(fmt.Sprintf("\n"))

	// Open a log file
		fileHandleDefault, err1 := os.OpenFile("dataLog-From_calculate-pi-and-friends.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file
			check(err1)                              // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
				defer fileHandleDefault.Close()     // It’s idiomatic to defer a Close immediately after opening a file.

	Hostname, _ := os.Hostname()
	current_time := time.Now()
	TotalRun = elapsed.String()

	// print to ::: file
		_, err0 := fmt.Fprintf(fileHandleDefault, "\n  -- Nilakantha Somayaji -- on %s \n", Hostname)
		check(err0)
		_, err6 := fmt.Fprint(fileHandleDefault, "was run on: ", current_time.Format(time.ANSIC), "\n")
		check(err6)
		_, err5 := fmt.Fprintf(fileHandleDefault, "%d was total Iterations; %d was precision setting for the big.Float types \n", iterBig, precision)
		check(err5)
		_, err7 := fmt.Fprintf(fileHandleDefault, "Total run was %s \n ", TotalRun)
		check(err7)
		_, err2prslc2da := fmt.Fprint(fileHandleDefault, "\nResults from running Nilakantha can be viewed in dataLog-From_Nilakantha_Method_lengthy_prints.txt\n")
		check(err2prslc2da)

			err := fileHandleDefault.Close()
			if err != nil {
				return
			}

	// print to ::: screen
	fyneFunc(fmt.Sprintf(" via Nilakantha with big floats. Written entirely by Richard Woolley\n"))
/*
	// ::: Prepare to exit the Gregory Nilakantha method functions
	calculating2 = false // Allow another method to be selected.
	for _, btn := range buttons2 { // ok to only Enable buttons2, because I expect to only ever execute this from window2
		btn.Enable() // ::: Enable
	}
 */

	// written entirely by Richard Woolley
} 
