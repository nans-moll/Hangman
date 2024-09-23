package main

import (
	"fmt"
	"os"
)

func main() {

	filePath := "C:/Users/nans-/GolandProjects/Hangman/Hangman/mots.txt"

	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		return

	}

	fmt.Println(string(data))
}
