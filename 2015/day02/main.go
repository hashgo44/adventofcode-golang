package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func StringToInt(numberInString string) (int, error) {
	result, err := strconv.Atoi(numberInString)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func main() {
	file, err := os.Open("2015/day02/input.txt")
	check(err)
	defer file.Close()

	var all_text_split [][]string
	var final_result int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()

		text_split := strings.Split(text, "x")

		all_text_split = append(all_text_split, text_split)
	}

	for _, v := range all_text_split {
		numberone, err := StringToInt(v[0])
		check(err)
		numbertwo, err := StringToInt(v[1])
		check(err)
		numberthree, err := StringToInt(v[2])
		check(err)

		x := numbertwo * numberone
		y := numberthree * numbertwo
		z := numberthree * numberone

		smallest := x
		if y < smallest {
			smallest = y
		}
		if z < smallest {
			smallest = z
		}

		calcul := 2*x + 2*y + 2*z

		final_result += calcul
		final_result += smallest
	}
	fmt.Println(final_result)
}
