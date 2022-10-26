package main

// faire la sauvegarde (la pofiner)

import (
	"bufio"
	"fmt"
	"hangman"
	"io"
	"io/ioutil"
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

			if len(os.Args) == 2 && os.Args[1] == "Save.txt" {
				// Lire le fichier txt
				data, err := ioutil.ReadFile("Save.txt") // lire le fichier
				if err != nil {
					fmt.Println(err)
				}

				unmarshsave := hangman.Detransformation(data)
				if len(unmarshsave.Asccii) == 0 {
					hangman.JeuBase(unmarshsave.Cptt, unmarshsave.Lettremanquante, unmarshsave.MotATrouve, unmarshsave.Mots, unmarshsave.Pend, unmarshsave.Lett, unmarshsave.Affiche, unmarshsave.Asccii)
				} else {
					hangman.JeuAscii(unmarshsave.Cptt, unmarshsave.Lettremanquante, unmarshsave.Asccii, unmarshsave.Mots, unmarshsave.Pend, unmarshsave.Lett, unmarshsave.Affiche, unmarshsave.MotATrouve)
				}
			} else {

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

					// Initialisation des variables
					cpt := -1
					var lettremanque int
					MotATrouver := []string{}
					Ascci := [][]string{}
					Mot := ""
					var Pendu []string
					Lettre := []string{} // pour les lettres deja dites
					var Stop bool
					var Affichage []string
					// Determiner le mot a deviner le mot à deviner
					Mot = hangman.Findword()
					//fmt.Println(Mot)

					switch choix {
					case "1":

						// on teste que le mot soit bon
						if Mot == "Veuillez relancer le jeux" {
							fmt.Println("Veuillez relancer le jeux")
							return
						}

						Pendu = hangman.Hangmanpose()

						// creation des "_" en Ascii
						for i := 0; i < len(Mot); i++ {
							MotATrouver = append(MotATrouver, "_ ")
						}

						// affichage des n lettre Ascii
						n := (len(Mot) / 2) - 1
						lettremanque = len(Mot) - n
						MotATrouver, Affichage = hangman.NLetterBase(Mot, MotATrouver, Affichage)

						cpt, lettremanque, MotATrouver, Mot, Pendu, Lettre, Stop, Affichage, Ascci = hangman.JeuBase(cpt, lettremanque, MotATrouver, Mot, Pendu, Lettre, Affichage, Ascci)

						// Sauvegarde
						if Stop {
							boolean := true
							for boolean {
								fmt.Println("Voulez vous sauvegarder votre partie : oui / non ")
								choice := ""
								fmt.Scan(&choice)
								switch choice {
								case "oui":

									hangman.Sauvegarder(cpt, lettremanque, Ascci, Mot, Pendu, Lettre, Affichage, MotATrouver)

									fmt.Println("Votre Partie a été sauvegarder")

									boolean = false
								case "non":
									fmt.Println("Votre partie n'a pas été sauvegarder")
									boolean = false
								default:
									fmt.Println("Veuillez mettre une réponse correcte svp")
								}
							}
						}
					case "2":

						// creation des "_" en Ascii
						for i := 0; i < len(Mot); i++ {
							Ascci = append(Ascci, []string{"             ", "             ", "             ", "             ", " ___________ ", "|___________|", "              "}) // 7 éléments
						}
						// on teste que le mot soit bon
						if Mot == "Veuillez relancer le jeux" {
							fmt.Println("Veuillez relancer le jeux")
							return
						}

						Pendu = hangman.Hangmanpose()

						// affichage des n lettre Ascii
						n := (len(Mot) / 2) - 1
						lettremanque = len(Mot) - n
						Ascci, Affichage = hangman.NLetterAscii(Mot, Ascci, Affichage)

						cpt, lettremanque, Ascci, Mot, Pendu, Lettre, Stop, Affichage, MotATrouver = hangman.JeuAscii(cpt, lettremanque, Ascci, Mot, Pendu, Lettre, Affichage, MotATrouver)

						// Sauvegarde
						if Stop {
							boolean := true
							for boolean {
								fmt.Println("Voulez vous sauvegarder votre partie : oui / non ")
								choice := ""
								fmt.Scan(&choice)
								switch choice {
								case "oui":
									hangman.Sauvegarder(cpt, lettremanque, Ascci, Mot, Pendu, Lettre, Affichage, MotATrouver)

									fmt.Println("Votre Partie a été sauvegarder")

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
