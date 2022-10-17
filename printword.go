package hangman

import (
	"fmt"
	"math/rand"
	"time"
)

func Prtword(tab [][]string) {
	for i := 0; i < len(tab[0]); i++ {
		for j := 0; j < len(tab); j++ {
			fmt.Printf(tab[j][i])
		}
		fmt.Printf("\n")
	}
}

func NLetter(Mot string, ascii [][]string, lettres []string) ([][]string, []string) {
	// afficher aléatoirement n lettres
	n := len(Mot)/2 - 1
	tab := []int{}
	boolean := false
	if n == 0 {
		return ascii, lettres
	}
	for !boolean {
		for i := 0; i < n; i++ {
			rand.Seed(time.Now().UnixNano())
			radomInt2 := rand.Intn(len(Mot))
			tab = append(tab, radomInt2)
		}
		boolean = Alldiff(tab)
	}
	for i := 0; i < len(tab); i++ {
		ascii[tab[i]] = Lettertoascii(string(Mot[tab[i]]))
		lettres = append(lettres, string(Mot[tab[i]]))
		//tableau[tab[i]] = string(Mot[tab[i]])
	}
	//fmt.Print(tab)
	return ascii, lettres
}

func Alldiff(tab []int) bool {
	for i := 0; i < len(tab)-1; i++ {
		for j := i + 1; j < len(tab); j++ {
			if tab[i] == tab[j] {
				return false
			}
		}
	}
	return true
}
