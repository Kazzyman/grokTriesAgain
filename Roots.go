package main

import (
	"fmt"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

// @formatter:off

// SetupRootsDemo sets up the roots demo UI and returns the button for window2
func SetupRootsDemo(mgr *TrafficManager, radicalEntry, workEntry *widget.Entry) *ColoredButton {
	rootsBtn := NewColoredButton(
		"Roots demo usage: enter an integer in each of the above fields\n" +
			"2 or 3 in the first field, then any positive integer in the second\n" +
			"then click this button to run the calculation of square or cube root\n" +
			"                   -*-*- Rick's own-favorite method -*-*-     ",
		color.RGBA{255, 255, 100, 235},
		func() {
			if calculating2 {return}
			// if mgr.IsCalculating() {return} // bail if coast not clear // ::: remove if we are not planning to use this
			
			trimmedRadicalString := strings.TrimRight(radicalEntry.Text, " ")
			radical, err := strconv.Atoi(trimmedRadicalString)
			if err != nil || (radical != 2 && radical != 3) {
				if radical == 0 {
					updateOutput2("\nPlease read the usage instructions on the button that you clicked\n")
					return
				}
				updateOutput2("Invalid radical: enter 2 or 3\n")
				return
			}
			trimmedWorkPieceString := strings.TrimRight(workEntry.Text, " ")
			workPiece, err := strconv.Atoi(trimmedWorkPieceString)
			if err != nil || workPiece < 0 {
				updateOutput2("Invalid number: enter a non-negative integer\n")
				return
			}
			fmt.Printf(" ::: - Radical is set to: %d\n", radical) // debug
			fmt.Printf(" ::: - Work Piece is set to: %d\n", workPiece) // debug
			mgr.SetRadical(radical) // instead of passing these variables we have elected to try a little OOP 
			mgr.SetWorkPiece(workPiece)
			// mgr.SetCalculating(true) // ::: remove if we are not planning to use this
			for _, btn := range buttons2 {
				btn.Disable()
			}
			for _, btn := range rootBut2 {
				btn.Enable()
			}
			currentDone = make(chan bool) // ::: New channel per run
			go func(done chan bool) {
					defer func() {
						mgr.Reset()
						for _, btn := range buttons2 {
							btn.Enable()
						}
					}()
				xRootOfy(done) // ::: formatted to highlight the meat
					mgr.SetCalculating(false)
			}(currentDone) // ::: pass via closure
		},
	)
	return rootsBtn
}

func xRootOfy(done chan bool) {
	sortedResults = nil 
	usingBigFloats = false
	TimeOfStartFromTop := time.Now()

	radical2or3 := mgr.GetRadical() // trying out some OOP here 
	workPiece := mgr.GetWorkPiece()

	setPrecisionForSquareOrCubeRoot(radical2or3, workPiece) // sets precision only, basis is radical2or3 and workPiece

	updateOutput2("\n\nBuilding table...\n")
	buildPairsSlice(radical2or3)
	updateOutput2("Table built, starting calculation...\n")
	startBeforeCall := time.Now()

	var indx int
	breakOutLabel1:
	for i := 0; i < 400000; i += 2 { // this is meant to be a pretty big loop 825,000 is the number of 
		select {
		case <-done: // ::: here an attempt is made to read from the channel (a closed channel can be read from successfully; but what is read will be the null/zero value of the type of chan (0, false, "", 0.0, etc.)
			// in the case of this particular channel (which is of type bool) we get the value false from having received from the channel when it is already closed. 
			// ::: if the channel known by the moniker "done" is already closed, that/it is to be interpreted as the abort signal by all listening processes. 
			fmt.Println("Goroutine xRootOfy-1 for-loop (1 of 1) is being terminated by select case finding the done channel to be already closed")
			updateOutput2("\nProcess terminated, ready for another selection\n")
			return // Exit the goroutine
		default:
			if mgr.ShouldStop() {
				updateOutput2("Calculation of a root aborted\n")
				return
			}
			returnedEarly := readPairsSlice(i, startBeforeCall, radical2or3, workPiece, done)
			if returnedEarly {
				break breakOutLabel1 // break out of select and show the final 'early' result
			}
			handlePerfectSquaresAndCubes(TimeOfStartFromTop, radical2or3, workPiece, mgr)
			if diffOfLarger == 0 || diffOfSmaller == 0 {
				updateOutput2("\nbreakOut\n")
				break breakOutLabel1 // because we have a perfect square or cube; need to break out of the for loop which is parent to the select (a simple break would only break out of the select) 
			}
			if i%80000 == 0 && i > 0 { // if remainder of div is 0 (every 80,000 iterations) conditional progress updates print
				stringVindx := formatInt64WithThousandSeparators(int64(indx))
				updateOutput2(fmt.Sprintf("\n%s iterations completed... of 400,000\n", stringVindx))
				updateOutput2(fmt.Sprintf("\n... still working ...\n")) // ok
	
				fmt.Printf("%s iterations completed... of 400,000\n", stringVindx)
				fmt.Println(i, "... still working ...")
			}
			indx = i // save/copy to a wider scope for later use outside this loop
		}
	}
	fmt.Println("Loop completed at index:", indx) // Debug

	// ::: Show the final result
	fmt.Println("Entering result block, mathSqrtCheat 'square':", mathSqrtCheat, "mathCbrtCheat 'cube':", mathCbrtCheat) // Debug
	// ::: "Entering result block ... "

	t_s2 := time.Now()
	elapsed_s2 := t_s2.Sub(TimeOfStartFromTop)
	if diffOfLarger != 0 || diffOfSmaller != 0 { // if not a perfect square or cube do this else skip due to detection of perfect result
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Panic in result block:", r)
				updateOutput2("\nError calculating result\n")
			}
		}()
		fileHandle, err31 := os.OpenFile("dataLog-From_calculate-pi-and-friends.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		check(err31)
		defer fileHandle.Close()

		Hostname, _ := os.Hostname()
		fmt.Fprintf(fileHandle, "\n  -- %d root of %d by a ratio of perfect Products -- on %s \n", radical2or3, workPiece, Hostname)
		fmt.Fprint(fileHandle, "was run on: ", time.Now().Format(time.ANSIC), "\n")
		fmt.Fprintf(fileHandle, "%d was total Iterations \n", indx)

		fmt.Println("Sorting results...") // Debug
			sort.Slice(sortedResults, func(i, j int) bool { return sortedResults[i].pdiff < sortedResults[j].pdiff })
		fmt.Println("Sorted results, length:", len(sortedResults)) // Debug

		if len(sortedResults) > 0 {
			if radical2or3 == 2 {
				updateOutput2(fmt.Sprintf("\n%0.9f, it's the best approximation for the Square Root of %d", sortedResults[0].result, workPiece))
				fmt.Println("GUI updated, printing to console...") // Debug
				fmt.Printf("%s\n", sortedResults[0].result)
				fmt.Println("Writing to file...") // Debug
				fmt.Fprintf(fileHandle, "%s \n", sortedResults[0].result)
				fmt.Println("File written") // Debug
			}
			if radical2or3 == 3 {
				updateOutput2(fmt.Sprintf("\n%0.9f, it's the best approximation for the Cube Root of %d", sortedResults[0].result, workPiece))
				fmt.Println("GUI updated, printing to console...") // Debug
				fmt.Printf("%s\n", sortedResults[0].result)
				fmt.Println("Writing to file...") // Debug
				fmt.Fprintf(fileHandle, "%s \n", sortedResults[0].result)
				fmt.Println("File written") // Debug
			}
		}

		TotalRun := elapsed_s2.String()
		fmt.Fprintf(fileHandle, "Total run was %s \n ", TotalRun)
		fmt.Printf("Calculation completed in %s\n", elapsed_s2)
		updateOutput2(fmt.Sprintf("\nCalculation completed in %s\n", elapsed_s2))
	} else {
		fmt.Println("Skipped result block due to perfect result detection") // Debug
	}
}

