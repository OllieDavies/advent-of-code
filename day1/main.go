package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	amountToFind := 2020

	content, _ := ioutil.ReadFile("input.txt")
	items := strings.Split(string(content), "\n")

	for _, firstStar := range items {
		firstStar, _ := strconv.Atoi(firstStar)

		for _, secondStar := range items {
			secondStar, _ := strconv.Atoi(secondStar)

			if firstStar+secondStar == amountToFind {
				fmt.Printf("Part 1 = %v\n", firstStar*secondStar)
			}

			for _, thirdStar := range items {
				thirdStar, _ := strconv.Atoi(thirdStar)

				if firstStar+secondStar+thirdStar == amountToFind {
					fmt.Printf("Part 2 = %v\n", firstStar*secondStar*thirdStar)
					return
				}
			}
		}
	}
}
