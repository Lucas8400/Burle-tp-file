package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	var fragment1 string
	var fragment2 string
	var fragment3 string
	var fragment4 string
	readFile, _ := os.Open("File.txt")        // Ouverture du fichier
	fileScanner := bufio.NewScanner(readFile) // Création d'un scanner pour lire le fichier
	fileScanner.Split(bufio.ScanLines)        // On scanne ligne par ligne
	var lines []string                        //Création d'un tableau conteant chaque ligne du fichier(mots)
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text()) //Ajout de chaque ligne du fichier dans le tableau
	}
	fragment1 += lines[0]
	fragment2 += lines[len(lines)-1]
	for index, word := range lines { // On parcourt le tableau et on utilise des conditions pour effectuer ce que l'on nous demande dans l'énigme pour les fragment 3 et 4
		if word == "before" {
			strIt := lines[index+1]
			intIt, _ := strconv.Atoi(strIt) // On convertit le string en int
			intIt -= 1                      // On soustrait 1 à l'entier
			fragment3 += lines[intIt]       // On récupére le mot à l'index intIt
		}
		if word == "now " {
			result := lines[index-1]               // On récupère le mot avant "now"
			result1 := int(result[1]) / len(lines) // On récupère le 2ème caractère du mot et on le divise par le nombre de mots dans le fichier
			fragment4 += lines[result1-1]          // On récupère le mot correspondant à la division
		}

	}
	fmt.Println(fragment1+"", fragment2+"", fragment3+"", fragment4) // Affichage des résultats séparés par des espaces
	rand.Seed(time.Now().UnixNano())                                 // On initialise le générateur de nombres aléatoires
	fmt.Println(rand.Intn(100))                                      // On affiche un nombre aléatoire entre 0 et 100
}
