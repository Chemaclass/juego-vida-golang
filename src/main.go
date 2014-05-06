package main

import (
	"fmt"
	"time"
)

//Tamaño del tablero
const TAMANO_TABLERO = 15
const TIEMPO_ACTUALIZACION = time.Second

//Estados de una célula
const VIVA = "*"
const MUERTA = " "

type Celula string
type Tablero [][]Celula

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

/* Rellenar todo el tablero */
func (self *Juego) Fill() {
	//Inicializamos el tablero
	self.Tablero = make([][]Celula, TAMANO_TABLERO)
	for i, _ := range self.Tablero {
		self.Tablero[i] = make([]Celula, TAMANO_TABLERO)
	}
	//rellenamos con células muertas
	for i, v := range self.Tablero {
		for j, _ := range v {
			self.Tablero[i][j] = MUERTA
		}
	}
}


func main() {
	j := new(Juego)
	j.Fill()
	j.fillTest()
	for {
		fmt.Println(j)
		
		time.Sleep(TIEMPO_ACTUALIZACION)
	}
}

/* Rellenar todo el tablero */
func (j *Juego) fillTest() {
	j.Tablero[0][0] = VIVA
	j.Tablero[0][1] = VIVA
	j.Tablero[1][0] = VIVA
	j.Tablero[1][2] = VIVA
	j.Tablero[2][1] = VIVA
	j.Tablero[3][5] = VIVA
	j.Tablero[2][2] = VIVA
	j.Tablero[3][4] = VIVA
}
