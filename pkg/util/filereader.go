package util

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Read list of integers from file specified by `filename`
func ReadIntegersFromFile(filename string) []int {
	file, err := os.Open(filename)
	must(err)

	defer func() {
		err = file.Close()
		must(err)
	}()

	scanner := bufio.NewScanner(file)

	var integersInput []int
	for scanner.Scan() { // internally, it advances token based on sperator
		integerString := scanner.Text()
		integerInt, err := strconv.Atoi(integerString)
		must(err)
		integersInput = append(integersInput, integerInt)
		must(err)
	}

	return integersInput
}
