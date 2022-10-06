package main

import (
	"fmt"
	"hangman"
	"math/rand"
	"os"
	"time"
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
		switch str {
		case "1":
			fmt.Println("Une nouvelle partie a été lancée")
			fmt.Printf("\n")

			Mot := ""
			// Determiner le mot a deviner le mot à deviner
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
				Mot = hangman.Findword(contents)
				//fmt.Println(mot)

				// Initialisation des varibles
				var cpt int
				var lettremanque int
				Ascci := [][]string{}
				for i := 0; i < len(Mot); i++ {
					Ascci = append(Ascci, []string{})
				}

				// afficher aléatoirement n lettres
				n := len(Mot)/2 - 1
				tab := []int{}
				boolean := false
				for !boolean {
					for i := 0; i < n; i++ {
						rand.Seed(time.Now().UnixNano())
						radomInt2 := rand.Intn(len(Mot))
						tab = append(tab, radomInt2)
					}
					boolean = hangman.Alldiff(tab)
				}
				for i := 0; i < len(tab); i++ {
					//Ascci[tab[i]] = lettertoascci(string(Mot[tab[i]]))
				}

				// Debut du jeu
				for (cpt != 10) || (lettremanque != 0) {
					fmt.Printf("Good Luck, you have %v attempts.", 10-cpt)
					fmt.Printf("\n")
					hangman.Prtword(Ascci)
					//hangman.Posehang(cpt)

					fmt.Print("Merci d'écrire des lettres en minuscules s'il vous plaît : ")
					var lettre string
					fmt.Scan(&lettre)
					for i := 0; i < len(Mot); i++ {
						if lettre == string(Mot[i]) {
							//Ascci[i] = hangman.Lettertoascci(lettre)
							lettremanque--
						} else {
							cpt++
						}
					}
				}
				if cpt == 10 {
					//hangman.Posehang(10)
					fmt.Println("Dommage, lance une nouvelle partie pour réessayer")
					break
				} else if lettremanque == 0 {
					hangman.Prtword(Ascci)
					fmt.Printf("Bravo, tu as trouvé le bon mot qui était %v", Mot)
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
