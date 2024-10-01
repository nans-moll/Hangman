package jeu

import (
	"fmt"
	"math/rand"
	"strings"
)

type Jeu struct {
	MotSecret           string
	MotAffiche          []rune
	TentativesRestantes int
	LettresDevinees     []rune
}

var etapesPendu = []string{
	`
     -----
     |   |
         |
         |
         |
         |
    =======`,
	`
     -----
     |   |
     O   |
         |
         |
         |
    =======`,
	`
     -----
     |   |
     O   |
     |   |
         |
         |
    =======`,
	`
     -----
     |   |
     O   |
    /|   |
         |
         |
    =======`,
	`
     -----
     |   |
     O   |
    /|\\  |
         |
         |
    =======`,
	`
     -----
     |   |
     O   |
    /|\\  |
    /    |
         |
    =======`,
	`
     -----
     |   |
     O   |
    /|\\  |
    / \\  |
         |
    =======`,
}

func Demarrer(mots []string) {
	motSecret := mots[rand.Intn(len(mots))]

	jeu := &Jeu{
		MotSecret:           strings.ToLower(motSecret),
		MotAffiche:          make([]rune, len(motSecret)),
		TentativesRestantes: 5,
		LettresDevinees:     []rune{},
	}

	for i := range jeu.MotAffiche {
		jeu.MotAffiche[i] = '_'
	}

	for !jeu.estTermine() {
		jeu.afficherEtat()
		fmt.Print("Entrez une lettre : ")
		var input string
		fmt.Scanln(&input)
		if len(input) > 0 {
			jeu.devinerLettre(rune(input[0]))
		}

	}

	if jeu.estGagne() {
		fmt.Printf("Félicitations, vous avez gagné ! Le mot était : %s\n", jeu.MotSecret)
	} else {
		fmt.Printf("Dommage, vous avez perdu. Le mot était : %s\n", jeu.MotSecret)
	}
}

func (j *Jeu) devinerLettre(lettre rune) {
	lettre = rune(strings.ToLower(string(lettre))[0])

	// Vérifier si la lettre est bien une lettre alphabétique
	if lettre < 'a' || lettre > 'z' {
		fmt.Println("chef choisi une lettre !")
		return
	}

	if strings.ContainsRune(string(j.LettresDevinees), lettre) {
		fmt.Println("Vous avez déjà deviné cette lettre.")
		return
	}

	j.LettresDevinees = append(j.LettresDevinees, lettre)

	if strings.ContainsRune(j.MotSecret, lettre) {
		for i, l := range j.MotSecret {
			if l == lettre {
				j.MotAffiche[i] = lettre
			}
		}
	} else {
		j.TentativesRestantes--
	}
}

func (j *Jeu) afficherEtat() {
	fmt.Println(etapesPendu[5-j.TentativesRestantes])
	fmt.Printf("\nMot à deviner : %s\n", string(j.MotAffiche))
	fmt.Printf("Tentatives restantes : %d\n", j.TentativesRestantes)
	fmt.Printf("Lettres déjà tentées : %s\n\n", string(j.LettresDevinees))
}

func (j *Jeu) estGagne() bool {
	return !strings.ContainsRune(string(j.MotAffiche), '_')
}

func (j *Jeu) estTermine() bool {
	return j.TentativesRestantes <= 0 || j.estGagne()
}
