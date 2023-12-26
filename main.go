package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

		if len(singleLineMatches) > 1 {
			fmt.Println(singleLineMatches)
		} else if len(singleLineMatches) == 1 {
			fmt.Println(singleLineMatches)
		} else {
			linesWithNoMatches = append(linesWithNoMatches, line)
		}
	}

	fmt.Println(result)
	//fmt.Println(linesWithNoMatches)
}

func buildRegexPattern(valueMap map[string]int) string {
	var pattern string
	counter := 0

	for k := range valueMap {
		pattern += k
		counter++
		if counter < len(valueMap) {
			pattern += "|"
		}
	}
	return pattern
}

//returns the required number for a single line as a string aka "22"
func matchWordsWithDigits(valueMap map[string]int, words []string) string {
	numberOfWords := len(words)
	if numberOfWords > 1 {
		return valueMap[words[0]] valueMap[words[numberOfWords-1]]
	} else if numberOfWords == 1 {

	} else {
		return "-1"
	}
}
