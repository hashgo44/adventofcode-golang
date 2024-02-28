package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	input, err := os.ReadFile("2015/day05/input.txt")
	if err != nil {
		panic(err)
	}
	result1 := part1(string(input))
	result2 := part2(string(input))
	fmt.Println(result1)
	fmt.Println(result2)
}

func part1(input string) int {
	var nice int

	disallowPattern := regexp.MustCompile("(ab|cd|pq|xy)")
	for _, line := range strings.Split(input, "\n") {

		var hasDouble bool
		for i := 0; i < len(line)-1; i++ {
			if line[i] == line[i+1] {
				hasDouble = true
				break
			}
		}

		if countVowels(line) >= 3 && !disallowPattern.MatchString(line) && hasDouble {
			nice++
		}

	}
	return nice
}

func part2(input string) int {
	var nice int
	for _, line := range strings.Split(input, "\n") {
		oneLetterWhichRepeats := false

		for i := 0; i < len(line)-2; i++ {
			if line[i] == line[i+2] {
				oneLetterWhichRepeats = true
			}
		}

		pairInLine := false
		for i := 0; i < len(line)-1; i++ {
			pair := line[i : i+2]
			if strings.Count(line, pair) >= 2 {
				pairInLine = true
				break
			}
		}

		if oneLetterWhichRepeats && pairInLine {
			nice++
		}
	}
	return nice
}

func countVowels(str string) int {
	count := 0
	vowels := "aeiou"
	for _, v := range str {
		if strings.ContainsRune(vowels, v) {
			count++
		}
	}
	return count
}
