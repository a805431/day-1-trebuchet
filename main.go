package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {

	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var fileLines []string
	result := 0

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()

	for _, line := range fileLines {
		re := regexp.MustCompile("[0-9]")
		singleLineMatches := re.FindAllString(line, -1)
		if len(singleLineMatches) > 1 {
			matchNumber, _ := strconv.Atoi(singleLineMatches[0] + singleLineMatches[len(singleLineMatches)-1])
			result += matchNumber
		} else if len(singleLineMatches) == 1 {
			matchNumber, _ := strconv.Atoi(singleLineMatches[0] + singleLineMatches[0])
			result += matchNumber
		} else {
			fmt.Println("No matches found in this line.")
		}
	}

	fmt.Println(result)
}
