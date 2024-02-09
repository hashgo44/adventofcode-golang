package main

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	part1()
	part2()
}

func part1() {
	dat, err := os.ReadFile("2015/day01/input.txt")
	check(err)
	number_of_floor := 0
	for _, v := range dat {
		if string(v) == "(" {
			number_of_floor += 1
		} else {
			number_of_floor -= 1
		}
	}
	fmt.Println("The instructions take Santa is :", number_of_floor)
}

func part2() {
	dat, err := os.ReadFile("2015/day01/input.txt")
	check(err)
	number_of_floor := 0
	for i, v := range dat {
		if string(v) == "(" {
			number_of_floor += 1
		} else {
			if number_of_floor == -1 {
				fmt.Println("The position of the character that causes Santa to first enter the basement is :", i)
				return
			}
			number_of_floor -= 1
		}
	}
}
