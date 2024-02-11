package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Utilisation de os.ReadFile
	file, err := os.ReadFile("2015/day02/input.txt")
	// Gestion des erreurs
	if err != nil {
		panic(err)
	}
	// Appel des functions part1 et part2 et conversion du file en string. return le résultat
	result1 := part1(string(file))
	result2 := part2(string(file))

	// Affichage des résultats
	fmt.Println("Result 1 :", result1)
	fmt.Println("Result 2 :", result2)
}

func part1(input string) int {
	// Déclaration de la variable totalSqFt pour stocker le nombre total de pieds carrés nécessaires
	var totalSqFt int

	// Parcourir chaque ligne de la chaîne d'entrée (chaque ligne représente les dimensions d'une boîte)
	for _, line := range strings.Split(input, "\n") {
		// Déclaration des variables pour stocker les dimensions de la boîte
		var x, y, z int
		// Lecture des dimensions de la boîte à partir de la ligne, en utilisant fmt.Sscanf qui analyse la chaîne selon un format spécifié
		// Le return des valeurs x, y, z doivent etre des pointeurs pour modifier les vrais variables et non une copie en mémoire
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

func part2(input string) int {
	// Déclaration de la variable TotalLen pour stocker la taille total du ruban
	var TotalLen int

	// Parcourir chaque ligne des dimensions de la boîte
	for _, line := range strings.Split(input, "\n") {
		// Déclaration des variables pour stocker les dimensions de la boîte
		var x, y, z int
		// Utilisation de Sscanf pour enregister les valeurs dans les variables x, y, z
		_, err := fmt.Sscanf(line, "%dx%dx%d", &x, &y, &z)
		// Gestion des erreurs
		if err != nil {
			panic(err)
		}
		// Calcule du volumes des cubes
		cubic := x * y * z
		// Ajout du volume des cubes dans TotalLen
		TotalLen += cubic

		// Calcule du nœud pour le ruban
		// Additions des deux faces et multiplications par 2
		a := 2 * (x + y)
		b := 2 * (z + y)
		c := 2 * (x + z)

		// Ajouter la valeurs la plus petite a la taille Total
		TotalLen += min(min(a, b), c)
	}
	return TotalLen
}

// Comparer deux entiers pour trouver le plus petit
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
