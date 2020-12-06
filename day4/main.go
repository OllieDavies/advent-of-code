package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type passport struct {
	fields map[string]string
}

func isStringOutOfRange(stringYear string, min int, max int) bool {
	year, _ := strconv.Atoi(stringYear)
	if year < min || year > max {
		return true
	}
	return false
}

func getRegexResult(s string, regex string) string {
	exp := regexp.MustCompile(regex)
	return exp.FindString(s)
}

func (p *passport) isValid() bool {
	fields := p.fields

	// Check fields are present
	requiredFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	for _, requiredField := range requiredFields {
		_, ok := fields[requiredField]
		if ok == false {
			return false
		}
	}

	// Validate Birth Year
	if isStringOutOfRange(fields["byr"], 1920, 2002) {
		return false
	}

	// Validate Issue Year
	if isStringOutOfRange(fields["iyr"], 2010, 2020) {
		return false
	}

	// Validate Expiration Year
	if isStringOutOfRange(fields["eyr"], 2020, 2030) {
		return false
	}

	// Validate Height
	heightUnit := getRegexResult(fields["hgt"], `\D+`)
	height := getRegexResult(fields["hgt"], `\d+`)

	if heightUnit == "cm" && isStringOutOfRange(height, 150, 193) {
		return false
	} else if heightUnit == "in" && isStringOutOfRange(height, 59, 76) {
		return false
	} else if heightUnit == "" {
		return false
	}

	// Validate Hair Color
	hairColor := getRegexResult(fields["hcl"], `(#[0-9a-f]{6})\b`)
	if hairColor == "" {
		return false
	}

	// Validate Eye Color
	eyeColor := getRegexResult(fields["ecl"], `(amb|blu|brn|gry|grn|hzl|oth)\b`)
	if eyeColor == "" {
		return false
	}

	// Validate Passport Number
	passportNumber := getRegexResult(fields["pid"], `^(\d{9})$`)
	if passportNumber == "" {
		return false
	}

	return true
}

func main() {
	content, _ := ioutil.ReadFile("input.txt")
	data := strings.Split(string(content), "\n\n") // passports end with two newlines

	validCount := 0
	for _, passportData := range data {
		r := regexp.MustCompile(`(\w+):(\#?\w+)`)
		matches := r.FindAllStringSubmatch(passportData, 8)

		p := passport{fields: make(map[string]string)}
		for _, match := range matches {
			p.fields[match[1]] = match[2]
		}

		if p.isValid() {
			validCount++
		}
	}
	fmt.Println("Valid Passports:", validCount)
}
