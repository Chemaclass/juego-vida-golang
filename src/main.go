package main

import (
	"juego"
)

func main() {
	j := new(juego.Juego)
	j.Fill()
	j.FillInitTest()
	for {
		j.Print() // Pintar
		j.Check() // Actualizar
		j.Sleep() // Dormir
	}
}
