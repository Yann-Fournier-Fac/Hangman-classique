package main

// régler bug ajout des lettres
// faire une fonction clear()

import (
	"fmt"
	"hangman"
	"os"
	"os/exec"
	//"reflect"
)

func main() {
	for {
		str := ""
		fmt.Println("1 : Lancer une nouvelle partie")
		fmt.Println("2 : Continuer la partie sauvegarder")
		fmt.Println("3 : En savoir à propos du jeu")
		fmt.Println("4 : Quitter le jeu")
		fmt.Printf("Ecris le numéro de ce que tu veux faire : ")
		fmt.Scan(&str)
		fmt.Printf("\n")

		// Clear le terminal
		c := exec.Command("clear")
		c.Stdout = os.Stdout
		c.Run()

		switch str {
		case "1":

			Version := true

			for Version {
				choix := ""
				fmt.Println("A quelle version du jeu voulez-vous jouer : ")
				fmt.Println("1 : Version de base")
				fmt.Println("2 : Version Ascii")
				fmt.Println("3 : Sortir du jeu")
				fmt.Scan(&choix)

				// Clear le terminal
				c := exec.Command("clear")
				c.Stdout = os.Stdout
				c.Run()

				switch choix {
				case "1":
					hangman.JeuBase()
				case "2":
					hangman.JeuAscii()
				case "3":
					Version = false
				default:
					fmt.Println("Merci de saisir un chiffre correct")
				}
			}

		case "2":
			fmt.Println("Continuer la partie sauvegardée")
			fmt.Printf("\n")
		case "3":
			fmt.Println("readme")
			fmt.Printf("\n")
		case "4":
			return
		}
	}
}
