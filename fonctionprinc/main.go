package main

// Mettre toutes les lettres de la même longueur 

import (
	"fmt"
	"hangman"
	"io/ioutil"
	"os"
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
			//Motdev := []string{}

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
				fmt.Println(Mot)

				for i := 0; i < len(Mot); i++ {
					Ascci = append(Ascci, []string{"         ", "         ", "         ", "         ", " _______ ", "|_______|"}) // 8 de longueur
					//Motdev = append(Motdev, "_")
				}

				// affichage des n lettre Ascii
				n := (len(Mot) / 2) - 1
				lettremanque = len(Mot) - n
				//fmt.Print(lettremanque)
				Ascci = hangman.NLetter(Mot, Ascci)
				//Motdev = hangman.NLetter(Mot, Motdev)

				// Creation des niveux d'erreurs
				content, err := ioutil.ReadFile("hangman.txt")
				if err != nil {
					return
				}
				//fmt.Print(content)
				pendu := hangman.Pendu(content)

				// Debut du jeu
				fmt.Printf("Good luck")
				fmt.Printf("\n")
				for (cpt != 9) && (lettremanque != 0) {
					fmt.Printf("You have %v possible errors left", 10-(cpt+1))
					fmt.Printf("\n")
					fmt.Print(cpt)
					fmt.Printf("\n")
					hangman.Prtword(Ascci)
					/*for i := 0; i < len(Mot); i++ {
						fmt.Printf(Motdev[i])
					}*/
					//fmt.Print(Motdev)
					fmt.Printf("\n")

					fmt.Printf("Merci d'écrire des lettres en minuscules s'il vous plaît : ")
					var lettre string
					fmt.Scan(&lettre)
					let := 0
					for i := 0; i < len(Mot); i++ {
						if lettre == string(Mot[i]) {
							Ascci[i] = hangman.Lettertoascii(lettre)
							//Motdev[i] = string(Mot[i])
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
							fmt.Print(pendu[cpt])
						}
						fmt.Printf("\n")
						fmt.Println("José ce chie dessus")
						fmt.Printf("\n")
					}
				}
			}
			if cpt == 9 {
				//hangman.Posehang(10)
				fmt.Println("Dommage, vous avez tué José. lance une nouvelle partie pour réessayer")
				fmt.Printf("Le mot à trouver était : %v", Mot)
				fmt.Printf("\n")
			} else if lettremanque == 0 {
				hangman.Prtword(Ascci)
				/*for i := 0; i < len(Mot); i++ {
					fmt.Printf(Motdev[i])
				}*/
				fmt.Println("Bravo, Tu as sauvé José")
				fmt.Printf("\n")
				fmt.Printf("Tu as trouvé le bon mot qui était %v", Mot)
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
