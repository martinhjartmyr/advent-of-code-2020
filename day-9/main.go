package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
)

const inputFilePath = "input.txt"

func main() {
	log.Println("Loading:", inputFilePath)
	data, err := readLines(inputFilePath)
	if err != nil {
		panic(err)
	}

	log.Println("Found", len(data), "lines of input data")

	invalidNumber := 0
	for i := 25; i < len(data); i++ {
		preamble := data[i-25 : i]
		valid := false
		for x := 0; x < 25; x++ {
			for y := 0; y < 25; y++ {
				if preamble[x]+preamble[y] == data[i] && preamble[x] != preamble[y] {
					valid = true
					break
				}
			}
			if valid {
				break
			}
		}

		if !valid {
			invalidNumber = data[i]
			log.Println(invalidNumber, "is invalid")
			break
		}
	}

	// Part 2
	done := false
	for i := 0; i < len(data); i++ {
		sum := 0
		numbers := make([]int, 0)
		for x := i; x < len(data); x++ {
			numbers = append(numbers, data[x])
			sum += data[x]
			// log.Println(sum)
			if sum == invalidNumber {
				sort.Ints(numbers)
				log.Println("Weakness", numbers[0]+numbers[len(numbers)-1])
				done = true
				break
			}
			if sum > invalidNumber {
				break
			}
		}
		if done {
			break
		}
	}
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
