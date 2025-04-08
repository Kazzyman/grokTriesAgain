package main

import (
	"fmt"
	"math/big"
	"time"
)

// @formatter:off

func ArchimedesBig(fyneFunc func(string), done chan bool) { // ::: - -
	
	fyneFunc(fmt.Sprintf("\n\nYou've selected a demonstration of Rick's improved version of Archimedes' method for aproximating the value of Pi : 3.14159...\n\n"))

	fyneFunc(fmt.Sprintf("The goal is to accurately calculate over 2,700 correct digits of Pi. We'll need to use floating-point numbers with thousands of decimal places.\n"))
	fyneFunc(fmt.Sprintf("This can be done using Rick's most-favoured language: go.lang, or simply Go. GoLand (by JetBrains) wil be our IDE.\n"))
	fyneFunc(fmt.Sprintf("We'll also be using the Fyne.io package to create a graphical (windowed) user interface.\n\n"))

	fyneFunc(fmt.Sprintf("All of our variables must be big.Floats (as in the above code, this we now do)\n"))

	r := big.NewFloat(1)
	s1 := big.NewFloat(1)
	numberOfSides := big.NewFloat(6)

	a := new(big.Float)
	b := new(big.Float)
	p := new(big.Float)
	s2 := new(big.Float)
	p_d := new(big.Float)
	s1_2 := new(big.Float)
	
	// ::: screen
		fyneFunc(fmt.Sprintf("\nGo's precision is set to 55000 on all of our variables (as per the above code).\n"))

	precision := 55000
	p_d.SetPrec(uint(precision))
	a.SetPrec(uint(precision))
	s1_2.SetPrec(uint(precision))
	s2.SetPrec(uint(precision))
	b.SetPrec(uint(precision))
	p.SetPrec(uint(precision))
	r.SetPrec(uint(precision))
	s1.SetPrec(uint(precision))
	numberOfSides.SetPrec(uint(precision))
	fyneFunc(fmt.Sprintf("\nThen, we do some initial asignments and calculations (as per above):\n"))

	// ::: initial calculation 
	numberOfSides.Mul(numberOfSides, big.NewFloat(2))
	s1_2.Quo(s1, big.NewFloat(2))
	a.Sqrt(new(big.Float).Sub(r, new(big.Float).Mul(s1_2, s1_2))) // a = Sqrt(r-(s1_2)^2) where r is always 1  [all values must be big floats, including the const 1]

	// ::: screen
	fyneFunc(fmt.Sprintf("\n\n\t\tTo calculate the height (a) of a right triangle formed by bisecting a side of a polygon inscribed in a unit circle (radius r = 1). \n\n" +
		"\t\t\tThe polygon’s side length (s1) is halved (s1_2 = s1 / 2), and this computation helps refine the polygon’s perimeter to approximate \n\n\t\t\tπ as the number of sides increases.")) 

	b.Sub(r, a) // b = r-a  where a: is the Height of the bisected triangle (from the prior step).\n"))
	s2.Sqrt(new(big.Float).Add(new(big.Float).Mul(b, b), new(big.Float).Mul(s1_2, s1_2))) // s2 = Sqrt(b*b + s1_2*s1_2)\n"))
	/* The above statement is explained below:
	Inputs:
	  b: short side from midpoint to circle edge (a big float)
	  s1_2: half the current side length (s1 / 2, a big float)
	Output:
	  s2: new side length of the polygon (a big float)

	1. Compute b^2
	temp1 = b * b

	2. Compute (s1_2)^2
	temp2 = s1_2 * s1_2

	3. Add the two squares
	temp3 = temp1 + temp2

	4. Take the square root to get the new side length
	s2 = square_root(temp3)
	*/
	// ::: screen
	fyneFunc(fmt.Sprintf("\n\n\t\t\tHere is some seudo code for the algorithm:\n\n\n\n" +
		"\t\t\tInputs:\n\n\t\t\t\t  b: short side from midpoint to circle edge (a big float)\n\n" +
		"\t\t\t\t  s1_2: half the current side length (s1 / 2, a big float)\n\n\t\t\tOutput:\n\n" +
		"\t\t\t\t  s2: new side length of the polygon (a big float)\n\n\t\t\n\n\t\t\t1. Compute b^2\n\n" +
		"\t\t\t\ttemp1 = b * b\n\n\t\t\n\n\t\t\t2. Compute (s1_2)^2\n\n\t\t\t\ttemp2 = s1_2 * s1_2\n\n" +
		"\t\t\n\n\t\t\t3. Add the two squares\n\n\t\t\t\ttemp3 = temp1 + temp2\n\n\t\t\n\n" +
		"\t\t\t4. Take the square root to get the new side length\n\n\t\t\t\ts2 = square_root(temp3)\n\n\nNow we get to work!!\n\n"))
	
	s1.Set(s2) // Use big.Float method Set, to asign s2 to s1\n"))
	p.Mul(numberOfSides, s1) // p = numberOfSides * s1\n"))
	p_d.Set(p) // Use big.Float method Set, to asign p to p_d  'read p sub d'\n"))
	
	fyneFunc(fmt.Sprintf("\n\n"))

		for i := 0; i < 5001; i++ { // ::: calculation, result, and notification loop < - - - - - - - - - - - - - - - - - < -
			select {
			case <-done: // ::: here an attempt is made to read from the channel (a closed channel can be read from successfully; but what is read will be the null/zero value of the type of chan (0, false, "", 0.0, etc.)
			// in the case of this particular channel (which is of type bool) we get the value false from having received from the channel when it is already closed. 
				// ::: if the channel known by the moniker "done" is already closed, that/it is to be interpreted as the abort signal by all listening processes. 
				fmt.Println("Goroutine Archimedes for-loop (1 of 1) is being terminated by select case finding the done channel to be already closed")
				return // Exit the goroutine
			default:
				// the calculation 
					numberOfSides.Mul(numberOfSides, big.NewFloat(2))
					s1_2.Quo(s1, big.NewFloat(2)) 
					a.Sqrt(new(big.Float).Sub(r, new(big.Float).Mul(s1_2, s1_2))) 
					b.Sub(r, a)
					s2.Sqrt(new(big.Float).Add(new(big.Float).Mul(b, b), new(big.Float).Mul(s1_2, s1_2))) // ; 
					s1.Set(s2)
					p.Mul(numberOfSides, s1)
					p_d.Set(p)
					p_d.Quo(p_d, big.NewFloat(2)) 
		
				// ::: conditional progressive results and notifications 
				if i == 24 {
					fyneFunc(fmt.Sprintf("----------------------------------------------------------------------------------------------------------------------\n" +
						"  %d iterations were completed in order to yeild the following digits of π\n\n", i))
		
					fyneFunc(fmt.Sprintf("    %.20f is the big.Float of what we have calculated  ----- per Archimedes' at 24 iterations, formatted: 20f\n", p_d))
					fyneFunc("    3.141592653589793238  vs the value of π from the web\n")
		
					formattedNum := formatWithThousandSeparators(numberOfSides) // Manually format with thousand separators
					fyneFunc(fmt.Sprintf("\nthe above was estimated from a %s  --- sided polygon\n\n", formattedNum))
		
					// stringOfSum := p_d.Text('f', 999999) // create a string version of a big result, up to a million digits can be handled, if not verified, here-with
					_, lenOfPi := checkPiTo59766(p_d) // sets global var lenOfPi
					fyneFunc(fmt.Sprintf("... And, it has been verified that we actually calculated pi correctly to %d digits!\n\n", lenOfPi))
		
					fyneFunc("... Mister A. would have wept!\n\n\n\n")
				}
		
				if i == 50 {
					fyneFunc(fmt.Sprintf("----------------------------------------------------------------------------------------------------------------------\n" +
						"  %d iterations were completed in order to yeild the following digits of π\n\n", i))
		
					fyneFunc(fmt.Sprintf("    %.33f is the big.Float of what we have calculated  ----- per Archimedes' at 50 iters, formatted: 33f\n", p_d))
					fyneFunc("    3.141592653589793238462643383279502  ----- is the value of π from the web\n")
		
					// stringOfSum := p_d.Text('f', 999999) // create a string version of a big result, up to a million digits can be handled, if not verified, here-with
					_, lenOfPi := checkPiTo59766(p_d) // sets global var lenOfPi
					fyneFunc(fmt.Sprintf("... And, it has been verified that we actually calculated pi correctly to %d digits!\n\n", lenOfPi))
		
					formattedNum := formatWithThousandSeparators(numberOfSides) // Manually format with thousand separators
					fyneFunc(fmt.Sprintf(" the above was estimated from a %s  --- sided polygon\n\n\n\n", formattedNum))
				}
		
				if i == 150 {
					fyneFunc(fmt.Sprintf("----------------------------------------------------------------------------------------------------------------------\n" +
						"  %d iterations were completed in order to yeild the following digits of π\n\n", i))
		
					fyneFunc(fmt.Sprintf("   %.95f   ----- per Rick's modified Archimedes' method, formatted 95f\n", p_d))
					fyneFunc("   3.14159265358979323846264338327950288419716939937510582097494459230781640628620899862803482534211  ----- is from web\n")
		
					// stringOfSum := p_d.Text('f', 999999) // create a string version of a big result, up to a million digits can be handled, if not verified, here-with
					_, lenOfPi := checkPiTo59766(p_d) // sets global var lenOfPi
					fyneFunc(fmt.Sprintf("... And, it has been verified that we actually calculated pi correctly to %d digits!\n\n", lenOfPi))
		
					formattedNum := formatWithThousandSeparators(numberOfSides) // Manually format with thousand separators
					fyneFunc(fmt.Sprintf(" the above was estimated from a %s  --- sided polygon\n\n\n\n", formattedNum))
				}
		
				if i == 200 {
					fyneFunc(fmt.Sprintf("----------------------------------------------------------------------------------------------------------------------\n" +
						"  %d iterations were completed in order to yeild the following digits of π\n\n", i))
		
					fyneFunc(fmt.Sprintf("   %.122f   ---- ... Archimedes' method, formatted: 122f\n", p_d))
					fyneFunc("   3.14159265358979323846264338327950288419716939937510582097494459230781640628620899862803482534211706798214808651328230664709  ----- is from web\n")
		
					formattedNum := formatWithThousandSeparators(numberOfSides) // Manually format with thousand separators
					fyneFunc(fmt.Sprintf("\n\nour figure was estimated from a %s  --- sided polygon\n\n", formattedNum))
		
					// stringOfSum := p_d.Text('f', 999999) // create a string version of a big result, up to a million digits can be handled, if not verified, here-with
					_, lenOfPi := checkPiTo59766(p_d) // sets global var lenOfPi
					fyneFunc(fmt.Sprintf("... And, it has been verified that we actually calculated pi correctly to %d digits!\n\n", lenOfPi))
		
					fyneFunc(" ... working ...\n\n")
				}
		
				if i == 1200 || i == 2200 || i == 3200 || i == 4200 {
					fyneFunc(fmt.Sprintf("... still working, %d iterations completed ...\n\n", i))
				}
		
				if i == 5000 { // was 5500
					fyneFunc(fmt.Sprintf("------------------------------------------------------------------------------------------------------------------------------------------\n\n"))
		
					formattedNum := formatWithThousandSeparators(numberOfSides) // Manually format with thousand separators
					fyneFunc(fmt.Sprintf("All Done! So, how many sides does our polygon have now? A lot:\n\nA staggering:\n\n%s\n\nSIDED POLYGON !!!\n\n\n", formattedNum))
		
					fyneFunc(fmt.Sprintf("%d iterations were completed to yeild well over 2,700 correct digits of π!!!\n\n", i))
					fyneFunc(fmt.Sprintf("Go's math/big objects were set to a precision value of: %d  --- here is your GIANT slice of pie:\n\n", precision))
					fyneFunc(fmt.Sprintf("  %.3020f \n\n\nper Rick's modified Archimedes' method, formatted: 3020f\n\n\n", p_d))
		
		
					_, lenOfPi := checkPiTo59766(p_d) // sets global var lenOfPi // Sets the global lenOfPi [the calculated and verified quantity of digits of pi]

					fyneFunc(fmt.Sprintf("... And, it has been verified that we actually calculated pi correctly to %d digits!\n\n\n\n by Richard (Rick) H. Woolley\n\n\n\n\n\n", lenOfPi))
					// my code says: pi correctly to 3012 digits (3023 unconfirmed digits, including the decimal, were printed to the terminal)
					// Gpt said: The number you provided contains 3021 digits (when given the 3023 unconfirmed digits, including the decimal, that were printed to the terminal)
					// Gemini said: After using a character counting tool, I can confirm that there are 3012 digits after the decimal point. But Gemini hallucinates a lot!
				}
				if i < 24 {
					time.Sleep(135 * time.Millisecond) // Slow it down slightly for visibility
					if i == 2 {fyneFunc(fmt.Sprintf("\t\tSleeping each iteration for 135 milliseconds...\n\n"))}
				}
				if i > 23 && i < 50 {
					if i == 26 {fyneFunc(fmt.Sprintf("\t\tSleeping each iteration for 55 milliseconds...\n\n"))}
					time.Sleep(55 * time.Millisecond) // Slow it down slightly for visibility
				}
				if i > 49 && i < 150 {
					if i == 52{fyneFunc(fmt.Sprintf("\t\tSleeping each iteration for 35 milliseconds...\n\n"))}
					time.Sleep(35 * time.Millisecond) // Slow it down slightly for visibility
				}
				if i > 149 && i < 400 {
					if i == 152 {fyneFunc(fmt.Sprintf("\t\tSleeping each iteration for 7 milliseconds...\n\n"))}
					time.Sleep(7 * time.Millisecond) // Slow it down slightly for visibility
				}
				if i > 399 && i < 1100 {
					if i == 402 {fyneFunc(fmt.Sprintf("\t\tSleeping each iteration for 2 milliseconds...\n\n"))}
					time.Sleep(2 * time.Millisecond) // Slow it down slightly for visibility
				}
				if i > 1099 && i < 2000 {
					if i == 1102 {fyneFunc(fmt.Sprintf("\t\tSleeping each iteration for 1 milliseconds...\n\n"))}
					time.Sleep(time.Millisecond) // Slow it down slightly for visibility
				}
				if i > 1999 {
					time.Sleep(0 * time.Millisecond)
					if i == 2002 {fyneFunc(fmt.Sprintf("\t\tNo more sleeping!!!...\n\n"))}}		
			} // end of select
		} // end of for loop, only one way out 
		/*
		// ::: Prepare to exit the Archimedes method function, (set coast-clear flag and enable all buttons)
			calculating1 = false
			for _, btn := range buttons1 {
				btn.Enable()
			}
		 */

}
