package main

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strconv"
	"strings"
)

const inputFilePath = "input.txt"

type command struct {
	operation string
	argument  int
}

func main() {
	log.Println("Loading:", inputFilePath)
	data, err := readLines(inputFilePath)
	if err != nil {
		panic(err)
	}

	log.Println("Found", len(data), "lines of input data")

	commands := parseCommands(data)
	acc, err := executeCommands(commands)
	log.Println("accumulator:", acc, err)
	accFixed, err := fixBug(commands)
	log.Println("accumulator:", accFixed, err)
}

func executeCommands(commands []command) (int, error) {
	accumulator := 0
	commandHistory := make(map[int]bool)
	crashed := false
	commandPos := 0
	for !crashed {
		command := commands[commandPos]
		if _, ok := commandHistory[commandPos]; ok {
			return accumulator, errors.New("Crashed")
		}
		commandHistory[commandPos] = true
		switch command.operation {
		case "acc":
			accumulator += command.argument
			commandPos++
		case "jmp":
			commandPos += command.argument
		case "nop":
			commandPos++
		}

		if commandPos >= len(commands)-1 {
			return accumulator, nil
		}
	}
	return accumulator, errors.New("Crashed")
}

func fixBug(commands []command) (int, error) {
	for i := 0; i < len(commands); i++ {
		tmpCommands := make([]command, len(commands))
		copy(tmpCommands, commands)
		switch tmpCommands[i].operation {
		case "jmp":
			tmpCommands[i].operation = "nop"
		case "nop":
			tmpCommands[i].operation = "jmp"
		}
		acc, err := executeCommands(tmpCommands)
		if err == nil {
			return acc, nil
		}
	}
	return 0, errors.New("Not fixed")
}

func parseCommands(data []string) []command {
	var commands []command
	for _, line := range data {
		commandArray := strings.Split(line, " ")
		argument, _ := strconv.Atoi(commandArray[1])
		commands = append(commands, command{commandArray[0], argument})
	}

	return commands
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
