package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"regexp"
	"sort"
	"strings"
)

func search(input string, nextPosition int, min int, max int, lowerLetter rune, upperLetter rune) int {
	// Reached end of input. Return either min or max
	if nextPosition+1 >= len(input) {
		if rune(input[nextPosition]) == lowerLetter {
			return min
		}
		return max
	}

	switch rune(input[nextPosition]) {
	case lowerLetter:
		middle := (max + min) / 2
		return search(input, nextPosition+1, min, middle, lowerLetter, upperLetter)
	case upperLetter:
		middle := math.Ceil(float64(max+min) / 2)
		return search(input, nextPosition+1, int(middle), max, lowerLetter, upperLetter)
	}

	return 0
}

func getRowNumber(input string, nextPosition int, min int, max int) int {
	return search(input, nextPosition, min, max, 'F', 'B')
}

func getColumnNumber(input string, nextPosition int, min int, max int) int {
	return search(input, nextPosition, min, max, 'L', 'R')
}

func main() {
	content, _ := ioutil.ReadFile("input.txt")
	passes := strings.Split(string(content), "\n")

	var takenSeats []int
	for _, boardingPass := range passes {
		r := regexp.MustCompile(`([FB]+)([RL]+)`)
		matches := r.FindAllStringSubmatch(boardingPass, 2)[0]

		row, column := getRowNumber(matches[1], 0, 0, 127), getColumnNumber(matches[2], 0, 0, 7)

		seat := row*8 + column
		takenSeats = append(takenSeats, seat)
	}

	sort.Ints(takenSeats)

	lowestSeatNumber, highestSeatNumber := takenSeats[0], takenSeats[len(takenSeats)-1]

	fmt.Println("Highest Seat Number (Part 1):", highestSeatNumber)
	fmt.Println("Lowest Seat Number:", lowestSeatNumber)

	// Find my seat
	for index, seat := range takenSeats {
		if takenSeats[index+1] != seat+1 {
			fmt.Println("My Seat ID (Part 2):", seat+1)
			break
		}
	}
}
