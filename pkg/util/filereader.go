package util

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func Must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Read list of integers from file specified by `filename`
func ReadIntegersFromFile(filename string) []int {
	file, err := os.Open(filename)
	Must(err)

	defer func() {
		err = file.Close()
		Must(err)
	}()

	scanner := bufio.NewScanner(file)

	var integersInput []int
	for scanner.Scan() { // internally, it advances token based on sperator
		integerString := scanner.Text()
		integerInt, err := strconv.Atoi(integerString)
		Must(err)
		integersInput = append(integersInput, integerInt)
		Must(err)
	}

	return integersInput
}
