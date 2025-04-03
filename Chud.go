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

// This will be a little-bit tricky. We want to use callbacks etc. so that we can use the smoother-scrolling fyneFunc(fmt.Sprintf("")) way of doing prints ... 
// ... but this chudnovsky section is a cascade of functions: chudnovskyBig()-->calcPi()-->finishChudIfsAndPrint() 
func chudnovskyBig(fyneFunc func(string), digits int, done chan bool) { // ::: - -

	// fyneFunc(fmt.Sprintf("\nThe forgoing is the entire code for this method.\n\n"))

	// ::: fyneFunc will use updateOutput[1-4] depending on from which window called -- so we pass fyneFunc to calcPi(fyneFunc, float64(digits), start, loops) thusly 
	fyneFunc(fmt.Sprintf("\n... working ...\n"))
	usingBigFloats = true
	var loops int
	start := time.Now() // start will be passed, and then passed back, in order to be compared with end time t

	pi := new(big.Float)

	// ::: calcPi  <---- runs from here: v v v v v v v  
	loops, pi, start = calcPi(fyneFunc, float64(digits), start, done)
	// ::: calcPi ----- ^ ^ ^ 

	/*
		if loops < 100 {
			fyneFunc(fmt.Sprintf("\nLess than 100 loops, so here is a peek at the prospective value of Pi as a big float, and formatted 0.122f, is : \n%0.122f \n\n", pi))
		}
	 */
	
	// The following runs ::: after calcPi 
	fyneFunc(fmt.Sprintf("\n loops were: %d, and digits requested was: %d \n", loops, digits))

	fyneFunc(fmt.Sprintf("\n 	The Chudnovsky algorithm is an incredibly-fast algorithm for calculating the digits of pi. It was developed by Gregory Chudnovsky and his "))
	fyneFunc(fmt.Sprintf("brother David Chudnovsky in the 1980s. It is more efficient than other algorithms and is based on the theory of modular equations. It has been "))
	fyneFunc(fmt.Sprintf("used to calculate pi to over 62 trillion digits.\n\n"))

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
	// ::: Prepare to exit the Chud method functions 
		calculating = false // Allow another method to be selected.
		for _, btn := range buttons1 { // ok to only Enable buttons1, because I expect to only ever execute this from window1
			btn.Enable() // ::: Enable
		}
}
/*
.
.
.
.
 */
