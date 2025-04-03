package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

// @formatter:off

func TheSpigot(fyneFunc func(string), numberOfDigitsToCalc int, done chan bool) {

	fyneFunc(fmt.Sprintf("\n\n... from A trick I mooched off of GitHub ... and re-wrote\n\n"))
	fyneFunc(fmt.Sprintf("Spigot executed with a request for %d digits, and produced:\n\n", numberOfDigitsToCalc)) // ::: pi is then printed one char at a time in the loop below

	usingBigFloats = false

	SpigotCalculation(fyneFunc, numberOfDigitsToCalc, done)

}

var piWithInsertedDecimalPoint []string

func SpigotCalculation(fyneFunc func(string), numberOfDigits int, done chan bool) { // Rick's version does not return a string // called by the previous func
	start := time.Now()
	pi := "" // allocate a string var "pi" which will end up being pi sans the decimal point
	boxes := numberOfDigits * 10 / 3
	remainders := make([]int, boxes)
	for i := 0; i < boxes; i++ {
		select {
		case <-done: // ::: here an attempt is made to read from the channel (a closed channel can be read from successfully; but what is read will be the null/zero value of the type of chan (0, false, "", 0.0, etc.)
			// in the case of this particular channel (which is of type bool) we get the value false from having received from the channel when it is already closed. 
			// ::: if the channel known by the moniker "done" is already closed, that/it is to be interpreted as the abort signal by all listening processes. 
			fmt.Println("Goroutine Spigot-func for-loop (1 of 4) is being terminated by select case finding the done channel to be already closed")
			return // Exit the goroutine
		default:
		remainders[i] = 2
	}
	}
	digitsHeld := 0
	for i := 0; i < numberOfDigits; i++ { // ::: loop: ----------------------------------------------------- < < < < < 
		select {
		case <-done: // ::: here an attempt is made to read from the channel (a closed channel can be read from successfully; but what is read will be the null/zero value of the type of chan (0, false, "", 0.0, etc.)
			// in the case of this particular channel (which is of type bool) we get the value false from having received from the channel when it is already closed. 
			// ::: if the channel known by the moniker "done" is already closed, that/it is to be interpreted as the abort signal by all listening processes. 
			fmt.Println("Goroutine Spigot-func for-loop (2 of 4) is being terminated by select case finding the done channel to be already closed")
			return // Exit the goroutine
		default:
		carriedOver := 0
		sum := 0
		for j := boxes - 1; j >= 0; j-- {
			select {
			case <-done: // ::: here an attempt is made to read from the channel (a closed channel can be read from successfully; but what is read will be the null/zero value of the type of chan (0, false, "", 0.0, etc.)
				// in the case of this particular channel (which is of type bool) we get the value false from having received from the channel when it is already closed. 
				// ::: if the channel known by the moniker "done" is already closed, that/it is to be interpreted as the abort signal by all listening processes. 
				fmt.Println("Goroutine Spigot-func for-loop (3 of 4) is being terminated by select case finding the done channel to be already closed")
				return // Exit the goroutine
			default:
			remainders[j] *= 10
			sum = remainders[j] + carriedOver
			quotient := sum / (j*2 + 1)
			remainders[j] = sum % (j*2 + 1)
			carriedOver = quotient * j
		}
		}
		remainders[0] = sum % 10
		q := sum / 10
		switch q {
		case 9:
			digitsHeld++
		case 10:
			q = 0
			for k := 1; k <= digitsHeld; k++ {
				select {
				case <-done: // ::: here an attempt is made to read from the channel (a closed channel can be read from successfully; but what is read will be the null/zero value of the type of chan (0, false, "", 0.0, etc.)
					// in the case of this particular channel (which is of type bool) we get the value false from having received from the channel when it is already closed. 
					// ::: if the channel known by the moniker "done" is already closed, that/it is to be interpreted as the abort signal by all listening processes. 
					fmt.Println("Goroutine Spigot-func for-loop (4 of 4) is being terminated by select case finding the done channel to be already closed")
					return // Exit the goroutine
				default:
				replaced, _ := strconv.Atoi(pi[i-k : i-k+1])
				if replaced == 9 {
					replaced = 0
				} else {
					replaced++
				}
				pi = delChar(pi, i-k) // ::: delChar func 
				pi = pi[:i-k] + strconv.Itoa(replaced) + pi[i-k:]
			}
			}
			digitsHeld = 1
		default:
			digitsHeld = 1
		}
		// Rick's code : File prints
		if i == 1 {
			fyneFunc(fmt.Sprintf(".")) // ::: print the decimal between the three and the 1, i.e., 3.1 
			piWithInsertedDecimalPoint = append(piWithInsertedDecimalPoint, ".")
		} // insert the decimal point between the 3 and the 1

		fyneFunc(fmt.Sprintf(strconv.Itoa(q))) // ::: Rick's new method of displaying pi, one digit at a time

		piWithInsertedDecimalPoint = append(piWithInsertedDecimalPoint, strconv.Itoa(q))

		// ::: this is still a loop ... so if the calculation is a really long one we want to log the progress as it unfolds < - - - - v v v v v v v v v v v v v still the loop v v v v 
		if i == 50 || i == 250 || i == 450 || i == 1200 || i == 5000 || i == 10000 || i == 20000 || i == 40000 || i == 80000 || i == 140000 || i == 200000 {
			t := time.Now()
			elapsed := t.Sub(start)
			// log stats to a log file
			fileHandle, err1 := os.OpenFile("dataLog-From_calculate-pi-and-friends.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file
			check(err1)                                                                                                             // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
			defer fileHandle.Close()                                                                                                // It’s idiomatic to defer a Close immediately after opening a file.
			Hostname, _ := os.Hostname()
			_, err0 := fmt.Fprintf(fileHandle, "\n  -- partial Spigot in process -- on %s \n", Hostname)
			check(err0)
			current_time := time.Now()
			_, err6 := fmt.Fprint(fileHandle, "... running at: ", current_time.Format(time.ANSIC), "\n")
			check(err6)
			TotalRun := elapsed.String() // cast time durations to a String type for Fprintf "formatted print"
			_, err7 := fmt.Fprintf(fileHandle, "Runtime this far is %s \n", TotalRun)
			check(err7)
			_, err8 := fmt.Fprintf(fileHandle, "... while calculating Pi to %d digits, having completed %d digits\n", numberOfDigits, i)
			check(err8)
		}
		// end Rick's code
		pi += strconv.Itoa(q) // in orriginal method, entire string accumulator was being returned
		}
	} // - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -  - - - - - - - ::: ^ ^ ^ ^ ^ ^ ^ ^ ^  the loop ends here ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ 

	/*
		// ::: pi calculations finished and displayed
		fyneFunc(fmt.Sprintf("\nFinished.\n%s\n\n", pi))

		// ::: here comes our calculated pi with the decimal inserted; printed from an array that we accumulated for this purpose:
		fyneFunc(fmt.Sprintf("\nHere comes our calculated pi with the decimal inserted:\n\nSpigot has ended ... \n\n"))
		for _, character := range piWithInsertedDecimalPoint { // ok, because I will only execute this from window1
			fyneFunc(fmt.Sprintf("%s", character))
		}
	 */


	// Rick's code ::: File prints
		t := time.Now()
		elapsed := t.Sub(start)
		// log stats to a log ::: file
			fileHandle, err1 := os.OpenFile("dataLog-From_calculate-pi-and-friends.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file
			check(err1)                                                                                                             // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
			defer fileHandle.Close()                                                                                                // It’s idiomatic to defer a Close immediately after opening a file.
			Hostname, _ := os.Hostname()
			_, err0 := fmt.Fprintf(fileHandle, "\n  -- Spigot -- on %s \n", Hostname)
			check(err0)
			current_time := time.Now()
			_, err6 := fmt.Fprint(fileHandle, "was run on: ", current_time.Format(time.ANSIC), "\n")
			check(err6)
			TotalRun := elapsed.String() // cast time durations to a String type for Fprintf "formatted print"
			_, err7 := fmt.Fprintf(fileHandle, "Total run was %s \n", TotalRun)
			check(err7)
			_, err8 := fmt.Fprintf(fileHandle, "To calculate Pi to %d digits\n", numberOfDigits)
			check(err8)
		// end Rick's code
}

func delChar(s string, index int) string {
	tmp := []rune(s)
	return string(append(tmp[0:index], tmp[index+1:]...))
	// written largely by Richard Woolley
}
