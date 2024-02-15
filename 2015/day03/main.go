package main

import (
	"fmt"
	"os"
)

func main() {
	// Utilisation de os.ReadFile pour lire le contenue du fichier
	file, err := os.ReadFile("2015/day03/input.txt")
	// Gestion des erreurs
	if err != nil {
		panic(err)
	}
	result1 := part1(string(file))
	fmt.Println("Part1 : ", result1)
	part2(string(file))
}

func part1(input string) int {
	visited := make(map[string]int)
	y, x := 0, 0
	for _, v := range input {
		switch string(v) {
		case "^":
			y++
		case "v":
			y--
		case ">":
			x++
		case "<":
			x--
		}
		coords := fmt.Sprintf("%d,%d", x, y)
		visited[coords]++
	}

	count := 1
	for _, v := range visited {
		if v >= 1 {
			count++
		}
	}
	return count
}

func part2(input string) {
	var slices1 string
	var slices2 string
	houseCount := make(map[string]int)

	for i, v := range input {
		if i%2 == 0 {
			slices1 += string(v)
		} else {
			slices2 += string(v)
		}
	}

	// COORDENNER PERE NOEL
	// visited1 := make(map[string]int)
	y, x := 0, 0
	for _, v := range slices1 {
		switch string(v) {
		case "^":
			y++
		case "v":
			y--
		case ">":
			x++
		case "<":
			x--
		}
		coords1 := fmt.Sprintf("%d,%d", x, y)
		houseCount[coords1]++
	}

	// COORDENNER ROBOT PERE NOEL
	// visited2 := make(map[string]int)
	z, w := 0, 0
	for _, v := range slices2 {
		switch string(v) {
		case "^":
			z++
		case "v":
			z--
		case ">":
			w++
		case "<":
			w--
		}
		coords2 := fmt.Sprintf("%d,%d", z, w)
		houseCount[coords2]++
	}
	fmt.Println(len(houseCount))
}