// calculate Pi for n number of digits
func calcPi(fyneFunc func(string), digits float64, start time.Time, done chan bool) (int, *big.Float, time.Time) {
	// ::: fyneFunc will be the proper one to match the calling window[1-4] therefore prints will go where they should 
runeToPrint := `
	/**
	 *   This is an implementation for https://en.wikipedia.org/wiki/Chudnovsky_algorithm
	 *   "It can be improved using binary splitting http://numbers.computation.free.fr/Constants/Algorithms/splitting.html
	 *   if we were to split it into two independent parts and simplify the formula." For more details, visit:
	 *             https://www.craig-wood.com/nick/articles/pi-chudnovsky
	 */
`
	fyneFunc(fmt.Sprintf(runeToPrint))

	usingBigFloats = true
	var i int

	// ::: n ...
	// ... apparently, n, is the expected number of loops we may need to produce digits number of digits
	n := int64(2 + int(float64(digits)/14.181647462))
	// comments re: n := int64(2 + int(float64(digits)/12))  // I tried this, and may try something like it again someday?? like /14.0 ?
	
	// set precision 
		// comments re: precision := uint(int(math.Ceil(math.Log2(10)*digits)) + int(math.Ceil(math.Log10(digits))) + 2) // the original
		// comments re: precision := uint(digits) // not good, not large enough, so ...
			digitsPlus := digits + digits*0.10 // because we needed a little more than the orriginal programmer had figured on :)
			precision := uint(int(math.Ceil(math.Log2(10)*digitsPlus)) + int(math.Ceil(math.Log10(digitsPlus))) + 2)

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

	// this is a flag; if it is set to zero we exit
	queryIfTimeToDie := 1
	i = 1 // a secondary dedicated loop counter


		if n > 8998 {
			fyneFunc(fmt.Sprintf("\n Well, this is going to take a while, because you asked for too much pie (> 8990)\n"))
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
		finishChudIfsAndPrint(fyneFunc, pi, "no")
		}
		if i == 4 {
			finishChudIfsAndPrint(fyneFunc, pi, "no")
		}
		if i == 8 {
			finishChudIfsAndPrint(fyneFunc, pi, "no")
		}
		if i == 16 {
			finishChudIfsAndPrint(fyneFunc, pi, "no")
		}
		if i == 32 {
			finishChudIfsAndPrint(fyneFunc, pi, "no")
		}
		if i == 44 {
			finishChudIfsAndPrint(fyneFunc, pi, "no")
		}
		if i == 52 {
			finishChudIfsAndPrint(fyneFunc, pi, "no")
		}
		if i == 62 {
			finishChudIfsAndPrint(fyneFunc, pi, "no")
		}
		if i == 72 {
			finishChudIfsAndPrint(fyneFunc, pi, "no")
		}
		if i == 82 {
			finishChudIfsAndPrint(fyneFunc, pi, "no")
		}
		if i == 92 {
			finishChudIfsAndPrint(fyneFunc, pi, "no")
		}

		useAlternateFile := "no" // no means to use the standard log file rather than some special one
		// the compiler is not happy unless it sees this created outside of an if
		// But, wait. Why is the compiler allowing me to violate the no new var left of the := assignment ??? This IS in a loop !!!!
		if i == 100 {
			// useAlternateFile := "no" // the compiler is not happy unless it sees this created outside of an if
			fyneFunc(fmt.Sprintf("\n we are at %d loops\n", i))
		}
		if i == 200 {
			// useAlternateFile = "no" // still no
			fyneFunc(fmt.Sprintf("\n we are at %d loops: \n", i))
			finishChudIfsAndPrint(fyneFunc, pi, useAlternateFile)
		}
		if i == 400 {
			// useAlternateFile = "no" // still no ::: based on this flag ...
			fyneFunc(fmt.Sprintf("\n we are at %d loops: \n", i))
			finishChudIfsAndPrint(fyneFunc, pi, useAlternateFile)
		}
		// ::: ... up to this point the user will be shown the verified pi message
		//
		// note below the: useAlternateFile = "chudDid800orMoreLoops"
		if i == 800 {
			useAlternateFile = "chudDid800orMoreLoops"
			fyneFunc(fmt.Sprintf("\n we are at %d loops: \n", i))
			finishChudIfsAndPrint(fyneFunc, pi, useAlternateFile)
		}
		if i == 1600 {
			useAlternateFile = "chudDid800orMoreLoops"
			fyneFunc(fmt.Sprintf("\n\n we are at %d loops: \n", i))
			finishChudIfsAndPrint(fyneFunc, pi, useAlternateFile)
		}
		if i == 2000 {
			useAlternateFile = "chudDid800orMoreLoops"
			fyneFunc(fmt.Sprintf("\n\n we are at %d loops: \n", i))
			finishChudIfsAndPrint(fyneFunc, pi, useAlternateFile)
		}
		if i == 2400 {
			useAlternateFile = "chudDid800orMoreLoops"
			fyneFunc(fmt.Sprintf("\n\n we are at %d loops: \n", i))
			finishChudIfsAndPrint(fyneFunc, pi, useAlternateFile)
		}
		if i == 2800 {
			useAlternateFile = "chudDid800orMoreLoops"
			fyneFunc(fmt.Sprintf("\n\n we are at %d loops: \n", i))
			finishChudIfsAndPrint(fyneFunc, pi, useAlternateFile)
		}
		if i == 3200 {
			useAlternateFile = "chudDid800orMoreLoops"
			fyneFunc(fmt.Sprintf("\n\n we are at %d loops: \n", i))
			finishChudIfsAndPrint(fyneFunc, pi, useAlternateFile)
		}
		if i == 4000 {
			useAlternateFile = "chudDid800orMoreLoops"
			fyneFunc(fmt.Sprintf("\n\n we are at %d loops: \n", i))
			finishChudIfsAndPrint(fyneFunc, pi, useAlternateFile)
		}
		if i == 6000 {
			useAlternateFile = "chudDid800orMoreLoops"
			fyneFunc(fmt.Sprintf("\n\n we are at %d loops: \n", i))
			finishChudIfsAndPrint(fyneFunc, pi, useAlternateFile)
		}
		if i == 8000 {
			useAlternateFile = "chudDid800orMoreLoops"
			fyneFunc(fmt.Sprintf("\n\n we are at %d loops: \n", i))
			finishChudIfsAndPrint(fyneFunc, pi, useAlternateFile)
		}
		if queryIfTimeToDie == 0 {
			fyneFunc(fmt.Sprintf("if queryIfTimeToDie is 0, time to die"))
			fyneFunc(fmt.Sprintf("\nprecisionision was: %d \n", precision))
			break
		}
		// 1,000,000 digits requires 70516 loops, per the run on May 7 2023 at 10:30
		//  was run on: Sun May  7 08:50:23 2023
		//  Total run was 8h4m39.7847064s
		// AND THE CALCULATION WAS INDEPENDANTLY VERIFIED !!!!!!!!!!!
		} // end of select
	} // end of for loop way up thar :: it prompts periodically to continue or die

	// ::: we are out of the loop, so we do the following just once:

		// obtain file handle
			fileHandleBig, err1prslc2c := os.OpenFile("big_pie_is_in_here.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file
				check(err1prslc2c)                                                                                             // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
					defer fileHandleBig.Close()                                                                                    // It’s idiomatic to defer a Close immediately after opening a file.
		
			// to ::: file		
			_, err9bigpie := fmt.Fprint(fileHandleBig, pi)                               // dump this big-assed pie to a special log file
				check(err9bigpie)
			_, err9bigpie = fmt.Fprint(fileHandleBig, "\nwas pi as a big.Float\n")  // add a suffix 
				check(err9bigpie)
		
					_, errGoesHere := fmt.Fprint(fileHandleBig, "\n\n")
						check(errGoesHere)
		
	fileHandleBig.Close()

	return i, pi, start // assigning i to loops in caller
}
/*
.
.
.
.
.
 */
