package hangman

// faire un dictionnaire

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

// compresser le dossier en .zip (pas en .rar(sinon modification de ))

/*func Findword(s []byte) string {
	mot := ""
	tab := bytetoword(s)
	long := len(tab)
	rand.Seed(time.Now().UnixNano())
	nbr := rand.Intn(long)
	mot = tab[nbr]
	return mot
}

func bytetoword(tab []byte) []string {
	tabl := []string{}
	str := ""
	if len(tab) == 0 {
		return tabl
	}
	for i := 0; i < len(tab); i++ {
		if tab[i] != 10 {
			str += string(rune(tab[i]))
		} else if tab[i] == 10 {
			tabl = append(tabl, str)
			str = ""
		}
	}
	return tabl
}*/

func Findword() string {

	// Initialisation Variables
	var Mot string
	var words []string
	Dictionnaire := [][]string{}

	// Si aucun fichier txt est placé en paramettre
	if len(os.Args[1:]) != 1 {

		Dictionnaire = append(Dictionnaire, readfile("words.txt"))
		Dictionnaire = append(Dictionnaire, readfile("words2.txt"))
		Dictionnaire = append(Dictionnaire, readfile("words3.txt"))

		for _, txt := range Dictionnaire {
			for _, i := range txt {
				words = append(words, i)
			}
		}
		
		if len(words) != 0 {
			rand.Seed(time.Now().UnixNano())
			nbr := rand.Intn(len(words))
			Mot = words[nbr]
			return Mot
		} else {
			return "Veuillez relancer le jeux"
		}
		// Sinon ...
	} else {
		words = readfile(os.Args[1])
		if len(words) != 0 {
			rand.Seed(time.Now().UnixNano())
			nbr := rand.Intn(len(words))
			Mot = words[nbr]
			return Mot
		} else {
			return "Veuillez relancer le jeux"
		}
	}
}

func readfile(name string) []string {
	words := []string{}
	contents, err := os.Open(name)
	if err != nil {
		fmt.Println("File reading error", err)
		fmt.Println("Veuillez relencer le jeux")
		return words
	}
	defer contents.Close()

	scanner := bufio.NewScanner(contents)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return words
}
