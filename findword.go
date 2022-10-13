package hangman

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
	var Mot string
	var words []string
	if len(os.Args[1:]) == 0 {
		fmt.Printf("Il manque le nom d'un fichier text en argument\n")
		return "Veuillez relancer le jeux"
	} else if len(os.Args[1:]) > 1 {
		fmt.Printf("Il y a trop d'argument\n")
		return "Veuillez relancer le jeux"
	} else {
		contents, err := os.Open(os.Args[1])
		if err != nil {
			fmt.Println("File reading error", err)
			return "Veuillez relancer le jeux"
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

		rand.Seed(time.Now().UnixNano())
		nbr := rand.Intn(len(words))
		Mot = words[nbr]
	}
	return Mot
}
