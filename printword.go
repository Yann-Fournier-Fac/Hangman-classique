package hangman

import (
	"fmt"
	"math/rand"
	"time"
)

// Affichage des lettres en Ascii
// Elles sont afficher ligne par ligne pour pouvoire les afficher cote à cote
func Prtword(tab [][]string) {
	for i := 0; i < len(tab[0]); i++ {
		for j := 0; j < len(tab); j++ {
			fmt.Printf(tab[j][i])
		}
		fmt.Printf("\n")
	}
}

// Affichages de N lettres aléatoires pour deviner le Mot au début
func NLetterAscii(Mot string, ascii [][]string, Affich []string) ([][]string, []string) {

	n := len(Mot)/2 - 1 // Le nombre de lettre à afficher
	tab := []int{}
	//cpt := []int{} //bug ajout des lettres
	boolean := false

	// On affiche aucune lettre car le mot est trop petit
	if n == 0 {
		return ascii, Affich
	}

	// On recupere n numéros < len(Mot)
	// Ce sont les indices des lettres à afficher
	for !boolean {
		for i := 0; i < n; i++ {
			rand.Seed(time.Now().UnixNano())
			radomInt2 := rand.Intn(len(Mot))
			tab = append(tab, radomInt2)
		}
		boolean = Alldiff(tab) // tous les indices doivent évidemment etre different
	}

	// Puis on met les lettres en Ascii-Art
	for i := 0; i < len(tab); i++ {
		ascii[tab[i]] = Lettertoascii(string(Mot[tab[i]]))
		Affich = append(Affich, string(Mot[tab[i]]))
	}

	//fmt.Print(tab)
	return ascii, Affich
}

func NLetterBase(Mot string, MotATrouver []string, Affich []string) ([]string, []string){

	n := len(Mot)/2 - 1 // Le nombre de lettre à afficher
	tab := []int{}
	//cpt := []int{} //bug ajout des lettres
	boolean := false

	// On affiche aucune lettre car le mot est trop petit
	if n == 0 {
		return MotATrouver, Affich
	}

	// On recupere n numéros < len(Mot)
	// Ce sont les indices des lettres à afficher
	for !boolean {
		for i := 0; i < n; i++ {
			rand.Seed(time.Now().UnixNano())
			radomInt2 := rand.Intn(len(Mot))
			tab = append(tab, radomInt2)
		}
		boolean = Alldiff(tab) // tous les indices doivent évidemment etre different
	}

	// Puis on met les lettres en Ascii-Art
	for i := 0; i < len(tab); i++ {
		MotATrouver[tab[i]] = string(Mot[tab[i]])
		Affich = append(Affich, string(Mot[tab[i]]))
	}

	//fmt.Print(tab)
	return MotATrouver, Affich
}

// retourne true si tous les élément du tableau tab sont différents
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
