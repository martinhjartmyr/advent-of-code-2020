package main

import (
	"bufio"
	"log"
	"os"
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

	log.Println("Part 1:")
	missionComplete := false
	for _, value := range data {
		for _, valueComp := range data {
			valueInt, _ := strconv.Atoi(value)
			valueCompInt, _ := strconv.Atoi(valueComp)
			if valueInt+valueCompInt == 2020 {
				log.Println(valueInt, "+", valueCompInt, "=", "2020")
				log.Println(valueInt, "*", valueCompInt, "=", valueInt*valueCompInt)
				missionComplete = true
				break
			}
		}
		if missionComplete {
			break
		}
	}

	log.Println("Part 2:")
	missionComplete = false
	for _, value := range data {
		for _, valueComp := range data {
			for _, valueComp2 := range data {
				valueInt, _ := strconv.Atoi(value)
				valueCompInt, _ := strconv.Atoi(valueComp)
				valueComp2Int, _ := strconv.Atoi(valueComp2)
				if valueInt+valueCompInt+valueComp2Int == 2020 {
					log.Println(valueInt, "+", valueCompInt, "+", valueComp2Int, "=", "2020")
					log.Println(valueInt, "*", valueCompInt, "*", valueComp2Int, "=", valueInt*valueCompInt*valueComp2Int)
					missionComplete = true
					break
				}
			}
			if missionComplete {
				break
			}
		}
		if missionComplete {
			break
		}
	}
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
