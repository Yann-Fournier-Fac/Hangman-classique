package hangman

func Pendu(pendu []byte) []string {
	cpt := 0
	tab := []string{}
	str := ""
	for i := 0; i < len(pendu); i++ {
		//fmt.Println(str)
		if cpt == 8 {
			tab = append(tab, str)
			str = " "
			cpt = 0
		} else if string(pendu[i]) == "\n" {
			str = str + "\n"
			cpt++
		} else if string(pendu[i]) != "\n" {
			str = str + string(pendu[i])
		}
	}
	tab = append(tab, str)
	return tab
}
