package main

import (
	"fmt"
	"math"

	"github.com/leo-ventura/advent-of-code-2021/pkg/util"
)

func main() {
	binaryValues := util.ReadStringsFromFile("day-03/input.txt")
	nDigits := len(binaryValues[0])

	// Assumes every binary value will have the same number of digits
	// Maps every entry in the array to the number of occurrences of the digit 1
	onesOccurrences := make([]int, nDigits)
	for _, value := range binaryValues {
		for pos, digit := range value {
			if digit == '1' {
				onesOccurrences[pos]++
			}
		}
	}

	gammaRate := 0
	epsilonRate := 0
	power := nDigits - 1
	requiredMajority := len(binaryValues) / 2
	for _, oneAmount := range onesOccurrences {
		if oneAmount > requiredMajority {
			gammaRate += int(math.Pow(2, float64(power)))
		} else {
			epsilonRate += int(math.Pow(2, float64(power)))
		}

		power--
	}

	fmt.Println(gammaRate * epsilonRate)
}
