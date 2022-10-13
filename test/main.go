package main

import (
	"bufio"
	"fmt"

	//"io"
	"os"
	//"devt.de/krotik/common/termutil/getch"
)

func hangman(hp int) {
	f, _ := os.Open("hangman.txt")
	scanner := bufio.NewScanner(f)
	var line int
	a := 71 - hp*8
	b := 78 - hp*8
	for scanner.Scan() {
		if line >= a && line <= b {
			fmt.Println(scanner.Text())
		}
		line++
	}
}

/*func Clear() {
	var KeyPress *getch.KeyEvent

	if err := getch.Start(); err != nil {
		fmt.Println(err)
		return
	}
	defer getch.Stop()

	KeyPress, _ = getch.Getch()
	fmt.Print(KeyPress)
}*/

func main() {

	hangman(3)

	/*file, err := os.Open("hangman.txt")
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
		fmt.Println(line)
		fmt.Printf("%s \n", line)
	}*/

	// Creation des niveux d'erreurs
	/*content, err := os.Open("hangman.txt")
	if err != nil {
		return
	}
	defer content.Close()

	scanner := bufio.NewScanner(content)
	//scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}*/
	//pendu := hangman.Pendu(content)

	/*f, err := os.Open("words.txt")

	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}*/

	/*nbr := []int{}
	for i := 0; i < len(nbr); i++ {
		fmt.Print(nbr[i])
	}*/

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
	for w := "ab"; e < len(w); {
		Ascci = append(Ascci, hangman.Lettertoascii(string(rune(w[e]))))
		e++
	}
	for i := 0; i < len(Ascci[0]); i++ {
		for j := 0; j < len(Ascci); j++ {
			fmt.Printf(Ascci[j][i])
			//fmt.Printf("@")
		}
		fmt.Printf("\n")
	}*/

	/*fmt.Println("Merci d'entrer un caractère à convertir :")
	  letter := ""
	  fmt.Scan(&letter)
	  tab := hangman.Lettertoascii(letter)
	  fmt.Print(tab)*/

}
