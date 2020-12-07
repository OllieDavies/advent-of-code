package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func getParents(data map[string]map[string]int, toFind string, seen map[string]bool) map[string]bool {
	for name, bag := range data {
		for key := range bag {
			if key == toFind {
				seen[name] = true
				getParents(data, name, seen)
			}
		}
	}
	return seen
}

func main() {
	content, _ := ioutil.ReadFile("input.txt")
	rules := strings.Split(string(content), "\n")

	data := make(map[string]map[string]int)

	for _, rule := range rules {

		// contains
		containsRegex := regexp.MustCompile(`(\d) (\w+ \w+)`)
		contains := containsRegex.FindAllStringSubmatch(rule, -1)

		bags := make(map[string]int)

		for _, contain := range contains {
			num, _ := strconv.Atoi(contain[1])
			bags[contain[2]] = num
		}

		// bag
		bagRegex := regexp.MustCompile(`(\w+ \w+)`)
		bag := bagRegex.FindString(rule)

		data[bag] = bags
	}

	seen := make(map[string]bool)
	count := getParents(data, "shiny gold", seen)
	fmt.Println("Count (Part 1 ):", len(count))
}
