package main

import (
	"fmt"

	"github.com/leo-ventura/advent-of-code-2021/pkg/util"
)

func main() {
	integersInput := util.ReadIntegersFromFile("day-01/input.txt")

	countIncreases := 0
	previousEntry := integersInput[len(integersInput)-1]
	for _, entry := range integersInput {
		if entry > previousEntry {
			countIncreases++
		}
		previousEntry = entry
	}

	fmt.Println(countIncreases)
}
