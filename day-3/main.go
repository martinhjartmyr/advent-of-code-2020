package main

import (
	"bufio"
	"log"
	"os"
)

const inputFilePath = "input.txt"

func main() {
	log.Println("Loading:", inputFilePath)
	data, err := readLines(inputFilePath)
	if err != nil {
		panic(err)
	}

	log.Println("Found", len(data), "lines of input data")

	log.Println("Trees hit part 1:", travelMap(data, 3, 1))

	log.Println(
		"Trees hit part 2:",
		travelMap(data, 1, 1)*travelMap(data, 3, 1)*
			travelMap(data, 5, 1)*travelMap(data, 7, 1)*
			travelMap(data, 1, 2))
}

func travelMap(data []string, xPosInc int, yPosInc int) int {
	xPos, yPos, treesHit := 0, 1, 0 // we skip the first row according to the rules
	mapHeight := len(data)

	for yPos < mapHeight {
		mapRow := data[yPos]

		// travel withing the map
		xPos += xPosInc
		yPos += yPosInc

		// map is growing unlimited in x space and is 31 chars wide
		xPosMap := xPos % 31

		// check if we hit a tree (#)
		char := string(mapRow[xPosMap])
		if char == "#" {
			treesHit++
		}
	}

	return treesHit
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
