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
	result2 := part2(string(file))

	fmt.Println("Part1 : ", result1)
	fmt.Println("Part2 : ", result2)
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

func part2(input string) int {
	visited := make(map[[2]int]bool)
	currentPos := [2]int{0, 0}

	santaPos := currentPos
	robotsantaPos := currentPos

	visited[currentPos] = true

	for i, char := range input {
		if i%2 == 0 {
			currentPos = santaPos
		} else {
			currentPos = robotsantaPos
		}

		switch char {
		case '^':
			currentPos[1]++
		case 'v':
			currentPos[1]--
		case '>':
			currentPos[0]++
		case '<':
			currentPos[0]--
		}

		visited[currentPos] = true

		if i%2 == 0 {
			santaPos = currentPos
		} else {
			robotsantaPos = currentPos
		}
	}

	return len(visited)
}
