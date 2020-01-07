// find how many different passwords match betweem puzzle range
//
// part 1 rules
// It is a six-digit number.
// The value is within the range given in your puzzle input.
// Two adjacent digits are the same (like 22 in 122345).
// Going from left to right, the digits never decrease; they only ever increase or stay the same (like 111123 or 135679).
//
// part 2 additional rules
// 112233 meets these criteria because the digits never decrease and all repeated digits are exactly two digits long.
// 123444 no longer meets the criteria (the repeated 44 is part of a larger group of 444).
// 111122 meets the criteria (even though 1 is repeated more than twice, it still contains a double 22).
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
	var largergroup int

	startingpass = 136760
	endingpass = 595730
	//startingpass = 136760
	//endingpass = 137000

	// loop through values between start and end range
	for password := startingpass; password <= endingpass; password++ {
		// convert integer to string in order to iterate through digits
		converted := strconv.Itoa(password)
		largergroup = 0
		dupfound = false
		// compare digits to check for duplicate values and confirm if part of a larger group
		for digit := 1; digit < len(converted); digit++ {
			if converted[digit] == converted[digit-1] {
				largergroup = largergroup + 1
			} else {
				if largergroup == 1 {
					dupfound = true
				} else {
					largergroup = 0
				}

			}
			if digit == (len(converted) - 1) {
				if largergroup == 1 {
					dupfound = true
				}
			}
		}
		// 112233 meets these criteria because the digits never decrease and all repeated digits are exactly two digits long.
		// 123444 no longer meets the criteria (the repeated 44 is part of a larger group of 444).
		// 111122 meets the criteria (even though 1 is repeated more than twice, it still contains a double 22).
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
			fmt.Println(password)
		}
	}
	fmt.Println(runningtotal)
}
