package hangman

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/*func Pendu(pendu []byte) []string {
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
}*/

func Hangmanpose() []string {
	var pose []string
	file, err := os.Open("hangman.txt")
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
		pose = append(pose, string(line))
		//fmt.Printf("%s \n", line)
	}
	return pose
}

func Hang(hp int, pose []string) {
	a := 0 + hp*8
	b := 7 + hp*8
	for i := a; i <= b; i++ {
		fmt.Print(pose[i])
		fmt.Printf("\n")
	}

}

/*func Hangmanpose(hp int) {
	f, _ := os.Open("hangman.txt")
	scanner := bufio.NewScanner(f)
	var line int
	a := 0 + hp*8
	b := 7 + hp*8
	for scanner.Scan() {
		if line >= a && line <= b {
			fmt.Println(scanner.Text())
		}
		line++
	}
}*/
