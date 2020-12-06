package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func isValidSledRentalPassword(letter string, minOccurances int, maxOccurances int, password string) bool {
	charRegex := regexp.MustCompile(letter)
	matches := charRegex.FindAllStringIndex(password, -1)

	occurances := len(matches)

	if occurances > maxOccurances || occurances < minOccurances {
		return false
	}

	return true
}

func isValidPassword(letter string, firstPosition int, secondPosition int, password string) bool {
	doesFirstPositionMatch := string(password[firstPosition]) == letter
	doesSecondPositionMatch := string(password[secondPosition]) == letter

	// XOR
	if doesFirstPositionMatch != doesSecondPositionMatch {
		return true
	}
	return false
}

func main() {
	content, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(content), "\n")

	part1Count := 0
	for _, line := range lines {
		r := regexp.MustCompile(`(\d+)-(\d+) (\w){1}: (\w+)`)
		matches := r.FindStringSubmatch(line)

		minOccurancesStr, maxOccurancesStr, letter, password := matches[1], matches[2], matches[3], matches[4]

		minOccurances, _ := strconv.Atoi(minOccurancesStr)
		maxOccurances, _ := strconv.Atoi(maxOccurancesStr)

		valid := isValidSledRentalPassword(letter, minOccurances, maxOccurances, password)

		if valid {
			part1Count++
		}
	}
	fmt.Println("Part 1 - Valid Passwords:", part1Count)

	part2Count := 0
	for _, line := range lines {
		r := regexp.MustCompile(`(\d+)-(\d+) (\w): (\w+)`)
		matches := r.FindStringSubmatch(line)

		firstPositionStr, secondPositionStr, letter, password := matches[1], matches[2], matches[3], matches[4]

		firstPosition, _ := strconv.Atoi(firstPositionStr)
		secondPosition, _ := strconv.Atoi(secondPositionStr)

		valid := isValidPassword(letter, firstPosition-1, secondPosition-1, password)

		if valid {
			part2Count++
		}
	}
	fmt.Println("Part 2 - Valid Passwords:", part2Count)
}
