package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

func getFullLine(linePart string, minX int) string {
	minLength := float64(minX/len(linePart) + 1)
	duplications := math.Max(math.Ceil(minLength), 1)
	return strings.Repeat(linePart, int(duplications))
}

func incrementValueIfTree(value *int, line string, x int) {
	fullLine := getFullLine(line, x)
	if string(fullLine[x]) == "#" {
		*value++
	}
}

func main() {
	content, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(content), "\n")

	a, b, c, d, e := 0, 0, 0, 0, 0
	for index, line := range lines {

		// Slope A
		incrementValueIfTree(&a, line, index)

		// Slope B (Part 1)
		incrementValueIfTree(&b, line, 3*index)

		// Slope C
		incrementValueIfTree(&c, line, 5*index)

		// Slope D
		incrementValueIfTree(&d, line, 7*index)

		// Slope E
		if index%2 == 0 {
			incrementValueIfTree(&e, line, index/2)
		}

	}
	fmt.Println("Slope A Trees Encountered:", a)
	fmt.Println("Slope B Trees Encountered:", b, "(Part 1)")
	fmt.Println("Slope C Trees Encountered:", c)
	fmt.Println("Slope D Trees Encountered:", d)
	fmt.Println("Slope E Trees Encountered:", e)
	fmt.Println("Part 2 Result:", a*b*c*d*e)
}
