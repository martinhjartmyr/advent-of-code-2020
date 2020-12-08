package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

const inputFilePath = "input.txt"

var bagColors = make(map[string]bool)

func main() {
	log.Println("Loading:", inputFilePath)
	data, err := readLines(inputFilePath)
	if err != nil {
		panic(err)
	}

	log.Println("Found", len(data), "lines of input data")
	bags := parseBags(data)
	getBagColors("shiny gold", bags)
	log.Println("Unique colors", len(bagColors))
	bagCount := getBagCount("shiny gold", bags)
	log.Println("Bag count", bagCount)
}

func getBagColors(bagName string, bags map[string]map[string]int) {
	for key, bagRules := range bags {
		if _, ok := bagRules[bagName]; ok {
			bagColors[key] = true
			getBagColors(key, bags)
		}
	}
}

func getBagCount(bagName string, bags map[string]map[string]int) int {
	bagCounter := 0
	for key, count := range bags[bagName] {
		bagCounter += count
		bagCounter += count * getBagCount(key, bags)
	}
	return bagCounter
}

func parseBags(data []string) map[string]map[string]int {
	var bags = map[string]map[string]int{}

	for _, line := range data {
		bagName := regexp.MustCompile(`^([a-z]+\ [a-z]+){1}\ bags\ contain`).FindStringSubmatch(line)
		bagRules := regexp.MustCompile(`(\d+\ [a-z]+\ [a-z]+) bag[s]?`).FindAllStringSubmatch(line, -1)

		if len(bagName) == 2 {
			bags[bagName[1]] = map[string]int{}
			if len(bagRules) > 0 {
				for _, bagRule := range bagRules {
					bagRuleArray := regexp.MustCompile(`^(\d+)+\ ([a-z]+\ [a-z]+)`).FindStringSubmatch(bagRule[1])
					if len(bagRuleArray) == 3 {
						count, _ := strconv.Atoi(bagRuleArray[1])
						bags[bagName[1]][bagRuleArray[2]] = count
					}
				}
			}
		} else {
			continue
		}

	}
	return bags
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
