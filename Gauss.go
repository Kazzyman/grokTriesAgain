package main

import (
	"fmt"
	"math"
	"math/big"
	"time"
)

func Gauss_Legendre(fyneFunc func(string), done chan bool) {
	about := `
	The Gauss-Legendre algorithm is an iterative method to compute π
	based on the arithmetic-geometric mean (AGM) of two numbers. It was developed by Carl Friedrich Gauss and later refined by Adrien-Marie Legendre.
	Unlike the Wallis product (a slow-converging product) or Gregory-Leibniz (a slow-converging series), Gauss-Legendre converges quadratically, meaning the
	number of correct digits roughly doubles with each iteration—making it highly efficient for high-precision calculations.

	Convergence: The values a and b approach the same limit (the AGM), and t sub n adjusts to refine the estimate. 
	The quadratic convergence comes from the squaring in the error terms.

`
	fyneFunc(fmt.Sprintf(about))

	usingBigFloats = false
	start := time.Now()
	var Ricks_value float64 // Rick's code
	// var exterior_catcher int           // Rick's code
	// initialize lists
	pin := []float64{}
	an := []float64{1}
	pn := []float64{1}

	tn := []float64{float64(1) / float64(4)} // Rick's improved code, instead of the following ...
	// ...
	// tn := []float64{}                      // should have been able to accomplish the next line here as in initialize, see above
	// tn = append(tn, float64(1)/float64(4)) // append the quotient of 1.0/4.0 to list tn

	bn := []float64{float64(1) / math.Sqrt(2)} // Rick's improved code, instead of the following ...
	// ...
	// bn := []float64{}                         // should have been able to accomplish the next line here as in initialize, as above
	// bn = append(bn, float64(1)/math.Sqrt(2))  // append the quotient of 1.0/sqrt of 2 to list bn

	iters := 3
	// run the algorithm iters times
	for i := 0; i < iters; i++ { // call the 5 funcs (a,b,t,p, and pi) defined below, each of which returns just one []float64
		select {
		case <-done: // ::: here an attempt is made to read from the channel (a closed channel can be read from successfully; but what is read will be the null/zero value of the type of chan (0, false, "", 0.0, etc.)
			// in the case of this particular channel (which is of type bool) we get the value false from having received from the channel when it is already closed. 
			// ::: if the channel known by the moniker "done" is already closed, that/it is to be interpreted as the abort signal by all listening processes. 
			fmt.Println("Goroutine Gauss for-loop (1 of 2) is being terminated by select case finding the done channel to be already closed")
			return // Exit the goroutine
		default:
			an = a(an, bn)
			bn = b(an, bn)
			tn = t(an, bn, tn, pn) // calls func t and passes to it 4 []float64s
			pn = p(pn)
			pin = pi(an, bn, tn, pin) // pin array ends up containing the 4 values that were calculated for Pi
			// each time pi is called it returns a new version of a []float64 array
		}
	}

	// print results
	// this for loop runs 4 times, therefore range pin yielded _ and 4 ???? There are no ';'s after the for, so, ??
	// ... no, 'range' looks to be a partner to the for itself
	// ran as ...
	// for exterior_catcher, value := range pin { // with the two test prints below
	for _, value := range pin { // skip initializing a counter? no!, not sure exactly all that 'range' does, but when appllied to pin fetches calculated Pi from pin and apparently also another return that is being tossed via _
		select {
		case <-done: // ::: here an attempt is made to read from the channel (a closed channel can be read from successfully; but what is read will be the null/zero value of the type of chan (0, false, "", 0.0, etc.)
			// in the case of this particular channel (which is of type bool) we get the value false from having received from the channel when it is already closed. 
			// ::: if the channel known by the moniker "done" is already closed, that/it is to be interpreted as the abort signal by all listening processes. 
			fmt.Println("Goroutine Gauss for-loop (2 of 2) is being terminated by select case finding the done channel to be already closed")
			return // Exit the goroutine
		default:
			// above 'range' seems to tell the for to "range" accross pin and assign a successive element to 'value', there were 4 elements, so it runs 4 times -- the _ catches the unneeded return from 'range' which starts at 0 and goes to 3 in this loop
			fyneFunc(fmt.Sprintf("pin is %.16f, and ... ", pin))            // pin is an aray of calculated values for Pi  // Rick's code to discover same
			fyneFunc(fmt.Sprintf("%.16f Was calculated herewith\n", value)) // 'value' created on prior 'for' line and is set 4 times to a successive element of pin
			Ricks_value = value                                             // Rick's code to grab that final 'value' from last iteration
			// fmt.Printf("\n\nTop underscore is %d \n\n", exterior_catcher) it starts at 0 and goes to 3

		}
	}
	// fyneFunc(fmt.Sprintf(Ricks_value)) // Rick's code
	// fmt.Printf("\n\nBottom underscore is %d \n\n", exterior_catcher) this exterior_catcher var is never touched by the for loop
	fyneFunc(fmt.Sprintf("3.1415926535897932 <-- compared to the actual value of Pi"))
	fyneFunc(fmt.Sprintf("1 23456789012345 counting to fifteen \n"))
	fyneFunc(fmt.Sprintf("   ... via the Gauss–Legendre algorithm ... \n"))

	piAsBF := new(big.Float)
	piAsBF = big.NewFloat(Ricks_value) // pi is being cast to big from float64

	t := time.Now()
	elapsed := t.Sub(start)
	TotalRun := elapsed.String() // cast time durations to a String type for Fprintf "formatted print"

	fyneFunc(fmt.Sprintf("Pi is %0.9f, and run was %s\n", piAsBF, TotalRun))
	// printResultStatsLong(piAsBF, 0, "Gauss–Legendre", iters, TotalRun, selection)
}

// helper functions follow:
func a(an, bn []float64) []float64 { // func a accepts an and bn of type []float64, and returns a []float64
	a := (an[len(an)-1] + bn[len(bn)-1]) / float64(2) // create local 'a' = (element of an indexed by len of an-1) + (element of bn indexed by len of bn-1) ??
	an = append(an, a)                                // append a to an
	return an
}
func b(an, bn []float64) []float64 {
	b := math.Sqrt(an[len(an)-2] * bn[len(bn)-1]) // create local 'b' = (sqrt of element an indexed by len of an-2) * (element of bn indexed by len of bn-1) ??
	bn = append(bn, b)
	return bn
}
func t(an, bn, tn, pn []float64) []float64 { // accepts 4 parameters of type []float64, and returns a []float64
	t := tn[len(tn)-1] - pn[len(pn)-1]*math.Pow((an[len(an)-2]-an[len(an)-1]), 2)
	tn = append(tn, t)
	return tn
}
func p(pn []float64) []float64 { // accepts one []float64
	p := 2 * pn[len(pn)-1] // create local p as 2 * the element of pn indexed by len(pn)-1
	pn = append(pn, p)
	return pn
}
func pi(an, bn, tn, pin []float64) []float64 { // this func is all about appending just one value to pin ...
	pi := math.Pow((an[len(an)-1]+bn[len(bn)-1]), 2) / (4 * tn[len(tn)-1]) // ... and this is that value ... x^2 / (4 * an element from tn) ... x = two elements added together
	//                                             ^2
	pin = append(pin, pi)
	return pin
	// adapted by Richard Woolley
} // End of Gauss_Legendre Set // case 37: // -- AMFGauss_LegendreB
