package main

import (
	"Hangman/Hangman/utils"
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

		jeu.jeu.Demarrer(mots)

		fmt.Println("\nVoulez-vous :")
		fmt.Println("1. Facile")
		fmt.Println("2. moyen")
		fmt.Println("3. Difficile")
		fmt.Println("4. Quitter")

		var choix int
		fmt.Print("Votre choix : ")
		fmt.Scanln(&choix)

		if choix == 4 {
			fmt.Println("Merci d'avoir joué ! À bientôt.")
			break
		} else if choix == 1 {
			fmt.Println("Choix non valide. Fermeture du jeu.")
			break
		}
		if choix == 3 {
			fmt.Println("A vous de jouer ! : ")
			break
		}
	}
}
