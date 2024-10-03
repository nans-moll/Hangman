package main

import (
	jeu "Hangman/game"
	"Hangman/utils"
	"fmt"
	"github.com/fatih/color"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Veuillez fournir un fichier contenant des mots en paramètre.")
		return
	}

	fichierMots := os.Args[1]
	mots, err := utils.ChargerMots(fichierMots)
	if err != nil {
		fmt.Printf("Erreur lors du chargement des mots : %v\n", err)
		return
	}

	for {

		fmt.Println("\nVoulez-vous :")
		fmt.Println("					")
		color.Green("1. Facile")
		fmt.Println("					")
		color.Yellow("2. moyen")
		fmt.Println("					")
		color.Red("3. Difficile")
		fmt.Println("					")
		fmt.Println("4. Quitter")

		var choix int
		fmt.Print("Votre choix : ")
		fmt.Scanln(&choix)

		if choix == 4 {
			fmt.Println("Merci d'avoir joué ! À bientôt.")
			break
		}
		switch choix {
		case 1:
			jeu.Demarrer(mots, "facile")

		case 2:
			jeu.Demarrer(mots, "moyen")

		case 3:
			jeu.Demarrer(mots, "difficile")
		default:
			fmt.Println("veuillez réessayer")
		}
	}
}
