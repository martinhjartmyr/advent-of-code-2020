package main

import (
	"bufio"
	"log"
	"os"
	"sort"
)

const inputFilePath = "input.txt"

type ticket struct {
	row    int
	column int
	seatID int
}

func main() {
	log.Println("Loading:", inputFilePath)
	data, err := readLines(inputFilePath)
	if err != nil {
		panic(err)
	}

	log.Println("Found", len(data), "lines of input data")

	tickets := generateTickets(data)

	// sort by the highest seatID first
	sort.SliceStable(tickets, func(i, j int) bool {
		return tickets[i].seatID > tickets[j].seatID
	})

	log.Println("Highest seatID:", tickets[0].seatID)
	totalTickets := len(tickets)
	log.Println("Number of tickets:", totalTickets)

	// loop all tickets to check for the missing one
	for i := 0; i < totalTickets-1; i++ {
		if tickets[i].seatID != (tickets[i+1].seatID + 1) {
			log.Println(tickets[i].seatID-1, "is missing")
		}
	}
}

func generateTickets(data []string) []ticket {
	var tickets []ticket
	for _, d := range data {
		ticket := ticket{}
		ticket.row = getRow(d)
		ticket.column = getColumn(d)
		ticket.seatID = getSeatID(ticket.row, ticket.column)
		tickets = append(tickets, ticket)
	}
	return tickets
}

func getRow(data string) int {
	data = data[:7]
	rows := make([]int, 128)
	for i := 0; i < 128; i++ {
		rows[i] = i
	}

	for i := 0; i < 7; i++ {
		char := string(data[i])
		centerPos := len(rows) / 2
		if char == "F" {
			rows = rows[:centerPos]
		} else {
			rows = rows[centerPos:]
		}
	}

	return rows[0]
}

func getColumn(data string) int {
	data = data[7:]
	columns := make([]int, 8)
	for i := 0; i < 8; i++ {
		columns[i] = i
	}

	for i := 0; i < 3; i++ {
		char := string(data[i])
		centerPos := len(columns) / 2
		if char == "L" {
			columns = columns[:centerPos]
		} else {
			columns = columns[centerPos:]
		}
	}

	return columns[0]
}

func getSeatID(row int, column int) int {
	return row*8 + column
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
