// find how many different passwords match betweem puzzle range
//
// It is a six-digit number.
// The value is within the range given in your puzzle input.
// Two adjacent digits are the same (like 22 in 122345).
// Going from left to right, the digits never decrease; they only ever increase or stay the same (like 111123 or 135679).
//
// puzzle range: 136760-595730

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var startingpass int
	var endingpass int
	var runningtotal int
	var dupfound bool
	var ascending bool

	startingpass = 136760
	endingpass = 595730

	// loop through values between start and end range
	for password := startingpass; password <= endingpass; password++ {
		// convert integer to string in order to iterate through digits
		converted := strconv.Itoa(password)
		// compare digits to check for duplicate values
		for digit := 1; digit < len(converted); digit++ {
			if converted[digit] == converted[digit-1] {
				dupfound = true
			}
		}
		if dupfound == true {
			// check for ascending numbers
			for digit := 1; digit < len(converted); digit++ {
				if converted[digit] < converted[digit-1] {
					ascending = false
					break
				} else {
					ascending = true
				}
			}
		}
		if ascending == true {
			runningtotal = runningtotal + 1
			ascending = false
		}
		dupfound = false
	}

	fmt.Println(runningtotal)
}
