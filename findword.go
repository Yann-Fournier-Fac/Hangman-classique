package hangman

import (
	"math/rand"
	"time"
)

// compresser le dossier en .zip (pas en .rar(sinon modification de ))

func Findword(s []byte) string {
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
}
