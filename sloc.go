package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"
)

// @formatter:off

func countAndLogSLOC() { // ::: - -

	if err := extractAndPrintGoFilenames(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	numberOfFilesExplored := 0

	filenameOfThisFile1 := "/Users/quasar/piAndFriendsGUI/main.go"
	blankLines1, singleComments1, commentBlock11, commentBlock21, commentBlock31, runes11, runes21, runes31, totalLines1, nonEmptyLines1 := reportSLOCstats(filenameOfThisFile1)
	numberOfFilesExplored++
	
	filenameOfThisFile2 := "/Users/quasar/piAndFriendsGUI/constants.go"
	blankLines2, singleComments2, commentBlock12, commentBlock22, commentBlock32, runes12, runes22, runes32, totalLines2, nonEmptyLines2 := reportSLOCstats(filenameOfThisFile2)
	numberOfFilesExplored++

	filenameOfThisFile3 := "/Users/quasar/piAndFriendsGUI/Roots.go"
	blankLines3, singleComments3, commentBlock13, commentBlock23, commentBlock33, runes13, runes23, runes33, totalLines3, nonEmptyLines3 := reportSLOCstats(filenameOfThisFile3)
	numberOfFilesExplored++

	filenameOfThisFile4 := "/Users/quasar/piAndFriendsGUI/manager.go"
	blankLines4, singleComments4, commentBlock14, commentBlock24, commentBlock34, runes14, runes24, runes34, totalLines4, nonEmptyLines4 := reportSLOCstats(filenameOfThisFile4)
	numberOfFilesExplored++
	
	filenameOfThisFile5 := "/Users/quasar/piAndFriendsGUI/sloc.go"
	blankLines5, singleComments5, commentBlock15, commentBlock25, commentBlock35, runes15, runes25, runes35, totalLines5, nonEmptyLines5 := reportSLOCstats(filenameOfThisFile5)
	numberOfFilesExplored++
	
	filenameOfThisFile6 := "/Users/quasar/piAndFriendsGUI/functions.go"
	blankLines6, singleComments6, commentBlock16, commentBlock26, commentBlock36, runes16, runes26, runes36, totalLines6, nonEmptyLines6 := reportSLOCstats(filenameOfThisFile6)
	numberOfFilesExplored++
	
	filenameOfThisFile7 := "/Users/quasar/piAndFriendsGUI/globals.go"
	blankLines7, singleComments7, commentBlock17, commentBlock27, commentBlock37, runes17, runes27, runes37, totalLines7, nonEmptyLines7 := reportSLOCstats(filenameOfThisFile7)
	numberOfFilesExplored++
	
	filenameOfThisFile8 := "/Users/quasar/piAndFriendsGUI/Archimedes.go"
	blankLines8, singleComments8, commentBlock18, commentBlock28, commentBlock38, runes18, runes28, runes38, totalLines8, nonEmptyLines8 := reportSLOCstats(filenameOfThisFile8)
	numberOfFilesExplored++
	
	filenameOfThisFile9 := "/Users/quasar/piAndFriendsGUI/Chud.go"
	blankLines9, singleComments9, commentBlock19, commentBlock29, commentBlock39, runes19, runes29, runes39, totalLines9, nonEmptyLines9 := reportSLOCstats(filenameOfThisFile9)
	numberOfFilesExplored++
	
	filenameOfThisFile10 := "/Users/quasar/piAndFriendsGUI/coloredButton.go"
	blankLines10, singleComments10, commentBlock110, commentBlock210, commentBlock310, runes110, runes210, runes310, totalLines10, nonEmptyLines10 := reportSLOCstats(filenameOfThisFile10)
	numberOfFilesExplored++
	
	filenameOfThisFile11 := "/Users/quasar/piAndFriendsGUI/go.mod"
	blankLines11, singleComments11, commentBlock111, commentBlock211, commentBlock311, runes111, runes211, runes311, totalLines11, nonEmptyLines11 := reportSLOCstats(filenameOfThisFile11)
	numberOfFilesExplored++
	
	filenameOfThisFile12 := "/Users/quasar/piAndFriendsGUI/Wallis.go"
	blankLines12, singleComments12, commentBlock112, commentBlock212, commentBlock312, runes112, runes212, runes312, totalLines12, nonEmptyLines12 := reportSLOCstats(filenameOfThisFile12)
	numberOfFilesExplored++
	
	filenameOfThisFile13 := "/Users/quasar/piAndFriendsGUI/BBPmax.go"
	blankLines13, singleComments13, commentBlock113, commentBlock213, commentBlock313, runes113, runes213, runes313, totalLines13, nonEmptyLines13 := reportSLOCstats(filenameOfThisFile13)
	numberOfFilesExplored++
	
	filenameOfThisFile14 := "/Users/quasar/piAndFriendsGUI/GregoryLeibniz.go"
	blankLines14, singleComments14, commentBlock114, commentBlock214, commentBlock314, runes114, runes214, runes314, totalLines14, nonEmptyLines14 := reportSLOCstats(filenameOfThisFile14)
	numberOfFilesExplored++
	
	filenameOfThisFile15 := "/Users/quasar/piAndFriendsGUI/CustomSeries.go"
	blankLines15, singleComments15, commentBlock115, commentBlock215, commentBlock315, runes115, runes215, runes315, totalLines15, nonEmptyLines15 := reportSLOCstats(filenameOfThisFile15)
	numberOfFilesExplored++
	
	filenameOfThisFile16 := "/Users/quasar/piAndFriendsGUI/Nilakantha.go"
	blankLines16, singleComments16, commentBlock116, commentBlock216, commentBlock316, runes116, runes216, runes316, totalLines16, nonEmptyLines16 := reportSLOCstats(filenameOfThisFile16)
	numberOfFilesExplored++
	
	filenameOfThisFile17 := "/Users/quasar/piAndFriendsGUI/Gauss.go"
	blankLines17, singleComments17, commentBlock117, commentBlock217, commentBlock317, runes117, runes217, runes317, totalLines17, nonEmptyLines17 := reportSLOCstats(filenameOfThisFile17)
	numberOfFilesExplored++
	
	filenameOfThisFile18 := "/Users/quasar/piAndFriendsGUI/window2.go"
	blankLines18, singleComments18, commentBlock118, commentBlock218, commentBlock318, runes118, runes218, runes318, totalLines18, nonEmptyLines18 := reportSLOCstats(filenameOfThisFile18)
	numberOfFilesExplored++
	
	filenameOfThisFile19 := "/Users/quasar/piAndFriendsGUI/Spigot.go"
	blankLines19, singleComments19, commentBlock119, commentBlock219, commentBlock319, runes119, runes219, runes319, totalLines19, nonEmptyLines19 := reportSLOCstats(filenameOfThisFile19)
	numberOfFilesExplored++
	
	filenameOfThisFile20 := "/Users/quasar/piAndFriendsGUI/Monty.go"
	blankLines20, singleComments20, commentBlock120, commentBlock220, commentBlock320, runes120, runes220, runes320, totalLines20, nonEmptyLines20 := reportSLOCstats(filenameOfThisFile20)
	numberOfFilesExplored++

	// ::: was double, now empty 
	filenameOfThisFile21 := "/Users/quasar/piAndFriendsGUI/1_emptyFile.no"
	blankLines21, singleComments21, commentBlock121, commentBlock221, commentBlock321, runes121, runes221, runes321, totalLines21, nonEmptyLines21 := reportSLOCstats(filenameOfThisFile21)
	numberOfFilesExplored++

	filenameOfThisFile22 := "/Users/quasar/piAndFriendsGUI/methods.go"
	blankLines22, singleComments22, commentBlock122, commentBlock222, commentBlock322, runes122, runes222, runes322, totalLines22, nonEmptyLines22 := reportSLOCstats(filenameOfThisFile22)
	numberOfFilesExplored++

	filenameOfThisFile23 := "/Users/quasar/piAndFriendsGUI/ScoreBoardNila.go"
	blankLines23, singleComments23, commentBlock123, commentBlock223, commentBlock323, runes123, runes223, runes323, totalLines23, nonEmptyLines23 := reportSLOCstats(filenameOfThisFile23)
	numberOfFilesExplored++
	
	// numberOfFilesExplored // Used only in countSLOC() and the associated about_app()

	// Total lines of code (exclusive of data).
	totalLines := totalLines1 + totalLines2 + totalLines3 + totalLines4 + totalLines5 + totalLines6 + totalLines7 + totalLines8 + totalLines9 + totalLines10 +
		totalLines11 + totalLines12 + totalLines13 + totalLines14 + totalLines15 + totalLines16 + totalLines17 + totalLines18 + totalLines19 + totalLines20 +
		totalLines21 + totalLines22 + totalLines23

	// Total lines of executable code.
	nonEmptyLines := nonEmptyLines1 + nonEmptyLines2 + nonEmptyLines3 + nonEmptyLines4 + nonEmptyLines5 + nonEmptyLines6 + nonEmptyLines7 + nonEmptyLines8 + nonEmptyLines9 + nonEmptyLines10 +
		nonEmptyLines11 + nonEmptyLines12 + nonEmptyLines13 + nonEmptyLines14 + nonEmptyLines15 + nonEmptyLines16 + nonEmptyLines17 + nonEmptyLines18 + nonEmptyLines19 + nonEmptyLines20 +
		nonEmptyLines21 + nonEmptyLines22 + nonEmptyLines23

	// Blank lines and comment lines.
	blankLinesTotal := blankLines1 + blankLines2 + blankLines3 + blankLines4 + blankLines5 + blankLines6 + blankLines7 + blankLines8 + blankLines9 + blankLines10 +
		blankLines11 + blankLines12 + blankLines13 + blankLines14 + blankLines15 + blankLines16 + blankLines17 + blankLines18 + blankLines19 + blankLines20 +
		blankLines21 + blankLines22 + blankLines23

	singleCommentsTotal := singleComments1 + singleComments2 + singleComments3 + singleComments4 + singleComments5 + singleComments6 + singleComments7 + singleComments8 + singleComments9 + singleComments10 +
		singleComments11 + singleComments12 + singleComments13 + singleComments14 + singleComments15 + singleComments16 + singleComments17 + singleComments18 + singleComments19 + singleComments20 +
		singleComments21 + singleComments22 + singleComments23

	commentBlock1Total := commentBlock11 + commentBlock12 + commentBlock13 + commentBlock14 + commentBlock15 + commentBlock16 + commentBlock17 + commentBlock18 + commentBlock19 + commentBlock110 +
		commentBlock111 + commentBlock112 + commentBlock113 + commentBlock114 + commentBlock115 + commentBlock116 + commentBlock117 + commentBlock118 + commentBlock119 + commentBlock120 +
		commentBlock121 + commentBlock122 + commentBlock123

	commentBlock2Total := commentBlock21 + commentBlock22 + commentBlock23 + commentBlock24 + commentBlock25 + commentBlock26 + commentBlock27 + commentBlock28 + commentBlock29 + commentBlock210 +
		commentBlock211 + commentBlock212 + commentBlock213 + commentBlock214 + commentBlock215 + commentBlock216 + commentBlock217 + commentBlock218 + commentBlock219 + commentBlock220 +
		commentBlock221 + commentBlock222 + commentBlock223

	commentBlock3Total := commentBlock31 + commentBlock32 + commentBlock33 + commentBlock34 + commentBlock35 + commentBlock36 + commentBlock37 + commentBlock38 + commentBlock39 + commentBlock310 +
		commentBlock311 + commentBlock312 + commentBlock313 + commentBlock314 + commentBlock315 + commentBlock316 + commentBlock317 + commentBlock318 + commentBlock319 + commentBlock320 +
		commentBlock321 + commentBlock322 + commentBlock323

	runes1Total := runes11 + runes12 + runes13 + runes14 + runes15 + runes16 + runes17 + runes18 + runes19 + runes110 +
		runes111 + runes112 + runes113 + runes114 + runes115 + runes116 + runes117 + runes118 + runes119 + runes120 +
		runes121 + runes122 + runes123

	runes2Total := runes21 + runes22 + runes23 + runes24 + runes25 + runes26 + runes27 + runes28 + runes29 + runes210 +
		runes211 + runes212 + runes213 + runes214 + runes215 + runes216 + runes217 + runes218 + runes219 + runes220 +
		runes221 + runes222 + runes223

	runes3Total := runes31 + runes32 + runes33 + runes34 + runes35 + runes36 + runes37 + runes38 + runes39 + runes310 +
		runes311 + runes312 + runes313 + runes314 + runes315 + runes316 + runes317 + runes318 + runes319 + runes320 +
		runes321 + runes322 + runes323
	
	/*
		// totalLines is ::: The Total lines of Code (exclusive of data) -- assume this means data files such as are found in Jap Language apps
		totalLines := totalLines1 + totalLines2 + totalLines3 + totalLines4 + totalLines5 + totalLines6 + totalLines7 + totalLines8 + totalLines9 + totalLines10 + totalLines11 + totalLines12 + totalLines13 + totalLines14 +
		+ totalLines15 + totalLines16 + totalLines17 + totalLines18 + totalLines19 + totalLines20 + totalLines21 + totalLines22


		// ::: the Total lines of executable Code:
		nonEmptyLines := nonEmptyLines1 + nonEmptyLines2 + nonEmptyLines3 + nonEmptyLines4 + nonEmptyLines5 + nonEmptyLines6 + nonEmptyLines7 + nonEmptyLines8 + nonEmptyLines9 + nonEmptyLines10 + nonEmptyLines11 + nonEmptyLines12 +
			nonEmptyLines13 + nonEmptyLines14 + nonEmptyLines15 + nonEmptyLines16 + nonEmptyLines17 + nonEmptyLines18 + nonEmptyLines19 + nonEmptyLines20 + nonEmptyLines21 + nonEmptyLines22


		// black lines and all forms of comment lines :
		blankLinesTotal := blankLines1 + blankLines2 + blankLines3 + blankLines4 + blankLines5 + blankLines6 + blankLines7 + blankLines8 + blankLines9 + blankLines10 + blankLines11 + blankLines12 + blankLines13 + blankLines14 +
			blankLines15 + blankLines16 + blankLines17 + blankLines18 + blankLines19 + blankLines20 + blankLines21 + blankLines22

		singleCommentsTotal := singleComments1 + singleComments2 + singleComments3 + singleComments4 + singleComments5 + singleComments6 + singleComments7 + singleComments8 + singleComments9 + singleComments10 + singleComments11 +
			singleComments12 + singleComments13 + singleComments14 + singleComments15 + singleComments16 + singleComments17 + singleComments18 + singleComments19 + singleComments20 + singleComments21 + singleComments22

		commentBlock1Total := commentBlock11 + commentBlock12 + commentBlock13 + commentBlock14 + commentBlock15 + commentBlock16 + commentBlock17 + commentBlock18 + commentBlock19 + commentBlock110 +
			commentBlock111 + commentBlock112 + commentBlock113 + commentBlock114 + commentBlock115 + commentBlock116 + commentBlock117 + commentBlock118 + commentBlock119 + commentBlock120 + commentBlock121 + commentBlock122

		commentBlock2Total := commentBlock21 + commentBlock22 + commentBlock23 + commentBlock24 + commentBlock25 + commentBlock26 + commentBlock27 + commentBlock28 + commentBlock29 + commentBlock210 + commentBlock211 + commentBlock212 +
		commentBlock213 + commentBlock214 + commentBlock215 + commentBlock216 + commentBlock217 + commentBlock218 + commentBlock219 + commentBlock220 + commentBlock221 + commentBlock222

		commentBlock3Total := commentBlock31 + commentBlock32 + commentBlock33 + commentBlock34 + commentBlock35 + commentBlock36 + commentBlock37 + commentBlock38 + commentBlock39 + commentBlock310 + commentBlock311 + commentBlock312 +
			commentBlock313 + commentBlock314 + commentBlock315 + commentBlock316 + commentBlock317 + commentBlock318 + commentBlock319 + commentBlock320 + commentBlock321 + commentBlock322

		runes1Total := runes11 + runes12 + runes13 + runes14 + runes15 + runes16 + runes17 + runes18 + runes19 + runes110 + runes111 +
			runes112 + runes113 + runes114 + runes115 + runes116 + runes117 + runes118 + runes119 + runes120 + runes121 + runes122
		runes2Total := runes21 + runes22 + runes23 + runes24 + runes25 + runes26 + runes27 + runes28 + runes29 + runes210 + runes211 +
			runes212 + runes213 + runes214 + runes215 + runes216 + runes217 + runes218 + runes219 + runes220 + runes221 + runes222
		runes3Total := runes31 + runes32 + runes33 + runes34 + runes35 + runes36 + runes37 + runes38 + runes39 + runes310 + runes311 +
			runes312 + runes313 + runes314 + runes315 + runes316 + runes317 + runes318 + runes319 + runes320 + runes321 + runes322

	*/

	fmt.Printf("\nrune1-22: %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d", runes11, runes12, runes13, runes14, runes15, runes16, runes17, runes18, runes19, runes110, runes111 +
		runes112, runes113, runes114, runes115, runes116, runes117, runes118, runes119, runes120, runes121, runes122)

	fmt.Printf("\nrune2-22: %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d", runes21, runes22, runes23, runes24, runes25, runes26, runes27, runes28, runes29, runes210, runes211 +
		runes212, runes213, runes214, runes215, runes216, runes217, runes218, runes219, runes220, runes221, runes222)

	fmt.Printf("\nrune3-22: %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d\n", runes31, runes32, runes33, runes34, runes35, runes36, runes37, runes38, runes39, runes310, runes311 +
		runes312, runes313, runes314, runes315, runes316, runes317, runes318, runes319, runes320, runes321, runes322)

	// grand total is really just of non-coding lines 
	grandTotal := blankLinesTotal + singleCommentsTotal +
		commentBlock1Total + commentBlock2Total + commentBlock3Total +
		runes1Total + runes2Total + runes3Total

	sumOfCodePlusNon := grandTotal + nonEmptyLines // the actual grand total 

	fileHandle, err := os.OpenFile("slocLog.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	check_error(err)

	currentTime := time.Now()

	_, err29 := fmt.Fprintf(fileHandle, "\nWhen the App began at: %s", currentTime.Format("15:04:05 on Monday 01-02-2006"))
	check_error(err29)
	_, err1 := fmt.Fprintf(fileHandle, "\nThe Total lines of Code (exclusive of data) = %d t-SLOC\n", totalLines)
	check_error(err1)
	_, err19 := fmt.Fprintf(fileHandle, "and, the Total lines of executable Code = %d e-SLOC\n\n", nonEmptyLines)
	check_error(err19)

	fmt.Printf("Total lines of Code (exclusive of data) = %d t-SLOC\n\n", totalLines)

	fmt.Printf("Blank lines: %d\n\n", blankLinesTotal)

	fmt.Printf("Single comment lines: %d\n\n", singleCommentsTotal)


	fmt.Printf("Total lines of executable Code = %d e-SLOC\n\n", nonEmptyLines)
	fmt.Printf("BlnkLns:%d + SnglCmLns:%d + ComBlks:%d + runes:%d = total of cmnts+spc = %d lines\n\n", blankLinesTotal, singleCommentsTotal, commentBlock2Total, runes2Total, grandTotal)

	fmt.Printf("Total of comments etc. in %d files + e-SLOC = %d = t-SLOC\n\n", numberOfFilesExplored, sumOfCodePlusNon)

	if runes3Total > 0 || runes1Total > 0 || commentBlock3Total > 0 || commentBlock1Total > 0 { // if any of these was > 0
		fmt.Println("\n\n === hey we actually got something from where there should not have been anything === \n\n")
	}

}

/*
.
*/
func reportSLOCstats(filepath string) (blankLines, singleComments, commentBlock1, commentBlock2, commentBlock3, runes1, runes2, runes3, totalLines, sloc int) { // ::: - -
	// Patterns to identify comments, blank lines, and strings
	singleLineCommentPattern := `^\s*//`
	multiLineCommentPattern := `(?s)/\*.*?\*/`
	blankLinePattern := `^\s*$`
	stringLiteralPattern := `(?s)"(?:\\.|[^\\"])*?"|` + "`(?:\\.|[^`])*?`"

	// Compile regular expressions
	singleLineCommentRe := regexp.MustCompile(singleLineCommentPattern)
	multiLineCommentRe := regexp.MustCompile(multiLineCommentPattern)
	blankLineRe := regexp.MustCompile(blankLinePattern)
	stringLiteralRe := regexp.MustCompile(stringLiteralPattern)

	// Open the file
	file, err := os.Open(filepath)
	if err != nil {
		// return 0, 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalLines = 0
	sloc = 0
	inMultiLineComment := false
	inMultiLineString := false

	for scanner.Scan() {
		line := scanner.Text()
		totalLines++

		// ::: Check for blank lines
		if blankLineRe.MatchString(line) {
			blankLines++
			continue
		}

		// ::: Check for single-line comments
		if singleLineCommentRe.MatchString(line) {
			singleComments++
			continue
		}

		// ::: Check for multi-line comment blocks
		if inMultiLineComment {
			if strings.Contains(line, "*/") {
				inMultiLineComment = false
				line = multiLineCommentRe.ReplaceAllString(line, "")
				if blankLineRe.MatchString(line) || singleLineCommentRe.MatchString(line) {
					commentBlock1++ // Does not normally accumulate anything.
					continue
				}
			} else {
				commentBlock2++ // This is where we find lines that match.
				continue
			}
		}
		if strings.Contains(line, "/*") {
			inMultiLineComment = true
			line = multiLineCommentRe.ReplaceAllString(line, "")
			if blankLineRe.MatchString(line) || singleLineCommentRe.MatchString(line) { // blankLines, singleComments, commentBlock1, commentBlock2, commentBlock3, runes1, runes2, runes3
				commentBlock3++ // Does not normally accumulate anything.
				continue
			}
		}

		// ::: Check for multi-line strings // string literals // Runes
		if inMultiLineString {
			if strings.Count(line, "`")%2 != 0 || strings.Count(line, "\"")%2 != 0 {
				inMultiLineString = false
				line = stringLiteralRe.ReplaceAllString(line, "")
				if blankLineRe.MatchString(line) || singleLineCommentRe.MatchString(line) {
					runes1++ // Does not normally accumulate anything.
					continue
				}
			} else {
				runes2++ // This is where we find lines that match.
				continue
			}
		}
		if strings.Count(line, "`")%2 != 0 || strings.Count(line, "\"")%2 != 0 {
			inMultiLineString = true
			line = stringLiteralRe.ReplaceAllString(line, "")
			if blankLineRe.MatchString(line) || singleLineCommentRe.MatchString(line) {
				runes3++ // Does not normally accumulate anything.
				continue
			}
		}

		sloc++
	}

	if err := scanner.Err(); err != nil {
		// return 0, 0, err
	}

	return blankLines, singleComments, commentBlock1, commentBlock2, commentBlock3, runes1, runes2, runes3, totalLines, sloc
}

// Creates a func named check_error which takes one parameter 'e' of type error
func check_error(e error) { // ::: - -
	if e != nil {
		panic(e) // use panic() to display error code
	}
}
