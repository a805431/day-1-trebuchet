package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
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
	searchValues := map[int]string{
		1: "one",
		2: "two",
		3: "three",
		4: "four",
		5: "five",
		6: "six",
		7: "seven",
		8: "eight",
		9: "nine",
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

func buildRegexPattern(valueMap map[int]string) string {
	var pattern string
	var keys []int
	for k := range valueMap {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	for _, k := range keys {
		pattern += valueMap[k]
		if k != keys[len(keys)-1] {
			pattern += "|"
		}
	}
	return pattern
}

// func matchWordsWithDigits(valueMap map[int]string, words[]string) string {
// 	numberOfWords := len(words)
// 	if numberOfWords > 1 {
// 		words[0] 
// 		words[numberOfWords - 1]
// 	} else if numberOfWords == 1 {

// 	} else {
// 		return "-1"
// 	}
// }