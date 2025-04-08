package main

import (
	"fmt"
	"math"
	"math/big"
	"os"
	"strings"
	"time"
)

// @formatter:off

// Chudnovsky method, based on https://arxiv.org/pdf/1809.00533.pdf
/*
	   The Chudnovsky algorithm is an incredibly-fast algorithm for calculating the digits of pi. It was developed by Gregory Chudnovsky and his
	brother David Chudnovsky in the 1980s. It is more efficient than other algorithms and is based on the theory of modular equations. It has
	been used to calculate pi to over 62 trillion digits.
*/
//  Using this procedure, calculating 1,000,000 digits requires 70516 loops, per the run on:
//  Sun May  7 08:50:23 2023
//  Total run was 8h4m39.7847064s
// AND, THAT CALCULATION WAS INDEPENDENTLY VERIFIED !!!!!!!!!!!

func chudnovskyBig(digits int, done chan bool) { // ::: - -

	updateOutput1(fmt.Sprintf("\n... working ...\n"))
	usingBigFloats = true
	var loops int
	start := time.Now() // start will be passed, and then passed back, in order to be compared with end time t

	pi := new(big.Float)

	// ::: calcPi  <---- runs from here: v v v v v v v  
	loops, pi, start = calcPi(float64(digits), start, done)
	// ::: calcPi ----- ^ ^ ^ 

	
	// The following runs ::: after calcPi 
	updateOutput1(fmt.Sprintf("\n loops were: %d, and digits requested was: %d \n", loops, digits))

	updateOutput1(fmt.Sprintf("\n 	The Chudnovsky algorithm is an incredibly-fast algorithm for calculating the digits of pi. It was developed by Gregory Chudnovsky and his "))
	updateOutput1(fmt.Sprintf("brother David Chudnovsky in the 1980s. It is more efficient than other algorithms and is based on the theory of modular equations. It has been "))
	updateOutput1(fmt.Sprintf("used to calculate pi to over 62 trillion digits.\n\n"))

	file1 := "/Users/quasar/grokTriesAgain/big_pie_is_in_here.txt" // Replace with your first file path
	file2 := "/Users/quasar/grokTriesAgain/piOneMil.txt" // Replace with your second file path

	count, err := compareFiles1(file1, file2)
	if err != nil {
		fmt.Println("Error:", err)
	}
	updateOutput1(fmt.Sprintf("\n\nMatched %d characters in sequence from the start.\n", count))
	
	// determine elapsed timme:
	t := time.Now()
	elapsed := t.Sub(start)
	// The following print section is conditional upon some time having elapsed: at least one second. In this way we avoid logging to a file the smallest of run times. 
	if int(elapsed.Seconds()) != 0 { // ::: Note, that if runtime is less than one second this will be 0 : if, as a whole int, elapsed seconds is not zero. 
		// obtain file handle
			fileHandle, err1 := os.OpenFile("dataLog-From_calculate-pi-and-friends.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file
				check(err1)                   // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
					defer fileHandle.Close()     // It’s idiomatic to defer a Close immediately after opening a file.

		// print to ::: file			
			Hostname, _ := os.Hostname()
			current_time := time.Now()
			TotalRun := elapsed.String() // cast time durations to a String type for Fprintf "formatted print"
			
			_, err0 := fmt.Fprintf(fileHandle, "\n  --  pi-via-chudnovsky  --  on %s \n", Hostname)
				check(err0)
			_, err6 := fmt.Fprint(fileHandle, "was run on: ", current_time.Format(time.ANSIC), "\n")
				check(err6)
					// the whole pi would be printed to the datalog file on the line below
					// _ , err8 := fmt.Fprintf(fileHandle, "pi was %1.[1]*[2]f \n", digits, pi)
					//    check(err8)
					// ... after printing the whole pi, some nice stats are appended to the file's log entry
			_, err7 := fmt.Fprintf(fileHandle, "Total run was %s, and digits requested was %d , and at 80f, pi: %0.80f\n ", TotalRun, digits, pi)
				check(err7)
	}
}
/*
.
.
.
.
 */
// calculate Pi for n number of digits
func calcPi(digits float64, start time.Time, done chan bool) (int, *big.Float, time.Time) {
runeToPrint := `
	/**
	 *   This is an implementation for https://en.wikipedia.org/wiki/Chudnovsky_algorithm
	 *   "It can be improved using binary splitting http://numbers.computation.free.fr/Constants/Algorithms/splitting.html
	 *   if we were to split it into two independent parts and simplify the formula." For more details, visit:
	 *             https://www.craig-wood.com/nick/articles/pi-chudnovsky
	 */
`
	updateOutput1(fmt.Sprintf(runeToPrint))

	usingBigFloats = true
	var i int
	var n int

	/*
	requested 250,000 and got 275,002 matching chars with 
	n = int(2 + int(float64(digits)/12.3))
	digits*0.04
	16m22s 
	 */

	// ::: re 'n' ... setting the loop counter
	if digits < 60000 {
		// ... apparently, n, is the expected number of loops we may need in order to produce digits number of digits
		n = int(2 + int(float64(digits)/14.181647462))
		// comments re: n := int64(2 + int(float64(digits)/12))  // I tried this, and may try something like it again someday?? like /14.0 ?
	} else if digits < 100000 {
		n = int(2 + int(float64(digits)/14)) // 14
	} else if digits < 150000 {
		n = int(2 + int(float64(digits)/13.0)) // 13
	} else if digits < 200000 {
		n = int(2 + int(float64(digits)/12.5)) // 12.5
	} else {
		n = int(2 + int(float64(digits)/12.3)) // 12.3
	}

	// ::: set precision 
		// comments re: precision := uint(int(math.Ceil(math.Log2(10)*digits)) + int(math.Ceil(math.Log10(digits))) + 2) // the original
		// comments re: precision := uint(digits) // not good, not large enough, so ...
		digitsPlus := digits + digits*0.10 // because we needed a little more than the original programmer had figured on :)
		precision := uint(int(math.Ceil(math.Log2(10)*digitsPlus)) + int(math.Ceil(math.Log10(digitsPlus))) + 2)
	if digits < 60000 {
		digitsPlus := digits + digits*0.10 // because we needed a little more than the original programmer had figured on :)
		precision = uint(int(math.Ceil(math.Log2(10)*digitsPlus)) + int(math.Ceil(math.Log10(digitsPlus))) + 2)
	} else if digits < 100000 {
		digitsPlus := digits + digits*0.07 // because we needed a little more than the original programmer had figured on :)
		precision = uint(int(math.Ceil(math.Log2(10)*digitsPlus)) + int(math.Ceil(math.Log10(digitsPlus))) + 2)
	} else if digits < 150000 {
		digitsPlus := digits + digits*0.06 // because we needed a little more than the original programmer had figured on :)
		precision = uint(int(math.Ceil(math.Log2(10)*digitsPlus)) + int(math.Ceil(math.Log10(digitsPlus))) + 2)
	} else if digits < 200000 {
		digitsPlus := digits + digits*0.05 // because we needed a little more than the original programmer had figured on :)
		precision = uint(int(math.Ceil(math.Log2(10)*digitsPlus)) + int(math.Ceil(math.Log10(digitsPlus))) + 2)
	} else {
		digitsPlus := digits + digits*0.04 // because we needed a little more than the original programmer had figured on :)
		precision = uint(int(math.Ceil(math.Log2(10)*digitsPlus)) + int(math.Ceil(math.Log10(digitsPlus))) + 2)
	}

	c := new(big.Float).Mul(
		big.NewFloat(float64(426880)),
		new(big.Float).SetPrec(precision).Sqrt(big.NewFloat(float64(10005))),
	)

	k := big.NewInt(int64(6))
	k12 := big.NewInt(int64(12))
	l := big.NewFloat(float64(13591409))
	lc := big.NewFloat(float64(545140134))
	x := big.NewFloat(float64(1))
	xc := big.NewFloat(float64(-262537412640768000))
	m := big.NewFloat(float64(1))
	sum := big.NewFloat(float64(13591409))

	pi := big.NewFloat(0)

	x.SetPrec(precision)
	m.SetPrec(precision)
	sum.SetPrec(precision)
	pi.SetPrec(precision)

	bigI := big.NewInt(0)
	bigOne := big.NewInt(1)

	i = 1 // a secondary dedicated loop counter

	updateOutput1(fmt.Sprintf("\n\n	Digits of Pi is set at: %.0f, so, 'n' (the number of iterations to perform is actually set to: %d \n\n", digits, n))

		if n > 1000 || digits > 9000 {
			updateOutput1(fmt.Sprintf("\n Well, this is going to take a while, because you asked for too much pie (> 9,000 digits), actually %d iterattions were prescribed.\n", n))
		}


	for ; n > 0; n-- {
		select {
		case <-done: // ::: here an attempt is made to read from the channel (a closed channel can be read from successfully; but what is read will be the null/zero value of the type of chan (0, false, "", 0.0, etc.)
			// in the case of this particular channel (which is of type bool) we get the value false from having received from the channel when it is already closed. 
			// ::: if the channel known by the moniker "done" is already closed, that/it is to be interpreted as the abort signal by all listening processes. 
			fmt.Println("Goroutine chud-func-calcPi for-loop (1 of 1) is being terminated by select case finding the done channel to be already closed")
			return i, pi, start // Exit the goroutine
		default:
		i++

		// L calculation
		l.Add(l, lc)

		// X calculation
		x.Mul(x, xc)

		// M calculation
		kpower3 := big.NewInt(0)
		kpower3.Exp(k, big.NewInt(3), nil)
		ktimes16 := new(big.Int).Mul(k, big.NewInt(16))
		mtop := big.NewFloat(0).SetPrec(precision)
		mtop.SetInt(new(big.Int).Sub(kpower3, ktimes16))
		mbot := big.NewFloat(0).SetPrec(precision)
		mbot.SetInt(new(big.Int).Exp(new(big.Int).Add(bigI, bigOne), big.NewInt(3), nil))
		mtmp := big.NewFloat(0).SetPrec(precision)
		mtmp.Quo(mtop, mbot)
		m.Mul(m, mtmp)

		// Sum calculation
		t := big.NewFloat(0).SetPrec(precision)
		t.Mul(m, l)
		t.Quo(t, x)
		sum.Add(sum, t)

		// Pi calculation
		pi.Quo(c, sum)
		k.Add(k, k12)
		bigI.Add(bigI, bigOne)

		if i == 2 {
			updateOutput1("\nhere at 2\n")
		// finishChudIfsAndPrint(pi, "no", done, digits)
		}
				if i == 4 {
					updateOutput1("\nhere at 4\n")
				// finishChudIfsAndPrint(pi, "no", done, digits)
			}
			if i == 8 {
					updateOutput1("\nhere at 8\n")
				// finishChudIfsAndPrint(pi, "no", done, digits)
			}
			if i == 16 {
					updateOutput1("\nhere at 16\n")
				// finishChudIfsAndPrint(pi, "no", done, digits)
			}
			if i == 32 {
					updateOutput1("\nhere at 32\n")
				// finishChudIfsAndPrint(pi, "no", done, digits)
			}
			if i == 44 {
					updateOutput1("\nhere at 44\n")
				// finishChudIfsAndPrint(pi, "no", done, digits)
				// ::: Check pi and convert to []string -- and, set lenOfPi
				_, lenOfPi := checkPiTo59766(pi)
				updateOutput1(fmt.Sprintf("\n\nWe have confirmation via checkPiTo59766 that %d digits have been verified\n", lenOfPi))
			}
			if i == 52 {
					updateOutput1("\nhere at 52\n")
				// finishChudIfsAndPrint(pi, "no", done, digits)
			}
			if i == 62 {
					updateOutput1("\nhere at 62\n")
				// finishChudIfsAndPrint(pi, "no", done, digits)
			}
			if i == 72 {
					updateOutput1("\nhere at 72\n")
				// finishChudIfsAndPrint(pi, "no", done, digits)
			}
			if i == 82 {
					updateOutput1("\nhere at 82\n")
				// finishChudIfsAndPrint(pi, "no", done, digits)
			}
			if i == 92 {
					updateOutput1("\nhere at 92\n")
				// finishChudIfsAndPrint(pi, "no", done, digits)
			}
		
		if i == 100 {
			// useAlternateFile := "no" // the compiler is not happy unless it sees this created outside of an if
			updateOutput1(fmt.Sprintf("\n we have done %d loops and have %d loops to go\n", i, n))
		}
		if i == 200 {
			// useAlternateFile = "no" // still no
			updateOutput1(fmt.Sprintf("\n we have done %d loops and have %d loops to go\n", i, n))
			// finishChudIfsAndPrint(pi, useAlternateFile, done, digits)
			// ::: Check pi and convert to []string -- and, set lenOfPi
			// _, lenOfPi := checkPiTo59766(pi)
			// updateOutput1(fmt.Sprintf("\n\nWe have confirmation via checkPiTo59766 that %d digits have been verified\n", lenOfPi))
		}
		if i == 400 {
			// useAlternateFile = "no" // still no ::: based on this flag ...
			updateOutput1(fmt.Sprintf("\n we have done %d loops and have %d loops to go\n", i, n))
			// finishChudIfsAndPrint(pi, useAlternateFile, done, digits)
		}
			// ::: ... up to this point the user will be shown the verified pi message
			//
			// note below the: useAlternateFile = "chudDid800orMoreLoops"
			if i == 800 {
				// useAlternateFile = "chudDid800orMoreLoops"
				updateOutput1(fmt.Sprintf("\n we have done %d loops and have %d loops to go\n", i, n))
				// finishChudIfsAndPrint(pi, useAlternateFile, done, digits)
			}
			if i == 1600 {
				// useAlternateFile = "chudDid800orMoreLoops"
				updateOutput1(fmt.Sprintf("\n\n we have done %d loops and have %d loops to go\n", i, n))
				// finishChudIfsAndPrint(pi, useAlternateFile, done, digits)
			}
			if i == 2000 {
				// useAlternateFile = "chudDid800orMoreLoops"
				updateOutput1(fmt.Sprintf("\n\n we have done %d loops and have %d loops to go\n", i, n))
				// finishChudIfsAndPrint(pi, useAlternateFile, done, digits)
			}
			if i == 2400 {
				// useAlternateFile = "chudDid800orMoreLoops"
				updateOutput1(fmt.Sprintf("\n\n we have done %d loops and have %d loops to go\n", i, n))
				// finishChudIfsAndPrint(pi, useAlternateFile, done, digits)
				// ::: Check pi and convert to []string -- and, set lenOfPi
				// _, lenOfPi := checkPiTo59766(pi)
				// updateOutput1(fmt.Sprintf("\n\nWe have confirmation via checkPiTo59766 that %d digits have been verified\n", lenOfPi))
			}
			if i == 2800 {
				// useAlternateFile = "chudDid800orMoreLoops"
				updateOutput1(fmt.Sprintf("\n\n we have done %d loops and have %d loops to go\n", i, n))
				// finishChudIfsAndPrint(pi, useAlternateFile, done, digits)
			}
			if i == 3200 {
				// useAlternateFile = "chudDid800orMoreLoops"
				updateOutput1(fmt.Sprintf("\n\n we have done %d loops and have %d loops to go\n", i, n))
				// finishChudIfsAndPrint(pi, useAlternateFile, done, digits)
			}
			if i == 4000 {
				// useAlternateFile = "chudDid800orMoreLoops"
				updateOutput1(fmt.Sprintf("\n\n we have done %d loops and have %d loops to go\n", i, n))
				// finishChudIfsAndPrint(pi, useAlternateFile, done, digits)
			}
			if i == 6000 {
				// useAlternateFile = "chudDid800orMoreLoops"
				updateOutput1(fmt.Sprintf("\n\n we have done %d loops and have %d loops to go\n", i, n))
				// finishChudIfsAndPrint(pi, useAlternateFile, done, digits)
			}
				if i == 8000 {
					// useAlternateFile = "chudDid800orMoreLoops"
					updateOutput1(fmt.Sprintf("\n\n we have done %d loops and have %d loops to go\n", i, n))
					// finishChudIfsAndPrint(pi, useAlternateFile, done, digits)
				}
				if i == 10000 {
					// useAlternateFile = "chudDid800orMoreLoops"
					updateOutput1(fmt.Sprintf("\n\n we have done %d loops and have %d loops to go\n", i, n))
					// finishChudIfsAndPrint(pi, useAlternateFile, done, digits)
				}
				if i == 13000 {
					// useAlternateFile = "chudDid800orMoreLoops"
					updateOutput1(fmt.Sprintf("\n\n we have done %d loops and have %d loops to go\n", i, n))
					// finishChudIfsAndPrint(pi, useAlternateFile, done, digits)
				}
				if i ==17000 {
					// useAlternateFile = "chudDid800orMoreLoops"
					updateOutput1(fmt.Sprintf("\n\n we have done %d loops and have %d loops to go\n", i, n))
					// finishChudIfsAndPrint(pi, useAlternateFile, done, digits)
				}
				if i == 21000 {
					// useAlternateFile = "chudDid800orMoreLoops"
					updateOutput1(fmt.Sprintf("\n\n we have done %d loops and have %d loops to go\n", i, n))
					// finishChudIfsAndPrint(pi, useAlternateFile, done, digits)
				}
				if i == 25000 {
					// useAlternateFile = "chudDid800orMoreLoops"
					updateOutput1(fmt.Sprintf("\n\n we have done %d loops and have %d loops to go\n", i, n))
					// finishChudIfsAndPrint(pi, useAlternateFile, done, digits)
					// ::: Check pi and convert to []string -- and, set lenOfPi
					// _, lenOfPi := checkPiTo59766(pi)
					// updateOutput1(fmt.Sprintf("\n\nWe have confirmation via checkPiTo59766 that %d digits have been verified\n", lenOfPi))
				}

			if i == 30000 {
				// useAlternateFile = "chudDid800orMoreLoops"
				updateOutput1(fmt.Sprintf("\n\n we have done %d loops and have %d loops to go\n", i, n))
				// finishChudIfsAndPrint(pi, useAlternateFile, done, digits)
			}
				if i == 35000 {
					// useAlternateFile = "chudDid800orMoreLoops"
					updateOutput1(fmt.Sprintf("\n\n we have done %d loops and have %d loops to go\n", i, n))
					// finishChudIfsAndPrint(pi, useAlternateFile, done, digits)
				}
				if i == 40000 {
					// useAlternateFile = "chudDid800orMoreLoops"
					updateOutput1(fmt.Sprintf("\n\n we have done %d loops and have %d loops to go\n", i, n))
					// finishChudIfsAndPrint(pi, useAlternateFile, done, digits)
				}
				if i == 50000 {
					// useAlternateFile = "chudDid800orMoreLoops"
					updateOutput1(fmt.Sprintf("\n\n we have done %d loops and have %d loops to go\n", i, n))
					// finishChudIfsAndPrint(pi, useAlternateFile, done, digits)
				}
				if i == 60000 {
					// useAlternateFile = "chudDid800orMoreLoops"
					updateOutput1(fmt.Sprintf("\n\n we have done %d loops and have %d loops to go\n", i, n))
					// finishChudIfsAndPrint(pi, useAlternateFile, done, digits)
					// ::: Check pi and convert to []string -- and, set lenOfPi
					// _, lenOfPi := checkPiTo59766(pi)
					// updateOutput1(fmt.Sprintf("\n\nWe have confirmation via checkPiTo59766 that %d digits have been verified\n", lenOfPi))
				}
				if i == 65000 {
					// useAlternateFile = "chudDid800orMoreLoops"
					updateOutput1(fmt.Sprintf("\n\n we have done %d loops and have %d loops to go\n", i, n))
					// finishChudIfsAndPrint(pi, useAlternateFile, done, digits)
				}

			if i == 70000 {
				// useAlternateFile = "chudDid800orMoreLoops"
				updateOutput1(fmt.Sprintf("\n\n we have done %d loops and have %d loops to go\n", i, n))
				// finishChudIfsAndPrint(pi, useAlternateFile, done, digits)
				// ::: Check pi and convert to []string -- and, set lenOfPi
				// _, lenOfPi := checkPiTo59766(pi)
				// updateOutput1(fmt.Sprintf("\n\nWe have confirmation via checkPiTo59766 that %d digits have been verified\n", lenOfPi))
			}
			
			if n == 1 {
			updateOutput1(fmt.Sprintf("\n\nn has gone to 1, so we are finished, precision was: %d \n", precision))
			break
		}
		// 1,000,000 digits requires 70516 loops, per the run on May 7 2023 at 10:30
		//  was run on: Sun May  7 08:50:23 2023
		//  Total run was 8h4m39.7847064s
		// AND THE CALCULATION WAS INDEPENDENTLY VERIFIED !!!!!!!!!!!
		} // end of select
	} // end of for loop way up thar :: it prompts periodically to continue or die

	// ::: we are out of the loop, so we do the following just once:
	// finishChudIfsAndPrint(pi, "no", done, digits)
		// obtain file handle
			fileHandleBig, err1prslc2c := os.OpenFile("big_pie_is_in_here.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file
				check(err1prslc2c)                 // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
					// defer fileHandleBig.Close()   // It’s idiomatic to defer a Close immediately after opening a file.
		
			// to ::: file		
			_, err9bigpie := fmt.Fprint(fileHandleBig, pi)        // ::: dump this big-assed pie to a special log file
				check(err9bigpie)
			_, err9bigpie = fmt.Fprint(fileHandleBig, "\n was pi as a big.Float\n")  // add a suffix 
				check(err9bigpie)
	// also to the file, add ID and timestamp: :::file
	Hostname, _ := os.Hostname()
	_, err0 := fmt.Fprintf(fileHandleBig, "\n  -- Chud -- on %s \n", Hostname)
	check(err0)
	current_time := time.Now()
	_, err6 := fmt.Fprint(fileHandleBig, "was run on: ", current_time.Format(time.ANSIC), "\n")
	check(err6)
					_, errGoesHere := fmt.Fprint(fileHandleBig, "\n\n")
						check(errGoesHere)
		
	fileHandleBig.Close()
	// ::: Check pi and convert to []string -- and, set lenOfPi
	// _, lenOfPi := checkPiTo59766(pi)
	// updateOutput1(fmt.Sprintf("\n\nWe have confirmation via checkPiTo59766 that %d digits have been verified\n", lenOfPi))



	return i, pi, start // assigning i to 'loops' var in caller
}
/*
.
.
.
.
.
 */
// a helper func   
func finishChudIfsAndPrint(pi *big.Float, useAlternateFile string, done chan bool, digits float64) { // ::: - -

	// ::: Check pi and convert to []string -- and, set lenOfPi
	_, lenOfPi := checkPiTo59766(pi)
	updateOutput1(fmt.Sprintf("\n\nWe have confirmation via checkPiTo59766 that %d digits have been verified\n", lenOfPi))
	
	if digits > 39700 {
		// if lenOfPi > 46000 { // if length of pi is > 48,000 digits we have something really big
			// print to ::: screen
			// updateOutput1(fmt.Sprintf("\n\n\nWe have been tasked with making a lot of pie and it was sooo big it needed its own file ...\n"))
			// updateOutput1(fmt.Sprintf("... Go have a look in /.big_pie_is_in_here.txt to find all the digits of π you had requested. \n\n"))

			// print (log) to a special ::: file
			// obtain file handle
			fileHandleBig, err1prslc2c := os.OpenFile("big_pie_is_in_here.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file
			check(err1prslc2c)             // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
			defer fileHandleBig.Close()           // It’s idiomatic to defer a Close immediately after opening a file.

			// print to file
			_, err2prslc2c := fmt.Fprintf(fileHandleBig, "\n\nrick > 46,000 2 Here are %d calculated digits that we have NOT verified by reference: \n", 47000)
			check(err2prslc2c)

			// add ID and time stamp to ::: file 
			Hostname, _ := os.Hostname()
			current_time := time.Now()

			_, err0 := fmt.Fprintf(fileHandleBig, "\n  -- Chud -- on %s \n", Hostname)
			check(err0)
			_, err6 := fmt.Fprint(fileHandleBig, "was run on: ", current_time.Format(time.ANSIC), "\n")
			check(err6)

		// to ::: file		
		_, err9bigpie := fmt.Fprint(fileHandleBig, pi)                               // dump this big-assed pie to a special log file
		check(err9bigpie)
			/*
						for _, oneChar := range stringVerOfOurCorrectDigits {
					select {
					case <-done: // ::: here an attempt is made to read from the channel (a closed channel can be read from successfully; but what is read will be the null/zero value of the type of chan (0, false, "", 0.0, etc.)
						// in the case of this particular channel (which is of type bool) we get the value false from having received from the channel when it is already closed.
						// ::: if the channel known by the moniker "done" is already closed, that/it is to be interpreted as the abort signal by all listening processes.
						fmt.Println("Goroutine chud-func-calcPi for-loop (1 of 1) is being terminated by select case finding the done channel to be already closed")
						return // Exit the goroutine
					default:

						// fmt.Print(oneChar) // to the console // the whole point of using an alternate file is to not clutter up the console or the default file
						// *************************************** this is the one and only logging loop ******************************************************************************
						_, err8prslc2c := fmt.Fprint(fileHandleBig, oneChar) // to a file
						check(err8prslc2c)
					}
				}
						_, err9prslc2c := fmt.Fprintf(fileHandleBig, "\n...the preceding was logged/printed one char at a time\n")
				check(err9prslc2c) // ::: ...the preceding was logged/printed one char at a time
			 */


			fileHandleBig.Close()
		// }
	} else {
		stringVerOfOurCorrectDigits, lenOfPi := checkPiTo59766(pi)
		if lenOfPi < 600 {
			// obtain file handle
			fileHandleDefault, err91prslc2c := os.OpenFile("dataLog-From_Chudnovsky_Method_lengthy_prints.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file
			check(err91prslc2c)                // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
			defer fileHandleDefault.Close()

			//	print to ::: screen
			updateOutput1(fmt.Sprintf("\n\nlenOfPi < 600, so, Here are %d calculated digits that we have verified by reference (one at a time): \n", lenOfPi))

			// print via range to ::: screen	
			for _, oneChar := range stringVerOfOurCorrectDigits { // pi is finally ::: printed here via ranging 

				// to screen:
				updateOutput1(fmt.Sprintf("%s", oneChar)) // ::: to screen
			}

			// dump array as string to a ::: file 
			asString := strings.Join(stringVerOfOurCorrectDigits, "")
			_, lastError := fmt.Fprint(fileHandleDefault, asString) // to a file
			check(lastError)

			// also to the file, add ID and timestamp: :::file
			Hostname, _ := os.Hostname()
			_, err0 := fmt.Fprintf(fileHandleDefault, "\n  -- Chud -- on %s \n", Hostname)
			check(err0)
			current_time := time.Now()
			_, err6 := fmt.Fprint(fileHandleDefault, "was run on: ", current_time.Format(time.ANSIC), "\n")
			check(err6)

			// print to ::: screen	
			updateOutput1(fmt.Sprintf("\n\n"))
		}
	}
	

	/*

		// else {

			// } else { continues below: (in other words, the following if-else conditions are only checked if length of pi was < 55,000 digits)
			if useAlternateFile == "chudDid800orMoreLoops" {
				// obtain file handle
				fileHandleChud, err1prslc2c := os.OpenFile("dataLog-From_Chudnovsky_Method_lengthy_prints.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file
					check(err1prslc2c)                   // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
					defer fileHandleChud.Close()        // It’s idiomatic to defer a Close immediately after opening a file.

				// print to ::: file
					_, err2prslc2c := fmt.Fprintf(fileHandleChud, "\n\n800+ 3 Here are %d calculated digits that we have verified by reference: \n", lenOfPi)
						check(err2prslc2c)

				// add ID and time stamp to ::: file
				Hostname, _ := os.Hostname()
				current_time := time.Now()

				_, err0 := fmt.Fprintf(fileHandleChud, "\n  -- Chud -- on %s \n", Hostname)
				check(err0)
				_, err6 := fmt.Fprint(fileHandleChud, "was run on: ", current_time.Format(time.ANSIC), "\n")
				check(err6)

				// dump array as string to a ::: file
				asString := strings.Join(stringVerOfOurCorrectDigits, "")
					_, lastError := fmt.Fprint(fileHandleChud, asString) // to a file
						check(lastError)

			} else if useAlternateFile == "ChudDidLessThanOneHundredLoops" {
				// obtain file handle
					fileHandleDefault, err1prslc2d := os.OpenFile("dataLog-From_calculate-pi-and-friends.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file
						check(err1prslc2d)           // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
							defer fileHandleDefault.Close()      // It’s idiomatic to defer a Close immediately after opening a file.

				// print to ::: file
					_, err2prslc2d := fmt.Fprintf(fileHandleDefault, "\n\n<100  4 Here are %d calculated digits that we have verified by reference: \n", lenOfPi)
						check(err2prslc2d)

				// add ID and time stamp to ::: file
				Hostname, _ := os.Hostname()
				current_time := time.Now()

				_, err0 := fmt.Fprintf(fileHandleDefault, "\n  -- Chud -- on %s \n", Hostname)
				check(err0)
				_, err6 := fmt.Fprint(fileHandleDefault, "was run on: ", current_time.Format(time.ANSIC), "\n")
				check(err6)

				// print to ::: screen
					updateOutput1(fmt.Sprintf("\n\n Here are %d calculated digits that we have verified by reference: \n", lenOfPi))

				// print one char at a time to ::: screen & file
					for _, oneChar := range stringVerOfOurCorrectDigits {
						// to screen
						fmt.Print(oneChar)

						// to file
						_, err8prslc2d := fmt.Fprint(fileHandleDefault, oneChar)
							check(err8prslc2d)
					}
					fileHandleDefault.Close()

				// add ID and time stamp to ::: file
					Hostname, _ = os.Hostname()
					current_time = time.Now()

					_, err0 = fmt.Fprintf(fileHandleDefault, "\n  -- Chud -- on %s \n", Hostname)
						check(err0)
					_, err6 = fmt.Fprint(fileHandleDefault, "was run on: ", current_time.Format(time.ANSIC), "\n")
						check(err6)


				// ::: this final else handles any instances of useAlternateFile not caught above
			} else {
				// obtain file handle to pi-and-friends.txt
					fileHandleDefault, err1prslc2d := os.OpenFile("dataLog-From_calculate-pi-and-friends.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file
						check(err1prslc2d)           // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
							defer fileHandleDefault.Close() // It’s idiomatic to defer a Close immediately after opening a file.

				// to ::: file
					_, err2prslc2d := fmt.Fprintf(fileHandleDefault, "\n\nelse 6 Here are %d calculated digits that we have verified by reference: ChudDidLessThanOneHundredLoops ::\n", lenOfPi)
						check(err2prslc2d)

					Hostname, _ := os.Hostname()
					current_time := time.Now()

					_, err0 := fmt.Fprintf(fileHandleDefault, "\n  -- Chud -- on %s \n", Hostname)
						check(err0)
					_, err6 := fmt.Fprint(fileHandleDefault, "was run on: ", current_time.Format(time.ANSIC), "\n")
						check(err6)

				// to ::: screen
					updateOutput1(fmt.Sprintf("\n Here are %d calculated digits that we have verified by reference:\n", lenOfPi))

					asString := strings.Join(stringVerOfOurCorrectDigits, "")
						updateOutput1(fmt.Sprintf("\n catch-all, asString: %s\n", asString))

				// obtain file handel to ...lengthy_prints.txt
					fileHandleChud, err1prslc2c := os.OpenFile("dataLog-From_Chudnovsky_Method_lengthy_prints.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file
						check(err1prslc2c)                // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
							defer fileHandleChud.Close()

					// to ::: file
					_, err0 = fmt.Fprintf(fileHandleDefault, "\n  -- Chud -- on %s \n", Hostname)
						check(err0)
					_, err6 = fmt.Fprint(fileHandleDefault, "was run on: ", current_time.Format(time.ANSIC), "\n")
						check(err6)

					_, err2prslc2da := fmt.Fprint(fileHandleDefault, "\nResults from running Chud can be viewed in a file\n")
						check(err2prslc2da)

			fileHandleDefault.Close()

			}
	 */


	// } // end of if's else, way up thar "if lenOfPi > 46000 {} else {"   so, this has been the instance where pi is shorter than 55,000
}