package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/leo-ventura/advent-of-code-2021/pkg/util"
)

type command struct {
	cmd string
	mov int
}

func readCommands(filename string) []command {
	file, err := os.Open(filename)
	util.Must(err)

	defer func() {
		err = file.Close()
		util.Must(err)
	}()

	scanner := bufio.NewScanner(file)

	var commands []command
	for scanner.Scan() { // internally, it advances token based on sperator
		line := strings.Split(scanner.Text(), " ")
		cmd := line[0]
		movementString := line[1]
		movement, err := strconv.Atoi(movementString)
		util.Must(err)

		command := command{cmd, movement}

		commands = append(commands, command)
	}

	return commands
}

func main() {
	commandsInput := readCommands("day-02/input.txt")

	horizontalMov := 0
	depthMov := 0
	for _, command := range commandsInput {
		switch command.cmd {
		case "forward":
			horizontalMov += command.mov
		case "down":
			depthMov += command.mov
		case "up":
			depthMov -= command.mov
		}
	}

	fmt.Println(horizontalMov * depthMov)
}
