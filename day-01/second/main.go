package main

import (
	"fmt"

	"github.com/leo-ventura/advent-of-code-2021/pkg/util"
)

const WINDOW = 3

func sumWindow(values []int, startIndex int) int {
	sum := 0
	for i := startIndex; i < startIndex+WINDOW; i++ {
		sum += values[i]
	}
	return sum
}

func main() {
	integersInput := util.ReadIntegersFromFile("day-01/input.txt")

	sumCountIncreases := 0
	previousSum := sumWindow(integersInput, 0)
	var currentSum int
	for i := 3; i <= len(integersInput)-WINDOW; i++ {
		currentSum = sumWindow(integersInput, i)
		if currentSum > previousSum {
			sumCountIncreases++
		}
		previousSum = currentSum
	}

	fmt.Println(sumCountIncreases)
}
