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

func main() {
	file, err := os.Open("2015/day02/input_test.txt")
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
		numberone, err := strconv.Atoi(v[0])
		if err != nil {
			fmt.Println(err)
		}
		numbertwo, err := strconv.Atoi(v[1])
		if err != nil {
			fmt.Println(err)
		}
		numberthree, err := strconv.Atoi(v[2])
		if err != nil {
			fmt.Println(err)
		}

		smallest := numberone
		if numbertwo < smallest {
			smallest = numbertwo
		}

		if numberthree < smallest {
			smallest = numberthree
		}

		fmt.Println(smallest, "here")
		calcul := 2*(numbertwo*numberone) + 2*(numberthree*numbertwo) + 2*(numberthree*numberone)

		final_result += calcul
		final_result += smallest
		// fmt.Println(calcul, smallest, calcul+smallest)
	}
	fmt.Println(final_result)
}
