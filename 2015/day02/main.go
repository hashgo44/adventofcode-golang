package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Utilisation de os.ReadFile
	file, err := os.ReadFile("2015/day02/input_test.txt")
	// Gestion des erreurs
	if err != nil {
		panic(err)
	}
	// Appel de la method part1 et conversion du file en string. return le résultat
	result := part1(string(file))
	// Affichage du résultat de la part1
	fmt.Println(result)
}

// Comparer deux entiers pour trouver le plus petit
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func part1(input string) int {
	// Déclaration de la variable totalSqFt pour stocker le nombre total de pieds carrés nécessaires
	var totalSqFt int

	// Parcourir chaque ligne de la chaîne d'entrée (chaque ligne représente les dimensions d'une boîte)
	for _, line := range strings.Split(input, "\n") {
		// Déclaration des variables pour stocker les dimensions de la boîte
		var x, y, z int
		// Lecture des dimensions de la boîte à partir de la ligne, en utilisant fmt.Sscanf qui analyse la chaîne selon un format spécifié
		_, err := fmt.Sscanf(line, "%dx%dx%d", &x, &y, &z)
		// Gestion des erreurs
		if err != nil {
			panic(err)
		}
		// Calcul de la surface de la boîte en ajoutant la surface de chaque face (longueur x largeur) multipliée par 2 (car il y a deux côtés identiques par face)
		totalSqFt += x * y * 2
		totalSqFt += x * z * 2
		totalSqFt += z * y * 2
		// Calcul de la surface minimale nécessaire pour emballer la boîte. Cela correspond à la plus petite surface parmi les six surfaces possibles.
		totalSqFt += min(min(x*y, y*z), x*z)
	}
	// Retourner le nombre total de pieds carrés
	return totalSqFt
}
