package main

// faire un tableau des lettres déjà entrer
// donner la possibilité au joueur de rentrer un mot
// faire une bibliothèque de mots

import (
	"fmt"
	"hangman"
	"strings"
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

			// Initialisation des variables
			cpt := -1
			var lettremanque int
			Ascci := [][]string{}
			Mot := ""
			var Pendu []string
			//Lettre := []string{} // pour les lettres deja dites
			//Dictionnaire := []string{}

			// Determiner le mot a deviner le mot à deviner
			Mot = hangman.Findword()
			fmt.Println(Mot)

			if Mot == "Veuillez relancer le jeux" {
				fmt.Println("Veuillez relancer le jeux")
				return
			}

			for i := 0; i < len(Mot); i++ {
				Ascci = append(Ascci, []string{"             ", "             ", "             ", "             ", " ___________ ", "|___________|", "              "}) // 7 éléments
			}

			// affichage des n lettre Ascii
			n := (len(Mot) / 2) - 1
			lettremanque = len(Mot) - n
			Ascci = hangman.NLetter(Mot, Ascci)

			Pendu = hangman.Hangmanpose()

			// Debut du jeu
			fmt.Printf("Good luck")
			fmt.Printf("\n")
			for (cpt != 9) && (lettremanque != 0) {
				fmt.Printf("You have %v possible errors left", 10-(cpt+1))
				fmt.Printf("\n")
				hangman.Prtword(Ascci)
				fmt.Printf("\n")

				fmt.Printf("Ecrivez une lettre ou un mot s'il vous plaît : ")
				var lettre string
				fmt.Scan(&lettre)
				lettre = strings.ToLower(lettre)
				/*if len(Lettre) != 0 {
					for i := 0; i < len(Lettre); i++ {
						if
					}
				}*/
				let := 0
				for i := 0; i < len(Mot); i++ {
					if lettre == string(Mot[i]) {
						Ascci[i] = hangman.Lettertoascii(lettre)
						lettremanque--
						let++
					}
				}
				if let == 0 {
					cpt++
				}
				if cpt == -1 {
					fmt.Printf("José ce porte bien")
					fmt.Printf("\n")
				} else {
					if cpt != 10 {
						hangman.Hang(cpt, Pendu)
					}
					fmt.Printf("\n")
					fmt.Println("José se chie dessus")
					fmt.Printf("\n")
				}

			}
			if cpt == 9 {
				fmt.Println("Dommage, vous avez tué José. lance une nouvelle partie pour réessayer")
				fmt.Printf("Le mot à trouver était : %v", Mot)
				fmt.Printf("\n")
			} else if lettremanque == 0 {
				hangman.Prtword(Ascci)
				fmt.Println("Bravo, Tu as sauvé José")
				fmt.Printf("\n")
				fmt.Printf("Tu as trouvé le bon mot qui était %v", Mot)
				fmt.Printf("\n")
				fmt.Printf("\n")
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
