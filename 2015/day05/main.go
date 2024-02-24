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
	result := part1(string(input))
	fmt.Println(result)
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
