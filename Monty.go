package main

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
	"time"
)

// @formatter:off

func Monty(fyneFunc func(string), gridSizeAsString string, done chan bool) {
	// Produce an alternate string suitable for printing, with commas every three digits from the right
		withCommas := ""
		for i, char := range gridSizeAsString {
			select {
			case <-done: // ::: here an attempt is made to read from the channel (a closed channel can be read from successfully; but what is read will be the null/zero value of the type of chan (0, false, "", 0.0, etc.)
				// in the case of this particular channel (which is of type bool) we get the value false from having received from the channel when it is already closed. 
				// ::: if the channel known by the moniker "done" is already closed, that/it is to be interpreted as the abort signal by all listening processes. 
				fmt.Println("Goroutine Monty for-loop (1 of 2) is being terminated by select case finding the done channel to be already closed")
				return // Exit the goroutine
			default:
			if i > 0 && (len(gridSizeAsString)-i)%3 == 0 {
				withCommas += ","
			}
			withCommas += string(char)
			}
		}
		// ::: screen
		fyneFunc(fmt.Sprintf("\n\nSize of the grid has been set to: %s\n", withCommas))

	// convert gridSize to an int
	gridSize, err := strconv.Atoi(gridSizeAsString)
	if err != nil {
		fyneFunc(fmt.Sprintf("Invalid input, please enter an integer in string form."))
		return
	}
		// ::: screen
		if gridSize < 5 {
			fyneFunc(fmt.Sprintf("\n A grid that small makes me puke! \n"))
			return
		}
		if gridSize > 6000 && gridSize <= 8500 {
			fyneFunc(fmt.Sprintf("\n ... working ... expect 15s\n"))
		} else if gridSize > 8500 && gridSize <= 11000 {
			fyneFunc(fmt.Sprintf("\n ... really working ... expect 25s\n"))
		} else if gridSize > 11000 && gridSize <= 12000 {
			fyneFunc(fmt.Sprintf("\n ... I will be working really hard ...expect 30s\n"))
		} else if gridSize > 12000 && gridSize <= 13000 {
			fyneFunc(fmt.Sprintf("\n ... working really really hard...expect 40s\n"))
		} else if gridSize > 13000 && gridSize <= 15000 {
			fyneFunc(fmt.Sprintf("\n ... for so very long, I'll be working ...expect 50s\n"))
		} else if gridSize > 15000 && gridSize <= 18000 {
			fyneFunc(fmt.Sprintf("\n ... Yikes, I'll be working, for too long ...expect 1m5s\n"))
		} else if gridSize > 18000 && gridSize <= 24000 {
			fyneFunc(fmt.Sprintf("\n ... while you take a nap, I'll still be working ... expect 1m25s\n"))
		} else if gridSize > 24000 && gridSize <= 34000 {
			fyneFunc(fmt.Sprintf("\n ... Brace yourself for how long I'll be working ... expect 4min\n"))
		} else if gridSize > 34000 && gridSize <= 100000 {
			fyneFunc(fmt.Sprintf("\n ... Expect 5–15 minutes for ~4–5 digits ...\n"))
			fyneFunc(fmt.Sprintf("\n ... and be advised that 120k, or more, will make me puke! ...\n"))
		} else if gridSize > 100000 && gridSize <= 119999 {
			fyneFunc(fmt.Sprintf("\n ... Working insanely hard, expect 15–30 minutes for ~5 digits ...\n"))
		} else if gridSize > 119999 {
			fyneFunc(fmt.Sprintf("\n ... I have puked! \n"))
			return
		}
		
	piApprox := GridPi(fyneFunc, gridSize, done) // ::: run GridPi < - - - - - - - - - - < -

		// ::: screen
		fyneFunc(fmt.Sprintf("\nSize of the grid was set at: %s\n", withCommas))
		fyneFunc(fmt.Sprintf("Approximated Pi as big float: %s\n", piApprox.Text('f', 30)))
			piApproxFloat64, _ := piApprox.Float64()
		fyneFunc(fmt.Sprintf("Approximated Pi as float64:   %f\n", piApproxFloat64))
			piFromMathLib := math.Pi
			piFromMathLibBF := big.NewFloat(piFromMathLib)
		fyneFunc(fmt.Sprintf("Pi from Math Library:         %s\n", piFromMathLibBF.Text('f', 30)))
		fyneFunc(fmt.Sprintf("Difference: %f\n", math.Abs(piApproxFloat64-math.Pi)))
			_, digitCount := checkPiTo100(piApprox)
		fyneFunc(fmt.Sprintf("\nWe verified Pi to %d digits\n\n", digitCount))
}
/*
.
 */
func GridPi(fyneFunc func(string), gridSize int, done chan bool) *big.Float {
	start := time.Now()
		insideCircle := big.NewInt(0)
		totalPoints := big.NewInt(int64(gridSize * gridSize))
		increment := big.NewFloat(1.0 / float64(gridSize)).SetPrec(256)
		halfIncrement := new(big.Float).Quo(increment, big.NewFloat(2.0)).SetPrec(256)
	for i := 0; i < gridSize; i++ {
		select {
		case <-done: // ::: here an attempt is made to read from the channel (a closed channel can be read from successfully; but what is read will be the null/zero value of the type of chan (0, false, "", 0.0, etc.)
			// in the case of this particular channel (which is of type bool) we get the value false from having received from the channel when it is already closed. 
			// ::: if the channel known by the moniker "done" is already closed, that/it is to be interpreted as the abort signal by all listening processes. 
			fmt.Println("Goroutine Monty for-loop (2 of 2) is being terminated by select case finding the done channel to be already closed")
			return increment // Exit the goroutine ::: We had to return some kind of a big float ... 
		default:
		for j := 0; j < gridSize; j++ {
			// ::: x = (i * increment) + halfIncrement
				x := new(big.Float).SetPrec(256)
				x.Mul(increment, big.NewFloat(float64(i)))
				x.Add(x, halfIncrement)
			// ::: y = (j * increment) + halfIncrement
				y := new(big.Float).SetPrec(256)
				y.Mul(increment, big.NewFloat(float64(j)))
				y.Add(y, halfIncrement)
			xSquared := new(big.Float).Mul(x, x)
			ySquared := new(big.Float).Mul(y, y)
			sum := new(big.Float).Add(xSquared, ySquared)
				if sum.Cmp(big.NewFloat(1.0)) <= 0 {
					insideCircle.Add(insideCircle, big.NewInt(1))
				}
			iterationsForMonte16j = j
		}
		iterationsForMonte16i = i
		}
	}
	iterationsForMonteTotal = iterationsForMonte16j * iterationsForMonte16i
		four := big.NewFloat(4.0).SetPrec(256)
		insideCircleF := new(big.Float).SetPrec(256).SetInt(insideCircle)
		totalPointsF := new(big.Float).SetPrec(256).SetInt(totalPoints)
		piApprox := new(big.Float).SetPrec(256)
		piApprox.Quo(insideCircleF, totalPointsF)
		piApprox.Mul(piApprox, four)
	t := time.Now()
	elapsed := t.Sub(start)
	TotalRun := elapsed.String()
		// ::: put commas into the Total-iterations var
		numStr := strconv.FormatInt(int64(iterationsForMonteTotal), 10)
		result := ""
		for i, char := range numStr {
			if i > 0 && (len(numStr)-i)%3 == 0 {
				result += ","
			}
			result += string(char)
		}
	// ::: screen	
	fyneFunc(fmt.Sprintf("\nTotal iterations: %s \nElapsed time: %s \n", result, TotalRun))
	return piApprox
}
