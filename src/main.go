package main

import (
	"juego"
	"time"
)

func main() {
	j := new(juego.Juego)
	j.Fill()
	j.FillInitTest()
	for {
		j.Print()
		j.Check()
		time.Sleep(juego.TIEMPO_ACTUALIZACION)
	}
}
