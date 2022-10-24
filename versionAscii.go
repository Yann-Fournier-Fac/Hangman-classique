package hangman

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func JeuAscii(cpt int, lettremanque int, Ascci [][]string, Mot string, Pendu []string, Lettre []string) (int, int, [][]string, string, []string, []string, bool) {

	var Stop bool = false
	var stop string = "stop"
	// Determiner le mot a deviner le mot à deviner
	Mot = Findword()
	//fmt.Println(Mot)

	// creation des "_" en Ascii
	for i := 0; i < len(Mot); i++ {
		Ascci = append(Ascci, []string{"             ", "             ", "             ", "             ", " ___________ ", "|___________|", "              "}) // 7 éléments
	}

	// affichage des n lettre Ascii
	n := (len(Mot) / 2) - 1
	lettremanque = len(Mot) - n
	Ascci = NLetterAscii(Mot, Ascci)

	// Debut du jeu
	fmt.Println("Une nouvelle partie a été lancée")
	fmt.Printf("\n")
	fmt.Printf("Good luck")
	fmt.Printf("\n")

	for (cpt < 9) && (lettremanque != 0) {

		//fmt.Print(Lettre)
		//fmt.Printf("\n")

		// Affichage du nombre d'essais restant et du mot en Ascci-art
		if cpt >= 7 {
			fmt.Printf(Red+"You have %v possible errors left\n", 10-(cpt+1))
			fmt.Println(Reset)
		} else if cpt >= 4 && cpt < 7 {
			fmt.Printf(Yellow+"You have %v possible errors left\n", 10-(cpt+1))
			fmt.Println(Reset)
		} else {
			fmt.Printf(Green+"You have %v possible errors left\n", 10-(cpt+1))
			fmt.Println(Reset)
		}

		fmt.Printf("\n")
		Prtword(Ascci)
		fmt.Printf("\n")

		fmt.Println("Voici les lettres déjà entrée :")
		fmt.Print(Lettre)
		fmt.Printf("\n")
		fmt.Printf("\n")

		// le joueur entre quelque chose
		fmt.Printf("Ecrivez une lettre ou un mot s'il vous plaît : ")
		var lettre string
		fmt.Scan(&lettre)

		// Clear le terminal
		c := exec.Command("clear")
		c.Stdout = os.Stdout
		c.Run()

		// comparer le mot rentrer par le joueur
		if len(lettre) >= 2 {
			if len(lettre) == 4 {
				for i := 0; i < len(lettre); i++ {
					if strings.ToLower(string(lettre[i])) != string(stop[i]) {
						break
					} else {
						Stop = true
						return cpt, lettremanque, Ascci, Mot, Pendu, Lettre, Stop
					}
				}
			} else if len(lettre) == len(Mot) {
				var cpt2 int // compteur de lettre correspondentes
				for i := 0; i < len(Mot); i++ {
					let := strings.ToLower(string(lettre[i]))
					if let == string(Mot[i]) {
						cpt2++
					}
				}
				if cpt2 == len(Mot) { // verification du nbr de lettre correspondentes
					lettremanque = 0
				} else {
					cpt += 2
					//hangman.Hang(cpt, Pendu)
					fmt.Println(Purple + "Ce n'était pas le bon mot" + Reset)
				}
			} else if len(lettre) != len(Mot) {
				cpt += 2
				//hangman.Hang(cpt, Pendu)
				fmt.Println(Purple + "Ce n'était pas le bon mot" + Reset)
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
						Ascci[i] = Lettertoascii(lettre)
						lettremanque--
						let++
					}
				}
				cpt4 = 0
				if let == 0 {
					cpt++
				}

			} else {
				fmt.Println(Purple + "Cette lettre à déjà été rentrée" + Reset)
			}
		}

		// Afficher le pendu
		if cpt == -1 {
			fmt.Printf(Green + "José se porte bien" + Reset)
			fmt.Printf("\n")
		} else {
			if cpt != 10 {
				Hang(cpt, Pendu)
			}
			fmt.Printf("\n")
			if cpt < 6 {
				fmt.Println(Yellow + "José est en danger" + Reset)
				fmt.Printf("\n")
			} else {
				fmt.Println(Red + "José est en danger" + Reset)
				fmt.Printf("\n")
			}
		}
	}

	// fin du jeu
	if cpt >= 9 { // les 10 essays ont été utilisé
		fmt.Println(Red + "Dommage, vous avez tué José. lance une nouvelle partie pour réessayer" + Reset)
		fmt.Printf("Le mot à trouver était : %v \n", Mot)
		for i := 0; i < len(Mot); i++ {
			Ascci[i] = Lettertoascii(string(Mot[i]))
		}
		Prtword(Ascci)
		fmt.Printf("\n")
		fmt.Printf("\n")
		return cpt, lettremanque, Ascci, Mot, Pendu, Lettre, Stop

	} else if lettremanque == 0 { // Toutes les lettres ont été trouvée

		// Affichage du Mot en Ascci
		for i := 0; i < len(Mot); i++ {
			Ascci[i] = Lettertoascii(string(Mot[i]))
		}
		Prtword(Ascci)

		fmt.Println(Green + "Bravo, Tu as sauvé José" + Reset)
		fmt.Printf("\n")
		fmt.Printf("Tu as trouvé le bon mot qui était %v", Mot)
		fmt.Printf("\n")
		fmt.Printf("\n")
		return cpt, lettremanque, Ascci, Mot, Pendu, Lettre, Stop
	}
	return cpt, lettremanque, Ascci, Mot, Pendu, Lettre, Stop
}
