package main

import (
	"fmt"
	"math"
	"math/big"
	"os"
	"runtime"
	"time"
)

func bbpMax(fyneFunc func(string), digits int, done chan bool) { // ::: - -
	/*
				π = Σ(k=0 to ∞) [ (1/16^k) ( 4/(8k + 1) - 2/(8k + 4) - 1/(8k + 5) - 1/(8k + 6) ) ]

			In this algorithm for the BBP formula we use a channel with a buffer size equal to the requested number of digits of Pi : n. And, since we will also deploy n parallel
		    workers the buffer len of n is intended to prevent stalling of the overall process; we won't have to worry about the receiver keeping up. Note that the workers aren’t
		    calculating individual digits of π directly -- they’re only computing terms of the BBP formula, which get added together, and addition doesn’t care about order.

			In the workers closure, each goroutine computes a term of the BBP formula (e.g., for k = id) and sends it via result <- R. The buffer lets up to n of these pile up.

			The second for loop in bbpMax pulls n values from result with pi.Add(pi, <-result), adding them to pi. The buffer ensures workers aren’t stuck waiting for this loop to grab a result.

			Blocking: Normally, if the buffer fills (more than n sends occur before there are any receives done), senders are forced to pause: the last sender to succeed is said to
			have blocked. But in our algorithm the buffer is sized perfectly (n), so with n workers and n receives, continuous flow is assured unless something weird happens (like done closing early).

			By marshalling the optimal number of workers, and setting the channel buffer to match; along with using runtime.GOMAXPROCS(numCPU) we assure maximum speed.
	*/
	start := time.Now()

	numCPU := runtime.NumCPU() // Discover the number of CPUs we have 
	runtime.GOMAXPROCS(numCPU) // Have this function and all assigned closures use all available CPUs

	usingBigFloats = true
	iters_bbp := 1
	n := digits // of Pi to calculate [digits is a passed-in value] . this copy is made for later brevity

	// Determine and set the precision of our big floats based on the number of digits of Pi 'n' that we are to calculate ...
	// ... figures out how many bits of precision (p) we need for n digits of π. Since each digit needs about 3.32 bits (log₂(10)), we multiply by n and add 
	// ... a few extra bits for safety. It’s like saying, “Give me enough room to store this many digits accurately.”
	p := uint((int(math.Log2(10)))*n + 3) // A calculation is performed, cast as int, re-cast as in unassigned int, and the result is assigned to a new var 'p' 

	result := make(chan *big.Float, n) // Create a channel though which to pass pointers to big float values; with a buffer size that is set via a named value 'n'
	// i.e., a buffered channel that can hold n *big.Float values BEFORE BLOCKING. A channel is like a pipe between workers. The buffer (n) means it can store n results 
	// ... (pieces of π) WITHOUT WAITING (before needing to wait) for someone to pick one or more of them up. It’s A QUEUE for the math bits being calculated.

	worker := workers(p) // Here is where we assign a closure to a reusable func value : worker . A closure is a function that can remember the state of its values from the last time it was called. 

	pi := new(big.Float).SetPrec(p).SetInt64(0) // that last part: SetInt64(0) is entirely redundant and included only for the sake of being explicit 

	for i := 0; i < n; i++ {
		select {
		case <-done: // ::: here an attempt is made to read from the channel (a closed channel can be read from successfully; but what is read will be the null/zero value of the type of chan (0, false, "", 0.0, etc.)
			// in the case of this particular channel (which is of type bool) we get the value false from having received from the channel when it is already closed. 
			// ::: if the channel known by the moniker "done" is already closed, that/it is to be interpreted as the abort signal by all listening processes. 
			fmt.Println("Goroutine bbpMax for-loop (1 of 2) is being terminated by select case finding the done channel to be already closed")
			return // Exit the goroutine
		default:
			go worker(i, result) // Call the closure worker as an independent go routine (a separate thread) and pass i and our chan named result
			iters_bbp = i        // Track the number of iterations done by this method 
		}
	}

	for i := 0; i < n; i++ {
		select {
		case <-done: // ::: here an attempt is made to read from the channel (a closed channel can be read from successfully; but what is read will be the null/zero value of the type of chan (0, false, "", 0.0, etc.)
			// in the case of this particular channel (which is of type bool) we get the value false from having received from the channel when it is already closed. 
			// ::: if the channel known by the moniker "done" is already closed, that/it is to be interpreted as the abort signal by all listening processes. 
			fmt.Println("Goroutine bbpMax for-loop (2 of 2) is being terminated by select case finding the done channel to be already closed")
			return // Exit the goroutine
		default:
			pi.Add(pi, <-result) // Add is a method intrinsic to big float types. It appears that we may be accumulating at pi via the result chan ?
			iters_bbp = i        // Track the number of iterations done by this method [in a similar but separate loop]
		}
	}

	dur := time.Since(start)
	fyneFunc(fmt.Sprintf("bbpMax executed with %d digits and ", n))
	fyneFunc(fmt.Sprintf("took %v to calculate %d digits of pi \n\n", dur, n))
	/* or:
	output := fmt.Sprintf("%s\nIt only took BBP %v to calculate the following %d digits of pi\n", codeSnippet, dur, n) // my codeSnippet was a rune of this entire method ::: setup w codeSnippet
	output := fmt.Sprintf("\nIt only took BBP %v to calculate the following %d digits of pi\n", dur, n) // ::: setup sans codeSnippet
	fyneFunc(output) // ::: display
	*/

	// n was the number of digits of pi to calculate, and here below n specifies the number of digits past the decimal to print using the indexed ver of the %.nf verb 
	// fmt.Printf("%[1]*.[2]*[3]f \n", 1, n, pi) // original from CLI version
	// updateChan <- updateData{text:"%[1]*.[2]*[3]f \n", 1, n, pi} // does not work, even with the correct signature for updateChan <- updateData{text:"
	fiveK := new(big.Float)
	if pi.Cmp(fiveK) == 1 { // if pi > 5,000
		// obtain file handle
		fileHandleBig, err1prslc2c := os.OpenFile("big_pie_is_in_here.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file
		check(err1prslc2c)                                                                                             // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
		// defer fileHandleBig.Close()   // It’s idiomatic to defer a Close immediately after opening a file.

		// to ::: file
		_, err9bigpie := fmt.Fprint(fileHandleBig, pi) // ::: dump this big-assed pie to a special log file
		check(err9bigpie)
		_, err9bigpie = fmt.Fprint(fileHandleBig, "\n was pi as a big.Float\n") // add a suffix
		check(err9bigpie)
		fileHandleBig.Close()

		file1 := "/Users/quasar/grokTriesAgain/big_pie_is_in_here.txt" // Replace with your first file path
		file2 := "/Users/quasar/grokTriesAgain/piOneMil.txt"           // Replace with your second file path

		count, err := compareFiles2(file1, file2)
		if err != nil {
			fmt.Println("Error:", err)
		}
		updateOutput2(fmt.Sprintf("\n\nMatched %d characters in sequence from the start.\n", count))
	} else {
		fyneFunc(fmt.Sprintf("%[1]*.[2]*[3]f \n", 1, n, pi)) // a more explicit ver of %.nf  [here we also specify the width of the number to the left of the decimal]
	}

	// Write run-stats to a log file
	t := time.Now()
	elapsed := t.Sub(start)
	fileHandle, err1 := os.OpenFile("dataLog-From_calculate-pi-and-friends.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file
	check(err1)                                                                                                             // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
	defer fileHandle.Close()                                                                                                // It’s idiomatic to defer a Close immediately after opening a file.
	Hostname, _ := os.Hostname()
	_, err0 := fmt.Fprintf(fileHandle, "\n  -- calculate pi using the bbp formula -- on %s \n", Hostname)
	check(err0)
	current_time := time.Now()
	_, err6 := fmt.Fprint(fileHandle, "was run on: ", current_time.Format(time.ANSIC), "\n")
	check(err6)
	_, err4 := fmt.Fprintf(fileHandle, "%.02f was Iterations/Seconds \n", float64(iters_bbp)/elapsed.Seconds())
	check(err4)
	_, err5 := fmt.Fprintf(fileHandle, "%d was total Iterations \n", iters_bbp)
	check(err5)
	TotalRun := elapsed.String() // cast time durations to a String type for Fprintf "formatted print"
	_, err7 := fmt.Fprintf(fileHandle, "Total run was %s \n ", TotalRun)
	check(err7)

	done <- true // trying to signal a clean and complete exit on a buffered chan :
	/*
			If we were to use an Unbuffered Channel: Sender (done <- true) waits for receiver (<-done). No buffer means they handshake directly -- describing how communication works between a
		sender and a receiver in Go when there’s no buffer to hold values. A buffered channel (e.g., make(chan bool, 1)) has space to store values—like a mailbox with slots. A sender can drop
		something in (e.g., done <- true) and keep going, even if no one’s there to pick it up yet, as long as the buffer isn’t full. An unbuffered channel (e.g., make(chan bool)) has no storage
		— it’s like a direct phone line instead of a mailbox. The sender and receiver have to connect at the same time for the value to pass.

		Signal Meaning: Define true = success, false = failure in bbpMax, then check the value, not just the send.
	*/

}

// Create a closure that will be assigned to worker
func workers(p uint) func(id int, result chan *big.Float) { //  ::: - -
	// Captured variables ?
	B1 := new(big.Float).SetPrec(p).SetInt64(1) // Initialize these new big float values to: 1, 2, 4, 5, etc. [a strange-looking limited series]
	B2 := new(big.Float).SetPrec(p).SetInt64(2) // All of these will become part of the closure function value below 
	B4 := new(big.Float).SetPrec(p).SetInt64(4)
	B5 := new(big.Float).SetPrec(p).SetInt64(5)
	B6 := new(big.Float).SetPrec(p).SetInt64(6)
	B8 := new(big.Float).SetPrec(p).SetInt64(8)
	B16 := new(big.Float).SetPrec(p).SetInt64(16)

	return func(id int, result chan *big.Float) { // This is the closure function/value specified in the signature of the constructor func workers() it is passed our requested n of digits & a chan named result
		Bn := new(big.Float).SetPrec(p).SetInt64(int64(id))
		C1 := new(big.Float).SetPrec(p).SetInt64(1)

		for i := 0; i < id; i++ {
			C1.Mul(C1, B16) // first use of a captured value: B16
		}

		C2 := new(big.Float).SetPrec(p)
		C2.Mul(B8, Bn) // cumulative calculations continue using the captured high-precision versions of the integers: 1, 2, 4, 5, 6, 8, and 16

		T1 := new(big.Float).SetPrec(p)
		T1.Add(C2, B1)
		T1.Quo(B4, T1)

		T2 := new(big.Float).SetPrec(p)
		T2.Add(C2, B4)
		T2.Quo(B2, T2)

		T3 := new(big.Float).SetPrec(p)
		T3.Add(C2, B5)
		T3.Quo(B1, T3)

		T4 := new(big.Float).SetPrec(p)
		T4.Add(C2, B6)
		T4.Quo(B1, T4)

		R := new(big.Float).SetPrec(p)
		R.Sub(T1, T2)
		R.Sub(R, T3)
		R.Sub(R, T4)
		R.Quo(R, C1)

		result <- R // Several threads of worker instances will use this same result channel to collect the ingredients needed to assemble/bake the requested Pi
	}
	// adapted by Richard Woolley
}
