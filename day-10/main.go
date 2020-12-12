package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
)

const inputFilePath = "input.txt"

// const inputFilePath = "example-data.txt"

func main() {
	log.Println("Loading:", inputFilePath)
	data, err := readLines(inputFilePath)
	if err != nil {
		panic(err)
	}

	log.Println("Found", len(data), "lines of input data")

	sort.Ints(data)

	data = append([]int{0}, data...)         // add first
	data = append(data, data[len(data)-1]+3) // add biggest plus 3

	joltCount1 := 0
	joltCount3 := 0
	combiCount := make(map[int]int)
	combiCount[0] = 1
	log.Println(data)
	for i := 1; i < len(data); i++ {
		if data[i]-data[i-1] == 1 {
			joltCount1++
		} else if data[i]-data[i-1] == 3 {
			joltCount3++
		} else {
			log.Println("should not exist according to the rules")
		}

		// Accumilate existing combinatinos forward
		for x := -3; x <= -1; x++ {
			combiCount[data[i]] += combiCount[data[i]+x]
		}
	}

	log.Println("Result", joltCount1*joltCount3)
	log.Println("Result part 2:", combiCount[data[len(data)-1]])
}

func readLines(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, _ := strconv.Atoi(scanner.Text())
		lines = append(lines, number)
	}
	return lines, scanner.Err()
}
