package main

// faire la sauvegarde

import (
	"bufio"
	"fmt"
	"hangman"
	"io"
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

			// Initialisation des variables
			cpt := -1
			var lettremanque int
			MotATrouver := []string{}
			Ascci := [][]string{}
			Mot := ""
			var Pendu []string
			Lettre := []string{} // pour les lettres deja dites
			Version := true
			var Stop bool

			// on teste que le mot soit bon
			if Mot == "Veuillez relancer le jeux" {
				fmt.Println("Veuillez relancer le jeux")
				return
			}

			Pendu = hangman.Hangmanpose()

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

					cpt, lettremanque, MotATrouver, Mot, Pendu, Lettre, Stop = hangman.JeuBase(cpt, lettremanque, MotATrouver, Mot, Pendu, Lettre)

					// Sauvegarde
					if Stop {
						boolean := true
						for boolean {
							fmt.Println("Voulez vous sauvegarder votre partie : oui / non ")
							choice := ""
							fmt.Scan(&choice)
							switch choice {
							case "oui":
								fmt.Println("Votre Partie a été sauvegarder")
								// sauvegarde du pendu basic
								boolean = false
							case "non":
								fmt.Println("Votre parie n'a pas été sauvegarder")
								boolean = false
							default:
								fmt.Println("Veuillez mettre une réponse correcte svp")
							}
						}
					}
				case "2":

					cpt, lettremanque, Ascci, Mot, Pendu, Lettre, Stop = hangman.JeuAscii(cpt, lettremanque, Ascci, Mot, Pendu, Lettre)

					// Sauvegarde
					if Stop {
						boolean := true
						for boolean {
							fmt.Println("Voulez vous sauvegarder votre partie : oui / non ")
							choice := ""
							fmt.Scan(&choice)
							switch choice {
							case "oui":
								fmt.Println("Votre Partie a été sauvegarder")
								// sauvegarde du pendu avec l'Ascii-Art
								boolean = false
							case "non":
								fmt.Println("Votre parie n'a pas été sauvegarder")
								boolean = false
							default:
								fmt.Println("Veuillez mettre une réponse correcte svp")
							}
						}
					}
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
			file, err := os.Open("readme.txt")
			if err != nil {
				panic(err)
			}

			defer file.Close()

			reader := bufio.NewReader(file)

			for {
				line, _, err := reader.ReadLine()

				if err == io.EOF {
					break
				}

				fmt.Println(string(line))
			}
		case "4":
			return
		}
	}
}
