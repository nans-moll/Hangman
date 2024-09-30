package main

import (
	"Hangman/game"
	"Hangman/utils"
	"fmt"
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

		jeu.Demarrer(mots)

		fmt.Println("\nVoulez-vous :")
		fmt.Println("1. Nouvelle Partie")
		fmt.Println("2. Quitter")

		var choix int
		fmt.Print("Votre choix : ")
		fmt.Scanln(&choix)

		if choix == 2 {
			fmt.Println("Merci d'avoir joué ! À bientôt.")
			break
		} else if choix != 1 {
			fmt.Println("Choix non valide. Fermeture du jeu.")
			break
		}
	}
}
