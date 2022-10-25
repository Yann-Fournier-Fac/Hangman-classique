package hangman

import (
	"encoding/json"
	"fmt"
)

type Save struct {
	Cptt            int
	Lettremanquante int
	MotATrouve      []string
	Asccii          [][]string
	Mots            string
	Pend            []string
	Lett            []string
	Affiche         []string
}

func Unmarshal(data []byte) Save {
	var emp Save
	json.Unmarshal(data, &emp)
	return emp
}

func Marshal(emp Save) []byte {
	empData, err := json.Marshal(emp)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return empData
}
