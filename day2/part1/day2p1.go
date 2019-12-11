// position 0 - opcode
// 	- either 1, 2 or 99
// 	1 = addition of two integers read from position 1 and 2 and then stores result in positon identified by position 3
// 	2 = multiply two integers read from position 1 and 2 and then stores result in positon identified by position 3
// 	99 = program end
//
// when done processing move forward 4 postions to the next set of four

// find starting opcode in position 0
// find value in from locations in pos 1, 2 and 3
// run calculation of values from 1 and 2 based on opcode and store in value from 3
// start again at starting opcode plus 4
// loop until hit 99
// when 99 found output value in postion 0

package main

import (
	"fmt"
)

func main() {
	intcode := []int{1, 12, 2, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 10, 1, 19, 1, 5, 19, 23, 1, 23, 5, 27, 1, 27, 13, 31, 1, 31, 5, 35, 1, 9, 35, 39, 2, 13, 39, 43, 1, 43, 10, 47, 1, 47, 13, 51, 2, 10, 51, 55, 1, 55, 5, 59, 1, 59, 5, 63, 1, 63, 13, 67, 1, 13, 67, 71, 1, 71, 10, 75, 1, 6, 75, 79, 1, 6, 79, 83, 2, 10, 83, 87, 1, 87, 5, 91, 1, 5, 91, 95, 2, 95, 10, 99, 1, 9, 99, 103, 1, 103, 13, 107, 2, 10, 107, 111, 2, 13, 111, 115, 1, 6, 115, 119, 1, 119, 10, 123, 2, 9, 123, 127, 2, 127, 9, 131, 1, 131, 10, 135, 1, 135, 2, 139, 1, 10, 139, 0, 99, 2, 0, 14, 0}
	//intcode := []int{1, 1, 1, 4, 99, 5, 6, 0, 99}
	//intcode := []int{1, 0, 0, 0, 99}

	var opcode int
	var startingopcode int
	var pos1, pos2, pos3 int

	startingopcode = 0

	opcode = intcode[startingopcode]
loop:
	for {
		switch opcode {
		case 1:
			// find value in from locations in pos 1, 2 and 3
			// add values from 1 and 2 based on opcode and store in value from 3
			// start again at starting opcode plus 4
			pos1 = intcode[startingopcode+1]
			pos2 = intcode[startingopcode+2]
			pos3 = intcode[startingopcode+3]
			intcode[pos3] = intcode[pos1] + intcode[pos2]
			startingopcode += 4
			opcode = intcode[startingopcode]

		case 2:
			// find value in from locations in pos 1, 2 and 3
			// multiply values from 1 and 2 based on opcode and store in value from 3
			// start again at starting opcode plus 4
			pos1 = intcode[startingopcode+1]
			pos2 = intcode[startingopcode+2]
			pos3 = intcode[startingopcode+3]
			intcode[pos3] = intcode[pos1] * intcode[pos2]
			startingopcode += 4
			opcode = intcode[startingopcode]

		case 99:
			break loop

		default:
			break loop
		}

	}

	fmt.Println("result =", intcode[0])

}
