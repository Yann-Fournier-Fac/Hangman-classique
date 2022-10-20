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
			// Initialisation des variables
			cpt := -1
			var lettremanque int
			Ascci := [][]string{}
			Mot := ""
			var Pendu []string
			Lettre := []string{} // pour les lettres deja dites

			// Determiner le mot a deviner le mot à deviner
			Mot = hangman.Findword()
			//fmt.Println(Mot)

			// on teste que le mot soit bon
			if Mot == "Veuillez relancer le jeux" {
				fmt.Println("Veuillez relancer le jeux")
				return
			}

			// creation des "_" en Ascii
			for i := 0; i < len(Mot); i++ {
				Ascci = append(Ascci, []string{"             ", "             ", "             ", "             ", " ___________ ", "|___________|", "              "}) // 7 éléments
			}

			// affichage des n lettre Ascii
			n := (len(Mot) / 2) - 1
			lettremanque = len(Mot) - n
			Ascci, Lettre = hangman.NLetter(Mot, Ascci, Lettre)

			Pendu = hangman.Hangmanpose()

			// Debut du jeu
			fmt.Println("Une nouvelle partie a été lancée")
			fmt.Printf("\n")
			fmt.Printf("Good luck")
			fmt.Printf("\n")

			for (cpt < 9) && (lettremanque != 0) {

				//fmt.Print(Lettre)
				//fmt.Printf("\n")

				// Affichage du nombre d'essais restant et du mot en Ascci-art
				fmt.Printf("You have %v possible errors left", 10-(cpt+1))
				fmt.Printf("\n")
				hangman.Prtword(Ascci)
				fmt.Printf("\n")

				// le joueur entre quelque chose
				fmt.Printf("Ecrivez une lettre ou un mot s'il vous plaît : ")
				var lettre string
				fmt.Scan(&lettre)

				// comparer le mot rentrer par le joueur
				if len(lettre) >= 2 {
					if len(lettre) == len(Mot) {
						var cpt2 int // compteur de lettre correspondentes
						for i := 0; i < len(Mot); i++ {
							if lettre[i] == Mot[i] {
								cpt2++
							}
						}
						if cpt2 == len(Mot) { // verification du nbr de lettre correspondentes
							lettremanque = 0
						} else {
							cpt += 2
							//hangman.Hang(cpt, Pendu)
							fmt.Println("Ce n'était pas le bon mot")
						}
					} else if len(lettre) != len(Mot) {
						cpt += 2
						//hangman.Hang(cpt, Pendu)
						fmt.Println("Ce n'était pas le bon mot")
					}

				} else { // Sinon comparer la lettre rentrer par le joueur

					var cpt4 int = 0 // compteur de lettre correspondentes

					lettre = strings.ToLower(lettre) // mettre la lettre en minuscule (si besoin)

					// On cherche si la lettre a déjà été rentrée
					for i := 0; i < len(Lettre); i++ {
						if lettre == Lettre[i] {
							cpt4++
							break
						}
					}
					if cpt4 == 0 { // Si elle est nouvelle

						Lettre = append(Lettre, lettre) // ajout à lettre déjà entrer à une liste

						// on transforme la lettre en Ascii
						let := 0
						for i := 0; i < len(Mot); i++ {
							if lettre == string(Mot[i]) {
								Ascci[i] = hangman.Lettertoascii(lettre)
								lettremanque--
								let++
							}
						}
						cpt4 = 0
						if let == 0 {
							cpt++
						}

					} else {
						fmt.Println("Cette lettre à déjà été rentrée")
					}
				}

				// Afficher le pendu
				if cpt == -1 {
					fmt.Printf("José ce porte bien")
					fmt.Printf("\n")
				} else {
					if cpt != 10 {
						hangman.Hang(cpt, Pendu)
					}
					fmt.Printf("\n")
					fmt.Println("José est en danger")
					fmt.Printf("\n")
				}

			}

			// fin du jeu
			if cpt >= 9 { // les 10 essays ont été utilisé
				fmt.Println("Dommage, vous avez tué José. lance une nouvelle partie pour réessayer")
				fmt.Printf("Le mot à trouver était : %v", Mot)
				fmt.Printf("\n")
				fmt.Printf("\n")

			} else if lettremanque == 0 { // Toutes les lettres ont été trouvée

				// Affichage du Mot en Ascci
				for i := 0; i < len(Mot); i++ {
					Ascci[i] = hangman.Lettertoascii(string(Mot[i]))
				}
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
