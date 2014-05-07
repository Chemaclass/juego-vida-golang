package main

import (
	"juego"
)

func main() {
	j := new(juego.Juego)
	j.Fill()
	j.FillInitTest()
	j.Run()
}