// a helper func   
func finishChudIfsAndPrint(fyneFunc func(string), pi *big.Float, useAlternateFile string) { // ::: - -

	// ::: Check pi and convert to []string -- and, set lenOfPi
		stringVerOfOurCorrectDigits, lenOfPi := checkPiTo59766(pi) 
	
	if lenOfPi < 600 {
		// obtain file handle
			fileHandleDefault, err91prslc2c := os.OpenFile("dataLog-From_Chudnovsky_Method_lengthy_prints.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file
				check(err91prslc2c)                // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
					defer fileHandleDefault.Close()
			
		//	print to ::: screen
			fyneFunc(fmt.Sprintf("\n\nlenOfPi < 600, so, Here are %d calculated digits that we have verified by reference (one at a time): \n", lenOfPi))
			
		// print via range to ::: screen	
			for _, oneChar := range stringVerOfOurCorrectDigits { // pi is finally ::: printed here via ranging 
				// to screen:
					fyneFunc(fmt.Sprintf("%s", oneChar)) // ::: to screen
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
		fyneFunc(fmt.Sprintf("\n\n"))
	}

	if lenOfPi > 46000 { // if length of pi is > 48,000 digits we have something really big
		// print to ::: screen
		fyneFunc(fmt.Sprintf("\n\n\nWe have been tasked with making a lot of pie and it was sooo big it needed its own file ...\n"))
		fyneFunc(fmt.Sprintf("... Go have a look in /.big_pie_is_in_here.txt to find all the digits of π you had requested. \n\n"))

		// print (log) to a special ::: file
				// obtain file handle
				fileHandleBig, err1prslc2c := os.OpenFile("big_pie_is_in_here.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file
				check(err1prslc2c)             // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
				defer fileHandleBig.Close()           // It’s idiomatic to defer a Close immediately after opening a file.
			
				// print to file
				_, err2prslc2c := fmt.Fprintf(fileHandleBig, "\n\nHere are %d calculated digits that we have verified by reference: \n", lenOfPi)
				check(err2prslc2c)
			
				for _, oneChar := range stringVerOfOurCorrectDigits {
					// fmt.Print(oneChar) // to the console // the whole point of using an alternate file is to not clutter up the console or the default file
					// *************************************** this is the one and only logging loop ******************************************************************************
					_, err8prslc2c := fmt.Fprint(fileHandleBig, oneChar) // to a file
					check(err8prslc2c)
				}
				_, err9prslc2c := fmt.Fprintf(fileHandleBig, "\n...the precisioneding was logged one char at a time\n")
				check(err9prslc2c)
				fileHandleBig.Close()
	} else {

		// } else { continues below: (in other words, the following if-else conditions are only checked if length of pi was < 55,000 digits)
		if useAlternateFile == "chudDid800orMoreLoops" {
			// obtain file handle 
			fileHandleChud, err1prslc2c := os.OpenFile("dataLog-From_Chudnovsky_Method_lengthy_prints.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file
				check(err1prslc2c)                   // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
				defer fileHandleChud.Close()        // It’s idiomatic to defer a Close immediately after opening a file.
			
			// print to ::: file
				_, err2prslc2c := fmt.Fprintf(fileHandleChud, "\n\nHere are %d calculated digits that we have verified by reference: \n", lenOfPi)
					check(err2prslc2c)
			
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
				_, err2prslc2d := fmt.Fprintf(fileHandleDefault, "\n\nHere are %d calculated digits that we have verified by reference: \n", lenOfPi)
					check(err2prslc2d)

			// print to ::: screen
				fyneFunc(fmt.Sprintf("\n\n Here are %d calculated digits that we have verified by reference: \n", lenOfPi))
			
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
				Hostname, _ := os.Hostname()
				current_time := time.Now()
				
				_, err0 := fmt.Fprintf(fileHandleDefault, "\n  -- Chud -- on %s \n", Hostname)
					check(err0)
				_, err6 := fmt.Fprint(fileHandleDefault, "was run on: ", current_time.Format(time.ANSIC), "\n")
					check(err6)


			// ::: this final else handles any instances of useAlternateFile not caught above
		} else {
			// obtain file handle to pi-and-friends.txt
				fileHandleDefault, err1prslc2d := os.OpenFile("dataLog-From_calculate-pi-and-friends.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file
					check(err1prslc2d)           // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
						defer fileHandleDefault.Close() // It’s idiomatic to defer a Close immediately after opening a file.
			
			// to ::: file
				_, err2prslc2d := fmt.Fprintf(fileHandleDefault, "\n\nHere are %d calculated digits that we have verified by reference: ChudDidLessThanOneHundredLoops ::\n", lenOfPi)
					check(err2prslc2d)

				Hostname, _ := os.Hostname()
				current_time := time.Now()
				
				_, err0 := fmt.Fprintf(fileHandleDefault, "\n  -- Chud -- on %s \n", Hostname)
					check(err0)
				_, err6 := fmt.Fprint(fileHandleDefault, "was run on: ", current_time.Format(time.ANSIC), "\n")
					check(err6)
			
			// to ::: screen
				fyneFunc(fmt.Sprintf("\n Here are %d calculated digits that we have verified by reference:\n", lenOfPi))
					
				asString := strings.Join(stringVerOfOurCorrectDigits, "")
					fyneFunc(fmt.Sprintf("\n catch-all, asString: %s\n", asString))
			
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
	} // end of if's else, way up thar "if lenOfPi > 46000 {} else {"   so, this has been the instance where pi is shorter than 55,000
}