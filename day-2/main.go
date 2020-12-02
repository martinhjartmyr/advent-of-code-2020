package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const inputFilePath = "input.txt"

func main() {
	log.Println("Loading:", inputFilePath)
	data, err := readLines(inputFilePath)
	if err != nil {
		panic(err)
	}

	log.Println("Found", len(data), "lines of input data")

	r := regexp.MustCompile(`^(\d+)-(\d+) ([a-z]{1}): ([a-z]+)`)
	matchingPasswordPolicy1Count := 0
	matchingPasswordPolicy2Count := 0

	for _, rowData := range data {
		sample := r.FindStringSubmatch(rowData)

		countMin, _ := strconv.Atoi(sample[1])
		countMax, _ := strconv.Atoi(sample[2])
		char := string(sample[3])
		password := string(sample[4])

		matchCount := strings.Count(password, char)
		if matchCount >= countMin && matchCount <= countMax {
			matchingPasswordPolicy1Count++
		}

		// re-use countMin countMax for positions as in policy 2
		charPosition1 := string(password[countMin-1])
		charPosition2 := string(password[countMax-1])
		if charPosition1 != charPosition2 && (charPosition1 == char || charPosition2 == char) {
			matchingPasswordPolicy2Count++
		}
	}

	log.Println("Matching passwords policy 1:", matchingPasswordPolicy1Count)
	log.Println("Matching passwords policy 2:", matchingPasswordPolicy2Count)
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
