package main

import (
	"fmt"
)

const TAMANO_TABLERO = 10

type Tablero [][]int

type Juego struct {
	Tablero
}

func (j Juego) String() string {
	str := "\n"
	for _, v := range j.Tablero {
		str += fmt.Sprintf("%v\n", v)
	}
	return str
}

func (j *Juego) Fill(cant int) {
	j.Tablero = make([][]int, cant)
	for i, _ := range j.Tablero {
		j.Tablero[i] = make([]int, cant)
	}
}

func main() {
	fmt.Println("> main()")
	j := new(Juego)
	j.Fill(TAMANO_TABLERO)
	fmt.Println("Fill:", j)
	fmt.Println("< main()")
}
