package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

func main() {

	content, _ := ioutil.ReadFile("input.txt")
	groups := strings.Split(string(content), "\n\n") // groups end with two newlines

	count := 0
	for _, groupAnswers := range groups {
		peopleInGroup := len(strings.Split(groupAnswers, "\n"))

		r := regexp.MustCompile(`([a-z])`)
		chars := r.FindAllString(groupAnswers, len(groupAnswers))

		uniqueAnswers := make(map[string]int)
		for _, char := range chars {
			uniqueAnswers[string(char)]++
		}

		for _, answers := range uniqueAnswers {
			if answers == peopleInGroup {
				count++
			}
		}
	}
	fmt.Println("Total (Part 2):", count)

}