func readPairsSlice(i int, startBeforeCall time.Time, radical2or3, workPiece int, done chan bool) bool { // ::: - -
	// each time that readPairsSlice is called we do two initial reads of pairsSlice prior to entering a loop in which four reads are done many many times ...
	// ... these next two lines are the two initial reads (done using a passed index: i)
	oneReadOfSmallerRoot := pairsSlice[i].root // Read a smaller PP and its root (just once) for each time readPairsSlice is called
	oneReadOfSmallerPP := pairsSlice[i].product // pairsSlice is a slice of two-element structs, i.e., pairs 
		breakOutLabel2: // ::: use of this label DOES NOT reenter the following for loop -- seems like kind of shitty syntax, but it's a good thing!
		for iter := 0; iter < 410000 && i < len(pairsSlice); iter++ { // go big, '410,000', but not so big that we would read past the end of the pairsSlice
				select {
				case <-done: // ::: here an attempt is made to read from the channel (a closed channel can be read from successfully; but what is read will be the null/zero value of the type of chan (0, false, "", 0.0, etc.)
					// in the case of this particular channel (which is of type bool) we get the value false from having received from the channel when it is already closed. 
					// ::: if the channel known by the moniker "done" is already closed, that/it is to be interpreted as the abort signal by all listening processes. 
					fmt.Println("Goroutine readPairsSlice-1 for-loop (1 of 1) is being terminated by select case finding the done channel to be already closed")
					return false // Exit the goroutine
				default:
				i++ // Note how we increment the passed index i here within the loop rather than in the for loop header ... 
				// note also that the for loop's header (the for clause) only sets and increments the loop counter -- secret sauce is what that is. 
				largerPerfectProduct := pairsSlice[i].product // i has been incremented since the initial 'one-time' read of oneReadOfSmallerPP ...
		
				// ... and, we keep incrementing the i until largerPerfectProduct is greater than (oneReadOfSmallerPP * workPiece)
				if largerPerfectProduct > oneReadOfSmallerPP * workPiece { // For example: workPiece may be 11: 3.32*3.32.   Larger PP may be 49: 7*7.   Smaller oneReadPP may be 4: 2*2. 
		
					ProspectivePHitOnLargeSide := largerPerfectProduct // rename it, badly; ::: 3
					rootOfProspectivePHitOnLargeSide := pairsSlice[i].root // grab larger side's root ::: 4
		
					ProspectivePHitOnSmallerSide := pairsSlice[i-1].product // these are reads 5 & 6 (initial comprise 1&2, larger comprise 3&4) 
					rootOfProspectivePHitOnSmallerSide := pairsSlice[i-1].root
		
		
					diffOfLarger = ProspectivePHitOnLargeSide - (workPiece * oneReadOfSmallerPP) // ::: PH_larger - (WP * _once)     7 - (11 * 4)
					// What does it tell us if we find that the sum of one of the larger roots from the table : ProspectivePHitOnLargeSide
					// 'plus' the negative of another smaller root from the table (times our WP) turns out to be zero?
					diffOfSmaller = (workPiece * oneReadOfSmallerPP) - ProspectivePHitOnSmallerSide // ::: (WP * _once) - PH_smaller    (11 * 4) - 
		
					if diffOfLarger == 0 {
						fmt.Println(colorCyan, "\n The", radical2or3, "root of", workPiece, "is", colorGreen, float64(rootOfProspectivePHitOnLargeSide)/float64(oneReadOfSmallerRoot), colorReset, "\n")
						updateOutput2(fmt.Sprintf("\n The %d root of %d is %0.33f\n\n", radical2or3, workPiece, float64(rootOfProspectivePHitOnLargeSide)/float64(oneReadOfSmallerRoot)))
		
						mathSqrtCheat = math.Sqrt(float64(workPiece)) // ::: this line was initially missing, but me thinks it must belong here, as it mirrors what follows in the next if 
						mathCbrtCheat = math.Cbrt(float64(workPiece)) // I cheated? Yea, a bit. But only in order to generate verbiage to print re a perfect root having been found
						break breakOutLabel2 // because we have a perfect square or cube; need to break out of the for loop which is parent to the select (a simple break would only break out of the select) 
					}
					if diffOfSmaller == 0 {
						fmt.Println(colorCyan, "\n The", radical2or3, "root of", workPiece, "is", colorGreen, float64(rootOfProspectivePHitOnSmallerSide)/float64(oneReadOfSmallerRoot), colorReset, "\n")
						updateOutput2(fmt.Sprintf("\n The %d root of %d is %0.33f\n\n", radical2or3, workPiece, float64(rootOfProspectivePHitOnSmallerSide)/float64(oneReadOfSmallerRoot)))
		
						mathSqrtCheat = math.Sqrt(float64(workPiece)) // ::: I cheated? Yea, a bit. But only in order to generate verbiage to print re a perfect root having been found 
						mathCbrtCheat = math.Cbrt(float64(workPiece))
						break breakOutLabel2 // because we have a perfect square or cube; need to break out of the for loop which is parent to the select (a simple break would only break out of the select) 
					}
					
					// 'large' element of a couplet: 
					if diffOfLarger < precisionOfRoot {
						result := float64(rootOfProspectivePHitOnLargeSide) / float64(oneReadOfSmallerRoot) // :::    root/root
						pdiff := float64(diffOfLarger) / float64(ProspectivePHitOnLargeSide) // :::                   diff/PP
							sortedResults = append(sortedResults, Results{result: result, pdiff: pdiff}) // collect results into a slice that will eventually be sorted 
						fmt.Printf("Found large prospect at index %d: result=%f, diff=%d\n", i, result, diffOfLarger) // Debug, ditto
						updateOutput2(fmt.Sprintf("Found large prospect at index %d: result=%f, diff=%d\n", i, result, diffOfLarger)) // update, info
						if diffOfLarger < 2 {break breakOutLabel2} // reconsider this later addition to the code base?
					}
					// 'small' element of a couplet: 
					if diffOfSmaller < precisionOfRoot {
						result := float64(rootOfProspectivePHitOnSmallerSide) / float64(oneReadOfSmallerRoot)
						pdiff := float64(diffOfSmaller) / float64(ProspectivePHitOnSmallerSide)
							sortedResults = append(sortedResults, Results{result: result, pdiff: pdiff}) // collect results into a slice that will eventually be sorted 
						fmt.Printf("Found small prospect at index %d: result=%f, diff=%d\n", i, result, diffOfSmaller) // Debug
						updateOutput2(fmt.Sprintf("Found small prospect at index %d: result=%f, diff=%d\n", i, result, diffOfSmaller)) // Debug
						if diffOfSmaller < 2 {break breakOutLabel2} // reconsider this later addition to the code base?
					}
		
					// ::: we will be potentially duplicating Results struct -> slice 
					// larger side section: ----------------------------------------------------------------------------------------------------------------------------------------
					// ---------------------------------------------------------------------------------------------------------------------------------------------------------------
						if diffOfLarger < precisionOfRoot { // report the prospects, their differences, and the calculated result for the Sqrt or Cbrt
							fmt.Println("small PP is", colorCyan, oneReadOfSmallerPP, colorReset, "and, slightly on the higher side of", workPiece,
								"* that we found a PP of", colorCyan, ProspectivePHitOnLargeSide, colorReset, "a difference of", diffOfLarger)
							updateOutput2(fmt.Sprintf("\nsmall PP is %d and, slightly on the higher side of %d * that we found a PP of %d a difference of %d\n", 
								oneReadOfSmallerPP, workPiece, ProspectivePHitOnLargeSide, diffOfLarger))
			
							result := float64(rootOfProspectivePHitOnLargeSide)/float64(oneReadOfSmallerRoot)
							
							fmt.Println("the ", radical2or3, " root of ", workPiece, " is calculated as ", colorGreen,
								result, colorReset)
							updateOutput2(fmt.Sprintf("\nthe %d root of %d is calculated as %0.9f \n", 
								radical2or3, workPiece, result))
			
							pdiff := (float64(diffOfLarger) / float64(ProspectivePHitOnLargeSide)) * 100000
							fmt.Printf("with pdiff of %0.4f \n", pdiff )
							updateOutput2(fmt.Sprintf("with pdiff of %0.4f \n", pdiff ))
							if pdiff < 0.0002 {
								updateOutput2("\npdiff was less than 0.001 so we are calling it\n")
								sortedResults = append(sortedResults, Results{result: result, pdiff: pdiff}) // collect results into a slice that will eventually be sorted
								return true // calledItEarly is true
							} 
							
							// save the result to an accumulator array so we can Fprint all such hits at the very end
							// List_of_2_results_case18 = append(List_of_2_results_case18, float64(rootOfProspectivePHitOnLargeSide) / float64(oneReadOfSmallerRoot) )
							// corresponding_diffs = append(corresponding_diffs, diffOfLarger)
							// diffs_as_percent = append(diffs_as_percent, float64(diffOfLarger)/float64(ProspectivePHitOnLargeSide))
			
							// ***** ^^^^ ****** the preceding was replaced with the following five lines *******************************************
							
							// in the next five lines we load (append) a record into/to the slice of Results
							Result1 := Results{
								result: float64(rootOfProspectivePHitOnLargeSide) / float64(oneReadOfSmallerRoot),
								pdiff:  float64(diffOfLarger) / float64(ProspectivePHitOnLargeSide),
							}
							sortedResults = append(sortedResults, Result1)
			
							t2 := time.Now()
							elapsed2 := t2.Sub(startBeforeCall)
							// if needed, notify the user that we are still working
							Tim_win = 0.178
							if radical2or3 == 3 {
								if workPiece > 13 {
									Tim_win = 0.0012
								} else {
									Tim_win = 0.003
								}
							}
							if elapsed2.Seconds() > Tim_win {
								fmt.Println(elapsed2.Seconds(), "Seconds have elapsed ... working ...\n")
								updateOutput2(fmt.Sprintf("\n%0.4f Seconds have elapsed ... working ...\n\n", elapsed2.Seconds()))
							}
						}
		
					// smaller side section: ----------------------------------------------------------------------------------------------------------------------------------------
					// ---------------------------------------------------------------------------------------------------------------------------------------------------------------
						if diffOfSmaller < precisionOfRoot { // report the prospects, their differences, and the calculated result for the Sqrt or Cbrt
							fmt.Println("small PP is", colorCyan, oneReadOfSmallerPP, colorReset, "and, slightly on the lesser side of", workPiece,
								"* that we found a PP of", colorCyan, ProspectivePHitOnSmallerSide, colorReset, "a difference of", diffOfSmaller)
							updateOutput2(fmt.Sprintf("\nsmall PP is %d and, slightly on the higher side of %d * that we found a PP of %d a difference of %d\n", 
								oneReadOfSmallerPP, workPiece, ProspectivePHitOnSmallerSide, diffOfSmaller))
							
							result := float64(rootOfProspectivePHitOnSmallerSide)/float64(oneReadOfSmallerRoot)
							
							fmt.Println("the ", radical2or3, " root of ", workPiece, " is calculated as ", colorGreen,
								result, colorReset)
							updateOutput2(fmt.Sprintf("\nthe %d root of %d is calculated as %0.9f \n", 
								radical2or3, workPiece, result))
			
							pdiff := (float64(diffOfSmaller) / float64(ProspectivePHitOnSmallerSide)) * 100000 // even within an if block, variables are local 
							fmt.Printf("with pdiff of %0.4f \n", pdiff )
							updateOutput2(fmt.Sprintf("with pdiff of %0.4f \n", pdiff ))
							if pdiff < 0.0002 {
								updateOutput2("\npdiff was less than 0.001 so we are calling it\n")
								sortedResults = append(sortedResults, Results{result: result, pdiff: pdiff}) // collect results into a slice that will eventually be sorted
								return true // calledItEarly is true
							}
							
							// save the result to three accumulator arrays so we can Fprint all such hits, diffs, and p-diffs, at the very end of run
							// List_of_2_results_case18 = append(List_of_2_results_case18, float64(rootOfProspectivePHitOnSmallerSide) / float64(oneReadOfSmallerRoot) )
							// corresponding_diffs = append(corresponding_diffs, diffOfSmaller)
							// diffs_as_percent = append(diffs_as_percent, float64(diffOfSmaller)/float64(ProspectivePHitOnSmallerSide))
							
							// ***** ^^^^ ****** the preceding was replaced with the following five lines *******************************************
			
							// in the next five lines we load (append) a record into/to the slice of Results
							Result1 := Results{
								result: float64(rootOfProspectivePHitOnSmallerSide) / float64(oneReadOfSmallerRoot),
								pdiff:  float64(diffOfSmaller) / float64(ProspectivePHitOnSmallerSide),
							}
							sortedResults = append(sortedResults, Result1)
			
							t2 := time.Now()
							elapsed2 := t2.Sub(startBeforeCall)
							// if needed, notify the user that we are still working
							Tim_win = 0.178
							if radical2or3 == 3 {
								if workPiece > 13 {
									Tim_win = 0.0012
								} else {
									Tim_win = 0.003
								}
							}
							if elapsed2.Seconds() > Tim_win {
								fmt.Println(elapsed2.Seconds(), "Seconds have elapsed ... working ...\n")
								updateOutput2(fmt.Sprintf("\n%0.4f Seconds have elapsed ... working ...\n\n", elapsed2.Seconds()))
							}
						}
					break breakOutLabel2
				}
			}
		}
	return false
} 

