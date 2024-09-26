package main

import (
	"Hangman/game"
	"Hangman/utils"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("ajouter un fichier existant avec")
		return
	}
	wordsFile := os.Args[1]
	words, err := utils.LoadWords(wordsFile)
	if err != nil {
		fmt.Printf("Erreur lors du chargement des mots: %v\n", err)
		return
	}
	{
		game.start(words)

		fmt.Println("\nVous voulez ?:")
		fmt.Println("							")
		fmt.Println("1. New game")
		fmt.Println("							")
		fmt.Println("2. Exit")
	}
	var choice int
	fmt.Print("Votre choix : ")
	fmt.Scanln(&choice)

	if choice == 2 {
		fmt.Println("Merci d'avoir joué ! À bientôt.")
		break
	} else if choice != 1 {
		fmt.Println("Choix non autorise. aurevoir.")
		break
	}

}
