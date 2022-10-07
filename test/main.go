package main

import (
	"fmt"
	"hangman"
	"os"
)

func main() {

	Motdev := []string{}

	if len(os.Args[1:]) == 0 {
		fmt.Printf("File name missing\n")
	} else if len(os.Args[1:]) > 1 {
		fmt.Printf("Too many arguments\n")
	} else {
		contents, err := os.ReadFile(os.Args[1])
		if err != nil {
			fmt.Println("File reading error", err)
			return
		}

		Mot := hangman.Findword(contents)
		fmt.Printf("Le mot à trouver est : %v", Mot)
		fmt.Printf("\n")

		for i := 0; i < len(Mot); i++ {
			//Ascci = append(Ascci, []string{})
			Motdev = append(Motdev, "_")
		}

		lettremanque := 0
		// affichage des n lettre Ascii
		n := (len(Mot) / 2) - 1
		fmt.Print((len(Mot) / 2) - 1)

		fmt.Printf("Les n lettres à afficher : %v", n)
		fmt.Printf("\n")

		lettremanque = len(Mot) - n
		fmt.Printf("Lettres manquante : %v", lettremanque)
		fmt.Printf("\n")

		Motdev = hangman.NLetter(Mot, Motdev)
		fmt.Print(Motdev)
		fmt.Printf("\n")
	}
}
