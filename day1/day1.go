// for each value
// 	divide by 3
// 	round down
// 	subtract 2
// done
//
// ===========================================================================

package main

import (
	"fmt"
)

func main() {
	var a = [100]int{89822, 149236, 106135, 147663, 91417, 59765, 66470, 121156, 148632, 116660, 90316, 111666, 142111, 72595, 139673, 145157, 77572, 83741, 79815, 74693, 139077, 106066, 125817, 127827, 103884, 147289, 81588, 142651, 69916, 147214, 71501, 130067, 60182, 139195, 115502, 127751, 95013, 73411, 125294, 79809, 118110, 122547, 145141, 72231, 138853, 108119, 139960, 128665, 107228, 73416, 54608, 63811, 72363, 130546, 61055, 56786, 127718, 144953, 149284, 137318, 109566, 112866, 148063, 130570, 67536, 84011, 123795, 128098, 51687, 83758, 59867, 103122, 77339, 72126, 71446, 67162, 112342, 120248, 137629, 135736, 139781, 92512, 105922, 85458, 148571, 51173, 135047, 110175, 93722, 82611, 128288, 125225, 104177, 115081, 78470, 96167, 138445, 117778, 100133, 140047}
	var sum int
	var runningtotal int
	var zerocheck bool
	var current int
	var part1total int

	// part 1

	for i := 0; i < len(a); i++ {
		current = a[i]
		sum = current / 3
		sum = sum - 2
		part1total = part1total + sum
	}

	// part 2

	for item := 0; item < len(a); item++ {
		zerocheck = true
		current = a[item]
		for zerocheck == true {
			sum = current / 3
			sum = sum - 2
			if sum > 0 {
				runningtotal = runningtotal + sum
				current = sum
			} else {
				zerocheck = false
			}
		}
	}
	fmt.Println("Part 1 answer is", part1total)
	fmt.Println("Part 2 answer is", runningtotal)

}