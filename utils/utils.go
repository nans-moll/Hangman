package utils

import (
	"bufio"
	"os"
)

func ChargerMots(cheminFichier string) ([]string, error) {
	fichier, err := os.Open(cheminFichier)
	if err != nil {
		return nil, err
	}
	defer fichier.Close()

	var mots []string
	scanner := bufio.NewScanner(fichier)

	for scanner.Scan() {
		mots = append(mots, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return mots, nil
}
