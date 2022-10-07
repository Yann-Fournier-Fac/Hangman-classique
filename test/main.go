package main

func main() {

	/*Motdev := []string{}

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


	}*/

	//hangman := ReadFile(hangman.txt)
	//tab := [][]string{[]string{" _ ", "| |", "| |", "| |", "|_|", "(_)"}, []string{"   _  _    ", " _| || |_  ", "|_  __  _| ", " _| || |_  ", "|_  __  _| ", "  |_||_|   "}}
	//Ascci := [][]string{}
	/*for i := 97; i <= 112; i++ {
	      Ascci = append(Ascci, hangman.Lettertoascii(string(rune(i))))
	  }
	*/
	//fmt.Print(Ascci)
	/*var e int
	for w := "send nudes"; e < len(w); {
		Ascci = append(Ascci, hangman.Lettertoascii(string(rune(w[e]))))
		e++
	}
	for i := 0; i < len(Ascci[0]); i++ {
		for j := 0; j < len(Ascci); j++ {
			fmt.Printf(Ascci[j][i])
		}
		fmt.Printf("\n")
	}*/

	/*fmt.Println("Merci d'entrer un caractère à convertir :")
	  letter := ""
	  fmt.Scan(&letter)
	  tab := hangman.Lettertoascii(letter)
	  fmt.Print(tab)*/

}
