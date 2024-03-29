package juego

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
	NumVivas int
}

func (j Juego) String() string {
	str := "\n"
	for _, v := range j.Tablero {
		str += fmt.Sprintf("%v\n", v)
	}
	str += fmt.Sprintf("Número de células vivas: %d\n", j.NumVivas)
	return str
}

func (j Juego) Print() {
	fmt.Println(j)
}

// Rellenar todo el tablero
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
	self.NumVivas = 0
}

// Comprobar todas las casillas del tablero
func (self *Juego) Check() {
	numVivas := 0
	var flag bool
	for i, v := range self.Tablero {
		for j, celula := range v {

			//Si la célula está muerta
			if celula == MUERTA {
				//Celula.Check(false,i,j)
				flag = self.checkCelula(false, i, j)
			} else {
				//Si la célula está viva
				//Puede sobrevivir o morir
				flag = self.checkCelula(true, i, j)
			}
			if flag {
				self.Tablero[i][j] = VIVA
				numVivas++
			} else {
				self.Tablero[i][j] = MUERTA
			}
			//fmt.Printf("%d,%d -> %d\n", i, j, celula)
		}
	}
	self.NumVivas = numVivas
}

// Comprobar las células hermanas de una célula
func (self Juego) checkCelula(flag bool, i int, j int) bool {

	//Para controlar los 'runtime error: index out of range'
	defer func() {
		if r := recover(); r != nil {
		}
	}()

	c := 0 //count

	if self.Tablero[i-1][j-1] == VIVA {
		c++
	}
	if self.Tablero[i-1][j] == VIVA {
		c++
	}
	if self.Tablero[i-1][j+1] == VIVA {
		c++
	}

	if self.Tablero[i][j-1] == VIVA {
		c++
	}
	if self.Tablero[i][j+1] == VIVA {
		c++
	}

	if self.Tablero[i+1][j-1] == VIVA {
		c++
	}
	if self.Tablero[i+1][j] == VIVA {
		c++
	}
	if self.Tablero[i+1][j+1] == VIVA {
		c++
	}

	//Si está muerta y tiene 3 o más vivas
	// o si está viva y tiene 2 o 3 vivas sigue viva
	if (!flag && c >= 3) || (flag && (c == 2 || c == 3)) {
		return true
	}
	return false
}

// Rellenar tablero con algunas casillas vivas
func (self *Juego) FillInitTest() {
	self.Tablero[1][0] = VIVA
	self.Tablero[1][2] = VIVA
	self.Tablero[2][1] = VIVA
	self.Tablero[3][5] = VIVA
	self.Tablero[2][2] = VIVA
	self.Tablero[3][4] = VIVA
	self.Tablero[1][4] = VIVA
	self.NumVivas = 7
}

// Parar el Juego para que podamos ver los cambios
func (self *Juego) Sleep() {
	time.Sleep(TIEMPO_ACTUALIZACION)
}

//Ejecutar el Juego
func (self *Juego) Run() {
	for {
		self.Print() // Pintar
		self.Check() // Actualizar
		self.Sleep() // Dormir
	}
}
