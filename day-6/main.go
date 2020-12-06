package main

import (
	"io/ioutil"
	"log"
	// "strconv"
	"strings"
)

const inputFilePath = "input.txt"

func main() {
	log.Println("Loading:", inputFilePath)
	byteArray, err := ioutil.ReadFile(inputFilePath)
	if err != nil {
		panic(err)
	}
	data := string(byteArray)

	answerGroups := strings.Split(data, "\n\n")
	log.Println("Group count:", len(answerGroups))

	// Part 1

	var uniqueAnswers []int
	for _, groupData := range answerGroups {
		groupData = strings.ReplaceAll(groupData, "\n", "")

		var charMap map[string]int
		charMap = make(map[string]int)

		for i := 0; i < len(groupData); i++ {
			char := string(groupData[i])
			if _, ok := charMap[char]; ok {
				charMap[char]++
			} else {
				charMap[char] = 1
			}
		}
		uniqueAnswers = append(uniqueAnswers, len(charMap))
	}

	sumOfUniqueAnswers := 0
	for _, d := range uniqueAnswers {
		sumOfUniqueAnswers += d
	}

	log.Println("The sum of all unique answers is:", sumOfUniqueAnswers)

	// Part 2

	var sameAnswers []int
	for _, groupData := range answerGroups {
		userGroupData := strings.Split(groupData, "\n")
		userCount := len(userGroupData)

		var charMap map[string]int
		charMap = make(map[string]int)

		for _, userData := range userGroupData {
			for i := 0; i < len(userData); i++ {
				char := string(userData[i])
				if _, ok := charMap[char]; ok {
					charMap[char]++
				} else {
					charMap[char] = 1
				}
			}
		}

		groupSameAnswers := 0
		for _, count := range charMap {
			if count == userCount {
				groupSameAnswers++
			}
		}

		sameAnswers = append(sameAnswers, groupSameAnswers)
	}

	sumOfSameAnswers := 0
	for _, d := range sameAnswers {
		sumOfSameAnswers += d
	}

	log.Println("The sum of all same answers is:", sumOfSameAnswers)
}