// handlePerfectSquaresAndCubes reports/logs perfect squares/cubes to file and UI
func handlePerfectSquaresAndCubes(TimeOfStartFromTop time.Time, radical2or3, workPiece int, mgr *TrafficManager) {
	if diffOfLarger == 0 || diffOfSmaller == 0 {
		t_s1 := time.Now()
		elapsed_s1 := t_s1.Sub(TimeOfStartFromTop)

		fileHandle, err1 := os.OpenFile("dataLog-From_calculate-pi-and-friends.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		check(err1)
		defer fileHandle.Close()

		Hostname, _ := os.Hostname()
		fmt.Fprintf(fileHandle, "\n  -- %d root of %d by a ratio of PerfectProducts -- selection #%d on %s \n", radical2or3, workPiece, 1, Hostname)
		fmt.Fprint(fileHandle, "was run on: ", time.Now().Format(time.ANSIC), "\n")
		fmt.Fprintf(fileHandle, "Total run was %s \n ", elapsed_s1.String())

		if radical2or3 == 2 {
			result := fmt.Sprintf("Perfect square: %0.0f is the %d root of %d", mathSqrtCheat, radical2or3, workPiece)
			updateOutput2(result)
			fmt.Fprintf(fileHandle, "the %d root of %d is %0.0f \n", radical2or3, workPiece, mathSqrtCheat)
		}
		if radical2or3 == 3 {
			result := fmt.Sprintf("Perfect cube: %0.0f is the %d root of %d", mathCbrtCheat, radical2or3, workPiece)
			updateOutput2(result)
			fmt.Fprintf(fileHandle, "the %d root of %d is %0.0f \n", radical2or3, workPiece, mathCbrtCheat)
		}
	}
}


// setPrecisionForSquareOrCubeRoot adjusts precision based on radical and workPiece  ::: setting the optimal precision in this way is a crude kluge
func setPrecisionForSquareOrCubeRoot(radical2or3, workPiece int) { // ::: - -
	//
	// exhaustive trials have proven that these three precision levels are optimal for these special cases of doing cube roots ...
		if radical2or3 == 3 { 
			if workPiece > 4 { // unless overridden below, precision will be 1700 when doing all cube roots 
				precisionOfRoot = 1700
				fmt.Println("\n Default precision is 1700 \n")
				updateOutput2(fmt.Sprintf("\n Default precision is 1700 \n"))
			}
			if workPiece == 2 || workPiece == 11 || workPiece == 17 { // ::: the logic expressed in this func is far more complex than is apparent at first glance
				precisionOfRoot = 600
				fmt.Println("\n resetting precision to 600 \n")
				updateOutput2(fmt.Sprintf("\n resetting precision to 600 \n"))
			}
			if workPiece == 3 || workPiece == 4 || workPiece == 14 {
				precisionOfRoot = 900
				fmt.Println("\n resetting precision to 900 \n")
				updateOutput2(fmt.Sprintf("\n resetting precision to 900 \n"))
			}
		}
	// ... while a precision level of 4 has been found to be optimal for doing ALL square roots, ergo, this we now do
	if radical2or3 == 2 {
		precisionOfRoot = 4 // squares are so two-dimensional 
	}
}


// build a slice containing 825,000 pairs: a table/slice of ::: perfect squares or cubes, depending on radical 2 or 3 
func buildPairsSlice(radical2or3 int) { // ::: - -
	pairsSlice = nil // Clear/reset the slice between runs
	//
	r := radical2or3
		var identityProduct int // descriptive names
		var sideLengthIsRoot = 2 // 2 being the smallest possible whole-number/perfect root, i.e., it's the square root of 4 & the cube root of 8 // I used to have this as root := 10 but I do not recall why : (how I had decided on 10?)
	p := identityProduct // tight names, p:identityProduct
	s := sideLengthIsRoot // ... s:sideLengthIsRoot
	// 
	for i := 0; i < 825000; i++ {
		s++
		if r == 3 {          // ::: perfect cubes for cube root
			p = s * s * s
		}
		if r == 2 {          // ::: perfect squares for square root 
			p = s * s
		}
		pairsSlice = append(pairsSlice, Pairs{ // stuff identityProduct:s and sideLengthIsRoot:s into pairsSlice (into the corresponding fields: product and root) 
			product: p,
			root:  s,
		})
	}
}
/* ::: and the struct for the pairs looks like this: 
type Pairs struct {
	product int
	root int
}
 */