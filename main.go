package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var fileLines []string
	var linesWithNoMatches []string
	result := 0
	searchValues := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
	re := regexp.MustCompile(buildRegexPattern(searchValues))

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()

	for _, line := range fileLines {
		var singleLineMatches []string
		singleLineMatches = re.FindAllString(line, -1)

		if len(singleLineMatches) > 0 {
			fmt.Println(singleLineMatches)
			result += getFinalNumberFromLine(searchValues, singleLineMatches)
		} else {
			linesWithNoMatches = append(linesWithNoMatches, line)
		}
	}

	fmt.Println(result)
	//fmt.Println(linesWithNoMatches)
}

func buildRegexPattern(valueMap map[string]int) string {
	pattern := "\\d|"
	counter := 0

	for i := range valueMap {
		pattern += i
		counter++
		if counter < len(valueMap) {
			pattern += "|"
		}
	}
	return pattern
}

// returns the required number for a single line
func getFinalNumberFromLine(valueMap map[string]int, lineMatches []string) int {
	numberOfMatches := len(lineMatches)
	var numberString string

	if numberOfMatches > 1 {
		numberString += getMatchStringValue(valueMap, lineMatches[0]) + getMatchStringValue(valueMap, lineMatches[numberOfMatches-1])
	} else if numberOfMatches == 1 {
		numberString += strings.Repeat(getMatchStringValue(valueMap, lineMatches[0]), 2)
	} else {
		return -1
	}

	numberValue, _ := strconv.Atoi(numberString)
	return numberValue
}

// get the string value for a single match from one line
func getMatchStringValue(valueMap map[string]int, match string) string {
	var singleElemStringValue string = ""

	if len(match) > 1 {
		singleElemStringValue += strconv.Itoa(valueMap[match])
	} else if len(match) == 1 {
		singleElemStringValue += match
	} else {
		return singleElemStringValue
	}

	return singleElemStringValue
}
